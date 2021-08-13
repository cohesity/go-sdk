# ArchivalTargetResultAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunType** | Pointer to **NullableString** | Type of Protection Group run. &#39;kRegular&#39; indicates an incremental (CBT) backup. Incremental backups utilizing CBT (if supported) are captured of the target protection objects. The first run of a kRegular schedule captures all the blocks. &#39;kFull&#39; indicates a full (no CBT) backup. A complete backup (all blocks) of the target protection objects are always captured and Change Block Tracking (CBT) is not utilized. &#39;kLog&#39; indicates a Database Log backup. Capture the database transaction logs to allow rolling back to a specific point in time. &#39;kSystem&#39; indicates system volume backup. It produces an image for bare metal recovery. | [optional] 
**IsSlaViolated** | Pointer to **NullableBool** | Indicated if SLA has been violated for this run. | [optional] 
**SnapshotId** | Pointer to **NullableString** | Snapshot id for a successful snapshot. This field will not be set if the archival Run fails to take the snapshot. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of replication run in Unix epoch Timestamp(in microseconds) for an archival target. | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time of replication run in Unix epoch Timestamp(in microseconds) for an archival target. | [optional] 
**QueuedTimeUsecs** | Pointer to **NullableInt64** | Specifies the time when the archival is queued for schedule in Unix epoch Timestamp(in microseconds) for a target. | [optional] 
**IsIncremental** | Pointer to **NullableBool** | Whether this is an incremental archive. If set to true, this is an incremental archive, otherwise this is a full archive. | [optional] 
**IsForeverIncremental** | Pointer to **NullableBool** | Whether this is forever incremental or not | [optional] 
**Status** | Pointer to **NullableString** | Status of the replication run for an archival target. &#39;Running&#39; indicates that the run is still running. &#39;Canceled&#39; indicates that the run has been canceled. &#39;Canceling&#39; indicates that the run is in the process of being canceled. &#39;Paused&#39; indicates that the ongoing run has been paused. &#39;Failed&#39; indicates that the run has failed. &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening. &#39;Succeeded&#39; indicates that the run has finished successfully. &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages. | [optional] 
**Message** | Pointer to **NullableString** | Message about the archival run. | [optional] 
**ProgressTaskId** | Pointer to **NullableString** | Progress monitor task id for archival. | [optional] 
**IndexingTaskId** | Pointer to **NullableString** | Progress monitor task for indexing. | [optional] 
**SuccessfulObjectsCount** | Pointer to **NullableInt64** | Specifies the count of objects for which backup was successful. | [optional] 
**FailedObjectsCount** | Pointer to **NullableInt64** | Specifies the count of objects for which backup failed. | [optional] 
**CancelledObjectsCount** | Pointer to **NullableInt64** | Specifies the count of objects for which backup was cancelled. | [optional] 
**SuccessfulAppObjectsCount** | Pointer to **NullableInt32** | Specifies the count of app objects for which backup was successful. | [optional] 
**FailedAppObjectsCount** | Pointer to **NullableInt32** | Specifies the count of app objects for which backup failed. | [optional] 
**CancelledAppObjectsCount** | Pointer to **NullableInt32** | Specifies the count of app objects for which backup was cancelled. | [optional] 
**Stats** | Pointer to [**ArchivalDataStats**](ArchivalDataStats.md) |  | [optional] 
**IsManuallyDeleted** | Pointer to **NullableBool** | Specifies whether the snapshot is deleted manually. | [optional] 
**ExpiryTimeUsecs** | Pointer to **NullableInt64** | Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds). | [optional] 
**DataLockConstraints** | Pointer to [**DataLockConstraints**](DataLockConstraints.md) |  | [optional] 
**OnLegalHold** | Pointer to **NullableBool** | Specifies the legal hold status for a archival target. | [optional] 

## Methods

### NewArchivalTargetResultAllOf

`func NewArchivalTargetResultAllOf() *ArchivalTargetResultAllOf`

