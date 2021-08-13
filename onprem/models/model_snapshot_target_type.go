/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package onprem
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/onprem/utils"
	"fmt"
)

var _ = NullableBool{}

// SnapshotTargetType Snapshot Target Type
type SnapshotTargetType struct {
	// Specifies the snapshot target type.
	SnapshotTargetType *string `json:"snapshotTargetType,omitempty"`
}

// NewSnapshotTargetType instantiates a new SnapshotTargetType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotTargetType() *SnapshotTargetType {
	this := SnapshotTargetType{}
	return &this
}

// NewSnapshotTargetTypeWithDefaults instantiates a new SnapshotTargetType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotTargetTypeWithDefaults() *SnapshotTargetType {
	this := SnapshotTargetType{}
	return &this
}

// GetSnapshotTargetType returns the SnapshotTargetType field value if set, zero value otherwise.
func (o *SnapshotTargetType) GetSnapshotTargetType() string {
	if o == nil || o.SnapshotTargetType == nil {
		var ret string
		return ret
	}
	return *o.SnapshotTargetType
}

// GetSnapshotTargetTypeOk returns a tuple with the SnapshotTargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotTargetType) GetSnapshotTargetTypeOk() (*string, bool) {
	if o == nil || o.SnapshotTargetType == nil {
		return nil, false
	}
	return o.SnapshotTargetType, true
}

// HasSnapshotTargetType returns a boolean if a field has been set.
func (o *SnapshotTargetType) HasSnapshotTargetType() bool {
	if o != nil && o.SnapshotTargetType != nil {
		return true
	}

	return false
}

// SetSnapshotTargetType gets a reference to the given string and assigns it to the SnapshotTargetType field.
func (o *SnapshotTargetType) SetSnapshotTargetType(v string) {
	o.SnapshotTargetType = &v
}

func (o SnapshotTargetType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SnapshotTargetType != nil {
		toSerialize["snapshotTargetType"] = o.SnapshotTargetType
	}
	return json.Marshal(toSerialize)
}

type NullableSnapshotTargetType struct {
	value *SnapshotTargetType
	isSet bool
}

func (v NullableSnapshotTargetType) Get() *SnapshotTargetType {
	return v.value
}

func (v *NullableSnapshotTargetType) Set(val *SnapshotTargetType) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotTargetType) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotTargetType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotTargetType(val *SnapshotTargetType) *NullableSnapshotTargetType {
	return &NullableSnapshotTargetType{value: val, isSet: true}
}

func (v NullableSnapshotTargetType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotTargetType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o SnapshotTargetType) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}