# LatestProtectionRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupRun** | Pointer to [**SourceBackupStatus**](SourceBackupStatus.md) |  | [optional] 
**ChangeEventId** | Pointer to **NullableInt64** | Specifies the event id which caused last update on this object. | [optional] 
**CopyRun** | Pointer to [**CopyRun**](CopyRun.md) |  | [optional] 
**JobRunId** | Pointer to **NullableInt64** | Specifies job run id of the latest successful Protection Job Run. | [optional] 
**ProtectionJobRunUid** | Pointer to [**RunUid**](RunUid.md) |  | [optional] 
**SnapshotTarget** | Pointer to **NullableString** | Specifies the cluster id in case of local or replication snapshots and name of location in case of archival snapshots. | [optional] 
**SnapshotTargetType** | Pointer to **NullableInt32** | Specifies the snapshot target type of the latest snapshot. | [optional] 
**TaskStatus** | Pointer to **NullableInt32** | Specifies the task status of the Protection Job Run in the final attempt. | [optional] 
**Uuid** | Pointer to **NullableString** | Specifies the unique id of the Protection Source for which a snapshot is taken. | [optional] 

## Methods

### NewLatestProtectionRun

`func NewLatestProtectionRun() *LatestProtectionRun`

NewLatestProtectionRun instantiates a new LatestProtectionRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLatestProtectionRunWithDefaults

`func NewLatestProtectionRunWithDefaults() *LatestProtectionRun`

NewLatestProtectionRunWithDefaults instantiates a new LatestProtectionRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupRun

`func (o *LatestProtectionRun) GetBackupRun() SourceBackupStatus`

GetBackupRun returns the BackupRun field if non-nil, zero value otherwise.

### GetBackupRunOk

`func (o *LatestProtectionRun) GetBackupRunOk() (*SourceBackupStatus, bool)`

GetBackupRunOk returns a tuple with the BackupRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRun

`func (o *LatestProtectionRun) SetBackupRun(v SourceBackupStatus)`

SetBackupRun sets BackupRun field to given value.

### HasBackupRun

`func (o *LatestProtectionRun) HasBackupRun() bool`

HasBackupRun returns a boolean if a field has been set.

### GetChangeEventId

`func (o *LatestProtectionRun) GetChangeEventId() int64`

GetChangeEventId returns the ChangeEventId field if non-nil, zero value otherwise.

### GetChangeEventIdOk

`func (o *LatestProtectionRun) GetChangeEventIdOk() (*int64, bool)`

GetChangeEventIdOk returns a tuple with the ChangeEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeEventId

`func (o *LatestProtectionRun) SetChangeEventId(v int64)`

SetChangeEventId sets ChangeEventId field to given value.

### HasChangeEventId

`func (o *LatestProtectionRun) HasChangeEventId() bool`

HasChangeEventId returns a boolean if a field has been set.

### SetChangeEventIdNil

`func (o *LatestProtectionRun) SetChangeEventIdNil(b bool)`

 SetChangeEventIdNil sets the value for ChangeEventId to be an explicit nil

### UnsetChangeEventId
`func (o *LatestProtectionRun) UnsetChangeEventId()`

UnsetChangeEventId ensures that no value is present for ChangeEventId, not even an explicit nil
### GetCopyRun

`func (o *LatestProtectionRun) GetCopyRun() CopyRun`

GetCopyRun returns the CopyRun field if non-nil, zero value otherwise.

### GetCopyRunOk

`func (o *LatestProtectionRun) GetCopyRunOk() (*CopyRun, bool)`

GetCopyRunOk returns a tuple with the CopyRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyRun

`func (o *LatestProtectionRun) SetCopyRun(v CopyRun)`

SetCopyRun sets CopyRun field to given value.

### HasCopyRun

`func (o *LatestProtectionRun) HasCopyRun() bool`

HasCopyRun returns a boolean if a field has been set.

### GetJobRunId

`func (o *LatestProtectionRun) GetJobRunId() int64`

GetJobRunId returns the JobRunId field if non-nil, zero value otherwise.

### GetJobRunIdOk

`func (o *LatestProtectionRun) GetJobRunIdOk() (*int64, bool)`

GetJobRunIdOk returns a tuple with the JobRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobRunId

`func (o *LatestProtectionRun) SetJobRunId(v int64)`

SetJobRunId sets JobRunId field to given value.

### HasJobRunId

`func (o *LatestProtectionRun) HasJobRunId() bool`

HasJobRunId returns a boolean if a field has been set.

### SetJobRunIdNil

`func (o *LatestProtectionRun) SetJobRunIdNil(b bool)`

 SetJobRunIdNil sets the value for JobRunId to be an explicit nil

