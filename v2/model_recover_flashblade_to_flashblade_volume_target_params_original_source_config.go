/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

// checks if the RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig{}

// RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig Specifies the Source configuration if volumes are being recovered to original Source. If not specified, all the configuration parameters will be retained.
type RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig struct {
	// Specifies whether to continue recovering other volumes if one of the volumes fails to recover. Default value is false.
	ContinueOnError NullableBool `json:"continueOnError,omitempty"`
	// Specifies whether encryption should be enabled during recovery.
	EncryptionEnabled NullableBool `json:"encryptionEnabled,omitempty"`
	FilterIpConfig *FilterIpConfig `json:"filterIpConfig,omitempty"`
	// Specifies whether to overwrite existing file/folder during recovery.
	OverwriteExistingFile NullableBool `json:"overwriteExistingFile,omitempty"`
	// Specifies whether to preserve file/folder attributes during recovery.
	PreserveFileAttributes NullableBool `json:"preserveFileAttributes,omitempty"`
	VlanConfig *RecoveryVlanConfig `json:"vlanConfig,omitempty"`
}

// NewRecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig instantiates a new RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig() *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig {
	this := RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig{}
	return &this
}

// NewRecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfigWithDefaults instantiates a new RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfigWithDefaults() *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig {
	this := RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig{}
	return &this
}

// GetContinueOnError returns the ContinueOnError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) GetContinueOnError() bool {
	if o == nil || IsNil(o.ContinueOnError.Get()) {
		var ret bool
		return ret
	}
	return *o.ContinueOnError.Get()
}

// GetContinueOnErrorOk returns a tuple with the ContinueOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) GetContinueOnErrorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContinueOnError.Get(), o.ContinueOnError.IsSet()
}

// HasContinueOnError returns a boolean if a field has been set.
func (o *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) HasContinueOnError() bool {
	if o != nil && o.ContinueOnError.IsSet() {
		return true
	}

	return false
}

// SetContinueOnError gets a reference to the given NullableBool and assigns it to the ContinueOnError field.
func (o *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) SetContinueOnError(v bool) {
	o.ContinueOnError.Set(&v)
}
// SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil
func (o *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) SetContinueOnErrorNil() {
	o.ContinueOnError.Set(nil)
}

// UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
func (o *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) UnsetContinueOnError() {
	o.ContinueOnError.Unset()
}

// GetEncryptionEnabled returns the EncryptionEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) GetEncryptionEnabled() bool {
	if o == nil || IsNil(o.EncryptionEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.EncryptionEnabled.Get()
}

// GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) GetEncryptionEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EncryptionEnabled.Get(), o.EncryptionEnabled.IsSet()
}

// HasEncryptionEnabled returns a boolean if a field has been set.
func (o *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) HasEncryptionEnabled() bool {
	if o != nil && o.EncryptionEnabled.IsSet() {
		return true
	}

	return false
}

// SetEncryptionEnabled gets a reference to the given NullableBool and assigns it to the EncryptionEnabled field.
func (o *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) SetEncryptionEnabled(v bool) {
	o.EncryptionEnabled.Set(&v)
}
// SetEncryptionEnabledNil sets the value for EncryptionEnabled to be an explicit nil
func (o *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) SetEncryptionEnabledNil() {
	o.EncryptionEnabled.Set(nil)
}

// UnsetEncryptionEnabled ensures that no value is present for EncryptionEnabled, not even an explicit nil
func (o *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) UnsetEncryptionEnabled() {
	o.EncryptionEnabled.Unset()
}

// GetFilterIpConfig returns the FilterIpConfig field value if set, zero value otherwise.
func (o *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) GetFilterIpConfig() FilterIpConfig {
	if o == nil || IsNil(o.FilterIpConfig) {
		var ret FilterIpConfig
		return ret
	}
	return *o.FilterIpConfig
}

// GetFilterIpConfigOk returns a tuple with the FilterIpConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) GetFilterIpConfigOk() (*FilterIpConfig, bool) {
	if o == nil || IsNil(o.FilterIpConfig) {
		return nil, false
	}
	return o.FilterIpConfig, true
}

// HasFilterIpConfig returns a boolean if a field has been set.
func (o *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) HasFilterIpConfig() bool {
	if o != nil && !IsNil(o.FilterIpConfig) {
		return true
	}

	return false
}

// SetFilterIpConfig gets a reference to the given FilterIpConfig and assigns it to the FilterIpConfig field.
func (o *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) SetFilterIpConfig(v FilterIpConfig) {
	o.FilterIpConfig = &v
}

// GetOverwriteExistingFile returns the OverwriteExistingFile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) GetOverwriteExistingFile() bool {
	if o == nil || IsNil(o.OverwriteExistingFile.Get()) {
		var ret bool
		return ret
	}
	return *o.OverwriteExistingFile.Get()
}

