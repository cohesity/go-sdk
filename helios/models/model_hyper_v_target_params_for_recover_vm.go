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

// HyperVTargetParamsForRecoverVm Specifies the parameters for a HyperV recovery target.
type HyperVTargetParamsForRecoverVm struct {
	// Specifies params to rename the VMs that are recovered. If not specified, the original names of the VMs are preserved.
	RenameRecoveredVmsParams NullableRecoveredOrClonedVmsRenameConfig `json:"renameRecoveredVmsParams,omitempty"`
	// Specifies the recovery target configuration if recovery has to be done to a different location which is different from original source or to original Source with different configuration. If not specified, then the recovery of the vms will be performed to original location with all configuration parameters retained.
	RecoveryTargetConfig NullableHyperVVmRecoveryTargetConfig `json:"recoveryTargetConfig,omitempty"`
	// Specifies whether to power on vms after recovery. If not specified, or false, recovered vms will be in powered off state.
	PowerOnVms NullableBool `json:"powerOnVms,omitempty"`
	// Specifies whether to continue recovering other vms if one of vms failed to recover. Default value is false.
	ContinueOnError NullableBool `json:"continueOnError,omitempty"`
	// Specifies whether to perform an instant recovery. By instant recovery, the recovered VM is available before files are completely copied to the recovered VM. Default is true.
	InstantRecovery NullableBool `json:"instantRecovery,omitempty"`
	// Specifies if the HyperV recovery is using the SMB service to perform the restore. On-prem, this is the case by default. However, as of today, DMaaS does not support SMB, and HyperV VM VM restores will employ an alternative restore method in this case.
	UseSmbService NullableBool `json:"useSmbService,omitempty"`
	// Specifies whether to preserve uuids of recovered VMs. Default is false.
	PreserveUuids NullableBool `json:"preserveUuids,omitempty"`
	// Specifies VLAN Params associated with the recovered. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client's (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery.
	VlanConfig NullableRecoveryVlanConfig `json:"vlanConfig,omitempty"`
}

// NewHyperVTargetParamsForRecoverVm instantiates a new HyperVTargetParamsForRecoverVm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperVTargetParamsForRecoverVm() *HyperVTargetParamsForRecoverVm {
	this := HyperVTargetParamsForRecoverVm{}
	return &this
}

// NewHyperVTargetParamsForRecoverVmWithDefaults instantiates a new HyperVTargetParamsForRecoverVm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperVTargetParamsForRecoverVmWithDefaults() *HyperVTargetParamsForRecoverVm {
	this := HyperVTargetParamsForRecoverVm{}
	return &this
}

// GetRenameRecoveredVmsParams returns the RenameRecoveredVmsParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperVTargetParamsForRecoverVm) GetRenameRecoveredVmsParams() RecoveredOrClonedVmsRenameConfig {
	if o == nil || o.RenameRecoveredVmsParams.Get() == nil {
		var ret RecoveredOrClonedVmsRenameConfig
		return ret
	}
	return *o.RenameRecoveredVmsParams.Get()
}

// GetRenameRecoveredVmsParamsOk returns a tuple with the RenameRecoveredVmsParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperVTargetParamsForRecoverVm) GetRenameRecoveredVmsParamsOk() (*RecoveredOrClonedVmsRenameConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RenameRecoveredVmsParams.Get(), o.RenameRecoveredVmsParams.IsSet()
}

// HasRenameRecoveredVmsParams returns a boolean if a field has been set.
func (o *HyperVTargetParamsForRecoverVm) HasRenameRecoveredVmsParams() bool {
	if o != nil && o.RenameRecoveredVmsParams.IsSet() {
		return true
	}

	return false
}

// SetRenameRecoveredVmsParams gets a reference to the given NullableRecoveredOrClonedVmsRenameConfig and assigns it to the RenameRecoveredVmsParams field.
func (o *HyperVTargetParamsForRecoverVm) SetRenameRecoveredVmsParams(v RecoveredOrClonedVmsRenameConfig) {
	o.RenameRecoveredVmsParams.Set(&v)
}
// SetRenameRecoveredVmsParamsNil sets the value for RenameRecoveredVmsParams to be an explicit nil
func (o *HyperVTargetParamsForRecoverVm) SetRenameRecoveredVmsParamsNil() {
	o.RenameRecoveredVmsParams.Set(nil)
}

