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

// QuotaPolicy Specifies a quota limit that can be optionally applied to Views and Storage Domains. At the View level, this quota defines a logical limit for usage on the View. At the Storage Domain level, this quota defines a physical limit or a default logical View limit. If a physical quota is specified for Storage Domain, this quota defines a physical limit for the usage on the Storage Domain. If a default logical View quota is specified for Storage Domain, this limit is inherited by all the Views in that Storage Domain. However, this inherited quota can be overwritten at the View level. A new write is not allowed if the resource will exceed the specified quota. However, it takes time for the Cohesity Cluster to calculate the usage across Nodes, so the limit may be exceeded by a small amount. In addition, if the limit is increased or data is removed, there may be a delay before the Cohesity Cluster allows more data to be written to the resource, as the Cluster calculates the usage across Nodes.
type QuotaPolicy struct {
	// Specifies if an alert should be triggered when the usage of this resource exceeds this quota limit. This limit is optional and is specified in bytes. If no value is specified, there is no limit.
	AlertLimitBytes NullableInt64 `json:"alertLimitBytes,omitempty"`
	// Supported only for user quota policy. Specifies when the usage goes above an alert threshold percentage which is: HardLimitBytes * AlertThresholdPercentage, eg: 80% of HardLimitBytes Can only be set if HardLimitBytes is set. Cannot be set if AlertLimitBytes is already set.
	AlertThresholdPercentage NullableInt64 `json:"alertThresholdPercentage,omitempty"`
	// Specifies an optional quota limit on the usage allowed for this resource. This limit is specified in bytes. If no value is specified, there is no limit.
	HardLimitBytes NullableInt64 `json:"hardLimitBytes,omitempty"`
}

// NewQuotaPolicy instantiates a new QuotaPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuotaPolicy() *QuotaPolicy {
	this := QuotaPolicy{}
	return &this
}

// NewQuotaPolicyWithDefaults instantiates a new QuotaPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotaPolicyWithDefaults() *QuotaPolicy {
	this := QuotaPolicy{}
	return &this
}

// GetAlertLimitBytes returns the AlertLimitBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QuotaPolicy) GetAlertLimitBytes() int64 {
	if o == nil || o.AlertLimitBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AlertLimitBytes.Get()
}

// GetAlertLimitBytesOk returns a tuple with the AlertLimitBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QuotaPolicy) GetAlertLimitBytesOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AlertLimitBytes.Get(), o.AlertLimitBytes.IsSet()
}

// HasAlertLimitBytes returns a boolean if a field has been set.
func (o *QuotaPolicy) HasAlertLimitBytes() bool {
	if o != nil && o.AlertLimitBytes.IsSet() {
		return true
	}

	return false
}

// SetAlertLimitBytes gets a reference to the given NullableInt64 and assigns it to the AlertLimitBytes field.
func (o *QuotaPolicy) SetAlertLimitBytes(v int64) {
	o.AlertLimitBytes.Set(&v)
}
// SetAlertLimitBytesNil sets the value for AlertLimitBytes to be an explicit nil
func (o *QuotaPolicy) SetAlertLimitBytesNil() {
	o.AlertLimitBytes.Set(nil)
}

// UnsetAlertLimitBytes ensures that no value is present for AlertLimitBytes, not even an explicit nil
func (o *QuotaPolicy) UnsetAlertLimitBytes() {
	o.AlertLimitBytes.Unset()
}

// GetAlertThresholdPercentage returns the AlertThresholdPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QuotaPolicy) GetAlertThresholdPercentage() int64 {
	if o == nil || o.AlertThresholdPercentage.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AlertThresholdPercentage.Get()
}

// GetAlertThresholdPercentageOk returns a tuple with the AlertThresholdPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QuotaPolicy) GetAlertThresholdPercentageOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AlertThresholdPercentage.Get(), o.AlertThresholdPercentage.IsSet()
}

// HasAlertThresholdPercentage returns a boolean if a field has been set.
func (o *QuotaPolicy) HasAlertThresholdPercentage() bool {
	if o != nil && o.AlertThresholdPercentage.IsSet() {
		return true
	}

	return false
}

// SetAlertThresholdPercentage gets a reference to the given NullableInt64 and assigns it to the AlertThresholdPercentage field.
func (o *QuotaPolicy) SetAlertThresholdPercentage(v int64) {
	o.AlertThresholdPercentage.Set(&v)
}
// SetAlertThresholdPercentageNil sets the value for AlertThresholdPercentage to be an explicit nil
func (o *QuotaPolicy) SetAlertThresholdPercentageNil() {
	o.AlertThresholdPercentage.Set(nil)
}

// UnsetAlertThresholdPercentage ensures that no value is present for AlertThresholdPercentage, not even an explicit nil
func (o *QuotaPolicy) UnsetAlertThresholdPercentage() {
	o.AlertThresholdPercentage.Unset()
}

// GetHardLimitBytes returns the HardLimitBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QuotaPolicy) GetHardLimitBytes() int64 {
	if o == nil || o.HardLimitBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.HardLimitBytes.Get()
}

// GetHardLimitBytesOk returns a tuple with the HardLimitBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QuotaPolicy) GetHardLimitBytesOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HardLimitBytes.Get(), o.HardLimitBytes.IsSet()
}

// HasHardLimitBytes returns a boolean if a field has been set.
func (o *QuotaPolicy) HasHardLimitBytes() bool {
	if o != nil && o.HardLimitBytes.IsSet() {
		return true
	}

	return false
}

// SetHardLimitBytes gets a reference to the given NullableInt64 and assigns it to the HardLimitBytes field.
func (o *QuotaPolicy) SetHardLimitBytes(v int64) {
	o.HardLimitBytes.Set(&v)
}
// SetHardLimitBytesNil sets the value for HardLimitBytes to be an explicit nil
func (o *QuotaPolicy) SetHardLimitBytesNil() {
	o.HardLimitBytes.Set(nil)
}

// UnsetHardLimitBytes ensures that no value is present for HardLimitBytes, not even an explicit nil
func (o *QuotaPolicy) UnsetHardLimitBytes() {
	o.HardLimitBytes.Unset()
}

func (o QuotaPolicy) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableQuotaPolicy struct {
	value *QuotaPolicy
	isSet bool
}

func (v NullableQuotaPolicy) Get() *QuotaPolicy {
	return v.value
}

func (v *NullableQuotaPolicy) Set(val *QuotaPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableQuotaPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableQuotaPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuotaPolicy(val *QuotaPolicy) *NullableQuotaPolicy {
	return &NullableQuotaPolicy{value: val, isSet: true}
}

func (v NullableQuotaPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuotaPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


