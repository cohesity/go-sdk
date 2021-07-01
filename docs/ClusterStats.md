# ClusterStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudUsagePerfStats** | Pointer to [**NullableUsageAndPerformanceStats**](UsageAndPerformanceStats.md) | Provides usage and performance statistics for the remote data stored on a Cloud Tier by the Cohesity Cluster. | [optional] 
**DataReductionRatio** | Pointer to **NullableFloat64** | Provides the ratio of Cluster Logical Data (totalLogicalUsageBytes) Managed to Cluster Storage Used (totalPhysicalUsageBytes) | [optional] 
**DataUsageStats** | Pointer to [**DataUsageStats**](DataUsageStats.md) |  | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies the id of the Cohesity Cluster. | [optional] 
**LocalUsagePerfStats** | Pointer to [**NullableUsageAndPerformanceStats**](UsageAndPerformanceStats.md) | Provides usage and performance statistics for local data stored directly on the Cohesity Cluster. | [optional] 
**LogicalStats** | Pointer to [**NullableLogicalStats**](LogicalStats.md) | Specifies the total logical data size of all the local and Cloud Tier data stored by the Cohesity Cluster before the data is reduced by change-block tracking, compression and deduplication. The size of the data if the data is fully hydrated or expanded. | [optional] 
**UsagePerfStats** | Pointer to [**NullableUsageAndPerformanceStats**](UsageAndPerformanceStats.md) | Provides usage and performance statistics about the local data stored directly on the Cohesity Cluster and the remote data stored on a Cloud Tier for the Cohesity Cluster. | [optional] 

## Methods

### NewClusterStats

`func NewClusterStats() *ClusterStats`

NewClusterStats instantiates a new ClusterStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterStatsWithDefaults

`func NewClusterStatsWithDefaults() *ClusterStats`

NewClusterStatsWithDefaults instantiates a new ClusterStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudUsagePerfStats

`func (o *ClusterStats) GetCloudUsagePerfStats() UsageAndPerformanceStats`

GetCloudUsagePerfStats returns the CloudUsagePerfStats field if non-nil, zero value otherwise.

### GetCloudUsagePerfStatsOk

`func (o *ClusterStats) GetCloudUsagePerfStatsOk() (*UsageAndPerformanceStats, bool)`

GetCloudUsagePerfStatsOk returns a tuple with the CloudUsagePerfStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudUsagePerfStats

`func (o *ClusterStats) SetCloudUsagePerfStats(v UsageAndPerformanceStats)`

SetCloudUsagePerfStats sets CloudUsagePerfStats field to given value.

### HasCloudUsagePerfStats

`func (o *ClusterStats) HasCloudUsagePerfStats() bool`

HasCloudUsagePerfStats returns a boolean if a field has been set.

### SetCloudUsagePerfStatsNil

`func (o *ClusterStats) SetCloudUsagePerfStatsNil(b bool)`

 SetCloudUsagePerfStatsNil sets the value for CloudUsagePerfStats to be an explicit nil

### UnsetCloudUsagePerfStats
`func (o *ClusterStats) UnsetCloudUsagePerfStats()`

UnsetCloudUsagePerfStats ensures that no value is present for CloudUsagePerfStats, not even an explicit nil
### GetDataReductionRatio

`func (o *ClusterStats) GetDataReductionRatio() float64`

GetDataReductionRatio returns the DataReductionRatio field if non-nil, zero value otherwise.

### GetDataReductionRatioOk

`func (o *ClusterStats) GetDataReductionRatioOk() (*float64, bool)`

GetDataReductionRatioOk returns a tuple with the DataReductionRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReductionRatio

`func (o *ClusterStats) SetDataReductionRatio(v float64)`

SetDataReductionRatio sets DataReductionRatio field to given value.

### HasDataReductionRatio

`func (o *ClusterStats) HasDataReductionRatio() bool`

HasDataReductionRatio returns a boolean if a field has been set.

### SetDataReductionRatioNil

`func (o *ClusterStats) SetDataReductionRatioNil(b bool)`

 SetDataReductionRatioNil sets the value for DataReductionRatio to be an explicit nil

### UnsetDataReductionRatio
`func (o *ClusterStats) UnsetDataReductionRatio()`

UnsetDataReductionRatio ensures that no value is present for DataReductionRatio, not even an explicit nil
### GetDataUsageStats

