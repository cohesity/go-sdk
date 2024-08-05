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

// checks if the UpdateOdpRemoteClusterParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOdpRemoteClusterParams{}

// UpdateOdpRemoteClusterParams Specifies the parameters to create an ODP Remote Cluster config.
type UpdateOdpRemoteClusterParams struct {
	// Specifies if all endpoints on ODP Remote Cluster are reachable.
	AllEndpointsReachable NullableBool `json:"allEndpointsReachable,omitempty"`
	// Specifies if the cluster id is stale and needs to be refreshed.
	ClusterIdStale NullableBool `json:"clusterIdStale,omitempty"`
	// Specifies the ODP Remote Cluster name.
	ClusterName NullableString `json:"clusterName"`
	// Specifies whether to compress the data transferred to ODP Remote Cluster.
	CompressionEnabled NullableBool `json:"compressionEnabled,omitempty"`
	// Specifies the interface group name of the ODP Remote Cluster.
	InterfaceGroupName NullableString `json:"interfaceGroupName,omitempty"`
	// Specifies the key used for encrypting the data transferred to ODP Remote Cluster.
	KeyEncryptionKey NullableString `json:"keyEncryptionKey,omitempty"`
	// Specifies the tenant id for ODP Remote Cluster.
	RemoteTenantId NullableString `json:"remoteTenantId,omitempty"`
	// Specifies a list of Storage Domain pairs.
	StorageDomainPairs []StorageDomainPair `json:"storageDomainPairs,omitempty"`
	// Specifies the tenant id.
	TenantId NullableString `json:"tenantId,omitempty"`
	// Specifies whether to use Bifrost Broker channel for remote connection.
	UseBifrostBrokerChannel NullableBool `json:"useBifrostBrokerChannel,omitempty"`
	// Specifies if the ODP Remote Cluster is used for replication.
	UsedForReplication NullableBool `json:"usedForReplication,omitempty"`
}

type _UpdateOdpRemoteClusterParams UpdateOdpRemoteClusterParams

// NewUpdateOdpRemoteClusterParams instantiates a new UpdateOdpRemoteClusterParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOdpRemoteClusterParams(clusterName NullableString) *UpdateOdpRemoteClusterParams {
	this := UpdateOdpRemoteClusterParams{}
	this.ClusterName = clusterName
	return &this
}

// NewUpdateOdpRemoteClusterParamsWithDefaults instantiates a new UpdateOdpRemoteClusterParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOdpRemoteClusterParamsWithDefaults() *UpdateOdpRemoteClusterParams {
	this := UpdateOdpRemoteClusterParams{}
	return &this
}

// GetAllEndpointsReachable returns the AllEndpointsReachable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOdpRemoteClusterParams) GetAllEndpointsReachable() bool {
	if o == nil || IsNil(o.AllEndpointsReachable.Get()) {
		var ret bool
		return ret
	}
	return *o.AllEndpointsReachable.Get()
}

// GetAllEndpointsReachableOk returns a tuple with the AllEndpointsReachable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOdpRemoteClusterParams) GetAllEndpointsReachableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllEndpointsReachable.Get(), o.AllEndpointsReachable.IsSet()
}

// HasAllEndpointsReachable returns a boolean if a field has been set.
func (o *UpdateOdpRemoteClusterParams) HasAllEndpointsReachable() bool {
	if o != nil && o.AllEndpointsReachable.IsSet() {
		return true
	}

	return false
}

// SetAllEndpointsReachable gets a reference to the given NullableBool and assigns it to the AllEndpointsReachable field.
func (o *UpdateOdpRemoteClusterParams) SetAllEndpointsReachable(v bool) {
	o.AllEndpointsReachable.Set(&v)
}
// SetAllEndpointsReachableNil sets the value for AllEndpointsReachable to be an explicit nil
func (o *UpdateOdpRemoteClusterParams) SetAllEndpointsReachableNil() {
	o.AllEndpointsReachable.Set(nil)
}

