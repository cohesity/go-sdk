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

// ViewDirectoryQuota Specifies a View directory quota.
type ViewDirectoryQuota struct {
	// Specifies the directory path.
	DirectoryPath NullableString `json:"directoryPath"`
	// Specifies the directory quota policy.
	QuotaPolicy ViewDirectoryQuotaPolicy `json:"quotaPolicy"`
	// Specifies the directory usage in bytes.
	UsageBytes NullableInt64 `json:"usageBytes,omitempty"`
	// Specifies whether directory quota walk is pending.
	DirectoryWalkPending NullableBool `json:"directoryWalkPending,omitempty"`
}

// NewViewDirectoryQuota instantiates a new ViewDirectoryQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewDirectoryQuota(directoryPath NullableString, quotaPolicy ViewDirectoryQuotaPolicy) *ViewDirectoryQuota {
	this := ViewDirectoryQuota{}
	this.DirectoryPath = directoryPath
	this.QuotaPolicy = quotaPolicy
	return &this
}

// NewViewDirectoryQuotaWithDefaults instantiates a new ViewDirectoryQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewDirectoryQuotaWithDefaults() *ViewDirectoryQuota {
	this := ViewDirectoryQuota{}
	return &this
}

// GetDirectoryPath returns the DirectoryPath field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ViewDirectoryQuota) GetDirectoryPath() string {
	if o == nil || o.DirectoryPath.Get() == nil {
		var ret string
		return ret
	}

	return *o.DirectoryPath.Get()
}

// GetDirectoryPathOk returns a tuple with the DirectoryPath field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ViewDirectoryQuota) GetDirectoryPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DirectoryPath.Get(), o.DirectoryPath.IsSet()
}

// SetDirectoryPath sets field value
func (o *ViewDirectoryQuota) SetDirectoryPath(v string) {
	o.DirectoryPath.Set(&v)
}

// GetQuotaPolicy returns the QuotaPolicy field value
func (o *ViewDirectoryQuota) GetQuotaPolicy() ViewDirectoryQuotaPolicy {
	if o == nil {
		var ret ViewDirectoryQuotaPolicy
		return ret
	}

	return o.QuotaPolicy
}

// GetQuotaPolicyOk returns a tuple with the QuotaPolicy field value
// and a boolean to check if the value has been set.
func (o *ViewDirectoryQuota) GetQuotaPolicyOk() (*ViewDirectoryQuotaPolicy, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.QuotaPolicy, true
}

// SetQuotaPolicy sets field value
func (o *ViewDirectoryQuota) SetQuotaPolicy(v ViewDirectoryQuotaPolicy) {
	o.QuotaPolicy = v
}

// GetUsageBytes returns the UsageBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ViewDirectoryQuota) GetUsageBytes() int64 {
	if o == nil || o.UsageBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.UsageBytes.Get()
}

// GetUsageBytesOk returns a tuple with the UsageBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ViewDirectoryQuota) GetUsageBytesOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UsageBytes.Get(), o.UsageBytes.IsSet()
}

// HasUsageBytes returns a boolean if a field has been set.
func (o *ViewDirectoryQuota) HasUsageBytes() bool {
	if o != nil && o.UsageBytes.IsSet() {
		return true
	}

	return false
}

// SetUsageBytes gets a reference to the given NullableInt64 and assigns it to the UsageBytes field.
func (o *ViewDirectoryQuota) SetUsageBytes(v int64) {
	o.UsageBytes.Set(&v)
}
// SetUsageBytesNil sets the value for UsageBytes to be an explicit nil
func (o *ViewDirectoryQuota) SetUsageBytesNil() {
	o.UsageBytes.Set(nil)
}

// UnsetUsageBytes ensures that no value is present for UsageBytes, not even an explicit nil
func (o *ViewDirectoryQuota) UnsetUsageBytes() {
	o.UsageBytes.Unset()
}

// GetDirectoryWalkPending returns the DirectoryWalkPending field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ViewDirectoryQuota) GetDirectoryWalkPending() bool {
	if o == nil || o.DirectoryWalkPending.Get() == nil {
		var ret bool
		return ret
	}
	return *o.DirectoryWalkPending.Get()
}

// GetDirectoryWalkPendingOk returns a tuple with the DirectoryWalkPending field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ViewDirectoryQuota) GetDirectoryWalkPendingOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DirectoryWalkPending.Get(), o.DirectoryWalkPending.IsSet()
}

// HasDirectoryWalkPending returns a boolean if a field has been set.
func (o *ViewDirectoryQuota) HasDirectoryWalkPending() bool {
	if o != nil && o.DirectoryWalkPending.IsSet() {
		return true
	}

	return false
}

// SetDirectoryWalkPending gets a reference to the given NullableBool and assigns it to the DirectoryWalkPending field.
func (o *ViewDirectoryQuota) SetDirectoryWalkPending(v bool) {
	o.DirectoryWalkPending.Set(&v)
}
// SetDirectoryWalkPendingNil sets the value for DirectoryWalkPending to be an explicit nil
func (o *ViewDirectoryQuota) SetDirectoryWalkPendingNil() {
	o.DirectoryWalkPending.Set(nil)
}

// UnsetDirectoryWalkPending ensures that no value is present for DirectoryWalkPending, not even an explicit nil
func (o *ViewDirectoryQuota) UnsetDirectoryWalkPending() {
	o.DirectoryWalkPending.Unset()
}

func (o ViewDirectoryQuota) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["directoryPath"] = o.DirectoryPath.Get()
	}
	if true {
		toSerialize["quotaPolicy"] = o.QuotaPolicy
	}
	if o.UsageBytes.IsSet() {
		toSerialize["usageBytes"] = o.UsageBytes.Get()
	}
	if o.DirectoryWalkPending.IsSet() {
		toSerialize["directoryWalkPending"] = o.DirectoryWalkPending.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableViewDirectoryQuota struct {
	value *ViewDirectoryQuota
	isSet bool
}

func (v NullableViewDirectoryQuota) Get() *ViewDirectoryQuota {
	return v.value
}

func (v *NullableViewDirectoryQuota) Set(val *ViewDirectoryQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableViewDirectoryQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableViewDirectoryQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewDirectoryQuota(val *ViewDirectoryQuota) *NullableViewDirectoryQuota {
	return &NullableViewDirectoryQuota{value: val, isSet: true}
}

func (v NullableViewDirectoryQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewDirectoryQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o ViewDirectoryQuota) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}