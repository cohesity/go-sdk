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

// GpfsProtectionGroupParams Specifies the parameters which are specific to GPFS related Protection Groups.
type GpfsProtectionGroupParams struct {
	// Specifies the objects to be included in the Protection Group.
	Objects []GpfsProtectionGroupObjectParams `json:"objects"`
	// Specifies whether or not to store the snapshots in this run directly in an Archive Target instead of on the Cluster. If this is set to true, the associated policy must have exactly one Archive Target associated with it and the policy must be set up to archive after every run. Also, a Storage Domain cannot be specified. Default behavior is 'false'.
	DirectCloudArchive NullableBool `json:"directCloudArchive,omitempty"`
	// Specifies whether or not to enable native format for direct archive job. This field is set to true if native format should be used for archiving.
	NativeFormat NullableBool `json:"nativeFormat,omitempty"`
	IndexingPolicy *IndexingPolicy `json:"indexingPolicy,omitempty"`
	// Specifies the preferred protocol to use if this device supports multiple protocols.
	Protocol NullableString `json:"protocol,omitempty"`
	// Specifies whether or not the Protection Group should continue regardless of whether or not an error was encountered during protection group run.
	ContinueOnError NullableBool `json:"continueOnError,omitempty"`
	// Specifies whether the protection group should use encryption while backup or not.
	EncryptionEnabled NullableBool `json:"encryptionEnabled,omitempty"`
	FileLockConfig *FileLevelDataLockConfig `json:"fileLockConfig,omitempty"`
	FileFilters *FileFilteringPolicy `json:"fileFilters,omitempty"`
	// Specifies the id of the parent of the objects.
	SourceId NullableInt64 `json:"sourceId,omitempty"`
	// Specifies the name of the parent of the objects.
	SourceName NullableString `json:"sourceName,omitempty"`
	PrePostScript *HostBasedBackupScriptParams `json:"prePostScript,omitempty"`
	FilterIpConfig *FilterIpConfig `json:"filterIpConfig,omitempty"`
	ThrottlingConfig *NasThrottlingConfig `json:"throttlingConfig,omitempty"`
}

// NewGpfsProtectionGroupParams instantiates a new GpfsProtectionGroupParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGpfsProtectionGroupParams(objects []GpfsProtectionGroupObjectParams) *GpfsProtectionGroupParams {
	this := GpfsProtectionGroupParams{}
	this.Objects = objects
	return &this
}

// NewGpfsProtectionGroupParamsWithDefaults instantiates a new GpfsProtectionGroupParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGpfsProtectionGroupParamsWithDefaults() *GpfsProtectionGroupParams {
	this := GpfsProtectionGroupParams{}
	return &this
}

// GetObjects returns the Objects field value
func (o *GpfsProtectionGroupParams) GetObjects() []GpfsProtectionGroupObjectParams {
	if o == nil {
		var ret []GpfsProtectionGroupObjectParams
		return ret
	}

	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value
// and a boolean to check if the value has been set.
func (o *GpfsProtectionGroupParams) GetObjectsOk() (*[]GpfsProtectionGroupObjectParams, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Objects, true
}

// SetObjects sets field value
func (o *GpfsProtectionGroupParams) SetObjects(v []GpfsProtectionGroupObjectParams) {
	o.Objects = v
}

// GetDirectCloudArchive returns the DirectCloudArchive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GpfsProtectionGroupParams) GetDirectCloudArchive() bool {
	if o == nil || o.DirectCloudArchive.Get() == nil {
		var ret bool
		return ret
	}
	return *o.DirectCloudArchive.Get()
}

// GetDirectCloudArchiveOk returns a tuple with the DirectCloudArchive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GpfsProtectionGroupParams) GetDirectCloudArchiveOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DirectCloudArchive.Get(), o.DirectCloudArchive.IsSet()
}

// HasDirectCloudArchive returns a boolean if a field has been set.
func (o *GpfsProtectionGroupParams) HasDirectCloudArchive() bool {
	if o != nil && o.DirectCloudArchive.IsSet() {
		return true
	}

	return false
}

