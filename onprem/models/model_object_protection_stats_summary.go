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

// ObjectProtectionStatsSummary Specifies the count and size of protected and unprotected objects for a given environment.
type ObjectProtectionStatsSummary struct {
	// Specifies the environment of the object.
	Environment NullableString `json:"environment,omitempty"`
	// Specifies the count of the protected leaf objects.
	ProtectedCount NullableInt64 `json:"protectedCount,omitempty"`
	// Specifies the count of the unprotected leaf objects.
	UnprotectedCount NullableInt64 `json:"unprotectedCount,omitempty"`
	// Specifies the protected logical size in bytes.
	ProtectedSizeBytes NullableInt64 `json:"protectedSizeBytes,omitempty"`
	// Specifies the unprotected logical size in bytes.
	UnprotectedSizeBytes NullableInt64 `json:"unprotectedSizeBytes,omitempty"`
}

// NewObjectProtectionStatsSummary instantiates a new ObjectProtectionStatsSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectProtectionStatsSummary() *ObjectProtectionStatsSummary {
	this := ObjectProtectionStatsSummary{}
	return &this
}

// NewObjectProtectionStatsSummaryWithDefaults instantiates a new ObjectProtectionStatsSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectProtectionStatsSummaryWithDefaults() *ObjectProtectionStatsSummary {
	this := ObjectProtectionStatsSummary{}
	return &this
}

// GetEnvironment returns the Environment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObjectProtectionStatsSummary) GetEnvironment() string {
	if o == nil || o.Environment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Environment.Get()
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectProtectionStatsSummary) GetEnvironmentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Environment.Get(), o.Environment.IsSet()
}

// HasEnvironment returns a boolean if a field has been set.
func (o *ObjectProtectionStatsSummary) HasEnvironment() bool {
	if o != nil && o.Environment.IsSet() {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given NullableString and assigns it to the Environment field.
func (o *ObjectProtectionStatsSummary) SetEnvironment(v string) {
	o.Environment.Set(&v)
}
// SetEnvironmentNil sets the value for Environment to be an explicit nil
func (o *ObjectProtectionStatsSummary) SetEnvironmentNil() {
	o.Environment.Set(nil)
}

// UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
func (o *ObjectProtectionStatsSummary) UnsetEnvironment() {
	o.Environment.Unset()
}

// GetProtectedCount returns the ProtectedCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObjectProtectionStatsSummary) GetProtectedCount() int64 {
	if o == nil || o.ProtectedCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ProtectedCount.Get()
}

// GetProtectedCountOk returns a tuple with the ProtectedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectProtectionStatsSummary) GetProtectedCountOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProtectedCount.Get(), o.ProtectedCount.IsSet()
}

// HasProtectedCount returns a boolean if a field has been set.
func (o *ObjectProtectionStatsSummary) HasProtectedCount() bool {
	if o != nil && o.ProtectedCount.IsSet() {
		return true
	}

	return false
}

// SetProtectedCount gets a reference to the given NullableInt64 and assigns it to the ProtectedCount field.
func (o *ObjectProtectionStatsSummary) SetProtectedCount(v int64) {
	o.ProtectedCount.Set(&v)
}
// SetProtectedCountNil sets the value for ProtectedCount to be an explicit nil
func (o *ObjectProtectionStatsSummary) SetProtectedCountNil() {
	o.ProtectedCount.Set(nil)
}

// UnsetProtectedCount ensures that no value is present for ProtectedCount, not even an explicit nil
func (o *ObjectProtectionStatsSummary) UnsetProtectedCount() {
	o.ProtectedCount.Unset()
}

// GetUnprotectedCount returns the UnprotectedCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObjectProtectionStatsSummary) GetUnprotectedCount() int64 {
	if o == nil || o.UnprotectedCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.UnprotectedCount.Get()
}

// GetUnprotectedCountOk returns a tuple with the UnprotectedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectProtectionStatsSummary) GetUnprotectedCountOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnprotectedCount.Get(), o.UnprotectedCount.IsSet()
}

// HasUnprotectedCount returns a boolean if a field has been set.
func (o *ObjectProtectionStatsSummary) HasUnprotectedCount() bool {
	if o != nil && o.UnprotectedCount.IsSet() {
		return true
	}

	return false
}

// SetUnprotectedCount gets a reference to the given NullableInt64 and assigns it to the UnprotectedCount field.
func (o *ObjectProtectionStatsSummary) SetUnprotectedCount(v int64) {
	o.UnprotectedCount.Set(&v)
}
// SetUnprotectedCountNil sets the value for UnprotectedCount to be an explicit nil
func (o *ObjectProtectionStatsSummary) SetUnprotectedCountNil() {
	o.UnprotectedCount.Set(nil)
}

