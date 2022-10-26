package docs

import (
	"bytes"
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
