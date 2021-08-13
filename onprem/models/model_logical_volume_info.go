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

// LogicalVolumeInfo Specifies the logical volume info for LVM or LDM volume.
type LogicalVolumeInfo struct {
	// Specifies the volume group uuid.
	VolumeGroupUuid NullableString `json:"volumeGroupUuid,omitempty"`
	// Specifies the volume group name.
	VolumeGroupName NullableString `json:"volumeGroupName,omitempty"`
	// Specifies the logical volume uuid.
	LogicalVolumeUuid NullableString `json:"logicalVolumeUuid,omitempty"`
	// Specifies the logical volume name.
	LogicalVolumeName NullableString `json:"logicalVolumeName,omitempty"`
	// Specifies the tree structure of the logical volume.
	DeviceTree *DeviceTreeNode `json:"deviceTree,omitempty"`
}

// NewLogicalVolumeInfo instantiates a new LogicalVolumeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogicalVolumeInfo() *LogicalVolumeInfo {
	this := LogicalVolumeInfo{}
	return &this
}

// NewLogicalVolumeInfoWithDefaults instantiates a new LogicalVolumeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogicalVolumeInfoWithDefaults() *LogicalVolumeInfo {
	this := LogicalVolumeInfo{}
	return &this
}

// GetVolumeGroupUuid returns the VolumeGroupUuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogicalVolumeInfo) GetVolumeGroupUuid() string {
	if o == nil || o.VolumeGroupUuid.Get() == nil {
		var ret string
		return ret
	}
	return *o.VolumeGroupUuid.Get()
}

// GetVolumeGroupUuidOk returns a tuple with the VolumeGroupUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogicalVolumeInfo) GetVolumeGroupUuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VolumeGroupUuid.Get(), o.VolumeGroupUuid.IsSet()
}

// HasVolumeGroupUuid returns a boolean if a field has been set.
func (o *LogicalVolumeInfo) HasVolumeGroupUuid() bool {
	if o != nil && o.VolumeGroupUuid.IsSet() {
		return true
	}

	return false
}

// SetVolumeGroupUuid gets a reference to the given NullableString and assigns it to the VolumeGroupUuid field.
func (o *LogicalVolumeInfo) SetVolumeGroupUuid(v string) {
	o.VolumeGroupUuid.Set(&v)
}
// SetVolumeGroupUuidNil sets the value for VolumeGroupUuid to be an explicit nil
func (o *LogicalVolumeInfo) SetVolumeGroupUuidNil() {
	o.VolumeGroupUuid.Set(nil)
}

// UnsetVolumeGroupUuid ensures that no value is present for VolumeGroupUuid, not even an explicit nil
func (o *LogicalVolumeInfo) UnsetVolumeGroupUuid() {
	o.VolumeGroupUuid.Unset()
}

// GetVolumeGroupName returns the VolumeGroupName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogicalVolumeInfo) GetVolumeGroupName() string {
	if o == nil || o.VolumeGroupName.Get() == nil {
		var ret string
		return ret
	}
	return *o.VolumeGroupName.Get()
}

// GetVolumeGroupNameOk returns a tuple with the VolumeGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogicalVolumeInfo) GetVolumeGroupNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VolumeGroupName.Get(), o.VolumeGroupName.IsSet()
}

// HasVolumeGroupName returns a boolean if a field has been set.
func (o *LogicalVolumeInfo) HasVolumeGroupName() bool {
	if o != nil && o.VolumeGroupName.IsSet() {
		return true
	}

	return false
}

// SetVolumeGroupName gets a reference to the given NullableString and assigns it to the VolumeGroupName field.
func (o *LogicalVolumeInfo) SetVolumeGroupName(v string) {
	o.VolumeGroupName.Set(&v)
}
// SetVolumeGroupNameNil sets the value for VolumeGroupName to be an explicit nil
func (o *LogicalVolumeInfo) SetVolumeGroupNameNil() {
	o.VolumeGroupName.Set(nil)
}

// UnsetVolumeGroupName ensures that no value is present for VolumeGroupName, not even an explicit nil
func (o *LogicalVolumeInfo) UnsetVolumeGroupName() {
	o.VolumeGroupName.Unset()
}

// GetLogicalVolumeUuid returns the LogicalVolumeUuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogicalVolumeInfo) GetLogicalVolumeUuid() string {
	if o == nil || o.LogicalVolumeUuid.Get() == nil {
		var ret string
		return ret
	}
	return *o.LogicalVolumeUuid.Get()
}

// GetLogicalVolumeUuidOk returns a tuple with the LogicalVolumeUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogicalVolumeInfo) GetLogicalVolumeUuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LogicalVolumeUuid.Get(), o.LogicalVolumeUuid.IsSet()
}

