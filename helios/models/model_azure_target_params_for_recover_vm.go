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

// AzureTargetParamsForRecoverVm Specifies the parameters for an Azure recovery target.
type AzureTargetParamsForRecoverVm struct {
	// Specifies params to rename the VMs that are recovered. If not specified, the original names of the VMs are preserved.
	RenameRecoveredVmsParams NullableRecoveredOrClonedVmsRenameConfig `json:"renameRecoveredVmsParams,omitempty"`
	// Specifies the recovery target configuration if recovery has to be done to a different location which is different from original source or to original Source with different configuration. If not specified, then the recovery of the vms will be performed to original location with all configuration parameters retained.
	RecoveryTargetConfig NullableAzureVmRecoveryTargetConfig `json:"recoveryTargetConfig,omitempty"`
	// Specifies whether to power on vms after recovery. If not specified, or false, recovered vms will be in powered off state.
	PowerOnVms NullableBool `json:"powerOnVms,omitempty"`
	// Specifies whether to continue recovering other vms if one of vms failed to recover. Default value is false.
	ContinueOnError NullableBool `json:"continueOnError,omitempty"`
}

// NewAzureTargetParamsForRecoverVm instantiates a new AzureTargetParamsForRecoverVm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureTargetParamsForRecoverVm() *AzureTargetParamsForRecoverVm {
	this := AzureTargetParamsForRecoverVm{}
	return &this
}

// NewAzureTargetParamsForRecoverVmWithDefaults instantiates a new AzureTargetParamsForRecoverVm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureTargetParamsForRecoverVmWithDefaults() *AzureTargetParamsForRecoverVm {
	this := AzureTargetParamsForRecoverVm{}
	return &this
}

// GetRenameRecoveredVmsParams returns the RenameRecoveredVmsParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureTargetParamsForRecoverVm) GetRenameRecoveredVmsParams() RecoveredOrClonedVmsRenameConfig {
	if o == nil || o.RenameRecoveredVmsParams.Get() == nil {
		var ret RecoveredOrClonedVmsRenameConfig
		return ret
	}
	return *o.RenameRecoveredVmsParams.Get()
}

// GetRenameRecoveredVmsParamsOk returns a tuple with the RenameRecoveredVmsParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureTargetParamsForRecoverVm) GetRenameRecoveredVmsParamsOk() (*RecoveredOrClonedVmsRenameConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RenameRecoveredVmsParams.Get(), o.RenameRecoveredVmsParams.IsSet()
}

// HasRenameRecoveredVmsParams returns a boolean if a field has been set.
func (o *AzureTargetParamsForRecoverVm) HasRenameRecoveredVmsParams() bool {
	if o != nil && o.RenameRecoveredVmsParams.IsSet() {
		return true
	}

	return false
}

// SetRenameRecoveredVmsParams gets a reference to the given NullableRecoveredOrClonedVmsRenameConfig and assigns it to the RenameRecoveredVmsParams field.
func (o *AzureTargetParamsForRecoverVm) SetRenameRecoveredVmsParams(v RecoveredOrClonedVmsRenameConfig) {
	o.RenameRecoveredVmsParams.Set(&v)
}
// SetRenameRecoveredVmsParamsNil sets the value for RenameRecoveredVmsParams to be an explicit nil
func (o *AzureTargetParamsForRecoverVm) SetRenameRecoveredVmsParamsNil() {
	o.RenameRecoveredVmsParams.Set(nil)
}

// UnsetRenameRecoveredVmsParams ensures that no value is present for RenameRecoveredVmsParams, not even an explicit nil
func (o *AzureTargetParamsForRecoverVm) UnsetRenameRecoveredVmsParams() {
	o.RenameRecoveredVmsParams.Unset()
}

// GetRecoveryTargetConfig returns the RecoveryTargetConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureTargetParamsForRecoverVm) GetRecoveryTargetConfig() AzureVmRecoveryTargetConfig {
	if o == nil || o.RecoveryTargetConfig.Get() == nil {
		var ret AzureVmRecoveryTargetConfig
		return ret
	}
	return *o.RecoveryTargetConfig.Get()
}

// GetRecoveryTargetConfigOk returns a tuple with the RecoveryTargetConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureTargetParamsForRecoverVm) GetRecoveryTargetConfigOk() (*AzureVmRecoveryTargetConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RecoveryTargetConfig.Get(), o.RecoveryTargetConfig.IsSet()
}

