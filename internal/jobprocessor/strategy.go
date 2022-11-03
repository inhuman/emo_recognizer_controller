package jobprocessor

import (
	"context"
	"fmt"
	"github.com/go-openapi/runtime"
	"github.com/inhuman/noise_wrapper/pkg/gen/client/noise_wrap"
	"github.com/inhuman/speech-recognizer/pkg/gen/client/speech_to_text"
	"go.uber.org/zap"
	"sync"
	"time"

	"github.com/inhuman/emo_recognizer_common/jobs"
	"github.com/inhuman/emo_recognizer_controller/internal/repository"
	nwclient "github.com/inhuman/noise_wrapper/pkg/gen/client"
	srclient "github.com/inhuman/speech-recognizer/pkg/gen/client"
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
	repo                   repository.Repository
	noiseWrapperClient     *nwclient.NoiseWrapper
	speechRecognizerClient *srclient.SpeechRecognizer
	storageClient          Storage
	logger                 *zap.Logger
}

type DefaultStrategyOps struct {
	Repo                   repository.Repository
	NoiseWrapperClient     *nwclient.NoiseWrapper
	SpeechRecognizerClient *srclient.SpeechRecognizer
	StorageClient          Storage
	Logger                 *zap.Logger
}

func NewDefaultStrategy(opts DefaultStrategyOps) *DefaultStrategy {
	return &DefaultStrategy{
		repo:                   opts.Repo,
		noiseWrapperClient:     opts.NoiseWrapperClient,
		speechRecognizerClient: opts.SpeechRecognizerClient,
		storageClient:          opts.StorageClient,
		logger:                 opts.Logger,
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
			return fmt.Errorf("can not download original file from storage: %w", err)
		}

		resp, err := d.noiseWrapperClient.NoiseWrap.UploadFile(&noise_wrap.UploadFileParams{
			File:    runtime.NamedReader("file", file),
			UUID:    jobToProcess.UUID,
			Context: contextWithTimeout,
		})
		if err != nil {
			return fmt.Errorf("can not upload file to noise wrapper: %w", err)
		}

		d.logger.Info("successful upload file to noise wrapper")

		status := jobs.JobStatusNoiseWrapError

		if resp.IsSuccess() {
			status = jobs.JobStatusNoiseWrapComplete
		}

		err = d.repo.UpdateStatusByUUID(ctx, jobToProcess.UUID, status)
		if err != nil {
			return fmt.Errorf("error update job status")
		}

		return nil

	case jobs.JobStatusNoiseWrapComplete:
		contextWithTimeout, cancel := context.WithTimeout(context.Background(), time.Second*120)
		defer cancel()

		err := d.repo.UpdateStatusByUUID(ctx, jobToProcess.UUID, jobs.JobStatusSpeechRecognizeStarted)
		if err != nil {
			return fmt.Errorf("error update job status")
		}

		file, err := d.storageClient.Read(ctx, jobToProcess.CleanFileName())
		if err != nil {
			return fmt.Errorf("can not download clean file from storage: %w", err)
		}

		resp, err := d.speechRecognizerClient.SpeechToText.UploadFile(&speech_to_text.UploadFileParams{
			File:    runtime.NamedReader("file", file),
			Context: contextWithTimeout,
			UUID:    jobToProcess.UUID,
		})
		if err != nil {
			return fmt.Errorf("can not upload file to speech recognizer: %w", err)
		}

		status := jobs.JobStatusSpeechRecognizeError

		if resp.IsSuccess() {
			status = jobs.JobStatusSpeechRecognizeComplete
		}

		err = d.repo.UpdateStatusByUUID(ctx, jobToProcess.UUID, status)
		if err != nil {
			return fmt.Errorf("error update job status")
		}

		err = d.repo.UpdateRecognizedText(ctx, jobToProcess.UUID, resp.GetPayload().Text)
		if err != nil {
			return fmt.Errorf("error update job recognized text")
		}

		return nil

	default:
		err := d.repo.UpdateStatusByUUID(ctx, jobToProcess.UUID, jobs.JobStatusComplete)
		if err != nil {
			return fmt.Errorf("error update job status")
		}
	}

	return nil
}

func (d *DefaultStrategy) Name() jobs.Strategy {
	return "default strategy"
}
