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

// ObjectProtectionSummaryAllOf struct for ObjectProtectionSummaryAllOf
type ObjectProtectionSummaryAllOf struct {
	// Specifies the type of error which occurred during creation of the object protection.
	ErrorType NullableString `json:"errorType,omitempty"`
	// Specifies the error message if an error occurred during creation of the object protection.
	Error NullableString `json:"error,omitempty"`
}

// NewObjectProtectionSummaryAllOf instantiates a new ObjectProtectionSummaryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectProtectionSummaryAllOf() *ObjectProtectionSummaryAllOf {
	this := ObjectProtectionSummaryAllOf{}
	return &this
}

// NewObjectProtectionSummaryAllOfWithDefaults instantiates a new ObjectProtectionSummaryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectProtectionSummaryAllOfWithDefaults() *ObjectProtectionSummaryAllOf {
	this := ObjectProtectionSummaryAllOf{}
	return &this
}

// GetErrorType returns the ErrorType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObjectProtectionSummaryAllOf) GetErrorType() string {
	if o == nil || o.ErrorType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ErrorType.Get()
}

// GetErrorTypeOk returns a tuple with the ErrorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectProtectionSummaryAllOf) GetErrorTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ErrorType.Get(), o.ErrorType.IsSet()
}

// HasErrorType returns a boolean if a field has been set.
func (o *ObjectProtectionSummaryAllOf) HasErrorType() bool {
	if o != nil && o.ErrorType.IsSet() {
		return true
	}

	return false
}

// SetErrorType gets a reference to the given NullableString and assigns it to the ErrorType field.
func (o *ObjectProtectionSummaryAllOf) SetErrorType(v string) {
	o.ErrorType.Set(&v)
}
// SetErrorTypeNil sets the value for ErrorType to be an explicit nil
func (o *ObjectProtectionSummaryAllOf) SetErrorTypeNil() {
	o.ErrorType.Set(nil)
}

// UnsetErrorType ensures that no value is present for ErrorType, not even an explicit nil
func (o *ObjectProtectionSummaryAllOf) UnsetErrorType() {
	o.ErrorType.Unset()
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObjectProtectionSummaryAllOf) GetError() string {
	if o == nil || o.Error.Get() == nil {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectProtectionSummaryAllOf) GetErrorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *ObjectProtectionSummaryAllOf) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *ObjectProtectionSummaryAllOf) SetError(v string) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *ObjectProtectionSummaryAllOf) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *ObjectProtectionSummaryAllOf) UnsetError() {
	o.Error.Unset()
}

func (o ObjectProtectionSummaryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorType.IsSet() {
		toSerialize["errorType"] = o.ErrorType.Get()
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableObjectProtectionSummaryAllOf struct {
	value *ObjectProtectionSummaryAllOf
	isSet bool
}

func (v NullableObjectProtectionSummaryAllOf) Get() *ObjectProtectionSummaryAllOf {
	return v.value
}

func (v *NullableObjectProtectionSummaryAllOf) Set(val *ObjectProtectionSummaryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectProtectionSummaryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectProtectionSummaryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectProtectionSummaryAllOf(val *ObjectProtectionSummaryAllOf) *NullableObjectProtectionSummaryAllOf {
	return &NullableObjectProtectionSummaryAllOf{value: val, isSet: true}
}

func (v NullableObjectProtectionSummaryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectProtectionSummaryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


