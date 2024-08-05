# SecurityConfig

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

### NewSecurityConfig

`func NewSecurityConfig() *SecurityConfig`

NewSecurityConfig instantiates a new SecurityConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityConfigWithDefaults

`func NewSecurityConfigWithDefaults() *SecurityConfig`

NewSecurityConfigWithDefaults instantiates a new SecurityConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountLockout

`func (o *SecurityConfig) GetAccountLockout() SecurityConfigAccountLockout`

GetAccountLockout returns the AccountLockout field if non-nil, zero value otherwise.

### GetAccountLockoutOk

`func (o *SecurityConfig) GetAccountLockoutOk() (*SecurityConfigAccountLockout, bool)`

GetAccountLockoutOk returns a tuple with the AccountLockout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLockout

`func (o *SecurityConfig) SetAccountLockout(v SecurityConfigAccountLockout)`

SetAccountLockout sets AccountLockout field to given value.

### HasAccountLockout

`func (o *SecurityConfig) HasAccountLockout() bool`

HasAccountLockout returns a boolean if a field has been set.

### GetAuthTokenTimeoutMinutes

`func (o *SecurityConfig) GetAuthTokenTimeoutMinutes() int32`

GetAuthTokenTimeoutMinutes returns the AuthTokenTimeoutMinutes field if non-nil, zero value otherwise.

### GetAuthTokenTimeoutMinutesOk

`func (o *SecurityConfig) GetAuthTokenTimeoutMinutesOk() (*int32, bool)`

GetAuthTokenTimeoutMinutesOk returns a tuple with the AuthTokenTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTokenTimeoutMinutes

`func (o *SecurityConfig) SetAuthTokenTimeoutMinutes(v int32)`

SetAuthTokenTimeoutMinutes sets AuthTokenTimeoutMinutes field to given value.

### HasAuthTokenTimeoutMinutes

`func (o *SecurityConfig) HasAuthTokenTimeoutMinutes() bool`

HasAuthTokenTimeoutMinutes returns a boolean if a field has been set.

### SetAuthTokenTimeoutMinutesNil

`func (o *SecurityConfig) SetAuthTokenTimeoutMinutesNil(b bool)`

 SetAuthTokenTimeoutMinutesNil sets the value for AuthTokenTimeoutMinutes to be an explicit nil

### UnsetAuthTokenTimeoutMinutes
`func (o *SecurityConfig) UnsetAuthTokenTimeoutMinutes()`

UnsetAuthTokenTimeoutMinutes ensures that no value is present for AuthTokenTimeoutMinutes, not even an explicit nil
### GetCertificateBasedAuth

`func (o *SecurityConfig) GetCertificateBasedAuth() SecurityConfigCertificateBasedAuth`

GetCertificateBasedAuth returns the CertificateBasedAuth field if non-nil, zero value otherwise.

### GetCertificateBasedAuthOk

`func (o *SecurityConfig) GetCertificateBasedAuthOk() (*SecurityConfigCertificateBasedAuth, bool)`

GetCertificateBasedAuthOk returns a tuple with the CertificateBasedAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateBasedAuth

`func (o *SecurityConfig) SetCertificateBasedAuth(v SecurityConfigCertificateBasedAuth)`

SetCertificateBasedAuth sets CertificateBasedAuth field to given value.

### HasCertificateBasedAuth

`func (o *SecurityConfig) HasCertificateBasedAuth() bool`

HasCertificateBasedAuth returns a boolean if a field has been set.

### GetDataClassification

`func (o *SecurityConfig) GetDataClassification() SecurityConfigDataClassification`

GetDataClassification returns the DataClassification field if non-nil, zero value otherwise.

### GetDataClassificationOk

`func (o *SecurityConfig) GetDataClassificationOk() (*SecurityConfigDataClassification, bool)`

GetDataClassificationOk returns a tuple with the DataClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataClassification

`func (o *SecurityConfig) SetDataClassification(v SecurityConfigDataClassification)`

SetDataClassification sets DataClassification field to given value.

### HasDataClassification

`func (o *SecurityConfig) HasDataClassification() bool`

HasDataClassification returns a boolean if a field has been set.

### GetInactivityTimeoutMSecs

`func (o *SecurityConfig) GetInactivityTimeoutMSecs() int64`

GetInactivityTimeoutMSecs returns the InactivityTimeoutMSecs field if non-nil, zero value otherwise.

### GetInactivityTimeoutMSecsOk

`func (o *SecurityConfig) GetInactivityTimeoutMSecsOk() (*int64, bool)`

GetInactivityTimeoutMSecsOk returns a tuple with the InactivityTimeoutMSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactivityTimeoutMSecs

