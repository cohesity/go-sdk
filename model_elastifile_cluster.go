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

// ElastifileCluster Specifies information about a Elastifile Cluster.
type ElastifileCluster struct {
	// IP addresses of Elastifile nodes.
	EnodeIpAddressVec []string `json:"enodeIpAddressVec,omitempty"`
	// Specifies the load balancer VIP if present.
	LoadBalancerVip NullableString `json:"loadBalancerVip,omitempty"`
	// Specifies name of a Elastifile Cluster
	Name NullableString `json:"name,omitempty"`
	// Specifies the UUID of a Elastifile Cluster.
	Uuid NullableString `json:"uuid,omitempty"`
	// Specifies the version of a Elastifile Cluster.
	Version NullableString `json:"version,omitempty"`
}

// NewElastifileCluster instantiates a new ElastifileCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewElastifileCluster() *ElastifileCluster {
	this := ElastifileCluster{}
	return &this
}

// NewElastifileClusterWithDefaults instantiates a new ElastifileCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewElastifileClusterWithDefaults() *ElastifileCluster {
	this := ElastifileCluster{}
	return &this
}

// GetEnodeIpAddressVec returns the EnodeIpAddressVec field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElastifileCluster) GetEnodeIpAddressVec() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.EnodeIpAddressVec
}

// GetEnodeIpAddressVecOk returns a tuple with the EnodeIpAddressVec field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ElastifileCluster) GetEnodeIpAddressVecOk() (*[]string, bool) {
	if o == nil || o.EnodeIpAddressVec == nil {
		return nil, false
	}
	return &o.EnodeIpAddressVec, true
}

// HasEnodeIpAddressVec returns a boolean if a field has been set.
func (o *ElastifileCluster) HasEnodeIpAddressVec() bool {
	if o != nil && o.EnodeIpAddressVec != nil {
		return true
	}

	return false
}

// SetEnodeIpAddressVec gets a reference to the given []string and assigns it to the EnodeIpAddressVec field.
func (o *ElastifileCluster) SetEnodeIpAddressVec(v []string) {
	o.EnodeIpAddressVec = v
}

// GetLoadBalancerVip returns the LoadBalancerVip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElastifileCluster) GetLoadBalancerVip() string {
	if o == nil || o.LoadBalancerVip.Get() == nil {
		var ret string
		return ret
	}
	return *o.LoadBalancerVip.Get()
}

// GetLoadBalancerVipOk returns a tuple with the LoadBalancerVip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ElastifileCluster) GetLoadBalancerVipOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LoadBalancerVip.Get(), o.LoadBalancerVip.IsSet()
}

// HasLoadBalancerVip returns a boolean if a field has been set.
func (o *ElastifileCluster) HasLoadBalancerVip() bool {
	if o != nil && o.LoadBalancerVip.IsSet() {
		return true
	}

	return false
}

// SetLoadBalancerVip gets a reference to the given NullableString and assigns it to the LoadBalancerVip field.
func (o *ElastifileCluster) SetLoadBalancerVip(v string) {
	o.LoadBalancerVip.Set(&v)
}
// SetLoadBalancerVipNil sets the value for LoadBalancerVip to be an explicit nil
func (o *ElastifileCluster) SetLoadBalancerVipNil() {
	o.LoadBalancerVip.Set(nil)
}

// UnsetLoadBalancerVip ensures that no value is present for LoadBalancerVip, not even an explicit nil
func (o *ElastifileCluster) UnsetLoadBalancerVip() {
	o.LoadBalancerVip.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElastifileCluster) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ElastifileCluster) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ElastifileCluster) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ElastifileCluster) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ElastifileCluster) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ElastifileCluster) UnsetName() {
	o.Name.Unset()
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElastifileCluster) GetUuid() string {
	if o == nil || o.Uuid.Get() == nil {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ElastifileCluster) GetUuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *ElastifileCluster) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *ElastifileCluster) SetUuid(v string) {
	o.Uuid.Set(&v)
}
// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *ElastifileCluster) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *ElastifileCluster) UnsetUuid() {
	o.Uuid.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElastifileCluster) GetVersion() string {
	if o == nil || o.Version.Get() == nil {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ElastifileCluster) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *ElastifileCluster) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *ElastifileCluster) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *ElastifileCluster) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *ElastifileCluster) UnsetVersion() {
	o.Version.Unset()
}

func (o ElastifileCluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EnodeIpAddressVec != nil {
		toSerialize["enodeIpAddressVec"] = o.EnodeIpAddressVec
	}
	if o.LoadBalancerVip.IsSet() {
		toSerialize["loadBalancerVip"] = o.LoadBalancerVip.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Uuid.IsSet() {
		toSerialize["uuid"] = o.Uuid.Get()
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableElastifileCluster struct {
	value *ElastifileCluster
	isSet bool
}

func (v NullableElastifileCluster) Get() *ElastifileCluster {
	return v.value
}

func (v *NullableElastifileCluster) Set(val *ElastifileCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableElastifileCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableElastifileCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableElastifileCluster(val *ElastifileCluster) *NullableElastifileCluster {
	return &NullableElastifileCluster{value: val, isSet: true}
}

func (v NullableElastifileCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableElastifileCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


