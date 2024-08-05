# ImportCertificateByClientcsrResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateServer** | Pointer to **NullableString** | Specifies the server certificate. | [optional] 
**FileServerCert** | Pointer to **NullableString** | Specifies the path to the file to be uploaded to server. This file has the server cert, id and encrypted private key | [optional] 
**PrivateKey** | Pointer to **NullableString** | Specifies the private key of agent. | [optional] 

## Methods

### NewImportCertificateByClientcsrResponseBody

`func NewImportCertificateByClientcsrResponseBody() *ImportCertificateByClientcsrResponseBody`

NewImportCertificateByClientcsrResponseBody instantiates a new ImportCertificateByClientcsrResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportCertificateByClientcsrResponseBodyWithDefaults

`func NewImportCertificateByClientcsrResponseBodyWithDefaults() *ImportCertificateByClientcsrResponseBody`

NewImportCertificateByClientcsrResponseBodyWithDefaults instantiates a new ImportCertificateByClientcsrResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateServer

`func (o *ImportCertificateByClientcsrResponseBody) GetCertificateServer() string`

GetCertificateServer returns the CertificateServer field if non-nil, zero value otherwise.

### GetCertificateServerOk

`func (o *ImportCertificateByClientcsrResponseBody) GetCertificateServerOk() (*string, bool)`

GetCertificateServerOk returns a tuple with the CertificateServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateServer

`func (o *ImportCertificateByClientcsrResponseBody) SetCertificateServer(v string)`

SetCertificateServer sets CertificateServer field to given value.

### HasCertificateServer

`func (o *ImportCertificateByClientcsrResponseBody) HasCertificateServer() bool`

HasCertificateServer returns a boolean if a field has been set.

### SetCertificateServerNil

`func (o *ImportCertificateByClientcsrResponseBody) SetCertificateServerNil(b bool)`

 SetCertificateServerNil sets the value for CertificateServer to be an explicit nil

### UnsetCertificateServer
`func (o *ImportCertificateByClientcsrResponseBody) UnsetCertificateServer()`

UnsetCertificateServer ensures that no value is present for CertificateServer, not even an explicit nil
### GetFileServerCert

`func (o *ImportCertificateByClientcsrResponseBody) GetFileServerCert() string`

GetFileServerCert returns the FileServerCert field if non-nil, zero value otherwise.

### GetFileServerCertOk

`func (o *ImportCertificateByClientcsrResponseBody) GetFileServerCertOk() (*string, bool)`

GetFileServerCertOk returns a tuple with the FileServerCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileServerCert

`func (o *ImportCertificateByClientcsrResponseBody) SetFileServerCert(v string)`

SetFileServerCert sets FileServerCert field to given value.

### HasFileServerCert

`func (o *ImportCertificateByClientcsrResponseBody) HasFileServerCert() bool`

HasFileServerCert returns a boolean if a field has been set.

### SetFileServerCertNil

`func (o *ImportCertificateByClientcsrResponseBody) SetFileServerCertNil(b bool)`

 SetFileServerCertNil sets the value for FileServerCert to be an explicit nil

### UnsetFileServerCert
`func (o *ImportCertificateByClientcsrResponseBody) UnsetFileServerCert()`

UnsetFileServerCert ensures that no value is present for FileServerCert, not even an explicit nil
### GetPrivateKey

`func (o *ImportCertificateByClientcsrResponseBody) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *ImportCertificateByClientcsrResponseBody) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *ImportCertificateByClientcsrResponseBody) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *ImportCertificateByClientcsrResponseBody) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### SetPrivateKeyNil

`func (o *ImportCertificateByClientcsrResponseBody) SetPrivateKeyNil(b bool)`

 SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil

### UnsetPrivateKey
`func (o *ImportCertificateByClientcsrResponseBody) UnsetPrivateKey()`

UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


