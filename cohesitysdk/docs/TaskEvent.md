# TaskEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventMessage** | Pointer to **NullableString** | Specifies the message associated with an event. | [optional] 
**PercentFinished** | Pointer to **NullableFloat32** | Specifies the completion percentage of the task attached to this event. | [optional] 
**RemainingWorkCount** | Pointer to **NullableInt64** | Specifies the amount of work remaining for the task attached to this event. | [optional] 
**TimestampSeconds** | Pointer to **NullableInt64** | Specifies the Unix timestamp that the event occurred. | [optional] 

## Methods

### NewTaskEvent

`func NewTaskEvent() *TaskEvent`

NewTaskEvent instantiates a new TaskEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskEventWithDefaults

`func NewTaskEventWithDefaults() *TaskEvent`

NewTaskEventWithDefaults instantiates a new TaskEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventMessage

`func (o *TaskEvent) GetEventMessage() string`

GetEventMessage returns the EventMessage field if non-nil, zero value otherwise.

### GetEventMessageOk

`func (o *TaskEvent) GetEventMessageOk() (*string, bool)`

GetEventMessageOk returns a tuple with the EventMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventMessage

`func (o *TaskEvent) SetEventMessage(v string)`

SetEventMessage sets EventMessage field to given value.

### HasEventMessage

`func (o *TaskEvent) HasEventMessage() bool`

HasEventMessage returns a boolean if a field has been set.

### SetEventMessageNil

`func (o *TaskEvent) SetEventMessageNil(b bool)`

 SetEventMessageNil sets the value for EventMessage to be an explicit nil

### UnsetEventMessage
`func (o *TaskEvent) UnsetEventMessage()`

UnsetEventMessage ensures that no value is present for EventMessage, not even an explicit nil
### GetPercentFinished

`func (o *TaskEvent) GetPercentFinished() float32`

GetPercentFinished returns the PercentFinished field if non-nil, zero value otherwise.

### GetPercentFinishedOk

`func (o *TaskEvent) GetPercentFinishedOk() (*float32, bool)`

GetPercentFinishedOk returns a tuple with the PercentFinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentFinished

`func (o *TaskEvent) SetPercentFinished(v float32)`

SetPercentFinished sets PercentFinished field to given value.

### HasPercentFinished

`func (o *TaskEvent) HasPercentFinished() bool`

HasPercentFinished returns a boolean if a field has been set.

### SetPercentFinishedNil

`func (o *TaskEvent) SetPercentFinishedNil(b bool)`

 SetPercentFinishedNil sets the value for PercentFinished to be an explicit nil

### UnsetPercentFinished
`func (o *TaskEvent) UnsetPercentFinished()`

UnsetPercentFinished ensures that no value is present for PercentFinished, not even an explicit nil
### GetRemainingWorkCount

`func (o *TaskEvent) GetRemainingWorkCount() int64`

GetRemainingWorkCount returns the RemainingWorkCount field if non-nil, zero value otherwise.

### GetRemainingWorkCountOk

`func (o *TaskEvent) GetRemainingWorkCountOk() (*int64, bool)`

GetRemainingWorkCountOk returns a tuple with the RemainingWorkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingWorkCount

`func (o *TaskEvent) SetRemainingWorkCount(v int64)`

SetRemainingWorkCount sets RemainingWorkCount field to given value.

### HasRemainingWorkCount

`func (o *TaskEvent) HasRemainingWorkCount() bool`

HasRemainingWorkCount returns a boolean if a field has been set.

### SetRemainingWorkCountNil

`func (o *TaskEvent) SetRemainingWorkCountNil(b bool)`

 SetRemainingWorkCountNil sets the value for RemainingWorkCount to be an explicit nil

### UnsetRemainingWorkCount
`func (o *TaskEvent) UnsetRemainingWorkCount()`

UnsetRemainingWorkCount ensures that no value is present for RemainingWorkCount, not even an explicit nil
### GetTimestampSeconds

`func (o *TaskEvent) GetTimestampSeconds() int64`

GetTimestampSeconds returns the TimestampSeconds field if non-nil, zero value otherwise.

### GetTimestampSecondsOk

`func (o *TaskEvent) GetTimestampSecondsOk() (*int64, bool)`

GetTimestampSecondsOk returns a tuple with the TimestampSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampSeconds

`func (o *TaskEvent) SetTimestampSeconds(v int64)`

SetTimestampSeconds sets TimestampSeconds field to given value.

### HasTimestampSeconds

`func (o *TaskEvent) HasTimestampSeconds() bool`

HasTimestampSeconds returns a boolean if a field has been set.

### SetTimestampSecondsNil

`func (o *TaskEvent) SetTimestampSecondsNil(b bool)`

 SetTimestampSecondsNil sets the value for TimestampSeconds to be an explicit nil

### UnsetTimestampSeconds
`func (o *TaskEvent) UnsetTimestampSeconds()`

UnsetTimestampSeconds ensures that no value is present for TimestampSeconds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


