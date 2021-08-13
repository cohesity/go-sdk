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

// RetryOptions Retry Options of a Protection Policy when a Protection Group run fails.
type RetryOptions struct {
	// Specifies the number of times to retry capturing Snapshots before the Protection Group Run fails.
	Retries NullableInt32 `json:"retries,omitempty"`
	// Specifies the number of minutes before retrying a failed Protection Group.
	RetryIntervalMins NullableInt32 `json:"retryIntervalMins,omitempty"`
}

// NewRetryOptions instantiates a new RetryOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetryOptions() *RetryOptions {
	this := RetryOptions{}
	return &this
}

// NewRetryOptionsWithDefaults instantiates a new RetryOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetryOptionsWithDefaults() *RetryOptions {
	this := RetryOptions{}
	return &this
}

// GetRetries returns the Retries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RetryOptions) GetRetries() int32 {
	if o == nil || o.Retries.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Retries.Get()
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RetryOptions) GetRetriesOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Retries.Get(), o.Retries.IsSet()
}

// HasRetries returns a boolean if a field has been set.
func (o *RetryOptions) HasRetries() bool {
	if o != nil && o.Retries.IsSet() {
		return true
	}

	return false
}

// SetRetries gets a reference to the given NullableInt32 and assigns it to the Retries field.
func (o *RetryOptions) SetRetries(v int32) {
	o.Retries.Set(&v)
}
// SetRetriesNil sets the value for Retries to be an explicit nil
func (o *RetryOptions) SetRetriesNil() {
	o.Retries.Set(nil)
}

// UnsetRetries ensures that no value is present for Retries, not even an explicit nil
func (o *RetryOptions) UnsetRetries() {
	o.Retries.Unset()
}

// GetRetryIntervalMins returns the RetryIntervalMins field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RetryOptions) GetRetryIntervalMins() int32 {
	if o == nil || o.RetryIntervalMins.Get() == nil {
		var ret int32
		return ret
	}
	return *o.RetryIntervalMins.Get()
}

// GetRetryIntervalMinsOk returns a tuple with the RetryIntervalMins field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RetryOptions) GetRetryIntervalMinsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RetryIntervalMins.Get(), o.RetryIntervalMins.IsSet()
}

// HasRetryIntervalMins returns a boolean if a field has been set.
func (o *RetryOptions) HasRetryIntervalMins() bool {
	if o != nil && o.RetryIntervalMins.IsSet() {
		return true
	}

	return false
}

// SetRetryIntervalMins gets a reference to the given NullableInt32 and assigns it to the RetryIntervalMins field.
func (o *RetryOptions) SetRetryIntervalMins(v int32) {
	o.RetryIntervalMins.Set(&v)
}
// SetRetryIntervalMinsNil sets the value for RetryIntervalMins to be an explicit nil
func (o *RetryOptions) SetRetryIntervalMinsNil() {
	o.RetryIntervalMins.Set(nil)
}

// UnsetRetryIntervalMins ensures that no value is present for RetryIntervalMins, not even an explicit nil
func (o *RetryOptions) UnsetRetryIntervalMins() {
	o.RetryIntervalMins.Unset()
}

func (o RetryOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Retries.IsSet() {
		toSerialize["retries"] = o.Retries.Get()
	}
	if o.RetryIntervalMins.IsSet() {
		toSerialize["retryIntervalMins"] = o.RetryIntervalMins.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRetryOptions struct {
	value *RetryOptions
	isSet bool
}

func (v NullableRetryOptions) Get() *RetryOptions {
	return v.value
}

func (v *NullableRetryOptions) Set(val *RetryOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableRetryOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableRetryOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetryOptions(val *RetryOptions) *NullableRetryOptions {
	return &NullableRetryOptions{value: val, isSet: true}
}

func (v NullableRetryOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetryOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


