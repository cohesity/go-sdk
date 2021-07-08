# ProtectionSourceSnapshotInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CopyTasks** | Pointer to [**[]SnapshotCopyTask**](SnapshotCopyTask.md) | Array of Snapshot Copy Tasks.  Specifies a list of copy tasks (such as replication and archival tasks). | [optional] 
**JobId** | Pointer to **NullableInt64** | Specifies the id of the Protection Job. | [optional] 
**JobName** | Pointer to **NullableString** | Specifies the name of the Protection Job. | [optional] 
**JobRunId** | Pointer to **NullableInt64** | Specifies the id of the Job Run. | [optional] 
**JobRunStartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of the Job which this object is part of. The time is specified in Unix epoch Timestamp (in microseconds). | [optional] 
**LastRunEndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time of the last Run of this object&#39;s snapshot. The time is specified in Unix epoch Timestamp (in microseconds). | [optional] 
**LastRunStartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of the last Run of this object&#39;s snapshot. The time is specified in Unix epoch Timestamp (in microseconds). | [optional] 
**Message** | Pointer to **NullableString** | Specifies warning or error information when the Job Run is not successful. | [optional] 
**NumBytesRead** | Pointer to **NullableInt64** | Specifies the total number of bytes read. | [optional] 
**NumLogicalBytesProtected** | Pointer to **NullableInt64** | Specifies the total number of logical bytes that are protected. The logical size is when the data is fully hydrated or expanded. | [optional] 
**PaginationCookie** | Pointer to **NullableInt32** | Specifies an opaque string to pass into the next request to get the next set of Snapshots for pagination purposes. If null, this is the last set of Snapshots or the number of Snapshots returned is equal to or less than the specified pageCount. | [optional] 
**RunStatus** | Pointer to **NullableString** | Specifies the type of the Job Run. &#39;kSuccess&#39; indicates that the Job Run was successful. &#39;kRunning&#39; indicates that the Job Run is currently running. &#39;kWarning&#39; indicates that the Job Run was successful but warnings were issued. &#39;kCancelled&#39; indicates that the Job Run was canceled. &#39;kError&#39; indicates the Job Run encountered an error and did not run to completion. | [optional] 
**RunType** | Pointer to **NullableString** | Specifies the status of the Job Run. &#39;kRegular&#39; indicates an incremental (CBT) backup. Incremental backups utilizing CBT (if supported) are captured of the target protection objects. The first run of a kRegular schedule captures all the blocks. &#39;kFull&#39; indicates a full (no CBT) backup. A complete backup (all blocks) of the target protection objects are always captured and Change Block Tracking (CBT) is not utilized. &#39;kLog&#39; indicates a Database Log backup. Capture the database transaction logs to allow rolling back to a specific point in time. &#39;kSystem&#39; indicates system volume backup. It produces an image for bare metal recovery. | [optional] 

## Methods

### NewProtectionSourceSnapshotInformation

`func NewProtectionSourceSnapshotInformation() *ProtectionSourceSnapshotInformation`

NewProtectionSourceSnapshotInformation instantiates a new ProtectionSourceSnapshotInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionSourceSnapshotInformationWithDefaults

`func NewProtectionSourceSnapshotInformationWithDefaults() *ProtectionSourceSnapshotInformation`

NewProtectionSourceSnapshotInformationWithDefaults instantiates a new ProtectionSourceSnapshotInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCopyTasks

`func (o *ProtectionSourceSnapshotInformation) GetCopyTasks() []SnapshotCopyTask`

GetCopyTasks returns the CopyTasks field if non-nil, zero value otherwise.

### GetCopyTasksOk

`func (o *ProtectionSourceSnapshotInformation) GetCopyTasksOk() (*[]SnapshotCopyTask, bool)`

GetCopyTasksOk returns a tuple with the CopyTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyTasks

`func (o *ProtectionSourceSnapshotInformation) SetCopyTasks(v []SnapshotCopyTask)`

SetCopyTasks sets CopyTasks field to given value.

### HasCopyTasks

`func (o *ProtectionSourceSnapshotInformation) HasCopyTasks() bool`

HasCopyTasks returns a boolean if a field has been set.

### SetCopyTasksNil

`func (o *ProtectionSourceSnapshotInformation) SetCopyTasksNil(b bool)`

 SetCopyTasksNil sets the value for CopyTasks to be an explicit nil

### UnsetCopyTasks
`func (o *ProtectionSourceSnapshotInformation) UnsetCopyTasks()`

