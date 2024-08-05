# RecoveryTaskInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time in Unix timestamp epoch in microseconds. | [optional] [readonly] 
**ProgressTaskId** | Pointer to **NullableString** | Specifies the progress monitor path. | [optional] [readonly] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time in Unix timestamp epoch in microseconds. | [optional] [readonly] 
**Status** | Pointer to **NullableString** | Specifies the status of the recovery. | [optional] [readonly] 

## Methods

### NewRecoveryTaskInfo

`func NewRecoveryTaskInfo() *RecoveryTaskInfo`

NewRecoveryTaskInfo instantiates a new RecoveryTaskInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryTaskInfoWithDefaults

`func NewRecoveryTaskInfoWithDefaults() *RecoveryTaskInfo`

NewRecoveryTaskInfoWithDefaults instantiates a new RecoveryTaskInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTimeUsecs

`func (o *RecoveryTaskInfo) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *RecoveryTaskInfo) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *RecoveryTaskInfo) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *RecoveryTaskInfo) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *RecoveryTaskInfo) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *RecoveryTaskInfo) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetProgressTaskId

`func (o *RecoveryTaskInfo) GetProgressTaskId() string`

GetProgressTaskId returns the ProgressTaskId field if non-nil, zero value otherwise.

### GetProgressTaskIdOk

`func (o *RecoveryTaskInfo) GetProgressTaskIdOk() (*string, bool)`

GetProgressTaskIdOk returns a tuple with the ProgressTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressTaskId

`func (o *RecoveryTaskInfo) SetProgressTaskId(v string)`

SetProgressTaskId sets ProgressTaskId field to given value.

### HasProgressTaskId

`func (o *RecoveryTaskInfo) HasProgressTaskId() bool`

HasProgressTaskId returns a boolean if a field has been set.

### SetProgressTaskIdNil

`func (o *RecoveryTaskInfo) SetProgressTaskIdNil(b bool)`

 SetProgressTaskIdNil sets the value for ProgressTaskId to be an explicit nil

### UnsetProgressTaskId
`func (o *RecoveryTaskInfo) UnsetProgressTaskId()`

UnsetProgressTaskId ensures that no value is present for ProgressTaskId, not even an explicit nil
### GetStartTimeUsecs

`func (o *RecoveryTaskInfo) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *RecoveryTaskInfo) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *RecoveryTaskInfo) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *RecoveryTaskInfo) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *RecoveryTaskInfo) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *RecoveryTaskInfo) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetStatus

`func (o *RecoveryTaskInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RecoveryTaskInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RecoveryTaskInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RecoveryTaskInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *RecoveryTaskInfo) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *RecoveryTaskInfo) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


