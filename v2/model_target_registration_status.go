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

// checks if the TargetRegistrationStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetRegistrationStatus{}

// TargetRegistrationStatus Specifies the registration status of the external target.
type TargetRegistrationStatus struct {
	// Specifies the registration status of the external target.
	Enum *string `json:"enum,omitempty"`
}

// NewTargetRegistrationStatus instantiates a new TargetRegistrationStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetRegistrationStatus() *TargetRegistrationStatus {
	this := TargetRegistrationStatus{}
	return &this
}

// NewTargetRegistrationStatusWithDefaults instantiates a new TargetRegistrationStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetRegistrationStatusWithDefaults() *TargetRegistrationStatus {
	this := TargetRegistrationStatus{}
	return &this
}

// GetEnum returns the Enum field value if set, zero value otherwise.
func (o *TargetRegistrationStatus) GetEnum() string {
	if o == nil || IsNil(o.Enum) {
		var ret string
		return ret
	}
	return *o.Enum
}

// GetEnumOk returns a tuple with the Enum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetRegistrationStatus) GetEnumOk() (*string, bool) {
	if o == nil || IsNil(o.Enum) {
		return nil, false
	}
	return o.Enum, true
}

// HasEnum returns a boolean if a field has been set.
func (o *TargetRegistrationStatus) HasEnum() bool {
	if o != nil && !IsNil(o.Enum) {
		return true
	}

	return false
}

// SetEnum gets a reference to the given string and assigns it to the Enum field.
func (o *TargetRegistrationStatus) SetEnum(v string) {
	o.Enum = &v
}

func (o TargetRegistrationStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TargetRegistrationStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enum) {
		toSerialize["enum"] = o.Enum
	}
	return toSerialize, nil
}

type NullableTargetRegistrationStatus struct {
	value *TargetRegistrationStatus
	isSet bool
}

func (v NullableTargetRegistrationStatus) Get() *TargetRegistrationStatus {
	return v.value
}

func (v *NullableTargetRegistrationStatus) Set(val *TargetRegistrationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetRegistrationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetRegistrationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetRegistrationStatus(val *TargetRegistrationStatus) *NullableTargetRegistrationStatus {
	return &NullableTargetRegistrationStatus{value: val, isSet: true}
}

func (v NullableTargetRegistrationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetRegistrationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


