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

// KeyValuePair Specifies a map structure to store Key and Values.
type KeyValuePair struct {
	// key
	Key NullableString `json:"key"`
	// value
	Value NullableString `json:"value"`
}

// NewKeyValuePair instantiates a new KeyValuePair object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyValuePair(key NullableString, value NullableString) *KeyValuePair {
	this := KeyValuePair{}
	this.Key = key
	this.Value = value
	return &this
}

// NewKeyValuePairWithDefaults instantiates a new KeyValuePair object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyValuePairWithDefaults() *KeyValuePair {
	this := KeyValuePair{}
	return &this
}

// GetKey returns the Key field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KeyValuePair) GetKey() string {
	if o == nil || o.Key.Get() == nil {
		var ret string
		return ret
	}

	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyValuePair) GetKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// SetKey sets field value
func (o *KeyValuePair) SetKey(v string) {
	o.Key.Set(&v)
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KeyValuePair) GetValue() string {
	if o == nil || o.Value.Get() == nil {
		var ret string
		return ret
	}

	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyValuePair) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// SetValue sets field value
func (o *KeyValuePair) SetValue(v string) {
	o.Value.Set(&v)
}

func (o KeyValuePair) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key.Get()
	}
	if true {
		toSerialize["value"] = o.Value.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableKeyValuePair struct {
	value *KeyValuePair
	isSet bool
}

func (v NullableKeyValuePair) Get() *KeyValuePair {
	return v.value
}

func (v *NullableKeyValuePair) Set(val *KeyValuePair) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyValuePair) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyValuePair) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyValuePair(val *KeyValuePair) *NullableKeyValuePair {
	return &NullableKeyValuePair{value: val, isSet: true}
}

func (v NullableKeyValuePair) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyValuePair) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o KeyValuePair) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}