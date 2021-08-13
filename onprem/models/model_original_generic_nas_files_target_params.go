/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package onprem
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/onprem/utils"
	"fmt"
)

var _ = NullableBool{}

// OriginalGenericNasFilesTargetParams Specifies the params of the original Generic Nas recovery target.
type OriginalGenericNasFilesTargetParams struct {
	// Specifies whether to recover files and folders to the original path location. If false, alternatePath must be specified.
	RecoverToOriginalPath NullableBool `json:"recoverToOriginalPath"`
	// Specifies the alternate path location to recover files to.
	AlternatePath NullableString `json:"alternatePath,omitempty"`
	// Specifies whether to overwrite existing file/folder during recovery.
	OverwriteExistingFile NullableBool `json:"overwriteExistingFile,omitempty"`
	// Specifies whether to preserve file/folder attributes during recovery.
	PreserveFileAttributes NullableBool `json:"preserveFileAttributes,omitempty"`
	// Specifies whether to continue recovering other files if one of the files fails to recover. Default value is false.
	ContinueOnError NullableBool `json:"continueOnError,omitempty"`
	// Specifies whether encryption should be enabled during recovery.
	EncryptionEnabled NullableBool `json:"encryptionEnabled,omitempty"`
	VlanConfig *RecoveryVlanConfig `json:"vlanConfig,omitempty"`
}

// NewOriginalGenericNasFilesTargetParams instantiates a new OriginalGenericNasFilesTargetParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginalGenericNasFilesTargetParams(recoverToOriginalPath NullableBool) *OriginalGenericNasFilesTargetParams {
	this := OriginalGenericNasFilesTargetParams{}
	this.RecoverToOriginalPath = recoverToOriginalPath
	return &this
}

// NewOriginalGenericNasFilesTargetParamsWithDefaults instantiates a new OriginalGenericNasFilesTargetParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginalGenericNasFilesTargetParamsWithDefaults() *OriginalGenericNasFilesTargetParams {
	this := OriginalGenericNasFilesTargetParams{}
	return &this
}

// GetRecoverToOriginalPath returns the RecoverToOriginalPath field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *OriginalGenericNasFilesTargetParams) GetRecoverToOriginalPath() bool {
	if o == nil || o.RecoverToOriginalPath.Get() == nil {
		var ret bool
		return ret
	}

	return *o.RecoverToOriginalPath.Get()
}

// GetRecoverToOriginalPathOk returns a tuple with the RecoverToOriginalPath field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OriginalGenericNasFilesTargetParams) GetRecoverToOriginalPathOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RecoverToOriginalPath.Get(), o.RecoverToOriginalPath.IsSet()
}

// SetRecoverToOriginalPath sets field value
func (o *OriginalGenericNasFilesTargetParams) SetRecoverToOriginalPath(v bool) {
	o.RecoverToOriginalPath.Set(&v)
}

// GetAlternatePath returns the AlternatePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OriginalGenericNasFilesTargetParams) GetAlternatePath() string {
	if o == nil || o.AlternatePath.Get() == nil {
		var ret string
		return ret
	}
	return *o.AlternatePath.Get()
}

// GetAlternatePathOk returns a tuple with the AlternatePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OriginalGenericNasFilesTargetParams) GetAlternatePathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AlternatePath.Get(), o.AlternatePath.IsSet()
}

// HasAlternatePath returns a boolean if a field has been set.
func (o *OriginalGenericNasFilesTargetParams) HasAlternatePath() bool {
	if o != nil && o.AlternatePath.IsSet() {
		return true
	}

	return false
}

// SetAlternatePath gets a reference to the given NullableString and assigns it to the AlternatePath field.
func (o *OriginalGenericNasFilesTargetParams) SetAlternatePath(v string) {
	o.AlternatePath.Set(&v)
}
// SetAlternatePathNil sets the value for AlternatePath to be an explicit nil
func (o *OriginalGenericNasFilesTargetParams) SetAlternatePathNil() {
	o.AlternatePath.Set(nil)
}

// UnsetAlternatePath ensures that no value is present for AlternatePath, not even an explicit nil
func (o *OriginalGenericNasFilesTargetParams) UnsetAlternatePath() {
	o.AlternatePath.Unset()
}

// GetOverwriteExistingFile returns the OverwriteExistingFile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OriginalGenericNasFilesTargetParams) GetOverwriteExistingFile() bool {
	if o == nil || o.OverwriteExistingFile.Get() == nil {
		var ret bool
		return ret
	}
	return *o.OverwriteExistingFile.Get()
}