UnsetCopyTasks ensures that no value is present for CopyTasks, not even an explicit nil
### GetJobId

`func (o *ProtectionSourceSnapshotInformation) GetJobId() int64`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *ProtectionSourceSnapshotInformation) GetJobIdOk() (*int64, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *ProtectionSourceSnapshotInformation) SetJobId(v int64)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *ProtectionSourceSnapshotInformation) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### SetJobIdNil

`func (o *ProtectionSourceSnapshotInformation) SetJobIdNil(b bool)`

 SetJobIdNil sets the value for JobId to be an explicit nil

### UnsetJobId
`func (o *ProtectionSourceSnapshotInformation) UnsetJobId()`

UnsetJobId ensures that no value is present for JobId, not even an explicit nil
### GetJobName

`func (o *ProtectionSourceSnapshotInformation) GetJobName() string`

GetJobName returns the JobName field if non-nil, zero value otherwise.

### GetJobNameOk

`func (o *ProtectionSourceSnapshotInformation) GetJobNameOk() (*string, bool)`

GetJobNameOk returns a tuple with the JobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobName

`func (o *ProtectionSourceSnapshotInformation) SetJobName(v string)`

SetJobName sets JobName field to given value.

### HasJobName

`func (o *ProtectionSourceSnapshotInformation) HasJobName() bool`

HasJobName returns a boolean if a field has been set.

### SetJobNameNil

`func (o *ProtectionSourceSnapshotInformation) SetJobNameNil(b bool)`

 SetJobNameNil sets the value for JobName to be an explicit nil

### UnsetJobName
`func (o *ProtectionSourceSnapshotInformation) UnsetJobName()`

UnsetJobName ensures that no value is present for JobName, not even an explicit nil
### GetJobRunId

`func (o *ProtectionSourceSnapshotInformation) GetJobRunId() int64`

GetJobRunId returns the JobRunId field if non-nil, zero value otherwise.

### GetJobRunIdOk

`func (o *ProtectionSourceSnapshotInformation) GetJobRunIdOk() (*int64, bool)`

GetJobRunIdOk returns a tuple with the JobRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobRunId

`func (o *ProtectionSourceSnapshotInformation) SetJobRunId(v int64)`

SetJobRunId sets JobRunId field to given value.

### HasJobRunId

`func (o *ProtectionSourceSnapshotInformation) HasJobRunId() bool`

HasJobRunId returns a boolean if a field has been set.

### SetJobRunIdNil

`func (o *ProtectionSourceSnapshotInformation) SetJobRunIdNil(b bool)`

 SetJobRunIdNil sets the value for JobRunId to be an explicit nil

### UnsetJobRunId
`func (o *ProtectionSourceSnapshotInformation) UnsetJobRunId()`

UnsetJobRunId ensures that no value is present for JobRunId, not even an explicit nil
### GetJobRunStartTimeUsecs

`func (o *ProtectionSourceSnapshotInformation) GetJobRunStartTimeUsecs() int64`

GetJobRunStartTimeUsecs returns the JobRunStartTimeUsecs field if non-nil, zero value otherwise.

### GetJobRunStartTimeUsecsOk

`func (o *ProtectionSourceSnapshotInformation) GetJobRunStartTimeUsecsOk() (*int64, bool)`

GetJobRunStartTimeUsecsOk returns a tuple with the JobRunStartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobRunStartTimeUsecs

`func (o *ProtectionSourceSnapshotInformation) SetJobRunStartTimeUsecs(v int64)`

SetJobRunStartTimeUsecs sets JobRunStartTimeUsecs field to given value.

### HasJobRunStartTimeUsecs

`func (o *ProtectionSourceSnapshotInformation) HasJobRunStartTimeUsecs() bool`

HasJobRunStartTimeUsecs returns a boolean if a field has been set.

### SetJobRunStartTimeUsecsNil

`func (o *ProtectionSourceSnapshotInformation) SetJobRunStartTimeUsecsNil(b bool)`

 SetJobRunStartTimeUsecsNil sets the value for JobRunStartTimeUsecs to be an explicit nil

### UnsetJobRunStartTimeUsecs
`func (o *ProtectionSourceSnapshotInformation) UnsetJobRunStartTimeUsecs()`

UnsetJobRunStartTimeUsecs ensures that no value is present for JobRunStartTimeUsecs, not even an explicit nil
### GetLastRunEndTimeUsecs

`func (o *ProtectionSourceSnapshotInformation) GetLastRunEndTimeUsecs() int64`

GetLastRunEndTimeUsecs returns the LastRunEndTimeUsecs field if non-nil, zero value otherwise.

