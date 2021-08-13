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

// RecoverOtherNasToNetappFilesTargetParams Specifies the params of the Netapp recovery target.
type RecoverOtherNasToNetappFilesTargetParams struct {
	// Specifies the id and name of the parent NAS to recover to. This volume will be the target of the recovery.
	Volume RecoverTarget `json:"volume"`
	// Specifies the id of the parent source of the recovery target.
	ParentSource *RecoveryObjectIdentifier `json:"parentSource,omitempty"`
	// Specifies the path location to recover files to.
	AlternatePath NullableString `json:"alternatePath"`
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

// NewRecoverOtherNasToNetappFilesTargetParams instantiates a new RecoverOtherNasToNetappFilesTargetParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverOtherNasToNetappFilesTargetParams(volume RecoverTarget, alternatePath NullableString) *RecoverOtherNasToNetappFilesTargetParams {
	this := RecoverOtherNasToNetappFilesTargetParams{}
	this.Volume = volume
	this.AlternatePath = alternatePath
	return &this
}

// NewRecoverOtherNasToNetappFilesTargetParamsWithDefaults instantiates a new RecoverOtherNasToNetappFilesTargetParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverOtherNasToNetappFilesTargetParamsWithDefaults() *RecoverOtherNasToNetappFilesTargetParams {
	this := RecoverOtherNasToNetappFilesTargetParams{}
	return &this
}

// GetVolume returns the Volume field value
func (o *RecoverOtherNasToNetappFilesTargetParams) GetVolume() RecoverTarget {
	if o == nil {
		var ret RecoverTarget
		return ret
	}

	return o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value
// and a boolean to check if the value has been set.
func (o *RecoverOtherNasToNetappFilesTargetParams) GetVolumeOk() (*RecoverTarget, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Volume, true
}

// SetVolume sets field value
func (o *RecoverOtherNasToNetappFilesTargetParams) SetVolume(v RecoverTarget) {
	o.Volume = v
}

// GetParentSource returns the ParentSource field value if set, zero value otherwise.
func (o *RecoverOtherNasToNetappFilesTargetParams) GetParentSource() RecoveryObjectIdentifier {
	if o == nil || o.ParentSource == nil {
		var ret RecoveryObjectIdentifier
		return ret
	}
	return *o.ParentSource
}

// GetParentSourceOk returns a tuple with the ParentSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoverOtherNasToNetappFilesTargetParams) GetParentSourceOk() (*RecoveryObjectIdentifier, bool) {
	if o == nil || o.ParentSource == nil {
		return nil, false
	}
	return o.ParentSource, true
}

// HasParentSource returns a boolean if a field has been set.
func (o *RecoverOtherNasToNetappFilesTargetParams) HasParentSource() bool {
	if o != nil && o.ParentSource != nil {
		return true
	}

	return false
}

// SetParentSource gets a reference to the given RecoveryObjectIdentifier and assigns it to the ParentSource field.
func (o *RecoverOtherNasToNetappFilesTargetParams) SetParentSource(v RecoveryObjectIdentifier) {
	o.ParentSource = &v
}

// GetAlternatePath returns the AlternatePath field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RecoverOtherNasToNetappFilesTargetParams) GetAlternatePath() string {
	if o == nil || o.AlternatePath.Get() == nil {
		var ret string
		return ret
	}

	return *o.AlternatePath.Get()
}

// GetAlternatePathOk returns a tuple with the AlternatePath field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverOtherNasToNetappFilesTargetParams) GetAlternatePathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AlternatePath.Get(), o.AlternatePath.IsSet()
}

// SetAlternatePath sets field value
func (o *RecoverOtherNasToNetappFilesTargetParams) SetAlternatePath(v string) {
	o.AlternatePath.Set(&v)
}

// GetOverwriteExistingFile returns the OverwriteExistingFile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverOtherNasToNetappFilesTargetParams) GetOverwriteExistingFile() bool {
	if o == nil || o.OverwriteExistingFile.Get() == nil {
		var ret bool
		return ret
	}
	return *o.OverwriteExistingFile.Get()
}

// GetOverwriteExistingFileOk returns a tuple with the OverwriteExistingFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverOtherNasToNetappFilesTargetParams) GetOverwriteExistingFileOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OverwriteExistingFile.Get(), o.OverwriteExistingFile.IsSet()
}

