/*
 * Cohesity REST API
 *
 * This API list provides operations for interfacing with the Cohesity Cluster.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cohesitysdk

import (
	"encoding/json"
)

// DocError struct for DocError
type DocError struct {
	// DocumentId is document which caused the error.
	DocumentId NullableString `json:"documentId,omitempty"`
	// ErrorString is the error converted to string.
	ErrorString NullableString `json:"errorString,omitempty"`
}

// NewDocError instantiates a new DocError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocError() *DocError {
	this := DocError{}
	return &this
}

// NewDocErrorWithDefaults instantiates a new DocError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocErrorWithDefaults() *DocError {
	this := DocError{}
	return &this
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocError) GetDocumentId() string {
	if o == nil || o.DocumentId.Get() == nil {
		var ret string
		return ret
	}
	return *o.DocumentId.Get()
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocError) GetDocumentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DocumentId.Get(), o.DocumentId.IsSet()
}

// HasDocumentId returns a boolean if a field has been set.
func (o *DocError) HasDocumentId() bool {
	if o != nil && o.DocumentId.IsSet() {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given NullableString and assigns it to the DocumentId field.
func (o *DocError) SetDocumentId(v string) {
	o.DocumentId.Set(&v)
}
// SetDocumentIdNil sets the value for DocumentId to be an explicit nil
func (o *DocError) SetDocumentIdNil() {
	o.DocumentId.Set(nil)
}

// UnsetDocumentId ensures that no value is present for DocumentId, not even an explicit nil
func (o *DocError) UnsetDocumentId() {
	o.DocumentId.Unset()
}

// GetErrorString returns the ErrorString field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocError) GetErrorString() string {
	if o == nil || o.ErrorString.Get() == nil {
		var ret string
		return ret
	}
	return *o.ErrorString.Get()
}

// GetErrorStringOk returns a tuple with the ErrorString field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocError) GetErrorStringOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ErrorString.Get(), o.ErrorString.IsSet()
}

// HasErrorString returns a boolean if a field has been set.
func (o *DocError) HasErrorString() bool {
	if o != nil && o.ErrorString.IsSet() {
		return true
	}

	return false
}

// SetErrorString gets a reference to the given NullableString and assigns it to the ErrorString field.
func (o *DocError) SetErrorString(v string) {
	o.ErrorString.Set(&v)
}
// SetErrorStringNil sets the value for ErrorString to be an explicit nil
func (o *DocError) SetErrorStringNil() {
	o.ErrorString.Set(nil)
}

// UnsetErrorString ensures that no value is present for ErrorString, not even an explicit nil
func (o *DocError) UnsetErrorString() {
	o.ErrorString.Unset()
}

func (o DocError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DocumentId.IsSet() {
		toSerialize["documentId"] = o.DocumentId.Get()
	}
	if o.ErrorString.IsSet() {
		toSerialize["errorString"] = o.ErrorString.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableDocError struct {
	value *DocError
	isSet bool
}

func (v NullableDocError) Get() *DocError {
	return v.value
}

func (v *NullableDocError) Set(val *DocError) {
	v.value = val
	v.isSet = true
}

func (v NullableDocError) IsSet() bool {
	return v.isSet
}

func (v *NullableDocError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocError(val *DocError) *NullableDocError {
	return &NullableDocError{value: val, isSet: true}
}

func (v NullableDocError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


