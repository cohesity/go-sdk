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

// GcpProtectionGroupParams Specifies the parameters which are specific to GCP related Protection Groups.
type GcpProtectionGroupParams struct {
	// Specifies the GCP Protection Group type.
	ProtectionType string `json:"protectionType"`
	NativeProtectionTypeParams *GcpNativeProtectionGroupParams `json:"nativeProtectionTypeParams,omitempty"`
}

// NewGcpProtectionGroupParams instantiates a new GcpProtectionGroupParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGcpProtectionGroupParams(protectionType string) *GcpProtectionGroupParams {
	this := GcpProtectionGroupParams{}
	this.ProtectionType = protectionType
	return &this
}

// NewGcpProtectionGroupParamsWithDefaults instantiates a new GcpProtectionGroupParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGcpProtectionGroupParamsWithDefaults() *GcpProtectionGroupParams {
	this := GcpProtectionGroupParams{}
	return &this
}

// GetProtectionType returns the ProtectionType field value
func (o *GcpProtectionGroupParams) GetProtectionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProtectionType
}

// GetProtectionTypeOk returns a tuple with the ProtectionType field value
// and a boolean to check if the value has been set.
func (o *GcpProtectionGroupParams) GetProtectionTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProtectionType, true
}

// SetProtectionType sets field value
func (o *GcpProtectionGroupParams) SetProtectionType(v string) {
	o.ProtectionType = v
}

// GetNativeProtectionTypeParams returns the NativeProtectionTypeParams field value if set, zero value otherwise.
func (o *GcpProtectionGroupParams) GetNativeProtectionTypeParams() GcpNativeProtectionGroupParams {
	if o == nil || o.NativeProtectionTypeParams == nil {
		var ret GcpNativeProtectionGroupParams
		return ret
	}
	return *o.NativeProtectionTypeParams
}

// GetNativeProtectionTypeParamsOk returns a tuple with the NativeProtectionTypeParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpProtectionGroupParams) GetNativeProtectionTypeParamsOk() (*GcpNativeProtectionGroupParams, bool) {
	if o == nil || o.NativeProtectionTypeParams == nil {
		return nil, false
	}
	return o.NativeProtectionTypeParams, true
}

// HasNativeProtectionTypeParams returns a boolean if a field has been set.
func (o *GcpProtectionGroupParams) HasNativeProtectionTypeParams() bool {
	if o != nil && o.NativeProtectionTypeParams != nil {
		return true
	}

	return false
}

// SetNativeProtectionTypeParams gets a reference to the given GcpNativeProtectionGroupParams and assigns it to the NativeProtectionTypeParams field.
func (o *GcpProtectionGroupParams) SetNativeProtectionTypeParams(v GcpNativeProtectionGroupParams) {
	o.NativeProtectionTypeParams = &v
}

func (o GcpProtectionGroupParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["protectionType"] = o.ProtectionType
	}
	if o.NativeProtectionTypeParams != nil {
		toSerialize["nativeProtectionTypeParams"] = o.NativeProtectionTypeParams
	}
	return json.Marshal(toSerialize)
}

type NullableGcpProtectionGroupParams struct {
	value *GcpProtectionGroupParams
	isSet bool
}

func (v NullableGcpProtectionGroupParams) Get() *GcpProtectionGroupParams {
	return v.value
}

func (v *NullableGcpProtectionGroupParams) Set(val *GcpProtectionGroupParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpProtectionGroupParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpProtectionGroupParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpProtectionGroupParams(val *GcpProtectionGroupParams) *NullableGcpProtectionGroupParams {
	return &NullableGcpProtectionGroupParams{value: val, isSet: true}
}

func (v NullableGcpProtectionGroupParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpProtectionGroupParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o GcpProtectionGroupParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}