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

// checks if the CassandraParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CassandraParams{}

// CassandraParams Specifies the recovery options specific to Cassandra environment.
type CassandraParams struct {
	// Specifies whether to continue recovering other objects if one of Object failed to recover. Default value is false.
	ContinueOnError NullableBool `json:"continueOnError,omitempty"`
	// Specifies whether the current recovery operation is a multi-stage restore operation.
	IsMultiStageRestore NullableBool `json:"isMultiStageRestore,omitempty"`
	RecoverCassandraParams RecoverCassandraParams `json:"recoverCassandraParams"`
	// Specifies the type of recover action to be performed.
	RecoveryAction string `json:"recoveryAction"`
}

type _CassandraParams CassandraParams

// NewCassandraParams instantiates a new CassandraParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCassandraParams(recoverCassandraParams RecoverCassandraParams, recoveryAction string) *CassandraParams {
	this := CassandraParams{}
	this.RecoverCassandraParams = recoverCassandraParams
	this.RecoveryAction = recoveryAction
	return &this
}

// NewCassandraParamsWithDefaults instantiates a new CassandraParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCassandraParamsWithDefaults() *CassandraParams {
	this := CassandraParams{}
	return &this
}

// GetContinueOnError returns the ContinueOnError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CassandraParams) GetContinueOnError() bool {
	if o == nil || IsNil(o.ContinueOnError.Get()) {
		var ret bool
		return ret
	}
	return *o.ContinueOnError.Get()
}

// GetContinueOnErrorOk returns a tuple with the ContinueOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CassandraParams) GetContinueOnErrorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContinueOnError.Get(), o.ContinueOnError.IsSet()
}

// HasContinueOnError returns a boolean if a field has been set.
func (o *CassandraParams) HasContinueOnError() bool {
	if o != nil && o.ContinueOnError.IsSet() {
		return true
	}

	return false
}

// SetContinueOnError gets a reference to the given NullableBool and assigns it to the ContinueOnError field.
func (o *CassandraParams) SetContinueOnError(v bool) {
	o.ContinueOnError.Set(&v)
}
// SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil
func (o *CassandraParams) SetContinueOnErrorNil() {
	o.ContinueOnError.Set(nil)
}

// UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
func (o *CassandraParams) UnsetContinueOnError() {
	o.ContinueOnError.Unset()
}

// GetIsMultiStageRestore returns the IsMultiStageRestore field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CassandraParams) GetIsMultiStageRestore() bool {
	if o == nil || IsNil(o.IsMultiStageRestore.Get()) {
		var ret bool
		return ret
	}
	return *o.IsMultiStageRestore.Get()
}

// GetIsMultiStageRestoreOk returns a tuple with the IsMultiStageRestore field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CassandraParams) GetIsMultiStageRestoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsMultiStageRestore.Get(), o.IsMultiStageRestore.IsSet()
}

// HasIsMultiStageRestore returns a boolean if a field has been set.
func (o *CassandraParams) HasIsMultiStageRestore() bool {
	if o != nil && o.IsMultiStageRestore.IsSet() {
		return true
	}

	return false
}

// SetIsMultiStageRestore gets a reference to the given NullableBool and assigns it to the IsMultiStageRestore field.
func (o *CassandraParams) SetIsMultiStageRestore(v bool) {
	o.IsMultiStageRestore.Set(&v)
}
// SetIsMultiStageRestoreNil sets the value for IsMultiStageRestore to be an explicit nil
func (o *CassandraParams) SetIsMultiStageRestoreNil() {
	o.IsMultiStageRestore.Set(nil)
}

// UnsetIsMultiStageRestore ensures that no value is present for IsMultiStageRestore, not even an explicit nil
func (o *CassandraParams) UnsetIsMultiStageRestore() {
	o.IsMultiStageRestore.Unset()
}

// GetRecoverCassandraParams returns the RecoverCassandraParams field value
func (o *CassandraParams) GetRecoverCassandraParams() RecoverCassandraParams {
	if o == nil {
		var ret RecoverCassandraParams
		return ret
	}

	return o.RecoverCassandraParams
}

// GetRecoverCassandraParamsOk returns a tuple with the RecoverCassandraParams field value
// and a boolean to check if the value has been set.
func (o *CassandraParams) GetRecoverCassandraParamsOk() (*RecoverCassandraParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecoverCassandraParams, true
}

// SetRecoverCassandraParams sets field value
func (o *CassandraParams) SetRecoverCassandraParams(v RecoverCassandraParams) {
	o.RecoverCassandraParams = v
}

// GetRecoveryAction returns the RecoveryAction field value
func (o *CassandraParams) GetRecoveryAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecoveryAction
}

// GetRecoveryActionOk returns a tuple with the RecoveryAction field value
// and a boolean to check if the value has been set.
func (o *CassandraParams) GetRecoveryActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecoveryAction, true
}

// SetRecoveryAction sets field value
func (o *CassandraParams) SetRecoveryAction(v string) {
	o.RecoveryAction = v
}

func (o CassandraParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CassandraParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ContinueOnError.IsSet() {
		toSerialize["continueOnError"] = o.ContinueOnError.Get()
	}
	if o.IsMultiStageRestore.IsSet() {
		toSerialize["isMultiStageRestore"] = o.IsMultiStageRestore.Get()
	}
	toSerialize["recoverCassandraParams"] = o.RecoverCassandraParams
	toSerialize["recoveryAction"] = o.RecoveryAction
	return toSerialize, nil
}

func (o *CassandraParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"recoverCassandraParams",
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

	varCassandraParams := _CassandraParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCassandraParams)

	if err != nil {
		return err
	}

	*o = CassandraParams(varCassandraParams)

	return err
}

type NullableCassandraParams struct {
	value *CassandraParams
	isSet bool
}

func (v NullableCassandraParams) Get() *CassandraParams {
	return v.value
}

func (v *NullableCassandraParams) Set(val *CassandraParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCassandraParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCassandraParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCassandraParams(val *CassandraParams) *NullableCassandraParams {
	return &NullableCassandraParams{value: val, isSet: true}
}

func (v NullableCassandraParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCassandraParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


