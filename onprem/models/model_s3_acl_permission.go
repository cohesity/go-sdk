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

// S3AclPermission Specifies S3 ACL permission type.
type S3AclPermission struct {
	// Specifies S3 ACL permission type.
	Enum *string `json:"enum,omitempty"`
}

// NewS3AclPermission instantiates a new S3AclPermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3AclPermission() *S3AclPermission {
	this := S3AclPermission{}
	return &this
}

// NewS3AclPermissionWithDefaults instantiates a new S3AclPermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3AclPermissionWithDefaults() *S3AclPermission {
	this := S3AclPermission{}
	return &this
}

// GetEnum returns the Enum field value if set, zero value otherwise.
func (o *S3AclPermission) GetEnum() string {
	if o == nil || o.Enum == nil {
		var ret string
		return ret
	}
	return *o.Enum
}

// GetEnumOk returns a tuple with the Enum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3AclPermission) GetEnumOk() (*string, bool) {
	if o == nil || o.Enum == nil {
		return nil, false
	}
	return o.Enum, true
}

// HasEnum returns a boolean if a field has been set.
func (o *S3AclPermission) HasEnum() bool {
	if o != nil && o.Enum != nil {
		return true
	}

	return false
}

// SetEnum gets a reference to the given string and assigns it to the Enum field.
func (o *S3AclPermission) SetEnum(v string) {
	o.Enum = &v
}

func (o S3AclPermission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enum != nil {
		toSerialize["enum"] = o.Enum
	}
	return json.Marshal(toSerialize)
}

type NullableS3AclPermission struct {
	value *S3AclPermission
	isSet bool
}

func (v NullableS3AclPermission) Get() *S3AclPermission {
	return v.value
}

func (v *NullableS3AclPermission) Set(val *S3AclPermission) {
	v.value = val
	v.isSet = true
}

func (v NullableS3AclPermission) IsSet() bool {
	return v.isSet
}

func (v *NullableS3AclPermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3AclPermission(val *S3AclPermission) *NullableS3AclPermission {
	return &NullableS3AclPermission{value: val, isSet: true}
}

func (v NullableS3AclPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3AclPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o S3AclPermission) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}