// UnsetUnprotectedCount ensures that no value is present for UnprotectedCount, not even an explicit nil
func (o *ObjectProtectionStatsSummary) UnsetUnprotectedCount() {
	o.UnprotectedCount.Unset()
}

// GetProtectedSizeBytes returns the ProtectedSizeBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObjectProtectionStatsSummary) GetProtectedSizeBytes() int64 {
	if o == nil || o.ProtectedSizeBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ProtectedSizeBytes.Get()
}

// GetProtectedSizeBytesOk returns a tuple with the ProtectedSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectProtectionStatsSummary) GetProtectedSizeBytesOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProtectedSizeBytes.Get(), o.ProtectedSizeBytes.IsSet()
}

// HasProtectedSizeBytes returns a boolean if a field has been set.
func (o *ObjectProtectionStatsSummary) HasProtectedSizeBytes() bool {
	if o != nil && o.ProtectedSizeBytes.IsSet() {
		return true
	}

	return false
}

// SetProtectedSizeBytes gets a reference to the given NullableInt64 and assigns it to the ProtectedSizeBytes field.
func (o *ObjectProtectionStatsSummary) SetProtectedSizeBytes(v int64) {
	o.ProtectedSizeBytes.Set(&v)
}
// SetProtectedSizeBytesNil sets the value for ProtectedSizeBytes to be an explicit nil
func (o *ObjectProtectionStatsSummary) SetProtectedSizeBytesNil() {
	o.ProtectedSizeBytes.Set(nil)
}

// UnsetProtectedSizeBytes ensures that no value is present for ProtectedSizeBytes, not even an explicit nil
func (o *ObjectProtectionStatsSummary) UnsetProtectedSizeBytes() {
	o.ProtectedSizeBytes.Unset()
}

// GetUnprotectedSizeBytes returns the UnprotectedSizeBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObjectProtectionStatsSummary) GetUnprotectedSizeBytes() int64 {
	if o == nil || o.UnprotectedSizeBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.UnprotectedSizeBytes.Get()
}

// GetUnprotectedSizeBytesOk returns a tuple with the UnprotectedSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectProtectionStatsSummary) GetUnprotectedSizeBytesOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnprotectedSizeBytes.Get(), o.UnprotectedSizeBytes.IsSet()
}

// HasUnprotectedSizeBytes returns a boolean if a field has been set.
func (o *ObjectProtectionStatsSummary) HasUnprotectedSizeBytes() bool {
	if o != nil && o.UnprotectedSizeBytes.IsSet() {
		return true
	}

	return false
}

// SetUnprotectedSizeBytes gets a reference to the given NullableInt64 and assigns it to the UnprotectedSizeBytes field.
func (o *ObjectProtectionStatsSummary) SetUnprotectedSizeBytes(v int64) {
	o.UnprotectedSizeBytes.Set(&v)
}
// SetUnprotectedSizeBytesNil sets the value for UnprotectedSizeBytes to be an explicit nil
func (o *ObjectProtectionStatsSummary) SetUnprotectedSizeBytesNil() {
	o.UnprotectedSizeBytes.Set(nil)
}

// UnsetUnprotectedSizeBytes ensures that no value is present for UnprotectedSizeBytes, not even an explicit nil
func (o *ObjectProtectionStatsSummary) UnsetUnprotectedSizeBytes() {
	o.UnprotectedSizeBytes.Unset()
}

func (o ObjectProtectionStatsSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Environment.IsSet() {
		toSerialize["environment"] = o.Environment.Get()
	}
	if o.ProtectedCount.IsSet() {
		toSerialize["protectedCount"] = o.ProtectedCount.Get()
	}
	if o.UnprotectedCount.IsSet() {
		toSerialize["unprotectedCount"] = o.UnprotectedCount.Get()
	}
	if o.ProtectedSizeBytes.IsSet() {
		toSerialize["protectedSizeBytes"] = o.ProtectedSizeBytes.Get()
	}
	if o.UnprotectedSizeBytes.IsSet() {
		toSerialize["unprotectedSizeBytes"] = o.UnprotectedSizeBytes.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableObjectProtectionStatsSummary struct {
	value *ObjectProtectionStatsSummary
	isSet bool
}

func (v NullableObjectProtectionStatsSummary) Get() *ObjectProtectionStatsSummary {
	return v.value
}

func (v *NullableObjectProtectionStatsSummary) Set(val *ObjectProtectionStatsSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectProtectionStatsSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectProtectionStatsSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectProtectionStatsSummary(val *ObjectProtectionStatsSummary) *NullableObjectProtectionStatsSummary {
	return &NullableObjectProtectionStatsSummary{value: val, isSet: true}
}

func (v NullableObjectProtectionStatsSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectProtectionStatsSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o ObjectProtectionStatsSummary) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}