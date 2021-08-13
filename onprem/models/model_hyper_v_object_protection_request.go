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

// HyperVObjectProtectionRequest Specifies the HyperV object level settings for object protection.
type HyperVObjectProtectionRequest struct {
	// Specifies the id of the object being protected. This can be a leaf level or non leaf level object.
	Id NullableInt64 `json:"id"`
	// Specifies the list of IDs of the objects to not be protected by this Protection Group. This can be used to ignore specific objects under a parent object which has been included for protection.
	ExcludeObjectIds *[]int64 `json:"excludeObjectIds,omitempty"`
	// Specifies whether the backups should be copy-only.
	BackupsCopyOnly NullableBool `json:"backupsCopyOnly,omitempty"`
}

// NewHyperVObjectProtectionRequest instantiates a new HyperVObjectProtectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperVObjectProtectionRequest(id NullableInt64) *HyperVObjectProtectionRequest {
	this := HyperVObjectProtectionRequest{}
	this.Id = id
	return &this
}

// NewHyperVObjectProtectionRequestWithDefaults instantiates a new HyperVObjectProtectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperVObjectProtectionRequestWithDefaults() *HyperVObjectProtectionRequest {
	this := HyperVObjectProtectionRequest{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *HyperVObjectProtectionRequest) GetId() int64 {
	if o == nil || o.Id.Get() == nil {
		var ret int64
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperVObjectProtectionRequest) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *HyperVObjectProtectionRequest) SetId(v int64) {
	o.Id.Set(&v)
}

// GetExcludeObjectIds returns the ExcludeObjectIds field value if set, zero value otherwise.
func (o *HyperVObjectProtectionRequest) GetExcludeObjectIds() []int64 {
	if o == nil || o.ExcludeObjectIds == nil {
		var ret []int64
		return ret
	}
	return *o.ExcludeObjectIds
}

// GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperVObjectProtectionRequest) GetExcludeObjectIdsOk() (*[]int64, bool) {
	if o == nil || o.ExcludeObjectIds == nil {
		return nil, false
	}
	return o.ExcludeObjectIds, true
}

// HasExcludeObjectIds returns a boolean if a field has been set.
func (o *HyperVObjectProtectionRequest) HasExcludeObjectIds() bool {
	if o != nil && o.ExcludeObjectIds != nil {
		return true
	}

	return false
}

// SetExcludeObjectIds gets a reference to the given []int64 and assigns it to the ExcludeObjectIds field.
func (o *HyperVObjectProtectionRequest) SetExcludeObjectIds(v []int64) {
	o.ExcludeObjectIds = &v
}

// GetBackupsCopyOnly returns the BackupsCopyOnly field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperVObjectProtectionRequest) GetBackupsCopyOnly() bool {
	if o == nil || o.BackupsCopyOnly.Get() == nil {
		var ret bool
		return ret
	}
	return *o.BackupsCopyOnly.Get()
}

// GetBackupsCopyOnlyOk returns a tuple with the BackupsCopyOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperVObjectProtectionRequest) GetBackupsCopyOnlyOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BackupsCopyOnly.Get(), o.BackupsCopyOnly.IsSet()
}

// HasBackupsCopyOnly returns a boolean if a field has been set.
func (o *HyperVObjectProtectionRequest) HasBackupsCopyOnly() bool {
	if o != nil && o.BackupsCopyOnly.IsSet() {
		return true
	}

	return false
}

// SetBackupsCopyOnly gets a reference to the given NullableBool and assigns it to the BackupsCopyOnly field.
func (o *HyperVObjectProtectionRequest) SetBackupsCopyOnly(v bool) {
	o.BackupsCopyOnly.Set(&v)
}
// SetBackupsCopyOnlyNil sets the value for BackupsCopyOnly to be an explicit nil
func (o *HyperVObjectProtectionRequest) SetBackupsCopyOnlyNil() {
	o.BackupsCopyOnly.Set(nil)
}

// UnsetBackupsCopyOnly ensures that no value is present for BackupsCopyOnly, not even an explicit nil
func (o *HyperVObjectProtectionRequest) UnsetBackupsCopyOnly() {
	o.BackupsCopyOnly.Unset()
}

func (o HyperVObjectProtectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id.Get()
	}
	if o.ExcludeObjectIds != nil {
		toSerialize["excludeObjectIds"] = o.ExcludeObjectIds
	}
	if o.BackupsCopyOnly.IsSet() {
		toSerialize["backupsCopyOnly"] = o.BackupsCopyOnly.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableHyperVObjectProtectionRequest struct {
	value *HyperVObjectProtectionRequest
	isSet bool
}

func (v NullableHyperVObjectProtectionRequest) Get() *HyperVObjectProtectionRequest {
	return v.value
}

func (v *NullableHyperVObjectProtectionRequest) Set(val *HyperVObjectProtectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperVObjectProtectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperVObjectProtectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperVObjectProtectionRequest(val *HyperVObjectProtectionRequest) *NullableHyperVObjectProtectionRequest {
	return &NullableHyperVObjectProtectionRequest{value: val, isSet: true}
}

func (v NullableHyperVObjectProtectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperVObjectProtectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o HyperVObjectProtectionRequest) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}