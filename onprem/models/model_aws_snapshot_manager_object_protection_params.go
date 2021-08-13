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

// AwsSnapshotManagerObjectProtectionParams Specifies the parameters which are specific to AWS Object Protection using AWS native snapshot orchestration with snapshot manager. Atlease one of tags or objects must be specified.
type AwsSnapshotManagerObjectProtectionParams struct {
	// Specifies the objects to be protected.
	Objects *[]AwsObjectLevelParams `json:"objects,omitempty"`
	VolumeExclusionParams NullableEbsVolumeExclusionParams `json:"volumeExclusionParams,omitempty"`
	// Specifies whether AMI should be created after taking snapshots of the instance.
	CreateAmi NullableBool `json:"createAmi,omitempty"`
	// The frequency of AMI creation. This should be set if the option to create AMI is set. A value of n creates an AMI from the snapshots after every n runs. eg. n = 2 implies every alternate backup run starting from the first will create an AMI.
	AmiCreationFrequency NullableInt32 `json:"amiCreationFrequency,omitempty"`
}

// NewAwsSnapshotManagerObjectProtectionParams instantiates a new AwsSnapshotManagerObjectProtectionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsSnapshotManagerObjectProtectionParams() *AwsSnapshotManagerObjectProtectionParams {
	this := AwsSnapshotManagerObjectProtectionParams{}
	return &this
}

// NewAwsSnapshotManagerObjectProtectionParamsWithDefaults instantiates a new AwsSnapshotManagerObjectProtectionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsSnapshotManagerObjectProtectionParamsWithDefaults() *AwsSnapshotManagerObjectProtectionParams {
	this := AwsSnapshotManagerObjectProtectionParams{}
	return &this
}

// GetObjects returns the Objects field value if set, zero value otherwise.
func (o *AwsSnapshotManagerObjectProtectionParams) GetObjects() []AwsObjectLevelParams {
	if o == nil || o.Objects == nil {
		var ret []AwsObjectLevelParams
		return ret
	}
	return *o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsSnapshotManagerObjectProtectionParams) GetObjectsOk() (*[]AwsObjectLevelParams, bool) {
	if o == nil || o.Objects == nil {
		return nil, false
	}
	return o.Objects, true
}

// HasObjects returns a boolean if a field has been set.
func (o *AwsSnapshotManagerObjectProtectionParams) HasObjects() bool {
	if o != nil && o.Objects != nil {
		return true
	}

	return false
}

// SetObjects gets a reference to the given []AwsObjectLevelParams and assigns it to the Objects field.
func (o *AwsSnapshotManagerObjectProtectionParams) SetObjects(v []AwsObjectLevelParams) {
	o.Objects = &v
}

// GetVolumeExclusionParams returns the VolumeExclusionParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsSnapshotManagerObjectProtectionParams) GetVolumeExclusionParams() EbsVolumeExclusionParams {
	if o == nil || o.VolumeExclusionParams.Get() == nil {
		var ret EbsVolumeExclusionParams
		return ret
	}
	return *o.VolumeExclusionParams.Get()
}

// GetVolumeExclusionParamsOk returns a tuple with the VolumeExclusionParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsSnapshotManagerObjectProtectionParams) GetVolumeExclusionParamsOk() (*EbsVolumeExclusionParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VolumeExclusionParams.Get(), o.VolumeExclusionParams.IsSet()
}

// HasVolumeExclusionParams returns a boolean if a field has been set.
func (o *AwsSnapshotManagerObjectProtectionParams) HasVolumeExclusionParams() bool {
	if o != nil && o.VolumeExclusionParams.IsSet() {
		return true
	}

	return false
}

// SetVolumeExclusionParams gets a reference to the given NullableEbsVolumeExclusionParams and assigns it to the VolumeExclusionParams field.
func (o *AwsSnapshotManagerObjectProtectionParams) SetVolumeExclusionParams(v EbsVolumeExclusionParams) {
	o.VolumeExclusionParams.Set(&v)
}
// SetVolumeExclusionParamsNil sets the value for VolumeExclusionParams to be an explicit nil
func (o *AwsSnapshotManagerObjectProtectionParams) SetVolumeExclusionParamsNil() {
	o.VolumeExclusionParams.Set(nil)
}

// UnsetVolumeExclusionParams ensures that no value is present for VolumeExclusionParams, not even an explicit nil
func (o *AwsSnapshotManagerObjectProtectionParams) UnsetVolumeExclusionParams() {
	o.VolumeExclusionParams.Unset()
}

