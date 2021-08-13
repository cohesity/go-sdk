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

// ElastifileObjectProtectionResponseParams Specifies the parameters which are specific to Elastifile object protection.
type ElastifileObjectProtectionResponseParams struct {
	IndexingPolicy *IndexingPolicy `json:"indexingPolicy,omitempty"`
	// Specifies whether or not the backup should continue regardless of whether or not an error was encountered.
	ContinueOnError NullableBool `json:"continueOnError,omitempty"`
	// Specifies whether the encryption should be used while backup or not.
	EncryptionEnabled NullableBool `json:"encryptionEnabled,omitempty"`
	FileLockConfig *FileLevelDataLockConfig `json:"fileLockConfig,omitempty"`
	FileFilters *FileFilteringPolicy `json:"fileFilters,omitempty"`
	PrePostScript *HostBasedBackupScriptParams `json:"prePostScript,omitempty"`
	ThrottlingConfig *NasThrottlingConfig `json:"throttlingConfig,omitempty"`
	// Specifies the protocol of the NAS device being backed up.
	Protocol NullableString `json:"protocol,omitempty"`
}

// NewElastifileObjectProtectionResponseParams instantiates a new ElastifileObjectProtectionResponseParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewElastifileObjectProtectionResponseParams() *ElastifileObjectProtectionResponseParams {
	this := ElastifileObjectProtectionResponseParams{}
	return &this
}

// NewElastifileObjectProtectionResponseParamsWithDefaults instantiates a new ElastifileObjectProtectionResponseParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewElastifileObjectProtectionResponseParamsWithDefaults() *ElastifileObjectProtectionResponseParams {
	this := ElastifileObjectProtectionResponseParams{}
	return &this
}

// GetIndexingPolicy returns the IndexingPolicy field value if set, zero value otherwise.
func (o *ElastifileObjectProtectionResponseParams) GetIndexingPolicy() IndexingPolicy {
	if o == nil || o.IndexingPolicy == nil {
		var ret IndexingPolicy
		return ret
	}
	return *o.IndexingPolicy
}

// GetIndexingPolicyOk returns a tuple with the IndexingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElastifileObjectProtectionResponseParams) GetIndexingPolicyOk() (*IndexingPolicy, bool) {
	if o == nil || o.IndexingPolicy == nil {
		return nil, false
	}
	return o.IndexingPolicy, true
}

// HasIndexingPolicy returns a boolean if a field has been set.
func (o *ElastifileObjectProtectionResponseParams) HasIndexingPolicy() bool {
	if o != nil && o.IndexingPolicy != nil {
		return true
	}

	return false
}

// SetIndexingPolicy gets a reference to the given IndexingPolicy and assigns it to the IndexingPolicy field.
func (o *ElastifileObjectProtectionResponseParams) SetIndexingPolicy(v IndexingPolicy) {
	o.IndexingPolicy = &v
}

// GetContinueOnError returns the ContinueOnError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElastifileObjectProtectionResponseParams) GetContinueOnError() bool {
	if o == nil || o.ContinueOnError.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ContinueOnError.Get()
}

// GetContinueOnErrorOk returns a tuple with the ContinueOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ElastifileObjectProtectionResponseParams) GetContinueOnErrorOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContinueOnError.Get(), o.ContinueOnError.IsSet()
}

// HasContinueOnError returns a boolean if a field has been set.
func (o *ElastifileObjectProtectionResponseParams) HasContinueOnError() bool {
	if o != nil && o.ContinueOnError.IsSet() {
		return true
	}

	return false
}

// SetContinueOnError gets a reference to the given NullableBool and assigns it to the ContinueOnError field.
func (o *ElastifileObjectProtectionResponseParams) SetContinueOnError(v bool) {
	o.ContinueOnError.Set(&v)
}
// SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil
func (o *ElastifileObjectProtectionResponseParams) SetContinueOnErrorNil() {
	o.ContinueOnError.Set(nil)
}

// UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
func (o *ElastifileObjectProtectionResponseParams) UnsetContinueOnError() {
	o.ContinueOnError.Unset()
}

