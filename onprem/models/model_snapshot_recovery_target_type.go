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

// SnapshotRecoveryTargetType Snapshot Recovery Target Type
type SnapshotRecoveryTargetType struct {
	// Specifies the snapshot recovery target type.
	SnapshotRecoveryTargetType *string `json:"snapshotRecoveryTargetType,omitempty"`
}

// NewSnapshotRecoveryTargetType instantiates a new SnapshotRecoveryTargetType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotRecoveryTargetType() *SnapshotRecoveryTargetType {
	this := SnapshotRecoveryTargetType{}
	return &this
}

// NewSnapshotRecoveryTargetTypeWithDefaults instantiates a new SnapshotRecoveryTargetType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotRecoveryTargetTypeWithDefaults() *SnapshotRecoveryTargetType {
	this := SnapshotRecoveryTargetType{}
	return &this
}

// GetSnapshotRecoveryTargetType returns the SnapshotRecoveryTargetType field value if set, zero value otherwise.
func (o *SnapshotRecoveryTargetType) GetSnapshotRecoveryTargetType() string {
	if o == nil || o.SnapshotRecoveryTargetType == nil {
		var ret string
		return ret
	}
	return *o.SnapshotRecoveryTargetType
}

// GetSnapshotRecoveryTargetTypeOk returns a tuple with the SnapshotRecoveryTargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotRecoveryTargetType) GetSnapshotRecoveryTargetTypeOk() (*string, bool) {
	if o == nil || o.SnapshotRecoveryTargetType == nil {
		return nil, false
	}
	return o.SnapshotRecoveryTargetType, true
}

// HasSnapshotRecoveryTargetType returns a boolean if a field has been set.
func (o *SnapshotRecoveryTargetType) HasSnapshotRecoveryTargetType() bool {
	if o != nil && o.SnapshotRecoveryTargetType != nil {
		return true
	}

	return false
}

// SetSnapshotRecoveryTargetType gets a reference to the given string and assigns it to the SnapshotRecoveryTargetType field.
func (o *SnapshotRecoveryTargetType) SetSnapshotRecoveryTargetType(v string) {
	o.SnapshotRecoveryTargetType = &v
}

func (o SnapshotRecoveryTargetType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SnapshotRecoveryTargetType != nil {
		toSerialize["snapshotRecoveryTargetType"] = o.SnapshotRecoveryTargetType
	}
	return json.Marshal(toSerialize)
}

type NullableSnapshotRecoveryTargetType struct {
	value *SnapshotRecoveryTargetType
	isSet bool
}

func (v NullableSnapshotRecoveryTargetType) Get() *SnapshotRecoveryTargetType {
	return v.value
}

func (v *NullableSnapshotRecoveryTargetType) Set(val *SnapshotRecoveryTargetType) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotRecoveryTargetType) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotRecoveryTargetType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotRecoveryTargetType(val *SnapshotRecoveryTargetType) *NullableSnapshotRecoveryTargetType {
	return &NullableSnapshotRecoveryTargetType{value: val, isSet: true}
}

func (v NullableSnapshotRecoveryTargetType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotRecoveryTargetType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o SnapshotRecoveryTargetType) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}