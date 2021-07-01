# SslVerification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaCertificate** | Pointer to **NullableString** | Contains the contents of CA cert/cert chain. | [optional] 
**IsEnabled** | Pointer to **NullableBool** | Whether SSL verification should be performed. | [optional] 

## Methods

### NewSslVerification

`func NewSslVerification() *SslVerification`

NewSslVerification instantiates a new SslVerification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslVerificationWithDefaults

`func NewSslVerificationWithDefaults() *SslVerification`

NewSslVerificationWithDefaults instantiates a new SslVerification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaCertificate

`func (o *SslVerification) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *SslVerification) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *SslVerification) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *SslVerification) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.

### SetCaCertificateNil

`func (o *SslVerification) SetCaCertificateNil(b bool)`

 SetCaCertificateNil sets the value for CaCertificate to be an explicit nil

### UnsetCaCertificate
`func (o *SslVerification) UnsetCaCertificate()`

UnsetCaCertificate ensures that no value is present for CaCertificate, not even an explicit nil
### GetIsEnabled

`func (o *SslVerification) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *SslVerification) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *SslVerification) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *SslVerification) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *SslVerification) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *SslVerification) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


