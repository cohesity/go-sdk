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

// CommonPostBackupScriptParams Specifies the common params for PostBackup scripts.
type CommonPostBackupScriptParams struct {
	// Specifies the absolute path to the script on the remote host.
	Path string `json:"path"`
	// Specifies the arguments or parameters and values to pass into the remote script. For example if the script expects values for the 'database' and 'user' parameters, specify the parameters and values using the following string: \"database=myDatabase user=me\".
	Params NullableString `json:"params,omitempty"`
	// Specifies the timeout of the script in seconds. The script will be killed if it exceeds this value. By default, no timeout will occur if left empty.
	TimeoutSecs NullableInt32 `json:"timeoutSecs,omitempty"`
	// Specifies whether the script should be enabled, default value set to true.
	IsActive NullableBool `json:"isActive,omitempty"`
}

// NewCommonPostBackupScriptParams instantiates a new CommonPostBackupScriptParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonPostBackupScriptParams(path string) *CommonPostBackupScriptParams {
	this := CommonPostBackupScriptParams{}
	this.Path = path
	return &this
}

// NewCommonPostBackupScriptParamsWithDefaults instantiates a new CommonPostBackupScriptParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonPostBackupScriptParamsWithDefaults() *CommonPostBackupScriptParams {
	this := CommonPostBackupScriptParams{}
	return &this
}

// GetPath returns the Path field value
func (o *CommonPostBackupScriptParams) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *CommonPostBackupScriptParams) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *CommonPostBackupScriptParams) SetPath(v string) {
	o.Path = v
}

// GetParams returns the Params field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonPostBackupScriptParams) GetParams() string {
	if o == nil || o.Params.Get() == nil {
		var ret string
		return ret
	}
	return *o.Params.Get()
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonPostBackupScriptParams) GetParamsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Params.Get(), o.Params.IsSet()
}

// HasParams returns a boolean if a field has been set.
func (o *CommonPostBackupScriptParams) HasParams() bool {
	if o != nil && o.Params.IsSet() {
		return true
	}

	return false
}

// SetParams gets a reference to the given NullableString and assigns it to the Params field.
func (o *CommonPostBackupScriptParams) SetParams(v string) {
	o.Params.Set(&v)
}
// SetParamsNil sets the value for Params to be an explicit nil
func (o *CommonPostBackupScriptParams) SetParamsNil() {
	o.Params.Set(nil)
}

// UnsetParams ensures that no value is present for Params, not even an explicit nil
func (o *CommonPostBackupScriptParams) UnsetParams() {
	o.Params.Unset()
}

// GetTimeoutSecs returns the TimeoutSecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonPostBackupScriptParams) GetTimeoutSecs() int32 {
	if o == nil || o.TimeoutSecs.Get() == nil {
		var ret int32
		return ret
	}
	return *o.TimeoutSecs.Get()
}

// GetTimeoutSecsOk returns a tuple with the TimeoutSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonPostBackupScriptParams) GetTimeoutSecsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TimeoutSecs.Get(), o.TimeoutSecs.IsSet()
}

// HasTimeoutSecs returns a boolean if a field has been set.
func (o *CommonPostBackupScriptParams) HasTimeoutSecs() bool {
	if o != nil && o.TimeoutSecs.IsSet() {
		return true
	}

	return false
}

// SetTimeoutSecs gets a reference to the given NullableInt32 and assigns it to the TimeoutSecs field.
func (o *CommonPostBackupScriptParams) SetTimeoutSecs(v int32) {
	o.TimeoutSecs.Set(&v)
}
// SetTimeoutSecsNil sets the value for TimeoutSecs to be an explicit nil
func (o *CommonPostBackupScriptParams) SetTimeoutSecsNil() {
	o.TimeoutSecs.Set(nil)
}

// UnsetTimeoutSecs ensures that no value is present for TimeoutSecs, not even an explicit nil
func (o *CommonPostBackupScriptParams) UnsetTimeoutSecs() {
	o.TimeoutSecs.Unset()
}

// GetIsActive returns the IsActive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonPostBackupScriptParams) GetIsActive() bool {
	if o == nil || o.IsActive.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsActive.Get()
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonPostBackupScriptParams) GetIsActiveOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsActive.Get(), o.IsActive.IsSet()
}

// HasIsActive returns a boolean if a field has been set.
func (o *CommonPostBackupScriptParams) HasIsActive() bool {
	if o != nil && o.IsActive.IsSet() {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given NullableBool and assigns it to the IsActive field.
func (o *CommonPostBackupScriptParams) SetIsActive(v bool) {
	o.IsActive.Set(&v)
}
// SetIsActiveNil sets the value for IsActive to be an explicit nil
func (o *CommonPostBackupScriptParams) SetIsActiveNil() {
	o.IsActive.Set(nil)
}

// UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
func (o *CommonPostBackupScriptParams) UnsetIsActive() {
	o.IsActive.Unset()
}

func (o CommonPostBackupScriptParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["path"] = o.Path
	}
	if o.Params.IsSet() {
		toSerialize["params"] = o.Params.Get()
	}
	if o.TimeoutSecs.IsSet() {
		toSerialize["timeoutSecs"] = o.TimeoutSecs.Get()
	}
	if o.IsActive.IsSet() {
		toSerialize["isActive"] = o.IsActive.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCommonPostBackupScriptParams struct {
	value *CommonPostBackupScriptParams
	isSet bool
}

func (v NullableCommonPostBackupScriptParams) Get() *CommonPostBackupScriptParams {
	return v.value
}

func (v *NullableCommonPostBackupScriptParams) Set(val *CommonPostBackupScriptParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonPostBackupScriptParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonPostBackupScriptParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonPostBackupScriptParams(val *CommonPostBackupScriptParams) *NullableCommonPostBackupScriptParams {
	return &NullableCommonPostBackupScriptParams{value: val, isSet: true}
}

func (v NullableCommonPostBackupScriptParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonPostBackupScriptParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


