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

// ClusterConfigProtoSubnet struct for ClusterConfigProtoSubnet
type ClusterConfigProtoSubnet struct {
	// The component that has claimed this subnet.
	Component NullableInt32 `json:"component,omitempty"`
	// Description of the subnet.
	Description NullableString `json:"description,omitempty"`
	// Gateway for the subnet.
	Gateway NullableString `json:"gateway,omitempty"`
	// ID for this subnet.
	Id NullableInt32 `json:"id,omitempty"`
	// ip is subnet IP address given either in v4 or v6. Netmask is specified by giving CIDR length in netmask_bits for ipv6. For IPv4 addresses, netmask_ip4 field is set in dotted decimal.
	Ip NullableString `json:"ip,omitempty"`
	NetmaskBits NullableInt32 `json:"netmaskBits,omitempty"`
	NetmaskIp4 NullableString `json:"netmaskIp4,omitempty"`
	// Whether clients from this subnet can mount using NFS protocol.
	NfsAccess NullableInt32 `json:"nfsAccess,omitempty"`
	// Whether all clients from this subnet can map view with view_all_squash_uid/view_all_squash_gid configured in the view.
	NfsAllSquash NullableBool `json:"nfsAllSquash,omitempty"`
	// Whether clients from this subnet can mount as root on NFS.
	NfsRootSquash NullableBool `json:"nfsRootSquash,omitempty"`
	// Whether clients from this subnet can mount using SMB protocol.
	SmbAccess NullableInt32 `json:"smbAccess,omitempty"`
}

// NewClusterConfigProtoSubnet instantiates a new ClusterConfigProtoSubnet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterConfigProtoSubnet() *ClusterConfigProtoSubnet {
	this := ClusterConfigProtoSubnet{}
	return &this
}

// NewClusterConfigProtoSubnetWithDefaults instantiates a new ClusterConfigProtoSubnet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterConfigProtoSubnetWithDefaults() *ClusterConfigProtoSubnet {
	this := ClusterConfigProtoSubnet{}
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterConfigProtoSubnet) GetComponent() int32 {
	if o == nil || o.Component.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Component.Get()
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterConfigProtoSubnet) GetComponentOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Component.Get(), o.Component.IsSet()
}

// HasComponent returns a boolean if a field has been set.
func (o *ClusterConfigProtoSubnet) HasComponent() bool {
	if o != nil && o.Component.IsSet() {
		return true
	}

	return false
}

// SetComponent gets a reference to the given NullableInt32 and assigns it to the Component field.
func (o *ClusterConfigProtoSubnet) SetComponent(v int32) {
	o.Component.Set(&v)
}
// SetComponentNil sets the value for Component to be an explicit nil
func (o *ClusterConfigProtoSubnet) SetComponentNil() {
	o.Component.Set(nil)
}

// UnsetComponent ensures that no value is present for Component, not even an explicit nil
func (o *ClusterConfigProtoSubnet) UnsetComponent() {
	o.Component.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterConfigProtoSubnet) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterConfigProtoSubnet) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ClusterConfigProtoSubnet) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ClusterConfigProtoSubnet) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ClusterConfigProtoSubnet) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ClusterConfigProtoSubnet) UnsetDescription() {
	o.Description.Unset()
}

// GetGateway returns the Gateway field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterConfigProtoSubnet) GetGateway() string {
	if o == nil || o.Gateway.Get() == nil {
		var ret string
		return ret
	}
	return *o.Gateway.Get()
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterConfigProtoSubnet) GetGatewayOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Gateway.Get(), o.Gateway.IsSet()
}

// HasGateway returns a boolean if a field has been set.
func (o *ClusterConfigProtoSubnet) HasGateway() bool {
	if o != nil && o.Gateway.IsSet() {
		return true
	}

	return false
}

