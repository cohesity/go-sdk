# ProtectionShellInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time of the Protection Run. The end time is specified as a Unix epoch Timestamp (in microseconds). | [optional] 
**Error** | Pointer to **NullableString** | Specifies if an error occurred (if any) while running this task. This field is populated when the status is equal to &#39;kFailure&#39;. | [optional] 
**JobRunId** | Pointer to **NullableInt64** | Specifies the id of the Job Run that ran the backup task and the copy tasks. | [optional] 
**RunType** | Pointer to **NullableString** | Specifies the type of backup such as &#39;kRegular&#39;, &#39;kFull&#39;, &#39;kLog&#39; or &#39;kSystem&#39;. &#39;kRegular&#39; indicates a incremental (CBT) backup. Incremental backups utilizing CBT (if supported) are captured of the target protection objects. The first run of a kRegular schedule captures all the blocks. &#39;kFull&#39; indicates a full (no CBT) backup. A complete backup (all blocks) of the target protection objects are always captured and Change Block Tracking (CBT) is not utilized. &#39;kLog&#39; indicates a Database Log backup. Capture the database transaction logs to allow rolling back to a specific point in time. &#39;kSystem&#39; indicates a system backup. System backups are used to do bare metal recovery of the system to a specific point in time. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of the Protection Run. The start time is specified as a Unix epoch Timestamp (in microseconds). This time is when the task is queued to an internal queue where tasks are waiting to run. | [optional] 
**Status** | Pointer to **NullableString** | Specifies the status of Backup task such as &#39;kRunning&#39;, &#39;kSuccess&#39;, &#39;kFailure&#39; etc. &#39;kAccepted&#39; indicates the task is queued to run but not yet running. &#39;kRunning&#39; indicates the task is running. &#39;kCanceling&#39; indicates a request to cancel the task has occurred but the task is not yet canceled. &#39;kCanceled&#39; indicates the task has been canceled. &#39;kSuccess&#39; indicates the task was successful. &#39;kFailure&#39; indicates the task failed. &#39;kWarning&#39; indicates the task has finished with warning. &#39;kOnHold&#39; indicates the task is kept onHold. &#39;kMissed&#39; indicates the task is missed. | [optional] 

## Methods

### NewProtectionShellInfo

`func NewProtectionShellInfo() *ProtectionShellInfo`

NewProtectionShellInfo instantiates a new ProtectionShellInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionShellInfoWithDefaults

`func NewProtectionShellInfoWithDefaults() *ProtectionShellInfo`

NewProtectionShellInfoWithDefaults instantiates a new ProtectionShellInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTimeUsecs

`func (o *ProtectionShellInfo) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *ProtectionShellInfo) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *ProtectionShellInfo) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *ProtectionShellInfo) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *ProtectionShellInfo) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *ProtectionShellInfo) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetError

`func (o *ProtectionShellInfo) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ProtectionShellInfo) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ProtectionShellInfo) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ProtectionShellInfo) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *ProtectionShellInfo) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ProtectionShellInfo) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetJobRunId

`func (o *ProtectionShellInfo) GetJobRunId() int64`

GetJobRunId returns the JobRunId field if non-nil, zero value otherwise.

### GetJobRunIdOk

`func (o *ProtectionShellInfo) GetJobRunIdOk() (*int64, bool)`

GetJobRunIdOk returns a tuple with the JobRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobRunId

`func (o *ProtectionShellInfo) SetJobRunId(v int64)`

SetJobRunId sets JobRunId field to given value.

### HasJobRunId

`func (o *ProtectionShellInfo) HasJobRunId() bool`

HasJobRunId returns a boolean if a field has been set.

### SetJobRunIdNil

`func (o *ProtectionShellInfo) SetJobRunIdNil(b bool)`

 SetJobRunIdNil sets the value for JobRunId to be an explicit nil

### UnsetJobRunId
`func (o *ProtectionShellInfo) UnsetJobRunId()`

UnsetJobRunId ensures that no value is present for JobRunId, not even an explicit nil
### GetRunType

`func (o *ProtectionShellInfo) GetRunType() string`

GetRunType returns the RunType field if non-nil, zero value otherwise.

### GetRunTypeOk

`func (o *ProtectionShellInfo) GetRunTypeOk() (*string, bool)`

GetRunTypeOk returns a tuple with the RunType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunType

`func (o *ProtectionShellInfo) SetRunType(v string)`

SetRunType sets RunType field to given value.

### HasRunType

`func (o *ProtectionShellInfo) HasRunType() bool`

HasRunType returns a boolean if a field has been set.

### SetRunTypeNil

`func (o *ProtectionShellInfo) SetRunTypeNil(b bool)`

 SetRunTypeNil sets the value for RunType to be an explicit nil

### UnsetRunType
`func (o *ProtectionShellInfo) UnsetRunType()`

UnsetRunType ensures that no value is present for RunType, not even an explicit nil
### GetStartTimeUsecs

`func (o *ProtectionShellInfo) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *ProtectionShellInfo) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *ProtectionShellInfo) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *ProtectionShellInfo) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *ProtectionShellInfo) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *ProtectionShellInfo) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetStatus

`func (o *ProtectionShellInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProtectionShellInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProtectionShellInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProtectionShellInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ProtectionShellInfo) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ProtectionShellInfo) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


