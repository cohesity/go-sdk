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

// checks if the GlacierFlrRestoreType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlacierFlrRestoreType{}

// GlacierFlrRestoreType Glacier Retrieval Type
type GlacierFlrRestoreType struct {
	// Specifies the glacier retrieval type.
	GlacierFlrRestoreType *string `json:"glacierFlrRestoreType,omitempty"`
}

// NewGlacierFlrRestoreType instantiates a new GlacierFlrRestoreType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlacierFlrRestoreType() *GlacierFlrRestoreType {
	this := GlacierFlrRestoreType{}
	return &this
}

// NewGlacierFlrRestoreTypeWithDefaults instantiates a new GlacierFlrRestoreType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlacierFlrRestoreTypeWithDefaults() *GlacierFlrRestoreType {
	this := GlacierFlrRestoreType{}
	return &this
}

// GetGlacierFlrRestoreType returns the GlacierFlrRestoreType field value if set, zero value otherwise.
func (o *GlacierFlrRestoreType) GetGlacierFlrRestoreType() string {
	if o == nil || IsNil(o.GlacierFlrRestoreType) {
		var ret string
		return ret
	}
	return *o.GlacierFlrRestoreType
}

// GetGlacierFlrRestoreTypeOk returns a tuple with the GlacierFlrRestoreType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlacierFlrRestoreType) GetGlacierFlrRestoreTypeOk() (*string, bool) {
	if o == nil || IsNil(o.GlacierFlrRestoreType) {
		return nil, false
	}
	return o.GlacierFlrRestoreType, true
}

// HasGlacierFlrRestoreType returns a boolean if a field has been set.
func (o *GlacierFlrRestoreType) HasGlacierFlrRestoreType() bool {
	if o != nil && !IsNil(o.GlacierFlrRestoreType) {
		return true
	}

	return false
}

// SetGlacierFlrRestoreType gets a reference to the given string and assigns it to the GlacierFlrRestoreType field.
func (o *GlacierFlrRestoreType) SetGlacierFlrRestoreType(v string) {
	o.GlacierFlrRestoreType = &v
}

func (o GlacierFlrRestoreType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GlacierFlrRestoreType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GlacierFlrRestoreType) {
		toSerialize["glacierFlrRestoreType"] = o.GlacierFlrRestoreType
	}
	return toSerialize, nil
}

type NullableGlacierFlrRestoreType struct {
	value *GlacierFlrRestoreType
	isSet bool
}

func (v NullableGlacierFlrRestoreType) Get() *GlacierFlrRestoreType {
	return v.value
}

func (v *NullableGlacierFlrRestoreType) Set(val *GlacierFlrRestoreType) {
	v.value = val
	v.isSet = true
}

func (v NullableGlacierFlrRestoreType) IsSet() bool {
	return v.isSet
}

func (v *NullableGlacierFlrRestoreType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlacierFlrRestoreType(val *GlacierFlrRestoreType) *NullableGlacierFlrRestoreType {
	return &NullableGlacierFlrRestoreType{value: val, isSet: true}
}

func (v NullableGlacierFlrRestoreType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlacierFlrRestoreType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


