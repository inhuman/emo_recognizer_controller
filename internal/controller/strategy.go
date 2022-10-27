package controller

import (
	"github.com/inhuman/emo_recognizer_common/jobs"
	"github.com/inhuman/emo_recognizer_controller/internal/repository"
	"sync"
)

type StrategyName string

const (
	StrategyDefault  StrategyName = "default"
	StrategyLongFile StrategyName = "long-file"
)

type ProcessStrategyChooser interface {
	ChooseStrategy(jobToProcess *jobs.Job) ProcessStrategy
	AddStrategy(name StrategyName, strategy ProcessStrategy)
}

type StrategyChooser struct {
	strategies map[StrategyName]ProcessStrategy
	mu         sync.RWMutex
}

func (s *StrategyChooser) AddStrategy(name StrategyName, strategy ProcessStrategy) {
	s.mu.Lock()
	s.strategies[name] = strategy
	s.mu.Unlock()
}

func NewStrategyChooser() *StrategyChooser {
	return &StrategyChooser{}
}

func (s *StrategyChooser) ChooseStrategy(jobToProcess *jobs.Job) ProcessStrategy {
	s.mu.RLock()
	strategy := s.strategies[StrategyDefault]
	s.mu.RUnlock()

	return strategy
}

type ProcessStrategy interface {
	Process(jobToProcess *jobs.Job) error
	Name() StrategyName
}

type DefaultStrategy struct {
	repo repository.Repository
}

func NewDefaultStrategy() *DefaultStrategy {
	return &DefaultStrategy{}
}

func (d *DefaultStrategy) Process(jobToProcess *jobs.Job) error {
	switch jobToProcess.Status {
	case jobs.JobStatusPlanned:
		//TODO: to noise wrapper
	case jobs.JobStatusNoiseWrapComplete:
		//TODO: to speech recognizer
	}

	return nil
}

func (d *DefaultStrategy) Name() StrategyName {
	return "default strategy"
}
