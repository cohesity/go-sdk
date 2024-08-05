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

// checks if the AuthHeader type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthHeader{}

// AuthHeader Specifies structure of request header
type AuthHeader struct {
	// Specifies key for the header.
	Key *string `json:"key,omitempty"`
	// Specifies value for the header.
	Value *string `json:"value,omitempty"`
}

// NewAuthHeader instantiates a new AuthHeader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthHeader() *AuthHeader {
	this := AuthHeader{}
	return &this
}

// NewAuthHeaderWithDefaults instantiates a new AuthHeader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthHeaderWithDefaults() *AuthHeader {
	this := AuthHeader{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *AuthHeader) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthHeader) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *AuthHeader) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *AuthHeader) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AuthHeader) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthHeader) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AuthHeader) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *AuthHeader) SetValue(v string) {
	o.Value = &v
}

func (o AuthHeader) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthHeader) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableAuthHeader struct {
	value *AuthHeader
	isSet bool
}

func (v NullableAuthHeader) Get() *AuthHeader {
	return v.value
}

func (v *NullableAuthHeader) Set(val *AuthHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthHeader(val *AuthHeader) *NullableAuthHeader {
	return &NullableAuthHeader{value: val, isSet: true}
}

func (v NullableAuthHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


