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

// RecoverAcropolisVmOriginalSourceConfig Specifies the Source configuration if VM's are being recovered to Original Source.
type RecoverAcropolisVmOriginalSourceConfig struct {
	// Specifies the networking configuration to be applied to the recovered VMs.
	NetworkConfig NullableRecoverAcropolisVmOriginalSourceNetworkConfig `json:"networkConfig,omitempty"`
}

// NewRecoverAcropolisVmOriginalSourceConfig instantiates a new RecoverAcropolisVmOriginalSourceConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverAcropolisVmOriginalSourceConfig() *RecoverAcropolisVmOriginalSourceConfig {
	this := RecoverAcropolisVmOriginalSourceConfig{}
	return &this
}

// NewRecoverAcropolisVmOriginalSourceConfigWithDefaults instantiates a new RecoverAcropolisVmOriginalSourceConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverAcropolisVmOriginalSourceConfigWithDefaults() *RecoverAcropolisVmOriginalSourceConfig {
	this := RecoverAcropolisVmOriginalSourceConfig{}
	return &this
}

// GetNetworkConfig returns the NetworkConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverAcropolisVmOriginalSourceConfig) GetNetworkConfig() RecoverAcropolisVmOriginalSourceNetworkConfig {
	if o == nil || o.NetworkConfig.Get() == nil {
		var ret RecoverAcropolisVmOriginalSourceNetworkConfig
		return ret
	}
	return *o.NetworkConfig.Get()
}

// GetNetworkConfigOk returns a tuple with the NetworkConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverAcropolisVmOriginalSourceConfig) GetNetworkConfigOk() (*RecoverAcropolisVmOriginalSourceNetworkConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NetworkConfig.Get(), o.NetworkConfig.IsSet()
}

// HasNetworkConfig returns a boolean if a field has been set.
func (o *RecoverAcropolisVmOriginalSourceConfig) HasNetworkConfig() bool {
	if o != nil && o.NetworkConfig.IsSet() {
		return true
	}

	return false
}

// SetNetworkConfig gets a reference to the given NullableRecoverAcropolisVmOriginalSourceNetworkConfig and assigns it to the NetworkConfig field.
func (o *RecoverAcropolisVmOriginalSourceConfig) SetNetworkConfig(v RecoverAcropolisVmOriginalSourceNetworkConfig) {
	o.NetworkConfig.Set(&v)
}
// SetNetworkConfigNil sets the value for NetworkConfig to be an explicit nil
func (o *RecoverAcropolisVmOriginalSourceConfig) SetNetworkConfigNil() {
	o.NetworkConfig.Set(nil)
}

// UnsetNetworkConfig ensures that no value is present for NetworkConfig, not even an explicit nil
func (o *RecoverAcropolisVmOriginalSourceConfig) UnsetNetworkConfig() {
	o.NetworkConfig.Unset()
}

func (o RecoverAcropolisVmOriginalSourceConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NetworkConfig.IsSet() {
		toSerialize["networkConfig"] = o.NetworkConfig.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRecoverAcropolisVmOriginalSourceConfig struct {
	value *RecoverAcropolisVmOriginalSourceConfig
	isSet bool
}

func (v NullableRecoverAcropolisVmOriginalSourceConfig) Get() *RecoverAcropolisVmOriginalSourceConfig {
	return v.value
}

func (v *NullableRecoverAcropolisVmOriginalSourceConfig) Set(val *RecoverAcropolisVmOriginalSourceConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverAcropolisVmOriginalSourceConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverAcropolisVmOriginalSourceConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverAcropolisVmOriginalSourceConfig(val *RecoverAcropolisVmOriginalSourceConfig) *NullableRecoverAcropolisVmOriginalSourceConfig {
	return &NullableRecoverAcropolisVmOriginalSourceConfig{value: val, isSet: true}
}

func (v NullableRecoverAcropolisVmOriginalSourceConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverAcropolisVmOriginalSourceConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


