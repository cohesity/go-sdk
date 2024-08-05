# NodeStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** | Id is the Id of the Node. | [optional] 
**UsagePerfStats** | Pointer to [**UsageAndPerformanceStats**](UsageAndPerformanceStats.md) |  | [optional] 

## Methods

### NewNodeStats

`func NewNodeStats() *NodeStats`

NewNodeStats instantiates a new NodeStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeStatsWithDefaults

`func NewNodeStatsWithDefaults() *NodeStats`

NewNodeStatsWithDefaults instantiates a new NodeStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NodeStats) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodeStats) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodeStats) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NodeStats) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *NodeStats) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *NodeStats) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetUsagePerfStats

`func (o *NodeStats) GetUsagePerfStats() UsageAndPerformanceStats`

GetUsagePerfStats returns the UsagePerfStats field if non-nil, zero value otherwise.

### GetUsagePerfStatsOk

`func (o *NodeStats) GetUsagePerfStatsOk() (*UsageAndPerformanceStats, bool)`

GetUsagePerfStatsOk returns a tuple with the UsagePerfStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePerfStats

`func (o *NodeStats) SetUsagePerfStats(v UsageAndPerformanceStats)`

SetUsagePerfStats sets UsagePerfStats field to given value.

### HasUsagePerfStats

`func (o *NodeStats) HasUsagePerfStats() bool`

HasUsagePerfStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


