# CopyRunStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the time when this replication ended. If not set, then the replication has not ended yet. | [optional] 
**IsIncremental** | Pointer to **NullableBool** | Specifies whether this archival is incremental for archival targets. | [optional] 
**LogicalBytesTransferred** | Pointer to **NullableInt64** | Specifies the number of logical bytes transferred for this replication so far. This value can never exceed the total logical size of the replicated view. | [optional] 
**LogicalSizeBytes** | Pointer to **NullableInt64** | Specifies the total amount of logical data to be transferred for this replication. | [optional] 
**LogicalTransferRateBps** | Pointer to **NullableInt64** | Specifies average logical bytes transfer rate in bytes per second for archchival targets. | [optional] 
**PhysicalBytesTransferred** | Pointer to **NullableInt64** | Specifies the number of physical bytes sent over the wire for replication targets. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the time when this replication was started. If not set, then replication has not been started yet. | [optional] 

## Methods

### NewCopyRunStats

`func NewCopyRunStats() *CopyRunStats`

NewCopyRunStats instantiates a new CopyRunStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCopyRunStatsWithDefaults

`func NewCopyRunStatsWithDefaults() *CopyRunStats`

NewCopyRunStatsWithDefaults instantiates a new CopyRunStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTimeUsecs

`func (o *CopyRunStats) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *CopyRunStats) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *CopyRunStats) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *CopyRunStats) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *CopyRunStats) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *CopyRunStats) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetIsIncremental

`func (o *CopyRunStats) GetIsIncremental() bool`

GetIsIncremental returns the IsIncremental field if non-nil, zero value otherwise.

### GetIsIncrementalOk

`func (o *CopyRunStats) GetIsIncrementalOk() (*bool, bool)`

GetIsIncrementalOk returns a tuple with the IsIncremental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIncremental

`func (o *CopyRunStats) SetIsIncremental(v bool)`

SetIsIncremental sets IsIncremental field to given value.

### HasIsIncremental

`func (o *CopyRunStats) HasIsIncremental() bool`

HasIsIncremental returns a boolean if a field has been set.

### SetIsIncrementalNil

`func (o *CopyRunStats) SetIsIncrementalNil(b bool)`

 SetIsIncrementalNil sets the value for IsIncremental to be an explicit nil

### UnsetIsIncremental
`func (o *CopyRunStats) UnsetIsIncremental()`

UnsetIsIncremental ensures that no value is present for IsIncremental, not even an explicit nil
### GetLogicalBytesTransferred

`func (o *CopyRunStats) GetLogicalBytesTransferred() int64`

GetLogicalBytesTransferred returns the LogicalBytesTransferred field if non-nil, zero value otherwise.

### GetLogicalBytesTransferredOk

`func (o *CopyRunStats) GetLogicalBytesTransferredOk() (*int64, bool)`

GetLogicalBytesTransferredOk returns a tuple with the LogicalBytesTransferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalBytesTransferred

`func (o *CopyRunStats) SetLogicalBytesTransferred(v int64)`

SetLogicalBytesTransferred sets LogicalBytesTransferred field to given value.

### HasLogicalBytesTransferred

`func (o *CopyRunStats) HasLogicalBytesTransferred() bool`

HasLogicalBytesTransferred returns a boolean if a field has been set.

### SetLogicalBytesTransferredNil

`func (o *CopyRunStats) SetLogicalBytesTransferredNil(b bool)`

 SetLogicalBytesTransferredNil sets the value for LogicalBytesTransferred to be an explicit nil

### UnsetLogicalBytesTransferred
`func (o *CopyRunStats) UnsetLogicalBytesTransferred()`

UnsetLogicalBytesTransferred ensures that no value is present for LogicalBytesTransferred, not even an explicit nil
### GetLogicalSizeBytes

`func (o *CopyRunStats) GetLogicalSizeBytes() int64`

GetLogicalSizeBytes returns the LogicalSizeBytes field if non-nil, zero value otherwise.

### GetLogicalSizeBytesOk

`func (o *CopyRunStats) GetLogicalSizeBytesOk() (*int64, bool)`

GetLogicalSizeBytesOk returns a tuple with the LogicalSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalSizeBytes

`func (o *CopyRunStats) SetLogicalSizeBytes(v int64)`

SetLogicalSizeBytes sets LogicalSizeBytes field to given value.

### HasLogicalSizeBytes

`func (o *CopyRunStats) HasLogicalSizeBytes() bool`

HasLogicalSizeBytes returns a boolean if a field has been set.

### SetLogicalSizeBytesNil

`func (o *CopyRunStats) SetLogicalSizeBytesNil(b bool)`

 SetLogicalSizeBytesNil sets the value for LogicalSizeBytes to be an explicit nil

### UnsetLogicalSizeBytes
`func (o *CopyRunStats) UnsetLogicalSizeBytes()`

UnsetLogicalSizeBytes ensures that no value is present for LogicalSizeBytes, not even an explicit nil
### GetLogicalTransferRateBps

`func (o *CopyRunStats) GetLogicalTransferRateBps() int64`

GetLogicalTransferRateBps returns the LogicalTransferRateBps field if non-nil, zero value otherwise.

### GetLogicalTransferRateBpsOk

`func (o *CopyRunStats) GetLogicalTransferRateBpsOk() (*int64, bool)`

GetLogicalTransferRateBpsOk returns a tuple with the LogicalTransferRateBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalTransferRateBps

`func (o *CopyRunStats) SetLogicalTransferRateBps(v int64)`

SetLogicalTransferRateBps sets LogicalTransferRateBps field to given value.

### HasLogicalTransferRateBps

`func (o *CopyRunStats) HasLogicalTransferRateBps() bool`

HasLogicalTransferRateBps returns a boolean if a field has been set.

### SetLogicalTransferRateBpsNil

`func (o *CopyRunStats) SetLogicalTransferRateBpsNil(b bool)`

 SetLogicalTransferRateBpsNil sets the value for LogicalTransferRateBps to be an explicit nil

### UnsetLogicalTransferRateBps
`func (o *CopyRunStats) UnsetLogicalTransferRateBps()`

UnsetLogicalTransferRateBps ensures that no value is present for LogicalTransferRateBps, not even an explicit nil
### GetPhysicalBytesTransferred

`func (o *CopyRunStats) GetPhysicalBytesTransferred() int64`

GetPhysicalBytesTransferred returns the PhysicalBytesTransferred field if non-nil, zero value otherwise.

### GetPhysicalBytesTransferredOk

`func (o *CopyRunStats) GetPhysicalBytesTransferredOk() (*int64, bool)`

GetPhysicalBytesTransferredOk returns a tuple with the PhysicalBytesTransferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalBytesTransferred

`func (o *CopyRunStats) SetPhysicalBytesTransferred(v int64)`

SetPhysicalBytesTransferred sets PhysicalBytesTransferred field to given value.

### HasPhysicalBytesTransferred

`func (o *CopyRunStats) HasPhysicalBytesTransferred() bool`

HasPhysicalBytesTransferred returns a boolean if a field has been set.

### SetPhysicalBytesTransferredNil

`func (o *CopyRunStats) SetPhysicalBytesTransferredNil(b bool)`

 SetPhysicalBytesTransferredNil sets the value for PhysicalBytesTransferred to be an explicit nil

### UnsetPhysicalBytesTransferred
`func (o *CopyRunStats) UnsetPhysicalBytesTransferred()`

UnsetPhysicalBytesTransferred ensures that no value is present for PhysicalBytesTransferred, not even an explicit nil
### GetStartTimeUsecs

`func (o *CopyRunStats) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *CopyRunStats) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *CopyRunStats) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *CopyRunStats) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *CopyRunStats) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *CopyRunStats) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


