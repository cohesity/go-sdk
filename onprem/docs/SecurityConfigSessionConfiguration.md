# SecurityConfigSessionConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbsoluteTimeout** | Pointer to **NullableInt64** | Specifies absolute session expiration time in seconds. | [optional] 
**InactivityTimeout** | Pointer to **NullableInt64** | Specifies inactivity session expiration time in seconds. | [optional] 
**LimitSessions** | Pointer to **NullableBool** | Specifies if limitations on number of active sessions is enabled or not. | [optional] 
**SessionLimitPerUser** | Pointer to **NullableInt64** | Specifies maximum number of active sessions allowed per user. This applies only when limitSessions is enabled. | [optional] 
**SessionLimitSystemWide** | Pointer to **NullableInt64** | Specifies maximum number of active sessions allowed system wide. This applies only when limitSessions is enabled. | [optional] 

## Methods

### NewSecurityConfigSessionConfiguration

`func NewSecurityConfigSessionConfiguration() *SecurityConfigSessionConfiguration`

NewSecurityConfigSessionConfiguration instantiates a new SecurityConfigSessionConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityConfigSessionConfigurationWithDefaults

`func NewSecurityConfigSessionConfigurationWithDefaults() *SecurityConfigSessionConfiguration`

NewSecurityConfigSessionConfigurationWithDefaults instantiates a new SecurityConfigSessionConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsoluteTimeout

`func (o *SecurityConfigSessionConfiguration) GetAbsoluteTimeout() int64`

GetAbsoluteTimeout returns the AbsoluteTimeout field if non-nil, zero value otherwise.

### GetAbsoluteTimeoutOk

`func (o *SecurityConfigSessionConfiguration) GetAbsoluteTimeoutOk() (*int64, bool)`

GetAbsoluteTimeoutOk returns a tuple with the AbsoluteTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteTimeout

`func (o *SecurityConfigSessionConfiguration) SetAbsoluteTimeout(v int64)`

SetAbsoluteTimeout sets AbsoluteTimeout field to given value.

### HasAbsoluteTimeout

`func (o *SecurityConfigSessionConfiguration) HasAbsoluteTimeout() bool`

HasAbsoluteTimeout returns a boolean if a field has been set.

### SetAbsoluteTimeoutNil

`func (o *SecurityConfigSessionConfiguration) SetAbsoluteTimeoutNil(b bool)`

 SetAbsoluteTimeoutNil sets the value for AbsoluteTimeout to be an explicit nil

### UnsetAbsoluteTimeout
`func (o *SecurityConfigSessionConfiguration) UnsetAbsoluteTimeout()`

UnsetAbsoluteTimeout ensures that no value is present for AbsoluteTimeout, not even an explicit nil
### GetInactivityTimeout

`func (o *SecurityConfigSessionConfiguration) GetInactivityTimeout() int64`

GetInactivityTimeout returns the InactivityTimeout field if non-nil, zero value otherwise.

### GetInactivityTimeoutOk

`func (o *SecurityConfigSessionConfiguration) GetInactivityTimeoutOk() (*int64, bool)`

GetInactivityTimeoutOk returns a tuple with the InactivityTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactivityTimeout

`func (o *SecurityConfigSessionConfiguration) SetInactivityTimeout(v int64)`

SetInactivityTimeout sets InactivityTimeout field to given value.

### HasInactivityTimeout

`func (o *SecurityConfigSessionConfiguration) HasInactivityTimeout() bool`

HasInactivityTimeout returns a boolean if a field has been set.

### SetInactivityTimeoutNil

`func (o *SecurityConfigSessionConfiguration) SetInactivityTimeoutNil(b bool)`

 SetInactivityTimeoutNil sets the value for InactivityTimeout to be an explicit nil

### UnsetInactivityTimeout
`func (o *SecurityConfigSessionConfiguration) UnsetInactivityTimeout()`

