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

// HeliosRegConfig Specifies the Helios Registration Config.
type HeliosRegConfig struct {
	// Specifies the type of entity that is registered on Helios.
	EntityType NullableString `json:"entityType,omitempty"`
	RigelRegConfig *RigelRegConfig `json:"rigelRegConfig,omitempty"`
}

// NewHeliosRegConfig instantiates a new HeliosRegConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeliosRegConfig() *HeliosRegConfig {
	this := HeliosRegConfig{}
	return &this
}

// NewHeliosRegConfigWithDefaults instantiates a new HeliosRegConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeliosRegConfigWithDefaults() *HeliosRegConfig {
	this := HeliosRegConfig{}
	return &this
}

// GetEntityType returns the EntityType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosRegConfig) GetEntityType() string {
	if o == nil || o.EntityType.Get() == nil {
		var ret string
		return ret
	}
	return *o.EntityType.Get()
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosRegConfig) GetEntityTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EntityType.Get(), o.EntityType.IsSet()
}

// HasEntityType returns a boolean if a field has been set.
func (o *HeliosRegConfig) HasEntityType() bool {
	if o != nil && o.EntityType.IsSet() {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given NullableString and assigns it to the EntityType field.
func (o *HeliosRegConfig) SetEntityType(v string) {
	o.EntityType.Set(&v)
}
// SetEntityTypeNil sets the value for EntityType to be an explicit nil
func (o *HeliosRegConfig) SetEntityTypeNil() {
	o.EntityType.Set(nil)
}

// UnsetEntityType ensures that no value is present for EntityType, not even an explicit nil
func (o *HeliosRegConfig) UnsetEntityType() {
	o.EntityType.Unset()
}

// GetRigelRegConfig returns the RigelRegConfig field value if set, zero value otherwise.
func (o *HeliosRegConfig) GetRigelRegConfig() RigelRegConfig {
	if o == nil || o.RigelRegConfig == nil {
		var ret RigelRegConfig
		return ret
	}
	return *o.RigelRegConfig
}

// GetRigelRegConfigOk returns a tuple with the RigelRegConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeliosRegConfig) GetRigelRegConfigOk() (*RigelRegConfig, bool) {
	if o == nil || o.RigelRegConfig == nil {
		return nil, false
	}
	return o.RigelRegConfig, true
}

// HasRigelRegConfig returns a boolean if a field has been set.
func (o *HeliosRegConfig) HasRigelRegConfig() bool {
	if o != nil && o.RigelRegConfig != nil {
		return true
	}

	return false
}

// SetRigelRegConfig gets a reference to the given RigelRegConfig and assigns it to the RigelRegConfig field.
func (o *HeliosRegConfig) SetRigelRegConfig(v RigelRegConfig) {
	o.RigelRegConfig = &v
}

func (o HeliosRegConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EntityType.IsSet() {
		toSerialize["entityType"] = o.EntityType.Get()
	}
	if o.RigelRegConfig != nil {
		toSerialize["rigelRegConfig"] = o.RigelRegConfig
	}
	return json.Marshal(toSerialize)
}

type NullableHeliosRegConfig struct {
	value *HeliosRegConfig
	isSet bool
}

func (v NullableHeliosRegConfig) Get() *HeliosRegConfig {
	return v.value
}

func (v *NullableHeliosRegConfig) Set(val *HeliosRegConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableHeliosRegConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableHeliosRegConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeliosRegConfig(val *HeliosRegConfig) *NullableHeliosRegConfig {
	return &NullableHeliosRegConfig{value: val, isSet: true}
}

func (v NullableHeliosRegConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeliosRegConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o HeliosRegConfig) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}