// SetGateway gets a reference to the given NullableString and assigns it to the Gateway field.
func (o *ClusterConfigProtoSubnet) SetGateway(v string) {
	o.Gateway.Set(&v)
}
// SetGatewayNil sets the value for Gateway to be an explicit nil
func (o *ClusterConfigProtoSubnet) SetGatewayNil() {
	o.Gateway.Set(nil)
}

// UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
func (o *ClusterConfigProtoSubnet) UnsetGateway() {
	o.Gateway.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterConfigProtoSubnet) GetId() int32 {
	if o == nil || o.Id.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterConfigProtoSubnet) GetIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ClusterConfigProtoSubnet) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *ClusterConfigProtoSubnet) SetId(v int32) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ClusterConfigProtoSubnet) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ClusterConfigProtoSubnet) UnsetId() {
	o.Id.Unset()
}

// GetIp returns the Ip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterConfigProtoSubnet) GetIp() string {
	if o == nil || o.Ip.Get() == nil {
		var ret string
		return ret
	}
	return *o.Ip.Get()
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterConfigProtoSubnet) GetIpOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Ip.Get(), o.Ip.IsSet()
}

// HasIp returns a boolean if a field has been set.
func (o *ClusterConfigProtoSubnet) HasIp() bool {
	if o != nil && o.Ip.IsSet() {
		return true
	}

	return false
}

// SetIp gets a reference to the given NullableString and assigns it to the Ip field.
func (o *ClusterConfigProtoSubnet) SetIp(v string) {
	o.Ip.Set(&v)
}
// SetIpNil sets the value for Ip to be an explicit nil
func (o *ClusterConfigProtoSubnet) SetIpNil() {
	o.Ip.Set(nil)
}

// UnsetIp ensures that no value is present for Ip, not even an explicit nil
func (o *ClusterConfigProtoSubnet) UnsetIp() {
	o.Ip.Unset()
}

// GetNetmaskBits returns the NetmaskBits field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterConfigProtoSubnet) GetNetmaskBits() int32 {
	if o == nil || o.NetmaskBits.Get() == nil {
		var ret int32
		return ret
	}
	return *o.NetmaskBits.Get()
}

// GetNetmaskBitsOk returns a tuple with the NetmaskBits field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterConfigProtoSubnet) GetNetmaskBitsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NetmaskBits.Get(), o.NetmaskBits.IsSet()
}

// HasNetmaskBits returns a boolean if a field has been set.
func (o *ClusterConfigProtoSubnet) HasNetmaskBits() bool {
	if o != nil && o.NetmaskBits.IsSet() {
		return true
	}

	return false
}

// SetNetmaskBits gets a reference to the given NullableInt32 and assigns it to the NetmaskBits field.
func (o *ClusterConfigProtoSubnet) SetNetmaskBits(v int32) {
	o.NetmaskBits.Set(&v)
}
// SetNetmaskBitsNil sets the value for NetmaskBits to be an explicit nil
func (o *ClusterConfigProtoSubnet) SetNetmaskBitsNil() {
	o.NetmaskBits.Set(nil)
}

// UnsetNetmaskBits ensures that no value is present for NetmaskBits, not even an explicit nil
func (o *ClusterConfigProtoSubnet) UnsetNetmaskBits() {
	o.NetmaskBits.Unset()
}

// GetNetmaskIp4 returns the NetmaskIp4 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterConfigProtoSubnet) GetNetmaskIp4() string {
	if o == nil || o.NetmaskIp4.Get() == nil {
		var ret string
		return ret
	}
	return *o.NetmaskIp4.Get()
}

// GetNetmaskIp4Ok returns a tuple with the NetmaskIp4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterConfigProtoSubnet) GetNetmaskIp4Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NetmaskIp4.Get(), o.NetmaskIp4.IsSet()
}

// HasNetmaskIp4 returns a boolean if a field has been set.
func (o *ClusterConfigProtoSubnet) HasNetmaskIp4() bool {
	if o != nil && o.NetmaskIp4.IsSet() {
		return true
	}

	return false
}