// UnsetAllEndpointsReachable ensures that no value is present for AllEndpointsReachable, not even an explicit nil
func (o *UpdateOdpRemoteClusterParams) UnsetAllEndpointsReachable() {
	o.AllEndpointsReachable.Unset()
}

// GetClusterIdStale returns the ClusterIdStale field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOdpRemoteClusterParams) GetClusterIdStale() bool {
	if o == nil || IsNil(o.ClusterIdStale.Get()) {
		var ret bool
		return ret
	}
	return *o.ClusterIdStale.Get()
}

// GetClusterIdStaleOk returns a tuple with the ClusterIdStale field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOdpRemoteClusterParams) GetClusterIdStaleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterIdStale.Get(), o.ClusterIdStale.IsSet()
}

// HasClusterIdStale returns a boolean if a field has been set.
func (o *UpdateOdpRemoteClusterParams) HasClusterIdStale() bool {
	if o != nil && o.ClusterIdStale.IsSet() {
		return true
	}

	return false
}

// SetClusterIdStale gets a reference to the given NullableBool and assigns it to the ClusterIdStale field.
func (o *UpdateOdpRemoteClusterParams) SetClusterIdStale(v bool) {
	o.ClusterIdStale.Set(&v)
}
// SetClusterIdStaleNil sets the value for ClusterIdStale to be an explicit nil
func (o *UpdateOdpRemoteClusterParams) SetClusterIdStaleNil() {
	o.ClusterIdStale.Set(nil)
}

// UnsetClusterIdStale ensures that no value is present for ClusterIdStale, not even an explicit nil
func (o *UpdateOdpRemoteClusterParams) UnsetClusterIdStale() {
	o.ClusterIdStale.Unset()
}

// GetClusterName returns the ClusterName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UpdateOdpRemoteClusterParams) GetClusterName() string {
	if o == nil || o.ClusterName.Get() == nil {
		var ret string
		return ret
	}

	return *o.ClusterName.Get()
}

// GetClusterNameOk returns a tuple with the ClusterName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOdpRemoteClusterParams) GetClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterName.Get(), o.ClusterName.IsSet()
}

// SetClusterName sets field value
func (o *UpdateOdpRemoteClusterParams) SetClusterName(v string) {
	o.ClusterName.Set(&v)
}

// GetCompressionEnabled returns the CompressionEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOdpRemoteClusterParams) GetCompressionEnabled() bool {
	if o == nil || IsNil(o.CompressionEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.CompressionEnabled.Get()
}

// GetCompressionEnabledOk returns a tuple with the CompressionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOdpRemoteClusterParams) GetCompressionEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompressionEnabled.Get(), o.CompressionEnabled.IsSet()
}

// HasCompressionEnabled returns a boolean if a field has been set.
func (o *UpdateOdpRemoteClusterParams) HasCompressionEnabled() bool {
	if o != nil && o.CompressionEnabled.IsSet() {
		return true
	}

	return false
}

// SetCompressionEnabled gets a reference to the given NullableBool and assigns it to the CompressionEnabled field.
func (o *UpdateOdpRemoteClusterParams) SetCompressionEnabled(v bool) {
	o.CompressionEnabled.Set(&v)
}
// SetCompressionEnabledNil sets the value for CompressionEnabled to be an explicit nil
func (o *UpdateOdpRemoteClusterParams) SetCompressionEnabledNil() {
	o.CompressionEnabled.Set(nil)
}

// UnsetCompressionEnabled ensures that no value is present for CompressionEnabled, not even an explicit nil
func (o *UpdateOdpRemoteClusterParams) UnsetCompressionEnabled() {
	o.CompressionEnabled.Unset()
}

