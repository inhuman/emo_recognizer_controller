package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-openapi/loads"
	"github.com/inhuman/emo_recognizer_controller/internal/config"
	"github.com/inhuman/emo_recognizer_controller/internal/controller"
	"github.com/inhuman/emo_recognizer_controller/internal/handlers"
	"github.com/inhuman/emo_recognizer_controller/pkg/gen/restapi"
	"github.com/inhuman/emo_recognizer_controller/pkg/gen/restapi/operations"
	"github.com/jessevdk/go-flags"
	"go.uber.org/zap"
)

func main() {

	ctx := context.Background()

	logger := zap.NewExample()

	appConfig, err := config.InitConfig(ctx, logger)
	panicErr(err, logger)

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	panicErr(err, logger)

	jobsProcesor := &controller.JobProcessor{}

	api := operations.NewNoiseWrapperAPI(swaggerSpec)

	handlers.SetupAPI(api, &handlers.SetupOpts{
		Logger:        logger,
		JobsProcessor: jobsProcesor,
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

	err = server.Serve()
	if err != nil {
		panic(fmt.Errorf("error serve server: %w", err))
	}
}

func panicErr(err error, logger *zap.Logger) {
	if err != nil {
		logger.Panic("can not start app", zap.Error(err))
	}
}
