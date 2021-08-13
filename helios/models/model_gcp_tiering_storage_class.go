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

// GCPTieringStorageClass Specifies the storage class of GCP External Target of tiering purpose type.
type GCPTieringStorageClass struct {
	// Specifies the storage class of GCP External Target of tiering purpose type.
	Enum *string `json:"enum,omitempty"`
}

// NewGCPTieringStorageClass instantiates a new GCPTieringStorageClass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGCPTieringStorageClass() *GCPTieringStorageClass {
	this := GCPTieringStorageClass{}
	return &this
}

// NewGCPTieringStorageClassWithDefaults instantiates a new GCPTieringStorageClass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGCPTieringStorageClassWithDefaults() *GCPTieringStorageClass {
	this := GCPTieringStorageClass{}
	return &this
}

// GetEnum returns the Enum field value if set, zero value otherwise.
func (o *GCPTieringStorageClass) GetEnum() string {
	if o == nil || o.Enum == nil {
		var ret string
		return ret
	}
	return *o.Enum
}

// GetEnumOk returns a tuple with the Enum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPTieringStorageClass) GetEnumOk() (*string, bool) {
	if o == nil || o.Enum == nil {
		return nil, false
	}
	return o.Enum, true
}

// HasEnum returns a boolean if a field has been set.
func (o *GCPTieringStorageClass) HasEnum() bool {
	if o != nil && o.Enum != nil {
		return true
	}

	return false
}

// SetEnum gets a reference to the given string and assigns it to the Enum field.
func (o *GCPTieringStorageClass) SetEnum(v string) {
	o.Enum = &v
}

func (o GCPTieringStorageClass) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enum != nil {
		toSerialize["enum"] = o.Enum
	}
	return json.Marshal(toSerialize)
}

type NullableGCPTieringStorageClass struct {
	value *GCPTieringStorageClass
	isSet bool
}

func (v NullableGCPTieringStorageClass) Get() *GCPTieringStorageClass {
	return v.value
}

func (v *NullableGCPTieringStorageClass) Set(val *GCPTieringStorageClass) {
	v.value = val
	v.isSet = true
}

func (v NullableGCPTieringStorageClass) IsSet() bool {
	return v.isSet
}

func (v *NullableGCPTieringStorageClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGCPTieringStorageClass(val *GCPTieringStorageClass) *NullableGCPTieringStorageClass {
	return &NullableGCPTieringStorageClass{value: val, isSet: true}
}

func (v NullableGCPTieringStorageClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGCPTieringStorageClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


