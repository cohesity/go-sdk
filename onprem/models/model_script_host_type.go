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

// ScriptHostType Script Host Type
type ScriptHostType struct {
	// Specifies the host type of the pre/post script.
	ScriptHostType *string `json:"scriptHostType,omitempty"`
}

// NewScriptHostType instantiates a new ScriptHostType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScriptHostType() *ScriptHostType {
	this := ScriptHostType{}
	return &this
}

// NewScriptHostTypeWithDefaults instantiates a new ScriptHostType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScriptHostTypeWithDefaults() *ScriptHostType {
	this := ScriptHostType{}
	return &this
}

// GetScriptHostType returns the ScriptHostType field value if set, zero value otherwise.
func (o *ScriptHostType) GetScriptHostType() string {
	if o == nil || o.ScriptHostType == nil {
		var ret string
		return ret
	}
	return *o.ScriptHostType
}

// GetScriptHostTypeOk returns a tuple with the ScriptHostType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptHostType) GetScriptHostTypeOk() (*string, bool) {
	if o == nil || o.ScriptHostType == nil {
		return nil, false
	}
	return o.ScriptHostType, true
}

// HasScriptHostType returns a boolean if a field has been set.
func (o *ScriptHostType) HasScriptHostType() bool {
	if o != nil && o.ScriptHostType != nil {
		return true
	}

	return false
}

// SetScriptHostType gets a reference to the given string and assigns it to the ScriptHostType field.
func (o *ScriptHostType) SetScriptHostType(v string) {
	o.ScriptHostType = &v
}

func (o ScriptHostType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ScriptHostType != nil {
		toSerialize["scriptHostType"] = o.ScriptHostType
	}
	return json.Marshal(toSerialize)
}

type NullableScriptHostType struct {
	value *ScriptHostType
	isSet bool
}

func (v NullableScriptHostType) Get() *ScriptHostType {
	return v.value
}

func (v *NullableScriptHostType) Set(val *ScriptHostType) {
	v.value = val
	v.isSet = true
}

func (v NullableScriptHostType) IsSet() bool {
	return v.isSet
}

func (v *NullableScriptHostType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScriptHostType(val *ScriptHostType) *NullableScriptHostType {
	return &NullableScriptHostType{value: val, isSet: true}
}

func (v NullableScriptHostType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScriptHostType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o ScriptHostType) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}