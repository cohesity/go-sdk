/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the HyperVTargetParamsForMountVolumeNewTargetConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperVTargetParamsForMountVolumeNewTargetConfig{}

// HyperVTargetParamsForMountVolumeNewTargetConfig Specifies the configuration for mounting to a new target.
type HyperVTargetParamsForMountVolumeNewTargetConfig struct {
	// Specifies whether the volumes need to be online within the target environment after attaching the disks. For linux VMs, this should always be set to false because bring disks online is only supported for Windows VM. If this is set to true, HyperV Integration Services must be installed on the VM.
	BringDisksOnline NullableBool `json:"bringDisksOnline"`
	MountTarget RecoverTarget `json:"mountTarget"`
	TargetVmCredentials NullableHyperVMountVolumesNewTargetConfigTargetVmCredentials `json:"targetVmCredentials,omitempty"`
}

type _HyperVTargetParamsForMountVolumeNewTargetConfig HyperVTargetParamsForMountVolumeNewTargetConfig

// NewHyperVTargetParamsForMountVolumeNewTargetConfig instantiates a new HyperVTargetParamsForMountVolumeNewTargetConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperVTargetParamsForMountVolumeNewTargetConfig(bringDisksOnline NullableBool, mountTarget RecoverTarget) *HyperVTargetParamsForMountVolumeNewTargetConfig {
	this := HyperVTargetParamsForMountVolumeNewTargetConfig{}
	this.BringDisksOnline = bringDisksOnline
	this.MountTarget = mountTarget
	return &this
}

// NewHyperVTargetParamsForMountVolumeNewTargetConfigWithDefaults instantiates a new HyperVTargetParamsForMountVolumeNewTargetConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperVTargetParamsForMountVolumeNewTargetConfigWithDefaults() *HyperVTargetParamsForMountVolumeNewTargetConfig {
	this := HyperVTargetParamsForMountVolumeNewTargetConfig{}
	return &this
}

// GetBringDisksOnline returns the BringDisksOnline field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *HyperVTargetParamsForMountVolumeNewTargetConfig) GetBringDisksOnline() bool {
	if o == nil || o.BringDisksOnline.Get() == nil {
		var ret bool
		return ret
	}

	return *o.BringDisksOnline.Get()
}

// GetBringDisksOnlineOk returns a tuple with the BringDisksOnline field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperVTargetParamsForMountVolumeNewTargetConfig) GetBringDisksOnlineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.BringDisksOnline.Get(), o.BringDisksOnline.IsSet()
}

// SetBringDisksOnline sets field value
func (o *HyperVTargetParamsForMountVolumeNewTargetConfig) SetBringDisksOnline(v bool) {
	o.BringDisksOnline.Set(&v)
}

// GetMountTarget returns the MountTarget field value
func (o *HyperVTargetParamsForMountVolumeNewTargetConfig) GetMountTarget() RecoverTarget {
	if o == nil {
		var ret RecoverTarget
		return ret
	}

	return o.MountTarget
}

// GetMountTargetOk returns a tuple with the MountTarget field value
// and a boolean to check if the value has been set.
func (o *HyperVTargetParamsForMountVolumeNewTargetConfig) GetMountTargetOk() (*RecoverTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MountTarget, true
}

// SetMountTarget sets field value
func (o *HyperVTargetParamsForMountVolumeNewTargetConfig) SetMountTarget(v RecoverTarget) {
	o.MountTarget = v
}

// GetTargetVmCredentials returns the TargetVmCredentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperVTargetParamsForMountVolumeNewTargetConfig) GetTargetVmCredentials() HyperVMountVolumesNewTargetConfigTargetVmCredentials {
	if o == nil || IsNil(o.TargetVmCredentials.Get()) {
		var ret HyperVMountVolumesNewTargetConfigTargetVmCredentials
		return ret
	}
	return *o.TargetVmCredentials.Get()
}

// GetTargetVmCredentialsOk returns a tuple with the TargetVmCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperVTargetParamsForMountVolumeNewTargetConfig) GetTargetVmCredentialsOk() (*HyperVMountVolumesNewTargetConfigTargetVmCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetVmCredentials.Get(), o.TargetVmCredentials.IsSet()
}

// HasTargetVmCredentials returns a boolean if a field has been set.
func (o *HyperVTargetParamsForMountVolumeNewTargetConfig) HasTargetVmCredentials() bool {
	if o != nil && o.TargetVmCredentials.IsSet() {
		return true
	}

	return false
}

// SetTargetVmCredentials gets a reference to the given NullableHyperVMountVolumesNewTargetConfigTargetVmCredentials and assigns it to the TargetVmCredentials field.
func (o *HyperVTargetParamsForMountVolumeNewTargetConfig) SetTargetVmCredentials(v HyperVMountVolumesNewTargetConfigTargetVmCredentials) {
	o.TargetVmCredentials.Set(&v)
}
// SetTargetVmCredentialsNil sets the value for TargetVmCredentials to be an explicit nil
func (o *HyperVTargetParamsForMountVolumeNewTargetConfig) SetTargetVmCredentialsNil() {
	o.TargetVmCredentials.Set(nil)
}

// UnsetTargetVmCredentials ensures that no value is present for TargetVmCredentials, not even an explicit nil
func (o *HyperVTargetParamsForMountVolumeNewTargetConfig) UnsetTargetVmCredentials() {
	o.TargetVmCredentials.Unset()
}

func (o HyperVTargetParamsForMountVolumeNewTargetConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperVTargetParamsForMountVolumeNewTargetConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bringDisksOnline"] = o.BringDisksOnline.Get()
	toSerialize["mountTarget"] = o.MountTarget
	if o.TargetVmCredentials.IsSet() {
		toSerialize["targetVmCredentials"] = o.TargetVmCredentials.Get()
	}
	return toSerialize, nil
}

func (o *HyperVTargetParamsForMountVolumeNewTargetConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bringDisksOnline",
		"mountTarget",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varHyperVTargetParamsForMountVolumeNewTargetConfig := _HyperVTargetParamsForMountVolumeNewTargetConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHyperVTargetParamsForMountVolumeNewTargetConfig)

	if err != nil {
		return err
	}

	*o = HyperVTargetParamsForMountVolumeNewTargetConfig(varHyperVTargetParamsForMountVolumeNewTargetConfig)

	return err
}

type NullableHyperVTargetParamsForMountVolumeNewTargetConfig struct {
	value *HyperVTargetParamsForMountVolumeNewTargetConfig
	isSet bool
}

func (v NullableHyperVTargetParamsForMountVolumeNewTargetConfig) Get() *HyperVTargetParamsForMountVolumeNewTargetConfig {
	return v.value
}

func (v *NullableHyperVTargetParamsForMountVolumeNewTargetConfig) Set(val *HyperVTargetParamsForMountVolumeNewTargetConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperVTargetParamsForMountVolumeNewTargetConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperVTargetParamsForMountVolumeNewTargetConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperVTargetParamsForMountVolumeNewTargetConfig(val *HyperVTargetParamsForMountVolumeNewTargetConfig) *NullableHyperVTargetParamsForMountVolumeNewTargetConfig {
	return &NullableHyperVTargetParamsForMountVolumeNewTargetConfig{value: val, isSet: true}
}

func (v NullableHyperVTargetParamsForMountVolumeNewTargetConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperVTargetParamsForMountVolumeNewTargetConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


