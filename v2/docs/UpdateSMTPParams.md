# UpdateSMTPParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** | Specifies the IP address or the FQDN of the SMTP server. | 
**IsActive** | Pointer to **NullableBool** | Specifies if the SMTP configuration is active. | [optional] [default to true]
**Port** | **int32** | Specifies the SMTP port. Usually 465 or 587. For authenticated connection, it is generally 587. | 
**UseSSL** | Pointer to **NullableBool** | This is set to true when the SMTP server uses SSL/TLS without supporting STARTTLS. Typically, this is used for port 465. | [optional] [default to false]
**Username** | Pointer to **NullableString** | Specifies the username which will be used to connect to the SMTP server. If username is not specified, then it would imply that SMTP server is set up for unauthenticated access. | [optional] 
**Password** | Pointer to **NullableString** | Specifies the password of the SMTP user. This is required if username is specified in the request. | [optional] 

## Methods

### NewUpdateSMTPParams

`func NewUpdateSMTPParams(hostname string, port int32, ) *UpdateSMTPParams`

NewUpdateSMTPParams instantiates a new UpdateSMTPParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSMTPParamsWithDefaults

`func NewUpdateSMTPParamsWithDefaults() *UpdateSMTPParams`

NewUpdateSMTPParamsWithDefaults instantiates a new UpdateSMTPParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *UpdateSMTPParams) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *UpdateSMTPParams) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *UpdateSMTPParams) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetIsActive

`func (o *UpdateSMTPParams) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UpdateSMTPParams) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UpdateSMTPParams) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *UpdateSMTPParams) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *UpdateSMTPParams) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *UpdateSMTPParams) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetPort

`func (o *UpdateSMTPParams) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UpdateSMTPParams) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UpdateSMTPParams) SetPort(v int32)`

SetPort sets Port field to given value.


### GetUseSSL

`func (o *UpdateSMTPParams) GetUseSSL() bool`

GetUseSSL returns the UseSSL field if non-nil, zero value otherwise.

### GetUseSSLOk

`func (o *UpdateSMTPParams) GetUseSSLOk() (*bool, bool)`

GetUseSSLOk returns a tuple with the UseSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSSL

`func (o *UpdateSMTPParams) SetUseSSL(v bool)`

SetUseSSL sets UseSSL field to given value.

### HasUseSSL

`func (o *UpdateSMTPParams) HasUseSSL() bool`

HasUseSSL returns a boolean if a field has been set.

### SetUseSSLNil

`func (o *UpdateSMTPParams) SetUseSSLNil(b bool)`

 SetUseSSLNil sets the value for UseSSL to be an explicit nil

### UnsetUseSSL
`func (o *UpdateSMTPParams) UnsetUseSSL()`

UnsetUseSSL ensures that no value is present for UseSSL, not even an explicit nil
### GetUsername

`func (o *UpdateSMTPParams) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateSMTPParams) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateSMTPParams) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateSMTPParams) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *UpdateSMTPParams) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *UpdateSMTPParams) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *UpdateSMTPParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateSMTPParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateSMTPParams) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateSMTPParams) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *UpdateSMTPParams) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *UpdateSMTPParams) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


