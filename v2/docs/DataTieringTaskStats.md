# DataTieringTaskStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BytesRead** | Pointer to **NullableInt64** | Specifies total logical bytes read for creating the snapshot. | [optional] 
**BytesWritten** | Pointer to **NullableInt64** | Specifies total size of data in bytes written after taking backup. | [optional] 
**LogicalSizeBytes** | Pointer to **NullableInt64** | Specifies total logical size of object(s) in bytes. | [optional] 
**ChangedEntityCount** | Pointer to **NullableInt64** | Specifies changed entity count. | [optional] 
**EntityCount** | Pointer to **NullableInt64** | Specifies total entity count. | [optional] 
**IsTieringGoalMet** | Pointer to **NullableBool** | Specifies whether tiering goal has been met. | [optional] [default to false]
**TotalTieredBytes** | Pointer to **NullableInt64** | Specifies total amount of data successfully tiered from the NAS source. | [optional] 

## Methods

### NewDataTieringTaskStats

`func NewDataTieringTaskStats() *DataTieringTaskStats`

NewDataTieringTaskStats instantiates a new DataTieringTaskStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTieringTaskStatsWithDefaults

`func NewDataTieringTaskStatsWithDefaults() *DataTieringTaskStats`

NewDataTieringTaskStatsWithDefaults instantiates a new DataTieringTaskStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytesRead

`func (o *DataTieringTaskStats) GetBytesRead() int64`

GetBytesRead returns the BytesRead field if non-nil, zero value otherwise.

### GetBytesReadOk

`func (o *DataTieringTaskStats) GetBytesReadOk() (*int64, bool)`

GetBytesReadOk returns a tuple with the BytesRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesRead

`func (o *DataTieringTaskStats) SetBytesRead(v int64)`

SetBytesRead sets BytesRead field to given value.

### HasBytesRead

`func (o *DataTieringTaskStats) HasBytesRead() bool`

HasBytesRead returns a boolean if a field has been set.

### SetBytesReadNil

`func (o *DataTieringTaskStats) SetBytesReadNil(b bool)`

 SetBytesReadNil sets the value for BytesRead to be an explicit nil

### UnsetBytesRead
`func (o *DataTieringTaskStats) UnsetBytesRead()`

UnsetBytesRead ensures that no value is present for BytesRead, not even an explicit nil
### GetBytesWritten

`func (o *DataTieringTaskStats) GetBytesWritten() int64`

GetBytesWritten returns the BytesWritten field if non-nil, zero value otherwise.

### GetBytesWrittenOk

`func (o *DataTieringTaskStats) GetBytesWrittenOk() (*int64, bool)`

GetBytesWrittenOk returns a tuple with the BytesWritten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesWritten

`func (o *DataTieringTaskStats) SetBytesWritten(v int64)`

SetBytesWritten sets BytesWritten field to given value.

### HasBytesWritten

`func (o *DataTieringTaskStats) HasBytesWritten() bool`

HasBytesWritten returns a boolean if a field has been set.

### SetBytesWrittenNil

`func (o *DataTieringTaskStats) SetBytesWrittenNil(b bool)`

 SetBytesWrittenNil sets the value for BytesWritten to be an explicit nil

### UnsetBytesWritten
`func (o *DataTieringTaskStats) UnsetBytesWritten()`

UnsetBytesWritten ensures that no value is present for BytesWritten, not even an explicit nil
### GetLogicalSizeBytes

`func (o *DataTieringTaskStats) GetLogicalSizeBytes() int64`

GetLogicalSizeBytes returns the LogicalSizeBytes field if non-nil, zero value otherwise.

### GetLogicalSizeBytesOk

`func (o *DataTieringTaskStats) GetLogicalSizeBytesOk() (*int64, bool)`

GetLogicalSizeBytesOk returns a tuple with the LogicalSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalSizeBytes

`func (o *DataTieringTaskStats) SetLogicalSizeBytes(v int64)`

SetLogicalSizeBytes sets LogicalSizeBytes field to given value.

### HasLogicalSizeBytes

`func (o *DataTieringTaskStats) HasLogicalSizeBytes() bool`

HasLogicalSizeBytes returns a boolean if a field has been set.

### SetLogicalSizeBytesNil

`func (o *DataTieringTaskStats) SetLogicalSizeBytesNil(b bool)`

 SetLogicalSizeBytesNil sets the value for LogicalSizeBytes to be an explicit nil

### UnsetLogicalSizeBytes
`func (o *DataTieringTaskStats) UnsetLogicalSizeBytes()`

UnsetLogicalSizeBytes ensures that no value is present for LogicalSizeBytes, not even an explicit nil
### GetChangedEntityCount

