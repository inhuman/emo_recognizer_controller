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
	"github.com/go-openapi/swag"
)

// NewGetJobsParams creates a new GetJobsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetJobsParams() *GetJobsParams {
	return &GetJobsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetJobsParamsWithTimeout creates a new GetJobsParams object
// with the ability to set a timeout on a request.
func NewGetJobsParamsWithTimeout(timeout time.Duration) *GetJobsParams {
	return &GetJobsParams{
		timeout: timeout,
	}
}

// NewGetJobsParamsWithContext creates a new GetJobsParams object
// with the ability to set a context for a request.
func NewGetJobsParamsWithContext(ctx context.Context) *GetJobsParams {
	return &GetJobsParams{
		Context: ctx,
	}
}

// NewGetJobsParamsWithHTTPClient creates a new GetJobsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetJobsParamsWithHTTPClient(client *http.Client) *GetJobsParams {
	return &GetJobsParams{
		HTTPClient: client,
	}
}

/*
GetJobsParams contains all the parameters to send to the API endpoint

	for the get jobs operation.

	Typically these are written to a http.Request.
*/
type GetJobsParams struct {

	/* Limit.

	   Limit

	   Format: int64
	*/
	Limit *int64

	/* Offset.

	   Offset

	   Format: int64
	*/
	Offset *int64

	/* Status.

	   Status
	*/
	Status *string

	/* Strategy.

	   Strategy
	*/
	Strategy *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get jobs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetJobsParams) WithDefaults() *GetJobsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get jobs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetJobsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get jobs params
func (o *GetJobsParams) WithTimeout(timeout time.Duration) *GetJobsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get jobs params
func (o *GetJobsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get jobs params
func (o *GetJobsParams) WithContext(ctx context.Context) *GetJobsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get jobs params
func (o *GetJobsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get jobs params
func (o *GetJobsParams) WithHTTPClient(client *http.Client) *GetJobsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get jobs params
func (o *GetJobsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get jobs params
func (o *GetJobsParams) WithLimit(limit *int64) *GetJobsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get jobs params
func (o *GetJobsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get jobs params
func (o *GetJobsParams) WithOffset(offset *int64) *GetJobsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get jobs params
func (o *GetJobsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithStatus adds the status to the get jobs params
func (o *GetJobsParams) WithStatus(status *string) *GetJobsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get jobs params
func (o *GetJobsParams) SetStatus(status *string) {
	o.Status = status
}

// WithStrategy adds the strategy to the get jobs params
func (o *GetJobsParams) WithStrategy(strategy *string) *GetJobsParams {
	o.SetStrategy(strategy)
	return o
}

// SetStrategy adds the strategy to the get jobs params
func (o *GetJobsParams) SetStrategy(strategy *string) {
	o.Strategy = strategy
}

// WriteToRequest writes these params to a swagger request
func (o *GetJobsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Status != nil {

		// query param status
		var qrStatus string

		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {

			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}
	}

	if o.Strategy != nil {

		// query param strategy
		var qrStrategy string

		if o.Strategy != nil {
			qrStrategy = *o.Strategy
		}
		qStrategy := qrStrategy
		if qStrategy != "" {

			if err := r.SetQueryParam("strategy", qStrategy); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
