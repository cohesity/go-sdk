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

// checks if the TieringType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TieringType{}

// TieringType Data tiering task type.
type TieringType struct {
	// Specifies the data tiering task type.
	Type *string `json:"type,omitempty"`
}

// NewTieringType instantiates a new TieringType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTieringType() *TieringType {
	this := TieringType{}
	return &this
}

// NewTieringTypeWithDefaults instantiates a new TieringType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTieringTypeWithDefaults() *TieringType {
	this := TieringType{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TieringType) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TieringType) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TieringType) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TieringType) SetType(v string) {
	o.Type = &v
}

func (o TieringType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TieringType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableTieringType struct {
	value *TieringType
	isSet bool
}

func (v NullableTieringType) Get() *TieringType {
	return v.value
}

func (v *NullableTieringType) Set(val *TieringType) {
	v.value = val
	v.isSet = true
}

func (v NullableTieringType) IsSet() bool {
	return v.isSet
}

func (v *NullableTieringType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTieringType(val *TieringType) *NullableTieringType {
	return &NullableTieringType{value: val, isSet: true}
}

func (v NullableTieringType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTieringType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


