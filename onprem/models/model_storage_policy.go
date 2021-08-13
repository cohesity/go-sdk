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

// StoragePolicy Specifies the storage policy of a Storage Domain.
type StoragePolicy struct {
	// Specifies deduplication settings for a Storage Domain.
	DeduplicationParams *DeduplicationParams `json:"deduplicationParams,omitempty"`
	// Specifies compression settings for a Storage Domain.
	CompressionParams *CompressionParams `json:"compressionParams,omitempty"`
	// Specifies the time in seconds when deduplication and compression of the Storage Domain starts.
	DeduplicationCompressionDelaySecs NullableInt32 `json:"deduplicationCompressionDelaySecs,omitempty"`
	// Specifies the erasure coding parameters for a Storage Domain.
	ErasureCodingParams *ErasureCodingParams `json:"erasureCodingParams,omitempty"`
	// Specifies the encryption type for a Storage Domain.
	EncryptionType NullableString `json:"encryptionType,omitempty"`
	// Specifies the vault id assigned for cloud spill for a Storage Domain.
	CloudSpillVaultId NullableInt64 `json:"cloudSpillVaultId,omitempty"`
	// Specifies the number of disk failures to tolerate for a Storage Domain. By default, this field is 1 for cluster with three or more nodes. If erasure coding is enabled, this field will be the same as numCodedStripes.
	NumDiskFailuresTolerated NullableInt32 `json:"numDiskFailuresTolerated,omitempty"`
	// Specifies the number of node failures to tolerate for a Storage Domain. By default this field is replication factor minus 1 for replication chunk files and is the same as numCodedStripes for erasure coding chunk files.
	NumNodeFailuresTolerated NullableInt32 `json:"numNodeFailuresTolerated,omitempty"`
	// Specifies whether app marker detection is enabled. When enabled, app markers will be removed from data and put in separate chunks.
	AppMarkerDetectionEnabled NullableBool `json:"appMarkerDetectionEnabled,omitempty"`
}

// NewStoragePolicy instantiates a new StoragePolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoragePolicy() *StoragePolicy {
	this := StoragePolicy{}
	return &this
}

// NewStoragePolicyWithDefaults instantiates a new StoragePolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoragePolicyWithDefaults() *StoragePolicy {
	this := StoragePolicy{}
	return &this
}

// GetDeduplicationParams returns the DeduplicationParams field value if set, zero value otherwise.
func (o *StoragePolicy) GetDeduplicationParams() DeduplicationParams {
	if o == nil || o.DeduplicationParams == nil {
		var ret DeduplicationParams
		return ret
	}
	return *o.DeduplicationParams
}

// GetDeduplicationParamsOk returns a tuple with the DeduplicationParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePolicy) GetDeduplicationParamsOk() (*DeduplicationParams, bool) {
	if o == nil || o.DeduplicationParams == nil {
		return nil, false
	}
	return o.DeduplicationParams, true
}

// HasDeduplicationParams returns a boolean if a field has been set.
func (o *StoragePolicy) HasDeduplicationParams() bool {
	if o != nil && o.DeduplicationParams != nil {
		return true
	}

	return false
}

// SetDeduplicationParams gets a reference to the given DeduplicationParams and assigns it to the DeduplicationParams field.
func (o *StoragePolicy) SetDeduplicationParams(v DeduplicationParams) {
	o.DeduplicationParams = &v
}

// GetCompressionParams returns the CompressionParams field value if set, zero value otherwise.
func (o *StoragePolicy) GetCompressionParams() CompressionParams {
	if o == nil || o.CompressionParams == nil {
		var ret CompressionParams
		return ret
	}
	return *o.CompressionParams
}

// GetCompressionParamsOk returns a tuple with the CompressionParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePolicy) GetCompressionParamsOk() (*CompressionParams, bool) {
	if o == nil || o.CompressionParams == nil {
		return nil, false
	}
	return o.CompressionParams, true
}

// HasCompressionParams returns a boolean if a field has been set.
func (o *StoragePolicy) HasCompressionParams() bool {
	if o != nil && o.CompressionParams != nil {
		return true
	}

	return false
}