// UnsetRenameRecoveredVmsParams ensures that no value is present for RenameRecoveredVmsParams, not even an explicit nil
func (o *HyperVTargetParamsForRecoverVm) UnsetRenameRecoveredVmsParams() {
	o.RenameRecoveredVmsParams.Unset()
}

// GetRecoveryTargetConfig returns the RecoveryTargetConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperVTargetParamsForRecoverVm) GetRecoveryTargetConfig() HyperVVmRecoveryTargetConfig {
	if o == nil || o.RecoveryTargetConfig.Get() == nil {
		var ret HyperVVmRecoveryTargetConfig
		return ret
	}
	return *o.RecoveryTargetConfig.Get()
}

// GetRecoveryTargetConfigOk returns a tuple with the RecoveryTargetConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperVTargetParamsForRecoverVm) GetRecoveryTargetConfigOk() (*HyperVVmRecoveryTargetConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RecoveryTargetConfig.Get(), o.RecoveryTargetConfig.IsSet()
}

// HasRecoveryTargetConfig returns a boolean if a field has been set.
func (o *HyperVTargetParamsForRecoverVm) HasRecoveryTargetConfig() bool {
	if o != nil && o.RecoveryTargetConfig.IsSet() {
		return true
	}

	return false
}

// SetRecoveryTargetConfig gets a reference to the given NullableHyperVVmRecoveryTargetConfig and assigns it to the RecoveryTargetConfig field.
func (o *HyperVTargetParamsForRecoverVm) SetRecoveryTargetConfig(v HyperVVmRecoveryTargetConfig) {
	o.RecoveryTargetConfig.Set(&v)
}
// SetRecoveryTargetConfigNil sets the value for RecoveryTargetConfig to be an explicit nil
func (o *HyperVTargetParamsForRecoverVm) SetRecoveryTargetConfigNil() {
	o.RecoveryTargetConfig.Set(nil)
}

// UnsetRecoveryTargetConfig ensures that no value is present for RecoveryTargetConfig, not even an explicit nil
func (o *HyperVTargetParamsForRecoverVm) UnsetRecoveryTargetConfig() {
	o.RecoveryTargetConfig.Unset()
}

// GetPowerOnVms returns the PowerOnVms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperVTargetParamsForRecoverVm) GetPowerOnVms() bool {
	if o == nil || o.PowerOnVms.Get() == nil {
		var ret bool
		return ret
	}
	return *o.PowerOnVms.Get()
}

// GetPowerOnVmsOk returns a tuple with the PowerOnVms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperVTargetParamsForRecoverVm) GetPowerOnVmsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PowerOnVms.Get(), o.PowerOnVms.IsSet()
}

// HasPowerOnVms returns a boolean if a field has been set.
func (o *HyperVTargetParamsForRecoverVm) HasPowerOnVms() bool {
	if o != nil && o.PowerOnVms.IsSet() {
		return true
	}

	return false
}

// SetPowerOnVms gets a reference to the given NullableBool and assigns it to the PowerOnVms field.
func (o *HyperVTargetParamsForRecoverVm) SetPowerOnVms(v bool) {
	o.PowerOnVms.Set(&v)
}
// SetPowerOnVmsNil sets the value for PowerOnVms to be an explicit nil
func (o *HyperVTargetParamsForRecoverVm) SetPowerOnVmsNil() {
	o.PowerOnVms.Set(nil)
}

// UnsetPowerOnVms ensures that no value is present for PowerOnVms, not even an explicit nil
func (o *HyperVTargetParamsForRecoverVm) UnsetPowerOnVms() {
	o.PowerOnVms.Unset()
}

// GetContinueOnError returns the ContinueOnError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperVTargetParamsForRecoverVm) GetContinueOnError() bool {
	if o == nil || o.ContinueOnError.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ContinueOnError.Get()
}

// GetContinueOnErrorOk returns a tuple with the ContinueOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperVTargetParamsForRecoverVm) GetContinueOnErrorOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContinueOnError.Get(), o.ContinueOnError.IsSet()
}

