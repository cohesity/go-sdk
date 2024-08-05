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

// checks if the HyperVTargetParamsForRecoverFileAndFolderNewTargetConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperVTargetParamsForRecoverFileAndFolderNewTargetConfig{}

// HyperVTargetParamsForRecoverFileAndFolderNewTargetConfig Specifies the configuration for recovering to a new target.
type HyperVTargetParamsForRecoverFileAndFolderNewTargetConfig struct {
	// Specifies the path location to recover files to.
	AbsolutePath NullableString `json:"absolutePath"`
	TargetVm NullableAcropolisRecoverFilesNewTargetConfigTargetVm `json:"targetVm"`
	TargetVmCredentials NullableAcropolisRecoverFilesNewTargetConfigTargetVmCredentials `json:"targetVmCredentials"`
}

type _HyperVTargetParamsForRecoverFileAndFolderNewTargetConfig HyperVTargetParamsForRecoverFileAndFolderNewTargetConfig

// NewHyperVTargetParamsForRecoverFileAndFolderNewTargetConfig instantiates a new HyperVTargetParamsForRecoverFileAndFolderNewTargetConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperVTargetParamsForRecoverFileAndFolderNewTargetConfig(absolutePath NullableString, targetVm NullableAcropolisRecoverFilesNewTargetConfigTargetVm, targetVmCredentials NullableAcropolisRecoverFilesNewTargetConfigTargetVmCredentials) *HyperVTargetParamsForRecoverFileAndFolderNewTargetConfig {
	this := HyperVTargetParamsForRecoverFileAndFolderNewTargetConfig{}
	this.AbsolutePath = absolutePath
	this.TargetVm = targetVm
	this.TargetVmCredentials = targetVmCredentials
	return &this
}

// NewHyperVTargetParamsForRecoverFileAndFolderNewTargetConfigWithDefaults instantiates a new HyperVTargetParamsForRecoverFileAndFolderNewTargetConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperVTargetParamsForRecoverFileAndFolderNewTargetConfigWithDefaults() *HyperVTargetParamsForRecoverFileAndFolderNewTargetConfig {
	this := HyperVTargetParamsForRecoverFileAndFolderNewTargetConfig{}
	return &this
}

// GetAbsolutePath returns the AbsolutePath field value
// If the value is explicit nil, the zero value for string will be returned
func (o *HyperVTargetParamsForRecoverFileAndFolderNewTargetConfig) GetAbsolutePath() string {
	if o == nil || o.AbsolutePath.Get() == nil {
		var ret string
		return ret
	}

	return *o.AbsolutePath.Get()
}

// GetAbsolutePathOk returns a tuple with the AbsolutePath field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperVTargetParamsForRecoverFileAndFolderNewTargetConfig) GetAbsolutePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AbsolutePath.Get(), o.AbsolutePath.IsSet()
}

// SetAbsolutePath sets field value
func (o *HyperVTargetParamsForRecoverFileAndFolderNewTargetConfig) SetAbsolutePath(v string) {
	o.AbsolutePath.Set(&v)
}

// GetTargetVm returns the TargetVm field value
// If the value is explicit nil, the zero value for AcropolisRecoverFilesNewTargetConfigTargetVm will be returned
func (o *HyperVTargetParamsForRecoverFileAndFolderNewTargetConfig) GetTargetVm() AcropolisRecoverFilesNewTargetConfigTargetVm {
	if o == nil || o.TargetVm.Get() == nil {
		var ret AcropolisRecoverFilesNewTargetConfigTargetVm
		return ret
	}

	return *o.TargetVm.Get()
}

// GetTargetVmOk returns a tuple with the TargetVm field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperVTargetParamsForRecoverFileAndFolderNewTargetConfig) GetTargetVmOk() (*AcropolisRecoverFilesNewTargetConfigTargetVm, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetVm.Get(), o.TargetVm.IsSet()
}

// SetTargetVm sets field value
func (o *HyperVTargetParamsForRecoverFileAndFolderNewTargetConfig) SetTargetVm(v AcropolisRecoverFilesNewTargetConfigTargetVm) {
	o.TargetVm.Set(&v)
}

// GetTargetVmCredentials returns the TargetVmCredentials field value
// If the value is explicit nil, the zero value for AcropolisRecoverFilesNewTargetConfigTargetVmCredentials will be returned
func (o *HyperVTargetParamsForRecoverFileAndFolderNewTargetConfig) GetTargetVmCredentials() AcropolisRecoverFilesNewTargetConfigTargetVmCredentials {
	if o == nil || o.TargetVmCredentials.Get() == nil {
		var ret AcropolisRecoverFilesNewTargetConfigTargetVmCredentials
		return ret
	}

	return *o.TargetVmCredentials.Get()
}

// GetTargetVmCredentialsOk returns a tuple with the TargetVmCredentials field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperVTargetParamsForRecoverFileAndFolderNewTargetConfig) GetTargetVmCredentialsOk() (*AcropolisRecoverFilesNewTargetConfigTargetVmCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetVmCredentials.Get(), o.TargetVmCredentials.IsSet()
}

// SetTargetVmCredentials sets field value
func (o *HyperVTargetParamsForRecoverFileAndFolderNewTargetConfig) SetTargetVmCredentials(v AcropolisRecoverFilesNewTargetConfigTargetVmCredentials) {
	o.TargetVmCredentials.Set(&v)
}

func (o HyperVTargetParamsForRecoverFileAndFolderNewTargetConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperVTargetParamsForRecoverFileAndFolderNewTargetConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["absolutePath"] = o.AbsolutePath.Get()
	toSerialize["targetVm"] = o.TargetVm.Get()
	toSerialize["targetVmCredentials"] = o.TargetVmCredentials.Get()
	return toSerialize, nil
}

func (o *HyperVTargetParamsForRecoverFileAndFolderNewTargetConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"absolutePath",
		"targetVm",
		"targetVmCredentials",
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

	varHyperVTargetParamsForRecoverFileAndFolderNewTargetConfig := _HyperVTargetParamsForRecoverFileAndFolderNewTargetConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHyperVTargetParamsForRecoverFileAndFolderNewTargetConfig)

	if err != nil {
		return err
	}

	*o = HyperVTargetParamsForRecoverFileAndFolderNewTargetConfig(varHyperVTargetParamsForRecoverFileAndFolderNewTargetConfig)

	return err
}

type NullableHyperVTargetParamsForRecoverFileAndFolderNewTargetConfig struct {
	value *HyperVTargetParamsForRecoverFileAndFolderNewTargetConfig
	isSet bool
}

func (v NullableHyperVTargetParamsForRecoverFileAndFolderNewTargetConfig) Get() *HyperVTargetParamsForRecoverFileAndFolderNewTargetConfig {
	return v.value
}

func (v *NullableHyperVTargetParamsForRecoverFileAndFolderNewTargetConfig) Set(val *HyperVTargetParamsForRecoverFileAndFolderNewTargetConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperVTargetParamsForRecoverFileAndFolderNewTargetConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperVTargetParamsForRecoverFileAndFolderNewTargetConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperVTargetParamsForRecoverFileAndFolderNewTargetConfig(val *HyperVTargetParamsForRecoverFileAndFolderNewTargetConfig) *NullableHyperVTargetParamsForRecoverFileAndFolderNewTargetConfig {
	return &NullableHyperVTargetParamsForRecoverFileAndFolderNewTargetConfig{value: val, isSet: true}
}

func (v NullableHyperVTargetParamsForRecoverFileAndFolderNewTargetConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperVTargetParamsForRecoverFileAndFolderNewTargetConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


