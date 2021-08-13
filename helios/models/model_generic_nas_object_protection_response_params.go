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

// GenericNasObjectProtectionResponseParams Specifies the parameters which are specific to Generic NAS object protection.
type GenericNasObjectProtectionResponseParams struct {
	IndexingPolicy *IndexingPolicy `json:"indexingPolicy,omitempty"`
	// Specifies whether or not the backup should continue regardless of whether or not an error was encountered.
	ContinueOnError NullableBool `json:"continueOnError,omitempty"`
	// Specifies whether the encryption should be used while backup or not.
	EncryptionEnabled NullableBool `json:"encryptionEnabled,omitempty"`
	FileLockConfig *FileLevelDataLockConfig `json:"fileLockConfig,omitempty"`
	FileFilters *FileFilteringPolicy `json:"fileFilters,omitempty"`
	PrePostScript *HostBasedBackupScriptParams `json:"prePostScript,omitempty"`
	ThrottlingConfig *NasThrottlingConfig `json:"throttlingConfig,omitempty"`
}

// NewGenericNasObjectProtectionResponseParams instantiates a new GenericNasObjectProtectionResponseParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericNasObjectProtectionResponseParams() *GenericNasObjectProtectionResponseParams {
	this := GenericNasObjectProtectionResponseParams{}
	return &this
}

// NewGenericNasObjectProtectionResponseParamsWithDefaults instantiates a new GenericNasObjectProtectionResponseParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericNasObjectProtectionResponseParamsWithDefaults() *GenericNasObjectProtectionResponseParams {
	this := GenericNasObjectProtectionResponseParams{}
	return &this
}

// GetIndexingPolicy returns the IndexingPolicy field value if set, zero value otherwise.
func (o *GenericNasObjectProtectionResponseParams) GetIndexingPolicy() IndexingPolicy {
	if o == nil || o.IndexingPolicy == nil {
		var ret IndexingPolicy
		return ret
	}
	return *o.IndexingPolicy
}

// GetIndexingPolicyOk returns a tuple with the IndexingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericNasObjectProtectionResponseParams) GetIndexingPolicyOk() (*IndexingPolicy, bool) {
	if o == nil || o.IndexingPolicy == nil {
		return nil, false
	}
	return o.IndexingPolicy, true
}

// HasIndexingPolicy returns a boolean if a field has been set.
func (o *GenericNasObjectProtectionResponseParams) HasIndexingPolicy() bool {
	if o != nil && o.IndexingPolicy != nil {
		return true
	}

	return false
}

// SetIndexingPolicy gets a reference to the given IndexingPolicy and assigns it to the IndexingPolicy field.
func (o *GenericNasObjectProtectionResponseParams) SetIndexingPolicy(v IndexingPolicy) {
	o.IndexingPolicy = &v
}

// GetContinueOnError returns the ContinueOnError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GenericNasObjectProtectionResponseParams) GetContinueOnError() bool {
	if o == nil || o.ContinueOnError.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ContinueOnError.Get()
}

// GetContinueOnErrorOk returns a tuple with the ContinueOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenericNasObjectProtectionResponseParams) GetContinueOnErrorOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContinueOnError.Get(), o.ContinueOnError.IsSet()
}

// HasContinueOnError returns a boolean if a field has been set.
func (o *GenericNasObjectProtectionResponseParams) HasContinueOnError() bool {
	if o != nil && o.ContinueOnError.IsSet() {
		return true
	}

	return false
}

// SetContinueOnError gets a reference to the given NullableBool and assigns it to the ContinueOnError field.
func (o *GenericNasObjectProtectionResponseParams) SetContinueOnError(v bool) {
	o.ContinueOnError.Set(&v)
}
// SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil
func (o *GenericNasObjectProtectionResponseParams) SetContinueOnErrorNil() {
	o.ContinueOnError.Set(nil)
}

// UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
func (o *GenericNasObjectProtectionResponseParams) UnsetContinueOnError() {
	o.ContinueOnError.Unset()
}

