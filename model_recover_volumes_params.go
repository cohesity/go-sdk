/*
 * Cohesity REST API
 *
 * This API list provides operations for interfacing with the Cohesity Cluster.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cohesitysdk

import (
	"encoding/json"
)

// RecoverVolumesParams struct for RecoverVolumesParams
type RecoverVolumesParams struct {
	// Whether volume would be dismounted first during LockVolume failure
	ForceUnmountVolume NullableBool `json:"forceUnmountVolume,omitempty"`
	// Contains the volume mapping data that defines the restore task.
	MappingVec []RecoverVolumesParamsMapping `json:"mappingVec,omitempty"`
	TargetEntity *EntityProto `json:"targetEntity,omitempty"`
}

// NewRecoverVolumesParams instantiates a new RecoverVolumesParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverVolumesParams() *RecoverVolumesParams {
	this := RecoverVolumesParams{}
	return &this
}

// NewRecoverVolumesParamsWithDefaults instantiates a new RecoverVolumesParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverVolumesParamsWithDefaults() *RecoverVolumesParams {
	this := RecoverVolumesParams{}
	return &this
}

// GetForceUnmountVolume returns the ForceUnmountVolume field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverVolumesParams) GetForceUnmountVolume() bool {
	if o == nil || o.ForceUnmountVolume.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ForceUnmountVolume.Get()
}

// GetForceUnmountVolumeOk returns a tuple with the ForceUnmountVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverVolumesParams) GetForceUnmountVolumeOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ForceUnmountVolume.Get(), o.ForceUnmountVolume.IsSet()
}

// HasForceUnmountVolume returns a boolean if a field has been set.
func (o *RecoverVolumesParams) HasForceUnmountVolume() bool {
	if o != nil && o.ForceUnmountVolume.IsSet() {
		return true
	}

	return false
}

// SetForceUnmountVolume gets a reference to the given NullableBool and assigns it to the ForceUnmountVolume field.
func (o *RecoverVolumesParams) SetForceUnmountVolume(v bool) {
	o.ForceUnmountVolume.Set(&v)
}
// SetForceUnmountVolumeNil sets the value for ForceUnmountVolume to be an explicit nil
func (o *RecoverVolumesParams) SetForceUnmountVolumeNil() {
	o.ForceUnmountVolume.Set(nil)
}

// UnsetForceUnmountVolume ensures that no value is present for ForceUnmountVolume, not even an explicit nil
func (o *RecoverVolumesParams) UnsetForceUnmountVolume() {
	o.ForceUnmountVolume.Unset()
}

// GetMappingVec returns the MappingVec field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverVolumesParams) GetMappingVec() []RecoverVolumesParamsMapping {
	if o == nil  {
		var ret []RecoverVolumesParamsMapping
		return ret
	}
	return o.MappingVec
}

// GetMappingVecOk returns a tuple with the MappingVec field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverVolumesParams) GetMappingVecOk() (*[]RecoverVolumesParamsMapping, bool) {
	if o == nil || o.MappingVec == nil {
		return nil, false
	}
	return &o.MappingVec, true
}

// HasMappingVec returns a boolean if a field has been set.
func (o *RecoverVolumesParams) HasMappingVec() bool {
	if o != nil && o.MappingVec != nil {
		return true
	}

	return false
}

// SetMappingVec gets a reference to the given []RecoverVolumesParamsMapping and assigns it to the MappingVec field.
func (o *RecoverVolumesParams) SetMappingVec(v []RecoverVolumesParamsMapping) {
	o.MappingVec = v
}

// GetTargetEntity returns the TargetEntity field value if set, zero value otherwise.
func (o *RecoverVolumesParams) GetTargetEntity() EntityProto {
	if o == nil || o.TargetEntity == nil {
		var ret EntityProto
		return ret
	}
	return *o.TargetEntity
}

// GetTargetEntityOk returns a tuple with the TargetEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoverVolumesParams) GetTargetEntityOk() (*EntityProto, bool) {
	if o == nil || o.TargetEntity == nil {
		return nil, false
	}
	return o.TargetEntity, true
}

// HasTargetEntity returns a boolean if a field has been set.
func (o *RecoverVolumesParams) HasTargetEntity() bool {
	if o != nil && o.TargetEntity != nil {
		return true
	}

	return false
}

// SetTargetEntity gets a reference to the given EntityProto and assigns it to the TargetEntity field.
func (o *RecoverVolumesParams) SetTargetEntity(v EntityProto) {
	o.TargetEntity = &v
}

func (o RecoverVolumesParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ForceUnmountVolume.IsSet() {
		toSerialize["forceUnmountVolume"] = o.ForceUnmountVolume.Get()
	}
	if o.MappingVec != nil {
		toSerialize["mappingVec"] = o.MappingVec
	}
	if o.TargetEntity != nil {
		toSerialize["targetEntity"] = o.TargetEntity
	}
	return json.Marshal(toSerialize)
}

type NullableRecoverVolumesParams struct {
	value *RecoverVolumesParams
	isSet bool
}

func (v NullableRecoverVolumesParams) Get() *RecoverVolumesParams {
	return v.value
}

func (v *NullableRecoverVolumesParams) Set(val *RecoverVolumesParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverVolumesParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverVolumesParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverVolumesParams(val *RecoverVolumesParams) *NullableRecoverVolumesParams {
	return &NullableRecoverVolumesParams{value: val, isSet: true}
}

func (v NullableRecoverVolumesParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverVolumesParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


