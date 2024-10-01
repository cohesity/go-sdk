// Code generated by go-swagger; DO NOT EDIT.

package vlan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new vlan API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new vlan API client with basic auth credentials.
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

// New creates a new vlan API client with a bearer token for authentication.
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
Client for vlan API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateBifrostVlan(params *CreateBifrostVlanParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateBifrostVlanOK, error)

	CreateVlan(params *CreateVlanParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateVlanOK, error)

	GetBifrostTenantVlans(params *GetBifrostTenantVlansParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBifrostTenantVlansOK, error)

	GetVlanByID(params *GetVlanByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVlanByIDOK, error)

	GetVlans(params *GetVlansParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVlansOK, error)

	RemoveVlan(params *RemoveVlanParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RemoveVlanNoContent, error)

	UpdateVlan(params *UpdateVlanParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateVlanOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateBifrostVlan creates a bifrost v l a n of the cohesity cluster

Returns the created Bifrost VLAN on the Cohesity Cluster.
*/
func (a *Client) CreateBifrostVlan(params *CreateBifrostVlanParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateBifrostVlanOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateBifrostVlanParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateBifrostVlan",
		Method:             "POST",
		PathPattern:        "/public/bifrost/vlans",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateBifrostVlanReader{formats: a.formats},
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
	success, ok := result.(*CreateBifrostVlanOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateBifrostVlanDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CreateVlan creates a v l a n of the cohesity cluster

Returns the created VLAN on the Cohesity Cluster.
*/
func (a *Client) CreateVlan(params *CreateVlanParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateVlanOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateVlanParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateVlan",
		Method:             "POST",
		PathPattern:        "/public/vlans",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateVlanReader{formats: a.formats},
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
	success, ok := result.(*CreateVlanOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateVlanDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetBifrostTenantVlans lists the bifrost tenant v l a ns for the cohesity cluster

Returns the Bifrost Tenant VLANs for the Cohesity Cluster.
*/
func (a *Client) GetBifrostTenantVlans(params *GetBifrostTenantVlansParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBifrostTenantVlansOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBifrostTenantVlansParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetBifrostTenantVlans",
		Method:             "GET",
		PathPattern:        "/public/bifrost/vlans",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBifrostTenantVlansReader{formats: a.formats},
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
	success, ok := result.(*GetBifrostTenantVlansOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetBifrostTenantVlansDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
	GetVlanByID lists the details about a single v l a n

	Returns the VLAN corresponding to the specified VLAN ID or a specified

vlan interface group name.
Example: /public/vlans/intf_group1.20
*/
func (a *Client) GetVlanByID(params *GetVlanByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVlanByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVlanByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetVlanById",
		Method:             "GET",
		PathPattern:        "/public/vlans/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVlanByIDReader{formats: a.formats},
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
	success, ok := result.(*GetVlanByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetVlanByIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetVlans lists the v l a ns for the cohesity cluster

Returns the VLANs for the Cohesity Cluster.
*/
func (a *Client) GetVlans(params *GetVlansParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVlansOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVlansParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetVlans",
		Method:             "GET",
		PathPattern:        "/public/vlans",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVlansReader{formats: a.formats},
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
	success, ok := result.(*GetVlansOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetVlansDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RemoveVlan removes the specified v l a n from the cohesity cluster

Returns the delete status upon completion.
*/
func (a *Client) RemoveVlan(params *RemoveVlanParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RemoveVlanNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveVlanParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RemoveVlan",
		Method:             "DELETE",
		PathPattern:        "/public/vlans/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RemoveVlanReader{formats: a.formats},
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
	success, ok := result.(*RemoveVlanNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RemoveVlanDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateVlan creates or updates a v l a n of the cohesity cluster

Returns the created or updated VLAN on the Cohesity Cluster.
*/
func (a *Client) UpdateVlan(params *UpdateVlanParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateVlanOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateVlanParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateVlan",
		Method:             "PUT",
		PathPattern:        "/public/vlans/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateVlanReader{formats: a.formats},
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
	success, ok := result.(*UpdateVlanOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateVlanDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
