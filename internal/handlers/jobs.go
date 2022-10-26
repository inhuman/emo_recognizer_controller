package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/inhuman/emo_recognizer_controller/internal/controller"

	"github.com/inhuman/emo_recognizer_controller/pkg/gen/restapi/operations/job"
	"go.uber.org/zap"
)

type UploadFileHandler struct {
	CommonHandler
}

func NewUploadFileHandler(logger *zap.Logger, jobProcessor *controller.JobProcessor) *UploadFileHandler {
	return &UploadFileHandler{
		CommonHandler{
			logger:       logger,
			jobProcessor: jobProcessor,
		},
	}
}

func (u *UploadFileHandler) Handle(params job.CreateJobParams) middleware.Responder {
	//result, err := u.CommonHandler.denoiser.CleanSound(params.File)
	//if err != nil {
	//	return noise_wrap.NewUploadFileInternalServerError().
	//		WithPayload(&models.CommonErrorResponse{
	//			Code:  http.StatusInternalServerError,
	//			Error: err.Error(),
	//		})
	//}

	return job.NewCreateJobOK()
}
