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

// RecoverKvmVmNewSourceNetworkConfig Specifies the network config parameters to be applied for KVM VMs if recovering to new Source.
type RecoverKvmVmNewSourceNetworkConfig struct {
	// If this is set to true, then the network will be detached from the recovered VMs. All the other networking parameters set will be ignored if set to true. Default value is false.
	DetachNetwork NullableBool `json:"detachNetwork,omitempty"`
	// Specifies the new network configuration of the Kvm recovery.
	NewNetworkConfig *RecoverKvmVmNewNetworkConfig `json:"newNetworkConfig,omitempty"`
}

// NewRecoverKvmVmNewSourceNetworkConfig instantiates a new RecoverKvmVmNewSourceNetworkConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverKvmVmNewSourceNetworkConfig() *RecoverKvmVmNewSourceNetworkConfig {
	this := RecoverKvmVmNewSourceNetworkConfig{}
	return &this
}

// NewRecoverKvmVmNewSourceNetworkConfigWithDefaults instantiates a new RecoverKvmVmNewSourceNetworkConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverKvmVmNewSourceNetworkConfigWithDefaults() *RecoverKvmVmNewSourceNetworkConfig {
	this := RecoverKvmVmNewSourceNetworkConfig{}
	return &this
}

// GetDetachNetwork returns the DetachNetwork field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverKvmVmNewSourceNetworkConfig) GetDetachNetwork() bool {
	if o == nil || o.DetachNetwork.Get() == nil {
		var ret bool
		return ret
	}
	return *o.DetachNetwork.Get()
}

// GetDetachNetworkOk returns a tuple with the DetachNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverKvmVmNewSourceNetworkConfig) GetDetachNetworkOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DetachNetwork.Get(), o.DetachNetwork.IsSet()
}

// HasDetachNetwork returns a boolean if a field has been set.
func (o *RecoverKvmVmNewSourceNetworkConfig) HasDetachNetwork() bool {
	if o != nil && o.DetachNetwork.IsSet() {
		return true
	}

	return false
}

// SetDetachNetwork gets a reference to the given NullableBool and assigns it to the DetachNetwork field.
func (o *RecoverKvmVmNewSourceNetworkConfig) SetDetachNetwork(v bool) {
	o.DetachNetwork.Set(&v)
}
// SetDetachNetworkNil sets the value for DetachNetwork to be an explicit nil
func (o *RecoverKvmVmNewSourceNetworkConfig) SetDetachNetworkNil() {
	o.DetachNetwork.Set(nil)
}

// UnsetDetachNetwork ensures that no value is present for DetachNetwork, not even an explicit nil
func (o *RecoverKvmVmNewSourceNetworkConfig) UnsetDetachNetwork() {
	o.DetachNetwork.Unset()
}

// GetNewNetworkConfig returns the NewNetworkConfig field value if set, zero value otherwise.
func (o *RecoverKvmVmNewSourceNetworkConfig) GetNewNetworkConfig() RecoverKvmVmNewNetworkConfig {
	if o == nil || o.NewNetworkConfig == nil {
		var ret RecoverKvmVmNewNetworkConfig
		return ret
	}
	return *o.NewNetworkConfig
}

// GetNewNetworkConfigOk returns a tuple with the NewNetworkConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoverKvmVmNewSourceNetworkConfig) GetNewNetworkConfigOk() (*RecoverKvmVmNewNetworkConfig, bool) {
	if o == nil || o.NewNetworkConfig == nil {
		return nil, false
	}
	return o.NewNetworkConfig, true
}

// HasNewNetworkConfig returns a boolean if a field has been set.
func (o *RecoverKvmVmNewSourceNetworkConfig) HasNewNetworkConfig() bool {
	if o != nil && o.NewNetworkConfig != nil {
		return true
	}

	return false
}

// SetNewNetworkConfig gets a reference to the given RecoverKvmVmNewNetworkConfig and assigns it to the NewNetworkConfig field.
func (o *RecoverKvmVmNewSourceNetworkConfig) SetNewNetworkConfig(v RecoverKvmVmNewNetworkConfig) {
	o.NewNetworkConfig = &v
}

func (o RecoverKvmVmNewSourceNetworkConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DetachNetwork.IsSet() {
		toSerialize["detachNetwork"] = o.DetachNetwork.Get()
	}
	if o.NewNetworkConfig != nil {
		toSerialize["newNetworkConfig"] = o.NewNetworkConfig
	}
	return json.Marshal(toSerialize)
}

type NullableRecoverKvmVmNewSourceNetworkConfig struct {
	value *RecoverKvmVmNewSourceNetworkConfig
	isSet bool
}

func (v NullableRecoverKvmVmNewSourceNetworkConfig) Get() *RecoverKvmVmNewSourceNetworkConfig {
	return v.value
}

func (v *NullableRecoverKvmVmNewSourceNetworkConfig) Set(val *RecoverKvmVmNewSourceNetworkConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverKvmVmNewSourceNetworkConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverKvmVmNewSourceNetworkConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverKvmVmNewSourceNetworkConfig(val *RecoverKvmVmNewSourceNetworkConfig) *NullableRecoverKvmVmNewSourceNetworkConfig {
	return &NullableRecoverKvmVmNewSourceNetworkConfig{value: val, isSet: true}
}

func (v NullableRecoverKvmVmNewSourceNetworkConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverKvmVmNewSourceNetworkConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o RecoverKvmVmNewSourceNetworkConfig) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}