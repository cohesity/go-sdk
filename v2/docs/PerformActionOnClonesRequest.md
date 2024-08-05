# PerformActionOnClonesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **NullableString** | Specifies the action to be performed. | 
**CleanupParams** | Pointer to [**CloneActionCleanupParams**](CloneActionCleanupParams.md) |  | [optional] 

## Methods

### NewPerformActionOnClonesRequest

`func NewPerformActionOnClonesRequest(action NullableString, ) *PerformActionOnClonesRequest`

NewPerformActionOnClonesRequest instantiates a new PerformActionOnClonesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerformActionOnClonesRequestWithDefaults

`func NewPerformActionOnClonesRequestWithDefaults() *PerformActionOnClonesRequest`

NewPerformActionOnClonesRequestWithDefaults instantiates a new PerformActionOnClonesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PerformActionOnClonesRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PerformActionOnClonesRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PerformActionOnClonesRequest) SetAction(v string)`

SetAction sets Action field to given value.


### SetActionNil

`func (o *PerformActionOnClonesRequest) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *PerformActionOnClonesRequest) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetCleanupParams

`func (o *PerformActionOnClonesRequest) GetCleanupParams() CloneActionCleanupParams`

GetCleanupParams returns the CleanupParams field if non-nil, zero value otherwise.

### GetCleanupParamsOk

`func (o *PerformActionOnClonesRequest) GetCleanupParamsOk() (*CloneActionCleanupParams, bool)`

GetCleanupParamsOk returns a tuple with the CleanupParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupParams

`func (o *PerformActionOnClonesRequest) SetCleanupParams(v CloneActionCleanupParams)`

SetCleanupParams sets CleanupParams field to given value.

### HasCleanupParams

`func (o *PerformActionOnClonesRequest) HasCleanupParams() bool`

HasCleanupParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