// SetCompressionParams gets a reference to the given CompressionParams and assigns it to the CompressionParams field.
func (o *StoragePolicy) SetCompressionParams(v CompressionParams) {
	o.CompressionParams = &v
}

// GetDeduplicationCompressionDelaySecs returns the DeduplicationCompressionDelaySecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePolicy) GetDeduplicationCompressionDelaySecs() int32 {
	if o == nil || o.DeduplicationCompressionDelaySecs.Get() == nil {
		var ret int32
		return ret
	}
	return *o.DeduplicationCompressionDelaySecs.Get()
}

// GetDeduplicationCompressionDelaySecsOk returns a tuple with the DeduplicationCompressionDelaySecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePolicy) GetDeduplicationCompressionDelaySecsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeduplicationCompressionDelaySecs.Get(), o.DeduplicationCompressionDelaySecs.IsSet()
}

// HasDeduplicationCompressionDelaySecs returns a boolean if a field has been set.
func (o *StoragePolicy) HasDeduplicationCompressionDelaySecs() bool {
	if o != nil && o.DeduplicationCompressionDelaySecs.IsSet() {
		return true
	}

	return false
}

// SetDeduplicationCompressionDelaySecs gets a reference to the given NullableInt32 and assigns it to the DeduplicationCompressionDelaySecs field.
func (o *StoragePolicy) SetDeduplicationCompressionDelaySecs(v int32) {
	o.DeduplicationCompressionDelaySecs.Set(&v)
}
// SetDeduplicationCompressionDelaySecsNil sets the value for DeduplicationCompressionDelaySecs to be an explicit nil
func (o *StoragePolicy) SetDeduplicationCompressionDelaySecsNil() {
	o.DeduplicationCompressionDelaySecs.Set(nil)
}

// UnsetDeduplicationCompressionDelaySecs ensures that no value is present for DeduplicationCompressionDelaySecs, not even an explicit nil
func (o *StoragePolicy) UnsetDeduplicationCompressionDelaySecs() {
	o.DeduplicationCompressionDelaySecs.Unset()
}

// GetErasureCodingParams returns the ErasureCodingParams field value if set, zero value otherwise.
func (o *StoragePolicy) GetErasureCodingParams() ErasureCodingParams {
	if o == nil || o.ErasureCodingParams == nil {
		var ret ErasureCodingParams
		return ret
	}
	return *o.ErasureCodingParams
}

// GetErasureCodingParamsOk returns a tuple with the ErasureCodingParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePolicy) GetErasureCodingParamsOk() (*ErasureCodingParams, bool) {
	if o == nil || o.ErasureCodingParams == nil {
		return nil, false
	}
	return o.ErasureCodingParams, true
}

// HasErasureCodingParams returns a boolean if a field has been set.
func (o *StoragePolicy) HasErasureCodingParams() bool {
	if o != nil && o.ErasureCodingParams != nil {
		return true
	}

	return false
}

// SetErasureCodingParams gets a reference to the given ErasureCodingParams and assigns it to the ErasureCodingParams field.
func (o *StoragePolicy) SetErasureCodingParams(v ErasureCodingParams) {
	o.ErasureCodingParams = &v
}

// GetEncryptionType returns the EncryptionType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePolicy) GetEncryptionType() string {
	if o == nil || o.EncryptionType.Get() == nil {
		var ret string
		return ret
	}
	return *o.EncryptionType.Get()
}

// GetEncryptionTypeOk returns a tuple with the EncryptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePolicy) GetEncryptionTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EncryptionType.Get(), o.EncryptionType.IsSet()
}

// HasEncryptionType returns a boolean if a field has been set.
func (o *StoragePolicy) HasEncryptionType() bool {
	if o != nil && o.EncryptionType.IsSet() {
		return true
	}

	return false
}

// SetEncryptionType gets a reference to the given NullableString and assigns it to the EncryptionType field.
func (o *StoragePolicy) SetEncryptionType(v string) {
	o.EncryptionType.Set(&v)
}
// SetEncryptionTypeNil sets the value for EncryptionType to be an explicit nil
func (o *StoragePolicy) SetEncryptionTypeNil() {
	o.EncryptionType.Set(nil)
}

