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

// VcdRegistrationParamsAllOf struct for VcdRegistrationParamsAllOf
type VcdRegistrationParamsAllOf struct {
	// Specifies the credentials information for all the vcenters in vcloud director.
	VcenterCredentialInfoList []VcenterCredentialInfo `json:"vcenterCredentialInfoList"`
}

// NewVcdRegistrationParamsAllOf instantiates a new VcdRegistrationParamsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcdRegistrationParamsAllOf(vcenterCredentialInfoList []VcenterCredentialInfo) *VcdRegistrationParamsAllOf {
	this := VcdRegistrationParamsAllOf{}
	this.VcenterCredentialInfoList = vcenterCredentialInfoList
	return &this
}

// NewVcdRegistrationParamsAllOfWithDefaults instantiates a new VcdRegistrationParamsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcdRegistrationParamsAllOfWithDefaults() *VcdRegistrationParamsAllOf {
	this := VcdRegistrationParamsAllOf{}
	return &this
}

// GetVcenterCredentialInfoList returns the VcenterCredentialInfoList field value
// If the value is explicit nil, the zero value for []VcenterCredentialInfo will be returned
func (o *VcdRegistrationParamsAllOf) GetVcenterCredentialInfoList() []VcenterCredentialInfo {
	if o == nil {
		var ret []VcenterCredentialInfo
		return ret
	}

	return o.VcenterCredentialInfoList
}

// GetVcenterCredentialInfoListOk returns a tuple with the VcenterCredentialInfoList field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VcdRegistrationParamsAllOf) GetVcenterCredentialInfoListOk() (*[]VcenterCredentialInfo, bool) {
	if o == nil || o.VcenterCredentialInfoList == nil {
		return nil, false
	}
	return &o.VcenterCredentialInfoList, true
}

// SetVcenterCredentialInfoList sets field value
func (o *VcdRegistrationParamsAllOf) SetVcenterCredentialInfoList(v []VcenterCredentialInfo) {
	o.VcenterCredentialInfoList = v
}

func (o VcdRegistrationParamsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VcenterCredentialInfoList != nil {
		toSerialize["vcenterCredentialInfoList"] = o.VcenterCredentialInfoList
	}
	return json.Marshal(toSerialize)
}

type NullableVcdRegistrationParamsAllOf struct {
	value *VcdRegistrationParamsAllOf
	isSet bool
}

func (v NullableVcdRegistrationParamsAllOf) Get() *VcdRegistrationParamsAllOf {
	return v.value
}

func (v *NullableVcdRegistrationParamsAllOf) Set(val *VcdRegistrationParamsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVcdRegistrationParamsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVcdRegistrationParamsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcdRegistrationParamsAllOf(val *VcdRegistrationParamsAllOf) *NullableVcdRegistrationParamsAllOf {
	return &NullableVcdRegistrationParamsAllOf{value: val, isSet: true}
}

func (v NullableVcdRegistrationParamsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcdRegistrationParamsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