// GetInterfaceGroupName returns the InterfaceGroupName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOdpRemoteClusterParams) GetInterfaceGroupName() string {
	if o == nil || IsNil(o.InterfaceGroupName.Get()) {
		var ret string
		return ret
	}
	return *o.InterfaceGroupName.Get()
}

// GetInterfaceGroupNameOk returns a tuple with the InterfaceGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOdpRemoteClusterParams) GetInterfaceGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InterfaceGroupName.Get(), o.InterfaceGroupName.IsSet()
}

// HasInterfaceGroupName returns a boolean if a field has been set.
func (o *UpdateOdpRemoteClusterParams) HasInterfaceGroupName() bool {
	if o != nil && o.InterfaceGroupName.IsSet() {
		return true
	}

	return false
}

// SetInterfaceGroupName gets a reference to the given NullableString and assigns it to the InterfaceGroupName field.
func (o *UpdateOdpRemoteClusterParams) SetInterfaceGroupName(v string) {
	o.InterfaceGroupName.Set(&v)
}
// SetInterfaceGroupNameNil sets the value for InterfaceGroupName to be an explicit nil
func (o *UpdateOdpRemoteClusterParams) SetInterfaceGroupNameNil() {
	o.InterfaceGroupName.Set(nil)
}

// UnsetInterfaceGroupName ensures that no value is present for InterfaceGroupName, not even an explicit nil
func (o *UpdateOdpRemoteClusterParams) UnsetInterfaceGroupName() {
	o.InterfaceGroupName.Unset()
}

// GetKeyEncryptionKey returns the KeyEncryptionKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOdpRemoteClusterParams) GetKeyEncryptionKey() string {
	if o == nil || IsNil(o.KeyEncryptionKey.Get()) {
		var ret string
		return ret
	}
	return *o.KeyEncryptionKey.Get()
}

// GetKeyEncryptionKeyOk returns a tuple with the KeyEncryptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOdpRemoteClusterParams) GetKeyEncryptionKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KeyEncryptionKey.Get(), o.KeyEncryptionKey.IsSet()
}

// HasKeyEncryptionKey returns a boolean if a field has been set.
func (o *UpdateOdpRemoteClusterParams) HasKeyEncryptionKey() bool {
	if o != nil && o.KeyEncryptionKey.IsSet() {
		return true
	}

	return false
}

// SetKeyEncryptionKey gets a reference to the given NullableString and assigns it to the KeyEncryptionKey field.
func (o *UpdateOdpRemoteClusterParams) SetKeyEncryptionKey(v string) {
	o.KeyEncryptionKey.Set(&v)
}
// SetKeyEncryptionKeyNil sets the value for KeyEncryptionKey to be an explicit nil
func (o *UpdateOdpRemoteClusterParams) SetKeyEncryptionKeyNil() {
	o.KeyEncryptionKey.Set(nil)
}

// UnsetKeyEncryptionKey ensures that no value is present for KeyEncryptionKey, not even an explicit nil
func (o *UpdateOdpRemoteClusterParams) UnsetKeyEncryptionKey() {
	o.KeyEncryptionKey.Unset()
}

// GetRemoteTenantId returns the RemoteTenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOdpRemoteClusterParams) GetRemoteTenantId() string {
	if o == nil || IsNil(o.RemoteTenantId.Get()) {
		var ret string
		return ret
	}
	return *o.RemoteTenantId.Get()
}

// GetRemoteTenantIdOk returns a tuple with the RemoteTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOdpRemoteClusterParams) GetRemoteTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemoteTenantId.Get(), o.RemoteTenantId.IsSet()
}

// HasRemoteTenantId returns a boolean if a field has been set.
func (o *UpdateOdpRemoteClusterParams) HasRemoteTenantId() bool {
	if o != nil && o.RemoteTenantId.IsSet() {
		return true
	}

	return false
}

