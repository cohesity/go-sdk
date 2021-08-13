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

// ViewProtectionType Specifies the view protection type.
type ViewProtectionType struct {
	// Specifies the view protection type.
	Enum *string `json:"enum,omitempty"`
}

// NewViewProtectionType instantiates a new ViewProtectionType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewProtectionType() *ViewProtectionType {
	this := ViewProtectionType{}
	return &this
}

// NewViewProtectionTypeWithDefaults instantiates a new ViewProtectionType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewProtectionTypeWithDefaults() *ViewProtectionType {
	this := ViewProtectionType{}
	return &this
}

// GetEnum returns the Enum field value if set, zero value otherwise.
func (o *ViewProtectionType) GetEnum() string {
	if o == nil || o.Enum == nil {
		var ret string
		return ret
	}
	return *o.Enum
}

// GetEnumOk returns a tuple with the Enum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProtectionType) GetEnumOk() (*string, bool) {
	if o == nil || o.Enum == nil {
		return nil, false
	}
	return o.Enum, true
}

// HasEnum returns a boolean if a field has been set.
func (o *ViewProtectionType) HasEnum() bool {
	if o != nil && o.Enum != nil {
		return true
	}

	return false
}

// SetEnum gets a reference to the given string and assigns it to the Enum field.
func (o *ViewProtectionType) SetEnum(v string) {
	o.Enum = &v
}

func (o ViewProtectionType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enum != nil {
		toSerialize["enum"] = o.Enum
	}
	return json.Marshal(toSerialize)
}

type NullableViewProtectionType struct {
	value *ViewProtectionType
	isSet bool
}

func (v NullableViewProtectionType) Get() *ViewProtectionType {
	return v.value
}

func (v *NullableViewProtectionType) Set(val *ViewProtectionType) {
	v.value = val
	v.isSet = true
}

func (v NullableViewProtectionType) IsSet() bool {
	return v.isSet
}

func (v *NullableViewProtectionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewProtectionType(val *ViewProtectionType) *NullableViewProtectionType {
	return &NullableViewProtectionType{value: val, isSet: true}
}

func (v NullableViewProtectionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewProtectionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


