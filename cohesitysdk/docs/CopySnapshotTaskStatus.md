# CopySnapshotTaskStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **NullableString** | Specifies if an error occurred (if any) while running this task. This field is populated when the status is equal to &#39;kFailure&#39;. | [optional] 
**Source** | Pointer to [**ProtectionSource**](ProtectionSource.md) |  | [optional] 
**Stats** | Pointer to [**CopyRunStats**](CopyRunStats.md) |  | [optional] 
**Status** | Pointer to **NullableString** | Specifies the status of the source object being protected. &#39;kAccepted&#39; indicates the task is queued to run but not yet running. &#39;kRunning&#39; indicates the task is running. &#39;kCanceling&#39; indicates a request to cancel the task has occurred but the task is not yet canceled. &#39;kCanceled&#39; indicates the task has been canceled. &#39;kSuccess&#39; indicates the task was successful. &#39;kFailure&#39; indicates the task failed. &#39;kWarning&#39; indicates the task has finished with warning. &#39;kOnHold&#39; indicates the task is kept onHold. &#39;kMissed&#39; indicates the task is missed. | [optional] 
**TaskEndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time of the copy task. The end time is specified as a Unix epoch Timestamp (in microseconds). | [optional] 
**TaskStartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of the copy task. The start time is specified as a Unix epoch Timestamp (in microseconds). Copy run task is started after completing backup tasks. It may spawn sub-tasks to copy or replicate individual snapshots. | [optional] 

## Methods

### NewCopySnapshotTaskStatus

`func NewCopySnapshotTaskStatus() *CopySnapshotTaskStatus`

NewCopySnapshotTaskStatus instantiates a new CopySnapshotTaskStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCopySnapshotTaskStatusWithDefaults

`func NewCopySnapshotTaskStatusWithDefaults() *CopySnapshotTaskStatus`

NewCopySnapshotTaskStatusWithDefaults instantiates a new CopySnapshotTaskStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *CopySnapshotTaskStatus) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CopySnapshotTaskStatus) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CopySnapshotTaskStatus) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CopySnapshotTaskStatus) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *CopySnapshotTaskStatus) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *CopySnapshotTaskStatus) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetSource

`func (o *CopySnapshotTaskStatus) GetSource() ProtectionSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CopySnapshotTaskStatus) GetSourceOk() (*ProtectionSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CopySnapshotTaskStatus) SetSource(v ProtectionSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *CopySnapshotTaskStatus) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStats

`func (o *CopySnapshotTaskStatus) GetStats() CopyRunStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *CopySnapshotTaskStatus) GetStatsOk() (*CopyRunStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *CopySnapshotTaskStatus) SetStats(v CopyRunStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *CopySnapshotTaskStatus) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetStatus

`func (o *CopySnapshotTaskStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CopySnapshotTaskStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CopySnapshotTaskStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CopySnapshotTaskStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *CopySnapshotTaskStatus) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *CopySnapshotTaskStatus) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTaskEndTimeUsecs

`func (o *CopySnapshotTaskStatus) GetTaskEndTimeUsecs() int64`

GetTaskEndTimeUsecs returns the TaskEndTimeUsecs field if non-nil, zero value otherwise.

### GetTaskEndTimeUsecsOk

`func (o *CopySnapshotTaskStatus) GetTaskEndTimeUsecsOk() (*int64, bool)`

GetTaskEndTimeUsecsOk returns a tuple with the TaskEndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskEndTimeUsecs

`func (o *CopySnapshotTaskStatus) SetTaskEndTimeUsecs(v int64)`

SetTaskEndTimeUsecs sets TaskEndTimeUsecs field to given value.

### HasTaskEndTimeUsecs

`func (o *CopySnapshotTaskStatus) HasTaskEndTimeUsecs() bool`

HasTaskEndTimeUsecs returns a boolean if a field has been set.

### SetTaskEndTimeUsecsNil

`func (o *CopySnapshotTaskStatus) SetTaskEndTimeUsecsNil(b bool)`

 SetTaskEndTimeUsecsNil sets the value for TaskEndTimeUsecs to be an explicit nil

### UnsetTaskEndTimeUsecs
`func (o *CopySnapshotTaskStatus) UnsetTaskEndTimeUsecs()`

UnsetTaskEndTimeUsecs ensures that no value is present for TaskEndTimeUsecs, not even an explicit nil
### GetTaskStartTimeUsecs

`func (o *CopySnapshotTaskStatus) GetTaskStartTimeUsecs() int64`

GetTaskStartTimeUsecs returns the TaskStartTimeUsecs field if non-nil, zero value otherwise.

### GetTaskStartTimeUsecsOk

`func (o *CopySnapshotTaskStatus) GetTaskStartTimeUsecsOk() (*int64, bool)`

GetTaskStartTimeUsecsOk returns a tuple with the TaskStartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskStartTimeUsecs

`func (o *CopySnapshotTaskStatus) SetTaskStartTimeUsecs(v int64)`

SetTaskStartTimeUsecs sets TaskStartTimeUsecs field to given value.

### HasTaskStartTimeUsecs

`func (o *CopySnapshotTaskStatus) HasTaskStartTimeUsecs() bool`

HasTaskStartTimeUsecs returns a boolean if a field has been set.

### SetTaskStartTimeUsecsNil

`func (o *CopySnapshotTaskStatus) SetTaskStartTimeUsecsNil(b bool)`

 SetTaskStartTimeUsecsNil sets the value for TaskStartTimeUsecs to be an explicit nil

### UnsetTaskStartTimeUsecs
`func (o *CopySnapshotTaskStatus) UnsetTaskStartTimeUsecs()`

UnsetTaskStartTimeUsecs ensures that no value is present for TaskStartTimeUsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


