// Code generated by go-swagger; DO NOT EDIT.

package backup_jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new backup jobs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new backup jobs API client with basic auth credentials.
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

// New creates a new backup jobs API client with a bearer token for authentication.
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
Client for backup jobs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ActivateBackupJob(params *ActivateBackupJobParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ActivateBackupJobOK, error)

	CreateBackupJob(params *CreateBackupJobParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateBackupJobCreated, error)

	DeactivateBackupJob(params *DeactivateBackupJobParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeactivateBackupJobOK, error)

	DeleteBackupJob(params *DeleteBackupJobParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteBackupJobOK, error)

	DeleteBackupJobRuns(params *DeleteBackupJobRunsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteBackupJobRunsOK, error)

	GetBackupJobAudit(params *GetBackupJobAuditParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBackupJobAuditOK, error)

	GetBackupJobByID(params *GetBackupJobByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBackupJobByIDOK, error)

	GetBackupJobHistory(params *GetBackupJobHistoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBackupJobHistoryOK, error)

	GetBackupJobRuns(params *GetBackupJobRunsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBackupJobRunsOK, error)

	GetBackupJobs(params *GetBackupJobsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBackupJobsOK, error)

	GetBackupJobsSummary(params *GetBackupJobsSummaryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBackupJobsSummaryOK, error)

	UpdateBackupJob(params *UpdateBackupJobParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateBackupJobOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ActivateBackupJob activates a backup job

This is used for failback.
*/
func (a *Client) ActivateBackupJob(params *ActivateBackupJobParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ActivateBackupJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewActivateBackupJobParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ActivateBackupJob",
		Method:             "POST",
		PathPattern:        "/activateBackupJob/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ActivateBackupJobReader{formats: a.formats},
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
	success, ok := result.(*ActivateBackupJobOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ActivateBackupJobDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CreateBackupJob creates a backup job

Returns the created backup job.
*/
func (a *Client) CreateBackupJob(params *CreateBackupJobParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateBackupJobCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateBackupJobParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateBackupJob",
		Method:             "POST",
		PathPattern:        "/backupjobs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateBackupJobReader{formats: a.formats},
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
	success, ok := result.(*CreateBackupJobCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateBackupJobDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeactivateBackupJob deactivates a backup job

Deactivate a backup job. This is used for failover on a remote cluster.
*/
func (a *Client) DeactivateBackupJob(params *DeactivateBackupJobParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeactivateBackupJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeactivateBackupJobParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeactivateBackupJob",
		Method:             "POST",
		PathPattern:        "/deactivateBackupJob/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeactivateBackupJobReader{formats: a.formats},
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
	success, ok := result.(*DeactivateBackupJobOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeactivateBackupJobDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteBackupJob deletes a backup job

Return success if the backup job is deleted.
*/
func (a *Client) DeleteBackupJob(params *DeleteBackupJobParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteBackupJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBackupJobParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteBackupJob",
		Method:             "DELETE",
		PathPattern:        "/backupjobs/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteBackupJobReader{formats: a.formats},
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
	success, ok := result.(*DeleteBackupJobOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteBackupJobDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteBackupJobRuns deletes backup job runs

Returns success if the backup job runs are deleted.
*/
func (a *Client) DeleteBackupJobRuns(params *DeleteBackupJobRunsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteBackupJobRunsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBackupJobRunsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteBackupJobRuns",
		Method:             "DELETE",
		PathPattern:        "/backupjobruns",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteBackupJobRunsReader{formats: a.formats},
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
	success, ok := result.(*DeleteBackupJobRunsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteBackupJobRunsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetBackupJobAudit lists a backup job audit

Returns the audit of specific backup job history.
*/
func (a *Client) GetBackupJobAudit(params *GetBackupJobAuditParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBackupJobAuditOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBackupJobAuditParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetBackupJobAudit",
		Method:             "GET",
		PathPattern:        "/backupjobaudit/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBackupJobAuditReader{formats: a.formats},
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
	success, ok := result.(*GetBackupJobAuditOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetBackupJobAuditDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetBackupJobByID lists details about single backup job

Return the Backup Job corresponding to the specified id.
*/
func (a *Client) GetBackupJobByID(params *GetBackupJobByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBackupJobByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBackupJobByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetBackupJobById",
		Method:             "GET",
		PathPattern:        "/backupjobs/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBackupJobByIDReader{formats: a.formats},
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
	success, ok := result.(*GetBackupJobByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetBackupJobByIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetBackupJobHistory lists a backup job history

Returns the history of specific backup job history.
*/
func (a *Client) GetBackupJobHistory(params *GetBackupJobHistoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBackupJobHistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBackupJobHistoryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetBackupJobHistory",
		Method:             "GET",
		PathPattern:        "/backupjobhistory/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBackupJobHistoryReader{formats: a.formats},
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
	success, ok := result.(*GetBackupJobHistoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetBackupJobHistoryDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetBackupJobRuns lists the backup job runs

Returns the list of backup job returns.
*/
func (a *Client) GetBackupJobRuns(params *GetBackupJobRunsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBackupJobRunsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBackupJobRunsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetBackupJobRuns",
		Method:             "GET",
		PathPattern:        "/backupjobruns",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBackupJobRunsReader{formats: a.formats},
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
	success, ok := result.(*GetBackupJobRunsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetBackupJobRunsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
	GetBackupJobs lists backup jobs filtered by the specifed parameters

	If no parameters are specified, all Backup Jobs currently

on the Cohesity Cluster are returned.
Specifying parameters filters the results that are returned.
*/
func (a *Client) GetBackupJobs(params *GetBackupJobsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBackupJobsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBackupJobsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetBackupJobs",
		Method:             "GET",
		PathPattern:        "/backupjobs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBackupJobsReader{formats: a.formats},
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
	success, ok := result.(*GetBackupJobsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetBackupJobsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetBackupJobsSummary lists backup jobs summary according to specified summary

Returns the backup jobs summary filtered by the mentioned criteria.
*/
func (a *Client) GetBackupJobsSummary(params *GetBackupJobsSummaryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBackupJobsSummaryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBackupJobsSummaryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetBackupJobsSummary",
		Method:             "GET",
		PathPattern:        "/backupjobssummary",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBackupJobsSummaryReader{formats: a.formats},
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
	success, ok := result.(*GetBackupJobsSummaryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetBackupJobsSummaryDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateBackupJob updates a backup job

Returns the updated backup job.
*/
func (a *Client) UpdateBackupJob(params *UpdateBackupJobParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateBackupJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateBackupJobParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateBackupJob",
		Method:             "PUT",
		PathPattern:        "/backupjobs/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateBackupJobReader{formats: a.formats},
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
	success, ok := result.(*UpdateBackupJobOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateBackupJobDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
