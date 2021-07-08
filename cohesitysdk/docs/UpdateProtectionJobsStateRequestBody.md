# UpdateProtectionJobsStateRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **NullableString** | Specifies the action to be performed on all the specified Protection Jobs. Specifies the type of action to be performed on Protection Job. &#39;kActivate&#39; specifies that Protection Job should be activated. &#39;kDeactivate&#39; specifies that Protection Job should be deactivated. &#39;kPause&#39; specifies that Protection Job should be paused. &#39;kResume&#39; specifies that Protection Job should be resumed. | [optional] 
**JobIds** | Pointer to **[]int64** | Specifies a list of Protection Job ids for which the state should change. | [optional] 

## Methods

### NewUpdateProtectionJobsStateRequestBody

`func NewUpdateProtectionJobsStateRequestBody() *UpdateProtectionJobsStateRequestBody`

NewUpdateProtectionJobsStateRequestBody instantiates a new UpdateProtectionJobsStateRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProtectionJobsStateRequestBodyWithDefaults

`func NewUpdateProtectionJobsStateRequestBodyWithDefaults() *UpdateProtectionJobsStateRequestBody`

NewUpdateProtectionJobsStateRequestBodyWithDefaults instantiates a new UpdateProtectionJobsStateRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *UpdateProtectionJobsStateRequestBody) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UpdateProtectionJobsStateRequestBody) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UpdateProtectionJobsStateRequestBody) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *UpdateProtectionJobsStateRequestBody) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *UpdateProtectionJobsStateRequestBody) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *UpdateProtectionJobsStateRequestBody) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetJobIds

`func (o *UpdateProtectionJobsStateRequestBody) GetJobIds() []int64`

GetJobIds returns the JobIds field if non-nil, zero value otherwise.

### GetJobIdsOk

`func (o *UpdateProtectionJobsStateRequestBody) GetJobIdsOk() (*[]int64, bool)`

GetJobIdsOk returns a tuple with the JobIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobIds

`func (o *UpdateProtectionJobsStateRequestBody) SetJobIds(v []int64)`

SetJobIds sets JobIds field to given value.

### HasJobIds

`func (o *UpdateProtectionJobsStateRequestBody) HasJobIds() bool`

HasJobIds returns a boolean if a field has been set.

### SetJobIdsNil

`func (o *UpdateProtectionJobsStateRequestBody) SetJobIdsNil(b bool)`

 SetJobIdsNil sets the value for JobIds to be an explicit nil

### UnsetJobIds
`func (o *UpdateProtectionJobsStateRequestBody) UnsetJobIds()`

UnsetJobIds ensures that no value is present for JobIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


