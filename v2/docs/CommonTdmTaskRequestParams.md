# CommonTdmTaskRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **NullableString** | Specifies the TDM Task action. | 
**Name** | Pointer to **NullableString** | Specifies the name of the task. | [optional] 

## Methods

### NewCommonTdmTaskRequestParams

`func NewCommonTdmTaskRequestParams(action NullableString, ) *CommonTdmTaskRequestParams`

NewCommonTdmTaskRequestParams instantiates a new CommonTdmTaskRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonTdmTaskRequestParamsWithDefaults

`func NewCommonTdmTaskRequestParamsWithDefaults() *CommonTdmTaskRequestParams`

NewCommonTdmTaskRequestParamsWithDefaults instantiates a new CommonTdmTaskRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *CommonTdmTaskRequestParams) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CommonTdmTaskRequestParams) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CommonTdmTaskRequestParams) SetAction(v string)`

SetAction sets Action field to given value.


### SetActionNil

`func (o *CommonTdmTaskRequestParams) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *CommonTdmTaskRequestParams) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetName

`func (o *CommonTdmTaskRequestParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommonTdmTaskRequestParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommonTdmTaskRequestParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CommonTdmTaskRequestParams) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CommonTdmTaskRequestParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CommonTdmTaskRequestParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


