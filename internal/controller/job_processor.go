package controller

import (
	"context"
	"github.com/inhuman/emo_recognizer_controller/internal/config"
	"github.com/inhuman/emo_recognizer_controller/internal/repository"
	"go.uber.org/zap"
	"time"
)

type JobProcessor struct {
	repo            repository.Repository
	logger          *zap.Logger
	strategyChooser ProcessStrategyChooser
}

func (jp *JobProcessor) Run(ctx context.Context, conf config.JobProcessor) {

	ticker := time.NewTicker(conf.FetchJobsPeriod)

	for {
		select {
		case <-ticker.C:

			jobForProcess, err := jp.repo.GetJobToProcess(ctx)
			if err != nil {
				jp.logger.Error("error get job to process: %w", zap.Error(err))

				continue
			}

			if jobForProcess.UUID == "" {
				continue
			}

			jp.logger.Info("found job to process",
				zap.String("status", string(jobForProcess.Status)),
				zap.String("uuid", jobForProcess.UUID),
				zap.String("file name", jobForProcess.Filename),
			)

			strategy := jp.strategyChooser.ChooseStrategy(jobForProcess)
			jp.logger.Info("picked strategy to process",
				zap.String("job status", string(jobForProcess.Status)),
				zap.String("job uuid", jobForProcess.UUID),
				zap.String("job file name", jobForProcess.Filename),
				zap.String("strategy name", string(strategy.Name())),
			)

			err = strategy.Process(jobForProcess)
			if err != nil {
				jp.logger.Error("job process error",
					zap.String("job status", string(jobForProcess.Status)),
					zap.String("job uuid", jobForProcess.UUID),
					zap.String("job file name", jobForProcess.Filename),
					zap.String("strategy name", string(strategy.Name())),
					zap.Error(err),
				)
			}

		case <-ctx.Done():
			return
		}
	}
}
