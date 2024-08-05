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

// checks if the RecoverAwsParamsRecoverFileAndFolderParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoverAwsParamsRecoverFileAndFolderParams{}

// RecoverAwsParamsRecoverFileAndFolderParams Specifies the parameters to recover files and folders.
type RecoverAwsParamsRecoverFileAndFolderParams struct {
	AwsTargetParams NullableRecoverAwsFileAndFolderParamsAwsTargetParams `json:"awsTargetParams,omitempty"`
	// Specifies the info about the files and folders to be recovered.
	FilesAndFolders []CommonRecoverFileAndFolderInfo `json:"filesAndFolders"`
	// Specifies the environment of the recovery target. The corresponding params below must be filled out.
	TargetEnvironment string `json:"targetEnvironment"`
}

type _RecoverAwsParamsRecoverFileAndFolderParams RecoverAwsParamsRecoverFileAndFolderParams

// NewRecoverAwsParamsRecoverFileAndFolderParams instantiates a new RecoverAwsParamsRecoverFileAndFolderParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverAwsParamsRecoverFileAndFolderParams(filesAndFolders []CommonRecoverFileAndFolderInfo, targetEnvironment string) *RecoverAwsParamsRecoverFileAndFolderParams {
	this := RecoverAwsParamsRecoverFileAndFolderParams{}
	this.FilesAndFolders = filesAndFolders
	this.TargetEnvironment = targetEnvironment
	return &this
}

// NewRecoverAwsParamsRecoverFileAndFolderParamsWithDefaults instantiates a new RecoverAwsParamsRecoverFileAndFolderParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverAwsParamsRecoverFileAndFolderParamsWithDefaults() *RecoverAwsParamsRecoverFileAndFolderParams {
	this := RecoverAwsParamsRecoverFileAndFolderParams{}
	return &this
}

// GetAwsTargetParams returns the AwsTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverAwsParamsRecoverFileAndFolderParams) GetAwsTargetParams() RecoverAwsFileAndFolderParamsAwsTargetParams {
	if o == nil || IsNil(o.AwsTargetParams.Get()) {
		var ret RecoverAwsFileAndFolderParamsAwsTargetParams
		return ret
	}
	return *o.AwsTargetParams.Get()
}

// GetAwsTargetParamsOk returns a tuple with the AwsTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverAwsParamsRecoverFileAndFolderParams) GetAwsTargetParamsOk() (*RecoverAwsFileAndFolderParamsAwsTargetParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.AwsTargetParams.Get(), o.AwsTargetParams.IsSet()
}

// HasAwsTargetParams returns a boolean if a field has been set.
func (o *RecoverAwsParamsRecoverFileAndFolderParams) HasAwsTargetParams() bool {
	if o != nil && o.AwsTargetParams.IsSet() {
		return true
	}

	return false
}

// SetAwsTargetParams gets a reference to the given NullableRecoverAwsFileAndFolderParamsAwsTargetParams and assigns it to the AwsTargetParams field.
func (o *RecoverAwsParamsRecoverFileAndFolderParams) SetAwsTargetParams(v RecoverAwsFileAndFolderParamsAwsTargetParams) {
	o.AwsTargetParams.Set(&v)
}
// SetAwsTargetParamsNil sets the value for AwsTargetParams to be an explicit nil
func (o *RecoverAwsParamsRecoverFileAndFolderParams) SetAwsTargetParamsNil() {
	o.AwsTargetParams.Set(nil)
}

// UnsetAwsTargetParams ensures that no value is present for AwsTargetParams, not even an explicit nil
func (o *RecoverAwsParamsRecoverFileAndFolderParams) UnsetAwsTargetParams() {
	o.AwsTargetParams.Unset()
}

// GetFilesAndFolders returns the FilesAndFolders field value
// If the value is explicit nil, the zero value for []CommonRecoverFileAndFolderInfo will be returned
func (o *RecoverAwsParamsRecoverFileAndFolderParams) GetFilesAndFolders() []CommonRecoverFileAndFolderInfo {
	if o == nil {
		var ret []CommonRecoverFileAndFolderInfo
		return ret
	}

	return o.FilesAndFolders
}

// GetFilesAndFoldersOk returns a tuple with the FilesAndFolders field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverAwsParamsRecoverFileAndFolderParams) GetFilesAndFoldersOk() ([]CommonRecoverFileAndFolderInfo, bool) {
	if o == nil || IsNil(o.FilesAndFolders) {
		return nil, false
	}
	return o.FilesAndFolders, true
}

// SetFilesAndFolders sets field value
func (o *RecoverAwsParamsRecoverFileAndFolderParams) SetFilesAndFolders(v []CommonRecoverFileAndFolderInfo) {
	o.FilesAndFolders = v
}

// GetTargetEnvironment returns the TargetEnvironment field value
func (o *RecoverAwsParamsRecoverFileAndFolderParams) GetTargetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetEnvironment
}

// GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field value
// and a boolean to check if the value has been set.
func (o *RecoverAwsParamsRecoverFileAndFolderParams) GetTargetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetEnvironment, true
}

// SetTargetEnvironment sets field value
func (o *RecoverAwsParamsRecoverFileAndFolderParams) SetTargetEnvironment(v string) {
	o.TargetEnvironment = v
}

func (o RecoverAwsParamsRecoverFileAndFolderParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoverAwsParamsRecoverFileAndFolderParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AwsTargetParams.IsSet() {
		toSerialize["awsTargetParams"] = o.AwsTargetParams.Get()
	}
	if o.FilesAndFolders != nil {
		toSerialize["filesAndFolders"] = o.FilesAndFolders
	}
	toSerialize["targetEnvironment"] = o.TargetEnvironment
	return toSerialize, nil
}

func (o *RecoverAwsParamsRecoverFileAndFolderParams) UnmarshalJSON(data []byte) (err error) {
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

	varRecoverAwsParamsRecoverFileAndFolderParams := _RecoverAwsParamsRecoverFileAndFolderParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRecoverAwsParamsRecoverFileAndFolderParams)

	if err != nil {
		return err
	}

	*o = RecoverAwsParamsRecoverFileAndFolderParams(varRecoverAwsParamsRecoverFileAndFolderParams)

	return err
}

type NullableRecoverAwsParamsRecoverFileAndFolderParams struct {
	value *RecoverAwsParamsRecoverFileAndFolderParams
	isSet bool
}

func (v NullableRecoverAwsParamsRecoverFileAndFolderParams) Get() *RecoverAwsParamsRecoverFileAndFolderParams {
	return v.value
}

func (v *NullableRecoverAwsParamsRecoverFileAndFolderParams) Set(val *RecoverAwsParamsRecoverFileAndFolderParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverAwsParamsRecoverFileAndFolderParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverAwsParamsRecoverFileAndFolderParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverAwsParamsRecoverFileAndFolderParams(val *RecoverAwsParamsRecoverFileAndFolderParams) *NullableRecoverAwsParamsRecoverFileAndFolderParams {
	return &NullableRecoverAwsParamsRecoverFileAndFolderParams{value: val, isSet: true}
}

func (v NullableRecoverAwsParamsRecoverFileAndFolderParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverAwsParamsRecoverFileAndFolderParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


