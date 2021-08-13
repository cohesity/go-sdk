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

// VmwareTargetParamsForRecoverDisk Specifies the parameters for a VMware recovery target.
type VmwareTargetParamsForRecoverDisk struct {
	OriginalSourceConfig *VmwareRecoverDisksOriginalSourceConfig `json:"originalSourceConfig,omitempty"`
	TargetSourceConfig *VmwareRecoverDisksTargetSourceConfig `json:"targetSourceConfig,omitempty"`
	// Specifies VLAN Params associated with the recovered. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client's (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery.
	VlanConfig NullableRecoveryVlanConfig `json:"vlanConfig,omitempty"`
	// Specifies whether or not to power off VMs before performing the recovery.
	PowerOffVms NullableBool `json:"powerOffVms,omitempty"`
	// Specifies whether or not to power on VMs after performing the recovery.
	PowerOnVms NullableBool `json:"powerOnVms,omitempty"`
	// Specifies whether or not to continue performing the recovery in the event that an error is encountered.
	ContinueOnError NullableBool `json:"continueOnError,omitempty"`
}

// NewVmwareTargetParamsForRecoverDisk instantiates a new VmwareTargetParamsForRecoverDisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmwareTargetParamsForRecoverDisk() *VmwareTargetParamsForRecoverDisk {
	this := VmwareTargetParamsForRecoverDisk{}
	return &this
}

// NewVmwareTargetParamsForRecoverDiskWithDefaults instantiates a new VmwareTargetParamsForRecoverDisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmwareTargetParamsForRecoverDiskWithDefaults() *VmwareTargetParamsForRecoverDisk {
	this := VmwareTargetParamsForRecoverDisk{}
	return &this
}

// GetOriginalSourceConfig returns the OriginalSourceConfig field value if set, zero value otherwise.
func (o *VmwareTargetParamsForRecoverDisk) GetOriginalSourceConfig() VmwareRecoverDisksOriginalSourceConfig {
	if o == nil || o.OriginalSourceConfig == nil {
		var ret VmwareRecoverDisksOriginalSourceConfig
		return ret
	}
	return *o.OriginalSourceConfig
}

// GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmwareTargetParamsForRecoverDisk) GetOriginalSourceConfigOk() (*VmwareRecoverDisksOriginalSourceConfig, bool) {
	if o == nil || o.OriginalSourceConfig == nil {
		return nil, false
	}
	return o.OriginalSourceConfig, true
}

// HasOriginalSourceConfig returns a boolean if a field has been set.
func (o *VmwareTargetParamsForRecoverDisk) HasOriginalSourceConfig() bool {
	if o != nil && o.OriginalSourceConfig != nil {
		return true
	}

	return false
}

// SetOriginalSourceConfig gets a reference to the given VmwareRecoverDisksOriginalSourceConfig and assigns it to the OriginalSourceConfig field.
func (o *VmwareTargetParamsForRecoverDisk) SetOriginalSourceConfig(v VmwareRecoverDisksOriginalSourceConfig) {
	o.OriginalSourceConfig = &v
}

// GetTargetSourceConfig returns the TargetSourceConfig field value if set, zero value otherwise.
func (o *VmwareTargetParamsForRecoverDisk) GetTargetSourceConfig() VmwareRecoverDisksTargetSourceConfig {
	if o == nil || o.TargetSourceConfig == nil {
		var ret VmwareRecoverDisksTargetSourceConfig
		return ret
	}
	return *o.TargetSourceConfig
}

// GetTargetSourceConfigOk returns a tuple with the TargetSourceConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmwareTargetParamsForRecoverDisk) GetTargetSourceConfigOk() (*VmwareRecoverDisksTargetSourceConfig, bool) {
	if o == nil || o.TargetSourceConfig == nil {
		return nil, false
	}
	return o.TargetSourceConfig, true
}

