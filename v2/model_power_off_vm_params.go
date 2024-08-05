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

// checks if the PowerOffVmParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PowerOffVmParams{}

// PowerOffVmParams Specifies the params for powering off the VMs.
type PowerOffVmParams struct {
	// Specifies the Entity IDs of the VMs to be powered off.
	VmIds []float32 `json:"vmIds,omitempty"`
}

// NewPowerOffVmParams instantiates a new PowerOffVmParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPowerOffVmParams() *PowerOffVmParams {
	this := PowerOffVmParams{}
	return &this
}

// NewPowerOffVmParamsWithDefaults instantiates a new PowerOffVmParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPowerOffVmParamsWithDefaults() *PowerOffVmParams {
	this := PowerOffVmParams{}
	return &this
}

// GetVmIds returns the VmIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerOffVmParams) GetVmIds() []float32 {
	if o == nil {
		var ret []float32
		return ret
	}
	return o.VmIds
}

// GetVmIdsOk returns a tuple with the VmIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerOffVmParams) GetVmIdsOk() ([]float32, bool) {
	if o == nil || IsNil(o.VmIds) {
		return nil, false
	}
	return o.VmIds, true
}

// HasVmIds returns a boolean if a field has been set.
func (o *PowerOffVmParams) HasVmIds() bool {
	if o != nil && !IsNil(o.VmIds) {
		return true
	}

	return false
}

// SetVmIds gets a reference to the given []float32 and assigns it to the VmIds field.
func (o *PowerOffVmParams) SetVmIds(v []float32) {
	o.VmIds = v
}

func (o PowerOffVmParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PowerOffVmParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.VmIds != nil {
		toSerialize["vmIds"] = o.VmIds
	}
	return toSerialize, nil
}

type NullablePowerOffVmParams struct {
	value *PowerOffVmParams
	isSet bool
}

func (v NullablePowerOffVmParams) Get() *PowerOffVmParams {
	return v.value
}

func (v *NullablePowerOffVmParams) Set(val *PowerOffVmParams) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerOffVmParams) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerOffVmParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerOffVmParams(val *PowerOffVmParams) *NullablePowerOffVmParams {
	return &NullablePowerOffVmParams{value: val, isSet: true}
}

func (v NullablePowerOffVmParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerOffVmParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


