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

// NasFullThrottlingConfig Specifies the throttling configuration during full backup run.
type NasFullThrottlingConfig struct {
	// Specifies the percentage value of maximum concurrent metadata to be fetched during full backup of the source.
	MaxMetadataFetchPercentage NullableInt32 `json:"maxMetadataFetchPercentage,omitempty"`
	// Specifies the percentage value of maximum concurrent read/write during full backup of the source.
	MaxReadWritePercentage NullableInt32 `json:"maxReadWritePercentage,omitempty"`
}

// NewNasFullThrottlingConfig instantiates a new NasFullThrottlingConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNasFullThrottlingConfig() *NasFullThrottlingConfig {
	this := NasFullThrottlingConfig{}
	return &this
}

// NewNasFullThrottlingConfigWithDefaults instantiates a new NasFullThrottlingConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNasFullThrottlingConfigWithDefaults() *NasFullThrottlingConfig {
	this := NasFullThrottlingConfig{}
	return &this
}

// GetMaxMetadataFetchPercentage returns the MaxMetadataFetchPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NasFullThrottlingConfig) GetMaxMetadataFetchPercentage() int32 {
	if o == nil || o.MaxMetadataFetchPercentage.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MaxMetadataFetchPercentage.Get()
}

// GetMaxMetadataFetchPercentageOk returns a tuple with the MaxMetadataFetchPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NasFullThrottlingConfig) GetMaxMetadataFetchPercentageOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MaxMetadataFetchPercentage.Get(), o.MaxMetadataFetchPercentage.IsSet()
}

// HasMaxMetadataFetchPercentage returns a boolean if a field has been set.
func (o *NasFullThrottlingConfig) HasMaxMetadataFetchPercentage() bool {
	if o != nil && o.MaxMetadataFetchPercentage.IsSet() {
		return true
	}

	return false
}

// SetMaxMetadataFetchPercentage gets a reference to the given NullableInt32 and assigns it to the MaxMetadataFetchPercentage field.
func (o *NasFullThrottlingConfig) SetMaxMetadataFetchPercentage(v int32) {
	o.MaxMetadataFetchPercentage.Set(&v)
}
// SetMaxMetadataFetchPercentageNil sets the value for MaxMetadataFetchPercentage to be an explicit nil
func (o *NasFullThrottlingConfig) SetMaxMetadataFetchPercentageNil() {
	o.MaxMetadataFetchPercentage.Set(nil)
}

// UnsetMaxMetadataFetchPercentage ensures that no value is present for MaxMetadataFetchPercentage, not even an explicit nil
func (o *NasFullThrottlingConfig) UnsetMaxMetadataFetchPercentage() {
	o.MaxMetadataFetchPercentage.Unset()
}

// GetMaxReadWritePercentage returns the MaxReadWritePercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NasFullThrottlingConfig) GetMaxReadWritePercentage() int32 {
	if o == nil || o.MaxReadWritePercentage.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MaxReadWritePercentage.Get()
}

// GetMaxReadWritePercentageOk returns a tuple with the MaxReadWritePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NasFullThrottlingConfig) GetMaxReadWritePercentageOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MaxReadWritePercentage.Get(), o.MaxReadWritePercentage.IsSet()
}

// HasMaxReadWritePercentage returns a boolean if a field has been set.
func (o *NasFullThrottlingConfig) HasMaxReadWritePercentage() bool {
	if o != nil && o.MaxReadWritePercentage.IsSet() {
		return true
	}

	return false
}

// SetMaxReadWritePercentage gets a reference to the given NullableInt32 and assigns it to the MaxReadWritePercentage field.
func (o *NasFullThrottlingConfig) SetMaxReadWritePercentage(v int32) {
	o.MaxReadWritePercentage.Set(&v)
}
// SetMaxReadWritePercentageNil sets the value for MaxReadWritePercentage to be an explicit nil
func (o *NasFullThrottlingConfig) SetMaxReadWritePercentageNil() {
	o.MaxReadWritePercentage.Set(nil)
}

// UnsetMaxReadWritePercentage ensures that no value is present for MaxReadWritePercentage, not even an explicit nil
func (o *NasFullThrottlingConfig) UnsetMaxReadWritePercentage() {
	o.MaxReadWritePercentage.Unset()
}

func (o NasFullThrottlingConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MaxMetadataFetchPercentage.IsSet() {
		toSerialize["maxMetadataFetchPercentage"] = o.MaxMetadataFetchPercentage.Get()
	}
	if o.MaxReadWritePercentage.IsSet() {
		toSerialize["maxReadWritePercentage"] = o.MaxReadWritePercentage.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableNasFullThrottlingConfig struct {
	value *NasFullThrottlingConfig
	isSet bool
}

func (v NullableNasFullThrottlingConfig) Get() *NasFullThrottlingConfig {
	return v.value
}

func (v *NullableNasFullThrottlingConfig) Set(val *NasFullThrottlingConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableNasFullThrottlingConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableNasFullThrottlingConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNasFullThrottlingConfig(val *NasFullThrottlingConfig) *NullableNasFullThrottlingConfig {
	return &NullableNasFullThrottlingConfig{value: val, isSet: true}
}

func (v NullableNasFullThrottlingConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNasFullThrottlingConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o NasFullThrottlingConfig) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}