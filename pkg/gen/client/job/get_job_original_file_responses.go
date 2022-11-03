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

// GetJobOriginalFileReader is a Reader for the GetJobOriginalFile structure.
type GetJobOriginalFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetJobOriginalFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetJobOriginalFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetJobOriginalFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetJobOriginalFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetJobOriginalFileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetJobOriginalFileOK creates a GetJobOriginalFileOK with default headers values
func NewGetJobOriginalFileOK() *GetJobOriginalFileOK {
	return &GetJobOriginalFileOK{}
}

/*
GetJobOriginalFileOK describes a response with status code 200, with default header values.

GetJobOriginalFileOK get job original file o k
*/
type GetJobOriginalFileOK struct {
	Payload models.Reader
}

// IsSuccess returns true when this get job original file o k response has a 2xx status code
func (o *GetJobOriginalFileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get job original file o k response has a 3xx status code
func (o *GetJobOriginalFileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get job original file o k response has a 4xx status code
func (o *GetJobOriginalFileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get job original file o k response has a 5xx status code
func (o *GetJobOriginalFileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get job original file o k response a status code equal to that given
func (o *GetJobOriginalFileOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetJobOriginalFileOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/jobs/{Uuid}/file/original][%d] getJobOriginalFileOK  %+v", 200, o.Payload)
}

func (o *GetJobOriginalFileOK) String() string {
	return fmt.Sprintf("[GET /api/v1/jobs/{Uuid}/file/original][%d] getJobOriginalFileOK  %+v", 200, o.Payload)
}

func (o *GetJobOriginalFileOK) GetPayload() models.Reader {
	return o.Payload
}

func (o *GetJobOriginalFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJobOriginalFileBadRequest creates a GetJobOriginalFileBadRequest with default headers values
func NewGetJobOriginalFileBadRequest() *GetJobOriginalFileBadRequest {
	return &GetJobOriginalFileBadRequest{}
}

/*
GetJobOriginalFileBadRequest describes a response with status code 400, with default header values.

Bad data (400)
*/
type GetJobOriginalFileBadRequest struct {
	Payload *models.CommonErrorResponse
}

// IsSuccess returns true when this get job original file bad request response has a 2xx status code
func (o *GetJobOriginalFileBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get job original file bad request response has a 3xx status code
func (o *GetJobOriginalFileBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get job original file bad request response has a 4xx status code
func (o *GetJobOriginalFileBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get job original file bad request response has a 5xx status code
func (o *GetJobOriginalFileBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get job original file bad request response a status code equal to that given
func (o *GetJobOriginalFileBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetJobOriginalFileBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/jobs/{Uuid}/file/original][%d] getJobOriginalFileBadRequest  %+v", 400, o.Payload)
}

func (o *GetJobOriginalFileBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/jobs/{Uuid}/file/original][%d] getJobOriginalFileBadRequest  %+v", 400, o.Payload)
}

func (o *GetJobOriginalFileBadRequest) GetPayload() *models.CommonErrorResponse {
	return o.Payload
}

func (o *GetJobOriginalFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJobOriginalFileNotFound creates a GetJobOriginalFileNotFound with default headers values
func NewGetJobOriginalFileNotFound() *GetJobOriginalFileNotFound {
	return &GetJobOriginalFileNotFound{}
}

/*
GetJobOriginalFileNotFound describes a response with status code 404, with default header values.

Not found (404)
*/
type GetJobOriginalFileNotFound struct {
	Payload *models.CommonErrorResponse
}

// IsSuccess returns true when this get job original file not found response has a 2xx status code
func (o *GetJobOriginalFileNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get job original file not found response has a 3xx status code
func (o *GetJobOriginalFileNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get job original file not found response has a 4xx status code
func (o *GetJobOriginalFileNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get job original file not found response has a 5xx status code
func (o *GetJobOriginalFileNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get job original file not found response a status code equal to that given
func (o *GetJobOriginalFileNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetJobOriginalFileNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/jobs/{Uuid}/file/original][%d] getJobOriginalFileNotFound  %+v", 404, o.Payload)
}

func (o *GetJobOriginalFileNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/jobs/{Uuid}/file/original][%d] getJobOriginalFileNotFound  %+v", 404, o.Payload)
}

func (o *GetJobOriginalFileNotFound) GetPayload() *models.CommonErrorResponse {
	return o.Payload
}

func (o *GetJobOriginalFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJobOriginalFileInternalServerError creates a GetJobOriginalFileInternalServerError with default headers values
func NewGetJobOriginalFileInternalServerError() *GetJobOriginalFileInternalServerError {
	return &GetJobOriginalFileInternalServerError{}
}

/*
GetJobOriginalFileInternalServerError describes a response with status code 500, with default header values.

Internal server error (500)
*/
type GetJobOriginalFileInternalServerError struct {
	Payload *models.CommonErrorResponse
}

// IsSuccess returns true when this get job original file internal server error response has a 2xx status code
func (o *GetJobOriginalFileInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get job original file internal server error response has a 3xx status code
func (o *GetJobOriginalFileInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get job original file internal server error response has a 4xx status code
func (o *GetJobOriginalFileInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get job original file internal server error response has a 5xx status code
func (o *GetJobOriginalFileInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get job original file internal server error response a status code equal to that given
func (o *GetJobOriginalFileInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetJobOriginalFileInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/jobs/{Uuid}/file/original][%d] getJobOriginalFileInternalServerError  %+v", 500, o.Payload)
}

func (o *GetJobOriginalFileInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/jobs/{Uuid}/file/original][%d] getJobOriginalFileInternalServerError  %+v", 500, o.Payload)
}

func (o *GetJobOriginalFileInternalServerError) GetPayload() *models.CommonErrorResponse {
	return o.Payload
}

func (o *GetJobOriginalFileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