### GetLastRunEndTimeUsecsOk

`func (o *ProtectionSourceSnapshotInformation) GetLastRunEndTimeUsecsOk() (*int64, bool)`

GetLastRunEndTimeUsecsOk returns a tuple with the LastRunEndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunEndTimeUsecs

`func (o *ProtectionSourceSnapshotInformation) SetLastRunEndTimeUsecs(v int64)`

SetLastRunEndTimeUsecs sets LastRunEndTimeUsecs field to given value.

### HasLastRunEndTimeUsecs

`func (o *ProtectionSourceSnapshotInformation) HasLastRunEndTimeUsecs() bool`

HasLastRunEndTimeUsecs returns a boolean if a field has been set.

### SetLastRunEndTimeUsecsNil

`func (o *ProtectionSourceSnapshotInformation) SetLastRunEndTimeUsecsNil(b bool)`

 SetLastRunEndTimeUsecsNil sets the value for LastRunEndTimeUsecs to be an explicit nil

### UnsetLastRunEndTimeUsecs
`func (o *ProtectionSourceSnapshotInformation) UnsetLastRunEndTimeUsecs()`

UnsetLastRunEndTimeUsecs ensures that no value is present for LastRunEndTimeUsecs, not even an explicit nil
### GetLastRunStartTimeUsecs

`func (o *ProtectionSourceSnapshotInformation) GetLastRunStartTimeUsecs() int64`

GetLastRunStartTimeUsecs returns the LastRunStartTimeUsecs field if non-nil, zero value otherwise.

### GetLastRunStartTimeUsecsOk

`func (o *ProtectionSourceSnapshotInformation) GetLastRunStartTimeUsecsOk() (*int64, bool)`

GetLastRunStartTimeUsecsOk returns a tuple with the LastRunStartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunStartTimeUsecs

`func (o *ProtectionSourceSnapshotInformation) SetLastRunStartTimeUsecs(v int64)`

SetLastRunStartTimeUsecs sets LastRunStartTimeUsecs field to given value.

### HasLastRunStartTimeUsecs

`func (o *ProtectionSourceSnapshotInformation) HasLastRunStartTimeUsecs() bool`

HasLastRunStartTimeUsecs returns a boolean if a field has been set.

### SetLastRunStartTimeUsecsNil

`func (o *ProtectionSourceSnapshotInformation) SetLastRunStartTimeUsecsNil(b bool)`

 SetLastRunStartTimeUsecsNil sets the value for LastRunStartTimeUsecs to be an explicit nil

### UnsetLastRunStartTimeUsecs
`func (o *ProtectionSourceSnapshotInformation) UnsetLastRunStartTimeUsecs()`

UnsetLastRunStartTimeUsecs ensures that no value is present for LastRunStartTimeUsecs, not even an explicit nil
### GetMessage

