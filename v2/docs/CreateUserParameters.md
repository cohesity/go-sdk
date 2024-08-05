# CreateUserParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**S3AccessKeys** | Pointer to [**CreateUserParametersS3AccessKeys**](CreateUserParametersS3AccessKeys.md) |  | [optional] 
**AllowSmbAccessToken** | Pointer to **NullableBool** | Specifies whether the SMB access token is to be set for the user. | [optional] 
**Description** | Pointer to **NullableString** | Specifies the description of the User. | [optional] 
**Domain** | **string** | Specifies the domain of the user. For active directories, this is the fully qualified domain name (FQDN). It is &#39;LOCAL&#39; for local users on the Cohesity Cluster. A user is uniquely identified by combination of the username and the domain. | 
**EffectiveTimeMsecs** | Pointer to **NullableInt64** | Specifies the epoch time in milliseconds since when the user can login. | [optional] 
**ExpiryTimeMsecs** | Pointer to **NullableInt64** | Specifies the epoch time in milliseconds when the user expires. Post expiry the user cannot access Cohesity cluster. | [optional] 
**LocalUserParams** | Pointer to [**CreateUserParametersLocalUserParams**](CreateUserParametersLocalUserParams.md) |  | [optional] 
**Locked** | Pointer to **NullableBool** | Specifies whether the User is locked. | [optional] 
**OtherGroups** | Pointer to **[]string** | Specifies additional groups the User may belong to. | [optional] [readonly] 
**PrimaryGroup** | Pointer to **NullableString** | Specifies the primary group of the User. Primary group is used for file access. | [optional] [readonly] 
**Restricted** | Pointer to **NullableBool** | Specifies whether the User is restricted. A restricted user can only view &amp; manage the objects it has permissions to. | [optional] 
**Roles** | Pointer to **[]string** | Specifies the Cohesity roles to associate with the user. The Cohesity roles determine privileges on the Cohesity Cluster for this user. | [optional] 
**Username** | **string** | Specifies the username. | 

## Methods

### NewCreateUserParameters

`func NewCreateUserParameters(domain string, username string, ) *CreateUserParameters`

NewCreateUserParameters instantiates a new CreateUserParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserParametersWithDefaults

`func NewCreateUserParametersWithDefaults() *CreateUserParameters`

NewCreateUserParametersWithDefaults instantiates a new CreateUserParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetS3AccessKeys

`func (o *CreateUserParameters) GetS3AccessKeys() CreateUserParametersS3AccessKeys`

GetS3AccessKeys returns the S3AccessKeys field if non-nil, zero value otherwise.

### GetS3AccessKeysOk

`func (o *CreateUserParameters) GetS3AccessKeysOk() (*CreateUserParametersS3AccessKeys, bool)`

GetS3AccessKeysOk returns a tuple with the S3AccessKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3AccessKeys

`func (o *CreateUserParameters) SetS3AccessKeys(v CreateUserParametersS3AccessKeys)`

SetS3AccessKeys sets S3AccessKeys field to given value.

### HasS3AccessKeys

`func (o *CreateUserParameters) HasS3AccessKeys() bool`

HasS3AccessKeys returns a boolean if a field has been set.

### GetAllowSmbAccessToken

`func (o *CreateUserParameters) GetAllowSmbAccessToken() bool`

GetAllowSmbAccessToken returns the AllowSmbAccessToken field if non-nil, zero value otherwise.

### GetAllowSmbAccessTokenOk

`func (o *CreateUserParameters) GetAllowSmbAccessTokenOk() (*bool, bool)`

GetAllowSmbAccessTokenOk returns a tuple with the AllowSmbAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSmbAccessToken

`func (o *CreateUserParameters) SetAllowSmbAccessToken(v bool)`

SetAllowSmbAccessToken sets AllowSmbAccessToken field to given value.

### HasAllowSmbAccessToken

`func (o *CreateUserParameters) HasAllowSmbAccessToken() bool`

HasAllowSmbAccessToken returns a boolean if a field has been set.

### SetAllowSmbAccessTokenNil

`func (o *CreateUserParameters) SetAllowSmbAccessTokenNil(b bool)`

 SetAllowSmbAccessTokenNil sets the value for AllowSmbAccessToken to be an explicit nil

### UnsetAllowSmbAccessToken
`func (o *CreateUserParameters) UnsetAllowSmbAccessToken()`

UnsetAllowSmbAccessToken ensures that no value is present for AllowSmbAccessToken, not even an explicit nil
### GetDescription

`func (o *CreateUserParameters) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateUserParameters) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateUserParameters) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateUserParameters) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateUserParameters) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateUserParameters) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDomain

`func (o *CreateUserParameters) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CreateUserParameters) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CreateUserParameters) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetEffectiveTimeMsecs

`func (o *CreateUserParameters) GetEffectiveTimeMsecs() int64`

GetEffectiveTimeMsecs returns the EffectiveTimeMsecs field if non-nil, zero value otherwise.

### GetEffectiveTimeMsecsOk

`func (o *CreateUserParameters) GetEffectiveTimeMsecsOk() (*int64, bool)`

GetEffectiveTimeMsecsOk returns a tuple with the EffectiveTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTimeMsecs

`func (o *CreateUserParameters) SetEffectiveTimeMsecs(v int64)`

SetEffectiveTimeMsecs sets EffectiveTimeMsecs field to given value.

### HasEffectiveTimeMsecs

`func (o *CreateUserParameters) HasEffectiveTimeMsecs() bool`

HasEffectiveTimeMsecs returns a boolean if a field has been set.

### SetEffectiveTimeMsecsNil

`func (o *CreateUserParameters) SetEffectiveTimeMsecsNil(b bool)`

 SetEffectiveTimeMsecsNil sets the value for EffectiveTimeMsecs to be an explicit nil

