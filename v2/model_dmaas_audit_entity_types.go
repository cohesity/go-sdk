/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

// checks if the DmaasAuditEntityTypes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DmaasAuditEntityTypes{}

// DmaasAuditEntityTypes Dmaas Audit Entity Types
type DmaasAuditEntityTypes struct {
	// Dmaas Audit Entity Types
	EntityTypes *string `json:"entityTypes,omitempty"`
}

// NewDmaasAuditEntityTypes instantiates a new DmaasAuditEntityTypes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDmaasAuditEntityTypes() *DmaasAuditEntityTypes {
	this := DmaasAuditEntityTypes{}
	return &this
}

// NewDmaasAuditEntityTypesWithDefaults instantiates a new DmaasAuditEntityTypes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDmaasAuditEntityTypesWithDefaults() *DmaasAuditEntityTypes {
	this := DmaasAuditEntityTypes{}
	return &this
}

// GetEntityTypes returns the EntityTypes field value if set, zero value otherwise.
func (o *DmaasAuditEntityTypes) GetEntityTypes() string {
	if o == nil || IsNil(o.EntityTypes) {
		var ret string
		return ret
	}
	return *o.EntityTypes
}

// GetEntityTypesOk returns a tuple with the EntityTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmaasAuditEntityTypes) GetEntityTypesOk() (*string, bool) {
	if o == nil || IsNil(o.EntityTypes) {
		return nil, false
	}
	return o.EntityTypes, true
}

// HasEntityTypes returns a boolean if a field has been set.
func (o *DmaasAuditEntityTypes) HasEntityTypes() bool {
	if o != nil && !IsNil(o.EntityTypes) {
		return true
	}

	return false
}

// SetEntityTypes gets a reference to the given string and assigns it to the EntityTypes field.
func (o *DmaasAuditEntityTypes) SetEntityTypes(v string) {
	o.EntityTypes = &v
}

func (o DmaasAuditEntityTypes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DmaasAuditEntityTypes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EntityTypes) {
		toSerialize["entityTypes"] = o.EntityTypes
	}
	return toSerialize, nil
}

type NullableDmaasAuditEntityTypes struct {
	value *DmaasAuditEntityTypes
	isSet bool
}

func (v NullableDmaasAuditEntityTypes) Get() *DmaasAuditEntityTypes {
	return v.value
}

func (v *NullableDmaasAuditEntityTypes) Set(val *DmaasAuditEntityTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableDmaasAuditEntityTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableDmaasAuditEntityTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDmaasAuditEntityTypes(val *DmaasAuditEntityTypes) *NullableDmaasAuditEntityTypes {
	return &NullableDmaasAuditEntityTypes{value: val, isSet: true}
}

func (v NullableDmaasAuditEntityTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDmaasAuditEntityTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


