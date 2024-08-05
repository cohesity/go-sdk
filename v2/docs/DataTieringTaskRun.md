# DataTieringTaskRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time of task run in Unix epoch Timestamp(in microseconds). | [optional] 
**Id** | Pointer to **NullableString** | Specifies the id of the data tiering task run. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of task run in Unix epoch Timestamp(in microseconds). | [optional] 
**TieringInfo** | Pointer to [**TieringInfo**](TieringInfo.md) |  | [optional] 

## Methods

### NewDataTieringTaskRun

`func NewDataTieringTaskRun() *DataTieringTaskRun`

NewDataTieringTaskRun instantiates a new DataTieringTaskRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTieringTaskRunWithDefaults

`func NewDataTieringTaskRunWithDefaults() *DataTieringTaskRun`

NewDataTieringTaskRunWithDefaults instantiates a new DataTieringTaskRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTimeUsecs

`func (o *DataTieringTaskRun) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *DataTieringTaskRun) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *DataTieringTaskRun) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *DataTieringTaskRun) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *DataTieringTaskRun) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *DataTieringTaskRun) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetId

`func (o *DataTieringTaskRun) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataTieringTaskRun) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataTieringTaskRun) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataTieringTaskRun) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DataTieringTaskRun) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DataTieringTaskRun) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetStartTimeUsecs

`func (o *DataTieringTaskRun) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *DataTieringTaskRun) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *DataTieringTaskRun) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *DataTieringTaskRun) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *DataTieringTaskRun) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *DataTieringTaskRun) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetTieringInfo

`func (o *DataTieringTaskRun) GetTieringInfo() TieringInfo`

GetTieringInfo returns the TieringInfo field if non-nil, zero value otherwise.

### GetTieringInfoOk

`func (o *DataTieringTaskRun) GetTieringInfoOk() (*TieringInfo, bool)`

GetTieringInfoOk returns a tuple with the TieringInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTieringInfo

`func (o *DataTieringTaskRun) SetTieringInfo(v TieringInfo)`

SetTieringInfo sets TieringInfo field to given value.

### HasTieringInfo

`func (o *DataTieringTaskRun) HasTieringInfo() bool`

HasTieringInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


