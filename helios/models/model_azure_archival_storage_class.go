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

// AzureArchivalStorageClass Specifies the storage class of Azure External Target of archival purpose type.
type AzureArchivalStorageClass struct {
	// Specifies the storage class of Azure External Target of archival purpose type.
	Enum *string `json:"enum,omitempty"`
}

// NewAzureArchivalStorageClass instantiates a new AzureArchivalStorageClass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureArchivalStorageClass() *AzureArchivalStorageClass {
	this := AzureArchivalStorageClass{}
	return &this
}

// NewAzureArchivalStorageClassWithDefaults instantiates a new AzureArchivalStorageClass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureArchivalStorageClassWithDefaults() *AzureArchivalStorageClass {
	this := AzureArchivalStorageClass{}
	return &this
}

// GetEnum returns the Enum field value if set, zero value otherwise.
func (o *AzureArchivalStorageClass) GetEnum() string {
	if o == nil || o.Enum == nil {
		var ret string
		return ret
	}
	return *o.Enum
}

// GetEnumOk returns a tuple with the Enum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureArchivalStorageClass) GetEnumOk() (*string, bool) {
	if o == nil || o.Enum == nil {
		return nil, false
	}
	return o.Enum, true
}

// HasEnum returns a boolean if a field has been set.
func (o *AzureArchivalStorageClass) HasEnum() bool {
	if o != nil && o.Enum != nil {
		return true
	}

	return false
}

// SetEnum gets a reference to the given string and assigns it to the Enum field.
func (o *AzureArchivalStorageClass) SetEnum(v string) {
	o.Enum = &v
}

func (o AzureArchivalStorageClass) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enum != nil {
		toSerialize["enum"] = o.Enum
	}
	return json.Marshal(toSerialize)
}

type NullableAzureArchivalStorageClass struct {
	value *AzureArchivalStorageClass
	isSet bool
}

func (v NullableAzureArchivalStorageClass) Get() *AzureArchivalStorageClass {
	return v.value
}

func (v *NullableAzureArchivalStorageClass) Set(val *AzureArchivalStorageClass) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureArchivalStorageClass) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureArchivalStorageClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureArchivalStorageClass(val *AzureArchivalStorageClass) *NullableAzureArchivalStorageClass {
	return &NullableAzureArchivalStorageClass{value: val, isSet: true}
}

func (v NullableAzureArchivalStorageClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureArchivalStorageClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


