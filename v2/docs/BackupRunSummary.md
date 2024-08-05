# BackupRunSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancelledAppObjectsCount** | Pointer to **NullableInt32** | Specifies the count of app objects for which backup was cancelled. | [optional] 
**CancelledObjectsCount** | Pointer to **NullableInt64** | Specifies the count of objects for which backup was cancelled. | [optional] 
**DataLock** | Pointer to **NullableString** | This field is deprecated. Use DataLockConstraints field instead. | [optional] 
**DataLockConstraints** | Pointer to [**DataLockConstraints**](DataLockConstraints.md) |  | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time of backup run in Unix epoch Timestamp(in microseconds). | [optional] 
**FailedAppObjectsCount** | Pointer to **NullableInt32** | Specifies the count of app objects for which backup failed. | [optional] 
**FailedObjectsCount** | Pointer to **NullableInt64** | Specifies the count of objects for which backup failed. | [optional] 
**IndexingTaskId** | Pointer to **NullableString** | Progress monitor task for indexing. | [optional] 
**IsSlaViolated** | Pointer to **NullableBool** | Indicated if SLA has been violated for this run. | [optional] 
**LocalSnapshotStats** | Pointer to [**BackupDataStats**](BackupDataStats.md) |  | [optional] 
**LocalTaskId** | Pointer to **NullableString** | Task ID for a local protection run. | [optional] 
**Messages** | Pointer to **[]string** | Message about the backup run. | [optional] 
**ProgressTaskId** | Pointer to **NullableString** | Progress monitor task id for local backup run. | [optional] 
**RunType** | Pointer to **NullableString** | Type of Protection Group run. &#39;kRegular&#39; indicates an incremental (CBT) backup. Incremental backups utilizing CBT (if supported) are captured of the target protection objects. The first run of a kRegular schedule captures all the blocks. &#39;kFull&#39; indicates a full (no CBT) backup. A complete backup (all blocks) of the target protection objects are always captured and Change Block Tracking (CBT) is not utilized. &#39;kLog&#39; indicates a Database Log backup. Capture the database transaction logs to allow rolling back to a specific point in time. &#39;kSystem&#39; indicates system volume backup. It produces an image for bare metal recovery. &#39;kStorageArraySnapshot&#39; indicates storage array snapshot backup. | [optional] 
**SkippedObjectsCount** | Pointer to **NullableInt64** | Specifies the count of objects for which backup was skipped. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of backup run in Unix epoch Timestamp(in microseconds). | [optional] 
**StatsTaskId** | Pointer to **NullableString** | Stats task id for local backup run. | [optional] 
**Status** | Pointer to **NullableString** | Status of the backup run. &#39;Running&#39; indicates that the run is still running. &#39;Canceled&#39; indicates that the run has been canceled. &#39;Canceling&#39; indicates that the run is in the process of being canceled. &#39;Paused&#39; indicates that the ongoing run has been paused. &#39;Failed&#39; indicates that the run has failed. &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening. &#39;Succeeded&#39; indicates that the run has finished successfully. &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages. &#39;Skipped&#39; indicates that the run was skipped. | [optional] 
**SuccessfulAppObjectsCount** | Pointer to **NullableInt32** | Specifies the count of app objects for which backup was successful. | [optional] 
**SuccessfulObjectsCount** | Pointer to **NullableInt64** | Specifies the count of objects for which backup was successful. | [optional] 

## Methods

### NewBackupRunSummary

`func NewBackupRunSummary() *BackupRunSummary`

NewBackupRunSummary instantiates a new BackupRunSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRunSummaryWithDefaults

`func NewBackupRunSummaryWithDefaults() *BackupRunSummary`

NewBackupRunSummaryWithDefaults instantiates a new BackupRunSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelledAppObjectsCount

`func (o *BackupRunSummary) GetCancelledAppObjectsCount() int32`

GetCancelledAppObjectsCount returns the CancelledAppObjectsCount field if non-nil, zero value otherwise.

### GetCancelledAppObjectsCountOk

`func (o *BackupRunSummary) GetCancelledAppObjectsCountOk() (*int32, bool)`

GetCancelledAppObjectsCountOk returns a tuple with the CancelledAppObjectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledAppObjectsCount

`func (o *BackupRunSummary) SetCancelledAppObjectsCount(v int32)`

SetCancelledAppObjectsCount sets CancelledAppObjectsCount field to given value.

### HasCancelledAppObjectsCount

`func (o *BackupRunSummary) HasCancelledAppObjectsCount() bool`

