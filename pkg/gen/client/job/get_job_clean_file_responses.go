// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/inhuman/emo_recognizer_controller/pkg/gen/models"
)

// GetJobCleanFileReader is a Reader for the GetJobCleanFile structure.
type GetJobCleanFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetJobCleanFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetJobCleanFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetJobCleanFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetJobCleanFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetJobCleanFileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetJobCleanFileOK creates a GetJobCleanFileOK with default headers values
func NewGetJobCleanFileOK() *GetJobCleanFileOK {
	return &GetJobCleanFileOK{}
}

/*
GetJobCleanFileOK describes a response with status code 200, with default header values.

GetJobCleanFileOK get job clean file o k
*/
type GetJobCleanFileOK struct {
	Payload models.Reader
}

// IsSuccess returns true when this get job clean file o k response has a 2xx status code
func (o *GetJobCleanFileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get job clean file o k response has a 3xx status code
func (o *GetJobCleanFileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get job clean file o k response has a 4xx status code
func (o *GetJobCleanFileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get job clean file o k response has a 5xx status code
func (o *GetJobCleanFileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get job clean file o k response a status code equal to that given
func (o *GetJobCleanFileOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetJobCleanFileOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/jobs/{Uuid}/file/clean][%d] getJobCleanFileOK  %+v", 200, o.Payload)
}

func (o *GetJobCleanFileOK) String() string {
	return fmt.Sprintf("[GET /api/v1/jobs/{Uuid}/file/clean][%d] getJobCleanFileOK  %+v", 200, o.Payload)
}

func (o *GetJobCleanFileOK) GetPayload() models.Reader {
	return o.Payload
}

func (o *GetJobCleanFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJobCleanFileBadRequest creates a GetJobCleanFileBadRequest with default headers values
func NewGetJobCleanFileBadRequest() *GetJobCleanFileBadRequest {
	return &GetJobCleanFileBadRequest{}
}

/*
GetJobCleanFileBadRequest describes a response with status code 400, with default header values.

Bad data (400)
*/
type GetJobCleanFileBadRequest struct {
	Payload *models.CommonErrorResponse
}

// IsSuccess returns true when this get job clean file bad request response has a 2xx status code
func (o *GetJobCleanFileBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get job clean file bad request response has a 3xx status code
func (o *GetJobCleanFileBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get job clean file bad request response has a 4xx status code
func (o *GetJobCleanFileBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get job clean file bad request response has a 5xx status code
func (o *GetJobCleanFileBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get job clean file bad request response a status code equal to that given
func (o *GetJobCleanFileBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetJobCleanFileBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/jobs/{Uuid}/file/clean][%d] getJobCleanFileBadRequest  %+v", 400, o.Payload)
}

func (o *GetJobCleanFileBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/jobs/{Uuid}/file/clean][%d] getJobCleanFileBadRequest  %+v", 400, o.Payload)
}

func (o *GetJobCleanFileBadRequest) GetPayload() *models.CommonErrorResponse {
	return o.Payload
}

func (o *GetJobCleanFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJobCleanFileNotFound creates a GetJobCleanFileNotFound with default headers values
func NewGetJobCleanFileNotFound() *GetJobCleanFileNotFound {
	return &GetJobCleanFileNotFound{}
}

/*
GetJobCleanFileNotFound describes a response with status code 404, with default header values.

Not found (404)
*/
type GetJobCleanFileNotFound struct {
	Payload *models.CommonErrorResponse
}

// IsSuccess returns true when this get job clean file not found response has a 2xx status code
func (o *GetJobCleanFileNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get job clean file not found response has a 3xx status code
func (o *GetJobCleanFileNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get job clean file not found response has a 4xx status code
func (o *GetJobCleanFileNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get job clean file not found response has a 5xx status code
func (o *GetJobCleanFileNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get job clean file not found response a status code equal to that given
func (o *GetJobCleanFileNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetJobCleanFileNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/jobs/{Uuid}/file/clean][%d] getJobCleanFileNotFound  %+v", 404, o.Payload)
}

func (o *GetJobCleanFileNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/jobs/{Uuid}/file/clean][%d] getJobCleanFileNotFound  %+v", 404, o.Payload)
}

func (o *GetJobCleanFileNotFound) GetPayload() *models.CommonErrorResponse {
	return o.Payload
}

func (o *GetJobCleanFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJobCleanFileInternalServerError creates a GetJobCleanFileInternalServerError with default headers values
func NewGetJobCleanFileInternalServerError() *GetJobCleanFileInternalServerError {
	return &GetJobCleanFileInternalServerError{}
}

/*
GetJobCleanFileInternalServerError describes a response with status code 500, with default header values.

Internal server error (500)
*/
type GetJobCleanFileInternalServerError struct {
	Payload *models.CommonErrorResponse
}

// IsSuccess returns true when this get job clean file internal server error response has a 2xx status code
func (o *GetJobCleanFileInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get job clean file internal server error response has a 3xx status code
func (o *GetJobCleanFileInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get job clean file internal server error response has a 4xx status code
func (o *GetJobCleanFileInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get job clean file internal server error response has a 5xx status code
func (o *GetJobCleanFileInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get job clean file internal server error response a status code equal to that given
func (o *GetJobCleanFileInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetJobCleanFileInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/jobs/{Uuid}/file/clean][%d] getJobCleanFileInternalServerError  %+v", 500, o.Payload)
}

func (o *GetJobCleanFileInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/jobs/{Uuid}/file/clean][%d] getJobCleanFileInternalServerError  %+v", 500, o.Payload)
}

func (o *GetJobCleanFileInternalServerError) GetPayload() *models.CommonErrorResponse {
	return o.Payload
}

func (o *GetJobCleanFileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
