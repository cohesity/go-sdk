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

// VmwareRecoverDisksTargetSourceConfig Specifies the configuration for restoring disks to a different VM than the one from which the snapshot was taken.
type VmwareRecoverDisksTargetSourceConfig struct {
	// Specifies the source ID of the VM to which the disks will be restored.
	SourceId NullableInt64 `json:"sourceId"`
	// Specifies the disks to be recovered and the location to which they will be recovered.
	Disks []VmwareRecoverTargetSourceDiskParams `json:"disks"`
}

// NewVmwareRecoverDisksTargetSourceConfig instantiates a new VmwareRecoverDisksTargetSourceConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmwareRecoverDisksTargetSourceConfig(sourceId NullableInt64, disks []VmwareRecoverTargetSourceDiskParams) *VmwareRecoverDisksTargetSourceConfig {
	this := VmwareRecoverDisksTargetSourceConfig{}
	this.SourceId = sourceId
	this.Disks = disks
	return &this
}

// NewVmwareRecoverDisksTargetSourceConfigWithDefaults instantiates a new VmwareRecoverDisksTargetSourceConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmwareRecoverDisksTargetSourceConfigWithDefaults() *VmwareRecoverDisksTargetSourceConfig {
	this := VmwareRecoverDisksTargetSourceConfig{}
	return &this
}

// GetSourceId returns the SourceId field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *VmwareRecoverDisksTargetSourceConfig) GetSourceId() int64 {
	if o == nil || o.SourceId.Get() == nil {
		var ret int64
		return ret
	}

	return *o.SourceId.Get()
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareRecoverDisksTargetSourceConfig) GetSourceIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SourceId.Get(), o.SourceId.IsSet()
}

// SetSourceId sets field value
func (o *VmwareRecoverDisksTargetSourceConfig) SetSourceId(v int64) {
	o.SourceId.Set(&v)
}

// GetDisks returns the Disks field value
func (o *VmwareRecoverDisksTargetSourceConfig) GetDisks() []VmwareRecoverTargetSourceDiskParams {
	if o == nil {
		var ret []VmwareRecoverTargetSourceDiskParams
		return ret
	}

	return o.Disks
}

// GetDisksOk returns a tuple with the Disks field value
// and a boolean to check if the value has been set.
func (o *VmwareRecoverDisksTargetSourceConfig) GetDisksOk() (*[]VmwareRecoverTargetSourceDiskParams, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Disks, true
}

// SetDisks sets field value
func (o *VmwareRecoverDisksTargetSourceConfig) SetDisks(v []VmwareRecoverTargetSourceDiskParams) {
	o.Disks = v
}

func (o VmwareRecoverDisksTargetSourceConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sourceId"] = o.SourceId.Get()
	}
	if true {
		toSerialize["disks"] = o.Disks
	}
	return json.Marshal(toSerialize)
}

type NullableVmwareRecoverDisksTargetSourceConfig struct {
	value *VmwareRecoverDisksTargetSourceConfig
	isSet bool
}

func (v NullableVmwareRecoverDisksTargetSourceConfig) Get() *VmwareRecoverDisksTargetSourceConfig {
	return v.value
}

func (v *NullableVmwareRecoverDisksTargetSourceConfig) Set(val *VmwareRecoverDisksTargetSourceConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableVmwareRecoverDisksTargetSourceConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableVmwareRecoverDisksTargetSourceConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmwareRecoverDisksTargetSourceConfig(val *VmwareRecoverDisksTargetSourceConfig) *NullableVmwareRecoverDisksTargetSourceConfig {
	return &NullableVmwareRecoverDisksTargetSourceConfig{value: val, isSet: true}
}

func (v NullableVmwareRecoverDisksTargetSourceConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmwareRecoverDisksTargetSourceConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


