package main

import (
	"context"
	"errors"
	"fmt"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	nwclient "github.com/inhuman/noise_wrapper/pkg/gen/client"
	"time"

	"github.com/go-openapi/loads"
	"github.com/golang-migrate/migrate/v4"

	// drivers needed
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	libpgx "github.com/inhuman/emo_recognizer_common/pgx"
	"github.com/inhuman/emo_recognizer_common/storage/s3"
	"github.com/inhuman/emo_recognizer_controller/internal/config"
	"github.com/inhuman/emo_recognizer_controller/internal/handlers"
	"github.com/inhuman/emo_recognizer_controller/internal/jobprocessor"
	"github.com/inhuman/emo_recognizer_controller/internal/repository/db"
	"github.com/inhuman/emo_recognizer_controller/pkg/gen/restapi"
	"github.com/inhuman/emo_recognizer_controller/pkg/gen/restapi/operations"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jessevdk/go-flags"
	"go.uber.org/zap"
)

func main() {

	time.Sleep(5 * time.Second)

	ctx := context.Background()

	logger := zap.NewExample()

	appConfig, err := config.InitConfig(ctx, logger)
	panicErr(err, logger)

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	panicErr(err, logger)

	pgxPool := preparePg(ctx, appConfig.Db, logger)

	if appConfig.ApplyDbMigrations {
		logger.Info("applying db migrations",
			zap.String("migrations path", appConfig.MigrationsPath),
		)

		databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
			appConfig.Db.User,
			appConfig.Db.Password,
			appConfig.Db.Host,
			appConfig.Db.Port,
			appConfig.Db.DbName,
			appConfig.Db.SslMode,
		)

		migrator, err := migrate.New(appConfig.MigrationsPath, databaseURL)
		panicErr(err, logger)

		err = migrator.Up()
		if !errors.Is(err, migrate.ErrNoChange) {
			panicErr(err, logger)
		}

		logger.Info("migrations applied")
	}

	dbRepo := db.NewRepository(pgxPool, logger)

	fileStorage, err := s3.NewS3(
		s3.WithEndpoint(appConfig.S3.Endpoint),
		s3.WithBucketName(appConfig.S3.Bucket),
		s3.WithAccessKey(appConfig.S3.AccessKey),
		s3.WithSecretKey(appConfig.S3.SecretKey),
		s3.WithSecure(appConfig.S3.Secure),
		s3.WithCreateBucket(true),
	)
	panicErr(err, logger)

	transport := httptransport.New(appConfig.Services.NoiseWrapper.Address, "", []string{"http"})
	noiseWrapperClient := nwclient.New(transport, strfmt.Default)

	defaultStrategy := jobprocessor.NewDefaultStrategy(jobprocessor.DefaultStrategyOps{
		Repo:               dbRepo,
		NoiseWrapperClient: noiseWrapperClient,
		StorageClient:      fileStorage,
		Logger:             logger,
	})

	strategyChooser := jobprocessor.NewStrategyChooser()
	strategyChooser.AddStrategy(jobprocessor.StrategyDefault, defaultStrategy)

	jobsProcessor := jobprocessor.NewJobProcessor(jobprocessor.Opts{
		Repo:            dbRepo,
		Logger:          logger,
		StrategyChooser: strategyChooser,
		FileStorage:     fileStorage,
	})

	api := operations.NewEmotionsRecognizerAPI(swaggerSpec)

	handlers.SetupAPI(api, &handlers.SetupOpts{
		Logger:        logger,
		JobsProcessor: jobsProcessor,
	})

	server := restapi.NewServer(api)
	server.Port = appConfig.Port

	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Controller"
	parser.LongDescription = "Controller for Emo Recognizer"

	server.ConfigureFlags()

	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			panic(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1

		flagErr := &flags.Error{}

		if errors.As(err, &flagErr) {
			if flagErr.Type == flags.ErrHelp {
				code = 0
			}
		}

		panic(fmt.Errorf("exit code: %d", code))
	}

	server.ConfigureAPI()

	go jobsProcessor.Run(ctx, appConfig.JobProcessor)

	err = server.Serve()
	if err != nil {
		panic(fmt.Errorf("error serve server: %w", err))
	}
}

func preparePg(ctx context.Context, conf config.Database, logger *zap.Logger) *pgxpool.Pool {
	pgxConf := libpgx.PgArgs{
		Host:     conf.Host,
		Port:     conf.Port,
		DB:       conf.DbName,
		User:     conf.User,
		Password: conf.Password,
		SslMode:  conf.SslMode,
	}

	pgDsn, err := pgxConf.String()
	panicErr(err, logger)

	pgConf, err := pgxpool.ParseConfig(pgDsn)
	panicErr(err, logger)

	pgConf.LazyConnect = true

	pgxPool, err := pgxpool.ConnectConfig(ctx, pgConf)
	panicErr(err, logger)

	return pgxPool
}

func panicErr(err error, logger *zap.Logger) {
	if err != nil {
		logger.Panic("can not start app", zap.Error(err))
	}
}