// SetRemoteTenantId gets a reference to the given NullableString and assigns it to the RemoteTenantId field.
func (o *UpdateOdpRemoteClusterParams) SetRemoteTenantId(v string) {
	o.RemoteTenantId.Set(&v)
}
// SetRemoteTenantIdNil sets the value for RemoteTenantId to be an explicit nil
func (o *UpdateOdpRemoteClusterParams) SetRemoteTenantIdNil() {
	o.RemoteTenantId.Set(nil)
}

// UnsetRemoteTenantId ensures that no value is present for RemoteTenantId, not even an explicit nil
func (o *UpdateOdpRemoteClusterParams) UnsetRemoteTenantId() {
	o.RemoteTenantId.Unset()
}

// GetStorageDomainPairs returns the StorageDomainPairs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOdpRemoteClusterParams) GetStorageDomainPairs() []StorageDomainPair {
	if o == nil {
		var ret []StorageDomainPair
		return ret
	}
	return o.StorageDomainPairs
}

// GetStorageDomainPairsOk returns a tuple with the StorageDomainPairs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOdpRemoteClusterParams) GetStorageDomainPairsOk() ([]StorageDomainPair, bool) {
	if o == nil || IsNil(o.StorageDomainPairs) {
		return nil, false
	}
	return o.StorageDomainPairs, true
}

// HasStorageDomainPairs returns a boolean if a field has been set.
func (o *UpdateOdpRemoteClusterParams) HasStorageDomainPairs() bool {
	if o != nil && !IsNil(o.StorageDomainPairs) {
		return true
	}

	return false
}

// SetStorageDomainPairs gets a reference to the given []StorageDomainPair and assigns it to the StorageDomainPairs field.
func (o *UpdateOdpRemoteClusterParams) SetStorageDomainPairs(v []StorageDomainPair) {
	o.StorageDomainPairs = v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOdpRemoteClusterParams) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOdpRemoteClusterParams) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *UpdateOdpRemoteClusterParams) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *UpdateOdpRemoteClusterParams) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *UpdateOdpRemoteClusterParams) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *UpdateOdpRemoteClusterParams) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetUseBifrostBrokerChannel returns the UseBifrostBrokerChannel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOdpRemoteClusterParams) GetUseBifrostBrokerChannel() bool {
	if o == nil || IsNil(o.UseBifrostBrokerChannel.Get()) {
		var ret bool
		return ret
	}
	return *o.UseBifrostBrokerChannel.Get()
}

// GetUseBifrostBrokerChannelOk returns a tuple with the UseBifrostBrokerChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOdpRemoteClusterParams) GetUseBifrostBrokerChannelOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.UseBifrostBrokerChannel.Get(), o.UseBifrostBrokerChannel.IsSet()
}

// HasUseBifrostBrokerChannel returns a boolean if a field has been set.
func (o *UpdateOdpRemoteClusterParams) HasUseBifrostBrokerChannel() bool {
	if o != nil && o.UseBifrostBrokerChannel.IsSet() {
		return true
	}

	return false
}

// SetUseBifrostBrokerChannel gets a reference to the given NullableBool and assigns it to the UseBifrostBrokerChannel field.
func (o *UpdateOdpRemoteClusterParams) SetUseBifrostBrokerChannel(v bool) {
	o.UseBifrostBrokerChannel.Set(&v)
}
// SetUseBifrostBrokerChannelNil sets the value for UseBifrostBrokerChannel to be an explicit nil
func (o *UpdateOdpRemoteClusterParams) SetUseBifrostBrokerChannelNil() {
	o.UseBifrostBrokerChannel.Set(nil)
}

// UnsetUseBifrostBrokerChannel ensures that no value is present for UseBifrostBrokerChannel, not even an explicit nil
func (o *UpdateOdpRemoteClusterParams) UnsetUseBifrostBrokerChannel() {
	o.UseBifrostBrokerChannel.Unset()
}

// GetUsedForReplication returns the UsedForReplication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOdpRemoteClusterParams) GetUsedForReplication() bool {
	if o == nil || IsNil(o.UsedForReplication.Get()) {
		var ret bool
		return ret
	}
	return *o.UsedForReplication.Get()
}

