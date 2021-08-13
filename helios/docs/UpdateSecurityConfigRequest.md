# UpdateSecurityConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PasswordStrength** | Pointer to [**SecurityConfigPasswordStrength**](SecurityConfigPasswordStrength.md) |  | [optional] 
**PasswordReuse** | Pointer to [**SecurityConfigPasswordReuse**](SecurityConfigPasswordReuse.md) |  | [optional] 
**PasswordLifetime** | Pointer to [**SecurityConfigPasswordLifetime**](SecurityConfigPasswordLifetime.md) |  | [optional] 
**AccountLockout** | Pointer to [**SecurityConfigAccountLockout**](SecurityConfigAccountLockout.md) |  | [optional] 
**DataClassification** | Pointer to [**SecurityConfigDataClassification**](SecurityConfigDataClassification.md) |  | [optional] 
**SessionConfiguration** | Pointer to [**SecurityConfigSessionConfiguration**](SecurityConfigSessionConfiguration.md) |  | [optional] 
**CertificateBasedAuth** | Pointer to [**SecurityConfigCertificateBasedAuth**](SecurityConfigCertificateBasedAuth.md) |  | [optional] 

## Methods

### NewUpdateSecurityConfigRequest

`func NewUpdateSecurityConfigRequest() *UpdateSecurityConfigRequest`

NewUpdateSecurityConfigRequest instantiates a new UpdateSecurityConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSecurityConfigRequestWithDefaults

`func NewUpdateSecurityConfigRequestWithDefaults() *UpdateSecurityConfigRequest`

NewUpdateSecurityConfigRequestWithDefaults instantiates a new UpdateSecurityConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPasswordStrength

`func (o *UpdateSecurityConfigRequest) GetPasswordStrength() SecurityConfigPasswordStrength`

GetPasswordStrength returns the PasswordStrength field if non-nil, zero value otherwise.

### GetPasswordStrengthOk

`func (o *UpdateSecurityConfigRequest) GetPasswordStrengthOk() (*SecurityConfigPasswordStrength, bool)`

GetPasswordStrengthOk returns a tuple with the PasswordStrength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordStrength

`func (o *UpdateSecurityConfigRequest) SetPasswordStrength(v SecurityConfigPasswordStrength)`

SetPasswordStrength sets PasswordStrength field to given value.

### HasPasswordStrength

`func (o *UpdateSecurityConfigRequest) HasPasswordStrength() bool`

HasPasswordStrength returns a boolean if a field has been set.

### GetPasswordReuse

`func (o *UpdateSecurityConfigRequest) GetPasswordReuse() SecurityConfigPasswordReuse`

GetPasswordReuse returns the PasswordReuse field if non-nil, zero value otherwise.

### GetPasswordReuseOk

`func (o *UpdateSecurityConfigRequest) GetPasswordReuseOk() (*SecurityConfigPasswordReuse, bool)`

GetPasswordReuseOk returns a tuple with the PasswordReuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordReuse

`func (o *UpdateSecurityConfigRequest) SetPasswordReuse(v SecurityConfigPasswordReuse)`

SetPasswordReuse sets PasswordReuse field to given value.

### HasPasswordReuse

`func (o *UpdateSecurityConfigRequest) HasPasswordReuse() bool`

HasPasswordReuse returns a boolean if a field has been set.

### GetPasswordLifetime

`func (o *UpdateSecurityConfigRequest) GetPasswordLifetime() SecurityConfigPasswordLifetime`

GetPasswordLifetime returns the PasswordLifetime field if non-nil, zero value otherwise.

### GetPasswordLifetimeOk

`func (o *UpdateSecurityConfigRequest) GetPasswordLifetimeOk() (*SecurityConfigPasswordLifetime, bool)`

GetPasswordLifetimeOk returns a tuple with the PasswordLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLifetime

`func (o *UpdateSecurityConfigRequest) SetPasswordLifetime(v SecurityConfigPasswordLifetime)`

SetPasswordLifetime sets PasswordLifetime field to given value.

### HasPasswordLifetime

`func (o *UpdateSecurityConfigRequest) HasPasswordLifetime() bool`

HasPasswordLifetime returns a boolean if a field has been set.

### GetAccountLockout

`func (o *UpdateSecurityConfigRequest) GetAccountLockout() SecurityConfigAccountLockout`

GetAccountLockout returns the AccountLockout field if non-nil, zero value otherwise.

### GetAccountLockoutOk

`func (o *UpdateSecurityConfigRequest) GetAccountLockoutOk() (*SecurityConfigAccountLockout, bool)`

GetAccountLockoutOk returns a tuple with the AccountLockout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLockout

`func (o *UpdateSecurityConfigRequest) SetAccountLockout(v SecurityConfigAccountLockout)`

SetAccountLockout sets AccountLockout field to given value.

### HasAccountLockout

`func (o *UpdateSecurityConfigRequest) HasAccountLockout() bool`

HasAccountLockout returns a boolean if a field has been set.

### GetDataClassification

`func (o *UpdateSecurityConfigRequest) GetDataClassification() SecurityConfigDataClassification`

GetDataClassification returns the DataClassification field if non-nil, zero value otherwise.

### GetDataClassificationOk

`func (o *UpdateSecurityConfigRequest) GetDataClassificationOk() (*SecurityConfigDataClassification, bool)`

GetDataClassificationOk returns a tuple with the DataClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataClassification

`func (o *UpdateSecurityConfigRequest) SetDataClassification(v SecurityConfigDataClassification)`

SetDataClassification sets DataClassification field to given value.

### HasDataClassification

`func (o *UpdateSecurityConfigRequest) HasDataClassification() bool`

HasDataClassification returns a boolean if a field has been set.

### GetSessionConfiguration

`func (o *UpdateSecurityConfigRequest) GetSessionConfiguration() SecurityConfigSessionConfiguration`

GetSessionConfiguration returns the SessionConfiguration field if non-nil, zero value otherwise.

### GetSessionConfigurationOk

`func (o *UpdateSecurityConfigRequest) GetSessionConfigurationOk() (*SecurityConfigSessionConfiguration, bool)`

GetSessionConfigurationOk returns a tuple with the SessionConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionConfiguration

`func (o *UpdateSecurityConfigRequest) SetSessionConfiguration(v SecurityConfigSessionConfiguration)`

SetSessionConfiguration sets SessionConfiguration field to given value.

### HasSessionConfiguration

`func (o *UpdateSecurityConfigRequest) HasSessionConfiguration() bool`

HasSessionConfiguration returns a boolean if a field has been set.

### GetCertificateBasedAuth

`func (o *UpdateSecurityConfigRequest) GetCertificateBasedAuth() SecurityConfigCertificateBasedAuth`

GetCertificateBasedAuth returns the CertificateBasedAuth field if non-nil, zero value otherwise.

### GetCertificateBasedAuthOk

`func (o *UpdateSecurityConfigRequest) GetCertificateBasedAuthOk() (*SecurityConfigCertificateBasedAuth, bool)`

GetCertificateBasedAuthOk returns a tuple with the CertificateBasedAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateBasedAuth

`func (o *UpdateSecurityConfigRequest) SetCertificateBasedAuth(v SecurityConfigCertificateBasedAuth)`

SetCertificateBasedAuth sets CertificateBasedAuth field to given value.

### HasCertificateBasedAuth

`func (o *UpdateSecurityConfigRequest) HasCertificateBasedAuth() bool`

HasCertificateBasedAuth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


