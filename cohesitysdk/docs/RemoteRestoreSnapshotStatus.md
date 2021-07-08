# RemoteRestoreSnapshotStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveTaskUid** | Pointer to [**NullableUniversalId**](UniversalId.md) | Specifies the globally unique id of the archival task that archived the Snapshots to the remote Vault. | [optional] 
**Error** | Pointer to **NullableString** | Specifies the error message if the indexing task fails. | [optional] 
**ExpiryTimeUsecs** | Pointer to **NullableInt64** | Specifies the time when the Snapshot expires on the remote Vault. This field is recorded as a Unix epoch Timestamp (in microseconds). | [optional] 
**JobRunId** | Pointer to **NullableInt64** | Specifies the id of the Job Run that originally captured the Snapshot. | [optional] 
**ProgressMonitorTask** | Pointer to **NullableString** | Specifies the path to the progress monitor task that tracks the progress of building the index. | [optional] 
**SnapshotTaskStatus** | Pointer to **NullableString** | Specifies the status of the indexing task. &#39;kJobRunning&#39; indicates that the Job/task is currently running. &#39;kJobFinished&#39; indicates that the Job/task completed and finished. &#39;kJobFailed&#39; indicates that the Job/task failed and did not complete. &#39;kJobCanceled&#39; indicates that the Job/task was canceled. &#39;kJobPaused&#39; indicates the Job/task is paused. | [optional] 
**SnapshotTaskUid** | Pointer to [**NullableUniversalId**](UniversalId.md) | Specifies the globally unique id of the task capturing the Snapshot. | [optional] 
**SnapshotTimeUsecs** | Pointer to **NullableInt64** | Specify the time the Snapshot was captured. This time is recorded as a Unix epoch Timestamp (in microseconds). | [optional] 

## Methods

### NewRemoteRestoreSnapshotStatus

`func NewRemoteRestoreSnapshotStatus() *RemoteRestoreSnapshotStatus`

NewRemoteRestoreSnapshotStatus instantiates a new RemoteRestoreSnapshotStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteRestoreSnapshotStatusWithDefaults

`func NewRemoteRestoreSnapshotStatusWithDefaults() *RemoteRestoreSnapshotStatus`

NewRemoteRestoreSnapshotStatusWithDefaults instantiates a new RemoteRestoreSnapshotStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveTaskUid

`func (o *RemoteRestoreSnapshotStatus) GetArchiveTaskUid() UniversalId`

GetArchiveTaskUid returns the ArchiveTaskUid field if non-nil, zero value otherwise.

### GetArchiveTaskUidOk

`func (o *RemoteRestoreSnapshotStatus) GetArchiveTaskUidOk() (*UniversalId, bool)`

GetArchiveTaskUidOk returns a tuple with the ArchiveTaskUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveTaskUid

`func (o *RemoteRestoreSnapshotStatus) SetArchiveTaskUid(v UniversalId)`

SetArchiveTaskUid sets ArchiveTaskUid field to given value.

### HasArchiveTaskUid

`func (o *RemoteRestoreSnapshotStatus) HasArchiveTaskUid() bool`

HasArchiveTaskUid returns a boolean if a field has been set.

### SetArchiveTaskUidNil

`func (o *RemoteRestoreSnapshotStatus) SetArchiveTaskUidNil(b bool)`

 SetArchiveTaskUidNil sets the value for ArchiveTaskUid to be an explicit nil

### UnsetArchiveTaskUid
`func (o *RemoteRestoreSnapshotStatus) UnsetArchiveTaskUid()`

UnsetArchiveTaskUid ensures that no value is present for ArchiveTaskUid, not even an explicit nil
### GetError

