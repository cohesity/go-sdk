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

// checks if the HyperVObjectProtectionUpdateRequestParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperVObjectProtectionUpdateRequestParams{}

// HyperVObjectProtectionUpdateRequestParams Specifies the parameters which are specific to HyperV object protection.
type HyperVObjectProtectionUpdateRequestParams struct {
	// Specifies whether or not to quiesce apps and the file system in order to take app consistent snapshots. If not specified or false then snapshots will not be app consistent.
	AppConsistentSnapshot NullableBool `json:"appConsistentSnapshot,omitempty"`
	// Specifies whether or not to fallback to a crash consistent snapshot in the event that an app consistent snapshot fails.
	FallbackToCrashConsistentSnapshot NullableBool `json:"fallbackToCrashConsistentSnapshot,omitempty"`
	IndexingPolicy *IndexingPolicy `json:"indexingPolicy,omitempty"`
	// Specifies whether the backups should be copy-only.
	BackupsCopyOnly NullableBool `json:"backupsCopyOnly,omitempty"`
	// Specifies the list of IDs of the objects to not be protected by this Protection Group. This can be used to ignore specific objects under a parent object which has been included for protection.
	ExcludeObjectIds []*int64 `json:"excludeObjectIds,omitempty"`
}

// NewHyperVObjectProtectionUpdateRequestParams instantiates a new HyperVObjectProtectionUpdateRequestParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperVObjectProtectionUpdateRequestParams() *HyperVObjectProtectionUpdateRequestParams {
	this := HyperVObjectProtectionUpdateRequestParams{}
	return &this
}

// NewHyperVObjectProtectionUpdateRequestParamsWithDefaults instantiates a new HyperVObjectProtectionUpdateRequestParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperVObjectProtectionUpdateRequestParamsWithDefaults() *HyperVObjectProtectionUpdateRequestParams {
	this := HyperVObjectProtectionUpdateRequestParams{}
	return &this
}

// GetAppConsistentSnapshot returns the AppConsistentSnapshot field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperVObjectProtectionUpdateRequestParams) GetAppConsistentSnapshot() bool {
	if o == nil || IsNil(o.AppConsistentSnapshot.Get()) {
		var ret bool
		return ret
	}
	return *o.AppConsistentSnapshot.Get()
}

// GetAppConsistentSnapshotOk returns a tuple with the AppConsistentSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperVObjectProtectionUpdateRequestParams) GetAppConsistentSnapshotOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppConsistentSnapshot.Get(), o.AppConsistentSnapshot.IsSet()
}

// HasAppConsistentSnapshot returns a boolean if a field has been set.
func (o *HyperVObjectProtectionUpdateRequestParams) HasAppConsistentSnapshot() bool {
	if o != nil && o.AppConsistentSnapshot.IsSet() {
		return true
	}

	return false
}

// SetAppConsistentSnapshot gets a reference to the given NullableBool and assigns it to the AppConsistentSnapshot field.
func (o *HyperVObjectProtectionUpdateRequestParams) SetAppConsistentSnapshot(v bool) {
	o.AppConsistentSnapshot.Set(&v)
}
// SetAppConsistentSnapshotNil sets the value for AppConsistentSnapshot to be an explicit nil
func (o *HyperVObjectProtectionUpdateRequestParams) SetAppConsistentSnapshotNil() {
	o.AppConsistentSnapshot.Set(nil)
}

// UnsetAppConsistentSnapshot ensures that no value is present for AppConsistentSnapshot, not even an explicit nil
func (o *HyperVObjectProtectionUpdateRequestParams) UnsetAppConsistentSnapshot() {
	o.AppConsistentSnapshot.Unset()
}

// GetFallbackToCrashConsistentSnapshot returns the FallbackToCrashConsistentSnapshot field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperVObjectProtectionUpdateRequestParams) GetFallbackToCrashConsistentSnapshot() bool {
	if o == nil || IsNil(o.FallbackToCrashConsistentSnapshot.Get()) {
		var ret bool
		return ret
	}
	return *o.FallbackToCrashConsistentSnapshot.Get()
}