`func (o *ClusterStats) GetDataUsageStats() DataUsageStats`

GetDataUsageStats returns the DataUsageStats field if non-nil, zero value otherwise.

### GetDataUsageStatsOk

`func (o *ClusterStats) GetDataUsageStatsOk() (*DataUsageStats, bool)`

GetDataUsageStatsOk returns a tuple with the DataUsageStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataUsageStats

`func (o *ClusterStats) SetDataUsageStats(v DataUsageStats)`

SetDataUsageStats sets DataUsageStats field to given value.

### HasDataUsageStats

`func (o *ClusterStats) HasDataUsageStats() bool`

HasDataUsageStats returns a boolean if a field has been set.

### GetId

`func (o *ClusterStats) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterStats) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterStats) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterStats) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ClusterStats) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ClusterStats) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLocalUsagePerfStats

`func (o *ClusterStats) GetLocalUsagePerfStats() UsageAndPerformanceStats`

GetLocalUsagePerfStats returns the LocalUsagePerfStats field if non-nil, zero value otherwise.

### GetLocalUsagePerfStatsOk

`func (o *ClusterStats) GetLocalUsagePerfStatsOk() (*UsageAndPerformanceStats, bool)`

GetLocalUsagePerfStatsOk returns a tuple with the LocalUsagePerfStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalUsagePerfStats

`func (o *ClusterStats) SetLocalUsagePerfStats(v UsageAndPerformanceStats)`

SetLocalUsagePerfStats sets LocalUsagePerfStats field to given value.

### HasLocalUsagePerfStats

`func (o *ClusterStats) HasLocalUsagePerfStats() bool`

HasLocalUsagePerfStats returns a boolean if a field has been set.

### SetLocalUsagePerfStatsNil

`func (o *ClusterStats) SetLocalUsagePerfStatsNil(b bool)`

 SetLocalUsagePerfStatsNil sets the value for LocalUsagePerfStats to be an explicit nil

### UnsetLocalUsagePerfStats
`func (o *ClusterStats) UnsetLocalUsagePerfStats()`

UnsetLocalUsagePerfStats ensures that no value is present for LocalUsagePerfStats, not even an explicit nil
### GetLogicalStats

`func (o *ClusterStats) GetLogicalStats() LogicalStats`

GetLogicalStats returns the LogicalStats field if non-nil, zero value otherwise.

### GetLogicalStatsOk

`func (o *ClusterStats) GetLogicalStatsOk() (*LogicalStats, bool)`

GetLogicalStatsOk returns a tuple with the LogicalStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalStats

`func (o *ClusterStats) SetLogicalStats(v LogicalStats)`

SetLogicalStats sets LogicalStats field to given value.

### HasLogicalStats

`func (o *ClusterStats) HasLogicalStats() bool`

HasLogicalStats returns a boolean if a field has been set.

### SetLogicalStatsNil

`func (o *ClusterStats) SetLogicalStatsNil(b bool)`

 SetLogicalStatsNil sets the value for LogicalStats to be an explicit nil

### UnsetLogicalStats
`func (o *ClusterStats) UnsetLogicalStats()`

UnsetLogicalStats ensures that no value is present for LogicalStats, not even an explicit nil
### GetUsagePerfStats

`func (o *ClusterStats) GetUsagePerfStats() UsageAndPerformanceStats`

GetUsagePerfStats returns the UsagePerfStats field if non-nil, zero value otherwise.

### GetUsagePerfStatsOk

`func (o *ClusterStats) GetUsagePerfStatsOk() (*UsageAndPerformanceStats, bool)`

GetUsagePerfStatsOk returns a tuple with the UsagePerfStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePerfStats

`func (o *ClusterStats) SetUsagePerfStats(v UsageAndPerformanceStats)`

SetUsagePerfStats sets UsagePerfStats field to given value.

### HasUsagePerfStats

`func (o *ClusterStats) HasUsagePerfStats() bool`

HasUsagePerfStats returns a boolean if a field has been set.

### SetUsagePerfStatsNil

`func (o *ClusterStats) SetUsagePerfStatsNil(b bool)`

 SetUsagePerfStatsNil sets the value for UsagePerfStats to be an explicit nil

### UnsetUsagePerfStats
`func (o *ClusterStats) UnsetUsagePerfStats()`

UnsetUsagePerfStats ensures that no value is present for UsagePerfStats, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


