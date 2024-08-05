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

// checks if the ClusterCreateVirtualParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterCreateVirtualParams{}

// ClusterCreateVirtualParams Params for Virtual Edition Cluster Creation
type ClusterCreateVirtualParams struct {
	Nodes []ClusterCreateNodeParams `json:"nodes,omitempty"`
}

// NewClusterCreateVirtualParams instantiates a new ClusterCreateVirtualParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterCreateVirtualParams() *ClusterCreateVirtualParams {
	this := ClusterCreateVirtualParams{}
	return &this
}

// NewClusterCreateVirtualParamsWithDefaults instantiates a new ClusterCreateVirtualParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterCreateVirtualParamsWithDefaults() *ClusterCreateVirtualParams {
	this := ClusterCreateVirtualParams{}
	return &this
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *ClusterCreateVirtualParams) GetNodes() []ClusterCreateNodeParams {
	if o == nil || IsNil(o.Nodes) {
		var ret []ClusterCreateNodeParams
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterCreateVirtualParams) GetNodesOk() ([]ClusterCreateNodeParams, bool) {
	if o == nil || IsNil(o.Nodes) {
		return nil, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *ClusterCreateVirtualParams) HasNodes() bool {
	if o != nil && !IsNil(o.Nodes) {
		return true
	}

	return false
}

// SetNodes gets a reference to the given []ClusterCreateNodeParams and assigns it to the Nodes field.
func (o *ClusterCreateVirtualParams) SetNodes(v []ClusterCreateNodeParams) {
	o.Nodes = v
}

func (o ClusterCreateVirtualParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterCreateVirtualParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Nodes) {
		toSerialize["nodes"] = o.Nodes
	}
	return toSerialize, nil
}

type NullableClusterCreateVirtualParams struct {
	value *ClusterCreateVirtualParams
	isSet bool
}

func (v NullableClusterCreateVirtualParams) Get() *ClusterCreateVirtualParams {
	return v.value
}

func (v *NullableClusterCreateVirtualParams) Set(val *ClusterCreateVirtualParams) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterCreateVirtualParams) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterCreateVirtualParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterCreateVirtualParams(val *ClusterCreateVirtualParams) *NullableClusterCreateVirtualParams {
	return &NullableClusterCreateVirtualParams{value: val, isSet: true}
}

func (v NullableClusterCreateVirtualParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterCreateVirtualParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


