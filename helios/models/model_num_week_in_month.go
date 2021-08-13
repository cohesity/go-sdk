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

// NumWeekInMonth Num Week In Month type.
type NumWeekInMonth struct {
	// Specifies Num Week In Month type.
	NumWeekInMonth *string `json:"numWeekInMonth,omitempty"`
}

// NewNumWeekInMonth instantiates a new NumWeekInMonth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNumWeekInMonth() *NumWeekInMonth {
	this := NumWeekInMonth{}
	return &this
}

// NewNumWeekInMonthWithDefaults instantiates a new NumWeekInMonth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNumWeekInMonthWithDefaults() *NumWeekInMonth {
	this := NumWeekInMonth{}
	return &this
}

// GetNumWeekInMonth returns the NumWeekInMonth field value if set, zero value otherwise.
func (o *NumWeekInMonth) GetNumWeekInMonth() string {
	if o == nil || o.NumWeekInMonth == nil {
		var ret string
		return ret
	}
	return *o.NumWeekInMonth
}

// GetNumWeekInMonthOk returns a tuple with the NumWeekInMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumWeekInMonth) GetNumWeekInMonthOk() (*string, bool) {
	if o == nil || o.NumWeekInMonth == nil {
		return nil, false
	}
	return o.NumWeekInMonth, true
}

// HasNumWeekInMonth returns a boolean if a field has been set.
func (o *NumWeekInMonth) HasNumWeekInMonth() bool {
	if o != nil && o.NumWeekInMonth != nil {
		return true
	}

	return false
}

// SetNumWeekInMonth gets a reference to the given string and assigns it to the NumWeekInMonth field.
func (o *NumWeekInMonth) SetNumWeekInMonth(v string) {
	o.NumWeekInMonth = &v
}

func (o NumWeekInMonth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NumWeekInMonth != nil {
		toSerialize["numWeekInMonth"] = o.NumWeekInMonth
	}
	return json.Marshal(toSerialize)
}

type NullableNumWeekInMonth struct {
	value *NumWeekInMonth
	isSet bool
}

func (v NullableNumWeekInMonth) Get() *NumWeekInMonth {
	return v.value
}

func (v *NullableNumWeekInMonth) Set(val *NumWeekInMonth) {
	v.value = val
	v.isSet = true
}

func (v NullableNumWeekInMonth) IsSet() bool {
	return v.isSet
}

func (v *NullableNumWeekInMonth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNumWeekInMonth(val *NumWeekInMonth) *NullableNumWeekInMonth {
	return &NullableNumWeekInMonth{value: val, isSet: true}
}

func (v NullableNumWeekInMonth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNumWeekInMonth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


