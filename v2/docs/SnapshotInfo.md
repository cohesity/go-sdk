# SnapshotInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdmittedTimeUsecs** | Pointer to **NullableInt64** | Specifies the time at which the backup task was admitted to run in Unix epoch Timestamp(in microseconds) for an object. | [optional] 
**BackupFileCount** | Pointer to **NullableInt64** | The total number of file and directory entities that are backed up in this run. Only applicable to file based backups. | [optional] 
**DataLockConstraints** | Pointer to [**DataLockConstraints**](DataLockConstraints.md) |  | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time of attempt in Unix epoch Timestamp(in microseconds) for an object. | [optional] 
**ExpiryTimeUsecs** | Pointer to **NullableInt64** | Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds) for an object. | [optional] 
**IndexingTaskId** | Pointer to **NullableString** | Progress monitor task for the indexing of documents in an object. | [optional] 
**IsManuallyDeleted** | Pointer to **NullableBool** | Specifies whether the snapshot is deleted manually. | [optional] 
**PermitGrantTimeUsecs** | Pointer to **NullableInt64** | Specifies the time when gatekeeper permit is granted to the backup task. If the backup task is rescheduled due to errors, the field is updated to the time when permit is granted again. | [optional] 
**ProgressTaskId** | Pointer to **NullableString** | Progress monitor task for backup of the object. | [optional] 
**QueueDurationUsecs** | Pointer to **NullableInt64** | Specifies the duration between the startTime and when gatekeeper permit is granted to the backup task. If the backup task is rescheduled due to errors, the field is updated considering the time when permit is granted again. Queue duration &#x3D; PermitGrantTimeUsecs - StartTimeUsecs | [optional] 
**SnapshotCreationTimeUsecs** | Pointer to **NullableInt64** | Specifies the time at which the source snapshot was taken in Unix epoch Timestamp(in microseconds) for an object. | [optional] 
**SnapshotId** | Pointer to **NullableString** | Snapshot id for a successful snapshot. This field will not be set if the Protection Group Run has no successful attempt. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of attempt in Unix epoch Timestamp(in microseconds) for an object. | [optional] 
**Stats** | Pointer to [**BackupDataStats**](BackupDataStats.md) |  | [optional] 
**StatsTaskId** | Pointer to **NullableString** | Stats task for an object. | [optional] 
**Status** | Pointer to **NullableString** | Status of snapshot. | [optional] 
**StatusMessage** | Pointer to **NullableString** | A message decribing the status. This will be populated currently only for kWaitingForOlderBackupRun status. | [optional] 
**TotalFileCount** | Pointer to **NullableInt64** | The total number of file and directory entities visited in this backup. Only applicable to file based backups. | [optional] 
**Warnings** | Pointer to **[]string** | Specifies a list of warning messages. | [optional] 

## Methods

### NewSnapshotInfo

`func NewSnapshotInfo() *SnapshotInfo`

NewSnapshotInfo instantiates a new SnapshotInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotInfoWithDefaults

`func NewSnapshotInfoWithDefaults() *SnapshotInfo`

NewSnapshotInfoWithDefaults instantiates a new SnapshotInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmittedTimeUsecs

`func (o *SnapshotInfo) GetAdmittedTimeUsecs() int64`

GetAdmittedTimeUsecs returns the AdmittedTimeUsecs field if non-nil, zero value otherwise.

### GetAdmittedTimeUsecsOk

`func (o *SnapshotInfo) GetAdmittedTimeUsecsOk() (*int64, bool)`

GetAdmittedTimeUsecsOk returns a tuple with the AdmittedTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmittedTimeUsecs

`func (o *SnapshotInfo) SetAdmittedTimeUsecs(v int64)`

SetAdmittedTimeUsecs sets AdmittedTimeUsecs field to given value.

### HasAdmittedTimeUsecs

`func (o *SnapshotInfo) HasAdmittedTimeUsecs() bool`

HasAdmittedTimeUsecs returns a boolean if a field has been set.

### SetAdmittedTimeUsecsNil

`func (o *SnapshotInfo) SetAdmittedTimeUsecsNil(b bool)`

 SetAdmittedTimeUsecsNil sets the value for AdmittedTimeUsecs to be an explicit nil

### UnsetAdmittedTimeUsecs
`func (o *SnapshotInfo) UnsetAdmittedTimeUsecs()`

UnsetAdmittedTimeUsecs ensures that no value is present for AdmittedTimeUsecs, not even an explicit nil
### GetBackupFileCount

`func (o *SnapshotInfo) GetBackupFileCount() int64`

GetBackupFileCount returns the BackupFileCount field if non-nil, zero value otherwise.

### GetBackupFileCountOk

`func (o *SnapshotInfo) GetBackupFileCountOk() (*int64, bool)`

