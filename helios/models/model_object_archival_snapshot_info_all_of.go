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

// ObjectArchivalSnapshotInfoAllOf struct for ObjectArchivalSnapshotInfoAllOf
type ObjectArchivalSnapshotInfoAllOf struct {
	// Specifies the id of the archival snapshot for the object.
	SnapshotId NullableString `json:"snapshotId,omitempty"`
	// Specifies the logical size of this snapshot in bytes.
	LogicalSizeBytes NullableInt64 `json:"logicalSizeBytes,omitempty"`
}

// NewObjectArchivalSnapshotInfoAllOf instantiates a new ObjectArchivalSnapshotInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectArchivalSnapshotInfoAllOf() *ObjectArchivalSnapshotInfoAllOf {
	this := ObjectArchivalSnapshotInfoAllOf{}
	return &this
}

// NewObjectArchivalSnapshotInfoAllOfWithDefaults instantiates a new ObjectArchivalSnapshotInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectArchivalSnapshotInfoAllOfWithDefaults() *ObjectArchivalSnapshotInfoAllOf {
	this := ObjectArchivalSnapshotInfoAllOf{}
	return &this
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObjectArchivalSnapshotInfoAllOf) GetSnapshotId() string {
	if o == nil || o.SnapshotId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SnapshotId.Get()
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectArchivalSnapshotInfoAllOf) GetSnapshotIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SnapshotId.Get(), o.SnapshotId.IsSet()
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *ObjectArchivalSnapshotInfoAllOf) HasSnapshotId() bool {
	if o != nil && o.SnapshotId.IsSet() {
		return true
	}

	return false
}

// SetSnapshotId gets a reference to the given NullableString and assigns it to the SnapshotId field.
func (o *ObjectArchivalSnapshotInfoAllOf) SetSnapshotId(v string) {
	o.SnapshotId.Set(&v)
}
// SetSnapshotIdNil sets the value for SnapshotId to be an explicit nil
func (o *ObjectArchivalSnapshotInfoAllOf) SetSnapshotIdNil() {
	o.SnapshotId.Set(nil)
}

// UnsetSnapshotId ensures that no value is present for SnapshotId, not even an explicit nil
func (o *ObjectArchivalSnapshotInfoAllOf) UnsetSnapshotId() {
	o.SnapshotId.Unset()
}

// GetLogicalSizeBytes returns the LogicalSizeBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObjectArchivalSnapshotInfoAllOf) GetLogicalSizeBytes() int64 {
	if o == nil || o.LogicalSizeBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LogicalSizeBytes.Get()
}

// GetLogicalSizeBytesOk returns a tuple with the LogicalSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectArchivalSnapshotInfoAllOf) GetLogicalSizeBytesOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LogicalSizeBytes.Get(), o.LogicalSizeBytes.IsSet()
}

// HasLogicalSizeBytes returns a boolean if a field has been set.
func (o *ObjectArchivalSnapshotInfoAllOf) HasLogicalSizeBytes() bool {
	if o != nil && o.LogicalSizeBytes.IsSet() {
		return true
	}

	return false
}

// SetLogicalSizeBytes gets a reference to the given NullableInt64 and assigns it to the LogicalSizeBytes field.
func (o *ObjectArchivalSnapshotInfoAllOf) SetLogicalSizeBytes(v int64) {
	o.LogicalSizeBytes.Set(&v)
}
// SetLogicalSizeBytesNil sets the value for LogicalSizeBytes to be an explicit nil
func (o *ObjectArchivalSnapshotInfoAllOf) SetLogicalSizeBytesNil() {
	o.LogicalSizeBytes.Set(nil)
}

// UnsetLogicalSizeBytes ensures that no value is present for LogicalSizeBytes, not even an explicit nil
func (o *ObjectArchivalSnapshotInfoAllOf) UnsetLogicalSizeBytes() {
	o.LogicalSizeBytes.Unset()
}

func (o ObjectArchivalSnapshotInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SnapshotId.IsSet() {
		toSerialize["snapshotId"] = o.SnapshotId.Get()
	}
	if o.LogicalSizeBytes.IsSet() {
		toSerialize["logicalSizeBytes"] = o.LogicalSizeBytes.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableObjectArchivalSnapshotInfoAllOf struct {
	value *ObjectArchivalSnapshotInfoAllOf
	isSet bool
}

func (v NullableObjectArchivalSnapshotInfoAllOf) Get() *ObjectArchivalSnapshotInfoAllOf {
	return v.value
}

func (v *NullableObjectArchivalSnapshotInfoAllOf) Set(val *ObjectArchivalSnapshotInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectArchivalSnapshotInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectArchivalSnapshotInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectArchivalSnapshotInfoAllOf(val *ObjectArchivalSnapshotInfoAllOf) *NullableObjectArchivalSnapshotInfoAllOf {
	return &NullableObjectArchivalSnapshotInfoAllOf{value: val, isSet: true}
}

func (v NullableObjectArchivalSnapshotInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectArchivalSnapshotInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


