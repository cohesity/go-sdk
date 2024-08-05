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

// checks if the AwsProtectionGroupType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsProtectionGroupType{}

// AwsProtectionGroupType AWS Protection Group type.
type AwsProtectionGroupType struct {
	// Specifies AWS Protection Group type.
	Environment *string `json:"environment,omitempty"`
}

// NewAwsProtectionGroupType instantiates a new AwsProtectionGroupType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsProtectionGroupType() *AwsProtectionGroupType {
	this := AwsProtectionGroupType{}
	return &this
}

// NewAwsProtectionGroupTypeWithDefaults instantiates a new AwsProtectionGroupType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsProtectionGroupTypeWithDefaults() *AwsProtectionGroupType {
	this := AwsProtectionGroupType{}
	return &this
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *AwsProtectionGroupType) GetEnvironment() string {
	if o == nil || IsNil(o.Environment) {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsProtectionGroupType) GetEnvironmentOk() (*string, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *AwsProtectionGroupType) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *AwsProtectionGroupType) SetEnvironment(v string) {
	o.Environment = &v
}

func (o AwsProtectionGroupType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsProtectionGroupType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	return toSerialize, nil
}

type NullableAwsProtectionGroupType struct {
	value *AwsProtectionGroupType
	isSet bool
}

func (v NullableAwsProtectionGroupType) Get() *AwsProtectionGroupType {
	return v.value
}

func (v *NullableAwsProtectionGroupType) Set(val *AwsProtectionGroupType) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsProtectionGroupType) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsProtectionGroupType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsProtectionGroupType(val *AwsProtectionGroupType) *NullableAwsProtectionGroupType {
	return &NullableAwsProtectionGroupType{value: val, isSet: true}
}

func (v NullableAwsProtectionGroupType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsProtectionGroupType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


