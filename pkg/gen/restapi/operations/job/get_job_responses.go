// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/inhuman/emo_recognizer_controller/pkg/gen/models"
)

// GetJobOKCode is the HTTP code returned for type GetJobOK
const GetJobOKCode int = 200

/*
GetJobOK get job o k

swagger:response getJobOK
*/
type GetJobOK struct {

	/*
	  In: Body
	*/
	Payload *models.Job `json:"body,omitempty"`
}

// NewGetJobOK creates GetJobOK with default headers values
func NewGetJobOK() *GetJobOK {

	return &GetJobOK{}
}

// WithPayload adds the payload to the get job o k response
func (o *GetJobOK) WithPayload(payload *models.Job) *GetJobOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get job o k response
func (o *GetJobOK) SetPayload(payload *models.Job) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetJobOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetJobBadRequestCode is the HTTP code returned for type GetJobBadRequest
const GetJobBadRequestCode int = 400

/*
GetJobBadRequest Bad data (400)

swagger:response getJobBadRequest
*/
type GetJobBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.CommonErrorResponse `json:"body,omitempty"`
}

// NewGetJobBadRequest creates GetJobBadRequest with default headers values
func NewGetJobBadRequest() *GetJobBadRequest {

	return &GetJobBadRequest{}
}

// WithPayload adds the payload to the get job bad request response
func (o *GetJobBadRequest) WithPayload(payload *models.CommonErrorResponse) *GetJobBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get job bad request response
func (o *GetJobBadRequest) SetPayload(payload *models.CommonErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetJobBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetJobNotFoundCode is the HTTP code returned for type GetJobNotFound
const GetJobNotFoundCode int = 404

/*
GetJobNotFound Not found (404)

swagger:response getJobNotFound
*/
type GetJobNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.CommonErrorResponse `json:"body,omitempty"`
}

// NewGetJobNotFound creates GetJobNotFound with default headers values
func NewGetJobNotFound() *GetJobNotFound {

	return &GetJobNotFound{}
}

// WithPayload adds the payload to the get job not found response
func (o *GetJobNotFound) WithPayload(payload *models.CommonErrorResponse) *GetJobNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get job not found response
func (o *GetJobNotFound) SetPayload(payload *models.CommonErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetJobNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetJobInternalServerErrorCode is the HTTP code returned for type GetJobInternalServerError
const GetJobInternalServerErrorCode int = 500

/*
GetJobInternalServerError Internal server error (500)

swagger:response getJobInternalServerError
*/
type GetJobInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.CommonErrorResponse `json:"body,omitempty"`
}

// NewGetJobInternalServerError creates GetJobInternalServerError with default headers values
func NewGetJobInternalServerError() *GetJobInternalServerError {

	return &GetJobInternalServerError{}
}

// WithPayload adds the payload to the get job internal server error response
func (o *GetJobInternalServerError) WithPayload(payload *models.CommonErrorResponse) *GetJobInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get job internal server error response
func (o *GetJobInternalServerError) SetPayload(payload *models.CommonErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetJobInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
