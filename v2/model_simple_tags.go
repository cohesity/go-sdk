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

// checks if the SimpleTags type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SimpleTags{}

// SimpleTags Specifies the simple tag parameters.
type SimpleTags struct {
	// Specifies key for the tag.
	Key NullableString `json:"key"`
	// Specifies value for the tag.
	Value NullableString `json:"value"`
}

type _SimpleTags SimpleTags

// NewSimpleTags instantiates a new SimpleTags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimpleTags(key NullableString, value NullableString) *SimpleTags {
	this := SimpleTags{}
	this.Key = key
	this.Value = value
	return &this
}

// NewSimpleTagsWithDefaults instantiates a new SimpleTags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimpleTagsWithDefaults() *SimpleTags {
	this := SimpleTags{}
	return &this
}

// GetKey returns the Key field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SimpleTags) GetKey() string {
	if o == nil || o.Key.Get() == nil {
		var ret string
		return ret
	}

	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimpleTags) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// SetKey sets field value
func (o *SimpleTags) SetKey(v string) {
	o.Key.Set(&v)
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SimpleTags) GetValue() string {
	if o == nil || o.Value.Get() == nil {
		var ret string
		return ret
	}

	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimpleTags) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// SetValue sets field value
func (o *SimpleTags) SetValue(v string) {
	o.Value.Set(&v)
}

func (o SimpleTags) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimpleTags) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key.Get()
	toSerialize["value"] = o.Value.Get()
	return toSerialize, nil
}

func (o *SimpleTags) UnmarshalJSON(data []byte) (err error) {
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

	varSimpleTags := _SimpleTags{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSimpleTags)

	if err != nil {
		return err
	}

	*o = SimpleTags(varSimpleTags)

	return err
}

type NullableSimpleTags struct {
	value *SimpleTags
	isSet bool
}

func (v NullableSimpleTags) Get() *SimpleTags {
	return v.value
}

func (v *NullableSimpleTags) Set(val *SimpleTags) {
	v.value = val
	v.isSet = true
}

func (v NullableSimpleTags) IsSet() bool {
	return v.isSet
}

func (v *NullableSimpleTags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimpleTags(val *SimpleTags) *NullableSimpleTags {
	return &NullableSimpleTags{value: val, isSet: true}
}

func (v NullableSimpleTags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimpleTags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


