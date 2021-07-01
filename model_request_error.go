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

// RequestError Details about the Error.
type RequestError struct {
	// Operation response error code.
	ErrorCode NullableInt64 `json:"errorCode,omitempty"`
	// Description of the error.
	Message NullableString `json:"message,omitempty"`
}

// NewRequestError instantiates a new RequestError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestError() *RequestError {
	this := RequestError{}
	return &this
}

// NewRequestErrorWithDefaults instantiates a new RequestError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestErrorWithDefaults() *RequestError {
	this := RequestError{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestError) GetErrorCode() int64 {
	if o == nil || o.ErrorCode.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ErrorCode.Get()
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestError) GetErrorCodeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ErrorCode.Get(), o.ErrorCode.IsSet()
}

// HasErrorCode returns a boolean if a field has been set.
func (o *RequestError) HasErrorCode() bool {
	if o != nil && o.ErrorCode.IsSet() {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given NullableInt64 and assigns it to the ErrorCode field.
func (o *RequestError) SetErrorCode(v int64) {
	o.ErrorCode.Set(&v)
}
// SetErrorCodeNil sets the value for ErrorCode to be an explicit nil
func (o *RequestError) SetErrorCodeNil() {
	o.ErrorCode.Set(nil)
}

// UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
func (o *RequestError) UnsetErrorCode() {
	o.ErrorCode.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestError) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestError) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *RequestError) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *RequestError) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *RequestError) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *RequestError) UnsetMessage() {
	o.Message.Unset()
}

func (o RequestError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorCode.IsSet() {
		toSerialize["errorCode"] = o.ErrorCode.Get()
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRequestError struct {
	value *RequestError
	isSet bool
}

func (v NullableRequestError) Get() *RequestError {
	return v.value
}

func (v *NullableRequestError) Set(val *RequestError) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestError) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestError(val *RequestError) *NullableRequestError {
	return &NullableRequestError{value: val, isSet: true}
}

func (v NullableRequestError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


