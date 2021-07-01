# RemoteRestoreIndexingStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time of the time range that is being indexed. The indexing task is creating an index of the Job Runs that occurred between the startTimeUsecs and this endTimeUsecs. This field is recorded as a Unix epoch Timestamp (in microseconds). | [optional] 
**Error** | Pointer to **NullableString** | Specifies the error message if the indexing Job/task fails. | [optional] 
**IndexingTaskEndTimeUsecs** | Pointer to **NullableInt64** | Specifies when the indexing task completed. This time is recorded as a Unix epoch Timestamp (in microseconds). This field is not set if the indexing task is still in progress. | [optional] 
**IndexingTaskStartTimeUsecs** | Pointer to **NullableInt64** | Specifies when the indexing task started. This time is recorded as a Unix epoch Timestamp (in microseconds). | [optional] 
**IndexingTaskStatus** | Pointer to **NullableString** | Specifies the status of the indexing Job/task. &#39;kJobRunning&#39; indicates that the Job/task is currently running. &#39;kJobFinished&#39; indicates that the Job/task completed and finished. &#39;kJobFailed&#39; indicates that the Job/task failed and did not complete. &#39;kJobCanceled&#39; indicates that the Job/task was canceled. &#39;kJobPaused&#39; indicates the Job/task is paused. | [optional] 
**IndexingTaskUid** | Pointer to [**NullableUniversalId**](UniversalId.md) | Specifies the unique id of the indexing task assigned by this Cluster. | [optional] 
**LatestExpiryTimeUsecs** | Pointer to **NullableInt64** | For all the Snapshots retrieved by this Job, specifies the latest time when a Snapshot expires. | [optional] 
**ProgressMonitorTask** | Pointer to **NullableString** | Specifies the path to progress monitor task to track the progress of building the index. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of the time range that is being indexed. The indexing task is creating an index of the Job Runs that occurred between this startTimeUsecs and the endTimeUsecs. This field is recorded as a Unix epoch Timestamp (in microseconds). | [optional] 

## Methods

### NewRemoteRestoreIndexingStatus

`func NewRemoteRestoreIndexingStatus() *RemoteRestoreIndexingStatus`

NewRemoteRestoreIndexingStatus instantiates a new RemoteRestoreIndexingStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteRestoreIndexingStatusWithDefaults

`func NewRemoteRestoreIndexingStatusWithDefaults() *RemoteRestoreIndexingStatus`

NewRemoteRestoreIndexingStatusWithDefaults instantiates a new RemoteRestoreIndexingStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTimeUsecs

`func (o *RemoteRestoreIndexingStatus) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *RemoteRestoreIndexingStatus) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *RemoteRestoreIndexingStatus) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *RemoteRestoreIndexingStatus) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *RemoteRestoreIndexingStatus) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *RemoteRestoreIndexingStatus) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetError

`func (o *RemoteRestoreIndexingStatus) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RemoteRestoreIndexingStatus) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RemoteRestoreIndexingStatus) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *RemoteRestoreIndexingStatus) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *RemoteRestoreIndexingStatus) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *RemoteRestoreIndexingStatus) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetIndexingTaskEndTimeUsecs

`func (o *RemoteRestoreIndexingStatus) GetIndexingTaskEndTimeUsecs() int64`

GetIndexingTaskEndTimeUsecs returns the IndexingTaskEndTimeUsecs field if non-nil, zero value otherwise.

### GetIndexingTaskEndTimeUsecsOk

`func (o *RemoteRestoreIndexingStatus) GetIndexingTaskEndTimeUsecsOk() (*int64, bool)`

GetIndexingTaskEndTimeUsecsOk returns a tuple with the IndexingTaskEndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingTaskEndTimeUsecs

`func (o *RemoteRestoreIndexingStatus) SetIndexingTaskEndTimeUsecs(v int64)`

SetIndexingTaskEndTimeUsecs sets IndexingTaskEndTimeUsecs field to given value.

### HasIndexingTaskEndTimeUsecs

`func (o *RemoteRestoreIndexingStatus) HasIndexingTaskEndTimeUsecs() bool`