// GetFallbackToCrashConsistentSnapshotOk returns a tuple with the FallbackToCrashConsistentSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperVObjectProtectionUpdateRequestParams) GetFallbackToCrashConsistentSnapshotOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.FallbackToCrashConsistentSnapshot.Get(), o.FallbackToCrashConsistentSnapshot.IsSet()
}

// HasFallbackToCrashConsistentSnapshot returns a boolean if a field has been set.
func (o *HyperVObjectProtectionUpdateRequestParams) HasFallbackToCrashConsistentSnapshot() bool {
	if o != nil && o.FallbackToCrashConsistentSnapshot.IsSet() {
		return true
	}

	return false
}

// SetFallbackToCrashConsistentSnapshot gets a reference to the given NullableBool and assigns it to the FallbackToCrashConsistentSnapshot field.
func (o *HyperVObjectProtectionUpdateRequestParams) SetFallbackToCrashConsistentSnapshot(v bool) {
	o.FallbackToCrashConsistentSnapshot.Set(&v)
}
// SetFallbackToCrashConsistentSnapshotNil sets the value for FallbackToCrashConsistentSnapshot to be an explicit nil
func (o *HyperVObjectProtectionUpdateRequestParams) SetFallbackToCrashConsistentSnapshotNil() {
	o.FallbackToCrashConsistentSnapshot.Set(nil)
}

// UnsetFallbackToCrashConsistentSnapshot ensures that no value is present for FallbackToCrashConsistentSnapshot, not even an explicit nil
func (o *HyperVObjectProtectionUpdateRequestParams) UnsetFallbackToCrashConsistentSnapshot() {
	o.FallbackToCrashConsistentSnapshot.Unset()
}

// GetIndexingPolicy returns the IndexingPolicy field value if set, zero value otherwise.
func (o *HyperVObjectProtectionUpdateRequestParams) GetIndexingPolicy() IndexingPolicy {
	if o == nil || IsNil(o.IndexingPolicy) {
		var ret IndexingPolicy
		return ret
	}
	return *o.IndexingPolicy
}

// GetIndexingPolicyOk returns a tuple with the IndexingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperVObjectProtectionUpdateRequestParams) GetIndexingPolicyOk() (*IndexingPolicy, bool) {
	if o == nil || IsNil(o.IndexingPolicy) {
		return nil, false
	}
	return o.IndexingPolicy, true
}

// HasIndexingPolicy returns a boolean if a field has been set.
func (o *HyperVObjectProtectionUpdateRequestParams) HasIndexingPolicy() bool {
	if o != nil && !IsNil(o.IndexingPolicy) {
		return true
	}

	return false
}

// SetIndexingPolicy gets a reference to the given IndexingPolicy and assigns it to the IndexingPolicy field.
func (o *HyperVObjectProtectionUpdateRequestParams) SetIndexingPolicy(v IndexingPolicy) {
	o.IndexingPolicy = &v
}

// GetBackupsCopyOnly returns the BackupsCopyOnly field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperVObjectProtectionUpdateRequestParams) GetBackupsCopyOnly() bool {
	if o == nil || IsNil(o.BackupsCopyOnly.Get()) {
		var ret bool
		return ret
	}
	return *o.BackupsCopyOnly.Get()
}

// GetBackupsCopyOnlyOk returns a tuple with the BackupsCopyOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperVObjectProtectionUpdateRequestParams) GetBackupsCopyOnlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.BackupsCopyOnly.Get(), o.BackupsCopyOnly.IsSet()
}

// HasBackupsCopyOnly returns a boolean if a field has been set.
func (o *HyperVObjectProtectionUpdateRequestParams) HasBackupsCopyOnly() bool {
	if o != nil && o.BackupsCopyOnly.IsSet() {
		return true
	}

	return false
}

