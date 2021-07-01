/*
 * Cohesity REST API
 *
 * This API list provides operations for interfacing with the Cohesity Cluster.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cohesitysdk

import (
	"encoding/json"
)

// RestoredObjectNetworkConfigProto struct for RestoredObjectNetworkConfigProto
type RestoredObjectNetworkConfigProto struct {
	// If this is set to true, then the network will be detached from the recovered or cloned VMs. NOTE: If this is set to true, then all the following fields will be ignored.
	DetachNetwork NullableBool `json:"detachNetwork,omitempty"`
	// This can be set to true to indicate that the attached network should be left in disabled state.
	DisableNetwork NullableBool `json:"disableNetwork,omitempty"`
	// The network mappings to be applied to the target object.
	Mappings []NetworkMappingProto `json:"mappings,omitempty"`
	NetworkEntity *EntityProto `json:"networkEntity,omitempty"`
	// If this is true and we are attaching to a new network entity, then the VM's MAC address will be preserved on the new network.
	PreserveMacAddressOnNewNetwork NullableBool `json:"preserveMacAddressOnNewNetwork,omitempty"`
	VnicEntity *EntityProto `json:"vnicEntity,omitempty"`
}

// NewRestoredObjectNetworkConfigProto instantiates a new RestoredObjectNetworkConfigProto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestoredObjectNetworkConfigProto() *RestoredObjectNetworkConfigProto {
	this := RestoredObjectNetworkConfigProto{}
	return &this
}

// NewRestoredObjectNetworkConfigProtoWithDefaults instantiates a new RestoredObjectNetworkConfigProto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestoredObjectNetworkConfigProtoWithDefaults() *RestoredObjectNetworkConfigProto {
	this := RestoredObjectNetworkConfigProto{}
	return &this
}

// GetDetachNetwork returns the DetachNetwork field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestoredObjectNetworkConfigProto) GetDetachNetwork() bool {
	if o == nil || o.DetachNetwork.Get() == nil {
		var ret bool
		return ret
	}
	return *o.DetachNetwork.Get()
}

// GetDetachNetworkOk returns a tuple with the DetachNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestoredObjectNetworkConfigProto) GetDetachNetworkOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DetachNetwork.Get(), o.DetachNetwork.IsSet()
}

// HasDetachNetwork returns a boolean if a field has been set.
func (o *RestoredObjectNetworkConfigProto) HasDetachNetwork() bool {
	if o != nil && o.DetachNetwork.IsSet() {
		return true
	}

	return false
}

// SetDetachNetwork gets a reference to the given NullableBool and assigns it to the DetachNetwork field.
func (o *RestoredObjectNetworkConfigProto) SetDetachNetwork(v bool) {
	o.DetachNetwork.Set(&v)
}
// SetDetachNetworkNil sets the value for DetachNetwork to be an explicit nil
func (o *RestoredObjectNetworkConfigProto) SetDetachNetworkNil() {
	o.DetachNetwork.Set(nil)
}

// UnsetDetachNetwork ensures that no value is present for DetachNetwork, not even an explicit nil
func (o *RestoredObjectNetworkConfigProto) UnsetDetachNetwork() {
	o.DetachNetwork.Unset()
}

// GetDisableNetwork returns the DisableNetwork field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestoredObjectNetworkConfigProto) GetDisableNetwork() bool {
	if o == nil || o.DisableNetwork.Get() == nil {
		var ret bool
		return ret
	}
	return *o.DisableNetwork.Get()
}

// GetDisableNetworkOk returns a tuple with the DisableNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestoredObjectNetworkConfigProto) GetDisableNetworkOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisableNetwork.Get(), o.DisableNetwork.IsSet()
}

// HasDisableNetwork returns a boolean if a field has been set.
func (o *RestoredObjectNetworkConfigProto) HasDisableNetwork() bool {
	if o != nil && o.DisableNetwork.IsSet() {
		return true
	}

	return false
}

// SetDisableNetwork gets a reference to the given NullableBool and assigns it to the DisableNetwork field.
func (o *RestoredObjectNetworkConfigProto) SetDisableNetwork(v bool) {
	o.DisableNetwork.Set(&v)
}
// SetDisableNetworkNil sets the value for DisableNetwork to be an explicit nil
func (o *RestoredObjectNetworkConfigProto) SetDisableNetworkNil() {
	o.DisableNetwork.Set(nil)
}

// UnsetDisableNetwork ensures that no value is present for DisableNetwork, not even an explicit nil
func (o *RestoredObjectNetworkConfigProto) UnsetDisableNetwork() {
	o.DisableNetwork.Unset()
}

// GetMappings returns the Mappings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestoredObjectNetworkConfigProto) GetMappings() []NetworkMappingProto {
	if o == nil  {
		var ret []NetworkMappingProto
		return ret
	}
	return o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestoredObjectNetworkConfigProto) GetMappingsOk() (*[]NetworkMappingProto, bool) {
	if o == nil || o.Mappings == nil {
		return nil, false
	}
	return &o.Mappings, true
}

// HasMappings returns a boolean if a field has been set.
func (o *RestoredObjectNetworkConfigProto) HasMappings() bool {
	if o != nil && o.Mappings != nil {
		return true
	}

	return false
}

// SetMappings gets a reference to the given []NetworkMappingProto and assigns it to the Mappings field.
func (o *RestoredObjectNetworkConfigProto) SetMappings(v []NetworkMappingProto) {
	o.Mappings = v
}

// GetNetworkEntity returns the NetworkEntity field value if set, zero value otherwise.
func (o *RestoredObjectNetworkConfigProto) GetNetworkEntity() EntityProto {
	if o == nil || o.NetworkEntity == nil {
		var ret EntityProto
		return ret
	}
	return *o.NetworkEntity
}

// GetNetworkEntityOk returns a tuple with the NetworkEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoredObjectNetworkConfigProto) GetNetworkEntityOk() (*EntityProto, bool) {
	if o == nil || o.NetworkEntity == nil {
		return nil, false
	}
	return o.NetworkEntity, true
}

// HasNetworkEntity returns a boolean if a field has been set.
func (o *RestoredObjectNetworkConfigProto) HasNetworkEntity() bool {
	if o != nil && o.NetworkEntity != nil {
		return true
	}

	return false
}

// SetNetworkEntity gets a reference to the given EntityProto and assigns it to the NetworkEntity field.
func (o *RestoredObjectNetworkConfigProto) SetNetworkEntity(v EntityProto) {
	o.NetworkEntity = &v
}

// GetPreserveMacAddressOnNewNetwork returns the PreserveMacAddressOnNewNetwork field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestoredObjectNetworkConfigProto) GetPreserveMacAddressOnNewNetwork() bool {
	if o == nil || o.PreserveMacAddressOnNewNetwork.Get() == nil {
		var ret bool
		return ret
	}
	return *o.PreserveMacAddressOnNewNetwork.Get()
}

// GetPreserveMacAddressOnNewNetworkOk returns a tuple with the PreserveMacAddressOnNewNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestoredObjectNetworkConfigProto) GetPreserveMacAddressOnNewNetworkOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PreserveMacAddressOnNewNetwork.Get(), o.PreserveMacAddressOnNewNetwork.IsSet()
}

// HasPreserveMacAddressOnNewNetwork returns a boolean if a field has been set.
func (o *RestoredObjectNetworkConfigProto) HasPreserveMacAddressOnNewNetwork() bool {
	if o != nil && o.PreserveMacAddressOnNewNetwork.IsSet() {
		return true
	}

	return false
}

// SetPreserveMacAddressOnNewNetwork gets a reference to the given NullableBool and assigns it to the PreserveMacAddressOnNewNetwork field.
func (o *RestoredObjectNetworkConfigProto) SetPreserveMacAddressOnNewNetwork(v bool) {
	o.PreserveMacAddressOnNewNetwork.Set(&v)
}
// SetPreserveMacAddressOnNewNetworkNil sets the value for PreserveMacAddressOnNewNetwork to be an explicit nil
func (o *RestoredObjectNetworkConfigProto) SetPreserveMacAddressOnNewNetworkNil() {
	o.PreserveMacAddressOnNewNetwork.Set(nil)
}

// UnsetPreserveMacAddressOnNewNetwork ensures that no value is present for PreserveMacAddressOnNewNetwork, not even an explicit nil
func (o *RestoredObjectNetworkConfigProto) UnsetPreserveMacAddressOnNewNetwork() {
	o.PreserveMacAddressOnNewNetwork.Unset()
}

// GetVnicEntity returns the VnicEntity field value if set, zero value otherwise.
func (o *RestoredObjectNetworkConfigProto) GetVnicEntity() EntityProto {
	if o == nil || o.VnicEntity == nil {
		var ret EntityProto
		return ret
	}
	return *o.VnicEntity
}

// GetVnicEntityOk returns a tuple with the VnicEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoredObjectNetworkConfigProto) GetVnicEntityOk() (*EntityProto, bool) {
	if o == nil || o.VnicEntity == nil {
		return nil, false
	}
	return o.VnicEntity, true
}

// HasVnicEntity returns a boolean if a field has been set.
func (o *RestoredObjectNetworkConfigProto) HasVnicEntity() bool {
	if o != nil && o.VnicEntity != nil {
		return true
	}

	return false
}

// SetVnicEntity gets a reference to the given EntityProto and assigns it to the VnicEntity field.
func (o *RestoredObjectNetworkConfigProto) SetVnicEntity(v EntityProto) {
	o.VnicEntity = &v
}

func (o RestoredObjectNetworkConfigProto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DetachNetwork.IsSet() {
		toSerialize["detachNetwork"] = o.DetachNetwork.Get()
	}
	if o.DisableNetwork.IsSet() {
		toSerialize["disableNetwork"] = o.DisableNetwork.Get()
	}
	if o.Mappings != nil {
		toSerialize["mappings"] = o.Mappings
	}
	if o.NetworkEntity != nil {
		toSerialize["networkEntity"] = o.NetworkEntity
	}
	if o.PreserveMacAddressOnNewNetwork.IsSet() {
		toSerialize["preserveMacAddressOnNewNetwork"] = o.PreserveMacAddressOnNewNetwork.Get()
	}
	if o.VnicEntity != nil {
		toSerialize["vnicEntity"] = o.VnicEntity
	}
	return json.Marshal(toSerialize)
}

type NullableRestoredObjectNetworkConfigProto struct {
	value *RestoredObjectNetworkConfigProto
	isSet bool
}

func (v NullableRestoredObjectNetworkConfigProto) Get() *RestoredObjectNetworkConfigProto {
	return v.value
}

func (v *NullableRestoredObjectNetworkConfigProto) Set(val *RestoredObjectNetworkConfigProto) {
	v.value = val
	v.isSet = true
}

func (v NullableRestoredObjectNetworkConfigProto) IsSet() bool {
	return v.isSet
}

func (v *NullableRestoredObjectNetworkConfigProto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestoredObjectNetworkConfigProto(val *RestoredObjectNetworkConfigProto) *NullableRestoredObjectNetworkConfigProto {
	return &NullableRestoredObjectNetworkConfigProto{value: val, isSet: true}
}

func (v NullableRestoredObjectNetworkConfigProto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestoredObjectNetworkConfigProto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


