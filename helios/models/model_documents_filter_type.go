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

// DocumentsFilterType Enumeration of all the document filter option.
type DocumentsFilterType struct {
	// Enumeration of all the document filter option.
	DocumentsFilterType *string `json:"DocumentsFilterType,omitempty"`
}

// NewDocumentsFilterType instantiates a new DocumentsFilterType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentsFilterType() *DocumentsFilterType {
	this := DocumentsFilterType{}
	return &this
}

// NewDocumentsFilterTypeWithDefaults instantiates a new DocumentsFilterType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentsFilterTypeWithDefaults() *DocumentsFilterType {
	this := DocumentsFilterType{}
	return &this
}

// GetDocumentsFilterType returns the DocumentsFilterType field value if set, zero value otherwise.
func (o *DocumentsFilterType) GetDocumentsFilterType() string {
	if o == nil || o.DocumentsFilterType == nil {
		var ret string
		return ret
	}
	return *o.DocumentsFilterType
}

// GetDocumentsFilterTypeOk returns a tuple with the DocumentsFilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentsFilterType) GetDocumentsFilterTypeOk() (*string, bool) {
	if o == nil || o.DocumentsFilterType == nil {
		return nil, false
	}
	return o.DocumentsFilterType, true
}

// HasDocumentsFilterType returns a boolean if a field has been set.
func (o *DocumentsFilterType) HasDocumentsFilterType() bool {
	if o != nil && o.DocumentsFilterType != nil {
		return true
	}

	return false
}

// SetDocumentsFilterType gets a reference to the given string and assigns it to the DocumentsFilterType field.
func (o *DocumentsFilterType) SetDocumentsFilterType(v string) {
	o.DocumentsFilterType = &v
}

func (o DocumentsFilterType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DocumentsFilterType != nil {
		toSerialize["DocumentsFilterType"] = o.DocumentsFilterType
	}
	return json.Marshal(toSerialize)
}

type NullableDocumentsFilterType struct {
	value *DocumentsFilterType
	isSet bool
}

func (v NullableDocumentsFilterType) Get() *DocumentsFilterType {
	return v.value
}

func (v *NullableDocumentsFilterType) Set(val *DocumentsFilterType) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentsFilterType) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentsFilterType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentsFilterType(val *DocumentsFilterType) *NullableDocumentsFilterType {
	return &NullableDocumentsFilterType{value: val, isSet: true}
}

func (v NullableDocumentsFilterType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentsFilterType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


