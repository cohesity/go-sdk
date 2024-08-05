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

// checks if the RecoverVmwareDiskParamsVmwareTargetParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoverVmwareDiskParamsVmwareTargetParams{}

// RecoverVmwareDiskParamsVmwareTargetParams Specifies the params for recovering to a VMware target.
type RecoverVmwareDiskParamsVmwareTargetParams struct {
	// Specifies whether or not to continue performing the recovery in the event that an error is encountered.
	ContinueOnError NullableBool `json:"continueOnError,omitempty"`
	OriginalSourceConfig *VmwareRecoverDisksOriginalSourceConfig `json:"originalSourceConfig,omitempty"`
	// Specifies whether or not to power off VMs before performing the recovery.
	PowerOffVms NullableBool `json:"powerOffVms,omitempty"`
	// Specifies whether or not to power on VMs after performing the recovery.
	PowerOnVms NullableBool `json:"powerOnVms,omitempty"`
	TargetSourceConfig *VmwareRecoverDisksTargetSourceConfig `json:"targetSourceConfig,omitempty"`
	VlanConfig NullableRecoverKubernetesNamespaceParamsVlanConfig `json:"vlanConfig,omitempty"`
}

// NewRecoverVmwareDiskParamsVmwareTargetParams instantiates a new RecoverVmwareDiskParamsVmwareTargetParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverVmwareDiskParamsVmwareTargetParams() *RecoverVmwareDiskParamsVmwareTargetParams {
	this := RecoverVmwareDiskParamsVmwareTargetParams{}
	return &this
}

// NewRecoverVmwareDiskParamsVmwareTargetParamsWithDefaults instantiates a new RecoverVmwareDiskParamsVmwareTargetParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverVmwareDiskParamsVmwareTargetParamsWithDefaults() *RecoverVmwareDiskParamsVmwareTargetParams {
	this := RecoverVmwareDiskParamsVmwareTargetParams{}
	return &this
}

// GetContinueOnError returns the ContinueOnError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverVmwareDiskParamsVmwareTargetParams) GetContinueOnError() bool {
	if o == nil || IsNil(o.ContinueOnError.Get()) {
		var ret bool
		return ret
	}
	return *o.ContinueOnError.Get()
}

// GetContinueOnErrorOk returns a tuple with the ContinueOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverVmwareDiskParamsVmwareTargetParams) GetContinueOnErrorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContinueOnError.Get(), o.ContinueOnError.IsSet()
}

// HasContinueOnError returns a boolean if a field has been set.
func (o *RecoverVmwareDiskParamsVmwareTargetParams) HasContinueOnError() bool {
	if o != nil && o.ContinueOnError.IsSet() {
		return true
	}

	return false
}

// SetContinueOnError gets a reference to the given NullableBool and assigns it to the ContinueOnError field.
func (o *RecoverVmwareDiskParamsVmwareTargetParams) SetContinueOnError(v bool) {
	o.ContinueOnError.Set(&v)
}
// SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil
func (o *RecoverVmwareDiskParamsVmwareTargetParams) SetContinueOnErrorNil() {
	o.ContinueOnError.Set(nil)
}

// UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
func (o *RecoverVmwareDiskParamsVmwareTargetParams) UnsetContinueOnError() {
	o.ContinueOnError.Unset()
}

// GetOriginalSourceConfig returns the OriginalSourceConfig field value if set, zero value otherwise.
func (o *RecoverVmwareDiskParamsVmwareTargetParams) GetOriginalSourceConfig() VmwareRecoverDisksOriginalSourceConfig {
	if o == nil || IsNil(o.OriginalSourceConfig) {
		var ret VmwareRecoverDisksOriginalSourceConfig
		return ret
	}
	return *o.OriginalSourceConfig
}

// GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoverVmwareDiskParamsVmwareTargetParams) GetOriginalSourceConfigOk() (*VmwareRecoverDisksOriginalSourceConfig, bool) {
	if o == nil || IsNil(o.OriginalSourceConfig) {
		return nil, false
	}
	return o.OriginalSourceConfig, true
}

// HasOriginalSourceConfig returns a boolean if a field has been set.
func (o *RecoverVmwareDiskParamsVmwareTargetParams) HasOriginalSourceConfig() bool {
	if o != nil && !IsNil(o.OriginalSourceConfig) {
		return true
	}

	return false
}

