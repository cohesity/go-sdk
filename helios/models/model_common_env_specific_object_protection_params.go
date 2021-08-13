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

// CommonEnvSpecificObjectProtectionParams Specifies common properties of Object Protection Operations
type CommonEnvSpecificObjectProtectionParams struct {
	// Specifies the environment for current object.
	Environment NullableString `json:"environment,omitempty"`
}

// NewCommonEnvSpecificObjectProtectionParams instantiates a new CommonEnvSpecificObjectProtectionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonEnvSpecificObjectProtectionParams() *CommonEnvSpecificObjectProtectionParams {
	this := CommonEnvSpecificObjectProtectionParams{}
	return &this
}

// NewCommonEnvSpecificObjectProtectionParamsWithDefaults instantiates a new CommonEnvSpecificObjectProtectionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonEnvSpecificObjectProtectionParamsWithDefaults() *CommonEnvSpecificObjectProtectionParams {
	this := CommonEnvSpecificObjectProtectionParams{}
	return &this
}

// GetEnvironment returns the Environment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonEnvSpecificObjectProtectionParams) GetEnvironment() string {
	if o == nil || o.Environment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Environment.Get()
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonEnvSpecificObjectProtectionParams) GetEnvironmentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Environment.Get(), o.Environment.IsSet()
}

// HasEnvironment returns a boolean if a field has been set.
func (o *CommonEnvSpecificObjectProtectionParams) HasEnvironment() bool {
	if o != nil && o.Environment.IsSet() {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given NullableString and assigns it to the Environment field.
func (o *CommonEnvSpecificObjectProtectionParams) SetEnvironment(v string) {
	o.Environment.Set(&v)
}
// SetEnvironmentNil sets the value for Environment to be an explicit nil
func (o *CommonEnvSpecificObjectProtectionParams) SetEnvironmentNil() {
	o.Environment.Set(nil)
}

// UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
func (o *CommonEnvSpecificObjectProtectionParams) UnsetEnvironment() {
	o.Environment.Unset()
}

func (o CommonEnvSpecificObjectProtectionParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Environment.IsSet() {
		toSerialize["environment"] = o.Environment.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCommonEnvSpecificObjectProtectionParams struct {
	value *CommonEnvSpecificObjectProtectionParams
	isSet bool
}

func (v NullableCommonEnvSpecificObjectProtectionParams) Get() *CommonEnvSpecificObjectProtectionParams {
	return v.value
}

func (v *NullableCommonEnvSpecificObjectProtectionParams) Set(val *CommonEnvSpecificObjectProtectionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonEnvSpecificObjectProtectionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonEnvSpecificObjectProtectionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonEnvSpecificObjectProtectionParams(val *CommonEnvSpecificObjectProtectionParams) *NullableCommonEnvSpecificObjectProtectionParams {
	return &NullableCommonEnvSpecificObjectProtectionParams{value: val, isSet: true}
}

func (v NullableCommonEnvSpecificObjectProtectionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonEnvSpecificObjectProtectionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


