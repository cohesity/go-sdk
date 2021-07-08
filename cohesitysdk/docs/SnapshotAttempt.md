# SnapshotAttempt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttemptNumber** | Pointer to **NullableInt64** | Specifies the number of the attempts made by the Job Run to capture a snapshot of the object. For example, if an snapshot is successfully captured after three attempts, this field equals 3. | [optional] 
**JobRunId** | Pointer to **NullableInt64** | Specifies the id of the Job Run that captured the snapshot. | [optional] 
**StartedTimeUsecs** | Pointer to **NullableInt64** | Specifies the time when the Job Run starts capturing a snapshot. Specified as a Unix epoch Timestamp (in microseconds). | [optional] 

## Methods

### NewSnapshotAttempt

`func NewSnapshotAttempt() *SnapshotAttempt`

NewSnapshotAttempt instantiates a new SnapshotAttempt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotAttemptWithDefaults

`func NewSnapshotAttemptWithDefaults() *SnapshotAttempt`

NewSnapshotAttemptWithDefaults instantiates a new SnapshotAttempt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttemptNumber

`func (o *SnapshotAttempt) GetAttemptNumber() int64`

GetAttemptNumber returns the AttemptNumber field if non-nil, zero value otherwise.

### GetAttemptNumberOk

`func (o *SnapshotAttempt) GetAttemptNumberOk() (*int64, bool)`

GetAttemptNumberOk returns a tuple with the AttemptNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptNumber

`func (o *SnapshotAttempt) SetAttemptNumber(v int64)`

SetAttemptNumber sets AttemptNumber field to given value.

### HasAttemptNumber

`func (o *SnapshotAttempt) HasAttemptNumber() bool`

HasAttemptNumber returns a boolean if a field has been set.

### SetAttemptNumberNil

`func (o *SnapshotAttempt) SetAttemptNumberNil(b bool)`

 SetAttemptNumberNil sets the value for AttemptNumber to be an explicit nil

### UnsetAttemptNumber
`func (o *SnapshotAttempt) UnsetAttemptNumber()`

UnsetAttemptNumber ensures that no value is present for AttemptNumber, not even an explicit nil
### GetJobRunId

`func (o *SnapshotAttempt) GetJobRunId() int64`

GetJobRunId returns the JobRunId field if non-nil, zero value otherwise.

### GetJobRunIdOk

`func (o *SnapshotAttempt) GetJobRunIdOk() (*int64, bool)`

GetJobRunIdOk returns a tuple with the JobRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobRunId

`func (o *SnapshotAttempt) SetJobRunId(v int64)`

SetJobRunId sets JobRunId field to given value.

### HasJobRunId

`func (o *SnapshotAttempt) HasJobRunId() bool`

HasJobRunId returns a boolean if a field has been set.

### SetJobRunIdNil

`func (o *SnapshotAttempt) SetJobRunIdNil(b bool)`

 SetJobRunIdNil sets the value for JobRunId to be an explicit nil

### UnsetJobRunId
`func (o *SnapshotAttempt) UnsetJobRunId()`

UnsetJobRunId ensures that no value is present for JobRunId, not even an explicit nil
### GetStartedTimeUsecs

`func (o *SnapshotAttempt) GetStartedTimeUsecs() int64`

GetStartedTimeUsecs returns the StartedTimeUsecs field if non-nil, zero value otherwise.

### GetStartedTimeUsecsOk

`func (o *SnapshotAttempt) GetStartedTimeUsecsOk() (*int64, bool)`

GetStartedTimeUsecsOk returns a tuple with the StartedTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedTimeUsecs

`func (o *SnapshotAttempt) SetStartedTimeUsecs(v int64)`

SetStartedTimeUsecs sets StartedTimeUsecs field to given value.

### HasStartedTimeUsecs

`func (o *SnapshotAttempt) HasStartedTimeUsecs() bool`

HasStartedTimeUsecs returns a boolean if a field has been set.

### SetStartedTimeUsecsNil

`func (o *SnapshotAttempt) SetStartedTimeUsecsNil(b bool)`

 SetStartedTimeUsecsNil sets the value for StartedTimeUsecs to be an explicit nil

### UnsetStartedTimeUsecs
`func (o *SnapshotAttempt) UnsetStartedTimeUsecs()`

UnsetStartedTimeUsecs ensures that no value is present for StartedTimeUsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


