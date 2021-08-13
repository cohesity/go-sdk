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

// Racks Specifies info about list of racks.
type Racks struct {
	// Specifies list of racks
	Racks []Rack `json:"racks,omitempty"`
}

// NewRacks instantiates a new Racks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRacks() *Racks {
	this := Racks{}
	return &this
}

// NewRacksWithDefaults instantiates a new Racks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRacksWithDefaults() *Racks {
	this := Racks{}
	return &this
}

// GetRacks returns the Racks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Racks) GetRacks() []Rack {
	if o == nil  {
		var ret []Rack
		return ret
	}
	return o.Racks
}

// GetRacksOk returns a tuple with the Racks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Racks) GetRacksOk() (*[]Rack, bool) {
	if o == nil || o.Racks == nil {
		return nil, false
	}
	return &o.Racks, true
}

// HasRacks returns a boolean if a field has been set.
func (o *Racks) HasRacks() bool {
	if o != nil && o.Racks != nil {
		return true
	}

	return false
}

// SetRacks gets a reference to the given []Rack and assigns it to the Racks field.
func (o *Racks) SetRacks(v []Rack) {
	o.Racks = v
}

func (o Racks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Racks != nil {
		toSerialize["racks"] = o.Racks
	}
	return json.Marshal(toSerialize)
}

type NullableRacks struct {
	value *Racks
	isSet bool
}

func (v NullableRacks) Get() *Racks {
	return v.value
}

func (v *NullableRacks) Set(val *Racks) {
	v.value = val
	v.isSet = true
}

func (v NullableRacks) IsSet() bool {
	return v.isSet
}

func (v *NullableRacks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRacks(val *Racks) *NullableRacks {
	return &NullableRacks{value: val, isSet: true}
}

func (v NullableRacks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRacks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o Racks) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}