// HasTargetSourceConfig returns a boolean if a field has been set.
func (o *VmwareTargetParamsForRecoverDisk) HasTargetSourceConfig() bool {
	if o != nil && o.TargetSourceConfig != nil {
		return true
	}

	return false
}

// SetTargetSourceConfig gets a reference to the given VmwareRecoverDisksTargetSourceConfig and assigns it to the TargetSourceConfig field.
func (o *VmwareTargetParamsForRecoverDisk) SetTargetSourceConfig(v VmwareRecoverDisksTargetSourceConfig) {
	o.TargetSourceConfig = &v
}

// GetVlanConfig returns the VlanConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmwareTargetParamsForRecoverDisk) GetVlanConfig() RecoveryVlanConfig {
	if o == nil || o.VlanConfig.Get() == nil {
		var ret RecoveryVlanConfig
		return ret
	}
	return *o.VlanConfig.Get()
}

// GetVlanConfigOk returns a tuple with the VlanConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareTargetParamsForRecoverDisk) GetVlanConfigOk() (*RecoveryVlanConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VlanConfig.Get(), o.VlanConfig.IsSet()
}

// HasVlanConfig returns a boolean if a field has been set.
func (o *VmwareTargetParamsForRecoverDisk) HasVlanConfig() bool {
	if o != nil && o.VlanConfig.IsSet() {
		return true
	}

	return false
}

// SetVlanConfig gets a reference to the given NullableRecoveryVlanConfig and assigns it to the VlanConfig field.
func (o *VmwareTargetParamsForRecoverDisk) SetVlanConfig(v RecoveryVlanConfig) {
	o.VlanConfig.Set(&v)
}
// SetVlanConfigNil sets the value for VlanConfig to be an explicit nil
func (o *VmwareTargetParamsForRecoverDisk) SetVlanConfigNil() {
	o.VlanConfig.Set(nil)
}

// UnsetVlanConfig ensures that no value is present for VlanConfig, not even an explicit nil
func (o *VmwareTargetParamsForRecoverDisk) UnsetVlanConfig() {
	o.VlanConfig.Unset()
}

// GetPowerOffVms returns the PowerOffVms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmwareTargetParamsForRecoverDisk) GetPowerOffVms() bool {
	if o == nil || o.PowerOffVms.Get() == nil {
		var ret bool
		return ret
	}
	return *o.PowerOffVms.Get()
}

// GetPowerOffVmsOk returns a tuple with the PowerOffVms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareTargetParamsForRecoverDisk) GetPowerOffVmsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PowerOffVms.Get(), o.PowerOffVms.IsSet()
}

// HasPowerOffVms returns a boolean if a field has been set.
func (o *VmwareTargetParamsForRecoverDisk) HasPowerOffVms() bool {
	if o != nil && o.PowerOffVms.IsSet() {
		return true
	}

	return false
}

// SetPowerOffVms gets a reference to the given NullableBool and assigns it to the PowerOffVms field.
func (o *VmwareTargetParamsForRecoverDisk) SetPowerOffVms(v bool) {
	o.PowerOffVms.Set(&v)
}
// SetPowerOffVmsNil sets the value for PowerOffVms to be an explicit nil
func (o *VmwareTargetParamsForRecoverDisk) SetPowerOffVmsNil() {
	o.PowerOffVms.Set(nil)
}

// UnsetPowerOffVms ensures that no value is present for PowerOffVms, not even an explicit nil
func (o *VmwareTargetParamsForRecoverDisk) UnsetPowerOffVms() {
	o.PowerOffVms.Unset()
}

// GetPowerOnVms returns the PowerOnVms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmwareTargetParamsForRecoverDisk) GetPowerOnVms() bool {
	if o == nil || o.PowerOnVms.Get() == nil {
		var ret bool
		return ret
	}
	return *o.PowerOnVms.Get()
}

// GetPowerOnVmsOk returns a tuple with the PowerOnVms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareTargetParamsForRecoverDisk) GetPowerOnVmsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PowerOnVms.Get(), o.PowerOnVms.IsSet()
}

