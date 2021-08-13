# ObjectsActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **NullableString** | Specifies the action type that need to be performed. | [optional] 
**LinkParams** | Pointer to [**ObjectLinkingParams**](ObjectLinkingParams.md) |  | [optional] 
**UnLinkParams** | Pointer to [**ObjectUnLinkingParams**](ObjectUnLinkingParams.md) |  | [optional] 

## Methods

### NewObjectsActionRequest

`func NewObjectsActionRequest() *ObjectsActionRequest`

NewObjectsActionRequest instantiates a new ObjectsActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectsActionRequestWithDefaults

`func NewObjectsActionRequestWithDefaults() *ObjectsActionRequest`

NewObjectsActionRequestWithDefaults instantiates a new ObjectsActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ObjectsActionRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ObjectsActionRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ObjectsActionRequest) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ObjectsActionRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *ObjectsActionRequest) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *ObjectsActionRequest) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetLinkParams

`func (o *ObjectsActionRequest) GetLinkParams() ObjectLinkingParams`

GetLinkParams returns the LinkParams field if non-nil, zero value otherwise.

### GetLinkParamsOk

`func (o *ObjectsActionRequest) GetLinkParamsOk() (*ObjectLinkingParams, bool)`

GetLinkParamsOk returns a tuple with the LinkParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkParams

`func (o *ObjectsActionRequest) SetLinkParams(v ObjectLinkingParams)`

SetLinkParams sets LinkParams field to given value.

### HasLinkParams

`func (o *ObjectsActionRequest) HasLinkParams() bool`

HasLinkParams returns a boolean if a field has been set.

### GetUnLinkParams

`func (o *ObjectsActionRequest) GetUnLinkParams() ObjectUnLinkingParams`

GetUnLinkParams returns the UnLinkParams field if non-nil, zero value otherwise.

### GetUnLinkParamsOk

`func (o *ObjectsActionRequest) GetUnLinkParamsOk() (*ObjectUnLinkingParams, bool)`

GetUnLinkParamsOk returns a tuple with the UnLinkParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnLinkParams

`func (o *ObjectsActionRequest) SetUnLinkParams(v ObjectUnLinkingParams)`

SetUnLinkParams sets UnLinkParams field to given value.

### HasUnLinkParams

`func (o *ObjectsActionRequest) HasUnLinkParams() bool`

HasUnLinkParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


