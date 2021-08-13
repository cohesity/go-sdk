/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package onprem
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/onprem/utils"
	"fmt"
)

var _ = NullableBool{}

// PauseStatusRunLevel Specifies the various statuses assciated with pause and resume operation on protection run.
type PauseStatusRunLevel struct {
	// Specifies the pause status at run level.
	PauseStatusRunLevel *string `json:"pauseStatusRunLevel,omitempty"`
}

// NewPauseStatusRunLevel instantiates a new PauseStatusRunLevel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPauseStatusRunLevel() *PauseStatusRunLevel {
	this := PauseStatusRunLevel{}
	return &this
}

// NewPauseStatusRunLevelWithDefaults instantiates a new PauseStatusRunLevel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPauseStatusRunLevelWithDefaults() *PauseStatusRunLevel {
	this := PauseStatusRunLevel{}
	return &this
}

// GetPauseStatusRunLevel returns the PauseStatusRunLevel field value if set, zero value otherwise.
func (o *PauseStatusRunLevel) GetPauseStatusRunLevel() string {
	if o == nil || o.PauseStatusRunLevel == nil {
		var ret string
		return ret
	}
	return *o.PauseStatusRunLevel
}

// GetPauseStatusRunLevelOk returns a tuple with the PauseStatusRunLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PauseStatusRunLevel) GetPauseStatusRunLevelOk() (*string, bool) {
	if o == nil || o.PauseStatusRunLevel == nil {
		return nil, false
	}
	return o.PauseStatusRunLevel, true
}

// HasPauseStatusRunLevel returns a boolean if a field has been set.
func (o *PauseStatusRunLevel) HasPauseStatusRunLevel() bool {
	if o != nil && o.PauseStatusRunLevel != nil {
		return true
	}

	return false
}

// SetPauseStatusRunLevel gets a reference to the given string and assigns it to the PauseStatusRunLevel field.
func (o *PauseStatusRunLevel) SetPauseStatusRunLevel(v string) {
	o.PauseStatusRunLevel = &v
}

func (o PauseStatusRunLevel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PauseStatusRunLevel != nil {
		toSerialize["pauseStatusRunLevel"] = o.PauseStatusRunLevel
	}
	return json.Marshal(toSerialize)
}

type NullablePauseStatusRunLevel struct {
	value *PauseStatusRunLevel
	isSet bool
}

func (v NullablePauseStatusRunLevel) Get() *PauseStatusRunLevel {
	return v.value
}

func (v *NullablePauseStatusRunLevel) Set(val *PauseStatusRunLevel) {
	v.value = val
	v.isSet = true
}

func (v NullablePauseStatusRunLevel) IsSet() bool {
	return v.isSet
}

func (v *NullablePauseStatusRunLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePauseStatusRunLevel(val *PauseStatusRunLevel) *NullablePauseStatusRunLevel {
	return &NullablePauseStatusRunLevel{value: val, isSet: true}
}

func (v NullablePauseStatusRunLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePauseStatusRunLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o PauseStatusRunLevel) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}