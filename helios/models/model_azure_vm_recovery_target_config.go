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

// AzureVmRecoveryTargetConfig Specifies the target object parameters to recover Azure vms.
type AzureVmRecoveryTargetConfig struct {
	// Specifies the parameter whether the recovery should be performed to a new or an existing Source Target.
	RecoverToNewSource bool `json:"recoverToNewSource"`
	// Specifies the new destination Source configuration parameters where the VMs will be recovered. This is mandatory if recoverToNewSource is set to true.
	NewSourceConfig NullableRecoverAzureVmNewSourceConfig `json:"newSourceConfig,omitempty"`
}

// NewAzureVmRecoveryTargetConfig instantiates a new AzureVmRecoveryTargetConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureVmRecoveryTargetConfig(recoverToNewSource bool) *AzureVmRecoveryTargetConfig {
	this := AzureVmRecoveryTargetConfig{}
	this.RecoverToNewSource = recoverToNewSource
	return &this
}

// NewAzureVmRecoveryTargetConfigWithDefaults instantiates a new AzureVmRecoveryTargetConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureVmRecoveryTargetConfigWithDefaults() *AzureVmRecoveryTargetConfig {
	this := AzureVmRecoveryTargetConfig{}
	return &this
}

// GetRecoverToNewSource returns the RecoverToNewSource field value
func (o *AzureVmRecoveryTargetConfig) GetRecoverToNewSource() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RecoverToNewSource
}

// GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field value
// and a boolean to check if the value has been set.
func (o *AzureVmRecoveryTargetConfig) GetRecoverToNewSourceOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RecoverToNewSource, true
}

// SetRecoverToNewSource sets field value
func (o *AzureVmRecoveryTargetConfig) SetRecoverToNewSource(v bool) {
	o.RecoverToNewSource = v
}

// GetNewSourceConfig returns the NewSourceConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureVmRecoveryTargetConfig) GetNewSourceConfig() RecoverAzureVmNewSourceConfig {
	if o == nil || o.NewSourceConfig.Get() == nil {
		var ret RecoverAzureVmNewSourceConfig
		return ret
	}
	return *o.NewSourceConfig.Get()
}

// GetNewSourceConfigOk returns a tuple with the NewSourceConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureVmRecoveryTargetConfig) GetNewSourceConfigOk() (*RecoverAzureVmNewSourceConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NewSourceConfig.Get(), o.NewSourceConfig.IsSet()
}

// HasNewSourceConfig returns a boolean if a field has been set.
func (o *AzureVmRecoveryTargetConfig) HasNewSourceConfig() bool {
	if o != nil && o.NewSourceConfig.IsSet() {
		return true
	}

	return false
}

// SetNewSourceConfig gets a reference to the given NullableRecoverAzureVmNewSourceConfig and assigns it to the NewSourceConfig field.
func (o *AzureVmRecoveryTargetConfig) SetNewSourceConfig(v RecoverAzureVmNewSourceConfig) {
	o.NewSourceConfig.Set(&v)
}
// SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil
func (o *AzureVmRecoveryTargetConfig) SetNewSourceConfigNil() {
	o.NewSourceConfig.Set(nil)
}

// UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
func (o *AzureVmRecoveryTargetConfig) UnsetNewSourceConfig() {
	o.NewSourceConfig.Unset()
}

func (o AzureVmRecoveryTargetConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["recoverToNewSource"] = o.RecoverToNewSource
	}
	if o.NewSourceConfig.IsSet() {
		toSerialize["newSourceConfig"] = o.NewSourceConfig.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAzureVmRecoveryTargetConfig struct {
	value *AzureVmRecoveryTargetConfig
	isSet bool
}

func (v NullableAzureVmRecoveryTargetConfig) Get() *AzureVmRecoveryTargetConfig {
	return v.value
}

func (v *NullableAzureVmRecoveryTargetConfig) Set(val *AzureVmRecoveryTargetConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureVmRecoveryTargetConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureVmRecoveryTargetConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureVmRecoveryTargetConfig(val *AzureVmRecoveryTargetConfig) *NullableAzureVmRecoveryTargetConfig {
	return &NullableAzureVmRecoveryTargetConfig{value: val, isSet: true}
}

func (v NullableAzureVmRecoveryTargetConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureVmRecoveryTargetConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


