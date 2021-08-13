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

// OriginalNetappFilesTargetParams Specifies the params of the original Netapp recovery target.
type OriginalNetappFilesTargetParams struct {
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
	FilterIpConfig *FilterIpConfig `json:"filterIpConfig,omitempty"`
	VlanConfig *RecoveryVlanConfig `json:"vlanConfig,omitempty"`
}

// NewOriginalNetappFilesTargetParams instantiates a new OriginalNetappFilesTargetParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginalNetappFilesTargetParams(recoverToOriginalPath NullableBool) *OriginalNetappFilesTargetParams {
	this := OriginalNetappFilesTargetParams{}
	this.RecoverToOriginalPath = recoverToOriginalPath
	return &this
}

// NewOriginalNetappFilesTargetParamsWithDefaults instantiates a new OriginalNetappFilesTargetParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginalNetappFilesTargetParamsWithDefaults() *OriginalNetappFilesTargetParams {
	this := OriginalNetappFilesTargetParams{}
	return &this
}

// GetRecoverToOriginalPath returns the RecoverToOriginalPath field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *OriginalNetappFilesTargetParams) GetRecoverToOriginalPath() bool {
	if o == nil || o.RecoverToOriginalPath.Get() == nil {
		var ret bool
		return ret
	}

	return *o.RecoverToOriginalPath.Get()
}

// GetRecoverToOriginalPathOk returns a tuple with the RecoverToOriginalPath field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OriginalNetappFilesTargetParams) GetRecoverToOriginalPathOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RecoverToOriginalPath.Get(), o.RecoverToOriginalPath.IsSet()
}

// SetRecoverToOriginalPath sets field value
func (o *OriginalNetappFilesTargetParams) SetRecoverToOriginalPath(v bool) {
	o.RecoverToOriginalPath.Set(&v)
}

// GetAlternatePath returns the AlternatePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OriginalNetappFilesTargetParams) GetAlternatePath() string {
	if o == nil || o.AlternatePath.Get() == nil {
		var ret string
		return ret
	}
	return *o.AlternatePath.Get()
}

// GetAlternatePathOk returns a tuple with the AlternatePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OriginalNetappFilesTargetParams) GetAlternatePathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AlternatePath.Get(), o.AlternatePath.IsSet()
}

// HasAlternatePath returns a boolean if a field has been set.
func (o *OriginalNetappFilesTargetParams) HasAlternatePath() bool {
	if o != nil && o.AlternatePath.IsSet() {
		return true
	}

	return false
}

// SetAlternatePath gets a reference to the given NullableString and assigns it to the AlternatePath field.
func (o *OriginalNetappFilesTargetParams) SetAlternatePath(v string) {
	o.AlternatePath.Set(&v)
}
// SetAlternatePathNil sets the value for AlternatePath to be an explicit nil
func (o *OriginalNetappFilesTargetParams) SetAlternatePathNil() {
	o.AlternatePath.Set(nil)
}

// UnsetAlternatePath ensures that no value is present for AlternatePath, not even an explicit nil
func (o *OriginalNetappFilesTargetParams) UnsetAlternatePath() {
	o.AlternatePath.Unset()
}

// GetOverwriteExistingFile returns the OverwriteExistingFile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OriginalNetappFilesTargetParams) GetOverwriteExistingFile() bool {
	if o == nil || o.OverwriteExistingFile.Get() == nil {
		var ret bool
		return ret
	}
	return *o.OverwriteExistingFile.Get()
}

// GetOverwriteExistingFileOk returns a tuple with the OverwriteExistingFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OriginalNetappFilesTargetParams) GetOverwriteExistingFileOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OverwriteExistingFile.Get(), o.OverwriteExistingFile.IsSet()
}

// HasOverwriteExistingFile returns a boolean if a field has been set.
func (o *OriginalNetappFilesTargetParams) HasOverwriteExistingFile() bool {
	if o != nil && o.OverwriteExistingFile.IsSet() {
		return true
	}

	return false
}