HasCancelledAppObjectsCount returns a boolean if a field has been set.

### SetCancelledAppObjectsCountNil

`func (o *BackupRunSummary) SetCancelledAppObjectsCountNil(b bool)`

 SetCancelledAppObjectsCountNil sets the value for CancelledAppObjectsCount to be an explicit nil

### UnsetCancelledAppObjectsCount
`func (o *BackupRunSummary) UnsetCancelledAppObjectsCount()`

UnsetCancelledAppObjectsCount ensures that no value is present for CancelledAppObjectsCount, not even an explicit nil
### GetCancelledObjectsCount

`func (o *BackupRunSummary) GetCancelledObjectsCount() int64`

GetCancelledObjectsCount returns the CancelledObjectsCount field if non-nil, zero value otherwise.

### GetCancelledObjectsCountOk

`func (o *BackupRunSummary) GetCancelledObjectsCountOk() (*int64, bool)`

GetCancelledObjectsCountOk returns a tuple with the CancelledObjectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledObjectsCount

`func (o *BackupRunSummary) SetCancelledObjectsCount(v int64)`

SetCancelledObjectsCount sets CancelledObjectsCount field to given value.

### HasCancelledObjectsCount

`func (o *BackupRunSummary) HasCancelledObjectsCount() bool`

HasCancelledObjectsCount returns a boolean if a field has been set.

### SetCancelledObjectsCountNil

`func (o *BackupRunSummary) SetCancelledObjectsCountNil(b bool)`

 SetCancelledObjectsCountNil sets the value for CancelledObjectsCount to be an explicit nil

### UnsetCancelledObjectsCount
`func (o *BackupRunSummary) UnsetCancelledObjectsCount()`

UnsetCancelledObjectsCount ensures that no value is present for CancelledObjectsCount, not even an explicit nil
### GetDataLock

`func (o *BackupRunSummary) GetDataLock() string`

GetDataLock returns the DataLock field if non-nil, zero value otherwise.

### GetDataLockOk

`func (o *BackupRunSummary) GetDataLockOk() (*string, bool)`

GetDataLockOk returns a tuple with the DataLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLock

`func (o *BackupRunSummary) SetDataLock(v string)`

SetDataLock sets DataLock field to given value.

### HasDataLock

`func (o *BackupRunSummary) HasDataLock() bool`

HasDataLock returns a boolean if a field has been set.

### SetDataLockNil

`func (o *BackupRunSummary) SetDataLockNil(b bool)`

 SetDataLockNil sets the value for DataLock to be an explicit nil

### UnsetDataLock
`func (o *BackupRunSummary) UnsetDataLock()`

UnsetDataLock ensures that no value is present for DataLock, not even an explicit nil
### GetDataLockConstraints

`func (o *BackupRunSummary) GetDataLockConstraints() DataLockConstraints`

GetDataLockConstraints returns the DataLockConstraints field if non-nil, zero value otherwise.

### GetDataLockConstraintsOk

`func (o *BackupRunSummary) GetDataLockConstraintsOk() (*DataLockConstraints, bool)`

GetDataLockConstraintsOk returns a tuple with the DataLockConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLockConstraints

`func (o *BackupRunSummary) SetDataLockConstraints(v DataLockConstraints)`

SetDataLockConstraints sets DataLockConstraints field to given value.

### HasDataLockConstraints

`func (o *BackupRunSummary) HasDataLockConstraints() bool`

HasDataLockConstraints returns a boolean if a field has been set.

### GetEndTimeUsecs

