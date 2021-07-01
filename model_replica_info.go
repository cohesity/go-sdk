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

// ReplicaInfo Specifies the Replication information about a snapshot.
type ReplicaInfo struct {
	// Specifies the expiration time of the snapshot within the target location.
	ExpiryTimeUsecs NullableInt64 `json:"expiryTimeUsecs,omitempty"`
	SnapshotTargetSettings *SnapshotTargetSettings `json:"snapshotTargetSettings,omitempty"`
	Uid *UniversalId `json:"uid,omitempty"`
}

// NewReplicaInfo instantiates a new ReplicaInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplicaInfo() *ReplicaInfo {
	this := ReplicaInfo{}
	return &this
}

// NewReplicaInfoWithDefaults instantiates a new ReplicaInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplicaInfoWithDefaults() *ReplicaInfo {
	this := ReplicaInfo{}
	return &this
}

// GetExpiryTimeUsecs returns the ExpiryTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReplicaInfo) GetExpiryTimeUsecs() int64 {
	if o == nil || o.ExpiryTimeUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ExpiryTimeUsecs.Get()
}

// GetExpiryTimeUsecsOk returns a tuple with the ExpiryTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReplicaInfo) GetExpiryTimeUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExpiryTimeUsecs.Get(), o.ExpiryTimeUsecs.IsSet()
}

// HasExpiryTimeUsecs returns a boolean if a field has been set.
func (o *ReplicaInfo) HasExpiryTimeUsecs() bool {
	if o != nil && o.ExpiryTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetExpiryTimeUsecs gets a reference to the given NullableInt64 and assigns it to the ExpiryTimeUsecs field.
func (o *ReplicaInfo) SetExpiryTimeUsecs(v int64) {
	o.ExpiryTimeUsecs.Set(&v)
}
// SetExpiryTimeUsecsNil sets the value for ExpiryTimeUsecs to be an explicit nil
func (o *ReplicaInfo) SetExpiryTimeUsecsNil() {
	o.ExpiryTimeUsecs.Set(nil)
}

// UnsetExpiryTimeUsecs ensures that no value is present for ExpiryTimeUsecs, not even an explicit nil
func (o *ReplicaInfo) UnsetExpiryTimeUsecs() {
	o.ExpiryTimeUsecs.Unset()
}

// GetSnapshotTargetSettings returns the SnapshotTargetSettings field value if set, zero value otherwise.
func (o *ReplicaInfo) GetSnapshotTargetSettings() SnapshotTargetSettings {
	if o == nil || o.SnapshotTargetSettings == nil {
		var ret SnapshotTargetSettings
		return ret
	}
	return *o.SnapshotTargetSettings
}

// GetSnapshotTargetSettingsOk returns a tuple with the SnapshotTargetSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicaInfo) GetSnapshotTargetSettingsOk() (*SnapshotTargetSettings, bool) {
	if o == nil || o.SnapshotTargetSettings == nil {
		return nil, false
	}
	return o.SnapshotTargetSettings, true
}

// HasSnapshotTargetSettings returns a boolean if a field has been set.
func (o *ReplicaInfo) HasSnapshotTargetSettings() bool {
	if o != nil && o.SnapshotTargetSettings != nil {
		return true
	}

	return false
}

// SetSnapshotTargetSettings gets a reference to the given SnapshotTargetSettings and assigns it to the SnapshotTargetSettings field.
func (o *ReplicaInfo) SetSnapshotTargetSettings(v SnapshotTargetSettings) {
	o.SnapshotTargetSettings = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *ReplicaInfo) GetUid() UniversalId {
	if o == nil || o.Uid == nil {
		var ret UniversalId
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicaInfo) GetUidOk() (*UniversalId, bool) {
	if o == nil || o.Uid == nil {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *ReplicaInfo) HasUid() bool {
	if o != nil && o.Uid != nil {
		return true
	}

	return false
}

// SetUid gets a reference to the given UniversalId and assigns it to the Uid field.
func (o *ReplicaInfo) SetUid(v UniversalId) {
	o.Uid = &v
}

func (o ReplicaInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExpiryTimeUsecs.IsSet() {
		toSerialize["expiryTimeUsecs"] = o.ExpiryTimeUsecs.Get()
	}
	if o.SnapshotTargetSettings != nil {
		toSerialize["snapshotTargetSettings"] = o.SnapshotTargetSettings
	}
	if o.Uid != nil {
		toSerialize["uid"] = o.Uid
	}
	return json.Marshal(toSerialize)
}

type NullableReplicaInfo struct {
	value *ReplicaInfo
	isSet bool
}

func (v NullableReplicaInfo) Get() *ReplicaInfo {
	return v.value
}

func (v *NullableReplicaInfo) Set(val *ReplicaInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableReplicaInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableReplicaInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplicaInfo(val *ReplicaInfo) *NullableReplicaInfo {
	return &NullableReplicaInfo{value: val, isSet: true}
}

func (v NullableReplicaInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplicaInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


