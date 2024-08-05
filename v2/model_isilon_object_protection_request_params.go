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

// checks if the IsilonObjectProtectionRequestParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IsilonObjectProtectionRequestParams{}

// IsilonObjectProtectionRequestParams Specifies the parameters which are specific to Isilon object protection.
type IsilonObjectProtectionRequestParams struct {
	// Specifies whether or not the backup should continue regardless of whether or not an error was encountered.
	ContinueOnError NullableBool `json:"continueOnError,omitempty"`
	// Specifies whether the encryption should be used while backup or not.
	EncryptionEnabled NullableBool `json:"encryptionEnabled,omitempty"`
	FileFilters *FileFilteringPolicy `json:"fileFilters,omitempty"`
	FileLockConfig *FileLevelDataLockConfig `json:"fileLockConfig,omitempty"`
	IndexingPolicy *IndexingPolicy `json:"indexingPolicy,omitempty"`
	PrePostScript *HostBasedBackupScriptParams `json:"prePostScript,omitempty"`
	ThrottlingConfig *NasThrottlingConfig `json:"throttlingConfig,omitempty"`
	ContinuousSnapshots *ContinuousSnapshotParams `json:"continuousSnapshots,omitempty"`
	// Specifies the protocol of the NAS device being backed up.
	Protocol NullableString `json:"protocol,omitempty"`
	// Specify whether to use the Isilon Changelist API to directly discover changed files/directories for faster incremental backup. Cohesity will keep an extra snapshot which will be deleted by the next successful backup.
	UseChangelist NullableBool `json:"useChangelist,omitempty"`
	// Specifies the objects to be protected.
	Objects []ProtectionObjectInput `json:"objects"`
}

type _IsilonObjectProtectionRequestParams IsilonObjectProtectionRequestParams

// NewIsilonObjectProtectionRequestParams instantiates a new IsilonObjectProtectionRequestParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIsilonObjectProtectionRequestParams(objects []ProtectionObjectInput) *IsilonObjectProtectionRequestParams {
	this := IsilonObjectProtectionRequestParams{}
	this.Objects = objects
	return &this
}

// NewIsilonObjectProtectionRequestParamsWithDefaults instantiates a new IsilonObjectProtectionRequestParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIsilonObjectProtectionRequestParamsWithDefaults() *IsilonObjectProtectionRequestParams {
	this := IsilonObjectProtectionRequestParams{}
	return &this
}

// GetContinueOnError returns the ContinueOnError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IsilonObjectProtectionRequestParams) GetContinueOnError() bool {
	if o == nil || IsNil(o.ContinueOnError.Get()) {
		var ret bool
		return ret
	}
	return *o.ContinueOnError.Get()
}

// GetContinueOnErrorOk returns a tuple with the ContinueOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IsilonObjectProtectionRequestParams) GetContinueOnErrorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContinueOnError.Get(), o.ContinueOnError.IsSet()
}

// HasContinueOnError returns a boolean if a field has been set.
func (o *IsilonObjectProtectionRequestParams) HasContinueOnError() bool {
	if o != nil && o.ContinueOnError.IsSet() {
		return true
	}

	return false
}

// SetContinueOnError gets a reference to the given NullableBool and assigns it to the ContinueOnError field.
func (o *IsilonObjectProtectionRequestParams) SetContinueOnError(v bool) {
	o.ContinueOnError.Set(&v)
}
// SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil
func (o *IsilonObjectProtectionRequestParams) SetContinueOnErrorNil() {
	o.ContinueOnError.Set(nil)
}

// UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
func (o *IsilonObjectProtectionRequestParams) UnsetContinueOnError() {
	o.ContinueOnError.Unset()
}

// GetEncryptionEnabled returns the EncryptionEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IsilonObjectProtectionRequestParams) GetEncryptionEnabled() bool {
	if o == nil || IsNil(o.EncryptionEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.EncryptionEnabled.Get()
}

// GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IsilonObjectProtectionRequestParams) GetEncryptionEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EncryptionEnabled.Get(), o.EncryptionEnabled.IsSet()
}

// HasEncryptionEnabled returns a boolean if a field has been set.
func (o *IsilonObjectProtectionRequestParams) HasEncryptionEnabled() bool {
	if o != nil && o.EncryptionEnabled.IsSet() {
		return true
	}

	return false
}

