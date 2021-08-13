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

// InfectedFileState Specifies the state of infected file.
type InfectedFileState struct {
	// Specifies the state of infected file.
	Type *string `json:"type,omitempty"`
}

// NewInfectedFileState instantiates a new InfectedFileState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInfectedFileState() *InfectedFileState {
	this := InfectedFileState{}
	return &this
}

// NewInfectedFileStateWithDefaults instantiates a new InfectedFileState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInfectedFileStateWithDefaults() *InfectedFileState {
	this := InfectedFileState{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InfectedFileState) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfectedFileState) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InfectedFileState) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InfectedFileState) SetType(v string) {
	o.Type = &v
}

func (o InfectedFileState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableInfectedFileState struct {
	value *InfectedFileState
	isSet bool
}

func (v NullableInfectedFileState) Get() *InfectedFileState {
	return v.value
}

func (v *NullableInfectedFileState) Set(val *InfectedFileState) {
	v.value = val
	v.isSet = true
}

func (v NullableInfectedFileState) IsSet() bool {
	return v.isSet
}

func (v *NullableInfectedFileState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInfectedFileState(val *InfectedFileState) *NullableInfectedFileState {
	return &NullableInfectedFileState{value: val, isSet: true}
}

func (v NullableInfectedFileState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInfectedFileState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