// GetEncryptionEnabled returns the EncryptionEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GenericNasObjectProtectionResponseParams) GetEncryptionEnabled() bool {
	if o == nil || o.EncryptionEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.EncryptionEnabled.Get()
}

// GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenericNasObjectProtectionResponseParams) GetEncryptionEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EncryptionEnabled.Get(), o.EncryptionEnabled.IsSet()
}

// HasEncryptionEnabled returns a boolean if a field has been set.
func (o *GenericNasObjectProtectionResponseParams) HasEncryptionEnabled() bool {
	if o != nil && o.EncryptionEnabled.IsSet() {
		return true
	}

	return false
}

// SetEncryptionEnabled gets a reference to the given NullableBool and assigns it to the EncryptionEnabled field.
func (o *GenericNasObjectProtectionResponseParams) SetEncryptionEnabled(v bool) {
	o.EncryptionEnabled.Set(&v)
}
// SetEncryptionEnabledNil sets the value for EncryptionEnabled to be an explicit nil
func (o *GenericNasObjectProtectionResponseParams) SetEncryptionEnabledNil() {
	o.EncryptionEnabled.Set(nil)
}

// UnsetEncryptionEnabled ensures that no value is present for EncryptionEnabled, not even an explicit nil
func (o *GenericNasObjectProtectionResponseParams) UnsetEncryptionEnabled() {
	o.EncryptionEnabled.Unset()
}

// GetFileLockConfig returns the FileLockConfig field value if set, zero value otherwise.
func (o *GenericNasObjectProtectionResponseParams) GetFileLockConfig() FileLevelDataLockConfig {
	if o == nil || o.FileLockConfig == nil {
		var ret FileLevelDataLockConfig
		return ret
	}
	return *o.FileLockConfig
}

// GetFileLockConfigOk returns a tuple with the FileLockConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericNasObjectProtectionResponseParams) GetFileLockConfigOk() (*FileLevelDataLockConfig, bool) {
	if o == nil || o.FileLockConfig == nil {
		return nil, false
	}
	return o.FileLockConfig, true
}

// HasFileLockConfig returns a boolean if a field has been set.
func (o *GenericNasObjectProtectionResponseParams) HasFileLockConfig() bool {
	if o != nil && o.FileLockConfig != nil {
		return true
	}

	return false
}

// SetFileLockConfig gets a reference to the given FileLevelDataLockConfig and assigns it to the FileLockConfig field.
func (o *GenericNasObjectProtectionResponseParams) SetFileLockConfig(v FileLevelDataLockConfig) {
	o.FileLockConfig = &v
}

// GetFileFilters returns the FileFilters field value if set, zero value otherwise.
func (o *GenericNasObjectProtectionResponseParams) GetFileFilters() FileFilteringPolicy {
	if o == nil || o.FileFilters == nil {
		var ret FileFilteringPolicy
		return ret
	}
	return *o.FileFilters
}

// GetFileFiltersOk returns a tuple with the FileFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericNasObjectProtectionResponseParams) GetFileFiltersOk() (*FileFilteringPolicy, bool) {
	if o == nil || o.FileFilters == nil {
		return nil, false
	}
	return o.FileFilters, true
}

// HasFileFilters returns a boolean if a field has been set.
func (o *GenericNasObjectProtectionResponseParams) HasFileFilters() bool {
	if o != nil && o.FileFilters != nil {
		return true
	}

	return false
}

// SetFileFilters gets a reference to the given FileFilteringPolicy and assigns it to the FileFilters field.
func (o *GenericNasObjectProtectionResponseParams) SetFileFilters(v FileFilteringPolicy) {
	o.FileFilters = &v
}

// GetPrePostScript returns the PrePostScript field value if set, zero value otherwise.
func (o *GenericNasObjectProtectionResponseParams) GetPrePostScript() HostBasedBackupScriptParams {
	if o == nil || o.PrePostScript == nil {
		var ret HostBasedBackupScriptParams
		return ret
	}
	return *o.PrePostScript
}