// HasPowerOnVms returns a boolean if a field has been set.
func (o *VmwareTargetParamsForRecoverDisk) HasPowerOnVms() bool {
	if o != nil && o.PowerOnVms.IsSet() {
		return true
	}

	return false
}

// SetPowerOnVms gets a reference to the given NullableBool and assigns it to the PowerOnVms field.
func (o *VmwareTargetParamsForRecoverDisk) SetPowerOnVms(v bool) {
	o.PowerOnVms.Set(&v)
}
// SetPowerOnVmsNil sets the value for PowerOnVms to be an explicit nil
func (o *VmwareTargetParamsForRecoverDisk) SetPowerOnVmsNil() {
	o.PowerOnVms.Set(nil)
}

// UnsetPowerOnVms ensures that no value is present for PowerOnVms, not even an explicit nil
func (o *VmwareTargetParamsForRecoverDisk) UnsetPowerOnVms() {
	o.PowerOnVms.Unset()
}

// GetContinueOnError returns the ContinueOnError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmwareTargetParamsForRecoverDisk) GetContinueOnError() bool {
	if o == nil || o.ContinueOnError.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ContinueOnError.Get()
}

// GetContinueOnErrorOk returns a tuple with the ContinueOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareTargetParamsForRecoverDisk) GetContinueOnErrorOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContinueOnError.Get(), o.ContinueOnError.IsSet()
}

// HasContinueOnError returns a boolean if a field has been set.
func (o *VmwareTargetParamsForRecoverDisk) HasContinueOnError() bool {
	if o != nil && o.ContinueOnError.IsSet() {
		return true
	}

	return false
}

// SetContinueOnError gets a reference to the given NullableBool and assigns it to the ContinueOnError field.
func (o *VmwareTargetParamsForRecoverDisk) SetContinueOnError(v bool) {
	o.ContinueOnError.Set(&v)
}
// SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil
func (o *VmwareTargetParamsForRecoverDisk) SetContinueOnErrorNil() {
	o.ContinueOnError.Set(nil)
}

// UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
func (o *VmwareTargetParamsForRecoverDisk) UnsetContinueOnError() {
	o.ContinueOnError.Unset()
}

func (o VmwareTargetParamsForRecoverDisk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OriginalSourceConfig != nil {
		toSerialize["originalSourceConfig"] = o.OriginalSourceConfig
	}
	if o.TargetSourceConfig != nil {
		toSerialize["targetSourceConfig"] = o.TargetSourceConfig
	}
	if o.VlanConfig.IsSet() {
		toSerialize["vlanConfig"] = o.VlanConfig.Get()
	}
	if o.PowerOffVms.IsSet() {
		toSerialize["powerOffVms"] = o.PowerOffVms.Get()
	}
	if o.PowerOnVms.IsSet() {
		toSerialize["powerOnVms"] = o.PowerOnVms.Get()
	}
	if o.ContinueOnError.IsSet() {
		toSerialize["continueOnError"] = o.ContinueOnError.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableVmwareTargetParamsForRecoverDisk struct {
	value *VmwareTargetParamsForRecoverDisk
	isSet bool
}

func (v NullableVmwareTargetParamsForRecoverDisk) Get() *VmwareTargetParamsForRecoverDisk {
	return v.value
}

func (v *NullableVmwareTargetParamsForRecoverDisk) Set(val *VmwareTargetParamsForRecoverDisk) {
	v.value = val
	v.isSet = true
}

func (v NullableVmwareTargetParamsForRecoverDisk) IsSet() bool {
	return v.isSet
}

func (v *NullableVmwareTargetParamsForRecoverDisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmwareTargetParamsForRecoverDisk(val *VmwareTargetParamsForRecoverDisk) *NullableVmwareTargetParamsForRecoverDisk {
	return &NullableVmwareTargetParamsForRecoverDisk{value: val, isSet: true}
}

func (v NullableVmwareTargetParamsForRecoverDisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmwareTargetParamsForRecoverDisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


