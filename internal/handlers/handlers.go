package handlers

import (
	"github.com/inhuman/emo_recognizer_controller/internal/controller"
	"github.com/inhuman/emo_recognizer_controller/pkg/gen/restapi/operations"
)
import "go.uber.org/zap"

type SetupOpts struct {
	Logger        *zap.Logger
	JobsProcessor *controller.JobProcessor
}

func SetupAPI(api *operations.NoiseWrapperAPI, opts *SetupOpts) {
	api.JobCreateJobHandler = NewUploadFileHandler(opts.Logger, opts.JobsProcessor)
}

type CommonHandler struct {
	logger       *zap.Logger
	jobProcessor *controller.JobProcessor
}
