// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new security API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new security API client with basic auth credentials.
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

// New creates a new security API client with a bearer token for authentication.
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
Client for security API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateClientcsr(params *CreateClientcsrParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateClientcsrCreated, error)

	CreateCsr(params *CreateCsrParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateCsrCreated, error)

	DeleteCsr(params *DeleteCsrParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteCsrNoContent, error)

	GetCiphers(params *GetCiphersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCiphersOK, error)

	GetCsrByID(params *GetCsrByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCsrByIDOK, error)

	GetCsrList(params *GetCsrListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCsrListOK, error)

	GetObjectStoreCiphers(params *GetObjectStoreCiphersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetObjectStoreCiphersOK, error)

	GetSecurityConfig(params *GetSecurityConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSecurityConfigOK, error)

	ImportCertificateByClientcsr(params *ImportCertificateByClientcsrParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImportCertificateByClientcsrOK, error)

	ListTrustedCaByID(params *ListTrustedCaByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListTrustedCaByIDOK, error)

	ListTrustedCas(params *ListTrustedCasParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListTrustedCasOK, error)

	ModifyCiphers(params *ModifyCiphersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ModifyCiphersOK, error)

	ModifyObjectStoreCiphers(params *ModifyObjectStoreCiphersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ModifyObjectStoreCiphersOK, error)

	RegisterTrustedCas(params *RegisterTrustedCasParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RegisterTrustedCasCreated, error)

	UnregisterTrustedCa(params *UnregisterTrustedCaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UnregisterTrustedCaNoContent, error)

	UpdateCertificateByCsr(params *UpdateCertificateByCsrParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateCertificateByCsrOK, error)

	UpdateSecurityConfig(params *UpdateSecurityConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateSecurityConfigOK, error)

	ValidateTrustedCaByID(params *ValidateTrustedCaByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ValidateTrustedCaByIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateClientcsr creates certificate signing requests on the cluster

Create two Certificate Signing Request on the cluster with the given details one each for client and server. Each service can have at most one outstanding pair of CSR.
*/
func (a *Client) CreateClientcsr(params *CreateClientcsrParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateClientcsrCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateClientcsrParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateClientcsr",
		Method:             "POST",
		PathPattern:        "/client-csr",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateClientcsrReader{formats: a.formats},
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
	success, ok := result.(*CreateClientcsrCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateClientcsrDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CreateCsr creates a certificate signing request on the cluster

Create a Certificate Signing Request on the cluster with the given details. Each service has at most one outstanding CSR.
*/
func (a *Client) CreateCsr(params *CreateCsrParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateCsrCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateCsrParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateCsr",
		Method:             "POST",
		PathPattern:        "/csr",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateCsrReader{formats: a.formats},
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
	success, ok := result.(*CreateCsrCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateCsrDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteCsr deletes a certificate signing request on the cluster

Delete a Certificate Signing Request on the cluster.
*/
func (a *Client) DeleteCsr(params *DeleteCsrParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteCsrNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCsrParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteCsr",
		Method:             "DELETE",
		PathPattern:        "/csr/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteCsrReader{formats: a.formats},
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
	success, ok := result.(*DeleteCsrNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteCsrDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetCiphers gets the list of ciphers enabled on the cluster

Gets the list of ciphers enabled on the cluster.
*/
func (a *Client) GetCiphers(params *GetCiphersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCiphersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCiphersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetCiphers",
		Method:             "GET",
		PathPattern:        "/security/ciphers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCiphersReader{formats: a.formats},
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
	success, ok := result.(*GetCiphersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCiphersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetCsrByID lists the specified certificate signing request

List the specified Certificate Signing Request.
*/
func (a *Client) GetCsrByID(params *GetCsrByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCsrByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCsrByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetCsrById",
		Method:             "GET",
		PathPattern:        "/csr/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCsrByIDReader{formats: a.formats},
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
	success, ok := result.(*GetCsrByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCsrByIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetCsrList lists certificate signing requests on the cluster

List Certificate Signing Requests on the cluster with service name filtering.
*/
func (a *Client) GetCsrList(params *GetCsrListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCsrListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCsrListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetCsrList",
		Method:             "GET",
		PathPattern:        "/csr",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCsrListReader{formats: a.formats},
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
	success, ok := result.(*GetCsrListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCsrListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetObjectStoreCiphers gets the list of object store ciphers enabled on the cluster

Gets the list of object store ciphers enabled on the cluster.
*/
func (a *Client) GetObjectStoreCiphers(params *GetObjectStoreCiphersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetObjectStoreCiphersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetObjectStoreCiphersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetObjectStoreCiphers",
		Method:             "GET",
		PathPattern:        "/security/object-store-ciphers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetObjectStoreCiphersReader{formats: a.formats},
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
	success, ok := result.(*GetObjectStoreCiphersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetObjectStoreCiphersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetSecurityConfig gets cluster security settings

Get cluster security settings.
*/
func (a *Client) GetSecurityConfig(params *GetSecurityConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSecurityConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSecurityConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSecurityConfig",
		Method:             "GET",
		PathPattern:        "/security-config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSecurityConfigReader{formats: a.formats},
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
	success, ok := result.(*GetSecurityConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSecurityConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ImportCertificateByClientcsr imports the signed certificates on the cluster after the certificate signing requests are created

Import the signed certificates on the cluster after the Certificate Signing Requests are created.
*/
func (a *Client) ImportCertificateByClientcsr(params *ImportCertificateByClientcsrParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImportCertificateByClientcsrOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImportCertificateByClientcsrParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImportCertificateByClientcsr",
		Method:             "POST",
		PathPattern:        "/client-csr/certificate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImportCertificateByClientcsrReader{formats: a.formats},
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
	success, ok := result.(*ImportCertificateByClientcsrOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ImportCertificateByClientcsrDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListTrustedCaByID lists the specified certificate

List the specified Certificate.
*/
func (a *Client) ListTrustedCaByID(params *ListTrustedCaByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListTrustedCaByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListTrustedCaByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListTrustedCaById",
		Method:             "GET",
		PathPattern:        "/trusted-cas/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListTrustedCaByIDReader{formats: a.formats},
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
	success, ok := result.(*ListTrustedCaByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListTrustedCaByIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListTrustedCas lists all certificates with cluster trust store

List all trusted certificates in cluster trust store.
*/
func (a *Client) ListTrustedCas(params *ListTrustedCasParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListTrustedCasOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListTrustedCasParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListTrustedCas",
		Method:             "GET",
		PathPattern:        "/trusted-cas",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListTrustedCasReader{formats: a.formats},
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
	success, ok := result.(*ListTrustedCasOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListTrustedCasDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ModifyCiphers enables disable a list of ciphers on the cluster iris must be restarted for the change to take effect

Enable/Disable a list of ciphers on the cluster.
*/
func (a *Client) ModifyCiphers(params *ModifyCiphersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ModifyCiphersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyCiphersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ModifyCiphers",
		Method:             "POST",
		PathPattern:        "/security/ciphers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ModifyCiphersReader{formats: a.formats},
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
	success, ok := result.(*ModifyCiphersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ModifyCiphersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ModifyObjectStoreCiphers enables disable a list of object store ciphers on the cluster bridge must be restarted for the change to take effect

Enable/Disable a list of object store ciphers on the cluster.
*/
func (a *Client) ModifyObjectStoreCiphers(params *ModifyObjectStoreCiphersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ModifyObjectStoreCiphersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyObjectStoreCiphersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ModifyObjectStoreCiphers",
		Method:             "POST",
		PathPattern:        "/security/object-store-ciphers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ModifyObjectStoreCiphersReader{formats: a.formats},
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
	success, ok := result.(*ModifyObjectStoreCiphersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ModifyObjectStoreCiphersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RegisterTrustedCas registers c a certificate to the cluster trust store

Register CA Certificate to the cluster trust store.
*/
func (a *Client) RegisterTrustedCas(params *RegisterTrustedCasParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RegisterTrustedCasCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegisterTrustedCasParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RegisterTrustedCas",
		Method:             "POST",
		PathPattern:        "/trusted-cas",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RegisterTrustedCasReader{formats: a.formats},
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
	success, ok := result.(*RegisterTrustedCasCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RegisterTrustedCasDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UnregisterTrustedCa unregisters c a certificate from the cluster trust store

Unregister CA Certificate from the cluster trust store.
*/
func (a *Client) UnregisterTrustedCa(params *UnregisterTrustedCaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UnregisterTrustedCaNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUnregisterTrustedCaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UnregisterTrustedCa",
		Method:             "DELETE",
		PathPattern:        "/trusted-cas/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UnregisterTrustedCaReader{formats: a.formats},
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
	success, ok := result.(*UnregisterTrustedCaNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UnregisterTrustedCaDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateCertificateByCsr updates the signed certificate on the cluster after a certificate signing request is created

Update the signed certificate on the cluster after a Certificate Signing Request is created.
*/
func (a *Client) UpdateCertificateByCsr(params *UpdateCertificateByCsrParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateCertificateByCsrOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCertificateByCsrParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateCertificateByCsr",
		Method:             "POST",
		PathPattern:        "/csr/certificate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateCertificateByCsrReader{formats: a.formats},
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
	success, ok := result.(*UpdateCertificateByCsrOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateCertificateByCsrDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateSecurityConfig updates cluster security settings

Update cluster security settings.
*/
func (a *Client) UpdateSecurityConfig(params *UpdateSecurityConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateSecurityConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSecurityConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateSecurityConfig",
		Method:             "PUT",
		PathPattern:        "/security-config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateSecurityConfigReader{formats: a.formats},
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
	success, ok := result.(*UpdateSecurityConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateSecurityConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ValidateTrustedCaByID validates c a certificate

Certificate will be checked for Expiration and Revocation.
*/
func (a *Client) ValidateTrustedCaByID(params *ValidateTrustedCaByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ValidateTrustedCaByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateTrustedCaByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ValidateTrustedCaById",
		Method:             "POST",
		PathPattern:        "/trusted-cas/{id}/validate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateTrustedCaByIDReader{formats: a.formats},
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
	success, ok := result.(*ValidateTrustedCaByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ValidateTrustedCaByIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
