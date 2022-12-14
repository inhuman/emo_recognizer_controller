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

// NewCreateJobParams creates a new CreateJobParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateJobParams() *CreateJobParams {
	return &CreateJobParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateJobParamsWithTimeout creates a new CreateJobParams object
// with the ability to set a timeout on a request.
func NewCreateJobParamsWithTimeout(timeout time.Duration) *CreateJobParams {
	return &CreateJobParams{
		timeout: timeout,
	}
}

// NewCreateJobParamsWithContext creates a new CreateJobParams object
// with the ability to set a context for a request.
func NewCreateJobParamsWithContext(ctx context.Context) *CreateJobParams {
	return &CreateJobParams{
		Context: ctx,
	}
}

// NewCreateJobParamsWithHTTPClient creates a new CreateJobParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateJobParamsWithHTTPClient(client *http.Client) *CreateJobParams {
	return &CreateJobParams{
		HTTPClient: client,
	}
}

/*
CreateJobParams contains all the parameters to send to the API endpoint

	for the create job operation.

	Typically these are written to a http.Request.
*/
type CreateJobParams struct {

	/* File.

	   Звуковой файл в формате wav
	*/
	File runtime.NamedReadCloser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateJobParams) WithDefaults() *CreateJobParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateJobParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create job params
func (o *CreateJobParams) WithTimeout(timeout time.Duration) *CreateJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create job params
func (o *CreateJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create job params
func (o *CreateJobParams) WithContext(ctx context.Context) *CreateJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create job params
func (o *CreateJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create job params
func (o *CreateJobParams) WithHTTPClient(client *http.Client) *CreateJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create job params
func (o *CreateJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFile adds the file to the create job params
func (o *CreateJobParams) WithFile(file runtime.NamedReadCloser) *CreateJobParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the create job params
func (o *CreateJobParams) SetFile(file runtime.NamedReadCloser) {
	o.File = file
}

// WriteToRequest writes these params to a swagger request
func (o *CreateJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.File != nil {

		if o.File != nil {
			// form file param file
			if err := r.SetFileParam("file", o.File); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