// GetUsedForReplicationOk returns a tuple with the UsedForReplication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOdpRemoteClusterParams) GetUsedForReplicationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.UsedForReplication.Get(), o.UsedForReplication.IsSet()
}

// HasUsedForReplication returns a boolean if a field has been set.
func (o *UpdateOdpRemoteClusterParams) HasUsedForReplication() bool {
	if o != nil && o.UsedForReplication.IsSet() {
		return true
	}

	return false
}

// SetUsedForReplication gets a reference to the given NullableBool and assigns it to the UsedForReplication field.
func (o *UpdateOdpRemoteClusterParams) SetUsedForReplication(v bool) {
	o.UsedForReplication.Set(&v)
}
// SetUsedForReplicationNil sets the value for UsedForReplication to be an explicit nil
func (o *UpdateOdpRemoteClusterParams) SetUsedForReplicationNil() {
	o.UsedForReplication.Set(nil)
}

// UnsetUsedForReplication ensures that no value is present for UsedForReplication, not even an explicit nil
func (o *UpdateOdpRemoteClusterParams) UnsetUsedForReplication() {
	o.UsedForReplication.Unset()
}

func (o UpdateOdpRemoteClusterParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOdpRemoteClusterParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AllEndpointsReachable.IsSet() {
		toSerialize["allEndpointsReachable"] = o.AllEndpointsReachable.Get()
	}
	if o.ClusterIdStale.IsSet() {
		toSerialize["clusterIdStale"] = o.ClusterIdStale.Get()
	}
	toSerialize["clusterName"] = o.ClusterName.Get()
	if o.CompressionEnabled.IsSet() {
		toSerialize["compressionEnabled"] = o.CompressionEnabled.Get()
	}
	if o.InterfaceGroupName.IsSet() {
		toSerialize["interfaceGroupName"] = o.InterfaceGroupName.Get()
	}
	if o.KeyEncryptionKey.IsSet() {
		toSerialize["keyEncryptionKey"] = o.KeyEncryptionKey.Get()
	}
	if o.RemoteTenantId.IsSet() {
		toSerialize["remoteTenantId"] = o.RemoteTenantId.Get()
	}
	if o.StorageDomainPairs != nil {
		toSerialize["storageDomainPairs"] = o.StorageDomainPairs
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.UseBifrostBrokerChannel.IsSet() {
		toSerialize["useBifrostBrokerChannel"] = o.UseBifrostBrokerChannel.Get()
	}
	if o.UsedForReplication.IsSet() {
		toSerialize["usedForReplication"] = o.UsedForReplication.Get()
	}
	return toSerialize, nil
}

func (o *UpdateOdpRemoteClusterParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"clusterName",
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

	varUpdateOdpRemoteClusterParams := _UpdateOdpRemoteClusterParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateOdpRemoteClusterParams)

	if err != nil {
		return err
	}

	*o = UpdateOdpRemoteClusterParams(varUpdateOdpRemoteClusterParams)

	return err
}

type NullableUpdateOdpRemoteClusterParams struct {
	value *UpdateOdpRemoteClusterParams
	isSet bool
}

func (v NullableUpdateOdpRemoteClusterParams) Get() *UpdateOdpRemoteClusterParams {
	return v.value
}

func (v *NullableUpdateOdpRemoteClusterParams) Set(val *UpdateOdpRemoteClusterParams) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOdpRemoteClusterParams) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOdpRemoteClusterParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOdpRemoteClusterParams(val *UpdateOdpRemoteClusterParams) *NullableUpdateOdpRemoteClusterParams {
	return &NullableUpdateOdpRemoteClusterParams{value: val, isSet: true}
}

func (v NullableUpdateOdpRemoteClusterParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOdpRemoteClusterParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


