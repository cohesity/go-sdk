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

// AwsNativeObjectProtectionParams Specifies the parameters which are specific to AWS Object Protection Groups using AWS native snapshot APIs. Atlease one of tags or objects must be specified.
type AwsNativeObjectProtectionParams struct {
	// Specifies the objects to be protected.
	Objects *[]AwsObjectLevelParams `json:"objects,omitempty"`
	VolumeExclusionParams NullableEbsVolumeExclusionParams `json:"volumeExclusionParams,omitempty"`
}

// NewAwsNativeObjectProtectionParams instantiates a new AwsNativeObjectProtectionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsNativeObjectProtectionParams() *AwsNativeObjectProtectionParams {
	this := AwsNativeObjectProtectionParams{}
	return &this
}

// NewAwsNativeObjectProtectionParamsWithDefaults instantiates a new AwsNativeObjectProtectionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsNativeObjectProtectionParamsWithDefaults() *AwsNativeObjectProtectionParams {
	this := AwsNativeObjectProtectionParams{}
	return &this
}

// GetObjects returns the Objects field value if set, zero value otherwise.
func (o *AwsNativeObjectProtectionParams) GetObjects() []AwsObjectLevelParams {
	if o == nil || o.Objects == nil {
		var ret []AwsObjectLevelParams
		return ret
	}
	return *o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsNativeObjectProtectionParams) GetObjectsOk() (*[]AwsObjectLevelParams, bool) {
	if o == nil || o.Objects == nil {
		return nil, false
	}
	return o.Objects, true
}

// HasObjects returns a boolean if a field has been set.
func (o *AwsNativeObjectProtectionParams) HasObjects() bool {
	if o != nil && o.Objects != nil {
		return true
	}

	return false
}

// SetObjects gets a reference to the given []AwsObjectLevelParams and assigns it to the Objects field.
func (o *AwsNativeObjectProtectionParams) SetObjects(v []AwsObjectLevelParams) {
	o.Objects = &v
}

// GetVolumeExclusionParams returns the VolumeExclusionParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsNativeObjectProtectionParams) GetVolumeExclusionParams() EbsVolumeExclusionParams {
	if o == nil || o.VolumeExclusionParams.Get() == nil {
		var ret EbsVolumeExclusionParams
		return ret
	}
	return *o.VolumeExclusionParams.Get()
}

// GetVolumeExclusionParamsOk returns a tuple with the VolumeExclusionParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsNativeObjectProtectionParams) GetVolumeExclusionParamsOk() (*EbsVolumeExclusionParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VolumeExclusionParams.Get(), o.VolumeExclusionParams.IsSet()
}

// HasVolumeExclusionParams returns a boolean if a field has been set.
func (o *AwsNativeObjectProtectionParams) HasVolumeExclusionParams() bool {
	if o != nil && o.VolumeExclusionParams.IsSet() {
		return true
	}

	return false
}

// SetVolumeExclusionParams gets a reference to the given NullableEbsVolumeExclusionParams and assigns it to the VolumeExclusionParams field.
func (o *AwsNativeObjectProtectionParams) SetVolumeExclusionParams(v EbsVolumeExclusionParams) {
	o.VolumeExclusionParams.Set(&v)
}
// SetVolumeExclusionParamsNil sets the value for VolumeExclusionParams to be an explicit nil
func (o *AwsNativeObjectProtectionParams) SetVolumeExclusionParamsNil() {
	o.VolumeExclusionParams.Set(nil)
}

// UnsetVolumeExclusionParams ensures that no value is present for VolumeExclusionParams, not even an explicit nil
func (o *AwsNativeObjectProtectionParams) UnsetVolumeExclusionParams() {
	o.VolumeExclusionParams.Unset()
}

func (o AwsNativeObjectProtectionParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Objects != nil {
		toSerialize["objects"] = o.Objects
	}
	if o.VolumeExclusionParams.IsSet() {
		toSerialize["volumeExclusionParams"] = o.VolumeExclusionParams.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAwsNativeObjectProtectionParams struct {
	value *AwsNativeObjectProtectionParams
	isSet bool
}

func (v NullableAwsNativeObjectProtectionParams) Get() *AwsNativeObjectProtectionParams {
	return v.value
}

func (v *NullableAwsNativeObjectProtectionParams) Set(val *AwsNativeObjectProtectionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsNativeObjectProtectionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsNativeObjectProtectionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsNativeObjectProtectionParams(val *AwsNativeObjectProtectionParams) *NullableAwsNativeObjectProtectionParams {
	return &NullableAwsNativeObjectProtectionParams{value: val, isSet: true}
}

func (v NullableAwsNativeObjectProtectionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsNativeObjectProtectionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


