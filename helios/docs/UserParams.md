# UserParams

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
**Sid** | Pointer to **NullableString** | Specifies the sid of the User. | [optional] 
**CreatedTimeMsecs** | Pointer to **NullableInt64** | Specifies the epoch time in milliseconds when the user account was created. | [optional] [readonly] 
**LastUpdatedTimeMsecs** | Pointer to **NullableInt64** | Specifies the epoch time in milliseconds when the user account was last modified. | [optional] [readonly] 
**TenantId** | Pointer to **NullableString** | Specifies the tenant id of the User. | [optional] 
**S3AccountParams** | Pointer to [**S3AccountParams**](S3AccountParams.md) | Specifies the S3 Account parameters of the User. | [optional] 
**LockedReason** | Pointer to **NullableString** | Specifies the reason for locking the User. | [optional] [readonly] 

## Methods

### NewUserParams

`func NewUserParams(username NullableString, domain NullableString, password NullableString, ) *UserParams`

NewUserParams instantiates a new UserParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserParamsWithDefaults

`func NewUserParamsWithDefaults() *UserParams`

NewUserParamsWithDefaults instantiates a new UserParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UserParams) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserParams) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserParams) SetUsername(v string)`

SetUsername sets Username field to given value.


### SetUsernameNil

`func (o *UserParams) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *UserParams) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetDomain

`func (o *UserParams) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *UserParams) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *UserParams) SetDomain(v string)`

SetDomain sets Domain field to given value.


### SetDomainNil

`func (o *UserParams) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *UserParams) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetPassword

`func (o *UserParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserParams) SetPassword(v string)`

SetPassword sets Password field to given value.


### SetPasswordNil

`func (o *UserParams) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *UserParams) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetDescription

`func (o *UserParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UserParams) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UserParams) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEmail

`func (o *UserParams) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserParams) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserParams) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserParams) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *UserParams) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UserParams) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetRoles

