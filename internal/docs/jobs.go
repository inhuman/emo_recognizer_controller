package docs

import (
	"bytes"
	"io"

	"github.com/inhuman/emo_recognizer_common/jobs"
)

// swagger:route POST /api/v1/jobs Job createJob
// Эндпоинт для загрузки звукового файла (.wav)
//
// Consumes:
// - multipart/form-data
//
// responses:
//   200: createJobResponse
//   400: badDataResponse
//   500: internalErrorResponse

// swagger:parameters createJob
type createJobParams struct {
	// Звуковой файл в формате wav
	// in:formData
	//
	// swagger:file
	File *bytes.Buffer `json:"file"`
}

// swagger:response createJobResponse
type createJobResponse struct {
	// Результат создания задания на распознавание эмоций
	// in:body
	Body struct {
		UUID string `json:"UUID"`
	}
}

// swagger:route GET /api/v1/jobs Job getJobs
// Эндпоинт для получения списка задач на обработку
//
// responses:
//   200: getJobsResponse
//   400: badDataResponse
//   500: internalErrorResponse

// swagger:parameters getJobs
type getJobsParams struct {
	// Status
	// in:query
	Status string `json:"status"`

	// Strategy
	// in:query
	Strategy string `json:"strategy"`

	// Limit
	// in:query
	Limit int `json:"limit"`

	// Offset
	// in:query
	Offset int `json:"offset"`
}

// swagger:response getJobsResponse
type getJobsResponse struct {
	// Задания на обработку
	// in:body
	Body []jobs.Job
}

// swagger:route GET /api/v1/jobs/{Uuid} Job getJob
// Эндпоинт для получения задачи по UUID
//
//
// responses:
//   200: getJobResponse
//   400: badDataResponse
//	 404: notFoundResponse
//   500: internalErrorResponse

// swagger:parameters getJob
type getJobParams struct {
	// Uuid задания
	// in:path
	Uuid string
}

// swagger:response getJobResponse
type getJobResponse struct {
	// Задание на обработку
	// in:body
	Body jobs.Job
}

// swagger:route GET /api/v1/jobs/{Uuid}/file/original Job getJobOriginalFile
// Эндпоинт для получения файла задачи по UUID
//
// Produces:
// - application/octet-stream
//
// responses:
//   200: getJobOriginalFileResponse
//   400: badDataResponse
//	 404: notFoundResponse
//   500: internalErrorResponse

// swagger:parameters getJobOriginalFile
type getJobFileParams struct {
	// Uuid задания
	// in:path
	Uuid string
}

// swagger:response getJobOriginalFileResponse
type getJobFileResponse struct {
	// Задание на обработку
	// in:body
	Body io.Reader
}