### UnsetEffectiveTimeMsecs
`func (o *CreateUserParameters) UnsetEffectiveTimeMsecs()`

UnsetEffectiveTimeMsecs ensures that no value is present for EffectiveTimeMsecs, not even an explicit nil
### GetExpiryTimeMsecs

`func (o *CreateUserParameters) GetExpiryTimeMsecs() int64`

GetExpiryTimeMsecs returns the ExpiryTimeMsecs field if non-nil, zero value otherwise.

### GetExpiryTimeMsecsOk

`func (o *CreateUserParameters) GetExpiryTimeMsecsOk() (*int64, bool)`

GetExpiryTimeMsecsOk returns a tuple with the ExpiryTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimeMsecs

`func (o *CreateUserParameters) SetExpiryTimeMsecs(v int64)`

SetExpiryTimeMsecs sets ExpiryTimeMsecs field to given value.

### HasExpiryTimeMsecs

`func (o *CreateUserParameters) HasExpiryTimeMsecs() bool`

HasExpiryTimeMsecs returns a boolean if a field has been set.

### SetExpiryTimeMsecsNil

`func (o *CreateUserParameters) SetExpiryTimeMsecsNil(b bool)`

 SetExpiryTimeMsecsNil sets the value for ExpiryTimeMsecs to be an explicit nil

### UnsetExpiryTimeMsecs
`func (o *CreateUserParameters) UnsetExpiryTimeMsecs()`

UnsetExpiryTimeMsecs ensures that no value is present for ExpiryTimeMsecs, not even an explicit nil
### GetLocalUserParams

`func (o *CreateUserParameters) GetLocalUserParams() CreateUserParametersLocalUserParams`

GetLocalUserParams returns the LocalUserParams field if non-nil, zero value otherwise.

### GetLocalUserParamsOk

`func (o *CreateUserParameters) GetLocalUserParamsOk() (*CreateUserParametersLocalUserParams, bool)`

GetLocalUserParamsOk returns a tuple with the LocalUserParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalUserParams

`func (o *CreateUserParameters) SetLocalUserParams(v CreateUserParametersLocalUserParams)`

SetLocalUserParams sets LocalUserParams field to given value.

### HasLocalUserParams

`func (o *CreateUserParameters) HasLocalUserParams() bool`

HasLocalUserParams returns a boolean if a field has been set.

### GetLocked

`func (o *CreateUserParameters) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *CreateUserParameters) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *CreateUserParameters) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *CreateUserParameters) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### SetLockedNil

`func (o *CreateUserParameters) SetLockedNil(b bool)`

 SetLockedNil sets the value for Locked to be an explicit nil

### UnsetLocked
`func (o *CreateUserParameters) UnsetLocked()`

UnsetLocked ensures that no value is present for Locked, not even an explicit nil
### GetOtherGroups

`func (o *CreateUserParameters) GetOtherGroups() []string`

GetOtherGroups returns the OtherGroups field if non-nil, zero value otherwise.

### GetOtherGroupsOk

`func (o *CreateUserParameters) GetOtherGroupsOk() (*[]string, bool)`

GetOtherGroupsOk returns a tuple with the OtherGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherGroups

`func (o *CreateUserParameters) SetOtherGroups(v []string)`

SetOtherGroups sets OtherGroups field to given value.

### HasOtherGroups

`func (o *CreateUserParameters) HasOtherGroups() bool`

HasOtherGroups returns a boolean if a field has been set.

### GetPrimaryGroup

`func (o *CreateUserParameters) GetPrimaryGroup() string`

GetPrimaryGroup returns the PrimaryGroup field if non-nil, zero value otherwise.

### GetPrimaryGroupOk

`func (o *CreateUserParameters) GetPrimaryGroupOk() (*string, bool)`

GetPrimaryGroupOk returns a tuple with the PrimaryGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryGroup

`func (o *CreateUserParameters) SetPrimaryGroup(v string)`

SetPrimaryGroup sets PrimaryGroup field to given value.

### HasPrimaryGroup

`func (o *CreateUserParameters) HasPrimaryGroup() bool`

HasPrimaryGroup returns a boolean if a field has been set.

### SetPrimaryGroupNil

`func (o *CreateUserParameters) SetPrimaryGroupNil(b bool)`

 SetPrimaryGroupNil sets the value for PrimaryGroup to be an explicit nil

### UnsetPrimaryGroup
`func (o *CreateUserParameters) UnsetPrimaryGroup()`

UnsetPrimaryGroup ensures that no value is present for PrimaryGroup, not even an explicit nil
### GetRestricted

`func (o *CreateUserParameters) GetRestricted() bool`

GetRestricted returns the Restricted field if non-nil, zero value otherwise.

### GetRestrictedOk

`func (o *CreateUserParameters) GetRestrictedOk() (*bool, bool)`

GetRestrictedOk returns a tuple with the Restricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestricted

`func (o *CreateUserParameters) SetRestricted(v bool)`

SetRestricted sets Restricted field to given value.

### HasRestricted

`func (o *CreateUserParameters) HasRestricted() bool`

HasRestricted returns a boolean if a field has been set.

### SetRestrictedNil

`func (o *CreateUserParameters) SetRestrictedNil(b bool)`

 SetRestrictedNil sets the value for Restricted to be an explicit nil

### UnsetRestricted
`func (o *CreateUserParameters) UnsetRestricted()`

UnsetRestricted ensures that no value is present for Restricted, not even an explicit nil
### GetRoles

`func (o *CreateUserParameters) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CreateUserParameters) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CreateUserParameters) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *CreateUserParameters) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *CreateUserParameters) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *CreateUserParameters) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetUsername

`func (o *CreateUserParameters) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateUserParameters) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateUserParameters) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