// HasContinueOnError returns a boolean if a field has been set.
func (o *HyperVTargetParamsForRecoverVm) HasContinueOnError() bool {
	if o != nil && o.ContinueOnError.IsSet() {
		return true
	}

	return false
}

// SetContinueOnError gets a reference to the given NullableBool and assigns it to the ContinueOnError field.
func (o *HyperVTargetParamsForRecoverVm) SetContinueOnError(v bool) {
	o.ContinueOnError.Set(&v)
}
// SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil
func (o *HyperVTargetParamsForRecoverVm) SetContinueOnErrorNil() {
	o.ContinueOnError.Set(nil)
}

// UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
func (o *HyperVTargetParamsForRecoverVm) UnsetContinueOnError() {
	o.ContinueOnError.Unset()
}

// GetInstantRecovery returns the InstantRecovery field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperVTargetParamsForRecoverVm) GetInstantRecovery() bool {
	if o == nil || o.InstantRecovery.Get() == nil {
		var ret bool
		return ret
	}
	return *o.InstantRecovery.Get()
}

// GetInstantRecoveryOk returns a tuple with the InstantRecovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperVTargetParamsForRecoverVm) GetInstantRecoveryOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InstantRecovery.Get(), o.InstantRecovery.IsSet()
}

// HasInstantRecovery returns a boolean if a field has been set.
func (o *HyperVTargetParamsForRecoverVm) HasInstantRecovery() bool {
	if o != nil && o.InstantRecovery.IsSet() {
		return true
	}

	return false
}

// SetInstantRecovery gets a reference to the given NullableBool and assigns it to the InstantRecovery field.
func (o *HyperVTargetParamsForRecoverVm) SetInstantRecovery(v bool) {
	o.InstantRecovery.Set(&v)
}
// SetInstantRecoveryNil sets the value for InstantRecovery to be an explicit nil
func (o *HyperVTargetParamsForRecoverVm) SetInstantRecoveryNil() {
	o.InstantRecovery.Set(nil)
}

// UnsetInstantRecovery ensures that no value is present for InstantRecovery, not even an explicit nil
func (o *HyperVTargetParamsForRecoverVm) UnsetInstantRecovery() {
	o.InstantRecovery.Unset()
}

// GetUseSmbService returns the UseSmbService field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperVTargetParamsForRecoverVm) GetUseSmbService() bool {
	if o == nil || o.UseSmbService.Get() == nil {
		var ret bool
		return ret
	}
	return *o.UseSmbService.Get()
}

// GetUseSmbServiceOk returns a tuple with the UseSmbService field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperVTargetParamsForRecoverVm) GetUseSmbServiceOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UseSmbService.Get(), o.UseSmbService.IsSet()
}

// HasUseSmbService returns a boolean if a field has been set.
func (o *HyperVTargetParamsForRecoverVm) HasUseSmbService() bool {
	if o != nil && o.UseSmbService.IsSet() {
		return true
	}

	return false
}

// SetUseSmbService gets a reference to the given NullableBool and assigns it to the UseSmbService field.
func (o *HyperVTargetParamsForRecoverVm) SetUseSmbService(v bool) {
	o.UseSmbService.Set(&v)
}
// SetUseSmbServiceNil sets the value for UseSmbService to be an explicit nil
func (o *HyperVTargetParamsForRecoverVm) SetUseSmbServiceNil() {
	o.UseSmbService.Set(nil)
}

// UnsetUseSmbService ensures that no value is present for UseSmbService, not even an explicit nil
func (o *HyperVTargetParamsForRecoverVm) UnsetUseSmbService() {
	o.UseSmbService.Unset()
}

// GetPreserveUuids returns the PreserveUuids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperVTargetParamsForRecoverVm) GetPreserveUuids() bool {
	if o == nil || o.PreserveUuids.Get() == nil {
		var ret bool
		return ret
	}
	return *o.PreserveUuids.Get()
}

// GetPreserveUuidsOk returns a tuple with the PreserveUuids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperVTargetParamsForRecoverVm) GetPreserveUuidsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PreserveUuids.Get(), o.PreserveUuids.IsSet()
}

// HasPreserveUuids returns a boolean if a field has been set.
func (o *HyperVTargetParamsForRecoverVm) HasPreserveUuids() bool {
	if o != nil && o.PreserveUuids.IsSet() {
		return true
	}

	return false
}

