package handlers

import (
	"github.com/inhuman/emo_recognizer_controller/internal/jobprocessor"
	"github.com/inhuman/emo_recognizer_controller/pkg/gen/restapi/operations"
	"go.uber.org/zap"
)

type SetupOpts struct {
	Logger        *zap.Logger
	JobsProcessor *jobprocessor.JobProcessor
}

func SetupAPI(api *operations.EmotionsRecognizerAPI, opts *SetupOpts) {
	api.JobCreateJobHandler = NewUploadFileHandler(opts.Logger, opts.JobsProcessor)
	api.JobGetJobsHandler = NewGetJobsHandler(opts.Logger, opts.JobsProcessor)
	api.JobGetJobHandler = NewGetJobHandler(opts.Logger, opts.JobsProcessor)
}

type CommonHandler struct {
	logger       *zap.Logger
	jobProcessor *jobprocessor.JobProcessor
}