// HasOverwriteExistingFile returns a boolean if a field has been set.
func (o *RecoverOtherNasToNetappFilesTargetParams) HasOverwriteExistingFile() bool {
	if o != nil && o.OverwriteExistingFile.IsSet() {
		return true
	}

	return false
}

// SetOverwriteExistingFile gets a reference to the given NullableBool and assigns it to the OverwriteExistingFile field.
func (o *RecoverOtherNasToNetappFilesTargetParams) SetOverwriteExistingFile(v bool) {
	o.OverwriteExistingFile.Set(&v)
}
// SetOverwriteExistingFileNil sets the value for OverwriteExistingFile to be an explicit nil
func (o *RecoverOtherNasToNetappFilesTargetParams) SetOverwriteExistingFileNil() {
	o.OverwriteExistingFile.Set(nil)
}

// UnsetOverwriteExistingFile ensures that no value is present for OverwriteExistingFile, not even an explicit nil
func (o *RecoverOtherNasToNetappFilesTargetParams) UnsetOverwriteExistingFile() {
	o.OverwriteExistingFile.Unset()
}

// GetPreserveFileAttributes returns the PreserveFileAttributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverOtherNasToNetappFilesTargetParams) GetPreserveFileAttributes() bool {
	if o == nil || o.PreserveFileAttributes.Get() == nil {
		var ret bool
		return ret
	}
	return *o.PreserveFileAttributes.Get()
}

// GetPreserveFileAttributesOk returns a tuple with the PreserveFileAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverOtherNasToNetappFilesTargetParams) GetPreserveFileAttributesOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PreserveFileAttributes.Get(), o.PreserveFileAttributes.IsSet()
}

// HasPreserveFileAttributes returns a boolean if a field has been set.
func (o *RecoverOtherNasToNetappFilesTargetParams) HasPreserveFileAttributes() bool {
	if o != nil && o.PreserveFileAttributes.IsSet() {
		return true
	}

	return false
}

// SetPreserveFileAttributes gets a reference to the given NullableBool and assigns it to the PreserveFileAttributes field.
func (o *RecoverOtherNasToNetappFilesTargetParams) SetPreserveFileAttributes(v bool) {
	o.PreserveFileAttributes.Set(&v)
}
// SetPreserveFileAttributesNil sets the value for PreserveFileAttributes to be an explicit nil
func (o *RecoverOtherNasToNetappFilesTargetParams) SetPreserveFileAttributesNil() {
	o.PreserveFileAttributes.Set(nil)
}

// UnsetPreserveFileAttributes ensures that no value is present for PreserveFileAttributes, not even an explicit nil
func (o *RecoverOtherNasToNetappFilesTargetParams) UnsetPreserveFileAttributes() {
	o.PreserveFileAttributes.Unset()
}

// GetContinueOnError returns the ContinueOnError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverOtherNasToNetappFilesTargetParams) GetContinueOnError() bool {
	if o == nil || o.ContinueOnError.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ContinueOnError.Get()
}

// GetContinueOnErrorOk returns a tuple with the ContinueOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverOtherNasToNetappFilesTargetParams) GetContinueOnErrorOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContinueOnError.Get(), o.ContinueOnError.IsSet()
}

// HasContinueOnError returns a boolean if a field has been set.
func (o *RecoverOtherNasToNetappFilesTargetParams) HasContinueOnError() bool {
	if o != nil && o.ContinueOnError.IsSet() {
		return true
	}

	return false
}

// SetContinueOnError gets a reference to the given NullableBool and assigns it to the ContinueOnError field.
func (o *RecoverOtherNasToNetappFilesTargetParams) SetContinueOnError(v bool) {
	o.ContinueOnError.Set(&v)
}
// SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil
func (o *RecoverOtherNasToNetappFilesTargetParams) SetContinueOnErrorNil() {
	o.ContinueOnError.Set(nil)
}

// UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
func (o *RecoverOtherNasToNetappFilesTargetParams) UnsetContinueOnError() {
	o.ContinueOnError.Unset()
}

// GetEncryptionEnabled returns the EncryptionEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverOtherNasToNetappFilesTargetParams) GetEncryptionEnabled() bool {
	if o == nil || o.EncryptionEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.EncryptionEnabled.Get()
}

// GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverOtherNasToNetappFilesTargetParams) GetEncryptionEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EncryptionEnabled.Get(), o.EncryptionEnabled.IsSet()
}

// HasEncryptionEnabled returns a boolean if a field has been set.
func (o *RecoverOtherNasToNetappFilesTargetParams) HasEncryptionEnabled() bool {
	if o != nil && o.EncryptionEnabled.IsSet() {
		return true
	}

	return false
}

