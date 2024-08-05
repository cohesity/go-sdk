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

// checks if the IndexedGroupItemType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndexedGroupItemType{}

// IndexedGroupItemType Microsoft365 Group item types for search.
type IndexedGroupItemType struct {
	// Specifies the Group item types for search.
	Type *string `json:"type,omitempty"`
}

// NewIndexedGroupItemType instantiates a new IndexedGroupItemType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexedGroupItemType() *IndexedGroupItemType {
	this := IndexedGroupItemType{}
	return &this
}

// NewIndexedGroupItemTypeWithDefaults instantiates a new IndexedGroupItemType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexedGroupItemTypeWithDefaults() *IndexedGroupItemType {
	this := IndexedGroupItemType{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IndexedGroupItemType) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexedGroupItemType) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IndexedGroupItemType) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IndexedGroupItemType) SetType(v string) {
	o.Type = &v
}

func (o IndexedGroupItemType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndexedGroupItemType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableIndexedGroupItemType struct {
	value *IndexedGroupItemType
	isSet bool
}

func (v NullableIndexedGroupItemType) Get() *IndexedGroupItemType {
	return v.value
}

func (v *NullableIndexedGroupItemType) Set(val *IndexedGroupItemType) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexedGroupItemType) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexedGroupItemType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexedGroupItemType(val *IndexedGroupItemType) *NullableIndexedGroupItemType {
	return &NullableIndexedGroupItemType{value: val, isSet: true}
}

func (v NullableIndexedGroupItemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexedGroupItemType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


