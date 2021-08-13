/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package onprem
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/onprem/utils"
	"fmt"
)

var _ = NullableBool{}

// RecoverAwsAuroraNewSourceConfig Specifies the new destination Source configuration where the Aurora instances will be recovered.
type RecoverAwsAuroraNewSourceConfig struct {
	// Specifies the id of the parent source to recover the Aurora.
	Source NullableRecoveryObjectIdentifier `json:"source"`
	// Specifies the AWS region in which to deploy the Aurora instance.
	Region NullableRecoveryObjectIdentifier `json:"region"`
	// Specifies the networking configuration to be applied to the recovered VMs.
	NetworkConfig NullableRecoverAwsAuroraNewSourceNetworkConfig `json:"networkConfig"`
}

// NewRecoverAwsAuroraNewSourceConfig instantiates a new RecoverAwsAuroraNewSourceConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverAwsAuroraNewSourceConfig(source NullableRecoveryObjectIdentifier, region NullableRecoveryObjectIdentifier, networkConfig NullableRecoverAwsAuroraNewSourceNetworkConfig) *RecoverAwsAuroraNewSourceConfig {
	this := RecoverAwsAuroraNewSourceConfig{}
	this.Source = source
	this.Region = region
	this.NetworkConfig = networkConfig
	return &this
}

// NewRecoverAwsAuroraNewSourceConfigWithDefaults instantiates a new RecoverAwsAuroraNewSourceConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverAwsAuroraNewSourceConfigWithDefaults() *RecoverAwsAuroraNewSourceConfig {
	this := RecoverAwsAuroraNewSourceConfig{}
	return &this
}

// GetSource returns the Source field value
// If the value is explicit nil, the zero value for RecoveryObjectIdentifier will be returned
func (o *RecoverAwsAuroraNewSourceConfig) GetSource() RecoveryObjectIdentifier {
	if o == nil || o.Source.Get() == nil {
		var ret RecoveryObjectIdentifier
		return ret
	}

	return *o.Source.Get()
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverAwsAuroraNewSourceConfig) GetSourceOk() (*RecoveryObjectIdentifier, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Source.Get(), o.Source.IsSet()
}

// SetSource sets field value
func (o *RecoverAwsAuroraNewSourceConfig) SetSource(v RecoveryObjectIdentifier) {
	o.Source.Set(&v)
}

// GetRegion returns the Region field value
// If the value is explicit nil, the zero value for RecoveryObjectIdentifier will be returned
func (o *RecoverAwsAuroraNewSourceConfig) GetRegion() RecoveryObjectIdentifier {
	if o == nil || o.Region.Get() == nil {
		var ret RecoveryObjectIdentifier
		return ret
	}

	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverAwsAuroraNewSourceConfig) GetRegionOk() (*RecoveryObjectIdentifier, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// SetRegion sets field value
func (o *RecoverAwsAuroraNewSourceConfig) SetRegion(v RecoveryObjectIdentifier) {
	o.Region.Set(&v)
}

// GetNetworkConfig returns the NetworkConfig field value
// If the value is explicit nil, the zero value for RecoverAwsAuroraNewSourceNetworkConfig will be returned
func (o *RecoverAwsAuroraNewSourceConfig) GetNetworkConfig() RecoverAwsAuroraNewSourceNetworkConfig {
	if o == nil || o.NetworkConfig.Get() == nil {
		var ret RecoverAwsAuroraNewSourceNetworkConfig
		return ret
	}

	return *o.NetworkConfig.Get()
}

// GetNetworkConfigOk returns a tuple with the NetworkConfig field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverAwsAuroraNewSourceConfig) GetNetworkConfigOk() (*RecoverAwsAuroraNewSourceNetworkConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NetworkConfig.Get(), o.NetworkConfig.IsSet()
}

// SetNetworkConfig sets field value
func (o *RecoverAwsAuroraNewSourceConfig) SetNetworkConfig(v RecoverAwsAuroraNewSourceNetworkConfig) {
	o.NetworkConfig.Set(&v)
}

func (o RecoverAwsAuroraNewSourceConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["source"] = o.Source.Get()
	}
	if true {
		toSerialize["region"] = o.Region.Get()
	}
	if true {
		toSerialize["networkConfig"] = o.NetworkConfig.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRecoverAwsAuroraNewSourceConfig struct {
	value *RecoverAwsAuroraNewSourceConfig
	isSet bool
}

func (v NullableRecoverAwsAuroraNewSourceConfig) Get() *RecoverAwsAuroraNewSourceConfig {
	return v.value
}

func (v *NullableRecoverAwsAuroraNewSourceConfig) Set(val *RecoverAwsAuroraNewSourceConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverAwsAuroraNewSourceConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverAwsAuroraNewSourceConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverAwsAuroraNewSourceConfig(val *RecoverAwsAuroraNewSourceConfig) *NullableRecoverAwsAuroraNewSourceConfig {
	return &NullableRecoverAwsAuroraNewSourceConfig{value: val, isSet: true}
}

func (v NullableRecoverAwsAuroraNewSourceConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverAwsAuroraNewSourceConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o RecoverAwsAuroraNewSourceConfig) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}