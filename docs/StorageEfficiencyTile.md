# StorageEfficiencyTile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataInBytes** | Pointer to **NullableInt64** | Specifies the size of data brought into the cluster. This is the usage before data reduction if we ignore the zeroes and effects of cloning. | [optional] 
**DataInBytesSamples** | Pointer to [**[]Sample**](Sample.md) | Specifies the samples taken for Data brought into the cluster in bytes in ascending order of time. | [optional] 
**DataInDedupedBytes** | Pointer to **NullableInt64** | Specifies the size of data after compression and or dedupe operations just before the data is replicated to other nodes as per RF or Erasure Coding policy. | [optional] 
**DataInDedupedBytesSamples** | Pointer to [**[]Sample**](Sample.md) | Specifies the samples taken for morphed data in bytes in ascending order of time. | [optional] 
**DedupeRatio** | Pointer to **NullableFloat64** | Specifies the current dedupe ratio on the cluster. It is the ratio of DataInBytes to DataInDedupedBytes. | [optional] 
**DedupeRatioSamples** | Pointer to [**[]Sample**](Sample.md) | Specifies the samples for data reduction ratio in ascending order of time. | [optional] 
**DurationDays** | Pointer to **NullableInt32** | Specifies the duration in days in which the samples were taken. For this tile, it is 7 days. | [optional] 
**IntervalSeconds** | Pointer to **NullableInt32** | Specifies the interval between the samples in seconds. For this tile, it is 1 day which is 86400 seconds. | [optional] 
**LogicalUsedBytes** | Pointer to **NullableInt64** | Specifies the size of logical data currently represented on the cluster. | [optional] 
**LogicalUsedBytesSamples** | Pointer to [**[]Sample**](Sample.md) | Specifies the samples taken for logical data represented in bytes in ascending order of time. | [optional] 
**PhysicalUsedBytes** | Pointer to **NullableInt64** | Specifies the size of physical data currently consumed on the cluster. | [optional] 
**PhysicalUsedBytesSamples** | Pointer to [**[]Sample**](Sample.md) | Specifies the samples taken for physical data consumed in bytes in ascending order of time. | [optional] 
**StorageReductionRatio** | Pointer to **NullableFloat64** | Specifies the current storage reduction ratio on the cluster. It is the ratio of LogicalUsedBytes to PhysicalUsedBytes. | [optional] 
**StorageReductionSamples** | Pointer to [**[]Sample**](Sample.md) | Specifies the samples for storage reduction ratio in ascending order of time. | [optional] 

## Methods

### NewStorageEfficiencyTile

`func NewStorageEfficiencyTile() *StorageEfficiencyTile`

NewStorageEfficiencyTile instantiates a new StorageEfficiencyTile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageEfficiencyTileWithDefaults

`func NewStorageEfficiencyTileWithDefaults() *StorageEfficiencyTile`

NewStorageEfficiencyTileWithDefaults instantiates a new StorageEfficiencyTile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataInBytes

`func (o *StorageEfficiencyTile) GetDataInBytes() int64`

GetDataInBytes returns the DataInBytes field if non-nil, zero value otherwise.

### GetDataInBytesOk

`func (o *StorageEfficiencyTile) GetDataInBytesOk() (*int64, bool)`

GetDataInBytesOk returns a tuple with the DataInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataInBytes

`func (o *StorageEfficiencyTile) SetDataInBytes(v int64)`

SetDataInBytes sets DataInBytes field to given value.

### HasDataInBytes

`func (o *StorageEfficiencyTile) HasDataInBytes() bool`

HasDataInBytes returns a boolean if a field has been set.

### SetDataInBytesNil

`func (o *StorageEfficiencyTile) SetDataInBytesNil(b bool)`

 SetDataInBytesNil sets the value for DataInBytes to be an explicit nil

### UnsetDataInBytes
`func (o *StorageEfficiencyTile) UnsetDataInBytes()`

UnsetDataInBytes ensures that no value is present for DataInBytes, not even an explicit nil
### GetDataInBytesSamples

`func (o *StorageEfficiencyTile) GetDataInBytesSamples() []Sample`

GetDataInBytesSamples returns the DataInBytesSamples field if non-nil, zero value otherwise.