// SetDirectCloudArchive gets a reference to the given NullableBool and assigns it to the DirectCloudArchive field.
func (o *GpfsProtectionGroupParams) SetDirectCloudArchive(v bool) {
	o.DirectCloudArchive.Set(&v)
}
// SetDirectCloudArchiveNil sets the value for DirectCloudArchive to be an explicit nil
func (o *GpfsProtectionGroupParams) SetDirectCloudArchiveNil() {
	o.DirectCloudArchive.Set(nil)
}

// UnsetDirectCloudArchive ensures that no value is present for DirectCloudArchive, not even an explicit nil
func (o *GpfsProtectionGroupParams) UnsetDirectCloudArchive() {
	o.DirectCloudArchive.Unset()
}

// GetNativeFormat returns the NativeFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GpfsProtectionGroupParams) GetNativeFormat() bool {
	if o == nil || o.NativeFormat.Get() == nil {
		var ret bool
		return ret
	}
	return *o.NativeFormat.Get()
}

// GetNativeFormatOk returns a tuple with the NativeFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GpfsProtectionGroupParams) GetNativeFormatOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NativeFormat.Get(), o.NativeFormat.IsSet()
}

// HasNativeFormat returns a boolean if a field has been set.
func (o *GpfsProtectionGroupParams) HasNativeFormat() bool {
	if o != nil && o.NativeFormat.IsSet() {
		return true
	}

	return false
}

// SetNativeFormat gets a reference to the given NullableBool and assigns it to the NativeFormat field.
func (o *GpfsProtectionGroupParams) SetNativeFormat(v bool) {
	o.NativeFormat.Set(&v)
}
// SetNativeFormatNil sets the value for NativeFormat to be an explicit nil
func (o *GpfsProtectionGroupParams) SetNativeFormatNil() {
	o.NativeFormat.Set(nil)
}

// UnsetNativeFormat ensures that no value is present for NativeFormat, not even an explicit nil
func (o *GpfsProtectionGroupParams) UnsetNativeFormat() {
	o.NativeFormat.Unset()
}

// GetIndexingPolicy returns the IndexingPolicy field value if set, zero value otherwise.
func (o *GpfsProtectionGroupParams) GetIndexingPolicy() IndexingPolicy {
	if o == nil || o.IndexingPolicy == nil {
		var ret IndexingPolicy
		return ret
	}
	return *o.IndexingPolicy
}

// GetIndexingPolicyOk returns a tuple with the IndexingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GpfsProtectionGroupParams) GetIndexingPolicyOk() (*IndexingPolicy, bool) {
	if o == nil || o.IndexingPolicy == nil {
		return nil, false
	}
	return o.IndexingPolicy, true
}

// HasIndexingPolicy returns a boolean if a field has been set.
func (o *GpfsProtectionGroupParams) HasIndexingPolicy() bool {
	if o != nil && o.IndexingPolicy != nil {
		return true
	}

	return false
}

// SetIndexingPolicy gets a reference to the given IndexingPolicy and assigns it to the IndexingPolicy field.
func (o *GpfsProtectionGroupParams) SetIndexingPolicy(v IndexingPolicy) {
	o.IndexingPolicy = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GpfsProtectionGroupParams) GetProtocol() string {
	if o == nil || o.Protocol.Get() == nil {
		var ret string
		return ret
	}
	return *o.Protocol.Get()
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GpfsProtectionGroupParams) GetProtocolOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Protocol.Get(), o.Protocol.IsSet()
}

// HasProtocol returns a boolean if a field has been set.
func (o *GpfsProtectionGroupParams) HasProtocol() bool {
	if o != nil && o.Protocol.IsSet() {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given NullableString and assigns it to the Protocol field.
func (o *GpfsProtectionGroupParams) SetProtocol(v string) {
	o.Protocol.Set(&v)
}
// SetProtocolNil sets the value for Protocol to be an explicit nil
func (o *GpfsProtectionGroupParams) SetProtocolNil() {
	o.Protocol.Set(nil)
}

// UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
func (o *GpfsProtectionGroupParams) UnsetProtocol() {
	o.Protocol.Unset()
}

// GetContinueOnError returns the ContinueOnError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GpfsProtectionGroupParams) GetContinueOnError() bool {
	if o == nil || o.ContinueOnError.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ContinueOnError.Get()
}

// GetContinueOnErrorOk returns a tuple with the ContinueOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GpfsProtectionGroupParams) GetContinueOnErrorOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContinueOnError.Get(), o.ContinueOnError.IsSet()
}

