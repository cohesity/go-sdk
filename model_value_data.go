/*
 * Cohesity REST API
 *
 * This API list provides operations for interfacing with the Cohesity Cluster.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cohesitysdk

import (
	"encoding/json"
)

// ValueData Specifies the fields to store data of a given type. Specify data in the appropriate field for the current data type.
type ValueData struct {
	// Specifies the field to store an array of bytes if the current data type is bytes. Specify a value for this field when type is equal to 4.
	BytesValue []int32 `json:"bytesValue,omitempty"`
	// Specifies the field to store data if the current data type is a double value. Specify a value for this field when type is equal to 2.
	DoubleValue NullableFloat64 `json:"doubleValue,omitempty"`
	// Specifies the field to store data if the current data type is a int64 value. Specify a value for this field when type is equal to 1.
	Int64Value NullableInt64 `json:"int64Value,omitempty"`
	// Specifies the field to store data if the current data type is a string value. Specify a value for this field when type is equal to 3.
	StringValue NullableString `json:"stringValue,omitempty"`
}

// NewValueData instantiates a new ValueData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValueData() *ValueData {
	this := ValueData{}
	return &this
}

// NewValueDataWithDefaults instantiates a new ValueData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValueDataWithDefaults() *ValueData {
	this := ValueData{}
	return &this
}

// GetBytesValue returns the BytesValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ValueData) GetBytesValue() []int32 {
	if o == nil  {
		var ret []int32
		return ret
	}
	return o.BytesValue
}

// GetBytesValueOk returns a tuple with the BytesValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ValueData) GetBytesValueOk() (*[]int32, bool) {
	if o == nil || o.BytesValue == nil {
		return nil, false
	}
	return &o.BytesValue, true
}

// HasBytesValue returns a boolean if a field has been set.
func (o *ValueData) HasBytesValue() bool {
	if o != nil && o.BytesValue != nil {
		return true
	}

	return false
}

// SetBytesValue gets a reference to the given []int32 and assigns it to the BytesValue field.
func (o *ValueData) SetBytesValue(v []int32) {
	o.BytesValue = v
}

// GetDoubleValue returns the DoubleValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ValueData) GetDoubleValue() float64 {
	if o == nil || o.DoubleValue.Get() == nil {
		var ret float64
		return ret
	}
	return *o.DoubleValue.Get()
}

// GetDoubleValueOk returns a tuple with the DoubleValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ValueData) GetDoubleValueOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DoubleValue.Get(), o.DoubleValue.IsSet()
}

// HasDoubleValue returns a boolean if a field has been set.
func (o *ValueData) HasDoubleValue() bool {
	if o != nil && o.DoubleValue.IsSet() {
		return true
	}

	return false
}

// SetDoubleValue gets a reference to the given NullableFloat64 and assigns it to the DoubleValue field.
func (o *ValueData) SetDoubleValue(v float64) {
	o.DoubleValue.Set(&v)
}
// SetDoubleValueNil sets the value for DoubleValue to be an explicit nil
func (o *ValueData) SetDoubleValueNil() {
	o.DoubleValue.Set(nil)
}

// UnsetDoubleValue ensures that no value is present for DoubleValue, not even an explicit nil
func (o *ValueData) UnsetDoubleValue() {
	o.DoubleValue.Unset()
}

// GetInt64Value returns the Int64Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ValueData) GetInt64Value() int64 {
	if o == nil || o.Int64Value.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Int64Value.Get()
}

// GetInt64ValueOk returns a tuple with the Int64Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ValueData) GetInt64ValueOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Int64Value.Get(), o.Int64Value.IsSet()
}

// HasInt64Value returns a boolean if a field has been set.
func (o *ValueData) HasInt64Value() bool {
	if o != nil && o.Int64Value.IsSet() {
		return true
	}

	return false
}

// SetInt64Value gets a reference to the given NullableInt64 and assigns it to the Int64Value field.
func (o *ValueData) SetInt64Value(v int64) {
	o.Int64Value.Set(&v)
}
// SetInt64ValueNil sets the value for Int64Value to be an explicit nil
func (o *ValueData) SetInt64ValueNil() {
	o.Int64Value.Set(nil)
}

// UnsetInt64Value ensures that no value is present for Int64Value, not even an explicit nil
func (o *ValueData) UnsetInt64Value() {
	o.Int64Value.Unset()
}

// GetStringValue returns the StringValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ValueData) GetStringValue() string {
	if o == nil || o.StringValue.Get() == nil {
		var ret string
		return ret
	}
	return *o.StringValue.Get()
}

// GetStringValueOk returns a tuple with the StringValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ValueData) GetStringValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StringValue.Get(), o.StringValue.IsSet()
}

// HasStringValue returns a boolean if a field has been set.
func (o *ValueData) HasStringValue() bool {
	if o != nil && o.StringValue.IsSet() {
		return true
	}

	return false
}

// SetStringValue gets a reference to the given NullableString and assigns it to the StringValue field.
func (o *ValueData) SetStringValue(v string) {
	o.StringValue.Set(&v)
}
// SetStringValueNil sets the value for StringValue to be an explicit nil
func (o *ValueData) SetStringValueNil() {
	o.StringValue.Set(nil)
}

// UnsetStringValue ensures that no value is present for StringValue, not even an explicit nil
func (o *ValueData) UnsetStringValue() {
	o.StringValue.Unset()
}

func (o ValueData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BytesValue != nil {
		toSerialize["bytesValue"] = o.BytesValue
	}
	if o.DoubleValue.IsSet() {
		toSerialize["doubleValue"] = o.DoubleValue.Get()
	}
	if o.Int64Value.IsSet() {
		toSerialize["int64Value"] = o.Int64Value.Get()
	}
	if o.StringValue.IsSet() {
		toSerialize["stringValue"] = o.StringValue.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableValueData struct {
	value *ValueData
	isSet bool
}

func (v NullableValueData) Get() *ValueData {
	return v.value
}

func (v *NullableValueData) Set(val *ValueData) {
	v.value = val
	v.isSet = true
}

func (v NullableValueData) IsSet() bool {
	return v.isSet
}

func (v *NullableValueData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValueData(val *ValueData) *NullableValueData {
	return &NullableValueData{value: val, isSet: true}
}

func (v NullableValueData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValueData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


