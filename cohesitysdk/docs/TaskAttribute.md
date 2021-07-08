# TaskAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Specifies the name of this Task Attribute. | [optional] 
**Value** | Pointer to **NullableString** | Specifies the value of this Task Attribute. | [optional] 
**ValueType** | Pointer to **NullableString** | Specifies the type of the value contained here. All values are returned as pointers to strings, but they can be casted to the type indicated here. &#39;kInt64&#39; indicates that the value stored in the Task Attribute is a 64-bit integer. &#39;kDouble&#39; indicates that the value stored in the Task Attribute is a 64 bit floating point number. &#39;kString&#39; indicates that the value stored in the Task Attribute is a string. &#39;kBytes&#39; indicates that the value stored in the Task Attribute is an array of bytes. | [optional] 

## Methods

### NewTaskAttribute

`func NewTaskAttribute() *TaskAttribute`

NewTaskAttribute instantiates a new TaskAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskAttributeWithDefaults

`func NewTaskAttributeWithDefaults() *TaskAttribute`

NewTaskAttributeWithDefaults instantiates a new TaskAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TaskAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaskAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaskAttribute) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaskAttribute) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TaskAttribute) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TaskAttribute) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetValue

`func (o *TaskAttribute) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TaskAttribute) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TaskAttribute) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *TaskAttribute) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *TaskAttribute) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *TaskAttribute) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetValueType

`func (o *TaskAttribute) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *TaskAttribute) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *TaskAttribute) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *TaskAttribute) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### SetValueTypeNil

`func (o *TaskAttribute) SetValueTypeNil(b bool)`

 SetValueTypeNil sets the value for ValueType to be an explicit nil

### UnsetValueType
`func (o *TaskAttribute) UnsetValueType()`

UnsetValueType ensures that no value is present for ValueType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


