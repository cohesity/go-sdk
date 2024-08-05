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

// checks if the ConstructMetaInfoResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConstructMetaInfoResult{}

// ConstructMetaInfoResult Result to store meta-info from an object snapshot and additional information.
type ConstructMetaInfoResult struct {
	// Specifies the environment type for fetching the meta Info.
	Environment NullableString `json:"environment,omitempty"`
	OracleParams *ConstructMetaInfoResultOracleParams `json:"oracleParams,omitempty"`
	SfdcParams *ConstructMetaInfoResultSfdcParams `json:"sfdcParams,omitempty"`
}

// NewConstructMetaInfoResult instantiates a new ConstructMetaInfoResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConstructMetaInfoResult() *ConstructMetaInfoResult {
	this := ConstructMetaInfoResult{}
	return &this
}

// NewConstructMetaInfoResultWithDefaults instantiates a new ConstructMetaInfoResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConstructMetaInfoResultWithDefaults() *ConstructMetaInfoResult {
	this := ConstructMetaInfoResult{}
	return &this
}

// GetEnvironment returns the Environment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConstructMetaInfoResult) GetEnvironment() string {
	if o == nil || IsNil(o.Environment.Get()) {
		var ret string
		return ret
	}
	return *o.Environment.Get()
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConstructMetaInfoResult) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Environment.Get(), o.Environment.IsSet()
}

// HasEnvironment returns a boolean if a field has been set.
func (o *ConstructMetaInfoResult) HasEnvironment() bool {
	if o != nil && o.Environment.IsSet() {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given NullableString and assigns it to the Environment field.
func (o *ConstructMetaInfoResult) SetEnvironment(v string) {
	o.Environment.Set(&v)
}
// SetEnvironmentNil sets the value for Environment to be an explicit nil
func (o *ConstructMetaInfoResult) SetEnvironmentNil() {
	o.Environment.Set(nil)
}

// UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
func (o *ConstructMetaInfoResult) UnsetEnvironment() {
	o.Environment.Unset()
}

// GetOracleParams returns the OracleParams field value if set, zero value otherwise.
func (o *ConstructMetaInfoResult) GetOracleParams() ConstructMetaInfoResultOracleParams {
	if o == nil || IsNil(o.OracleParams) {
		var ret ConstructMetaInfoResultOracleParams
		return ret
	}
	return *o.OracleParams
}

// GetOracleParamsOk returns a tuple with the OracleParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstructMetaInfoResult) GetOracleParamsOk() (*ConstructMetaInfoResultOracleParams, bool) {
	if o == nil || IsNil(o.OracleParams) {
		return nil, false
	}
	return o.OracleParams, true
}

// HasOracleParams returns a boolean if a field has been set.
func (o *ConstructMetaInfoResult) HasOracleParams() bool {
	if o != nil && !IsNil(o.OracleParams) {
		return true
	}

	return false
}

// SetOracleParams gets a reference to the given ConstructMetaInfoResultOracleParams and assigns it to the OracleParams field.
func (o *ConstructMetaInfoResult) SetOracleParams(v ConstructMetaInfoResultOracleParams) {
	o.OracleParams = &v
}

// GetSfdcParams returns the SfdcParams field value if set, zero value otherwise.
func (o *ConstructMetaInfoResult) GetSfdcParams() ConstructMetaInfoResultSfdcParams {
	if o == nil || IsNil(o.SfdcParams) {
		var ret ConstructMetaInfoResultSfdcParams
		return ret
	}
	return *o.SfdcParams
}

// GetSfdcParamsOk returns a tuple with the SfdcParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstructMetaInfoResult) GetSfdcParamsOk() (*ConstructMetaInfoResultSfdcParams, bool) {
	if o == nil || IsNil(o.SfdcParams) {
		return nil, false
	}
	return o.SfdcParams, true
}

// HasSfdcParams returns a boolean if a field has been set.
func (o *ConstructMetaInfoResult) HasSfdcParams() bool {
	if o != nil && !IsNil(o.SfdcParams) {
		return true
	}

	return false
}

// SetSfdcParams gets a reference to the given ConstructMetaInfoResultSfdcParams and assigns it to the SfdcParams field.
func (o *ConstructMetaInfoResult) SetSfdcParams(v ConstructMetaInfoResultSfdcParams) {
	o.SfdcParams = &v
}

func (o ConstructMetaInfoResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConstructMetaInfoResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Environment.IsSet() {
		toSerialize["environment"] = o.Environment.Get()
	}
	if !IsNil(o.OracleParams) {
		toSerialize["oracleParams"] = o.OracleParams
	}
	if !IsNil(o.SfdcParams) {
		toSerialize["sfdcParams"] = o.SfdcParams
	}
	return toSerialize, nil
}

type NullableConstructMetaInfoResult struct {
	value *ConstructMetaInfoResult
	isSet bool
}

func (v NullableConstructMetaInfoResult) Get() *ConstructMetaInfoResult {
	return v.value
}

func (v *NullableConstructMetaInfoResult) Set(val *ConstructMetaInfoResult) {
	v.value = val
	v.isSet = true
}

func (v NullableConstructMetaInfoResult) IsSet() bool {
	return v.isSet
}

func (v *NullableConstructMetaInfoResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConstructMetaInfoResult(val *ConstructMetaInfoResult) *NullableConstructMetaInfoResult {
	return &NullableConstructMetaInfoResult{value: val, isSet: true}
}

func (v NullableConstructMetaInfoResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConstructMetaInfoResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


