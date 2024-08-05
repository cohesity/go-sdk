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

// checks if the NisNetgroups type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NisNetgroups{}

// NisNetgroups Response of NIS Netgroups.
type NisNetgroups struct {
	// A list of NIS Netgroups.
	NisNetgroups []NisNetgroup `json:"nisNetgroups,omitempty"`
}

// NewNisNetgroups instantiates a new NisNetgroups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNisNetgroups() *NisNetgroups {
	this := NisNetgroups{}
	return &this
}

// NewNisNetgroupsWithDefaults instantiates a new NisNetgroups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNisNetgroupsWithDefaults() *NisNetgroups {
	this := NisNetgroups{}
	return &this
}

// GetNisNetgroups returns the NisNetgroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NisNetgroups) GetNisNetgroups() []NisNetgroup {
	if o == nil {
		var ret []NisNetgroup
		return ret
	}
	return o.NisNetgroups
}

// GetNisNetgroupsOk returns a tuple with the NisNetgroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NisNetgroups) GetNisNetgroupsOk() ([]NisNetgroup, bool) {
	if o == nil || IsNil(o.NisNetgroups) {
		return nil, false
	}
	return o.NisNetgroups, true
}

// HasNisNetgroups returns a boolean if a field has been set.
func (o *NisNetgroups) HasNisNetgroups() bool {
	if o != nil && !IsNil(o.NisNetgroups) {
		return true
	}

	return false
}

// SetNisNetgroups gets a reference to the given []NisNetgroup and assigns it to the NisNetgroups field.
func (o *NisNetgroups) SetNisNetgroups(v []NisNetgroup) {
	o.NisNetgroups = v
}

func (o NisNetgroups) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NisNetgroups) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.NisNetgroups != nil {
		toSerialize["nisNetgroups"] = o.NisNetgroups
	}
	return toSerialize, nil
}

type NullableNisNetgroups struct {
	value *NisNetgroups
	isSet bool
}

func (v NullableNisNetgroups) Get() *NisNetgroups {
	return v.value
}

func (v *NullableNisNetgroups) Set(val *NisNetgroups) {
	v.value = val
	v.isSet = true
}

func (v NullableNisNetgroups) IsSet() bool {
	return v.isSet
}

func (v *NullableNisNetgroups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNisNetgroups(val *NisNetgroups) *NullableNisNetgroups {
	return &NullableNisNetgroups{value: val, isSet: true}
}

func (v NullableNisNetgroups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNisNetgroups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


