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

// checks if the KubernetesNamespaceRecoveryTargetConfigNewSourceConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesNamespaceRecoveryTargetConfigNewSourceConfig{}

// KubernetesNamespaceRecoveryTargetConfigNewSourceConfig Specifies the new source configuration if a Kubernetes Namespace is being restored to a different source than the one from which it was protected.
type KubernetesNamespaceRecoveryTargetConfigNewSourceConfig struct {
	Source NullableKubernetesNamespaceRecoveryNewSourceConfigSource `json:"source"`
}

type _KubernetesNamespaceRecoveryTargetConfigNewSourceConfig KubernetesNamespaceRecoveryTargetConfigNewSourceConfig

// NewKubernetesNamespaceRecoveryTargetConfigNewSourceConfig instantiates a new KubernetesNamespaceRecoveryTargetConfigNewSourceConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesNamespaceRecoveryTargetConfigNewSourceConfig(source NullableKubernetesNamespaceRecoveryNewSourceConfigSource) *KubernetesNamespaceRecoveryTargetConfigNewSourceConfig {
	this := KubernetesNamespaceRecoveryTargetConfigNewSourceConfig{}
	this.Source = source
	return &this
}

// NewKubernetesNamespaceRecoveryTargetConfigNewSourceConfigWithDefaults instantiates a new KubernetesNamespaceRecoveryTargetConfigNewSourceConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesNamespaceRecoveryTargetConfigNewSourceConfigWithDefaults() *KubernetesNamespaceRecoveryTargetConfigNewSourceConfig {
	this := KubernetesNamespaceRecoveryTargetConfigNewSourceConfig{}
	return &this
}

// GetSource returns the Source field value
// If the value is explicit nil, the zero value for KubernetesNamespaceRecoveryNewSourceConfigSource will be returned
func (o *KubernetesNamespaceRecoveryTargetConfigNewSourceConfig) GetSource() KubernetesNamespaceRecoveryNewSourceConfigSource {
	if o == nil || o.Source.Get() == nil {
		var ret KubernetesNamespaceRecoveryNewSourceConfigSource
		return ret
	}

	return *o.Source.Get()
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNamespaceRecoveryTargetConfigNewSourceConfig) GetSourceOk() (*KubernetesNamespaceRecoveryNewSourceConfigSource, bool) {
	if o == nil {
		return nil, false
	}
	return o.Source.Get(), o.Source.IsSet()
}

// SetSource sets field value
func (o *KubernetesNamespaceRecoveryTargetConfigNewSourceConfig) SetSource(v KubernetesNamespaceRecoveryNewSourceConfigSource) {
	o.Source.Set(&v)
}

func (o KubernetesNamespaceRecoveryTargetConfigNewSourceConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesNamespaceRecoveryTargetConfigNewSourceConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source"] = o.Source.Get()
	return toSerialize, nil
}

func (o *KubernetesNamespaceRecoveryTargetConfigNewSourceConfig) UnmarshalJSON(data []byte) (err error) {
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

	varKubernetesNamespaceRecoveryTargetConfigNewSourceConfig := _KubernetesNamespaceRecoveryTargetConfigNewSourceConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKubernetesNamespaceRecoveryTargetConfigNewSourceConfig)

	if err != nil {
		return err
	}

	*o = KubernetesNamespaceRecoveryTargetConfigNewSourceConfig(varKubernetesNamespaceRecoveryTargetConfigNewSourceConfig)

	return err
}

type NullableKubernetesNamespaceRecoveryTargetConfigNewSourceConfig struct {
	value *KubernetesNamespaceRecoveryTargetConfigNewSourceConfig
	isSet bool
}

func (v NullableKubernetesNamespaceRecoveryTargetConfigNewSourceConfig) Get() *KubernetesNamespaceRecoveryTargetConfigNewSourceConfig {
	return v.value
}

func (v *NullableKubernetesNamespaceRecoveryTargetConfigNewSourceConfig) Set(val *KubernetesNamespaceRecoveryTargetConfigNewSourceConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNamespaceRecoveryTargetConfigNewSourceConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNamespaceRecoveryTargetConfigNewSourceConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNamespaceRecoveryTargetConfigNewSourceConfig(val *KubernetesNamespaceRecoveryTargetConfigNewSourceConfig) *NullableKubernetesNamespaceRecoveryTargetConfigNewSourceConfig {
	return &NullableKubernetesNamespaceRecoveryTargetConfigNewSourceConfig{value: val, isSet: true}
}

func (v NullableKubernetesNamespaceRecoveryTargetConfigNewSourceConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNamespaceRecoveryTargetConfigNewSourceConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


