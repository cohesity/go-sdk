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

// checks if the ResourceEndpoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceEndpoint{}

// ResourceEndpoint Specifies the details about the resource endpoint.
type ResourceEndpoint struct {
	// Specifies the fqdn of this endpoint.
	Fqdn NullableString `json:"fqdn,omitempty"`
	// Specifies the ipv4 address of this endpoint.
	Ipv4addr NullableString `json:"ipv4addr,omitempty"`
	// Specifies the ipv6 address of this endpoint.
	Ipv6addr NullableString `json:"ipv6addr,omitempty"`
	// Whether to use this endpoint to connect.
	IsPreferredEndpoint NullableBool `json:"isPreferredEndpoint,omitempty"`
	// Specifies the preferred address to use for connecting.
	PreferredAddress NullableString `json:"preferredAddress,omitempty"`
	// Specifies the subnet Ip4 address of this endpoint.
	SubnetIp4addr NullableString `json:"subnetIp4addr,omitempty"`
}

// NewResourceEndpoint instantiates a new ResourceEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceEndpoint() *ResourceEndpoint {
	this := ResourceEndpoint{}
	return &this
}

// NewResourceEndpointWithDefaults instantiates a new ResourceEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceEndpointWithDefaults() *ResourceEndpoint {
	this := ResourceEndpoint{}
	return &this
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceEndpoint) GetFqdn() string {
	if o == nil || IsNil(o.Fqdn.Get()) {
		var ret string
		return ret
	}
	return *o.Fqdn.Get()
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceEndpoint) GetFqdnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fqdn.Get(), o.Fqdn.IsSet()
}

// HasFqdn returns a boolean if a field has been set.
func (o *ResourceEndpoint) HasFqdn() bool {
	if o != nil && o.Fqdn.IsSet() {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given NullableString and assigns it to the Fqdn field.
func (o *ResourceEndpoint) SetFqdn(v string) {
	o.Fqdn.Set(&v)
}
// SetFqdnNil sets the value for Fqdn to be an explicit nil
func (o *ResourceEndpoint) SetFqdnNil() {
	o.Fqdn.Set(nil)
}

// UnsetFqdn ensures that no value is present for Fqdn, not even an explicit nil
func (o *ResourceEndpoint) UnsetFqdn() {
	o.Fqdn.Unset()
}

// GetIpv4addr returns the Ipv4addr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceEndpoint) GetIpv4addr() string {
	if o == nil || IsNil(o.Ipv4addr.Get()) {
		var ret string
		return ret
	}
	return *o.Ipv4addr.Get()
}

// GetIpv4addrOk returns a tuple with the Ipv4addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceEndpoint) GetIpv4addrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ipv4addr.Get(), o.Ipv4addr.IsSet()
}

// HasIpv4addr returns a boolean if a field has been set.
func (o *ResourceEndpoint) HasIpv4addr() bool {
	if o != nil && o.Ipv4addr.IsSet() {
		return true
	}

	return false
}

// SetIpv4addr gets a reference to the given NullableString and assigns it to the Ipv4addr field.
func (o *ResourceEndpoint) SetIpv4addr(v string) {
	o.Ipv4addr.Set(&v)
}
// SetIpv4addrNil sets the value for Ipv4addr to be an explicit nil
func (o *ResourceEndpoint) SetIpv4addrNil() {
	o.Ipv4addr.Set(nil)
}

// UnsetIpv4addr ensures that no value is present for Ipv4addr, not even an explicit nil
func (o *ResourceEndpoint) UnsetIpv4addr() {
	o.Ipv4addr.Unset()
}

// GetIpv6addr returns the Ipv6addr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceEndpoint) GetIpv6addr() string {
	if o == nil || IsNil(o.Ipv6addr.Get()) {
		var ret string
		return ret
	}
	return *o.Ipv6addr.Get()
}

// GetIpv6addrOk returns a tuple with the Ipv6addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceEndpoint) GetIpv6addrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ipv6addr.Get(), o.Ipv6addr.IsSet()
}

// HasIpv6addr returns a boolean if a field has been set.
func (o *ResourceEndpoint) HasIpv6addr() bool {
	if o != nil && o.Ipv6addr.IsSet() {
		return true
	}

	return false
}

// SetIpv6addr gets a reference to the given NullableString and assigns it to the Ipv6addr field.
func (o *ResourceEndpoint) SetIpv6addr(v string) {
	o.Ipv6addr.Set(&v)
}
// SetIpv6addrNil sets the value for Ipv6addr to be an explicit nil
func (o *ResourceEndpoint) SetIpv6addrNil() {
	o.Ipv6addr.Set(nil)
}

// UnsetIpv6addr ensures that no value is present for Ipv6addr, not even an explicit nil
func (o *ResourceEndpoint) UnsetIpv6addr() {
	o.Ipv6addr.Unset()
}

// GetIsPreferredEndpoint returns the IsPreferredEndpoint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceEndpoint) GetIsPreferredEndpoint() bool {
	if o == nil || IsNil(o.IsPreferredEndpoint.Get()) {
		var ret bool
		return ret
	}
	return *o.IsPreferredEndpoint.Get()
}