`func (o *UserParams) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserParams) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserParams) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UserParams) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *UserParams) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *UserParams) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetPrimaryGroup

`func (o *UserParams) GetPrimaryGroup() string`

GetPrimaryGroup returns the PrimaryGroup field if non-nil, zero value otherwise.

### GetPrimaryGroupOk

`func (o *UserParams) GetPrimaryGroupOk() (*string, bool)`

GetPrimaryGroupOk returns a tuple with the PrimaryGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryGroup

`func (o *UserParams) SetPrimaryGroup(v string)`

SetPrimaryGroup sets PrimaryGroup field to given value.

### HasPrimaryGroup

`func (o *UserParams) HasPrimaryGroup() bool`

HasPrimaryGroup returns a boolean if a field has been set.

### SetPrimaryGroupNil

`func (o *UserParams) SetPrimaryGroupNil(b bool)`

 SetPrimaryGroupNil sets the value for PrimaryGroup to be an explicit nil

### UnsetPrimaryGroup
`func (o *UserParams) UnsetPrimaryGroup()`

UnsetPrimaryGroup ensures that no value is present for PrimaryGroup, not even an explicit nil
### GetOtherGroups

`func (o *UserParams) GetOtherGroups() []string`

GetOtherGroups returns the OtherGroups field if non-nil, zero value otherwise.

### GetOtherGroupsOk

`func (o *UserParams) GetOtherGroupsOk() (*[]string, bool)`

GetOtherGroupsOk returns a tuple with the OtherGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherGroups

`func (o *UserParams) SetOtherGroups(v []string)`

SetOtherGroups sets OtherGroups field to given value.

### HasOtherGroups

`func (o *UserParams) HasOtherGroups() bool`

HasOtherGroups returns a boolean if a field has been set.

### SetOtherGroupsNil

`func (o *UserParams) SetOtherGroupsNil(b bool)`

 SetOtherGroupsNil sets the value for OtherGroups to be an explicit nil

### UnsetOtherGroups
`func (o *UserParams) UnsetOtherGroups()`

UnsetOtherGroups ensures that no value is present for OtherGroups, not even an explicit nil
### GetRestricted

`func (o *UserParams) GetRestricted() bool`

GetRestricted returns the Restricted field if non-nil, zero value otherwise.

### GetRestrictedOk

`func (o *UserParams) GetRestrictedOk() (*bool, bool)`

GetRestrictedOk returns a tuple with the Restricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestricted

`func (o *UserParams) SetRestricted(v bool)`

SetRestricted sets Restricted field to given value.

### HasRestricted

`func (o *UserParams) HasRestricted() bool`

HasRestricted returns a boolean if a field has been set.

### SetRestrictedNil

`func (o *UserParams) SetRestrictedNil(b bool)`

 SetRestrictedNil sets the value for Restricted to be an explicit nil

### UnsetRestricted
`func (o *UserParams) UnsetRestricted()`

UnsetRestricted ensures that no value is present for Restricted, not even an explicit nil
### GetEffectiveTimeMsecs

`func (o *UserParams) GetEffectiveTimeMsecs() int64`

GetEffectiveTimeMsecs returns the EffectiveTimeMsecs field if non-nil, zero value otherwise.

### GetEffectiveTimeMsecsOk

`func (o *UserParams) GetEffectiveTimeMsecsOk() (*int64, bool)`

GetEffectiveTimeMsecsOk returns a tuple with the EffectiveTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTimeMsecs

`func (o *UserParams) SetEffectiveTimeMsecs(v int64)`

SetEffectiveTimeMsecs sets EffectiveTimeMsecs field to given value.

### HasEffectiveTimeMsecs

`func (o *UserParams) HasEffectiveTimeMsecs() bool`

HasEffectiveTimeMsecs returns a boolean if a field has been set.

### SetEffectiveTimeMsecsNil

`func (o *UserParams) SetEffectiveTimeMsecsNil(b bool)`

 SetEffectiveTimeMsecsNil sets the value for EffectiveTimeMsecs to be an explicit nil

### UnsetEffectiveTimeMsecs
`func (o *UserParams) UnsetEffectiveTimeMsecs()`

UnsetEffectiveTimeMsecs ensures that no value is present for EffectiveTimeMsecs, not even an explicit nil
### GetExpiredTimeMsecs

`func (o *UserParams) GetExpiredTimeMsecs() int64`

GetExpiredTimeMsecs returns the ExpiredTimeMsecs field if non-nil, zero value otherwise.

### GetExpiredTimeMsecsOk

`func (o *UserParams) GetExpiredTimeMsecsOk() (*int64, bool)`

GetExpiredTimeMsecsOk returns a tuple with the ExpiredTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredTimeMsecs

`func (o *UserParams) SetExpiredTimeMsecs(v int64)`

SetExpiredTimeMsecs sets ExpiredTimeMsecs field to given value.

### HasExpiredTimeMsecs

`func (o *UserParams) HasExpiredTimeMsecs() bool`

HasExpiredTimeMsecs returns a boolean if a field has been set.

### SetExpiredTimeMsecsNil

`func (o *UserParams) SetExpiredTimeMsecsNil(b bool)`

 SetExpiredTimeMsecsNil sets the value for ExpiredTimeMsecs to be an explicit nil

### UnsetExpiredTimeMsecs
`func (o *UserParams) UnsetExpiredTimeMsecs()`

UnsetExpiredTimeMsecs ensures that no value is present for ExpiredTimeMsecs, not even an explicit nil
### GetLocked

`func (o *UserParams) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *UserParams) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *UserParams) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *UserParams) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### SetLockedNil

`func (o *UserParams) SetLockedNil(b bool)`

 SetLockedNil sets the value for Locked to be an explicit nil

### UnsetLocked
`func (o *UserParams) UnsetLocked()`

UnsetLocked ensures that no value is present for Locked, not even an explicit nil
### GetSid

