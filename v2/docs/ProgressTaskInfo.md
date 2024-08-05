# ProgressTaskInfo

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

## Methods

### NewProgressTaskInfo

`func NewProgressTaskInfo() *ProgressTaskInfo`

NewProgressTaskInfo instantiates a new ProgressTaskInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgressTaskInfoWithDefaults

`func NewProgressTaskInfoWithDefaults() *ProgressTaskInfo`

NewProgressTaskInfoWithDefaults instantiates a new ProgressTaskInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTimeUsecs

`func (o *ProgressTaskInfo) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *ProgressTaskInfo) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *ProgressTaskInfo) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *ProgressTaskInfo) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *ProgressTaskInfo) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *ProgressTaskInfo) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetEvents

`func (o *ProgressTaskInfo) GetEvents() []ProgressTaskEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ProgressTaskInfo) GetEventsOk() (*[]ProgressTaskEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ProgressTaskInfo) SetEvents(v []ProgressTaskEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ProgressTaskInfo) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetExpectedRemainingTimeUsecs

`func (o *ProgressTaskInfo) GetExpectedRemainingTimeUsecs() int64`

GetExpectedRemainingTimeUsecs returns the ExpectedRemainingTimeUsecs field if non-nil, zero value otherwise.

### GetExpectedRemainingTimeUsecsOk

`func (o *ProgressTaskInfo) GetExpectedRemainingTimeUsecsOk() (*int64, bool)`

GetExpectedRemainingTimeUsecsOk returns a tuple with the ExpectedRemainingTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedRemainingTimeUsecs

`func (o *ProgressTaskInfo) SetExpectedRemainingTimeUsecs(v int64)`

SetExpectedRemainingTimeUsecs sets ExpectedRemainingTimeUsecs field to given value.

### HasExpectedRemainingTimeUsecs

`func (o *ProgressTaskInfo) HasExpectedRemainingTimeUsecs() bool`

HasExpectedRemainingTimeUsecs returns a boolean if a field has been set.

### SetExpectedRemainingTimeUsecsNil

`func (o *ProgressTaskInfo) SetExpectedRemainingTimeUsecsNil(b bool)`

 SetExpectedRemainingTimeUsecsNil sets the value for ExpectedRemainingTimeUsecs to be an explicit nil

### UnsetExpectedRemainingTimeUsecs
`func (o *ProgressTaskInfo) UnsetExpectedRemainingTimeUsecs()`

UnsetExpectedRemainingTimeUsecs ensures that no value is present for ExpectedRemainingTimeUsecs, not even an explicit nil
### GetPercentageCompleted

`func (o *ProgressTaskInfo) GetPercentageCompleted() float32`

GetPercentageCompleted returns the PercentageCompleted field if non-nil, zero value otherwise.

### GetPercentageCompletedOk

`func (o *ProgressTaskInfo) GetPercentageCompletedOk() (*float32, bool)`

GetPercentageCompletedOk returns a tuple with the PercentageCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageCompleted

`func (o *ProgressTaskInfo) SetPercentageCompleted(v float32)`

SetPercentageCompleted sets PercentageCompleted field to given value.

### HasPercentageCompleted

`func (o *ProgressTaskInfo) HasPercentageCompleted() bool`

HasPercentageCompleted returns a boolean if a field has been set.

### SetPercentageCompletedNil

`func (o *ProgressTaskInfo) SetPercentageCompletedNil(b bool)`

 SetPercentageCompletedNil sets the value for PercentageCompleted to be an explicit nil

### UnsetPercentageCompleted
`func (o *ProgressTaskInfo) UnsetPercentageCompleted()`

UnsetPercentageCompleted ensures that no value is present for PercentageCompleted, not even an explicit nil
### GetStartTimeUsecs

`func (o *ProgressTaskInfo) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *ProgressTaskInfo) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *ProgressTaskInfo) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *ProgressTaskInfo) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *ProgressTaskInfo) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *ProgressTaskInfo) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetStats

`func (o *ProgressTaskInfo) GetStats() ProgressStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ProgressTaskInfo) GetStatsOk() (*ProgressStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ProgressTaskInfo) SetStats(v ProgressStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ProgressTaskInfo) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetStatus

`func (o *ProgressTaskInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProgressTaskInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProgressTaskInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProgressTaskInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ProgressTaskInfo) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ProgressTaskInfo) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


