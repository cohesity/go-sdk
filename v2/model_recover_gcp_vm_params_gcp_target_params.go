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

// checks if the RecoverGcpVmParamsGcpTargetParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoverGcpVmParamsGcpTargetParams{}

// RecoverGcpVmParamsGcpTargetParams Specifies the params for recovering to a GCP target.
type RecoverGcpVmParamsGcpTargetParams struct {
	// Specifies whether to continue recovering other vms if one of vms failed to recover. Default value is false.
	ContinueOnError NullableBool `json:"continueOnError,omitempty"`
	// Specifies whether to power on vms after recovery. If not specified, or false, recovered vms will be in powered off state.
	PowerOnVms NullableBool `json:"powerOnVms,omitempty"`
	RecoveryTargetConfig NullableGcpTargetParamsForRecoverVmRecoveryTargetConfig `json:"recoveryTargetConfig,omitempty"`
	RenameRecoveredVmsParams NullableAcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams `json:"renameRecoveredVmsParams,omitempty"`
}

// NewRecoverGcpVmParamsGcpTargetParams instantiates a new RecoverGcpVmParamsGcpTargetParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverGcpVmParamsGcpTargetParams() *RecoverGcpVmParamsGcpTargetParams {
	this := RecoverGcpVmParamsGcpTargetParams{}
	return &this
}

// NewRecoverGcpVmParamsGcpTargetParamsWithDefaults instantiates a new RecoverGcpVmParamsGcpTargetParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverGcpVmParamsGcpTargetParamsWithDefaults() *RecoverGcpVmParamsGcpTargetParams {
	this := RecoverGcpVmParamsGcpTargetParams{}
	return &this
}

// GetContinueOnError returns the ContinueOnError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverGcpVmParamsGcpTargetParams) GetContinueOnError() bool {
	if o == nil || IsNil(o.ContinueOnError.Get()) {
		var ret bool
		return ret
	}
	return *o.ContinueOnError.Get()
}

// GetContinueOnErrorOk returns a tuple with the ContinueOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverGcpVmParamsGcpTargetParams) GetContinueOnErrorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContinueOnError.Get(), o.ContinueOnError.IsSet()
}

// HasContinueOnError returns a boolean if a field has been set.
func (o *RecoverGcpVmParamsGcpTargetParams) HasContinueOnError() bool {
	if o != nil && o.ContinueOnError.IsSet() {
		return true
	}

	return false
}

// SetContinueOnError gets a reference to the given NullableBool and assigns it to the ContinueOnError field.
func (o *RecoverGcpVmParamsGcpTargetParams) SetContinueOnError(v bool) {
	o.ContinueOnError.Set(&v)
}
// SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil
func (o *RecoverGcpVmParamsGcpTargetParams) SetContinueOnErrorNil() {
	o.ContinueOnError.Set(nil)
}

// UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
func (o *RecoverGcpVmParamsGcpTargetParams) UnsetContinueOnError() {
	o.ContinueOnError.Unset()
}

// GetPowerOnVms returns the PowerOnVms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverGcpVmParamsGcpTargetParams) GetPowerOnVms() bool {
	if o == nil || IsNil(o.PowerOnVms.Get()) {
		var ret bool
		return ret
	}
	return *o.PowerOnVms.Get()
}

// GetPowerOnVmsOk returns a tuple with the PowerOnVms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverGcpVmParamsGcpTargetParams) GetPowerOnVmsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PowerOnVms.Get(), o.PowerOnVms.IsSet()
}

// HasPowerOnVms returns a boolean if a field has been set.
func (o *RecoverGcpVmParamsGcpTargetParams) HasPowerOnVms() bool {
	if o != nil && o.PowerOnVms.IsSet() {
		return true
	}

	return false
}

// SetPowerOnVms gets a reference to the given NullableBool and assigns it to the PowerOnVms field.
func (o *RecoverGcpVmParamsGcpTargetParams) SetPowerOnVms(v bool) {
	o.PowerOnVms.Set(&v)
}
// SetPowerOnVmsNil sets the value for PowerOnVms to be an explicit nil
func (o *RecoverGcpVmParamsGcpTargetParams) SetPowerOnVmsNil() {
	o.PowerOnVms.Set(nil)
}

// UnsetPowerOnVms ensures that no value is present for PowerOnVms, not even an explicit nil
func (o *RecoverGcpVmParamsGcpTargetParams) UnsetPowerOnVms() {
	o.PowerOnVms.Unset()
}

// GetRecoveryTargetConfig returns the RecoveryTargetConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverGcpVmParamsGcpTargetParams) GetRecoveryTargetConfig() GcpTargetParamsForRecoverVmRecoveryTargetConfig {
	if o == nil || IsNil(o.RecoveryTargetConfig.Get()) {
		var ret GcpTargetParamsForRecoverVmRecoveryTargetConfig
		return ret
	}
	return *o.RecoveryTargetConfig.Get()
}

