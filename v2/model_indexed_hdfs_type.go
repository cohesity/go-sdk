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

// checks if the IndexedHDFSType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndexedHDFSType{}

// IndexedHDFSType HDFS object Types for search.
type IndexedHDFSType struct {
	// Specifies the HDFS object Types for search.
	Type *string `json:"type,omitempty"`
}

// NewIndexedHDFSType instantiates a new IndexedHDFSType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexedHDFSType() *IndexedHDFSType {
	this := IndexedHDFSType{}
	return &this
}

// NewIndexedHDFSTypeWithDefaults instantiates a new IndexedHDFSType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexedHDFSTypeWithDefaults() *IndexedHDFSType {
	this := IndexedHDFSType{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IndexedHDFSType) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexedHDFSType) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IndexedHDFSType) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IndexedHDFSType) SetType(v string) {
	o.Type = &v
}

func (o IndexedHDFSType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndexedHDFSType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableIndexedHDFSType struct {
	value *IndexedHDFSType
	isSet bool
}

func (v NullableIndexedHDFSType) Get() *IndexedHDFSType {
	return v.value
}

func (v *NullableIndexedHDFSType) Set(val *IndexedHDFSType) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexedHDFSType) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexedHDFSType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexedHDFSType(val *IndexedHDFSType) *NullableIndexedHDFSType {
	return &NullableIndexedHDFSType{value: val, isSet: true}
}

func (v NullableIndexedHDFSType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexedHDFSType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


