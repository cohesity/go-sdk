# LogicalStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalLogicalUsageBytes** | Pointer to **NullableInt64** | Provides the combined data residing on protected objects. The size of data before reduction by deduplication and compression. | [optional] 

## Methods

### NewLogicalStats

`func NewLogicalStats() *LogicalStats`

NewLogicalStats instantiates a new LogicalStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogicalStatsWithDefaults

`func NewLogicalStatsWithDefaults() *LogicalStats`

NewLogicalStatsWithDefaults instantiates a new LogicalStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalLogicalUsageBytes

`func (o *LogicalStats) GetTotalLogicalUsageBytes() int64`

GetTotalLogicalUsageBytes returns the TotalLogicalUsageBytes field if non-nil, zero value otherwise.

### GetTotalLogicalUsageBytesOk

`func (o *LogicalStats) GetTotalLogicalUsageBytesOk() (*int64, bool)`

GetTotalLogicalUsageBytesOk returns a tuple with the TotalLogicalUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLogicalUsageBytes

`func (o *LogicalStats) SetTotalLogicalUsageBytes(v int64)`

SetTotalLogicalUsageBytes sets TotalLogicalUsageBytes field to given value.

### HasTotalLogicalUsageBytes

`func (o *LogicalStats) HasTotalLogicalUsageBytes() bool`

HasTotalLogicalUsageBytes returns a boolean if a field has been set.

### SetTotalLogicalUsageBytesNil

`func (o *LogicalStats) SetTotalLogicalUsageBytesNil(b bool)`

 SetTotalLogicalUsageBytesNil sets the value for TotalLogicalUsageBytes to be an explicit nil

### UnsetTotalLogicalUsageBytes
`func (o *LogicalStats) UnsetTotalLogicalUsageBytes()`

UnsetTotalLogicalUsageBytes ensures that no value is present for TotalLogicalUsageBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


