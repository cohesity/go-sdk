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

// UptieringPolicyAllOf struct for UptieringPolicyAllOf
type UptieringPolicyAllOf struct {
	FileAge *UptieringFileAgePolicy `json:"fileAge,omitempty"`
	// If set, all files in the view will be uptiered regardless of file_select_policy, num_file_access, hot_file_window, file_size constraints.
	IncludeAllFiles NullableBool `json:"includeAllFiles,omitempty"`
	Target NullableUptieringTarget `json:"target,omitempty"`
}

// NewUptieringPolicyAllOf instantiates a new UptieringPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUptieringPolicyAllOf() *UptieringPolicyAllOf {
	this := UptieringPolicyAllOf{}
	var includeAllFiles bool = false
	this.IncludeAllFiles = *NewNullableBool(&includeAllFiles)
	return &this
}

// NewUptieringPolicyAllOfWithDefaults instantiates a new UptieringPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUptieringPolicyAllOfWithDefaults() *UptieringPolicyAllOf {
	this := UptieringPolicyAllOf{}
	var includeAllFiles bool = false
	this.IncludeAllFiles = *NewNullableBool(&includeAllFiles) // model_simple false
	return &this
}

// GetFileAge returns the FileAge field value if set, zero value otherwise.
func (o *UptieringPolicyAllOf) GetFileAge() UptieringFileAgePolicy {
	if o == nil || o.FileAge == nil {
		var ret UptieringFileAgePolicy
		return ret
	}
	return *o.FileAge
}

// GetFileAgeOk returns a tuple with the FileAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UptieringPolicyAllOf) GetFileAgeOk() (*UptieringFileAgePolicy, bool) {
	if o == nil || o.FileAge == nil {
		return nil, false
	}
	return o.FileAge, true
}

// HasFileAge returns a boolean if a field has been set.
func (o *UptieringPolicyAllOf) HasFileAge() bool {
	if o != nil && o.FileAge != nil {
		return true
	}

	return false
}

// SetFileAge gets a reference to the given UptieringFileAgePolicy and assigns it to the FileAge field.
func (o *UptieringPolicyAllOf) SetFileAge(v UptieringFileAgePolicy) {
	o.FileAge = &v
}

// GetIncludeAllFiles returns the IncludeAllFiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UptieringPolicyAllOf) GetIncludeAllFiles() bool {
	if o == nil || o.IncludeAllFiles.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IncludeAllFiles.Get()
}

// GetIncludeAllFilesOk returns a tuple with the IncludeAllFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UptieringPolicyAllOf) GetIncludeAllFilesOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IncludeAllFiles.Get(), o.IncludeAllFiles.IsSet()
}

// HasIncludeAllFiles returns a boolean if a field has been set.
func (o *UptieringPolicyAllOf) HasIncludeAllFiles() bool {
	if o != nil && o.IncludeAllFiles.IsSet() {
		return true
	}

	return false
}

// SetIncludeAllFiles gets a reference to the given NullableBool and assigns it to the IncludeAllFiles field.
func (o *UptieringPolicyAllOf) SetIncludeAllFiles(v bool) {
	o.IncludeAllFiles.Set(&v)
}
// SetIncludeAllFilesNil sets the value for IncludeAllFiles to be an explicit nil
func (o *UptieringPolicyAllOf) SetIncludeAllFilesNil() {
	o.IncludeAllFiles.Set(nil)
}

// UnsetIncludeAllFiles ensures that no value is present for IncludeAllFiles, not even an explicit nil
func (o *UptieringPolicyAllOf) UnsetIncludeAllFiles() {
	o.IncludeAllFiles.Unset()
}

// GetTarget returns the Target field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UptieringPolicyAllOf) GetTarget() UptieringTarget {
	if o == nil || o.Target.Get() == nil {
		var ret UptieringTarget
		return ret
	}
	return *o.Target.Get()
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UptieringPolicyAllOf) GetTargetOk() (*UptieringTarget, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Target.Get(), o.Target.IsSet()
}

// HasTarget returns a boolean if a field has been set.
func (o *UptieringPolicyAllOf) HasTarget() bool {
	if o != nil && o.Target.IsSet() {
		return true
	}

	return false
}

// SetTarget gets a reference to the given NullableUptieringTarget and assigns it to the Target field.
func (o *UptieringPolicyAllOf) SetTarget(v UptieringTarget) {
	o.Target.Set(&v)
}
// SetTargetNil sets the value for Target to be an explicit nil
func (o *UptieringPolicyAllOf) SetTargetNil() {
	o.Target.Set(nil)
}

// UnsetTarget ensures that no value is present for Target, not even an explicit nil
func (o *UptieringPolicyAllOf) UnsetTarget() {
	o.Target.Unset()
}

func (o UptieringPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FileAge != nil {
		toSerialize["fileAge"] = o.FileAge
	}
	if o.IncludeAllFiles.IsSet() {
		toSerialize["includeAllFiles"] = o.IncludeAllFiles.Get()
	}
	if o.Target.IsSet() {
		toSerialize["target"] = o.Target.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableUptieringPolicyAllOf struct {
	value *UptieringPolicyAllOf
	isSet bool
}

func (v NullableUptieringPolicyAllOf) Get() *UptieringPolicyAllOf {
	return v.value
}

func (v *NullableUptieringPolicyAllOf) Set(val *UptieringPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUptieringPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUptieringPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUptieringPolicyAllOf(val *UptieringPolicyAllOf) *NullableUptieringPolicyAllOf {
	return &NullableUptieringPolicyAllOf{value: val, isSet: true}
}

func (v NullableUptieringPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUptieringPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


