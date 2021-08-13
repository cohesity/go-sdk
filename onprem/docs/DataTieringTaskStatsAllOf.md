# DataTieringTaskStatsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityCount** | Pointer to **NullableInt64** | Specifies total entity count. | [optional] 
**ChangedEntityCount** | Pointer to **NullableInt64** | Specifies changed entity count. | [optional] 
**IsTieringGoalMet** | Pointer to **NullableBool** | Specifies whether tiering goal has been met. | [optional] [default to false]
**TotalTieredBytes** | Pointer to **NullableInt64** | Specifies total amount of data successfully tiered from the NAS source. | [optional] 

## Methods

### NewDataTieringTaskStatsAllOf

`func NewDataTieringTaskStatsAllOf() *DataTieringTaskStatsAllOf`

NewDataTieringTaskStatsAllOf instantiates a new DataTieringTaskStatsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTieringTaskStatsAllOfWithDefaults

`func NewDataTieringTaskStatsAllOfWithDefaults() *DataTieringTaskStatsAllOf`

NewDataTieringTaskStatsAllOfWithDefaults instantiates a new DataTieringTaskStatsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityCount

`func (o *DataTieringTaskStatsAllOf) GetEntityCount() int64`

GetEntityCount returns the EntityCount field if non-nil, zero value otherwise.

### GetEntityCountOk

`func (o *DataTieringTaskStatsAllOf) GetEntityCountOk() (*int64, bool)`

GetEntityCountOk returns a tuple with the EntityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityCount

`func (o *DataTieringTaskStatsAllOf) SetEntityCount(v int64)`

SetEntityCount sets EntityCount field to given value.

### HasEntityCount

`func (o *DataTieringTaskStatsAllOf) HasEntityCount() bool`

HasEntityCount returns a boolean if a field has been set.

### SetEntityCountNil

`func (o *DataTieringTaskStatsAllOf) SetEntityCountNil(b bool)`

 SetEntityCountNil sets the value for EntityCount to be an explicit nil

### UnsetEntityCount
`func (o *DataTieringTaskStatsAllOf) UnsetEntityCount()`

UnsetEntityCount ensures that no value is present for EntityCount, not even an explicit nil
### GetChangedEntityCount

`func (o *DataTieringTaskStatsAllOf) GetChangedEntityCount() int64`

GetChangedEntityCount returns the ChangedEntityCount field if non-nil, zero value otherwise.

### GetChangedEntityCountOk

`func (o *DataTieringTaskStatsAllOf) GetChangedEntityCountOk() (*int64, bool)`

GetChangedEntityCountOk returns a tuple with the ChangedEntityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedEntityCount

`func (o *DataTieringTaskStatsAllOf) SetChangedEntityCount(v int64)`

SetChangedEntityCount sets ChangedEntityCount field to given value.

### HasChangedEntityCount

`func (o *DataTieringTaskStatsAllOf) HasChangedEntityCount() bool`

HasChangedEntityCount returns a boolean if a field has been set.

### SetChangedEntityCountNil

`func (o *DataTieringTaskStatsAllOf) SetChangedEntityCountNil(b bool)`

 SetChangedEntityCountNil sets the value for ChangedEntityCount to be an explicit nil

### UnsetChangedEntityCount
`func (o *DataTieringTaskStatsAllOf) UnsetChangedEntityCount()`

UnsetChangedEntityCount ensures that no value is present for ChangedEntityCount, not even an explicit nil
### GetIsTieringGoalMet

`func (o *DataTieringTaskStatsAllOf) GetIsTieringGoalMet() bool`

GetIsTieringGoalMet returns the IsTieringGoalMet field if non-nil, zero value otherwise.

### GetIsTieringGoalMetOk

`func (o *DataTieringTaskStatsAllOf) GetIsTieringGoalMetOk() (*bool, bool)`

GetIsTieringGoalMetOk returns a tuple with the IsTieringGoalMet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTieringGoalMet

`func (o *DataTieringTaskStatsAllOf) SetIsTieringGoalMet(v bool)`

SetIsTieringGoalMet sets IsTieringGoalMet field to given value.

### HasIsTieringGoalMet

`func (o *DataTieringTaskStatsAllOf) HasIsTieringGoalMet() bool`

HasIsTieringGoalMet returns a boolean if a field has been set.

### SetIsTieringGoalMetNil

`func (o *DataTieringTaskStatsAllOf) SetIsTieringGoalMetNil(b bool)`

 SetIsTieringGoalMetNil sets the value for IsTieringGoalMet to be an explicit nil

### UnsetIsTieringGoalMet
`func (o *DataTieringTaskStatsAllOf) UnsetIsTieringGoalMet()`

UnsetIsTieringGoalMet ensures that no value is present for IsTieringGoalMet, not even an explicit nil
### GetTotalTieredBytes

`func (o *DataTieringTaskStatsAllOf) GetTotalTieredBytes() int64`

GetTotalTieredBytes returns the TotalTieredBytes field if non-nil, zero value otherwise.

### GetTotalTieredBytesOk

`func (o *DataTieringTaskStatsAllOf) GetTotalTieredBytesOk() (*int64, bool)`

GetTotalTieredBytesOk returns a tuple with the TotalTieredBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTieredBytes

`func (o *DataTieringTaskStatsAllOf) SetTotalTieredBytes(v int64)`

SetTotalTieredBytes sets TotalTieredBytes field to given value.

### HasTotalTieredBytes

`func (o *DataTieringTaskStatsAllOf) HasTotalTieredBytes() bool`

HasTotalTieredBytes returns a boolean if a field has been set.

### SetTotalTieredBytesNil

`func (o *DataTieringTaskStatsAllOf) SetTotalTieredBytesNil(b bool)`

 SetTotalTieredBytesNil sets the value for TotalTieredBytes to be an explicit nil

### UnsetTotalTieredBytes
`func (o *DataTieringTaskStatsAllOf) UnsetTotalTieredBytes()`

UnsetTotalTieredBytes ensures that no value is present for TotalTieredBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


