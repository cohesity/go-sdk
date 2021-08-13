# TenantAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **NullableString** | Specifies the Tenant Id of the tenant access. | 
**Roles** | **[]string** | Specifies a list of roles associated with this access. | 
**Clusters** | [**[]McmClusterIdentifier**](McmClusterIdentifier.md) | Specifies the list of clusters accessible by this access. | 
**CreatedTimeMsecs** | Pointer to **NullableInt64** | Specifies the timestamp in milliseconds since the epoch when this access was created. | [optional] [readonly] 
**LastUpdatedTimeMsecs** | Pointer to **NullableInt64** | Specifies the timestamp in milliseconds since the epoch when this access was updated. | [optional] [readonly] 
**IsAccessActive** | Pointer to **NullableBool** | Specifies whether or not this access is currently active. | [optional] [readonly] 
**EffectiveTimeMsecs** | Pointer to **NullableInt64** | Specifies the starting timestamp in milliseconds since the epoch when this access will be able allowed. | [optional] [readonly] 
**ExpiredTimeMsecs** | Pointer to **NullableInt64** | Specifies the timestamp in milliseconds since the epoch when this access will no longer be allowed. | [optional] [readonly] 
**TenantName** | Pointer to **NullableString** | Name of the Tenant. | [optional] [readonly] 
**TenantStatus** | Pointer to **NullableString** | Specifies the Tenant status. | [optional] [readonly] 
**TenantType** | Pointer to **NullableString** | Specifies the type of the tenant. | [optional] [readonly] 

## Methods

### NewTenantAccess

`func NewTenantAccess(tenantId NullableString, roles []string, clusters []McmClusterIdentifier, ) *TenantAccess`

NewTenantAccess instantiates a new TenantAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantAccessWithDefaults

`func NewTenantAccessWithDefaults() *TenantAccess`

NewTenantAccessWithDefaults instantiates a new TenantAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *TenantAccess) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantAccess) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantAccess) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### SetTenantIdNil

`func (o *TenantAccess) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TenantAccess) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetRoles

`func (o *TenantAccess) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *TenantAccess) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *TenantAccess) SetRoles(v []string)`

SetRoles sets Roles field to given value.


### GetClusters

`func (o *TenantAccess) GetClusters() []McmClusterIdentifier`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *TenantAccess) GetClustersOk() (*[]McmClusterIdentifier, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *TenantAccess) SetClusters(v []McmClusterIdentifier)`

SetClusters sets Clusters field to given value.


### GetCreatedTimeMsecs

`func (o *TenantAccess) GetCreatedTimeMsecs() int64`

GetCreatedTimeMsecs returns the CreatedTimeMsecs field if non-nil, zero value otherwise.

### GetCreatedTimeMsecsOk

`func (o *TenantAccess) GetCreatedTimeMsecsOk() (*int64, bool)`

GetCreatedTimeMsecsOk returns a tuple with the CreatedTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimeMsecs

`func (o *TenantAccess) SetCreatedTimeMsecs(v int64)`

SetCreatedTimeMsecs sets CreatedTimeMsecs field to given value.

### HasCreatedTimeMsecs

`func (o *TenantAccess) HasCreatedTimeMsecs() bool`

HasCreatedTimeMsecs returns a boolean if a field has been set.

### SetCreatedTimeMsecsNil

`func (o *TenantAccess) SetCreatedTimeMsecsNil(b bool)`

 SetCreatedTimeMsecsNil sets the value for CreatedTimeMsecs to be an explicit nil

### UnsetCreatedTimeMsecs
`func (o *TenantAccess) UnsetCreatedTimeMsecs()`

UnsetCreatedTimeMsecs ensures that no value is present for CreatedTimeMsecs, not even an explicit nil
### GetLastUpdatedTimeMsecs

`func (o *TenantAccess) GetLastUpdatedTimeMsecs() int64`

GetLastUpdatedTimeMsecs returns the LastUpdatedTimeMsecs field if non-nil, zero value otherwise.

### GetLastUpdatedTimeMsecsOk

`func (o *TenantAccess) GetLastUpdatedTimeMsecsOk() (*int64, bool)`

GetLastUpdatedTimeMsecsOk returns a tuple with the LastUpdatedTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTimeMsecs

`func (o *TenantAccess) SetLastUpdatedTimeMsecs(v int64)`

SetLastUpdatedTimeMsecs sets LastUpdatedTimeMsecs field to given value.

### HasLastUpdatedTimeMsecs

`func (o *TenantAccess) HasLastUpdatedTimeMsecs() bool`

HasLastUpdatedTimeMsecs returns a boolean if a field has been set.

### SetLastUpdatedTimeMsecsNil

`func (o *TenantAccess) SetLastUpdatedTimeMsecsNil(b bool)`

 SetLastUpdatedTimeMsecsNil sets the value for LastUpdatedTimeMsecs to be an explicit nil

### UnsetLastUpdatedTimeMsecs
`func (o *TenantAccess) UnsetLastUpdatedTimeMsecs()`

UnsetLastUpdatedTimeMsecs ensures that no value is present for LastUpdatedTimeMsecs, not even an explicit nil
### GetIsAccessActive

`func (o *TenantAccess) GetIsAccessActive() bool`

GetIsAccessActive returns the IsAccessActive field if non-nil, zero value otherwise.

### GetIsAccessActiveOk

`func (o *TenantAccess) GetIsAccessActiveOk() (*bool, bool)`

GetIsAccessActiveOk returns a tuple with the IsAccessActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccessActive

`func (o *TenantAccess) SetIsAccessActive(v bool)`

SetIsAccessActive sets IsAccessActive field to given value.

### HasIsAccessActive

`func (o *TenantAccess) HasIsAccessActive() bool`

