# BackupAttempt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdmittedTimeUsecs** | Pointer to **NullableInt64** | Specifies the time at which the backup task was admitted to run in Unix epoch Timestamp(in microseconds) for an object. | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time of attempt in Unix epoch Timestamp(in microseconds) for an object. | [optional] 
**Message** | Pointer to **NullableString** | A message about the error if encountered while performing backup. | [optional] 
**PermitGrantTimeUsecs** | Pointer to **NullableInt64** | Specifies the time when gatekeeper permit is granted to the backup task. If the backup task is rescheduled due to errors, the field is updated to the time when permit is granted again. | [optional] 
**ProgressTaskId** | Pointer to **NullableString** | Progress monitor task for an object.. | [optional] 
**QueueDurationUsecs** | Pointer to **NullableInt64** | Specifies the duration between the startTime and when gatekeeper permit is granted to the backup task. If the backup task is rescheduled due to errors, the field is updated considering the time when permit is granted again. Queue duration &#x3D; PermitGrantTimeUsecs - StartTimeUsecs | [optional] 
**SnapshotCreationTimeUsecs** | Pointer to **NullableInt64** | Specifies the time at which the source snapshot was taken in Unix epoch Timestamp(in microseconds) for an object. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of attempt in Unix epoch Timestamp(in microseconds) for an object. | [optional] 
**Stats** | Pointer to [**BackupDataStats**](BackupDataStats.md) |  | [optional] 
**Status** | Pointer to **NullableString** | Status of the attempt for an object. &#39;Running&#39; indicates that the run is still running. &#39;Canceled&#39; indicates that the run has been canceled. &#39;Canceling&#39; indicates that the run is in the process of being canceled. &#39;Paused&#39; indicates that the ongoing run has been paused. &#39;Pausing&#39; indicates that the ongoing run is in the process of being paused. &#39;Resuming&#39; indicates that the already paused run is in the process of being running again. &#39;Failed&#39; indicates that the run has failed. &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening. &#39;Succeeded&#39; indicates that the run has finished successfully. &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages. &#39;Skipped&#39; indicates that the run was skipped. | [optional] 

## Methods

### NewBackupAttempt

`func NewBackupAttempt() *BackupAttempt`

NewBackupAttempt instantiates a new BackupAttempt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupAttemptWithDefaults

`func NewBackupAttemptWithDefaults() *BackupAttempt`

NewBackupAttemptWithDefaults instantiates a new BackupAttempt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmittedTimeUsecs

`func (o *BackupAttempt) GetAdmittedTimeUsecs() int64`

GetAdmittedTimeUsecs returns the AdmittedTimeUsecs field if non-nil, zero value otherwise.

### GetAdmittedTimeUsecsOk

`func (o *BackupAttempt) GetAdmittedTimeUsecsOk() (*int64, bool)`

GetAdmittedTimeUsecsOk returns a tuple with the AdmittedTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmittedTimeUsecs

`func (o *BackupAttempt) SetAdmittedTimeUsecs(v int64)`

SetAdmittedTimeUsecs sets AdmittedTimeUsecs field to given value.

### HasAdmittedTimeUsecs

`func (o *BackupAttempt) HasAdmittedTimeUsecs() bool`

HasAdmittedTimeUsecs returns a boolean if a field has been set.

### SetAdmittedTimeUsecsNil

`func (o *BackupAttempt) SetAdmittedTimeUsecsNil(b bool)`

 SetAdmittedTimeUsecsNil sets the value for AdmittedTimeUsecs to be an explicit nil

### UnsetAdmittedTimeUsecs
`func (o *BackupAttempt) UnsetAdmittedTimeUsecs()`

UnsetAdmittedTimeUsecs ensures that no value is present for AdmittedTimeUsecs, not even an explicit nil
### GetEndTimeUsecs

`func (o *BackupAttempt) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *BackupAttempt) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *BackupAttempt) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *BackupAttempt) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *BackupAttempt) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *BackupAttempt) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetMessage

`func (o *BackupAttempt) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BackupAttempt) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BackupAttempt) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BackupAttempt) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *BackupAttempt) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *BackupAttempt) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetPermitGrantTimeUsecs

`func (o *BackupAttempt) GetPermitGrantTimeUsecs() int64`

GetPermitGrantTimeUsecs returns the PermitGrantTimeUsecs field if non-nil, zero value otherwise.

### GetPermitGrantTimeUsecsOk

`func (o *BackupAttempt) GetPermitGrantTimeUsecsOk() (*int64, bool)`

GetPermitGrantTimeUsecsOk returns a tuple with the PermitGrantTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermitGrantTimeUsecs

`func (o *BackupAttempt) SetPermitGrantTimeUsecs(v int64)`

SetPermitGrantTimeUsecs sets PermitGrantTimeUsecs field to given value.

### HasPermitGrantTimeUsecs

`func (o *BackupAttempt) HasPermitGrantTimeUsecs() bool`

