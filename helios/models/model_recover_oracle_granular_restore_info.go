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

// RecoverOracleGranularRestoreInfo Specifies information about list of objects (PDBs) to restore.
type RecoverOracleGranularRestoreInfo struct {
	// Specifies type of granular restore.
	GranularityType NullableString `json:"granularityType,omitempty"`
	// Specifies information about the list of pdbs to be restored.
	PdbRestoreParams *OraclePdbRestoreParams `json:"pdbRestoreParams,omitempty"`
}

// NewRecoverOracleGranularRestoreInfo instantiates a new RecoverOracleGranularRestoreInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverOracleGranularRestoreInfo() *RecoverOracleGranularRestoreInfo {
	this := RecoverOracleGranularRestoreInfo{}
	return &this
}

// NewRecoverOracleGranularRestoreInfoWithDefaults instantiates a new RecoverOracleGranularRestoreInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverOracleGranularRestoreInfoWithDefaults() *RecoverOracleGranularRestoreInfo {
	this := RecoverOracleGranularRestoreInfo{}
	return &this
}

// GetGranularityType returns the GranularityType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverOracleGranularRestoreInfo) GetGranularityType() string {
	if o == nil || o.GranularityType.Get() == nil {
		var ret string
		return ret
	}
	return *o.GranularityType.Get()
}

// GetGranularityTypeOk returns a tuple with the GranularityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverOracleGranularRestoreInfo) GetGranularityTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GranularityType.Get(), o.GranularityType.IsSet()
}

// HasGranularityType returns a boolean if a field has been set.
func (o *RecoverOracleGranularRestoreInfo) HasGranularityType() bool {
	if o != nil && o.GranularityType.IsSet() {
		return true
	}

	return false
}

// SetGranularityType gets a reference to the given NullableString and assigns it to the GranularityType field.
func (o *RecoverOracleGranularRestoreInfo) SetGranularityType(v string) {
	o.GranularityType.Set(&v)
}
// SetGranularityTypeNil sets the value for GranularityType to be an explicit nil
func (o *RecoverOracleGranularRestoreInfo) SetGranularityTypeNil() {
	o.GranularityType.Set(nil)
}

// UnsetGranularityType ensures that no value is present for GranularityType, not even an explicit nil
func (o *RecoverOracleGranularRestoreInfo) UnsetGranularityType() {
	o.GranularityType.Unset()
}

// GetPdbRestoreParams returns the PdbRestoreParams field value if set, zero value otherwise.
func (o *RecoverOracleGranularRestoreInfo) GetPdbRestoreParams() OraclePdbRestoreParams {
	if o == nil || o.PdbRestoreParams == nil {
		var ret OraclePdbRestoreParams
		return ret
	}
	return *o.PdbRestoreParams
}

// GetPdbRestoreParamsOk returns a tuple with the PdbRestoreParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoverOracleGranularRestoreInfo) GetPdbRestoreParamsOk() (*OraclePdbRestoreParams, bool) {
	if o == nil || o.PdbRestoreParams == nil {
		return nil, false
	}
	return o.PdbRestoreParams, true
}

// HasPdbRestoreParams returns a boolean if a field has been set.
func (o *RecoverOracleGranularRestoreInfo) HasPdbRestoreParams() bool {
	if o != nil && o.PdbRestoreParams != nil {
		return true
	}

	return false
}

// SetPdbRestoreParams gets a reference to the given OraclePdbRestoreParams and assigns it to the PdbRestoreParams field.
func (o *RecoverOracleGranularRestoreInfo) SetPdbRestoreParams(v OraclePdbRestoreParams) {
	o.PdbRestoreParams = &v
}

func (o RecoverOracleGranularRestoreInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GranularityType.IsSet() {
		toSerialize["granularityType"] = o.GranularityType.Get()
	}
	if o.PdbRestoreParams != nil {
		toSerialize["pdbRestoreParams"] = o.PdbRestoreParams
	}
	return json.Marshal(toSerialize)
}

type NullableRecoverOracleGranularRestoreInfo struct {
	value *RecoverOracleGranularRestoreInfo
	isSet bool
}

func (v NullableRecoverOracleGranularRestoreInfo) Get() *RecoverOracleGranularRestoreInfo {
	return v.value
}

func (v *NullableRecoverOracleGranularRestoreInfo) Set(val *RecoverOracleGranularRestoreInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverOracleGranularRestoreInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverOracleGranularRestoreInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverOracleGranularRestoreInfo(val *RecoverOracleGranularRestoreInfo) *NullableRecoverOracleGranularRestoreInfo {
	return &NullableRecoverOracleGranularRestoreInfo{value: val, isSet: true}
}

func (v NullableRecoverOracleGranularRestoreInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverOracleGranularRestoreInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


