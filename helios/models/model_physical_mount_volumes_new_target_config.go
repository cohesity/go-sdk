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

// PhysicalMountVolumesNewTargetConfig Specifies the configuration for mounting volumes to a new target.
type PhysicalMountVolumesNewTargetConfig struct {
	MountTarget RecoverTarget `json:"mountTarget"`
	// Specifies credentials to access the target server. This is required if the server is of Linux OS.
	ServerCredentials NullableCredentials `json:"serverCredentials,omitempty"`
}

// NewPhysicalMountVolumesNewTargetConfig instantiates a new PhysicalMountVolumesNewTargetConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhysicalMountVolumesNewTargetConfig(mountTarget RecoverTarget) *PhysicalMountVolumesNewTargetConfig {
	this := PhysicalMountVolumesNewTargetConfig{}
	this.MountTarget = mountTarget
	return &this
}

// NewPhysicalMountVolumesNewTargetConfigWithDefaults instantiates a new PhysicalMountVolumesNewTargetConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhysicalMountVolumesNewTargetConfigWithDefaults() *PhysicalMountVolumesNewTargetConfig {
	this := PhysicalMountVolumesNewTargetConfig{}
	return &this
}

// GetMountTarget returns the MountTarget field value
func (o *PhysicalMountVolumesNewTargetConfig) GetMountTarget() RecoverTarget {
	if o == nil {
		var ret RecoverTarget
		return ret
	}

	return o.MountTarget
}

// GetMountTargetOk returns a tuple with the MountTarget field value
// and a boolean to check if the value has been set.
func (o *PhysicalMountVolumesNewTargetConfig) GetMountTargetOk() (*RecoverTarget, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MountTarget, true
}

// SetMountTarget sets field value
func (o *PhysicalMountVolumesNewTargetConfig) SetMountTarget(v RecoverTarget) {
	o.MountTarget = v
}

// GetServerCredentials returns the ServerCredentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PhysicalMountVolumesNewTargetConfig) GetServerCredentials() Credentials {
	if o == nil || o.ServerCredentials.Get() == nil {
		var ret Credentials
		return ret
	}
	return *o.ServerCredentials.Get()
}

// GetServerCredentialsOk returns a tuple with the ServerCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhysicalMountVolumesNewTargetConfig) GetServerCredentialsOk() (*Credentials, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServerCredentials.Get(), o.ServerCredentials.IsSet()
}

// HasServerCredentials returns a boolean if a field has been set.
func (o *PhysicalMountVolumesNewTargetConfig) HasServerCredentials() bool {
	if o != nil && o.ServerCredentials.IsSet() {
		return true
	}

	return false
}

// SetServerCredentials gets a reference to the given NullableCredentials and assigns it to the ServerCredentials field.
func (o *PhysicalMountVolumesNewTargetConfig) SetServerCredentials(v Credentials) {
	o.ServerCredentials.Set(&v)
}
// SetServerCredentialsNil sets the value for ServerCredentials to be an explicit nil
func (o *PhysicalMountVolumesNewTargetConfig) SetServerCredentialsNil() {
	o.ServerCredentials.Set(nil)
}

// UnsetServerCredentials ensures that no value is present for ServerCredentials, not even an explicit nil
func (o *PhysicalMountVolumesNewTargetConfig) UnsetServerCredentials() {
	o.ServerCredentials.Unset()
}

func (o PhysicalMountVolumesNewTargetConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mountTarget"] = o.MountTarget
	}
	if o.ServerCredentials.IsSet() {
		toSerialize["serverCredentials"] = o.ServerCredentials.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePhysicalMountVolumesNewTargetConfig struct {
	value *PhysicalMountVolumesNewTargetConfig
	isSet bool
}

func (v NullablePhysicalMountVolumesNewTargetConfig) Get() *PhysicalMountVolumesNewTargetConfig {
	return v.value
}

func (v *NullablePhysicalMountVolumesNewTargetConfig) Set(val *PhysicalMountVolumesNewTargetConfig) {
	v.value = val
	v.isSet = true
}

func (v NullablePhysicalMountVolumesNewTargetConfig) IsSet() bool {
	return v.isSet
}

func (v *NullablePhysicalMountVolumesNewTargetConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhysicalMountVolumesNewTargetConfig(val *PhysicalMountVolumesNewTargetConfig) *NullablePhysicalMountVolumesNewTargetConfig {
	return &NullablePhysicalMountVolumesNewTargetConfig{value: val, isSet: true}
}

func (v NullablePhysicalMountVolumesNewTargetConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhysicalMountVolumesNewTargetConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