HasPermitGrantTimeUsecs returns a boolean if a field has been set.

### SetPermitGrantTimeUsecsNil

`func (o *BackupAttempt) SetPermitGrantTimeUsecsNil(b bool)`

 SetPermitGrantTimeUsecsNil sets the value for PermitGrantTimeUsecs to be an explicit nil

### UnsetPermitGrantTimeUsecs
`func (o *BackupAttempt) UnsetPermitGrantTimeUsecs()`

UnsetPermitGrantTimeUsecs ensures that no value is present for PermitGrantTimeUsecs, not even an explicit nil
### GetProgressTaskId

`func (o *BackupAttempt) GetProgressTaskId() string`

GetProgressTaskId returns the ProgressTaskId field if non-nil, zero value otherwise.

### GetProgressTaskIdOk

`func (o *BackupAttempt) GetProgressTaskIdOk() (*string, bool)`

GetProgressTaskIdOk returns a tuple with the ProgressTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressTaskId

`func (o *BackupAttempt) SetProgressTaskId(v string)`

SetProgressTaskId sets ProgressTaskId field to given value.

### HasProgressTaskId

`func (o *BackupAttempt) HasProgressTaskId() bool`

HasProgressTaskId returns a boolean if a field has been set.

### SetProgressTaskIdNil

`func (o *BackupAttempt) SetProgressTaskIdNil(b bool)`

 SetProgressTaskIdNil sets the value for ProgressTaskId to be an explicit nil

### UnsetProgressTaskId
`func (o *BackupAttempt) UnsetProgressTaskId()`

UnsetProgressTaskId ensures that no value is present for ProgressTaskId, not even an explicit nil
### GetQueueDurationUsecs

`func (o *BackupAttempt) GetQueueDurationUsecs() int64`

GetQueueDurationUsecs returns the QueueDurationUsecs field if non-nil, zero value otherwise.

### GetQueueDurationUsecsOk

`func (o *BackupAttempt) GetQueueDurationUsecsOk() (*int64, bool)`

GetQueueDurationUsecsOk returns a tuple with the QueueDurationUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueDurationUsecs

`func (o *BackupAttempt) SetQueueDurationUsecs(v int64)`

SetQueueDurationUsecs sets QueueDurationUsecs field to given value.

### HasQueueDurationUsecs

`func (o *BackupAttempt) HasQueueDurationUsecs() bool`

HasQueueDurationUsecs returns a boolean if a field has been set.

### SetQueueDurationUsecsNil

`func (o *BackupAttempt) SetQueueDurationUsecsNil(b bool)`

 SetQueueDurationUsecsNil sets the value for QueueDurationUsecs to be an explicit nil

### UnsetQueueDurationUsecs
`func (o *BackupAttempt) UnsetQueueDurationUsecs()`

UnsetQueueDurationUsecs ensures that no value is present for QueueDurationUsecs, not even an explicit nil
### GetSnapshotCreationTimeUsecs

`func (o *BackupAttempt) GetSnapshotCreationTimeUsecs() int64`

GetSnapshotCreationTimeUsecs returns the SnapshotCreationTimeUsecs field if non-nil, zero value otherwise.

### GetSnapshotCreationTimeUsecsOk

`func (o *BackupAttempt) GetSnapshotCreationTimeUsecsOk() (*int64, bool)`

GetSnapshotCreationTimeUsecsOk returns a tuple with the SnapshotCreationTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotCreationTimeUsecs

`func (o *BackupAttempt) SetSnapshotCreationTimeUsecs(v int64)`

SetSnapshotCreationTimeUsecs sets SnapshotCreationTimeUsecs field to given value.

### HasSnapshotCreationTimeUsecs

`func (o *BackupAttempt) HasSnapshotCreationTimeUsecs() bool`

HasSnapshotCreationTimeUsecs returns a boolean if a field has been set.

### SetSnapshotCreationTimeUsecsNil

`func (o *BackupAttempt) SetSnapshotCreationTimeUsecsNil(b bool)`

 SetSnapshotCreationTimeUsecsNil sets the value for SnapshotCreationTimeUsecs to be an explicit nil

### UnsetSnapshotCreationTimeUsecs
`func (o *BackupAttempt) UnsetSnapshotCreationTimeUsecs()`

UnsetSnapshotCreationTimeUsecs ensures that no value is present for SnapshotCreationTimeUsecs, not even an explicit nil
### GetStartTimeUsecs

`func (o *BackupAttempt) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *BackupAttempt) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *BackupAttempt) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *BackupAttempt) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *BackupAttempt) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *BackupAttempt) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetStats

`func (o *BackupAttempt) GetStats() BackupDataStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *BackupAttempt) GetStatsOk() (*BackupDataStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *BackupAttempt) SetStats(v BackupDataStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *BackupAttempt) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetStatus

`func (o *BackupAttempt) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BackupAttempt) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BackupAttempt) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BackupAttempt) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *BackupAttempt) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *BackupAttempt) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


