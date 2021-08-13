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

// GetPITRangesProtectedObjectResponseBodyAllOf struct for GetPITRangesProtectedObjectResponseBodyAllOf
type GetPITRangesProtectedObjectResponseBodyAllOf struct {
	OracleRestoreRangeInfo *OracleRestoreRangeInfo `json:"oracleRestoreRangeInfo,omitempty"`
}

// NewGetPITRangesProtectedObjectResponseBodyAllOf instantiates a new GetPITRangesProtectedObjectResponseBodyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPITRangesProtectedObjectResponseBodyAllOf() *GetPITRangesProtectedObjectResponseBodyAllOf {
	this := GetPITRangesProtectedObjectResponseBodyAllOf{}
	return &this
}

// NewGetPITRangesProtectedObjectResponseBodyAllOfWithDefaults instantiates a new GetPITRangesProtectedObjectResponseBodyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPITRangesProtectedObjectResponseBodyAllOfWithDefaults() *GetPITRangesProtectedObjectResponseBodyAllOf {
	this := GetPITRangesProtectedObjectResponseBodyAllOf{}
	return &this
}

// GetOracleRestoreRangeInfo returns the OracleRestoreRangeInfo field value if set, zero value otherwise.
func (o *GetPITRangesProtectedObjectResponseBodyAllOf) GetOracleRestoreRangeInfo() OracleRestoreRangeInfo {
	if o == nil || o.OracleRestoreRangeInfo == nil {
		var ret OracleRestoreRangeInfo
		return ret
	}
	return *o.OracleRestoreRangeInfo
}

// GetOracleRestoreRangeInfoOk returns a tuple with the OracleRestoreRangeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPITRangesProtectedObjectResponseBodyAllOf) GetOracleRestoreRangeInfoOk() (*OracleRestoreRangeInfo, bool) {
	if o == nil || o.OracleRestoreRangeInfo == nil {
		return nil, false
	}
	return o.OracleRestoreRangeInfo, true
}

// HasOracleRestoreRangeInfo returns a boolean if a field has been set.
func (o *GetPITRangesProtectedObjectResponseBodyAllOf) HasOracleRestoreRangeInfo() bool {
	if o != nil && o.OracleRestoreRangeInfo != nil {
		return true
	}

	return false
}

// SetOracleRestoreRangeInfo gets a reference to the given OracleRestoreRangeInfo and assigns it to the OracleRestoreRangeInfo field.
func (o *GetPITRangesProtectedObjectResponseBodyAllOf) SetOracleRestoreRangeInfo(v OracleRestoreRangeInfo) {
	o.OracleRestoreRangeInfo = &v
}

func (o GetPITRangesProtectedObjectResponseBodyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OracleRestoreRangeInfo != nil {
		toSerialize["oracleRestoreRangeInfo"] = o.OracleRestoreRangeInfo
	}
	return json.Marshal(toSerialize)
}

type NullableGetPITRangesProtectedObjectResponseBodyAllOf struct {
	value *GetPITRangesProtectedObjectResponseBodyAllOf
	isSet bool
}

func (v NullableGetPITRangesProtectedObjectResponseBodyAllOf) Get() *GetPITRangesProtectedObjectResponseBodyAllOf {
	return v.value
}

func (v *NullableGetPITRangesProtectedObjectResponseBodyAllOf) Set(val *GetPITRangesProtectedObjectResponseBodyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPITRangesProtectedObjectResponseBodyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPITRangesProtectedObjectResponseBodyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPITRangesProtectedObjectResponseBodyAllOf(val *GetPITRangesProtectedObjectResponseBodyAllOf) *NullableGetPITRangesProtectedObjectResponseBodyAllOf {
	return &NullableGetPITRangesProtectedObjectResponseBodyAllOf{value: val, isSet: true}
}

func (v NullableGetPITRangesProtectedObjectResponseBodyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPITRangesProtectedObjectResponseBodyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


