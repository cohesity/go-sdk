/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RecoverAwsVmNewSourceConfigNetworkConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoverAwsVmNewSourceConfigNetworkConfig{}

// RecoverAwsVmNewSourceConfigNetworkConfig Specifies the networking configuration to be applied to the recovered VMs.
type RecoverAwsVmNewSourceConfigNetworkConfig struct {
	// Specifies the network security groups within above VPC.
	SecurityGroups []RecoveryObjectIdentifier `json:"securityGroups"`
	Subnet NullableRecoverAwsAuroraNewSourceNetworkConfigSubnet `json:"subnet"`
	Vpc NullableRecoverAwsAuroraNewSourceNetworkConfigVpc `json:"vpc"`
}

type _RecoverAwsVmNewSourceConfigNetworkConfig RecoverAwsVmNewSourceConfigNetworkConfig

// NewRecoverAwsVmNewSourceConfigNetworkConfig instantiates a new RecoverAwsVmNewSourceConfigNetworkConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverAwsVmNewSourceConfigNetworkConfig(securityGroups []RecoveryObjectIdentifier, subnet NullableRecoverAwsAuroraNewSourceNetworkConfigSubnet, vpc NullableRecoverAwsAuroraNewSourceNetworkConfigVpc) *RecoverAwsVmNewSourceConfigNetworkConfig {
	this := RecoverAwsVmNewSourceConfigNetworkConfig{}
	this.SecurityGroups = securityGroups
	this.Subnet = subnet
	this.Vpc = vpc
	return &this
}

// NewRecoverAwsVmNewSourceConfigNetworkConfigWithDefaults instantiates a new RecoverAwsVmNewSourceConfigNetworkConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverAwsVmNewSourceConfigNetworkConfigWithDefaults() *RecoverAwsVmNewSourceConfigNetworkConfig {
	this := RecoverAwsVmNewSourceConfigNetworkConfig{}
	return &this
}

// GetSecurityGroups returns the SecurityGroups field value
// If the value is explicit nil, the zero value for []RecoveryObjectIdentifier will be returned
func (o *RecoverAwsVmNewSourceConfigNetworkConfig) GetSecurityGroups() []RecoveryObjectIdentifier {
	if o == nil {
		var ret []RecoveryObjectIdentifier
		return ret
	}

	return o.SecurityGroups
}

// GetSecurityGroupsOk returns a tuple with the SecurityGroups field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverAwsVmNewSourceConfigNetworkConfig) GetSecurityGroupsOk() ([]RecoveryObjectIdentifier, bool) {
	if o == nil || IsNil(o.SecurityGroups) {
		return nil, false
	}
	return o.SecurityGroups, true
}

// SetSecurityGroups sets field value
func (o *RecoverAwsVmNewSourceConfigNetworkConfig) SetSecurityGroups(v []RecoveryObjectIdentifier) {
	o.SecurityGroups = v
}

// GetSubnet returns the Subnet field value
// If the value is explicit nil, the zero value for RecoverAwsAuroraNewSourceNetworkConfigSubnet will be returned
func (o *RecoverAwsVmNewSourceConfigNetworkConfig) GetSubnet() RecoverAwsAuroraNewSourceNetworkConfigSubnet {
	if o == nil || o.Subnet.Get() == nil {
		var ret RecoverAwsAuroraNewSourceNetworkConfigSubnet
		return ret
	}

	return *o.Subnet.Get()
}

// GetSubnetOk returns a tuple with the Subnet field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverAwsVmNewSourceConfigNetworkConfig) GetSubnetOk() (*RecoverAwsAuroraNewSourceNetworkConfigSubnet, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subnet.Get(), o.Subnet.IsSet()
}

// SetSubnet sets field value
func (o *RecoverAwsVmNewSourceConfigNetworkConfig) SetSubnet(v RecoverAwsAuroraNewSourceNetworkConfigSubnet) {
	o.Subnet.Set(&v)
}

// GetVpc returns the Vpc field value
// If the value is explicit nil, the zero value for RecoverAwsAuroraNewSourceNetworkConfigVpc will be returned
func (o *RecoverAwsVmNewSourceConfigNetworkConfig) GetVpc() RecoverAwsAuroraNewSourceNetworkConfigVpc {
	if o == nil || o.Vpc.Get() == nil {
		var ret RecoverAwsAuroraNewSourceNetworkConfigVpc
		return ret
	}

	return *o.Vpc.Get()
}

// GetVpcOk returns a tuple with the Vpc field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverAwsVmNewSourceConfigNetworkConfig) GetVpcOk() (*RecoverAwsAuroraNewSourceNetworkConfigVpc, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vpc.Get(), o.Vpc.IsSet()
}

// SetVpc sets field value
func (o *RecoverAwsVmNewSourceConfigNetworkConfig) SetVpc(v RecoverAwsAuroraNewSourceNetworkConfigVpc) {
	o.Vpc.Set(&v)
}

func (o RecoverAwsVmNewSourceConfigNetworkConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoverAwsVmNewSourceConfigNetworkConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SecurityGroups != nil {
		toSerialize["securityGroups"] = o.SecurityGroups
	}
	toSerialize["subnet"] = o.Subnet.Get()
	toSerialize["vpc"] = o.Vpc.Get()
	return toSerialize, nil
}

func (o *RecoverAwsVmNewSourceConfigNetworkConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"securityGroups",
		"subnet",
		"vpc",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRecoverAwsVmNewSourceConfigNetworkConfig := _RecoverAwsVmNewSourceConfigNetworkConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRecoverAwsVmNewSourceConfigNetworkConfig)

	if err != nil {
		return err
	}

	*o = RecoverAwsVmNewSourceConfigNetworkConfig(varRecoverAwsVmNewSourceConfigNetworkConfig)

	return err
}

type NullableRecoverAwsVmNewSourceConfigNetworkConfig struct {
	value *RecoverAwsVmNewSourceConfigNetworkConfig
	isSet bool
}

func (v NullableRecoverAwsVmNewSourceConfigNetworkConfig) Get() *RecoverAwsVmNewSourceConfigNetworkConfig {
	return v.value
}

func (v *NullableRecoverAwsVmNewSourceConfigNetworkConfig) Set(val *RecoverAwsVmNewSourceConfigNetworkConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverAwsVmNewSourceConfigNetworkConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverAwsVmNewSourceConfigNetworkConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverAwsVmNewSourceConfigNetworkConfig(val *RecoverAwsVmNewSourceConfigNetworkConfig) *NullableRecoverAwsVmNewSourceConfigNetworkConfig {
	return &NullableRecoverAwsVmNewSourceConfigNetworkConfig{value: val, isSet: true}
}

func (v NullableRecoverAwsVmNewSourceConfigNetworkConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverAwsVmNewSourceConfigNetworkConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