`func (o *SecurityConfig) SetInactivityTimeoutMSecs(v int64)`

SetInactivityTimeoutMSecs sets InactivityTimeoutMSecs field to given value.

### HasInactivityTimeoutMSecs

`func (o *SecurityConfig) HasInactivityTimeoutMSecs() bool`

HasInactivityTimeoutMSecs returns a boolean if a field has been set.

### SetInactivityTimeoutMSecsNil

`func (o *SecurityConfig) SetInactivityTimeoutMSecsNil(b bool)`

 SetInactivityTimeoutMSecsNil sets the value for InactivityTimeoutMSecs to be an explicit nil

### UnsetInactivityTimeoutMSecs
`func (o *SecurityConfig) UnsetInactivityTimeoutMSecs()`

UnsetInactivityTimeoutMSecs ensures that no value is present for InactivityTimeoutMSecs, not even an explicit nil
### GetPasswordLifetime

`func (o *SecurityConfig) GetPasswordLifetime() SecurityConfigPasswordLifetime`

GetPasswordLifetime returns the PasswordLifetime field if non-nil, zero value otherwise.

### GetPasswordLifetimeOk

`func (o *SecurityConfig) GetPasswordLifetimeOk() (*SecurityConfigPasswordLifetime, bool)`

GetPasswordLifetimeOk returns a tuple with the PasswordLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLifetime

`func (o *SecurityConfig) SetPasswordLifetime(v SecurityConfigPasswordLifetime)`

SetPasswordLifetime sets PasswordLifetime field to given value.

### HasPasswordLifetime

`func (o *SecurityConfig) HasPasswordLifetime() bool`

HasPasswordLifetime returns a boolean if a field has been set.

### GetPasswordReuse

`func (o *SecurityConfig) GetPasswordReuse() SecurityConfigPasswordReuse`

GetPasswordReuse returns the PasswordReuse field if non-nil, zero value otherwise.

### GetPasswordReuseOk

`func (o *SecurityConfig) GetPasswordReuseOk() (*SecurityConfigPasswordReuse, bool)`

GetPasswordReuseOk returns a tuple with the PasswordReuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordReuse

`func (o *SecurityConfig) SetPasswordReuse(v SecurityConfigPasswordReuse)`

SetPasswordReuse sets PasswordReuse field to given value.

### HasPasswordReuse

`func (o *SecurityConfig) HasPasswordReuse() bool`

HasPasswordReuse returns a boolean if a field has been set.

### GetPasswordStrength

`func (o *SecurityConfig) GetPasswordStrength() SecurityConfigPasswordStrength`

GetPasswordStrength returns the PasswordStrength field if non-nil, zero value otherwise.

### GetPasswordStrengthOk

`func (o *SecurityConfig) GetPasswordStrengthOk() (*SecurityConfigPasswordStrength, bool)`

GetPasswordStrengthOk returns a tuple with the PasswordStrength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordStrength

`func (o *SecurityConfig) SetPasswordStrength(v SecurityConfigPasswordStrength)`

SetPasswordStrength sets PasswordStrength field to given value.

### HasPasswordStrength

`func (o *SecurityConfig) HasPasswordStrength() bool`

HasPasswordStrength returns a boolean if a field has been set.

### GetSessionConfiguration

`func (o *SecurityConfig) GetSessionConfiguration() SecurityConfigSessionConfiguration`

GetSessionConfiguration returns the SessionConfiguration field if non-nil, zero value otherwise.

### GetSessionConfigurationOk

`func (o *SecurityConfig) GetSessionConfigurationOk() (*SecurityConfigSessionConfiguration, bool)`

GetSessionConfigurationOk returns a tuple with the SessionConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionConfiguration

`func (o *SecurityConfig) SetSessionConfiguration(v SecurityConfigSessionConfiguration)`

SetSessionConfiguration sets SessionConfiguration field to given value.

### HasSessionConfiguration

`func (o *SecurityConfig) HasSessionConfiguration() bool`

HasSessionConfiguration returns a boolean if a field has been set.

### GetSshConfiguration

`func (o *SecurityConfig) GetSshConfiguration() SecurityConfigSshConfiguration`

GetSshConfiguration returns the SshConfiguration field if non-nil, zero value otherwise.

### GetSshConfigurationOk

`func (o *SecurityConfig) GetSshConfigurationOk() (*SecurityConfigSshConfiguration, bool)`

GetSshConfigurationOk returns a tuple with the SshConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshConfiguration

`func (o *SecurityConfig) SetSshConfiguration(v SecurityConfigSshConfiguration)`

SetSshConfiguration sets SshConfiguration field to given value.

### HasSshConfiguration

`func (o *SecurityConfig) HasSshConfiguration() bool`

HasSshConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


