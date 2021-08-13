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

// ObjectTypeVCenterParams struct for ObjectTypeVCenterParams
type ObjectTypeVCenterParams struct {
	// Specifies that registered vCenter source is a VMC (VMware Cloud) environment or not.
	IsCloudEnv NullableBool `json:"isCloudEnv,omitempty"`
}

// NewObjectTypeVCenterParams instantiates a new ObjectTypeVCenterParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectTypeVCenterParams() *ObjectTypeVCenterParams {
	this := ObjectTypeVCenterParams{}
	return &this
}

// NewObjectTypeVCenterParamsWithDefaults instantiates a new ObjectTypeVCenterParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectTypeVCenterParamsWithDefaults() *ObjectTypeVCenterParams {
	this := ObjectTypeVCenterParams{}
	return &this
}

// GetIsCloudEnv returns the IsCloudEnv field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObjectTypeVCenterParams) GetIsCloudEnv() bool {
	if o == nil || o.IsCloudEnv.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsCloudEnv.Get()
}

// GetIsCloudEnvOk returns a tuple with the IsCloudEnv field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectTypeVCenterParams) GetIsCloudEnvOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsCloudEnv.Get(), o.IsCloudEnv.IsSet()
}

// HasIsCloudEnv returns a boolean if a field has been set.
func (o *ObjectTypeVCenterParams) HasIsCloudEnv() bool {
	if o != nil && o.IsCloudEnv.IsSet() {
		return true
	}

	return false
}

// SetIsCloudEnv gets a reference to the given NullableBool and assigns it to the IsCloudEnv field.
func (o *ObjectTypeVCenterParams) SetIsCloudEnv(v bool) {
	o.IsCloudEnv.Set(&v)
}
// SetIsCloudEnvNil sets the value for IsCloudEnv to be an explicit nil
func (o *ObjectTypeVCenterParams) SetIsCloudEnvNil() {
	o.IsCloudEnv.Set(nil)
}

// UnsetIsCloudEnv ensures that no value is present for IsCloudEnv, not even an explicit nil
func (o *ObjectTypeVCenterParams) UnsetIsCloudEnv() {
	o.IsCloudEnv.Unset()
}

func (o ObjectTypeVCenterParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsCloudEnv.IsSet() {
		toSerialize["isCloudEnv"] = o.IsCloudEnv.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableObjectTypeVCenterParams struct {
	value *ObjectTypeVCenterParams
	isSet bool
}

func (v NullableObjectTypeVCenterParams) Get() *ObjectTypeVCenterParams {
	return v.value
}

func (v *NullableObjectTypeVCenterParams) Set(val *ObjectTypeVCenterParams) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectTypeVCenterParams) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectTypeVCenterParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectTypeVCenterParams(val *ObjectTypeVCenterParams) *NullableObjectTypeVCenterParams {
	return &NullableObjectTypeVCenterParams{value: val, isSet: true}
}

func (v NullableObjectTypeVCenterParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectTypeVCenterParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