// SetOriginalSourceConfig gets a reference to the given VmwareRecoverDisksOriginalSourceConfig and assigns it to the OriginalSourceConfig field.
func (o *RecoverVmwareDiskParamsVmwareTargetParams) SetOriginalSourceConfig(v VmwareRecoverDisksOriginalSourceConfig) {
	o.OriginalSourceConfig = &v
}

// GetPowerOffVms returns the PowerOffVms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverVmwareDiskParamsVmwareTargetParams) GetPowerOffVms() bool {
	if o == nil || IsNil(o.PowerOffVms.Get()) {
		var ret bool
		return ret
	}
	return *o.PowerOffVms.Get()
}

// GetPowerOffVmsOk returns a tuple with the PowerOffVms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverVmwareDiskParamsVmwareTargetParams) GetPowerOffVmsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PowerOffVms.Get(), o.PowerOffVms.IsSet()
}

// HasPowerOffVms returns a boolean if a field has been set.
func (o *RecoverVmwareDiskParamsVmwareTargetParams) HasPowerOffVms() bool {
	if o != nil && o.PowerOffVms.IsSet() {
		return true
	}

	return false
}

// SetPowerOffVms gets a reference to the given NullableBool and assigns it to the PowerOffVms field.
func (o *RecoverVmwareDiskParamsVmwareTargetParams) SetPowerOffVms(v bool) {
	o.PowerOffVms.Set(&v)
}
// SetPowerOffVmsNil sets the value for PowerOffVms to be an explicit nil
func (o *RecoverVmwareDiskParamsVmwareTargetParams) SetPowerOffVmsNil() {
	o.PowerOffVms.Set(nil)
}

// UnsetPowerOffVms ensures that no value is present for PowerOffVms, not even an explicit nil
func (o *RecoverVmwareDiskParamsVmwareTargetParams) UnsetPowerOffVms() {
	o.PowerOffVms.Unset()
}

// GetPowerOnVms returns the PowerOnVms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverVmwareDiskParamsVmwareTargetParams) GetPowerOnVms() bool {
	if o == nil || IsNil(o.PowerOnVms.Get()) {
		var ret bool
		return ret
	}
	return *o.PowerOnVms.Get()
}

// GetPowerOnVmsOk returns a tuple with the PowerOnVms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverVmwareDiskParamsVmwareTargetParams) GetPowerOnVmsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PowerOnVms.Get(), o.PowerOnVms.IsSet()
}

// HasPowerOnVms returns a boolean if a field has been set.
func (o *RecoverVmwareDiskParamsVmwareTargetParams) HasPowerOnVms() bool {
	if o != nil && o.PowerOnVms.IsSet() {
		return true
	}

	return false
}

// SetPowerOnVms gets a reference to the given NullableBool and assigns it to the PowerOnVms field.
func (o *RecoverVmwareDiskParamsVmwareTargetParams) SetPowerOnVms(v bool) {
	o.PowerOnVms.Set(&v)
}
// SetPowerOnVmsNil sets the value for PowerOnVms to be an explicit nil
func (o *RecoverVmwareDiskParamsVmwareTargetParams) SetPowerOnVmsNil() {
	o.PowerOnVms.Set(nil)
}

// UnsetPowerOnVms ensures that no value is present for PowerOnVms, not even an explicit nil
func (o *RecoverVmwareDiskParamsVmwareTargetParams) UnsetPowerOnVms() {
	o.PowerOnVms.Unset()
}

// GetTargetSourceConfig returns the TargetSourceConfig field value if set, zero value otherwise.
func (o *RecoverVmwareDiskParamsVmwareTargetParams) GetTargetSourceConfig() VmwareRecoverDisksTargetSourceConfig {
	if o == nil || IsNil(o.TargetSourceConfig) {
		var ret VmwareRecoverDisksTargetSourceConfig
		return ret
	}
	return *o.TargetSourceConfig
}

// GetTargetSourceConfigOk returns a tuple with the TargetSourceConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoverVmwareDiskParamsVmwareTargetParams) GetTargetSourceConfigOk() (*VmwareRecoverDisksTargetSourceConfig, bool) {
	if o == nil || IsNil(o.TargetSourceConfig) {
		return nil, false
	}
	return o.TargetSourceConfig, true
}

