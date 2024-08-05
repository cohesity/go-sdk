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

// checks if the KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig{}

// KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig Specifies the configuration for recovering to a new target.
type KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig struct {
	// Specifies the path location to recover files to.
	AbsolutePath NullableString `json:"absolutePath"`
	TargetNamespace NullableKubernetesRecoverFilesNewTargetConfigTargetNamespace `json:"targetNamespace,omitempty"`
	TargetPvc NullableKubernetesRecoverFilesNewTargetConfigTargetPvc `json:"targetPvc"`
	TargetSource NullableKubernetesRecoverFilesNewTargetConfigTargetSource `json:"targetSource,omitempty"`
}

type _KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig

// NewKubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig instantiates a new KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig(absolutePath NullableString, targetPvc NullableKubernetesRecoverFilesNewTargetConfigTargetPvc) *KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig {
	this := KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig{}
	this.AbsolutePath = absolutePath
	this.TargetPvc = targetPvc
	return &this
}

// NewKubernetesTargetParamsForRecoverFileAndFolderNewTargetConfigWithDefaults instantiates a new KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesTargetParamsForRecoverFileAndFolderNewTargetConfigWithDefaults() *KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig {
	this := KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig{}
	return &this
}

// GetAbsolutePath returns the AbsolutePath field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig) GetAbsolutePath() string {
	if o == nil || o.AbsolutePath.Get() == nil {
		var ret string
		return ret
	}

	return *o.AbsolutePath.Get()
}

// GetAbsolutePathOk returns a tuple with the AbsolutePath field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig) GetAbsolutePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AbsolutePath.Get(), o.AbsolutePath.IsSet()
}

// SetAbsolutePath sets field value
func (o *KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig) SetAbsolutePath(v string) {
	o.AbsolutePath.Set(&v)
}

// GetTargetNamespace returns the TargetNamespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig) GetTargetNamespace() KubernetesRecoverFilesNewTargetConfigTargetNamespace {
	if o == nil || IsNil(o.TargetNamespace.Get()) {
		var ret KubernetesRecoverFilesNewTargetConfigTargetNamespace
		return ret
	}
	return *o.TargetNamespace.Get()
}

// GetTargetNamespaceOk returns a tuple with the TargetNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig) GetTargetNamespaceOk() (*KubernetesRecoverFilesNewTargetConfigTargetNamespace, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetNamespace.Get(), o.TargetNamespace.IsSet()
}

// HasTargetNamespace returns a boolean if a field has been set.
func (o *KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig) HasTargetNamespace() bool {
	if o != nil && o.TargetNamespace.IsSet() {
		return true
	}

	return false
}

// SetTargetNamespace gets a reference to the given NullableKubernetesRecoverFilesNewTargetConfigTargetNamespace and assigns it to the TargetNamespace field.
func (o *KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig) SetTargetNamespace(v KubernetesRecoverFilesNewTargetConfigTargetNamespace) {
	o.TargetNamespace.Set(&v)
}
// SetTargetNamespaceNil sets the value for TargetNamespace to be an explicit nil
func (o *KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig) SetTargetNamespaceNil() {
	o.TargetNamespace.Set(nil)
}

// UnsetTargetNamespace ensures that no value is present for TargetNamespace, not even an explicit nil
func (o *KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig) UnsetTargetNamespace() {
	o.TargetNamespace.Unset()
}

// GetTargetPvc returns the TargetPvc field value
// If the value is explicit nil, the zero value for KubernetesRecoverFilesNewTargetConfigTargetPvc will be returned
func (o *KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig) GetTargetPvc() KubernetesRecoverFilesNewTargetConfigTargetPvc {
	if o == nil || o.TargetPvc.Get() == nil {
		var ret KubernetesRecoverFilesNewTargetConfigTargetPvc
		return ret
	}

	return *o.TargetPvc.Get()
}

// GetTargetPvcOk returns a tuple with the TargetPvc field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig) GetTargetPvcOk() (*KubernetesRecoverFilesNewTargetConfigTargetPvc, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetPvc.Get(), o.TargetPvc.IsSet()
}

// SetTargetPvc sets field value
func (o *KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig) SetTargetPvc(v KubernetesRecoverFilesNewTargetConfigTargetPvc) {
	o.TargetPvc.Set(&v)
}

// GetTargetSource returns the TargetSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig) GetTargetSource() KubernetesRecoverFilesNewTargetConfigTargetSource {
	if o == nil || IsNil(o.TargetSource.Get()) {
		var ret KubernetesRecoverFilesNewTargetConfigTargetSource
		return ret
	}
	return *o.TargetSource.Get()
}

// GetTargetSourceOk returns a tuple with the TargetSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig) GetTargetSourceOk() (*KubernetesRecoverFilesNewTargetConfigTargetSource, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetSource.Get(), o.TargetSource.IsSet()
}

// HasTargetSource returns a boolean if a field has been set.
func (o *KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig) HasTargetSource() bool {
	if o != nil && o.TargetSource.IsSet() {
		return true
	}

	return false
}

// SetTargetSource gets a reference to the given NullableKubernetesRecoverFilesNewTargetConfigTargetSource and assigns it to the TargetSource field.
func (o *KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig) SetTargetSource(v KubernetesRecoverFilesNewTargetConfigTargetSource) {
	o.TargetSource.Set(&v)
}
// SetTargetSourceNil sets the value for TargetSource to be an explicit nil
func (o *KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig) SetTargetSourceNil() {
	o.TargetSource.Set(nil)
}

// UnsetTargetSource ensures that no value is present for TargetSource, not even an explicit nil
func (o *KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig) UnsetTargetSource() {
	o.TargetSource.Unset()
}

func (o KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["absolutePath"] = o.AbsolutePath.Get()
	if o.TargetNamespace.IsSet() {
		toSerialize["targetNamespace"] = o.TargetNamespace.Get()
	}
	toSerialize["targetPvc"] = o.TargetPvc.Get()
	if o.TargetSource.IsSet() {
		toSerialize["targetSource"] = o.TargetSource.Get()
	}
	return toSerialize, nil
}

func (o *KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"absolutePath",
		"targetPvc",
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

	varKubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig := _KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig)

	if err != nil {
		return err
	}

	*o = KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig(varKubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig)

	return err
}

type NullableKubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig struct {
	value *KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig
	isSet bool
}

func (v NullableKubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig) Get() *KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig {
	return v.value
}

func (v *NullableKubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig) Set(val *KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig(val *KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig) *NullableKubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig {
	return &NullableKubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig{value: val, isSet: true}
}

func (v NullableKubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