// GetCreateAmi returns the CreateAmi field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsSnapshotManagerObjectProtectionParams) GetCreateAmi() bool {
	if o == nil || o.CreateAmi.Get() == nil {
		var ret bool
		return ret
	}
	return *o.CreateAmi.Get()
}

// GetCreateAmiOk returns a tuple with the CreateAmi field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsSnapshotManagerObjectProtectionParams) GetCreateAmiOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreateAmi.Get(), o.CreateAmi.IsSet()
}

// HasCreateAmi returns a boolean if a field has been set.
func (o *AwsSnapshotManagerObjectProtectionParams) HasCreateAmi() bool {
	if o != nil && o.CreateAmi.IsSet() {
		return true
	}

	return false
}

// SetCreateAmi gets a reference to the given NullableBool and assigns it to the CreateAmi field.
func (o *AwsSnapshotManagerObjectProtectionParams) SetCreateAmi(v bool) {
	o.CreateAmi.Set(&v)
}
// SetCreateAmiNil sets the value for CreateAmi to be an explicit nil
func (o *AwsSnapshotManagerObjectProtectionParams) SetCreateAmiNil() {
	o.CreateAmi.Set(nil)
}

// UnsetCreateAmi ensures that no value is present for CreateAmi, not even an explicit nil
func (o *AwsSnapshotManagerObjectProtectionParams) UnsetCreateAmi() {
	o.CreateAmi.Unset()
}

// GetAmiCreationFrequency returns the AmiCreationFrequency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsSnapshotManagerObjectProtectionParams) GetAmiCreationFrequency() int32 {
	if o == nil || o.AmiCreationFrequency.Get() == nil {
		var ret int32
		return ret
	}
	return *o.AmiCreationFrequency.Get()
}

// GetAmiCreationFrequencyOk returns a tuple with the AmiCreationFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsSnapshotManagerObjectProtectionParams) GetAmiCreationFrequencyOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AmiCreationFrequency.Get(), o.AmiCreationFrequency.IsSet()
}

// HasAmiCreationFrequency returns a boolean if a field has been set.
func (o *AwsSnapshotManagerObjectProtectionParams) HasAmiCreationFrequency() bool {
	if o != nil && o.AmiCreationFrequency.IsSet() {
		return true
	}

	return false
}

// SetAmiCreationFrequency gets a reference to the given NullableInt32 and assigns it to the AmiCreationFrequency field.
func (o *AwsSnapshotManagerObjectProtectionParams) SetAmiCreationFrequency(v int32) {
	o.AmiCreationFrequency.Set(&v)
}
// SetAmiCreationFrequencyNil sets the value for AmiCreationFrequency to be an explicit nil
func (o *AwsSnapshotManagerObjectProtectionParams) SetAmiCreationFrequencyNil() {
	o.AmiCreationFrequency.Set(nil)
}

// UnsetAmiCreationFrequency ensures that no value is present for AmiCreationFrequency, not even an explicit nil
func (o *AwsSnapshotManagerObjectProtectionParams) UnsetAmiCreationFrequency() {
	o.AmiCreationFrequency.Unset()
}

func (o AwsSnapshotManagerObjectProtectionParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Objects != nil {
		toSerialize["objects"] = o.Objects
	}
	if o.VolumeExclusionParams.IsSet() {
		toSerialize["volumeExclusionParams"] = o.VolumeExclusionParams.Get()
	}
	if o.CreateAmi.IsSet() {
		toSerialize["createAmi"] = o.CreateAmi.Get()
	}
	if o.AmiCreationFrequency.IsSet() {
		toSerialize["amiCreationFrequency"] = o.AmiCreationFrequency.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAwsSnapshotManagerObjectProtectionParams struct {
	value *AwsSnapshotManagerObjectProtectionParams
	isSet bool
}

func (v NullableAwsSnapshotManagerObjectProtectionParams) Get() *AwsSnapshotManagerObjectProtectionParams {
	return v.value
}

func (v *NullableAwsSnapshotManagerObjectProtectionParams) Set(val *AwsSnapshotManagerObjectProtectionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsSnapshotManagerObjectProtectionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsSnapshotManagerObjectProtectionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsSnapshotManagerObjectProtectionParams(val *AwsSnapshotManagerObjectProtectionParams) *NullableAwsSnapshotManagerObjectProtectionParams {
	return &NullableAwsSnapshotManagerObjectProtectionParams{value: val, isSet: true}
}

func (v NullableAwsSnapshotManagerObjectProtectionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsSnapshotManagerObjectProtectionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o AwsSnapshotManagerObjectProtectionParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}