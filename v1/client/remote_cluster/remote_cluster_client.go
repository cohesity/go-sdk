// Code generated by go-swagger; DO NOT EDIT.

package remote_cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new remote cluster API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new remote cluster API client with basic auth credentials.
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

// New creates a new remote cluster API client with a bearer token for authentication.
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
Client for remote cluster API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateRemoteCluster(params *CreateRemoteClusterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateRemoteClusterCreated, error)

	DeleteRemoteCluster(params *DeleteRemoteClusterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteRemoteClusterNoContent, error)

	GetRemoteClusterByID(params *GetRemoteClusterByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRemoteClusterByIDOK, error)

	GetRemoteClusters(params *GetRemoteClustersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRemoteClustersOK, error)

	GetReplicationEncryptionKey(params *GetReplicationEncryptionKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetReplicationEncryptionKeyOK, error)

	UpdateRemoteCluster(params *UpdateRemoteClusterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateRemoteClusterOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	CreateRemoteCluster registers a remote cluster on this local cluster for replication

	For a Protection Job to replicate Snapshots from one Cluster

to another Cluster, the Clusters must be paired together by
registering each Cluster on the other Cluster.
For example, Cluster A must be registered on Cluster B
and Cluster B must be registered on Cluster A.
*/
func (a *Client) CreateRemoteCluster(params *CreateRemoteClusterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateRemoteClusterCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRemoteClusterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateRemoteCluster",
		Method:             "POST",
		PathPattern:        "/public/remoteClusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateRemoteClusterReader{formats: a.formats},
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
	success, ok := result.(*CreateRemoteClusterCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateRemoteClusterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
	DeleteRemoteCluster deletes a remote cluster registration connection

	Delete the specified remote Cluster registration connection

on this Cluster.
*/
func (a *Client) DeleteRemoteCluster(params *DeleteRemoteClusterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteRemoteClusterNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRemoteClusterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteRemoteCluster",
		Method:             "DELETE",
		PathPattern:        "/public/remoteClusters/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteRemoteClusterReader{formats: a.formats},
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
	success, ok := result.(*DeleteRemoteClusterNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteRemoteClusterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
	GetRemoteClusterByID lists details about a single remote cluster registered on this local cluster

	Returns the details about the remote Cluster with the specified Cluster id

that is registered on this local Cluster.
*/
func (a *Client) GetRemoteClusterByID(params *GetRemoteClusterByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRemoteClusterByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRemoteClusterByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetRemoteClusterById",
		Method:             "GET",
		PathPattern:        "/public/remoteClusters/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRemoteClusterByIDReader{formats: a.formats},
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
	success, ok := result.(*GetRemoteClusterByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetRemoteClusterByIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
	GetRemoteClusters lists the remote cohesity clusters that are registered on this local cohesity cluster that match the filter criteria specified using parameters

	Cohesity Clusters involved in replication, must be registered to each other.

For example, if Cluster A is replicating Snapshots to Cluster B, Cluster
B must be registered on Cluster A and Cluster B must be registered
on Cluster A.
*/
func (a *Client) GetRemoteClusters(params *GetRemoteClustersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRemoteClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRemoteClustersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetRemoteClusters",
		Method:             "GET",
		PathPattern:        "/public/remoteClusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRemoteClustersReader{formats: a.formats},
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
	success, ok := result.(*GetRemoteClustersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetRemoteClustersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
	GetReplicationEncryptionKey gets the encryption key on this cluster

	Get the encryption key that is used for encrypting replication data

between this Cluster and a remote Cluster.
*/
func (a *Client) GetReplicationEncryptionKey(params *GetReplicationEncryptionKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetReplicationEncryptionKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetReplicationEncryptionKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetReplicationEncryptionKey",
		Method:             "GET",
		PathPattern:        "/public/replicationEncryptionKey",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetReplicationEncryptionKeyReader{formats: a.formats},
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
	success, ok := result.(*GetReplicationEncryptionKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetReplicationEncryptionKeyDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
	UpdateRemoteCluster updates the connection settings of the registered remote cluster

	Update the connection settings of the specified remote Cluster that is

registered on this Cluster.
*/
func (a *Client) UpdateRemoteCluster(params *UpdateRemoteClusterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateRemoteClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRemoteClusterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateRemoteCluster",
		Method:             "PUT",
		PathPattern:        "/public/remoteClusters/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateRemoteClusterReader{formats: a.formats},
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
	success, ok := result.(*UpdateRemoteClusterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateRemoteClusterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