// SetNetmaskIp4 gets a reference to the given NullableString and assigns it to the NetmaskIp4 field.
func (o *ClusterConfigProtoSubnet) SetNetmaskIp4(v string) {
	o.NetmaskIp4.Set(&v)
}
// SetNetmaskIp4Nil sets the value for NetmaskIp4 to be an explicit nil
func (o *ClusterConfigProtoSubnet) SetNetmaskIp4Nil() {
	o.NetmaskIp4.Set(nil)
}

// UnsetNetmaskIp4 ensures that no value is present for NetmaskIp4, not even an explicit nil
func (o *ClusterConfigProtoSubnet) UnsetNetmaskIp4() {
	o.NetmaskIp4.Unset()
}

// GetNfsAccess returns the NfsAccess field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterConfigProtoSubnet) GetNfsAccess() int32 {
	if o == nil || o.NfsAccess.Get() == nil {
		var ret int32
		return ret
	}
	return *o.NfsAccess.Get()
}

// GetNfsAccessOk returns a tuple with the NfsAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterConfigProtoSubnet) GetNfsAccessOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NfsAccess.Get(), o.NfsAccess.IsSet()
}

// HasNfsAccess returns a boolean if a field has been set.
func (o *ClusterConfigProtoSubnet) HasNfsAccess() bool {
	if o != nil && o.NfsAccess.IsSet() {
		return true
	}

	return false
}

// SetNfsAccess gets a reference to the given NullableInt32 and assigns it to the NfsAccess field.
func (o *ClusterConfigProtoSubnet) SetNfsAccess(v int32) {
	o.NfsAccess.Set(&v)
}
// SetNfsAccessNil sets the value for NfsAccess to be an explicit nil
func (o *ClusterConfigProtoSubnet) SetNfsAccessNil() {
	o.NfsAccess.Set(nil)
}

// UnsetNfsAccess ensures that no value is present for NfsAccess, not even an explicit nil
func (o *ClusterConfigProtoSubnet) UnsetNfsAccess() {
	o.NfsAccess.Unset()
}

// GetNfsAllSquash returns the NfsAllSquash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterConfigProtoSubnet) GetNfsAllSquash() bool {
	if o == nil || o.NfsAllSquash.Get() == nil {
		var ret bool
		return ret
	}
	return *o.NfsAllSquash.Get()
}

// GetNfsAllSquashOk returns a tuple with the NfsAllSquash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterConfigProtoSubnet) GetNfsAllSquashOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NfsAllSquash.Get(), o.NfsAllSquash.IsSet()
}

// HasNfsAllSquash returns a boolean if a field has been set.
func (o *ClusterConfigProtoSubnet) HasNfsAllSquash() bool {
	if o != nil && o.NfsAllSquash.IsSet() {
		return true
	}

	return false
}

// SetNfsAllSquash gets a reference to the given NullableBool and assigns it to the NfsAllSquash field.
func (o *ClusterConfigProtoSubnet) SetNfsAllSquash(v bool) {
	o.NfsAllSquash.Set(&v)
}
// SetNfsAllSquashNil sets the value for NfsAllSquash to be an explicit nil
func (o *ClusterConfigProtoSubnet) SetNfsAllSquashNil() {
	o.NfsAllSquash.Set(nil)
}

// UnsetNfsAllSquash ensures that no value is present for NfsAllSquash, not even an explicit nil
func (o *ClusterConfigProtoSubnet) UnsetNfsAllSquash() {
	o.NfsAllSquash.Unset()
}

// GetNfsRootSquash returns the NfsRootSquash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterConfigProtoSubnet) GetNfsRootSquash() bool {
	if o == nil || o.NfsRootSquash.Get() == nil {
		var ret bool
		return ret
	}
	return *o.NfsRootSquash.Get()
}

// GetNfsRootSquashOk returns a tuple with the NfsRootSquash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterConfigProtoSubnet) GetNfsRootSquashOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NfsRootSquash.Get(), o.NfsRootSquash.IsSet()
}