HasIndexingTaskEndTimeUsecs returns a boolean if a field has been set.

### SetIndexingTaskEndTimeUsecsNil

`func (o *RemoteRestoreIndexingStatus) SetIndexingTaskEndTimeUsecsNil(b bool)`

 SetIndexingTaskEndTimeUsecsNil sets the value for IndexingTaskEndTimeUsecs to be an explicit nil

### UnsetIndexingTaskEndTimeUsecs
`func (o *RemoteRestoreIndexingStatus) UnsetIndexingTaskEndTimeUsecs()`

UnsetIndexingTaskEndTimeUsecs ensures that no value is present for IndexingTaskEndTimeUsecs, not even an explicit nil
### GetIndexingTaskStartTimeUsecs

`func (o *RemoteRestoreIndexingStatus) GetIndexingTaskStartTimeUsecs() int64`

GetIndexingTaskStartTimeUsecs returns the IndexingTaskStartTimeUsecs field if non-nil, zero value otherwise.

### GetIndexingTaskStartTimeUsecsOk

`func (o *RemoteRestoreIndexingStatus) GetIndexingTaskStartTimeUsecsOk() (*int64, bool)`

GetIndexingTaskStartTimeUsecsOk returns a tuple with the IndexingTaskStartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingTaskStartTimeUsecs

`func (o *RemoteRestoreIndexingStatus) SetIndexingTaskStartTimeUsecs(v int64)`

SetIndexingTaskStartTimeUsecs sets IndexingTaskStartTimeUsecs field to given value.

### HasIndexingTaskStartTimeUsecs

`func (o *RemoteRestoreIndexingStatus) HasIndexingTaskStartTimeUsecs() bool`

HasIndexingTaskStartTimeUsecs returns a boolean if a field has been set.

### SetIndexingTaskStartTimeUsecsNil

`func (o *RemoteRestoreIndexingStatus) SetIndexingTaskStartTimeUsecsNil(b bool)`

 SetIndexingTaskStartTimeUsecsNil sets the value for IndexingTaskStartTimeUsecs to be an explicit nil

### UnsetIndexingTaskStartTimeUsecs
`func (o *RemoteRestoreIndexingStatus) UnsetIndexingTaskStartTimeUsecs()`

UnsetIndexingTaskStartTimeUsecs ensures that no value is present for IndexingTaskStartTimeUsecs, not even an explicit nil
### GetIndexingTaskStatus

`func (o *RemoteRestoreIndexingStatus) GetIndexingTaskStatus() string`

GetIndexingTaskStatus returns the IndexingTaskStatus field if non-nil, zero value otherwise.

### GetIndexingTaskStatusOk

`func (o *RemoteRestoreIndexingStatus) GetIndexingTaskStatusOk() (*string, bool)`

GetIndexingTaskStatusOk returns a tuple with the IndexingTaskStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingTaskStatus

`func (o *RemoteRestoreIndexingStatus) SetIndexingTaskStatus(v string)`

SetIndexingTaskStatus sets IndexingTaskStatus field to given value.

### HasIndexingTaskStatus

`func (o *RemoteRestoreIndexingStatus) HasIndexingTaskStatus() bool`

HasIndexingTaskStatus returns a boolean if a field has been set.

### SetIndexingTaskStatusNil

`func (o *RemoteRestoreIndexingStatus) SetIndexingTaskStatusNil(b bool)`

 SetIndexingTaskStatusNil sets the value for IndexingTaskStatus to be an explicit nil

### UnsetIndexingTaskStatus
`func (o *RemoteRestoreIndexingStatus) UnsetIndexingTaskStatus()`

UnsetIndexingTaskStatus ensures that no value is present for IndexingTaskStatus, not even an explicit nil
### GetIndexingTaskUid

`func (o *RemoteRestoreIndexingStatus) GetIndexingTaskUid() UniversalId`

GetIndexingTaskUid returns the IndexingTaskUid field if non-nil, zero value otherwise.

### GetIndexingTaskUidOk

`func (o *RemoteRestoreIndexingStatus) GetIndexingTaskUidOk() (*UniversalId, bool)`