// SetEncryptionEnabled gets a reference to the given NullableBool and assigns it to the EncryptionEnabled field.
func (o *RecoverOtherNasToNetappFilesTargetParams) SetEncryptionEnabled(v bool) {
	o.EncryptionEnabled.Set(&v)
}
// SetEncryptionEnabledNil sets the value for EncryptionEnabled to be an explicit nil
func (o *RecoverOtherNasToNetappFilesTargetParams) SetEncryptionEnabledNil() {
	o.EncryptionEnabled.Set(nil)
}

// UnsetEncryptionEnabled ensures that no value is present for EncryptionEnabled, not even an explicit nil
func (o *RecoverOtherNasToNetappFilesTargetParams) UnsetEncryptionEnabled() {
	o.EncryptionEnabled.Unset()
}

// GetFilterIpConfig returns the FilterIpConfig field value if set, zero value otherwise.
func (o *RecoverOtherNasToNetappFilesTargetParams) GetFilterIpConfig() FilterIpConfig {
	if o == nil || o.FilterIpConfig == nil {
		var ret FilterIpConfig
		return ret
	}
	return *o.FilterIpConfig
}

// GetFilterIpConfigOk returns a tuple with the FilterIpConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoverOtherNasToNetappFilesTargetParams) GetFilterIpConfigOk() (*FilterIpConfig, bool) {
	if o == nil || o.FilterIpConfig == nil {
		return nil, false
	}
	return o.FilterIpConfig, true
}

// HasFilterIpConfig returns a boolean if a field has been set.
func (o *RecoverOtherNasToNetappFilesTargetParams) HasFilterIpConfig() bool {
	if o != nil && o.FilterIpConfig != nil {
		return true
	}

	return false
}

// SetFilterIpConfig gets a reference to the given FilterIpConfig and assigns it to the FilterIpConfig field.
func (o *RecoverOtherNasToNetappFilesTargetParams) SetFilterIpConfig(v FilterIpConfig) {
	o.FilterIpConfig = &v
}

// GetVlanConfig returns the VlanConfig field value if set, zero value otherwise.
func (o *RecoverOtherNasToNetappFilesTargetParams) GetVlanConfig() RecoveryVlanConfig {
	if o == nil || o.VlanConfig == nil {
		var ret RecoveryVlanConfig
		return ret
	}
	return *o.VlanConfig
}

// GetVlanConfigOk returns a tuple with the VlanConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoverOtherNasToNetappFilesTargetParams) GetVlanConfigOk() (*RecoveryVlanConfig, bool) {
	if o == nil || o.VlanConfig == nil {
		return nil, false
	}
	return o.VlanConfig, true
}

// HasVlanConfig returns a boolean if a field has been set.
func (o *RecoverOtherNasToNetappFilesTargetParams) HasVlanConfig() bool {
	if o != nil && o.VlanConfig != nil {
		return true
	}

	return false
}

// SetVlanConfig gets a reference to the given RecoveryVlanConfig and assigns it to the VlanConfig field.
func (o *RecoverOtherNasToNetappFilesTargetParams) SetVlanConfig(v RecoveryVlanConfig) {
	o.VlanConfig = &v
}

func (o RecoverOtherNasToNetappFilesTargetParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["volume"] = o.Volume
	}
	if o.ParentSource != nil {
		toSerialize["parentSource"] = o.ParentSource
	}
	if true {
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

type NullableRecoverOtherNasToNetappFilesTargetParams struct {
	value *RecoverOtherNasToNetappFilesTargetParams
	isSet bool
}

func (v NullableRecoverOtherNasToNetappFilesTargetParams) Get() *RecoverOtherNasToNetappFilesTargetParams {
	return v.value
}

func (v *NullableRecoverOtherNasToNetappFilesTargetParams) Set(val *RecoverOtherNasToNetappFilesTargetParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverOtherNasToNetappFilesTargetParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverOtherNasToNetappFilesTargetParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverOtherNasToNetappFilesTargetParams(val *RecoverOtherNasToNetappFilesTargetParams) *NullableRecoverOtherNasToNetappFilesTargetParams {
	return &NullableRecoverOtherNasToNetappFilesTargetParams{value: val, isSet: true}
}

func (v NullableRecoverOtherNasToNetappFilesTargetParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverOtherNasToNetappFilesTargetParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


