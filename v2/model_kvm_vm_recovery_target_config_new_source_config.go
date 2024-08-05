/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the KvmVmRecoveryTargetConfigNewSourceConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KvmVmRecoveryTargetConfigNewSourceConfig{}

// KvmVmRecoveryTargetConfigNewSourceConfig Specifies the new destination Source configuration parameters where the VMs will be recovered. This is mandatory if recoverToNewSource is set to true.
type KvmVmRecoveryTargetConfigNewSourceConfig struct {
	Cluster RecoverKvmVmNewSourceConfigCluster `json:"cluster"`
	DataCenter RecoverKvmVmNewSourceConfigDataCenter `json:"dataCenter"`
	NetworkConfig *RecoverKvmVmNewSourceConfigNetworkConfig `json:"networkConfig,omitempty"`
	Source RecoverKvmVmNewSourceConfigSource `json:"source"`
	StorageDomain RecoverKvmVmNewSourceConfigStorageDomain `json:"storageDomain"`
}

type _KvmVmRecoveryTargetConfigNewSourceConfig KvmVmRecoveryTargetConfigNewSourceConfig

// NewKvmVmRecoveryTargetConfigNewSourceConfig instantiates a new KvmVmRecoveryTargetConfigNewSourceConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKvmVmRecoveryTargetConfigNewSourceConfig(cluster RecoverKvmVmNewSourceConfigCluster, dataCenter RecoverKvmVmNewSourceConfigDataCenter, source RecoverKvmVmNewSourceConfigSource, storageDomain RecoverKvmVmNewSourceConfigStorageDomain) *KvmVmRecoveryTargetConfigNewSourceConfig {
	this := KvmVmRecoveryTargetConfigNewSourceConfig{}
	this.Cluster = cluster
	this.DataCenter = dataCenter
	this.Source = source
	this.StorageDomain = storageDomain
	return &this
}

// NewKvmVmRecoveryTargetConfigNewSourceConfigWithDefaults instantiates a new KvmVmRecoveryTargetConfigNewSourceConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvmVmRecoveryTargetConfigNewSourceConfigWithDefaults() *KvmVmRecoveryTargetConfigNewSourceConfig {
	this := KvmVmRecoveryTargetConfigNewSourceConfig{}
	return &this
}

// GetCluster returns the Cluster field value
func (o *KvmVmRecoveryTargetConfigNewSourceConfig) GetCluster() RecoverKvmVmNewSourceConfigCluster {
	if o == nil {
		var ret RecoverKvmVmNewSourceConfigCluster
		return ret
	}

	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *KvmVmRecoveryTargetConfigNewSourceConfig) GetClusterOk() (*RecoverKvmVmNewSourceConfigCluster, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value
func (o *KvmVmRecoveryTargetConfigNewSourceConfig) SetCluster(v RecoverKvmVmNewSourceConfigCluster) {
	o.Cluster = v
}

// GetDataCenter returns the DataCenter field value
func (o *KvmVmRecoveryTargetConfigNewSourceConfig) GetDataCenter() RecoverKvmVmNewSourceConfigDataCenter {
	if o == nil {
		var ret RecoverKvmVmNewSourceConfigDataCenter
		return ret
	}

	return o.DataCenter
}

// GetDataCenterOk returns a tuple with the DataCenter field value
// and a boolean to check if the value has been set.
func (o *KvmVmRecoveryTargetConfigNewSourceConfig) GetDataCenterOk() (*RecoverKvmVmNewSourceConfigDataCenter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataCenter, true
}

// SetDataCenter sets field value
func (o *KvmVmRecoveryTargetConfigNewSourceConfig) SetDataCenter(v RecoverKvmVmNewSourceConfigDataCenter) {
	o.DataCenter = v
}

// GetNetworkConfig returns the NetworkConfig field value if set, zero value otherwise.
func (o *KvmVmRecoveryTargetConfigNewSourceConfig) GetNetworkConfig() RecoverKvmVmNewSourceConfigNetworkConfig {
	if o == nil || IsNil(o.NetworkConfig) {
		var ret RecoverKvmVmNewSourceConfigNetworkConfig
		return ret
	}
	return *o.NetworkConfig
}

// GetNetworkConfigOk returns a tuple with the NetworkConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmVmRecoveryTargetConfigNewSourceConfig) GetNetworkConfigOk() (*RecoverKvmVmNewSourceConfigNetworkConfig, bool) {
	if o == nil || IsNil(o.NetworkConfig) {
		return nil, false
	}
	return o.NetworkConfig, true
}

