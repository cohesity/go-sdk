# ChangeProtectionJobStateParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pause** | Pointer to **NullableBool** | If true, the specified Protection Job is paused and no new Runs of the Job are started. Any Runs that were executing continue to run. If false and the Protection Job was in a paused state, the Protection Job resumes and new Runs are started according to the schedule defined in the associated Policy. | [optional] 
**PauseReason** | Pointer to **NullableInt32** | Specifies the reason of pausing the job so that depending on the pause reason, only specific jobs can be resumed. All the jobs paused manually by the user will be identified by nil PauseReason. | [optional] 

## Methods

### NewChangeProtectionJobStateParam

`func NewChangeProtectionJobStateParam() *ChangeProtectionJobStateParam`

NewChangeProtectionJobStateParam instantiates a new ChangeProtectionJobStateParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeProtectionJobStateParamWithDefaults

`func NewChangeProtectionJobStateParamWithDefaults() *ChangeProtectionJobStateParam`

NewChangeProtectionJobStateParamWithDefaults instantiates a new ChangeProtectionJobStateParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPause

`func (o *ChangeProtectionJobStateParam) GetPause() bool`

GetPause returns the Pause field if non-nil, zero value otherwise.

### GetPauseOk

`func (o *ChangeProtectionJobStateParam) GetPauseOk() (*bool, bool)`

GetPauseOk returns a tuple with the Pause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPause

`func (o *ChangeProtectionJobStateParam) SetPause(v bool)`

SetPause sets Pause field to given value.

### HasPause

`func (o *ChangeProtectionJobStateParam) HasPause() bool`

HasPause returns a boolean if a field has been set.

### SetPauseNil

`func (o *ChangeProtectionJobStateParam) SetPauseNil(b bool)`

 SetPauseNil sets the value for Pause to be an explicit nil

### UnsetPause
`func (o *ChangeProtectionJobStateParam) UnsetPause()`

UnsetPause ensures that no value is present for Pause, not even an explicit nil
### GetPauseReason

`func (o *ChangeProtectionJobStateParam) GetPauseReason() int32`

GetPauseReason returns the PauseReason field if non-nil, zero value otherwise.

### GetPauseReasonOk

`func (o *ChangeProtectionJobStateParam) GetPauseReasonOk() (*int32, bool)`

GetPauseReasonOk returns a tuple with the PauseReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseReason

`func (o *ChangeProtectionJobStateParam) SetPauseReason(v int32)`

SetPauseReason sets PauseReason field to given value.

### HasPauseReason

`func (o *ChangeProtectionJobStateParam) HasPauseReason() bool`

HasPauseReason returns a boolean if a field has been set.

### SetPauseReasonNil

`func (o *ChangeProtectionJobStateParam) SetPauseReasonNil(b bool)`

 SetPauseReasonNil sets the value for PauseReason to be an explicit nil

### UnsetPauseReason
`func (o *ChangeProtectionJobStateParam) UnsetPauseReason()`

UnsetPauseReason ensures that no value is present for PauseReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


