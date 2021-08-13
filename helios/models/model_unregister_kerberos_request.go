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

// UnregisterKerberosRequest Specifies the request to unregister a Kerberos Provider.
type UnregisterKerberosRequest struct {
	// Specifies the password
	AdminPassword NullableString `json:"adminPassword"`
}

// NewUnregisterKerberosRequest instantiates a new UnregisterKerberosRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnregisterKerberosRequest(adminPassword NullableString) *UnregisterKerberosRequest {
	this := UnregisterKerberosRequest{}
	this.AdminPassword = adminPassword
	return &this
}

// NewUnregisterKerberosRequestWithDefaults instantiates a new UnregisterKerberosRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnregisterKerberosRequestWithDefaults() *UnregisterKerberosRequest {
	this := UnregisterKerberosRequest{}
	return &this
}

// GetAdminPassword returns the AdminPassword field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UnregisterKerberosRequest) GetAdminPassword() string {
	if o == nil || o.AdminPassword.Get() == nil {
		var ret string
		return ret
	}

	return *o.AdminPassword.Get()
}

// GetAdminPasswordOk returns a tuple with the AdminPassword field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnregisterKerberosRequest) GetAdminPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AdminPassword.Get(), o.AdminPassword.IsSet()
}

// SetAdminPassword sets field value
func (o *UnregisterKerberosRequest) SetAdminPassword(v string) {
	o.AdminPassword.Set(&v)
}

func (o UnregisterKerberosRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["adminPassword"] = o.AdminPassword.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableUnregisterKerberosRequest struct {
	value *UnregisterKerberosRequest
	isSet bool
}

func (v NullableUnregisterKerberosRequest) Get() *UnregisterKerberosRequest {
	return v.value
}

func (v *NullableUnregisterKerberosRequest) Set(val *UnregisterKerberosRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUnregisterKerberosRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUnregisterKerberosRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnregisterKerberosRequest(val *UnregisterKerberosRequest) *NullableUnregisterKerberosRequest {
	return &NullableUnregisterKerberosRequest{value: val, isSet: true}
}

func (v NullableUnregisterKerberosRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnregisterKerberosRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


