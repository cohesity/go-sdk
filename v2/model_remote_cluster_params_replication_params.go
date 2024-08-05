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

// checks if the RemoteClusterParamsReplicationParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoteClusterParamsReplicationParams{}

// RemoteClusterParamsReplicationParams Specifies the replication config for a Remote Cluster. Required when usedForReplication is set to true.
type RemoteClusterParamsReplicationParams struct {
	// Specifies if all endpoints on Remote Cluster are reachable.
	AllEndpointsReachable NullableBool `json:"allEndpointsReachable,omitempty"`
	BandwidthLimit *ReplicationParamsBandwidthLimit `json:"bandwidthLimit,omitempty"`
	// Specifies whether to compress the outbound data when transferring the replication data over the network to the Remote Cluster.
	CompressionEnabled NullableBool `json:"compressionEnabled,omitempty"`
	// Specifies the encryption key used for encrypting the replication data from a local Cluster to a Remote Cluster. If a key is not specified, replication traffic encryption is disabled. When Snapshots are replicated from a local Cluster to a Remote Cluster, the encryption key specified on the local Cluster must be the same as the key specified on the Remote Cluster.
	EncryptionKey NullableString `json:"encryptionKey,omitempty"`
	// Specifies a list of Storage Domain pairs.
	StorageDomainPairs []StorageDomainPair `json:"storageDomainPairs,omitempty"`
}

// NewRemoteClusterParamsReplicationParams instantiates a new RemoteClusterParamsReplicationParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteClusterParamsReplicationParams() *RemoteClusterParamsReplicationParams {
	this := RemoteClusterParamsReplicationParams{}
	var allEndpointsReachable bool = false
	this.AllEndpointsReachable = *NewNullableBool(&allEndpointsReachable)
	var compressionEnabled bool = true
	this.CompressionEnabled = *NewNullableBool(&compressionEnabled)
	return &this
}

// NewRemoteClusterParamsReplicationParamsWithDefaults instantiates a new RemoteClusterParamsReplicationParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteClusterParamsReplicationParamsWithDefaults() *RemoteClusterParamsReplicationParams {
	this := RemoteClusterParamsReplicationParams{}
	var allEndpointsReachable bool = false
	this.AllEndpointsReachable = *NewNullableBool(&allEndpointsReachable)
	var compressionEnabled bool = true
	this.CompressionEnabled = *NewNullableBool(&compressionEnabled)
	return &this
}

// GetAllEndpointsReachable returns the AllEndpointsReachable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteClusterParamsReplicationParams) GetAllEndpointsReachable() bool {
	if o == nil || IsNil(o.AllEndpointsReachable.Get()) {
		var ret bool
		return ret
	}
	return *o.AllEndpointsReachable.Get()
}

// GetAllEndpointsReachableOk returns a tuple with the AllEndpointsReachable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteClusterParamsReplicationParams) GetAllEndpointsReachableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllEndpointsReachable.Get(), o.AllEndpointsReachable.IsSet()
}

// HasAllEndpointsReachable returns a boolean if a field has been set.
func (o *RemoteClusterParamsReplicationParams) HasAllEndpointsReachable() bool {
	if o != nil && o.AllEndpointsReachable.IsSet() {
		return true
	}

	return false
}

// SetAllEndpointsReachable gets a reference to the given NullableBool and assigns it to the AllEndpointsReachable field.
func (o *RemoteClusterParamsReplicationParams) SetAllEndpointsReachable(v bool) {
	o.AllEndpointsReachable.Set(&v)
}
// SetAllEndpointsReachableNil sets the value for AllEndpointsReachable to be an explicit nil
func (o *RemoteClusterParamsReplicationParams) SetAllEndpointsReachableNil() {
	o.AllEndpointsReachable.Set(nil)
}

// UnsetAllEndpointsReachable ensures that no value is present for AllEndpointsReachable, not even an explicit nil
func (o *RemoteClusterParamsReplicationParams) UnsetAllEndpointsReachable() {
	o.AllEndpointsReachable.Unset()
}

// GetBandwidthLimit returns the BandwidthLimit field value if set, zero value otherwise.
func (o *RemoteClusterParamsReplicationParams) GetBandwidthLimit() ReplicationParamsBandwidthLimit {
	if o == nil || IsNil(o.BandwidthLimit) {
		var ret ReplicationParamsBandwidthLimit
		return ret
	}
	return *o.BandwidthLimit
}

// GetBandwidthLimitOk returns a tuple with the BandwidthLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteClusterParamsReplicationParams) GetBandwidthLimitOk() (*ReplicationParamsBandwidthLimit, bool) {
	if o == nil || IsNil(o.BandwidthLimit) {
		return nil, false
	}
	return o.BandwidthLimit, true
}

// HasBandwidthLimit returns a boolean if a field has been set.
func (o *RemoteClusterParamsReplicationParams) HasBandwidthLimit() bool {
	if o != nil && !IsNil(o.BandwidthLimit) {
		return true
	}

	return false
}

// SetBandwidthLimit gets a reference to the given ReplicationParamsBandwidthLimit and assigns it to the BandwidthLimit field.
func (o *RemoteClusterParamsReplicationParams) SetBandwidthLimit(v ReplicationParamsBandwidthLimit) {
	o.BandwidthLimit = &v
}

// GetCompressionEnabled returns the CompressionEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteClusterParamsReplicationParams) GetCompressionEnabled() bool {
	if o == nil || IsNil(o.CompressionEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.CompressionEnabled.Get()
}