### GetDataInBytesSamplesOk

`func (o *StorageEfficiencyTile) GetDataInBytesSamplesOk() (*[]Sample, bool)`

GetDataInBytesSamplesOk returns a tuple with the DataInBytesSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataInBytesSamples

`func (o *StorageEfficiencyTile) SetDataInBytesSamples(v []Sample)`

SetDataInBytesSamples sets DataInBytesSamples field to given value.

### HasDataInBytesSamples

`func (o *StorageEfficiencyTile) HasDataInBytesSamples() bool`

HasDataInBytesSamples returns a boolean if a field has been set.

### SetDataInBytesSamplesNil

`func (o *StorageEfficiencyTile) SetDataInBytesSamplesNil(b bool)`

 SetDataInBytesSamplesNil sets the value for DataInBytesSamples to be an explicit nil

### UnsetDataInBytesSamples
`func (o *StorageEfficiencyTile) UnsetDataInBytesSamples()`

UnsetDataInBytesSamples ensures that no value is present for DataInBytesSamples, not even an explicit nil
### GetDataInDedupedBytes

`func (o *StorageEfficiencyTile) GetDataInDedupedBytes() int64`

GetDataInDedupedBytes returns the DataInDedupedBytes field if non-nil, zero value otherwise.

### GetDataInDedupedBytesOk

`func (o *StorageEfficiencyTile) GetDataInDedupedBytesOk() (*int64, bool)`

GetDataInDedupedBytesOk returns a tuple with the DataInDedupedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataInDedupedBytes

`func (o *StorageEfficiencyTile) SetDataInDedupedBytes(v int64)`

SetDataInDedupedBytes sets DataInDedupedBytes field to given value.

### HasDataInDedupedBytes

`func (o *StorageEfficiencyTile) HasDataInDedupedBytes() bool`

HasDataInDedupedBytes returns a boolean if a field has been set.

### SetDataInDedupedBytesNil

`func (o *StorageEfficiencyTile) SetDataInDedupedBytesNil(b bool)`

 SetDataInDedupedBytesNil sets the value for DataInDedupedBytes to be an explicit nil

### UnsetDataInDedupedBytes
`func (o *StorageEfficiencyTile) UnsetDataInDedupedBytes()`

UnsetDataInDedupedBytes ensures that no value is present for DataInDedupedBytes, not even an explicit nil
### GetDataInDedupedBytesSamples

`func (o *StorageEfficiencyTile) GetDataInDedupedBytesSamples() []Sample`

GetDataInDedupedBytesSamples returns the DataInDedupedBytesSamples field if non-nil, zero value otherwise.

### GetDataInDedupedBytesSamplesOk

`func (o *StorageEfficiencyTile) GetDataInDedupedBytesSamplesOk() (*[]Sample, bool)`

GetDataInDedupedBytesSamplesOk returns a tuple with the DataInDedupedBytesSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataInDedupedBytesSamples

`func (o *StorageEfficiencyTile) SetDataInDedupedBytesSamples(v []Sample)`

SetDataInDedupedBytesSamples sets DataInDedupedBytesSamples field to given value.

### HasDataInDedupedBytesSamples

`func (o *StorageEfficiencyTile) HasDataInDedupedBytesSamples() bool`

HasDataInDedupedBytesSamples returns a boolean if a field has been set.

### SetDataInDedupedBytesSamplesNil

`func (o *StorageEfficiencyTile) SetDataInDedupedBytesSamplesNil(b bool)`

 SetDataInDedupedBytesSamplesNil sets the value for DataInDedupedBytesSamples to be an explicit nil

### UnsetDataInDedupedBytesSamples
`func (o *StorageEfficiencyTile) UnsetDataInDedupedBytesSamples()`

UnsetDataInDedupedBytesSamples ensures that no value is present for DataInDedupedBytesSamples, not even an explicit nil
### GetDedupeRatio

`func (o *StorageEfficiencyTile) GetDedupeRatio() float64`

GetDedupeRatio returns the DedupeRatio field if non-nil, zero value otherwise.

### GetDedupeRatioOk

