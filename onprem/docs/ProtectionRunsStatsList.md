# ProtectionRunsStatsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **NullableInt64** | Specifies a Unix epoch Timestamp (in microseconds) of this statistics. | [optional] 
**Stats** | Pointer to [**[]ProtectionRunsStats**](ProtectionRunsStats.md) | Specifies the protection runs stats. | [optional] 

## Methods

### NewProtectionRunsStatsList

`func NewProtectionRunsStatsList() *ProtectionRunsStatsList`

NewProtectionRunsStatsList instantiates a new ProtectionRunsStatsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionRunsStatsListWithDefaults

`func NewProtectionRunsStatsListWithDefaults() *ProtectionRunsStatsList`

NewProtectionRunsStatsListWithDefaults instantiates a new ProtectionRunsStatsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *ProtectionRunsStatsList) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ProtectionRunsStatsList) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ProtectionRunsStatsList) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ProtectionRunsStatsList) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ProtectionRunsStatsList) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ProtectionRunsStatsList) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetStats

`func (o *ProtectionRunsStatsList) GetStats() []ProtectionRunsStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ProtectionRunsStatsList) GetStatsOk() (*[]ProtectionRunsStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ProtectionRunsStatsList) SetStats(v []ProtectionRunsStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ProtectionRunsStatsList) HasStats() bool`

HasStats returns a boolean if a field has been set.

### SetStatsNil

`func (o *ProtectionRunsStatsList) SetStatsNil(b bool)`

 SetStatsNil sets the value for Stats to be an explicit nil

### UnsetStats
`func (o *ProtectionRunsStatsList) UnsetStats()`

UnsetStats ensures that no value is present for Stats, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


