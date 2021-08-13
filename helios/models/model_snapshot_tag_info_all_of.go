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

// SnapshotTagInfoAllOf struct for SnapshotTagInfoAllOf
type SnapshotTagInfoAllOf struct {
	// Specifies runs the tags are applied to.
	RunIds []string `json:"runIds,omitempty"`
}

// NewSnapshotTagInfoAllOf instantiates a new SnapshotTagInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotTagInfoAllOf() *SnapshotTagInfoAllOf {
	this := SnapshotTagInfoAllOf{}
	return &this
}

// NewSnapshotTagInfoAllOfWithDefaults instantiates a new SnapshotTagInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotTagInfoAllOfWithDefaults() *SnapshotTagInfoAllOf {
	this := SnapshotTagInfoAllOf{}
	return &this
}

// GetRunIds returns the RunIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnapshotTagInfoAllOf) GetRunIds() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.RunIds
}

// GetRunIdsOk returns a tuple with the RunIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotTagInfoAllOf) GetRunIdsOk() (*[]string, bool) {
	if o == nil || o.RunIds == nil {
		return nil, false
	}
	return &o.RunIds, true
}

// HasRunIds returns a boolean if a field has been set.
func (o *SnapshotTagInfoAllOf) HasRunIds() bool {
	if o != nil && o.RunIds != nil {
		return true
	}

	return false
}

// SetRunIds gets a reference to the given []string and assigns it to the RunIds field.
func (o *SnapshotTagInfoAllOf) SetRunIds(v []string) {
	o.RunIds = v
}

func (o SnapshotTagInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RunIds != nil {
		toSerialize["runIds"] = o.RunIds
	}
	return json.Marshal(toSerialize)
}

type NullableSnapshotTagInfoAllOf struct {
	value *SnapshotTagInfoAllOf
	isSet bool
}

func (v NullableSnapshotTagInfoAllOf) Get() *SnapshotTagInfoAllOf {
	return v.value
}

func (v *NullableSnapshotTagInfoAllOf) Set(val *SnapshotTagInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotTagInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotTagInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotTagInfoAllOf(val *SnapshotTagInfoAllOf) *NullableSnapshotTagInfoAllOf {
	return &NullableSnapshotTagInfoAllOf{value: val, isSet: true}
}

func (v NullableSnapshotTagInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotTagInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


