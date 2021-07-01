# Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**[]TaskAttribute**](TaskAttribute.md) | The latest attributes reported for this task. | [optional] 
**EndTimeSeconds** | Pointer to **NullableInt64** | Specifies the end time of the task. | [optional] 
**ErrorMessage** | Pointer to **NullableString** | Specifies an optional error message for this task. | [optional] 
**Events** | Pointer to [**[]TaskEvent**](TaskEvent.md) | Specifies the events reported for this task. | [optional] 
**ExpectedEndTimeSeconds** | Pointer to **NullableInt64** | Specifies the estimated end time of the task. | [optional] 
**ExpectedSecondsRemaining** | Pointer to **NullableInt64** | Specifies the expected remaining time for this task in seconds. | [optional] 
**ExpectedTotalWorkCount** | Pointer to **NullableInt64** | The expected raw count of the total work remaining. This is the highest work count value reported by the client. This field can be set to let pulse compute percentFinished by looking at the currently reported remainingWorkCount and the expectedTotalWorkCount. | [optional] 
**LastUpdateTimeSeconds** | Pointer to **NullableInt64** | Specifies the timestamp when the last progress was reported. | [optional] 
**PercentFinished** | Pointer to **NullableFloat32** | Specifies the reported progress on the task. | [optional] 
**StartTimeSeconds** | Pointer to **NullableInt64** | Specifies the start time of the task. | [optional] 
**Status** | Pointer to **NullableString** | Specifies the status of the task being queried. &#39;kActive&#39; indicates that the task is still active. &#39;kFinished&#39; indicates that the task has finished without any errors. &#39;kFinishedWithError&#39; indicates that the task has finished, but that there was an errror of some kind. &#39;kCancelled&#39; indicates that the task was cancelled. &#39;kFinishedGarbageCollected&#39; indicates that the task was garbage collected due to its subtasks not finishing within the allotted time. | [optional] 
**SubTasks** | Pointer to **[]map[string]interface{}** | Specifies a list of subtasks belonging to this task. | [optional] 
**TaskPath** | Pointer to **NullableString** | Specifes the path of this task. | [optional] 

## Methods

### NewTask

`func NewTask() *Task`

NewTask instantiates a new Task object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWithDefaults

`func NewTaskWithDefaults() *Task`

NewTaskWithDefaults instantiates a new Task object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *Task) GetAttributes() []TaskAttribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Task) GetAttributesOk() (*[]TaskAttribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Task) SetAttributes(v []TaskAttribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Task) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *Task) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *Task) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetEndTimeSeconds

`func (o *Task) GetEndTimeSeconds() int64`

GetEndTimeSeconds returns the EndTimeSeconds field if non-nil, zero value otherwise.

### GetEndTimeSecondsOk

`func (o *Task) GetEndTimeSecondsOk() (*int64, bool)`

GetEndTimeSecondsOk returns a tuple with the EndTimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeSeconds

`func (o *Task) SetEndTimeSeconds(v int64)`

SetEndTimeSeconds sets EndTimeSeconds field to given value.

### HasEndTimeSeconds

`func (o *Task) HasEndTimeSeconds() bool`

HasEndTimeSeconds returns a boolean if a field has been set.

### SetEndTimeSecondsNil

`func (o *Task) SetEndTimeSecondsNil(b bool)`

 SetEndTimeSecondsNil sets the value for EndTimeSeconds to be an explicit nil

### UnsetEndTimeSeconds
`func (o *Task) UnsetEndTimeSeconds()`

UnsetEndTimeSeconds ensures that no value is present for EndTimeSeconds, not even an explicit nil
### GetErrorMessage

`func (o *Task) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *Task) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *Task) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *Task) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *Task) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *Task) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetEvents

`func (o *Task) GetEvents() []TaskEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Task) GetEventsOk() (*[]TaskEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Task) SetEvents(v []TaskEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *Task) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *Task) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *Task) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetExpectedEndTimeSeconds

`func (o *Task) GetExpectedEndTimeSeconds() int64`

GetExpectedEndTimeSeconds returns the ExpectedEndTimeSeconds field if non-nil, zero value otherwise.

### GetExpectedEndTimeSecondsOk

`func (o *Task) GetExpectedEndTimeSecondsOk() (*int64, bool)`

