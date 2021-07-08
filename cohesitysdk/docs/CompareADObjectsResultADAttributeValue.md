# CompareADObjectsResultADAttributeValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValueFlags** | Pointer to **NullableInt32** | Object result flags of type ADAttributeValueFlags. | [optional] 
**ValueVec** | Pointer to **[]string** | String representation of attribute value. For single valued property, only one value will be present here. For multi-valued properties such as group membership, this field will contain values that are in same order as contained in AD. Each AD attribute value will be converted to string. If this property is not set, then the property has null value. | [optional] 

## Methods

### NewCompareADObjectsResultADAttributeValue

`func NewCompareADObjectsResultADAttributeValue() *CompareADObjectsResultADAttributeValue`

NewCompareADObjectsResultADAttributeValue instantiates a new CompareADObjectsResultADAttributeValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompareADObjectsResultADAttributeValueWithDefaults

`func NewCompareADObjectsResultADAttributeValueWithDefaults() *CompareADObjectsResultADAttributeValue`

NewCompareADObjectsResultADAttributeValueWithDefaults instantiates a new CompareADObjectsResultADAttributeValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValueFlags

`func (o *CompareADObjectsResultADAttributeValue) GetValueFlags() int32`

GetValueFlags returns the ValueFlags field if non-nil, zero value otherwise.

### GetValueFlagsOk

`func (o *CompareADObjectsResultADAttributeValue) GetValueFlagsOk() (*int32, bool)`

GetValueFlagsOk returns a tuple with the ValueFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueFlags

`func (o *CompareADObjectsResultADAttributeValue) SetValueFlags(v int32)`

SetValueFlags sets ValueFlags field to given value.

### HasValueFlags

`func (o *CompareADObjectsResultADAttributeValue) HasValueFlags() bool`

HasValueFlags returns a boolean if a field has been set.

### SetValueFlagsNil

`func (o *CompareADObjectsResultADAttributeValue) SetValueFlagsNil(b bool)`

 SetValueFlagsNil sets the value for ValueFlags to be an explicit nil

### UnsetValueFlags
`func (o *CompareADObjectsResultADAttributeValue) UnsetValueFlags()`

UnsetValueFlags ensures that no value is present for ValueFlags, not even an explicit nil
### GetValueVec

`func (o *CompareADObjectsResultADAttributeValue) GetValueVec() []string`

GetValueVec returns the ValueVec field if non-nil, zero value otherwise.

### GetValueVecOk

`func (o *CompareADObjectsResultADAttributeValue) GetValueVecOk() (*[]string, bool)`

GetValueVecOk returns a tuple with the ValueVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueVec

`func (o *CompareADObjectsResultADAttributeValue) SetValueVec(v []string)`

SetValueVec sets ValueVec field to given value.

### HasValueVec

`func (o *CompareADObjectsResultADAttributeValue) HasValueVec() bool`

HasValueVec returns a boolean if a field has been set.

### SetValueVecNil

`func (o *CompareADObjectsResultADAttributeValue) SetValueVecNil(b bool)`

 SetValueVecNil sets the value for ValueVec to be an explicit nil

### UnsetValueVec
`func (o *CompareADObjectsResultADAttributeValue) UnsetValueVec()`

UnsetValueVec ensures that no value is present for ValueVec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


