# PerformActionOnProtectionGroupRunRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **NullableString** | Specifies the type of the action which will be performed on protection runs. | 
**PauseParams** | Pointer to [**[]PauseProtectionRunActionParams**](PauseProtectionRunActionParams.md) | Specifies the pause action params for a protection run. | [optional] 
**ResumeParams** | Pointer to [**[]ResumeProtectionRunActionParams**](ResumeProtectionRunActionParams.md) | Specifies the resume action params for a protection run. | [optional] 
**CancelParams** | Pointer to [**[]CancelProtectionGroupRunRequest**](CancelProtectionGroupRunRequest.md) | Specifies the cancel action params for a protection run. | [optional] 

## Methods

### NewPerformActionOnProtectionGroupRunRequest

`func NewPerformActionOnProtectionGroupRunRequest(action NullableString, ) *PerformActionOnProtectionGroupRunRequest`

NewPerformActionOnProtectionGroupRunRequest instantiates a new PerformActionOnProtectionGroupRunRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerformActionOnProtectionGroupRunRequestWithDefaults

`func NewPerformActionOnProtectionGroupRunRequestWithDefaults() *PerformActionOnProtectionGroupRunRequest`

NewPerformActionOnProtectionGroupRunRequestWithDefaults instantiates a new PerformActionOnProtectionGroupRunRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PerformActionOnProtectionGroupRunRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PerformActionOnProtectionGroupRunRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PerformActionOnProtectionGroupRunRequest) SetAction(v string)`

SetAction sets Action field to given value.


### SetActionNil

`func (o *PerformActionOnProtectionGroupRunRequest) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *PerformActionOnProtectionGroupRunRequest) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetPauseParams

`func (o *PerformActionOnProtectionGroupRunRequest) GetPauseParams() []PauseProtectionRunActionParams`

GetPauseParams returns the PauseParams field if non-nil, zero value otherwise.

### GetPauseParamsOk

`func (o *PerformActionOnProtectionGroupRunRequest) GetPauseParamsOk() (*[]PauseProtectionRunActionParams, bool)`

GetPauseParamsOk returns a tuple with the PauseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseParams

`func (o *PerformActionOnProtectionGroupRunRequest) SetPauseParams(v []PauseProtectionRunActionParams)`

SetPauseParams sets PauseParams field to given value.

### HasPauseParams

`func (o *PerformActionOnProtectionGroupRunRequest) HasPauseParams() bool`

HasPauseParams returns a boolean if a field has been set.

### GetResumeParams

`func (o *PerformActionOnProtectionGroupRunRequest) GetResumeParams() []ResumeProtectionRunActionParams`

GetResumeParams returns the ResumeParams field if non-nil, zero value otherwise.

### GetResumeParamsOk

`func (o *PerformActionOnProtectionGroupRunRequest) GetResumeParamsOk() (*[]ResumeProtectionRunActionParams, bool)`

GetResumeParamsOk returns a tuple with the ResumeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeParams

`func (o *PerformActionOnProtectionGroupRunRequest) SetResumeParams(v []ResumeProtectionRunActionParams)`

SetResumeParams sets ResumeParams field to given value.

### HasResumeParams

`func (o *PerformActionOnProtectionGroupRunRequest) HasResumeParams() bool`

HasResumeParams returns a boolean if a field has been set.

### GetCancelParams

`func (o *PerformActionOnProtectionGroupRunRequest) GetCancelParams() []CancelProtectionGroupRunRequest`

GetCancelParams returns the CancelParams field if non-nil, zero value otherwise.

### GetCancelParamsOk

`func (o *PerformActionOnProtectionGroupRunRequest) GetCancelParamsOk() (*[]CancelProtectionGroupRunRequest, bool)`

GetCancelParamsOk returns a tuple with the CancelParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelParams

`func (o *PerformActionOnProtectionGroupRunRequest) SetCancelParams(v []CancelProtectionGroupRunRequest)`

SetCancelParams sets CancelParams field to given value.

### HasCancelParams

`func (o *PerformActionOnProtectionGroupRunRequest) HasCancelParams() bool`

HasCancelParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


