# CommonUpdatableUserParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Specifies the description of the User. | [optional] 
**EffectiveTimeMsecs** | Pointer to **NullableInt64** | Specifies the epoch time in milliseconds since when the user can login. | [optional] 
**ExpiryTimeMsecs** | Pointer to **NullableInt64** | Specifies the epoch time in milliseconds when the user expires. Post expiry the user cannot access Cohesity cluster. | [optional] 
**Locked** | Pointer to **NullableBool** | Specifies whether the User is locked. | [optional] 
**Restricted** | Pointer to **NullableBool** | Specifies whether the User is restricted. A restricted user can only view &amp; manage the objects it has permissions to. | [optional] 
**Roles** | Pointer to **[]string** | Specifies the Cohesity roles to associate with the user. The Cohesity roles determine privileges on the Cohesity Cluster for this user. | [optional] 

## Methods

### NewCommonUpdatableUserParams

`func NewCommonUpdatableUserParams() *CommonUpdatableUserParams`

NewCommonUpdatableUserParams instantiates a new CommonUpdatableUserParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonUpdatableUserParamsWithDefaults

`func NewCommonUpdatableUserParamsWithDefaults() *CommonUpdatableUserParams`

NewCommonUpdatableUserParamsWithDefaults instantiates a new CommonUpdatableUserParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CommonUpdatableUserParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CommonUpdatableUserParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CommonUpdatableUserParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CommonUpdatableUserParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CommonUpdatableUserParams) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CommonUpdatableUserParams) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEffectiveTimeMsecs

`func (o *CommonUpdatableUserParams) GetEffectiveTimeMsecs() int64`

GetEffectiveTimeMsecs returns the EffectiveTimeMsecs field if non-nil, zero value otherwise.

### GetEffectiveTimeMsecsOk

`func (o *CommonUpdatableUserParams) GetEffectiveTimeMsecsOk() (*int64, bool)`

GetEffectiveTimeMsecsOk returns a tuple with the EffectiveTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTimeMsecs

`func (o *CommonUpdatableUserParams) SetEffectiveTimeMsecs(v int64)`

SetEffectiveTimeMsecs sets EffectiveTimeMsecs field to given value.

### HasEffectiveTimeMsecs

`func (o *CommonUpdatableUserParams) HasEffectiveTimeMsecs() bool`

HasEffectiveTimeMsecs returns a boolean if a field has been set.

### SetEffectiveTimeMsecsNil

`func (o *CommonUpdatableUserParams) SetEffectiveTimeMsecsNil(b bool)`

 SetEffectiveTimeMsecsNil sets the value for EffectiveTimeMsecs to be an explicit nil

### UnsetEffectiveTimeMsecs
`func (o *CommonUpdatableUserParams) UnsetEffectiveTimeMsecs()`

UnsetEffectiveTimeMsecs ensures that no value is present for EffectiveTimeMsecs, not even an explicit nil
### GetExpiryTimeMsecs

`func (o *CommonUpdatableUserParams) GetExpiryTimeMsecs() int64`

GetExpiryTimeMsecs returns the ExpiryTimeMsecs field if non-nil, zero value otherwise.

### GetExpiryTimeMsecsOk

`func (o *CommonUpdatableUserParams) GetExpiryTimeMsecsOk() (*int64, bool)`

GetExpiryTimeMsecsOk returns a tuple with the ExpiryTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimeMsecs

`func (o *CommonUpdatableUserParams) SetExpiryTimeMsecs(v int64)`

SetExpiryTimeMsecs sets ExpiryTimeMsecs field to given value.

### HasExpiryTimeMsecs

`func (o *CommonUpdatableUserParams) HasExpiryTimeMsecs() bool`

HasExpiryTimeMsecs returns a boolean if a field has been set.

### SetExpiryTimeMsecsNil

`func (o *CommonUpdatableUserParams) SetExpiryTimeMsecsNil(b bool)`

 SetExpiryTimeMsecsNil sets the value for ExpiryTimeMsecs to be an explicit nil

### UnsetExpiryTimeMsecs
`func (o *CommonUpdatableUserParams) UnsetExpiryTimeMsecs()`

UnsetExpiryTimeMsecs ensures that no value is present for ExpiryTimeMsecs, not even an explicit nil
### GetLocked

`func (o *CommonUpdatableUserParams) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *CommonUpdatableUserParams) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *CommonUpdatableUserParams) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *CommonUpdatableUserParams) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### SetLockedNil

`func (o *CommonUpdatableUserParams) SetLockedNil(b bool)`

 SetLockedNil sets the value for Locked to be an explicit nil

### UnsetLocked
`func (o *CommonUpdatableUserParams) UnsetLocked()`

UnsetLocked ensures that no value is present for Locked, not even an explicit nil
### GetRestricted

`func (o *CommonUpdatableUserParams) GetRestricted() bool`

GetRestricted returns the Restricted field if non-nil, zero value otherwise.

### GetRestrictedOk

`func (o *CommonUpdatableUserParams) GetRestrictedOk() (*bool, bool)`

GetRestrictedOk returns a tuple with the Restricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestricted

`func (o *CommonUpdatableUserParams) SetRestricted(v bool)`

SetRestricted sets Restricted field to given value.

### HasRestricted

`func (o *CommonUpdatableUserParams) HasRestricted() bool`

HasRestricted returns a boolean if a field has been set.

### SetRestrictedNil

`func (o *CommonUpdatableUserParams) SetRestrictedNil(b bool)`

 SetRestrictedNil sets the value for Restricted to be an explicit nil

### UnsetRestricted
`func (o *CommonUpdatableUserParams) UnsetRestricted()`

UnsetRestricted ensures that no value is present for Restricted, not even an explicit nil
### GetRoles

`func (o *CommonUpdatableUserParams) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CommonUpdatableUserParams) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CommonUpdatableUserParams) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *CommonUpdatableUserParams) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *CommonUpdatableUserParams) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *CommonUpdatableUserParams) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


