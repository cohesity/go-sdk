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

// CommonObjectActionRequest Specifies the common request parameters for performing an action on Object.
type CommonObjectActionRequest struct {
	// Specifies the environment type of the Object.
	Environment NullableString `json:"environment"`
}

// NewCommonObjectActionRequest instantiates a new CommonObjectActionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonObjectActionRequest(environment NullableString) *CommonObjectActionRequest {
	this := CommonObjectActionRequest{}
	this.Environment = environment
	return &this
}

// NewCommonObjectActionRequestWithDefaults instantiates a new CommonObjectActionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonObjectActionRequestWithDefaults() *CommonObjectActionRequest {
	this := CommonObjectActionRequest{}
	return &this
}

// GetEnvironment returns the Environment field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommonObjectActionRequest) GetEnvironment() string {
	if o == nil || o.Environment.Get() == nil {
		var ret string
		return ret
	}

	return *o.Environment.Get()
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonObjectActionRequest) GetEnvironmentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Environment.Get(), o.Environment.IsSet()
}

// SetEnvironment sets field value
func (o *CommonObjectActionRequest) SetEnvironment(v string) {
	o.Environment.Set(&v)
}

func (o CommonObjectActionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["environment"] = o.Environment.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCommonObjectActionRequest struct {
	value *CommonObjectActionRequest
	isSet bool
}

func (v NullableCommonObjectActionRequest) Get() *CommonObjectActionRequest {
	return v.value
}

func (v *NullableCommonObjectActionRequest) Set(val *CommonObjectActionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonObjectActionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonObjectActionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonObjectActionRequest(val *CommonObjectActionRequest) *NullableCommonObjectActionRequest {
	return &NullableCommonObjectActionRequest{value: val, isSet: true}
}

func (v NullableCommonObjectActionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonObjectActionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


