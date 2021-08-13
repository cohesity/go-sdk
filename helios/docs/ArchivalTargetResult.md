# ArchivalTargetResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetId** | Pointer to **NullableInt64** | Specifies the archival target ID. | [optional] 
**ArchivalTaskId** | Pointer to **NullableString** | Specifies the archival task id. This is a protection group UID which only applies when archival type is &#39;Tape&#39;. | [optional] 
**TargetName** | Pointer to **NullableString** | Specifies the archival target name. | [optional] 
**TargetType** | Pointer to **NullableString** | Specifies the archival target type. | [optional] 
**TierSettings** | Pointer to [**ArchivalTargetTierInfo**](ArchivalTargetTierInfo.md) |  | [optional] 
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

### NewArchivalTargetResult

`func NewArchivalTargetResult() *ArchivalTargetResult`

NewArchivalTargetResult instantiates a new ArchivalTargetResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchivalTargetResultWithDefaults

`func NewArchivalTargetResultWithDefaults() *ArchivalTargetResult`

NewArchivalTargetResultWithDefaults instantiates a new ArchivalTargetResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetId

`func (o *ArchivalTargetResult) GetTargetId() int64`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *ArchivalTargetResult) GetTargetIdOk() (*int64, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *ArchivalTargetResult) SetTargetId(v int64)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *ArchivalTargetResult) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *ArchivalTargetResult) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *ArchivalTargetResult) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetArchivalTaskId

`func (o *ArchivalTargetResult) GetArchivalTaskId() string`

GetArchivalTaskId returns the ArchivalTaskId field if non-nil, zero value otherwise.

### GetArchivalTaskIdOk

`func (o *ArchivalTargetResult) GetArchivalTaskIdOk() (*string, bool)`

GetArchivalTaskIdOk returns a tuple with the ArchivalTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalTaskId

`func (o *ArchivalTargetResult) SetArchivalTaskId(v string)`

SetArchivalTaskId sets ArchivalTaskId field to given value.

### HasArchivalTaskId

`func (o *ArchivalTargetResult) HasArchivalTaskId() bool`

HasArchivalTaskId returns a boolean if a field has been set.

### SetArchivalTaskIdNil

`func (o *ArchivalTargetResult) SetArchivalTaskIdNil(b bool)`

 SetArchivalTaskIdNil sets the value for ArchivalTaskId to be an explicit nil

### UnsetArchivalTaskId
`func (o *ArchivalTargetResult) UnsetArchivalTaskId()`

UnsetArchivalTaskId ensures that no value is present for ArchivalTaskId, not even an explicit nil
### GetTargetName

`func (o *ArchivalTargetResult) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *ArchivalTargetResult) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *ArchivalTargetResult) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *ArchivalTargetResult) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### SetTargetNameNil

`func (o *ArchivalTargetResult) SetTargetNameNil(b bool)`

 SetTargetNameNil sets the value for TargetName to be an explicit nil

### UnsetTargetName
`func (o *ArchivalTargetResult) UnsetTargetName()`

UnsetTargetName ensures that no value is present for TargetName, not even an explicit nil
### GetTargetType

`func (o *ArchivalTargetResult) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *ArchivalTargetResult) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *ArchivalTargetResult) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *ArchivalTargetResult) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### SetTargetTypeNil

`func (o *ArchivalTargetResult) SetTargetTypeNil(b bool)`

 SetTargetTypeNil sets the value for TargetType to be an explicit nil

### UnsetTargetType
`func (o *ArchivalTargetResult) UnsetTargetType()`

UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil
### GetTierSettings

`func (o *ArchivalTargetResult) GetTierSettings() ArchivalTargetTierInfo`

GetTierSettings returns the TierSettings field if non-nil, zero value otherwise.

### GetTierSettingsOk

`func (o *ArchivalTargetResult) GetTierSettingsOk() (*ArchivalTargetTierInfo, bool)`

GetTierSettingsOk returns a tuple with the TierSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierSettings

`func (o *ArchivalTargetResult) SetTierSettings(v ArchivalTargetTierInfo)`

SetTierSettings sets TierSettings field to given value.

### HasTierSettings

`func (o *ArchivalTargetResult) HasTierSettings() bool`

HasTierSettings returns a boolean if a field has been set.

### GetRunType

`func (o *ArchivalTargetResult) GetRunType() string`

GetRunType returns the RunType field if non-nil, zero value otherwise.