`func (o *StorageEfficiencyTile) GetDedupeRatioOk() (*float64, bool)`

GetDedupeRatioOk returns a tuple with the DedupeRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedupeRatio

`func (o *StorageEfficiencyTile) SetDedupeRatio(v float64)`

SetDedupeRatio sets DedupeRatio field to given value.

### HasDedupeRatio

`func (o *StorageEfficiencyTile) HasDedupeRatio() bool`

HasDedupeRatio returns a boolean if a field has been set.

### SetDedupeRatioNil

`func (o *StorageEfficiencyTile) SetDedupeRatioNil(b bool)`

 SetDedupeRatioNil sets the value for DedupeRatio to be an explicit nil

### UnsetDedupeRatio
`func (o *StorageEfficiencyTile) UnsetDedupeRatio()`

UnsetDedupeRatio ensures that no value is present for DedupeRatio, not even an explicit nil
### GetDedupeRatioSamples

`func (o *StorageEfficiencyTile) GetDedupeRatioSamples() []Sample`

GetDedupeRatioSamples returns the DedupeRatioSamples field if non-nil, zero value otherwise.

### GetDedupeRatioSamplesOk

`func (o *StorageEfficiencyTile) GetDedupeRatioSamplesOk() (*[]Sample, bool)`

GetDedupeRatioSamplesOk returns a tuple with the DedupeRatioSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedupeRatioSamples

`func (o *StorageEfficiencyTile) SetDedupeRatioSamples(v []Sample)`

SetDedupeRatioSamples sets DedupeRatioSamples field to given value.

### HasDedupeRatioSamples

`func (o *StorageEfficiencyTile) HasDedupeRatioSamples() bool`

HasDedupeRatioSamples returns a boolean if a field has been set.

### SetDedupeRatioSamplesNil

`func (o *StorageEfficiencyTile) SetDedupeRatioSamplesNil(b bool)`

 SetDedupeRatioSamplesNil sets the value for DedupeRatioSamples to be an explicit nil

### UnsetDedupeRatioSamples
`func (o *StorageEfficiencyTile) UnsetDedupeRatioSamples()`

UnsetDedupeRatioSamples ensures that no value is present for DedupeRatioSamples, not even an explicit nil
### GetDurationDays

`func (o *StorageEfficiencyTile) GetDurationDays() int32`

GetDurationDays returns the DurationDays field if non-nil, zero value otherwise.

### GetDurationDaysOk

`func (o *StorageEfficiencyTile) GetDurationDaysOk() (*int32, bool)`

GetDurationDaysOk returns a tuple with the DurationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationDays

`func (o *StorageEfficiencyTile) SetDurationDays(v int32)`

SetDurationDays sets DurationDays field to given value.

### HasDurationDays

`func (o *StorageEfficiencyTile) HasDurationDays() bool`

HasDurationDays returns a boolean if a field has been set.

### SetDurationDaysNil

`func (o *StorageEfficiencyTile) SetDurationDaysNil(b bool)`

 SetDurationDaysNil sets the value for DurationDays to be an explicit nil

### UnsetDurationDays
`func (o *StorageEfficiencyTile) UnsetDurationDays()`

UnsetDurationDays ensures that no value is present for DurationDays, not even an explicit nil
### GetIntervalSeconds

`func (o *StorageEfficiencyTile) GetIntervalSeconds() int32`

GetIntervalSeconds returns the IntervalSeconds field if non-nil, zero value otherwise.

### GetIntervalSecondsOk

`func (o *StorageEfficiencyTile) GetIntervalSecondsOk() (*int32, bool)`

GetIntervalSecondsOk returns a tuple with the IntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalSeconds

`func (o *StorageEfficiencyTile) SetIntervalSeconds(v int32)`

SetIntervalSeconds sets IntervalSeconds field to given value.

### HasIntervalSeconds

`func (o *StorageEfficiencyTile) HasIntervalSeconds() bool`

HasIntervalSeconds returns a boolean if a field has been set.

### SetIntervalSecondsNil

`func (o *StorageEfficiencyTile) SetIntervalSecondsNil(b bool)`

 SetIntervalSecondsNil sets the value for IntervalSeconds to be an explicit nil