// HasContinueOnError returns a boolean if a field has been set.
func (o *GpfsProtectionGroupParams) HasContinueOnError() bool {
	if o != nil && o.ContinueOnError.IsSet() {
		return true
	}

	return false
}

// SetContinueOnError gets a reference to the given NullableBool and assigns it to the ContinueOnError field.
func (o *GpfsProtectionGroupParams) SetContinueOnError(v bool) {
	o.ContinueOnError.Set(&v)
}
// SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil
func (o *GpfsProtectionGroupParams) SetContinueOnErrorNil() {
	o.ContinueOnError.Set(nil)
}

// UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
func (o *GpfsProtectionGroupParams) UnsetContinueOnError() {
	o.ContinueOnError.Unset()
}

// GetEncryptionEnabled returns the EncryptionEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GpfsProtectionGroupParams) GetEncryptionEnabled() bool {
	if o == nil || o.EncryptionEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.EncryptionEnabled.Get()
}

// GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GpfsProtectionGroupParams) GetEncryptionEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EncryptionEnabled.Get(), o.EncryptionEnabled.IsSet()
}

// HasEncryptionEnabled returns a boolean if a field has been set.
func (o *GpfsProtectionGroupParams) HasEncryptionEnabled() bool {
	if o != nil && o.EncryptionEnabled.IsSet() {
		return true
	}

	return false
}

// SetEncryptionEnabled gets a reference to the given NullableBool and assigns it to the EncryptionEnabled field.
func (o *GpfsProtectionGroupParams) SetEncryptionEnabled(v bool) {
	o.EncryptionEnabled.Set(&v)
}
// SetEncryptionEnabledNil sets the value for EncryptionEnabled to be an explicit nil
func (o *GpfsProtectionGroupParams) SetEncryptionEnabledNil() {
	o.EncryptionEnabled.Set(nil)
}

// UnsetEncryptionEnabled ensures that no value is present for EncryptionEnabled, not even an explicit nil
func (o *GpfsProtectionGroupParams) UnsetEncryptionEnabled() {
	o.EncryptionEnabled.Unset()
}

// GetFileLockConfig returns the FileLockConfig field value if set, zero value otherwise.
func (o *GpfsProtectionGroupParams) GetFileLockConfig() FileLevelDataLockConfig {
	if o == nil || o.FileLockConfig == nil {
		var ret FileLevelDataLockConfig
		return ret
	}
	return *o.FileLockConfig
}

// GetFileLockConfigOk returns a tuple with the FileLockConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GpfsProtectionGroupParams) GetFileLockConfigOk() (*FileLevelDataLockConfig, bool) {
	if o == nil || o.FileLockConfig == nil {
		return nil, false
	}
	return o.FileLockConfig, true
}

// HasFileLockConfig returns a boolean if a field has been set.
func (o *GpfsProtectionGroupParams) HasFileLockConfig() bool {
	if o != nil && o.FileLockConfig != nil {
		return true
	}

	return false
}

// SetFileLockConfig gets a reference to the given FileLevelDataLockConfig and assigns it to the FileLockConfig field.
func (o *GpfsProtectionGroupParams) SetFileLockConfig(v FileLevelDataLockConfig) {
	o.FileLockConfig = &v
}

// GetFileFilters returns the FileFilters field value if set, zero value otherwise.
func (o *GpfsProtectionGroupParams) GetFileFilters() FileFilteringPolicy {
	if o == nil || o.FileFilters == nil {
		var ret FileFilteringPolicy
		return ret
	}
	return *o.FileFilters
}