// SetOverwriteExistingFile gets a reference to the given NullableBool and assigns it to the OverwriteExistingFile field.
func (o *OriginalNetappFilesTargetParams) SetOverwriteExistingFile(v bool) {
	o.OverwriteExistingFile.Set(&v)
}
// SetOverwriteExistingFileNil sets the value for OverwriteExistingFile to be an explicit nil
func (o *OriginalNetappFilesTargetParams) SetOverwriteExistingFileNil() {
	o.OverwriteExistingFile.Set(nil)
}

// UnsetOverwriteExistingFile ensures that no value is present for OverwriteExistingFile, not even an explicit nil
func (o *OriginalNetappFilesTargetParams) UnsetOverwriteExistingFile() {
	o.OverwriteExistingFile.Unset()
}

// GetPreserveFileAttributes returns the PreserveFileAttributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OriginalNetappFilesTargetParams) GetPreserveFileAttributes() bool {
	if o == nil || o.PreserveFileAttributes.Get() == nil {
		var ret bool
		return ret
	}
	return *o.PreserveFileAttributes.Get()
}

// GetPreserveFileAttributesOk returns a tuple with the PreserveFileAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OriginalNetappFilesTargetParams) GetPreserveFileAttributesOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PreserveFileAttributes.Get(), o.PreserveFileAttributes.IsSet()
}

// HasPreserveFileAttributes returns a boolean if a field has been set.
func (o *OriginalNetappFilesTargetParams) HasPreserveFileAttributes() bool {
	if o != nil && o.PreserveFileAttributes.IsSet() {
		return true
	}

	return false
}

// SetPreserveFileAttributes gets a reference to the given NullableBool and assigns it to the PreserveFileAttributes field.
func (o *OriginalNetappFilesTargetParams) SetPreserveFileAttributes(v bool) {
	o.PreserveFileAttributes.Set(&v)
}
// SetPreserveFileAttributesNil sets the value for PreserveFileAttributes to be an explicit nil
func (o *OriginalNetappFilesTargetParams) SetPreserveFileAttributesNil() {
	o.PreserveFileAttributes.Set(nil)
}

// UnsetPreserveFileAttributes ensures that no value is present for PreserveFileAttributes, not even an explicit nil
func (o *OriginalNetappFilesTargetParams) UnsetPreserveFileAttributes() {
	o.PreserveFileAttributes.Unset()
}

// GetContinueOnError returns the ContinueOnError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OriginalNetappFilesTargetParams) GetContinueOnError() bool {
	if o == nil || o.ContinueOnError.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ContinueOnError.Get()
}

// GetContinueOnErrorOk returns a tuple with the ContinueOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OriginalNetappFilesTargetParams) GetContinueOnErrorOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContinueOnError.Get(), o.ContinueOnError.IsSet()
}

// HasContinueOnError returns a boolean if a field has been set.
func (o *OriginalNetappFilesTargetParams) HasContinueOnError() bool {
	if o != nil && o.ContinueOnError.IsSet() {
		return true
	}

	return false
}

// SetContinueOnError gets a reference to the given NullableBool and assigns it to the ContinueOnError field.
func (o *OriginalNetappFilesTargetParams) SetContinueOnError(v bool) {
	o.ContinueOnError.Set(&v)
}
// SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil
func (o *OriginalNetappFilesTargetParams) SetContinueOnErrorNil() {
	o.ContinueOnError.Set(nil)
}

// UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
func (o *OriginalNetappFilesTargetParams) UnsetContinueOnError() {
	o.ContinueOnError.Unset()
}

// GetEncryptionEnabled returns the EncryptionEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OriginalNetappFilesTargetParams) GetEncryptionEnabled() bool {
	if o == nil || o.EncryptionEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.EncryptionEnabled.Get()
}

// GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OriginalNetappFilesTargetParams) GetEncryptionEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EncryptionEnabled.Get(), o.EncryptionEnabled.IsSet()
}

// HasEncryptionEnabled returns a boolean if a field has been set.
func (o *OriginalNetappFilesTargetParams) HasEncryptionEnabled() bool {
	if o != nil && o.EncryptionEnabled.IsSet() {
		return true
	}

	return false
}

