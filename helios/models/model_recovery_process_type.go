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

// RecoveryProcessType Recovery Process Type
type RecoveryProcessType struct {
	// Specifies the recovery process type.
	RecoveryProcessType *string `json:"recoveryProcessType,omitempty"`
}

// NewRecoveryProcessType instantiates a new RecoveryProcessType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoveryProcessType() *RecoveryProcessType {
	this := RecoveryProcessType{}
	return &this
}

// NewRecoveryProcessTypeWithDefaults instantiates a new RecoveryProcessType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoveryProcessTypeWithDefaults() *RecoveryProcessType {
	this := RecoveryProcessType{}
	return &this
}

// GetRecoveryProcessType returns the RecoveryProcessType field value if set, zero value otherwise.
func (o *RecoveryProcessType) GetRecoveryProcessType() string {
	if o == nil || o.RecoveryProcessType == nil {
		var ret string
		return ret
	}
	return *o.RecoveryProcessType
}

// GetRecoveryProcessTypeOk returns a tuple with the RecoveryProcessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryProcessType) GetRecoveryProcessTypeOk() (*string, bool) {
	if o == nil || o.RecoveryProcessType == nil {
		return nil, false
	}
	return o.RecoveryProcessType, true
}

// HasRecoveryProcessType returns a boolean if a field has been set.
func (o *RecoveryProcessType) HasRecoveryProcessType() bool {
	if o != nil && o.RecoveryProcessType != nil {
		return true
	}

	return false
}

// SetRecoveryProcessType gets a reference to the given string and assigns it to the RecoveryProcessType field.
func (o *RecoveryProcessType) SetRecoveryProcessType(v string) {
	o.RecoveryProcessType = &v
}

func (o RecoveryProcessType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RecoveryProcessType != nil {
		toSerialize["recoveryProcessType"] = o.RecoveryProcessType
	}
	return json.Marshal(toSerialize)
}

type NullableRecoveryProcessType struct {
	value *RecoveryProcessType
	isSet bool
}

func (v NullableRecoveryProcessType) Get() *RecoveryProcessType {
	return v.value
}

func (v *NullableRecoveryProcessType) Set(val *RecoveryProcessType) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoveryProcessType) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoveryProcessType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoveryProcessType(val *RecoveryProcessType) *NullableRecoveryProcessType {
	return &NullableRecoveryProcessType{value: val, isSet: true}
}

func (v NullableRecoveryProcessType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoveryProcessType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


