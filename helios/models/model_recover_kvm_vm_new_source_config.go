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

// RecoverKvmVmNewSourceConfig Specifies the new destination Source configuration where the VMs will be recovered.
type RecoverKvmVmNewSourceConfig struct {
	// Specifies the id of the parent source to recover the VMs.
	Source RecoveryObjectIdentifier `json:"source"`
	// Specifies the resource (KVMH host) to which the restored VM will be attached
	Cluster RecoveryObjectIdentifier `json:"cluster"`
	// Specifies the datacenter where the VM's files should be restored to.
	DataCenter RecoveryObjectIdentifier `json:"dataCenter"`
	// Specifies the Storage Domain where the VM's disk should be restored to.
	StorageDomain RecoveryObjectIdentifier `json:"storageDomain"`
	// Specifies the networking configuration to be applied to the recovered VMs.
	NetworkConfig *RecoverKvmVmNewSourceNetworkConfig `json:"networkConfig,omitempty"`
}

// NewRecoverKvmVmNewSourceConfig instantiates a new RecoverKvmVmNewSourceConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverKvmVmNewSourceConfig(source RecoveryObjectIdentifier, cluster RecoveryObjectIdentifier, dataCenter RecoveryObjectIdentifier, storageDomain RecoveryObjectIdentifier) *RecoverKvmVmNewSourceConfig {
	this := RecoverKvmVmNewSourceConfig{}
	this.Source = source
	this.Cluster = cluster
	this.DataCenter = dataCenter
	this.StorageDomain = storageDomain
	return &this
}

// NewRecoverKvmVmNewSourceConfigWithDefaults instantiates a new RecoverKvmVmNewSourceConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverKvmVmNewSourceConfigWithDefaults() *RecoverKvmVmNewSourceConfig {
	this := RecoverKvmVmNewSourceConfig{}
	return &this
}

// GetSource returns the Source field value
func (o *RecoverKvmVmNewSourceConfig) GetSource() RecoveryObjectIdentifier {
	if o == nil {
		var ret RecoveryObjectIdentifier
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *RecoverKvmVmNewSourceConfig) GetSourceOk() (*RecoveryObjectIdentifier, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *RecoverKvmVmNewSourceConfig) SetSource(v RecoveryObjectIdentifier) {
	o.Source = v
}

// GetCluster returns the Cluster field value
func (o *RecoverKvmVmNewSourceConfig) GetCluster() RecoveryObjectIdentifier {
	if o == nil {
		var ret RecoveryObjectIdentifier
		return ret
	}

	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *RecoverKvmVmNewSourceConfig) GetClusterOk() (*RecoveryObjectIdentifier, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value
func (o *RecoverKvmVmNewSourceConfig) SetCluster(v RecoveryObjectIdentifier) {
	o.Cluster = v
}

// GetDataCenter returns the DataCenter field value
func (o *RecoverKvmVmNewSourceConfig) GetDataCenter() RecoveryObjectIdentifier {
	if o == nil {
		var ret RecoveryObjectIdentifier
		return ret
	}

	return o.DataCenter
}

// GetDataCenterOk returns a tuple with the DataCenter field value
// and a boolean to check if the value has been set.
func (o *RecoverKvmVmNewSourceConfig) GetDataCenterOk() (*RecoveryObjectIdentifier, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DataCenter, true
}

// SetDataCenter sets field value
func (o *RecoverKvmVmNewSourceConfig) SetDataCenter(v RecoveryObjectIdentifier) {
	o.DataCenter = v
}

// GetStorageDomain returns the StorageDomain field value
func (o *RecoverKvmVmNewSourceConfig) GetStorageDomain() RecoveryObjectIdentifier {
	if o == nil {
		var ret RecoveryObjectIdentifier
		return ret
	}

	return o.StorageDomain
}

// GetStorageDomainOk returns a tuple with the StorageDomain field value
// and a boolean to check if the value has been set.
func (o *RecoverKvmVmNewSourceConfig) GetStorageDomainOk() (*RecoveryObjectIdentifier, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StorageDomain, true
}

// SetStorageDomain sets field value
func (o *RecoverKvmVmNewSourceConfig) SetStorageDomain(v RecoveryObjectIdentifier) {
	o.StorageDomain = v
}

// GetNetworkConfig returns the NetworkConfig field value if set, zero value otherwise.
func (o *RecoverKvmVmNewSourceConfig) GetNetworkConfig() RecoverKvmVmNewSourceNetworkConfig {
	if o == nil || o.NetworkConfig == nil {
		var ret RecoverKvmVmNewSourceNetworkConfig
		return ret
	}
	return *o.NetworkConfig
}

// GetNetworkConfigOk returns a tuple with the NetworkConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoverKvmVmNewSourceConfig) GetNetworkConfigOk() (*RecoverKvmVmNewSourceNetworkConfig, bool) {
	if o == nil || o.NetworkConfig == nil {
		return nil, false
	}
	return o.NetworkConfig, true
}

// HasNetworkConfig returns a boolean if a field has been set.
func (o *RecoverKvmVmNewSourceConfig) HasNetworkConfig() bool {
	if o != nil && o.NetworkConfig != nil {
		return true
	}

	return false
}

// SetNetworkConfig gets a reference to the given RecoverKvmVmNewSourceNetworkConfig and assigns it to the NetworkConfig field.
func (o *RecoverKvmVmNewSourceConfig) SetNetworkConfig(v RecoverKvmVmNewSourceNetworkConfig) {
	o.NetworkConfig = &v
}

func (o RecoverKvmVmNewSourceConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["cluster"] = o.Cluster
	}
	if true {
		toSerialize["dataCenter"] = o.DataCenter
	}
	if true {
		toSerialize["storageDomain"] = o.StorageDomain
	}
	if o.NetworkConfig != nil {
		toSerialize["networkConfig"] = o.NetworkConfig
	}
	return json.Marshal(toSerialize)
}

type NullableRecoverKvmVmNewSourceConfig struct {
	value *RecoverKvmVmNewSourceConfig
	isSet bool
}

func (v NullableRecoverKvmVmNewSourceConfig) Get() *RecoverKvmVmNewSourceConfig {
	return v.value
}

func (v *NullableRecoverKvmVmNewSourceConfig) Set(val *RecoverKvmVmNewSourceConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverKvmVmNewSourceConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverKvmVmNewSourceConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverKvmVmNewSourceConfig(val *RecoverKvmVmNewSourceConfig) *NullableRecoverKvmVmNewSourceConfig {
	return &NullableRecoverKvmVmNewSourceConfig{value: val, isSet: true}
}

func (v NullableRecoverKvmVmNewSourceConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverKvmVmNewSourceConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