HasIsAccessActive returns a boolean if a field has been set.

### SetIsAccessActiveNil

`func (o *TenantAccess) SetIsAccessActiveNil(b bool)`

 SetIsAccessActiveNil sets the value for IsAccessActive to be an explicit nil

### UnsetIsAccessActive
`func (o *TenantAccess) UnsetIsAccessActive()`

UnsetIsAccessActive ensures that no value is present for IsAccessActive, not even an explicit nil
### GetEffectiveTimeMsecs

`func (o *TenantAccess) GetEffectiveTimeMsecs() int64`

GetEffectiveTimeMsecs returns the EffectiveTimeMsecs field if non-nil, zero value otherwise.

### GetEffectiveTimeMsecsOk

`func (o *TenantAccess) GetEffectiveTimeMsecsOk() (*int64, bool)`

GetEffectiveTimeMsecsOk returns a tuple with the EffectiveTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTimeMsecs

`func (o *TenantAccess) SetEffectiveTimeMsecs(v int64)`

SetEffectiveTimeMsecs sets EffectiveTimeMsecs field to given value.

### HasEffectiveTimeMsecs

`func (o *TenantAccess) HasEffectiveTimeMsecs() bool`

HasEffectiveTimeMsecs returns a boolean if a field has been set.

### SetEffectiveTimeMsecsNil

`func (o *TenantAccess) SetEffectiveTimeMsecsNil(b bool)`

 SetEffectiveTimeMsecsNil sets the value for EffectiveTimeMsecs to be an explicit nil

### UnsetEffectiveTimeMsecs
`func (o *TenantAccess) UnsetEffectiveTimeMsecs()`

UnsetEffectiveTimeMsecs ensures that no value is present for EffectiveTimeMsecs, not even an explicit nil
### GetExpiredTimeMsecs

`func (o *TenantAccess) GetExpiredTimeMsecs() int64`

GetExpiredTimeMsecs returns the ExpiredTimeMsecs field if non-nil, zero value otherwise.

### GetExpiredTimeMsecsOk

`func (o *TenantAccess) GetExpiredTimeMsecsOk() (*int64, bool)`

GetExpiredTimeMsecsOk returns a tuple with the ExpiredTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredTimeMsecs

`func (o *TenantAccess) SetExpiredTimeMsecs(v int64)`

SetExpiredTimeMsecs sets ExpiredTimeMsecs field to given value.

### HasExpiredTimeMsecs

`func (o *TenantAccess) HasExpiredTimeMsecs() bool`

HasExpiredTimeMsecs returns a boolean if a field has been set.

### SetExpiredTimeMsecsNil

`func (o *TenantAccess) SetExpiredTimeMsecsNil(b bool)`

 SetExpiredTimeMsecsNil sets the value for ExpiredTimeMsecs to be an explicit nil

### UnsetExpiredTimeMsecs
`func (o *TenantAccess) UnsetExpiredTimeMsecs()`

UnsetExpiredTimeMsecs ensures that no value is present for ExpiredTimeMsecs, not even an explicit nil
### GetTenantName

`func (o *TenantAccess) GetTenantName() string`

GetTenantName returns the TenantName field if non-nil, zero value otherwise.

### GetTenantNameOk

`func (o *TenantAccess) GetTenantNameOk() (*string, bool)`

GetTenantNameOk returns a tuple with the TenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantName

`func (o *TenantAccess) SetTenantName(v string)`

SetTenantName sets TenantName field to given value.

### HasTenantName

`func (o *TenantAccess) HasTenantName() bool`

HasTenantName returns a boolean if a field has been set.

### SetTenantNameNil

`func (o *TenantAccess) SetTenantNameNil(b bool)`

 SetTenantNameNil sets the value for TenantName to be an explicit nil

### UnsetTenantName
`func (o *TenantAccess) UnsetTenantName()`

UnsetTenantName ensures that no value is present for TenantName, not even an explicit nil
### GetTenantStatus

`func (o *TenantAccess) GetTenantStatus() string`

GetTenantStatus returns the TenantStatus field if non-nil, zero value otherwise.

### GetTenantStatusOk

`func (o *TenantAccess) GetTenantStatusOk() (*string, bool)`

GetTenantStatusOk returns a tuple with the TenantStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantStatus

`func (o *TenantAccess) SetTenantStatus(v string)`

SetTenantStatus sets TenantStatus field to given value.

### HasTenantStatus

`func (o *TenantAccess) HasTenantStatus() bool`

HasTenantStatus returns a boolean if a field has been set.

### SetTenantStatusNil

`func (o *TenantAccess) SetTenantStatusNil(b bool)`

 SetTenantStatusNil sets the value for TenantStatus to be an explicit nil

### UnsetTenantStatus
`func (o *TenantAccess) UnsetTenantStatus()`

UnsetTenantStatus ensures that no value is present for TenantStatus, not even an explicit nil
### GetTenantType

`func (o *TenantAccess) GetTenantType() string`

GetTenantType returns the TenantType field if non-nil, zero value otherwise.

### GetTenantTypeOk

`func (o *TenantAccess) GetTenantTypeOk() (*string, bool)`

GetTenantTypeOk returns a tuple with the TenantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantType

`func (o *TenantAccess) SetTenantType(v string)`

SetTenantType sets TenantType field to given value.

### HasTenantType

`func (o *TenantAccess) HasTenantType() bool`

HasTenantType returns a boolean if a field has been set.

### SetTenantTypeNil

`func (o *TenantAccess) SetTenantTypeNil(b bool)`

 SetTenantTypeNil sets the value for TenantType to be an explicit nil

### UnsetTenantType
`func (o *TenantAccess) UnsetTenantType()`

UnsetTenantType ensures that no value is present for TenantType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