// GetPrePostScriptOk returns a tuple with the PrePostScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericNasObjectProtectionResponseParams) GetPrePostScriptOk() (*HostBasedBackupScriptParams, bool) {
	if o == nil || o.PrePostScript == nil {
		return nil, false
	}
	return o.PrePostScript, true
}

// HasPrePostScript returns a boolean if a field has been set.
func (o *GenericNasObjectProtectionResponseParams) HasPrePostScript() bool {
	if o != nil && o.PrePostScript != nil {
		return true
	}

	return false
}

// SetPrePostScript gets a reference to the given HostBasedBackupScriptParams and assigns it to the PrePostScript field.
func (o *GenericNasObjectProtectionResponseParams) SetPrePostScript(v HostBasedBackupScriptParams) {
	o.PrePostScript = &v
}

// GetThrottlingConfig returns the ThrottlingConfig field value if set, zero value otherwise.
func (o *GenericNasObjectProtectionResponseParams) GetThrottlingConfig() NasThrottlingConfig {
	if o == nil || o.ThrottlingConfig == nil {
		var ret NasThrottlingConfig
		return ret
	}
	return *o.ThrottlingConfig
}

// GetThrottlingConfigOk returns a tuple with the ThrottlingConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericNasObjectProtectionResponseParams) GetThrottlingConfigOk() (*NasThrottlingConfig, bool) {
	if o == nil || o.ThrottlingConfig == nil {
		return nil, false
	}
	return o.ThrottlingConfig, true
}

// HasThrottlingConfig returns a boolean if a field has been set.
func (o *GenericNasObjectProtectionResponseParams) HasThrottlingConfig() bool {
	if o != nil && o.ThrottlingConfig != nil {
		return true
	}

	return false
}

// SetThrottlingConfig gets a reference to the given NasThrottlingConfig and assigns it to the ThrottlingConfig field.
func (o *GenericNasObjectProtectionResponseParams) SetThrottlingConfig(v NasThrottlingConfig) {
	o.ThrottlingConfig = &v
}

func (o GenericNasObjectProtectionResponseParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IndexingPolicy != nil {
		toSerialize["indexingPolicy"] = o.IndexingPolicy
	}
	if o.ContinueOnError.IsSet() {
		toSerialize["continueOnError"] = o.ContinueOnError.Get()
	}
	if o.EncryptionEnabled.IsSet() {
		toSerialize["encryptionEnabled"] = o.EncryptionEnabled.Get()
	}
	if o.FileLockConfig != nil {
		toSerialize["fileLockConfig"] = o.FileLockConfig
	}
	if o.FileFilters != nil {
		toSerialize["fileFilters"] = o.FileFilters
	}
	if o.PrePostScript != nil {
		toSerialize["prePostScript"] = o.PrePostScript
	}
	if o.ThrottlingConfig != nil {
		toSerialize["throttlingConfig"] = o.ThrottlingConfig
	}
	return json.Marshal(toSerialize)
}

type NullableGenericNasObjectProtectionResponseParams struct {
	value *GenericNasObjectProtectionResponseParams
	isSet bool
}

func (v NullableGenericNasObjectProtectionResponseParams) Get() *GenericNasObjectProtectionResponseParams {
	return v.value
}

func (v *NullableGenericNasObjectProtectionResponseParams) Set(val *GenericNasObjectProtectionResponseParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericNasObjectProtectionResponseParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericNasObjectProtectionResponseParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericNasObjectProtectionResponseParams(val *GenericNasObjectProtectionResponseParams) *NullableGenericNasObjectProtectionResponseParams {
	return &NullableGenericNasObjectProtectionResponseParams{value: val, isSet: true}
}

func (v NullableGenericNasObjectProtectionResponseParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericNasObjectProtectionResponseParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