### UnsetIntervalSeconds
`func (o *StorageEfficiencyTile) UnsetIntervalSeconds()`

UnsetIntervalSeconds ensures that no value is present for IntervalSeconds, not even an explicit nil
### GetLogicalUsedBytes

`func (o *StorageEfficiencyTile) GetLogicalUsedBytes() int64`

GetLogicalUsedBytes returns the LogicalUsedBytes field if non-nil, zero value otherwise.

### GetLogicalUsedBytesOk

`func (o *StorageEfficiencyTile) GetLogicalUsedBytesOk() (*int64, bool)`

GetLogicalUsedBytesOk returns a tuple with the LogicalUsedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalUsedBytes

`func (o *StorageEfficiencyTile) SetLogicalUsedBytes(v int64)`

SetLogicalUsedBytes sets LogicalUsedBytes field to given value.

### HasLogicalUsedBytes

`func (o *StorageEfficiencyTile) HasLogicalUsedBytes() bool`

HasLogicalUsedBytes returns a boolean if a field has been set.

### SetLogicalUsedBytesNil

`func (o *StorageEfficiencyTile) SetLogicalUsedBytesNil(b bool)`

 SetLogicalUsedBytesNil sets the value for LogicalUsedBytes to be an explicit nil

### UnsetLogicalUsedBytes
`func (o *StorageEfficiencyTile) UnsetLogicalUsedBytes()`

UnsetLogicalUsedBytes ensures that no value is present for LogicalUsedBytes, not even an explicit nil
### GetLogicalUsedBytesSamples

`func (o *StorageEfficiencyTile) GetLogicalUsedBytesSamples() []Sample`

GetLogicalUsedBytesSamples returns the LogicalUsedBytesSamples field if non-nil, zero value otherwise.

### GetLogicalUsedBytesSamplesOk

`func (o *StorageEfficiencyTile) GetLogicalUsedBytesSamplesOk() (*[]Sample, bool)`

GetLogicalUsedBytesSamplesOk returns a tuple with the LogicalUsedBytesSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalUsedBytesSamples

`func (o *StorageEfficiencyTile) SetLogicalUsedBytesSamples(v []Sample)`

SetLogicalUsedBytesSamples sets LogicalUsedBytesSamples field to given value.

### HasLogicalUsedBytesSamples

`func (o *StorageEfficiencyTile) HasLogicalUsedBytesSamples() bool`

HasLogicalUsedBytesSamples returns a boolean if a field has been set.

### SetLogicalUsedBytesSamplesNil

`func (o *StorageEfficiencyTile) SetLogicalUsedBytesSamplesNil(b bool)`

 SetLogicalUsedBytesSamplesNil sets the value for LogicalUsedBytesSamples to be an explicit nil

### UnsetLogicalUsedBytesSamples
`func (o *StorageEfficiencyTile) UnsetLogicalUsedBytesSamples()`

UnsetLogicalUsedBytesSamples ensures that no value is present for LogicalUsedBytesSamples, not even an explicit nil
### GetPhysicalUsedBytes

`func (o *StorageEfficiencyTile) GetPhysicalUsedBytes() int64`

GetPhysicalUsedBytes returns the PhysicalUsedBytes field if non-nil, zero value otherwise.

### GetPhysicalUsedBytesOk

`func (o *StorageEfficiencyTile) GetPhysicalUsedBytesOk() (*int64, bool)`

GetPhysicalUsedBytesOk returns a tuple with the PhysicalUsedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalUsedBytes

`func (o *StorageEfficiencyTile) SetPhysicalUsedBytes(v int64)`

SetPhysicalUsedBytes sets PhysicalUsedBytes field to given value.

### HasPhysicalUsedBytes

`func (o *StorageEfficiencyTile) HasPhysicalUsedBytes() bool`

HasPhysicalUsedBytes returns a boolean if a field has been set.

### SetPhysicalUsedBytesNil

`func (o *StorageEfficiencyTile) SetPhysicalUsedBytesNil(b bool)`

 SetPhysicalUsedBytesNil sets the value for PhysicalUsedBytes to be an explicit nil

### UnsetPhysicalUsedBytes
`func (o *StorageEfficiencyTile) UnsetPhysicalUsedBytes()`

