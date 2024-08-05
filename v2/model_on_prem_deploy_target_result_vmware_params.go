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

// checks if the OnPremDeployTargetResultVmwareParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnPremDeployTargetResultVmwareParams{}

// OnPremDeployTargetResultVmwareParams Specifies information about a VMware OnPrem deploy target.
type OnPremDeployTargetResultVmwareParams struct {
	Moref *MOref `json:"moref,omitempty"`
	// UUID of the recovered VMware VM
	Uuid NullableString `json:"uuid,omitempty"`
}

// NewOnPremDeployTargetResultVmwareParams instantiates a new OnPremDeployTargetResultVmwareParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnPremDeployTargetResultVmwareParams() *OnPremDeployTargetResultVmwareParams {
	this := OnPremDeployTargetResultVmwareParams{}
	return &this
}

// NewOnPremDeployTargetResultVmwareParamsWithDefaults instantiates a new OnPremDeployTargetResultVmwareParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnPremDeployTargetResultVmwareParamsWithDefaults() *OnPremDeployTargetResultVmwareParams {
	this := OnPremDeployTargetResultVmwareParams{}
	return &this
}

// GetMoref returns the Moref field value if set, zero value otherwise.
func (o *OnPremDeployTargetResultVmwareParams) GetMoref() MOref {
	if o == nil || IsNil(o.Moref) {
		var ret MOref
		return ret
	}
	return *o.Moref
}

// GetMorefOk returns a tuple with the Moref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnPremDeployTargetResultVmwareParams) GetMorefOk() (*MOref, bool) {
	if o == nil || IsNil(o.Moref) {
		return nil, false
	}
	return o.Moref, true
}

// HasMoref returns a boolean if a field has been set.
func (o *OnPremDeployTargetResultVmwareParams) HasMoref() bool {
	if o != nil && !IsNil(o.Moref) {
		return true
	}

	return false
}

// SetMoref gets a reference to the given MOref and assigns it to the Moref field.
func (o *OnPremDeployTargetResultVmwareParams) SetMoref(v MOref) {
	o.Moref = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnPremDeployTargetResultVmwareParams) GetUuid() string {
	if o == nil || IsNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnPremDeployTargetResultVmwareParams) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *OnPremDeployTargetResultVmwareParams) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *OnPremDeployTargetResultVmwareParams) SetUuid(v string) {
	o.Uuid.Set(&v)
}
// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *OnPremDeployTargetResultVmwareParams) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *OnPremDeployTargetResultVmwareParams) UnsetUuid() {
	o.Uuid.Unset()
}

func (o OnPremDeployTargetResultVmwareParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnPremDeployTargetResultVmwareParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Moref) {
		toSerialize["moref"] = o.Moref
	}
	if o.Uuid.IsSet() {
		toSerialize["uuid"] = o.Uuid.Get()
	}
	return toSerialize, nil
}

type NullableOnPremDeployTargetResultVmwareParams struct {
	value *OnPremDeployTargetResultVmwareParams
	isSet bool
}

func (v NullableOnPremDeployTargetResultVmwareParams) Get() *OnPremDeployTargetResultVmwareParams {
	return v.value
}

func (v *NullableOnPremDeployTargetResultVmwareParams) Set(val *OnPremDeployTargetResultVmwareParams) {
	v.value = val
	v.isSet = true
}

func (v NullableOnPremDeployTargetResultVmwareParams) IsSet() bool {
	return v.isSet
}

func (v *NullableOnPremDeployTargetResultVmwareParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnPremDeployTargetResultVmwareParams(val *OnPremDeployTargetResultVmwareParams) *NullableOnPremDeployTargetResultVmwareParams {
	return &NullableOnPremDeployTargetResultVmwareParams{value: val, isSet: true}
}

func (v NullableOnPremDeployTargetResultVmwareParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnPremDeployTargetResultVmwareParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


