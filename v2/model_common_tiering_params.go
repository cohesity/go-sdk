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

// checks if the CommonTieringParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonTieringParams{}

// CommonTieringParams Params common to Uptiering and Downtiering params
type CommonTieringParams struct {
	// If set, all files in the view will be uptiered regardless of file_select_policy, num_file_access, hot_file_window, file_size constraints.
	IncludeAllFiles NullableBool `json:"includeAllFiles,omitempty"`
	Target NullableDataTieringTarget `json:"target,omitempty"`
}

// NewCommonTieringParams instantiates a new CommonTieringParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonTieringParams() *CommonTieringParams {
	this := CommonTieringParams{}
	var includeAllFiles bool = false
	this.IncludeAllFiles = *NewNullableBool(&includeAllFiles)
	return &this
}

// NewCommonTieringParamsWithDefaults instantiates a new CommonTieringParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonTieringParamsWithDefaults() *CommonTieringParams {
	this := CommonTieringParams{}
	var includeAllFiles bool = false
	this.IncludeAllFiles = *NewNullableBool(&includeAllFiles)
	return &this
}

// GetIncludeAllFiles returns the IncludeAllFiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonTieringParams) GetIncludeAllFiles() bool {
	if o == nil || IsNil(o.IncludeAllFiles.Get()) {
		var ret bool
		return ret
	}
	return *o.IncludeAllFiles.Get()
}

// GetIncludeAllFilesOk returns a tuple with the IncludeAllFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonTieringParams) GetIncludeAllFilesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IncludeAllFiles.Get(), o.IncludeAllFiles.IsSet()
}

// HasIncludeAllFiles returns a boolean if a field has been set.
func (o *CommonTieringParams) HasIncludeAllFiles() bool {
	if o != nil && o.IncludeAllFiles.IsSet() {
		return true
	}

	return false
}

// SetIncludeAllFiles gets a reference to the given NullableBool and assigns it to the IncludeAllFiles field.
func (o *CommonTieringParams) SetIncludeAllFiles(v bool) {
	o.IncludeAllFiles.Set(&v)
}
// SetIncludeAllFilesNil sets the value for IncludeAllFiles to be an explicit nil
func (o *CommonTieringParams) SetIncludeAllFilesNil() {
	o.IncludeAllFiles.Set(nil)
}

// UnsetIncludeAllFiles ensures that no value is present for IncludeAllFiles, not even an explicit nil
func (o *CommonTieringParams) UnsetIncludeAllFiles() {
	o.IncludeAllFiles.Unset()
}

// GetTarget returns the Target field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonTieringParams) GetTarget() DataTieringTarget {
	if o == nil || IsNil(o.Target.Get()) {
		var ret DataTieringTarget
		return ret
	}
	return *o.Target.Get()
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonTieringParams) GetTargetOk() (*DataTieringTarget, bool) {
	if o == nil {
		return nil, false
	}
	return o.Target.Get(), o.Target.IsSet()
}

// HasTarget returns a boolean if a field has been set.
func (o *CommonTieringParams) HasTarget() bool {
	if o != nil && o.Target.IsSet() {
		return true
	}

	return false
}

// SetTarget gets a reference to the given NullableDataTieringTarget and assigns it to the Target field.
func (o *CommonTieringParams) SetTarget(v DataTieringTarget) {
	o.Target.Set(&v)
}
// SetTargetNil sets the value for Target to be an explicit nil
func (o *CommonTieringParams) SetTargetNil() {
	o.Target.Set(nil)
}

// UnsetTarget ensures that no value is present for Target, not even an explicit nil
func (o *CommonTieringParams) UnsetTarget() {
	o.Target.Unset()
}

func (o CommonTieringParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonTieringParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.IncludeAllFiles.IsSet() {
		toSerialize["includeAllFiles"] = o.IncludeAllFiles.Get()
	}
	if o.Target.IsSet() {
		toSerialize["target"] = o.Target.Get()
	}
	return toSerialize, nil
}

type NullableCommonTieringParams struct {
	value *CommonTieringParams
	isSet bool
}

func (v NullableCommonTieringParams) Get() *CommonTieringParams {
	return v.value
}

func (v *NullableCommonTieringParams) Set(val *CommonTieringParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonTieringParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonTieringParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonTieringParams(val *CommonTieringParams) *NullableCommonTieringParams {
	return &NullableCommonTieringParams{value: val, isSet: true}
}

func (v NullableCommonTieringParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonTieringParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