// SetEncryptionEnabled gets a reference to the given NullableBool and assigns it to the EncryptionEnabled field.
func (o *IsilonObjectProtectionRequestParams) SetEncryptionEnabled(v bool) {
	o.EncryptionEnabled.Set(&v)
}
// SetEncryptionEnabledNil sets the value for EncryptionEnabled to be an explicit nil
func (o *IsilonObjectProtectionRequestParams) SetEncryptionEnabledNil() {
	o.EncryptionEnabled.Set(nil)
}

// UnsetEncryptionEnabled ensures that no value is present for EncryptionEnabled, not even an explicit nil
func (o *IsilonObjectProtectionRequestParams) UnsetEncryptionEnabled() {
	o.EncryptionEnabled.Unset()
}

// GetFileFilters returns the FileFilters field value if set, zero value otherwise.
func (o *IsilonObjectProtectionRequestParams) GetFileFilters() FileFilteringPolicy {
	if o == nil || IsNil(o.FileFilters) {
		var ret FileFilteringPolicy
		return ret
	}
	return *o.FileFilters
}

// GetFileFiltersOk returns a tuple with the FileFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsilonObjectProtectionRequestParams) GetFileFiltersOk() (*FileFilteringPolicy, bool) {
	if o == nil || IsNil(o.FileFilters) {
		return nil, false
	}
	return o.FileFilters, true
}

// HasFileFilters returns a boolean if a field has been set.
func (o *IsilonObjectProtectionRequestParams) HasFileFilters() bool {
	if o != nil && !IsNil(o.FileFilters) {
		return true
	}

	return false
}

// SetFileFilters gets a reference to the given FileFilteringPolicy and assigns it to the FileFilters field.
func (o *IsilonObjectProtectionRequestParams) SetFileFilters(v FileFilteringPolicy) {
	o.FileFilters = &v
}

// GetFileLockConfig returns the FileLockConfig field value if set, zero value otherwise.
func (o *IsilonObjectProtectionRequestParams) GetFileLockConfig() FileLevelDataLockConfig {
	if o == nil || IsNil(o.FileLockConfig) {
		var ret FileLevelDataLockConfig
		return ret
	}
	return *o.FileLockConfig
}

// GetFileLockConfigOk returns a tuple with the FileLockConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsilonObjectProtectionRequestParams) GetFileLockConfigOk() (*FileLevelDataLockConfig, bool) {
	if o == nil || IsNil(o.FileLockConfig) {
		return nil, false
	}
	return o.FileLockConfig, true
}

// HasFileLockConfig returns a boolean if a field has been set.
func (o *IsilonObjectProtectionRequestParams) HasFileLockConfig() bool {
	if o != nil && !IsNil(o.FileLockConfig) {
		return true
	}

	return false
}

// SetFileLockConfig gets a reference to the given FileLevelDataLockConfig and assigns it to the FileLockConfig field.
func (o *IsilonObjectProtectionRequestParams) SetFileLockConfig(v FileLevelDataLockConfig) {
	o.FileLockConfig = &v
}

// GetIndexingPolicy returns the IndexingPolicy field value if set, zero value otherwise.
func (o *IsilonObjectProtectionRequestParams) GetIndexingPolicy() IndexingPolicy {
	if o == nil || IsNil(o.IndexingPolicy) {
		var ret IndexingPolicy
		return ret
	}
	return *o.IndexingPolicy
}

// GetIndexingPolicyOk returns a tuple with the IndexingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsilonObjectProtectionRequestParams) GetIndexingPolicyOk() (*IndexingPolicy, bool) {
	if o == nil || IsNil(o.IndexingPolicy) {
		return nil, false
	}
	return o.IndexingPolicy, true
}

// HasIndexingPolicy returns a boolean if a field has been set.
func (o *IsilonObjectProtectionRequestParams) HasIndexingPolicy() bool {
	if o != nil && !IsNil(o.IndexingPolicy) {
		return true
	}

	return false
}

// SetIndexingPolicy gets a reference to the given IndexingPolicy and assigns it to the IndexingPolicy field.
func (o *IsilonObjectProtectionRequestParams) SetIndexingPolicy(v IndexingPolicy) {
	o.IndexingPolicy = &v
}

// GetPrePostScript returns the PrePostScript field value if set, zero value otherwise.
func (o *IsilonObjectProtectionRequestParams) GetPrePostScript() HostBasedBackupScriptParams {
	if o == nil || IsNil(o.PrePostScript) {
		var ret HostBasedBackupScriptParams
		return ret
	}
	return *o.PrePostScript
}