// UnsetEncryptionType ensures that no value is present for EncryptionType, not even an explicit nil
func (o *StoragePolicy) UnsetEncryptionType() {
	o.EncryptionType.Unset()
}

// GetCloudSpillVaultId returns the CloudSpillVaultId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePolicy) GetCloudSpillVaultId() int64 {
	if o == nil || o.CloudSpillVaultId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CloudSpillVaultId.Get()
}

// GetCloudSpillVaultIdOk returns a tuple with the CloudSpillVaultId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePolicy) GetCloudSpillVaultIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CloudSpillVaultId.Get(), o.CloudSpillVaultId.IsSet()
}

// HasCloudSpillVaultId returns a boolean if a field has been set.
func (o *StoragePolicy) HasCloudSpillVaultId() bool {
	if o != nil && o.CloudSpillVaultId.IsSet() {
		return true
	}

	return false
}

// SetCloudSpillVaultId gets a reference to the given NullableInt64 and assigns it to the CloudSpillVaultId field.
func (o *StoragePolicy) SetCloudSpillVaultId(v int64) {
	o.CloudSpillVaultId.Set(&v)
}
// SetCloudSpillVaultIdNil sets the value for CloudSpillVaultId to be an explicit nil
func (o *StoragePolicy) SetCloudSpillVaultIdNil() {
	o.CloudSpillVaultId.Set(nil)
}

// UnsetCloudSpillVaultId ensures that no value is present for CloudSpillVaultId, not even an explicit nil
func (o *StoragePolicy) UnsetCloudSpillVaultId() {
	o.CloudSpillVaultId.Unset()
}

// GetNumDiskFailuresTolerated returns the NumDiskFailuresTolerated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePolicy) GetNumDiskFailuresTolerated() int32 {
	if o == nil || o.NumDiskFailuresTolerated.Get() == nil {
		var ret int32
		return ret
	}
	return *o.NumDiskFailuresTolerated.Get()
}

// GetNumDiskFailuresToleratedOk returns a tuple with the NumDiskFailuresTolerated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePolicy) GetNumDiskFailuresToleratedOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NumDiskFailuresTolerated.Get(), o.NumDiskFailuresTolerated.IsSet()
}

// HasNumDiskFailuresTolerated returns a boolean if a field has been set.
func (o *StoragePolicy) HasNumDiskFailuresTolerated() bool {
	if o != nil && o.NumDiskFailuresTolerated.IsSet() {
		return true
	}

	return false
}

// SetNumDiskFailuresTolerated gets a reference to the given NullableInt32 and assigns it to the NumDiskFailuresTolerated field.
func (o *StoragePolicy) SetNumDiskFailuresTolerated(v int32) {
	o.NumDiskFailuresTolerated.Set(&v)
}
// SetNumDiskFailuresToleratedNil sets the value for NumDiskFailuresTolerated to be an explicit nil
func (o *StoragePolicy) SetNumDiskFailuresToleratedNil() {
	o.NumDiskFailuresTolerated.Set(nil)
}

// UnsetNumDiskFailuresTolerated ensures that no value is present for NumDiskFailuresTolerated, not even an explicit nil
func (o *StoragePolicy) UnsetNumDiskFailuresTolerated() {
	o.NumDiskFailuresTolerated.Unset()
}

// GetNumNodeFailuresTolerated returns the NumNodeFailuresTolerated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePolicy) GetNumNodeFailuresTolerated() int32 {
	if o == nil || o.NumNodeFailuresTolerated.Get() == nil {
		var ret int32
		return ret
	}
	return *o.NumNodeFailuresTolerated.Get()
}

// GetNumNodeFailuresToleratedOk returns a tuple with the NumNodeFailuresTolerated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePolicy) GetNumNodeFailuresToleratedOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NumNodeFailuresTolerated.Get(), o.NumNodeFailuresTolerated.IsSet()
}

// HasNumNodeFailuresTolerated returns a boolean if a field has been set.
func (o *StoragePolicy) HasNumNodeFailuresTolerated() bool {
	if o != nil && o.NumNodeFailuresTolerated.IsSet() {
		return true
	}

	return false
}

