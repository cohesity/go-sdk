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

// RecoverVmwareVmNewSourceNetworkConfig Specifies the network config parameters to be applied to VMware VMs if recovering to new Source.
type RecoverVmwareVmNewSourceNetworkConfig struct {
	// If this is set to true, then the network will be detached from the recovered VMs. All the other networking parameters set will be ignored if set to true. Default value is false.
	DetachNetwork NullableBool `json:"detachNetwork,omitempty"`
	// Specifies a new network configuration for the VM recovery.
	NewNetworkConfig NullableRecoverVmwareVmNewNetworkConfig `json:"newNetworkConfig,omitempty"`
}

// NewRecoverVmwareVmNewSourceNetworkConfig instantiates a new RecoverVmwareVmNewSourceNetworkConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverVmwareVmNewSourceNetworkConfig() *RecoverVmwareVmNewSourceNetworkConfig {
	this := RecoverVmwareVmNewSourceNetworkConfig{}
	return &this
}

// NewRecoverVmwareVmNewSourceNetworkConfigWithDefaults instantiates a new RecoverVmwareVmNewSourceNetworkConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverVmwareVmNewSourceNetworkConfigWithDefaults() *RecoverVmwareVmNewSourceNetworkConfig {
	this := RecoverVmwareVmNewSourceNetworkConfig{}
	return &this
}

// GetDetachNetwork returns the DetachNetwork field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverVmwareVmNewSourceNetworkConfig) GetDetachNetwork() bool {
	if o == nil || o.DetachNetwork.Get() == nil {
		var ret bool
		return ret
	}
	return *o.DetachNetwork.Get()
}

// GetDetachNetworkOk returns a tuple with the DetachNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverVmwareVmNewSourceNetworkConfig) GetDetachNetworkOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DetachNetwork.Get(), o.DetachNetwork.IsSet()
}

// HasDetachNetwork returns a boolean if a field has been set.
func (o *RecoverVmwareVmNewSourceNetworkConfig) HasDetachNetwork() bool {
	if o != nil && o.DetachNetwork.IsSet() {
		return true
	}

	return false
}

// SetDetachNetwork gets a reference to the given NullableBool and assigns it to the DetachNetwork field.
func (o *RecoverVmwareVmNewSourceNetworkConfig) SetDetachNetwork(v bool) {
	o.DetachNetwork.Set(&v)
}
// SetDetachNetworkNil sets the value for DetachNetwork to be an explicit nil
func (o *RecoverVmwareVmNewSourceNetworkConfig) SetDetachNetworkNil() {
	o.DetachNetwork.Set(nil)
}

// UnsetDetachNetwork ensures that no value is present for DetachNetwork, not even an explicit nil
func (o *RecoverVmwareVmNewSourceNetworkConfig) UnsetDetachNetwork() {
	o.DetachNetwork.Unset()
}

// GetNewNetworkConfig returns the NewNetworkConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverVmwareVmNewSourceNetworkConfig) GetNewNetworkConfig() RecoverVmwareVmNewNetworkConfig {
	if o == nil || o.NewNetworkConfig.Get() == nil {
		var ret RecoverVmwareVmNewNetworkConfig
		return ret
	}
	return *o.NewNetworkConfig.Get()
}

// GetNewNetworkConfigOk returns a tuple with the NewNetworkConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverVmwareVmNewSourceNetworkConfig) GetNewNetworkConfigOk() (*RecoverVmwareVmNewNetworkConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NewNetworkConfig.Get(), o.NewNetworkConfig.IsSet()
}

// HasNewNetworkConfig returns a boolean if a field has been set.
func (o *RecoverVmwareVmNewSourceNetworkConfig) HasNewNetworkConfig() bool {
	if o != nil && o.NewNetworkConfig.IsSet() {
		return true
	}

	return false
}

// SetNewNetworkConfig gets a reference to the given NullableRecoverVmwareVmNewNetworkConfig and assigns it to the NewNetworkConfig field.
func (o *RecoverVmwareVmNewSourceNetworkConfig) SetNewNetworkConfig(v RecoverVmwareVmNewNetworkConfig) {
	o.NewNetworkConfig.Set(&v)
}
// SetNewNetworkConfigNil sets the value for NewNetworkConfig to be an explicit nil
func (o *RecoverVmwareVmNewSourceNetworkConfig) SetNewNetworkConfigNil() {
	o.NewNetworkConfig.Set(nil)
}

// UnsetNewNetworkConfig ensures that no value is present for NewNetworkConfig, not even an explicit nil
func (o *RecoverVmwareVmNewSourceNetworkConfig) UnsetNewNetworkConfig() {
	o.NewNetworkConfig.Unset()
}

func (o RecoverVmwareVmNewSourceNetworkConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DetachNetwork.IsSet() {
		toSerialize["detachNetwork"] = o.DetachNetwork.Get()
	}
	if o.NewNetworkConfig.IsSet() {
		toSerialize["newNetworkConfig"] = o.NewNetworkConfig.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRecoverVmwareVmNewSourceNetworkConfig struct {
	value *RecoverVmwareVmNewSourceNetworkConfig
	isSet bool
}

func (v NullableRecoverVmwareVmNewSourceNetworkConfig) Get() *RecoverVmwareVmNewSourceNetworkConfig {
	return v.value
}

func (v *NullableRecoverVmwareVmNewSourceNetworkConfig) Set(val *RecoverVmwareVmNewSourceNetworkConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverVmwareVmNewSourceNetworkConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverVmwareVmNewSourceNetworkConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverVmwareVmNewSourceNetworkConfig(val *RecoverVmwareVmNewSourceNetworkConfig) *NullableRecoverVmwareVmNewSourceNetworkConfig {
	return &NullableRecoverVmwareVmNewSourceNetworkConfig{value: val, isSet: true}
}

func (v NullableRecoverVmwareVmNewSourceNetworkConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverVmwareVmNewSourceNetworkConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o RecoverVmwareVmNewSourceNetworkConfig) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}