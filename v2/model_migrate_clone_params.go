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

// checks if the MigrateCloneParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MigrateCloneParams{}

// MigrateCloneParams Specifies the DB migration parameters.
type MigrateCloneParams struct {
	// Specifies when the migration of the oracle instance should be started after successful recovery.
	DelaySecs NullableInt64 `json:"delaySecs,omitempty"`
	// Specifies the target paths to be used for DB migration.
	TargetPathVec []string `json:"targetPathVec,omitempty"`
}

// NewMigrateCloneParams instantiates a new MigrateCloneParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMigrateCloneParams() *MigrateCloneParams {
	this := MigrateCloneParams{}
	return &this
}

// NewMigrateCloneParamsWithDefaults instantiates a new MigrateCloneParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMigrateCloneParamsWithDefaults() *MigrateCloneParams {
	this := MigrateCloneParams{}
	return &this
}

// GetDelaySecs returns the DelaySecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MigrateCloneParams) GetDelaySecs() int64 {
	if o == nil || IsNil(o.DelaySecs.Get()) {
		var ret int64
		return ret
	}
	return *o.DelaySecs.Get()
}

// GetDelaySecsOk returns a tuple with the DelaySecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MigrateCloneParams) GetDelaySecsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DelaySecs.Get(), o.DelaySecs.IsSet()
}

// HasDelaySecs returns a boolean if a field has been set.
func (o *MigrateCloneParams) HasDelaySecs() bool {
	if o != nil && o.DelaySecs.IsSet() {
		return true
	}

	return false
}

// SetDelaySecs gets a reference to the given NullableInt64 and assigns it to the DelaySecs field.
func (o *MigrateCloneParams) SetDelaySecs(v int64) {
	o.DelaySecs.Set(&v)
}
// SetDelaySecsNil sets the value for DelaySecs to be an explicit nil
func (o *MigrateCloneParams) SetDelaySecsNil() {
	o.DelaySecs.Set(nil)
}

// UnsetDelaySecs ensures that no value is present for DelaySecs, not even an explicit nil
func (o *MigrateCloneParams) UnsetDelaySecs() {
	o.DelaySecs.Unset()
}

// GetTargetPathVec returns the TargetPathVec field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MigrateCloneParams) GetTargetPathVec() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TargetPathVec
}

// GetTargetPathVecOk returns a tuple with the TargetPathVec field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MigrateCloneParams) GetTargetPathVecOk() ([]string, bool) {
	if o == nil || IsNil(o.TargetPathVec) {
		return nil, false
	}
	return o.TargetPathVec, true
}

// HasTargetPathVec returns a boolean if a field has been set.
func (o *MigrateCloneParams) HasTargetPathVec() bool {
	if o != nil && !IsNil(o.TargetPathVec) {
		return true
	}

	return false
}

// SetTargetPathVec gets a reference to the given []string and assigns it to the TargetPathVec field.
func (o *MigrateCloneParams) SetTargetPathVec(v []string) {
	o.TargetPathVec = v
}

func (o MigrateCloneParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MigrateCloneParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DelaySecs.IsSet() {
		toSerialize["delaySecs"] = o.DelaySecs.Get()
	}
	if o.TargetPathVec != nil {
		toSerialize["targetPathVec"] = o.TargetPathVec
	}
	return toSerialize, nil
}

type NullableMigrateCloneParams struct {
	value *MigrateCloneParams
	isSet bool
}

func (v NullableMigrateCloneParams) Get() *MigrateCloneParams {
	return v.value
}

func (v *NullableMigrateCloneParams) Set(val *MigrateCloneParams) {
	v.value = val
	v.isSet = true
}

func (v NullableMigrateCloneParams) IsSet() bool {
	return v.isSet
}

func (v *NullableMigrateCloneParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMigrateCloneParams(val *MigrateCloneParams) *NullableMigrateCloneParams {
	return &NullableMigrateCloneParams{value: val, isSet: true}
}

func (v NullableMigrateCloneParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMigrateCloneParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


