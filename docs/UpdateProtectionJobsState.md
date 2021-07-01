# UpdateProtectionJobsState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedJobIds** | Pointer to **[]int64** | Specifies a list of Protection Job ids for which updation of state failed. | [optional] 
**SuccessfulJobIds** | Pointer to **[]int64** | Specifies a list of Protection Job ids for which updation of state is successful. | [optional] 

## Methods

### NewUpdateProtectionJobsState

`func NewUpdateProtectionJobsState() *UpdateProtectionJobsState`

NewUpdateProtectionJobsState instantiates a new UpdateProtectionJobsState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProtectionJobsStateWithDefaults

`func NewUpdateProtectionJobsStateWithDefaults() *UpdateProtectionJobsState`

NewUpdateProtectionJobsStateWithDefaults instantiates a new UpdateProtectionJobsState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailedJobIds

`func (o *UpdateProtectionJobsState) GetFailedJobIds() []int64`

GetFailedJobIds returns the FailedJobIds field if non-nil, zero value otherwise.

### GetFailedJobIdsOk

`func (o *UpdateProtectionJobsState) GetFailedJobIdsOk() (*[]int64, bool)`

GetFailedJobIdsOk returns a tuple with the FailedJobIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedJobIds

`func (o *UpdateProtectionJobsState) SetFailedJobIds(v []int64)`

SetFailedJobIds sets FailedJobIds field to given value.

### HasFailedJobIds

`func (o *UpdateProtectionJobsState) HasFailedJobIds() bool`

HasFailedJobIds returns a boolean if a field has been set.

### SetFailedJobIdsNil

`func (o *UpdateProtectionJobsState) SetFailedJobIdsNil(b bool)`

 SetFailedJobIdsNil sets the value for FailedJobIds to be an explicit nil

### UnsetFailedJobIds
`func (o *UpdateProtectionJobsState) UnsetFailedJobIds()`

UnsetFailedJobIds ensures that no value is present for FailedJobIds, not even an explicit nil
### GetSuccessfulJobIds

`func (o *UpdateProtectionJobsState) GetSuccessfulJobIds() []int64`

GetSuccessfulJobIds returns the SuccessfulJobIds field if non-nil, zero value otherwise.

### GetSuccessfulJobIdsOk

`func (o *UpdateProtectionJobsState) GetSuccessfulJobIdsOk() (*[]int64, bool)`

GetSuccessfulJobIdsOk returns a tuple with the SuccessfulJobIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulJobIds

`func (o *UpdateProtectionJobsState) SetSuccessfulJobIds(v []int64)`

SetSuccessfulJobIds sets SuccessfulJobIds field to given value.

### HasSuccessfulJobIds

`func (o *UpdateProtectionJobsState) HasSuccessfulJobIds() bool`

HasSuccessfulJobIds returns a boolean if a field has been set.

### SetSuccessfulJobIdsNil

`func (o *UpdateProtectionJobsState) SetSuccessfulJobIdsNil(b bool)`

 SetSuccessfulJobIdsNil sets the value for SuccessfulJobIds to be an explicit nil

### UnsetSuccessfulJobIds
`func (o *UpdateProtectionJobsState) UnsetSuccessfulJobIds()`

UnsetSuccessfulJobIds ensures that no value is present for SuccessfulJobIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


