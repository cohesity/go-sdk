/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package onprem
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/onprem/utils"
	"fmt"
)

var _ = NullableBool{}

// RecoverVmwareVAppOriginalSourceConfig Specifies the Source configuration if vApps are being recovered to Original Source.
type RecoverVmwareVAppOriginalSourceConfig struct {
	// Specifies the networking configuration to be applied to the recovered vApps.
	NetworkConfig NullableRecoverVmwareVmOriginalSourceNetworkConfig `json:"networkConfig,omitempty"`
}

// NewRecoverVmwareVAppOriginalSourceConfig instantiates a new RecoverVmwareVAppOriginalSourceConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverVmwareVAppOriginalSourceConfig() *RecoverVmwareVAppOriginalSourceConfig {
	this := RecoverVmwareVAppOriginalSourceConfig{}
	return &this
}

// NewRecoverVmwareVAppOriginalSourceConfigWithDefaults instantiates a new RecoverVmwareVAppOriginalSourceConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverVmwareVAppOriginalSourceConfigWithDefaults() *RecoverVmwareVAppOriginalSourceConfig {
	this := RecoverVmwareVAppOriginalSourceConfig{}
	return &this
}

// GetNetworkConfig returns the NetworkConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverVmwareVAppOriginalSourceConfig) GetNetworkConfig() RecoverVmwareVmOriginalSourceNetworkConfig {
	if o == nil || o.NetworkConfig.Get() == nil {
		var ret RecoverVmwareVmOriginalSourceNetworkConfig
		return ret
	}
	return *o.NetworkConfig.Get()
}

// GetNetworkConfigOk returns a tuple with the NetworkConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverVmwareVAppOriginalSourceConfig) GetNetworkConfigOk() (*RecoverVmwareVmOriginalSourceNetworkConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NetworkConfig.Get(), o.NetworkConfig.IsSet()
}

// HasNetworkConfig returns a boolean if a field has been set.
func (o *RecoverVmwareVAppOriginalSourceConfig) HasNetworkConfig() bool {
	if o != nil && o.NetworkConfig.IsSet() {
		return true
	}

	return false
}

// SetNetworkConfig gets a reference to the given NullableRecoverVmwareVmOriginalSourceNetworkConfig and assigns it to the NetworkConfig field.
func (o *RecoverVmwareVAppOriginalSourceConfig) SetNetworkConfig(v RecoverVmwareVmOriginalSourceNetworkConfig) {
	o.NetworkConfig.Set(&v)
}
// SetNetworkConfigNil sets the value for NetworkConfig to be an explicit nil
func (o *RecoverVmwareVAppOriginalSourceConfig) SetNetworkConfigNil() {
	o.NetworkConfig.Set(nil)
}

// UnsetNetworkConfig ensures that no value is present for NetworkConfig, not even an explicit nil
func (o *RecoverVmwareVAppOriginalSourceConfig) UnsetNetworkConfig() {
	o.NetworkConfig.Unset()
}

func (o RecoverVmwareVAppOriginalSourceConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NetworkConfig.IsSet() {
		toSerialize["networkConfig"] = o.NetworkConfig.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRecoverVmwareVAppOriginalSourceConfig struct {
	value *RecoverVmwareVAppOriginalSourceConfig
	isSet bool
}

func (v NullableRecoverVmwareVAppOriginalSourceConfig) Get() *RecoverVmwareVAppOriginalSourceConfig {
	return v.value
}

func (v *NullableRecoverVmwareVAppOriginalSourceConfig) Set(val *RecoverVmwareVAppOriginalSourceConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverVmwareVAppOriginalSourceConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverVmwareVAppOriginalSourceConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverVmwareVAppOriginalSourceConfig(val *RecoverVmwareVAppOriginalSourceConfig) *NullableRecoverVmwareVAppOriginalSourceConfig {
	return &NullableRecoverVmwareVAppOriginalSourceConfig{value: val, isSet: true}
}

func (v NullableRecoverVmwareVAppOriginalSourceConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverVmwareVAppOriginalSourceConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o RecoverVmwareVAppOriginalSourceConfig) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}