UnsetPhysicalUsedBytes ensures that no value is present for PhysicalUsedBytes, not even an explicit nil
### GetPhysicalUsedBytesSamples

`func (o *StorageEfficiencyTile) GetPhysicalUsedBytesSamples() []Sample`

GetPhysicalUsedBytesSamples returns the PhysicalUsedBytesSamples field if non-nil, zero value otherwise.

### GetPhysicalUsedBytesSamplesOk

`func (o *StorageEfficiencyTile) GetPhysicalUsedBytesSamplesOk() (*[]Sample, bool)`

GetPhysicalUsedBytesSamplesOk returns a tuple with the PhysicalUsedBytesSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalUsedBytesSamples

`func (o *StorageEfficiencyTile) SetPhysicalUsedBytesSamples(v []Sample)`

SetPhysicalUsedBytesSamples sets PhysicalUsedBytesSamples field to given value.

### HasPhysicalUsedBytesSamples

`func (o *StorageEfficiencyTile) HasPhysicalUsedBytesSamples() bool`

HasPhysicalUsedBytesSamples returns a boolean if a field has been set.

### SetPhysicalUsedBytesSamplesNil

`func (o *StorageEfficiencyTile) SetPhysicalUsedBytesSamplesNil(b bool)`

 SetPhysicalUsedBytesSamplesNil sets the value for PhysicalUsedBytesSamples to be an explicit nil

### UnsetPhysicalUsedBytesSamples
`func (o *StorageEfficiencyTile) UnsetPhysicalUsedBytesSamples()`

UnsetPhysicalUsedBytesSamples ensures that no value is present for PhysicalUsedBytesSamples, not even an explicit nil
### GetStorageReductionRatio

`func (o *StorageEfficiencyTile) GetStorageReductionRatio() float64`

GetStorageReductionRatio returns the StorageReductionRatio field if non-nil, zero value otherwise.

### GetStorageReductionRatioOk

`func (o *StorageEfficiencyTile) GetStorageReductionRatioOk() (*float64, bool)`

GetStorageReductionRatioOk returns a tuple with the StorageReductionRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageReductionRatio

`func (o *StorageEfficiencyTile) SetStorageReductionRatio(v float64)`

SetStorageReductionRatio sets StorageReductionRatio field to given value.

### HasStorageReductionRatio

`func (o *StorageEfficiencyTile) HasStorageReductionRatio() bool`

HasStorageReductionRatio returns a boolean if a field has been set.

### SetStorageReductionRatioNil

`func (o *StorageEfficiencyTile) SetStorageReductionRatioNil(b bool)`

 SetStorageReductionRatioNil sets the value for StorageReductionRatio to be an explicit nil

### UnsetStorageReductionRatio
`func (o *StorageEfficiencyTile) UnsetStorageReductionRatio()`

UnsetStorageReductionRatio ensures that no value is present for StorageReductionRatio, not even an explicit nil
### GetStorageReductionSamples

`func (o *StorageEfficiencyTile) GetStorageReductionSamples() []Sample`

GetStorageReductionSamples returns the StorageReductionSamples field if non-nil, zero value otherwise.

### GetStorageReductionSamplesOk

`func (o *StorageEfficiencyTile) GetStorageReductionSamplesOk() (*[]Sample, bool)`

GetStorageReductionSamplesOk returns a tuple with the StorageReductionSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageReductionSamples

`func (o *StorageEfficiencyTile) SetStorageReductionSamples(v []Sample)`

SetStorageReductionSamples sets StorageReductionSamples field to given value.

### HasStorageReductionSamples

`func (o *StorageEfficiencyTile) HasStorageReductionSamples() bool`

HasStorageReductionSamples returns a boolean if a field has been set.

### SetStorageReductionSamplesNil

`func (o *StorageEfficiencyTile) SetStorageReductionSamplesNil(b bool)`

 SetStorageReductionSamplesNil sets the value for StorageReductionSamples to be an explicit nil

### UnsetStorageReductionSamples
`func (o *StorageEfficiencyTile) UnsetStorageReductionSamples()`

UnsetStorageReductionSamples ensures that no value is present for StorageReductionSamples, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


