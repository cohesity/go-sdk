# ExecuteRunbookActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **NullableString** | Specifies the action to be performed. | 
**PowerOffVmParams** | Pointer to [**PowerOffVmParams**](PowerOffVmParams.md) |  | [optional] 

## Methods

### NewExecuteRunbookActionRequest

`func NewExecuteRunbookActionRequest(action NullableString, ) *ExecuteRunbookActionRequest`

NewExecuteRunbookActionRequest instantiates a new ExecuteRunbookActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteRunbookActionRequestWithDefaults

`func NewExecuteRunbookActionRequestWithDefaults() *ExecuteRunbookActionRequest`

NewExecuteRunbookActionRequestWithDefaults instantiates a new ExecuteRunbookActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ExecuteRunbookActionRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ExecuteRunbookActionRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ExecuteRunbookActionRequest) SetAction(v string)`

SetAction sets Action field to given value.


### SetActionNil

`func (o *ExecuteRunbookActionRequest) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *ExecuteRunbookActionRequest) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetPowerOffVmParams

`func (o *ExecuteRunbookActionRequest) GetPowerOffVmParams() PowerOffVmParams`

GetPowerOffVmParams returns the PowerOffVmParams field if non-nil, zero value otherwise.

### GetPowerOffVmParamsOk

`func (o *ExecuteRunbookActionRequest) GetPowerOffVmParamsOk() (*PowerOffVmParams, bool)`

GetPowerOffVmParamsOk returns a tuple with the PowerOffVmParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOffVmParams

`func (o *ExecuteRunbookActionRequest) SetPowerOffVmParams(v PowerOffVmParams)`

SetPowerOffVmParams sets PowerOffVmParams field to given value.

### HasPowerOffVmParams

`func (o *ExecuteRunbookActionRequest) HasPowerOffVmParams() bool`

HasPowerOffVmParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


