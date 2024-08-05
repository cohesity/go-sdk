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

// checks if the ErrorType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorType{}

// ErrorType Specifies type of error for faulty dbs.
type ErrorType struct {
	// Specifies type of error for faulty dbs.
	ErrorType *string `json:"errorType,omitempty"`
}

// NewErrorType instantiates a new ErrorType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorType() *ErrorType {
	this := ErrorType{}
	return &this
}

// NewErrorTypeWithDefaults instantiates a new ErrorType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorTypeWithDefaults() *ErrorType {
	this := ErrorType{}
	return &this
}

// GetErrorType returns the ErrorType field value if set, zero value otherwise.
func (o *ErrorType) GetErrorType() string {
	if o == nil || IsNil(o.ErrorType) {
		var ret string
		return ret
	}
	return *o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorType) GetErrorTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorType) {
		return nil, false
	}
	return o.ErrorType, true
}

// HasErrorType returns a boolean if a field has been set.
func (o *ErrorType) HasErrorType() bool {
	if o != nil && !IsNil(o.ErrorType) {
		return true
	}

	return false
}

// SetErrorType gets a reference to the given string and assigns it to the ErrorType field.
func (o *ErrorType) SetErrorType(v string) {
	o.ErrorType = &v
}

func (o ErrorType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorType) {
		toSerialize["errorType"] = o.ErrorType
	}
	return toSerialize, nil
}

type NullableErrorType struct {
	value *ErrorType
	isSet bool
}

func (v NullableErrorType) Get() *ErrorType {
	return v.value
}

func (v *NullableErrorType) Set(val *ErrorType) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorType) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorType(val *ErrorType) *NullableErrorType {
	return &NullableErrorType{value: val, isSet: true}
}

func (v NullableErrorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


