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

// RecoverHyperVVmSCVMMSourceConfig Specifies the new destination Source configuration where the VMs will be recovered.
type RecoverHyperVVmSCVMMSourceConfig struct {
	// Specifies the id of the parent source to recover the VMs.
	Source NullableRecoveryObjectIdentifier `json:"source"`
	// Specifies the HyperV host where the recovered VMs will be attached. For standalone host targets, the host must be the same as the source.
	Host NullableRecoveryObjectIdentifier `json:"host"`
	// Specifies the datastore object where the VMs' files should be recovered to.
	Volume NullableRecoveryObjectIdentifier `json:"volume"`
	// Specifies the networking configuration to be applied to the recovered VMs.
	NetworkConfig NullableRecoverHyperVVmNewSourceNetworkConfig `json:"networkConfig,omitempty"`
}

// NewRecoverHyperVVmSCVMMSourceConfig instantiates a new RecoverHyperVVmSCVMMSourceConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverHyperVVmSCVMMSourceConfig(source NullableRecoveryObjectIdentifier, host NullableRecoveryObjectIdentifier, volume NullableRecoveryObjectIdentifier) *RecoverHyperVVmSCVMMSourceConfig {
	this := RecoverHyperVVmSCVMMSourceConfig{}
	this.Source = source
	this.Host = host
	this.Volume = volume
	return &this
}

// NewRecoverHyperVVmSCVMMSourceConfigWithDefaults instantiates a new RecoverHyperVVmSCVMMSourceConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverHyperVVmSCVMMSourceConfigWithDefaults() *RecoverHyperVVmSCVMMSourceConfig {
	this := RecoverHyperVVmSCVMMSourceConfig{}
	return &this
}

// GetSource returns the Source field value
// If the value is explicit nil, the zero value for RecoveryObjectIdentifier will be returned
func (o *RecoverHyperVVmSCVMMSourceConfig) GetSource() RecoveryObjectIdentifier {
	if o == nil || o.Source.Get() == nil {
		var ret RecoveryObjectIdentifier
		return ret
	}

	return *o.Source.Get()
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverHyperVVmSCVMMSourceConfig) GetSourceOk() (*RecoveryObjectIdentifier, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Source.Get(), o.Source.IsSet()
}

// SetSource sets field value
func (o *RecoverHyperVVmSCVMMSourceConfig) SetSource(v RecoveryObjectIdentifier) {
	o.Source.Set(&v)
}

// GetHost returns the Host field value
// If the value is explicit nil, the zero value for RecoveryObjectIdentifier will be returned
func (o *RecoverHyperVVmSCVMMSourceConfig) GetHost() RecoveryObjectIdentifier {
	if o == nil || o.Host.Get() == nil {
		var ret RecoveryObjectIdentifier
		return ret
	}

	return *o.Host.Get()
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverHyperVVmSCVMMSourceConfig) GetHostOk() (*RecoveryObjectIdentifier, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Host.Get(), o.Host.IsSet()
}

// SetHost sets field value
func (o *RecoverHyperVVmSCVMMSourceConfig) SetHost(v RecoveryObjectIdentifier) {
	o.Host.Set(&v)
}

// GetVolume returns the Volume field value
// If the value is explicit nil, the zero value for RecoveryObjectIdentifier will be returned
func (o *RecoverHyperVVmSCVMMSourceConfig) GetVolume() RecoveryObjectIdentifier {
	if o == nil || o.Volume.Get() == nil {
		var ret RecoveryObjectIdentifier
		return ret
	}

	return *o.Volume.Get()
}

// GetVolumeOk returns a tuple with the Volume field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverHyperVVmSCVMMSourceConfig) GetVolumeOk() (*RecoveryObjectIdentifier, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Volume.Get(), o.Volume.IsSet()
}

// SetVolume sets field value
func (o *RecoverHyperVVmSCVMMSourceConfig) SetVolume(v RecoveryObjectIdentifier) {
	o.Volume.Set(&v)
}

// GetNetworkConfig returns the NetworkConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverHyperVVmSCVMMSourceConfig) GetNetworkConfig() RecoverHyperVVmNewSourceNetworkConfig {
	if o == nil || o.NetworkConfig.Get() == nil {
		var ret RecoverHyperVVmNewSourceNetworkConfig
		return ret
	}
	return *o.NetworkConfig.Get()
}

// GetNetworkConfigOk returns a tuple with the NetworkConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverHyperVVmSCVMMSourceConfig) GetNetworkConfigOk() (*RecoverHyperVVmNewSourceNetworkConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NetworkConfig.Get(), o.NetworkConfig.IsSet()
}

// HasNetworkConfig returns a boolean if a field has been set.
func (o *RecoverHyperVVmSCVMMSourceConfig) HasNetworkConfig() bool {
	if o != nil && o.NetworkConfig.IsSet() {
		return true
	}

	return false
}

// SetNetworkConfig gets a reference to the given NullableRecoverHyperVVmNewSourceNetworkConfig and assigns it to the NetworkConfig field.
func (o *RecoverHyperVVmSCVMMSourceConfig) SetNetworkConfig(v RecoverHyperVVmNewSourceNetworkConfig) {
	o.NetworkConfig.Set(&v)
}
// SetNetworkConfigNil sets the value for NetworkConfig to be an explicit nil
func (o *RecoverHyperVVmSCVMMSourceConfig) SetNetworkConfigNil() {
	o.NetworkConfig.Set(nil)
}

// UnsetNetworkConfig ensures that no value is present for NetworkConfig, not even an explicit nil
func (o *RecoverHyperVVmSCVMMSourceConfig) UnsetNetworkConfig() {
	o.NetworkConfig.Unset()
}

func (o RecoverHyperVVmSCVMMSourceConfig) MarshalJSON() ([]byte, error) {
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

type NullableRecoverHyperVVmSCVMMSourceConfig struct {
	value *RecoverHyperVVmSCVMMSourceConfig
	isSet bool
}

func (v NullableRecoverHyperVVmSCVMMSourceConfig) Get() *RecoverHyperVVmSCVMMSourceConfig {
	return v.value
}

func (v *NullableRecoverHyperVVmSCVMMSourceConfig) Set(val *RecoverHyperVVmSCVMMSourceConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverHyperVVmSCVMMSourceConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverHyperVVmSCVMMSourceConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverHyperVVmSCVMMSourceConfig(val *RecoverHyperVVmSCVMMSourceConfig) *NullableRecoverHyperVVmSCVMMSourceConfig {
	return &NullableRecoverHyperVVmSCVMMSourceConfig{value: val, isSet: true}
}

func (v NullableRecoverHyperVVmSCVMMSourceConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverHyperVVmSCVMMSourceConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


