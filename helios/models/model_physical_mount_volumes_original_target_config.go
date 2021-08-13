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

// PhysicalMountVolumesOriginalTargetConfig Specifies the configuration for mounting volumes to the original target.
type PhysicalMountVolumesOriginalTargetConfig struct {
	// Specifies credentials to access the target server. This is required if the server is of Linux OS.
	ServerCredentials NullableCredentials `json:"serverCredentials,omitempty"`
}

// NewPhysicalMountVolumesOriginalTargetConfig instantiates a new PhysicalMountVolumesOriginalTargetConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhysicalMountVolumesOriginalTargetConfig() *PhysicalMountVolumesOriginalTargetConfig {
	this := PhysicalMountVolumesOriginalTargetConfig{}
	return &this
}

// NewPhysicalMountVolumesOriginalTargetConfigWithDefaults instantiates a new PhysicalMountVolumesOriginalTargetConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhysicalMountVolumesOriginalTargetConfigWithDefaults() *PhysicalMountVolumesOriginalTargetConfig {
	this := PhysicalMountVolumesOriginalTargetConfig{}
	return &this
}

// GetServerCredentials returns the ServerCredentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PhysicalMountVolumesOriginalTargetConfig) GetServerCredentials() Credentials {
	if o == nil || o.ServerCredentials.Get() == nil {
		var ret Credentials
		return ret
	}
	return *o.ServerCredentials.Get()
}

// GetServerCredentialsOk returns a tuple with the ServerCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhysicalMountVolumesOriginalTargetConfig) GetServerCredentialsOk() (*Credentials, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServerCredentials.Get(), o.ServerCredentials.IsSet()
}

// HasServerCredentials returns a boolean if a field has been set.
func (o *PhysicalMountVolumesOriginalTargetConfig) HasServerCredentials() bool {
	if o != nil && o.ServerCredentials.IsSet() {
		return true
	}

	return false
}

// SetServerCredentials gets a reference to the given NullableCredentials and assigns it to the ServerCredentials field.
func (o *PhysicalMountVolumesOriginalTargetConfig) SetServerCredentials(v Credentials) {
	o.ServerCredentials.Set(&v)
}
// SetServerCredentialsNil sets the value for ServerCredentials to be an explicit nil
func (o *PhysicalMountVolumesOriginalTargetConfig) SetServerCredentialsNil() {
	o.ServerCredentials.Set(nil)
}

// UnsetServerCredentials ensures that no value is present for ServerCredentials, not even an explicit nil
func (o *PhysicalMountVolumesOriginalTargetConfig) UnsetServerCredentials() {
	o.ServerCredentials.Unset()
}

func (o PhysicalMountVolumesOriginalTargetConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ServerCredentials.IsSet() {
		toSerialize["serverCredentials"] = o.ServerCredentials.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePhysicalMountVolumesOriginalTargetConfig struct {
	value *PhysicalMountVolumesOriginalTargetConfig
	isSet bool
}

func (v NullablePhysicalMountVolumesOriginalTargetConfig) Get() *PhysicalMountVolumesOriginalTargetConfig {
	return v.value
}

func (v *NullablePhysicalMountVolumesOriginalTargetConfig) Set(val *PhysicalMountVolumesOriginalTargetConfig) {
	v.value = val
	v.isSet = true
}

func (v NullablePhysicalMountVolumesOriginalTargetConfig) IsSet() bool {
	return v.isSet
}

func (v *NullablePhysicalMountVolumesOriginalTargetConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhysicalMountVolumesOriginalTargetConfig(val *PhysicalMountVolumesOriginalTargetConfig) *NullablePhysicalMountVolumesOriginalTargetConfig {
	return &NullablePhysicalMountVolumesOriginalTargetConfig{value: val, isSet: true}
}

func (v NullablePhysicalMountVolumesOriginalTargetConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhysicalMountVolumesOriginalTargetConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