// GetIsPreferredEndpointOk returns a tuple with the IsPreferredEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceEndpoint) GetIsPreferredEndpointOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsPreferredEndpoint.Get(), o.IsPreferredEndpoint.IsSet()
}

// HasIsPreferredEndpoint returns a boolean if a field has been set.
func (o *ResourceEndpoint) HasIsPreferredEndpoint() bool {
	if o != nil && o.IsPreferredEndpoint.IsSet() {
		return true
	}

	return false
}

// SetIsPreferredEndpoint gets a reference to the given NullableBool and assigns it to the IsPreferredEndpoint field.
func (o *ResourceEndpoint) SetIsPreferredEndpoint(v bool) {
	o.IsPreferredEndpoint.Set(&v)
}
// SetIsPreferredEndpointNil sets the value for IsPreferredEndpoint to be an explicit nil
func (o *ResourceEndpoint) SetIsPreferredEndpointNil() {
	o.IsPreferredEndpoint.Set(nil)
}

// UnsetIsPreferredEndpoint ensures that no value is present for IsPreferredEndpoint, not even an explicit nil
func (o *ResourceEndpoint) UnsetIsPreferredEndpoint() {
	o.IsPreferredEndpoint.Unset()
}

// GetPreferredAddress returns the PreferredAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceEndpoint) GetPreferredAddress() string {
	if o == nil || IsNil(o.PreferredAddress.Get()) {
		var ret string
		return ret
	}
	return *o.PreferredAddress.Get()
}

// GetPreferredAddressOk returns a tuple with the PreferredAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceEndpoint) GetPreferredAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreferredAddress.Get(), o.PreferredAddress.IsSet()
}

// HasPreferredAddress returns a boolean if a field has been set.
func (o *ResourceEndpoint) HasPreferredAddress() bool {
	if o != nil && o.PreferredAddress.IsSet() {
		return true
	}

	return false
}

// SetPreferredAddress gets a reference to the given NullableString and assigns it to the PreferredAddress field.
func (o *ResourceEndpoint) SetPreferredAddress(v string) {
	o.PreferredAddress.Set(&v)
}
// SetPreferredAddressNil sets the value for PreferredAddress to be an explicit nil
func (o *ResourceEndpoint) SetPreferredAddressNil() {
	o.PreferredAddress.Set(nil)
}

// UnsetPreferredAddress ensures that no value is present for PreferredAddress, not even an explicit nil
func (o *ResourceEndpoint) UnsetPreferredAddress() {
	o.PreferredAddress.Unset()
}

// GetSubnetIp4addr returns the SubnetIp4addr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceEndpoint) GetSubnetIp4addr() string {
	if o == nil || IsNil(o.SubnetIp4addr.Get()) {
		var ret string
		return ret
	}
	return *o.SubnetIp4addr.Get()
}

// GetSubnetIp4addrOk returns a tuple with the SubnetIp4addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceEndpoint) GetSubnetIp4addrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubnetIp4addr.Get(), o.SubnetIp4addr.IsSet()
}

// HasSubnetIp4addr returns a boolean if a field has been set.
func (o *ResourceEndpoint) HasSubnetIp4addr() bool {
	if o != nil && o.SubnetIp4addr.IsSet() {
		return true
	}

	return false
}

// SetSubnetIp4addr gets a reference to the given NullableString and assigns it to the SubnetIp4addr field.
func (o *ResourceEndpoint) SetSubnetIp4addr(v string) {
	o.SubnetIp4addr.Set(&v)
}
// SetSubnetIp4addrNil sets the value for SubnetIp4addr to be an explicit nil
func (o *ResourceEndpoint) SetSubnetIp4addrNil() {
	o.SubnetIp4addr.Set(nil)
}

// UnsetSubnetIp4addr ensures that no value is present for SubnetIp4addr, not even an explicit nil
func (o *ResourceEndpoint) UnsetSubnetIp4addr() {
	o.SubnetIp4addr.Unset()
}

func (o ResourceEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceEndpoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Fqdn.IsSet() {
		toSerialize["fqdn"] = o.Fqdn.Get()
	}
	if o.Ipv4addr.IsSet() {
		toSerialize["ipv4addr"] = o.Ipv4addr.Get()
	}
	if o.Ipv6addr.IsSet() {
		toSerialize["ipv6addr"] = o.Ipv6addr.Get()
	}
	if o.IsPreferredEndpoint.IsSet() {
		toSerialize["isPreferredEndpoint"] = o.IsPreferredEndpoint.Get()
	}
	if o.PreferredAddress.IsSet() {
		toSerialize["preferredAddress"] = o.PreferredAddress.Get()
	}
	if o.SubnetIp4addr.IsSet() {
		toSerialize["subnetIp4addr"] = o.SubnetIp4addr.Get()
	}
	return toSerialize, nil
}

type NullableResourceEndpoint struct {
	value *ResourceEndpoint
	isSet bool
}

func (v NullableResourceEndpoint) Get() *ResourceEndpoint {
	return v.value
}

func (v *NullableResourceEndpoint) Set(val *ResourceEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceEndpoint(val *ResourceEndpoint) *NullableResourceEndpoint {
	return &NullableResourceEndpoint{value: val, isSet: true}
}

func (v NullableResourceEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