### GetRunTypeOk

`func (o *ArchivalTargetResult) GetRunTypeOk() (*string, bool)`

GetRunTypeOk returns a tuple with the RunType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunType

`func (o *ArchivalTargetResult) SetRunType(v string)`

SetRunType sets RunType field to given value.

### HasRunType

`func (o *ArchivalTargetResult) HasRunType() bool`

HasRunType returns a boolean if a field has been set.

### SetRunTypeNil

`func (o *ArchivalTargetResult) SetRunTypeNil(b bool)`

 SetRunTypeNil sets the value for RunType to be an explicit nil

### UnsetRunType
`func (o *ArchivalTargetResult) UnsetRunType()`

UnsetRunType ensures that no value is present for RunType, not even an explicit nil
### GetIsSlaViolated

`func (o *ArchivalTargetResult) GetIsSlaViolated() bool`

GetIsSlaViolated returns the IsSlaViolated field if non-nil, zero value otherwise.

### GetIsSlaViolatedOk

`func (o *ArchivalTargetResult) GetIsSlaViolatedOk() (*bool, bool)`

GetIsSlaViolatedOk returns a tuple with the IsSlaViolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSlaViolated

`func (o *ArchivalTargetResult) SetIsSlaViolated(v bool)`

SetIsSlaViolated sets IsSlaViolated field to given value.

### HasIsSlaViolated

`func (o *ArchivalTargetResult) HasIsSlaViolated() bool`

HasIsSlaViolated returns a boolean if a field has been set.

### SetIsSlaViolatedNil

`func (o *ArchivalTargetResult) SetIsSlaViolatedNil(b bool)`

 SetIsSlaViolatedNil sets the value for IsSlaViolated to be an explicit nil

### UnsetIsSlaViolated
`func (o *ArchivalTargetResult) UnsetIsSlaViolated()`

UnsetIsSlaViolated ensures that no value is present for IsSlaViolated, not even an explicit nil
### GetSnapshotId

`func (o *ArchivalTargetResult) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *ArchivalTargetResult) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *ArchivalTargetResult) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *ArchivalTargetResult) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### SetSnapshotIdNil

`func (o *ArchivalTargetResult) SetSnapshotIdNil(b bool)`

 SetSnapshotIdNil sets the value for SnapshotId to be an explicit nil

### UnsetSnapshotId
`func (o *ArchivalTargetResult) UnsetSnapshotId()`

UnsetSnapshotId ensures that no value is present for SnapshotId, not even an explicit nil
### GetStartTimeUsecs

`func (o *ArchivalTargetResult) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *ArchivalTargetResult) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *ArchivalTargetResult) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *ArchivalTargetResult) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *ArchivalTargetResult) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *ArchivalTargetResult) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetEndTimeUsecs

`func (o *ArchivalTargetResult) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *ArchivalTargetResult) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *ArchivalTargetResult) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *ArchivalTargetResult) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *ArchivalTargetResult) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *ArchivalTargetResult) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetQueuedTimeUsecs

`func (o *ArchivalTargetResult) GetQueuedTimeUsecs() int64`

GetQueuedTimeUsecs returns the QueuedTimeUsecs field if non-nil, zero value otherwise.

### GetQueuedTimeUsecsOk

`func (o *ArchivalTargetResult) GetQueuedTimeUsecsOk() (*int64, bool)`

GetQueuedTimeUsecsOk returns a tuple with the QueuedTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuedTimeUsecs

`func (o *ArchivalTargetResult) SetQueuedTimeUsecs(v int64)`

SetQueuedTimeUsecs sets QueuedTimeUsecs field to given value.

### HasQueuedTimeUsecs

`func (o *ArchivalTargetResult) HasQueuedTimeUsecs() bool`

HasQueuedTimeUsecs returns a boolean if a field has been set.

### SetQueuedTimeUsecsNil

`func (o *ArchivalTargetResult) SetQueuedTimeUsecsNil(b bool)`

 SetQueuedTimeUsecsNil sets the value for QueuedTimeUsecs to be an explicit nil

### UnsetQueuedTimeUsecs
`func (o *ArchivalTargetResult) UnsetQueuedTimeUsecs()`

UnsetQueuedTimeUsecs ensures that no value is present for QueuedTimeUsecs, not even an explicit nil
### GetIsIncremental

`func (o *ArchivalTargetResult) GetIsIncremental() bool`

GetIsIncremental returns the IsIncremental field if non-nil, zero value otherwise.

### GetIsIncrementalOk

`func (o *ArchivalTargetResult) GetIsIncrementalOk() (*bool, bool)`

GetIsIncrementalOk returns a tuple with the IsIncremental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIncremental

`func (o *ArchivalTargetResult) SetIsIncremental(v bool)`

SetIsIncremental sets IsIncremental field to given value.

### HasIsIncremental

`func (o *ArchivalTargetResult) HasIsIncremental() bool`

HasIsIncremental returns a boolean if a field has been set.

### SetIsIncrementalNil

`func (o *ArchivalTargetResult) SetIsIncrementalNil(b bool)`

 SetIsIncrementalNil sets the value for IsIncremental to be an explicit nil

### UnsetIsIncremental
`func (o *ArchivalTargetResult) UnsetIsIncremental()`

UnsetIsIncremental ensures that no value is present for IsIncremental, not even an explicit nil
### GetIsForeverIncremental

`func (o *ArchivalTargetResult) GetIsForeverIncremental() bool`

GetIsForeverIncremental returns the IsForeverIncremental field if non-nil, zero value otherwise.

### GetIsForeverIncrementalOk

`func (o *ArchivalTargetResult) GetIsForeverIncrementalOk() (*bool, bool)`

GetIsForeverIncrementalOk returns a tuple with the IsForeverIncremental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForeverIncremental

`func (o *ArchivalTargetResult) SetIsForeverIncremental(v bool)`

SetIsForeverIncremental sets IsForeverIncremental field to given value.

### HasIsForeverIncremental

`func (o *ArchivalTargetResult) HasIsForeverIncremental() bool`

HasIsForeverIncremental returns a boolean if a field has been set.

### SetIsForeverIncrementalNil

`func (o *ArchivalTargetResult) SetIsForeverIncrementalNil(b bool)`

 SetIsForeverIncrementalNil sets the value for IsForeverIncremental to be an explicit nil

### UnsetIsForeverIncremental
`func (o *ArchivalTargetResult) UnsetIsForeverIncremental()`

UnsetIsForeverIncremental ensures that no value is present for IsForeverIncremental, not even an explicit nil
### GetStatus

`func (o *ArchivalTargetResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ArchivalTargetResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ArchivalTargetResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ArchivalTargetResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ArchivalTargetResult) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ArchivalTargetResult) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetMessage

`func (o *ArchivalTargetResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ArchivalTargetResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ArchivalTargetResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ArchivalTargetResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ArchivalTargetResult) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ArchivalTargetResult) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetProgressTaskId

`func (o *ArchivalTargetResult) GetProgressTaskId() string`

GetProgressTaskId returns the ProgressTaskId field if non-nil, zero value otherwise.

### GetProgressTaskIdOk

`func (o *ArchivalTargetResult) GetProgressTaskIdOk() (*string, bool)`

GetProgressTaskIdOk returns a tuple with the ProgressTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressTaskId

`func (o *ArchivalTargetResult) SetProgressTaskId(v string)`

SetProgressTaskId sets ProgressTaskId field to given value.

### HasProgressTaskId

`func (o *ArchivalTargetResult) HasProgressTaskId() bool`

HasProgressTaskId returns a boolean if a field has been set.

### SetProgressTaskIdNil

`func (o *ArchivalTargetResult) SetProgressTaskIdNil(b bool)`

 SetProgressTaskIdNil sets the value for ProgressTaskId to be an explicit nil

### UnsetProgressTaskId
`func (o *ArchivalTargetResult) UnsetProgressTaskId()`

UnsetProgressTaskId ensures that no value is present for ProgressTaskId, not even an explicit nil
### GetIndexingTaskId

`func (o *ArchivalTargetResult) GetIndexingTaskId() string`

GetIndexingTaskId returns the IndexingTaskId field if non-nil, zero value otherwise.

### GetIndexingTaskIdOk

`func (o *ArchivalTargetResult) GetIndexingTaskIdOk() (*string, bool)`

GetIndexingTaskIdOk returns a tuple with the IndexingTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingTaskId

`func (o *ArchivalTargetResult) SetIndexingTaskId(v string)`

SetIndexingTaskId sets IndexingTaskId field to given value.

### HasIndexingTaskId

`func (o *ArchivalTargetResult) HasIndexingTaskId() bool`

HasIndexingTaskId returns a boolean if a field has been set.

### SetIndexingTaskIdNil

`func (o *ArchivalTargetResult) SetIndexingTaskIdNil(b bool)`

 SetIndexingTaskIdNil sets the value for IndexingTaskId to be an explicit nil

### UnsetIndexingTaskId
`func (o *ArchivalTargetResult) UnsetIndexingTaskId()`

UnsetIndexingTaskId ensures that no value is present for IndexingTaskId, not even an explicit nil
### GetSuccessfulObjectsCount

`func (o *ArchivalTargetResult) GetSuccessfulObjectsCount() int64`

GetSuccessfulObjectsCount returns the SuccessfulObjectsCount field if non-nil, zero value otherwise.

### GetSuccessfulObjectsCountOk

`func (o *ArchivalTargetResult) GetSuccessfulObjectsCountOk() (*int64, bool)`

GetSuccessfulObjectsCountOk returns a tuple with the SuccessfulObjectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulObjectsCount

`func (o *ArchivalTargetResult) SetSuccessfulObjectsCount(v int64)`

SetSuccessfulObjectsCount sets SuccessfulObjectsCount field to given value.

### HasSuccessfulObjectsCount

`func (o *ArchivalTargetResult) HasSuccessfulObjectsCount() bool`

HasSuccessfulObjectsCount returns a boolean if a field has been set.

### SetSuccessfulObjectsCountNil

`func (o *ArchivalTargetResult) SetSuccessfulObjectsCountNil(b bool)`

 SetSuccessfulObjectsCountNil sets the value for SuccessfulObjectsCount to be an explicit nil

### UnsetSuccessfulObjectsCount
`func (o *ArchivalTargetResult) UnsetSuccessfulObjectsCount()`

UnsetSuccessfulObjectsCount ensures that no value is present for SuccessfulObjectsCount, not even an explicit nil
### GetFailedObjectsCount

`func (o *ArchivalTargetResult) GetFailedObjectsCount() int64`

GetFailedObjectsCount returns the FailedObjectsCount field if non-nil, zero value otherwise.

### GetFailedObjectsCountOk

`func (o *ArchivalTargetResult) GetFailedObjectsCountOk() (*int64, bool)`

GetFailedObjectsCountOk returns a tuple with the FailedObjectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedObjectsCount

`func (o *ArchivalTargetResult) SetFailedObjectsCount(v int64)`

SetFailedObjectsCount sets FailedObjectsCount field to given value.

### HasFailedObjectsCount

`func (o *ArchivalTargetResult) HasFailedObjectsCount() bool`

HasFailedObjectsCount returns a boolean if a field has been set.

### SetFailedObjectsCountNil

`func (o *ArchivalTargetResult) SetFailedObjectsCountNil(b bool)`

 SetFailedObjectsCountNil sets the value for FailedObjectsCount to be an explicit nil

### UnsetFailedObjectsCount
`func (o *ArchivalTargetResult) UnsetFailedObjectsCount()`

UnsetFailedObjectsCount ensures that no value is present for FailedObjectsCount, not even an explicit nil
### GetCancelledObjectsCount

`func (o *ArchivalTargetResult) GetCancelledObjectsCount() int64`

GetCancelledObjectsCount returns the CancelledObjectsCount field if non-nil, zero value otherwise.

### GetCancelledObjectsCountOk

`func (o *ArchivalTargetResult) GetCancelledObjectsCountOk() (*int64, bool)`

GetCancelledObjectsCountOk returns a tuple with the CancelledObjectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledObjectsCount

`func (o *ArchivalTargetResult) SetCancelledObjectsCount(v int64)`

SetCancelledObjectsCount sets CancelledObjectsCount field to given value.

### HasCancelledObjectsCount

`func (o *ArchivalTargetResult) HasCancelledObjectsCount() bool`

HasCancelledObjectsCount returns a boolean if a field has been set.

### SetCancelledObjectsCountNil

`func (o *ArchivalTargetResult) SetCancelledObjectsCountNil(b bool)`

 SetCancelledObjectsCountNil sets the value for CancelledObjectsCount to be an explicit nil

### UnsetCancelledObjectsCount
`func (o *ArchivalTargetResult) UnsetCancelledObjectsCount()`

UnsetCancelledObjectsCount ensures that no value is present for CancelledObjectsCount, not even an explicit nil
### GetSuccessfulAppObjectsCount

`func (o *ArchivalTargetResult) GetSuccessfulAppObjectsCount() int32`

GetSuccessfulAppObjectsCount returns the SuccessfulAppObjectsCount field if non-nil, zero value otherwise.

### GetSuccessfulAppObjectsCountOk

`func (o *ArchivalTargetResult) GetSuccessfulAppObjectsCountOk() (*int32, bool)`

GetSuccessfulAppObjectsCountOk returns a tuple with the SuccessfulAppObjectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulAppObjectsCount

`func (o *ArchivalTargetResult) SetSuccessfulAppObjectsCount(v int32)`

SetSuccessfulAppObjectsCount sets SuccessfulAppObjectsCount field to given value.

### HasSuccessfulAppObjectsCount

`func (o *ArchivalTargetResult) HasSuccessfulAppObjectsCount() bool`

HasSuccessfulAppObjectsCount returns a boolean if a field has been set.

### SetSuccessfulAppObjectsCountNil

`func (o *ArchivalTargetResult) SetSuccessfulAppObjectsCountNil(b bool)`

 SetSuccessfulAppObjectsCountNil sets the value for SuccessfulAppObjectsCount to be an explicit nil

### UnsetSuccessfulAppObjectsCount
`func (o *ArchivalTargetResult) UnsetSuccessfulAppObjectsCount()`

UnsetSuccessfulAppObjectsCount ensures that no value is present for SuccessfulAppObjectsCount, not even an explicit nil
### GetFailedAppObjectsCount

`func (o *ArchivalTargetResult) GetFailedAppObjectsCount() int32`

GetFailedAppObjectsCount returns the FailedAppObjectsCount field if non-nil, zero value otherwise.

### GetFailedAppObjectsCountOk

`func (o *ArchivalTargetResult) GetFailedAppObjectsCountOk() (*int32, bool)`

GetFailedAppObjectsCountOk returns a tuple with the FailedAppObjectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAppObjectsCount

`func (o *ArchivalTargetResult) SetFailedAppObjectsCount(v int32)`

SetFailedAppObjectsCount sets FailedAppObjectsCount field to given value.

### HasFailedAppObjectsCount

`func (o *ArchivalTargetResult) HasFailedAppObjectsCount() bool`

HasFailedAppObjectsCount returns a boolean if a field has been set.

### SetFailedAppObjectsCountNil

`func (o *ArchivalTargetResult) SetFailedAppObjectsCountNil(b bool)`

 SetFailedAppObjectsCountNil sets the value for FailedAppObjectsCount to be an explicit nil

### UnsetFailedAppObjectsCount
`func (o *ArchivalTargetResult) UnsetFailedAppObjectsCount()`

UnsetFailedAppObjectsCount ensures that no value is present for FailedAppObjectsCount, not even an explicit nil
### GetCancelledAppObjectsCount

`func (o *ArchivalTargetResult) GetCancelledAppObjectsCount() int32`

GetCancelledAppObjectsCount returns the CancelledAppObjectsCount field if non-nil, zero value otherwise.

### GetCancelledAppObjectsCountOk

`func (o *ArchivalTargetResult) GetCancelledAppObjectsCountOk() (*int32, bool)`

GetCancelledAppObjectsCountOk returns a tuple with the CancelledAppObjectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledAppObjectsCount

`func (o *ArchivalTargetResult) SetCancelledAppObjectsCount(v int32)`

SetCancelledAppObjectsCount sets CancelledAppObjectsCount field to given value.

### HasCancelledAppObjectsCount

`func (o *ArchivalTargetResult) HasCancelledAppObjectsCount() bool`

HasCancelledAppObjectsCount returns a boolean if a field has been set.

### SetCancelledAppObjectsCountNil

`func (o *ArchivalTargetResult) SetCancelledAppObjectsCountNil(b bool)`

 SetCancelledAppObjectsCountNil sets the value for CancelledAppObjectsCount to be an explicit nil

### UnsetCancelledAppObjectsCount
`func (o *ArchivalTargetResult) UnsetCancelledAppObjectsCount()`

UnsetCancelledAppObjectsCount ensures that no value is present for CancelledAppObjectsCount, not even an explicit nil
### GetStats

`func (o *ArchivalTargetResult) GetStats() ArchivalDataStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ArchivalTargetResult) GetStatsOk() (*ArchivalDataStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ArchivalTargetResult) SetStats(v ArchivalDataStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ArchivalTargetResult) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetIsManuallyDeleted

`func (o *ArchivalTargetResult) GetIsManuallyDeleted() bool`

GetIsManuallyDeleted returns the IsManuallyDeleted field if non-nil, zero value otherwise.

### GetIsManuallyDeletedOk

`func (o *ArchivalTargetResult) GetIsManuallyDeletedOk() (*bool, bool)`

GetIsManuallyDeletedOk returns a tuple with the IsManuallyDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManuallyDeleted

`func (o *ArchivalTargetResult) SetIsManuallyDeleted(v bool)`

SetIsManuallyDeleted sets IsManuallyDeleted field to given value.

### HasIsManuallyDeleted

`func (o *ArchivalTargetResult) HasIsManuallyDeleted() bool`

HasIsManuallyDeleted returns a boolean if a field has been set.

### SetIsManuallyDeletedNil

`func (o *ArchivalTargetResult) SetIsManuallyDeletedNil(b bool)`

 SetIsManuallyDeletedNil sets the value for IsManuallyDeleted to be an explicit nil

### UnsetIsManuallyDeleted
`func (o *ArchivalTargetResult) UnsetIsManuallyDeleted()`

UnsetIsManuallyDeleted ensures that no value is present for IsManuallyDeleted, not even an explicit nil
### GetExpiryTimeUsecs

`func (o *ArchivalTargetResult) GetExpiryTimeUsecs() int64`