GetExpectedEndTimeSecondsOk returns a tuple with the ExpectedEndTimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedEndTimeSeconds

`func (o *Task) SetExpectedEndTimeSeconds(v int64)`

SetExpectedEndTimeSeconds sets ExpectedEndTimeSeconds field to given value.

### HasExpectedEndTimeSeconds

`func (o *Task) HasExpectedEndTimeSeconds() bool`

HasExpectedEndTimeSeconds returns a boolean if a field has been set.

### SetExpectedEndTimeSecondsNil

`func (o *Task) SetExpectedEndTimeSecondsNil(b bool)`

 SetExpectedEndTimeSecondsNil sets the value for ExpectedEndTimeSeconds to be an explicit nil

### UnsetExpectedEndTimeSeconds
`func (o *Task) UnsetExpectedEndTimeSeconds()`

UnsetExpectedEndTimeSeconds ensures that no value is present for ExpectedEndTimeSeconds, not even an explicit nil
### GetExpectedSecondsRemaining

`func (o *Task) GetExpectedSecondsRemaining() int64`

GetExpectedSecondsRemaining returns the ExpectedSecondsRemaining field if non-nil, zero value otherwise.

### GetExpectedSecondsRemainingOk

`func (o *Task) GetExpectedSecondsRemainingOk() (*int64, bool)`

GetExpectedSecondsRemainingOk returns a tuple with the ExpectedSecondsRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedSecondsRemaining

`func (o *Task) SetExpectedSecondsRemaining(v int64)`

SetExpectedSecondsRemaining sets ExpectedSecondsRemaining field to given value.

### HasExpectedSecondsRemaining

`func (o *Task) HasExpectedSecondsRemaining() bool`

HasExpectedSecondsRemaining returns a boolean if a field has been set.

### SetExpectedSecondsRemainingNil

`func (o *Task) SetExpectedSecondsRemainingNil(b bool)`

 SetExpectedSecondsRemainingNil sets the value for ExpectedSecondsRemaining to be an explicit nil

### UnsetExpectedSecondsRemaining
`func (o *Task) UnsetExpectedSecondsRemaining()`

UnsetExpectedSecondsRemaining ensures that no value is present for ExpectedSecondsRemaining, not even an explicit nil
### GetExpectedTotalWorkCount

`func (o *Task) GetExpectedTotalWorkCount() int64`

GetExpectedTotalWorkCount returns the ExpectedTotalWorkCount field if non-nil, zero value otherwise.

### GetExpectedTotalWorkCountOk

`func (o *Task) GetExpectedTotalWorkCountOk() (*int64, bool)`

GetExpectedTotalWorkCountOk returns a tuple with the ExpectedTotalWorkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedTotalWorkCount

`func (o *Task) SetExpectedTotalWorkCount(v int64)`

SetExpectedTotalWorkCount sets ExpectedTotalWorkCount field to given value.

### HasExpectedTotalWorkCount

`func (o *Task) HasExpectedTotalWorkCount() bool`

HasExpectedTotalWorkCount returns a boolean if a field has been set.

### SetExpectedTotalWorkCountNil

`func (o *Task) SetExpectedTotalWorkCountNil(b bool)`

 SetExpectedTotalWorkCountNil sets the value for ExpectedTotalWorkCount to be an explicit nil

### UnsetExpectedTotalWorkCount
`func (o *Task) UnsetExpectedTotalWorkCount()`

UnsetExpectedTotalWorkCount ensures that no value is present for ExpectedTotalWorkCount, not even an explicit nil
### GetLastUpdateTimeSeconds

`func (o *Task) GetLastUpdateTimeSeconds() int64`

GetLastUpdateTimeSeconds returns the LastUpdateTimeSeconds field if non-nil, zero value otherwise.

### GetLastUpdateTimeSecondsOk

`func (o *Task) GetLastUpdateTimeSecondsOk() (*int64, bool)`

GetLastUpdateTimeSecondsOk returns a tuple with the LastUpdateTimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimeSeconds

`func (o *Task) SetLastUpdateTimeSeconds(v int64)`

SetLastUpdateTimeSeconds sets LastUpdateTimeSeconds field to given value.

### HasLastUpdateTimeSeconds

`func (o *Task) HasLastUpdateTimeSeconds() bool`

HasLastUpdateTimeSeconds returns a boolean if a field has been set.

### SetLastUpdateTimeSecondsNil

