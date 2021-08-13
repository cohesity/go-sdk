# CreateUserParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **NullableString** | Specifies the username. | 
**Domain** | **NullableString** | Specifies the domain of the User. | 
**Password** | **NullableString** | Specifies the password of the User. | 
**Description** | Pointer to **NullableString** | Specifies the description of the User. | [optional] 
**Email** | Pointer to **NullableString** | Specifies the email address of the User. | [optional] 
**Roles** | Pointer to **[]string** | Specifies the Roles of the User. | [optional] 
**PrimaryGroup** | Pointer to **NullableString** | Specifies the primary group of the User. Primary group is used for file access. | [optional] 
**OtherGroups** | Pointer to **[]string** | Specifies other groups of the User. | [optional] 
**Restricted** | Pointer to **NullableBool** | Specifies whether the User is restricted. | [optional] 
**EffectiveTimeMsecs** | Pointer to **NullableInt64** | Specifies the epoch time in milliseconds when the user is effective. | [optional] 
**ExpiredTimeMsecs** | Pointer to **NullableInt64** | Specifies the epoch time in milliseconds when the user is expired. | [optional] 
**Locked** | Pointer to **NullableBool** | Specifies whether the User is locked. | [optional] 

## Methods

### NewCreateUserParams

`func NewCreateUserParams(username NullableString, domain NullableString, password NullableString, ) *CreateUserParams`

NewCreateUserParams instantiates a new CreateUserParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserParamsWithDefaults

`func NewCreateUserParamsWithDefaults() *CreateUserParams`

NewCreateUserParamsWithDefaults instantiates a new CreateUserParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *CreateUserParams) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateUserParams) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateUserParams) SetUsername(v string)`

SetUsername sets Username field to given value.


### SetUsernameNil

`func (o *CreateUserParams) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *CreateUserParams) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetDomain

`func (o *CreateUserParams) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CreateUserParams) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CreateUserParams) SetDomain(v string)`

SetDomain sets Domain field to given value.


### SetDomainNil

`func (o *CreateUserParams) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *CreateUserParams) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetPassword

`func (o *CreateUserParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateUserParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateUserParams) SetPassword(v string)`

SetPassword sets Password field to given value.


### SetPasswordNil

`func (o *CreateUserParams) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *CreateUserParams) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetDescription

`func (o *CreateUserParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateUserParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateUserParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateUserParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateUserParams) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateUserParams) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEmail

`func (o *CreateUserParams) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateUserParams) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateUserParams) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateUserParams) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *CreateUserParams) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CreateUserParams) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetRoles

`func (o *CreateUserParams) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CreateUserParams) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CreateUserParams) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *CreateUserParams) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *CreateUserParams) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *CreateUserParams) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetPrimaryGroup

`func (o *CreateUserParams) GetPrimaryGroup() string`

GetPrimaryGroup returns the PrimaryGroup field if non-nil, zero value otherwise.

### GetPrimaryGroupOk

`func (o *CreateUserParams) GetPrimaryGroupOk() (*string, bool)`

GetPrimaryGroupOk returns a tuple with the PrimaryGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryGroup

`func (o *CreateUserParams) SetPrimaryGroup(v string)`

SetPrimaryGroup sets PrimaryGroup field to given value.

### HasPrimaryGroup

`func (o *CreateUserParams) HasPrimaryGroup() bool`

HasPrimaryGroup returns a boolean if a field has been set.

### SetPrimaryGroupNil

`func (o *CreateUserParams) SetPrimaryGroupNil(b bool)`

 SetPrimaryGroupNil sets the value for PrimaryGroup to be an explicit nil

### UnsetPrimaryGroup
`func (o *CreateUserParams) UnsetPrimaryGroup()`

UnsetPrimaryGroup ensures that no value is present for PrimaryGroup, not even an explicit nil
### GetOtherGroups

`func (o *CreateUserParams) GetOtherGroups() []string`

GetOtherGroups returns the OtherGroups field if non-nil, zero value otherwise.

### GetOtherGroupsOk

`func (o *CreateUserParams) GetOtherGroupsOk() (*[]string, bool)`

GetOtherGroupsOk returns a tuple with the OtherGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherGroups

`func (o *CreateUserParams) SetOtherGroups(v []string)`

SetOtherGroups sets OtherGroups field to given value.

### HasOtherGroups

`func (o *CreateUserParams) HasOtherGroups() bool`

HasOtherGroups returns a boolean if a field has been set.

### SetOtherGroupsNil

`func (o *CreateUserParams) SetOtherGroupsNil(b bool)`

 SetOtherGroupsNil sets the value for OtherGroups to be an explicit nil

### UnsetOtherGroups
`func (o *CreateUserParams) UnsetOtherGroups()`

UnsetOtherGroups ensures that no value is present for OtherGroups, not even an explicit nil
### GetRestricted

`func (o *CreateUserParams) GetRestricted() bool`

GetRestricted returns the Restricted field if non-nil, zero value otherwise.

### GetRestrictedOk

`func (o *CreateUserParams) GetRestrictedOk() (*bool, bool)`

GetRestrictedOk returns a tuple with the Restricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestricted

`func (o *CreateUserParams) SetRestricted(v bool)`

SetRestricted sets Restricted field to given value.

### HasRestricted

`func (o *CreateUserParams) HasRestricted() bool`

HasRestricted returns a boolean if a field has been set.

### SetRestrictedNil

`func (o *CreateUserParams) SetRestrictedNil(b bool)`

 SetRestrictedNil sets the value for Restricted to be an explicit nil

### UnsetRestricted
`func (o *CreateUserParams) UnsetRestricted()`

UnsetRestricted ensures that no value is present for Restricted, not even an explicit nil
### GetEffectiveTimeMsecs

`func (o *CreateUserParams) GetEffectiveTimeMsecs() int64`

GetEffectiveTimeMsecs returns the EffectiveTimeMsecs field if non-nil, zero value otherwise.

### GetEffectiveTimeMsecsOk

`func (o *CreateUserParams) GetEffectiveTimeMsecsOk() (*int64, bool)`

GetEffectiveTimeMsecsOk returns a tuple with the EffectiveTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTimeMsecs

`func (o *CreateUserParams) SetEffectiveTimeMsecs(v int64)`

SetEffectiveTimeMsecs sets EffectiveTimeMsecs field to given value.

### HasEffectiveTimeMsecs

`func (o *CreateUserParams) HasEffectiveTimeMsecs() bool`

HasEffectiveTimeMsecs returns a boolean if a field has been set.

### SetEffectiveTimeMsecsNil

`func (o *CreateUserParams) SetEffectiveTimeMsecsNil(b bool)`

 SetEffectiveTimeMsecsNil sets the value for EffectiveTimeMsecs to be an explicit nil

### UnsetEffectiveTimeMsecs
`func (o *CreateUserParams) UnsetEffectiveTimeMsecs()`

UnsetEffectiveTimeMsecs ensures that no value is present for EffectiveTimeMsecs, not even an explicit nil
### GetExpiredTimeMsecs

`func (o *CreateUserParams) GetExpiredTimeMsecs() int64`

GetExpiredTimeMsecs returns the ExpiredTimeMsecs field if non-nil, zero value otherwise.

### GetExpiredTimeMsecsOk

`func (o *CreateUserParams) GetExpiredTimeMsecsOk() (*int64, bool)`

GetExpiredTimeMsecsOk returns a tuple with the ExpiredTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredTimeMsecs

`func (o *CreateUserParams) SetExpiredTimeMsecs(v int64)`

SetExpiredTimeMsecs sets ExpiredTimeMsecs field to given value.

### HasExpiredTimeMsecs

`func (o *CreateUserParams) HasExpiredTimeMsecs() bool`

HasExpiredTimeMsecs returns a boolean if a field has been set.

### SetExpiredTimeMsecsNil

`func (o *CreateUserParams) SetExpiredTimeMsecsNil(b bool)`

 SetExpiredTimeMsecsNil sets the value for ExpiredTimeMsecs to be an explicit nil

### UnsetExpiredTimeMsecs
`func (o *CreateUserParams) UnsetExpiredTimeMsecs()`

UnsetExpiredTimeMsecs ensures that no value is present for ExpiredTimeMsecs, not even an explicit nil
### GetLocked

`func (o *CreateUserParams) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *CreateUserParams) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *CreateUserParams) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *CreateUserParams) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### SetLockedNil

`func (o *CreateUserParams) SetLockedNil(b bool)`

 SetLockedNil sets the value for Locked to be an explicit nil

### UnsetLocked
`func (o *CreateUserParams) UnsetLocked()`

UnsetLocked ensures that no value is present for Locked, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


