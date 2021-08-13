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

// CommonStandbyResourceParams Specifies the params for standby resource.
type CommonStandbyResourceParams struct {
	// Specifies the recovery point objective time user expects for this standby resource.
	RecoveryPointObjectiveSecs NullableInt64 `json:"recoveryPointObjectiveSecs,omitempty"`
}

// NewCommonStandbyResourceParams instantiates a new CommonStandbyResourceParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonStandbyResourceParams() *CommonStandbyResourceParams {
	this := CommonStandbyResourceParams{}
	return &this
}

// NewCommonStandbyResourceParamsWithDefaults instantiates a new CommonStandbyResourceParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonStandbyResourceParamsWithDefaults() *CommonStandbyResourceParams {
	this := CommonStandbyResourceParams{}
	return &this
}

// GetRecoveryPointObjectiveSecs returns the RecoveryPointObjectiveSecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonStandbyResourceParams) GetRecoveryPointObjectiveSecs() int64 {
	if o == nil || o.RecoveryPointObjectiveSecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RecoveryPointObjectiveSecs.Get()
}

// GetRecoveryPointObjectiveSecsOk returns a tuple with the RecoveryPointObjectiveSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonStandbyResourceParams) GetRecoveryPointObjectiveSecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RecoveryPointObjectiveSecs.Get(), o.RecoveryPointObjectiveSecs.IsSet()
}

// HasRecoveryPointObjectiveSecs returns a boolean if a field has been set.
func (o *CommonStandbyResourceParams) HasRecoveryPointObjectiveSecs() bool {
	if o != nil && o.RecoveryPointObjectiveSecs.IsSet() {
		return true
	}

	return false
}

// SetRecoveryPointObjectiveSecs gets a reference to the given NullableInt64 and assigns it to the RecoveryPointObjectiveSecs field.
func (o *CommonStandbyResourceParams) SetRecoveryPointObjectiveSecs(v int64) {
	o.RecoveryPointObjectiveSecs.Set(&v)
}
// SetRecoveryPointObjectiveSecsNil sets the value for RecoveryPointObjectiveSecs to be an explicit nil
func (o *CommonStandbyResourceParams) SetRecoveryPointObjectiveSecsNil() {
	o.RecoveryPointObjectiveSecs.Set(nil)
}

// UnsetRecoveryPointObjectiveSecs ensures that no value is present for RecoveryPointObjectiveSecs, not even an explicit nil
func (o *CommonStandbyResourceParams) UnsetRecoveryPointObjectiveSecs() {
	o.RecoveryPointObjectiveSecs.Unset()
}

func (o CommonStandbyResourceParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RecoveryPointObjectiveSecs.IsSet() {
		toSerialize["recoveryPointObjectiveSecs"] = o.RecoveryPointObjectiveSecs.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCommonStandbyResourceParams struct {
	value *CommonStandbyResourceParams
	isSet bool
}

func (v NullableCommonStandbyResourceParams) Get() *CommonStandbyResourceParams {
	return v.value
}

func (v *NullableCommonStandbyResourceParams) Set(val *CommonStandbyResourceParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonStandbyResourceParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonStandbyResourceParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonStandbyResourceParams(val *CommonStandbyResourceParams) *NullableCommonStandbyResourceParams {
	return &NullableCommonStandbyResourceParams{value: val, isSet: true}
}

func (v NullableCommonStandbyResourceParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonStandbyResourceParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