// HasTargetSourceConfig returns a boolean if a field has been set.
func (o *RecoverVmwareDiskParamsVmwareTargetParams) HasTargetSourceConfig() bool {
	if o != nil && !IsNil(o.TargetSourceConfig) {
		return true
	}

	return false
}

// SetTargetSourceConfig gets a reference to the given VmwareRecoverDisksTargetSourceConfig and assigns it to the TargetSourceConfig field.
func (o *RecoverVmwareDiskParamsVmwareTargetParams) SetTargetSourceConfig(v VmwareRecoverDisksTargetSourceConfig) {
	o.TargetSourceConfig = &v
}

// GetVlanConfig returns the VlanConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverVmwareDiskParamsVmwareTargetParams) GetVlanConfig() RecoverKubernetesNamespaceParamsVlanConfig {
	if o == nil || IsNil(o.VlanConfig.Get()) {
		var ret RecoverKubernetesNamespaceParamsVlanConfig
		return ret
	}
	return *o.VlanConfig.Get()
}

// GetVlanConfigOk returns a tuple with the VlanConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverVmwareDiskParamsVmwareTargetParams) GetVlanConfigOk() (*RecoverKubernetesNamespaceParamsVlanConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.VlanConfig.Get(), o.VlanConfig.IsSet()
}

// HasVlanConfig returns a boolean if a field has been set.
func (o *RecoverVmwareDiskParamsVmwareTargetParams) HasVlanConfig() bool {
	if o != nil && o.VlanConfig.IsSet() {
		return true
	}

	return false
}

// SetVlanConfig gets a reference to the given NullableRecoverKubernetesNamespaceParamsVlanConfig and assigns it to the VlanConfig field.
func (o *RecoverVmwareDiskParamsVmwareTargetParams) SetVlanConfig(v RecoverKubernetesNamespaceParamsVlanConfig) {
	o.VlanConfig.Set(&v)
}
// SetVlanConfigNil sets the value for VlanConfig to be an explicit nil
func (o *RecoverVmwareDiskParamsVmwareTargetParams) SetVlanConfigNil() {
	o.VlanConfig.Set(nil)
}

// UnsetVlanConfig ensures that no value is present for VlanConfig, not even an explicit nil
func (o *RecoverVmwareDiskParamsVmwareTargetParams) UnsetVlanConfig() {
	o.VlanConfig.Unset()
}

func (o RecoverVmwareDiskParamsVmwareTargetParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoverVmwareDiskParamsVmwareTargetParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ContinueOnError.IsSet() {
		toSerialize["continueOnError"] = o.ContinueOnError.Get()
	}
	if !IsNil(o.OriginalSourceConfig) {
		toSerialize["originalSourceConfig"] = o.OriginalSourceConfig
	}
	if o.PowerOffVms.IsSet() {
		toSerialize["powerOffVms"] = o.PowerOffVms.Get()
	}
	if o.PowerOnVms.IsSet() {
		toSerialize["powerOnVms"] = o.PowerOnVms.Get()
	}
	if !IsNil(o.TargetSourceConfig) {
		toSerialize["targetSourceConfig"] = o.TargetSourceConfig
	}
	if o.VlanConfig.IsSet() {
		toSerialize["vlanConfig"] = o.VlanConfig.Get()
	}
	return toSerialize, nil
}

type NullableRecoverVmwareDiskParamsVmwareTargetParams struct {
	value *RecoverVmwareDiskParamsVmwareTargetParams
	isSet bool
}

func (v NullableRecoverVmwareDiskParamsVmwareTargetParams) Get() *RecoverVmwareDiskParamsVmwareTargetParams {
	return v.value
}

func (v *NullableRecoverVmwareDiskParamsVmwareTargetParams) Set(val *RecoverVmwareDiskParamsVmwareTargetParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverVmwareDiskParamsVmwareTargetParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverVmwareDiskParamsVmwareTargetParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverVmwareDiskParamsVmwareTargetParams(val *RecoverVmwareDiskParamsVmwareTargetParams) *NullableRecoverVmwareDiskParamsVmwareTargetParams {
	return &NullableRecoverVmwareDiskParamsVmwareTargetParams{value: val, isSet: true}
}

func (v NullableRecoverVmwareDiskParamsVmwareTargetParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverVmwareDiskParamsVmwareTargetParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


