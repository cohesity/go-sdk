/*
 * Cohesity REST API
 *
 * This API list provides operations for interfacing with the Cohesity Cluster.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cohesitysdk

import (
	"encoding/json"
)

// RestoreAppObjectParams struct for RestoreAppObjectParams
type RestoreAppObjectParams struct {
	AdRestoreParams *RestoreADAppObjectParams `json:"adRestoreParams,omitempty"`
	// Id of finished clone task which has to be refreshed with different data.
	CloneTaskId NullableInt64 `json:"cloneTaskId,omitempty"`
	ExchangeRestoreParams *RestoreExchangeParams `json:"exchangeRestoreParams,omitempty"`
	OracleRestoreParams *RestoreOracleAppObjectParams `json:"oracleRestoreParams,omitempty"`
	SqlRestoreParams *RestoreSqlAppObjectParams `json:"sqlRestoreParams,omitempty"`
	TargetHost *EntityProto `json:"targetHost,omitempty"`
	TargetHostParentSource *EntityProto `json:"targetHostParentSource,omitempty"`
}

// NewRestoreAppObjectParams instantiates a new RestoreAppObjectParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestoreAppObjectParams() *RestoreAppObjectParams {
	this := RestoreAppObjectParams{}
	return &this
}

// NewRestoreAppObjectParamsWithDefaults instantiates a new RestoreAppObjectParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestoreAppObjectParamsWithDefaults() *RestoreAppObjectParams {
	this := RestoreAppObjectParams{}
	return &this
}

// GetAdRestoreParams returns the AdRestoreParams field value if set, zero value otherwise.
func (o *RestoreAppObjectParams) GetAdRestoreParams() RestoreADAppObjectParams {
	if o == nil || o.AdRestoreParams == nil {
		var ret RestoreADAppObjectParams
		return ret
	}
	return *o.AdRestoreParams
}

// GetAdRestoreParamsOk returns a tuple with the AdRestoreParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreAppObjectParams) GetAdRestoreParamsOk() (*RestoreADAppObjectParams, bool) {
	if o == nil || o.AdRestoreParams == nil {
		return nil, false
	}
	return o.AdRestoreParams, true
}

// HasAdRestoreParams returns a boolean if a field has been set.
func (o *RestoreAppObjectParams) HasAdRestoreParams() bool {
	if o != nil && o.AdRestoreParams != nil {
		return true
	}

	return false
}

// SetAdRestoreParams gets a reference to the given RestoreADAppObjectParams and assigns it to the AdRestoreParams field.
func (o *RestoreAppObjectParams) SetAdRestoreParams(v RestoreADAppObjectParams) {
	o.AdRestoreParams = &v
}

// GetCloneTaskId returns the CloneTaskId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestoreAppObjectParams) GetCloneTaskId() int64 {
	if o == nil || o.CloneTaskId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CloneTaskId.Get()
}

// GetCloneTaskIdOk returns a tuple with the CloneTaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestoreAppObjectParams) GetCloneTaskIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CloneTaskId.Get(), o.CloneTaskId.IsSet()
}

// HasCloneTaskId returns a boolean if a field has been set.
func (o *RestoreAppObjectParams) HasCloneTaskId() bool {
	if o != nil && o.CloneTaskId.IsSet() {
		return true
	}

	return false
}

// SetCloneTaskId gets a reference to the given NullableInt64 and assigns it to the CloneTaskId field.
func (o *RestoreAppObjectParams) SetCloneTaskId(v int64) {
	o.CloneTaskId.Set(&v)
}
// SetCloneTaskIdNil sets the value for CloneTaskId to be an explicit nil
func (o *RestoreAppObjectParams) SetCloneTaskIdNil() {
	o.CloneTaskId.Set(nil)
}

// UnsetCloneTaskId ensures that no value is present for CloneTaskId, not even an explicit nil
func (o *RestoreAppObjectParams) UnsetCloneTaskId() {
	o.CloneTaskId.Unset()
}

// GetExchangeRestoreParams returns the ExchangeRestoreParams field value if set, zero value otherwise.
func (o *RestoreAppObjectParams) GetExchangeRestoreParams() RestoreExchangeParams {
	if o == nil || o.ExchangeRestoreParams == nil {
		var ret RestoreExchangeParams
		return ret
	}
	return *o.ExchangeRestoreParams
}

// GetExchangeRestoreParamsOk returns a tuple with the ExchangeRestoreParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreAppObjectParams) GetExchangeRestoreParamsOk() (*RestoreExchangeParams, bool) {
	if o == nil || o.ExchangeRestoreParams == nil {
		return nil, false
	}
	return o.ExchangeRestoreParams, true
}

// HasExchangeRestoreParams returns a boolean if a field has been set.
func (o *RestoreAppObjectParams) HasExchangeRestoreParams() bool {
	if o != nil && o.ExchangeRestoreParams != nil {
		return true
	}

	return false
}

// SetExchangeRestoreParams gets a reference to the given RestoreExchangeParams and assigns it to the ExchangeRestoreParams field.
func (o *RestoreAppObjectParams) SetExchangeRestoreParams(v RestoreExchangeParams) {
	o.ExchangeRestoreParams = &v
}

// GetOracleRestoreParams returns the OracleRestoreParams field value if set, zero value otherwise.
func (o *RestoreAppObjectParams) GetOracleRestoreParams() RestoreOracleAppObjectParams {
	if o == nil || o.OracleRestoreParams == nil {
		var ret RestoreOracleAppObjectParams
		return ret
	}
	return *o.OracleRestoreParams
}

// GetOracleRestoreParamsOk returns a tuple with the OracleRestoreParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreAppObjectParams) GetOracleRestoreParamsOk() (*RestoreOracleAppObjectParams, bool) {
	if o == nil || o.OracleRestoreParams == nil {
		return nil, false
	}
	return o.OracleRestoreParams, true
}

// HasOracleRestoreParams returns a boolean if a field has been set.
func (o *RestoreAppObjectParams) HasOracleRestoreParams() bool {
	if o != nil && o.OracleRestoreParams != nil {
		return true
	}

	return false
}

// SetOracleRestoreParams gets a reference to the given RestoreOracleAppObjectParams and assigns it to the OracleRestoreParams field.
func (o *RestoreAppObjectParams) SetOracleRestoreParams(v RestoreOracleAppObjectParams) {
	o.OracleRestoreParams = &v
}

// GetSqlRestoreParams returns the SqlRestoreParams field value if set, zero value otherwise.
func (o *RestoreAppObjectParams) GetSqlRestoreParams() RestoreSqlAppObjectParams {
	if o == nil || o.SqlRestoreParams == nil {
		var ret RestoreSqlAppObjectParams
		return ret
	}
	return *o.SqlRestoreParams
}

// GetSqlRestoreParamsOk returns a tuple with the SqlRestoreParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreAppObjectParams) GetSqlRestoreParamsOk() (*RestoreSqlAppObjectParams, bool) {
	if o == nil || o.SqlRestoreParams == nil {
		return nil, false
	}
	return o.SqlRestoreParams, true
}

// HasSqlRestoreParams returns a boolean if a field has been set.
func (o *RestoreAppObjectParams) HasSqlRestoreParams() bool {
	if o != nil && o.SqlRestoreParams != nil {
		return true
	}

	return false
}

// SetSqlRestoreParams gets a reference to the given RestoreSqlAppObjectParams and assigns it to the SqlRestoreParams field.
func (o *RestoreAppObjectParams) SetSqlRestoreParams(v RestoreSqlAppObjectParams) {
	o.SqlRestoreParams = &v
}

// GetTargetHost returns the TargetHost field value if set, zero value otherwise.
func (o *RestoreAppObjectParams) GetTargetHost() EntityProto {
	if o == nil || o.TargetHost == nil {
		var ret EntityProto
		return ret
	}
	return *o.TargetHost
}

// GetTargetHostOk returns a tuple with the TargetHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreAppObjectParams) GetTargetHostOk() (*EntityProto, bool) {
	if o == nil || o.TargetHost == nil {
		return nil, false
	}
	return o.TargetHost, true
}

// HasTargetHost returns a boolean if a field has been set.
func (o *RestoreAppObjectParams) HasTargetHost() bool {
	if o != nil && o.TargetHost != nil {
		return true
	}

	return false
}

// SetTargetHost gets a reference to the given EntityProto and assigns it to the TargetHost field.
func (o *RestoreAppObjectParams) SetTargetHost(v EntityProto) {
	o.TargetHost = &v
}

// GetTargetHostParentSource returns the TargetHostParentSource field value if set, zero value otherwise.
func (o *RestoreAppObjectParams) GetTargetHostParentSource() EntityProto {
	if o == nil || o.TargetHostParentSource == nil {
		var ret EntityProto
		return ret
	}
	return *o.TargetHostParentSource
}

// GetTargetHostParentSourceOk returns a tuple with the TargetHostParentSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreAppObjectParams) GetTargetHostParentSourceOk() (*EntityProto, bool) {
	if o == nil || o.TargetHostParentSource == nil {
		return nil, false
	}
	return o.TargetHostParentSource, true
}

// HasTargetHostParentSource returns a boolean if a field has been set.
func (o *RestoreAppObjectParams) HasTargetHostParentSource() bool {
	if o != nil && o.TargetHostParentSource != nil {
		return true
	}

	return false
}

// SetTargetHostParentSource gets a reference to the given EntityProto and assigns it to the TargetHostParentSource field.
func (o *RestoreAppObjectParams) SetTargetHostParentSource(v EntityProto) {
	o.TargetHostParentSource = &v
}

func (o RestoreAppObjectParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdRestoreParams != nil {
		toSerialize["adRestoreParams"] = o.AdRestoreParams
	}
	if o.CloneTaskId.IsSet() {
		toSerialize["cloneTaskId"] = o.CloneTaskId.Get()
	}
	if o.ExchangeRestoreParams != nil {
		toSerialize["exchangeRestoreParams"] = o.ExchangeRestoreParams
	}
	if o.OracleRestoreParams != nil {
		toSerialize["oracleRestoreParams"] = o.OracleRestoreParams
	}
	if o.SqlRestoreParams != nil {
		toSerialize["sqlRestoreParams"] = o.SqlRestoreParams
	}
	if o.TargetHost != nil {
		toSerialize["targetHost"] = o.TargetHost
	}
	if o.TargetHostParentSource != nil {
		toSerialize["targetHostParentSource"] = o.TargetHostParentSource
	}
	return json.Marshal(toSerialize)
}

type NullableRestoreAppObjectParams struct {
	value *RestoreAppObjectParams
	isSet bool
}

func (v NullableRestoreAppObjectParams) Get() *RestoreAppObjectParams {
	return v.value
}

func (v *NullableRestoreAppObjectParams) Set(val *RestoreAppObjectParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRestoreAppObjectParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRestoreAppObjectParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestoreAppObjectParams(val *RestoreAppObjectParams) *NullableRestoreAppObjectParams {
	return &NullableRestoreAppObjectParams{value: val, isSet: true}
}

func (v NullableRestoreAppObjectParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestoreAppObjectParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