// HasRecoveryTargetConfig returns a boolean if a field has been set.
func (o *AzureTargetParamsForRecoverVm) HasRecoveryTargetConfig() bool {
	if o != nil && o.RecoveryTargetConfig.IsSet() {
		return true
	}

	return false
}

// SetRecoveryTargetConfig gets a reference to the given NullableAzureVmRecoveryTargetConfig and assigns it to the RecoveryTargetConfig field.
func (o *AzureTargetParamsForRecoverVm) SetRecoveryTargetConfig(v AzureVmRecoveryTargetConfig) {
	o.RecoveryTargetConfig.Set(&v)
}
// SetRecoveryTargetConfigNil sets the value for RecoveryTargetConfig to be an explicit nil
func (o *AzureTargetParamsForRecoverVm) SetRecoveryTargetConfigNil() {
	o.RecoveryTargetConfig.Set(nil)
}

// UnsetRecoveryTargetConfig ensures that no value is present for RecoveryTargetConfig, not even an explicit nil
func (o *AzureTargetParamsForRecoverVm) UnsetRecoveryTargetConfig() {
	o.RecoveryTargetConfig.Unset()
}

// GetPowerOnVms returns the PowerOnVms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureTargetParamsForRecoverVm) GetPowerOnVms() bool {
	if o == nil || o.PowerOnVms.Get() == nil {
		var ret bool
		return ret
	}
	return *o.PowerOnVms.Get()
}

// GetPowerOnVmsOk returns a tuple with the PowerOnVms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureTargetParamsForRecoverVm) GetPowerOnVmsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PowerOnVms.Get(), o.PowerOnVms.IsSet()
}

// HasPowerOnVms returns a boolean if a field has been set.
func (o *AzureTargetParamsForRecoverVm) HasPowerOnVms() bool {
	if o != nil && o.PowerOnVms.IsSet() {
		return true
	}

	return false
}

// SetPowerOnVms gets a reference to the given NullableBool and assigns it to the PowerOnVms field.
func (o *AzureTargetParamsForRecoverVm) SetPowerOnVms(v bool) {
	o.PowerOnVms.Set(&v)
}
// SetPowerOnVmsNil sets the value for PowerOnVms to be an explicit nil
func (o *AzureTargetParamsForRecoverVm) SetPowerOnVmsNil() {
	o.PowerOnVms.Set(nil)
}

// UnsetPowerOnVms ensures that no value is present for PowerOnVms, not even an explicit nil
func (o *AzureTargetParamsForRecoverVm) UnsetPowerOnVms() {
	o.PowerOnVms.Unset()
}

// GetContinueOnError returns the ContinueOnError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureTargetParamsForRecoverVm) GetContinueOnError() bool {
	if o == nil || o.ContinueOnError.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ContinueOnError.Get()
}

// GetContinueOnErrorOk returns a tuple with the ContinueOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureTargetParamsForRecoverVm) GetContinueOnErrorOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContinueOnError.Get(), o.ContinueOnError.IsSet()
}

// HasContinueOnError returns a boolean if a field has been set.
func (o *AzureTargetParamsForRecoverVm) HasContinueOnError() bool {
	if o != nil && o.ContinueOnError.IsSet() {
		return true
	}

	return false
}

// SetContinueOnError gets a reference to the given NullableBool and assigns it to the ContinueOnError field.
func (o *AzureTargetParamsForRecoverVm) SetContinueOnError(v bool) {
	o.ContinueOnError.Set(&v)
}
// SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil
func (o *AzureTargetParamsForRecoverVm) SetContinueOnErrorNil() {
	o.ContinueOnError.Set(nil)
}

// UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
func (o *AzureTargetParamsForRecoverVm) UnsetContinueOnError() {
	o.ContinueOnError.Unset()
}

func (o AzureTargetParamsForRecoverVm) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableAzureTargetParamsForRecoverVm struct {
	value *AzureTargetParamsForRecoverVm
	isSet bool
}

func (v NullableAzureTargetParamsForRecoverVm) Get() *AzureTargetParamsForRecoverVm {
	return v.value
}

func (v *NullableAzureTargetParamsForRecoverVm) Set(val *AzureTargetParamsForRecoverVm) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureTargetParamsForRecoverVm) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureTargetParamsForRecoverVm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureTargetParamsForRecoverVm(val *AzureTargetParamsForRecoverVm) *NullableAzureTargetParamsForRecoverVm {
	return &NullableAzureTargetParamsForRecoverVm{value: val, isSet: true}
}

func (v NullableAzureTargetParamsForRecoverVm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureTargetParamsForRecoverVm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


