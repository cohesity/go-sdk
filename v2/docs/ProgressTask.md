# ProgressTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time of the progress task in Unix epoch Timestamp(in microseconds). | [optional] 
**Events** | Pointer to [**[]ProgressTaskEvent**](ProgressTaskEvent.md) | Specifies the event log created for progress Task. | [optional] 
**ExpectedRemainingTimeUsecs** | Pointer to **NullableInt64** | Specifies the expected remaining time of the progress task in Unix epoch Timestamp(in microseconds). | [optional] 
**PercentageCompleted** | Pointer to **NullableFloat32** | Specifies the current completed percentage of the progress task. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of the progress task in Unix epoch Timestamp(in microseconds). | [optional] 
**Stats** | Pointer to [**ProgressStats**](ProgressStats.md) |  | [optional] 
**Status** | Pointer to **NullableString** | Specifies the current status of the progress task. | [optional] 
**Id** | Pointer to **NullableString** | Specifies the task id of the Progress task. | [optional] 
**SubTasks** | Pointer to [**[]ProgressTask**](ProgressTask.md) | Specifies the list of Progress Task. | [optional] 

## Methods

### NewProgressTask

`func NewProgressTask() *ProgressTask`

NewProgressTask instantiates a new ProgressTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgressTaskWithDefaults

`func NewProgressTaskWithDefaults() *ProgressTask`

NewProgressTaskWithDefaults instantiates a new ProgressTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTimeUsecs

`func (o *ProgressTask) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *ProgressTask) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *ProgressTask) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *ProgressTask) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *ProgressTask) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *ProgressTask) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetEvents

`func (o *ProgressTask) GetEvents() []ProgressTaskEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ProgressTask) GetEventsOk() (*[]ProgressTaskEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ProgressTask) SetEvents(v []ProgressTaskEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ProgressTask) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetExpectedRemainingTimeUsecs

`func (o *ProgressTask) GetExpectedRemainingTimeUsecs() int64`

GetExpectedRemainingTimeUsecs returns the ExpectedRemainingTimeUsecs field if non-nil, zero value otherwise.

### GetExpectedRemainingTimeUsecsOk

`func (o *ProgressTask) GetExpectedRemainingTimeUsecsOk() (*int64, bool)`

GetExpectedRemainingTimeUsecsOk returns a tuple with the ExpectedRemainingTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedRemainingTimeUsecs

`func (o *ProgressTask) SetExpectedRemainingTimeUsecs(v int64)`

SetExpectedRemainingTimeUsecs sets ExpectedRemainingTimeUsecs field to given value.

### HasExpectedRemainingTimeUsecs

`func (o *ProgressTask) HasExpectedRemainingTimeUsecs() bool`

HasExpectedRemainingTimeUsecs returns a boolean if a field has been set.

### SetExpectedRemainingTimeUsecsNil

`func (o *ProgressTask) SetExpectedRemainingTimeUsecsNil(b bool)`

 SetExpectedRemainingTimeUsecsNil sets the value for ExpectedRemainingTimeUsecs to be an explicit nil

### UnsetExpectedRemainingTimeUsecs
`func (o *ProgressTask) UnsetExpectedRemainingTimeUsecs()`

UnsetExpectedRemainingTimeUsecs ensures that no value is present for ExpectedRemainingTimeUsecs, not even an explicit nil
### GetPercentageCompleted

`func (o *ProgressTask) GetPercentageCompleted() float32`

GetPercentageCompleted returns the PercentageCompleted field if non-nil, zero value otherwise.

### GetPercentageCompletedOk

`func (o *ProgressTask) GetPercentageCompletedOk() (*float32, bool)`

GetPercentageCompletedOk returns a tuple with the PercentageCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageCompleted

`func (o *ProgressTask) SetPercentageCompleted(v float32)`

SetPercentageCompleted sets PercentageCompleted field to given value.

### HasPercentageCompleted

`func (o *ProgressTask) HasPercentageCompleted() bool`

HasPercentageCompleted returns a boolean if a field has been set.

### SetPercentageCompletedNil

`func (o *ProgressTask) SetPercentageCompletedNil(b bool)`

 SetPercentageCompletedNil sets the value for PercentageCompleted to be an explicit nil

### UnsetPercentageCompleted
`func (o *ProgressTask) UnsetPercentageCompleted()`

UnsetPercentageCompleted ensures that no value is present for PercentageCompleted, not even an explicit nil
### GetStartTimeUsecs

`func (o *ProgressTask) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *ProgressTask) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *ProgressTask) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *ProgressTask) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *ProgressTask) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *ProgressTask) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetStats

`func (o *ProgressTask) GetStats() ProgressStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ProgressTask) GetStatsOk() (*ProgressStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ProgressTask) SetStats(v ProgressStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ProgressTask) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetStatus

`func (o *ProgressTask) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProgressTask) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProgressTask) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProgressTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ProgressTask) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ProgressTask) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetId

`func (o *ProgressTask) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProgressTask) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProgressTask) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProgressTask) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ProgressTask) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ProgressTask) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetSubTasks

`func (o *ProgressTask) GetSubTasks() []ProgressTask`

GetSubTasks returns the SubTasks field if non-nil, zero value otherwise.

### GetSubTasksOk

`func (o *ProgressTask) GetSubTasksOk() (*[]ProgressTask, bool)`

GetSubTasksOk returns a tuple with the SubTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTasks

`func (o *ProgressTask) SetSubTasks(v []ProgressTask)`

SetSubTasks sets SubTasks field to given value.

### HasSubTasks

`func (o *ProgressTask) HasSubTasks() bool`

HasSubTasks returns a boolean if a field has been set.

### SetSubTasksNil

`func (o *ProgressTask) SetSubTasksNil(b bool)`

 SetSubTasksNil sets the value for SubTasks to be an explicit nil

### UnsetSubTasks
`func (o *ProgressTask) UnsetSubTasks()`

UnsetSubTasks ensures that no value is present for SubTasks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


