# SecurityConfigAccountLockout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedLoginLockTimeDurationMins** | Pointer to **NullableInt32** | Specifies the time duration within which the consecutive failed login attempts causes a local user account to be locked and the lockout duration time due to that. | [optional] 
**InactivityTimeDays** | Pointer to **NullableInt32** | Specifies the lockout inactivity time range in days. | [optional] 
**MaxFailedLoginAttempts** | Pointer to **NullableInt32** | Specifies the maximum number of consecutive fail login attempts. | [optional] 

## Methods

### NewSecurityConfigAccountLockout

`func NewSecurityConfigAccountLockout() *SecurityConfigAccountLockout`

NewSecurityConfigAccountLockout instantiates a new SecurityConfigAccountLockout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityConfigAccountLockoutWithDefaults

`func NewSecurityConfigAccountLockoutWithDefaults() *SecurityConfigAccountLockout`

NewSecurityConfigAccountLockoutWithDefaults instantiates a new SecurityConfigAccountLockout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailedLoginLockTimeDurationMins

`func (o *SecurityConfigAccountLockout) GetFailedLoginLockTimeDurationMins() int32`

GetFailedLoginLockTimeDurationMins returns the FailedLoginLockTimeDurationMins field if non-nil, zero value otherwise.

### GetFailedLoginLockTimeDurationMinsOk

`func (o *SecurityConfigAccountLockout) GetFailedLoginLockTimeDurationMinsOk() (*int32, bool)`

GetFailedLoginLockTimeDurationMinsOk returns a tuple with the FailedLoginLockTimeDurationMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedLoginLockTimeDurationMins

`func (o *SecurityConfigAccountLockout) SetFailedLoginLockTimeDurationMins(v int32)`

SetFailedLoginLockTimeDurationMins sets FailedLoginLockTimeDurationMins field to given value.

### HasFailedLoginLockTimeDurationMins

`func (o *SecurityConfigAccountLockout) HasFailedLoginLockTimeDurationMins() bool`

HasFailedLoginLockTimeDurationMins returns a boolean if a field has been set.

### SetFailedLoginLockTimeDurationMinsNil

`func (o *SecurityConfigAccountLockout) SetFailedLoginLockTimeDurationMinsNil(b bool)`

 SetFailedLoginLockTimeDurationMinsNil sets the value for FailedLoginLockTimeDurationMins to be an explicit nil

### UnsetFailedLoginLockTimeDurationMins
`func (o *SecurityConfigAccountLockout) UnsetFailedLoginLockTimeDurationMins()`

UnsetFailedLoginLockTimeDurationMins ensures that no value is present for FailedLoginLockTimeDurationMins, not even an explicit nil
### GetInactivityTimeDays

`func (o *SecurityConfigAccountLockout) GetInactivityTimeDays() int32`

GetInactivityTimeDays returns the InactivityTimeDays field if non-nil, zero value otherwise.

### GetInactivityTimeDaysOk

`func (o *SecurityConfigAccountLockout) GetInactivityTimeDaysOk() (*int32, bool)`

GetInactivityTimeDaysOk returns a tuple with the InactivityTimeDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactivityTimeDays

`func (o *SecurityConfigAccountLockout) SetInactivityTimeDays(v int32)`

SetInactivityTimeDays sets InactivityTimeDays field to given value.

### HasInactivityTimeDays

`func (o *SecurityConfigAccountLockout) HasInactivityTimeDays() bool`

HasInactivityTimeDays returns a boolean if a field has been set.

### SetInactivityTimeDaysNil

`func (o *SecurityConfigAccountLockout) SetInactivityTimeDaysNil(b bool)`

 SetInactivityTimeDaysNil sets the value for InactivityTimeDays to be an explicit nil

### UnsetInactivityTimeDays
`func (o *SecurityConfigAccountLockout) UnsetInactivityTimeDays()`

UnsetInactivityTimeDays ensures that no value is present for InactivityTimeDays, not even an explicit nil
### GetMaxFailedLoginAttempts

`func (o *SecurityConfigAccountLockout) GetMaxFailedLoginAttempts() int32`

GetMaxFailedLoginAttempts returns the MaxFailedLoginAttempts field if non-nil, zero value otherwise.

### GetMaxFailedLoginAttemptsOk

`func (o *SecurityConfigAccountLockout) GetMaxFailedLoginAttemptsOk() (*int32, bool)`

GetMaxFailedLoginAttemptsOk returns a tuple with the MaxFailedLoginAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFailedLoginAttempts

`func (o *SecurityConfigAccountLockout) SetMaxFailedLoginAttempts(v int32)`

SetMaxFailedLoginAttempts sets MaxFailedLoginAttempts field to given value.

### HasMaxFailedLoginAttempts

`func (o *SecurityConfigAccountLockout) HasMaxFailedLoginAttempts() bool`

HasMaxFailedLoginAttempts returns a boolean if a field has been set.

### SetMaxFailedLoginAttemptsNil

`func (o *SecurityConfigAccountLockout) SetMaxFailedLoginAttemptsNil(b bool)`

 SetMaxFailedLoginAttemptsNil sets the value for MaxFailedLoginAttempts to be an explicit nil

### UnsetMaxFailedLoginAttempts
`func (o *SecurityConfigAccountLockout) UnsetMaxFailedLoginAttempts()`

UnsetMaxFailedLoginAttempts ensures that no value is present for MaxFailedLoginAttempts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


