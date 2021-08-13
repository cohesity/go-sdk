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

// RecoverGcpFileAndFolderParams Specifies the parameters to recover files and folders.
type RecoverGcpFileAndFolderParams struct {
	// Specifies the info about the files and folders to be recovered.
	FilesAndFolders []CommonRecoverFileAndFolderInfo `json:"filesAndFolders"`
	// Specifies the environment of the recovery target. The corresponding params below must be filled out.
	TargetEnvironment string `json:"targetEnvironment"`
	// Specifies the parameters to recover to a GCP target.
	GcpTargetParams NullableGcpTargetParamsForRecoverFileAndFolder `json:"gcpTargetParams,omitempty"`
}

// NewRecoverGcpFileAndFolderParams instantiates a new RecoverGcpFileAndFolderParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverGcpFileAndFolderParams(filesAndFolders []CommonRecoverFileAndFolderInfo, targetEnvironment string) *RecoverGcpFileAndFolderParams {
	this := RecoverGcpFileAndFolderParams{}
	this.FilesAndFolders = filesAndFolders
	this.TargetEnvironment = targetEnvironment
	return &this
}

// NewRecoverGcpFileAndFolderParamsWithDefaults instantiates a new RecoverGcpFileAndFolderParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverGcpFileAndFolderParamsWithDefaults() *RecoverGcpFileAndFolderParams {
	this := RecoverGcpFileAndFolderParams{}
	return &this
}

// GetFilesAndFolders returns the FilesAndFolders field value
// If the value is explicit nil, the zero value for []CommonRecoverFileAndFolderInfo will be returned
func (o *RecoverGcpFileAndFolderParams) GetFilesAndFolders() []CommonRecoverFileAndFolderInfo {
	if o == nil {
		var ret []CommonRecoverFileAndFolderInfo
		return ret
	}

	return o.FilesAndFolders
}

// GetFilesAndFoldersOk returns a tuple with the FilesAndFolders field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverGcpFileAndFolderParams) GetFilesAndFoldersOk() (*[]CommonRecoverFileAndFolderInfo, bool) {
	if o == nil || o.FilesAndFolders == nil {
		return nil, false
	}
	return &o.FilesAndFolders, true
}

// SetFilesAndFolders sets field value
func (o *RecoverGcpFileAndFolderParams) SetFilesAndFolders(v []CommonRecoverFileAndFolderInfo) {
	o.FilesAndFolders = v
}

// GetTargetEnvironment returns the TargetEnvironment field value
func (o *RecoverGcpFileAndFolderParams) GetTargetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetEnvironment
}

// GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field value
// and a boolean to check if the value has been set.
func (o *RecoverGcpFileAndFolderParams) GetTargetEnvironmentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TargetEnvironment, true
}

// SetTargetEnvironment sets field value
func (o *RecoverGcpFileAndFolderParams) SetTargetEnvironment(v string) {
	o.TargetEnvironment = v
}

// GetGcpTargetParams returns the GcpTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverGcpFileAndFolderParams) GetGcpTargetParams() GcpTargetParamsForRecoverFileAndFolder {
	if o == nil || o.GcpTargetParams.Get() == nil {
		var ret GcpTargetParamsForRecoverFileAndFolder
		return ret
	}
	return *o.GcpTargetParams.Get()
}

// GetGcpTargetParamsOk returns a tuple with the GcpTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverGcpFileAndFolderParams) GetGcpTargetParamsOk() (*GcpTargetParamsForRecoverFileAndFolder, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GcpTargetParams.Get(), o.GcpTargetParams.IsSet()
}

// HasGcpTargetParams returns a boolean if a field has been set.
func (o *RecoverGcpFileAndFolderParams) HasGcpTargetParams() bool {
	if o != nil && o.GcpTargetParams.IsSet() {
		return true
	}

	return false
}

// SetGcpTargetParams gets a reference to the given NullableGcpTargetParamsForRecoverFileAndFolder and assigns it to the GcpTargetParams field.
func (o *RecoverGcpFileAndFolderParams) SetGcpTargetParams(v GcpTargetParamsForRecoverFileAndFolder) {
	o.GcpTargetParams.Set(&v)
}
// SetGcpTargetParamsNil sets the value for GcpTargetParams to be an explicit nil
func (o *RecoverGcpFileAndFolderParams) SetGcpTargetParamsNil() {
	o.GcpTargetParams.Set(nil)
}

// UnsetGcpTargetParams ensures that no value is present for GcpTargetParams, not even an explicit nil
func (o *RecoverGcpFileAndFolderParams) UnsetGcpTargetParams() {
	o.GcpTargetParams.Unset()
}

func (o RecoverGcpFileAndFolderParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FilesAndFolders != nil {
		toSerialize["filesAndFolders"] = o.FilesAndFolders
	}
	if true {
		toSerialize["targetEnvironment"] = o.TargetEnvironment
	}
	if o.GcpTargetParams.IsSet() {
		toSerialize["gcpTargetParams"] = o.GcpTargetParams.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRecoverGcpFileAndFolderParams struct {
	value *RecoverGcpFileAndFolderParams
	isSet bool
}

func (v NullableRecoverGcpFileAndFolderParams) Get() *RecoverGcpFileAndFolderParams {
	return v.value
}

func (v *NullableRecoverGcpFileAndFolderParams) Set(val *RecoverGcpFileAndFolderParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverGcpFileAndFolderParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverGcpFileAndFolderParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverGcpFileAndFolderParams(val *RecoverGcpFileAndFolderParams) *NullableRecoverGcpFileAndFolderParams {
	return &NullableRecoverGcpFileAndFolderParams{value: val, isSet: true}
}

func (v NullableRecoverGcpFileAndFolderParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverGcpFileAndFolderParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


