package jobprocessor

import (
	"context"
	"fmt"
	"github.com/go-openapi/runtime"
	"github.com/inhuman/noise_wrapper/pkg/gen/client/noise_wrap"
	"go.uber.org/zap"
	"sync"
	"time"

	"github.com/inhuman/emo_recognizer_common/jobs"
	"github.com/inhuman/emo_recognizer_controller/internal/repository"
	nwclient "github.com/inhuman/noise_wrapper/pkg/gen/client"
)

const (
	StrategyDefault  jobs.Strategy = "default"
	StrategyLongFile jobs.Strategy = "long-file"
)

type ProcessStrategyChooser interface {
	ChooseStrategy(jobToProcess *jobs.Job) ProcessStrategy
	AddStrategy(name jobs.Strategy, strategy ProcessStrategy)
}

type StrategyChooser struct {
	strategies map[jobs.Strategy]ProcessStrategy
	mu         sync.RWMutex
}

func (s *StrategyChooser) AddStrategy(name jobs.Strategy, strategy ProcessStrategy) {
	s.mu.Lock()
	s.strategies[name] = strategy
	s.mu.Unlock()
}

func NewStrategyChooser() *StrategyChooser {
	return &StrategyChooser{
		strategies: make(map[jobs.Strategy]ProcessStrategy),
	}
}

func (s *StrategyChooser) ChooseStrategy(jobToProcess *jobs.Job) ProcessStrategy {
	// пока используем только дефолтовую стратегию
	s.mu.RLock()
	strategy := s.strategies[StrategyDefault]
	s.mu.RUnlock()

	return strategy
}

type ProcessStrategy interface {
	Process(ctx context.Context, jobToProcess *jobs.Job) error
	Name() jobs.Strategy
}

type DefaultStrategy struct {
	repo               repository.Repository
	noiseWrapperClient *nwclient.NoiseWrapper
	storageClient      Storage
	logger             *zap.Logger
}

type DefaultStrategyOps struct {
	Repo               repository.Repository
	NoiseWrapperClient *nwclient.NoiseWrapper
	StorageClient      Storage
	Logger             *zap.Logger
}

func NewDefaultStrategy(opts DefaultStrategyOps) *DefaultStrategy {
	return &DefaultStrategy{
		repo:               opts.Repo,
		noiseWrapperClient: opts.NoiseWrapperClient,
		storageClient:      opts.StorageClient,
		logger:             opts.Logger,
	}
}

func (d *DefaultStrategy) Process(ctx context.Context, jobToProcess *jobs.Job) error {
	switch jobToProcess.Status {
	case jobs.JobStatusFileUploaded:

		contextWithTimeout, cancel := context.WithTimeout(context.Background(), time.Second*120)
		defer cancel()

		err := d.repo.UpdateStatusByUUID(ctx, jobToProcess.UUID, jobs.JobStatusNoiseWrapStarted)
		if err != nil {
			return fmt.Errorf("error update job status")
		}

		file, err := d.storageClient.Read(ctx, jobToProcess.OriginalFileName())
		if err != nil {
			return fmt.Errorf("can not download file from storage: %w", err)
		}

		runtime.NamedReader("file", file)

		resp, err := d.noiseWrapperClient.NoiseWrap.UploadFile(&noise_wrap.UploadFileParams{
			File:    runtime.NamedReader("file", file),
			UUID:    jobToProcess.UUID,
			Context: contextWithTimeout,
		})
		if err != nil {
			return fmt.Errorf("can not upload file to noise wrapper: %w", err)
		}

		d.logger.Info("successful upload file to noise wrapper")

		if resp.IsSuccess() {
			err := d.repo.UpdateStatusByUUID(ctx, jobToProcess.UUID, jobs.JobStatusNoiseWrapComplete)
			if err != nil {
				return fmt.Errorf("error update job status")
			}
		}

	case jobs.JobStatusNoiseWrapComplete:
		// TODO: to speech recognizer
	}

	return nil
}

func (d *DefaultStrategy) Name() jobs.Strategy {
	return "default strategy"
}
