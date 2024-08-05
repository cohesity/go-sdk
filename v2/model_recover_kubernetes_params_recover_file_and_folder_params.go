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

// checks if the RecoverKubernetesParamsRecoverFileAndFolderParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoverKubernetesParamsRecoverFileAndFolderParams{}

// RecoverKubernetesParamsRecoverFileAndFolderParams Specifies the parameters to perform a file and folder recovery.
type RecoverKubernetesParamsRecoverFileAndFolderParams struct {
	// Specifies the information about the files and folders to be recovered.
	FilesAndFolders []CommonRecoverFileAndFolderInfo `json:"filesAndFolders"`
	KubernetesTargetParams NullableRecoverKubernetesFileAndFolderParamsKubernetesTargetParams `json:"kubernetesTargetParams,omitempty"`
	// Specifies the environment of the recovery target. The corresponding params below must be filled out.
	TargetEnvironment string `json:"targetEnvironment"`
}

type _RecoverKubernetesParamsRecoverFileAndFolderParams RecoverKubernetesParamsRecoverFileAndFolderParams

// NewRecoverKubernetesParamsRecoverFileAndFolderParams instantiates a new RecoverKubernetesParamsRecoverFileAndFolderParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverKubernetesParamsRecoverFileAndFolderParams(filesAndFolders []CommonRecoverFileAndFolderInfo, targetEnvironment string) *RecoverKubernetesParamsRecoverFileAndFolderParams {
	this := RecoverKubernetesParamsRecoverFileAndFolderParams{}
	this.FilesAndFolders = filesAndFolders
	this.TargetEnvironment = targetEnvironment
	return &this
}

// NewRecoverKubernetesParamsRecoverFileAndFolderParamsWithDefaults instantiates a new RecoverKubernetesParamsRecoverFileAndFolderParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverKubernetesParamsRecoverFileAndFolderParamsWithDefaults() *RecoverKubernetesParamsRecoverFileAndFolderParams {
	this := RecoverKubernetesParamsRecoverFileAndFolderParams{}
	return &this
}

// GetFilesAndFolders returns the FilesAndFolders field value
// If the value is explicit nil, the zero value for []CommonRecoverFileAndFolderInfo will be returned
func (o *RecoverKubernetesParamsRecoverFileAndFolderParams) GetFilesAndFolders() []CommonRecoverFileAndFolderInfo {
	if o == nil {
		var ret []CommonRecoverFileAndFolderInfo
		return ret
	}

	return o.FilesAndFolders
}

// GetFilesAndFoldersOk returns a tuple with the FilesAndFolders field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverKubernetesParamsRecoverFileAndFolderParams) GetFilesAndFoldersOk() ([]CommonRecoverFileAndFolderInfo, bool) {
	if o == nil || IsNil(o.FilesAndFolders) {
		return nil, false
	}
	return o.FilesAndFolders, true
}

// SetFilesAndFolders sets field value
func (o *RecoverKubernetesParamsRecoverFileAndFolderParams) SetFilesAndFolders(v []CommonRecoverFileAndFolderInfo) {
	o.FilesAndFolders = v
}

// GetKubernetesTargetParams returns the KubernetesTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverKubernetesParamsRecoverFileAndFolderParams) GetKubernetesTargetParams() RecoverKubernetesFileAndFolderParamsKubernetesTargetParams {
	if o == nil || IsNil(o.KubernetesTargetParams.Get()) {
		var ret RecoverKubernetesFileAndFolderParamsKubernetesTargetParams
		return ret
	}
	return *o.KubernetesTargetParams.Get()
}

// GetKubernetesTargetParamsOk returns a tuple with the KubernetesTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverKubernetesParamsRecoverFileAndFolderParams) GetKubernetesTargetParamsOk() (*RecoverKubernetesFileAndFolderParamsKubernetesTargetParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.KubernetesTargetParams.Get(), o.KubernetesTargetParams.IsSet()
}

