// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/inhuman/emo_recognizer_controller/pkg/gen/models"
)

// CreateJobOKCode is the HTTP code returned for type CreateJobOK
const CreateJobOKCode int = 200

/*
CreateJobOK create job o k

swagger:response createJobOK
*/
type CreateJobOK struct {

	/*
	  In: Body
	*/
	Payload *CreateJobOKBody `json:"body,omitempty"`
}

// NewCreateJobOK creates CreateJobOK with default headers values
func NewCreateJobOK() *CreateJobOK {

	return &CreateJobOK{}
}

// WithPayload adds the payload to the create job o k response
func (o *CreateJobOK) WithPayload(payload *CreateJobOKBody) *CreateJobOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create job o k response
func (o *CreateJobOK) SetPayload(payload *CreateJobOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateJobOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateJobBadRequestCode is the HTTP code returned for type CreateJobBadRequest
const CreateJobBadRequestCode int = 400

/*
CreateJobBadRequest Bad data (400)

swagger:response createJobBadRequest
*/
type CreateJobBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.CommonErrorResponse `json:"body,omitempty"`
}

// NewCreateJobBadRequest creates CreateJobBadRequest with default headers values
func NewCreateJobBadRequest() *CreateJobBadRequest {

	return &CreateJobBadRequest{}
}

// WithPayload adds the payload to the create job bad request response
func (o *CreateJobBadRequest) WithPayload(payload *models.CommonErrorResponse) *CreateJobBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create job bad request response
func (o *CreateJobBadRequest) SetPayload(payload *models.CommonErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateJobBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateJobInternalServerErrorCode is the HTTP code returned for type CreateJobInternalServerError
const CreateJobInternalServerErrorCode int = 500

/*
CreateJobInternalServerError Internal server error (500)

swagger:response createJobInternalServerError
*/
type CreateJobInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.CommonErrorResponse `json:"body,omitempty"`
}

// NewCreateJobInternalServerError creates CreateJobInternalServerError with default headers values
func NewCreateJobInternalServerError() *CreateJobInternalServerError {

	return &CreateJobInternalServerError{}
}

// WithPayload adds the payload to the create job internal server error response
func (o *CreateJobInternalServerError) WithPayload(payload *models.CommonErrorResponse) *CreateJobInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create job internal server error response
func (o *CreateJobInternalServerError) SetPayload(payload *models.CommonErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateJobInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}