// SetEncryptionEnabled gets a reference to the given NullableBool and assigns it to the EncryptionEnabled field.
func (o *OriginalNetappFilesTargetParams) SetEncryptionEnabled(v bool) {
	o.EncryptionEnabled.Set(&v)
}
// SetEncryptionEnabledNil sets the value for EncryptionEnabled to be an explicit nil
func (o *OriginalNetappFilesTargetParams) SetEncryptionEnabledNil() {
	o.EncryptionEnabled.Set(nil)
}

// UnsetEncryptionEnabled ensures that no value is present for EncryptionEnabled, not even an explicit nil
func (o *OriginalNetappFilesTargetParams) UnsetEncryptionEnabled() {
	o.EncryptionEnabled.Unset()
}

// GetFilterIpConfig returns the FilterIpConfig field value if set, zero value otherwise.
func (o *OriginalNetappFilesTargetParams) GetFilterIpConfig() FilterIpConfig {
	if o == nil || o.FilterIpConfig == nil {
		var ret FilterIpConfig
		return ret
	}
	return *o.FilterIpConfig
}

// GetFilterIpConfigOk returns a tuple with the FilterIpConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginalNetappFilesTargetParams) GetFilterIpConfigOk() (*FilterIpConfig, bool) {
	if o == nil || o.FilterIpConfig == nil {
		return nil, false
	}
	return o.FilterIpConfig, true
}

// HasFilterIpConfig returns a boolean if a field has been set.
func (o *OriginalNetappFilesTargetParams) HasFilterIpConfig() bool {
	if o != nil && o.FilterIpConfig != nil {
		return true
	}

	return false
}

// SetFilterIpConfig gets a reference to the given FilterIpConfig and assigns it to the FilterIpConfig field.
func (o *OriginalNetappFilesTargetParams) SetFilterIpConfig(v FilterIpConfig) {
	o.FilterIpConfig = &v
}

// GetVlanConfig returns the VlanConfig field value if set, zero value otherwise.
func (o *OriginalNetappFilesTargetParams) GetVlanConfig() RecoveryVlanConfig {
	if o == nil || o.VlanConfig == nil {
		var ret RecoveryVlanConfig
		return ret
	}
	return *o.VlanConfig
}

// GetVlanConfigOk returns a tuple with the VlanConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginalNetappFilesTargetParams) GetVlanConfigOk() (*RecoveryVlanConfig, bool) {
	if o == nil || o.VlanConfig == nil {
		return nil, false
	}
	return o.VlanConfig, true
}

// HasVlanConfig returns a boolean if a field has been set.
func (o *OriginalNetappFilesTargetParams) HasVlanConfig() bool {
	if o != nil && o.VlanConfig != nil {
		return true
	}

	return false
}

// SetVlanConfig gets a reference to the given RecoveryVlanConfig and assigns it to the VlanConfig field.
func (o *OriginalNetappFilesTargetParams) SetVlanConfig(v RecoveryVlanConfig) {
	o.VlanConfig = &v
}

func (o OriginalNetappFilesTargetParams) MarshalJSON() ([]byte, error) {
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
	if o.FilterIpConfig != nil {
		toSerialize["filterIpConfig"] = o.FilterIpConfig
	}
	if o.VlanConfig != nil {
		toSerialize["vlanConfig"] = o.VlanConfig
	}
	return json.Marshal(toSerialize)
}

type NullableOriginalNetappFilesTargetParams struct {
	value *OriginalNetappFilesTargetParams
	isSet bool
}

func (v NullableOriginalNetappFilesTargetParams) Get() *OriginalNetappFilesTargetParams {
	return v.value
}

func (v *NullableOriginalNetappFilesTargetParams) Set(val *OriginalNetappFilesTargetParams) {
	v.value = val
	v.isSet = true
}

func (v NullableOriginalNetappFilesTargetParams) IsSet() bool {
	return v.isSet
}

func (v *NullableOriginalNetappFilesTargetParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOriginalNetappFilesTargetParams(val *OriginalNetappFilesTargetParams) *NullableOriginalNetappFilesTargetParams {
	return &NullableOriginalNetappFilesTargetParams{value: val, isSet: true}
}

func (v NullableOriginalNetappFilesTargetParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOriginalNetappFilesTargetParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