// SetPreserveUuids gets a reference to the given NullableBool and assigns it to the PreserveUuids field.
func (o *HyperVTargetParamsForRecoverVm) SetPreserveUuids(v bool) {
	o.PreserveUuids.Set(&v)
}
// SetPreserveUuidsNil sets the value for PreserveUuids to be an explicit nil
func (o *HyperVTargetParamsForRecoverVm) SetPreserveUuidsNil() {
	o.PreserveUuids.Set(nil)
}

// UnsetPreserveUuids ensures that no value is present for PreserveUuids, not even an explicit nil
func (o *HyperVTargetParamsForRecoverVm) UnsetPreserveUuids() {
	o.PreserveUuids.Unset()
}

// GetVlanConfig returns the VlanConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperVTargetParamsForRecoverVm) GetVlanConfig() RecoveryVlanConfig {
	if o == nil || o.VlanConfig.Get() == nil {
		var ret RecoveryVlanConfig
		return ret
	}
	return *o.VlanConfig.Get()
}

// GetVlanConfigOk returns a tuple with the VlanConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperVTargetParamsForRecoverVm) GetVlanConfigOk() (*RecoveryVlanConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VlanConfig.Get(), o.VlanConfig.IsSet()
}

// HasVlanConfig returns a boolean if a field has been set.
func (o *HyperVTargetParamsForRecoverVm) HasVlanConfig() bool {
	if o != nil && o.VlanConfig.IsSet() {
		return true
	}

	return false
}

// SetVlanConfig gets a reference to the given NullableRecoveryVlanConfig and assigns it to the VlanConfig field.
func (o *HyperVTargetParamsForRecoverVm) SetVlanConfig(v RecoveryVlanConfig) {
	o.VlanConfig.Set(&v)
}
// SetVlanConfigNil sets the value for VlanConfig to be an explicit nil
func (o *HyperVTargetParamsForRecoverVm) SetVlanConfigNil() {
	o.VlanConfig.Set(nil)
}

// UnsetVlanConfig ensures that no value is present for VlanConfig, not even an explicit nil
func (o *HyperVTargetParamsForRecoverVm) UnsetVlanConfig() {
	o.VlanConfig.Unset()
}

func (o HyperVTargetParamsForRecoverVm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RenameRecoveredVmsParams.IsSet() {
		toSerialize["renameRecoveredVmsParams"] = o.RenameRecoveredVmsParams.Get()
	}
	if o.RecoveryTargetConfig.IsSet() {
		toSerialize["recoveryTargetConfig"] = o.RecoveryTargetConfig.Get()
	}
	if o.PowerOnVms.IsSet() {
		toSerialize["powerOnVms"] = o.PowerOnVms.Get()
	}
	if o.ContinueOnError.IsSet() {
		toSerialize["continueOnError"] = o.ContinueOnError.Get()
	}
	if o.InstantRecovery.IsSet() {
		toSerialize["instantRecovery"] = o.InstantRecovery.Get()
	}
	if o.UseSmbService.IsSet() {
		toSerialize["useSmbService"] = o.UseSmbService.Get()
	}
	if o.PreserveUuids.IsSet() {
		toSerialize["preserveUuids"] = o.PreserveUuids.Get()
	}
	if o.VlanConfig.IsSet() {
		toSerialize["vlanConfig"] = o.VlanConfig.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableHyperVTargetParamsForRecoverVm struct {
	value *HyperVTargetParamsForRecoverVm
	isSet bool
}

func (v NullableHyperVTargetParamsForRecoverVm) Get() *HyperVTargetParamsForRecoverVm {
	return v.value
}

func (v *NullableHyperVTargetParamsForRecoverVm) Set(val *HyperVTargetParamsForRecoverVm) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperVTargetParamsForRecoverVm) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperVTargetParamsForRecoverVm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperVTargetParamsForRecoverVm(val *HyperVTargetParamsForRecoverVm) *NullableHyperVTargetParamsForRecoverVm {
	return &NullableHyperVTargetParamsForRecoverVm{value: val, isSet: true}
}

func (v NullableHyperVTargetParamsForRecoverVm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperVTargetParamsForRecoverVm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