// GetPrePostScriptOk returns a tuple with the PrePostScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsilonObjectProtectionRequestParams) GetPrePostScriptOk() (*HostBasedBackupScriptParams, bool) {
	if o == nil || IsNil(o.PrePostScript) {
		return nil, false
	}
	return o.PrePostScript, true
}

// HasPrePostScript returns a boolean if a field has been set.
func (o *IsilonObjectProtectionRequestParams) HasPrePostScript() bool {
	if o != nil && !IsNil(o.PrePostScript) {
		return true
	}

	return false
}

// SetPrePostScript gets a reference to the given HostBasedBackupScriptParams and assigns it to the PrePostScript field.
func (o *IsilonObjectProtectionRequestParams) SetPrePostScript(v HostBasedBackupScriptParams) {
	o.PrePostScript = &v
}

// GetThrottlingConfig returns the ThrottlingConfig field value if set, zero value otherwise.
func (o *IsilonObjectProtectionRequestParams) GetThrottlingConfig() NasThrottlingConfig {
	if o == nil || IsNil(o.ThrottlingConfig) {
		var ret NasThrottlingConfig
		return ret
	}
	return *o.ThrottlingConfig
}

// GetThrottlingConfigOk returns a tuple with the ThrottlingConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsilonObjectProtectionRequestParams) GetThrottlingConfigOk() (*NasThrottlingConfig, bool) {
	if o == nil || IsNil(o.ThrottlingConfig) {
		return nil, false
	}
	return o.ThrottlingConfig, true
}

// HasThrottlingConfig returns a boolean if a field has been set.
func (o *IsilonObjectProtectionRequestParams) HasThrottlingConfig() bool {
	if o != nil && !IsNil(o.ThrottlingConfig) {
		return true
	}

	return false
}

// SetThrottlingConfig gets a reference to the given NasThrottlingConfig and assigns it to the ThrottlingConfig field.
func (o *IsilonObjectProtectionRequestParams) SetThrottlingConfig(v NasThrottlingConfig) {
	o.ThrottlingConfig = &v
}

// GetContinuousSnapshots returns the ContinuousSnapshots field value if set, zero value otherwise.
func (o *IsilonObjectProtectionRequestParams) GetContinuousSnapshots() ContinuousSnapshotParams {
	if o == nil || IsNil(o.ContinuousSnapshots) {
		var ret ContinuousSnapshotParams
		return ret
	}
	return *o.ContinuousSnapshots
}

// GetContinuousSnapshotsOk returns a tuple with the ContinuousSnapshots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsilonObjectProtectionRequestParams) GetContinuousSnapshotsOk() (*ContinuousSnapshotParams, bool) {
	if o == nil || IsNil(o.ContinuousSnapshots) {
		return nil, false
	}
	return o.ContinuousSnapshots, true
}

// HasContinuousSnapshots returns a boolean if a field has been set.
func (o *IsilonObjectProtectionRequestParams) HasContinuousSnapshots() bool {
	if o != nil && !IsNil(o.ContinuousSnapshots) {
		return true
	}

	return false
}

// SetContinuousSnapshots gets a reference to the given ContinuousSnapshotParams and assigns it to the ContinuousSnapshots field.
func (o *IsilonObjectProtectionRequestParams) SetContinuousSnapshots(v ContinuousSnapshotParams) {
	o.ContinuousSnapshots = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IsilonObjectProtectionRequestParams) GetProtocol() string {
	if o == nil || IsNil(o.Protocol.Get()) {
		var ret string
		return ret
	}
	return *o.Protocol.Get()
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IsilonObjectProtectionRequestParams) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Protocol.Get(), o.Protocol.IsSet()
}

// HasProtocol returns a boolean if a field has been set.
func (o *IsilonObjectProtectionRequestParams) HasProtocol() bool {
	if o != nil && o.Protocol.IsSet() {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given NullableString and assigns it to the Protocol field.
func (o *IsilonObjectProtectionRequestParams) SetProtocol(v string) {
	o.Protocol.Set(&v)
}
// SetProtocolNil sets the value for Protocol to be an explicit nil
func (o *IsilonObjectProtectionRequestParams) SetProtocolNil() {
	o.Protocol.Set(nil)
}

// UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
func (o *IsilonObjectProtectionRequestParams) UnsetProtocol() {
	o.Protocol.Unset()
}

// GetUseChangelist returns the UseChangelist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IsilonObjectProtectionRequestParams) GetUseChangelist() bool {
	if o == nil || IsNil(o.UseChangelist.Get()) {
		var ret bool
		return ret
	}
	return *o.UseChangelist.Get()
}

// GetUseChangelistOk returns a tuple with the UseChangelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IsilonObjectProtectionRequestParams) GetUseChangelistOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.UseChangelist.Get(), o.UseChangelist.IsSet()
}

