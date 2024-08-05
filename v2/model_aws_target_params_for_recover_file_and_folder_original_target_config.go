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

// checks if the AwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig{}

// AwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig Specifies the configuration for recovering to the original target.
type AwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig struct {
	// Specifies the alternate path location to recover files to.
	AlternatePath NullableString `json:"alternatePath,omitempty"`
	// Specifies whether to recover files and folders to the original path location. If false, alternatePath must be specified.
	RecoverToOriginalPath NullableBool `json:"recoverToOriginalPath"`
	TargetVmCredentials NullableAcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials `json:"targetVmCredentials,omitempty"`
}

type _AwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig AwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig

// NewAwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig instantiates a new AwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig(recoverToOriginalPath NullableBool) *AwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig {
	this := AwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig{}
	this.RecoverToOriginalPath = recoverToOriginalPath
	return &this
}

// NewAwsTargetParamsForRecoverFileAndFolderOriginalTargetConfigWithDefaults instantiates a new AwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsTargetParamsForRecoverFileAndFolderOriginalTargetConfigWithDefaults() *AwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig {
	this := AwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig{}
	return &this
}

// GetAlternatePath returns the AlternatePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig) GetAlternatePath() string {
	if o == nil || IsNil(o.AlternatePath.Get()) {
		var ret string
		return ret
	}
	return *o.AlternatePath.Get()
}

// GetAlternatePathOk returns a tuple with the AlternatePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig) GetAlternatePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlternatePath.Get(), o.AlternatePath.IsSet()
}

// HasAlternatePath returns a boolean if a field has been set.
func (o *AwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig) HasAlternatePath() bool {
	if o != nil && o.AlternatePath.IsSet() {
		return true
	}

	return false
}

// SetAlternatePath gets a reference to the given NullableString and assigns it to the AlternatePath field.
func (o *AwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig) SetAlternatePath(v string) {
	o.AlternatePath.Set(&v)
}
// SetAlternatePathNil sets the value for AlternatePath to be an explicit nil
func (o *AwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig) SetAlternatePathNil() {
	o.AlternatePath.Set(nil)
}

// UnsetAlternatePath ensures that no value is present for AlternatePath, not even an explicit nil
func (o *AwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig) UnsetAlternatePath() {
	o.AlternatePath.Unset()
}

// GetRecoverToOriginalPath returns the RecoverToOriginalPath field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *AwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig) GetRecoverToOriginalPath() bool {
	if o == nil || o.RecoverToOriginalPath.Get() == nil {
		var ret bool
		return ret
	}

	return *o.RecoverToOriginalPath.Get()
}

// GetRecoverToOriginalPathOk returns a tuple with the RecoverToOriginalPath field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig) GetRecoverToOriginalPathOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecoverToOriginalPath.Get(), o.RecoverToOriginalPath.IsSet()
}

// SetRecoverToOriginalPath sets field value
func (o *AwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig) SetRecoverToOriginalPath(v bool) {
	o.RecoverToOriginalPath.Set(&v)
}

// GetTargetVmCredentials returns the TargetVmCredentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig) GetTargetVmCredentials() AcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials {
	if o == nil || IsNil(o.TargetVmCredentials.Get()) {
		var ret AcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials
		return ret
	}
	return *o.TargetVmCredentials.Get()
}

// GetTargetVmCredentialsOk returns a tuple with the TargetVmCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig) GetTargetVmCredentialsOk() (*AcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetVmCredentials.Get(), o.TargetVmCredentials.IsSet()
}

// HasTargetVmCredentials returns a boolean if a field has been set.
func (o *AwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig) HasTargetVmCredentials() bool {
	if o != nil && o.TargetVmCredentials.IsSet() {
		return true
	}

	return false
}

// SetTargetVmCredentials gets a reference to the given NullableAcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials and assigns it to the TargetVmCredentials field.
func (o *AwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig) SetTargetVmCredentials(v AcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials) {
	o.TargetVmCredentials.Set(&v)
}
// SetTargetVmCredentialsNil sets the value for TargetVmCredentials to be an explicit nil
func (o *AwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig) SetTargetVmCredentialsNil() {
	o.TargetVmCredentials.Set(nil)
}

// UnsetTargetVmCredentials ensures that no value is present for TargetVmCredentials, not even an explicit nil
func (o *AwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig) UnsetTargetVmCredentials() {
	o.TargetVmCredentials.Unset()
}

func (o AwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AlternatePath.IsSet() {
		toSerialize["alternatePath"] = o.AlternatePath.Get()
	}
	toSerialize["recoverToOriginalPath"] = o.RecoverToOriginalPath.Get()
	if o.TargetVmCredentials.IsSet() {
		toSerialize["targetVmCredentials"] = o.TargetVmCredentials.Get()
	}
	return toSerialize, nil
}

func (o *AwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"recoverToOriginalPath",
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

	varAwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig := _AwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig)

	if err != nil {
		return err
	}

	*o = AwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig(varAwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig)

	return err
}

type NullableAwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig struct {
	value *AwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig
	isSet bool
}

func (v NullableAwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig) Get() *AwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig {
	return v.value
}

func (v *NullableAwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig) Set(val *AwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig(val *AwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig) *NullableAwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig {
	return &NullableAwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig{value: val, isSet: true}
}

func (v NullableAwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsTargetParamsForRecoverFileAndFolderOriginalTargetConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


