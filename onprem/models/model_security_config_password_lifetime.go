/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package onprem
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/onprem/utils"
	"fmt"
)

var _ = NullableBool{}

// SecurityConfigPasswordLifetime Specifies security config for password lifetime.
type SecurityConfigPasswordLifetime struct {
	// Specifies the minimum password lifetime in days.
	MinLifetimeDays NullableInt32 `json:"minLifetimeDays,omitempty"`
	// Specifies the maximum password lifetime in days.
	MaxLifetimeDays NullableInt32 `json:"maxLifetimeDays,omitempty"`
}

// NewSecurityConfigPasswordLifetime instantiates a new SecurityConfigPasswordLifetime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityConfigPasswordLifetime() *SecurityConfigPasswordLifetime {
	this := SecurityConfigPasswordLifetime{}
	return &this
}

// NewSecurityConfigPasswordLifetimeWithDefaults instantiates a new SecurityConfigPasswordLifetime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityConfigPasswordLifetimeWithDefaults() *SecurityConfigPasswordLifetime {
	this := SecurityConfigPasswordLifetime{}
	return &this
}

// GetMinLifetimeDays returns the MinLifetimeDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityConfigPasswordLifetime) GetMinLifetimeDays() int32 {
	if o == nil || o.MinLifetimeDays.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MinLifetimeDays.Get()
}

// GetMinLifetimeDaysOk returns a tuple with the MinLifetimeDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecurityConfigPasswordLifetime) GetMinLifetimeDaysOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MinLifetimeDays.Get(), o.MinLifetimeDays.IsSet()
}

// HasMinLifetimeDays returns a boolean if a field has been set.
func (o *SecurityConfigPasswordLifetime) HasMinLifetimeDays() bool {
	if o != nil && o.MinLifetimeDays.IsSet() {
		return true
	}

	return false
}

// SetMinLifetimeDays gets a reference to the given NullableInt32 and assigns it to the MinLifetimeDays field.
func (o *SecurityConfigPasswordLifetime) SetMinLifetimeDays(v int32) {
	o.MinLifetimeDays.Set(&v)
}
// SetMinLifetimeDaysNil sets the value for MinLifetimeDays to be an explicit nil
func (o *SecurityConfigPasswordLifetime) SetMinLifetimeDaysNil() {
	o.MinLifetimeDays.Set(nil)
}

// UnsetMinLifetimeDays ensures that no value is present for MinLifetimeDays, not even an explicit nil
func (o *SecurityConfigPasswordLifetime) UnsetMinLifetimeDays() {
	o.MinLifetimeDays.Unset()
}

// GetMaxLifetimeDays returns the MaxLifetimeDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityConfigPasswordLifetime) GetMaxLifetimeDays() int32 {
	if o == nil || o.MaxLifetimeDays.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MaxLifetimeDays.Get()
}

// GetMaxLifetimeDaysOk returns a tuple with the MaxLifetimeDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecurityConfigPasswordLifetime) GetMaxLifetimeDaysOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MaxLifetimeDays.Get(), o.MaxLifetimeDays.IsSet()
}

// HasMaxLifetimeDays returns a boolean if a field has been set.
func (o *SecurityConfigPasswordLifetime) HasMaxLifetimeDays() bool {
	if o != nil && o.MaxLifetimeDays.IsSet() {
		return true
	}

	return false
}

// SetMaxLifetimeDays gets a reference to the given NullableInt32 and assigns it to the MaxLifetimeDays field.
func (o *SecurityConfigPasswordLifetime) SetMaxLifetimeDays(v int32) {
	o.MaxLifetimeDays.Set(&v)
}
// SetMaxLifetimeDaysNil sets the value for MaxLifetimeDays to be an explicit nil
func (o *SecurityConfigPasswordLifetime) SetMaxLifetimeDaysNil() {
	o.MaxLifetimeDays.Set(nil)
}

// UnsetMaxLifetimeDays ensures that no value is present for MaxLifetimeDays, not even an explicit nil
func (o *SecurityConfigPasswordLifetime) UnsetMaxLifetimeDays() {
	o.MaxLifetimeDays.Unset()
}

func (o SecurityConfigPasswordLifetime) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MinLifetimeDays.IsSet() {
		toSerialize["minLifetimeDays"] = o.MinLifetimeDays.Get()
	}
	if o.MaxLifetimeDays.IsSet() {
		toSerialize["maxLifetimeDays"] = o.MaxLifetimeDays.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityConfigPasswordLifetime struct {
	value *SecurityConfigPasswordLifetime
	isSet bool
}

func (v NullableSecurityConfigPasswordLifetime) Get() *SecurityConfigPasswordLifetime {
	return v.value
}

func (v *NullableSecurityConfigPasswordLifetime) Set(val *SecurityConfigPasswordLifetime) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityConfigPasswordLifetime) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityConfigPasswordLifetime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityConfigPasswordLifetime(val *SecurityConfigPasswordLifetime) *NullableSecurityConfigPasswordLifetime {
	return &NullableSecurityConfigPasswordLifetime{value: val, isSet: true}
}

func (v NullableSecurityConfigPasswordLifetime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityConfigPasswordLifetime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o SecurityConfigPasswordLifetime) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}