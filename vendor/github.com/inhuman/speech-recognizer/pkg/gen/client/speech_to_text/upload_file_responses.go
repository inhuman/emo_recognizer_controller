// Code generated by go-swagger; DO NOT EDIT.

package speech_to_text

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/inhuman/speech-recognizer/pkg/gen/models"
)

// UploadFileReader is a Reader for the UploadFile structure.
type UploadFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUploadFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUploadFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUploadFileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUploadFileOK creates a UploadFileOK with default headers values
func NewUploadFileOK() *UploadFileOK {
	return &UploadFileOK{}
}

/*
UploadFileOK describes a response with status code 200, with default header values.

UploadFileOK upload file o k
*/
type UploadFileOK struct {
	Payload *models.RecognizeResult
}

// IsSuccess returns true when this upload file o k response has a 2xx status code
func (o *UploadFileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this upload file o k response has a 3xx status code
func (o *UploadFileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload file o k response has a 4xx status code
func (o *UploadFileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this upload file o k response has a 5xx status code
func (o *UploadFileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this upload file o k response a status code equal to that given
func (o *UploadFileOK) IsCode(code int) bool {
	return code == 200
}

func (o *UploadFileOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/upload/{Uuid}][%d] uploadFileOK  %+v", 200, o.Payload)
}

func (o *UploadFileOK) String() string {
	return fmt.Sprintf("[POST /api/v1/upload/{Uuid}][%d] uploadFileOK  %+v", 200, o.Payload)
}

func (o *UploadFileOK) GetPayload() *models.RecognizeResult {
	return o.Payload
}

func (o *UploadFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RecognizeResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadFileBadRequest creates a UploadFileBadRequest with default headers values
func NewUploadFileBadRequest() *UploadFileBadRequest {
	return &UploadFileBadRequest{}
}

/*
UploadFileBadRequest describes a response with status code 400, with default header values.

Bad data (400)
*/
type UploadFileBadRequest struct {
	Payload *models.CommonErrorResponse
}

// IsSuccess returns true when this upload file bad request response has a 2xx status code
func (o *UploadFileBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload file bad request response has a 3xx status code
func (o *UploadFileBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload file bad request response has a 4xx status code
func (o *UploadFileBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this upload file bad request response has a 5xx status code
func (o *UploadFileBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this upload file bad request response a status code equal to that given
func (o *UploadFileBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UploadFileBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/upload/{Uuid}][%d] uploadFileBadRequest  %+v", 400, o.Payload)
}

func (o *UploadFileBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/upload/{Uuid}][%d] uploadFileBadRequest  %+v", 400, o.Payload)
}

func (o *UploadFileBadRequest) GetPayload() *models.CommonErrorResponse {
	return o.Payload
}

func (o *UploadFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadFileInternalServerError creates a UploadFileInternalServerError with default headers values
func NewUploadFileInternalServerError() *UploadFileInternalServerError {
	return &UploadFileInternalServerError{}
}

/*
UploadFileInternalServerError describes a response with status code 500, with default header values.

Internal server error (500)
*/
type UploadFileInternalServerError struct {
	Payload *models.CommonErrorResponse
}

// IsSuccess returns true when this upload file internal server error response has a 2xx status code
func (o *UploadFileInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload file internal server error response has a 3xx status code
func (o *UploadFileInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload file internal server error response has a 4xx status code
func (o *UploadFileInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this upload file internal server error response has a 5xx status code
func (o *UploadFileInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this upload file internal server error response a status code equal to that given
func (o *UploadFileInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UploadFileInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/upload/{Uuid}][%d] uploadFileInternalServerError  %+v", 500, o.Payload)
}

func (o *UploadFileInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/upload/{Uuid}][%d] uploadFileInternalServerError  %+v", 500, o.Payload)
}

func (o *UploadFileInternalServerError) GetPayload() *models.CommonErrorResponse {
	return o.Payload
}

func (o *UploadFileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
