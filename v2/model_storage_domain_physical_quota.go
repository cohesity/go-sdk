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

// checks if the StorageDomainPhysicalQuota type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageDomainPhysicalQuota{}

// StorageDomainPhysicalQuota Specifies a quota limit for physical usage of this Storage Domain. This quota defines a limit of data that can be physically (after data size is reduced by block tracking, compression and deduplication) stored on this storage domain. A new write will not be allowed when the storage domain usage will exceeds the specified quota. Due to the latency of calculating usage across all nodes, the actual storage domain usage may exceed the quota limit by a little bit.
type StorageDomainPhysicalQuota struct {
	// Specifies if an alert should be triggered when the usage of this resource exceeds this quota limit. This limit is optional and is specified in bytes. If no value is specified, there is no limit.
	AlertLimitBytes NullableInt64 `json:"alertLimitBytes,omitempty"`
	// Supported only for user quota policy. Specifies when the usage goes above an alert threshold percentage which is: HardLimitBytes * AlertThresholdPercentage, eg: 80% of HardLimitBytes Can only be set if HardLimitBytes is set. Cannot be set if AlertLimitBytes is already set.
	AlertThresholdPercentage NullableInt64 `json:"alertThresholdPercentage,omitempty"`
	// Specifies an optional quota limit on the usage allowed for this resource. This limit is specified in bytes. If no value is specified, there is no limit.
	HardLimitBytes NullableInt64 `json:"hardLimitBytes,omitempty"`
}

// NewStorageDomainPhysicalQuota instantiates a new StorageDomainPhysicalQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageDomainPhysicalQuota() *StorageDomainPhysicalQuota {
	this := StorageDomainPhysicalQuota{}
	return &this
}

// NewStorageDomainPhysicalQuotaWithDefaults instantiates a new StorageDomainPhysicalQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageDomainPhysicalQuotaWithDefaults() *StorageDomainPhysicalQuota {
	this := StorageDomainPhysicalQuota{}
	return &this
}

// GetAlertLimitBytes returns the AlertLimitBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageDomainPhysicalQuota) GetAlertLimitBytes() int64 {
	if o == nil || IsNil(o.AlertLimitBytes.Get()) {
		var ret int64
		return ret
	}
	return *o.AlertLimitBytes.Get()
}

// GetAlertLimitBytesOk returns a tuple with the AlertLimitBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageDomainPhysicalQuota) GetAlertLimitBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlertLimitBytes.Get(), o.AlertLimitBytes.IsSet()
}

// HasAlertLimitBytes returns a boolean if a field has been set.
func (o *StorageDomainPhysicalQuota) HasAlertLimitBytes() bool {
	if o != nil && o.AlertLimitBytes.IsSet() {
		return true
	}

	return false
}

// SetAlertLimitBytes gets a reference to the given NullableInt64 and assigns it to the AlertLimitBytes field.
func (o *StorageDomainPhysicalQuota) SetAlertLimitBytes(v int64) {
	o.AlertLimitBytes.Set(&v)
}
// SetAlertLimitBytesNil sets the value for AlertLimitBytes to be an explicit nil
func (o *StorageDomainPhysicalQuota) SetAlertLimitBytesNil() {
	o.AlertLimitBytes.Set(nil)
}

// UnsetAlertLimitBytes ensures that no value is present for AlertLimitBytes, not even an explicit nil
func (o *StorageDomainPhysicalQuota) UnsetAlertLimitBytes() {
	o.AlertLimitBytes.Unset()
}

// GetAlertThresholdPercentage returns the AlertThresholdPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageDomainPhysicalQuota) GetAlertThresholdPercentage() int64 {
	if o == nil || IsNil(o.AlertThresholdPercentage.Get()) {
		var ret int64
		return ret
	}
	return *o.AlertThresholdPercentage.Get()
}

