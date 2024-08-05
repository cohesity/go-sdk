# ImportCertificateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertBundle** | Pointer to **string** | SSL identity file in PKCS12 format | [optional] 
**CertRequest** | Pointer to [**ImportCertRequest**](ImportCertRequest.md) |  | [optional] 
**Passphrase** | Pointer to **string** | passphrase for certBundle file | [optional] 

## Methods

### NewImportCertificateRequest

`func NewImportCertificateRequest() *ImportCertificateRequest`

NewImportCertificateRequest instantiates a new ImportCertificateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportCertificateRequestWithDefaults

`func NewImportCertificateRequestWithDefaults() *ImportCertificateRequest`

NewImportCertificateRequestWithDefaults instantiates a new ImportCertificateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertBundle

`func (o *ImportCertificateRequest) GetCertBundle() string`

GetCertBundle returns the CertBundle field if non-nil, zero value otherwise.

### GetCertBundleOk

`func (o *ImportCertificateRequest) GetCertBundleOk() (*string, bool)`

GetCertBundleOk returns a tuple with the CertBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertBundle

`func (o *ImportCertificateRequest) SetCertBundle(v string)`

SetCertBundle sets CertBundle field to given value.

### HasCertBundle

`func (o *ImportCertificateRequest) HasCertBundle() bool`

HasCertBundle returns a boolean if a field has been set.

### GetCertRequest

`func (o *ImportCertificateRequest) GetCertRequest() ImportCertRequest`

GetCertRequest returns the CertRequest field if non-nil, zero value otherwise.

### GetCertRequestOk

`func (o *ImportCertificateRequest) GetCertRequestOk() (*ImportCertRequest, bool)`

GetCertRequestOk returns a tuple with the CertRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertRequest

`func (o *ImportCertificateRequest) SetCertRequest(v ImportCertRequest)`

SetCertRequest sets CertRequest field to given value.

### HasCertRequest

`func (o *ImportCertificateRequest) HasCertRequest() bool`

HasCertRequest returns a boolean if a field has been set.

### GetPassphrase

`func (o *ImportCertificateRequest) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *ImportCertificateRequest) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *ImportCertificateRequest) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *ImportCertificateRequest) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