// GetOverwriteExistingFileOk returns a tuple with the OverwriteExistingFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OriginalGenericNasFilesTargetParams) GetOverwriteExistingFileOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OverwriteExistingFile.Get(), o.OverwriteExistingFile.IsSet()
}

// HasOverwriteExistingFile returns a boolean if a field has been set.
func (o *OriginalGenericNasFilesTargetParams) HasOverwriteExistingFile() bool {
	if o != nil && o.OverwriteExistingFile.IsSet() {
		return true
	}

	return false
}

// SetOverwriteExistingFile gets a reference to the given NullableBool and assigns it to the OverwriteExistingFile field.
func (o *OriginalGenericNasFilesTargetParams) SetOverwriteExistingFile(v bool) {
	o.OverwriteExistingFile.Set(&v)
}
// SetOverwriteExistingFileNil sets the value for OverwriteExistingFile to be an explicit nil
func (o *OriginalGenericNasFilesTargetParams) SetOverwriteExistingFileNil() {
	o.OverwriteExistingFile.Set(nil)
}

// UnsetOverwriteExistingFile ensures that no value is present for OverwriteExistingFile, not even an explicit nil
func (o *OriginalGenericNasFilesTargetParams) UnsetOverwriteExistingFile() {
	o.OverwriteExistingFile.Unset()
}

// GetPreserveFileAttributes returns the PreserveFileAttributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OriginalGenericNasFilesTargetParams) GetPreserveFileAttributes() bool {
	if o == nil || o.PreserveFileAttributes.Get() == nil {
		var ret bool
		return ret
	}
	return *o.PreserveFileAttributes.Get()
}

// GetPreserveFileAttributesOk returns a tuple with the PreserveFileAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OriginalGenericNasFilesTargetParams) GetPreserveFileAttributesOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PreserveFileAttributes.Get(), o.PreserveFileAttributes.IsSet()
}

// HasPreserveFileAttributes returns a boolean if a field has been set.
func (o *OriginalGenericNasFilesTargetParams) HasPreserveFileAttributes() bool {
	if o != nil && o.PreserveFileAttributes.IsSet() {
		return true
	}

	return false
}

// SetPreserveFileAttributes gets a reference to the given NullableBool and assigns it to the PreserveFileAttributes field.
func (o *OriginalGenericNasFilesTargetParams) SetPreserveFileAttributes(v bool) {
	o.PreserveFileAttributes.Set(&v)
}
// SetPreserveFileAttributesNil sets the value for PreserveFileAttributes to be an explicit nil
func (o *OriginalGenericNasFilesTargetParams) SetPreserveFileAttributesNil() {
	o.PreserveFileAttributes.Set(nil)
}

// UnsetPreserveFileAttributes ensures that no value is present for PreserveFileAttributes, not even an explicit nil
func (o *OriginalGenericNasFilesTargetParams) UnsetPreserveFileAttributes() {
	o.PreserveFileAttributes.Unset()
}

// GetContinueOnError returns the ContinueOnError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OriginalGenericNasFilesTargetParams) GetContinueOnError() bool {
	if o == nil || o.ContinueOnError.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ContinueOnError.Get()
}

// GetContinueOnErrorOk returns a tuple with the ContinueOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OriginalGenericNasFilesTargetParams) GetContinueOnErrorOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContinueOnError.Get(), o.ContinueOnError.IsSet()
}

// HasContinueOnError returns a boolean if a field has been set.
func (o *OriginalGenericNasFilesTargetParams) HasContinueOnError() bool {
	if o != nil && o.ContinueOnError.IsSet() {
		return true
	}

	return false
}

// SetContinueOnError gets a reference to the given NullableBool and assigns it to the ContinueOnError field.
func (o *OriginalGenericNasFilesTargetParams) SetContinueOnError(v bool) {
	o.ContinueOnError.Set(&v)
}
// SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil
func (o *OriginalGenericNasFilesTargetParams) SetContinueOnErrorNil() {
	o.ContinueOnError.Set(nil)
}

// UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
func (o *OriginalGenericNasFilesTargetParams) UnsetContinueOnError() {
	o.ContinueOnError.Unset()
}

// GetEncryptionEnabled returns the EncryptionEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OriginalGenericNasFilesTargetParams) GetEncryptionEnabled() bool {
	if o == nil || o.EncryptionEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.EncryptionEnabled.Get()
}

// GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OriginalGenericNasFilesTargetParams) GetEncryptionEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EncryptionEnabled.Get(), o.EncryptionEnabled.IsSet()
}

// HasEncryptionEnabled returns a boolean if a field has been set.
func (o *OriginalGenericNasFilesTargetParams) HasEncryptionEnabled() bool {
	if o != nil && o.EncryptionEnabled.IsSet() {
		return true
	}

	return false
}

// SetEncryptionEnabled gets a reference to the given NullableBool and assigns it to the EncryptionEnabled field.
func (o *OriginalGenericNasFilesTargetParams) SetEncryptionEnabled(v bool) {
	o.EncryptionEnabled.Set(&v)
}
// SetEncryptionEnabledNil sets the value for EncryptionEnabled to be an explicit nil
func (o *OriginalGenericNasFilesTargetParams) SetEncryptionEnabledNil() {
	o.EncryptionEnabled.Set(nil)
}

// UnsetEncryptionEnabled ensures that no value is present for EncryptionEnabled, not even an explicit nil
func (o *OriginalGenericNasFilesTargetParams) UnsetEncryptionEnabled() {
	o.EncryptionEnabled.Unset()
}

// GetVlanConfig returns the VlanConfig field value if set, zero value otherwise.
func (o *OriginalGenericNasFilesTargetParams) GetVlanConfig() RecoveryVlanConfig {
	if o == nil || o.VlanConfig == nil {
		var ret RecoveryVlanConfig
		return ret
	}
	return *o.VlanConfig
}

// GetVlanConfigOk returns a tuple with the VlanConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginalGenericNasFilesTargetParams) GetVlanConfigOk() (*RecoveryVlanConfig, bool) {
	if o == nil || o.VlanConfig == nil {
		return nil, false
	}
	return o.VlanConfig, true
}

// HasVlanConfig returns a boolean if a field has been set.
func (o *OriginalGenericNasFilesTargetParams) HasVlanConfig() bool {
	if o != nil && o.VlanConfig != nil {
		return true
	}

	return false
}

// SetVlanConfig gets a reference to the given RecoveryVlanConfig and assigns it to the VlanConfig field.
func (o *OriginalGenericNasFilesTargetParams) SetVlanConfig(v RecoveryVlanConfig) {
	o.VlanConfig = &v
}

func (o OriginalGenericNasFilesTargetParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["recoverToOriginalPath"] = o.RecoverToOriginalPath.Get()
	}
	if o.AlternatePath.IsSet() {
		toSerialize["alternatePath"] = o.AlternatePath.Get()
	}
	if o.OverwriteExistingFile.IsSet() {
		toSerialize["overwriteExistingFile"] = o.OverwriteExistingFile.Get()
	}
	if o.PreserveFileAttributes.IsSet() {
		toSerialize["preserveFileAttributes"] = o.PreserveFileAttributes.Get()
	}
	if o.ContinueOnError.IsSet() {
		toSerialize["continueOnError"] = o.ContinueOnError.Get()
	}
	if o.EncryptionEnabled.IsSet() {
		toSerialize["encryptionEnabled"] = o.EncryptionEnabled.Get()
	}
	if o.VlanConfig != nil {
		toSerialize["vlanConfig"] = o.VlanConfig
	}
	return json.Marshal(toSerialize)
}

type NullableOriginalGenericNasFilesTargetParams struct {
	value *OriginalGenericNasFilesTargetParams
	isSet bool
}

func (v NullableOriginalGenericNasFilesTargetParams) Get() *OriginalGenericNasFilesTargetParams {
	return v.value
}

func (v *NullableOriginalGenericNasFilesTargetParams) Set(val *OriginalGenericNasFilesTargetParams) {
	v.value = val
	v.isSet = true
}

func (v NullableOriginalGenericNasFilesTargetParams) IsSet() bool {
	return v.isSet
}

func (v *NullableOriginalGenericNasFilesTargetParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOriginalGenericNasFilesTargetParams(val *OriginalGenericNasFilesTargetParams) *NullableOriginalGenericNasFilesTargetParams {
	return &NullableOriginalGenericNasFilesTargetParams{value: val, isSet: true}
}

func (v NullableOriginalGenericNasFilesTargetParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOriginalGenericNasFilesTargetParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o OriginalGenericNasFilesTargetParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}