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

// StoragePolicyOverride Specifies if inline deduplication and compression settings inherited from Storage Domain (View Box) should be disabled for this View.
type StoragePolicyOverride struct {
	// If false, the inline deduplication and compression settings inherited from the Storage Domain (View Box) apply to this View. If true, both inline deduplication and compression are disabled for this View. This can only be set to true if inline deduplication is set for the Storage Domain (View Box).
	DisableInlineDedupAndCompression NullableBool `json:"disableInlineDedupAndCompression,omitempty"`
	// If it is set to true, we would disable dedup for writes made in this view irrespective of the view box's storage policy.
	DisableDedup NullableBool `json:"disableDedup,omitempty"`
}

// NewStoragePolicyOverride instantiates a new StoragePolicyOverride object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoragePolicyOverride() *StoragePolicyOverride {
	this := StoragePolicyOverride{}
	return &this
}

// NewStoragePolicyOverrideWithDefaults instantiates a new StoragePolicyOverride object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoragePolicyOverrideWithDefaults() *StoragePolicyOverride {
	this := StoragePolicyOverride{}
	return &this
}

// GetDisableInlineDedupAndCompression returns the DisableInlineDedupAndCompression field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePolicyOverride) GetDisableInlineDedupAndCompression() bool {
	if o == nil || o.DisableInlineDedupAndCompression.Get() == nil {
		var ret bool
		return ret
	}
	return *o.DisableInlineDedupAndCompression.Get()
}

// GetDisableInlineDedupAndCompressionOk returns a tuple with the DisableInlineDedupAndCompression field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePolicyOverride) GetDisableInlineDedupAndCompressionOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisableInlineDedupAndCompression.Get(), o.DisableInlineDedupAndCompression.IsSet()
}

// HasDisableInlineDedupAndCompression returns a boolean if a field has been set.
func (o *StoragePolicyOverride) HasDisableInlineDedupAndCompression() bool {
	if o != nil && o.DisableInlineDedupAndCompression.IsSet() {
		return true
	}

	return false
}

// SetDisableInlineDedupAndCompression gets a reference to the given NullableBool and assigns it to the DisableInlineDedupAndCompression field.
func (o *StoragePolicyOverride) SetDisableInlineDedupAndCompression(v bool) {
	o.DisableInlineDedupAndCompression.Set(&v)
}
// SetDisableInlineDedupAndCompressionNil sets the value for DisableInlineDedupAndCompression to be an explicit nil
func (o *StoragePolicyOverride) SetDisableInlineDedupAndCompressionNil() {
	o.DisableInlineDedupAndCompression.Set(nil)
}

// UnsetDisableInlineDedupAndCompression ensures that no value is present for DisableInlineDedupAndCompression, not even an explicit nil
func (o *StoragePolicyOverride) UnsetDisableInlineDedupAndCompression() {
	o.DisableInlineDedupAndCompression.Unset()
}

// GetDisableDedup returns the DisableDedup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePolicyOverride) GetDisableDedup() bool {
	if o == nil || o.DisableDedup.Get() == nil {
		var ret bool
		return ret
	}
	return *o.DisableDedup.Get()
}

// GetDisableDedupOk returns a tuple with the DisableDedup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePolicyOverride) GetDisableDedupOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisableDedup.Get(), o.DisableDedup.IsSet()
}

// HasDisableDedup returns a boolean if a field has been set.
func (o *StoragePolicyOverride) HasDisableDedup() bool {
	if o != nil && o.DisableDedup.IsSet() {
		return true
	}

	return false
}

// SetDisableDedup gets a reference to the given NullableBool and assigns it to the DisableDedup field.
func (o *StoragePolicyOverride) SetDisableDedup(v bool) {
	o.DisableDedup.Set(&v)
}
// SetDisableDedupNil sets the value for DisableDedup to be an explicit nil
func (o *StoragePolicyOverride) SetDisableDedupNil() {
	o.DisableDedup.Set(nil)
}

// UnsetDisableDedup ensures that no value is present for DisableDedup, not even an explicit nil
func (o *StoragePolicyOverride) UnsetDisableDedup() {
	o.DisableDedup.Unset()
}

func (o StoragePolicyOverride) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisableInlineDedupAndCompression.IsSet() {
		toSerialize["disableInlineDedupAndCompression"] = o.DisableInlineDedupAndCompression.Get()
	}
	if o.DisableDedup.IsSet() {
		toSerialize["disableDedup"] = o.DisableDedup.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableStoragePolicyOverride struct {
	value *StoragePolicyOverride
	isSet bool
}

func (v NullableStoragePolicyOverride) Get() *StoragePolicyOverride {
	return v.value
}

func (v *NullableStoragePolicyOverride) Set(val *StoragePolicyOverride) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePolicyOverride) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePolicyOverride) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePolicyOverride(val *StoragePolicyOverride) *NullableStoragePolicyOverride {
	return &NullableStoragePolicyOverride{value: val, isSet: true}
}

func (v NullableStoragePolicyOverride) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePolicyOverride) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o StoragePolicyOverride) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}