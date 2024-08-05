# PerformRunActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **NullableString** | Specifies the type of the action is performed on protection runs. | [optional] 
**CancelParams** | Pointer to [**[]CancelProtectionGroupRunResponseParams**](CancelProtectionGroupRunResponseParams.md) | Specifies the cancel action response params. | [optional] 
**PauseParams** | Pointer to [**[]PauseProtectionRunActionResponseParams**](PauseProtectionRunActionResponseParams.md) | Specifies the pause action response params. | [optional] 
**ResumeParams** | Pointer to [**[]ResumeProtectionRunActionResponseParams**](ResumeProtectionRunActionResponseParams.md) | Specifies the resume action response params. | [optional] 

## Methods

### NewPerformRunActionResponse

`func NewPerformRunActionResponse() *PerformRunActionResponse`

NewPerformRunActionResponse instantiates a new PerformRunActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerformRunActionResponseWithDefaults

`func NewPerformRunActionResponseWithDefaults() *PerformRunActionResponse`

NewPerformRunActionResponseWithDefaults instantiates a new PerformRunActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PerformRunActionResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PerformRunActionResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PerformRunActionResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *PerformRunActionResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *PerformRunActionResponse) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *PerformRunActionResponse) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetCancelParams

`func (o *PerformRunActionResponse) GetCancelParams() []CancelProtectionGroupRunResponseParams`

GetCancelParams returns the CancelParams field if non-nil, zero value otherwise.

### GetCancelParamsOk

`func (o *PerformRunActionResponse) GetCancelParamsOk() (*[]CancelProtectionGroupRunResponseParams, bool)`

GetCancelParamsOk returns a tuple with the CancelParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelParams

`func (o *PerformRunActionResponse) SetCancelParams(v []CancelProtectionGroupRunResponseParams)`

SetCancelParams sets CancelParams field to given value.

### HasCancelParams

`func (o *PerformRunActionResponse) HasCancelParams() bool`

HasCancelParams returns a boolean if a field has been set.

### GetPauseParams

`func (o *PerformRunActionResponse) GetPauseParams() []PauseProtectionRunActionResponseParams`

GetPauseParams returns the PauseParams field if non-nil, zero value otherwise.

### GetPauseParamsOk

`func (o *PerformRunActionResponse) GetPauseParamsOk() (*[]PauseProtectionRunActionResponseParams, bool)`

GetPauseParamsOk returns a tuple with the PauseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseParams

`func (o *PerformRunActionResponse) SetPauseParams(v []PauseProtectionRunActionResponseParams)`

SetPauseParams sets PauseParams field to given value.

### HasPauseParams

`func (o *PerformRunActionResponse) HasPauseParams() bool`

HasPauseParams returns a boolean if a field has been set.

### GetResumeParams

`func (o *PerformRunActionResponse) GetResumeParams() []ResumeProtectionRunActionResponseParams`

GetResumeParams returns the ResumeParams field if non-nil, zero value otherwise.

### GetResumeParamsOk

`func (o *PerformRunActionResponse) GetResumeParamsOk() (*[]ResumeProtectionRunActionResponseParams, bool)`

GetResumeParamsOk returns a tuple with the ResumeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeParams

`func (o *PerformRunActionResponse) SetResumeParams(v []ResumeProtectionRunActionResponseParams)`

SetResumeParams sets ResumeParams field to given value.

### HasResumeParams

`func (o *PerformRunActionResponse) HasResumeParams() bool`

HasResumeParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


