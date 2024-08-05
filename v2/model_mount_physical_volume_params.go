/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the MountPhysicalVolumeParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MountPhysicalVolumeParams{}

// MountPhysicalVolumeParams Specifies the parameters to Mount Physical Volumes.
type MountPhysicalVolumeParams struct {
	PhysicalTargetParams NullableMountPhysicalVolumeParamsPhysicalTargetParams `json:"physicalTargetParams,omitempty"`
	// Specifies the environment of the recovery target. The corresponding params below must be filled out.
	TargetEnvironment string `json:"targetEnvironment"`
}

type _MountPhysicalVolumeParams MountPhysicalVolumeParams

// NewMountPhysicalVolumeParams instantiates a new MountPhysicalVolumeParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMountPhysicalVolumeParams(targetEnvironment string) *MountPhysicalVolumeParams {
	this := MountPhysicalVolumeParams{}
	this.TargetEnvironment = targetEnvironment
	return &this
}

// NewMountPhysicalVolumeParamsWithDefaults instantiates a new MountPhysicalVolumeParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMountPhysicalVolumeParamsWithDefaults() *MountPhysicalVolumeParams {
	this := MountPhysicalVolumeParams{}
	return &this
}

// GetPhysicalTargetParams returns the PhysicalTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MountPhysicalVolumeParams) GetPhysicalTargetParams() MountPhysicalVolumeParamsPhysicalTargetParams {
	if o == nil || IsNil(o.PhysicalTargetParams.Get()) {
		var ret MountPhysicalVolumeParamsPhysicalTargetParams
		return ret
	}
	return *o.PhysicalTargetParams.Get()
}

// GetPhysicalTargetParamsOk returns a tuple with the PhysicalTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MountPhysicalVolumeParams) GetPhysicalTargetParamsOk() (*MountPhysicalVolumeParamsPhysicalTargetParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.PhysicalTargetParams.Get(), o.PhysicalTargetParams.IsSet()
}

// HasPhysicalTargetParams returns a boolean if a field has been set.
func (o *MountPhysicalVolumeParams) HasPhysicalTargetParams() bool {
	if o != nil && o.PhysicalTargetParams.IsSet() {
		return true
	}

	return false
}

// SetPhysicalTargetParams gets a reference to the given NullableMountPhysicalVolumeParamsPhysicalTargetParams and assigns it to the PhysicalTargetParams field.
func (o *MountPhysicalVolumeParams) SetPhysicalTargetParams(v MountPhysicalVolumeParamsPhysicalTargetParams) {
	o.PhysicalTargetParams.Set(&v)
}
// SetPhysicalTargetParamsNil sets the value for PhysicalTargetParams to be an explicit nil
func (o *MountPhysicalVolumeParams) SetPhysicalTargetParamsNil() {
	o.PhysicalTargetParams.Set(nil)
}

// UnsetPhysicalTargetParams ensures that no value is present for PhysicalTargetParams, not even an explicit nil
func (o *MountPhysicalVolumeParams) UnsetPhysicalTargetParams() {
	o.PhysicalTargetParams.Unset()
}

// GetTargetEnvironment returns the TargetEnvironment field value
func (o *MountPhysicalVolumeParams) GetTargetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetEnvironment
}

// GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field value
// and a boolean to check if the value has been set.
func (o *MountPhysicalVolumeParams) GetTargetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetEnvironment, true
}

// SetTargetEnvironment sets field value
func (o *MountPhysicalVolumeParams) SetTargetEnvironment(v string) {
	o.TargetEnvironment = v
}

func (o MountPhysicalVolumeParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MountPhysicalVolumeParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PhysicalTargetParams.IsSet() {
		toSerialize["physicalTargetParams"] = o.PhysicalTargetParams.Get()
	}
	toSerialize["targetEnvironment"] = o.TargetEnvironment
	return toSerialize, nil
}

func (o *MountPhysicalVolumeParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"targetEnvironment",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varMountPhysicalVolumeParams := _MountPhysicalVolumeParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMountPhysicalVolumeParams)

	if err != nil {
		return err
	}

	*o = MountPhysicalVolumeParams(varMountPhysicalVolumeParams)

	return err
}

type NullableMountPhysicalVolumeParams struct {
	value *MountPhysicalVolumeParams
	isSet bool
}

func (v NullableMountPhysicalVolumeParams) Get() *MountPhysicalVolumeParams {
	return v.value
}

func (v *NullableMountPhysicalVolumeParams) Set(val *MountPhysicalVolumeParams) {
	v.value = val
	v.isSet = true
}

func (v NullableMountPhysicalVolumeParams) IsSet() bool {
	return v.isSet
}

func (v *NullableMountPhysicalVolumeParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMountPhysicalVolumeParams(val *MountPhysicalVolumeParams) *NullableMountPhysicalVolumeParams {
	return &NullableMountPhysicalVolumeParams{value: val, isSet: true}
}

func (v NullableMountPhysicalVolumeParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMountPhysicalVolumeParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


