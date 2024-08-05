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

// checks if the RecoverExchangeParamsRecoverExchangeDbsParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoverExchangeParamsRecoverExchangeDbsParams{}

// RecoverExchangeParamsRecoverExchangeDbsParams Specifies the parameters to recover Exchange databases.
type RecoverExchangeParamsRecoverExchangeDbsParams struct {
	ExchangeTargetParams NullableRecoverExchangeDbsParamsExchangeTargetParams `json:"exchangeTargetParams,omitempty"`
	// Specifies the parameter whether the recovery should be performed to a new or an existing target.
	RecoverToNewSource bool `json:"recoverToNewSource"`
	// Specifies the environment of the recovery target. The corresponding params below must be filled out.
	TargetEnvironment string `json:"targetEnvironment"`
}

type _RecoverExchangeParamsRecoverExchangeDbsParams RecoverExchangeParamsRecoverExchangeDbsParams

// NewRecoverExchangeParamsRecoverExchangeDbsParams instantiates a new RecoverExchangeParamsRecoverExchangeDbsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverExchangeParamsRecoverExchangeDbsParams(recoverToNewSource bool, targetEnvironment string) *RecoverExchangeParamsRecoverExchangeDbsParams {
	this := RecoverExchangeParamsRecoverExchangeDbsParams{}
	this.RecoverToNewSource = recoverToNewSource
	this.TargetEnvironment = targetEnvironment
	return &this
}

// NewRecoverExchangeParamsRecoverExchangeDbsParamsWithDefaults instantiates a new RecoverExchangeParamsRecoverExchangeDbsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverExchangeParamsRecoverExchangeDbsParamsWithDefaults() *RecoverExchangeParamsRecoverExchangeDbsParams {
	this := RecoverExchangeParamsRecoverExchangeDbsParams{}
	return &this
}

// GetExchangeTargetParams returns the ExchangeTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverExchangeParamsRecoverExchangeDbsParams) GetExchangeTargetParams() RecoverExchangeDbsParamsExchangeTargetParams {
	if o == nil || IsNil(o.ExchangeTargetParams.Get()) {
		var ret RecoverExchangeDbsParamsExchangeTargetParams
		return ret
	}
	return *o.ExchangeTargetParams.Get()
}

// GetExchangeTargetParamsOk returns a tuple with the ExchangeTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverExchangeParamsRecoverExchangeDbsParams) GetExchangeTargetParamsOk() (*RecoverExchangeDbsParamsExchangeTargetParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExchangeTargetParams.Get(), o.ExchangeTargetParams.IsSet()
}

// HasExchangeTargetParams returns a boolean if a field has been set.
func (o *RecoverExchangeParamsRecoverExchangeDbsParams) HasExchangeTargetParams() bool {
	if o != nil && o.ExchangeTargetParams.IsSet() {
		return true
	}

	return false
}

// SetExchangeTargetParams gets a reference to the given NullableRecoverExchangeDbsParamsExchangeTargetParams and assigns it to the ExchangeTargetParams field.
func (o *RecoverExchangeParamsRecoverExchangeDbsParams) SetExchangeTargetParams(v RecoverExchangeDbsParamsExchangeTargetParams) {
	o.ExchangeTargetParams.Set(&v)
}
// SetExchangeTargetParamsNil sets the value for ExchangeTargetParams to be an explicit nil
func (o *RecoverExchangeParamsRecoverExchangeDbsParams) SetExchangeTargetParamsNil() {
	o.ExchangeTargetParams.Set(nil)
}

// UnsetExchangeTargetParams ensures that no value is present for ExchangeTargetParams, not even an explicit nil
func (o *RecoverExchangeParamsRecoverExchangeDbsParams) UnsetExchangeTargetParams() {
	o.ExchangeTargetParams.Unset()
}

// GetRecoverToNewSource returns the RecoverToNewSource field value
func (o *RecoverExchangeParamsRecoverExchangeDbsParams) GetRecoverToNewSource() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RecoverToNewSource
}

// GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field value
// and a boolean to check if the value has been set.
func (o *RecoverExchangeParamsRecoverExchangeDbsParams) GetRecoverToNewSourceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecoverToNewSource, true
}

// SetRecoverToNewSource sets field value
func (o *RecoverExchangeParamsRecoverExchangeDbsParams) SetRecoverToNewSource(v bool) {
	o.RecoverToNewSource = v
}

// GetTargetEnvironment returns the TargetEnvironment field value
func (o *RecoverExchangeParamsRecoverExchangeDbsParams) GetTargetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetEnvironment
}

// GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field value
// and a boolean to check if the value has been set.
func (o *RecoverExchangeParamsRecoverExchangeDbsParams) GetTargetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetEnvironment, true
}

// SetTargetEnvironment sets field value
func (o *RecoverExchangeParamsRecoverExchangeDbsParams) SetTargetEnvironment(v string) {
	o.TargetEnvironment = v
}

func (o RecoverExchangeParamsRecoverExchangeDbsParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoverExchangeParamsRecoverExchangeDbsParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ExchangeTargetParams.IsSet() {
		toSerialize["exchangeTargetParams"] = o.ExchangeTargetParams.Get()
	}
	toSerialize["recoverToNewSource"] = o.RecoverToNewSource
	toSerialize["targetEnvironment"] = o.TargetEnvironment
	return toSerialize, nil
}

func (o *RecoverExchangeParamsRecoverExchangeDbsParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"recoverToNewSource",
		"targetEnvironment",
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

	varRecoverExchangeParamsRecoverExchangeDbsParams := _RecoverExchangeParamsRecoverExchangeDbsParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRecoverExchangeParamsRecoverExchangeDbsParams)

	if err != nil {
		return err
	}

	*o = RecoverExchangeParamsRecoverExchangeDbsParams(varRecoverExchangeParamsRecoverExchangeDbsParams)

	return err
}

type NullableRecoverExchangeParamsRecoverExchangeDbsParams struct {
	value *RecoverExchangeParamsRecoverExchangeDbsParams
	isSet bool
}

func (v NullableRecoverExchangeParamsRecoverExchangeDbsParams) Get() *RecoverExchangeParamsRecoverExchangeDbsParams {
	return v.value
}

func (v *NullableRecoverExchangeParamsRecoverExchangeDbsParams) Set(val *RecoverExchangeParamsRecoverExchangeDbsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverExchangeParamsRecoverExchangeDbsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverExchangeParamsRecoverExchangeDbsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverExchangeParamsRecoverExchangeDbsParams(val *RecoverExchangeParamsRecoverExchangeDbsParams) *NullableRecoverExchangeParamsRecoverExchangeDbsParams {
	return &NullableRecoverExchangeParamsRecoverExchangeDbsParams{value: val, isSet: true}
}

func (v NullableRecoverExchangeParamsRecoverExchangeDbsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverExchangeParamsRecoverExchangeDbsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


