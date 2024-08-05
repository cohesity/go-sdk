# UpdateSecurityConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountLockout** | Pointer to [**SecurityConfigAccountLockout**](SecurityConfigAccountLockout.md) |  | [optional] 
**AuthTokenTimeoutMinutes** | Pointer to **NullableInt32** | Specifies the authentication token timeout in minutes. Applies both for API based access token and browser login cookie. | [optional] 
**CertificateBasedAuth** | Pointer to [**SecurityConfigCertificateBasedAuth**](SecurityConfigCertificateBasedAuth.md) |  | [optional] 
**DataClassification** | Pointer to [**SecurityConfigDataClassification**](SecurityConfigDataClassification.md) |  | [optional] 
**InactivityTimeoutMSecs** | Pointer to **NullableInt64** | Specifies the UI inactivity timeout in milliseconds. Default value is 30 minutes. | [optional] 
**PasswordLifetime** | Pointer to [**SecurityConfigPasswordLifetime**](SecurityConfigPasswordLifetime.md) |  | [optional] 
**PasswordReuse** | Pointer to [**SecurityConfigPasswordReuse**](SecurityConfigPasswordReuse.md) |  | [optional] 
**PasswordStrength** | Pointer to [**SecurityConfigPasswordStrength**](SecurityConfigPasswordStrength.md) |  | [optional] 
**SessionConfiguration** | Pointer to [**SecurityConfigSessionConfiguration**](SecurityConfigSessionConfiguration.md) |  | [optional] 
**SshConfiguration** | Pointer to [**SecurityConfigSshConfiguration**](SecurityConfigSshConfiguration.md) |  | [optional] 

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

### GetAuthTokenTimeoutMinutes

`func (o *UpdateSecurityConfigRequest) GetAuthTokenTimeoutMinutes() int32`

GetAuthTokenTimeoutMinutes returns the AuthTokenTimeoutMinutes field if non-nil, zero value otherwise.

### GetAuthTokenTimeoutMinutesOk

`func (o *UpdateSecurityConfigRequest) GetAuthTokenTimeoutMinutesOk() (*int32, bool)`

GetAuthTokenTimeoutMinutesOk returns a tuple with the AuthTokenTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTokenTimeoutMinutes

`func (o *UpdateSecurityConfigRequest) SetAuthTokenTimeoutMinutes(v int32)`

SetAuthTokenTimeoutMinutes sets AuthTokenTimeoutMinutes field to given value.

### HasAuthTokenTimeoutMinutes

`func (o *UpdateSecurityConfigRequest) HasAuthTokenTimeoutMinutes() bool`

HasAuthTokenTimeoutMinutes returns a boolean if a field has been set.

### SetAuthTokenTimeoutMinutesNil

`func (o *UpdateSecurityConfigRequest) SetAuthTokenTimeoutMinutesNil(b bool)`

 SetAuthTokenTimeoutMinutesNil sets the value for AuthTokenTimeoutMinutes to be an explicit nil

### UnsetAuthTokenTimeoutMinutes
`func (o *UpdateSecurityConfigRequest) UnsetAuthTokenTimeoutMinutes()`

UnsetAuthTokenTimeoutMinutes ensures that no value is present for AuthTokenTimeoutMinutes, not even an explicit nil
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

### GetInactivityTimeoutMSecs

`func (o *UpdateSecurityConfigRequest) GetInactivityTimeoutMSecs() int64`

GetInactivityTimeoutMSecs returns the InactivityTimeoutMSecs field if non-nil, zero value otherwise.

### GetInactivityTimeoutMSecsOk

`func (o *UpdateSecurityConfigRequest) GetInactivityTimeoutMSecsOk() (*int64, bool)`

GetInactivityTimeoutMSecsOk returns a tuple with the InactivityTimeoutMSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactivityTimeoutMSecs

`func (o *UpdateSecurityConfigRequest) SetInactivityTimeoutMSecs(v int64)`

SetInactivityTimeoutMSecs sets InactivityTimeoutMSecs field to given value.

### HasInactivityTimeoutMSecs

`func (o *UpdateSecurityConfigRequest) HasInactivityTimeoutMSecs() bool`

HasInactivityTimeoutMSecs returns a boolean if a field has been set.

### SetInactivityTimeoutMSecsNil

`func (o *UpdateSecurityConfigRequest) SetInactivityTimeoutMSecsNil(b bool)`

 SetInactivityTimeoutMSecsNil sets the value for InactivityTimeoutMSecs to be an explicit nil

### UnsetInactivityTimeoutMSecs
`func (o *UpdateSecurityConfigRequest) UnsetInactivityTimeoutMSecs()`

UnsetInactivityTimeoutMSecs ensures that no value is present for InactivityTimeoutMSecs, not even an explicit nil
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

### GetSshConfiguration

`func (o *UpdateSecurityConfigRequest) GetSshConfiguration() SecurityConfigSshConfiguration`

GetSshConfiguration returns the SshConfiguration field if non-nil, zero value otherwise.

### GetSshConfigurationOk

`func (o *UpdateSecurityConfigRequest) GetSshConfigurationOk() (*SecurityConfigSshConfiguration, bool)`

GetSshConfigurationOk returns a tuple with the SshConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshConfiguration

`func (o *UpdateSecurityConfigRequest) SetSshConfiguration(v SecurityConfigSshConfiguration)`

SetSshConfiguration sets SshConfiguration field to given value.

### HasSshConfiguration

`func (o *UpdateSecurityConfigRequest) HasSshConfiguration() bool`

HasSshConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