// GetAlertThresholdPercentageOk returns a tuple with the AlertThresholdPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageDomainPhysicalQuota) GetAlertThresholdPercentageOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlertThresholdPercentage.Get(), o.AlertThresholdPercentage.IsSet()
}

// HasAlertThresholdPercentage returns a boolean if a field has been set.
func (o *StorageDomainPhysicalQuota) HasAlertThresholdPercentage() bool {
	if o != nil && o.AlertThresholdPercentage.IsSet() {
		return true
	}

	return false
}

// SetAlertThresholdPercentage gets a reference to the given NullableInt64 and assigns it to the AlertThresholdPercentage field.
func (o *StorageDomainPhysicalQuota) SetAlertThresholdPercentage(v int64) {
	o.AlertThresholdPercentage.Set(&v)
}
// SetAlertThresholdPercentageNil sets the value for AlertThresholdPercentage to be an explicit nil
func (o *StorageDomainPhysicalQuota) SetAlertThresholdPercentageNil() {
	o.AlertThresholdPercentage.Set(nil)
}

// UnsetAlertThresholdPercentage ensures that no value is present for AlertThresholdPercentage, not even an explicit nil
func (o *StorageDomainPhysicalQuota) UnsetAlertThresholdPercentage() {
	o.AlertThresholdPercentage.Unset()
}

// GetHardLimitBytes returns the HardLimitBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageDomainPhysicalQuota) GetHardLimitBytes() int64 {
	if o == nil || IsNil(o.HardLimitBytes.Get()) {
		var ret int64
		return ret
	}
	return *o.HardLimitBytes.Get()
}

// GetHardLimitBytesOk returns a tuple with the HardLimitBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageDomainPhysicalQuota) GetHardLimitBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.HardLimitBytes.Get(), o.HardLimitBytes.IsSet()
}

// HasHardLimitBytes returns a boolean if a field has been set.
func (o *StorageDomainPhysicalQuota) HasHardLimitBytes() bool {
	if o != nil && o.HardLimitBytes.IsSet() {
		return true
	}

	return false
}

// SetHardLimitBytes gets a reference to the given NullableInt64 and assigns it to the HardLimitBytes field.
func (o *StorageDomainPhysicalQuota) SetHardLimitBytes(v int64) {
	o.HardLimitBytes.Set(&v)
}
// SetHardLimitBytesNil sets the value for HardLimitBytes to be an explicit nil
func (o *StorageDomainPhysicalQuota) SetHardLimitBytesNil() {
	o.HardLimitBytes.Set(nil)
}

// UnsetHardLimitBytes ensures that no value is present for HardLimitBytes, not even an explicit nil
func (o *StorageDomainPhysicalQuota) UnsetHardLimitBytes() {
	o.HardLimitBytes.Unset()
}

func (o StorageDomainPhysicalQuota) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageDomainPhysicalQuota) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AlertLimitBytes.IsSet() {
		toSerialize["alertLimitBytes"] = o.AlertLimitBytes.Get()
	}
	if o.AlertThresholdPercentage.IsSet() {
		toSerialize["alertThresholdPercentage"] = o.AlertThresholdPercentage.Get()
	}
	if o.HardLimitBytes.IsSet() {
		toSerialize["hardLimitBytes"] = o.HardLimitBytes.Get()
	}
	return toSerialize, nil
}

type NullableStorageDomainPhysicalQuota struct {
	value *StorageDomainPhysicalQuota
	isSet bool
}

func (v NullableStorageDomainPhysicalQuota) Get() *StorageDomainPhysicalQuota {
	return v.value
}

func (v *NullableStorageDomainPhysicalQuota) Set(val *StorageDomainPhysicalQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageDomainPhysicalQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageDomainPhysicalQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageDomainPhysicalQuota(val *StorageDomainPhysicalQuota) *NullableStorageDomainPhysicalQuota {
	return &NullableStorageDomainPhysicalQuota{value: val, isSet: true}
}

func (v NullableStorageDomainPhysicalQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageDomainPhysicalQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


