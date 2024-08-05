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

// checks if the RecoverExchangeParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoverExchangeParams{}

// RecoverExchangeParams Specifies the recovery options specific to Exchange environment.
type RecoverExchangeParams struct {
	RecoverAppParams NullableRecoverExchangeParamsRecoverAppParams `json:"recoverAppParams,omitempty"`
	RecoverExchangeDbsParams NullableRecoverExchangeParamsRecoverExchangeDbsParams `json:"recoverExchangeDbsParams,omitempty"`
	// Specifies the type of recover action to be performed.
	RecoveryAction string `json:"recoveryAction"`
}

type _RecoverExchangeParams RecoverExchangeParams

// NewRecoverExchangeParams instantiates a new RecoverExchangeParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverExchangeParams(recoveryAction string) *RecoverExchangeParams {
	this := RecoverExchangeParams{}
	this.RecoveryAction = recoveryAction
	return &this
}

// NewRecoverExchangeParamsWithDefaults instantiates a new RecoverExchangeParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverExchangeParamsWithDefaults() *RecoverExchangeParams {
	this := RecoverExchangeParams{}
	return &this
}

// GetRecoverAppParams returns the RecoverAppParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverExchangeParams) GetRecoverAppParams() RecoverExchangeParamsRecoverAppParams {
	if o == nil || IsNil(o.RecoverAppParams.Get()) {
		var ret RecoverExchangeParamsRecoverAppParams
		return ret
	}
	return *o.RecoverAppParams.Get()
}

// GetRecoverAppParamsOk returns a tuple with the RecoverAppParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverExchangeParams) GetRecoverAppParamsOk() (*RecoverExchangeParamsRecoverAppParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecoverAppParams.Get(), o.RecoverAppParams.IsSet()
}

// HasRecoverAppParams returns a boolean if a field has been set.
func (o *RecoverExchangeParams) HasRecoverAppParams() bool {
	if o != nil && o.RecoverAppParams.IsSet() {
		return true
	}

	return false
}

// SetRecoverAppParams gets a reference to the given NullableRecoverExchangeParamsRecoverAppParams and assigns it to the RecoverAppParams field.
func (o *RecoverExchangeParams) SetRecoverAppParams(v RecoverExchangeParamsRecoverAppParams) {
	o.RecoverAppParams.Set(&v)
}
// SetRecoverAppParamsNil sets the value for RecoverAppParams to be an explicit nil
func (o *RecoverExchangeParams) SetRecoverAppParamsNil() {
	o.RecoverAppParams.Set(nil)
}

// UnsetRecoverAppParams ensures that no value is present for RecoverAppParams, not even an explicit nil
func (o *RecoverExchangeParams) UnsetRecoverAppParams() {
	o.RecoverAppParams.Unset()
}

// GetRecoverExchangeDbsParams returns the RecoverExchangeDbsParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverExchangeParams) GetRecoverExchangeDbsParams() RecoverExchangeParamsRecoverExchangeDbsParams {
	if o == nil || IsNil(o.RecoverExchangeDbsParams.Get()) {
		var ret RecoverExchangeParamsRecoverExchangeDbsParams
		return ret
	}
	return *o.RecoverExchangeDbsParams.Get()
}

// GetRecoverExchangeDbsParamsOk returns a tuple with the RecoverExchangeDbsParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverExchangeParams) GetRecoverExchangeDbsParamsOk() (*RecoverExchangeParamsRecoverExchangeDbsParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecoverExchangeDbsParams.Get(), o.RecoverExchangeDbsParams.IsSet()
}

// HasRecoverExchangeDbsParams returns a boolean if a field has been set.
func (o *RecoverExchangeParams) HasRecoverExchangeDbsParams() bool {
	if o != nil && o.RecoverExchangeDbsParams.IsSet() {
		return true
	}

	return false
}

// SetRecoverExchangeDbsParams gets a reference to the given NullableRecoverExchangeParamsRecoverExchangeDbsParams and assigns it to the RecoverExchangeDbsParams field.
func (o *RecoverExchangeParams) SetRecoverExchangeDbsParams(v RecoverExchangeParamsRecoverExchangeDbsParams) {
	o.RecoverExchangeDbsParams.Set(&v)
}
// SetRecoverExchangeDbsParamsNil sets the value for RecoverExchangeDbsParams to be an explicit nil
func (o *RecoverExchangeParams) SetRecoverExchangeDbsParamsNil() {
	o.RecoverExchangeDbsParams.Set(nil)
}

// UnsetRecoverExchangeDbsParams ensures that no value is present for RecoverExchangeDbsParams, not even an explicit nil
func (o *RecoverExchangeParams) UnsetRecoverExchangeDbsParams() {
	o.RecoverExchangeDbsParams.Unset()
}

// GetRecoveryAction returns the RecoveryAction field value
func (o *RecoverExchangeParams) GetRecoveryAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecoveryAction
}

// GetRecoveryActionOk returns a tuple with the RecoveryAction field value
// and a boolean to check if the value has been set.
func (o *RecoverExchangeParams) GetRecoveryActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecoveryAction, true
}

// SetRecoveryAction sets field value
func (o *RecoverExchangeParams) SetRecoveryAction(v string) {
	o.RecoveryAction = v
}

func (o RecoverExchangeParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoverExchangeParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RecoverAppParams.IsSet() {
		toSerialize["recoverAppParams"] = o.RecoverAppParams.Get()
	}
	if o.RecoverExchangeDbsParams.IsSet() {
		toSerialize["recoverExchangeDbsParams"] = o.RecoverExchangeDbsParams.Get()
	}
	toSerialize["recoveryAction"] = o.RecoveryAction
	return toSerialize, nil
}

func (o *RecoverExchangeParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varRecoverExchangeParams := _RecoverExchangeParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRecoverExchangeParams)

	if err != nil {
		return err
	}

	*o = RecoverExchangeParams(varRecoverExchangeParams)

	return err
}

type NullableRecoverExchangeParams struct {
	value *RecoverExchangeParams
	isSet bool
}

func (v NullableRecoverExchangeParams) Get() *RecoverExchangeParams {
	return v.value
}

func (v *NullableRecoverExchangeParams) Set(val *RecoverExchangeParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverExchangeParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverExchangeParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverExchangeParams(val *RecoverExchangeParams) *NullableRecoverExchangeParams {
	return &NullableRecoverExchangeParams{value: val, isSet: true}
}

func (v NullableRecoverExchangeParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverExchangeParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


