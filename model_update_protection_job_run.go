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

// UpdateProtectionJobRun Specifies a Job Run and the expiration time to update. The expiration time defines the retention period for the Job Run and its snapshots.
type UpdateProtectionJobRun struct {
	// Specifies the retention for archival, replication or extended local retention.
	CopyRunTargets []RunJobSnapshotTarget `json:"copyRunTargets,omitempty"`
	// Specifies a unique universal id for the Job.
	JobUid NullableUniversalId `json:"jobUid,omitempty"`
	// Specifies the start time of the Job Run to update. The start time is specified as a Unix epoch Timestamp (in microseconds). This uniquely identifies a snapshot. This parameter is required.
	RunStartTimeUsecs NullableInt64 `json:"runStartTimeUsecs,omitempty"`
	// Ids of the Protection Sources. If this is specified, retention time will only be updated for the sources specified.
	SourceIds []int64 `json:"sourceIds,omitempty"`
}

// NewUpdateProtectionJobRun instantiates a new UpdateProtectionJobRun object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProtectionJobRun() *UpdateProtectionJobRun {
	this := UpdateProtectionJobRun{}
	return &this
}

// NewUpdateProtectionJobRunWithDefaults instantiates a new UpdateProtectionJobRun object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProtectionJobRunWithDefaults() *UpdateProtectionJobRun {
	this := UpdateProtectionJobRun{}
	return &this
}

// GetCopyRunTargets returns the CopyRunTargets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateProtectionJobRun) GetCopyRunTargets() []RunJobSnapshotTarget {
	if o == nil  {
		var ret []RunJobSnapshotTarget
		return ret
	}
	return o.CopyRunTargets
}

// GetCopyRunTargetsOk returns a tuple with the CopyRunTargets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateProtectionJobRun) GetCopyRunTargetsOk() (*[]RunJobSnapshotTarget, bool) {
	if o == nil || o.CopyRunTargets == nil {
		return nil, false
	}
	return &o.CopyRunTargets, true
}

// HasCopyRunTargets returns a boolean if a field has been set.
func (o *UpdateProtectionJobRun) HasCopyRunTargets() bool {
	if o != nil && o.CopyRunTargets != nil {
		return true
	}

	return false
}

// SetCopyRunTargets gets a reference to the given []RunJobSnapshotTarget and assigns it to the CopyRunTargets field.
func (o *UpdateProtectionJobRun) SetCopyRunTargets(v []RunJobSnapshotTarget) {
	o.CopyRunTargets = v
}

// GetJobUid returns the JobUid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateProtectionJobRun) GetJobUid() UniversalId {
	if o == nil || o.JobUid.Get() == nil {
		var ret UniversalId
		return ret
	}
	return *o.JobUid.Get()
}

// GetJobUidOk returns a tuple with the JobUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateProtectionJobRun) GetJobUidOk() (*UniversalId, bool) {
	if o == nil  {
		return nil, false
	}
	return o.JobUid.Get(), o.JobUid.IsSet()
}

// HasJobUid returns a boolean if a field has been set.
func (o *UpdateProtectionJobRun) HasJobUid() bool {
	if o != nil && o.JobUid.IsSet() {
		return true
	}

	return false
}

// SetJobUid gets a reference to the given NullableUniversalId and assigns it to the JobUid field.
func (o *UpdateProtectionJobRun) SetJobUid(v UniversalId) {
	o.JobUid.Set(&v)
}
// SetJobUidNil sets the value for JobUid to be an explicit nil
func (o *UpdateProtectionJobRun) SetJobUidNil() {
	o.JobUid.Set(nil)
}

// UnsetJobUid ensures that no value is present for JobUid, not even an explicit nil
func (o *UpdateProtectionJobRun) UnsetJobUid() {
	o.JobUid.Unset()
}

// GetRunStartTimeUsecs returns the RunStartTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateProtectionJobRun) GetRunStartTimeUsecs() int64 {
	if o == nil || o.RunStartTimeUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RunStartTimeUsecs.Get()
}

// GetRunStartTimeUsecsOk returns a tuple with the RunStartTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateProtectionJobRun) GetRunStartTimeUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RunStartTimeUsecs.Get(), o.RunStartTimeUsecs.IsSet()
}

// HasRunStartTimeUsecs returns a boolean if a field has been set.
func (o *UpdateProtectionJobRun) HasRunStartTimeUsecs() bool {
	if o != nil && o.RunStartTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetRunStartTimeUsecs gets a reference to the given NullableInt64 and assigns it to the RunStartTimeUsecs field.
func (o *UpdateProtectionJobRun) SetRunStartTimeUsecs(v int64) {
	o.RunStartTimeUsecs.Set(&v)
}
// SetRunStartTimeUsecsNil sets the value for RunStartTimeUsecs to be an explicit nil
func (o *UpdateProtectionJobRun) SetRunStartTimeUsecsNil() {
	o.RunStartTimeUsecs.Set(nil)
}

// UnsetRunStartTimeUsecs ensures that no value is present for RunStartTimeUsecs, not even an explicit nil
func (o *UpdateProtectionJobRun) UnsetRunStartTimeUsecs() {
	o.RunStartTimeUsecs.Unset()
}

// GetSourceIds returns the SourceIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateProtectionJobRun) GetSourceIds() []int64 {
	if o == nil  {
		var ret []int64
		return ret
	}
	return o.SourceIds
}

// GetSourceIdsOk returns a tuple with the SourceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateProtectionJobRun) GetSourceIdsOk() (*[]int64, bool) {
	if o == nil || o.SourceIds == nil {
		return nil, false
	}
	return &o.SourceIds, true
}

// HasSourceIds returns a boolean if a field has been set.
func (o *UpdateProtectionJobRun) HasSourceIds() bool {
	if o != nil && o.SourceIds != nil {
		return true
	}

	return false
}

// SetSourceIds gets a reference to the given []int64 and assigns it to the SourceIds field.
func (o *UpdateProtectionJobRun) SetSourceIds(v []int64) {
	o.SourceIds = v
}

func (o UpdateProtectionJobRun) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CopyRunTargets != nil {
		toSerialize["copyRunTargets"] = o.CopyRunTargets
	}
	if o.JobUid.IsSet() {
		toSerialize["jobUid"] = o.JobUid.Get()
	}
	if o.RunStartTimeUsecs.IsSet() {
		toSerialize["runStartTimeUsecs"] = o.RunStartTimeUsecs.Get()
	}
	if o.SourceIds != nil {
		toSerialize["sourceIds"] = o.SourceIds
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateProtectionJobRun struct {
	value *UpdateProtectionJobRun
	isSet bool
}

func (v NullableUpdateProtectionJobRun) Get() *UpdateProtectionJobRun {
	return v.value
}

func (v *NullableUpdateProtectionJobRun) Set(val *UpdateProtectionJobRun) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProtectionJobRun) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProtectionJobRun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProtectionJobRun(val *UpdateProtectionJobRun) *NullableUpdateProtectionJobRun {
	return &NullableUpdateProtectionJobRun{value: val, isSet: true}
}

func (v NullableUpdateProtectionJobRun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProtectionJobRun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


