# CloudRetrieveTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the task end time in microseconds. | [optional] 
**Message** | Pointer to **NullableString** | Message about the cloud retrieve task. | [optional] 
**PercentageFinished** | Pointer to **NullableInt64** | Specifies the percentage of the task. | [optional] 
**PulseTaskId** | Pointer to **NullableInt32** | Specifies the pulse task id. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the task start time in microseconds. | [optional] 
**Status** | Pointer to **NullableString** | Specifies the status of the retrieve task. | [optional] 
**TimeRemainingSeconds** | Pointer to **NullableInt64** | Specifies the time remaining to complete the retrieve task. | [optional] 

## Methods

### NewCloudRetrieveTask

`func NewCloudRetrieveTask() *CloudRetrieveTask`

NewCloudRetrieveTask instantiates a new CloudRetrieveTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRetrieveTaskWithDefaults

`func NewCloudRetrieveTaskWithDefaults() *CloudRetrieveTask`

NewCloudRetrieveTaskWithDefaults instantiates a new CloudRetrieveTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTimeUsecs

`func (o *CloudRetrieveTask) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *CloudRetrieveTask) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *CloudRetrieveTask) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *CloudRetrieveTask) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *CloudRetrieveTask) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *CloudRetrieveTask) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetMessage

`func (o *CloudRetrieveTask) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CloudRetrieveTask) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CloudRetrieveTask) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CloudRetrieveTask) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *CloudRetrieveTask) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *CloudRetrieveTask) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetPercentageFinished

`func (o *CloudRetrieveTask) GetPercentageFinished() int64`

GetPercentageFinished returns the PercentageFinished field if non-nil, zero value otherwise.

### GetPercentageFinishedOk

`func (o *CloudRetrieveTask) GetPercentageFinishedOk() (*int64, bool)`

GetPercentageFinishedOk returns a tuple with the PercentageFinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageFinished

`func (o *CloudRetrieveTask) SetPercentageFinished(v int64)`

SetPercentageFinished sets PercentageFinished field to given value.

### HasPercentageFinished

`func (o *CloudRetrieveTask) HasPercentageFinished() bool`

HasPercentageFinished returns a boolean if a field has been set.

### SetPercentageFinishedNil

`func (o *CloudRetrieveTask) SetPercentageFinishedNil(b bool)`

 SetPercentageFinishedNil sets the value for PercentageFinished to be an explicit nil

### UnsetPercentageFinished
`func (o *CloudRetrieveTask) UnsetPercentageFinished()`

UnsetPercentageFinished ensures that no value is present for PercentageFinished, not even an explicit nil
### GetPulseTaskId

`func (o *CloudRetrieveTask) GetPulseTaskId() int32`

GetPulseTaskId returns the PulseTaskId field if non-nil, zero value otherwise.

### GetPulseTaskIdOk

`func (o *CloudRetrieveTask) GetPulseTaskIdOk() (*int32, bool)`

GetPulseTaskIdOk returns a tuple with the PulseTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulseTaskId

`func (o *CloudRetrieveTask) SetPulseTaskId(v int32)`

SetPulseTaskId sets PulseTaskId field to given value.

### HasPulseTaskId

`func (o *CloudRetrieveTask) HasPulseTaskId() bool`

HasPulseTaskId returns a boolean if a field has been set.

### SetPulseTaskIdNil

`func (o *CloudRetrieveTask) SetPulseTaskIdNil(b bool)`

 SetPulseTaskIdNil sets the value for PulseTaskId to be an explicit nil

### UnsetPulseTaskId
`func (o *CloudRetrieveTask) UnsetPulseTaskId()`

UnsetPulseTaskId ensures that no value is present for PulseTaskId, not even an explicit nil
### GetStartTimeUsecs

`func (o *CloudRetrieveTask) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *CloudRetrieveTask) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *CloudRetrieveTask) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *CloudRetrieveTask) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *CloudRetrieveTask) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *CloudRetrieveTask) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetStatus

`func (o *CloudRetrieveTask) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudRetrieveTask) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudRetrieveTask) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudRetrieveTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *CloudRetrieveTask) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *CloudRetrieveTask) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTimeRemainingSeconds

`func (o *CloudRetrieveTask) GetTimeRemainingSeconds() int64`

GetTimeRemainingSeconds returns the TimeRemainingSeconds field if non-nil, zero value otherwise.

### GetTimeRemainingSecondsOk

`func (o *CloudRetrieveTask) GetTimeRemainingSecondsOk() (*int64, bool)`

GetTimeRemainingSecondsOk returns a tuple with the TimeRemainingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRemainingSeconds

`func (o *CloudRetrieveTask) SetTimeRemainingSeconds(v int64)`

SetTimeRemainingSeconds sets TimeRemainingSeconds field to given value.

### HasTimeRemainingSeconds

`func (o *CloudRetrieveTask) HasTimeRemainingSeconds() bool`

HasTimeRemainingSeconds returns a boolean if a field has been set.

### SetTimeRemainingSecondsNil

`func (o *CloudRetrieveTask) SetTimeRemainingSecondsNil(b bool)`

 SetTimeRemainingSecondsNil sets the value for TimeRemainingSeconds to be an explicit nil

### UnsetTimeRemainingSeconds
`func (o *CloudRetrieveTask) UnsetTimeRemainingSeconds()`

UnsetTimeRemainingSeconds ensures that no value is present for TimeRemainingSeconds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


