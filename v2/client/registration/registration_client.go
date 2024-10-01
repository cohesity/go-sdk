// Code generated by go-swagger; DO NOT EDIT.

package registration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new registration API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new registration API client with basic auth credentials.
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

// New creates a new registration API client with a bearer token for authentication.
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
Client for registration API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetHeliosRegConfig(params *GetHeliosRegConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHeliosRegConfigOK, error)

	HeliosClaim(params *HeliosClaimParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HeliosClaimNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetHeliosRegConfig lists the helios registration config

Lists the Helios Registration Config.
*/
func (a *Client) GetHeliosRegConfig(params *GetHeliosRegConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHeliosRegConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHeliosRegConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetHeliosRegConfig",
		Method:             "GET",
		PathPattern:        "/helios-registration-config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetHeliosRegConfigReader{formats: a.formats},
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
	success, ok := result.(*GetHeliosRegConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetHeliosRegConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
HeliosClaim registers to helios

Claim to Helios.
*/
func (a *Client) HeliosClaim(params *HeliosClaimParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*HeliosClaimNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHeliosClaimParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HeliosClaim",
		Method:             "POST",
		PathPattern:        "/helios-registration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HeliosClaimReader{formats: a.formats},
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
	success, ok := result.(*HeliosClaimNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*HeliosClaimDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
