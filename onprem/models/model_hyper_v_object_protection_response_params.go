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

// HyperVObjectProtectionResponseParams Specifies the parameters which are specific to HyperV object protection.
type HyperVObjectProtectionResponseParams struct {
	// Specifies whether the backups should be copy-only.
	BackupsCopyOnly NullableBool `json:"backupsCopyOnly,omitempty"`
	// Specifies the list of IDs of the objects to not be protected by this Protection Group. This can be used to ignore specific objects under a parent object which has been included for protection.
	ExcludeObjectIds *[]int64 `json:"excludeObjectIds,omitempty"`
	// Specifies whether or not to quiesce apps and the file system in order to take app consistent snapshots. If not specified or false then snapshots will not be app consistent.
	AppConsistentSnapshot NullableBool `json:"appConsistentSnapshot,omitempty"`
	// Specifies whether or not to fallback to a crash consistent snapshot in the event that an app consistent snapshot fails.
	FallbackToCrashConsistentSnapshot NullableBool `json:"fallbackToCrashConsistentSnapshot,omitempty"`
	IndexingPolicy *IndexingPolicy `json:"indexingPolicy,omitempty"`
}

// NewHyperVObjectProtectionResponseParams instantiates a new HyperVObjectProtectionResponseParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperVObjectProtectionResponseParams() *HyperVObjectProtectionResponseParams {
	this := HyperVObjectProtectionResponseParams{}
	return &this
}

// NewHyperVObjectProtectionResponseParamsWithDefaults instantiates a new HyperVObjectProtectionResponseParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperVObjectProtectionResponseParamsWithDefaults() *HyperVObjectProtectionResponseParams {
	this := HyperVObjectProtectionResponseParams{}
	return &this
}

// GetBackupsCopyOnly returns the BackupsCopyOnly field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperVObjectProtectionResponseParams) GetBackupsCopyOnly() bool {
	if o == nil || o.BackupsCopyOnly.Get() == nil {
		var ret bool
		return ret
	}
	return *o.BackupsCopyOnly.Get()
}

// GetBackupsCopyOnlyOk returns a tuple with the BackupsCopyOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperVObjectProtectionResponseParams) GetBackupsCopyOnlyOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BackupsCopyOnly.Get(), o.BackupsCopyOnly.IsSet()
}

// HasBackupsCopyOnly returns a boolean if a field has been set.
func (o *HyperVObjectProtectionResponseParams) HasBackupsCopyOnly() bool {
	if o != nil && o.BackupsCopyOnly.IsSet() {
		return true
	}

	return false
}

// SetBackupsCopyOnly gets a reference to the given NullableBool and assigns it to the BackupsCopyOnly field.
func (o *HyperVObjectProtectionResponseParams) SetBackupsCopyOnly(v bool) {
	o.BackupsCopyOnly.Set(&v)
}
// SetBackupsCopyOnlyNil sets the value for BackupsCopyOnly to be an explicit nil
func (o *HyperVObjectProtectionResponseParams) SetBackupsCopyOnlyNil() {
	o.BackupsCopyOnly.Set(nil)
}

// UnsetBackupsCopyOnly ensures that no value is present for BackupsCopyOnly, not even an explicit nil
func (o *HyperVObjectProtectionResponseParams) UnsetBackupsCopyOnly() {
	o.BackupsCopyOnly.Unset()
}

// GetExcludeObjectIds returns the ExcludeObjectIds field value if set, zero value otherwise.
func (o *HyperVObjectProtectionResponseParams) GetExcludeObjectIds() []int64 {
	if o == nil || o.ExcludeObjectIds == nil {
		var ret []int64
		return ret
	}
	return *o.ExcludeObjectIds
}

// GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperVObjectProtectionResponseParams) GetExcludeObjectIdsOk() (*[]int64, bool) {
	if o == nil || o.ExcludeObjectIds == nil {
		return nil, false
	}
	return o.ExcludeObjectIds, true
}

// HasExcludeObjectIds returns a boolean if a field has been set.
func (o *HyperVObjectProtectionResponseParams) HasExcludeObjectIds() bool {
	if o != nil && o.ExcludeObjectIds != nil {
		return true
	}

	return false
}

// SetExcludeObjectIds gets a reference to the given []int64 and assigns it to the ExcludeObjectIds field.
func (o *HyperVObjectProtectionResponseParams) SetExcludeObjectIds(v []int64) {
	o.ExcludeObjectIds = &v
}

// GetAppConsistentSnapshot returns the AppConsistentSnapshot field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperVObjectProtectionResponseParams) GetAppConsistentSnapshot() bool {
	if o == nil || o.AppConsistentSnapshot.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AppConsistentSnapshot.Get()
}

// GetAppConsistentSnapshotOk returns a tuple with the AppConsistentSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperVObjectProtectionResponseParams) GetAppConsistentSnapshotOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AppConsistentSnapshot.Get(), o.AppConsistentSnapshot.IsSet()
}

// HasAppConsistentSnapshot returns a boolean if a field has been set.
func (o *HyperVObjectProtectionResponseParams) HasAppConsistentSnapshot() bool {
	if o != nil && o.AppConsistentSnapshot.IsSet() {
		return true
	}

	return false
}

