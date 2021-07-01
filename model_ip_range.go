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

// IpRange IP Range for range of vip address addition
type IpRange struct {
	// Optional End IP of the range If not specified, EndIp is same as StartIp
	EndIp NullableString `json:"endIp,omitempty"`
	// Start IP of the range
	StartIp NullableString `json:"startIp,omitempty"`
}

// NewIpRange instantiates a new IpRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpRange() *IpRange {
	this := IpRange{}
	return &this
}

// NewIpRangeWithDefaults instantiates a new IpRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpRangeWithDefaults() *IpRange {
	this := IpRange{}
	return &this
}

// GetEndIp returns the EndIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpRange) GetEndIp() string {
	if o == nil || o.EndIp.Get() == nil {
		var ret string
		return ret
	}
	return *o.EndIp.Get()
}

// GetEndIpOk returns a tuple with the EndIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpRange) GetEndIpOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndIp.Get(), o.EndIp.IsSet()
}

// HasEndIp returns a boolean if a field has been set.
func (o *IpRange) HasEndIp() bool {
	if o != nil && o.EndIp.IsSet() {
		return true
	}

	return false
}

// SetEndIp gets a reference to the given NullableString and assigns it to the EndIp field.
func (o *IpRange) SetEndIp(v string) {
	o.EndIp.Set(&v)
}
// SetEndIpNil sets the value for EndIp to be an explicit nil
func (o *IpRange) SetEndIpNil() {
	o.EndIp.Set(nil)
}

// UnsetEndIp ensures that no value is present for EndIp, not even an explicit nil
func (o *IpRange) UnsetEndIp() {
	o.EndIp.Unset()
}

// GetStartIp returns the StartIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpRange) GetStartIp() string {
	if o == nil || o.StartIp.Get() == nil {
		var ret string
		return ret
	}
	return *o.StartIp.Get()
}

// GetStartIpOk returns a tuple with the StartIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpRange) GetStartIpOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartIp.Get(), o.StartIp.IsSet()
}

// HasStartIp returns a boolean if a field has been set.
func (o *IpRange) HasStartIp() bool {
	if o != nil && o.StartIp.IsSet() {
		return true
	}

	return false
}

// SetStartIp gets a reference to the given NullableString and assigns it to the StartIp field.
func (o *IpRange) SetStartIp(v string) {
	o.StartIp.Set(&v)
}
// SetStartIpNil sets the value for StartIp to be an explicit nil
func (o *IpRange) SetStartIpNil() {
	o.StartIp.Set(nil)
}

// UnsetStartIp ensures that no value is present for StartIp, not even an explicit nil
func (o *IpRange) UnsetStartIp() {
	o.StartIp.Unset()
}

func (o IpRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EndIp.IsSet() {
		toSerialize["endIp"] = o.EndIp.Get()
	}
	if o.StartIp.IsSet() {
		toSerialize["startIp"] = o.StartIp.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIpRange struct {
	value *IpRange
	isSet bool
}

func (v NullableIpRange) Get() *IpRange {
	return v.value
}

func (v *NullableIpRange) Set(val *IpRange) {
	v.value = val
	v.isSet = true
}

func (v NullableIpRange) IsSet() bool {
	return v.isSet
}

func (v *NullableIpRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpRange(val *IpRange) *NullableIpRange {
	return &NullableIpRange{value: val, isSet: true}
}

func (v NullableIpRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