// HasNetworkConfig returns a boolean if a field has been set.
func (o *KvmVmRecoveryTargetConfigNewSourceConfig) HasNetworkConfig() bool {
	if o != nil && !IsNil(o.NetworkConfig) {
		return true
	}

	return false
}

// SetNetworkConfig gets a reference to the given RecoverKvmVmNewSourceConfigNetworkConfig and assigns it to the NetworkConfig field.
func (o *KvmVmRecoveryTargetConfigNewSourceConfig) SetNetworkConfig(v RecoverKvmVmNewSourceConfigNetworkConfig) {
	o.NetworkConfig = &v
}

// GetSource returns the Source field value
func (o *KvmVmRecoveryTargetConfigNewSourceConfig) GetSource() RecoverKvmVmNewSourceConfigSource {
	if o == nil {
		var ret RecoverKvmVmNewSourceConfigSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *KvmVmRecoveryTargetConfigNewSourceConfig) GetSourceOk() (*RecoverKvmVmNewSourceConfigSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *KvmVmRecoveryTargetConfigNewSourceConfig) SetSource(v RecoverKvmVmNewSourceConfigSource) {
	o.Source = v
}

// GetStorageDomain returns the StorageDomain field value
func (o *KvmVmRecoveryTargetConfigNewSourceConfig) GetStorageDomain() RecoverKvmVmNewSourceConfigStorageDomain {
	if o == nil {
		var ret RecoverKvmVmNewSourceConfigStorageDomain
		return ret
	}

	return o.StorageDomain
}

// GetStorageDomainOk returns a tuple with the StorageDomain field value
// and a boolean to check if the value has been set.
func (o *KvmVmRecoveryTargetConfigNewSourceConfig) GetStorageDomainOk() (*RecoverKvmVmNewSourceConfigStorageDomain, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageDomain, true
}

// SetStorageDomain sets field value
func (o *KvmVmRecoveryTargetConfigNewSourceConfig) SetStorageDomain(v RecoverKvmVmNewSourceConfigStorageDomain) {
	o.StorageDomain = v
}

func (o KvmVmRecoveryTargetConfigNewSourceConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KvmVmRecoveryTargetConfigNewSourceConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cluster"] = o.Cluster
	toSerialize["dataCenter"] = o.DataCenter
	if !IsNil(o.NetworkConfig) {
		toSerialize["networkConfig"] = o.NetworkConfig
	}
	toSerialize["source"] = o.Source
	toSerialize["storageDomain"] = o.StorageDomain
	return toSerialize, nil
}

func (o *KvmVmRecoveryTargetConfigNewSourceConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cluster",
		"dataCenter",
		"source",
		"storageDomain",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varKvmVmRecoveryTargetConfigNewSourceConfig := _KvmVmRecoveryTargetConfigNewSourceConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKvmVmRecoveryTargetConfigNewSourceConfig)

	if err != nil {
		return err
	}

	*o = KvmVmRecoveryTargetConfigNewSourceConfig(varKvmVmRecoveryTargetConfigNewSourceConfig)

	return err
}

type NullableKvmVmRecoveryTargetConfigNewSourceConfig struct {
	value *KvmVmRecoveryTargetConfigNewSourceConfig
	isSet bool
}

func (v NullableKvmVmRecoveryTargetConfigNewSourceConfig) Get() *KvmVmRecoveryTargetConfigNewSourceConfig {
	return v.value
}

func (v *NullableKvmVmRecoveryTargetConfigNewSourceConfig) Set(val *KvmVmRecoveryTargetConfigNewSourceConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableKvmVmRecoveryTargetConfigNewSourceConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableKvmVmRecoveryTargetConfigNewSourceConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKvmVmRecoveryTargetConfigNewSourceConfig(val *KvmVmRecoveryTargetConfigNewSourceConfig) *NullableKvmVmRecoveryTargetConfigNewSourceConfig {
	return &NullableKvmVmRecoveryTargetConfigNewSourceConfig{value: val, isSet: true}
}

func (v NullableKvmVmRecoveryTargetConfigNewSourceConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKvmVmRecoveryTargetConfigNewSourceConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


