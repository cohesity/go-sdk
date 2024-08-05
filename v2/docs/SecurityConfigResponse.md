# SecurityConfigResponse

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

### NewSecurityConfigResponse

`func NewSecurityConfigResponse() *SecurityConfigResponse`

NewSecurityConfigResponse instantiates a new SecurityConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityConfigResponseWithDefaults

`func NewSecurityConfigResponseWithDefaults() *SecurityConfigResponse`

NewSecurityConfigResponseWithDefaults instantiates a new SecurityConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountLockout

`func (o *SecurityConfigResponse) GetAccountLockout() SecurityConfigAccountLockout`

GetAccountLockout returns the AccountLockout field if non-nil, zero value otherwise.

### GetAccountLockoutOk

`func (o *SecurityConfigResponse) GetAccountLockoutOk() (*SecurityConfigAccountLockout, bool)`

GetAccountLockoutOk returns a tuple with the AccountLockout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLockout

`func (o *SecurityConfigResponse) SetAccountLockout(v SecurityConfigAccountLockout)`

SetAccountLockout sets AccountLockout field to given value.

### HasAccountLockout

`func (o *SecurityConfigResponse) HasAccountLockout() bool`

HasAccountLockout returns a boolean if a field has been set.

### GetAuthTokenTimeoutMinutes

`func (o *SecurityConfigResponse) GetAuthTokenTimeoutMinutes() int32`

GetAuthTokenTimeoutMinutes returns the AuthTokenTimeoutMinutes field if non-nil, zero value otherwise.

### GetAuthTokenTimeoutMinutesOk

`func (o *SecurityConfigResponse) GetAuthTokenTimeoutMinutesOk() (*int32, bool)`

GetAuthTokenTimeoutMinutesOk returns a tuple with the AuthTokenTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTokenTimeoutMinutes

`func (o *SecurityConfigResponse) SetAuthTokenTimeoutMinutes(v int32)`

SetAuthTokenTimeoutMinutes sets AuthTokenTimeoutMinutes field to given value.

### HasAuthTokenTimeoutMinutes

`func (o *SecurityConfigResponse) HasAuthTokenTimeoutMinutes() bool`

HasAuthTokenTimeoutMinutes returns a boolean if a field has been set.

### SetAuthTokenTimeoutMinutesNil

`func (o *SecurityConfigResponse) SetAuthTokenTimeoutMinutesNil(b bool)`

 SetAuthTokenTimeoutMinutesNil sets the value for AuthTokenTimeoutMinutes to be an explicit nil

### UnsetAuthTokenTimeoutMinutes
`func (o *SecurityConfigResponse) UnsetAuthTokenTimeoutMinutes()`

UnsetAuthTokenTimeoutMinutes ensures that no value is present for AuthTokenTimeoutMinutes, not even an explicit nil
### GetCertificateBasedAuth

`func (o *SecurityConfigResponse) GetCertificateBasedAuth() SecurityConfigCertificateBasedAuth`

GetCertificateBasedAuth returns the CertificateBasedAuth field if non-nil, zero value otherwise.

### GetCertificateBasedAuthOk

`func (o *SecurityConfigResponse) GetCertificateBasedAuthOk() (*SecurityConfigCertificateBasedAuth, bool)`

GetCertificateBasedAuthOk returns a tuple with the CertificateBasedAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateBasedAuth

`func (o *SecurityConfigResponse) SetCertificateBasedAuth(v SecurityConfigCertificateBasedAuth)`

SetCertificateBasedAuth sets CertificateBasedAuth field to given value.

### HasCertificateBasedAuth

`func (o *SecurityConfigResponse) HasCertificateBasedAuth() bool`

HasCertificateBasedAuth returns a boolean if a field has been set.

### GetDataClassification

`func (o *SecurityConfigResponse) GetDataClassification() SecurityConfigDataClassification`

GetDataClassification returns the DataClassification field if non-nil, zero value otherwise.

### GetDataClassificationOk

`func (o *SecurityConfigResponse) GetDataClassificationOk() (*SecurityConfigDataClassification, bool)`

GetDataClassificationOk returns a tuple with the DataClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataClassification

`func (o *SecurityConfigResponse) SetDataClassification(v SecurityConfigDataClassification)`

SetDataClassification sets DataClassification field to given value.

### HasDataClassification

`func (o *SecurityConfigResponse) HasDataClassification() bool`

HasDataClassification returns a boolean if a field has been set.

### GetInactivityTimeoutMSecs

`func (o *SecurityConfigResponse) GetInactivityTimeoutMSecs() int64`

GetInactivityTimeoutMSecs returns the InactivityTimeoutMSecs field if non-nil, zero value otherwise.

### GetInactivityTimeoutMSecsOk

`func (o *SecurityConfigResponse) GetInactivityTimeoutMSecsOk() (*int64, bool)`

GetInactivityTimeoutMSecsOk returns a tuple with the InactivityTimeoutMSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactivityTimeoutMSecs

