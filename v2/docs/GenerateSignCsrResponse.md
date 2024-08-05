# GenerateSignCsrResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaCert** | Pointer to **[]string** | Specifies ca cert in pem format | [optional] 
**Certificate** | Pointer to **NullableString** | Specifies certificate in pem format. | [optional] 

## Methods

### NewGenerateSignCsrResponse

`func NewGenerateSignCsrResponse() *GenerateSignCsrResponse`

NewGenerateSignCsrResponse instantiates a new GenerateSignCsrResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateSignCsrResponseWithDefaults

`func NewGenerateSignCsrResponseWithDefaults() *GenerateSignCsrResponse`

NewGenerateSignCsrResponseWithDefaults instantiates a new GenerateSignCsrResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaCert

`func (o *GenerateSignCsrResponse) GetCaCert() []string`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *GenerateSignCsrResponse) GetCaCertOk() (*[]string, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *GenerateSignCsrResponse) SetCaCert(v []string)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *GenerateSignCsrResponse) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### GetCertificate

`func (o *GenerateSignCsrResponse) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *GenerateSignCsrResponse) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *GenerateSignCsrResponse) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *GenerateSignCsrResponse) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *GenerateSignCsrResponse) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *GenerateSignCsrResponse) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