// GetFileFiltersOk returns a tuple with the FileFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GpfsProtectionGroupParams) GetFileFiltersOk() (*FileFilteringPolicy, bool) {
	if o == nil || o.FileFilters == nil {
		return nil, false
	}
	return o.FileFilters, true
}

// HasFileFilters returns a boolean if a field has been set.
func (o *GpfsProtectionGroupParams) HasFileFilters() bool {
	if o != nil && o.FileFilters != nil {
		return true
	}

	return false
}

// SetFileFilters gets a reference to the given FileFilteringPolicy and assigns it to the FileFilters field.
func (o *GpfsProtectionGroupParams) SetFileFilters(v FileFilteringPolicy) {
	o.FileFilters = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GpfsProtectionGroupParams) GetSourceId() int64 {
	if o == nil || o.SourceId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SourceId.Get()
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GpfsProtectionGroupParams) GetSourceIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SourceId.Get(), o.SourceId.IsSet()
}

// HasSourceId returns a boolean if a field has been set.
func (o *GpfsProtectionGroupParams) HasSourceId() bool {
	if o != nil && o.SourceId.IsSet() {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given NullableInt64 and assigns it to the SourceId field.
func (o *GpfsProtectionGroupParams) SetSourceId(v int64) {
	o.SourceId.Set(&v)
}
// SetSourceIdNil sets the value for SourceId to be an explicit nil
func (o *GpfsProtectionGroupParams) SetSourceIdNil() {
	o.SourceId.Set(nil)
}

// UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
func (o *GpfsProtectionGroupParams) UnsetSourceId() {
	o.SourceId.Unset()
}

// GetSourceName returns the SourceName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GpfsProtectionGroupParams) GetSourceName() string {
	if o == nil || o.SourceName.Get() == nil {
		var ret string
		return ret
	}
	return *o.SourceName.Get()
}

// GetSourceNameOk returns a tuple with the SourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GpfsProtectionGroupParams) GetSourceNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SourceName.Get(), o.SourceName.IsSet()
}

// HasSourceName returns a boolean if a field has been set.
func (o *GpfsProtectionGroupParams) HasSourceName() bool {
	if o != nil && o.SourceName.IsSet() {
		return true
	}

	return false
}

// SetSourceName gets a reference to the given NullableString and assigns it to the SourceName field.
func (o *GpfsProtectionGroupParams) SetSourceName(v string) {
	o.SourceName.Set(&v)
}
// SetSourceNameNil sets the value for SourceName to be an explicit nil
func (o *GpfsProtectionGroupParams) SetSourceNameNil() {
	o.SourceName.Set(nil)
}

// UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
func (o *GpfsProtectionGroupParams) UnsetSourceName() {
	o.SourceName.Unset()
}

// GetPrePostScript returns the PrePostScript field value if set, zero value otherwise.
func (o *GpfsProtectionGroupParams) GetPrePostScript() HostBasedBackupScriptParams {
	if o == nil || o.PrePostScript == nil {
		var ret HostBasedBackupScriptParams
		return ret
	}
	return *o.PrePostScript
}

// GetPrePostScriptOk returns a tuple with the PrePostScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GpfsProtectionGroupParams) GetPrePostScriptOk() (*HostBasedBackupScriptParams, bool) {
	if o == nil || o.PrePostScript == nil {
		return nil, false
	}
	return o.PrePostScript, true
}

// HasPrePostScript returns a boolean if a field has been set.
func (o *GpfsProtectionGroupParams) HasPrePostScript() bool {
	if o != nil && o.PrePostScript != nil {
		return true
	}

	return false
}

// SetPrePostScript gets a reference to the given HostBasedBackupScriptParams and assigns it to the PrePostScript field.
func (o *GpfsProtectionGroupParams) SetPrePostScript(v HostBasedBackupScriptParams) {
	o.PrePostScript = &v
}

// GetFilterIpConfig returns the FilterIpConfig field value if set, zero value otherwise.
func (o *GpfsProtectionGroupParams) GetFilterIpConfig() FilterIpConfig {
	if o == nil || o.FilterIpConfig == nil {
		var ret FilterIpConfig
		return ret
	}
	return *o.FilterIpConfig
}