UnsetInactivityTimeout ensures that no value is present for InactivityTimeout, not even an explicit nil
### GetLimitSessions

`func (o *SecurityConfigSessionConfiguration) GetLimitSessions() bool`

GetLimitSessions returns the LimitSessions field if non-nil, zero value otherwise.

### GetLimitSessionsOk

`func (o *SecurityConfigSessionConfiguration) GetLimitSessionsOk() (*bool, bool)`

GetLimitSessionsOk returns a tuple with the LimitSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitSessions

`func (o *SecurityConfigSessionConfiguration) SetLimitSessions(v bool)`

SetLimitSessions sets LimitSessions field to given value.

### HasLimitSessions

`func (o *SecurityConfigSessionConfiguration) HasLimitSessions() bool`

HasLimitSessions returns a boolean if a field has been set.

### SetLimitSessionsNil

`func (o *SecurityConfigSessionConfiguration) SetLimitSessionsNil(b bool)`

 SetLimitSessionsNil sets the value for LimitSessions to be an explicit nil

### UnsetLimitSessions
`func (o *SecurityConfigSessionConfiguration) UnsetLimitSessions()`

UnsetLimitSessions ensures that no value is present for LimitSessions, not even an explicit nil
### GetSessionLimitPerUser

`func (o *SecurityConfigSessionConfiguration) GetSessionLimitPerUser() int64`

GetSessionLimitPerUser returns the SessionLimitPerUser field if non-nil, zero value otherwise.

### GetSessionLimitPerUserOk

`func (o *SecurityConfigSessionConfiguration) GetSessionLimitPerUserOk() (*int64, bool)`

GetSessionLimitPerUserOk returns a tuple with the SessionLimitPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionLimitPerUser

`func (o *SecurityConfigSessionConfiguration) SetSessionLimitPerUser(v int64)`

SetSessionLimitPerUser sets SessionLimitPerUser field to given value.

### HasSessionLimitPerUser

`func (o *SecurityConfigSessionConfiguration) HasSessionLimitPerUser() bool`

HasSessionLimitPerUser returns a boolean if a field has been set.

### SetSessionLimitPerUserNil

`func (o *SecurityConfigSessionConfiguration) SetSessionLimitPerUserNil(b bool)`

 SetSessionLimitPerUserNil sets the value for SessionLimitPerUser to be an explicit nil

### UnsetSessionLimitPerUser
`func (o *SecurityConfigSessionConfiguration) UnsetSessionLimitPerUser()`

UnsetSessionLimitPerUser ensures that no value is present for SessionLimitPerUser, not even an explicit nil
### GetSessionLimitSystemWide

`func (o *SecurityConfigSessionConfiguration) GetSessionLimitSystemWide() int64`

GetSessionLimitSystemWide returns the SessionLimitSystemWide field if non-nil, zero value otherwise.

### GetSessionLimitSystemWideOk

`func (o *SecurityConfigSessionConfiguration) GetSessionLimitSystemWideOk() (*int64, bool)`

GetSessionLimitSystemWideOk returns a tuple with the SessionLimitSystemWide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionLimitSystemWide

`func (o *SecurityConfigSessionConfiguration) SetSessionLimitSystemWide(v int64)`

SetSessionLimitSystemWide sets SessionLimitSystemWide field to given value.

### HasSessionLimitSystemWide

`func (o *SecurityConfigSessionConfiguration) HasSessionLimitSystemWide() bool`

HasSessionLimitSystemWide returns a boolean if a field has been set.

### SetSessionLimitSystemWideNil

`func (o *SecurityConfigSessionConfiguration) SetSessionLimitSystemWideNil(b bool)`

 SetSessionLimitSystemWideNil sets the value for SessionLimitSystemWide to be an explicit nil

### UnsetSessionLimitSystemWide
`func (o *SecurityConfigSessionConfiguration) UnsetSessionLimitSystemWide()`

UnsetSessionLimitSystemWide ensures that no value is present for SessionLimitSystemWide, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


