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

// RecoveryType Recovery Type
type RecoveryType struct {
	// Specifies the recovery types.
	RecoveryType *string `json:"recoveryType,omitempty"`
}

// NewRecoveryType instantiates a new RecoveryType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoveryType() *RecoveryType {
	this := RecoveryType{}
	return &this
}

// NewRecoveryTypeWithDefaults instantiates a new RecoveryType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoveryTypeWithDefaults() *RecoveryType {
	this := RecoveryType{}
	return &this
}

// GetRecoveryType returns the RecoveryType field value if set, zero value otherwise.
func (o *RecoveryType) GetRecoveryType() string {
	if o == nil || o.RecoveryType == nil {
		var ret string
		return ret
	}
	return *o.RecoveryType
}

// GetRecoveryTypeOk returns a tuple with the RecoveryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryType) GetRecoveryTypeOk() (*string, bool) {
	if o == nil || o.RecoveryType == nil {
		return nil, false
	}
	return o.RecoveryType, true
}

// HasRecoveryType returns a boolean if a field has been set.
func (o *RecoveryType) HasRecoveryType() bool {
	if o != nil && o.RecoveryType != nil {
		return true
	}

	return false
}

// SetRecoveryType gets a reference to the given string and assigns it to the RecoveryType field.
func (o *RecoveryType) SetRecoveryType(v string) {
	o.RecoveryType = &v
}

func (o RecoveryType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RecoveryType != nil {
		toSerialize["recoveryType"] = o.RecoveryType
	}
	return json.Marshal(toSerialize)
}

type NullableRecoveryType struct {
	value *RecoveryType
	isSet bool
}

func (v NullableRecoveryType) Get() *RecoveryType {
	return v.value
}

func (v *NullableRecoveryType) Set(val *RecoveryType) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoveryType) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoveryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoveryType(val *RecoveryType) *NullableRecoveryType {
	return &NullableRecoveryType{value: val, isSet: true}
}

func (v NullableRecoveryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoveryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


