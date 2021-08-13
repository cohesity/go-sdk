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

// FleetTags Specifies the fleet tag parameters.
type FleetTags struct {
	// Specifies key for the fleet tag.
	Key NullableString `json:"key"`
	// Specifies value for the fleet tag.
	Value NullableString `json:"value"`
}

// NewFleetTags instantiates a new FleetTags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFleetTags(key NullableString, value NullableString) *FleetTags {
	this := FleetTags{}
	this.Key = key
	this.Value = value
	return &this
}

// NewFleetTagsWithDefaults instantiates a new FleetTags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFleetTagsWithDefaults() *FleetTags {
	this := FleetTags{}
	return &this
}

// GetKey returns the Key field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FleetTags) GetKey() string {
	if o == nil || o.Key.Get() == nil {
		var ret string
		return ret
	}

	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FleetTags) GetKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// SetKey sets field value
func (o *FleetTags) SetKey(v string) {
	o.Key.Set(&v)
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FleetTags) GetValue() string {
	if o == nil || o.Value.Get() == nil {
		var ret string
		return ret
	}

	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FleetTags) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// SetValue sets field value
func (o *FleetTags) SetValue(v string) {
	o.Value.Set(&v)
}

func (o FleetTags) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key.Get()
	}
	if true {
		toSerialize["value"] = o.Value.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableFleetTags struct {
	value *FleetTags
	isSet bool
}

func (v NullableFleetTags) Get() *FleetTags {
	return v.value
}

func (v *NullableFleetTags) Set(val *FleetTags) {
	v.value = val
	v.isSet = true
}

func (v NullableFleetTags) IsSet() bool {
	return v.isSet
}

func (v *NullableFleetTags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFleetTags(val *FleetTags) *NullableFleetTags {
	return &NullableFleetTags{value: val, isSet: true}
}

func (v NullableFleetTags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFleetTags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


