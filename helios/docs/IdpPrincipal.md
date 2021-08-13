# IdpPrincipal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | Specifies the name of the Principal. | 
**IdpId** | **NullableInt64** | Specifies the IDP of the IDP with which this principal is associated. | 
**PrincipalType** | **NullableString** | Specifies the type of this principal. It can be a &#39;User&#39; or a &#39;Group&#39;. | 
**Roles** | Pointer to **[]string** | Specifies a list of roles associated with this Principal. | [optional] 
**Clusters** | Pointer to **[]string** | Specifies a list of clusters associated with this Principal. They should be in a string with the format &#39;{cluster ID}:{cluster incarnation ID}&#39;. | [optional] 
**IsActive** | Pointer to **NullableBool** | Specifies whether or not this principal is currently active. | [optional] 
**CreatedTimeUsecs** | Pointer to **NullableInt64** | Specifies the timestamp in microseconds since the epoch when this Principal was created. | [optional] [readonly] 
**LastUpdatedTimeUsecs** | Pointer to **NullableInt64** | Specifies the timestamp in microseconds since the epoch when this Principal was updated. | [optional] [readonly] 
**EffectiveTimeUsecs** | Pointer to **NullableInt64** | Specifies the starting timestamp in microseconds since the epoch when this principal will be able to log in. | [optional] 
**ExpiredTimeUsecs** | Pointer to **NullableInt64** | Specifies the timestamp in microseconds since the epoch when this principal will no longer be able to log in. | [optional] 
**Sid** | Pointer to **NullableString** | Specifies the unique SID of the principal. | [optional] [readonly] 
**Profiles** | Pointer to [**[]McmTenantProfile**](McmTenantProfile.md) | Specifies the list of tenant profiles associated to this principal if any. | [optional] 
**TenantAccesses** | Pointer to [**[]TenantAccess**](TenantAccess.md) | Specifies the list of tenant access associated to this principal. | [optional] 

## Methods

### NewIdpPrincipal

`func NewIdpPrincipal(name NullableString, idpId NullableInt64, principalType NullableString, ) *IdpPrincipal`

NewIdpPrincipal instantiates a new IdpPrincipal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpPrincipalWithDefaults

`func NewIdpPrincipalWithDefaults() *IdpPrincipal`

NewIdpPrincipalWithDefaults instantiates a new IdpPrincipal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IdpPrincipal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdpPrincipal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdpPrincipal) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *IdpPrincipal) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IdpPrincipal) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIdpId

`func (o *IdpPrincipal) GetIdpId() int64`

GetIdpId returns the IdpId field if non-nil, zero value otherwise.

### GetIdpIdOk

`func (o *IdpPrincipal) GetIdpIdOk() (*int64, bool)`

GetIdpIdOk returns a tuple with the IdpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpId

`func (o *IdpPrincipal) SetIdpId(v int64)`

SetIdpId sets IdpId field to given value.


### SetIdpIdNil

`func (o *IdpPrincipal) SetIdpIdNil(b bool)`

 SetIdpIdNil sets the value for IdpId to be an explicit nil

### UnsetIdpId
`func (o *IdpPrincipal) UnsetIdpId()`

UnsetIdpId ensures that no value is present for IdpId, not even an explicit nil
### GetPrincipalType

`func (o *IdpPrincipal) GetPrincipalType() string`

GetPrincipalType returns the PrincipalType field if non-nil, zero value otherwise.

### GetPrincipalTypeOk

`func (o *IdpPrincipal) GetPrincipalTypeOk() (*string, bool)`

GetPrincipalTypeOk returns a tuple with the PrincipalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalType

`func (o *IdpPrincipal) SetPrincipalType(v string)`

SetPrincipalType sets PrincipalType field to given value.


### SetPrincipalTypeNil

`func (o *IdpPrincipal) SetPrincipalTypeNil(b bool)`

 SetPrincipalTypeNil sets the value for PrincipalType to be an explicit nil

### UnsetPrincipalType
`func (o *IdpPrincipal) UnsetPrincipalType()`

UnsetPrincipalType ensures that no value is present for PrincipalType, not even an explicit nil
### GetRoles

`func (o *IdpPrincipal) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *IdpPrincipal) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *IdpPrincipal) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *IdpPrincipal) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetClusters

`func (o *IdpPrincipal) GetClusters() []string`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *IdpPrincipal) GetClustersOk() (*[]string, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *IdpPrincipal) SetClusters(v []string)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *IdpPrincipal) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetIsActive

`func (o *IdpPrincipal) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *IdpPrincipal) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *IdpPrincipal) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *IdpPrincipal) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *IdpPrincipal) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *IdpPrincipal) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetCreatedTimeUsecs

`func (o *IdpPrincipal) GetCreatedTimeUsecs() int64`

GetCreatedTimeUsecs returns the CreatedTimeUsecs field if non-nil, zero value otherwise.

### GetCreatedTimeUsecsOk

`func (o *IdpPrincipal) GetCreatedTimeUsecsOk() (*int64, bool)`

GetCreatedTimeUsecsOk returns a tuple with the CreatedTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimeUsecs

`func (o *IdpPrincipal) SetCreatedTimeUsecs(v int64)`

SetCreatedTimeUsecs sets CreatedTimeUsecs field to given value.

### HasCreatedTimeUsecs

`func (o *IdpPrincipal) HasCreatedTimeUsecs() bool`

HasCreatedTimeUsecs returns a boolean if a field has been set.

### SetCreatedTimeUsecsNil

