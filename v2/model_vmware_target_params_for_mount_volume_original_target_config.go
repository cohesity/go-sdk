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

// checks if the VmwareTargetParamsForMountVolumeOriginalTargetConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VmwareTargetParamsForMountVolumeOriginalTargetConfig{}

// VmwareTargetParamsForMountVolumeOriginalTargetConfig Specifies the configuration for mounting to the original target.
type VmwareTargetParamsForMountVolumeOriginalTargetConfig struct {
	// Specifies whether the volumes need to be online within the target environment after attaching the disks. For linux VMs, this should always be set to true. For Windows, this is optional. If this is set to true, VMware tools must be installed on the VM. If this is set to false, useExistingAgent and targetCredentials are not needed.
	BringDisksOnline NullableBool `json:"bringDisksOnline"`
	TargetVmCredentials NullableVMwareMountVolumesOriginalTargetConfigTargetVmCredentials `json:"targetVmCredentials,omitempty"`
	// Specifies whether this will use an existing agent on the target vm or will deploy a new agent. This is required if bringDisksOnline is set to true.
	UseExistingAgent NullableBool `json:"useExistingAgent,omitempty"`
}

type _VmwareTargetParamsForMountVolumeOriginalTargetConfig VmwareTargetParamsForMountVolumeOriginalTargetConfig

// NewVmwareTargetParamsForMountVolumeOriginalTargetConfig instantiates a new VmwareTargetParamsForMountVolumeOriginalTargetConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmwareTargetParamsForMountVolumeOriginalTargetConfig(bringDisksOnline NullableBool) *VmwareTargetParamsForMountVolumeOriginalTargetConfig {
	this := VmwareTargetParamsForMountVolumeOriginalTargetConfig{}
	this.BringDisksOnline = bringDisksOnline
	return &this
}

// NewVmwareTargetParamsForMountVolumeOriginalTargetConfigWithDefaults instantiates a new VmwareTargetParamsForMountVolumeOriginalTargetConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmwareTargetParamsForMountVolumeOriginalTargetConfigWithDefaults() *VmwareTargetParamsForMountVolumeOriginalTargetConfig {
	this := VmwareTargetParamsForMountVolumeOriginalTargetConfig{}
	return &this
}

// GetBringDisksOnline returns the BringDisksOnline field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *VmwareTargetParamsForMountVolumeOriginalTargetConfig) GetBringDisksOnline() bool {
	if o == nil || o.BringDisksOnline.Get() == nil {
		var ret bool
		return ret
	}

	return *o.BringDisksOnline.Get()
}

// GetBringDisksOnlineOk returns a tuple with the BringDisksOnline field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareTargetParamsForMountVolumeOriginalTargetConfig) GetBringDisksOnlineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.BringDisksOnline.Get(), o.BringDisksOnline.IsSet()
}

// SetBringDisksOnline sets field value
func (o *VmwareTargetParamsForMountVolumeOriginalTargetConfig) SetBringDisksOnline(v bool) {
	o.BringDisksOnline.Set(&v)
}

// GetTargetVmCredentials returns the TargetVmCredentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmwareTargetParamsForMountVolumeOriginalTargetConfig) GetTargetVmCredentials() VMwareMountVolumesOriginalTargetConfigTargetVmCredentials {
	if o == nil || IsNil(o.TargetVmCredentials.Get()) {
		var ret VMwareMountVolumesOriginalTargetConfigTargetVmCredentials
		return ret
	}
	return *o.TargetVmCredentials.Get()
}

// GetTargetVmCredentialsOk returns a tuple with the TargetVmCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareTargetParamsForMountVolumeOriginalTargetConfig) GetTargetVmCredentialsOk() (*VMwareMountVolumesOriginalTargetConfigTargetVmCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetVmCredentials.Get(), o.TargetVmCredentials.IsSet()
}

// HasTargetVmCredentials returns a boolean if a field has been set.
func (o *VmwareTargetParamsForMountVolumeOriginalTargetConfig) HasTargetVmCredentials() bool {
	if o != nil && o.TargetVmCredentials.IsSet() {
		return true
	}

	return false
}

// SetTargetVmCredentials gets a reference to the given NullableVMwareMountVolumesOriginalTargetConfigTargetVmCredentials and assigns it to the TargetVmCredentials field.
func (o *VmwareTargetParamsForMountVolumeOriginalTargetConfig) SetTargetVmCredentials(v VMwareMountVolumesOriginalTargetConfigTargetVmCredentials) {
	o.TargetVmCredentials.Set(&v)
}
// SetTargetVmCredentialsNil sets the value for TargetVmCredentials to be an explicit nil
func (o *VmwareTargetParamsForMountVolumeOriginalTargetConfig) SetTargetVmCredentialsNil() {
	o.TargetVmCredentials.Set(nil)
}

