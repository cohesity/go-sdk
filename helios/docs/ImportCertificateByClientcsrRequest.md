# ImportCertificateByClientcsrRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateServer** | **NullableString** | Specifies the server certificate to be imported. | 
**CertificateClient** | **NullableString** | Specifies the client certificate to be imported. | 

## Methods

### NewImportCertificateByClientcsrRequest

`func NewImportCertificateByClientcsrRequest(certificateServer NullableString, certificateClient NullableString, ) *ImportCertificateByClientcsrRequest`

NewImportCertificateByClientcsrRequest instantiates a new ImportCertificateByClientcsrRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportCertificateByClientcsrRequestWithDefaults

`func NewImportCertificateByClientcsrRequestWithDefaults() *ImportCertificateByClientcsrRequest`

NewImportCertificateByClientcsrRequestWithDefaults instantiates a new ImportCertificateByClientcsrRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateServer

`func (o *ImportCertificateByClientcsrRequest) GetCertificateServer() string`

GetCertificateServer returns the CertificateServer field if non-nil, zero value otherwise.

### GetCertificateServerOk

`func (o *ImportCertificateByClientcsrRequest) GetCertificateServerOk() (*string, bool)`

GetCertificateServerOk returns a tuple with the CertificateServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateServer

`func (o *ImportCertificateByClientcsrRequest) SetCertificateServer(v string)`

SetCertificateServer sets CertificateServer field to given value.


### SetCertificateServerNil

`func (o *ImportCertificateByClientcsrRequest) SetCertificateServerNil(b bool)`

 SetCertificateServerNil sets the value for CertificateServer to be an explicit nil

### UnsetCertificateServer
`func (o *ImportCertificateByClientcsrRequest) UnsetCertificateServer()`

UnsetCertificateServer ensures that no value is present for CertificateServer, not even an explicit nil
### GetCertificateClient

`func (o *ImportCertificateByClientcsrRequest) GetCertificateClient() string`

GetCertificateClient returns the CertificateClient field if non-nil, zero value otherwise.

### GetCertificateClientOk

`func (o *ImportCertificateByClientcsrRequest) GetCertificateClientOk() (*string, bool)`

GetCertificateClientOk returns a tuple with the CertificateClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateClient

`func (o *ImportCertificateByClientcsrRequest) SetCertificateClient(v string)`

SetCertificateClient sets CertificateClient field to given value.


### SetCertificateClientNil

`func (o *ImportCertificateByClientcsrRequest) SetCertificateClientNil(b bool)`

 SetCertificateClientNil sets the value for CertificateClient to be an explicit nil

### UnsetCertificateClient
`func (o *ImportCertificateByClientcsrRequest) UnsetCertificateClient()`

UnsetCertificateClient ensures that no value is present for CertificateClient, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


