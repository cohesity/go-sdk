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

// RecoverHyperVVmStandaloneClusterSourceConfig Specifies the new destination Source configuration where the VMs will be recovered.
type RecoverHyperVVmStandaloneClusterSourceConfig struct {
	// Specifies the id of the parent source to recover the VMs.
	Source NullableRecoveryObjectIdentifier `json:"source"`
	// Specifies the HyperV host where the recovered VMs will be attached. For standalone host targets, the host must be the same as the source.
	Host NullableRecoveryObjectIdentifier `json:"host"`
	// Specifies the datastore object where the VMs' files should be recovered to.
	Volume NullableRecoveryObjectIdentifier `json:"volume"`
	// Specifies the networking configuration to be applied to the recovered VMs.
	NetworkConfig NullableRecoverHyperVVmNewSourceNetworkConfig `json:"networkConfig,omitempty"`
}

// NewRecoverHyperVVmStandaloneClusterSourceConfig instantiates a new RecoverHyperVVmStandaloneClusterSourceConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverHyperVVmStandaloneClusterSourceConfig(source NullableRecoveryObjectIdentifier, host NullableRecoveryObjectIdentifier, volume NullableRecoveryObjectIdentifier) *RecoverHyperVVmStandaloneClusterSourceConfig {
	this := RecoverHyperVVmStandaloneClusterSourceConfig{}
	this.Source = source
	this.Host = host
	this.Volume = volume
	return &this
}

// NewRecoverHyperVVmStandaloneClusterSourceConfigWithDefaults instantiates a new RecoverHyperVVmStandaloneClusterSourceConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverHyperVVmStandaloneClusterSourceConfigWithDefaults() *RecoverHyperVVmStandaloneClusterSourceConfig {
	this := RecoverHyperVVmStandaloneClusterSourceConfig{}
	return &this
}

// GetSource returns the Source field value
// If the value is explicit nil, the zero value for RecoveryObjectIdentifier will be returned
func (o *RecoverHyperVVmStandaloneClusterSourceConfig) GetSource() RecoveryObjectIdentifier {
	if o == nil || o.Source.Get() == nil {
		var ret RecoveryObjectIdentifier
		return ret
	}

	return *o.Source.Get()
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverHyperVVmStandaloneClusterSourceConfig) GetSourceOk() (*RecoveryObjectIdentifier, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Source.Get(), o.Source.IsSet()
}

// SetSource sets field value
func (o *RecoverHyperVVmStandaloneClusterSourceConfig) SetSource(v RecoveryObjectIdentifier) {
	o.Source.Set(&v)
}

// GetHost returns the Host field value
// If the value is explicit nil, the zero value for RecoveryObjectIdentifier will be returned
func (o *RecoverHyperVVmStandaloneClusterSourceConfig) GetHost() RecoveryObjectIdentifier {
	if o == nil || o.Host.Get() == nil {
		var ret RecoveryObjectIdentifier
		return ret
	}

	return *o.Host.Get()
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverHyperVVmStandaloneClusterSourceConfig) GetHostOk() (*RecoveryObjectIdentifier, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Host.Get(), o.Host.IsSet()
}

// SetHost sets field value
func (o *RecoverHyperVVmStandaloneClusterSourceConfig) SetHost(v RecoveryObjectIdentifier) {
	o.Host.Set(&v)
}

// GetVolume returns the Volume field value
// If the value is explicit nil, the zero value for RecoveryObjectIdentifier will be returned
func (o *RecoverHyperVVmStandaloneClusterSourceConfig) GetVolume() RecoveryObjectIdentifier {
	if o == nil || o.Volume.Get() == nil {
		var ret RecoveryObjectIdentifier
		return ret
	}

	return *o.Volume.Get()
}

// GetVolumeOk returns a tuple with the Volume field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverHyperVVmStandaloneClusterSourceConfig) GetVolumeOk() (*RecoveryObjectIdentifier, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Volume.Get(), o.Volume.IsSet()
}

// SetVolume sets field value
func (o *RecoverHyperVVmStandaloneClusterSourceConfig) SetVolume(v RecoveryObjectIdentifier) {
	o.Volume.Set(&v)
}

// GetNetworkConfig returns the NetworkConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverHyperVVmStandaloneClusterSourceConfig) GetNetworkConfig() RecoverHyperVVmNewSourceNetworkConfig {
	if o == nil || o.NetworkConfig.Get() == nil {
		var ret RecoverHyperVVmNewSourceNetworkConfig
		return ret
	}
	return *o.NetworkConfig.Get()
}

// GetNetworkConfigOk returns a tuple with the NetworkConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverHyperVVmStandaloneClusterSourceConfig) GetNetworkConfigOk() (*RecoverHyperVVmNewSourceNetworkConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NetworkConfig.Get(), o.NetworkConfig.IsSet()
}

// HasNetworkConfig returns a boolean if a field has been set.
func (o *RecoverHyperVVmStandaloneClusterSourceConfig) HasNetworkConfig() bool {
	if o != nil && o.NetworkConfig.IsSet() {
		return true
	}

	return false
}

// SetNetworkConfig gets a reference to the given NullableRecoverHyperVVmNewSourceNetworkConfig and assigns it to the NetworkConfig field.
func (o *RecoverHyperVVmStandaloneClusterSourceConfig) SetNetworkConfig(v RecoverHyperVVmNewSourceNetworkConfig) {
	o.NetworkConfig.Set(&v)
}
// SetNetworkConfigNil sets the value for NetworkConfig to be an explicit nil
func (o *RecoverHyperVVmStandaloneClusterSourceConfig) SetNetworkConfigNil() {
	o.NetworkConfig.Set(nil)
}

// UnsetNetworkConfig ensures that no value is present for NetworkConfig, not even an explicit nil
func (o *RecoverHyperVVmStandaloneClusterSourceConfig) UnsetNetworkConfig() {
	o.NetworkConfig.Unset()
}

func (o RecoverHyperVVmStandaloneClusterSourceConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["source"] = o.Source.Get()
	}
	if true {
		toSerialize["host"] = o.Host.Get()
	}
	if true {
		toSerialize["volume"] = o.Volume.Get()
	}
	if o.NetworkConfig.IsSet() {
		toSerialize["networkConfig"] = o.NetworkConfig.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRecoverHyperVVmStandaloneClusterSourceConfig struct {
	value *RecoverHyperVVmStandaloneClusterSourceConfig
	isSet bool
}

func (v NullableRecoverHyperVVmStandaloneClusterSourceConfig) Get() *RecoverHyperVVmStandaloneClusterSourceConfig {
	return v.value
}

func (v *NullableRecoverHyperVVmStandaloneClusterSourceConfig) Set(val *RecoverHyperVVmStandaloneClusterSourceConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverHyperVVmStandaloneClusterSourceConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverHyperVVmStandaloneClusterSourceConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverHyperVVmStandaloneClusterSourceConfig(val *RecoverHyperVVmStandaloneClusterSourceConfig) *NullableRecoverHyperVVmStandaloneClusterSourceConfig {
	return &NullableRecoverHyperVVmStandaloneClusterSourceConfig{value: val, isSet: true}
}

func (v NullableRecoverHyperVVmStandaloneClusterSourceConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverHyperVVmStandaloneClusterSourceConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o RecoverHyperVVmStandaloneClusterSourceConfig) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}