# UpdateProtectionJobRunsParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobRuns** | Pointer to [**[]UpdateProtectionJobRun**](UpdateProtectionJobRun.md) | Array of Job Runs.  Specifies the Job Runs to update with a new expiration times. | [optional] 

## Methods

### NewUpdateProtectionJobRunsParam

`func NewUpdateProtectionJobRunsParam() *UpdateProtectionJobRunsParam`

NewUpdateProtectionJobRunsParam instantiates a new UpdateProtectionJobRunsParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProtectionJobRunsParamWithDefaults

`func NewUpdateProtectionJobRunsParamWithDefaults() *UpdateProtectionJobRunsParam`

NewUpdateProtectionJobRunsParamWithDefaults instantiates a new UpdateProtectionJobRunsParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobRuns

`func (o *UpdateProtectionJobRunsParam) GetJobRuns() []UpdateProtectionJobRun`

GetJobRuns returns the JobRuns field if non-nil, zero value otherwise.

### GetJobRunsOk

`func (o *UpdateProtectionJobRunsParam) GetJobRunsOk() (*[]UpdateProtectionJobRun, bool)`

GetJobRunsOk returns a tuple with the JobRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobRuns

`func (o *UpdateProtectionJobRunsParam) SetJobRuns(v []UpdateProtectionJobRun)`

SetJobRuns sets JobRuns field to given value.

### HasJobRuns

`func (o *UpdateProtectionJobRunsParam) HasJobRuns() bool`

HasJobRuns returns a boolean if a field has been set.

### SetJobRunsNil

`func (o *UpdateProtectionJobRunsParam) SetJobRunsNil(b bool)`

 SetJobRunsNil sets the value for JobRuns to be an explicit nil

### UnsetJobRuns
`func (o *UpdateProtectionJobRunsParam) UnsetJobRuns()`

UnsetJobRuns ensures that no value is present for JobRuns, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


