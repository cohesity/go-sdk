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

// IscsiSanPort Specifies an iSCSI SAN Port.
type IscsiSanPort struct {
	// Specifies the IP address of the SAN port.
	IpAddress NullableString `json:"ipAddress,omitempty"`
	// Specifies the qualified name of the iSCSI port (IQN).
	Iqn NullableString `json:"iqn,omitempty"`
	// Specifies the listening port(tcp) of the SAN port.
	TcpPort NullableInt32 `json:"tcpPort,omitempty"`
}

// NewIscsiSanPort instantiates a new IscsiSanPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIscsiSanPort() *IscsiSanPort {
	this := IscsiSanPort{}
	return &this
}

// NewIscsiSanPortWithDefaults instantiates a new IscsiSanPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIscsiSanPortWithDefaults() *IscsiSanPort {
	this := IscsiSanPort{}
	return &this
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IscsiSanPort) GetIpAddress() string {
	if o == nil || o.IpAddress.Get() == nil {
		var ret string
		return ret
	}
	return *o.IpAddress.Get()
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IscsiSanPort) GetIpAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IpAddress.Get(), o.IpAddress.IsSet()
}

// HasIpAddress returns a boolean if a field has been set.
func (o *IscsiSanPort) HasIpAddress() bool {
	if o != nil && o.IpAddress.IsSet() {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given NullableString and assigns it to the IpAddress field.
func (o *IscsiSanPort) SetIpAddress(v string) {
	o.IpAddress.Set(&v)
}
// SetIpAddressNil sets the value for IpAddress to be an explicit nil
func (o *IscsiSanPort) SetIpAddressNil() {
	o.IpAddress.Set(nil)
}

// UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
func (o *IscsiSanPort) UnsetIpAddress() {
	o.IpAddress.Unset()
}

// GetIqn returns the Iqn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IscsiSanPort) GetIqn() string {
	if o == nil || o.Iqn.Get() == nil {
		var ret string
		return ret
	}
	return *o.Iqn.Get()
}

// GetIqnOk returns a tuple with the Iqn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IscsiSanPort) GetIqnOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Iqn.Get(), o.Iqn.IsSet()
}

// HasIqn returns a boolean if a field has been set.
func (o *IscsiSanPort) HasIqn() bool {
	if o != nil && o.Iqn.IsSet() {
		return true
	}

	return false
}

// SetIqn gets a reference to the given NullableString and assigns it to the Iqn field.
func (o *IscsiSanPort) SetIqn(v string) {
	o.Iqn.Set(&v)
}
// SetIqnNil sets the value for Iqn to be an explicit nil
func (o *IscsiSanPort) SetIqnNil() {
	o.Iqn.Set(nil)
}

// UnsetIqn ensures that no value is present for Iqn, not even an explicit nil
func (o *IscsiSanPort) UnsetIqn() {
	o.Iqn.Unset()
}

// GetTcpPort returns the TcpPort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IscsiSanPort) GetTcpPort() int32 {
	if o == nil || o.TcpPort.Get() == nil {
		var ret int32
		return ret
	}
	return *o.TcpPort.Get()
}

// GetTcpPortOk returns a tuple with the TcpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IscsiSanPort) GetTcpPortOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TcpPort.Get(), o.TcpPort.IsSet()
}

// HasTcpPort returns a boolean if a field has been set.
func (o *IscsiSanPort) HasTcpPort() bool {
	if o != nil && o.TcpPort.IsSet() {
		return true
	}

	return false
}

// SetTcpPort gets a reference to the given NullableInt32 and assigns it to the TcpPort field.
func (o *IscsiSanPort) SetTcpPort(v int32) {
	o.TcpPort.Set(&v)
}
// SetTcpPortNil sets the value for TcpPort to be an explicit nil
func (o *IscsiSanPort) SetTcpPortNil() {
	o.TcpPort.Set(nil)
}

// UnsetTcpPort ensures that no value is present for TcpPort, not even an explicit nil
func (o *IscsiSanPort) UnsetTcpPort() {
	o.TcpPort.Unset()
}

func (o IscsiSanPort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IpAddress.IsSet() {
		toSerialize["ipAddress"] = o.IpAddress.Get()
	}
	if o.Iqn.IsSet() {
		toSerialize["iqn"] = o.Iqn.Get()
	}
	if o.TcpPort.IsSet() {
		toSerialize["tcpPort"] = o.TcpPort.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIscsiSanPort struct {
	value *IscsiSanPort
	isSet bool
}

func (v NullableIscsiSanPort) Get() *IscsiSanPort {
	return v.value
}

func (v *NullableIscsiSanPort) Set(val *IscsiSanPort) {
	v.value = val
	v.isSet = true
}

func (v NullableIscsiSanPort) IsSet() bool {
	return v.isSet
}

func (v *NullableIscsiSanPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIscsiSanPort(val *IscsiSanPort) *NullableIscsiSanPort {
	return &NullableIscsiSanPort{value: val, isSet: true}
}

func (v NullableIscsiSanPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIscsiSanPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


