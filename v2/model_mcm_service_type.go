/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

// checks if the McmServiceType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &McmServiceType{}

// McmServiceType Specifies the different helios services.
type McmServiceType struct {
	// Specifies the different helios services.
	Value *string `json:"value,omitempty"`
}

// NewMcmServiceType instantiates a new McmServiceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMcmServiceType() *McmServiceType {
	this := McmServiceType{}
	return &this
}

// NewMcmServiceTypeWithDefaults instantiates a new McmServiceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMcmServiceTypeWithDefaults() *McmServiceType {
	this := McmServiceType{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *McmServiceType) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *McmServiceType) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *McmServiceType) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *McmServiceType) SetValue(v string) {
	o.Value = &v
}

func (o McmServiceType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o McmServiceType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableMcmServiceType struct {
	value *McmServiceType
	isSet bool
}

func (v NullableMcmServiceType) Get() *McmServiceType {
	return v.value
}

func (v *NullableMcmServiceType) Set(val *McmServiceType) {
	v.value = val
	v.isSet = true
}

func (v NullableMcmServiceType) IsSet() bool {
	return v.isSet
}

func (v *NullableMcmServiceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMcmServiceType(val *McmServiceType) *NullableMcmServiceType {
	return &NullableMcmServiceType{value: val, isSet: true}
}

func (v NullableMcmServiceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMcmServiceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