GetExpiryTimeUsecs returns the ExpiryTimeUsecs field if non-nil, zero value otherwise.

### GetExpiryTimeUsecsOk

`func (o *ArchivalTargetResult) GetExpiryTimeUsecsOk() (*int64, bool)`

GetExpiryTimeUsecsOk returns a tuple with the ExpiryTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimeUsecs

`func (o *ArchivalTargetResult) SetExpiryTimeUsecs(v int64)`

SetExpiryTimeUsecs sets ExpiryTimeUsecs field to given value.

### HasExpiryTimeUsecs

`func (o *ArchivalTargetResult) HasExpiryTimeUsecs() bool`

HasExpiryTimeUsecs returns a boolean if a field has been set.

### SetExpiryTimeUsecsNil

`func (o *ArchivalTargetResult) SetExpiryTimeUsecsNil(b bool)`

 SetExpiryTimeUsecsNil sets the value for ExpiryTimeUsecs to be an explicit nil

### UnsetExpiryTimeUsecs
`func (o *ArchivalTargetResult) UnsetExpiryTimeUsecs()`

UnsetExpiryTimeUsecs ensures that no value is present for ExpiryTimeUsecs, not even an explicit nil
### GetDataLockConstraints

`func (o *ArchivalTargetResult) GetDataLockConstraints() DataLockConstraints`

GetDataLockConstraints returns the DataLockConstraints field if non-nil, zero value otherwise.

### GetDataLockConstraintsOk

`func (o *ArchivalTargetResult) GetDataLockConstraintsOk() (*DataLockConstraints, bool)`

GetDataLockConstraintsOk returns a tuple with the DataLockConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLockConstraints

`func (o *ArchivalTargetResult) SetDataLockConstraints(v DataLockConstraints)`

SetDataLockConstraints sets DataLockConstraints field to given value.

### HasDataLockConstraints

`func (o *ArchivalTargetResult) HasDataLockConstraints() bool`

HasDataLockConstraints returns a boolean if a field has been set.

### GetOnLegalHold

`func (o *ArchivalTargetResult) GetOnLegalHold() bool`

GetOnLegalHold returns the OnLegalHold field if non-nil, zero value otherwise.

### GetOnLegalHoldOk

`func (o *ArchivalTargetResult) GetOnLegalHoldOk() (*bool, bool)`

GetOnLegalHoldOk returns a tuple with the OnLegalHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnLegalHold

`func (o *ArchivalTargetResult) SetOnLegalHold(v bool)`

SetOnLegalHold sets OnLegalHold field to given value.

### HasOnLegalHold

`func (o *ArchivalTargetResult) HasOnLegalHold() bool`

HasOnLegalHold returns a boolean if a field has been set.

### SetOnLegalHoldNil

`func (o *ArchivalTargetResult) SetOnLegalHoldNil(b bool)`

 SetOnLegalHoldNil sets the value for OnLegalHold to be an explicit nil

### UnsetOnLegalHold
`func (o *ArchivalTargetResult) UnsetOnLegalHold()`

UnsetOnLegalHold ensures that no value is present for OnLegalHold, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