NewArchivalTargetResultAllOf instantiates a new ArchivalTargetResultAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchivalTargetResultAllOfWithDefaults

`func NewArchivalTargetResultAllOfWithDefaults() *ArchivalTargetResultAllOf`

NewArchivalTargetResultAllOfWithDefaults instantiates a new ArchivalTargetResultAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunType

`func (o *ArchivalTargetResultAllOf) GetRunType() string`

GetRunType returns the RunType field if non-nil, zero value otherwise.

### GetRunTypeOk

`func (o *ArchivalTargetResultAllOf) GetRunTypeOk() (*string, bool)`

GetRunTypeOk returns a tuple with the RunType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunType

`func (o *ArchivalTargetResultAllOf) SetRunType(v string)`

SetRunType sets RunType field to given value.

### HasRunType

`func (o *ArchivalTargetResultAllOf) HasRunType() bool`

HasRunType returns a boolean if a field has been set.

### SetRunTypeNil

`func (o *ArchivalTargetResultAllOf) SetRunTypeNil(b bool)`

 SetRunTypeNil sets the value for RunType to be an explicit nil

### UnsetRunType
`func (o *ArchivalTargetResultAllOf) UnsetRunType()`

UnsetRunType ensures that no value is present for RunType, not even an explicit nil
### GetIsSlaViolated

`func (o *ArchivalTargetResultAllOf) GetIsSlaViolated() bool`

GetIsSlaViolated returns the IsSlaViolated field if non-nil, zero value otherwise.

### GetIsSlaViolatedOk

`func (o *ArchivalTargetResultAllOf) GetIsSlaViolatedOk() (*bool, bool)`

GetIsSlaViolatedOk returns a tuple with the IsSlaViolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSlaViolated

`func (o *ArchivalTargetResultAllOf) SetIsSlaViolated(v bool)`

SetIsSlaViolated sets IsSlaViolated field to given value.

### HasIsSlaViolated

`func (o *ArchivalTargetResultAllOf) HasIsSlaViolated() bool`

HasIsSlaViolated returns a boolean if a field has been set.

### SetIsSlaViolatedNil

`func (o *ArchivalTargetResultAllOf) SetIsSlaViolatedNil(b bool)`

 SetIsSlaViolatedNil sets the value for IsSlaViolated to be an explicit nil

### UnsetIsSlaViolated
`func (o *ArchivalTargetResultAllOf) UnsetIsSlaViolated()`

UnsetIsSlaViolated ensures that no value is present for IsSlaViolated, not even an explicit nil
### GetSnapshotId

`func (o *ArchivalTargetResultAllOf) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *ArchivalTargetResultAllOf) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *ArchivalTargetResultAllOf) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *ArchivalTargetResultAllOf) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### SetSnapshotIdNil

`func (o *ArchivalTargetResultAllOf) SetSnapshotIdNil(b bool)`

 SetSnapshotIdNil sets the value for SnapshotId to be an explicit nil

### UnsetSnapshotId
`func (o *ArchivalTargetResultAllOf) UnsetSnapshotId()`

UnsetSnapshotId ensures that no value is present for SnapshotId, not even an explicit nil
### GetStartTimeUsecs

`func (o *ArchivalTargetResultAllOf) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *ArchivalTargetResultAllOf) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *ArchivalTargetResultAllOf) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *ArchivalTargetResultAllOf) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *ArchivalTargetResultAllOf) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *ArchivalTargetResultAllOf) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetEndTimeUsecs

`func (o *ArchivalTargetResultAllOf) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *ArchivalTargetResultAllOf) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *ArchivalTargetResultAllOf) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *ArchivalTargetResultAllOf) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *ArchivalTargetResultAllOf) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *ArchivalTargetResultAllOf) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetQueuedTimeUsecs

`func (o *ArchivalTargetResultAllOf) GetQueuedTimeUsecs() int64`

GetQueuedTimeUsecs returns the QueuedTimeUsecs field if non-nil, zero value otherwise.

### GetQueuedTimeUsecsOk

`func (o *ArchivalTargetResultAllOf) GetQueuedTimeUsecsOk() (*int64, bool)`

GetQueuedTimeUsecsOk returns a tuple with the QueuedTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuedTimeUsecs

`func (o *ArchivalTargetResultAllOf) SetQueuedTimeUsecs(v int64)`

SetQueuedTimeUsecs sets QueuedTimeUsecs field to given value.

### HasQueuedTimeUsecs

`func (o *ArchivalTargetResultAllOf) HasQueuedTimeUsecs() bool`

HasQueuedTimeUsecs returns a boolean if a field has been set.

### SetQueuedTimeUsecsNil

`func (o *ArchivalTargetResultAllOf) SetQueuedTimeUsecsNil(b bool)`

 SetQueuedTimeUsecsNil sets the value for QueuedTimeUsecs to be an explicit nil

### UnsetQueuedTimeUsecs
`func (o *ArchivalTargetResultAllOf) UnsetQueuedTimeUsecs()`

UnsetQueuedTimeUsecs ensures that no value is present for QueuedTimeUsecs, not even an explicit nil
### GetIsIncremental

`func (o *ArchivalTargetResultAllOf) GetIsIncremental() bool`

GetIsIncremental returns the IsIncremental field if non-nil, zero value otherwise.

### GetIsIncrementalOk

`func (o *ArchivalTargetResultAllOf) GetIsIncrementalOk() (*bool, bool)`

GetIsIncrementalOk returns a tuple with the IsIncremental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIncremental

`func (o *ArchivalTargetResultAllOf) SetIsIncremental(v bool)`

SetIsIncremental sets IsIncremental field to given value.

### HasIsIncremental

`func (o *ArchivalTargetResultAllOf) HasIsIncremental() bool`

HasIsIncremental returns a boolean if a field has been set.

### SetIsIncrementalNil

`func (o *ArchivalTargetResultAllOf) SetIsIncrementalNil(b bool)`

 SetIsIncrementalNil sets the value for IsIncremental to be an explicit nil

### UnsetIsIncremental
`func (o *ArchivalTargetResultAllOf) UnsetIsIncremental()`

UnsetIsIncremental ensures that no value is present for IsIncremental, not even an explicit nil
### GetIsForeverIncremental

`func (o *ArchivalTargetResultAllOf) GetIsForeverIncremental() bool`

GetIsForeverIncremental returns the IsForeverIncremental field if non-nil, zero value otherwise.

### GetIsForeverIncrementalOk

`func (o *ArchivalTargetResultAllOf) GetIsForeverIncrementalOk() (*bool, bool)`

GetIsForeverIncrementalOk returns a tuple with the IsForeverIncremental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForeverIncremental

`func (o *ArchivalTargetResultAllOf) SetIsForeverIncremental(v bool)`

SetIsForeverIncremental sets IsForeverIncremental field to given value.

### HasIsForeverIncremental

`func (o *ArchivalTargetResultAllOf) HasIsForeverIncremental() bool`

HasIsForeverIncremental returns a boolean if a field has been set.

### SetIsForeverIncrementalNil

`func (o *ArchivalTargetResultAllOf) SetIsForeverIncrementalNil(b bool)`

 SetIsForeverIncrementalNil sets the value for IsForeverIncremental to be an explicit nil

### UnsetIsForeverIncremental
`func (o *ArchivalTargetResultAllOf) UnsetIsForeverIncremental()`

UnsetIsForeverIncremental ensures that no value is present for IsForeverIncremental, not even an explicit nil
### GetStatus

`func (o *ArchivalTargetResultAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ArchivalTargetResultAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ArchivalTargetResultAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ArchivalTargetResultAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ArchivalTargetResultAllOf) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ArchivalTargetResultAllOf) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetMessage

`func (o *ArchivalTargetResultAllOf) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ArchivalTargetResultAllOf) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ArchivalTargetResultAllOf) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ArchivalTargetResultAllOf) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ArchivalTargetResultAllOf) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ArchivalTargetResultAllOf) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetProgressTaskId

`func (o *ArchivalTargetResultAllOf) GetProgressTaskId() string`

GetProgressTaskId returns the ProgressTaskId field if non-nil, zero value otherwise.

### GetProgressTaskIdOk

`func (o *ArchivalTargetResultAllOf) GetProgressTaskIdOk() (*string, bool)`

GetProgressTaskIdOk returns a tuple with the ProgressTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressTaskId

`func (o *ArchivalTargetResultAllOf) SetProgressTaskId(v string)`

SetProgressTaskId sets ProgressTaskId field to given value.

### HasProgressTaskId

`func (o *ArchivalTargetResultAllOf) HasProgressTaskId() bool`

HasProgressTaskId returns a boolean if a field has been set.

### SetProgressTaskIdNil

`func (o *ArchivalTargetResultAllOf) SetProgressTaskIdNil(b bool)`

 SetProgressTaskIdNil sets the value for ProgressTaskId to be an explicit nil

### UnsetProgressTaskId
`func (o *ArchivalTargetResultAllOf) UnsetProgressTaskId()`

UnsetProgressTaskId ensures that no value is present for ProgressTaskId, not even an explicit nil
### GetIndexingTaskId

`func (o *ArchivalTargetResultAllOf) GetIndexingTaskId() string`

GetIndexingTaskId returns the IndexingTaskId field if non-nil, zero value otherwise.

### GetIndexingTaskIdOk

`func (o *ArchivalTargetResultAllOf) GetIndexingTaskIdOk() (*string, bool)`

GetIndexingTaskIdOk returns a tuple with the IndexingTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingTaskId

`func (o *ArchivalTargetResultAllOf) SetIndexingTaskId(v string)`

SetIndexingTaskId sets IndexingTaskId field to given value.

### HasIndexingTaskId

`func (o *ArchivalTargetResultAllOf) HasIndexingTaskId() bool`

HasIndexingTaskId returns a boolean if a field has been set.

### SetIndexingTaskIdNil

`func (o *ArchivalTargetResultAllOf) SetIndexingTaskIdNil(b bool)`

 SetIndexingTaskIdNil sets the value for IndexingTaskId to be an explicit nil

### UnsetIndexingTaskId
`func (o *ArchivalTargetResultAllOf) UnsetIndexingTaskId()`

UnsetIndexingTaskId ensures that no value is present for IndexingTaskId, not even an explicit nil
### GetSuccessfulObjectsCount

`func (o *ArchivalTargetResultAllOf) GetSuccessfulObjectsCount() int64`

GetSuccessfulObjectsCount returns the SuccessfulObjectsCount field if non-nil, zero value otherwise.

### GetSuccessfulObjectsCountOk

`func (o *ArchivalTargetResultAllOf) GetSuccessfulObjectsCountOk() (*int64, bool)`

GetSuccessfulObjectsCountOk returns a tuple with the SuccessfulObjectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulObjectsCount

`func (o *ArchivalTargetResultAllOf) SetSuccessfulObjectsCount(v int64)`

SetSuccessfulObjectsCount sets SuccessfulObjectsCount field to given value.

### HasSuccessfulObjectsCount

`func (o *ArchivalTargetResultAllOf) HasSuccessfulObjectsCount() bool`

HasSuccessfulObjectsCount returns a boolean if a field has been set.

### SetSuccessfulObjectsCountNil

`func (o *ArchivalTargetResultAllOf) SetSuccessfulObjectsCountNil(b bool)`

 SetSuccessfulObjectsCountNil sets the value for SuccessfulObjectsCount to be an explicit nil

### UnsetSuccessfulObjectsCount
`func (o *ArchivalTargetResultAllOf) UnsetSuccessfulObjectsCount()`

UnsetSuccessfulObjectsCount ensures that no value is present for SuccessfulObjectsCount, not even an explicit nil
### GetFailedObjectsCount

`func (o *ArchivalTargetResultAllOf) GetFailedObjectsCount() int64`

GetFailedObjectsCount returns the FailedObjectsCount field if non-nil, zero value otherwise.

### GetFailedObjectsCountOk

`func (o *ArchivalTargetResultAllOf) GetFailedObjectsCountOk() (*int64, bool)`

GetFailedObjectsCountOk returns a tuple with the FailedObjectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedObjectsCount

`func (o *ArchivalTargetResultAllOf) SetFailedObjectsCount(v int64)`

SetFailedObjectsCount sets FailedObjectsCount field to given value.

### HasFailedObjectsCount

`func (o *ArchivalTargetResultAllOf) HasFailedObjectsCount() bool`

HasFailedObjectsCount returns a boolean if a field has been set.

### SetFailedObjectsCountNil

`func (o *ArchivalTargetResultAllOf) SetFailedObjectsCountNil(b bool)`

 SetFailedObjectsCountNil sets the value for FailedObjectsCount to be an explicit nil

### UnsetFailedObjectsCount
`func (o *ArchivalTargetResultAllOf) UnsetFailedObjectsCount()`

UnsetFailedObjectsCount ensures that no value is present for FailedObjectsCount, not even an explicit nil
### GetCancelledObjectsCount

`func (o *ArchivalTargetResultAllOf) GetCancelledObjectsCount() int64`

GetCancelledObjectsCount returns the CancelledObjectsCount field if non-nil, zero value otherwise.

### GetCancelledObjectsCountOk

`func (o *ArchivalTargetResultAllOf) GetCancelledObjectsCountOk() (*int64, bool)`

GetCancelledObjectsCountOk returns a tuple with the CancelledObjectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledObjectsCount

`func (o *ArchivalTargetResultAllOf) SetCancelledObjectsCount(v int64)`

SetCancelledObjectsCount sets CancelledObjectsCount field to given value.

### HasCancelledObjectsCount

`func (o *ArchivalTargetResultAllOf) HasCancelledObjectsCount() bool`

HasCancelledObjectsCount returns a boolean if a field has been set.

### SetCancelledObjectsCountNil

`func (o *ArchivalTargetResultAllOf) SetCancelledObjectsCountNil(b bool)`

 SetCancelledObjectsCountNil sets the value for CancelledObjectsCount to be an explicit nil

### UnsetCancelledObjectsCount
`func (o *ArchivalTargetResultAllOf) UnsetCancelledObjectsCount()`

UnsetCancelledObjectsCount ensures that no value is present for CancelledObjectsCount, not even an explicit nil
### GetSuccessfulAppObjectsCount

`func (o *ArchivalTargetResultAllOf) GetSuccessfulAppObjectsCount() int32`

GetSuccessfulAppObjectsCount returns the SuccessfulAppObjectsCount field if non-nil, zero value otherwise.

### GetSuccessfulAppObjectsCountOk

`func (o *ArchivalTargetResultAllOf) GetSuccessfulAppObjectsCountOk() (*int32, bool)`

GetSuccessfulAppObjectsCountOk returns a tuple with the SuccessfulAppObjectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulAppObjectsCount

`func (o *ArchivalTargetResultAllOf) SetSuccessfulAppObjectsCount(v int32)`

SetSuccessfulAppObjectsCount sets SuccessfulAppObjectsCount field to given value.

### HasSuccessfulAppObjectsCount

`func (o *ArchivalTargetResultAllOf) HasSuccessfulAppObjectsCount() bool`

HasSuccessfulAppObjectsCount returns a boolean if a field has been set.

### SetSuccessfulAppObjectsCountNil

`func (o *ArchivalTargetResultAllOf) SetSuccessfulAppObjectsCountNil(b bool)`

 SetSuccessfulAppObjectsCountNil sets the value for SuccessfulAppObjectsCount to be an explicit nil

### UnsetSuccessfulAppObjectsCount
`func (o *ArchivalTargetResultAllOf) UnsetSuccessfulAppObjectsCount()`

UnsetSuccessfulAppObjectsCount ensures that no value is present for SuccessfulAppObjectsCount, not even an explicit nil
### GetFailedAppObjectsCount

`func (o *ArchivalTargetResultAllOf) GetFailedAppObjectsCount() int32`

GetFailedAppObjectsCount returns the FailedAppObjectsCount field if non-nil, zero value otherwise.

### GetFailedAppObjectsCountOk

`func (o *ArchivalTargetResultAllOf) GetFailedAppObjectsCountOk() (*int32, bool)`

GetFailedAppObjectsCountOk returns a tuple with the FailedAppObjectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAppObjectsCount

`func (o *ArchivalTargetResultAllOf) SetFailedAppObjectsCount(v int32)`

SetFailedAppObjectsCount sets FailedAppObjectsCount field to given value.

### HasFailedAppObjectsCount

`func (o *ArchivalTargetResultAllOf) HasFailedAppObjectsCount() bool`

HasFailedAppObjectsCount returns a boolean if a field has been set.

### SetFailedAppObjectsCountNil

`func (o *ArchivalTargetResultAllOf) SetFailedAppObjectsCountNil(b bool)`

 SetFailedAppObjectsCountNil sets the value for FailedAppObjectsCount to be an explicit nil

### UnsetFailedAppObjectsCount
`func (o *ArchivalTargetResultAllOf) UnsetFailedAppObjectsCount()`

UnsetFailedAppObjectsCount ensures that no value is present for FailedAppObjectsCount, not even an explicit nil
### GetCancelledAppObjectsCount

`func (o *ArchivalTargetResultAllOf) GetCancelledAppObjectsCount() int32`

GetCancelledAppObjectsCount returns the CancelledAppObjectsCount field if non-nil, zero value otherwise.

### GetCancelledAppObjectsCountOk

`func (o *ArchivalTargetResultAllOf) GetCancelledAppObjectsCountOk() (*int32, bool)`

GetCancelledAppObjectsCountOk returns a tuple with the CancelledAppObjectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledAppObjectsCount

`func (o *ArchivalTargetResultAllOf) SetCancelledAppObjectsCount(v int32)`

SetCancelledAppObjectsCount sets CancelledAppObjectsCount field to given value.

### HasCancelledAppObjectsCount

`func (o *ArchivalTargetResultAllOf) HasCancelledAppObjectsCount() bool`

HasCancelledAppObjectsCount returns a boolean if a field has been set.

### SetCancelledAppObjectsCountNil

`func (o *ArchivalTargetResultAllOf) SetCancelledAppObjectsCountNil(b bool)`

 SetCancelledAppObjectsCountNil sets the value for CancelledAppObjectsCount to be an explicit nil

### UnsetCancelledAppObjectsCount
`func (o *ArchivalTargetResultAllOf) UnsetCancelledAppObjectsCount()`

UnsetCancelledAppObjectsCount ensures that no value is present for CancelledAppObjectsCount, not even an explicit nil
### GetStats

`func (o *ArchivalTargetResultAllOf) GetStats() ArchivalDataStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ArchivalTargetResultAllOf) GetStatsOk() (*ArchivalDataStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ArchivalTargetResultAllOf) SetStats(v ArchivalDataStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ArchivalTargetResultAllOf) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetIsManuallyDeleted

`func (o *ArchivalTargetResultAllOf) GetIsManuallyDeleted() bool`

GetIsManuallyDeleted returns the IsManuallyDeleted field if non-nil, zero value otherwise.

### GetIsManuallyDeletedOk

`func (o *ArchivalTargetResultAllOf) GetIsManuallyDeletedOk() (*bool, bool)`

GetIsManuallyDeletedOk returns a tuple with the IsManuallyDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManuallyDeleted

`func (o *ArchivalTargetResultAllOf) SetIsManuallyDeleted(v bool)`

SetIsManuallyDeleted sets IsManuallyDeleted field to given value.

### HasIsManuallyDeleted

`func (o *ArchivalTargetResultAllOf) HasIsManuallyDeleted() bool`

HasIsManuallyDeleted returns a boolean if a field has been set.

### SetIsManuallyDeletedNil

`func (o *ArchivalTargetResultAllOf) SetIsManuallyDeletedNil(b bool)`

 SetIsManuallyDeletedNil sets the value for IsManuallyDeleted to be an explicit nil

### UnsetIsManuallyDeleted
`func (o *ArchivalTargetResultAllOf) UnsetIsManuallyDeleted()`

UnsetIsManuallyDeleted ensures that no value is present for IsManuallyDeleted, not even an explicit nil
### GetExpiryTimeUsecs

