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

// checks if the MongodbParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MongodbParams{}

// MongodbParams Specifies the recovery options specific to MongoDB environment.
type MongodbParams struct {
	RecoverMongodbParams RecoverMongodbParams `json:"recoverMongodbParams"`
	// Specifies the type of recover action to be performed.
	RecoveryAction string `json:"recoveryAction"`
}

type _MongodbParams MongodbParams

// NewMongodbParams instantiates a new MongodbParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMongodbParams(recoverMongodbParams RecoverMongodbParams, recoveryAction string) *MongodbParams {
	this := MongodbParams{}
	this.RecoverMongodbParams = recoverMongodbParams
	this.RecoveryAction = recoveryAction
	return &this
}

// NewMongodbParamsWithDefaults instantiates a new MongodbParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMongodbParamsWithDefaults() *MongodbParams {
	this := MongodbParams{}
	return &this
}

// GetRecoverMongodbParams returns the RecoverMongodbParams field value
func (o *MongodbParams) GetRecoverMongodbParams() RecoverMongodbParams {
	if o == nil {
		var ret RecoverMongodbParams
		return ret
	}

	return o.RecoverMongodbParams
}

// GetRecoverMongodbParamsOk returns a tuple with the RecoverMongodbParams field value
// and a boolean to check if the value has been set.
func (o *MongodbParams) GetRecoverMongodbParamsOk() (*RecoverMongodbParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecoverMongodbParams, true
}

// SetRecoverMongodbParams sets field value
func (o *MongodbParams) SetRecoverMongodbParams(v RecoverMongodbParams) {
	o.RecoverMongodbParams = v
}

// GetRecoveryAction returns the RecoveryAction field value
func (o *MongodbParams) GetRecoveryAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecoveryAction
}

// GetRecoveryActionOk returns a tuple with the RecoveryAction field value
// and a boolean to check if the value has been set.
func (o *MongodbParams) GetRecoveryActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecoveryAction, true
}

// SetRecoveryAction sets field value
func (o *MongodbParams) SetRecoveryAction(v string) {
	o.RecoveryAction = v
}

func (o MongodbParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MongodbParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["recoverMongodbParams"] = o.RecoverMongodbParams
	toSerialize["recoveryAction"] = o.RecoveryAction
	return toSerialize, nil
}

func (o *MongodbParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"recoverMongodbParams",
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

	varMongodbParams := _MongodbParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMongodbParams)

	if err != nil {
		return err
	}

	*o = MongodbParams(varMongodbParams)

	return err
}

type NullableMongodbParams struct {
	value *MongodbParams
	isSet bool
}

func (v NullableMongodbParams) Get() *MongodbParams {
	return v.value
}

func (v *NullableMongodbParams) Set(val *MongodbParams) {
	v.value = val
	v.isSet = true
}

func (v NullableMongodbParams) IsSet() bool {
	return v.isSet
}

func (v *NullableMongodbParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMongodbParams(val *MongodbParams) *NullableMongodbParams {
	return &NullableMongodbParams{value: val, isSet: true}
}

func (v NullableMongodbParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMongodbParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