// GetOverwriteExistingFileOk returns a tuple with the OverwriteExistingFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) GetOverwriteExistingFileOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.OverwriteExistingFile.Get(), o.OverwriteExistingFile.IsSet()
}

// HasOverwriteExistingFile returns a boolean if a field has been set.
func (o *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) HasOverwriteExistingFile() bool {
	if o != nil && o.OverwriteExistingFile.IsSet() {
		return true
	}

	return false
}

// SetOverwriteExistingFile gets a reference to the given NullableBool and assigns it to the OverwriteExistingFile field.
func (o *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) SetOverwriteExistingFile(v bool) {
	o.OverwriteExistingFile.Set(&v)
}
// SetOverwriteExistingFileNil sets the value for OverwriteExistingFile to be an explicit nil
func (o *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) SetOverwriteExistingFileNil() {
	o.OverwriteExistingFile.Set(nil)
}

// UnsetOverwriteExistingFile ensures that no value is present for OverwriteExistingFile, not even an explicit nil
func (o *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) UnsetOverwriteExistingFile() {
	o.OverwriteExistingFile.Unset()
}

// GetPreserveFileAttributes returns the PreserveFileAttributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) GetPreserveFileAttributes() bool {
	if o == nil || IsNil(o.PreserveFileAttributes.Get()) {
		var ret bool
		return ret
	}
	return *o.PreserveFileAttributes.Get()
}

// GetPreserveFileAttributesOk returns a tuple with the PreserveFileAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) GetPreserveFileAttributesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreserveFileAttributes.Get(), o.PreserveFileAttributes.IsSet()
}

// HasPreserveFileAttributes returns a boolean if a field has been set.
func (o *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) HasPreserveFileAttributes() bool {
	if o != nil && o.PreserveFileAttributes.IsSet() {
		return true
	}

	return false
}

// SetPreserveFileAttributes gets a reference to the given NullableBool and assigns it to the PreserveFileAttributes field.
func (o *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) SetPreserveFileAttributes(v bool) {
	o.PreserveFileAttributes.Set(&v)
}
// SetPreserveFileAttributesNil sets the value for PreserveFileAttributes to be an explicit nil
func (o *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) SetPreserveFileAttributesNil() {
	o.PreserveFileAttributes.Set(nil)
}

// UnsetPreserveFileAttributes ensures that no value is present for PreserveFileAttributes, not even an explicit nil
func (o *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) UnsetPreserveFileAttributes() {
	o.PreserveFileAttributes.Unset()
}

// GetVlanConfig returns the VlanConfig field value if set, zero value otherwise.
func (o *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) GetVlanConfig() RecoveryVlanConfig {
	if o == nil || IsNil(o.VlanConfig) {
		var ret RecoveryVlanConfig
		return ret
	}
	return *o.VlanConfig
}

// GetVlanConfigOk returns a tuple with the VlanConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) GetVlanConfigOk() (*RecoveryVlanConfig, bool) {
	if o == nil || IsNil(o.VlanConfig) {
		return nil, false
	}
	return o.VlanConfig, true
}

// HasVlanConfig returns a boolean if a field has been set.
func (o *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) HasVlanConfig() bool {
	if o != nil && !IsNil(o.VlanConfig) {
		return true
	}

	return false
}

// SetVlanConfig gets a reference to the given RecoveryVlanConfig and assigns it to the VlanConfig field.
func (o *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) SetVlanConfig(v RecoveryVlanConfig) {
	o.VlanConfig = &v
}

func (o RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ContinueOnError.IsSet() {
		toSerialize["continueOnError"] = o.ContinueOnError.Get()
	}
	if o.EncryptionEnabled.IsSet() {
		toSerialize["encryptionEnabled"] = o.EncryptionEnabled.Get()
	}
	if !IsNil(o.FilterIpConfig) {
		toSerialize["filterIpConfig"] = o.FilterIpConfig
	}
	if o.OverwriteExistingFile.IsSet() {
		toSerialize["overwriteExistingFile"] = o.OverwriteExistingFile.Get()
	}
	if o.PreserveFileAttributes.IsSet() {
		toSerialize["preserveFileAttributes"] = o.PreserveFileAttributes.Get()
	}
	if !IsNil(o.VlanConfig) {
		toSerialize["vlanConfig"] = o.VlanConfig
	}
	return toSerialize, nil
}

type NullableRecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig struct {
	value *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig
	isSet bool
}

func (v NullableRecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) Get() *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig {
	return v.value
}

func (v *NullableRecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) Set(val *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig(val *RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) *NullableRecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig {
	return &NullableRecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig{value: val, isSet: true}
}

func (v NullableRecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