`func (o *BackupRunSummary) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *BackupRunSummary) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *BackupRunSummary) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *BackupRunSummary) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *BackupRunSummary) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *BackupRunSummary) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetFailedAppObjectsCount

`func (o *BackupRunSummary) GetFailedAppObjectsCount() int32`

GetFailedAppObjectsCount returns the FailedAppObjectsCount field if non-nil, zero value otherwise.

### GetFailedAppObjectsCountOk

`func (o *BackupRunSummary) GetFailedAppObjectsCountOk() (*int32, bool)`

GetFailedAppObjectsCountOk returns a tuple with the FailedAppObjectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAppObjectsCount

`func (o *BackupRunSummary) SetFailedAppObjectsCount(v int32)`

SetFailedAppObjectsCount sets FailedAppObjectsCount field to given value.

### HasFailedAppObjectsCount

`func (o *BackupRunSummary) HasFailedAppObjectsCount() bool`

HasFailedAppObjectsCount returns a boolean if a field has been set.

### SetFailedAppObjectsCountNil

`func (o *BackupRunSummary) SetFailedAppObjectsCountNil(b bool)`

 SetFailedAppObjectsCountNil sets the value for FailedAppObjectsCount to be an explicit nil

### UnsetFailedAppObjectsCount
`func (o *BackupRunSummary) UnsetFailedAppObjectsCount()`

UnsetFailedAppObjectsCount ensures that no value is present for FailedAppObjectsCount, not even an explicit nil
### GetFailedObjectsCount

`func (o *BackupRunSummary) GetFailedObjectsCount() int64`

GetFailedObjectsCount returns the FailedObjectsCount field if non-nil, zero value otherwise.

### GetFailedObjectsCountOk

`func (o *BackupRunSummary) GetFailedObjectsCountOk() (*int64, bool)`

GetFailedObjectsCountOk returns a tuple with the FailedObjectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedObjectsCount

`func (o *BackupRunSummary) SetFailedObjectsCount(v int64)`

SetFailedObjectsCount sets FailedObjectsCount field to given value.

### HasFailedObjectsCount

`func (o *BackupRunSummary) HasFailedObjectsCount() bool`

HasFailedObjectsCount returns a boolean if a field has been set.

### SetFailedObjectsCountNil

`func (o *BackupRunSummary) SetFailedObjectsCountNil(b bool)`

 SetFailedObjectsCountNil sets the value for FailedObjectsCount to be an explicit nil

### UnsetFailedObjectsCount
`func (o *BackupRunSummary) UnsetFailedObjectsCount()`

UnsetFailedObjectsCount ensures that no value is present for FailedObjectsCount, not even an explicit nil
### GetIndexingTaskId

`func (o *BackupRunSummary) GetIndexingTaskId() string`

GetIndexingTaskId returns the IndexingTaskId field if non-nil, zero value otherwise.

### GetIndexingTaskIdOk

`func (o *BackupRunSummary) GetIndexingTaskIdOk() (*string, bool)`

GetIndexingTaskIdOk returns a tuple with the IndexingTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingTaskId

`func (o *BackupRunSummary) SetIndexingTaskId(v string)`

SetIndexingTaskId sets IndexingTaskId field to given value.

### HasIndexingTaskId

`func (o *BackupRunSummary) HasIndexingTaskId() bool`

HasIndexingTaskId returns a boolean if a field has been set.

### SetIndexingTaskIdNil

`func (o *BackupRunSummary) SetIndexingTaskIdNil(b bool)`

 SetIndexingTaskIdNil sets the value for IndexingTaskId to be an explicit nil

### UnsetIndexingTaskId
`func (o *BackupRunSummary) UnsetIndexingTaskId()`

UnsetIndexingTaskId ensures that no value is present for IndexingTaskId, not even an explicit nil
### GetIsSlaViolated

`func (o *BackupRunSummary) GetIsSlaViolated() bool`

GetIsSlaViolated returns the IsSlaViolated field if non-nil, zero value otherwise.

### GetIsSlaViolatedOk

`func (o *BackupRunSummary) GetIsSlaViolatedOk() (*bool, bool)`

GetIsSlaViolatedOk returns a tuple with the IsSlaViolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSlaViolated

`func (o *BackupRunSummary) SetIsSlaViolated(v bool)`

SetIsSlaViolated sets IsSlaViolated field to given value.

### HasIsSlaViolated

`func (o *BackupRunSummary) HasIsSlaViolated() bool`

HasIsSlaViolated returns a boolean if a field has been set.

### SetIsSlaViolatedNil

`func (o *BackupRunSummary) SetIsSlaViolatedNil(b bool)`

 SetIsSlaViolatedNil sets the value for IsSlaViolated to be an explicit nil

### UnsetIsSlaViolated
`func (o *BackupRunSummary) UnsetIsSlaViolated()`

UnsetIsSlaViolated ensures that no value is present for IsSlaViolated, not even an explicit nil
### GetLocalSnapshotStats

`func (o *BackupRunSummary) GetLocalSnapshotStats() BackupDataStats`

GetLocalSnapshotStats returns the LocalSnapshotStats field if non-nil, zero value otherwise.

### GetLocalSnapshotStatsOk

`func (o *BackupRunSummary) GetLocalSnapshotStatsOk() (*BackupDataStats, bool)`

GetLocalSnapshotStatsOk returns a tuple with the LocalSnapshotStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSnapshotStats

`func (o *BackupRunSummary) SetLocalSnapshotStats(v BackupDataStats)`

SetLocalSnapshotStats sets LocalSnapshotStats field to given value.

### HasLocalSnapshotStats

`func (o *BackupRunSummary) HasLocalSnapshotStats() bool`

HasLocalSnapshotStats returns a boolean if a field has been set.

### GetLocalTaskId

`func (o *BackupRunSummary) GetLocalTaskId() string`

GetLocalTaskId returns the LocalTaskId field if non-nil, zero value otherwise.

### GetLocalTaskIdOk

`func (o *BackupRunSummary) GetLocalTaskIdOk() (*string, bool)`

GetLocalTaskIdOk returns a tuple with the LocalTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalTaskId

`func (o *BackupRunSummary) SetLocalTaskId(v string)`

SetLocalTaskId sets LocalTaskId field to given value.

### HasLocalTaskId

`func (o *BackupRunSummary) HasLocalTaskId() bool`

HasLocalTaskId returns a boolean if a field has been set.

### SetLocalTaskIdNil

`func (o *BackupRunSummary) SetLocalTaskIdNil(b bool)`

 SetLocalTaskIdNil sets the value for LocalTaskId to be an explicit nil

### UnsetLocalTaskId
`func (o *BackupRunSummary) UnsetLocalTaskId()`

UnsetLocalTaskId ensures that no value is present for LocalTaskId, not even an explicit nil
### GetMessages

`func (o *BackupRunSummary) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *BackupRunSummary) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *BackupRunSummary) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *BackupRunSummary) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### SetMessagesNil

