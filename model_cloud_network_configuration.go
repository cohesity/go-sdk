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

// CloudNetworkConfiguration Specifies all of the parameters needed for network configuration of the new Cloud Cluster.
type CloudNetworkConfiguration struct {
	// Specifies the default gateway IP address (or addresses) for the Cluster network.
	ClusterGateway NullableString `json:"clusterGateway,omitempty"`
	// Specifies the subnet mask (or masks) of the Cluster network.
	ClusterSubnetMask NullableString `json:"clusterSubnetMask,omitempty"`
	// Specifies the list of DNS Servers this cluster should be configured with.
	DnsServers []string `json:"dnsServers,omitempty"`
	// Specifies the list of domain names this cluster should be configured with.
	DomainNames []string `json:"domainNames,omitempty"`
	// Specifies the list of NTP Servers this cluster should be configured with.
	NtpServers []string `json:"ntpServers,omitempty"`
}

// NewCloudNetworkConfiguration instantiates a new CloudNetworkConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudNetworkConfiguration() *CloudNetworkConfiguration {
	this := CloudNetworkConfiguration{}
	return &this
}

// NewCloudNetworkConfigurationWithDefaults instantiates a new CloudNetworkConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudNetworkConfigurationWithDefaults() *CloudNetworkConfiguration {
	this := CloudNetworkConfiguration{}
	return &this
}

// GetClusterGateway returns the ClusterGateway field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudNetworkConfiguration) GetClusterGateway() string {
	if o == nil || o.ClusterGateway.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClusterGateway.Get()
}

// GetClusterGatewayOk returns a tuple with the ClusterGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudNetworkConfiguration) GetClusterGatewayOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClusterGateway.Get(), o.ClusterGateway.IsSet()
}

// HasClusterGateway returns a boolean if a field has been set.
func (o *CloudNetworkConfiguration) HasClusterGateway() bool {
	if o != nil && o.ClusterGateway.IsSet() {
		return true
	}

	return false
}

// SetClusterGateway gets a reference to the given NullableString and assigns it to the ClusterGateway field.
func (o *CloudNetworkConfiguration) SetClusterGateway(v string) {
	o.ClusterGateway.Set(&v)
}
// SetClusterGatewayNil sets the value for ClusterGateway to be an explicit nil
func (o *CloudNetworkConfiguration) SetClusterGatewayNil() {
	o.ClusterGateway.Set(nil)
}

// UnsetClusterGateway ensures that no value is present for ClusterGateway, not even an explicit nil
func (o *CloudNetworkConfiguration) UnsetClusterGateway() {
	o.ClusterGateway.Unset()
}

// GetClusterSubnetMask returns the ClusterSubnetMask field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudNetworkConfiguration) GetClusterSubnetMask() string {
	if o == nil || o.ClusterSubnetMask.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClusterSubnetMask.Get()
}

// GetClusterSubnetMaskOk returns a tuple with the ClusterSubnetMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudNetworkConfiguration) GetClusterSubnetMaskOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClusterSubnetMask.Get(), o.ClusterSubnetMask.IsSet()
}

// HasClusterSubnetMask returns a boolean if a field has been set.
func (o *CloudNetworkConfiguration) HasClusterSubnetMask() bool {
	if o != nil && o.ClusterSubnetMask.IsSet() {
		return true
	}

	return false
}

// SetClusterSubnetMask gets a reference to the given NullableString and assigns it to the ClusterSubnetMask field.
func (o *CloudNetworkConfiguration) SetClusterSubnetMask(v string) {
	o.ClusterSubnetMask.Set(&v)
}
// SetClusterSubnetMaskNil sets the value for ClusterSubnetMask to be an explicit nil
func (o *CloudNetworkConfiguration) SetClusterSubnetMaskNil() {
	o.ClusterSubnetMask.Set(nil)
}

// UnsetClusterSubnetMask ensures that no value is present for ClusterSubnetMask, not even an explicit nil
func (o *CloudNetworkConfiguration) UnsetClusterSubnetMask() {
	o.ClusterSubnetMask.Unset()
}

// GetDnsServers returns the DnsServers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudNetworkConfiguration) GetDnsServers() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.DnsServers
}

// GetDnsServersOk returns a tuple with the DnsServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudNetworkConfiguration) GetDnsServersOk() (*[]string, bool) {
	if o == nil || o.DnsServers == nil {
		return nil, false
	}
	return &o.DnsServers, true
}

// HasDnsServers returns a boolean if a field has been set.
func (o *CloudNetworkConfiguration) HasDnsServers() bool {
	if o != nil && o.DnsServers != nil {
		return true
	}

	return false
}

// SetDnsServers gets a reference to the given []string and assigns it to the DnsServers field.
func (o *CloudNetworkConfiguration) SetDnsServers(v []string) {
	o.DnsServers = v
}

// GetDomainNames returns the DomainNames field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudNetworkConfiguration) GetDomainNames() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.DomainNames
}

// GetDomainNamesOk returns a tuple with the DomainNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudNetworkConfiguration) GetDomainNamesOk() (*[]string, bool) {
	if o == nil || o.DomainNames == nil {
		return nil, false
	}
	return &o.DomainNames, true
}

// HasDomainNames returns a boolean if a field has been set.
func (o *CloudNetworkConfiguration) HasDomainNames() bool {
	if o != nil && o.DomainNames != nil {
		return true
	}

	return false
}

// SetDomainNames gets a reference to the given []string and assigns it to the DomainNames field.
func (o *CloudNetworkConfiguration) SetDomainNames(v []string) {
	o.DomainNames = v
}

// GetNtpServers returns the NtpServers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudNetworkConfiguration) GetNtpServers() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.NtpServers
}

// GetNtpServersOk returns a tuple with the NtpServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudNetworkConfiguration) GetNtpServersOk() (*[]string, bool) {
	if o == nil || o.NtpServers == nil {
		return nil, false
	}
	return &o.NtpServers, true
}

// HasNtpServers returns a boolean if a field has been set.
func (o *CloudNetworkConfiguration) HasNtpServers() bool {
	if o != nil && o.NtpServers != nil {
		return true
	}

	return false
}

// SetNtpServers gets a reference to the given []string and assigns it to the NtpServers field.
func (o *CloudNetworkConfiguration) SetNtpServers(v []string) {
	o.NtpServers = v
}

func (o CloudNetworkConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClusterGateway.IsSet() {
		toSerialize["clusterGateway"] = o.ClusterGateway.Get()
	}
	if o.ClusterSubnetMask.IsSet() {
		toSerialize["clusterSubnetMask"] = o.ClusterSubnetMask.Get()
	}
	if o.DnsServers != nil {
		toSerialize["dnsServers"] = o.DnsServers
	}
	if o.DomainNames != nil {
		toSerialize["domainNames"] = o.DomainNames
	}
	if o.NtpServers != nil {
		toSerialize["ntpServers"] = o.NtpServers
	}
	return json.Marshal(toSerialize)
}

type NullableCloudNetworkConfiguration struct {
	value *CloudNetworkConfiguration
	isSet bool
}

func (v NullableCloudNetworkConfiguration) Get() *CloudNetworkConfiguration {
	return v.value
}

func (v *NullableCloudNetworkConfiguration) Set(val *CloudNetworkConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudNetworkConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudNetworkConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudNetworkConfiguration(val *CloudNetworkConfiguration) *NullableCloudNetworkConfiguration {
	return &NullableCloudNetworkConfiguration{value: val, isSet: true}
}

func (v NullableCloudNetworkConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudNetworkConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


