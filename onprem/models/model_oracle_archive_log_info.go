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

// OracleArchiveLogInfo Specifies information related to archive log restore.
type OracleArchiveLogInfo struct {
	// Specifies the type of range.
	RangeType NullableString `json:"rangeType,omitempty"`
	// Specifies an array of oracle restore ranges.
	RangeInfoVec []OracleRangeMetaInfo `json:"rangeInfoVec,omitempty"`
	// Specifies destination where archive logs are to be restored.
	ArchiveLogRestoreDest NullableString `json:"archiveLogRestoreDest,omitempty"`
}

// NewOracleArchiveLogInfo instantiates a new OracleArchiveLogInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOracleArchiveLogInfo() *OracleArchiveLogInfo {
	this := OracleArchiveLogInfo{}
	return &this
}

// NewOracleArchiveLogInfoWithDefaults instantiates a new OracleArchiveLogInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOracleArchiveLogInfoWithDefaults() *OracleArchiveLogInfo {
	this := OracleArchiveLogInfo{}
	return &this
}

// GetRangeType returns the RangeType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OracleArchiveLogInfo) GetRangeType() string {
	if o == nil || o.RangeType.Get() == nil {
		var ret string
		return ret
	}
	return *o.RangeType.Get()
}

// GetRangeTypeOk returns a tuple with the RangeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OracleArchiveLogInfo) GetRangeTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RangeType.Get(), o.RangeType.IsSet()
}

// HasRangeType returns a boolean if a field has been set.
func (o *OracleArchiveLogInfo) HasRangeType() bool {
	if o != nil && o.RangeType.IsSet() {
		return true
	}

	return false
}

// SetRangeType gets a reference to the given NullableString and assigns it to the RangeType field.
func (o *OracleArchiveLogInfo) SetRangeType(v string) {
	o.RangeType.Set(&v)
}
// SetRangeTypeNil sets the value for RangeType to be an explicit nil
func (o *OracleArchiveLogInfo) SetRangeTypeNil() {
	o.RangeType.Set(nil)
}

// UnsetRangeType ensures that no value is present for RangeType, not even an explicit nil
func (o *OracleArchiveLogInfo) UnsetRangeType() {
	o.RangeType.Unset()
}

// GetRangeInfoVec returns the RangeInfoVec field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OracleArchiveLogInfo) GetRangeInfoVec() []OracleRangeMetaInfo {
	if o == nil  {
		var ret []OracleRangeMetaInfo
		return ret
	}
	return o.RangeInfoVec
}

// GetRangeInfoVecOk returns a tuple with the RangeInfoVec field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OracleArchiveLogInfo) GetRangeInfoVecOk() (*[]OracleRangeMetaInfo, bool) {
	if o == nil || o.RangeInfoVec == nil {
		return nil, false
	}
	return &o.RangeInfoVec, true
}

// HasRangeInfoVec returns a boolean if a field has been set.
func (o *OracleArchiveLogInfo) HasRangeInfoVec() bool {
	if o != nil && o.RangeInfoVec != nil {
		return true
	}

	return false
}

// SetRangeInfoVec gets a reference to the given []OracleRangeMetaInfo and assigns it to the RangeInfoVec field.
func (o *OracleArchiveLogInfo) SetRangeInfoVec(v []OracleRangeMetaInfo) {
	o.RangeInfoVec = v
}

// GetArchiveLogRestoreDest returns the ArchiveLogRestoreDest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OracleArchiveLogInfo) GetArchiveLogRestoreDest() string {
	if o == nil || o.ArchiveLogRestoreDest.Get() == nil {
		var ret string
		return ret
	}
	return *o.ArchiveLogRestoreDest.Get()
}

// GetArchiveLogRestoreDestOk returns a tuple with the ArchiveLogRestoreDest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OracleArchiveLogInfo) GetArchiveLogRestoreDestOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ArchiveLogRestoreDest.Get(), o.ArchiveLogRestoreDest.IsSet()
}

// HasArchiveLogRestoreDest returns a boolean if a field has been set.
func (o *OracleArchiveLogInfo) HasArchiveLogRestoreDest() bool {
	if o != nil && o.ArchiveLogRestoreDest.IsSet() {
		return true
	}

	return false
}

// SetArchiveLogRestoreDest gets a reference to the given NullableString and assigns it to the ArchiveLogRestoreDest field.
func (o *OracleArchiveLogInfo) SetArchiveLogRestoreDest(v string) {
	o.ArchiveLogRestoreDest.Set(&v)
}
// SetArchiveLogRestoreDestNil sets the value for ArchiveLogRestoreDest to be an explicit nil
func (o *OracleArchiveLogInfo) SetArchiveLogRestoreDestNil() {
	o.ArchiveLogRestoreDest.Set(nil)
}

// UnsetArchiveLogRestoreDest ensures that no value is present for ArchiveLogRestoreDest, not even an explicit nil
func (o *OracleArchiveLogInfo) UnsetArchiveLogRestoreDest() {
	o.ArchiveLogRestoreDest.Unset()
}

func (o OracleArchiveLogInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RangeType.IsSet() {
		toSerialize["rangeType"] = o.RangeType.Get()
	}
	if o.RangeInfoVec != nil {
		toSerialize["rangeInfoVec"] = o.RangeInfoVec
	}
	if o.ArchiveLogRestoreDest.IsSet() {
		toSerialize["archiveLogRestoreDest"] = o.ArchiveLogRestoreDest.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableOracleArchiveLogInfo struct {
	value *OracleArchiveLogInfo
	isSet bool
}

func (v NullableOracleArchiveLogInfo) Get() *OracleArchiveLogInfo {
	return v.value
}

func (v *NullableOracleArchiveLogInfo) Set(val *OracleArchiveLogInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOracleArchiveLogInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOracleArchiveLogInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOracleArchiveLogInfo(val *OracleArchiveLogInfo) *NullableOracleArchiveLogInfo {
	return &NullableOracleArchiveLogInfo{value: val, isSet: true}
}

func (v NullableOracleArchiveLogInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOracleArchiveLogInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o OracleArchiveLogInfo) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}