`func (o *ArchivalTargetResultAllOf) GetExpiryTimeUsecs() int64`

GetExpiryTimeUsecs returns the ExpiryTimeUsecs field if non-nil, zero value otherwise.

### GetExpiryTimeUsecsOk

`func (o *ArchivalTargetResultAllOf) GetExpiryTimeUsecsOk() (*int64, bool)`

GetExpiryTimeUsecsOk returns a tuple with the ExpiryTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimeUsecs

`func (o *ArchivalTargetResultAllOf) SetExpiryTimeUsecs(v int64)`

SetExpiryTimeUsecs sets ExpiryTimeUsecs field to given value.

### HasExpiryTimeUsecs

`func (o *ArchivalTargetResultAllOf) HasExpiryTimeUsecs() bool`

HasExpiryTimeUsecs returns a boolean if a field has been set.

### SetExpiryTimeUsecsNil

`func (o *ArchivalTargetResultAllOf) SetExpiryTimeUsecsNil(b bool)`

 SetExpiryTimeUsecsNil sets the value for ExpiryTimeUsecs to be an explicit nil

### UnsetExpiryTimeUsecs
`func (o *ArchivalTargetResultAllOf) UnsetExpiryTimeUsecs()`

UnsetExpiryTimeUsecs ensures that no value is present for ExpiryTimeUsecs, not even an explicit nil
### GetDataLockConstraints

