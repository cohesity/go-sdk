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

// FleetConfig Specifies various resources while deploying fleet params.
type FleetConfig struct {
	// Specifies the subnet type of the fleet.
	FleetSubnetType NullableString `json:"fleetSubnetType,omitempty"`
	// Specifies the network security groups within above VPC.
	FleetNetworkParams NullableFleetNetworkParams `json:"fleetNetworkParams,omitempty"`
	// Specifies the network security groups within above VPC.
	FleetTags []FleetTags `json:"fleetTags,omitempty"`
}

// NewFleetConfig instantiates a new FleetConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFleetConfig() *FleetConfig {
	this := FleetConfig{}
	return &this
}

// NewFleetConfigWithDefaults instantiates a new FleetConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFleetConfigWithDefaults() *FleetConfig {
	this := FleetConfig{}
	return &this
}

// GetFleetSubnetType returns the FleetSubnetType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FleetConfig) GetFleetSubnetType() string {
	if o == nil || o.FleetSubnetType.Get() == nil {
		var ret string
		return ret
	}
	return *o.FleetSubnetType.Get()
}

// GetFleetSubnetTypeOk returns a tuple with the FleetSubnetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FleetConfig) GetFleetSubnetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FleetSubnetType.Get(), o.FleetSubnetType.IsSet()
}

// HasFleetSubnetType returns a boolean if a field has been set.
func (o *FleetConfig) HasFleetSubnetType() bool {
	if o != nil && o.FleetSubnetType.IsSet() {
		return true
	}

	return false
}

// SetFleetSubnetType gets a reference to the given NullableString and assigns it to the FleetSubnetType field.
func (o *FleetConfig) SetFleetSubnetType(v string) {
	o.FleetSubnetType.Set(&v)
}
// SetFleetSubnetTypeNil sets the value for FleetSubnetType to be an explicit nil
func (o *FleetConfig) SetFleetSubnetTypeNil() {
	o.FleetSubnetType.Set(nil)
}

// UnsetFleetSubnetType ensures that no value is present for FleetSubnetType, not even an explicit nil
func (o *FleetConfig) UnsetFleetSubnetType() {
	o.FleetSubnetType.Unset()
}

// GetFleetNetworkParams returns the FleetNetworkParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FleetConfig) GetFleetNetworkParams() FleetNetworkParams {
	if o == nil || o.FleetNetworkParams.Get() == nil {
		var ret FleetNetworkParams
		return ret
	}
	return *o.FleetNetworkParams.Get()
}

// GetFleetNetworkParamsOk returns a tuple with the FleetNetworkParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FleetConfig) GetFleetNetworkParamsOk() (*FleetNetworkParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FleetNetworkParams.Get(), o.FleetNetworkParams.IsSet()
}

// HasFleetNetworkParams returns a boolean if a field has been set.
func (o *FleetConfig) HasFleetNetworkParams() bool {
	if o != nil && o.FleetNetworkParams.IsSet() {
		return true
	}

	return false
}

// SetFleetNetworkParams gets a reference to the given NullableFleetNetworkParams and assigns it to the FleetNetworkParams field.
func (o *FleetConfig) SetFleetNetworkParams(v FleetNetworkParams) {
	o.FleetNetworkParams.Set(&v)
}
// SetFleetNetworkParamsNil sets the value for FleetNetworkParams to be an explicit nil
func (o *FleetConfig) SetFleetNetworkParamsNil() {
	o.FleetNetworkParams.Set(nil)
}

// UnsetFleetNetworkParams ensures that no value is present for FleetNetworkParams, not even an explicit nil
func (o *FleetConfig) UnsetFleetNetworkParams() {
	o.FleetNetworkParams.Unset()
}

// GetFleetTags returns the FleetTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FleetConfig) GetFleetTags() []FleetTags {
	if o == nil  {
		var ret []FleetTags
		return ret
	}
	return o.FleetTags
}

// GetFleetTagsOk returns a tuple with the FleetTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FleetConfig) GetFleetTagsOk() (*[]FleetTags, bool) {
	if o == nil || o.FleetTags == nil {
		return nil, false
	}
	return &o.FleetTags, true
}

// HasFleetTags returns a boolean if a field has been set.
func (o *FleetConfig) HasFleetTags() bool {
	if o != nil && o.FleetTags != nil {
		return true
	}

	return false
}

// SetFleetTags gets a reference to the given []FleetTags and assigns it to the FleetTags field.
func (o *FleetConfig) SetFleetTags(v []FleetTags) {
	o.FleetTags = v
}

func (o FleetConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FleetSubnetType.IsSet() {
		toSerialize["fleetSubnetType"] = o.FleetSubnetType.Get()
	}
	if o.FleetNetworkParams.IsSet() {
		toSerialize["fleetNetworkParams"] = o.FleetNetworkParams.Get()
	}
	if o.FleetTags != nil {
		toSerialize["fleetTags"] = o.FleetTags
	}
	return json.Marshal(toSerialize)
}

type NullableFleetConfig struct {
	value *FleetConfig
	isSet bool
}

func (v NullableFleetConfig) Get() *FleetConfig {
	return v.value
}

func (v *NullableFleetConfig) Set(val *FleetConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableFleetConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableFleetConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFleetConfig(val *FleetConfig) *NullableFleetConfig {
	return &NullableFleetConfig{value: val, isSet: true}
}

func (v NullableFleetConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFleetConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o FleetConfig) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}