GetIndexingTaskUidOk returns a tuple with the IndexingTaskUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingTaskUid

`func (o *RemoteRestoreIndexingStatus) SetIndexingTaskUid(v UniversalId)`

SetIndexingTaskUid sets IndexingTaskUid field to given value.

### HasIndexingTaskUid

`func (o *RemoteRestoreIndexingStatus) HasIndexingTaskUid() bool`

HasIndexingTaskUid returns a boolean if a field has been set.

### SetIndexingTaskUidNil

`func (o *RemoteRestoreIndexingStatus) SetIndexingTaskUidNil(b bool)`

 SetIndexingTaskUidNil sets the value for IndexingTaskUid to be an explicit nil

### UnsetIndexingTaskUid
`func (o *RemoteRestoreIndexingStatus) UnsetIndexingTaskUid()`

UnsetIndexingTaskUid ensures that no value is present for IndexingTaskUid, not even an explicit nil
### GetLatestExpiryTimeUsecs

`func (o *RemoteRestoreIndexingStatus) GetLatestExpiryTimeUsecs() int64`

GetLatestExpiryTimeUsecs returns the LatestExpiryTimeUsecs field if non-nil, zero value otherwise.

### GetLatestExpiryTimeUsecsOk

`func (o *RemoteRestoreIndexingStatus) GetLatestExpiryTimeUsecsOk() (*int64, bool)`

GetLatestExpiryTimeUsecsOk returns a tuple with the LatestExpiryTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestExpiryTimeUsecs

`func (o *RemoteRestoreIndexingStatus) SetLatestExpiryTimeUsecs(v int64)`

SetLatestExpiryTimeUsecs sets LatestExpiryTimeUsecs field to given value.

### HasLatestExpiryTimeUsecs

`func (o *RemoteRestoreIndexingStatus) HasLatestExpiryTimeUsecs() bool`

HasLatestExpiryTimeUsecs returns a boolean if a field has been set.

### SetLatestExpiryTimeUsecsNil

`func (o *RemoteRestoreIndexingStatus) SetLatestExpiryTimeUsecsNil(b bool)`

 SetLatestExpiryTimeUsecsNil sets the value for LatestExpiryTimeUsecs to be an explicit nil

### UnsetLatestExpiryTimeUsecs
`func (o *RemoteRestoreIndexingStatus) UnsetLatestExpiryTimeUsecs()`

UnsetLatestExpiryTimeUsecs ensures that no value is present for LatestExpiryTimeUsecs, not even an explicit nil
### GetProgressMonitorTask

`func (o *RemoteRestoreIndexingStatus) GetProgressMonitorTask() string`

GetProgressMonitorTask returns the ProgressMonitorTask field if non-nil, zero value otherwise.

### GetProgressMonitorTaskOk

`func (o *RemoteRestoreIndexingStatus) GetProgressMonitorTaskOk() (*string, bool)`

GetProgressMonitorTaskOk returns a tuple with the ProgressMonitorTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressMonitorTask

`func (o *RemoteRestoreIndexingStatus) SetProgressMonitorTask(v string)`

SetProgressMonitorTask sets ProgressMonitorTask field to given value.

### HasProgressMonitorTask

`func (o *RemoteRestoreIndexingStatus) HasProgressMonitorTask() bool`

HasProgressMonitorTask returns a boolean if a field has been set.

### SetProgressMonitorTaskNil

`func (o *RemoteRestoreIndexingStatus) SetProgressMonitorTaskNil(b bool)`

 SetProgressMonitorTaskNil sets the value for ProgressMonitorTask to be an explicit nil

### UnsetProgressMonitorTask
`func (o *RemoteRestoreIndexingStatus) UnsetProgressMonitorTask()`

UnsetProgressMonitorTask ensures that no value is present for ProgressMonitorTask, not even an explicit nil
### GetStartTimeUsecs

`func (o *RemoteRestoreIndexingStatus) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *RemoteRestoreIndexingStatus) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *RemoteRestoreIndexingStatus) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *RemoteRestoreIndexingStatus) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *RemoteRestoreIndexingStatus) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *RemoteRestoreIndexingStatus) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


