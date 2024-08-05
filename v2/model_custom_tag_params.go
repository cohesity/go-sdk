/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CustomTagParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomTagParams{}

// CustomTagParams Specifies tag information of custom tags to be applied to various resources when converting and deploying a VM to AWS.
type CustomTagParams struct {
	// Specifies key of the custom tag.
	Key NullableString `json:"key"`
	// Specifies value of the custom tag.
	Value NullableString `json:"value"`
}

type _CustomTagParams CustomTagParams

// NewCustomTagParams instantiates a new CustomTagParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomTagParams(key NullableString, value NullableString) *CustomTagParams {
	this := CustomTagParams{}
	this.Key = key
	this.Value = value
	return &this
}

// NewCustomTagParamsWithDefaults instantiates a new CustomTagParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomTagParamsWithDefaults() *CustomTagParams {
	this := CustomTagParams{}
	return &this
}

// GetKey returns the Key field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CustomTagParams) GetKey() string {
	if o == nil || o.Key.Get() == nil {
		var ret string
		return ret
	}

	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomTagParams) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// SetKey sets field value
func (o *CustomTagParams) SetKey(v string) {
	o.Key.Set(&v)
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CustomTagParams) GetValue() string {
	if o == nil || o.Value.Get() == nil {
		var ret string
		return ret
	}

	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomTagParams) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// SetValue sets field value
func (o *CustomTagParams) SetValue(v string) {
	o.Value.Set(&v)
}

func (o CustomTagParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomTagParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key.Get()
	toSerialize["value"] = o.Value.Get()
	return toSerialize, nil
}

func (o *CustomTagParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"value",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCustomTagParams := _CustomTagParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomTagParams)

	if err != nil {
		return err
	}

	*o = CustomTagParams(varCustomTagParams)

	return err
}

type NullableCustomTagParams struct {
	value *CustomTagParams
	isSet bool
}

func (v NullableCustomTagParams) Get() *CustomTagParams {
	return v.value
}

func (v *NullableCustomTagParams) Set(val *CustomTagParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomTagParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomTagParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomTagParams(val *CustomTagParams) *NullableCustomTagParams {
	return &NullableCustomTagParams{value: val, isSet: true}
}

func (v NullableCustomTagParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomTagParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


