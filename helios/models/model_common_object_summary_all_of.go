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

// CommonObjectSummaryAllOf struct for CommonObjectSummaryAllOf
type CommonObjectSummaryAllOf struct {
	// Specifies the count and size of protected and unprotected objects for the size.
	ProtectionStats []ObjectProtectionStatsSummary `json:"protectionStats,omitempty"`
	Permissions *PermissionInfo `json:"permissions,omitempty"`
}

// NewCommonObjectSummaryAllOf instantiates a new CommonObjectSummaryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonObjectSummaryAllOf() *CommonObjectSummaryAllOf {
	this := CommonObjectSummaryAllOf{}
	return &this
}

// NewCommonObjectSummaryAllOfWithDefaults instantiates a new CommonObjectSummaryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonObjectSummaryAllOfWithDefaults() *CommonObjectSummaryAllOf {
	this := CommonObjectSummaryAllOf{}
	return &this
}

// GetProtectionStats returns the ProtectionStats field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonObjectSummaryAllOf) GetProtectionStats() []ObjectProtectionStatsSummary {
	if o == nil  {
		var ret []ObjectProtectionStatsSummary
		return ret
	}
	return o.ProtectionStats
}

// GetProtectionStatsOk returns a tuple with the ProtectionStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonObjectSummaryAllOf) GetProtectionStatsOk() (*[]ObjectProtectionStatsSummary, bool) {
	if o == nil || o.ProtectionStats == nil {
		return nil, false
	}
	return &o.ProtectionStats, true
}

// HasProtectionStats returns a boolean if a field has been set.
func (o *CommonObjectSummaryAllOf) HasProtectionStats() bool {
	if o != nil && o.ProtectionStats != nil {
		return true
	}

	return false
}

// SetProtectionStats gets a reference to the given []ObjectProtectionStatsSummary and assigns it to the ProtectionStats field.
func (o *CommonObjectSummaryAllOf) SetProtectionStats(v []ObjectProtectionStatsSummary) {
	o.ProtectionStats = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *CommonObjectSummaryAllOf) GetPermissions() PermissionInfo {
	if o == nil || o.Permissions == nil {
		var ret PermissionInfo
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonObjectSummaryAllOf) GetPermissionsOk() (*PermissionInfo, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *CommonObjectSummaryAllOf) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given PermissionInfo and assigns it to the Permissions field.
func (o *CommonObjectSummaryAllOf) SetPermissions(v PermissionInfo) {
	o.Permissions = &v
}

func (o CommonObjectSummaryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProtectionStats != nil {
		toSerialize["protectionStats"] = o.ProtectionStats
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableCommonObjectSummaryAllOf struct {
	value *CommonObjectSummaryAllOf
	isSet bool
}

func (v NullableCommonObjectSummaryAllOf) Get() *CommonObjectSummaryAllOf {
	return v.value
}

func (v *NullableCommonObjectSummaryAllOf) Set(val *CommonObjectSummaryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonObjectSummaryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonObjectSummaryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonObjectSummaryAllOf(val *CommonObjectSummaryAllOf) *NullableCommonObjectSummaryAllOf {
	return &NullableCommonObjectSummaryAllOf{value: val, isSet: true}
}

func (v NullableCommonObjectSummaryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonObjectSummaryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


