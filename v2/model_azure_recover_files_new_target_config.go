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

// checks if the AzureRecoverFilesNewTargetConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureRecoverFilesNewTargetConfig{}

// AzureRecoverFilesNewTargetConfig Specifies the configuration for recovering files and folders to the new target.
type AzureRecoverFilesNewTargetConfig struct {
	// Specifies the absolute path location to recover files to.
	AbsolutePath NullableString `json:"absolutePath"`
	TargetVm NullableAcropolisRecoverFilesNewTargetConfigTargetVm `json:"targetVm"`
	TargetVmCredentials NullableAcropolisRecoverFilesNewTargetConfigTargetVmCredentials `json:"targetVmCredentials,omitempty"`
}

type _AzureRecoverFilesNewTargetConfig AzureRecoverFilesNewTargetConfig

// NewAzureRecoverFilesNewTargetConfig instantiates a new AzureRecoverFilesNewTargetConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureRecoverFilesNewTargetConfig(absolutePath NullableString, targetVm NullableAcropolisRecoverFilesNewTargetConfigTargetVm) *AzureRecoverFilesNewTargetConfig {
	this := AzureRecoverFilesNewTargetConfig{}
	this.AbsolutePath = absolutePath
	this.TargetVm = targetVm
	return &this
}

// NewAzureRecoverFilesNewTargetConfigWithDefaults instantiates a new AzureRecoverFilesNewTargetConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureRecoverFilesNewTargetConfigWithDefaults() *AzureRecoverFilesNewTargetConfig {
	this := AzureRecoverFilesNewTargetConfig{}
	return &this
}

// GetAbsolutePath returns the AbsolutePath field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AzureRecoverFilesNewTargetConfig) GetAbsolutePath() string {
	if o == nil || o.AbsolutePath.Get() == nil {
		var ret string
		return ret
	}

	return *o.AbsolutePath.Get()
}

// GetAbsolutePathOk returns a tuple with the AbsolutePath field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureRecoverFilesNewTargetConfig) GetAbsolutePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AbsolutePath.Get(), o.AbsolutePath.IsSet()
}

// SetAbsolutePath sets field value
func (o *AzureRecoverFilesNewTargetConfig) SetAbsolutePath(v string) {
	o.AbsolutePath.Set(&v)
}

// GetTargetVm returns the TargetVm field value
// If the value is explicit nil, the zero value for AcropolisRecoverFilesNewTargetConfigTargetVm will be returned
func (o *AzureRecoverFilesNewTargetConfig) GetTargetVm() AcropolisRecoverFilesNewTargetConfigTargetVm {
	if o == nil || o.TargetVm.Get() == nil {
		var ret AcropolisRecoverFilesNewTargetConfigTargetVm
		return ret
	}

	return *o.TargetVm.Get()
}

// GetTargetVmOk returns a tuple with the TargetVm field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureRecoverFilesNewTargetConfig) GetTargetVmOk() (*AcropolisRecoverFilesNewTargetConfigTargetVm, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetVm.Get(), o.TargetVm.IsSet()
}

// SetTargetVm sets field value
func (o *AzureRecoverFilesNewTargetConfig) SetTargetVm(v AcropolisRecoverFilesNewTargetConfigTargetVm) {
	o.TargetVm.Set(&v)
}

// GetTargetVmCredentials returns the TargetVmCredentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureRecoverFilesNewTargetConfig) GetTargetVmCredentials() AcropolisRecoverFilesNewTargetConfigTargetVmCredentials {
	if o == nil || IsNil(o.TargetVmCredentials.Get()) {
		var ret AcropolisRecoverFilesNewTargetConfigTargetVmCredentials
		return ret
	}
	return *o.TargetVmCredentials.Get()
}

// GetTargetVmCredentialsOk returns a tuple with the TargetVmCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureRecoverFilesNewTargetConfig) GetTargetVmCredentialsOk() (*AcropolisRecoverFilesNewTargetConfigTargetVmCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetVmCredentials.Get(), o.TargetVmCredentials.IsSet()
}

// HasTargetVmCredentials returns a boolean if a field has been set.
func (o *AzureRecoverFilesNewTargetConfig) HasTargetVmCredentials() bool {
	if o != nil && o.TargetVmCredentials.IsSet() {
		return true
	}

	return false
}

// SetTargetVmCredentials gets a reference to the given NullableAcropolisRecoverFilesNewTargetConfigTargetVmCredentials and assigns it to the TargetVmCredentials field.
func (o *AzureRecoverFilesNewTargetConfig) SetTargetVmCredentials(v AcropolisRecoverFilesNewTargetConfigTargetVmCredentials) {
	o.TargetVmCredentials.Set(&v)
}
// SetTargetVmCredentialsNil sets the value for TargetVmCredentials to be an explicit nil
func (o *AzureRecoverFilesNewTargetConfig) SetTargetVmCredentialsNil() {
	o.TargetVmCredentials.Set(nil)
}

// UnsetTargetVmCredentials ensures that no value is present for TargetVmCredentials, not even an explicit nil
func (o *AzureRecoverFilesNewTargetConfig) UnsetTargetVmCredentials() {
	o.TargetVmCredentials.Unset()
}

func (o AzureRecoverFilesNewTargetConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureRecoverFilesNewTargetConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["absolutePath"] = o.AbsolutePath.Get()
	toSerialize["targetVm"] = o.TargetVm.Get()
	if o.TargetVmCredentials.IsSet() {
		toSerialize["targetVmCredentials"] = o.TargetVmCredentials.Get()
	}
	return toSerialize, nil
}

func (o *AzureRecoverFilesNewTargetConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"absolutePath",
		"targetVm",
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

	varAzureRecoverFilesNewTargetConfig := _AzureRecoverFilesNewTargetConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAzureRecoverFilesNewTargetConfig)

	if err != nil {
		return err
	}

	*o = AzureRecoverFilesNewTargetConfig(varAzureRecoverFilesNewTargetConfig)

	return err
}

type NullableAzureRecoverFilesNewTargetConfig struct {
	value *AzureRecoverFilesNewTargetConfig
	isSet bool
}

func (v NullableAzureRecoverFilesNewTargetConfig) Get() *AzureRecoverFilesNewTargetConfig {
	return v.value
}

func (v *NullableAzureRecoverFilesNewTargetConfig) Set(val *AzureRecoverFilesNewTargetConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureRecoverFilesNewTargetConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureRecoverFilesNewTargetConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureRecoverFilesNewTargetConfig(val *AzureRecoverFilesNewTargetConfig) *NullableAzureRecoverFilesNewTargetConfig {
	return &NullableAzureRecoverFilesNewTargetConfig{value: val, isSet: true}
}

func (v NullableAzureRecoverFilesNewTargetConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureRecoverFilesNewTargetConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


