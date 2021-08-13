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

// AuditEntityTypes Entity Types
type AuditEntityTypes struct {
	// Entity Types
	EntityTypes *string `json:"entityTypes,omitempty"`
}

// NewAuditEntityTypes instantiates a new AuditEntityTypes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditEntityTypes() *AuditEntityTypes {
	this := AuditEntityTypes{}
	return &this
}

// NewAuditEntityTypesWithDefaults instantiates a new AuditEntityTypes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditEntityTypesWithDefaults() *AuditEntityTypes {
	this := AuditEntityTypes{}
	return &this
}

// GetEntityTypes returns the EntityTypes field value if set, zero value otherwise.
func (o *AuditEntityTypes) GetEntityTypes() string {
	if o == nil || o.EntityTypes == nil {
		var ret string
		return ret
	}
	return *o.EntityTypes
}

// GetEntityTypesOk returns a tuple with the EntityTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditEntityTypes) GetEntityTypesOk() (*string, bool) {
	if o == nil || o.EntityTypes == nil {
		return nil, false
	}
	return o.EntityTypes, true
}

// HasEntityTypes returns a boolean if a field has been set.
func (o *AuditEntityTypes) HasEntityTypes() bool {
	if o != nil && o.EntityTypes != nil {
		return true
	}

	return false
}

// SetEntityTypes gets a reference to the given string and assigns it to the EntityTypes field.
func (o *AuditEntityTypes) SetEntityTypes(v string) {
	o.EntityTypes = &v
}

func (o AuditEntityTypes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EntityTypes != nil {
		toSerialize["entityTypes"] = o.EntityTypes
	}
	return json.Marshal(toSerialize)
}

type NullableAuditEntityTypes struct {
	value *AuditEntityTypes
	isSet bool
}

func (v NullableAuditEntityTypes) Get() *AuditEntityTypes {
	return v.value
}

func (v *NullableAuditEntityTypes) Set(val *AuditEntityTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditEntityTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditEntityTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditEntityTypes(val *AuditEntityTypes) *NullableAuditEntityTypes {
	return &NullableAuditEntityTypes{value: val, isSet: true}
}

func (v NullableAuditEntityTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditEntityTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


