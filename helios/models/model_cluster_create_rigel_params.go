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

// ClusterCreateRigelParams Params for Rigel Cluster Creation
type ClusterCreateRigelParams struct {
	Nodes []RigelClusterNode `json:"nodes"`
	// Specifies the token which will be used to claim this Cluster to Helios.
	ClaimToken NullableString `json:"claimToken"`
}

// NewClusterCreateRigelParams instantiates a new ClusterCreateRigelParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterCreateRigelParams(nodes []RigelClusterNode, claimToken NullableString) *ClusterCreateRigelParams {
	this := ClusterCreateRigelParams{}
	this.Nodes = nodes
	this.ClaimToken = claimToken
	return &this
}

// NewClusterCreateRigelParamsWithDefaults instantiates a new ClusterCreateRigelParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterCreateRigelParamsWithDefaults() *ClusterCreateRigelParams {
	this := ClusterCreateRigelParams{}
	return &this
}

// GetNodes returns the Nodes field value
func (o *ClusterCreateRigelParams) GetNodes() []RigelClusterNode {
	if o == nil {
		var ret []RigelClusterNode
		return ret
	}

	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value
// and a boolean to check if the value has been set.
func (o *ClusterCreateRigelParams) GetNodesOk() (*[]RigelClusterNode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Nodes, true
}

// SetNodes sets field value
func (o *ClusterCreateRigelParams) SetNodes(v []RigelClusterNode) {
	o.Nodes = v
}

// GetClaimToken returns the ClaimToken field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ClusterCreateRigelParams) GetClaimToken() string {
	if o == nil || o.ClaimToken.Get() == nil {
		var ret string
		return ret
	}

	return *o.ClaimToken.Get()
}

// GetClaimTokenOk returns a tuple with the ClaimToken field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterCreateRigelParams) GetClaimTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClaimToken.Get(), o.ClaimToken.IsSet()
}

// SetClaimToken sets field value
func (o *ClusterCreateRigelParams) SetClaimToken(v string) {
	o.ClaimToken.Set(&v)
}

func (o ClusterCreateRigelParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["nodes"] = o.Nodes
	}
	if true {
		toSerialize["claimToken"] = o.ClaimToken.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableClusterCreateRigelParams struct {
	value *ClusterCreateRigelParams
	isSet bool
}

func (v NullableClusterCreateRigelParams) Get() *ClusterCreateRigelParams {
	return v.value
}

func (v *NullableClusterCreateRigelParams) Set(val *ClusterCreateRigelParams) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterCreateRigelParams) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterCreateRigelParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterCreateRigelParams(val *ClusterCreateRigelParams) *NullableClusterCreateRigelParams {
	return &NullableClusterCreateRigelParams{value: val, isSet: true}
}

func (v NullableClusterCreateRigelParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterCreateRigelParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


