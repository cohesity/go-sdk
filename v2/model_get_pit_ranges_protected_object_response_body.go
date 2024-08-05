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

// checks if the GetPITRangesProtectedObjectResponseBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPITRangesProtectedObjectResponseBody{}

// GetPITRangesProtectedObjectResponseBody Specifies the points in time available for restore as a set of one or more time ranges. If the number of available ranges exceeds 1000, then the latest 1000 will be returned.
type GetPITRangesProtectedObjectResponseBody struct {
	// Specifies the environment for which restore ranges are returned.
	Environment NullableString `json:"environment,omitempty"`
	OracleRestoreRangeInfo *OracleRestoreRangeInfo `json:"oracleRestoreRangeInfo,omitempty"`
}

// NewGetPITRangesProtectedObjectResponseBody instantiates a new GetPITRangesProtectedObjectResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPITRangesProtectedObjectResponseBody() *GetPITRangesProtectedObjectResponseBody {
	this := GetPITRangesProtectedObjectResponseBody{}
	return &this
}

// NewGetPITRangesProtectedObjectResponseBodyWithDefaults instantiates a new GetPITRangesProtectedObjectResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPITRangesProtectedObjectResponseBodyWithDefaults() *GetPITRangesProtectedObjectResponseBody {
	this := GetPITRangesProtectedObjectResponseBody{}
	return &this
}

// GetEnvironment returns the Environment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetPITRangesProtectedObjectResponseBody) GetEnvironment() string {
	if o == nil || IsNil(o.Environment.Get()) {
		var ret string
		return ret
	}
	return *o.Environment.Get()
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetPITRangesProtectedObjectResponseBody) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Environment.Get(), o.Environment.IsSet()
}

// HasEnvironment returns a boolean if a field has been set.
func (o *GetPITRangesProtectedObjectResponseBody) HasEnvironment() bool {
	if o != nil && o.Environment.IsSet() {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given NullableString and assigns it to the Environment field.
func (o *GetPITRangesProtectedObjectResponseBody) SetEnvironment(v string) {
	o.Environment.Set(&v)
}
// SetEnvironmentNil sets the value for Environment to be an explicit nil
func (o *GetPITRangesProtectedObjectResponseBody) SetEnvironmentNil() {
	o.Environment.Set(nil)
}

// UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
func (o *GetPITRangesProtectedObjectResponseBody) UnsetEnvironment() {
	o.Environment.Unset()
}

// GetOracleRestoreRangeInfo returns the OracleRestoreRangeInfo field value if set, zero value otherwise.
func (o *GetPITRangesProtectedObjectResponseBody) GetOracleRestoreRangeInfo() OracleRestoreRangeInfo {
	if o == nil || IsNil(o.OracleRestoreRangeInfo) {
		var ret OracleRestoreRangeInfo
		return ret
	}
	return *o.OracleRestoreRangeInfo
}

// GetOracleRestoreRangeInfoOk returns a tuple with the OracleRestoreRangeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPITRangesProtectedObjectResponseBody) GetOracleRestoreRangeInfoOk() (*OracleRestoreRangeInfo, bool) {
	if o == nil || IsNil(o.OracleRestoreRangeInfo) {
		return nil, false
	}
	return o.OracleRestoreRangeInfo, true
}

// HasOracleRestoreRangeInfo returns a boolean if a field has been set.
func (o *GetPITRangesProtectedObjectResponseBody) HasOracleRestoreRangeInfo() bool {
	if o != nil && !IsNil(o.OracleRestoreRangeInfo) {
		return true
	}

	return false
}

// SetOracleRestoreRangeInfo gets a reference to the given OracleRestoreRangeInfo and assigns it to the OracleRestoreRangeInfo field.
func (o *GetPITRangesProtectedObjectResponseBody) SetOracleRestoreRangeInfo(v OracleRestoreRangeInfo) {
	o.OracleRestoreRangeInfo = &v
}

func (o GetPITRangesProtectedObjectResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPITRangesProtectedObjectResponseBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Environment.IsSet() {
		toSerialize["environment"] = o.Environment.Get()
	}
	if !IsNil(o.OracleRestoreRangeInfo) {
		toSerialize["oracleRestoreRangeInfo"] = o.OracleRestoreRangeInfo
	}
	return toSerialize, nil
}

type NullableGetPITRangesProtectedObjectResponseBody struct {
	value *GetPITRangesProtectedObjectResponseBody
	isSet bool
}

func (v NullableGetPITRangesProtectedObjectResponseBody) Get() *GetPITRangesProtectedObjectResponseBody {
	return v.value
}

func (v *NullableGetPITRangesProtectedObjectResponseBody) Set(val *GetPITRangesProtectedObjectResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPITRangesProtectedObjectResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPITRangesProtectedObjectResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPITRangesProtectedObjectResponseBody(val *GetPITRangesProtectedObjectResponseBody) *NullableGetPITRangesProtectedObjectResponseBody {
	return &NullableGetPITRangesProtectedObjectResponseBody{value: val, isSet: true}
}

func (v NullableGetPITRangesProtectedObjectResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPITRangesProtectedObjectResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


