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

// NetworkConnectionInfo Specify the network connection information.
type NetworkConnectionInfo struct {
	// Specifies the domain name of the network connection.
	DomainName NullableString `json:"domainName,omitempty"`
	// Specifies the network Gateway of the network connection.
	NetworkGateway NullableString `json:"networkGateway,omitempty"`
	// Specifies the DNS Server of the network connection.
	Dns NullableString `json:"dns,omitempty"`
	// Specifies the NTP Server of the network connection.
	Ntp NullableString `json:"ntp,omitempty"`
}

// NewNetworkConnectionInfo instantiates a new NetworkConnectionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkConnectionInfo() *NetworkConnectionInfo {
	this := NetworkConnectionInfo{}
	return &this
}

// NewNetworkConnectionInfoWithDefaults instantiates a new NetworkConnectionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkConnectionInfoWithDefaults() *NetworkConnectionInfo {
	this := NetworkConnectionInfo{}
	return &this
}

// GetDomainName returns the DomainName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkConnectionInfo) GetDomainName() string {
	if o == nil || o.DomainName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DomainName.Get()
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkConnectionInfo) GetDomainNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DomainName.Get(), o.DomainName.IsSet()
}

// HasDomainName returns a boolean if a field has been set.
func (o *NetworkConnectionInfo) HasDomainName() bool {
	if o != nil && o.DomainName.IsSet() {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given NullableString and assigns it to the DomainName field.
func (o *NetworkConnectionInfo) SetDomainName(v string) {
	o.DomainName.Set(&v)
}
// SetDomainNameNil sets the value for DomainName to be an explicit nil
func (o *NetworkConnectionInfo) SetDomainNameNil() {
	o.DomainName.Set(nil)
}

// UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
func (o *NetworkConnectionInfo) UnsetDomainName() {
	o.DomainName.Unset()
}

// GetNetworkGateway returns the NetworkGateway field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkConnectionInfo) GetNetworkGateway() string {
	if o == nil || o.NetworkGateway.Get() == nil {
		var ret string
		return ret
	}
	return *o.NetworkGateway.Get()
}

// GetNetworkGatewayOk returns a tuple with the NetworkGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkConnectionInfo) GetNetworkGatewayOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NetworkGateway.Get(), o.NetworkGateway.IsSet()
}

// HasNetworkGateway returns a boolean if a field has been set.
func (o *NetworkConnectionInfo) HasNetworkGateway() bool {
	if o != nil && o.NetworkGateway.IsSet() {
		return true
	}

	return false
}

// SetNetworkGateway gets a reference to the given NullableString and assigns it to the NetworkGateway field.
func (o *NetworkConnectionInfo) SetNetworkGateway(v string) {
	o.NetworkGateway.Set(&v)
}
// SetNetworkGatewayNil sets the value for NetworkGateway to be an explicit nil
func (o *NetworkConnectionInfo) SetNetworkGatewayNil() {
	o.NetworkGateway.Set(nil)
}

// UnsetNetworkGateway ensures that no value is present for NetworkGateway, not even an explicit nil
func (o *NetworkConnectionInfo) UnsetNetworkGateway() {
	o.NetworkGateway.Unset()
}

// GetDns returns the Dns field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkConnectionInfo) GetDns() string {
	if o == nil || o.Dns.Get() == nil {
		var ret string
		return ret
	}
	return *o.Dns.Get()
}

// GetDnsOk returns a tuple with the Dns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkConnectionInfo) GetDnsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Dns.Get(), o.Dns.IsSet()
}

// HasDns returns a boolean if a field has been set.
func (o *NetworkConnectionInfo) HasDns() bool {
	if o != nil && o.Dns.IsSet() {
		return true
	}

	return false
}

// SetDns gets a reference to the given NullableString and assigns it to the Dns field.
func (o *NetworkConnectionInfo) SetDns(v string) {
	o.Dns.Set(&v)
}
// SetDnsNil sets the value for Dns to be an explicit nil
func (o *NetworkConnectionInfo) SetDnsNil() {
	o.Dns.Set(nil)
}

// UnsetDns ensures that no value is present for Dns, not even an explicit nil
func (o *NetworkConnectionInfo) UnsetDns() {
	o.Dns.Unset()
}

// GetNtp returns the Ntp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkConnectionInfo) GetNtp() string {
	if o == nil || o.Ntp.Get() == nil {
		var ret string
		return ret
	}
	return *o.Ntp.Get()
}

// GetNtpOk returns a tuple with the Ntp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkConnectionInfo) GetNtpOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Ntp.Get(), o.Ntp.IsSet()
}

// HasNtp returns a boolean if a field has been set.
func (o *NetworkConnectionInfo) HasNtp() bool {
	if o != nil && o.Ntp.IsSet() {
		return true
	}

	return false
}

// SetNtp gets a reference to the given NullableString and assigns it to the Ntp field.
func (o *NetworkConnectionInfo) SetNtp(v string) {
	o.Ntp.Set(&v)
}
// SetNtpNil sets the value for Ntp to be an explicit nil
func (o *NetworkConnectionInfo) SetNtpNil() {
	o.Ntp.Set(nil)
}

// UnsetNtp ensures that no value is present for Ntp, not even an explicit nil
func (o *NetworkConnectionInfo) UnsetNtp() {
	o.Ntp.Unset()
}

func (o NetworkConnectionInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DomainName.IsSet() {
		toSerialize["domainName"] = o.DomainName.Get()
	}
	if o.NetworkGateway.IsSet() {
		toSerialize["networkGateway"] = o.NetworkGateway.Get()
	}
	if o.Dns.IsSet() {
		toSerialize["dns"] = o.Dns.Get()
	}
	if o.Ntp.IsSet() {
		toSerialize["ntp"] = o.Ntp.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkConnectionInfo struct {
	value *NetworkConnectionInfo
	isSet bool
}

func (v NullableNetworkConnectionInfo) Get() *NetworkConnectionInfo {
	return v.value
}

func (v *NullableNetworkConnectionInfo) Set(val *NetworkConnectionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkConnectionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkConnectionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkConnectionInfo(val *NetworkConnectionInfo) *NullableNetworkConnectionInfo {
	return &NullableNetworkConnectionInfo{value: val, isSet: true}
}

func (v NullableNetworkConnectionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkConnectionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


