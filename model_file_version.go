/*
 * Cohesity REST API
 *
 * This API list provides operations for interfacing with the Cohesity Cluster.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cohesitysdk

import (
	"encoding/json"
)

// FileVersion Specifies information about a single file or folder.
type FileVersion struct {
	// Specifies the time when the file or folder was last modified. Specified as a Unix epoch Timestamp (in microseconds).
	ModifiedTimeUsecs NullableInt64 `json:"modifiedTimeUsecs,omitempty"`
	// Specifies the size of the file or folder (in bytes) from the most recent snapshot.
	SizeBytes NullableInt64 `json:"sizeBytes,omitempty"`
	// Array of Snapshots.  Specifies the available snapshots of the file or folder. When a Job Run executes, it captures a snapshot of object (such as a VM) that contains the file or folder.
	Snapshots []SnapshotAttempt `json:"snapshots,omitempty"`
}

// NewFileVersion instantiates a new FileVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileVersion() *FileVersion {
	this := FileVersion{}
	return &this
}

// NewFileVersionWithDefaults instantiates a new FileVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileVersionWithDefaults() *FileVersion {
	this := FileVersion{}
	return &this
}

// GetModifiedTimeUsecs returns the ModifiedTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileVersion) GetModifiedTimeUsecs() int64 {
	if o == nil || o.ModifiedTimeUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ModifiedTimeUsecs.Get()
}

// GetModifiedTimeUsecsOk returns a tuple with the ModifiedTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileVersion) GetModifiedTimeUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ModifiedTimeUsecs.Get(), o.ModifiedTimeUsecs.IsSet()
}

// HasModifiedTimeUsecs returns a boolean if a field has been set.
func (o *FileVersion) HasModifiedTimeUsecs() bool {
	if o != nil && o.ModifiedTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetModifiedTimeUsecs gets a reference to the given NullableInt64 and assigns it to the ModifiedTimeUsecs field.
func (o *FileVersion) SetModifiedTimeUsecs(v int64) {
	o.ModifiedTimeUsecs.Set(&v)
}
// SetModifiedTimeUsecsNil sets the value for ModifiedTimeUsecs to be an explicit nil
func (o *FileVersion) SetModifiedTimeUsecsNil() {
	o.ModifiedTimeUsecs.Set(nil)
}

// UnsetModifiedTimeUsecs ensures that no value is present for ModifiedTimeUsecs, not even an explicit nil
func (o *FileVersion) UnsetModifiedTimeUsecs() {
	o.ModifiedTimeUsecs.Unset()
}

// GetSizeBytes returns the SizeBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileVersion) GetSizeBytes() int64 {
	if o == nil || o.SizeBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SizeBytes.Get()
}

// GetSizeBytesOk returns a tuple with the SizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileVersion) GetSizeBytesOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SizeBytes.Get(), o.SizeBytes.IsSet()
}

// HasSizeBytes returns a boolean if a field has been set.
func (o *FileVersion) HasSizeBytes() bool {
	if o != nil && o.SizeBytes.IsSet() {
		return true
	}

	return false
}

// SetSizeBytes gets a reference to the given NullableInt64 and assigns it to the SizeBytes field.
func (o *FileVersion) SetSizeBytes(v int64) {
	o.SizeBytes.Set(&v)
}
// SetSizeBytesNil sets the value for SizeBytes to be an explicit nil
func (o *FileVersion) SetSizeBytesNil() {
	o.SizeBytes.Set(nil)
}

// UnsetSizeBytes ensures that no value is present for SizeBytes, not even an explicit nil
func (o *FileVersion) UnsetSizeBytes() {
	o.SizeBytes.Unset()
}

// GetSnapshots returns the Snapshots field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileVersion) GetSnapshots() []SnapshotAttempt {
	if o == nil  {
		var ret []SnapshotAttempt
		return ret
	}
	return o.Snapshots
}

// GetSnapshotsOk returns a tuple with the Snapshots field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileVersion) GetSnapshotsOk() (*[]SnapshotAttempt, bool) {
	if o == nil || o.Snapshots == nil {
		return nil, false
	}
	return &o.Snapshots, true
}

// HasSnapshots returns a boolean if a field has been set.
func (o *FileVersion) HasSnapshots() bool {
	if o != nil && o.Snapshots != nil {
		return true
	}

	return false
}

// SetSnapshots gets a reference to the given []SnapshotAttempt and assigns it to the Snapshots field.
func (o *FileVersion) SetSnapshots(v []SnapshotAttempt) {
	o.Snapshots = v
}

func (o FileVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ModifiedTimeUsecs.IsSet() {
		toSerialize["modifiedTimeUsecs"] = o.ModifiedTimeUsecs.Get()
	}
	if o.SizeBytes.IsSet() {
		toSerialize["sizeBytes"] = o.SizeBytes.Get()
	}
	if o.Snapshots != nil {
		toSerialize["snapshots"] = o.Snapshots
	}
	return json.Marshal(toSerialize)
}

type NullableFileVersion struct {
	value *FileVersion
	isSet bool
}

func (v NullableFileVersion) Get() *FileVersion {
	return v.value
}

func (v *NullableFileVersion) Set(val *FileVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableFileVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableFileVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileVersion(val *FileVersion) *NullableFileVersion {
	return &NullableFileVersion{value: val, isSet: true}
}

func (v NullableFileVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