### UnsetJobRunId
`func (o *LatestProtectionRun) UnsetJobRunId()`

UnsetJobRunId ensures that no value is present for JobRunId, not even an explicit nil
### GetProtectionJobRunUid

`func (o *LatestProtectionRun) GetProtectionJobRunUid() RunUid`

GetProtectionJobRunUid returns the ProtectionJobRunUid field if non-nil, zero value otherwise.

### GetProtectionJobRunUidOk

`func (o *LatestProtectionRun) GetProtectionJobRunUidOk() (*RunUid, bool)`

GetProtectionJobRunUidOk returns a tuple with the ProtectionJobRunUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionJobRunUid

`func (o *LatestProtectionRun) SetProtectionJobRunUid(v RunUid)`

SetProtectionJobRunUid sets ProtectionJobRunUid field to given value.

### HasProtectionJobRunUid

`func (o *LatestProtectionRun) HasProtectionJobRunUid() bool`

HasProtectionJobRunUid returns a boolean if a field has been set.

### GetSnapshotTarget

`func (o *LatestProtectionRun) GetSnapshotTarget() string`

GetSnapshotTarget returns the SnapshotTarget field if non-nil, zero value otherwise.

### GetSnapshotTargetOk

`func (o *LatestProtectionRun) GetSnapshotTargetOk() (*string, bool)`

GetSnapshotTargetOk returns a tuple with the SnapshotTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTarget

`func (o *LatestProtectionRun) SetSnapshotTarget(v string)`

SetSnapshotTarget sets SnapshotTarget field to given value.

### HasSnapshotTarget

`func (o *LatestProtectionRun) HasSnapshotTarget() bool`

HasSnapshotTarget returns a boolean if a field has been set.

### SetSnapshotTargetNil

`func (o *LatestProtectionRun) SetSnapshotTargetNil(b bool)`

 SetSnapshotTargetNil sets the value for SnapshotTarget to be an explicit nil

### UnsetSnapshotTarget
`func (o *LatestProtectionRun) UnsetSnapshotTarget()`

UnsetSnapshotTarget ensures that no value is present for SnapshotTarget, not even an explicit nil
### GetSnapshotTargetType

`func (o *LatestProtectionRun) GetSnapshotTargetType() int32`

GetSnapshotTargetType returns the SnapshotTargetType field if non-nil, zero value otherwise.

### GetSnapshotTargetTypeOk

`func (o *LatestProtectionRun) GetSnapshotTargetTypeOk() (*int32, bool)`

GetSnapshotTargetTypeOk returns a tuple with the SnapshotTargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTargetType

`func (o *LatestProtectionRun) SetSnapshotTargetType(v int32)`

SetSnapshotTargetType sets SnapshotTargetType field to given value.

### HasSnapshotTargetType

`func (o *LatestProtectionRun) HasSnapshotTargetType() bool`

HasSnapshotTargetType returns a boolean if a field has been set.

### SetSnapshotTargetTypeNil

`func (o *LatestProtectionRun) SetSnapshotTargetTypeNil(b bool)`

 SetSnapshotTargetTypeNil sets the value for SnapshotTargetType to be an explicit nil

### UnsetSnapshotTargetType
`func (o *LatestProtectionRun) UnsetSnapshotTargetType()`

UnsetSnapshotTargetType ensures that no value is present for SnapshotTargetType, not even an explicit nil
### GetTaskStatus

`func (o *LatestProtectionRun) GetTaskStatus() int32`

GetTaskStatus returns the TaskStatus field if non-nil, zero value otherwise.

### GetTaskStatusOk

`func (o *LatestProtectionRun) GetTaskStatusOk() (*int32, bool)`

GetTaskStatusOk returns a tuple with the TaskStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskStatus

`func (o *LatestProtectionRun) SetTaskStatus(v int32)`

SetTaskStatus sets TaskStatus field to given value.

### HasTaskStatus

`func (o *LatestProtectionRun) HasTaskStatus() bool`

HasTaskStatus returns a boolean if a field has been set.

### SetTaskStatusNil

`func (o *LatestProtectionRun) SetTaskStatusNil(b bool)`

 SetTaskStatusNil sets the value for TaskStatus to be an explicit nil

### UnsetTaskStatus
`func (o *LatestProtectionRun) UnsetTaskStatus()`

UnsetTaskStatus ensures that no value is present for TaskStatus, not even an explicit nil
### GetUuid

`func (o *LatestProtectionRun) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *LatestProtectionRun) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *LatestProtectionRun) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *LatestProtectionRun) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *LatestProtectionRun) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *LatestProtectionRun) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