GetBackupFileCountOk returns a tuple with the BackupFileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupFileCount

`func (o *SnapshotInfo) SetBackupFileCount(v int64)`

SetBackupFileCount sets BackupFileCount field to given value.

### HasBackupFileCount

`func (o *SnapshotInfo) HasBackupFileCount() bool`

HasBackupFileCount returns a boolean if a field has been set.

### SetBackupFileCountNil

`func (o *SnapshotInfo) SetBackupFileCountNil(b bool)`

 SetBackupFileCountNil sets the value for BackupFileCount to be an explicit nil

### UnsetBackupFileCount
`func (o *SnapshotInfo) UnsetBackupFileCount()`

UnsetBackupFileCount ensures that no value is present for BackupFileCount, not even an explicit nil
### GetDataLockConstraints

`func (o *SnapshotInfo) GetDataLockConstraints() DataLockConstraints`

GetDataLockConstraints returns the DataLockConstraints field if non-nil, zero value otherwise.

### GetDataLockConstraintsOk

`func (o *SnapshotInfo) GetDataLockConstraintsOk() (*DataLockConstraints, bool)`

GetDataLockConstraintsOk returns a tuple with the DataLockConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLockConstraints

`func (o *SnapshotInfo) SetDataLockConstraints(v DataLockConstraints)`

SetDataLockConstraints sets DataLockConstraints field to given value.

### HasDataLockConstraints

`func (o *SnapshotInfo) HasDataLockConstraints() bool`

HasDataLockConstraints returns a boolean if a field has been set.

### GetEndTimeUsecs