// SetNumNodeFailuresTolerated gets a reference to the given NullableInt32 and assigns it to the NumNodeFailuresTolerated field.
func (o *StoragePolicy) SetNumNodeFailuresTolerated(v int32) {
	o.NumNodeFailuresTolerated.Set(&v)
}
// SetNumNodeFailuresToleratedNil sets the value for NumNodeFailuresTolerated to be an explicit nil
func (o *StoragePolicy) SetNumNodeFailuresToleratedNil() {
	o.NumNodeFailuresTolerated.Set(nil)
}

// UnsetNumNodeFailuresTolerated ensures that no value is present for NumNodeFailuresTolerated, not even an explicit nil
func (o *StoragePolicy) UnsetNumNodeFailuresTolerated() {
	o.NumNodeFailuresTolerated.Unset()
}

// GetAppMarkerDetectionEnabled returns the AppMarkerDetectionEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePolicy) GetAppMarkerDetectionEnabled() bool {
	if o == nil || o.AppMarkerDetectionEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AppMarkerDetectionEnabled.Get()
}

// GetAppMarkerDetectionEnabledOk returns a tuple with the AppMarkerDetectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePolicy) GetAppMarkerDetectionEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AppMarkerDetectionEnabled.Get(), o.AppMarkerDetectionEnabled.IsSet()
}

// HasAppMarkerDetectionEnabled returns a boolean if a field has been set.
func (o *StoragePolicy) HasAppMarkerDetectionEnabled() bool {
	if o != nil && o.AppMarkerDetectionEnabled.IsSet() {
		return true
	}

	return false
}

// SetAppMarkerDetectionEnabled gets a reference to the given NullableBool and assigns it to the AppMarkerDetectionEnabled field.
func (o *StoragePolicy) SetAppMarkerDetectionEnabled(v bool) {
	o.AppMarkerDetectionEnabled.Set(&v)
}
// SetAppMarkerDetectionEnabledNil sets the value for AppMarkerDetectionEnabled to be an explicit nil
func (o *StoragePolicy) SetAppMarkerDetectionEnabledNil() {
	o.AppMarkerDetectionEnabled.Set(nil)
}

// UnsetAppMarkerDetectionEnabled ensures that no value is present for AppMarkerDetectionEnabled, not even an explicit nil
func (o *StoragePolicy) UnsetAppMarkerDetectionEnabled() {
	o.AppMarkerDetectionEnabled.Unset()
}

func (o StoragePolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeduplicationParams != nil {
		toSerialize["deduplicationParams"] = o.DeduplicationParams
	}
	if o.CompressionParams != nil {
		toSerialize["compressionParams"] = o.CompressionParams
	}
	if o.DeduplicationCompressionDelaySecs.IsSet() {
		toSerialize["deduplicationCompressionDelaySecs"] = o.DeduplicationCompressionDelaySecs.Get()
	}
	if o.ErasureCodingParams != nil {
		toSerialize["erasureCodingParams"] = o.ErasureCodingParams
	}
	if o.EncryptionType.IsSet() {
		toSerialize["encryptionType"] = o.EncryptionType.Get()
	}
	if o.CloudSpillVaultId.IsSet() {
		toSerialize["cloudSpillVaultId"] = o.CloudSpillVaultId.Get()
	}
	if o.NumDiskFailuresTolerated.IsSet() {
		toSerialize["numDiskFailuresTolerated"] = o.NumDiskFailuresTolerated.Get()
	}
	if o.NumNodeFailuresTolerated.IsSet() {
		toSerialize["numNodeFailuresTolerated"] = o.NumNodeFailuresTolerated.Get()
	}
	if o.AppMarkerDetectionEnabled.IsSet() {
		toSerialize["appMarkerDetectionEnabled"] = o.AppMarkerDetectionEnabled.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableStoragePolicy struct {
	value *StoragePolicy
	isSet bool
}

func (v NullableStoragePolicy) Get() *StoragePolicy {
	return v.value
}

func (v *NullableStoragePolicy) Set(val *StoragePolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePolicy(val *StoragePolicy) *NullableStoragePolicy {
	return &NullableStoragePolicy{value: val, isSet: true}
}

func (v NullableStoragePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o StoragePolicy) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}