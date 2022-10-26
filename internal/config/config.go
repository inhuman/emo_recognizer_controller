package config

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
	"go.uber.org/zap"
)

type Config struct {
	Port int `env:"PORT"`
	//Services Services
	//Db       Database
	//S3       S3
}

type Services struct {
	NoiseWrapperAddress     string `env:"NOISE_WRAPPER_ADDRESS"`
	SpeechRecognizerAddress string `env:"SPEECH_RECOGNIZER_ADDRESS"`
}

type Database struct {
	Host   string `env:"PGHOST"`
	Port   int    `env:"PGPORT"`
	User   string `env:"PGUSER"`
	DbName string `env:"PGDATABASE"`
}

type S3 struct {
	Endpoint         string `env:"S3_ENDPOINT"`
	AccessKey        string `env:"=3BIaGF7ba4o8j5BltAtx"`
	SecretKey        string `env:"S3_SECRET_ACCESS_KEY"`
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
