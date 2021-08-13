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

// GcpProtectionGroupType GCP Protection Group type.
type GcpProtectionGroupType struct {
	// Specifies GCP Protection Group type.
	Environment *string `json:"environment,omitempty"`
}

// NewGcpProtectionGroupType instantiates a new GcpProtectionGroupType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGcpProtectionGroupType() *GcpProtectionGroupType {
	this := GcpProtectionGroupType{}
	return &this
}

// NewGcpProtectionGroupTypeWithDefaults instantiates a new GcpProtectionGroupType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGcpProtectionGroupTypeWithDefaults() *GcpProtectionGroupType {
	this := GcpProtectionGroupType{}
	return &this
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *GcpProtectionGroupType) GetEnvironment() string {
	if o == nil || o.Environment == nil {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpProtectionGroupType) GetEnvironmentOk() (*string, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *GcpProtectionGroupType) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *GcpProtectionGroupType) SetEnvironment(v string) {
	o.Environment = &v
}

func (o GcpProtectionGroupType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	return json.Marshal(toSerialize)
}

type NullableGcpProtectionGroupType struct {
	value *GcpProtectionGroupType
	isSet bool
}

func (v NullableGcpProtectionGroupType) Get() *GcpProtectionGroupType {
	return v.value
}

func (v *NullableGcpProtectionGroupType) Set(val *GcpProtectionGroupType) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpProtectionGroupType) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpProtectionGroupType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpProtectionGroupType(val *GcpProtectionGroupType) *NullableGcpProtectionGroupType {
	return &NullableGcpProtectionGroupType{value: val, isSet: true}
}

func (v NullableGcpProtectionGroupType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpProtectionGroupType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o GcpProtectionGroupType) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}