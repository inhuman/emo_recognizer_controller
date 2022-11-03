// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new job API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for job API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateJob(params *CreateJobParams, opts ...ClientOption) (*CreateJobOK, error)

	GetJob(params *GetJobParams, opts ...ClientOption) (*GetJobOK, error)

	GetJobCleanFile(params *GetJobCleanFileParams, opts ...ClientOption) (*GetJobCleanFileOK, error)

	GetJobOriginalFile(params *GetJobOriginalFileParams, opts ...ClientOption) (*GetJobOriginalFileOK, error)

	GetJobs(params *GetJobsParams, opts ...ClientOption) (*GetJobsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateJob Эндпоинт для загрузки звукового файла (.wav)
*/
func (a *Client) CreateJob(params *CreateJobParams, opts ...ClientOption) (*CreateJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateJobParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createJob",
		Method:             "POST",
		PathPattern:        "/api/v1/jobs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateJobReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateJobOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createJob: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetJob Эндпоинт для получения задачи по UUID
*/
func (a *Client) GetJob(params *GetJobParams, opts ...ClientOption) (*GetJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetJobParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getJob",
		Method:             "GET",
		PathPattern:        "/api/v1/jobs/{Uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetJobReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetJobOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getJob: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetJobCleanFile Эндпоинт для получение очищенного файла задачи по UUID
*/
func (a *Client) GetJobCleanFile(params *GetJobCleanFileParams, opts ...ClientOption) (*GetJobCleanFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetJobCleanFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getJobCleanFile",
		Method:             "GET",
		PathPattern:        "/api/v1/jobs/{Uuid}/file/clean",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetJobCleanFileReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetJobCleanFileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getJobCleanFile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetJobOriginalFile Эндпоинт для получения файла задачи по UUID
*/
func (a *Client) GetJobOriginalFile(params *GetJobOriginalFileParams, opts ...ClientOption) (*GetJobOriginalFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetJobOriginalFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getJobOriginalFile",
		Method:             "GET",
		PathPattern:        "/api/v1/jobs/{Uuid}/file/original",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetJobOriginalFileReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetJobOriginalFileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getJobOriginalFile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetJobs Эндпоинт для получения списка задач на обработку
*/
func (a *Client) GetJobs(params *GetJobsParams, opts ...ClientOption) (*GetJobsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetJobsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getJobs",
		Method:             "GET",
		PathPattern:        "/api/v1/jobs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetJobsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetJobsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getJobs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
