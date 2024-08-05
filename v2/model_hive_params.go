/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the HiveParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HiveParams{}

// HiveParams Specifies the recovery options specific to Hive environment.
type HiveParams struct {
	RecoverHiveParams RecoverHiveParams `json:"recoverHiveParams"`
	// Specifies the type of recover action to be performed.
	RecoveryAction string `json:"recoveryAction"`
}

type _HiveParams HiveParams

// NewHiveParams instantiates a new HiveParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHiveParams(recoverHiveParams RecoverHiveParams, recoveryAction string) *HiveParams {
	this := HiveParams{}
	this.RecoverHiveParams = recoverHiveParams
	this.RecoveryAction = recoveryAction
	return &this
}

// NewHiveParamsWithDefaults instantiates a new HiveParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHiveParamsWithDefaults() *HiveParams {
	this := HiveParams{}
	return &this
}

// GetRecoverHiveParams returns the RecoverHiveParams field value
func (o *HiveParams) GetRecoverHiveParams() RecoverHiveParams {
	if o == nil {
		var ret RecoverHiveParams
		return ret
	}

	return o.RecoverHiveParams
}

// GetRecoverHiveParamsOk returns a tuple with the RecoverHiveParams field value
// and a boolean to check if the value has been set.
func (o *HiveParams) GetRecoverHiveParamsOk() (*RecoverHiveParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecoverHiveParams, true
}

// SetRecoverHiveParams sets field value
func (o *HiveParams) SetRecoverHiveParams(v RecoverHiveParams) {
	o.RecoverHiveParams = v
}

// GetRecoveryAction returns the RecoveryAction field value
func (o *HiveParams) GetRecoveryAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecoveryAction
}

// GetRecoveryActionOk returns a tuple with the RecoveryAction field value
// and a boolean to check if the value has been set.
func (o *HiveParams) GetRecoveryActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecoveryAction, true
}

// SetRecoveryAction sets field value
func (o *HiveParams) SetRecoveryAction(v string) {
	o.RecoveryAction = v
}

func (o HiveParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HiveParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["recoverHiveParams"] = o.RecoverHiveParams
	toSerialize["recoveryAction"] = o.RecoveryAction
	return toSerialize, nil
}

func (o *HiveParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"recoverHiveParams",
		"recoveryAction",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varHiveParams := _HiveParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHiveParams)

	if err != nil {
		return err
	}

	*o = HiveParams(varHiveParams)

	return err
}

type NullableHiveParams struct {
	value *HiveParams
	isSet bool
}

func (v NullableHiveParams) Get() *HiveParams {
	return v.value
}

func (v *NullableHiveParams) Set(val *HiveParams) {
	v.value = val
	v.isSet = true
}

func (v NullableHiveParams) IsSet() bool {
	return v.isSet
}

func (v *NullableHiveParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHiveParams(val *HiveParams) *NullableHiveParams {
	return &NullableHiveParams{value: val, isSet: true}
}

func (v NullableHiveParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHiveParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


