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

// TdmEnvironment Supported environments in TDM.
type TdmEnvironment struct {
	// Specifies the TDM environment type.
	Type *string `json:"type,omitempty"`
}

// NewTdmEnvironment instantiates a new TdmEnvironment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTdmEnvironment() *TdmEnvironment {
	this := TdmEnvironment{}
	return &this
}

// NewTdmEnvironmentWithDefaults instantiates a new TdmEnvironment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTdmEnvironmentWithDefaults() *TdmEnvironment {
	this := TdmEnvironment{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TdmEnvironment) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TdmEnvironment) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TdmEnvironment) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TdmEnvironment) SetType(v string) {
	o.Type = &v
}

func (o TdmEnvironment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableTdmEnvironment struct {
	value *TdmEnvironment
	isSet bool
}

func (v NullableTdmEnvironment) Get() *TdmEnvironment {
	return v.value
}

func (v *NullableTdmEnvironment) Set(val *TdmEnvironment) {
	v.value = val
	v.isSet = true
}

func (v NullableTdmEnvironment) IsSet() bool {
	return v.isSet
}

func (v *NullableTdmEnvironment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTdmEnvironment(val *TdmEnvironment) *NullableTdmEnvironment {
	return &NullableTdmEnvironment{value: val, isSet: true}
}

func (v NullableTdmEnvironment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTdmEnvironment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o TdmEnvironment) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}