`func (o *BackupRunSummary) SetMessagesNil(b bool)`

 SetMessagesNil sets the value for Messages to be an explicit nil

### UnsetMessages
`func (o *BackupRunSummary) UnsetMessages()`

UnsetMessages ensures that no value is present for Messages, not even an explicit nil
### GetProgressTaskId

`func (o *BackupRunSummary) GetProgressTaskId() string`

GetProgressTaskId returns the ProgressTaskId field if non-nil, zero value otherwise.

### GetProgressTaskIdOk

`func (o *BackupRunSummary) GetProgressTaskIdOk() (*string, bool)`

GetProgressTaskIdOk returns a tuple with the ProgressTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressTaskId

`func (o *BackupRunSummary) SetProgressTaskId(v string)`

SetProgressTaskId sets ProgressTaskId field to given value.

### HasProgressTaskId

`func (o *BackupRunSummary) HasProgressTaskId() bool`

HasProgressTaskId returns a boolean if a field has been set.

### SetProgressTaskIdNil

`func (o *BackupRunSummary) SetProgressTaskIdNil(b bool)`

 SetProgressTaskIdNil sets the value for ProgressTaskId to be an explicit nil

### UnsetProgressTaskId
`func (o *BackupRunSummary) UnsetProgressTaskId()`

UnsetProgressTaskId ensures that no value is present for ProgressTaskId, not even an explicit nil
### GetRunType

`func (o *BackupRunSummary) GetRunType() string`

GetRunType returns the RunType field if non-nil, zero value otherwise.

### GetRunTypeOk

`func (o *BackupRunSummary) GetRunTypeOk() (*string, bool)`

GetRunTypeOk returns a tuple with the RunType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunType

`func (o *BackupRunSummary) SetRunType(v string)`

SetRunType sets RunType field to given value.

### HasRunType

`func (o *BackupRunSummary) HasRunType() bool`

HasRunType returns a boolean if a field has been set.

### SetRunTypeNil

`func (o *BackupRunSummary) SetRunTypeNil(b bool)`

 SetRunTypeNil sets the value for RunType to be an explicit nil

### UnsetRunType
`func (o *BackupRunSummary) UnsetRunType()`

UnsetRunType ensures that no value is present for RunType, not even an explicit nil
### GetSkippedObjectsCount

`func (o *BackupRunSummary) GetSkippedObjectsCount() int64`

GetSkippedObjectsCount returns the SkippedObjectsCount field if non-nil, zero value otherwise.

### GetSkippedObjectsCountOk

`func (o *BackupRunSummary) GetSkippedObjectsCountOk() (*int64, bool)`

GetSkippedObjectsCountOk returns a tuple with the SkippedObjectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedObjectsCount

`func (o *BackupRunSummary) SetSkippedObjectsCount(v int64)`

SetSkippedObjectsCount sets SkippedObjectsCount field to given value.

### HasSkippedObjectsCount

`func (o *BackupRunSummary) HasSkippedObjectsCount() bool`

HasSkippedObjectsCount returns a boolean if a field has been set.

### SetSkippedObjectsCountNil

`func (o *BackupRunSummary) SetSkippedObjectsCountNil(b bool)`

 SetSkippedObjectsCountNil sets the value for SkippedObjectsCount to be an explicit nil

