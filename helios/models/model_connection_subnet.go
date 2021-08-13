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

// ConnectionSubnet Specify the subnet used in connection.
type ConnectionSubnet struct {
	// Specifies the IP address part of the CIDR notation.
	Ip NullableString `json:"ip,omitempty"`
	// Specifies the number of set bits in the subnet mask.
	MaskBits NullableInt32 `json:"maskBits,omitempty"`
}

// NewConnectionSubnet instantiates a new ConnectionSubnet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionSubnet() *ConnectionSubnet {
	this := ConnectionSubnet{}
	return &this
}

// NewConnectionSubnetWithDefaults instantiates a new ConnectionSubnet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionSubnetWithDefaults() *ConnectionSubnet {
	this := ConnectionSubnet{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectionSubnet) GetIp() string {
	if o == nil || o.Ip.Get() == nil {
		var ret string
		return ret
	}
	return *o.Ip.Get()
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectionSubnet) GetIpOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Ip.Get(), o.Ip.IsSet()
}

// HasIp returns a boolean if a field has been set.
func (o *ConnectionSubnet) HasIp() bool {
	if o != nil && o.Ip.IsSet() {
		return true
	}

	return false
}

// SetIp gets a reference to the given NullableString and assigns it to the Ip field.
func (o *ConnectionSubnet) SetIp(v string) {
	o.Ip.Set(&v)
}
// SetIpNil sets the value for Ip to be an explicit nil
func (o *ConnectionSubnet) SetIpNil() {
	o.Ip.Set(nil)
}

// UnsetIp ensures that no value is present for Ip, not even an explicit nil
func (o *ConnectionSubnet) UnsetIp() {
	o.Ip.Unset()
}

// GetMaskBits returns the MaskBits field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectionSubnet) GetMaskBits() int32 {
	if o == nil || o.MaskBits.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MaskBits.Get()
}

// GetMaskBitsOk returns a tuple with the MaskBits field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectionSubnet) GetMaskBitsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MaskBits.Get(), o.MaskBits.IsSet()
}

// HasMaskBits returns a boolean if a field has been set.
func (o *ConnectionSubnet) HasMaskBits() bool {
	if o != nil && o.MaskBits.IsSet() {
		return true
	}

	return false
}

// SetMaskBits gets a reference to the given NullableInt32 and assigns it to the MaskBits field.
func (o *ConnectionSubnet) SetMaskBits(v int32) {
	o.MaskBits.Set(&v)
}
// SetMaskBitsNil sets the value for MaskBits to be an explicit nil
func (o *ConnectionSubnet) SetMaskBitsNil() {
	o.MaskBits.Set(nil)
}

// UnsetMaskBits ensures that no value is present for MaskBits, not even an explicit nil
func (o *ConnectionSubnet) UnsetMaskBits() {
	o.MaskBits.Unset()
}

func (o ConnectionSubnet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ip.IsSet() {
		toSerialize["ip"] = o.Ip.Get()
	}
	if o.MaskBits.IsSet() {
		toSerialize["maskBits"] = o.MaskBits.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableConnectionSubnet struct {
	value *ConnectionSubnet
	isSet bool
}

func (v NullableConnectionSubnet) Get() *ConnectionSubnet {
	return v.value
}

func (v *NullableConnectionSubnet) Set(val *ConnectionSubnet) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionSubnet) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionSubnet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionSubnet(val *ConnectionSubnet) *NullableConnectionSubnet {
	return &NullableConnectionSubnet{value: val, isSet: true}
}

func (v NullableConnectionSubnet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionSubnet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