`func (o *UserParams) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *UserParams) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *UserParams) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *UserParams) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *UserParams) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *UserParams) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetCreatedTimeMsecs

`func (o *UserParams) GetCreatedTimeMsecs() int64`

GetCreatedTimeMsecs returns the CreatedTimeMsecs field if non-nil, zero value otherwise.

### GetCreatedTimeMsecsOk

`func (o *UserParams) GetCreatedTimeMsecsOk() (*int64, bool)`

GetCreatedTimeMsecsOk returns a tuple with the CreatedTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimeMsecs

`func (o *UserParams) SetCreatedTimeMsecs(v int64)`

SetCreatedTimeMsecs sets CreatedTimeMsecs field to given value.

### HasCreatedTimeMsecs

`func (o *UserParams) HasCreatedTimeMsecs() bool`

HasCreatedTimeMsecs returns a boolean if a field has been set.

### SetCreatedTimeMsecsNil

`func (o *UserParams) SetCreatedTimeMsecsNil(b bool)`

 SetCreatedTimeMsecsNil sets the value for CreatedTimeMsecs to be an explicit nil

### UnsetCreatedTimeMsecs
`func (o *UserParams) UnsetCreatedTimeMsecs()`

UnsetCreatedTimeMsecs ensures that no value is present for CreatedTimeMsecs, not even an explicit nil
### GetLastUpdatedTimeMsecs

`func (o *UserParams) GetLastUpdatedTimeMsecs() int64`

GetLastUpdatedTimeMsecs returns the LastUpdatedTimeMsecs field if non-nil, zero value otherwise.

### GetLastUpdatedTimeMsecsOk

`func (o *UserParams) GetLastUpdatedTimeMsecsOk() (*int64, bool)`

GetLastUpdatedTimeMsecsOk returns a tuple with the LastUpdatedTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTimeMsecs

`func (o *UserParams) SetLastUpdatedTimeMsecs(v int64)`

SetLastUpdatedTimeMsecs sets LastUpdatedTimeMsecs field to given value.

### HasLastUpdatedTimeMsecs

`func (o *UserParams) HasLastUpdatedTimeMsecs() bool`

HasLastUpdatedTimeMsecs returns a boolean if a field has been set.

### SetLastUpdatedTimeMsecsNil

`func (o *UserParams) SetLastUpdatedTimeMsecsNil(b bool)`

 SetLastUpdatedTimeMsecsNil sets the value for LastUpdatedTimeMsecs to be an explicit nil

### UnsetLastUpdatedTimeMsecs
`func (o *UserParams) UnsetLastUpdatedTimeMsecs()`

UnsetLastUpdatedTimeMsecs ensures that no value is present for LastUpdatedTimeMsecs, not even an explicit nil
### GetTenantId

`func (o *UserParams) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *UserParams) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *UserParams) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *UserParams) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *UserParams) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *UserParams) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetS3AccountParams

`func (o *UserParams) GetS3AccountParams() S3AccountParams`

GetS3AccountParams returns the S3AccountParams field if non-nil, zero value otherwise.

### GetS3AccountParamsOk

`func (o *UserParams) GetS3AccountParamsOk() (*S3AccountParams, bool)`

GetS3AccountParamsOk returns a tuple with the S3AccountParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3AccountParams

`func (o *UserParams) SetS3AccountParams(v S3AccountParams)`

SetS3AccountParams sets S3AccountParams field to given value.

### HasS3AccountParams

`func (o *UserParams) HasS3AccountParams() bool`

HasS3AccountParams returns a boolean if a field has been set.

### GetLockedReason

`func (o *UserParams) GetLockedReason() string`

GetLockedReason returns the LockedReason field if non-nil, zero value otherwise.

### GetLockedReasonOk

`func (o *UserParams) GetLockedReasonOk() (*string, bool)`

GetLockedReasonOk returns a tuple with the LockedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedReason

`func (o *UserParams) SetLockedReason(v string)`

SetLockedReason sets LockedReason field to given value.

### HasLockedReason

`func (o *UserParams) HasLockedReason() bool`

HasLockedReason returns a boolean if a field has been set.

### SetLockedReasonNil

`func (o *UserParams) SetLockedReasonNil(b bool)`

 SetLockedReasonNil sets the value for LockedReason to be an explicit nil

### UnsetLockedReason
`func (o *UserParams) UnsetLockedReason()`

UnsetLockedReason ensures that no value is present for LockedReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


