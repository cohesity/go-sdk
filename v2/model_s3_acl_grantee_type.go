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

// checks if the S3AclGranteeType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &S3AclGranteeType{}

// S3AclGranteeType Specifies S3 Grantee Type.
type S3AclGranteeType struct {
	// Specifies S3 Grantee Type.
	Enum *string `json:"enum,omitempty"`
}

// NewS3AclGranteeType instantiates a new S3AclGranteeType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3AclGranteeType() *S3AclGranteeType {
	this := S3AclGranteeType{}
	return &this
}

// NewS3AclGranteeTypeWithDefaults instantiates a new S3AclGranteeType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3AclGranteeTypeWithDefaults() *S3AclGranteeType {
	this := S3AclGranteeType{}
	return &this
}

// GetEnum returns the Enum field value if set, zero value otherwise.
func (o *S3AclGranteeType) GetEnum() string {
	if o == nil || IsNil(o.Enum) {
		var ret string
		return ret
	}
	return *o.Enum
}

// GetEnumOk returns a tuple with the Enum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3AclGranteeType) GetEnumOk() (*string, bool) {
	if o == nil || IsNil(o.Enum) {
		return nil, false
	}
	return o.Enum, true
}

// HasEnum returns a boolean if a field has been set.
func (o *S3AclGranteeType) HasEnum() bool {
	if o != nil && !IsNil(o.Enum) {
		return true
	}

	return false
}

// SetEnum gets a reference to the given string and assigns it to the Enum field.
func (o *S3AclGranteeType) SetEnum(v string) {
	o.Enum = &v
}

func (o S3AclGranteeType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o S3AclGranteeType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enum) {
		toSerialize["enum"] = o.Enum
	}
	return toSerialize, nil
}

type NullableS3AclGranteeType struct {
	value *S3AclGranteeType
	isSet bool
}

func (v NullableS3AclGranteeType) Get() *S3AclGranteeType {
	return v.value
}

func (v *NullableS3AclGranteeType) Set(val *S3AclGranteeType) {
	v.value = val
	v.isSet = true
}

func (v NullableS3AclGranteeType) IsSet() bool {
	return v.isSet
}

func (v *NullableS3AclGranteeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3AclGranteeType(val *S3AclGranteeType) *NullableS3AclGranteeType {
	return &NullableS3AclGranteeType{value: val, isSet: true}
}

func (v NullableS3AclGranteeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3AclGranteeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


