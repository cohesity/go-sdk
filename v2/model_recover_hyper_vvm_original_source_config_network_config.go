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

// checks if the RecoverHyperVVmOriginalSourceConfigNetworkConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoverHyperVVmOriginalSourceConfigNetworkConfig{}

// RecoverHyperVVmOriginalSourceConfigNetworkConfig Specifies the networking configuration to be applied to the recovered VMs.
type RecoverHyperVVmOriginalSourceConfigNetworkConfig struct {
	// If this is set to true, then the network will be detached from the recovered VMs. Default value is false.
	DetachNetwork NullableBool `json:"detachNetwork,omitempty"`
}

// NewRecoverHyperVVmOriginalSourceConfigNetworkConfig instantiates a new RecoverHyperVVmOriginalSourceConfigNetworkConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverHyperVVmOriginalSourceConfigNetworkConfig() *RecoverHyperVVmOriginalSourceConfigNetworkConfig {
	this := RecoverHyperVVmOriginalSourceConfigNetworkConfig{}
	return &this
}

// NewRecoverHyperVVmOriginalSourceConfigNetworkConfigWithDefaults instantiates a new RecoverHyperVVmOriginalSourceConfigNetworkConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverHyperVVmOriginalSourceConfigNetworkConfigWithDefaults() *RecoverHyperVVmOriginalSourceConfigNetworkConfig {
	this := RecoverHyperVVmOriginalSourceConfigNetworkConfig{}
	return &this
}

// GetDetachNetwork returns the DetachNetwork field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverHyperVVmOriginalSourceConfigNetworkConfig) GetDetachNetwork() bool {
	if o == nil || IsNil(o.DetachNetwork.Get()) {
		var ret bool
		return ret
	}
	return *o.DetachNetwork.Get()
}

// GetDetachNetworkOk returns a tuple with the DetachNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverHyperVVmOriginalSourceConfigNetworkConfig) GetDetachNetworkOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DetachNetwork.Get(), o.DetachNetwork.IsSet()
}

// HasDetachNetwork returns a boolean if a field has been set.
func (o *RecoverHyperVVmOriginalSourceConfigNetworkConfig) HasDetachNetwork() bool {
	if o != nil && o.DetachNetwork.IsSet() {
		return true
	}

	return false
}

// SetDetachNetwork gets a reference to the given NullableBool and assigns it to the DetachNetwork field.
func (o *RecoverHyperVVmOriginalSourceConfigNetworkConfig) SetDetachNetwork(v bool) {
	o.DetachNetwork.Set(&v)
}
// SetDetachNetworkNil sets the value for DetachNetwork to be an explicit nil
func (o *RecoverHyperVVmOriginalSourceConfigNetworkConfig) SetDetachNetworkNil() {
	o.DetachNetwork.Set(nil)
}

// UnsetDetachNetwork ensures that no value is present for DetachNetwork, not even an explicit nil
func (o *RecoverHyperVVmOriginalSourceConfigNetworkConfig) UnsetDetachNetwork() {
	o.DetachNetwork.Unset()
}

func (o RecoverHyperVVmOriginalSourceConfigNetworkConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoverHyperVVmOriginalSourceConfigNetworkConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DetachNetwork.IsSet() {
		toSerialize["detachNetwork"] = o.DetachNetwork.Get()
	}
	return toSerialize, nil
}

type NullableRecoverHyperVVmOriginalSourceConfigNetworkConfig struct {
	value *RecoverHyperVVmOriginalSourceConfigNetworkConfig
	isSet bool
}

func (v NullableRecoverHyperVVmOriginalSourceConfigNetworkConfig) Get() *RecoverHyperVVmOriginalSourceConfigNetworkConfig {
	return v.value
}

func (v *NullableRecoverHyperVVmOriginalSourceConfigNetworkConfig) Set(val *RecoverHyperVVmOriginalSourceConfigNetworkConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverHyperVVmOriginalSourceConfigNetworkConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverHyperVVmOriginalSourceConfigNetworkConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverHyperVVmOriginalSourceConfigNetworkConfig(val *RecoverHyperVVmOriginalSourceConfigNetworkConfig) *NullableRecoverHyperVVmOriginalSourceConfigNetworkConfig {
	return &NullableRecoverHyperVVmOriginalSourceConfigNetworkConfig{value: val, isSet: true}
}

func (v NullableRecoverHyperVVmOriginalSourceConfigNetworkConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverHyperVVmOriginalSourceConfigNetworkConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


