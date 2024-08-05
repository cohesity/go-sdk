# FilesStatsForEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **NullableInt64** | Specifies the cluster id of the entity. | [optional] 
**ClusterIncarnationId** | Pointer to **NullableInt64** | Specifies the cluster incarnation id of the entity. | [optional] 
**EntityId** | Pointer to **NullableInt64** | Specifies the entity id. | [optional] 
**EntityName** | Pointer to **NullableString** | Specifies the entity name. | [optional] 
**EntityType** | Pointer to **NullableString** | Specifies the entity type. | [optional] 
**FilesStats** | Pointer to [**[]FileStats**](FileStats.md) | Specifies a list of files stats for the entity. | [optional] 

## Methods

### NewFilesStatsForEntity

`func NewFilesStatsForEntity() *FilesStatsForEntity`

NewFilesStatsForEntity instantiates a new FilesStatsForEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesStatsForEntityWithDefaults

`func NewFilesStatsForEntityWithDefaults() *FilesStatsForEntity`

NewFilesStatsForEntityWithDefaults instantiates a new FilesStatsForEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *FilesStatsForEntity) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *FilesStatsForEntity) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *FilesStatsForEntity) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *FilesStatsForEntity) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *FilesStatsForEntity) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *FilesStatsForEntity) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterIncarnationId

`func (o *FilesStatsForEntity) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *FilesStatsForEntity) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *FilesStatsForEntity) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.

### HasClusterIncarnationId

`func (o *FilesStatsForEntity) HasClusterIncarnationId() bool`

HasClusterIncarnationId returns a boolean if a field has been set.

### SetClusterIncarnationIdNil

`func (o *FilesStatsForEntity) SetClusterIncarnationIdNil(b bool)`

 SetClusterIncarnationIdNil sets the value for ClusterIncarnationId to be an explicit nil

### UnsetClusterIncarnationId
`func (o *FilesStatsForEntity) UnsetClusterIncarnationId()`

UnsetClusterIncarnationId ensures that no value is present for ClusterIncarnationId, not even an explicit nil
### GetEntityId

`func (o *FilesStatsForEntity) GetEntityId() int64`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *FilesStatsForEntity) GetEntityIdOk() (*int64, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *FilesStatsForEntity) SetEntityId(v int64)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *FilesStatsForEntity) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### SetEntityIdNil

`func (o *FilesStatsForEntity) SetEntityIdNil(b bool)`

 SetEntityIdNil sets the value for EntityId to be an explicit nil

### UnsetEntityId
`func (o *FilesStatsForEntity) UnsetEntityId()`

UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
### GetEntityName

`func (o *FilesStatsForEntity) GetEntityName() string`

GetEntityName returns the EntityName field if non-nil, zero value otherwise.

### GetEntityNameOk

`func (o *FilesStatsForEntity) GetEntityNameOk() (*string, bool)`

GetEntityNameOk returns a tuple with the EntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityName

`func (o *FilesStatsForEntity) SetEntityName(v string)`

SetEntityName sets EntityName field to given value.

### HasEntityName

`func (o *FilesStatsForEntity) HasEntityName() bool`

HasEntityName returns a boolean if a field has been set.

### SetEntityNameNil

`func (o *FilesStatsForEntity) SetEntityNameNil(b bool)`

 SetEntityNameNil sets the value for EntityName to be an explicit nil

### UnsetEntityName
`func (o *FilesStatsForEntity) UnsetEntityName()`

UnsetEntityName ensures that no value is present for EntityName, not even an explicit nil
### GetEntityType

`func (o *FilesStatsForEntity) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *FilesStatsForEntity) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *FilesStatsForEntity) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *FilesStatsForEntity) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### SetEntityTypeNil

`func (o *FilesStatsForEntity) SetEntityTypeNil(b bool)`

 SetEntityTypeNil sets the value for EntityType to be an explicit nil

### UnsetEntityType
`func (o *FilesStatsForEntity) UnsetEntityType()`

UnsetEntityType ensures that no value is present for EntityType, not even an explicit nil
### GetFilesStats

`func (o *FilesStatsForEntity) GetFilesStats() []FileStats`

GetFilesStats returns the FilesStats field if non-nil, zero value otherwise.

### GetFilesStatsOk

`func (o *FilesStatsForEntity) GetFilesStatsOk() (*[]FileStats, bool)`

GetFilesStatsOk returns a tuple with the FilesStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesStats

`func (o *FilesStatsForEntity) SetFilesStats(v []FileStats)`

SetFilesStats sets FilesStats field to given value.

### HasFilesStats

`func (o *FilesStatsForEntity) HasFilesStats() bool`

HasFilesStats returns a boolean if a field has been set.

### SetFilesStatsNil

`func (o *FilesStatsForEntity) SetFilesStatsNil(b bool)`

 SetFilesStatsNil sets the value for FilesStats to be an explicit nil

### UnsetFilesStats
`func (o *FilesStatsForEntity) UnsetFilesStats()`

UnsetFilesStats ensures that no value is present for FilesStats, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


