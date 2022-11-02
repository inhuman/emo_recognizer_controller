package handlers

import (
	"bytes"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/inhuman/emo_recognizer_common/jobs"
	"github.com/inhuman/emo_recognizer_controller/internal/jobprocessor"
	"github.com/inhuman/emo_recognizer_controller/internal/repository"
	"github.com/inhuman/emo_recognizer_controller/pkg/gen/models"
	"github.com/inhuman/emo_recognizer_controller/pkg/gen/restapi/operations/job"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type CreateJobHandler struct {
	CommonHandler
}

func NewCreateJobHandler(logger *zap.Logger, jobProcessor *jobprocessor.JobProcessor) *CreateJobHandler {
	return &CreateJobHandler{
		CommonHandler{
			logger:       logger,
			jobProcessor: jobProcessor,
		},
	}
}

func (u *CreateJobHandler) Handle(params job.CreateJobParams) middleware.Responder {
	jobToCreate := jobs.Job{
		Filename: "test.wav",
	}

	err := u.jobProcessor.Repo().CreateJob(params.HTTPRequest.Context(), &jobToCreate)
	if err != nil {
		return job.NewCreateJobInternalServerError().WithPayload(
			CommonErrorResponse().
				WithHTTPCode(http.StatusInternalServerError).
				WithError(err).
				Build(),
		)
	}

	buf := &bytes.Buffer{}
	length, err := io.Copy(buf, params.File)
	if err != nil {
		errRespBuilder := CommonErrorResponse().WithHTTPCode(http.StatusInternalServerError).WithError(err)

		updStatusErr := u.jobProcessor.Repo(). // TODO: fix status to obStatusFileUploadError
							UpdateStatusByUUID(params.HTTPRequest.Context(), jobToCreate.UUID, jobs.JobStatusSpeechRecognizeError)
		if err != nil {
			errRespBuilder.WithDetails(updStatusErr.Error())
		}

		return job.NewCreateJobInternalServerError().WithPayload(errRespBuilder.Build())
	}

	err = u.jobProcessor.FileStorage().Write(params.HTTPRequest.Context(), jobToCreate.OriginalFileName(), length, buf)
	if err != nil {
		return job.NewCreateJobInternalServerError().WithPayload(CommonErrorResponse().
			WithHTTPCode(http.StatusInternalServerError).
			WithError(err).
			Build(),
		)
	}

	updStatusErr := u.jobProcessor.Repo().
		UpdateStatusByUUID(params.HTTPRequest.Context(), jobToCreate.UUID, jobs.JobStatusFileUploaded)
	if updStatusErr != nil {
		return job.NewCreateJobInternalServerError().WithPayload(CommonErrorResponse().
			WithHTTPCode(http.StatusInternalServerError).
			WithError(updStatusErr).
			Build(),
		)
	}

	return job.NewCreateJobOK().WithPayload(&job.CreateJobOKBody{
		UUID: jobToCreate.UUID,
	})
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
		return job.NewGetJobsInternalServerError().WithPayload(
			CommonErrorResponse().
				WithHTTPCode(http.StatusInternalServerError).
				WithError(err).
				Build(),
		)
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
		return job.NewGetJobsInternalServerError().WithPayload(
			CommonErrorResponse().
				WithHTTPCode(http.StatusInternalServerError).
				WithError(err).
				Build(),
		)
	}

	if jobFromRepo == nil {
		return job.NewGetJobNotFound().WithPayload(
			CommonErrorResponse().
				WithHTTPCode(http.StatusNotFound).
				WithError(fmt.Errorf("job with uuid %s not found", params.UUID)).
				Build(),
		)
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

type GetJobFileHandler struct {
	CommonHandler
}

func NewGetJobFileHandler(logger *zap.Logger, jobProcessor *jobprocessor.JobProcessor) *GetJobFileHandler {
	return &GetJobFileHandler{
		CommonHandler{
			logger:       logger,
			jobProcessor: jobProcessor,
		},
	}
}

func (g *GetJobFileHandler) Handle(params job.GetJobOriginalFileParams) middleware.Responder {

	tempjob := &jobs.Job{UUID: params.UUID}

	rdr, err := g.jobProcessor.FileStorage().Read(params.HTTPRequest.Context(), tempjob.OriginalFileName())
	if err != nil {
		return job.NewGetJobOriginalFileInternalServerError().WithPayload(
			CommonErrorResponse().
				WithHTTPCode(http.StatusInternalServerError).
				WithError(err).
				Build(),
		)
	}

	return job.NewGetJobOriginalFileOK().WithPayload(rdr)
}
