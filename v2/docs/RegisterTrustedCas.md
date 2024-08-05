# RegisterTrustedCas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificates** | [**[]TrustedCaRequest**](TrustedCaRequest.md) | Specifies the certificates to be imported. | 
**OnlyValidate** | Pointer to **NullableBool** | Specifies if the certificates are only to be validated. | [optional] 

## Methods

### NewRegisterTrustedCas

`func NewRegisterTrustedCas(certificates []TrustedCaRequest, ) *RegisterTrustedCas`

NewRegisterTrustedCas instantiates a new RegisterTrustedCas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterTrustedCasWithDefaults

`func NewRegisterTrustedCasWithDefaults() *RegisterTrustedCas`

NewRegisterTrustedCasWithDefaults instantiates a new RegisterTrustedCas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificates

`func (o *RegisterTrustedCas) GetCertificates() []TrustedCaRequest`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *RegisterTrustedCas) GetCertificatesOk() (*[]TrustedCaRequest, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *RegisterTrustedCas) SetCertificates(v []TrustedCaRequest)`

SetCertificates sets Certificates field to given value.


### SetCertificatesNil

`func (o *RegisterTrustedCas) SetCertificatesNil(b bool)`

 SetCertificatesNil sets the value for Certificates to be an explicit nil

### UnsetCertificates
`func (o *RegisterTrustedCas) UnsetCertificates()`

UnsetCertificates ensures that no value is present for Certificates, not even an explicit nil
### GetOnlyValidate

`func (o *RegisterTrustedCas) GetOnlyValidate() bool`

GetOnlyValidate returns the OnlyValidate field if non-nil, zero value otherwise.

### GetOnlyValidateOk

`func (o *RegisterTrustedCas) GetOnlyValidateOk() (*bool, bool)`

GetOnlyValidateOk returns a tuple with the OnlyValidate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyValidate

`func (o *RegisterTrustedCas) SetOnlyValidate(v bool)`

SetOnlyValidate sets OnlyValidate field to given value.

### HasOnlyValidate

`func (o *RegisterTrustedCas) HasOnlyValidate() bool`

HasOnlyValidate returns a boolean if a field has been set.

### SetOnlyValidateNil

`func (o *RegisterTrustedCas) SetOnlyValidateNil(b bool)`

 SetOnlyValidateNil sets the value for OnlyValidate to be an explicit nil

### UnsetOnlyValidate
`func (o *RegisterTrustedCas) UnsetOnlyValidate()`

UnsetOnlyValidate ensures that no value is present for OnlyValidate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


