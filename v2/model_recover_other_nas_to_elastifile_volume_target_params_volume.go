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

// checks if the RecoverOtherNasToElastifileVolumeTargetParamsVolume type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoverOtherNasToElastifileVolumeTargetParamsVolume{}

// RecoverOtherNasToElastifileVolumeTargetParamsVolume Specifies the id and name of the parent volume to recover to. This volume will be the target of the recovery.
type RecoverOtherNasToElastifileVolumeTargetParamsVolume struct {
	// Specifies the id of the object.
	Id NullableInt64 `json:"id"`
	// Specifies the name of the object.
	Name NullableString `json:"name,omitempty"`
}

type _RecoverOtherNasToElastifileVolumeTargetParamsVolume RecoverOtherNasToElastifileVolumeTargetParamsVolume

// NewRecoverOtherNasToElastifileVolumeTargetParamsVolume instantiates a new RecoverOtherNasToElastifileVolumeTargetParamsVolume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverOtherNasToElastifileVolumeTargetParamsVolume(id NullableInt64) *RecoverOtherNasToElastifileVolumeTargetParamsVolume {
	this := RecoverOtherNasToElastifileVolumeTargetParamsVolume{}
	this.Id = id
	return &this
}

// NewRecoverOtherNasToElastifileVolumeTargetParamsVolumeWithDefaults instantiates a new RecoverOtherNasToElastifileVolumeTargetParamsVolume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverOtherNasToElastifileVolumeTargetParamsVolumeWithDefaults() *RecoverOtherNasToElastifileVolumeTargetParamsVolume {
	this := RecoverOtherNasToElastifileVolumeTargetParamsVolume{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *RecoverOtherNasToElastifileVolumeTargetParamsVolume) GetId() int64 {
	if o == nil || o.Id.Get() == nil {
		var ret int64
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverOtherNasToElastifileVolumeTargetParamsVolume) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *RecoverOtherNasToElastifileVolumeTargetParamsVolume) SetId(v int64) {
	o.Id.Set(&v)
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverOtherNasToElastifileVolumeTargetParamsVolume) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverOtherNasToElastifileVolumeTargetParamsVolume) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *RecoverOtherNasToElastifileVolumeTargetParamsVolume) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *RecoverOtherNasToElastifileVolumeTargetParamsVolume) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *RecoverOtherNasToElastifileVolumeTargetParamsVolume) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *RecoverOtherNasToElastifileVolumeTargetParamsVolume) UnsetName() {
	o.Name.Unset()
}

func (o RecoverOtherNasToElastifileVolumeTargetParamsVolume) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoverOtherNasToElastifileVolumeTargetParamsVolume) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return toSerialize, nil
}

func (o *RecoverOtherNasToElastifileVolumeTargetParamsVolume) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varRecoverOtherNasToElastifileVolumeTargetParamsVolume := _RecoverOtherNasToElastifileVolumeTargetParamsVolume{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRecoverOtherNasToElastifileVolumeTargetParamsVolume)

	if err != nil {
		return err
	}

	*o = RecoverOtherNasToElastifileVolumeTargetParamsVolume(varRecoverOtherNasToElastifileVolumeTargetParamsVolume)

	return err
}

type NullableRecoverOtherNasToElastifileVolumeTargetParamsVolume struct {
	value *RecoverOtherNasToElastifileVolumeTargetParamsVolume
	isSet bool
}

func (v NullableRecoverOtherNasToElastifileVolumeTargetParamsVolume) Get() *RecoverOtherNasToElastifileVolumeTargetParamsVolume {
	return v.value
}

func (v *NullableRecoverOtherNasToElastifileVolumeTargetParamsVolume) Set(val *RecoverOtherNasToElastifileVolumeTargetParamsVolume) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverOtherNasToElastifileVolumeTargetParamsVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverOtherNasToElastifileVolumeTargetParamsVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverOtherNasToElastifileVolumeTargetParamsVolume(val *RecoverOtherNasToElastifileVolumeTargetParamsVolume) *NullableRecoverOtherNasToElastifileVolumeTargetParamsVolume {
	return &NullableRecoverOtherNasToElastifileVolumeTargetParamsVolume{value: val, isSet: true}
}

func (v NullableRecoverOtherNasToElastifileVolumeTargetParamsVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverOtherNasToElastifileVolumeTargetParamsVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


