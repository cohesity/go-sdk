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

// GetProtectionRunProgressBody Specifies the progress of a protection run.
type GetProtectionRunProgressBody struct {
	LocalRun *BackupRunProgressInfo `json:"localRun,omitempty"`
	// Progress for the archival run.
	ArchivalRun []ArchivalTargetProgressInfo `json:"archivalRun,omitempty"`
	// Progress for the replication run.
	ReplicationRun []ReplicationTargetProgressInfo `json:"replicationRun,omitempty"`
}

// NewGetProtectionRunProgressBody instantiates a new GetProtectionRunProgressBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProtectionRunProgressBody() *GetProtectionRunProgressBody {
	this := GetProtectionRunProgressBody{}
	return &this
}

// NewGetProtectionRunProgressBodyWithDefaults instantiates a new GetProtectionRunProgressBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProtectionRunProgressBodyWithDefaults() *GetProtectionRunProgressBody {
	this := GetProtectionRunProgressBody{}
	return &this
}

// GetLocalRun returns the LocalRun field value if set, zero value otherwise.
func (o *GetProtectionRunProgressBody) GetLocalRun() BackupRunProgressInfo {
	if o == nil || o.LocalRun == nil {
		var ret BackupRunProgressInfo
		return ret
	}
	return *o.LocalRun
}

// GetLocalRunOk returns a tuple with the LocalRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProtectionRunProgressBody) GetLocalRunOk() (*BackupRunProgressInfo, bool) {
	if o == nil || o.LocalRun == nil {
		return nil, false
	}
	return o.LocalRun, true
}

// HasLocalRun returns a boolean if a field has been set.
func (o *GetProtectionRunProgressBody) HasLocalRun() bool {
	if o != nil && o.LocalRun != nil {
		return true
	}

	return false
}

// SetLocalRun gets a reference to the given BackupRunProgressInfo and assigns it to the LocalRun field.
func (o *GetProtectionRunProgressBody) SetLocalRun(v BackupRunProgressInfo) {
	o.LocalRun = &v
}

// GetArchivalRun returns the ArchivalRun field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetProtectionRunProgressBody) GetArchivalRun() []ArchivalTargetProgressInfo {
	if o == nil  {
		var ret []ArchivalTargetProgressInfo
		return ret
	}
	return o.ArchivalRun
}

// GetArchivalRunOk returns a tuple with the ArchivalRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetProtectionRunProgressBody) GetArchivalRunOk() (*[]ArchivalTargetProgressInfo, bool) {
	if o == nil || o.ArchivalRun == nil {
		return nil, false
	}
	return &o.ArchivalRun, true
}

// HasArchivalRun returns a boolean if a field has been set.
func (o *GetProtectionRunProgressBody) HasArchivalRun() bool {
	if o != nil && o.ArchivalRun != nil {
		return true
	}

	return false
}

// SetArchivalRun gets a reference to the given []ArchivalTargetProgressInfo and assigns it to the ArchivalRun field.
func (o *GetProtectionRunProgressBody) SetArchivalRun(v []ArchivalTargetProgressInfo) {
	o.ArchivalRun = v
}

// GetReplicationRun returns the ReplicationRun field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetProtectionRunProgressBody) GetReplicationRun() []ReplicationTargetProgressInfo {
	if o == nil  {
		var ret []ReplicationTargetProgressInfo
		return ret
	}
	return o.ReplicationRun
}

// GetReplicationRunOk returns a tuple with the ReplicationRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetProtectionRunProgressBody) GetReplicationRunOk() (*[]ReplicationTargetProgressInfo, bool) {
	if o == nil || o.ReplicationRun == nil {
		return nil, false
	}
	return &o.ReplicationRun, true
}

// HasReplicationRun returns a boolean if a field has been set.
func (o *GetProtectionRunProgressBody) HasReplicationRun() bool {
	if o != nil && o.ReplicationRun != nil {
		return true
	}

	return false
}

// SetReplicationRun gets a reference to the given []ReplicationTargetProgressInfo and assigns it to the ReplicationRun field.
func (o *GetProtectionRunProgressBody) SetReplicationRun(v []ReplicationTargetProgressInfo) {
	o.ReplicationRun = v
}

func (o GetProtectionRunProgressBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LocalRun != nil {
		toSerialize["localRun"] = o.LocalRun
	}
	if o.ArchivalRun != nil {
		toSerialize["archivalRun"] = o.ArchivalRun
	}
	if o.ReplicationRun != nil {
		toSerialize["replicationRun"] = o.ReplicationRun
	}
	return json.Marshal(toSerialize)
}

type NullableGetProtectionRunProgressBody struct {
	value *GetProtectionRunProgressBody
	isSet bool
}

func (v NullableGetProtectionRunProgressBody) Get() *GetProtectionRunProgressBody {
	return v.value
}

func (v *NullableGetProtectionRunProgressBody) Set(val *GetProtectionRunProgressBody) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProtectionRunProgressBody) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProtectionRunProgressBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProtectionRunProgressBody(val *GetProtectionRunProgressBody) *NullableGetProtectionRunProgressBody {
	return &NullableGetProtectionRunProgressBody{value: val, isSet: true}
}

func (v NullableGetProtectionRunProgressBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProtectionRunProgressBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


