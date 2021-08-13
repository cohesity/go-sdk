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

// UpdateReplicationSnapshotConfig Specifies the params to perform actions on replication snapshots taken by a Protection Group Run.
type UpdateReplicationSnapshotConfig struct {
	// Specifies the new configuration about adding Replication Snapshot to existing Protection Group Run.
	NewSnapshotConfig []RunReplicationConfig `json:"newSnapshotConfig,omitempty"`
	// Specifies the configuration about updating an existing Replication Snapshot Run.
	UpdateExistingSnapshotConfig []UpdateExistingReplicationSnapshotConfig `json:"updateExistingSnapshotConfig,omitempty"`
}

// NewUpdateReplicationSnapshotConfig instantiates a new UpdateReplicationSnapshotConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateReplicationSnapshotConfig() *UpdateReplicationSnapshotConfig {
	this := UpdateReplicationSnapshotConfig{}
	return &this
}

// NewUpdateReplicationSnapshotConfigWithDefaults instantiates a new UpdateReplicationSnapshotConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateReplicationSnapshotConfigWithDefaults() *UpdateReplicationSnapshotConfig {
	this := UpdateReplicationSnapshotConfig{}
	return &this
}

// GetNewSnapshotConfig returns the NewSnapshotConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateReplicationSnapshotConfig) GetNewSnapshotConfig() []RunReplicationConfig {
	if o == nil  {
		var ret []RunReplicationConfig
		return ret
	}
	return o.NewSnapshotConfig
}

// GetNewSnapshotConfigOk returns a tuple with the NewSnapshotConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateReplicationSnapshotConfig) GetNewSnapshotConfigOk() (*[]RunReplicationConfig, bool) {
	if o == nil || o.NewSnapshotConfig == nil {
		return nil, false
	}
	return &o.NewSnapshotConfig, true
}

// HasNewSnapshotConfig returns a boolean if a field has been set.
func (o *UpdateReplicationSnapshotConfig) HasNewSnapshotConfig() bool {
	if o != nil && o.NewSnapshotConfig != nil {
		return true
	}

	return false
}

// SetNewSnapshotConfig gets a reference to the given []RunReplicationConfig and assigns it to the NewSnapshotConfig field.
func (o *UpdateReplicationSnapshotConfig) SetNewSnapshotConfig(v []RunReplicationConfig) {
	o.NewSnapshotConfig = v
}

// GetUpdateExistingSnapshotConfig returns the UpdateExistingSnapshotConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateReplicationSnapshotConfig) GetUpdateExistingSnapshotConfig() []UpdateExistingReplicationSnapshotConfig {
	if o == nil  {
		var ret []UpdateExistingReplicationSnapshotConfig
		return ret
	}
	return o.UpdateExistingSnapshotConfig
}

// GetUpdateExistingSnapshotConfigOk returns a tuple with the UpdateExistingSnapshotConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateReplicationSnapshotConfig) GetUpdateExistingSnapshotConfigOk() (*[]UpdateExistingReplicationSnapshotConfig, bool) {
	if o == nil || o.UpdateExistingSnapshotConfig == nil {
		return nil, false
	}
	return &o.UpdateExistingSnapshotConfig, true
}

// HasUpdateExistingSnapshotConfig returns a boolean if a field has been set.
func (o *UpdateReplicationSnapshotConfig) HasUpdateExistingSnapshotConfig() bool {
	if o != nil && o.UpdateExistingSnapshotConfig != nil {
		return true
	}

	return false
}

// SetUpdateExistingSnapshotConfig gets a reference to the given []UpdateExistingReplicationSnapshotConfig and assigns it to the UpdateExistingSnapshotConfig field.
func (o *UpdateReplicationSnapshotConfig) SetUpdateExistingSnapshotConfig(v []UpdateExistingReplicationSnapshotConfig) {
	o.UpdateExistingSnapshotConfig = v
}

func (o UpdateReplicationSnapshotConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NewSnapshotConfig != nil {
		toSerialize["newSnapshotConfig"] = o.NewSnapshotConfig
	}
	if o.UpdateExistingSnapshotConfig != nil {
		toSerialize["updateExistingSnapshotConfig"] = o.UpdateExistingSnapshotConfig
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateReplicationSnapshotConfig struct {
	value *UpdateReplicationSnapshotConfig
	isSet bool
}

func (v NullableUpdateReplicationSnapshotConfig) Get() *UpdateReplicationSnapshotConfig {
	return v.value
}

func (v *NullableUpdateReplicationSnapshotConfig) Set(val *UpdateReplicationSnapshotConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateReplicationSnapshotConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateReplicationSnapshotConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateReplicationSnapshotConfig(val *UpdateReplicationSnapshotConfig) *NullableUpdateReplicationSnapshotConfig {
	return &NullableUpdateReplicationSnapshotConfig{value: val, isSet: true}
}

func (v NullableUpdateReplicationSnapshotConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateReplicationSnapshotConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


