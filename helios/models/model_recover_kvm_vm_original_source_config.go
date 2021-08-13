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

// RecoverKvmVmOriginalSourceConfig Specifies the Source configuration if VM's are being recovered to Original Source.
type RecoverKvmVmOriginalSourceConfig struct {
	// Specifies the networking configuration to be applied to the recovered VMs.
	NetworkConfig NullableRecoverKvmVmOriginalSourceNetworkConfig `json:"networkConfig,omitempty"`
}

// NewRecoverKvmVmOriginalSourceConfig instantiates a new RecoverKvmVmOriginalSourceConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverKvmVmOriginalSourceConfig() *RecoverKvmVmOriginalSourceConfig {
	this := RecoverKvmVmOriginalSourceConfig{}
	return &this
}

// NewRecoverKvmVmOriginalSourceConfigWithDefaults instantiates a new RecoverKvmVmOriginalSourceConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverKvmVmOriginalSourceConfigWithDefaults() *RecoverKvmVmOriginalSourceConfig {
	this := RecoverKvmVmOriginalSourceConfig{}
	return &this
}

// GetNetworkConfig returns the NetworkConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverKvmVmOriginalSourceConfig) GetNetworkConfig() RecoverKvmVmOriginalSourceNetworkConfig {
	if o == nil || o.NetworkConfig.Get() == nil {
		var ret RecoverKvmVmOriginalSourceNetworkConfig
		return ret
	}
	return *o.NetworkConfig.Get()
}

// GetNetworkConfigOk returns a tuple with the NetworkConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverKvmVmOriginalSourceConfig) GetNetworkConfigOk() (*RecoverKvmVmOriginalSourceNetworkConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NetworkConfig.Get(), o.NetworkConfig.IsSet()
}

// HasNetworkConfig returns a boolean if a field has been set.
func (o *RecoverKvmVmOriginalSourceConfig) HasNetworkConfig() bool {
	if o != nil && o.NetworkConfig.IsSet() {
		return true
	}

	return false
}

// SetNetworkConfig gets a reference to the given NullableRecoverKvmVmOriginalSourceNetworkConfig and assigns it to the NetworkConfig field.
func (o *RecoverKvmVmOriginalSourceConfig) SetNetworkConfig(v RecoverKvmVmOriginalSourceNetworkConfig) {
	o.NetworkConfig.Set(&v)
}
// SetNetworkConfigNil sets the value for NetworkConfig to be an explicit nil
func (o *RecoverKvmVmOriginalSourceConfig) SetNetworkConfigNil() {
	o.NetworkConfig.Set(nil)
}

// UnsetNetworkConfig ensures that no value is present for NetworkConfig, not even an explicit nil
func (o *RecoverKvmVmOriginalSourceConfig) UnsetNetworkConfig() {
	o.NetworkConfig.Unset()
}

func (o RecoverKvmVmOriginalSourceConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NetworkConfig.IsSet() {
		toSerialize["networkConfig"] = o.NetworkConfig.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRecoverKvmVmOriginalSourceConfig struct {
	value *RecoverKvmVmOriginalSourceConfig
	isSet bool
}

func (v NullableRecoverKvmVmOriginalSourceConfig) Get() *RecoverKvmVmOriginalSourceConfig {
	return v.value
}

func (v *NullableRecoverKvmVmOriginalSourceConfig) Set(val *RecoverKvmVmOriginalSourceConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverKvmVmOriginalSourceConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverKvmVmOriginalSourceConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverKvmVmOriginalSourceConfig(val *RecoverKvmVmOriginalSourceConfig) *NullableRecoverKvmVmOriginalSourceConfig {
	return &NullableRecoverKvmVmOriginalSourceConfig{value: val, isSet: true}
}

func (v NullableRecoverKvmVmOriginalSourceConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverKvmVmOriginalSourceConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