// HasLogicalVolumeUuid returns a boolean if a field has been set.
func (o *LogicalVolumeInfo) HasLogicalVolumeUuid() bool {
	if o != nil && o.LogicalVolumeUuid.IsSet() {
		return true
	}

	return false
}

// SetLogicalVolumeUuid gets a reference to the given NullableString and assigns it to the LogicalVolumeUuid field.
func (o *LogicalVolumeInfo) SetLogicalVolumeUuid(v string) {
	o.LogicalVolumeUuid.Set(&v)
}
// SetLogicalVolumeUuidNil sets the value for LogicalVolumeUuid to be an explicit nil
func (o *LogicalVolumeInfo) SetLogicalVolumeUuidNil() {
	o.LogicalVolumeUuid.Set(nil)
}

// UnsetLogicalVolumeUuid ensures that no value is present for LogicalVolumeUuid, not even an explicit nil
func (o *LogicalVolumeInfo) UnsetLogicalVolumeUuid() {
	o.LogicalVolumeUuid.Unset()
}

// GetLogicalVolumeName returns the LogicalVolumeName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogicalVolumeInfo) GetLogicalVolumeName() string {
	if o == nil || o.LogicalVolumeName.Get() == nil {
		var ret string
		return ret
	}
	return *o.LogicalVolumeName.Get()
}

// GetLogicalVolumeNameOk returns a tuple with the LogicalVolumeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogicalVolumeInfo) GetLogicalVolumeNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LogicalVolumeName.Get(), o.LogicalVolumeName.IsSet()
}

// HasLogicalVolumeName returns a boolean if a field has been set.
func (o *LogicalVolumeInfo) HasLogicalVolumeName() bool {
	if o != nil && o.LogicalVolumeName.IsSet() {
		return true
	}

	return false
}

// SetLogicalVolumeName gets a reference to the given NullableString and assigns it to the LogicalVolumeName field.
func (o *LogicalVolumeInfo) SetLogicalVolumeName(v string) {
	o.LogicalVolumeName.Set(&v)
}
// SetLogicalVolumeNameNil sets the value for LogicalVolumeName to be an explicit nil
func (o *LogicalVolumeInfo) SetLogicalVolumeNameNil() {
	o.LogicalVolumeName.Set(nil)
}

// UnsetLogicalVolumeName ensures that no value is present for LogicalVolumeName, not even an explicit nil
func (o *LogicalVolumeInfo) UnsetLogicalVolumeName() {
	o.LogicalVolumeName.Unset()
}

// GetDeviceTree returns the DeviceTree field value if set, zero value otherwise.
func (o *LogicalVolumeInfo) GetDeviceTree() DeviceTreeNode {
	if o == nil || o.DeviceTree == nil {
		var ret DeviceTreeNode
		return ret
	}
	return *o.DeviceTree
}

// GetDeviceTreeOk returns a tuple with the DeviceTree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogicalVolumeInfo) GetDeviceTreeOk() (*DeviceTreeNode, bool) {
	if o == nil || o.DeviceTree == nil {
		return nil, false
	}
	return o.DeviceTree, true
}

// HasDeviceTree returns a boolean if a field has been set.
func (o *LogicalVolumeInfo) HasDeviceTree() bool {
	if o != nil && o.DeviceTree != nil {
		return true
	}

	return false
}

// SetDeviceTree gets a reference to the given DeviceTreeNode and assigns it to the DeviceTree field.
func (o *LogicalVolumeInfo) SetDeviceTree(v DeviceTreeNode) {
	o.DeviceTree = &v
}

func (o LogicalVolumeInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VolumeGroupUuid.IsSet() {
		toSerialize["volumeGroupUuid"] = o.VolumeGroupUuid.Get()
	}
	if o.VolumeGroupName.IsSet() {
		toSerialize["volumeGroupName"] = o.VolumeGroupName.Get()
	}
	if o.LogicalVolumeUuid.IsSet() {
		toSerialize["logicalVolumeUuid"] = o.LogicalVolumeUuid.Get()
	}
	if o.LogicalVolumeName.IsSet() {
		toSerialize["logicalVolumeName"] = o.LogicalVolumeName.Get()
	}
	if o.DeviceTree != nil {
		toSerialize["deviceTree"] = o.DeviceTree
	}
	return json.Marshal(toSerialize)
}

type NullableLogicalVolumeInfo struct {
	value *LogicalVolumeInfo
	isSet bool
}

func (v NullableLogicalVolumeInfo) Get() *LogicalVolumeInfo {
	return v.value
}

func (v *NullableLogicalVolumeInfo) Set(val *LogicalVolumeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableLogicalVolumeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableLogicalVolumeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogicalVolumeInfo(val *LogicalVolumeInfo) *NullableLogicalVolumeInfo {
	return &NullableLogicalVolumeInfo{value: val, isSet: true}
}

func (v NullableLogicalVolumeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogicalVolumeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o LogicalVolumeInfo) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}