`func (o *IdpPrincipal) SetCreatedTimeUsecsNil(b bool)`

 SetCreatedTimeUsecsNil sets the value for CreatedTimeUsecs to be an explicit nil

### UnsetCreatedTimeUsecs
`func (o *IdpPrincipal) UnsetCreatedTimeUsecs()`

UnsetCreatedTimeUsecs ensures that no value is present for CreatedTimeUsecs, not even an explicit nil
### GetLastUpdatedTimeUsecs

`func (o *IdpPrincipal) GetLastUpdatedTimeUsecs() int64`

GetLastUpdatedTimeUsecs returns the LastUpdatedTimeUsecs field if non-nil, zero value otherwise.

### GetLastUpdatedTimeUsecsOk

`func (o *IdpPrincipal) GetLastUpdatedTimeUsecsOk() (*int64, bool)`

GetLastUpdatedTimeUsecsOk returns a tuple with the LastUpdatedTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTimeUsecs

`func (o *IdpPrincipal) SetLastUpdatedTimeUsecs(v int64)`

SetLastUpdatedTimeUsecs sets LastUpdatedTimeUsecs field to given value.

### HasLastUpdatedTimeUsecs

`func (o *IdpPrincipal) HasLastUpdatedTimeUsecs() bool`

HasLastUpdatedTimeUsecs returns a boolean if a field has been set.

### SetLastUpdatedTimeUsecsNil

`func (o *IdpPrincipal) SetLastUpdatedTimeUsecsNil(b bool)`

 SetLastUpdatedTimeUsecsNil sets the value for LastUpdatedTimeUsecs to be an explicit nil

### UnsetLastUpdatedTimeUsecs
`func (o *IdpPrincipal) UnsetLastUpdatedTimeUsecs()`

UnsetLastUpdatedTimeUsecs ensures that no value is present for LastUpdatedTimeUsecs, not even an explicit nil
### GetEffectiveTimeUsecs

`func (o *IdpPrincipal) GetEffectiveTimeUsecs() int64`

GetEffectiveTimeUsecs returns the EffectiveTimeUsecs field if non-nil, zero value otherwise.

### GetEffectiveTimeUsecsOk

`func (o *IdpPrincipal) GetEffectiveTimeUsecsOk() (*int64, bool)`

GetEffectiveTimeUsecsOk returns a tuple with the EffectiveTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTimeUsecs

`func (o *IdpPrincipal) SetEffectiveTimeUsecs(v int64)`

SetEffectiveTimeUsecs sets EffectiveTimeUsecs field to given value.

### HasEffectiveTimeUsecs

`func (o *IdpPrincipal) HasEffectiveTimeUsecs() bool`

HasEffectiveTimeUsecs returns a boolean if a field has been set.

### SetEffectiveTimeUsecsNil

`func (o *IdpPrincipal) SetEffectiveTimeUsecsNil(b bool)`

 SetEffectiveTimeUsecsNil sets the value for EffectiveTimeUsecs to be an explicit nil

### UnsetEffectiveTimeUsecs
`func (o *IdpPrincipal) UnsetEffectiveTimeUsecs()`

UnsetEffectiveTimeUsecs ensures that no value is present for EffectiveTimeUsecs, not even an explicit nil
### GetExpiredTimeUsecs

`func (o *IdpPrincipal) GetExpiredTimeUsecs() int64`

GetExpiredTimeUsecs returns the ExpiredTimeUsecs field if non-nil, zero value otherwise.

### GetExpiredTimeUsecsOk

`func (o *IdpPrincipal) GetExpiredTimeUsecsOk() (*int64, bool)`

GetExpiredTimeUsecsOk returns a tuple with the ExpiredTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredTimeUsecs

`func (o *IdpPrincipal) SetExpiredTimeUsecs(v int64)`

SetExpiredTimeUsecs sets ExpiredTimeUsecs field to given value.

### HasExpiredTimeUsecs

`func (o *IdpPrincipal) HasExpiredTimeUsecs() bool`

HasExpiredTimeUsecs returns a boolean if a field has been set.

### SetExpiredTimeUsecsNil

`func (o *IdpPrincipal) SetExpiredTimeUsecsNil(b bool)`

 SetExpiredTimeUsecsNil sets the value for ExpiredTimeUsecs to be an explicit nil

### UnsetExpiredTimeUsecs
`func (o *IdpPrincipal) UnsetExpiredTimeUsecs()`

UnsetExpiredTimeUsecs ensures that no value is present for ExpiredTimeUsecs, not even an explicit nil
### GetSid

`func (o *IdpPrincipal) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *IdpPrincipal) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *IdpPrincipal) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *IdpPrincipal) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *IdpPrincipal) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *IdpPrincipal) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetProfiles

`func (o *IdpPrincipal) GetProfiles() []McmTenantProfile`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *IdpPrincipal) GetProfilesOk() (*[]McmTenantProfile, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *IdpPrincipal) SetProfiles(v []McmTenantProfile)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *IdpPrincipal) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetTenantAccesses

`func (o *IdpPrincipal) GetTenantAccesses() []TenantAccess`

GetTenantAccesses returns the TenantAccesses field if non-nil, zero value otherwise.

### GetTenantAccessesOk

`func (o *IdpPrincipal) GetTenantAccessesOk() (*[]TenantAccess, bool)`

GetTenantAccessesOk returns a tuple with the TenantAccesses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantAccesses

`func (o *IdpPrincipal) SetTenantAccesses(v []TenantAccess)`

SetTenantAccesses sets TenantAccesses field to given value.

### HasTenantAccesses

`func (o *IdpPrincipal) HasTenantAccesses() bool`

HasTenantAccesses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