`func (o *SecurityConfigResponse) SetInactivityTimeoutMSecs(v int64)`

SetInactivityTimeoutMSecs sets InactivityTimeoutMSecs field to given value.

### HasInactivityTimeoutMSecs

`func (o *SecurityConfigResponse) HasInactivityTimeoutMSecs() bool`

HasInactivityTimeoutMSecs returns a boolean if a field has been set.

### SetInactivityTimeoutMSecsNil

`func (o *SecurityConfigResponse) SetInactivityTimeoutMSecsNil(b bool)`

 SetInactivityTimeoutMSecsNil sets the value for InactivityTimeoutMSecs to be an explicit nil

### UnsetInactivityTimeoutMSecs
`func (o *SecurityConfigResponse) UnsetInactivityTimeoutMSecs()`

UnsetInactivityTimeoutMSecs ensures that no value is present for InactivityTimeoutMSecs, not even an explicit nil
### GetPasswordLifetime

`func (o *SecurityConfigResponse) GetPasswordLifetime() SecurityConfigPasswordLifetime`

GetPasswordLifetime returns the PasswordLifetime field if non-nil, zero value otherwise.

### GetPasswordLifetimeOk

`func (o *SecurityConfigResponse) GetPasswordLifetimeOk() (*SecurityConfigPasswordLifetime, bool)`

GetPasswordLifetimeOk returns a tuple with the PasswordLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLifetime

`func (o *SecurityConfigResponse) SetPasswordLifetime(v SecurityConfigPasswordLifetime)`

SetPasswordLifetime sets PasswordLifetime field to given value.

### HasPasswordLifetime

`func (o *SecurityConfigResponse) HasPasswordLifetime() bool`

HasPasswordLifetime returns a boolean if a field has been set.

### GetPasswordReuse

`func (o *SecurityConfigResponse) GetPasswordReuse() SecurityConfigPasswordReuse`

GetPasswordReuse returns the PasswordReuse field if non-nil, zero value otherwise.

### GetPasswordReuseOk

`func (o *SecurityConfigResponse) GetPasswordReuseOk() (*SecurityConfigPasswordReuse, bool)`

GetPasswordReuseOk returns a tuple with the PasswordReuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordReuse

`func (o *SecurityConfigResponse) SetPasswordReuse(v SecurityConfigPasswordReuse)`

SetPasswordReuse sets PasswordReuse field to given value.

### HasPasswordReuse

`func (o *SecurityConfigResponse) HasPasswordReuse() bool`

HasPasswordReuse returns a boolean if a field has been set.

### GetPasswordStrength

`func (o *SecurityConfigResponse) GetPasswordStrength() SecurityConfigPasswordStrength`

GetPasswordStrength returns the PasswordStrength field if non-nil, zero value otherwise.

### GetPasswordStrengthOk

`func (o *SecurityConfigResponse) GetPasswordStrengthOk() (*SecurityConfigPasswordStrength, bool)`

GetPasswordStrengthOk returns a tuple with the PasswordStrength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordStrength

`func (o *SecurityConfigResponse) SetPasswordStrength(v SecurityConfigPasswordStrength)`

SetPasswordStrength sets PasswordStrength field to given value.

### HasPasswordStrength

`func (o *SecurityConfigResponse) HasPasswordStrength() bool`

HasPasswordStrength returns a boolean if a field has been set.

### GetSessionConfiguration

`func (o *SecurityConfigResponse) GetSessionConfiguration() SecurityConfigSessionConfiguration`

GetSessionConfiguration returns the SessionConfiguration field if non-nil, zero value otherwise.

### GetSessionConfigurationOk

`func (o *SecurityConfigResponse) GetSessionConfigurationOk() (*SecurityConfigSessionConfiguration, bool)`

GetSessionConfigurationOk returns a tuple with the SessionConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionConfiguration

`func (o *SecurityConfigResponse) SetSessionConfiguration(v SecurityConfigSessionConfiguration)`

SetSessionConfiguration sets SessionConfiguration field to given value.

### HasSessionConfiguration

`func (o *SecurityConfigResponse) HasSessionConfiguration() bool`

HasSessionConfiguration returns a boolean if a field has been set.

### GetSshConfiguration

`func (o *SecurityConfigResponse) GetSshConfiguration() SecurityConfigSshConfiguration`

GetSshConfiguration returns the SshConfiguration field if non-nil, zero value otherwise.

### GetSshConfigurationOk

`func (o *SecurityConfigResponse) GetSshConfigurationOk() (*SecurityConfigSshConfiguration, bool)`

GetSshConfigurationOk returns a tuple with the SshConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshConfiguration

`func (o *SecurityConfigResponse) SetSshConfiguration(v SecurityConfigSshConfiguration)`

SetSshConfiguration sets SshConfiguration field to given value.

### HasSshConfiguration

`func (o *SecurityConfigResponse) HasSshConfiguration() bool`

HasSshConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


