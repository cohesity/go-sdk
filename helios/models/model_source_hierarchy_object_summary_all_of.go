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

// SourceHierarchyObjectSummaryAllOf struct for SourceHierarchyObjectSummaryAllOf
type SourceHierarchyObjectSummaryAllOf struct {
	// Specifies the ID of the direct parent of this object in the source hierarchy.
	ParentId NullableInt64 `json:"parentId,omitempty"`
}

// NewSourceHierarchyObjectSummaryAllOf instantiates a new SourceHierarchyObjectSummaryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceHierarchyObjectSummaryAllOf() *SourceHierarchyObjectSummaryAllOf {
	this := SourceHierarchyObjectSummaryAllOf{}
	return &this
}

// NewSourceHierarchyObjectSummaryAllOfWithDefaults instantiates a new SourceHierarchyObjectSummaryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceHierarchyObjectSummaryAllOfWithDefaults() *SourceHierarchyObjectSummaryAllOf {
	this := SourceHierarchyObjectSummaryAllOf{}
	return &this
}

// GetParentId returns the ParentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceHierarchyObjectSummaryAllOf) GetParentId() int64 {
	if o == nil || o.ParentId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceHierarchyObjectSummaryAllOf) GetParentIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// HasParentId returns a boolean if a field has been set.
func (o *SourceHierarchyObjectSummaryAllOf) HasParentId() bool {
	if o != nil && o.ParentId.IsSet() {
		return true
	}

	return false
}

// SetParentId gets a reference to the given NullableInt64 and assigns it to the ParentId field.
func (o *SourceHierarchyObjectSummaryAllOf) SetParentId(v int64) {
	o.ParentId.Set(&v)
}
// SetParentIdNil sets the value for ParentId to be an explicit nil
func (o *SourceHierarchyObjectSummaryAllOf) SetParentIdNil() {
	o.ParentId.Set(nil)
}

// UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
func (o *SourceHierarchyObjectSummaryAllOf) UnsetParentId() {
	o.ParentId.Unset()
}

func (o SourceHierarchyObjectSummaryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ParentId.IsSet() {
		toSerialize["parentId"] = o.ParentId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSourceHierarchyObjectSummaryAllOf struct {
	value *SourceHierarchyObjectSummaryAllOf
	isSet bool
}

func (v NullableSourceHierarchyObjectSummaryAllOf) Get() *SourceHierarchyObjectSummaryAllOf {
	return v.value
}

func (v *NullableSourceHierarchyObjectSummaryAllOf) Set(val *SourceHierarchyObjectSummaryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceHierarchyObjectSummaryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceHierarchyObjectSummaryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceHierarchyObjectSummaryAllOf(val *SourceHierarchyObjectSummaryAllOf) *NullableSourceHierarchyObjectSummaryAllOf {
	return &NullableSourceHierarchyObjectSummaryAllOf{value: val, isSet: true}
}

func (v NullableSourceHierarchyObjectSummaryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceHierarchyObjectSummaryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


