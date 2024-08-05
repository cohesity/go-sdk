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

// checks if the RecoverVmwareVAppVCDSourceConfigNetworkConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoverVmwareVAppVCDSourceConfigNetworkConfig{}

// RecoverVmwareVAppVCDSourceConfigNetworkConfig Specifies the networking configuration to be applied to the recovered VMs.
type RecoverVmwareVAppVCDSourceConfigNetworkConfig struct {
	// If this is set to true, then the network will be detached from the recovered VMs. All the other networking parameters set will be ignored if set to true. Default value is false.
	DetachNetwork NullableBool `json:"detachNetwork,omitempty"`
	NewNetworkConfig NullableRecoverVmwareVmNewSourceNetworkConfigNewNetworkConfig `json:"newNetworkConfig,omitempty"`
}

// NewRecoverVmwareVAppVCDSourceConfigNetworkConfig instantiates a new RecoverVmwareVAppVCDSourceConfigNetworkConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverVmwareVAppVCDSourceConfigNetworkConfig() *RecoverVmwareVAppVCDSourceConfigNetworkConfig {
	this := RecoverVmwareVAppVCDSourceConfigNetworkConfig{}
	return &this
}

// NewRecoverVmwareVAppVCDSourceConfigNetworkConfigWithDefaults instantiates a new RecoverVmwareVAppVCDSourceConfigNetworkConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverVmwareVAppVCDSourceConfigNetworkConfigWithDefaults() *RecoverVmwareVAppVCDSourceConfigNetworkConfig {
	this := RecoverVmwareVAppVCDSourceConfigNetworkConfig{}
	return &this
}

// GetDetachNetwork returns the DetachNetwork field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverVmwareVAppVCDSourceConfigNetworkConfig) GetDetachNetwork() bool {
	if o == nil || IsNil(o.DetachNetwork.Get()) {
		var ret bool
		return ret
	}
	return *o.DetachNetwork.Get()
}

// GetDetachNetworkOk returns a tuple with the DetachNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverVmwareVAppVCDSourceConfigNetworkConfig) GetDetachNetworkOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DetachNetwork.Get(), o.DetachNetwork.IsSet()
}

// HasDetachNetwork returns a boolean if a field has been set.
func (o *RecoverVmwareVAppVCDSourceConfigNetworkConfig) HasDetachNetwork() bool {
	if o != nil && o.DetachNetwork.IsSet() {
		return true
	}

	return false
}

// SetDetachNetwork gets a reference to the given NullableBool and assigns it to the DetachNetwork field.
func (o *RecoverVmwareVAppVCDSourceConfigNetworkConfig) SetDetachNetwork(v bool) {
	o.DetachNetwork.Set(&v)
}
// SetDetachNetworkNil sets the value for DetachNetwork to be an explicit nil
func (o *RecoverVmwareVAppVCDSourceConfigNetworkConfig) SetDetachNetworkNil() {
	o.DetachNetwork.Set(nil)
}

// UnsetDetachNetwork ensures that no value is present for DetachNetwork, not even an explicit nil
func (o *RecoverVmwareVAppVCDSourceConfigNetworkConfig) UnsetDetachNetwork() {
	o.DetachNetwork.Unset()
}

// GetNewNetworkConfig returns the NewNetworkConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverVmwareVAppVCDSourceConfigNetworkConfig) GetNewNetworkConfig() RecoverVmwareVmNewSourceNetworkConfigNewNetworkConfig {
	if o == nil || IsNil(o.NewNetworkConfig.Get()) {
		var ret RecoverVmwareVmNewSourceNetworkConfigNewNetworkConfig
		return ret
	}
	return *o.NewNetworkConfig.Get()
}

// GetNewNetworkConfigOk returns a tuple with the NewNetworkConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverVmwareVAppVCDSourceConfigNetworkConfig) GetNewNetworkConfigOk() (*RecoverVmwareVmNewSourceNetworkConfigNewNetworkConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.NewNetworkConfig.Get(), o.NewNetworkConfig.IsSet()
}

// HasNewNetworkConfig returns a boolean if a field has been set.
func (o *RecoverVmwareVAppVCDSourceConfigNetworkConfig) HasNewNetworkConfig() bool {
	if o != nil && o.NewNetworkConfig.IsSet() {
		return true
	}

	return false
}

// SetNewNetworkConfig gets a reference to the given NullableRecoverVmwareVmNewSourceNetworkConfigNewNetworkConfig and assigns it to the NewNetworkConfig field.
func (o *RecoverVmwareVAppVCDSourceConfigNetworkConfig) SetNewNetworkConfig(v RecoverVmwareVmNewSourceNetworkConfigNewNetworkConfig) {
	o.NewNetworkConfig.Set(&v)
}
// SetNewNetworkConfigNil sets the value for NewNetworkConfig to be an explicit nil
func (o *RecoverVmwareVAppVCDSourceConfigNetworkConfig) SetNewNetworkConfigNil() {
	o.NewNetworkConfig.Set(nil)
}

// UnsetNewNetworkConfig ensures that no value is present for NewNetworkConfig, not even an explicit nil
func (o *RecoverVmwareVAppVCDSourceConfigNetworkConfig) UnsetNewNetworkConfig() {
	o.NewNetworkConfig.Unset()
}

func (o RecoverVmwareVAppVCDSourceConfigNetworkConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoverVmwareVAppVCDSourceConfigNetworkConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DetachNetwork.IsSet() {
		toSerialize["detachNetwork"] = o.DetachNetwork.Get()
	}
	if o.NewNetworkConfig.IsSet() {
		toSerialize["newNetworkConfig"] = o.NewNetworkConfig.Get()
	}
	return toSerialize, nil
}

type NullableRecoverVmwareVAppVCDSourceConfigNetworkConfig struct {
	value *RecoverVmwareVAppVCDSourceConfigNetworkConfig
	isSet bool
}

func (v NullableRecoverVmwareVAppVCDSourceConfigNetworkConfig) Get() *RecoverVmwareVAppVCDSourceConfigNetworkConfig {
	return v.value
}

func (v *NullableRecoverVmwareVAppVCDSourceConfigNetworkConfig) Set(val *RecoverVmwareVAppVCDSourceConfigNetworkConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverVmwareVAppVCDSourceConfigNetworkConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverVmwareVAppVCDSourceConfigNetworkConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverVmwareVAppVCDSourceConfigNetworkConfig(val *RecoverVmwareVAppVCDSourceConfigNetworkConfig) *NullableRecoverVmwareVAppVCDSourceConfigNetworkConfig {
	return &NullableRecoverVmwareVAppVCDSourceConfigNetworkConfig{value: val, isSet: true}
}

func (v NullableRecoverVmwareVAppVCDSourceConfigNetworkConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverVmwareVAppVCDSourceConfigNetworkConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


