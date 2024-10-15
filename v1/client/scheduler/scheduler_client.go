// Code generated by go-swagger; DO NOT EDIT.

package scheduler

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new scheduler API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new scheduler API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new scheduler API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for scheduler API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateSchedulerJob(params *CreateSchedulerJobParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSchedulerJobOK, *CreateSchedulerJobCreated, error)

	DeleteSchedulerJobs(params *DeleteSchedulerJobsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSchedulerJobsNoContent, error)

	GetSchedulerJobs(params *GetSchedulerJobsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSchedulerJobsOK, error)

	UpdateSchedulerJob(params *UpdateSchedulerJobParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateSchedulerJobOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	CreateSchedulerJob creates an email report scheduler job

	**Privileges:** ```SCHEDULER_MODIFY``` <br><br>Returns the created email report scheduler job.

An email report scheduler job generates a report with the specified parameters
and sends that report to the specified email accounts according to the defined
schedule.
This operation may also be used to send a report once (with no schedule),
by setting the EnableRecurring field to false.
*/
func (a *Client) CreateSchedulerJob(params *CreateSchedulerJobParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSchedulerJobOK, *CreateSchedulerJobCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSchedulerJobParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateSchedulerJob",
		Method:             "POST",
		PathPattern:        "/public/scheduler",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateSchedulerJobReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateSchedulerJobOK:
		return value, nil, nil
	case *CreateSchedulerJobCreated:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateSchedulerJobDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteSchedulerJobs deletes one or more email report schedule jobs

**Privileges:** ```SCHEDULER_MODIFY``` <br><br>Specify a list of email report schedule job ids to unschedule and delete.
*/
func (a *Client) DeleteSchedulerJobs(params *DeleteSchedulerJobsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSchedulerJobsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSchedulerJobsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteSchedulerJobs",
		Method:             "DELETE",
		PathPattern:        "/public/scheduler",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSchedulerJobsReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*DeleteSchedulerJobsNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteSchedulerJobsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
	GetSchedulerJobs lists the email report schedule jobs that are scheduled to run

	**Privileges:** ```SCHEDULER_VIEW``` <br><br>Returns all the email report scheduler jobs that are scheduled to run.

An email report scheduler job generates a report with the specified parameters
and sends that report to the specified email accounts according to the defined
schedule.
*/
func (a *Client) GetSchedulerJobs(params *GetSchedulerJobsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSchedulerJobsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSchedulerJobsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSchedulerJobs",
		Method:             "GET",
		PathPattern:        "/public/scheduler",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSchedulerJobsReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*GetSchedulerJobsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSchedulerJobsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateSchedulerJob updates an existing email report schedule job

**Privileges:** ```SCHEDULER_MODIFY``` <br><br>Returns the updated email report scheduler job.
*/
func (a *Client) UpdateSchedulerJob(params *UpdateSchedulerJobParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateSchedulerJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSchedulerJobParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateSchedulerJob",
		Method:             "PUT",
		PathPattern:        "/public/scheduler",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateSchedulerJobReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*UpdateSchedulerJobOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateSchedulerJobDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