`func (o *DataTieringTaskStats) GetChangedEntityCount() int64`

GetChangedEntityCount returns the ChangedEntityCount field if non-nil, zero value otherwise.

### GetChangedEntityCountOk

`func (o *DataTieringTaskStats) GetChangedEntityCountOk() (*int64, bool)`

GetChangedEntityCountOk returns a tuple with the ChangedEntityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedEntityCount

`func (o *DataTieringTaskStats) SetChangedEntityCount(v int64)`

SetChangedEntityCount sets ChangedEntityCount field to given value.

### HasChangedEntityCount

`func (o *DataTieringTaskStats) HasChangedEntityCount() bool`

HasChangedEntityCount returns a boolean if a field has been set.

### SetChangedEntityCountNil

`func (o *DataTieringTaskStats) SetChangedEntityCountNil(b bool)`

 SetChangedEntityCountNil sets the value for ChangedEntityCount to be an explicit nil

### UnsetChangedEntityCount
`func (o *DataTieringTaskStats) UnsetChangedEntityCount()`

UnsetChangedEntityCount ensures that no value is present for ChangedEntityCount, not even an explicit nil
### GetEntityCount

`func (o *DataTieringTaskStats) GetEntityCount() int64`

GetEntityCount returns the EntityCount field if non-nil, zero value otherwise.

### GetEntityCountOk

`func (o *DataTieringTaskStats) GetEntityCountOk() (*int64, bool)`

GetEntityCountOk returns a tuple with the EntityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityCount

`func (o *DataTieringTaskStats) SetEntityCount(v int64)`

SetEntityCount sets EntityCount field to given value.

### HasEntityCount

`func (o *DataTieringTaskStats) HasEntityCount() bool`

HasEntityCount returns a boolean if a field has been set.

### SetEntityCountNil

`func (o *DataTieringTaskStats) SetEntityCountNil(b bool)`

 SetEntityCountNil sets the value for EntityCount to be an explicit nil

### UnsetEntityCount
`func (o *DataTieringTaskStats) UnsetEntityCount()`

UnsetEntityCount ensures that no value is present for EntityCount, not even an explicit nil
### GetIsTieringGoalMet

`func (o *DataTieringTaskStats) GetIsTieringGoalMet() bool`

GetIsTieringGoalMet returns the IsTieringGoalMet field if non-nil, zero value otherwise.

### GetIsTieringGoalMetOk

`func (o *DataTieringTaskStats) GetIsTieringGoalMetOk() (*bool, bool)`

GetIsTieringGoalMetOk returns a tuple with the IsTieringGoalMet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTieringGoalMet

`func (o *DataTieringTaskStats) SetIsTieringGoalMet(v bool)`

SetIsTieringGoalMet sets IsTieringGoalMet field to given value.

### HasIsTieringGoalMet

`func (o *DataTieringTaskStats) HasIsTieringGoalMet() bool`

HasIsTieringGoalMet returns a boolean if a field has been set.

### SetIsTieringGoalMetNil

`func (o *DataTieringTaskStats) SetIsTieringGoalMetNil(b bool)`

 SetIsTieringGoalMetNil sets the value for IsTieringGoalMet to be an explicit nil

### UnsetIsTieringGoalMet
`func (o *DataTieringTaskStats) UnsetIsTieringGoalMet()`

UnsetIsTieringGoalMet ensures that no value is present for IsTieringGoalMet, not even an explicit nil
### GetTotalTieredBytes

`func (o *DataTieringTaskStats) GetTotalTieredBytes() int64`

GetTotalTieredBytes returns the TotalTieredBytes field if non-nil, zero value otherwise.

### GetTotalTieredBytesOk

`func (o *DataTieringTaskStats) GetTotalTieredBytesOk() (*int64, bool)`

GetTotalTieredBytesOk returns a tuple with the TotalTieredBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTieredBytes

`func (o *DataTieringTaskStats) SetTotalTieredBytes(v int64)`

SetTotalTieredBytes sets TotalTieredBytes field to given value.

### HasTotalTieredBytes

`func (o *DataTieringTaskStats) HasTotalTieredBytes() bool`

HasTotalTieredBytes returns a boolean if a field has been set.

### SetTotalTieredBytesNil

`func (o *DataTieringTaskStats) SetTotalTieredBytesNil(b bool)`

 SetTotalTieredBytesNil sets the value for TotalTieredBytes to be an explicit nil

### UnsetTotalTieredBytes
`func (o *DataTieringTaskStats) UnsetTotalTieredBytes()`

UnsetTotalTieredBytes ensures that no value is present for TotalTieredBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


