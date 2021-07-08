# FileDistributionStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **NullableInt64** | Specifies the cluster Id. | [optional] 
**ClusterIncarnationId** | Pointer to **NullableInt64** | Specifies the cluster Incarnation Id. | [optional] 
**EntityId** | Pointer to **NullableInt64** | Specifies the id of the entity for which file distribution stats are computed. | [optional] 
**EntityName** | Pointer to **NullableString** | Specifies the name of the entity for which file distribution stats are computed. | [optional] 
**MetricsList** | Pointer to [**[]FileDistributionMetrics**](FileDistributionMetrics.md) | Specifies the list of file stats for different file extensions. | [optional] 

## Methods

### NewFileDistributionStats

`func NewFileDistributionStats() *FileDistributionStats`

NewFileDistributionStats instantiates a new FileDistributionStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileDistributionStatsWithDefaults

`func NewFileDistributionStatsWithDefaults() *FileDistributionStats`

NewFileDistributionStatsWithDefaults instantiates a new FileDistributionStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *FileDistributionStats) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *FileDistributionStats) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *FileDistributionStats) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *FileDistributionStats) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *FileDistributionStats) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *FileDistributionStats) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterIncarnationId

`func (o *FileDistributionStats) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *FileDistributionStats) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *FileDistributionStats) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.

### HasClusterIncarnationId

`func (o *FileDistributionStats) HasClusterIncarnationId() bool`

HasClusterIncarnationId returns a boolean if a field has been set.

### SetClusterIncarnationIdNil

`func (o *FileDistributionStats) SetClusterIncarnationIdNil(b bool)`

 SetClusterIncarnationIdNil sets the value for ClusterIncarnationId to be an explicit nil

### UnsetClusterIncarnationId
`func (o *FileDistributionStats) UnsetClusterIncarnationId()`

UnsetClusterIncarnationId ensures that no value is present for ClusterIncarnationId, not even an explicit nil
### GetEntityId

`func (o *FileDistributionStats) GetEntityId() int64`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *FileDistributionStats) GetEntityIdOk() (*int64, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *FileDistributionStats) SetEntityId(v int64)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *FileDistributionStats) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### SetEntityIdNil

`func (o *FileDistributionStats) SetEntityIdNil(b bool)`

 SetEntityIdNil sets the value for EntityId to be an explicit nil

### UnsetEntityId
`func (o *FileDistributionStats) UnsetEntityId()`

UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
### GetEntityName

`func (o *FileDistributionStats) GetEntityName() string`

GetEntityName returns the EntityName field if non-nil, zero value otherwise.

### GetEntityNameOk

`func (o *FileDistributionStats) GetEntityNameOk() (*string, bool)`

GetEntityNameOk returns a tuple with the EntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityName

`func (o *FileDistributionStats) SetEntityName(v string)`

SetEntityName sets EntityName field to given value.

### HasEntityName

`func (o *FileDistributionStats) HasEntityName() bool`

HasEntityName returns a boolean if a field has been set.

### SetEntityNameNil

`func (o *FileDistributionStats) SetEntityNameNil(b bool)`

 SetEntityNameNil sets the value for EntityName to be an explicit nil

### UnsetEntityName
`func (o *FileDistributionStats) UnsetEntityName()`

UnsetEntityName ensures that no value is present for EntityName, not even an explicit nil
### GetMetricsList

`func (o *FileDistributionStats) GetMetricsList() []FileDistributionMetrics`

GetMetricsList returns the MetricsList field if non-nil, zero value otherwise.

### GetMetricsListOk

`func (o *FileDistributionStats) GetMetricsListOk() (*[]FileDistributionMetrics, bool)`

GetMetricsListOk returns a tuple with the MetricsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsList

`func (o *FileDistributionStats) SetMetricsList(v []FileDistributionMetrics)`

SetMetricsList sets MetricsList field to given value.

### HasMetricsList

`func (o *FileDistributionStats) HasMetricsList() bool`

HasMetricsList returns a boolean if a field has been set.

### SetMetricsListNil

`func (o *FileDistributionStats) SetMetricsListNil(b bool)`

 SetMetricsListNil sets the value for MetricsList to be an explicit nil

### UnsetMetricsList
`func (o *FileDistributionStats) UnsetMetricsList()`

UnsetMetricsList ensures that no value is present for MetricsList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


