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

// NodeIdentifyParams Specifies the parameter to identify node.
type NodeIdentifyParams struct {
	// Turn on/off node led light if set to true/false respectively.
	Identify NullableBool `json:"identify"`
}

// NewNodeIdentifyParams instantiates a new NodeIdentifyParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeIdentifyParams(identify NullableBool) *NodeIdentifyParams {
	this := NodeIdentifyParams{}
	this.Identify = identify
	return &this
}

// NewNodeIdentifyParamsWithDefaults instantiates a new NodeIdentifyParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeIdentifyParamsWithDefaults() *NodeIdentifyParams {
	this := NodeIdentifyParams{}
	return &this
}

// GetIdentify returns the Identify field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *NodeIdentifyParams) GetIdentify() bool {
	if o == nil || o.Identify.Get() == nil {
		var ret bool
		return ret
	}

	return *o.Identify.Get()
}

// GetIdentifyOk returns a tuple with the Identify field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodeIdentifyParams) GetIdentifyOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Identify.Get(), o.Identify.IsSet()
}

// SetIdentify sets field value
func (o *NodeIdentifyParams) SetIdentify(v bool) {
	o.Identify.Set(&v)
}

func (o NodeIdentifyParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["identify"] = o.Identify.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableNodeIdentifyParams struct {
	value *NodeIdentifyParams
	isSet bool
}

func (v NullableNodeIdentifyParams) Get() *NodeIdentifyParams {
	return v.value
}

func (v *NullableNodeIdentifyParams) Set(val *NodeIdentifyParams) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeIdentifyParams) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeIdentifyParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeIdentifyParams(val *NodeIdentifyParams) *NullableNodeIdentifyParams {
	return &NullableNodeIdentifyParams{value: val, isSet: true}
}

func (v NullableNodeIdentifyParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeIdentifyParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o NodeIdentifyParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}