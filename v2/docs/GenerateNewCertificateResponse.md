# GenerateNewCertificateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaCert** | Pointer to **[]string** | Specifies ca cert in pem format | [optional] 
**Certificate** | Pointer to **NullableString** | Specifies certificate in pem format. | [optional] 
**PrivateKey** | Pointer to **NullableString** | Actual private key of the certificate. | [optional] 

## Methods

### NewGenerateNewCertificateResponse

`func NewGenerateNewCertificateResponse() *GenerateNewCertificateResponse`

NewGenerateNewCertificateResponse instantiates a new GenerateNewCertificateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateNewCertificateResponseWithDefaults

`func NewGenerateNewCertificateResponseWithDefaults() *GenerateNewCertificateResponse`

NewGenerateNewCertificateResponseWithDefaults instantiates a new GenerateNewCertificateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaCert

`func (o *GenerateNewCertificateResponse) GetCaCert() []string`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *GenerateNewCertificateResponse) GetCaCertOk() (*[]string, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *GenerateNewCertificateResponse) SetCaCert(v []string)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *GenerateNewCertificateResponse) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### GetCertificate

`func (o *GenerateNewCertificateResponse) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *GenerateNewCertificateResponse) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *GenerateNewCertificateResponse) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *GenerateNewCertificateResponse) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *GenerateNewCertificateResponse) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *GenerateNewCertificateResponse) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetPrivateKey

`func (o *GenerateNewCertificateResponse) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *GenerateNewCertificateResponse) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *GenerateNewCertificateResponse) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *GenerateNewCertificateResponse) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### SetPrivateKeyNil

`func (o *GenerateNewCertificateResponse) SetPrivateKeyNil(b bool)`

 SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil

### UnsetPrivateKey
`func (o *GenerateNewCertificateResponse) UnsetPrivateKey()`

UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


