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

// CloudSpinDataStats Specifies statistics about Cloud Spin data.
type CloudSpinDataStats struct {
	// Specifies the physical bytes transferred.
	PhysicalBytesTransferred NullableInt64 `json:"physicalBytesTransferred,omitempty"`
}

// NewCloudSpinDataStats instantiates a new CloudSpinDataStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudSpinDataStats() *CloudSpinDataStats {
	this := CloudSpinDataStats{}
	return &this
}

// NewCloudSpinDataStatsWithDefaults instantiates a new CloudSpinDataStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudSpinDataStatsWithDefaults() *CloudSpinDataStats {
	this := CloudSpinDataStats{}
	return &this
}

// GetPhysicalBytesTransferred returns the PhysicalBytesTransferred field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudSpinDataStats) GetPhysicalBytesTransferred() int64 {
	if o == nil || o.PhysicalBytesTransferred.Get() == nil {
		var ret int64
		return ret
	}
	return *o.PhysicalBytesTransferred.Get()
}

// GetPhysicalBytesTransferredOk returns a tuple with the PhysicalBytesTransferred field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudSpinDataStats) GetPhysicalBytesTransferredOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PhysicalBytesTransferred.Get(), o.PhysicalBytesTransferred.IsSet()
}

// HasPhysicalBytesTransferred returns a boolean if a field has been set.
func (o *CloudSpinDataStats) HasPhysicalBytesTransferred() bool {
	if o != nil && o.PhysicalBytesTransferred.IsSet() {
		return true
	}

	return false
}

// SetPhysicalBytesTransferred gets a reference to the given NullableInt64 and assigns it to the PhysicalBytesTransferred field.
func (o *CloudSpinDataStats) SetPhysicalBytesTransferred(v int64) {
	o.PhysicalBytesTransferred.Set(&v)
}
// SetPhysicalBytesTransferredNil sets the value for PhysicalBytesTransferred to be an explicit nil
func (o *CloudSpinDataStats) SetPhysicalBytesTransferredNil() {
	o.PhysicalBytesTransferred.Set(nil)
}

// UnsetPhysicalBytesTransferred ensures that no value is present for PhysicalBytesTransferred, not even an explicit nil
func (o *CloudSpinDataStats) UnsetPhysicalBytesTransferred() {
	o.PhysicalBytesTransferred.Unset()
}

func (o CloudSpinDataStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PhysicalBytesTransferred.IsSet() {
		toSerialize["physicalBytesTransferred"] = o.PhysicalBytesTransferred.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCloudSpinDataStats struct {
	value *CloudSpinDataStats
	isSet bool
}

func (v NullableCloudSpinDataStats) Get() *CloudSpinDataStats {
	return v.value
}

func (v *NullableCloudSpinDataStats) Set(val *CloudSpinDataStats) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudSpinDataStats) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudSpinDataStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudSpinDataStats(val *CloudSpinDataStats) *NullableCloudSpinDataStats {
	return &NullableCloudSpinDataStats{value: val, isSet: true}
}

func (v NullableCloudSpinDataStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudSpinDataStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