// GetCompressionEnabledOk returns a tuple with the CompressionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteClusterParamsReplicationParams) GetCompressionEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompressionEnabled.Get(), o.CompressionEnabled.IsSet()
}

// HasCompressionEnabled returns a boolean if a field has been set.
func (o *RemoteClusterParamsReplicationParams) HasCompressionEnabled() bool {
	if o != nil && o.CompressionEnabled.IsSet() {
		return true
	}

	return false
}

// SetCompressionEnabled gets a reference to the given NullableBool and assigns it to the CompressionEnabled field.
func (o *RemoteClusterParamsReplicationParams) SetCompressionEnabled(v bool) {
	o.CompressionEnabled.Set(&v)
}
// SetCompressionEnabledNil sets the value for CompressionEnabled to be an explicit nil
func (o *RemoteClusterParamsReplicationParams) SetCompressionEnabledNil() {
	o.CompressionEnabled.Set(nil)
}

// UnsetCompressionEnabled ensures that no value is present for CompressionEnabled, not even an explicit nil
func (o *RemoteClusterParamsReplicationParams) UnsetCompressionEnabled() {
	o.CompressionEnabled.Unset()
}

// GetEncryptionKey returns the EncryptionKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteClusterParamsReplicationParams) GetEncryptionKey() string {
	if o == nil || IsNil(o.EncryptionKey.Get()) {
		var ret string
		return ret
	}
	return *o.EncryptionKey.Get()
}

// GetEncryptionKeyOk returns a tuple with the EncryptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteClusterParamsReplicationParams) GetEncryptionKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EncryptionKey.Get(), o.EncryptionKey.IsSet()
}

// HasEncryptionKey returns a boolean if a field has been set.
func (o *RemoteClusterParamsReplicationParams) HasEncryptionKey() bool {
	if o != nil && o.EncryptionKey.IsSet() {
		return true
	}

	return false
}

// SetEncryptionKey gets a reference to the given NullableString and assigns it to the EncryptionKey field.
func (o *RemoteClusterParamsReplicationParams) SetEncryptionKey(v string) {
	o.EncryptionKey.Set(&v)
}
// SetEncryptionKeyNil sets the value for EncryptionKey to be an explicit nil
func (o *RemoteClusterParamsReplicationParams) SetEncryptionKeyNil() {
	o.EncryptionKey.Set(nil)
}

// UnsetEncryptionKey ensures that no value is present for EncryptionKey, not even an explicit nil
func (o *RemoteClusterParamsReplicationParams) UnsetEncryptionKey() {
	o.EncryptionKey.Unset()
}

// GetStorageDomainPairs returns the StorageDomainPairs field value if set, zero value otherwise.
func (o *RemoteClusterParamsReplicationParams) GetStorageDomainPairs() []StorageDomainPair {
	if o == nil || IsNil(o.StorageDomainPairs) {
		var ret []StorageDomainPair
		return ret
	}
	return o.StorageDomainPairs
}

// GetStorageDomainPairsOk returns a tuple with the StorageDomainPairs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteClusterParamsReplicationParams) GetStorageDomainPairsOk() ([]StorageDomainPair, bool) {
	if o == nil || IsNil(o.StorageDomainPairs) {
		return nil, false
	}
	return o.StorageDomainPairs, true
}

// HasStorageDomainPairs returns a boolean if a field has been set.
func (o *RemoteClusterParamsReplicationParams) HasStorageDomainPairs() bool {
	if o != nil && !IsNil(o.StorageDomainPairs) {
		return true
	}

	return false
}

// SetStorageDomainPairs gets a reference to the given []StorageDomainPair and assigns it to the StorageDomainPairs field.
func (o *RemoteClusterParamsReplicationParams) SetStorageDomainPairs(v []StorageDomainPair) {
	o.StorageDomainPairs = v
}

func (o RemoteClusterParamsReplicationParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoteClusterParamsReplicationParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AllEndpointsReachable.IsSet() {
		toSerialize["allEndpointsReachable"] = o.AllEndpointsReachable.Get()
	}
	if !IsNil(o.BandwidthLimit) {
		toSerialize["bandwidthLimit"] = o.BandwidthLimit
	}
	if o.CompressionEnabled.IsSet() {
		toSerialize["compressionEnabled"] = o.CompressionEnabled.Get()
	}
	if o.EncryptionKey.IsSet() {
		toSerialize["encryptionKey"] = o.EncryptionKey.Get()
	}
	if !IsNil(o.StorageDomainPairs) {
		toSerialize["storageDomainPairs"] = o.StorageDomainPairs
	}
	return toSerialize, nil
}

type NullableRemoteClusterParamsReplicationParams struct {
	value *RemoteClusterParamsReplicationParams
	isSet bool
}

func (v NullableRemoteClusterParamsReplicationParams) Get() *RemoteClusterParamsReplicationParams {
	return v.value
}

func (v *NullableRemoteClusterParamsReplicationParams) Set(val *RemoteClusterParamsReplicationParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteClusterParamsReplicationParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteClusterParamsReplicationParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteClusterParamsReplicationParams(val *RemoteClusterParamsReplicationParams) *NullableRemoteClusterParamsReplicationParams {
	return &NullableRemoteClusterParamsReplicationParams{value: val, isSet: true}
}

func (v NullableRemoteClusterParamsReplicationParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteClusterParamsReplicationParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


