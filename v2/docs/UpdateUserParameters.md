# UpdateUserParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Specifies the description of the User. | [optional] 
**EffectiveTimeMsecs** | Pointer to **NullableInt64** | Specifies the epoch time in milliseconds since when the user can login. | [optional] 
**ExpiryTimeMsecs** | Pointer to **NullableInt64** | Specifies the epoch time in milliseconds when the user expires. Post expiry the user cannot access Cohesity cluster. | [optional] 
**Locked** | Pointer to **NullableBool** | Specifies whether the User is locked. | [optional] 
**Restricted** | Pointer to **NullableBool** | Specifies whether the User is restricted. A restricted user can only view &amp; manage the objects it has permissions to. | [optional] 
**Roles** | Pointer to **[]string** | Specifies the Cohesity roles to associate with the user. The Cohesity roles determine privileges on the Cohesity Cluster for this user. | [optional] 
**LocalUserParams** | Pointer to **map[string]interface{}** | Specifies the LOCAL user properties. This field is required when updating LOCAL Cohesity User params. | [optional] 
**Username** | Pointer to **NullableString** | Specifies the username. | [optional] 

## Methods

### NewUpdateUserParameters

`func NewUpdateUserParameters() *UpdateUserParameters`

NewUpdateUserParameters instantiates a new UpdateUserParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserParametersWithDefaults

`func NewUpdateUserParametersWithDefaults() *UpdateUserParameters`

NewUpdateUserParametersWithDefaults instantiates a new UpdateUserParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateUserParameters) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateUserParameters) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateUserParameters) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateUserParameters) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateUserParameters) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateUserParameters) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEffectiveTimeMsecs

`func (o *UpdateUserParameters) GetEffectiveTimeMsecs() int64`

GetEffectiveTimeMsecs returns the EffectiveTimeMsecs field if non-nil, zero value otherwise.

### GetEffectiveTimeMsecsOk

`func (o *UpdateUserParameters) GetEffectiveTimeMsecsOk() (*int64, bool)`

GetEffectiveTimeMsecsOk returns a tuple with the EffectiveTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTimeMsecs

`func (o *UpdateUserParameters) SetEffectiveTimeMsecs(v int64)`

SetEffectiveTimeMsecs sets EffectiveTimeMsecs field to given value.

### HasEffectiveTimeMsecs

`func (o *UpdateUserParameters) HasEffectiveTimeMsecs() bool`

HasEffectiveTimeMsecs returns a boolean if a field has been set.

### SetEffectiveTimeMsecsNil

`func (o *UpdateUserParameters) SetEffectiveTimeMsecsNil(b bool)`

 SetEffectiveTimeMsecsNil sets the value for EffectiveTimeMsecs to be an explicit nil

### UnsetEffectiveTimeMsecs
`func (o *UpdateUserParameters) UnsetEffectiveTimeMsecs()`

UnsetEffectiveTimeMsecs ensures that no value is present for EffectiveTimeMsecs, not even an explicit nil
### GetExpiryTimeMsecs

`func (o *UpdateUserParameters) GetExpiryTimeMsecs() int64`

GetExpiryTimeMsecs returns the ExpiryTimeMsecs field if non-nil, zero value otherwise.

### GetExpiryTimeMsecsOk

`func (o *UpdateUserParameters) GetExpiryTimeMsecsOk() (*int64, bool)`

GetExpiryTimeMsecsOk returns a tuple with the ExpiryTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimeMsecs

`func (o *UpdateUserParameters) SetExpiryTimeMsecs(v int64)`

SetExpiryTimeMsecs sets ExpiryTimeMsecs field to given value.

### HasExpiryTimeMsecs

`func (o *UpdateUserParameters) HasExpiryTimeMsecs() bool`

HasExpiryTimeMsecs returns a boolean if a field has been set.

### SetExpiryTimeMsecsNil

`func (o *UpdateUserParameters) SetExpiryTimeMsecsNil(b bool)`

 SetExpiryTimeMsecsNil sets the value for ExpiryTimeMsecs to be an explicit nil

### UnsetExpiryTimeMsecs
`func (o *UpdateUserParameters) UnsetExpiryTimeMsecs()`

UnsetExpiryTimeMsecs ensures that no value is present for ExpiryTimeMsecs, not even an explicit nil
### GetLocked

`func (o *UpdateUserParameters) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *UpdateUserParameters) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *UpdateUserParameters) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *UpdateUserParameters) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### SetLockedNil

`func (o *UpdateUserParameters) SetLockedNil(b bool)`

 SetLockedNil sets the value for Locked to be an explicit nil

### UnsetLocked
`func (o *UpdateUserParameters) UnsetLocked()`

UnsetLocked ensures that no value is present for Locked, not even an explicit nil
### GetRestricted

`func (o *UpdateUserParameters) GetRestricted() bool`

GetRestricted returns the Restricted field if non-nil, zero value otherwise.

### GetRestrictedOk

`func (o *UpdateUserParameters) GetRestrictedOk() (*bool, bool)`

GetRestrictedOk returns a tuple with the Restricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestricted

`func (o *UpdateUserParameters) SetRestricted(v bool)`

SetRestricted sets Restricted field to given value.

### HasRestricted

`func (o *UpdateUserParameters) HasRestricted() bool`

HasRestricted returns a boolean if a field has been set.

### SetRestrictedNil

`func (o *UpdateUserParameters) SetRestrictedNil(b bool)`

 SetRestrictedNil sets the value for Restricted to be an explicit nil

### UnsetRestricted
`func (o *UpdateUserParameters) UnsetRestricted()`

UnsetRestricted ensures that no value is present for Restricted, not even an explicit nil
### GetRoles

`func (o *UpdateUserParameters) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UpdateUserParameters) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UpdateUserParameters) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UpdateUserParameters) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *UpdateUserParameters) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *UpdateUserParameters) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetLocalUserParams

`func (o *UpdateUserParameters) GetLocalUserParams() map[string]interface{}`

GetLocalUserParams returns the LocalUserParams field if non-nil, zero value otherwise.

### GetLocalUserParamsOk

`func (o *UpdateUserParameters) GetLocalUserParamsOk() (*map[string]interface{}, bool)`

GetLocalUserParamsOk returns a tuple with the LocalUserParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalUserParams

`func (o *UpdateUserParameters) SetLocalUserParams(v map[string]interface{})`

SetLocalUserParams sets LocalUserParams field to given value.

### HasLocalUserParams

`func (o *UpdateUserParameters) HasLocalUserParams() bool`

HasLocalUserParams returns a boolean if a field has been set.

### GetUsername

`func (o *UpdateUserParameters) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateUserParameters) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateUserParameters) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateUserParameters) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *UpdateUserParameters) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *UpdateUserParameters) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


