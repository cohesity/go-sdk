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

// GoogleTieringLevels Google Tiering Levels
type GoogleTieringLevels struct {
	// Google Tiering Levels
	Type *string `json:"type,omitempty"`
}

// NewGoogleTieringLevels instantiates a new GoogleTieringLevels object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleTieringLevels() *GoogleTieringLevels {
	this := GoogleTieringLevels{}
	return &this
}

// NewGoogleTieringLevelsWithDefaults instantiates a new GoogleTieringLevels object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleTieringLevelsWithDefaults() *GoogleTieringLevels {
	this := GoogleTieringLevels{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GoogleTieringLevels) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleTieringLevels) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GoogleTieringLevels) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GoogleTieringLevels) SetType(v string) {
	o.Type = &v
}

func (o GoogleTieringLevels) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableGoogleTieringLevels struct {
	value *GoogleTieringLevels
	isSet bool
}

func (v NullableGoogleTieringLevels) Get() *GoogleTieringLevels {
	return v.value
}

func (v *NullableGoogleTieringLevels) Set(val *GoogleTieringLevels) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleTieringLevels) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleTieringLevels) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleTieringLevels(val *GoogleTieringLevels) *NullableGoogleTieringLevels {
	return &NullableGoogleTieringLevels{value: val, isSet: true}
}

func (v NullableGoogleTieringLevels) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleTieringLevels) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o GoogleTieringLevels) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}