`func (o *RemoteRestoreSnapshotStatus) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RemoteRestoreSnapshotStatus) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RemoteRestoreSnapshotStatus) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *RemoteRestoreSnapshotStatus) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *RemoteRestoreSnapshotStatus) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *RemoteRestoreSnapshotStatus) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetExpiryTimeUsecs

`func (o *RemoteRestoreSnapshotStatus) GetExpiryTimeUsecs() int64`

GetExpiryTimeUsecs returns the ExpiryTimeUsecs field if non-nil, zero value otherwise.

### GetExpiryTimeUsecsOk

`func (o *RemoteRestoreSnapshotStatus) GetExpiryTimeUsecsOk() (*int64, bool)`

GetExpiryTimeUsecsOk returns a tuple with the ExpiryTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimeUsecs

`func (o *RemoteRestoreSnapshotStatus) SetExpiryTimeUsecs(v int64)`

SetExpiryTimeUsecs sets ExpiryTimeUsecs field to given value.

### HasExpiryTimeUsecs

`func (o *RemoteRestoreSnapshotStatus) HasExpiryTimeUsecs() bool`

HasExpiryTimeUsecs returns a boolean if a field has been set.

### SetExpiryTimeUsecsNil

`func (o *RemoteRestoreSnapshotStatus) SetExpiryTimeUsecsNil(b bool)`

 SetExpiryTimeUsecsNil sets the value for ExpiryTimeUsecs to be an explicit nil

### UnsetExpiryTimeUsecs
`func (o *RemoteRestoreSnapshotStatus) UnsetExpiryTimeUsecs()`

UnsetExpiryTimeUsecs ensures that no value is present for ExpiryTimeUsecs, not even an explicit nil
### GetJobRunId

`func (o *RemoteRestoreSnapshotStatus) GetJobRunId() int64`

GetJobRunId returns the JobRunId field if non-nil, zero value otherwise.

### GetJobRunIdOk

`func (o *RemoteRestoreSnapshotStatus) GetJobRunIdOk() (*int64, bool)`

GetJobRunIdOk returns a tuple with the JobRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobRunId

`func (o *RemoteRestoreSnapshotStatus) SetJobRunId(v int64)`

SetJobRunId sets JobRunId field to given value.

### HasJobRunId

`func (o *RemoteRestoreSnapshotStatus) HasJobRunId() bool`

HasJobRunId returns a boolean if a field has been set.

### SetJobRunIdNil

`func (o *RemoteRestoreSnapshotStatus) SetJobRunIdNil(b bool)`

 SetJobRunIdNil sets the value for JobRunId to be an explicit nil

### UnsetJobRunId
`func (o *RemoteRestoreSnapshotStatus) UnsetJobRunId()`

UnsetJobRunId ensures that no value is present for JobRunId, not even an explicit nil
### GetProgressMonitorTask

`func (o *RemoteRestoreSnapshotStatus) GetProgressMonitorTask() string`

GetProgressMonitorTask returns the ProgressMonitorTask field if non-nil, zero value otherwise.

### GetProgressMonitorTaskOk

`func (o *RemoteRestoreSnapshotStatus) GetProgressMonitorTaskOk() (*string, bool)`

GetProgressMonitorTaskOk returns a tuple with the ProgressMonitorTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressMonitorTask

`func (o *RemoteRestoreSnapshotStatus) SetProgressMonitorTask(v string)`

SetProgressMonitorTask sets ProgressMonitorTask field to given value.

### HasProgressMonitorTask

`func (o *RemoteRestoreSnapshotStatus) HasProgressMonitorTask() bool`

HasProgressMonitorTask returns a boolean if a field has been set.

### SetProgressMonitorTaskNil

`func (o *RemoteRestoreSnapshotStatus) SetProgressMonitorTaskNil(b bool)`

 SetProgressMonitorTaskNil sets the value for ProgressMonitorTask to be an explicit nil

### UnsetProgressMonitorTask
`func (o *RemoteRestoreSnapshotStatus) UnsetProgressMonitorTask()`

UnsetProgressMonitorTask ensures that no value is present for ProgressMonitorTask, not even an explicit nil
### GetSnapshotTaskStatus

`func (o *RemoteRestoreSnapshotStatus) GetSnapshotTaskStatus() string`

GetSnapshotTaskStatus returns the SnapshotTaskStatus field if non-nil, zero value otherwise.

### GetSnapshotTaskStatusOk

`func (o *RemoteRestoreSnapshotStatus) GetSnapshotTaskStatusOk() (*string, bool)`

GetSnapshotTaskStatusOk returns a tuple with the SnapshotTaskStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTaskStatus

`func (o *RemoteRestoreSnapshotStatus) SetSnapshotTaskStatus(v string)`

SetSnapshotTaskStatus sets SnapshotTaskStatus field to given value.

### HasSnapshotTaskStatus

`func (o *RemoteRestoreSnapshotStatus) HasSnapshotTaskStatus() bool`

HasSnapshotTaskStatus returns a boolean if a field has been set.

### SetSnapshotTaskStatusNil

`func (o *RemoteRestoreSnapshotStatus) SetSnapshotTaskStatusNil(b bool)`

 SetSnapshotTaskStatusNil sets the value for SnapshotTaskStatus to be an explicit nil

### UnsetSnapshotTaskStatus
`func (o *RemoteRestoreSnapshotStatus) UnsetSnapshotTaskStatus()`

UnsetSnapshotTaskStatus ensures that no value is present for SnapshotTaskStatus, not even an explicit nil
### GetSnapshotTaskUid

`func (o *RemoteRestoreSnapshotStatus) GetSnapshotTaskUid() UniversalId`

GetSnapshotTaskUid returns the SnapshotTaskUid field if non-nil, zero value otherwise.

### GetSnapshotTaskUidOk

`func (o *RemoteRestoreSnapshotStatus) GetSnapshotTaskUidOk() (*UniversalId, bool)`

GetSnapshotTaskUidOk returns a tuple with the SnapshotTaskUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTaskUid

`func (o *RemoteRestoreSnapshotStatus) SetSnapshotTaskUid(v UniversalId)`

SetSnapshotTaskUid sets SnapshotTaskUid field to given value.

### HasSnapshotTaskUid

`func (o *RemoteRestoreSnapshotStatus) HasSnapshotTaskUid() bool`

HasSnapshotTaskUid returns a boolean if a field has been set.

### SetSnapshotTaskUidNil

`func (o *RemoteRestoreSnapshotStatus) SetSnapshotTaskUidNil(b bool)`

 SetSnapshotTaskUidNil sets the value for SnapshotTaskUid to be an explicit nil

### UnsetSnapshotTaskUid
`func (o *RemoteRestoreSnapshotStatus) UnsetSnapshotTaskUid()`

UnsetSnapshotTaskUid ensures that no value is present for SnapshotTaskUid, not even an explicit nil
### GetSnapshotTimeUsecs

`func (o *RemoteRestoreSnapshotStatus) GetSnapshotTimeUsecs() int64`

GetSnapshotTimeUsecs returns the SnapshotTimeUsecs field if non-nil, zero value otherwise.

### GetSnapshotTimeUsecsOk

`func (o *RemoteRestoreSnapshotStatus) GetSnapshotTimeUsecsOk() (*int64, bool)`

GetSnapshotTimeUsecsOk returns a tuple with the SnapshotTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTimeUsecs

`func (o *RemoteRestoreSnapshotStatus) SetSnapshotTimeUsecs(v int64)`

SetSnapshotTimeUsecs sets SnapshotTimeUsecs field to given value.

### HasSnapshotTimeUsecs

`func (o *RemoteRestoreSnapshotStatus) HasSnapshotTimeUsecs() bool`

HasSnapshotTimeUsecs returns a boolean if a field has been set.

### SetSnapshotTimeUsecsNil

`func (o *RemoteRestoreSnapshotStatus) SetSnapshotTimeUsecsNil(b bool)`

 SetSnapshotTimeUsecsNil sets the value for SnapshotTimeUsecs to be an explicit nil

### UnsetSnapshotTimeUsecs
`func (o *RemoteRestoreSnapshotStatus) UnsetSnapshotTimeUsecs()`

UnsetSnapshotTimeUsecs ensures that no value is present for SnapshotTimeUsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


