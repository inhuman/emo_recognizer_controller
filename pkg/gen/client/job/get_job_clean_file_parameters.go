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

// NewGetJobCleanFileParams creates a new GetJobCleanFileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetJobCleanFileParams() *GetJobCleanFileParams {
	return &GetJobCleanFileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetJobCleanFileParamsWithTimeout creates a new GetJobCleanFileParams object
// with the ability to set a timeout on a request.
func NewGetJobCleanFileParamsWithTimeout(timeout time.Duration) *GetJobCleanFileParams {
	return &GetJobCleanFileParams{
		timeout: timeout,
	}
}

// NewGetJobCleanFileParamsWithContext creates a new GetJobCleanFileParams object
// with the ability to set a context for a request.
func NewGetJobCleanFileParamsWithContext(ctx context.Context) *GetJobCleanFileParams {
	return &GetJobCleanFileParams{
		Context: ctx,
	}
}

// NewGetJobCleanFileParamsWithHTTPClient creates a new GetJobCleanFileParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetJobCleanFileParamsWithHTTPClient(client *http.Client) *GetJobCleanFileParams {
	return &GetJobCleanFileParams{
		HTTPClient: client,
	}
}

/*
GetJobCleanFileParams contains all the parameters to send to the API endpoint

	for the get job clean file operation.

	Typically these are written to a http.Request.
*/
type GetJobCleanFileParams struct {

	/* UUID.

	   Uuid задания
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get job clean file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetJobCleanFileParams) WithDefaults() *GetJobCleanFileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get job clean file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetJobCleanFileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get job clean file params
func (o *GetJobCleanFileParams) WithTimeout(timeout time.Duration) *GetJobCleanFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get job clean file params
func (o *GetJobCleanFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get job clean file params
func (o *GetJobCleanFileParams) WithContext(ctx context.Context) *GetJobCleanFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get job clean file params
func (o *GetJobCleanFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get job clean file params
func (o *GetJobCleanFileParams) WithHTTPClient(client *http.Client) *GetJobCleanFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get job clean file params
func (o *GetJobCleanFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUUID adds the uuid to the get job clean file params
func (o *GetJobCleanFileParams) WithUUID(uuid string) *GetJobCleanFileParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get job clean file params
func (o *GetJobCleanFileParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetJobCleanFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