// HasKubernetesTargetParams returns a boolean if a field has been set.
func (o *RecoverKubernetesParamsRecoverFileAndFolderParams) HasKubernetesTargetParams() bool {
	if o != nil && o.KubernetesTargetParams.IsSet() {
		return true
	}

	return false
}

// SetKubernetesTargetParams gets a reference to the given NullableRecoverKubernetesFileAndFolderParamsKubernetesTargetParams and assigns it to the KubernetesTargetParams field.
func (o *RecoverKubernetesParamsRecoverFileAndFolderParams) SetKubernetesTargetParams(v RecoverKubernetesFileAndFolderParamsKubernetesTargetParams) {
	o.KubernetesTargetParams.Set(&v)
}
// SetKubernetesTargetParamsNil sets the value for KubernetesTargetParams to be an explicit nil
func (o *RecoverKubernetesParamsRecoverFileAndFolderParams) SetKubernetesTargetParamsNil() {
	o.KubernetesTargetParams.Set(nil)
}

// UnsetKubernetesTargetParams ensures that no value is present for KubernetesTargetParams, not even an explicit nil
func (o *RecoverKubernetesParamsRecoverFileAndFolderParams) UnsetKubernetesTargetParams() {
	o.KubernetesTargetParams.Unset()
}

// GetTargetEnvironment returns the TargetEnvironment field value
func (o *RecoverKubernetesParamsRecoverFileAndFolderParams) GetTargetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetEnvironment
}

// GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field value
// and a boolean to check if the value has been set.
func (o *RecoverKubernetesParamsRecoverFileAndFolderParams) GetTargetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetEnvironment, true
}

// SetTargetEnvironment sets field value
func (o *RecoverKubernetesParamsRecoverFileAndFolderParams) SetTargetEnvironment(v string) {
	o.TargetEnvironment = v
}

func (o RecoverKubernetesParamsRecoverFileAndFolderParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoverKubernetesParamsRecoverFileAndFolderParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FilesAndFolders != nil {
		toSerialize["filesAndFolders"] = o.FilesAndFolders
	}
	if o.KubernetesTargetParams.IsSet() {
		toSerialize["kubernetesTargetParams"] = o.KubernetesTargetParams.Get()
	}
	toSerialize["targetEnvironment"] = o.TargetEnvironment
	return toSerialize, nil
}

func (o *RecoverKubernetesParamsRecoverFileAndFolderParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"filesAndFolders",
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

	varRecoverKubernetesParamsRecoverFileAndFolderParams := _RecoverKubernetesParamsRecoverFileAndFolderParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRecoverKubernetesParamsRecoverFileAndFolderParams)

	if err != nil {
		return err
	}

	*o = RecoverKubernetesParamsRecoverFileAndFolderParams(varRecoverKubernetesParamsRecoverFileAndFolderParams)

	return err
}

type NullableRecoverKubernetesParamsRecoverFileAndFolderParams struct {
	value *RecoverKubernetesParamsRecoverFileAndFolderParams
	isSet bool
}

func (v NullableRecoverKubernetesParamsRecoverFileAndFolderParams) Get() *RecoverKubernetesParamsRecoverFileAndFolderParams {
	return v.value
}

func (v *NullableRecoverKubernetesParamsRecoverFileAndFolderParams) Set(val *RecoverKubernetesParamsRecoverFileAndFolderParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverKubernetesParamsRecoverFileAndFolderParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverKubernetesParamsRecoverFileAndFolderParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverKubernetesParamsRecoverFileAndFolderParams(val *RecoverKubernetesParamsRecoverFileAndFolderParams) *NullableRecoverKubernetesParamsRecoverFileAndFolderParams {
	return &NullableRecoverKubernetesParamsRecoverFileAndFolderParams{value: val, isSet: true}
}

func (v NullableRecoverKubernetesParamsRecoverFileAndFolderParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverKubernetesParamsRecoverFileAndFolderParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