`func (o *ArchivalTargetResultAllOf) GetDataLockConstraints() DataLockConstraints`

GetDataLockConstraints returns the DataLockConstraints field if non-nil, zero value otherwise.

### GetDataLockConstraintsOk

`func (o *ArchivalTargetResultAllOf) GetDataLockConstraintsOk() (*DataLockConstraints, bool)`

GetDataLockConstraintsOk returns a tuple with the DataLockConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLockConstraints

`func (o *ArchivalTargetResultAllOf) SetDataLockConstraints(v DataLockConstraints)`

SetDataLockConstraints sets DataLockConstraints field to given value.

### HasDataLockConstraints

`func (o *ArchivalTargetResultAllOf) HasDataLockConstraints() bool`

HasDataLockConstraints returns a boolean if a field has been set.

### GetOnLegalHold

`func (o *ArchivalTargetResultAllOf) GetOnLegalHold() bool`

GetOnLegalHold returns the OnLegalHold field if non-nil, zero value otherwise.

### GetOnLegalHoldOk

`func (o *ArchivalTargetResultAllOf) GetOnLegalHoldOk() (*bool, bool)`

GetOnLegalHoldOk returns a tuple with the OnLegalHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnLegalHold

`func (o *ArchivalTargetResultAllOf) SetOnLegalHold(v bool)`

SetOnLegalHold sets OnLegalHold field to given value.

### HasOnLegalHold

`func (o *ArchivalTargetResultAllOf) HasOnLegalHold() bool`

HasOnLegalHold returns a boolean if a field has been set.

### SetOnLegalHoldNil

`func (o *ArchivalTargetResultAllOf) SetOnLegalHoldNil(b bool)`

 SetOnLegalHoldNil sets the value for OnLegalHold to be an explicit nil

### UnsetOnLegalHold
`func (o *ArchivalTargetResultAllOf) UnsetOnLegalHold()`

UnsetOnLegalHold ensures that no value is present for OnLegalHold, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