// GetFilterIpConfigOk returns a tuple with the FilterIpConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GpfsProtectionGroupParams) GetFilterIpConfigOk() (*FilterIpConfig, bool) {
	if o == nil || o.FilterIpConfig == nil {
		return nil, false
	}
	return o.FilterIpConfig, true
}

// HasFilterIpConfig returns a boolean if a field has been set.
func (o *GpfsProtectionGroupParams) HasFilterIpConfig() bool {
	if o != nil && o.FilterIpConfig != nil {
		return true
	}

	return false
}

// SetFilterIpConfig gets a reference to the given FilterIpConfig and assigns it to the FilterIpConfig field.
func (o *GpfsProtectionGroupParams) SetFilterIpConfig(v FilterIpConfig) {
	o.FilterIpConfig = &v
}

// GetThrottlingConfig returns the ThrottlingConfig field value if set, zero value otherwise.
func (o *GpfsProtectionGroupParams) GetThrottlingConfig() NasThrottlingConfig {
	if o == nil || o.ThrottlingConfig == nil {
		var ret NasThrottlingConfig
		return ret
	}
	return *o.ThrottlingConfig
}

// GetThrottlingConfigOk returns a tuple with the ThrottlingConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GpfsProtectionGroupParams) GetThrottlingConfigOk() (*NasThrottlingConfig, bool) {
	if o == nil || o.ThrottlingConfig == nil {
		return nil, false
	}
	return o.ThrottlingConfig, true
}

// HasThrottlingConfig returns a boolean if a field has been set.
func (o *GpfsProtectionGroupParams) HasThrottlingConfig() bool {
	if o != nil && o.ThrottlingConfig != nil {
		return true
	}

	return false
}

// SetThrottlingConfig gets a reference to the given NasThrottlingConfig and assigns it to the ThrottlingConfig field.
func (o *GpfsProtectionGroupParams) SetThrottlingConfig(v NasThrottlingConfig) {
	o.ThrottlingConfig = &v
}

func (o GpfsProtectionGroupParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["objects"] = o.Objects
	}
	if o.DirectCloudArchive.IsSet() {
		toSerialize["directCloudArchive"] = o.DirectCloudArchive.Get()
	}
	if o.NativeFormat.IsSet() {
		toSerialize["nativeFormat"] = o.NativeFormat.Get()
	}
	if o.IndexingPolicy != nil {
		toSerialize["indexingPolicy"] = o.IndexingPolicy
	}
	if o.Protocol.IsSet() {
		toSerialize["protocol"] = o.Protocol.Get()
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
	if o.SourceId.IsSet() {
		toSerialize["sourceId"] = o.SourceId.Get()
	}
	if o.SourceName.IsSet() {
		toSerialize["sourceName"] = o.SourceName.Get()
	}
	if o.PrePostScript != nil {
		toSerialize["prePostScript"] = o.PrePostScript
	}
	if o.FilterIpConfig != nil {
		toSerialize["filterIpConfig"] = o.FilterIpConfig
	}
	if o.ThrottlingConfig != nil {
		toSerialize["throttlingConfig"] = o.ThrottlingConfig
	}
	return json.Marshal(toSerialize)
}

type NullableGpfsProtectionGroupParams struct {
	value *GpfsProtectionGroupParams
	isSet bool
}

func (v NullableGpfsProtectionGroupParams) Get() *GpfsProtectionGroupParams {
	return v.value
}

func (v *NullableGpfsProtectionGroupParams) Set(val *GpfsProtectionGroupParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGpfsProtectionGroupParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGpfsProtectionGroupParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGpfsProtectionGroupParams(val *GpfsProtectionGroupParams) *NullableGpfsProtectionGroupParams {
	return &NullableGpfsProtectionGroupParams{value: val, isSet: true}
}

func (v NullableGpfsProtectionGroupParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGpfsProtectionGroupParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o GpfsProtectionGroupParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}