// GetRecoveryTargetConfigOk returns a tuple with the RecoveryTargetConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverGcpVmParamsGcpTargetParams) GetRecoveryTargetConfigOk() (*GcpTargetParamsForRecoverVmRecoveryTargetConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecoveryTargetConfig.Get(), o.RecoveryTargetConfig.IsSet()
}

// HasRecoveryTargetConfig returns a boolean if a field has been set.
func (o *RecoverGcpVmParamsGcpTargetParams) HasRecoveryTargetConfig() bool {
	if o != nil && o.RecoveryTargetConfig.IsSet() {
		return true
	}

	return false
}

// SetRecoveryTargetConfig gets a reference to the given NullableGcpTargetParamsForRecoverVmRecoveryTargetConfig and assigns it to the RecoveryTargetConfig field.
func (o *RecoverGcpVmParamsGcpTargetParams) SetRecoveryTargetConfig(v GcpTargetParamsForRecoverVmRecoveryTargetConfig) {
	o.RecoveryTargetConfig.Set(&v)
}
// SetRecoveryTargetConfigNil sets the value for RecoveryTargetConfig to be an explicit nil
func (o *RecoverGcpVmParamsGcpTargetParams) SetRecoveryTargetConfigNil() {
	o.RecoveryTargetConfig.Set(nil)
}

// UnsetRecoveryTargetConfig ensures that no value is present for RecoveryTargetConfig, not even an explicit nil
func (o *RecoverGcpVmParamsGcpTargetParams) UnsetRecoveryTargetConfig() {
	o.RecoveryTargetConfig.Unset()
}

// GetRenameRecoveredVmsParams returns the RenameRecoveredVmsParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverGcpVmParamsGcpTargetParams) GetRenameRecoveredVmsParams() AcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams {
	if o == nil || IsNil(o.RenameRecoveredVmsParams.Get()) {
		var ret AcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams
		return ret
	}
	return *o.RenameRecoveredVmsParams.Get()
}

// GetRenameRecoveredVmsParamsOk returns a tuple with the RenameRecoveredVmsParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverGcpVmParamsGcpTargetParams) GetRenameRecoveredVmsParamsOk() (*AcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.RenameRecoveredVmsParams.Get(), o.RenameRecoveredVmsParams.IsSet()
}

// HasRenameRecoveredVmsParams returns a boolean if a field has been set.
func (o *RecoverGcpVmParamsGcpTargetParams) HasRenameRecoveredVmsParams() bool {
	if o != nil && o.RenameRecoveredVmsParams.IsSet() {
		return true
	}

	return false
}

// SetRenameRecoveredVmsParams gets a reference to the given NullableAcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams and assigns it to the RenameRecoveredVmsParams field.
func (o *RecoverGcpVmParamsGcpTargetParams) SetRenameRecoveredVmsParams(v AcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams) {
	o.RenameRecoveredVmsParams.Set(&v)
}
// SetRenameRecoveredVmsParamsNil sets the value for RenameRecoveredVmsParams to be an explicit nil
func (o *RecoverGcpVmParamsGcpTargetParams) SetRenameRecoveredVmsParamsNil() {
	o.RenameRecoveredVmsParams.Set(nil)
}

// UnsetRenameRecoveredVmsParams ensures that no value is present for RenameRecoveredVmsParams, not even an explicit nil
func (o *RecoverGcpVmParamsGcpTargetParams) UnsetRenameRecoveredVmsParams() {
	o.RenameRecoveredVmsParams.Unset()
}

func (o RecoverGcpVmParamsGcpTargetParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoverGcpVmParamsGcpTargetParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ContinueOnError.IsSet() {
		toSerialize["continueOnError"] = o.ContinueOnError.Get()
	}
	if o.PowerOnVms.IsSet() {
		toSerialize["powerOnVms"] = o.PowerOnVms.Get()
	}
	if o.RecoveryTargetConfig.IsSet() {
		toSerialize["recoveryTargetConfig"] = o.RecoveryTargetConfig.Get()
	}
	if o.RenameRecoveredVmsParams.IsSet() {
		toSerialize["renameRecoveredVmsParams"] = o.RenameRecoveredVmsParams.Get()
	}
	return toSerialize, nil
}

type NullableRecoverGcpVmParamsGcpTargetParams struct {
	value *RecoverGcpVmParamsGcpTargetParams
	isSet bool
}

func (v NullableRecoverGcpVmParamsGcpTargetParams) Get() *RecoverGcpVmParamsGcpTargetParams {
	return v.value
}

func (v *NullableRecoverGcpVmParamsGcpTargetParams) Set(val *RecoverGcpVmParamsGcpTargetParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverGcpVmParamsGcpTargetParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverGcpVmParamsGcpTargetParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverGcpVmParamsGcpTargetParams(val *RecoverGcpVmParamsGcpTargetParams) *NullableRecoverGcpVmParamsGcpTargetParams {
	return &NullableRecoverGcpVmParamsGcpTargetParams{value: val, isSet: true}
}

func (v NullableRecoverGcpVmParamsGcpTargetParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverGcpVmParamsGcpTargetParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