// GetEncryptionEnabled returns the EncryptionEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElastifileObjectProtectionResponseParams) GetEncryptionEnabled() bool {
	if o == nil || o.EncryptionEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.EncryptionEnabled.Get()
}

// GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ElastifileObjectProtectionResponseParams) GetEncryptionEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EncryptionEnabled.Get(), o.EncryptionEnabled.IsSet()
}

// HasEncryptionEnabled returns a boolean if a field has been set.
func (o *ElastifileObjectProtectionResponseParams) HasEncryptionEnabled() bool {
	if o != nil && o.EncryptionEnabled.IsSet() {
		return true
	}

	return false
}

// SetEncryptionEnabled gets a reference to the given NullableBool and assigns it to the EncryptionEnabled field.
func (o *ElastifileObjectProtectionResponseParams) SetEncryptionEnabled(v bool) {
	o.EncryptionEnabled.Set(&v)
}
// SetEncryptionEnabledNil sets the value for EncryptionEnabled to be an explicit nil
func (o *ElastifileObjectProtectionResponseParams) SetEncryptionEnabledNil() {
	o.EncryptionEnabled.Set(nil)
}

// UnsetEncryptionEnabled ensures that no value is present for EncryptionEnabled, not even an explicit nil
func (o *ElastifileObjectProtectionResponseParams) UnsetEncryptionEnabled() {
	o.EncryptionEnabled.Unset()
}

// GetFileLockConfig returns the FileLockConfig field value if set, zero value otherwise.
func (o *ElastifileObjectProtectionResponseParams) GetFileLockConfig() FileLevelDataLockConfig {
	if o == nil || o.FileLockConfig == nil {
		var ret FileLevelDataLockConfig
		return ret
	}
	return *o.FileLockConfig
}

// GetFileLockConfigOk returns a tuple with the FileLockConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElastifileObjectProtectionResponseParams) GetFileLockConfigOk() (*FileLevelDataLockConfig, bool) {
	if o == nil || o.FileLockConfig == nil {
		return nil, false
	}
	return o.FileLockConfig, true
}

// HasFileLockConfig returns a boolean if a field has been set.
func (o *ElastifileObjectProtectionResponseParams) HasFileLockConfig() bool {
	if o != nil && o.FileLockConfig != nil {
		return true
	}

	return false
}

// SetFileLockConfig gets a reference to the given FileLevelDataLockConfig and assigns it to the FileLockConfig field.
func (o *ElastifileObjectProtectionResponseParams) SetFileLockConfig(v FileLevelDataLockConfig) {
	o.FileLockConfig = &v
}

// GetFileFilters returns the FileFilters field value if set, zero value otherwise.
func (o *ElastifileObjectProtectionResponseParams) GetFileFilters() FileFilteringPolicy {
	if o == nil || o.FileFilters == nil {
		var ret FileFilteringPolicy
		return ret
	}
	return *o.FileFilters
}

// GetFileFiltersOk returns a tuple with the FileFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElastifileObjectProtectionResponseParams) GetFileFiltersOk() (*FileFilteringPolicy, bool) {
	if o == nil || o.FileFilters == nil {
		return nil, false
	}
	return o.FileFilters, true
}

// HasFileFilters returns a boolean if a field has been set.
func (o *ElastifileObjectProtectionResponseParams) HasFileFilters() bool {
	if o != nil && o.FileFilters != nil {
		return true
	}

	return false
}

// SetFileFilters gets a reference to the given FileFilteringPolicy and assigns it to the FileFilters field.
func (o *ElastifileObjectProtectionResponseParams) SetFileFilters(v FileFilteringPolicy) {
	o.FileFilters = &v
}

// GetPrePostScript returns the PrePostScript field value if set, zero value otherwise.
func (o *ElastifileObjectProtectionResponseParams) GetPrePostScript() HostBasedBackupScriptParams {
	if o == nil || o.PrePostScript == nil {
		var ret HostBasedBackupScriptParams
		return ret
	}
	return *o.PrePostScript
}

