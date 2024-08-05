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

// checks if the AwsSnapshotManagerParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsSnapshotManagerParams{}

// AwsSnapshotManagerParams Specifies job parameters applicable for all 'kVMware' Environment type Protection Sources in a Protection Job.
type AwsSnapshotManagerParams struct {
	// Specifies the frequency of AMI creation. This should be set if the option to create AMI is set. A value of n creates an AMI from the snapshots after every n runs. eg. n = 2 implies every alternate backup run starting from the first will create an AMI.
	AmiCreationFrequency NullableInt32 `json:"amiCreationFrequency,omitempty"`
	// If true, creates an AMI after taking snapshots of the instance. It should be set only while backing up EC2 instances. CreateAmi creates AMI for the protection job.
	CreateAmi NullableBool `json:"createAmi,omitempty"`
	VolumeExclusionParams NullableEbsVolumeExclusionParams `json:"volumeExclusionParams,omitempty"`
}

// NewAwsSnapshotManagerParams instantiates a new AwsSnapshotManagerParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsSnapshotManagerParams() *AwsSnapshotManagerParams {
	this := AwsSnapshotManagerParams{}
	return &this
}

// NewAwsSnapshotManagerParamsWithDefaults instantiates a new AwsSnapshotManagerParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsSnapshotManagerParamsWithDefaults() *AwsSnapshotManagerParams {
	this := AwsSnapshotManagerParams{}
	return &this
}

// GetAmiCreationFrequency returns the AmiCreationFrequency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsSnapshotManagerParams) GetAmiCreationFrequency() int32 {
	if o == nil || IsNil(o.AmiCreationFrequency.Get()) {
		var ret int32
		return ret
	}
	return *o.AmiCreationFrequency.Get()
}

// GetAmiCreationFrequencyOk returns a tuple with the AmiCreationFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsSnapshotManagerParams) GetAmiCreationFrequencyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmiCreationFrequency.Get(), o.AmiCreationFrequency.IsSet()
}

// HasAmiCreationFrequency returns a boolean if a field has been set.
func (o *AwsSnapshotManagerParams) HasAmiCreationFrequency() bool {
	if o != nil && o.AmiCreationFrequency.IsSet() {
		return true
	}

	return false
}

// SetAmiCreationFrequency gets a reference to the given NullableInt32 and assigns it to the AmiCreationFrequency field.
func (o *AwsSnapshotManagerParams) SetAmiCreationFrequency(v int32) {
	o.AmiCreationFrequency.Set(&v)
}
// SetAmiCreationFrequencyNil sets the value for AmiCreationFrequency to be an explicit nil
func (o *AwsSnapshotManagerParams) SetAmiCreationFrequencyNil() {
	o.AmiCreationFrequency.Set(nil)
}

// UnsetAmiCreationFrequency ensures that no value is present for AmiCreationFrequency, not even an explicit nil
func (o *AwsSnapshotManagerParams) UnsetAmiCreationFrequency() {
	o.AmiCreationFrequency.Unset()
}

// GetCreateAmi returns the CreateAmi field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsSnapshotManagerParams) GetCreateAmi() bool {
	if o == nil || IsNil(o.CreateAmi.Get()) {
		var ret bool
		return ret
	}
	return *o.CreateAmi.Get()
}

// GetCreateAmiOk returns a tuple with the CreateAmi field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsSnapshotManagerParams) GetCreateAmiOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreateAmi.Get(), o.CreateAmi.IsSet()
}

// HasCreateAmi returns a boolean if a field has been set.
func (o *AwsSnapshotManagerParams) HasCreateAmi() bool {
	if o != nil && o.CreateAmi.IsSet() {
		return true
	}

	return false
}

// SetCreateAmi gets a reference to the given NullableBool and assigns it to the CreateAmi field.
func (o *AwsSnapshotManagerParams) SetCreateAmi(v bool) {
	o.CreateAmi.Set(&v)
}
// SetCreateAmiNil sets the value for CreateAmi to be an explicit nil
func (o *AwsSnapshotManagerParams) SetCreateAmiNil() {
	o.CreateAmi.Set(nil)
}

// UnsetCreateAmi ensures that no value is present for CreateAmi, not even an explicit nil
func (o *AwsSnapshotManagerParams) UnsetCreateAmi() {
	o.CreateAmi.Unset()
}

// GetVolumeExclusionParams returns the VolumeExclusionParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsSnapshotManagerParams) GetVolumeExclusionParams() EbsVolumeExclusionParams {
	if o == nil || IsNil(o.VolumeExclusionParams.Get()) {
		var ret EbsVolumeExclusionParams
		return ret
	}
	return *o.VolumeExclusionParams.Get()
}

// GetVolumeExclusionParamsOk returns a tuple with the VolumeExclusionParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsSnapshotManagerParams) GetVolumeExclusionParamsOk() (*EbsVolumeExclusionParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.VolumeExclusionParams.Get(), o.VolumeExclusionParams.IsSet()
}

// HasVolumeExclusionParams returns a boolean if a field has been set.
func (o *AwsSnapshotManagerParams) HasVolumeExclusionParams() bool {
	if o != nil && o.VolumeExclusionParams.IsSet() {
		return true
	}

	return false
}

// SetVolumeExclusionParams gets a reference to the given NullableEbsVolumeExclusionParams and assigns it to the VolumeExclusionParams field.
func (o *AwsSnapshotManagerParams) SetVolumeExclusionParams(v EbsVolumeExclusionParams) {
	o.VolumeExclusionParams.Set(&v)
}
// SetVolumeExclusionParamsNil sets the value for VolumeExclusionParams to be an explicit nil
func (o *AwsSnapshotManagerParams) SetVolumeExclusionParamsNil() {
	o.VolumeExclusionParams.Set(nil)
}

// UnsetVolumeExclusionParams ensures that no value is present for VolumeExclusionParams, not even an explicit nil
func (o *AwsSnapshotManagerParams) UnsetVolumeExclusionParams() {
	o.VolumeExclusionParams.Unset()
}

func (o AwsSnapshotManagerParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsSnapshotManagerParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AmiCreationFrequency.IsSet() {
		toSerialize["amiCreationFrequency"] = o.AmiCreationFrequency.Get()
	}
	if o.CreateAmi.IsSet() {
		toSerialize["createAmi"] = o.CreateAmi.Get()
	}
	if o.VolumeExclusionParams.IsSet() {
		toSerialize["volumeExclusionParams"] = o.VolumeExclusionParams.Get()
	}
	return toSerialize, nil
}

type NullableAwsSnapshotManagerParams struct {
	value *AwsSnapshotManagerParams
	isSet bool
}

func (v NullableAwsSnapshotManagerParams) Get() *AwsSnapshotManagerParams {
	return v.value
}

func (v *NullableAwsSnapshotManagerParams) Set(val *AwsSnapshotManagerParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsSnapshotManagerParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsSnapshotManagerParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsSnapshotManagerParams(val *AwsSnapshotManagerParams) *NullableAwsSnapshotManagerParams {
	return &NullableAwsSnapshotManagerParams{value: val, isSet: true}
}

func (v NullableAwsSnapshotManagerParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsSnapshotManagerParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