`func (o *ProtectionSourceSnapshotInformation) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ProtectionSourceSnapshotInformation) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ProtectionSourceSnapshotInformation) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ProtectionSourceSnapshotInformation) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ProtectionSourceSnapshotInformation) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ProtectionSourceSnapshotInformation) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetNumBytesRead

`func (o *ProtectionSourceSnapshotInformation) GetNumBytesRead() int64`

GetNumBytesRead returns the NumBytesRead field if non-nil, zero value otherwise.

### GetNumBytesReadOk

`func (o *ProtectionSourceSnapshotInformation) GetNumBytesReadOk() (*int64, bool)`

GetNumBytesReadOk returns a tuple with the NumBytesRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumBytesRead

`func (o *ProtectionSourceSnapshotInformation) SetNumBytesRead(v int64)`

SetNumBytesRead sets NumBytesRead field to given value.

### HasNumBytesRead

`func (o *ProtectionSourceSnapshotInformation) HasNumBytesRead() bool`

HasNumBytesRead returns a boolean if a field has been set.

### SetNumBytesReadNil

`func (o *ProtectionSourceSnapshotInformation) SetNumBytesReadNil(b bool)`

 SetNumBytesReadNil sets the value for NumBytesRead to be an explicit nil

### UnsetNumBytesRead
`func (o *ProtectionSourceSnapshotInformation) UnsetNumBytesRead()`

UnsetNumBytesRead ensures that no value is present for NumBytesRead, not even an explicit nil
### GetNumLogicalBytesProtected

`func (o *ProtectionSourceSnapshotInformation) GetNumLogicalBytesProtected() int64`

GetNumLogicalBytesProtected returns the NumLogicalBytesProtected field if non-nil, zero value otherwise.

### GetNumLogicalBytesProtectedOk

`func (o *ProtectionSourceSnapshotInformation) GetNumLogicalBytesProtectedOk() (*int64, bool)`

GetNumLogicalBytesProtectedOk returns a tuple with the NumLogicalBytesProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumLogicalBytesProtected

`func (o *ProtectionSourceSnapshotInformation) SetNumLogicalBytesProtected(v int64)`

SetNumLogicalBytesProtected sets NumLogicalBytesProtected field to given value.

### HasNumLogicalBytesProtected

`func (o *ProtectionSourceSnapshotInformation) HasNumLogicalBytesProtected() bool`

HasNumLogicalBytesProtected returns a boolean if a field has been set.

### SetNumLogicalBytesProtectedNil

`func (o *ProtectionSourceSnapshotInformation) SetNumLogicalBytesProtectedNil(b bool)`

 SetNumLogicalBytesProtectedNil sets the value for NumLogicalBytesProtected to be an explicit nil

### UnsetNumLogicalBytesProtected
`func (o *ProtectionSourceSnapshotInformation) UnsetNumLogicalBytesProtected()`

UnsetNumLogicalBytesProtected ensures that no value is present for NumLogicalBytesProtected, not even an explicit nil
### GetPaginationCookie

`func (o *ProtectionSourceSnapshotInformation) GetPaginationCookie() int32`

GetPaginationCookie returns the PaginationCookie field if non-nil, zero value otherwise.

### GetPaginationCookieOk

`func (o *ProtectionSourceSnapshotInformation) GetPaginationCookieOk() (*int32, bool)`

GetPaginationCookieOk returns a tuple with the PaginationCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaginationCookie

`func (o *ProtectionSourceSnapshotInformation) SetPaginationCookie(v int32)`

SetPaginationCookie sets PaginationCookie field to given value.

### HasPaginationCookie

`func (o *ProtectionSourceSnapshotInformation) HasPaginationCookie() bool`

HasPaginationCookie returns a boolean if a field has been set.

### SetPaginationCookieNil

`func (o *ProtectionSourceSnapshotInformation) SetPaginationCookieNil(b bool)`

 SetPaginationCookieNil sets the value for PaginationCookie to be an explicit nil

### UnsetPaginationCookie
`func (o *ProtectionSourceSnapshotInformation) UnsetPaginationCookie()`

UnsetPaginationCookie ensures that no value is present for PaginationCookie, not even an explicit nil
### GetRunStatus

`func (o *ProtectionSourceSnapshotInformation) GetRunStatus() string`

GetRunStatus returns the RunStatus field if non-nil, zero value otherwise.

### GetRunStatusOk

`func (o *ProtectionSourceSnapshotInformation) GetRunStatusOk() (*string, bool)`

GetRunStatusOk returns a tuple with the RunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunStatus

`func (o *ProtectionSourceSnapshotInformation) SetRunStatus(v string)`

SetRunStatus sets RunStatus field to given value.

### HasRunStatus

`func (o *ProtectionSourceSnapshotInformation) HasRunStatus() bool`

HasRunStatus returns a boolean if a field has been set.

### SetRunStatusNil

`func (o *ProtectionSourceSnapshotInformation) SetRunStatusNil(b bool)`

 SetRunStatusNil sets the value for RunStatus to be an explicit nil

### UnsetRunStatus
`func (o *ProtectionSourceSnapshotInformation) UnsetRunStatus()`

UnsetRunStatus ensures that no value is present for RunStatus, not even an explicit nil
### GetRunType

`func (o *ProtectionSourceSnapshotInformation) GetRunType() string`

GetRunType returns the RunType field if non-nil, zero value otherwise.

### GetRunTypeOk

`func (o *ProtectionSourceSnapshotInformation) GetRunTypeOk() (*string, bool)`

GetRunTypeOk returns a tuple with the RunType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunType

`func (o *ProtectionSourceSnapshotInformation) SetRunType(v string)`

SetRunType sets RunType field to given value.

### HasRunType

`func (o *ProtectionSourceSnapshotInformation) HasRunType() bool`

HasRunType returns a boolean if a field has been set.

### SetRunTypeNil

`func (o *ProtectionSourceSnapshotInformation) SetRunTypeNil(b bool)`

 SetRunTypeNil sets the value for RunType to be an explicit nil

### UnsetRunType
`func (o *ProtectionSourceSnapshotInformation) UnsetRunType()`

UnsetRunType ensures that no value is present for RunType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


