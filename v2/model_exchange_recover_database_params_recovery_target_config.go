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

// checks if the ExchangeRecoverDatabaseParamsRecoveryTargetConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExchangeRecoverDatabaseParamsRecoveryTargetConfig{}

// ExchangeRecoverDatabaseParamsRecoveryTargetConfig Specifies the recovery target configuration if recovery has to be done to a different location which is different from original source.
type ExchangeRecoverDatabaseParamsRecoveryTargetConfig struct {
	Source NullableExchangeDatabaseRecoveryTargetConfigSource `json:"source"`
}

type _ExchangeRecoverDatabaseParamsRecoveryTargetConfig ExchangeRecoverDatabaseParamsRecoveryTargetConfig

// NewExchangeRecoverDatabaseParamsRecoveryTargetConfig instantiates a new ExchangeRecoverDatabaseParamsRecoveryTargetConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeRecoverDatabaseParamsRecoveryTargetConfig(source NullableExchangeDatabaseRecoveryTargetConfigSource) *ExchangeRecoverDatabaseParamsRecoveryTargetConfig {
	this := ExchangeRecoverDatabaseParamsRecoveryTargetConfig{}
	this.Source = source
	return &this
}

// NewExchangeRecoverDatabaseParamsRecoveryTargetConfigWithDefaults instantiates a new ExchangeRecoverDatabaseParamsRecoveryTargetConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeRecoverDatabaseParamsRecoveryTargetConfigWithDefaults() *ExchangeRecoverDatabaseParamsRecoveryTargetConfig {
	this := ExchangeRecoverDatabaseParamsRecoveryTargetConfig{}
	return &this
}

// GetSource returns the Source field value
// If the value is explicit nil, the zero value for ExchangeDatabaseRecoveryTargetConfigSource will be returned
func (o *ExchangeRecoverDatabaseParamsRecoveryTargetConfig) GetSource() ExchangeDatabaseRecoveryTargetConfigSource {
	if o == nil || o.Source.Get() == nil {
		var ret ExchangeDatabaseRecoveryTargetConfigSource
		return ret
	}

	return *o.Source.Get()
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExchangeRecoverDatabaseParamsRecoveryTargetConfig) GetSourceOk() (*ExchangeDatabaseRecoveryTargetConfigSource, bool) {
	if o == nil {
		return nil, false
	}
	return o.Source.Get(), o.Source.IsSet()
}

// SetSource sets field value
func (o *ExchangeRecoverDatabaseParamsRecoveryTargetConfig) SetSource(v ExchangeDatabaseRecoveryTargetConfigSource) {
	o.Source.Set(&v)
}

func (o ExchangeRecoverDatabaseParamsRecoveryTargetConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeRecoverDatabaseParamsRecoveryTargetConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source"] = o.Source.Get()
	return toSerialize, nil
}

func (o *ExchangeRecoverDatabaseParamsRecoveryTargetConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source",
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

	varExchangeRecoverDatabaseParamsRecoveryTargetConfig := _ExchangeRecoverDatabaseParamsRecoveryTargetConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExchangeRecoverDatabaseParamsRecoveryTargetConfig)

	if err != nil {
		return err
	}

	*o = ExchangeRecoverDatabaseParamsRecoveryTargetConfig(varExchangeRecoverDatabaseParamsRecoveryTargetConfig)

	return err
}

type NullableExchangeRecoverDatabaseParamsRecoveryTargetConfig struct {
	value *ExchangeRecoverDatabaseParamsRecoveryTargetConfig
	isSet bool
}

func (v NullableExchangeRecoverDatabaseParamsRecoveryTargetConfig) Get() *ExchangeRecoverDatabaseParamsRecoveryTargetConfig {
	return v.value
}

func (v *NullableExchangeRecoverDatabaseParamsRecoveryTargetConfig) Set(val *ExchangeRecoverDatabaseParamsRecoveryTargetConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeRecoverDatabaseParamsRecoveryTargetConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeRecoverDatabaseParamsRecoveryTargetConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeRecoverDatabaseParamsRecoveryTargetConfig(val *ExchangeRecoverDatabaseParamsRecoveryTargetConfig) *NullableExchangeRecoverDatabaseParamsRecoveryTargetConfig {
	return &NullableExchangeRecoverDatabaseParamsRecoveryTargetConfig{value: val, isSet: true}
}

func (v NullableExchangeRecoverDatabaseParamsRecoveryTargetConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeRecoverDatabaseParamsRecoveryTargetConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