`func (o *SnapshotInfo) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *SnapshotInfo) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *SnapshotInfo) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *SnapshotInfo) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *SnapshotInfo) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *SnapshotInfo) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetExpiryTimeUsecs

`func (o *SnapshotInfo) GetExpiryTimeUsecs() int64`

GetExpiryTimeUsecs returns the ExpiryTimeUsecs field if non-nil, zero value otherwise.

### GetExpiryTimeUsecsOk

`func (o *SnapshotInfo) GetExpiryTimeUsecsOk() (*int64, bool)`

GetExpiryTimeUsecsOk returns a tuple with the ExpiryTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimeUsecs

`func (o *SnapshotInfo) SetExpiryTimeUsecs(v int64)`

SetExpiryTimeUsecs sets ExpiryTimeUsecs field to given value.

### HasExpiryTimeUsecs

`func (o *SnapshotInfo) HasExpiryTimeUsecs() bool`

HasExpiryTimeUsecs returns a boolean if a field has been set.

### SetExpiryTimeUsecsNil

`func (o *SnapshotInfo) SetExpiryTimeUsecsNil(b bool)`

 SetExpiryTimeUsecsNil sets the value for ExpiryTimeUsecs to be an explicit nil

### UnsetExpiryTimeUsecs
`func (o *SnapshotInfo) UnsetExpiryTimeUsecs()`

UnsetExpiryTimeUsecs ensures that no value is present for ExpiryTimeUsecs, not even an explicit nil
### GetIndexingTaskId

`func (o *SnapshotInfo) GetIndexingTaskId() string`

GetIndexingTaskId returns the IndexingTaskId field if non-nil, zero value otherwise.

### GetIndexingTaskIdOk

`func (o *SnapshotInfo) GetIndexingTaskIdOk() (*string, bool)`

GetIndexingTaskIdOk returns a tuple with the IndexingTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingTaskId

`func (o *SnapshotInfo) SetIndexingTaskId(v string)`

SetIndexingTaskId sets IndexingTaskId field to given value.

### HasIndexingTaskId

`func (o *SnapshotInfo) HasIndexingTaskId() bool`

HasIndexingTaskId returns a boolean if a field has been set.

### SetIndexingTaskIdNil

`func (o *SnapshotInfo) SetIndexingTaskIdNil(b bool)`

 SetIndexingTaskIdNil sets the value for IndexingTaskId to be an explicit nil

### UnsetIndexingTaskId
`func (o *SnapshotInfo) UnsetIndexingTaskId()`

UnsetIndexingTaskId ensures that no value is present for IndexingTaskId, not even an explicit nil
### GetIsManuallyDeleted

`func (o *SnapshotInfo) GetIsManuallyDeleted() bool`

GetIsManuallyDeleted returns the IsManuallyDeleted field if non-nil, zero value otherwise.

### GetIsManuallyDeletedOk

`func (o *SnapshotInfo) GetIsManuallyDeletedOk() (*bool, bool)`

GetIsManuallyDeletedOk returns a tuple with the IsManuallyDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManuallyDeleted

`func (o *SnapshotInfo) SetIsManuallyDeleted(v bool)`

SetIsManuallyDeleted sets IsManuallyDeleted field to given value.

### HasIsManuallyDeleted

`func (o *SnapshotInfo) HasIsManuallyDeleted() bool`

HasIsManuallyDeleted returns a boolean if a field has been set.

### SetIsManuallyDeletedNil

`func (o *SnapshotInfo) SetIsManuallyDeletedNil(b bool)`

 SetIsManuallyDeletedNil sets the value for IsManuallyDeleted to be an explicit nil

### UnsetIsManuallyDeleted
`func (o *SnapshotInfo) UnsetIsManuallyDeleted()`

UnsetIsManuallyDeleted ensures that no value is present for IsManuallyDeleted, not even an explicit nil
### GetPermitGrantTimeUsecs

`func (o *SnapshotInfo) GetPermitGrantTimeUsecs() int64`

GetPermitGrantTimeUsecs returns the PermitGrantTimeUsecs field if non-nil, zero value otherwise.

### GetPermitGrantTimeUsecsOk

`func (o *SnapshotInfo) GetPermitGrantTimeUsecsOk() (*int64, bool)`

GetPermitGrantTimeUsecsOk returns a tuple with the PermitGrantTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermitGrantTimeUsecs

`func (o *SnapshotInfo) SetPermitGrantTimeUsecs(v int64)`

SetPermitGrantTimeUsecs sets PermitGrantTimeUsecs field to given value.

### HasPermitGrantTimeUsecs

`func (o *SnapshotInfo) HasPermitGrantTimeUsecs() bool`

HasPermitGrantTimeUsecs returns a boolean if a field has been set.

### SetPermitGrantTimeUsecsNil

`func (o *SnapshotInfo) SetPermitGrantTimeUsecsNil(b bool)`

 SetPermitGrantTimeUsecsNil sets the value for PermitGrantTimeUsecs to be an explicit nil

### UnsetPermitGrantTimeUsecs
`func (o *SnapshotInfo) UnsetPermitGrantTimeUsecs()`

UnsetPermitGrantTimeUsecs ensures that no value is present for PermitGrantTimeUsecs, not even an explicit nil
### GetProgressTaskId

`func (o *SnapshotInfo) GetProgressTaskId() string`

GetProgressTaskId returns the ProgressTaskId field if non-nil, zero value otherwise.

### GetProgressTaskIdOk

`func (o *SnapshotInfo) GetProgressTaskIdOk() (*string, bool)`

GetProgressTaskIdOk returns a tuple with the ProgressTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressTaskId

`func (o *SnapshotInfo) SetProgressTaskId(v string)`

SetProgressTaskId sets ProgressTaskId field to given value.

### HasProgressTaskId

`func (o *SnapshotInfo) HasProgressTaskId() bool`

HasProgressTaskId returns a boolean if a field has been set.

### SetProgressTaskIdNil

`func (o *SnapshotInfo) SetProgressTaskIdNil(b bool)`

 SetProgressTaskIdNil sets the value for ProgressTaskId to be an explicit nil

### UnsetProgressTaskId
`func (o *SnapshotInfo) UnsetProgressTaskId()`

UnsetProgressTaskId ensures that no value is present for ProgressTaskId, not even an explicit nil
### GetQueueDurationUsecs

`func (o *SnapshotInfo) GetQueueDurationUsecs() int64`

GetQueueDurationUsecs returns the QueueDurationUsecs field if non-nil, zero value otherwise.

### GetQueueDurationUsecsOk

`func (o *SnapshotInfo) GetQueueDurationUsecsOk() (*int64, bool)`

GetQueueDurationUsecsOk returns a tuple with the QueueDurationUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueDurationUsecs

`func (o *SnapshotInfo) SetQueueDurationUsecs(v int64)`

SetQueueDurationUsecs sets QueueDurationUsecs field to given value.

### HasQueueDurationUsecs

`func (o *SnapshotInfo) HasQueueDurationUsecs() bool`

HasQueueDurationUsecs returns a boolean if a field has been set.

### SetQueueDurationUsecsNil

`func (o *SnapshotInfo) SetQueueDurationUsecsNil(b bool)`

 SetQueueDurationUsecsNil sets the value for QueueDurationUsecs to be an explicit nil

### UnsetQueueDurationUsecs
`func (o *SnapshotInfo) UnsetQueueDurationUsecs()`

UnsetQueueDurationUsecs ensures that no value is present for QueueDurationUsecs, not even an explicit nil
### GetSnapshotCreationTimeUsecs

`func (o *SnapshotInfo) GetSnapshotCreationTimeUsecs() int64`

GetSnapshotCreationTimeUsecs returns the SnapshotCreationTimeUsecs field if non-nil, zero value otherwise.

### GetSnapshotCreationTimeUsecsOk

`func (o *SnapshotInfo) GetSnapshotCreationTimeUsecsOk() (*int64, bool)`

GetSnapshotCreationTimeUsecsOk returns a tuple with the SnapshotCreationTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotCreationTimeUsecs

`func (o *SnapshotInfo) SetSnapshotCreationTimeUsecs(v int64)`

SetSnapshotCreationTimeUsecs sets SnapshotCreationTimeUsecs field to given value.

### HasSnapshotCreationTimeUsecs

`func (o *SnapshotInfo) HasSnapshotCreationTimeUsecs() bool`

HasSnapshotCreationTimeUsecs returns a boolean if a field has been set.

### SetSnapshotCreationTimeUsecsNil

`func (o *SnapshotInfo) SetSnapshotCreationTimeUsecsNil(b bool)`

 SetSnapshotCreationTimeUsecsNil sets the value for SnapshotCreationTimeUsecs to be an explicit nil

### UnsetSnapshotCreationTimeUsecs
`func (o *SnapshotInfo) UnsetSnapshotCreationTimeUsecs()`

UnsetSnapshotCreationTimeUsecs ensures that no value is present for SnapshotCreationTimeUsecs, not even an explicit nil
### GetSnapshotId

`func (o *SnapshotInfo) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *SnapshotInfo) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *SnapshotInfo) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *SnapshotInfo) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### SetSnapshotIdNil