// SetAppConsistentSnapshot gets a reference to the given NullableBool and assigns it to the AppConsistentSnapshot field.
func (o *HyperVObjectProtectionResponseParams) SetAppConsistentSnapshot(v bool) {
	o.AppConsistentSnapshot.Set(&v)
}
// SetAppConsistentSnapshotNil sets the value for AppConsistentSnapshot to be an explicit nil
func (o *HyperVObjectProtectionResponseParams) SetAppConsistentSnapshotNil() {
	o.AppConsistentSnapshot.Set(nil)
}

// UnsetAppConsistentSnapshot ensures that no value is present for AppConsistentSnapshot, not even an explicit nil
func (o *HyperVObjectProtectionResponseParams) UnsetAppConsistentSnapshot() {
	o.AppConsistentSnapshot.Unset()
}

// GetFallbackToCrashConsistentSnapshot returns the FallbackToCrashConsistentSnapshot field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperVObjectProtectionResponseParams) GetFallbackToCrashConsistentSnapshot() bool {
	if o == nil || o.FallbackToCrashConsistentSnapshot.Get() == nil {
		var ret bool
		return ret
	}
	return *o.FallbackToCrashConsistentSnapshot.Get()
}

// GetFallbackToCrashConsistentSnapshotOk returns a tuple with the FallbackToCrashConsistentSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperVObjectProtectionResponseParams) GetFallbackToCrashConsistentSnapshotOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FallbackToCrashConsistentSnapshot.Get(), o.FallbackToCrashConsistentSnapshot.IsSet()
}

// HasFallbackToCrashConsistentSnapshot returns a boolean if a field has been set.
func (o *HyperVObjectProtectionResponseParams) HasFallbackToCrashConsistentSnapshot() bool {
	if o != nil && o.FallbackToCrashConsistentSnapshot.IsSet() {
		return true
	}

	return false
}

// SetFallbackToCrashConsistentSnapshot gets a reference to the given NullableBool and assigns it to the FallbackToCrashConsistentSnapshot field.
func (o *HyperVObjectProtectionResponseParams) SetFallbackToCrashConsistentSnapshot(v bool) {
	o.FallbackToCrashConsistentSnapshot.Set(&v)
}
// SetFallbackToCrashConsistentSnapshotNil sets the value for FallbackToCrashConsistentSnapshot to be an explicit nil
func (o *HyperVObjectProtectionResponseParams) SetFallbackToCrashConsistentSnapshotNil() {
	o.FallbackToCrashConsistentSnapshot.Set(nil)
}

// UnsetFallbackToCrashConsistentSnapshot ensures that no value is present for FallbackToCrashConsistentSnapshot, not even an explicit nil
func (o *HyperVObjectProtectionResponseParams) UnsetFallbackToCrashConsistentSnapshot() {
	o.FallbackToCrashConsistentSnapshot.Unset()
}

// GetIndexingPolicy returns the IndexingPolicy field value if set, zero value otherwise.
func (o *HyperVObjectProtectionResponseParams) GetIndexingPolicy() IndexingPolicy {
	if o == nil || o.IndexingPolicy == nil {
		var ret IndexingPolicy
		return ret
	}
	return *o.IndexingPolicy
}

// GetIndexingPolicyOk returns a tuple with the IndexingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperVObjectProtectionResponseParams) GetIndexingPolicyOk() (*IndexingPolicy, bool) {
	if o == nil || o.IndexingPolicy == nil {
		return nil, false
	}
	return o.IndexingPolicy, true
}

// HasIndexingPolicy returns a boolean if a field has been set.
func (o *HyperVObjectProtectionResponseParams) HasIndexingPolicy() bool {
	if o != nil && o.IndexingPolicy != nil {
		return true
	}

	return false
}

// SetIndexingPolicy gets a reference to the given IndexingPolicy and assigns it to the IndexingPolicy field.
func (o *HyperVObjectProtectionResponseParams) SetIndexingPolicy(v IndexingPolicy) {
	o.IndexingPolicy = &v
}

func (o HyperVObjectProtectionResponseParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BackupsCopyOnly.IsSet() {
		toSerialize["backupsCopyOnly"] = o.BackupsCopyOnly.Get()
	}
	if o.ExcludeObjectIds != nil {
		toSerialize["excludeObjectIds"] = o.ExcludeObjectIds
	}
	if o.AppConsistentSnapshot.IsSet() {
		toSerialize["appConsistentSnapshot"] = o.AppConsistentSnapshot.Get()
	}
	if o.FallbackToCrashConsistentSnapshot.IsSet() {
		toSerialize["fallbackToCrashConsistentSnapshot"] = o.FallbackToCrashConsistentSnapshot.Get()
	}
	if o.IndexingPolicy != nil {
		toSerialize["indexingPolicy"] = o.IndexingPolicy
	}
	return json.Marshal(toSerialize)
}

type NullableHyperVObjectProtectionResponseParams struct {
	value *HyperVObjectProtectionResponseParams
	isSet bool
}

func (v NullableHyperVObjectProtectionResponseParams) Get() *HyperVObjectProtectionResponseParams {
	return v.value
}

func (v *NullableHyperVObjectProtectionResponseParams) Set(val *HyperVObjectProtectionResponseParams) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperVObjectProtectionResponseParams) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperVObjectProtectionResponseParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperVObjectProtectionResponseParams(val *HyperVObjectProtectionResponseParams) *NullableHyperVObjectProtectionResponseParams {
	return &NullableHyperVObjectProtectionResponseParams{value: val, isSet: true}
}

func (v NullableHyperVObjectProtectionResponseParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperVObjectProtectionResponseParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o HyperVObjectProtectionResponseParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}