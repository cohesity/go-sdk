# ViewBoxStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudUsagePerfStats** | Pointer to [**UsageAndPerformanceStats**](UsageAndPerformanceStats.md) |  | [optional] 
**DataUsageStats** | Pointer to [**DataUsageStats**](DataUsageStats.md) |  | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies the id of the Storage Domain (View Box). | [optional] 
**LocalUsagePerfStats** | Pointer to [**UsageAndPerformanceStats**](UsageAndPerformanceStats.md) |  | [optional] 
**LogicalStats** | Pointer to [**LogicalStats**](LogicalStats.md) |  | [optional] 
**UsagePerfStats** | Pointer to [**UsageAndPerformanceStats**](UsageAndPerformanceStats.md) |  | [optional] 

## Methods

### NewViewBoxStats

`func NewViewBoxStats() *ViewBoxStats`

NewViewBoxStats instantiates a new ViewBoxStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewBoxStatsWithDefaults

`func NewViewBoxStatsWithDefaults() *ViewBoxStats`

NewViewBoxStatsWithDefaults instantiates a new ViewBoxStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudUsagePerfStats

`func (o *ViewBoxStats) GetCloudUsagePerfStats() UsageAndPerformanceStats`

GetCloudUsagePerfStats returns the CloudUsagePerfStats field if non-nil, zero value otherwise.

### GetCloudUsagePerfStatsOk

`func (o *ViewBoxStats) GetCloudUsagePerfStatsOk() (*UsageAndPerformanceStats, bool)`

GetCloudUsagePerfStatsOk returns a tuple with the CloudUsagePerfStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudUsagePerfStats

`func (o *ViewBoxStats) SetCloudUsagePerfStats(v UsageAndPerformanceStats)`

SetCloudUsagePerfStats sets CloudUsagePerfStats field to given value.

### HasCloudUsagePerfStats

`func (o *ViewBoxStats) HasCloudUsagePerfStats() bool`

HasCloudUsagePerfStats returns a boolean if a field has been set.

### GetDataUsageStats

`func (o *ViewBoxStats) GetDataUsageStats() DataUsageStats`

GetDataUsageStats returns the DataUsageStats field if non-nil, zero value otherwise.

### GetDataUsageStatsOk

`func (o *ViewBoxStats) GetDataUsageStatsOk() (*DataUsageStats, bool)`

GetDataUsageStatsOk returns a tuple with the DataUsageStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataUsageStats

`func (o *ViewBoxStats) SetDataUsageStats(v DataUsageStats)`

SetDataUsageStats sets DataUsageStats field to given value.

### HasDataUsageStats

`func (o *ViewBoxStats) HasDataUsageStats() bool`

HasDataUsageStats returns a boolean if a field has been set.

### GetId

`func (o *ViewBoxStats) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewBoxStats) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewBoxStats) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ViewBoxStats) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ViewBoxStats) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ViewBoxStats) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLocalUsagePerfStats

`func (o *ViewBoxStats) GetLocalUsagePerfStats() UsageAndPerformanceStats`

GetLocalUsagePerfStats returns the LocalUsagePerfStats field if non-nil, zero value otherwise.

### GetLocalUsagePerfStatsOk

`func (o *ViewBoxStats) GetLocalUsagePerfStatsOk() (*UsageAndPerformanceStats, bool)`

GetLocalUsagePerfStatsOk returns a tuple with the LocalUsagePerfStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalUsagePerfStats

`func (o *ViewBoxStats) SetLocalUsagePerfStats(v UsageAndPerformanceStats)`

SetLocalUsagePerfStats sets LocalUsagePerfStats field to given value.

### HasLocalUsagePerfStats

`func (o *ViewBoxStats) HasLocalUsagePerfStats() bool`

HasLocalUsagePerfStats returns a boolean if a field has been set.

### GetLogicalStats

`func (o *ViewBoxStats) GetLogicalStats() LogicalStats`

GetLogicalStats returns the LogicalStats field if non-nil, zero value otherwise.

### GetLogicalStatsOk

`func (o *ViewBoxStats) GetLogicalStatsOk() (*LogicalStats, bool)`

GetLogicalStatsOk returns a tuple with the LogicalStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalStats

`func (o *ViewBoxStats) SetLogicalStats(v LogicalStats)`

SetLogicalStats sets LogicalStats field to given value.

### HasLogicalStats

`func (o *ViewBoxStats) HasLogicalStats() bool`

HasLogicalStats returns a boolean if a field has been set.

### GetUsagePerfStats

`func (o *ViewBoxStats) GetUsagePerfStats() UsageAndPerformanceStats`

GetUsagePerfStats returns the UsagePerfStats field if non-nil, zero value otherwise.

### GetUsagePerfStatsOk

`func (o *ViewBoxStats) GetUsagePerfStatsOk() (*UsageAndPerformanceStats, bool)`

GetUsagePerfStatsOk returns a tuple with the UsagePerfStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePerfStats

`func (o *ViewBoxStats) SetUsagePerfStats(v UsageAndPerformanceStats)`

SetUsagePerfStats sets UsagePerfStats field to given value.

### HasUsagePerfStats

`func (o *ViewBoxStats) HasUsagePerfStats() bool`

HasUsagePerfStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