// GetPrePostScriptOk returns a tuple with the PrePostScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElastifileObjectProtectionResponseParams) GetPrePostScriptOk() (*HostBasedBackupScriptParams, bool) {
	if o == nil || o.PrePostScript == nil {
		return nil, false
	}
	return o.PrePostScript, true
}

// HasPrePostScript returns a boolean if a field has been set.
func (o *ElastifileObjectProtectionResponseParams) HasPrePostScript() bool {
	if o != nil && o.PrePostScript != nil {
		return true
	}

	return false
}

// SetPrePostScript gets a reference to the given HostBasedBackupScriptParams and assigns it to the PrePostScript field.
func (o *ElastifileObjectProtectionResponseParams) SetPrePostScript(v HostBasedBackupScriptParams) {
	o.PrePostScript = &v
}

// GetThrottlingConfig returns the ThrottlingConfig field value if set, zero value otherwise.
func (o *ElastifileObjectProtectionResponseParams) GetThrottlingConfig() NasThrottlingConfig {
	if o == nil || o.ThrottlingConfig == nil {
		var ret NasThrottlingConfig
		return ret
	}
	return *o.ThrottlingConfig
}

// GetThrottlingConfigOk returns a tuple with the ThrottlingConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElastifileObjectProtectionResponseParams) GetThrottlingConfigOk() (*NasThrottlingConfig, bool) {
	if o == nil || o.ThrottlingConfig == nil {
		return nil, false
	}
	return o.ThrottlingConfig, true
}

// HasThrottlingConfig returns a boolean if a field has been set.
func (o *ElastifileObjectProtectionResponseParams) HasThrottlingConfig() bool {
	if o != nil && o.ThrottlingConfig != nil {
		return true
	}

	return false
}

// SetThrottlingConfig gets a reference to the given NasThrottlingConfig and assigns it to the ThrottlingConfig field.
func (o *ElastifileObjectProtectionResponseParams) SetThrottlingConfig(v NasThrottlingConfig) {
	o.ThrottlingConfig = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElastifileObjectProtectionResponseParams) GetProtocol() string {
	if o == nil || o.Protocol.Get() == nil {
		var ret string
		return ret
	}
	return *o.Protocol.Get()
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ElastifileObjectProtectionResponseParams) GetProtocolOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Protocol.Get(), o.Protocol.IsSet()
}

// HasProtocol returns a boolean if a field has been set.
func (o *ElastifileObjectProtectionResponseParams) HasProtocol() bool {
	if o != nil && o.Protocol.IsSet() {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given NullableString and assigns it to the Protocol field.
func (o *ElastifileObjectProtectionResponseParams) SetProtocol(v string) {
	o.Protocol.Set(&v)
}
// SetProtocolNil sets the value for Protocol to be an explicit nil
func (o *ElastifileObjectProtectionResponseParams) SetProtocolNil() {
	o.Protocol.Set(nil)
}

// UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
func (o *ElastifileObjectProtectionResponseParams) UnsetProtocol() {
	o.Protocol.Unset()
}

func (o ElastifileObjectProtectionResponseParams) MarshalJSON() ([]byte, error) {
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
	if o.Protocol.IsSet() {
		toSerialize["protocol"] = o.Protocol.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableElastifileObjectProtectionResponseParams struct {
	value *ElastifileObjectProtectionResponseParams
	isSet bool
}

func (v NullableElastifileObjectProtectionResponseParams) Get() *ElastifileObjectProtectionResponseParams {
	return v.value
}

func (v *NullableElastifileObjectProtectionResponseParams) Set(val *ElastifileObjectProtectionResponseParams) {
	v.value = val
	v.isSet = true
}

func (v NullableElastifileObjectProtectionResponseParams) IsSet() bool {
	return v.isSet
}

func (v *NullableElastifileObjectProtectionResponseParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableElastifileObjectProtectionResponseParams(val *ElastifileObjectProtectionResponseParams) *NullableElastifileObjectProtectionResponseParams {
	return &NullableElastifileObjectProtectionResponseParams{value: val, isSet: true}
}

func (v NullableElastifileObjectProtectionResponseParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableElastifileObjectProtectionResponseParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o ElastifileObjectProtectionResponseParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}