### UnsetSkippedObjectsCount
`func (o *BackupRunSummary) UnsetSkippedObjectsCount()`

UnsetSkippedObjectsCount ensures that no value is present for SkippedObjectsCount, not even an explicit nil
### GetStartTimeUsecs

`func (o *BackupRunSummary) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *BackupRunSummary) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *BackupRunSummary) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *BackupRunSummary) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *BackupRunSummary) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *BackupRunSummary) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetStatsTaskId

`func (o *BackupRunSummary) GetStatsTaskId() string`

GetStatsTaskId returns the StatsTaskId field if non-nil, zero value otherwise.

### GetStatsTaskIdOk

`func (o *BackupRunSummary) GetStatsTaskIdOk() (*string, bool)`

GetStatsTaskIdOk returns a tuple with the StatsTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsTaskId

`func (o *BackupRunSummary) SetStatsTaskId(v string)`

SetStatsTaskId sets StatsTaskId field to given value.

### HasStatsTaskId

`func (o *BackupRunSummary) HasStatsTaskId() bool`

HasStatsTaskId returns a boolean if a field has been set.

### SetStatsTaskIdNil

`func (o *BackupRunSummary) SetStatsTaskIdNil(b bool)`

 SetStatsTaskIdNil sets the value for StatsTaskId to be an explicit nil

### UnsetStatsTaskId
`func (o *BackupRunSummary) UnsetStatsTaskId()`

UnsetStatsTaskId ensures that no value is present for StatsTaskId, not even an explicit nil
### GetStatus

`func (o *BackupRunSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BackupRunSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BackupRunSummary) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BackupRunSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *BackupRunSummary) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *BackupRunSummary) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetSuccessfulAppObjectsCount

`func (o *BackupRunSummary) GetSuccessfulAppObjectsCount() int32`

GetSuccessfulAppObjectsCount returns the SuccessfulAppObjectsCount field if non-nil, zero value otherwise.

### GetSuccessfulAppObjectsCountOk

`func (o *BackupRunSummary) GetSuccessfulAppObjectsCountOk() (*int32, bool)`

GetSuccessfulAppObjectsCountOk returns a tuple with the SuccessfulAppObjectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulAppObjectsCount

`func (o *BackupRunSummary) SetSuccessfulAppObjectsCount(v int32)`

SetSuccessfulAppObjectsCount sets SuccessfulAppObjectsCount field to given value.

### HasSuccessfulAppObjectsCount

`func (o *BackupRunSummary) HasSuccessfulAppObjectsCount() bool`

HasSuccessfulAppObjectsCount returns a boolean if a field has been set.

### SetSuccessfulAppObjectsCountNil

`func (o *BackupRunSummary) SetSuccessfulAppObjectsCountNil(b bool)`

 SetSuccessfulAppObjectsCountNil sets the value for SuccessfulAppObjectsCount to be an explicit nil

### UnsetSuccessfulAppObjectsCount
`func (o *BackupRunSummary) UnsetSuccessfulAppObjectsCount()`

UnsetSuccessfulAppObjectsCount ensures that no value is present for SuccessfulAppObjectsCount, not even an explicit nil
### GetSuccessfulObjectsCount

`func (o *BackupRunSummary) GetSuccessfulObjectsCount() int64`

GetSuccessfulObjectsCount returns the SuccessfulObjectsCount field if non-nil, zero value otherwise.

### GetSuccessfulObjectsCountOk

`func (o *BackupRunSummary) GetSuccessfulObjectsCountOk() (*int64, bool)`

GetSuccessfulObjectsCountOk returns a tuple with the SuccessfulObjectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulObjectsCount

`func (o *BackupRunSummary) SetSuccessfulObjectsCount(v int64)`

SetSuccessfulObjectsCount sets SuccessfulObjectsCount field to given value.

### HasSuccessfulObjectsCount

`func (o *BackupRunSummary) HasSuccessfulObjectsCount() bool`

HasSuccessfulObjectsCount returns a boolean if a field has been set.

### SetSuccessfulObjectsCountNil

`func (o *BackupRunSummary) SetSuccessfulObjectsCountNil(b bool)`

 SetSuccessfulObjectsCountNil sets the value for SuccessfulObjectsCount to be an explicit nil

### UnsetSuccessfulObjectsCount
`func (o *BackupRunSummary) UnsetSuccessfulObjectsCount()`

UnsetSuccessfulObjectsCount ensures that no value is present for SuccessfulObjectsCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