// HasUseChangelist returns a boolean if a field has been set.
func (o *IsilonObjectProtectionRequestParams) HasUseChangelist() bool {
	if o != nil && o.UseChangelist.IsSet() {
		return true
	}

	return false
}

// SetUseChangelist gets a reference to the given NullableBool and assigns it to the UseChangelist field.
func (o *IsilonObjectProtectionRequestParams) SetUseChangelist(v bool) {
	o.UseChangelist.Set(&v)
}
// SetUseChangelistNil sets the value for UseChangelist to be an explicit nil
func (o *IsilonObjectProtectionRequestParams) SetUseChangelistNil() {
	o.UseChangelist.Set(nil)
}

// UnsetUseChangelist ensures that no value is present for UseChangelist, not even an explicit nil
func (o *IsilonObjectProtectionRequestParams) UnsetUseChangelist() {
	o.UseChangelist.Unset()
}

// GetObjects returns the Objects field value
func (o *IsilonObjectProtectionRequestParams) GetObjects() []ProtectionObjectInput {
	if o == nil {
		var ret []ProtectionObjectInput
		return ret
	}

	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value
// and a boolean to check if the value has been set.
func (o *IsilonObjectProtectionRequestParams) GetObjectsOk() ([]ProtectionObjectInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Objects, true
}

// SetObjects sets field value
func (o *IsilonObjectProtectionRequestParams) SetObjects(v []ProtectionObjectInput) {
	o.Objects = v
}

func (o IsilonObjectProtectionRequestParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IsilonObjectProtectionRequestParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ContinueOnError.IsSet() {
		toSerialize["continueOnError"] = o.ContinueOnError.Get()
	}
	if o.EncryptionEnabled.IsSet() {
		toSerialize["encryptionEnabled"] = o.EncryptionEnabled.Get()
	}
	if !IsNil(o.FileFilters) {
		toSerialize["fileFilters"] = o.FileFilters
	}
	if !IsNil(o.FileLockConfig) {
		toSerialize["fileLockConfig"] = o.FileLockConfig
	}
	if !IsNil(o.IndexingPolicy) {
		toSerialize["indexingPolicy"] = o.IndexingPolicy
	}
	if !IsNil(o.PrePostScript) {
		toSerialize["prePostScript"] = o.PrePostScript
	}
	if !IsNil(o.ThrottlingConfig) {
		toSerialize["throttlingConfig"] = o.ThrottlingConfig
	}
	if !IsNil(o.ContinuousSnapshots) {
		toSerialize["continuousSnapshots"] = o.ContinuousSnapshots
	}
	if o.Protocol.IsSet() {
		toSerialize["protocol"] = o.Protocol.Get()
	}
	if o.UseChangelist.IsSet() {
		toSerialize["useChangelist"] = o.UseChangelist.Get()
	}
	toSerialize["objects"] = o.Objects
	return toSerialize, nil
}

func (o *IsilonObjectProtectionRequestParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objects",
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

	varIsilonObjectProtectionRequestParams := _IsilonObjectProtectionRequestParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIsilonObjectProtectionRequestParams)

	if err != nil {
		return err
	}

	*o = IsilonObjectProtectionRequestParams(varIsilonObjectProtectionRequestParams)

	return err
}

type NullableIsilonObjectProtectionRequestParams struct {
	value *IsilonObjectProtectionRequestParams
	isSet bool
}

func (v NullableIsilonObjectProtectionRequestParams) Get() *IsilonObjectProtectionRequestParams {
	return v.value
}

func (v *NullableIsilonObjectProtectionRequestParams) Set(val *IsilonObjectProtectionRequestParams) {
	v.value = val
	v.isSet = true
}

func (v NullableIsilonObjectProtectionRequestParams) IsSet() bool {
	return v.isSet
}

func (v *NullableIsilonObjectProtectionRequestParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIsilonObjectProtectionRequestParams(val *IsilonObjectProtectionRequestParams) *NullableIsilonObjectProtectionRequestParams {
	return &NullableIsilonObjectProtectionRequestParams{value: val, isSet: true}
}

func (v NullableIsilonObjectProtectionRequestParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIsilonObjectProtectionRequestParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


