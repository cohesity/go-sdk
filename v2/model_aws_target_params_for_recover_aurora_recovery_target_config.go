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

// checks if the AwsTargetParamsForRecoverAuroraRecoveryTargetConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsTargetParamsForRecoverAuroraRecoveryTargetConfig{}

// AwsTargetParamsForRecoverAuroraRecoveryTargetConfig Specifies the recovery target configuration if recovery has to be done to a different location which is different from original source or to original Source with different configuration. If not specified, then the recovery of the vms will be performed to original location with all configuration parameters retained.
type AwsTargetParamsForRecoverAuroraRecoveryTargetConfig struct {
	NewSourceConfig NullableAwsAuroraRecoveryTargetConfigNewSourceConfig `json:"newSourceConfig,omitempty"`
	// Specifies the parameter whether the recovery should be performed to a new or an existing Source Target.
	RecoverToNewSource bool `json:"recoverToNewSource"`
}

type _AwsTargetParamsForRecoverAuroraRecoveryTargetConfig AwsTargetParamsForRecoverAuroraRecoveryTargetConfig

// NewAwsTargetParamsForRecoverAuroraRecoveryTargetConfig instantiates a new AwsTargetParamsForRecoverAuroraRecoveryTargetConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsTargetParamsForRecoverAuroraRecoveryTargetConfig(recoverToNewSource bool) *AwsTargetParamsForRecoverAuroraRecoveryTargetConfig {
	this := AwsTargetParamsForRecoverAuroraRecoveryTargetConfig{}
	this.RecoverToNewSource = recoverToNewSource
	return &this
}

// NewAwsTargetParamsForRecoverAuroraRecoveryTargetConfigWithDefaults instantiates a new AwsTargetParamsForRecoverAuroraRecoveryTargetConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsTargetParamsForRecoverAuroraRecoveryTargetConfigWithDefaults() *AwsTargetParamsForRecoverAuroraRecoveryTargetConfig {
	this := AwsTargetParamsForRecoverAuroraRecoveryTargetConfig{}
	return &this
}

// GetNewSourceConfig returns the NewSourceConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsTargetParamsForRecoverAuroraRecoveryTargetConfig) GetNewSourceConfig() AwsAuroraRecoveryTargetConfigNewSourceConfig {
	if o == nil || IsNil(o.NewSourceConfig.Get()) {
		var ret AwsAuroraRecoveryTargetConfigNewSourceConfig
		return ret
	}
	return *o.NewSourceConfig.Get()
}

// GetNewSourceConfigOk returns a tuple with the NewSourceConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsTargetParamsForRecoverAuroraRecoveryTargetConfig) GetNewSourceConfigOk() (*AwsAuroraRecoveryTargetConfigNewSourceConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.NewSourceConfig.Get(), o.NewSourceConfig.IsSet()
}

// HasNewSourceConfig returns a boolean if a field has been set.
func (o *AwsTargetParamsForRecoverAuroraRecoveryTargetConfig) HasNewSourceConfig() bool {
	if o != nil && o.NewSourceConfig.IsSet() {
		return true
	}

	return false
}

// SetNewSourceConfig gets a reference to the given NullableAwsAuroraRecoveryTargetConfigNewSourceConfig and assigns it to the NewSourceConfig field.
func (o *AwsTargetParamsForRecoverAuroraRecoveryTargetConfig) SetNewSourceConfig(v AwsAuroraRecoveryTargetConfigNewSourceConfig) {
	o.NewSourceConfig.Set(&v)
}
// SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil
func (o *AwsTargetParamsForRecoverAuroraRecoveryTargetConfig) SetNewSourceConfigNil() {
	o.NewSourceConfig.Set(nil)
}

// UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
func (o *AwsTargetParamsForRecoverAuroraRecoveryTargetConfig) UnsetNewSourceConfig() {
	o.NewSourceConfig.Unset()
}

// GetRecoverToNewSource returns the RecoverToNewSource field value
func (o *AwsTargetParamsForRecoverAuroraRecoveryTargetConfig) GetRecoverToNewSource() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RecoverToNewSource
}

// GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field value
// and a boolean to check if the value has been set.
func (o *AwsTargetParamsForRecoverAuroraRecoveryTargetConfig) GetRecoverToNewSourceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecoverToNewSource, true
}

// SetRecoverToNewSource sets field value
func (o *AwsTargetParamsForRecoverAuroraRecoveryTargetConfig) SetRecoverToNewSource(v bool) {
	o.RecoverToNewSource = v
}

func (o AwsTargetParamsForRecoverAuroraRecoveryTargetConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsTargetParamsForRecoverAuroraRecoveryTargetConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.NewSourceConfig.IsSet() {
		toSerialize["newSourceConfig"] = o.NewSourceConfig.Get()
	}
	toSerialize["recoverToNewSource"] = o.RecoverToNewSource
	return toSerialize, nil
}

func (o *AwsTargetParamsForRecoverAuroraRecoveryTargetConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"recoverToNewSource",
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

	varAwsTargetParamsForRecoverAuroraRecoveryTargetConfig := _AwsTargetParamsForRecoverAuroraRecoveryTargetConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAwsTargetParamsForRecoverAuroraRecoveryTargetConfig)

	if err != nil {
		return err
	}

	*o = AwsTargetParamsForRecoverAuroraRecoveryTargetConfig(varAwsTargetParamsForRecoverAuroraRecoveryTargetConfig)

	return err
}

type NullableAwsTargetParamsForRecoverAuroraRecoveryTargetConfig struct {
	value *AwsTargetParamsForRecoverAuroraRecoveryTargetConfig
	isSet bool
}

func (v NullableAwsTargetParamsForRecoverAuroraRecoveryTargetConfig) Get() *AwsTargetParamsForRecoverAuroraRecoveryTargetConfig {
	return v.value
}

func (v *NullableAwsTargetParamsForRecoverAuroraRecoveryTargetConfig) Set(val *AwsTargetParamsForRecoverAuroraRecoveryTargetConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsTargetParamsForRecoverAuroraRecoveryTargetConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsTargetParamsForRecoverAuroraRecoveryTargetConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsTargetParamsForRecoverAuroraRecoveryTargetConfig(val *AwsTargetParamsForRecoverAuroraRecoveryTargetConfig) *NullableAwsTargetParamsForRecoverAuroraRecoveryTargetConfig {
	return &NullableAwsTargetParamsForRecoverAuroraRecoveryTargetConfig{value: val, isSet: true}
}

func (v NullableAwsTargetParamsForRecoverAuroraRecoveryTargetConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsTargetParamsForRecoverAuroraRecoveryTargetConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


