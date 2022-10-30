package handlers

import (
	"github.com/inhuman/emo_recognizer_controller/internal/jobprocessor"
	"github.com/inhuman/emo_recognizer_controller/pkg/gen/models"
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
	api.JobGetJobOriginalFileHandler = NewGetJobFileHandler(opts.Logger, opts.JobsProcessor)
}

type CommonHandler struct {
	logger       *zap.Logger
	jobProcessor *jobprocessor.JobProcessor
}

type CommonErrorResponseBuilder struct {
	Code    int64
	Details []interface{}
	Error   string
}

func CommonErrorResponse() *CommonErrorResponseBuilder {
	return &CommonErrorResponseBuilder{}
}

func (b *CommonErrorResponseBuilder) WithHTTPCode(statusCode int64) *CommonErrorResponseBuilder {
	b.Code = statusCode

	return b
}

func (b *CommonErrorResponseBuilder) WithError(err error) *CommonErrorResponseBuilder {
	b.Error = err.Error()

	return b
}

func (b *CommonErrorResponseBuilder) WithDetails(i interface{}) *CommonErrorResponseBuilder {
	b.Details = append(b.Details, i)

	return b
}

func (b *CommonErrorResponseBuilder) Build() *models.CommonErrorResponse {
	resp := &models.CommonErrorResponse{}

	if b.Code != 0 {
		resp.Code = b.Code
	}

	if b.Error != "" {
		resp.Error = b.Error
	}

	if len(b.Details) > 0 {
		resp.Details = b.Details
	}

	return resp
}