// UnsetTargetVmCredentials ensures that no value is present for TargetVmCredentials, not even an explicit nil
func (o *VmwareTargetParamsForMountVolumeOriginalTargetConfig) UnsetTargetVmCredentials() {
	o.TargetVmCredentials.Unset()
}

// GetUseExistingAgent returns the UseExistingAgent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmwareTargetParamsForMountVolumeOriginalTargetConfig) GetUseExistingAgent() bool {
	if o == nil || IsNil(o.UseExistingAgent.Get()) {
		var ret bool
		return ret
	}
	return *o.UseExistingAgent.Get()
}

// GetUseExistingAgentOk returns a tuple with the UseExistingAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareTargetParamsForMountVolumeOriginalTargetConfig) GetUseExistingAgentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.UseExistingAgent.Get(), o.UseExistingAgent.IsSet()
}

// HasUseExistingAgent returns a boolean if a field has been set.
func (o *VmwareTargetParamsForMountVolumeOriginalTargetConfig) HasUseExistingAgent() bool {
	if o != nil && o.UseExistingAgent.IsSet() {
		return true
	}

	return false
}

// SetUseExistingAgent gets a reference to the given NullableBool and assigns it to the UseExistingAgent field.
func (o *VmwareTargetParamsForMountVolumeOriginalTargetConfig) SetUseExistingAgent(v bool) {
	o.UseExistingAgent.Set(&v)
}
// SetUseExistingAgentNil sets the value for UseExistingAgent to be an explicit nil
func (o *VmwareTargetParamsForMountVolumeOriginalTargetConfig) SetUseExistingAgentNil() {
	o.UseExistingAgent.Set(nil)
}

// UnsetUseExistingAgent ensures that no value is present for UseExistingAgent, not even an explicit nil
func (o *VmwareTargetParamsForMountVolumeOriginalTargetConfig) UnsetUseExistingAgent() {
	o.UseExistingAgent.Unset()
}

func (o VmwareTargetParamsForMountVolumeOriginalTargetConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VmwareTargetParamsForMountVolumeOriginalTargetConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bringDisksOnline"] = o.BringDisksOnline.Get()
	if o.TargetVmCredentials.IsSet() {
		toSerialize["targetVmCredentials"] = o.TargetVmCredentials.Get()
	}
	if o.UseExistingAgent.IsSet() {
		toSerialize["useExistingAgent"] = o.UseExistingAgent.Get()
	}
	return toSerialize, nil
}

func (o *VmwareTargetParamsForMountVolumeOriginalTargetConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bringDisksOnline",
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

	varVmwareTargetParamsForMountVolumeOriginalTargetConfig := _VmwareTargetParamsForMountVolumeOriginalTargetConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVmwareTargetParamsForMountVolumeOriginalTargetConfig)

	if err != nil {
		return err
	}

	*o = VmwareTargetParamsForMountVolumeOriginalTargetConfig(varVmwareTargetParamsForMountVolumeOriginalTargetConfig)

	return err
}

type NullableVmwareTargetParamsForMountVolumeOriginalTargetConfig struct {
	value *VmwareTargetParamsForMountVolumeOriginalTargetConfig
	isSet bool
}

func (v NullableVmwareTargetParamsForMountVolumeOriginalTargetConfig) Get() *VmwareTargetParamsForMountVolumeOriginalTargetConfig {
	return v.value
}

func (v *NullableVmwareTargetParamsForMountVolumeOriginalTargetConfig) Set(val *VmwareTargetParamsForMountVolumeOriginalTargetConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableVmwareTargetParamsForMountVolumeOriginalTargetConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableVmwareTargetParamsForMountVolumeOriginalTargetConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmwareTargetParamsForMountVolumeOriginalTargetConfig(val *VmwareTargetParamsForMountVolumeOriginalTargetConfig) *NullableVmwareTargetParamsForMountVolumeOriginalTargetConfig {
	return &NullableVmwareTargetParamsForMountVolumeOriginalTargetConfig{value: val, isSet: true}
}

func (v NullableVmwareTargetParamsForMountVolumeOriginalTargetConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmwareTargetParamsForMountVolumeOriginalTargetConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


