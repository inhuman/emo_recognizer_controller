// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetJobOriginalFileParams creates a new GetJobOriginalFileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetJobOriginalFileParams() *GetJobOriginalFileParams {
	return &GetJobOriginalFileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetJobOriginalFileParamsWithTimeout creates a new GetJobOriginalFileParams object
// with the ability to set a timeout on a request.
func NewGetJobOriginalFileParamsWithTimeout(timeout time.Duration) *GetJobOriginalFileParams {
	return &GetJobOriginalFileParams{
		timeout: timeout,
	}
}

// NewGetJobOriginalFileParamsWithContext creates a new GetJobOriginalFileParams object
// with the ability to set a context for a request.
func NewGetJobOriginalFileParamsWithContext(ctx context.Context) *GetJobOriginalFileParams {
	return &GetJobOriginalFileParams{
		Context: ctx,
	}
}

// NewGetJobOriginalFileParamsWithHTTPClient creates a new GetJobOriginalFileParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetJobOriginalFileParamsWithHTTPClient(client *http.Client) *GetJobOriginalFileParams {
	return &GetJobOriginalFileParams{
		HTTPClient: client,
	}
}

/*
GetJobOriginalFileParams contains all the parameters to send to the API endpoint

	for the get job original file operation.

	Typically these are written to a http.Request.
*/
type GetJobOriginalFileParams struct {

	/* UUID.

	   Uuid задания
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get job original file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetJobOriginalFileParams) WithDefaults() *GetJobOriginalFileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get job original file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetJobOriginalFileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get job original file params
func (o *GetJobOriginalFileParams) WithTimeout(timeout time.Duration) *GetJobOriginalFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get job original file params
func (o *GetJobOriginalFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get job original file params
func (o *GetJobOriginalFileParams) WithContext(ctx context.Context) *GetJobOriginalFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get job original file params
func (o *GetJobOriginalFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get job original file params
func (o *GetJobOriginalFileParams) WithHTTPClient(client *http.Client) *GetJobOriginalFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get job original file params
func (o *GetJobOriginalFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUUID adds the uuid to the get job original file params
func (o *GetJobOriginalFileParams) WithUUID(uuid string) *GetJobOriginalFileParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get job original file params
func (o *GetJobOriginalFileParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetJobOriginalFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Uuid
	if err := r.SetPathParam("Uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
