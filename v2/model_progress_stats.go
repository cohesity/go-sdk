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

// checks if the ProgressStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProgressStats{}

// ProgressStats Specifies the stats within progress.
type ProgressStats struct {
	// Specifies the total number of file and directory entities that are backed up in this run. Only applicable to file based backups.
	BackupFileCount NullableInt64 `json:"backupFileCount,omitempty"`
	// Specifies whether the file system walk is done. Only applicable to file based backups.
	FileWalkDone NullableBool `json:"fileWalkDone,omitempty"`
	// Specifies the total number of file and directory entities visited in this backup. Only applicable to file based backups.
	TotalFileCount NullableInt64 `json:"totalFileCount,omitempty"`
}

// NewProgressStats instantiates a new ProgressStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProgressStats() *ProgressStats {
	this := ProgressStats{}
	return &this
}

// NewProgressStatsWithDefaults instantiates a new ProgressStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProgressStatsWithDefaults() *ProgressStats {
	this := ProgressStats{}
	return &this
}

// GetBackupFileCount returns the BackupFileCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProgressStats) GetBackupFileCount() int64 {
	if o == nil || IsNil(o.BackupFileCount.Get()) {
		var ret int64
		return ret
	}
	return *o.BackupFileCount.Get()
}

// GetBackupFileCountOk returns a tuple with the BackupFileCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProgressStats) GetBackupFileCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BackupFileCount.Get(), o.BackupFileCount.IsSet()
}

// HasBackupFileCount returns a boolean if a field has been set.
func (o *ProgressStats) HasBackupFileCount() bool {
	if o != nil && o.BackupFileCount.IsSet() {
		return true
	}

	return false
}

// SetBackupFileCount gets a reference to the given NullableInt64 and assigns it to the BackupFileCount field.
func (o *ProgressStats) SetBackupFileCount(v int64) {
	o.BackupFileCount.Set(&v)
}
// SetBackupFileCountNil sets the value for BackupFileCount to be an explicit nil
func (o *ProgressStats) SetBackupFileCountNil() {
	o.BackupFileCount.Set(nil)
}

// UnsetBackupFileCount ensures that no value is present for BackupFileCount, not even an explicit nil
func (o *ProgressStats) UnsetBackupFileCount() {
	o.BackupFileCount.Unset()
}

// GetFileWalkDone returns the FileWalkDone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProgressStats) GetFileWalkDone() bool {
	if o == nil || IsNil(o.FileWalkDone.Get()) {
		var ret bool
		return ret
	}
	return *o.FileWalkDone.Get()
}

// GetFileWalkDoneOk returns a tuple with the FileWalkDone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProgressStats) GetFileWalkDoneOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileWalkDone.Get(), o.FileWalkDone.IsSet()
}

// HasFileWalkDone returns a boolean if a field has been set.
func (o *ProgressStats) HasFileWalkDone() bool {
	if o != nil && o.FileWalkDone.IsSet() {
		return true
	}

	return false
}

// SetFileWalkDone gets a reference to the given NullableBool and assigns it to the FileWalkDone field.
func (o *ProgressStats) SetFileWalkDone(v bool) {
	o.FileWalkDone.Set(&v)
}
// SetFileWalkDoneNil sets the value for FileWalkDone to be an explicit nil
func (o *ProgressStats) SetFileWalkDoneNil() {
	o.FileWalkDone.Set(nil)
}

// UnsetFileWalkDone ensures that no value is present for FileWalkDone, not even an explicit nil
func (o *ProgressStats) UnsetFileWalkDone() {
	o.FileWalkDone.Unset()
}

// GetTotalFileCount returns the TotalFileCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProgressStats) GetTotalFileCount() int64 {
	if o == nil || IsNil(o.TotalFileCount.Get()) {
		var ret int64
		return ret
	}
	return *o.TotalFileCount.Get()
}

// GetTotalFileCountOk returns a tuple with the TotalFileCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProgressStats) GetTotalFileCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalFileCount.Get(), o.TotalFileCount.IsSet()
}

// HasTotalFileCount returns a boolean if a field has been set.
func (o *ProgressStats) HasTotalFileCount() bool {
	if o != nil && o.TotalFileCount.IsSet() {
		return true
	}

	return false
}

// SetTotalFileCount gets a reference to the given NullableInt64 and assigns it to the TotalFileCount field.
func (o *ProgressStats) SetTotalFileCount(v int64) {
	o.TotalFileCount.Set(&v)
}
// SetTotalFileCountNil sets the value for TotalFileCount to be an explicit nil
func (o *ProgressStats) SetTotalFileCountNil() {
	o.TotalFileCount.Set(nil)
}

// UnsetTotalFileCount ensures that no value is present for TotalFileCount, not even an explicit nil
func (o *ProgressStats) UnsetTotalFileCount() {
	o.TotalFileCount.Unset()
}

func (o ProgressStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProgressStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.BackupFileCount.IsSet() {
		toSerialize["backupFileCount"] = o.BackupFileCount.Get()
	}
	if o.FileWalkDone.IsSet() {
		toSerialize["fileWalkDone"] = o.FileWalkDone.Get()
	}
	if o.TotalFileCount.IsSet() {
		toSerialize["totalFileCount"] = o.TotalFileCount.Get()
	}
	return toSerialize, nil
}

type NullableProgressStats struct {
	value *ProgressStats
	isSet bool
}

func (v NullableProgressStats) Get() *ProgressStats {
	return v.value
}

func (v *NullableProgressStats) Set(val *ProgressStats) {
	v.value = val
	v.isSet = true
}

func (v NullableProgressStats) IsSet() bool {
	return v.isSet
}

func (v *NullableProgressStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProgressStats(val *ProgressStats) *NullableProgressStats {
	return &NullableProgressStats{value: val, isSet: true}
}

func (v NullableProgressStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProgressStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


