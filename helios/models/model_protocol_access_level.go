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

// ProtocolAccessLevel Protocol access level
type ProtocolAccessLevel struct {
	// Specifies the level of access for any protocol.
	ProtocolAccessLevel *string `json:"protocolAccessLevel,omitempty"`
}

// NewProtocolAccessLevel instantiates a new ProtocolAccessLevel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtocolAccessLevel() *ProtocolAccessLevel {
	this := ProtocolAccessLevel{}
	return &this
}

// NewProtocolAccessLevelWithDefaults instantiates a new ProtocolAccessLevel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtocolAccessLevelWithDefaults() *ProtocolAccessLevel {
	this := ProtocolAccessLevel{}
	return &this
}

// GetProtocolAccessLevel returns the ProtocolAccessLevel field value if set, zero value otherwise.
func (o *ProtocolAccessLevel) GetProtocolAccessLevel() string {
	if o == nil || o.ProtocolAccessLevel == nil {
		var ret string
		return ret
	}
	return *o.ProtocolAccessLevel
}

// GetProtocolAccessLevelOk returns a tuple with the ProtocolAccessLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolAccessLevel) GetProtocolAccessLevelOk() (*string, bool) {
	if o == nil || o.ProtocolAccessLevel == nil {
		return nil, false
	}
	return o.ProtocolAccessLevel, true
}

// HasProtocolAccessLevel returns a boolean if a field has been set.
func (o *ProtocolAccessLevel) HasProtocolAccessLevel() bool {
	if o != nil && o.ProtocolAccessLevel != nil {
		return true
	}

	return false
}

// SetProtocolAccessLevel gets a reference to the given string and assigns it to the ProtocolAccessLevel field.
func (o *ProtocolAccessLevel) SetProtocolAccessLevel(v string) {
	o.ProtocolAccessLevel = &v
}

func (o ProtocolAccessLevel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProtocolAccessLevel != nil {
		toSerialize["protocolAccessLevel"] = o.ProtocolAccessLevel
	}
	return json.Marshal(toSerialize)
}

type NullableProtocolAccessLevel struct {
	value *ProtocolAccessLevel
	isSet bool
}

func (v NullableProtocolAccessLevel) Get() *ProtocolAccessLevel {
	return v.value
}

func (v *NullableProtocolAccessLevel) Set(val *ProtocolAccessLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableProtocolAccessLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableProtocolAccessLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtocolAccessLevel(val *ProtocolAccessLevel) *NullableProtocolAccessLevel {
	return &NullableProtocolAccessLevel{value: val, isSet: true}
}

func (v NullableProtocolAccessLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtocolAccessLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


