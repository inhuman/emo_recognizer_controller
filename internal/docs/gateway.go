package docs

import (
	"bytes"
)

// swagger:route POST /api/v1/upload NoiseWrap uploadFile
// Эндпоинт для загрузки звукового файла (.wav)
//
// Consumes:
// - multipart/form-data
//
// responses:
//   200: cleanSoundResponse
//   400: badDataResponse
//   500: internalErrorResponse

// swagger:parameters uploadFile
type uploadFileParams struct {
	// Звуковой файл в формате wav
	// in:formData
	//
	// swagger:file
	File *bytes.Buffer `json:"file"`
}

// swagger:response cleanSoundResponse
type cleanSoundResponse struct {
	// Результат очистки файла
	// in:body
	//Body noise_wrapper.ClearSoundResult
}
