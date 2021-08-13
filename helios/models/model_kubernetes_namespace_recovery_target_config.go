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

// KubernetesNamespaceRecoveryTargetConfig Specifies the recovery target configuration of the Namespace recovery.
type KubernetesNamespaceRecoveryTargetConfig struct {
	// Specifies whether or not to recover the Namespaces to a different source than they were backed up from.
	RecoverToNewSource NullableBool `json:"recoverToNewSource"`
	// Specifies the new source configuration if a Kubernetes Namespace is being restored to a different source than the one from which it was protected.
	NewSourceConfig NullableKubernetesNamespaceRecoveryNewSourceConfig `json:"newSourceConfig,omitempty"`
}

// NewKubernetesNamespaceRecoveryTargetConfig instantiates a new KubernetesNamespaceRecoveryTargetConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesNamespaceRecoveryTargetConfig(recoverToNewSource NullableBool) *KubernetesNamespaceRecoveryTargetConfig {
	this := KubernetesNamespaceRecoveryTargetConfig{}
	this.RecoverToNewSource = recoverToNewSource
	return &this
}

// NewKubernetesNamespaceRecoveryTargetConfigWithDefaults instantiates a new KubernetesNamespaceRecoveryTargetConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesNamespaceRecoveryTargetConfigWithDefaults() *KubernetesNamespaceRecoveryTargetConfig {
	this := KubernetesNamespaceRecoveryTargetConfig{}
	return &this
}

// GetRecoverToNewSource returns the RecoverToNewSource field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *KubernetesNamespaceRecoveryTargetConfig) GetRecoverToNewSource() bool {
	if o == nil || o.RecoverToNewSource.Get() == nil {
		var ret bool
		return ret
	}

	return *o.RecoverToNewSource.Get()
}

// GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNamespaceRecoveryTargetConfig) GetRecoverToNewSourceOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RecoverToNewSource.Get(), o.RecoverToNewSource.IsSet()
}

// SetRecoverToNewSource sets field value
func (o *KubernetesNamespaceRecoveryTargetConfig) SetRecoverToNewSource(v bool) {
	o.RecoverToNewSource.Set(&v)
}

// GetNewSourceConfig returns the NewSourceConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesNamespaceRecoveryTargetConfig) GetNewSourceConfig() KubernetesNamespaceRecoveryNewSourceConfig {
	if o == nil || o.NewSourceConfig.Get() == nil {
		var ret KubernetesNamespaceRecoveryNewSourceConfig
		return ret
	}
	return *o.NewSourceConfig.Get()
}

// GetNewSourceConfigOk returns a tuple with the NewSourceConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNamespaceRecoveryTargetConfig) GetNewSourceConfigOk() (*KubernetesNamespaceRecoveryNewSourceConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NewSourceConfig.Get(), o.NewSourceConfig.IsSet()
}

// HasNewSourceConfig returns a boolean if a field has been set.
func (o *KubernetesNamespaceRecoveryTargetConfig) HasNewSourceConfig() bool {
	if o != nil && o.NewSourceConfig.IsSet() {
		return true
	}

	return false
}

// SetNewSourceConfig gets a reference to the given NullableKubernetesNamespaceRecoveryNewSourceConfig and assigns it to the NewSourceConfig field.
func (o *KubernetesNamespaceRecoveryTargetConfig) SetNewSourceConfig(v KubernetesNamespaceRecoveryNewSourceConfig) {
	o.NewSourceConfig.Set(&v)
}
// SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil
func (o *KubernetesNamespaceRecoveryTargetConfig) SetNewSourceConfigNil() {
	o.NewSourceConfig.Set(nil)
}

// UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
func (o *KubernetesNamespaceRecoveryTargetConfig) UnsetNewSourceConfig() {
	o.NewSourceConfig.Unset()
}

func (o KubernetesNamespaceRecoveryTargetConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["recoverToNewSource"] = o.RecoverToNewSource.Get()
	}
	if o.NewSourceConfig.IsSet() {
		toSerialize["newSourceConfig"] = o.NewSourceConfig.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableKubernetesNamespaceRecoveryTargetConfig struct {
	value *KubernetesNamespaceRecoveryTargetConfig
	isSet bool
}

func (v NullableKubernetesNamespaceRecoveryTargetConfig) Get() *KubernetesNamespaceRecoveryTargetConfig {
	return v.value
}

func (v *NullableKubernetesNamespaceRecoveryTargetConfig) Set(val *KubernetesNamespaceRecoveryTargetConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNamespaceRecoveryTargetConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNamespaceRecoveryTargetConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNamespaceRecoveryTargetConfig(val *KubernetesNamespaceRecoveryTargetConfig) *NullableKubernetesNamespaceRecoveryTargetConfig {
	return &NullableKubernetesNamespaceRecoveryTargetConfig{value: val, isSet: true}
}

func (v NullableKubernetesNamespaceRecoveryTargetConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNamespaceRecoveryTargetConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