// SetBackupsCopyOnly gets a reference to the given NullableBool and assigns it to the BackupsCopyOnly field.
func (o *HyperVObjectProtectionUpdateRequestParams) SetBackupsCopyOnly(v bool) {
	o.BackupsCopyOnly.Set(&v)
}
// SetBackupsCopyOnlyNil sets the value for BackupsCopyOnly to be an explicit nil
func (o *HyperVObjectProtectionUpdateRequestParams) SetBackupsCopyOnlyNil() {
	o.BackupsCopyOnly.Set(nil)
}

// UnsetBackupsCopyOnly ensures that no value is present for BackupsCopyOnly, not even an explicit nil
func (o *HyperVObjectProtectionUpdateRequestParams) UnsetBackupsCopyOnly() {
	o.BackupsCopyOnly.Unset()
}

// GetExcludeObjectIds returns the ExcludeObjectIds field value if set, zero value otherwise.
func (o *HyperVObjectProtectionUpdateRequestParams) GetExcludeObjectIds() []*int64 {
	if o == nil || IsNil(o.ExcludeObjectIds) {
		var ret []*int64
		return ret
	}
	return o.ExcludeObjectIds
}

// GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperVObjectProtectionUpdateRequestParams) GetExcludeObjectIdsOk() ([]*int64, bool) {
	if o == nil || IsNil(o.ExcludeObjectIds) {
		return nil, false
	}
	return o.ExcludeObjectIds, true
}

// HasExcludeObjectIds returns a boolean if a field has been set.
func (o *HyperVObjectProtectionUpdateRequestParams) HasExcludeObjectIds() bool {
	if o != nil && !IsNil(o.ExcludeObjectIds) {
		return true
	}

	return false
}

// SetExcludeObjectIds gets a reference to the given []*int64 and assigns it to the ExcludeObjectIds field.
func (o *HyperVObjectProtectionUpdateRequestParams) SetExcludeObjectIds(v []*int64) {
	o.ExcludeObjectIds = v
}

func (o HyperVObjectProtectionUpdateRequestParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperVObjectProtectionUpdateRequestParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AppConsistentSnapshot.IsSet() {
		toSerialize["appConsistentSnapshot"] = o.AppConsistentSnapshot.Get()
	}
	if o.FallbackToCrashConsistentSnapshot.IsSet() {
		toSerialize["fallbackToCrashConsistentSnapshot"] = o.FallbackToCrashConsistentSnapshot.Get()
	}
	if !IsNil(o.IndexingPolicy) {
		toSerialize["indexingPolicy"] = o.IndexingPolicy
	}
	if o.BackupsCopyOnly.IsSet() {
		toSerialize["backupsCopyOnly"] = o.BackupsCopyOnly.Get()
	}
	if !IsNil(o.ExcludeObjectIds) {
		toSerialize["excludeObjectIds"] = o.ExcludeObjectIds
	}
	return toSerialize, nil
}

type NullableHyperVObjectProtectionUpdateRequestParams struct {
	value *HyperVObjectProtectionUpdateRequestParams
	isSet bool
}

func (v NullableHyperVObjectProtectionUpdateRequestParams) Get() *HyperVObjectProtectionUpdateRequestParams {
	return v.value
}

func (v *NullableHyperVObjectProtectionUpdateRequestParams) Set(val *HyperVObjectProtectionUpdateRequestParams) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperVObjectProtectionUpdateRequestParams) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperVObjectProtectionUpdateRequestParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperVObjectProtectionUpdateRequestParams(val *HyperVObjectProtectionUpdateRequestParams) *NullableHyperVObjectProtectionUpdateRequestParams {
	return &NullableHyperVObjectProtectionUpdateRequestParams{value: val, isSet: true}
}

func (v NullableHyperVObjectProtectionUpdateRequestParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperVObjectProtectionUpdateRequestParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


