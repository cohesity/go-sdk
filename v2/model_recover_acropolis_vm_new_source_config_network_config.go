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

// checks if the RecoverAcropolisVmNewSourceConfigNetworkConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoverAcropolisVmNewSourceConfigNetworkConfig{}

// RecoverAcropolisVmNewSourceConfigNetworkConfig Specifies the networking configuration to be applied to the recovered VMs.
type RecoverAcropolisVmNewSourceConfigNetworkConfig struct {
	// If this is set to true, then the network will be detached from the recovered VMs. All the other networking parameters set will be ignored if set to true. Default value is false.
	DetachNetwork NullableBool `json:"detachNetwork,omitempty"`
	NetworkPortGroup NullableRecoverAcropolisVmNewSourceNetworkConfigNetworkPortGroup `json:"networkPortGroup,omitempty"`
}

// NewRecoverAcropolisVmNewSourceConfigNetworkConfig instantiates a new RecoverAcropolisVmNewSourceConfigNetworkConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverAcropolisVmNewSourceConfigNetworkConfig() *RecoverAcropolisVmNewSourceConfigNetworkConfig {
	this := RecoverAcropolisVmNewSourceConfigNetworkConfig{}
	return &this
}

// NewRecoverAcropolisVmNewSourceConfigNetworkConfigWithDefaults instantiates a new RecoverAcropolisVmNewSourceConfigNetworkConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverAcropolisVmNewSourceConfigNetworkConfigWithDefaults() *RecoverAcropolisVmNewSourceConfigNetworkConfig {
	this := RecoverAcropolisVmNewSourceConfigNetworkConfig{}
	return &this
}

// GetDetachNetwork returns the DetachNetwork field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverAcropolisVmNewSourceConfigNetworkConfig) GetDetachNetwork() bool {
	if o == nil || IsNil(o.DetachNetwork.Get()) {
		var ret bool
		return ret
	}
	return *o.DetachNetwork.Get()
}

// GetDetachNetworkOk returns a tuple with the DetachNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverAcropolisVmNewSourceConfigNetworkConfig) GetDetachNetworkOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DetachNetwork.Get(), o.DetachNetwork.IsSet()
}

// HasDetachNetwork returns a boolean if a field has been set.
func (o *RecoverAcropolisVmNewSourceConfigNetworkConfig) HasDetachNetwork() bool {
	if o != nil && o.DetachNetwork.IsSet() {
		return true
	}

	return false
}

// SetDetachNetwork gets a reference to the given NullableBool and assigns it to the DetachNetwork field.
func (o *RecoverAcropolisVmNewSourceConfigNetworkConfig) SetDetachNetwork(v bool) {
	o.DetachNetwork.Set(&v)
}
// SetDetachNetworkNil sets the value for DetachNetwork to be an explicit nil
func (o *RecoverAcropolisVmNewSourceConfigNetworkConfig) SetDetachNetworkNil() {
	o.DetachNetwork.Set(nil)
}

// UnsetDetachNetwork ensures that no value is present for DetachNetwork, not even an explicit nil
func (o *RecoverAcropolisVmNewSourceConfigNetworkConfig) UnsetDetachNetwork() {
	o.DetachNetwork.Unset()
}

// GetNetworkPortGroup returns the NetworkPortGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverAcropolisVmNewSourceConfigNetworkConfig) GetNetworkPortGroup() RecoverAcropolisVmNewSourceNetworkConfigNetworkPortGroup {
	if o == nil || IsNil(o.NetworkPortGroup.Get()) {
		var ret RecoverAcropolisVmNewSourceNetworkConfigNetworkPortGroup
		return ret
	}
	return *o.NetworkPortGroup.Get()
}

// GetNetworkPortGroupOk returns a tuple with the NetworkPortGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverAcropolisVmNewSourceConfigNetworkConfig) GetNetworkPortGroupOk() (*RecoverAcropolisVmNewSourceNetworkConfigNetworkPortGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetworkPortGroup.Get(), o.NetworkPortGroup.IsSet()
}

// HasNetworkPortGroup returns a boolean if a field has been set.
func (o *RecoverAcropolisVmNewSourceConfigNetworkConfig) HasNetworkPortGroup() bool {
	if o != nil && o.NetworkPortGroup.IsSet() {
		return true
	}

	return false
}

// SetNetworkPortGroup gets a reference to the given NullableRecoverAcropolisVmNewSourceNetworkConfigNetworkPortGroup and assigns it to the NetworkPortGroup field.
func (o *RecoverAcropolisVmNewSourceConfigNetworkConfig) SetNetworkPortGroup(v RecoverAcropolisVmNewSourceNetworkConfigNetworkPortGroup) {
	o.NetworkPortGroup.Set(&v)
}
// SetNetworkPortGroupNil sets the value for NetworkPortGroup to be an explicit nil
func (o *RecoverAcropolisVmNewSourceConfigNetworkConfig) SetNetworkPortGroupNil() {
	o.NetworkPortGroup.Set(nil)
}

// UnsetNetworkPortGroup ensures that no value is present for NetworkPortGroup, not even an explicit nil
func (o *RecoverAcropolisVmNewSourceConfigNetworkConfig) UnsetNetworkPortGroup() {
	o.NetworkPortGroup.Unset()
}

func (o RecoverAcropolisVmNewSourceConfigNetworkConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoverAcropolisVmNewSourceConfigNetworkConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DetachNetwork.IsSet() {
		toSerialize["detachNetwork"] = o.DetachNetwork.Get()
	}
	if o.NetworkPortGroup.IsSet() {
		toSerialize["networkPortGroup"] = o.NetworkPortGroup.Get()
	}
	return toSerialize, nil
}

type NullableRecoverAcropolisVmNewSourceConfigNetworkConfig struct {
	value *RecoverAcropolisVmNewSourceConfigNetworkConfig
	isSet bool
}

func (v NullableRecoverAcropolisVmNewSourceConfigNetworkConfig) Get() *RecoverAcropolisVmNewSourceConfigNetworkConfig {
	return v.value
}

func (v *NullableRecoverAcropolisVmNewSourceConfigNetworkConfig) Set(val *RecoverAcropolisVmNewSourceConfigNetworkConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverAcropolisVmNewSourceConfigNetworkConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverAcropolisVmNewSourceConfigNetworkConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverAcropolisVmNewSourceConfigNetworkConfig(val *RecoverAcropolisVmNewSourceConfigNetworkConfig) *NullableRecoverAcropolisVmNewSourceConfigNetworkConfig {
	return &NullableRecoverAcropolisVmNewSourceConfigNetworkConfig{value: val, isSet: true}
}

func (v NullableRecoverAcropolisVmNewSourceConfigNetworkConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverAcropolisVmNewSourceConfigNetworkConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


