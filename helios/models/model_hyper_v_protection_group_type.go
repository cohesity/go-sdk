/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package helios
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/helios/utils"
)

var _ = NullableBool{}

// HyperVProtectionGroupType HyperV Protection Group type.
type HyperVProtectionGroupType struct {
	// Specifies HyperV Protection Group type.
	Environment *string `json:"environment,omitempty"`
}

// NewHyperVProtectionGroupType instantiates a new HyperVProtectionGroupType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperVProtectionGroupType() *HyperVProtectionGroupType {
	this := HyperVProtectionGroupType{}
	return &this
}

// NewHyperVProtectionGroupTypeWithDefaults instantiates a new HyperVProtectionGroupType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperVProtectionGroupTypeWithDefaults() *HyperVProtectionGroupType {
	this := HyperVProtectionGroupType{}
	return &this
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *HyperVProtectionGroupType) GetEnvironment() string {
	if o == nil || o.Environment == nil {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperVProtectionGroupType) GetEnvironmentOk() (*string, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *HyperVProtectionGroupType) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *HyperVProtectionGroupType) SetEnvironment(v string) {
	o.Environment = &v
}

func (o HyperVProtectionGroupType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	return json.Marshal(toSerialize)
}

type NullableHyperVProtectionGroupType struct {
	value *HyperVProtectionGroupType
	isSet bool
}

func (v NullableHyperVProtectionGroupType) Get() *HyperVProtectionGroupType {
	return v.value
}

func (v *NullableHyperVProtectionGroupType) Set(val *HyperVProtectionGroupType) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperVProtectionGroupType) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperVProtectionGroupType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperVProtectionGroupType(val *HyperVProtectionGroupType) *NullableHyperVProtectionGroupType {
	return &NullableHyperVProtectionGroupType{value: val, isSet: true}
}

func (v NullableHyperVProtectionGroupType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperVProtectionGroupType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