`func (o *Task) SetLastUpdateTimeSecondsNil(b bool)`

 SetLastUpdateTimeSecondsNil sets the value for LastUpdateTimeSeconds to be an explicit nil

### UnsetLastUpdateTimeSeconds
`func (o *Task) UnsetLastUpdateTimeSeconds()`

UnsetLastUpdateTimeSeconds ensures that no value is present for LastUpdateTimeSeconds, not even an explicit nil
### GetPercentFinished

`func (o *Task) GetPercentFinished() float32`

GetPercentFinished returns the PercentFinished field if non-nil, zero value otherwise.

### GetPercentFinishedOk

`func (o *Task) GetPercentFinishedOk() (*float32, bool)`

GetPercentFinishedOk returns a tuple with the PercentFinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentFinished

`func (o *Task) SetPercentFinished(v float32)`

SetPercentFinished sets PercentFinished field to given value.

### HasPercentFinished

`func (o *Task) HasPercentFinished() bool`

HasPercentFinished returns a boolean if a field has been set.

### SetPercentFinishedNil

`func (o *Task) SetPercentFinishedNil(b bool)`

 SetPercentFinishedNil sets the value for PercentFinished to be an explicit nil

### UnsetPercentFinished
`func (o *Task) UnsetPercentFinished()`

UnsetPercentFinished ensures that no value is present for PercentFinished, not even an explicit nil
### GetStartTimeSeconds

`func (o *Task) GetStartTimeSeconds() int64`

GetStartTimeSeconds returns the StartTimeSeconds field if non-nil, zero value otherwise.

### GetStartTimeSecondsOk

`func (o *Task) GetStartTimeSecondsOk() (*int64, bool)`

GetStartTimeSecondsOk returns a tuple with the StartTimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeSeconds

`func (o *Task) SetStartTimeSeconds(v int64)`

SetStartTimeSeconds sets StartTimeSeconds field to given value.

### HasStartTimeSeconds

`func (o *Task) HasStartTimeSeconds() bool`

HasStartTimeSeconds returns a boolean if a field has been set.

### SetStartTimeSecondsNil

`func (o *Task) SetStartTimeSecondsNil(b bool)`

 SetStartTimeSecondsNil sets the value for StartTimeSeconds to be an explicit nil

### UnsetStartTimeSeconds
`func (o *Task) UnsetStartTimeSeconds()`

UnsetStartTimeSeconds ensures that no value is present for StartTimeSeconds, not even an explicit nil
### GetStatus

`func (o *Task) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Task) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Task) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Task) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *Task) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Task) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetSubTasks

`func (o *Task) GetSubTasks() []map[string]interface{}`

GetSubTasks returns the SubTasks field if non-nil, zero value otherwise.

### GetSubTasksOk

`func (o *Task) GetSubTasksOk() (*[]map[string]interface{}, bool)`

GetSubTasksOk returns a tuple with the SubTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTasks

`func (o *Task) SetSubTasks(v []map[string]interface{})`

SetSubTasks sets SubTasks field to given value.

### HasSubTasks

`func (o *Task) HasSubTasks() bool`

HasSubTasks returns a boolean if a field has been set.

### SetSubTasksNil

`func (o *Task) SetSubTasksNil(b bool)`

 SetSubTasksNil sets the value for SubTasks to be an explicit nil

### UnsetSubTasks
`func (o *Task) UnsetSubTasks()`

UnsetSubTasks ensures that no value is present for SubTasks, not even an explicit nil
### GetTaskPath

`func (o *Task) GetTaskPath() string`

GetTaskPath returns the TaskPath field if non-nil, zero value otherwise.

### GetTaskPathOk

`func (o *Task) GetTaskPathOk() (*string, bool)`

GetTaskPathOk returns a tuple with the TaskPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskPath

`func (o *Task) SetTaskPath(v string)`

SetTaskPath sets TaskPath field to given value.

### HasTaskPath

`func (o *Task) HasTaskPath() bool`

HasTaskPath returns a boolean if a field has been set.

### SetTaskPathNil

`func (o *Task) SetTaskPathNil(b bool)`

 SetTaskPathNil sets the value for TaskPath to be an explicit nil

### UnsetTaskPath
`func (o *Task) UnsetTaskPath()`

UnsetTaskPath ensures that no value is present for TaskPath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


