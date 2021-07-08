# ValueData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BytesValue** | Pointer to **[]int32** | Specifies the field to store an array of bytes if the current data type is bytes. Specify a value for this field when type is equal to 4. | [optional] 
**DoubleValue** | Pointer to **NullableFloat64** | Specifies the field to store data if the current data type is a double value. Specify a value for this field when type is equal to 2. | [optional] 
**Int64Value** | Pointer to **NullableInt64** | Specifies the field to store data if the current data type is a int64 value. Specify a value for this field when type is equal to 1. | [optional] 
**StringValue** | Pointer to **NullableString** | Specifies the field to store data if the current data type is a string value. Specify a value for this field when type is equal to 3. | [optional] 

## Methods

### NewValueData

`func NewValueData() *ValueData`

NewValueData instantiates a new ValueData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValueDataWithDefaults

`func NewValueDataWithDefaults() *ValueData`

NewValueDataWithDefaults instantiates a new ValueData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytesValue

`func (o *ValueData) GetBytesValue() []int32`

GetBytesValue returns the BytesValue field if non-nil, zero value otherwise.

### GetBytesValueOk

`func (o *ValueData) GetBytesValueOk() (*[]int32, bool)`

GetBytesValueOk returns a tuple with the BytesValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesValue

`func (o *ValueData) SetBytesValue(v []int32)`

SetBytesValue sets BytesValue field to given value.

### HasBytesValue

`func (o *ValueData) HasBytesValue() bool`

HasBytesValue returns a boolean if a field has been set.

### SetBytesValueNil

`func (o *ValueData) SetBytesValueNil(b bool)`

 SetBytesValueNil sets the value for BytesValue to be an explicit nil

### UnsetBytesValue
`func (o *ValueData) UnsetBytesValue()`

UnsetBytesValue ensures that no value is present for BytesValue, not even an explicit nil
### GetDoubleValue

`func (o *ValueData) GetDoubleValue() float64`

GetDoubleValue returns the DoubleValue field if non-nil, zero value otherwise.

### GetDoubleValueOk

`func (o *ValueData) GetDoubleValueOk() (*float64, bool)`

GetDoubleValueOk returns a tuple with the DoubleValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoubleValue

`func (o *ValueData) SetDoubleValue(v float64)`

SetDoubleValue sets DoubleValue field to given value.

### HasDoubleValue

`func (o *ValueData) HasDoubleValue() bool`

HasDoubleValue returns a boolean if a field has been set.

### SetDoubleValueNil

`func (o *ValueData) SetDoubleValueNil(b bool)`

 SetDoubleValueNil sets the value for DoubleValue to be an explicit nil

### UnsetDoubleValue
`func (o *ValueData) UnsetDoubleValue()`

UnsetDoubleValue ensures that no value is present for DoubleValue, not even an explicit nil
### GetInt64Value

`func (o *ValueData) GetInt64Value() int64`

GetInt64Value returns the Int64Value field if non-nil, zero value otherwise.

### GetInt64ValueOk

`func (o *ValueData) GetInt64ValueOk() (*int64, bool)`

GetInt64ValueOk returns a tuple with the Int64Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInt64Value

`func (o *ValueData) SetInt64Value(v int64)`

SetInt64Value sets Int64Value field to given value.

### HasInt64Value

`func (o *ValueData) HasInt64Value() bool`

HasInt64Value returns a boolean if a field has been set.

### SetInt64ValueNil

`func (o *ValueData) SetInt64ValueNil(b bool)`

 SetInt64ValueNil sets the value for Int64Value to be an explicit nil

### UnsetInt64Value
`func (o *ValueData) UnsetInt64Value()`

UnsetInt64Value ensures that no value is present for Int64Value, not even an explicit nil
### GetStringValue

`func (o *ValueData) GetStringValue() string`

GetStringValue returns the StringValue field if non-nil, zero value otherwise.

### GetStringValueOk

`func (o *ValueData) GetStringValueOk() (*string, bool)`

GetStringValueOk returns a tuple with the StringValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringValue

`func (o *ValueData) SetStringValue(v string)`

SetStringValue sets StringValue field to given value.

### HasStringValue

`func (o *ValueData) HasStringValue() bool`

HasStringValue returns a boolean if a field has been set.

### SetStringValueNil

`func (o *ValueData) SetStringValueNil(b bool)`

 SetStringValueNil sets the value for StringValue to be an explicit nil

### UnsetStringValue
`func (o *ValueData) UnsetStringValue()`

UnsetStringValue ensures that no value is present for StringValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


