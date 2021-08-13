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

// RecoverAcropolisVmOriginalSourceNetworkConfig Specifies the network config parameters to be applied for Acropolis VMs if recovering to original Source.
type RecoverAcropolisVmOriginalSourceNetworkConfig struct {
	// If this is set to true, then the network will be detached from the recovered VMs. All the other networking parameters set will be ignored if set to true. Default value is false.
	DetachNetwork NullableBool `json:"detachNetwork,omitempty"`
}

// NewRecoverAcropolisVmOriginalSourceNetworkConfig instantiates a new RecoverAcropolisVmOriginalSourceNetworkConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverAcropolisVmOriginalSourceNetworkConfig() *RecoverAcropolisVmOriginalSourceNetworkConfig {
	this := RecoverAcropolisVmOriginalSourceNetworkConfig{}
	return &this
}

// NewRecoverAcropolisVmOriginalSourceNetworkConfigWithDefaults instantiates a new RecoverAcropolisVmOriginalSourceNetworkConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverAcropolisVmOriginalSourceNetworkConfigWithDefaults() *RecoverAcropolisVmOriginalSourceNetworkConfig {
	this := RecoverAcropolisVmOriginalSourceNetworkConfig{}
	return &this
}

// GetDetachNetwork returns the DetachNetwork field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverAcropolisVmOriginalSourceNetworkConfig) GetDetachNetwork() bool {
	if o == nil || o.DetachNetwork.Get() == nil {
		var ret bool
		return ret
	}
	return *o.DetachNetwork.Get()
}

// GetDetachNetworkOk returns a tuple with the DetachNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverAcropolisVmOriginalSourceNetworkConfig) GetDetachNetworkOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DetachNetwork.Get(), o.DetachNetwork.IsSet()
}

// HasDetachNetwork returns a boolean if a field has been set.
func (o *RecoverAcropolisVmOriginalSourceNetworkConfig) HasDetachNetwork() bool {
	if o != nil && o.DetachNetwork.IsSet() {
		return true
	}

	return false
}

// SetDetachNetwork gets a reference to the given NullableBool and assigns it to the DetachNetwork field.
func (o *RecoverAcropolisVmOriginalSourceNetworkConfig) SetDetachNetwork(v bool) {
	o.DetachNetwork.Set(&v)
}
// SetDetachNetworkNil sets the value for DetachNetwork to be an explicit nil
func (o *RecoverAcropolisVmOriginalSourceNetworkConfig) SetDetachNetworkNil() {
	o.DetachNetwork.Set(nil)
}

// UnsetDetachNetwork ensures that no value is present for DetachNetwork, not even an explicit nil
func (o *RecoverAcropolisVmOriginalSourceNetworkConfig) UnsetDetachNetwork() {
	o.DetachNetwork.Unset()
}

func (o RecoverAcropolisVmOriginalSourceNetworkConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DetachNetwork.IsSet() {
		toSerialize["detachNetwork"] = o.DetachNetwork.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRecoverAcropolisVmOriginalSourceNetworkConfig struct {
	value *RecoverAcropolisVmOriginalSourceNetworkConfig
	isSet bool
}

func (v NullableRecoverAcropolisVmOriginalSourceNetworkConfig) Get() *RecoverAcropolisVmOriginalSourceNetworkConfig {
	return v.value
}

func (v *NullableRecoverAcropolisVmOriginalSourceNetworkConfig) Set(val *RecoverAcropolisVmOriginalSourceNetworkConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverAcropolisVmOriginalSourceNetworkConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverAcropolisVmOriginalSourceNetworkConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverAcropolisVmOriginalSourceNetworkConfig(val *RecoverAcropolisVmOriginalSourceNetworkConfig) *NullableRecoverAcropolisVmOriginalSourceNetworkConfig {
	return &NullableRecoverAcropolisVmOriginalSourceNetworkConfig{value: val, isSet: true}
}

func (v NullableRecoverAcropolisVmOriginalSourceNetworkConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverAcropolisVmOriginalSourceNetworkConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


