# UserParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Specifies the description of the User. | [optional] 
**EffectiveTimeMsecs** | Pointer to **NullableInt64** | Specifies the epoch time in milliseconds since when the user can login. | [optional] 
**ExpiryTimeMsecs** | Pointer to **NullableInt64** | Specifies the epoch time in milliseconds when the user expires. Post expiry the user cannot access Cohesity cluster. | [optional] 
**Locked** | Pointer to **NullableBool** | Specifies whether the User is locked. | [optional] 
**Restricted** | Pointer to **NullableBool** | Specifies whether the User is restricted. A restricted user can only view &amp; manage the objects it has permissions to. | [optional] 
**Roles** | Pointer to **[]string** | Specifies the Cohesity roles to associate with the user. The Cohesity roles determine privileges on the Cohesity Cluster for this user. | [optional] 
**CreatedTimeMsecs** | Pointer to **NullableInt64** | Specifies the epoch time in milliseconds when the user account was created. | [optional] [readonly] 
**Domain** | Pointer to **string** | Specifies the domain of the user. For active directories, this is the fully qualified domain name (FQDN). It is &#39;LOCAL&#39; for local users on the Cohesity Cluster. A user is uniquely identified by combination of the username and the domain. | [optional] [readonly] 
**ForcePasswordChange** | Pointer to **NullableBool** | Specifies if the user must change password. | [optional] [readonly] 
**LastLoginTimeMsecs** | Pointer to **NullableInt64** | Specifies the epoch time in milliseconds when the user last logged in successfully. | [optional] [readonly] 
**LastUpdatedTimeMsecs** | Pointer to **NullableInt64** | Specifies the epoch time in milliseconds when the user account was last modified. | [optional] [readonly] 
**LocalUserParams** | Pointer to **map[string]interface{}** | Specifies the LOCAL user properties. This field is required when adding a new LOCAL Cohesity User. | [optional] 
**LockedReason** | Pointer to **NullableString** | Specifies the reason for locking the User. | [optional] [readonly] 
**OtherGroups** | Pointer to **[]string** | Specifies additional groups the User may belong to. | [optional] [readonly] 
**PrimaryGroup** | Pointer to **NullableString** | Specifies the primary group of the User. Primary group is used for file access. | [optional] [readonly] 
**S3AccountParams** | Pointer to **map[string]interface{}** | Specifies the S3 Account parameters of the User. | [optional] 
**Sid** | Pointer to **NullableString** | Specifies the sid of the User. | [optional] [readonly] 
**TenantId** | Pointer to **NullableString** | Specifies the tenant id of the User. | [optional] 
**Username** | Pointer to **string** | Specifies the username. | [optional] [readonly] 

## Methods

### NewUserParams

`func NewUserParams() *UserParams`

NewUserParams instantiates a new UserParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserParamsWithDefaults

`func NewUserParamsWithDefaults() *UserParams`

NewUserParamsWithDefaults instantiates a new UserParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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
### GetExpiryTimeMsecs

`func (o *UserParams) GetExpiryTimeMsecs() int64`

GetExpiryTimeMsecs returns the ExpiryTimeMsecs field if non-nil, zero value otherwise.

### GetExpiryTimeMsecsOk

`func (o *UserParams) GetExpiryTimeMsecsOk() (*int64, bool)`

GetExpiryTimeMsecsOk returns a tuple with the ExpiryTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimeMsecs

`func (o *UserParams) SetExpiryTimeMsecs(v int64)`

SetExpiryTimeMsecs sets ExpiryTimeMsecs field to given value.

### HasExpiryTimeMsecs

`func (o *UserParams) HasExpiryTimeMsecs() bool`

HasExpiryTimeMsecs returns a boolean if a field has been set.

### SetExpiryTimeMsecsNil

`func (o *UserParams) SetExpiryTimeMsecsNil(b bool)`

 SetExpiryTimeMsecsNil sets the value for ExpiryTimeMsecs to be an explicit nil

### UnsetExpiryTimeMsecs
`func (o *UserParams) UnsetExpiryTimeMsecs()`

