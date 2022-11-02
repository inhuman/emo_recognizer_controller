package config

import (
	"context"
	"time"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
	"go.uber.org/zap"
)

type Config struct {
	Port              int `env:"PORT,default=80"`
	Db                Database
	JobProcessor      JobProcessor
	S3                S3
	Services          Services
	ApplyDbMigrations bool   `env:"APPLY_MIGRATIONS,default=false"`
	MigrationsPath    string `env:"MIGRATIONS_PATH,default=file:///migrations/"`
}

type JobProcessor struct {
	FetchJobsPeriod time.Duration `env:"FETCH_JOBS_PERIOD,default=10s"`
}

type Services struct {
	NoiseWrapper     NoiseWrapper
	SpeechRecognizer SpeechRecognizer
}

type NoiseWrapper struct {
	Address string `env:"NOISE_WRAPPER_ADDRESS"`
}

type SpeechRecognizer struct {
	Address string `env:"SPEECH_RECOGNIZER_ADDRESS"`
}

type Database struct {
	Host     string `env:"PGHOST"`
	Port     int    `env:"PGPORT"`
	User     string `env:"PGUSER"`
	DbName   string `env:"PGDATABASE"`
	Password string `env:"PGPASSWORD"`
	SslMode  string `env:"PGSSLMODE,default=disable"`
}

type S3 struct {
	Endpoint         string `env:"S3_ENDPOINT"`
	AccessKey        string `env:"S3_ACCESS_KEY"`
	SecretKey        string `env:"S3_SECRET_KEY"`
	Bucket           string `env:"S3_BUCKET_NAME"`
	PublicHostAddr   string `env:"S3_PUBLIC_HOST_ADDR"`
	PublicHostSchema string `env:"S3_PUBLIC_HOST_SCHEMA"`
	Secure           bool   `env:"S3_SECURE"`
}

func InitConfig(ctx context.Context, logger *zap.Logger) (Config, error) {
	if godotenv.Load() != nil {
		logger.Info("envs loaded from OS")
	} else {
		logger.Info("env loaded from files")
	}

	conf := Config{}

	if err := envconfig.Process(ctx, &conf); err != nil {
		logger.Panic("failed load config", zap.Error(err))
	}

	return conf, nil
}