`func (o *SnapshotInfo) SetSnapshotIdNil(b bool)`

 SetSnapshotIdNil sets the value for SnapshotId to be an explicit nil

### UnsetSnapshotId
`func (o *SnapshotInfo) UnsetSnapshotId()`

UnsetSnapshotId ensures that no value is present for SnapshotId, not even an explicit nil
### GetStartTimeUsecs

`func (o *SnapshotInfo) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *SnapshotInfo) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *SnapshotInfo) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *SnapshotInfo) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *SnapshotInfo) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *SnapshotInfo) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetStats

`func (o *SnapshotInfo) GetStats() BackupDataStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *SnapshotInfo) GetStatsOk() (*BackupDataStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *SnapshotInfo) SetStats(v BackupDataStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *SnapshotInfo) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetStatsTaskId

`func (o *SnapshotInfo) GetStatsTaskId() string`

GetStatsTaskId returns the StatsTaskId field if non-nil, zero value otherwise.

### GetStatsTaskIdOk

`func (o *SnapshotInfo) GetStatsTaskIdOk() (*string, bool)`

GetStatsTaskIdOk returns a tuple with the StatsTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsTaskId

`func (o *SnapshotInfo) SetStatsTaskId(v string)`

SetStatsTaskId sets StatsTaskId field to given value.

### HasStatsTaskId

`func (o *SnapshotInfo) HasStatsTaskId() bool`

HasStatsTaskId returns a boolean if a field has been set.

### SetStatsTaskIdNil

`func (o *SnapshotInfo) SetStatsTaskIdNil(b bool)`

 SetStatsTaskIdNil sets the value for StatsTaskId to be an explicit nil

### UnsetStatsTaskId
`func (o *SnapshotInfo) UnsetStatsTaskId()`

UnsetStatsTaskId ensures that no value is present for StatsTaskId, not even an explicit nil
### GetStatus

`func (o *SnapshotInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SnapshotInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SnapshotInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SnapshotInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *SnapshotInfo) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *SnapshotInfo) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStatusMessage

`func (o *SnapshotInfo) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *SnapshotInfo) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *SnapshotInfo) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *SnapshotInfo) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *SnapshotInfo) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *SnapshotInfo) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetTotalFileCount

`func (o *SnapshotInfo) GetTotalFileCount() int64`

GetTotalFileCount returns the TotalFileCount field if non-nil, zero value otherwise.

### GetTotalFileCountOk

`func (o *SnapshotInfo) GetTotalFileCountOk() (*int64, bool)`

GetTotalFileCountOk returns a tuple with the TotalFileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFileCount

`func (o *SnapshotInfo) SetTotalFileCount(v int64)`

SetTotalFileCount sets TotalFileCount field to given value.

### HasTotalFileCount

`func (o *SnapshotInfo) HasTotalFileCount() bool`

HasTotalFileCount returns a boolean if a field has been set.

### SetTotalFileCountNil

`func (o *SnapshotInfo) SetTotalFileCountNil(b bool)`

 SetTotalFileCountNil sets the value for TotalFileCount to be an explicit nil

### UnsetTotalFileCount
`func (o *SnapshotInfo) UnsetTotalFileCount()`

UnsetTotalFileCount ensures that no value is present for TotalFileCount, not even an explicit nil
### GetWarnings

`func (o *SnapshotInfo) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *SnapshotInfo) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *SnapshotInfo) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *SnapshotInfo) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarningsNil

`func (o *SnapshotInfo) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *SnapshotInfo) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