UnsetExpiryTimeMsecs ensures that no value is present for ExpiryTimeMsecs, not even an explicit nil
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

### HasDomain

`func (o *UserParams) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetForcePasswordChange

`func (o *UserParams) GetForcePasswordChange() bool`

GetForcePasswordChange returns the ForcePasswordChange field if non-nil, zero value otherwise.

### GetForcePasswordChangeOk

`func (o *UserParams) GetForcePasswordChangeOk() (*bool, bool)`

GetForcePasswordChangeOk returns a tuple with the ForcePasswordChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcePasswordChange

`func (o *UserParams) SetForcePasswordChange(v bool)`

SetForcePasswordChange sets ForcePasswordChange field to given value.

### HasForcePasswordChange

`func (o *UserParams) HasForcePasswordChange() bool`

HasForcePasswordChange returns a boolean if a field has been set.

### SetForcePasswordChangeNil

`func (o *UserParams) SetForcePasswordChangeNil(b bool)`

 SetForcePasswordChangeNil sets the value for ForcePasswordChange to be an explicit nil

### UnsetForcePasswordChange
`func (o *UserParams) UnsetForcePasswordChange()`

UnsetForcePasswordChange ensures that no value is present for ForcePasswordChange, not even an explicit nil
### GetLastLoginTimeMsecs

`func (o *UserParams) GetLastLoginTimeMsecs() int64`

GetLastLoginTimeMsecs returns the LastLoginTimeMsecs field if non-nil, zero value otherwise.

### GetLastLoginTimeMsecsOk

`func (o *UserParams) GetLastLoginTimeMsecsOk() (*int64, bool)`

GetLastLoginTimeMsecsOk returns a tuple with the LastLoginTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginTimeMsecs

`func (o *UserParams) SetLastLoginTimeMsecs(v int64)`

SetLastLoginTimeMsecs sets LastLoginTimeMsecs field to given value.

### HasLastLoginTimeMsecs

`func (o *UserParams) HasLastLoginTimeMsecs() bool`

HasLastLoginTimeMsecs returns a boolean if a field has been set.

### SetLastLoginTimeMsecsNil

`func (o *UserParams) SetLastLoginTimeMsecsNil(b bool)`

 SetLastLoginTimeMsecsNil sets the value for LastLoginTimeMsecs to be an explicit nil

### UnsetLastLoginTimeMsecs
`func (o *UserParams) UnsetLastLoginTimeMsecs()`

UnsetLastLoginTimeMsecs ensures that no value is present for LastLoginTimeMsecs, not even an explicit nil
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
### GetLocalUserParams

`func (o *UserParams) GetLocalUserParams() map[string]interface{}`

GetLocalUserParams returns the LocalUserParams field if non-nil, zero value otherwise.

### GetLocalUserParamsOk

`func (o *UserParams) GetLocalUserParamsOk() (*map[string]interface{}, bool)`

GetLocalUserParamsOk returns a tuple with the LocalUserParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalUserParams

`func (o *UserParams) SetLocalUserParams(v map[string]interface{})`

SetLocalUserParams sets LocalUserParams field to given value.

### HasLocalUserParams

`func (o *UserParams) HasLocalUserParams() bool`

HasLocalUserParams returns a boolean if a field has been set.

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
### GetS3AccountParams

`func (o *UserParams) GetS3AccountParams() map[string]interface{}`

GetS3AccountParams returns the S3AccountParams field if non-nil, zero value otherwise.

### GetS3AccountParamsOk

`func (o *UserParams) GetS3AccountParamsOk() (*map[string]interface{}, bool)`

GetS3AccountParamsOk returns a tuple with the S3AccountParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3AccountParams

`func (o *UserParams) SetS3AccountParams(v map[string]interface{})`

SetS3AccountParams sets S3AccountParams field to given value.

### HasS3AccountParams

`func (o *UserParams) HasS3AccountParams() bool`

HasS3AccountParams returns a boolean if a field has been set.

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

### HasUsername

`func (o *UserParams) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


