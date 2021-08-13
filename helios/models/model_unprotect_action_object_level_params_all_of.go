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

// UnprotectActionObjectLevelParamsAllOf struct for UnprotectActionObjectLevelParamsAllOf
type UnprotectActionObjectLevelParamsAllOf struct {
	// Specifies whether to delete all snapshots along with unprotecting object protection. If set to true, all snapshots will be deleted and no more recoverable.
	DeleteAllSnapshots NullableBool `json:"deleteAllSnapshots,omitempty"`
	// If specified as true then it will cancel the ongoing runs and immediatly unprotect.
	ForceUnprotect NullableBool `json:"forceUnprotect,omitempty"`
}

// NewUnprotectActionObjectLevelParamsAllOf instantiates a new UnprotectActionObjectLevelParamsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnprotectActionObjectLevelParamsAllOf() *UnprotectActionObjectLevelParamsAllOf {
	this := UnprotectActionObjectLevelParamsAllOf{}
	return &this
}

// NewUnprotectActionObjectLevelParamsAllOfWithDefaults instantiates a new UnprotectActionObjectLevelParamsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnprotectActionObjectLevelParamsAllOfWithDefaults() *UnprotectActionObjectLevelParamsAllOf {
	this := UnprotectActionObjectLevelParamsAllOf{}
	return &this
}

// GetDeleteAllSnapshots returns the DeleteAllSnapshots field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnprotectActionObjectLevelParamsAllOf) GetDeleteAllSnapshots() bool {
	if o == nil || o.DeleteAllSnapshots.Get() == nil {
		var ret bool
		return ret
	}
	return *o.DeleteAllSnapshots.Get()
}

// GetDeleteAllSnapshotsOk returns a tuple with the DeleteAllSnapshots field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnprotectActionObjectLevelParamsAllOf) GetDeleteAllSnapshotsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeleteAllSnapshots.Get(), o.DeleteAllSnapshots.IsSet()
}

// HasDeleteAllSnapshots returns a boolean if a field has been set.
func (o *UnprotectActionObjectLevelParamsAllOf) HasDeleteAllSnapshots() bool {
	if o != nil && o.DeleteAllSnapshots.IsSet() {
		return true
	}

	return false
}

// SetDeleteAllSnapshots gets a reference to the given NullableBool and assigns it to the DeleteAllSnapshots field.
func (o *UnprotectActionObjectLevelParamsAllOf) SetDeleteAllSnapshots(v bool) {
	o.DeleteAllSnapshots.Set(&v)
}
// SetDeleteAllSnapshotsNil sets the value for DeleteAllSnapshots to be an explicit nil
func (o *UnprotectActionObjectLevelParamsAllOf) SetDeleteAllSnapshotsNil() {
	o.DeleteAllSnapshots.Set(nil)
}

// UnsetDeleteAllSnapshots ensures that no value is present for DeleteAllSnapshots, not even an explicit nil
func (o *UnprotectActionObjectLevelParamsAllOf) UnsetDeleteAllSnapshots() {
	o.DeleteAllSnapshots.Unset()
}

// GetForceUnprotect returns the ForceUnprotect field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnprotectActionObjectLevelParamsAllOf) GetForceUnprotect() bool {
	if o == nil || o.ForceUnprotect.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ForceUnprotect.Get()
}

// GetForceUnprotectOk returns a tuple with the ForceUnprotect field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnprotectActionObjectLevelParamsAllOf) GetForceUnprotectOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ForceUnprotect.Get(), o.ForceUnprotect.IsSet()
}

// HasForceUnprotect returns a boolean if a field has been set.
func (o *UnprotectActionObjectLevelParamsAllOf) HasForceUnprotect() bool {
	if o != nil && o.ForceUnprotect.IsSet() {
		return true
	}

	return false
}

// SetForceUnprotect gets a reference to the given NullableBool and assigns it to the ForceUnprotect field.
func (o *UnprotectActionObjectLevelParamsAllOf) SetForceUnprotect(v bool) {
	o.ForceUnprotect.Set(&v)
}
// SetForceUnprotectNil sets the value for ForceUnprotect to be an explicit nil
func (o *UnprotectActionObjectLevelParamsAllOf) SetForceUnprotectNil() {
	o.ForceUnprotect.Set(nil)
}

// UnsetForceUnprotect ensures that no value is present for ForceUnprotect, not even an explicit nil
func (o *UnprotectActionObjectLevelParamsAllOf) UnsetForceUnprotect() {
	o.ForceUnprotect.Unset()
}

func (o UnprotectActionObjectLevelParamsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeleteAllSnapshots.IsSet() {
		toSerialize["deleteAllSnapshots"] = o.DeleteAllSnapshots.Get()
	}
	if o.ForceUnprotect.IsSet() {
		toSerialize["forceUnprotect"] = o.ForceUnprotect.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableUnprotectActionObjectLevelParamsAllOf struct {
	value *UnprotectActionObjectLevelParamsAllOf
	isSet bool
}

func (v NullableUnprotectActionObjectLevelParamsAllOf) Get() *UnprotectActionObjectLevelParamsAllOf {
	return v.value
}

func (v *NullableUnprotectActionObjectLevelParamsAllOf) Set(val *UnprotectActionObjectLevelParamsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUnprotectActionObjectLevelParamsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUnprotectActionObjectLevelParamsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnprotectActionObjectLevelParamsAllOf(val *UnprotectActionObjectLevelParamsAllOf) *NullableUnprotectActionObjectLevelParamsAllOf {
	return &NullableUnprotectActionObjectLevelParamsAllOf{value: val, isSet: true}
}

func (v NullableUnprotectActionObjectLevelParamsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnprotectActionObjectLevelParamsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


