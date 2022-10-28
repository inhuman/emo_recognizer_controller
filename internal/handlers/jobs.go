package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/inhuman/emo_recognizer_common/jobs"
	"github.com/inhuman/emo_recognizer_controller/internal/jobprocessor"
	"github.com/inhuman/emo_recognizer_controller/internal/repository"
	"github.com/inhuman/emo_recognizer_controller/pkg/gen/models"
	"github.com/inhuman/emo_recognizer_controller/pkg/gen/restapi/operations/job"
	"go.uber.org/zap"
)

type UploadFileHandler struct {
	CommonHandler
}

func NewUploadFileHandler(logger *zap.Logger, jobProcessor *jobprocessor.JobProcessor) *UploadFileHandler {
	return &UploadFileHandler{
		CommonHandler{
			logger:       logger,
			jobProcessor: jobProcessor,
		},
	}
}

func (u *UploadFileHandler) Handle(params job.CreateJobParams) middleware.Responder {
	// TODO: upload to s3, change status

	return job.NewCreateJobOK()
}

type GetJobsHandler struct {
	CommonHandler
}

func NewGetJobsHandler(logger *zap.Logger, jobProcessor *jobprocessor.JobProcessor) *GetJobsHandler {
	return &GetJobsHandler{
		CommonHandler{
			logger:       logger,
			jobProcessor: jobProcessor,
		},
	}
}

func (g *GetJobsHandler) Handle(params job.GetJobsParams) middleware.Responder {
	filter := repository.GetJobsFilter{}

	if params.Status != nil {
		filter.Status = jobs.JobStatus(*params.Status)
	}

	if params.Limit != nil {
		filter.Limit = int(*params.Limit)
	}

	if params.Offset != nil {
		filter.Offset = int(*params.Offset)
	}

	jobsFromRepo, err := g.jobProcessor.Repo().GetJobs(params.HTTPRequest.Context(), filter)
	if err != nil {
		return job.NewGetJobsInternalServerError().WithPayload(&models.CommonErrorResponse{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		})
	}

	return job.NewGetJobsOK().WithPayload(jobsToDto(jobsFromRepo))
}

type GetJobHandler struct {
	CommonHandler
}

func NewGetJobHandler(logger *zap.Logger, jobProcessor *jobprocessor.JobProcessor) *GetJobHandler {
	return &GetJobHandler{
		CommonHandler{
			logger:       logger,
			jobProcessor: jobProcessor,
		},
	}
}

func (g *GetJobHandler) Handle(params job.GetJobParams) middleware.Responder {
	jobFromRepo, err := g.jobProcessor.Repo().GetJobByUUID(params.HTTPRequest.Context(), params.UUID)
	if err != nil {
		return job.NewGetJobsInternalServerError().WithPayload(&models.CommonErrorResponse{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		})
	}

	if jobFromRepo == nil {
		return job.NewGetJobNotFound().WithPayload(&models.CommonErrorResponse{
			Code:  http.StatusNotFound,
			Error: fmt.Sprintf("job with uuid %s not found", params.UUID),
		})
	}

	return job.NewGetJobOK().WithPayload(jobToDto(jobFromRepo))
}

func jobToDto(jobFromRepo *jobs.Job) *models.Job {
	return &models.Job{
		CreatedAt: strfmt.DateTime(jobFromRepo.CreatedAt),
		Filename:  jobFromRepo.Filename,
		Status:    models.JobStatus(jobFromRepo.Status),
		UUID:      jobFromRepo.UUID,
		UpdatedAt: strfmt.DateTime(jobFromRepo.UpdatedAt),
	}
}

func jobsToDto(jobsFromRepo []*jobs.Job) []*models.Job {
	var dtoJobs []*models.Job

	for i := range jobsFromRepo {
		dtoJobs = append(dtoJobs, jobToDto(jobsFromRepo[i]))
	}

	return dtoJobs
}
