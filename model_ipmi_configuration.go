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

// IpmiConfiguration Specifies the parameters for configuration of IPMI. This is only needed for physical clusters.
type IpmiConfiguration struct {
	// Specifies the default Gateway IP Address for the IPMI network.
	IpmiGateway NullableString `json:"ipmiGateway,omitempty"`
	// Specifies the IPMI Password.
	IpmiPassword NullableString `json:"ipmiPassword,omitempty"`
	// Specifies the subnet mask for the IPMI network.
	IpmiSubnetMask NullableString `json:"ipmiSubnetMask,omitempty"`
	// Specifies the IPMI Username.
	IpmiUsername NullableString `json:"ipmiUsername,omitempty"`
}

// NewIpmiConfiguration instantiates a new IpmiConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpmiConfiguration() *IpmiConfiguration {
	this := IpmiConfiguration{}
	return &this
}

// NewIpmiConfigurationWithDefaults instantiates a new IpmiConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpmiConfigurationWithDefaults() *IpmiConfiguration {
	this := IpmiConfiguration{}
	return &this
}

// GetIpmiGateway returns the IpmiGateway field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpmiConfiguration) GetIpmiGateway() string {
	if o == nil || o.IpmiGateway.Get() == nil {
		var ret string
		return ret
	}
	return *o.IpmiGateway.Get()
}

// GetIpmiGatewayOk returns a tuple with the IpmiGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpmiConfiguration) GetIpmiGatewayOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IpmiGateway.Get(), o.IpmiGateway.IsSet()
}

// HasIpmiGateway returns a boolean if a field has been set.
func (o *IpmiConfiguration) HasIpmiGateway() bool {
	if o != nil && o.IpmiGateway.IsSet() {
		return true
	}

	return false
}

// SetIpmiGateway gets a reference to the given NullableString and assigns it to the IpmiGateway field.
func (o *IpmiConfiguration) SetIpmiGateway(v string) {
	o.IpmiGateway.Set(&v)
}
// SetIpmiGatewayNil sets the value for IpmiGateway to be an explicit nil
func (o *IpmiConfiguration) SetIpmiGatewayNil() {
	o.IpmiGateway.Set(nil)
}

// UnsetIpmiGateway ensures that no value is present for IpmiGateway, not even an explicit nil
func (o *IpmiConfiguration) UnsetIpmiGateway() {
	o.IpmiGateway.Unset()
}

// GetIpmiPassword returns the IpmiPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpmiConfiguration) GetIpmiPassword() string {
	if o == nil || o.IpmiPassword.Get() == nil {
		var ret string
		return ret
	}
	return *o.IpmiPassword.Get()
}

// GetIpmiPasswordOk returns a tuple with the IpmiPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpmiConfiguration) GetIpmiPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IpmiPassword.Get(), o.IpmiPassword.IsSet()
}

// HasIpmiPassword returns a boolean if a field has been set.
func (o *IpmiConfiguration) HasIpmiPassword() bool {
	if o != nil && o.IpmiPassword.IsSet() {
		return true
	}

	return false
}

// SetIpmiPassword gets a reference to the given NullableString and assigns it to the IpmiPassword field.
func (o *IpmiConfiguration) SetIpmiPassword(v string) {
	o.IpmiPassword.Set(&v)
}
// SetIpmiPasswordNil sets the value for IpmiPassword to be an explicit nil
func (o *IpmiConfiguration) SetIpmiPasswordNil() {
	o.IpmiPassword.Set(nil)
}

// UnsetIpmiPassword ensures that no value is present for IpmiPassword, not even an explicit nil
func (o *IpmiConfiguration) UnsetIpmiPassword() {
	o.IpmiPassword.Unset()
}

// GetIpmiSubnetMask returns the IpmiSubnetMask field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpmiConfiguration) GetIpmiSubnetMask() string {
	if o == nil || o.IpmiSubnetMask.Get() == nil {
		var ret string
		return ret
	}
	return *o.IpmiSubnetMask.Get()
}

// GetIpmiSubnetMaskOk returns a tuple with the IpmiSubnetMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpmiConfiguration) GetIpmiSubnetMaskOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IpmiSubnetMask.Get(), o.IpmiSubnetMask.IsSet()
}

// HasIpmiSubnetMask returns a boolean if a field has been set.
func (o *IpmiConfiguration) HasIpmiSubnetMask() bool {
	if o != nil && o.IpmiSubnetMask.IsSet() {
		return true
	}

	return false
}

// SetIpmiSubnetMask gets a reference to the given NullableString and assigns it to the IpmiSubnetMask field.
func (o *IpmiConfiguration) SetIpmiSubnetMask(v string) {
	o.IpmiSubnetMask.Set(&v)
}
// SetIpmiSubnetMaskNil sets the value for IpmiSubnetMask to be an explicit nil
func (o *IpmiConfiguration) SetIpmiSubnetMaskNil() {
	o.IpmiSubnetMask.Set(nil)
}

// UnsetIpmiSubnetMask ensures that no value is present for IpmiSubnetMask, not even an explicit nil
func (o *IpmiConfiguration) UnsetIpmiSubnetMask() {
	o.IpmiSubnetMask.Unset()
}

// GetIpmiUsername returns the IpmiUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpmiConfiguration) GetIpmiUsername() string {
	if o == nil || o.IpmiUsername.Get() == nil {
		var ret string
		return ret
	}
	return *o.IpmiUsername.Get()
}

// GetIpmiUsernameOk returns a tuple with the IpmiUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpmiConfiguration) GetIpmiUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IpmiUsername.Get(), o.IpmiUsername.IsSet()
}

// HasIpmiUsername returns a boolean if a field has been set.
func (o *IpmiConfiguration) HasIpmiUsername() bool {
	if o != nil && o.IpmiUsername.IsSet() {
		return true
	}

	return false
}

// SetIpmiUsername gets a reference to the given NullableString and assigns it to the IpmiUsername field.
func (o *IpmiConfiguration) SetIpmiUsername(v string) {
	o.IpmiUsername.Set(&v)
}
// SetIpmiUsernameNil sets the value for IpmiUsername to be an explicit nil
func (o *IpmiConfiguration) SetIpmiUsernameNil() {
	o.IpmiUsername.Set(nil)
}

// UnsetIpmiUsername ensures that no value is present for IpmiUsername, not even an explicit nil
func (o *IpmiConfiguration) UnsetIpmiUsername() {
	o.IpmiUsername.Unset()
}

func (o IpmiConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IpmiGateway.IsSet() {
		toSerialize["ipmiGateway"] = o.IpmiGateway.Get()
	}
	if o.IpmiPassword.IsSet() {
		toSerialize["ipmiPassword"] = o.IpmiPassword.Get()
	}
	if o.IpmiSubnetMask.IsSet() {
		toSerialize["ipmiSubnetMask"] = o.IpmiSubnetMask.Get()
	}
	if o.IpmiUsername.IsSet() {
		toSerialize["ipmiUsername"] = o.IpmiUsername.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIpmiConfiguration struct {
	value *IpmiConfiguration
	isSet bool
}

func (v NullableIpmiConfiguration) Get() *IpmiConfiguration {
	return v.value
}

func (v *NullableIpmiConfiguration) Set(val *IpmiConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableIpmiConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableIpmiConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpmiConfiguration(val *IpmiConfiguration) *NullableIpmiConfiguration {
	return &NullableIpmiConfiguration{value: val, isSet: true}
}

func (v NullableIpmiConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpmiConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