// HasNfsRootSquash returns a boolean if a field has been set.
func (o *ClusterConfigProtoSubnet) HasNfsRootSquash() bool {
	if o != nil && o.NfsRootSquash.IsSet() {
		return true
	}

	return false
}

// SetNfsRootSquash gets a reference to the given NullableBool and assigns it to the NfsRootSquash field.
func (o *ClusterConfigProtoSubnet) SetNfsRootSquash(v bool) {
	o.NfsRootSquash.Set(&v)
}
// SetNfsRootSquashNil sets the value for NfsRootSquash to be an explicit nil
func (o *ClusterConfigProtoSubnet) SetNfsRootSquashNil() {
	o.NfsRootSquash.Set(nil)
}

// UnsetNfsRootSquash ensures that no value is present for NfsRootSquash, not even an explicit nil
func (o *ClusterConfigProtoSubnet) UnsetNfsRootSquash() {
	o.NfsRootSquash.Unset()
}

// GetSmbAccess returns the SmbAccess field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterConfigProtoSubnet) GetSmbAccess() int32 {
	if o == nil || o.SmbAccess.Get() == nil {
		var ret int32
		return ret
	}
	return *o.SmbAccess.Get()
}

// GetSmbAccessOk returns a tuple with the SmbAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterConfigProtoSubnet) GetSmbAccessOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SmbAccess.Get(), o.SmbAccess.IsSet()
}

// HasSmbAccess returns a boolean if a field has been set.
func (o *ClusterConfigProtoSubnet) HasSmbAccess() bool {
	if o != nil && o.SmbAccess.IsSet() {
		return true
	}

	return false
}

// SetSmbAccess gets a reference to the given NullableInt32 and assigns it to the SmbAccess field.
func (o *ClusterConfigProtoSubnet) SetSmbAccess(v int32) {
	o.SmbAccess.Set(&v)
}
// SetSmbAccessNil sets the value for SmbAccess to be an explicit nil
func (o *ClusterConfigProtoSubnet) SetSmbAccessNil() {
	o.SmbAccess.Set(nil)
}

// UnsetSmbAccess ensures that no value is present for SmbAccess, not even an explicit nil
func (o *ClusterConfigProtoSubnet) UnsetSmbAccess() {
	o.SmbAccess.Unset()
}

func (o ClusterConfigProtoSubnet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Component.IsSet() {
		toSerialize["component"] = o.Component.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Gateway.IsSet() {
		toSerialize["gateway"] = o.Gateway.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Ip.IsSet() {
		toSerialize["ip"] = o.Ip.Get()
	}
	if o.NetmaskBits.IsSet() {
		toSerialize["netmaskBits"] = o.NetmaskBits.Get()
	}
	if o.NetmaskIp4.IsSet() {
		toSerialize["netmaskIp4"] = o.NetmaskIp4.Get()
	}
	if o.NfsAccess.IsSet() {
		toSerialize["nfsAccess"] = o.NfsAccess.Get()
	}
	if o.NfsAllSquash.IsSet() {
		toSerialize["nfsAllSquash"] = o.NfsAllSquash.Get()
	}
	if o.NfsRootSquash.IsSet() {
		toSerialize["nfsRootSquash"] = o.NfsRootSquash.Get()
	}
	if o.SmbAccess.IsSet() {
		toSerialize["smbAccess"] = o.SmbAccess.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableClusterConfigProtoSubnet struct {
	value *ClusterConfigProtoSubnet
	isSet bool
}

func (v NullableClusterConfigProtoSubnet) Get() *ClusterConfigProtoSubnet {
	return v.value
}

func (v *NullableClusterConfigProtoSubnet) Set(val *ClusterConfigProtoSubnet) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterConfigProtoSubnet) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterConfigProtoSubnet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterConfigProtoSubnet(val *ClusterConfigProtoSubnet) *NullableClusterConfigProtoSubnet {
	return &NullableClusterConfigProtoSubnet{value: val, isSet: true}
}

func (v NullableClusterConfigProtoSubnet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterConfigProtoSubnet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


