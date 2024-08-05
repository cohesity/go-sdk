/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

// checks if the LogicalVolumeInfoDeviceTree type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogicalVolumeInfoDeviceTree{}

// LogicalVolumeInfoDeviceTree Specifies the tree structure of the logical volume.
type LogicalVolumeInfoDeviceTree struct {
	// Specifies if the node is a leaf node.
	IsLeaf NullableBool `json:"isLeaf,omitempty"`
	LeafNodeParams *DeviceTreeNodeLeafNodeParams `json:"leafNodeParams,omitempty"`
	NonLeafNodeParams *DeviceTreeNodeNonLeafNodeParams `json:"nonLeafNodeParams,omitempty"`
}

// NewLogicalVolumeInfoDeviceTree instantiates a new LogicalVolumeInfoDeviceTree object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogicalVolumeInfoDeviceTree() *LogicalVolumeInfoDeviceTree {
	this := LogicalVolumeInfoDeviceTree{}
	return &this
}

// NewLogicalVolumeInfoDeviceTreeWithDefaults instantiates a new LogicalVolumeInfoDeviceTree object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogicalVolumeInfoDeviceTreeWithDefaults() *LogicalVolumeInfoDeviceTree {
	this := LogicalVolumeInfoDeviceTree{}
	return &this
}

// GetIsLeaf returns the IsLeaf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogicalVolumeInfoDeviceTree) GetIsLeaf() bool {
	if o == nil || IsNil(o.IsLeaf.Get()) {
		var ret bool
		return ret
	}
	return *o.IsLeaf.Get()
}

// GetIsLeafOk returns a tuple with the IsLeaf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogicalVolumeInfoDeviceTree) GetIsLeafOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsLeaf.Get(), o.IsLeaf.IsSet()
}

// HasIsLeaf returns a boolean if a field has been set.
func (o *LogicalVolumeInfoDeviceTree) HasIsLeaf() bool {
	if o != nil && o.IsLeaf.IsSet() {
		return true
	}

	return false
}

// SetIsLeaf gets a reference to the given NullableBool and assigns it to the IsLeaf field.
func (o *LogicalVolumeInfoDeviceTree) SetIsLeaf(v bool) {
	o.IsLeaf.Set(&v)
}
// SetIsLeafNil sets the value for IsLeaf to be an explicit nil
func (o *LogicalVolumeInfoDeviceTree) SetIsLeafNil() {
	o.IsLeaf.Set(nil)
}

// UnsetIsLeaf ensures that no value is present for IsLeaf, not even an explicit nil
func (o *LogicalVolumeInfoDeviceTree) UnsetIsLeaf() {
	o.IsLeaf.Unset()
}

// GetLeafNodeParams returns the LeafNodeParams field value if set, zero value otherwise.
func (o *LogicalVolumeInfoDeviceTree) GetLeafNodeParams() DeviceTreeNodeLeafNodeParams {
	if o == nil || IsNil(o.LeafNodeParams) {
		var ret DeviceTreeNodeLeafNodeParams
		return ret
	}
	return *o.LeafNodeParams
}

// GetLeafNodeParamsOk returns a tuple with the LeafNodeParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogicalVolumeInfoDeviceTree) GetLeafNodeParamsOk() (*DeviceTreeNodeLeafNodeParams, bool) {
	if o == nil || IsNil(o.LeafNodeParams) {
		return nil, false
	}
	return o.LeafNodeParams, true
}

// HasLeafNodeParams returns a boolean if a field has been set.
func (o *LogicalVolumeInfoDeviceTree) HasLeafNodeParams() bool {
	if o != nil && !IsNil(o.LeafNodeParams) {
		return true
	}

	return false
}

// SetLeafNodeParams gets a reference to the given DeviceTreeNodeLeafNodeParams and assigns it to the LeafNodeParams field.
func (o *LogicalVolumeInfoDeviceTree) SetLeafNodeParams(v DeviceTreeNodeLeafNodeParams) {
	o.LeafNodeParams = &v
}

// GetNonLeafNodeParams returns the NonLeafNodeParams field value if set, zero value otherwise.
func (o *LogicalVolumeInfoDeviceTree) GetNonLeafNodeParams() DeviceTreeNodeNonLeafNodeParams {
	if o == nil || IsNil(o.NonLeafNodeParams) {
		var ret DeviceTreeNodeNonLeafNodeParams
		return ret
	}
	return *o.NonLeafNodeParams
}

// GetNonLeafNodeParamsOk returns a tuple with the NonLeafNodeParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogicalVolumeInfoDeviceTree) GetNonLeafNodeParamsOk() (*DeviceTreeNodeNonLeafNodeParams, bool) {
	if o == nil || IsNil(o.NonLeafNodeParams) {
		return nil, false
	}
	return o.NonLeafNodeParams, true
}

// HasNonLeafNodeParams returns a boolean if a field has been set.
func (o *LogicalVolumeInfoDeviceTree) HasNonLeafNodeParams() bool {
	if o != nil && !IsNil(o.NonLeafNodeParams) {
		return true
	}

	return false
}

// SetNonLeafNodeParams gets a reference to the given DeviceTreeNodeNonLeafNodeParams and assigns it to the NonLeafNodeParams field.
func (o *LogicalVolumeInfoDeviceTree) SetNonLeafNodeParams(v DeviceTreeNodeNonLeafNodeParams) {
	o.NonLeafNodeParams = &v
}

func (o LogicalVolumeInfoDeviceTree) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogicalVolumeInfoDeviceTree) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.IsLeaf.IsSet() {
		toSerialize["isLeaf"] = o.IsLeaf.Get()
	}
	if !IsNil(o.LeafNodeParams) {
		toSerialize["leafNodeParams"] = o.LeafNodeParams
	}
	if !IsNil(o.NonLeafNodeParams) {
		toSerialize["nonLeafNodeParams"] = o.NonLeafNodeParams
	}
	return toSerialize, nil
}

type NullableLogicalVolumeInfoDeviceTree struct {
	value *LogicalVolumeInfoDeviceTree
	isSet bool
}

func (v NullableLogicalVolumeInfoDeviceTree) Get() *LogicalVolumeInfoDeviceTree {
	return v.value
}

func (v *NullableLogicalVolumeInfoDeviceTree) Set(val *LogicalVolumeInfoDeviceTree) {
	v.value = val
	v.isSet = true
}

func (v NullableLogicalVolumeInfoDeviceTree) IsSet() bool {
	return v.isSet
}

func (v *NullableLogicalVolumeInfoDeviceTree) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogicalVolumeInfoDeviceTree(val *LogicalVolumeInfoDeviceTree) *NullableLogicalVolumeInfoDeviceTree {
	return &NullableLogicalVolumeInfoDeviceTree{value: val, isSet: true}
}

func (v NullableLogicalVolumeInfoDeviceTree) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogicalVolumeInfoDeviceTree) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


