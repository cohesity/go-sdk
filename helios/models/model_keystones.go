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

// Keystones Specifies a list of Keystones.
type Keystones struct {
	// Specifies a list of Keystones.
	Keystones *[]Keystone `json:"keystones,omitempty"`
}

// NewKeystones instantiates a new Keystones object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeystones() *Keystones {
	this := Keystones{}
	return &this
}

// NewKeystonesWithDefaults instantiates a new Keystones object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeystonesWithDefaults() *Keystones {
	this := Keystones{}
	return &this
}

// GetKeystones returns the Keystones field value if set, zero value otherwise.
func (o *Keystones) GetKeystones() []Keystone {
	if o == nil || o.Keystones == nil {
		var ret []Keystone
		return ret
	}
	return *o.Keystones
}

// GetKeystonesOk returns a tuple with the Keystones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Keystones) GetKeystonesOk() (*[]Keystone, bool) {
	if o == nil || o.Keystones == nil {
		return nil, false
	}
	return o.Keystones, true
}

// HasKeystones returns a boolean if a field has been set.
func (o *Keystones) HasKeystones() bool {
	if o != nil && o.Keystones != nil {
		return true
	}

	return false
}

// SetKeystones gets a reference to the given []Keystone and assigns it to the Keystones field.
func (o *Keystones) SetKeystones(v []Keystone) {
	o.Keystones = &v
}

func (o Keystones) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Keystones != nil {
		toSerialize["keystones"] = o.Keystones
	}
	return json.Marshal(toSerialize)
}

type NullableKeystones struct {
	value *Keystones
	isSet bool
}

func (v NullableKeystones) Get() *Keystones {
	return v.value
}

func (v *NullableKeystones) Set(val *Keystones) {
	v.value = val
	v.isSet = true
}

func (v NullableKeystones) IsSet() bool {
	return v.isSet
}

func (v *NullableKeystones) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeystones(val *Keystones) *NullableKeystones {
	return &NullableKeystones{value: val, isSet: true}
}

func (v NullableKeystones) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeystones) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


