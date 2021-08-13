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

// UdaSourceType Enumeration of all the supported source types for the Universal Data Adapter.
type UdaSourceType struct {
	// Enumeration of all the supported source types for the Universal Data Adapter.
	Type *string `json:"type,omitempty"`
}

// NewUdaSourceType instantiates a new UdaSourceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUdaSourceType() *UdaSourceType {
	this := UdaSourceType{}
	return &this
}

// NewUdaSourceTypeWithDefaults instantiates a new UdaSourceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUdaSourceTypeWithDefaults() *UdaSourceType {
	this := UdaSourceType{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UdaSourceType) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdaSourceType) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UdaSourceType) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UdaSourceType) SetType(v string) {
	o.Type = &v
}

func (o UdaSourceType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableUdaSourceType struct {
	value *UdaSourceType
	isSet bool
}

func (v NullableUdaSourceType) Get() *UdaSourceType {
	return v.value
}

func (v *NullableUdaSourceType) Set(val *UdaSourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableUdaSourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableUdaSourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUdaSourceType(val *UdaSourceType) *NullableUdaSourceType {
	return &NullableUdaSourceType{value: val, isSet: true}
}

func (v NullableUdaSourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUdaSourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


