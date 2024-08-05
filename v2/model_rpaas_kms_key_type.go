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

// checks if the RpaasKmsKeyType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RpaasKmsKeyType{}

// RpaasKmsKeyType Specifies the kms key type.
type RpaasKmsKeyType struct {
	// Specifies the kms key type.
	RpaasKmsKeyType *string `json:"rpaasKmsKeyType,omitempty"`
}

// NewRpaasKmsKeyType instantiates a new RpaasKmsKeyType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRpaasKmsKeyType() *RpaasKmsKeyType {
	this := RpaasKmsKeyType{}
	return &this
}

// NewRpaasKmsKeyTypeWithDefaults instantiates a new RpaasKmsKeyType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRpaasKmsKeyTypeWithDefaults() *RpaasKmsKeyType {
	this := RpaasKmsKeyType{}
	return &this
}

// GetRpaasKmsKeyType returns the RpaasKmsKeyType field value if set, zero value otherwise.
func (o *RpaasKmsKeyType) GetRpaasKmsKeyType() string {
	if o == nil || IsNil(o.RpaasKmsKeyType) {
		var ret string
		return ret
	}
	return *o.RpaasKmsKeyType
}

// GetRpaasKmsKeyTypeOk returns a tuple with the RpaasKmsKeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpaasKmsKeyType) GetRpaasKmsKeyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RpaasKmsKeyType) {
		return nil, false
	}
	return o.RpaasKmsKeyType, true
}

// HasRpaasKmsKeyType returns a boolean if a field has been set.
func (o *RpaasKmsKeyType) HasRpaasKmsKeyType() bool {
	if o != nil && !IsNil(o.RpaasKmsKeyType) {
		return true
	}

	return false
}

// SetRpaasKmsKeyType gets a reference to the given string and assigns it to the RpaasKmsKeyType field.
func (o *RpaasKmsKeyType) SetRpaasKmsKeyType(v string) {
	o.RpaasKmsKeyType = &v
}

func (o RpaasKmsKeyType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RpaasKmsKeyType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RpaasKmsKeyType) {
		toSerialize["rpaasKmsKeyType"] = o.RpaasKmsKeyType
	}
	return toSerialize, nil
}

type NullableRpaasKmsKeyType struct {
	value *RpaasKmsKeyType
	isSet bool
}

func (v NullableRpaasKmsKeyType) Get() *RpaasKmsKeyType {
	return v.value
}

func (v *NullableRpaasKmsKeyType) Set(val *RpaasKmsKeyType) {
	v.value = val
	v.isSet = true
}

func (v NullableRpaasKmsKeyType) IsSet() bool {
	return v.isSet
}

func (v *NullableRpaasKmsKeyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRpaasKmsKeyType(val *RpaasKmsKeyType) *NullableRpaasKmsKeyType {
	return &NullableRpaasKmsKeyType{value: val, isSet: true}
}

func (v NullableRpaasKmsKeyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRpaasKmsKeyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


