# CancelProtectionJobRunParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CopyTaskUid** | Pointer to [**UniversalId**](UniversalId.md) |  | [optional] 
**JobRunId** | Pointer to **NullableInt64** | Run Id of a Protection Job Run that needs to be cancelled. If this Run id does not match the id of an active Run in the Protection job, the job Run is not cancelled and an error will be returned. | [optional] 

## Methods

### NewCancelProtectionJobRunParam

`func NewCancelProtectionJobRunParam() *CancelProtectionJobRunParam`

NewCancelProtectionJobRunParam instantiates a new CancelProtectionJobRunParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelProtectionJobRunParamWithDefaults

`func NewCancelProtectionJobRunParamWithDefaults() *CancelProtectionJobRunParam`

NewCancelProtectionJobRunParamWithDefaults instantiates a new CancelProtectionJobRunParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCopyTaskUid

`func (o *CancelProtectionJobRunParam) GetCopyTaskUid() UniversalId`

GetCopyTaskUid returns the CopyTaskUid field if non-nil, zero value otherwise.

### GetCopyTaskUidOk

`func (o *CancelProtectionJobRunParam) GetCopyTaskUidOk() (*UniversalId, bool)`

GetCopyTaskUidOk returns a tuple with the CopyTaskUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyTaskUid

`func (o *CancelProtectionJobRunParam) SetCopyTaskUid(v UniversalId)`

SetCopyTaskUid sets CopyTaskUid field to given value.

### HasCopyTaskUid

`func (o *CancelProtectionJobRunParam) HasCopyTaskUid() bool`

HasCopyTaskUid returns a boolean if a field has been set.

### GetJobRunId

`func (o *CancelProtectionJobRunParam) GetJobRunId() int64`

GetJobRunId returns the JobRunId field if non-nil, zero value otherwise.

### GetJobRunIdOk

`func (o *CancelProtectionJobRunParam) GetJobRunIdOk() (*int64, bool)`

GetJobRunIdOk returns a tuple with the JobRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobRunId

`func (o *CancelProtectionJobRunParam) SetJobRunId(v int64)`

SetJobRunId sets JobRunId field to given value.

### HasJobRunId

`func (o *CancelProtectionJobRunParam) HasJobRunId() bool`

HasJobRunId returns a boolean if a field has been set.

### SetJobRunIdNil

`func (o *CancelProtectionJobRunParam) SetJobRunIdNil(b bool)`

 SetJobRunIdNil sets the value for JobRunId to be an explicit nil

### UnsetJobRunId
`func (o *CancelProtectionJobRunParam) UnsetJobRunId()`

UnsetJobRunId ensures that no value is present for JobRunId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


