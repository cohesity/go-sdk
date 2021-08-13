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

// McmObjectArchivalRunActivityParams Specifies the Object activity of an archival run.
type McmObjectArchivalRunActivityParams struct {
	// Specifies the ID of the Protection Run.
	RunId NullableString `json:"runId,omitempty"`
	// Specifies the type of the Protection Run.
	RunType NullableString `json:"runType,omitempty"`
	// Specifies the Protection Group Id.
	ProtectionGroupId NullableString `json:"protectionGroupId,omitempty"`
	// Specifies the Protection Group name.
	ProtectionGroupName NullableString `json:"protectionGroupName,omitempty"`
	// Specifies the id of the object snapshot that is created as a part of this Run. This field is only populated for runs which are successful.
	SnapshotId NullableString `json:"snapshotId,omitempty"`
	// Specifies the Protection Policy Id.
	PolicyId NullableString `json:"policyId,omitempty"`
	// Specifies the Protection Policy Name.
	PolicyName NullableString `json:"policyName,omitempty"`
	// Specifies the start time of Run in Unix epoch Timestamp(in microseconds).
	StartTimeUsecs NullableInt64 `json:"startTimeUsecs,omitempty"`
	// Specifies the end time of Run in Unix epoch Timestamp(in microseconds).
	EndTimeUsecs NullableInt64 `json:"endTimeUsecs,omitempty"`
	// Status of the Run. 'Running' indicates that the run is still running. 'Canceled' indicates that the run has been canceled. 'Canceling' indicates that the run is in the process of being canceled. 'Failed' indicates that the run has failed. 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening. 'Succeeded' indicates that the run has finished successfully. 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.
	Status NullableString `json:"status,omitempty"`
	// Progress monitor task id for the Run.
	ProgressTaskId NullableString `json:"progressTaskId,omitempty"`
	// Specifies the id of archival target.
	ArchivalTargetId NullableInt64 `json:"archivalTargetId,omitempty"`
	// Specifies the name of archival target.
	ArchivalTargetName NullableString `json:"archivalTargetName,omitempty"`
	// Specifies the type of protection environment.
	ProtectionEnvironmentType NullableString `json:"protectionEnvironmentType,omitempty"`
}

// NewMcmObjectArchivalRunActivityParams instantiates a new McmObjectArchivalRunActivityParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMcmObjectArchivalRunActivityParams() *McmObjectArchivalRunActivityParams {
	this := McmObjectArchivalRunActivityParams{}
	return &this
}

// NewMcmObjectArchivalRunActivityParamsWithDefaults instantiates a new McmObjectArchivalRunActivityParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMcmObjectArchivalRunActivityParamsWithDefaults() *McmObjectArchivalRunActivityParams {
	this := McmObjectArchivalRunActivityParams{}
	return &this
}

// GetRunId returns the RunId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *McmObjectArchivalRunActivityParams) GetRunId() string {
	if o == nil || o.RunId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RunId.Get()
}

// GetRunIdOk returns a tuple with the RunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *McmObjectArchivalRunActivityParams) GetRunIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RunId.Get(), o.RunId.IsSet()
}

// HasRunId returns a boolean if a field has been set.
func (o *McmObjectArchivalRunActivityParams) HasRunId() bool {
	if o != nil && o.RunId.IsSet() {
		return true
	}

	return false
}

// SetRunId gets a reference to the given NullableString and assigns it to the RunId field.
func (o *McmObjectArchivalRunActivityParams) SetRunId(v string) {
	o.RunId.Set(&v)
}
// SetRunIdNil sets the value for RunId to be an explicit nil
func (o *McmObjectArchivalRunActivityParams) SetRunIdNil() {
	o.RunId.Set(nil)
}

// UnsetRunId ensures that no value is present for RunId, not even an explicit nil
func (o *McmObjectArchivalRunActivityParams) UnsetRunId() {
	o.RunId.Unset()
}

// GetRunType returns the RunType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *McmObjectArchivalRunActivityParams) GetRunType() string {
	if o == nil || o.RunType.Get() == nil {
		var ret string
		return ret
	}
	return *o.RunType.Get()
}

// GetRunTypeOk returns a tuple with the RunType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *McmObjectArchivalRunActivityParams) GetRunTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RunType.Get(), o.RunType.IsSet()
}

// HasRunType returns a boolean if a field has been set.
func (o *McmObjectArchivalRunActivityParams) HasRunType() bool {
	if o != nil && o.RunType.IsSet() {
		return true
	}

	return false
}

// SetRunType gets a reference to the given NullableString and assigns it to the RunType field.
func (o *McmObjectArchivalRunActivityParams) SetRunType(v string) {
	o.RunType.Set(&v)
}
// SetRunTypeNil sets the value for RunType to be an explicit nil
func (o *McmObjectArchivalRunActivityParams) SetRunTypeNil() {
	o.RunType.Set(nil)
}

// UnsetRunType ensures that no value is present for RunType, not even an explicit nil
func (o *McmObjectArchivalRunActivityParams) UnsetRunType() {
	o.RunType.Unset()
}

// GetProtectionGroupId returns the ProtectionGroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *McmObjectArchivalRunActivityParams) GetProtectionGroupId() string {
	if o == nil || o.ProtectionGroupId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProtectionGroupId.Get()
}

// GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *McmObjectArchivalRunActivityParams) GetProtectionGroupIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProtectionGroupId.Get(), o.ProtectionGroupId.IsSet()
}

// HasProtectionGroupId returns a boolean if a field has been set.
func (o *McmObjectArchivalRunActivityParams) HasProtectionGroupId() bool {
	if o != nil && o.ProtectionGroupId.IsSet() {
		return true
	}

	return false
}

// SetProtectionGroupId gets a reference to the given NullableString and assigns it to the ProtectionGroupId field.
func (o *McmObjectArchivalRunActivityParams) SetProtectionGroupId(v string) {
	o.ProtectionGroupId.Set(&v)
}
// SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil
func (o *McmObjectArchivalRunActivityParams) SetProtectionGroupIdNil() {
	o.ProtectionGroupId.Set(nil)
}

// UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
func (o *McmObjectArchivalRunActivityParams) UnsetProtectionGroupId() {
	o.ProtectionGroupId.Unset()
}

// GetProtectionGroupName returns the ProtectionGroupName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *McmObjectArchivalRunActivityParams) GetProtectionGroupName() string {
	if o == nil || o.ProtectionGroupName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProtectionGroupName.Get()
}

// GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *McmObjectArchivalRunActivityParams) GetProtectionGroupNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProtectionGroupName.Get(), o.ProtectionGroupName.IsSet()
}

// HasProtectionGroupName returns a boolean if a field has been set.
func (o *McmObjectArchivalRunActivityParams) HasProtectionGroupName() bool {
	if o != nil && o.ProtectionGroupName.IsSet() {
		return true
	}

	return false
}

// SetProtectionGroupName gets a reference to the given NullableString and assigns it to the ProtectionGroupName field.
func (o *McmObjectArchivalRunActivityParams) SetProtectionGroupName(v string) {
	o.ProtectionGroupName.Set(&v)
}
// SetProtectionGroupNameNil sets the value for ProtectionGroupName to be an explicit nil
func (o *McmObjectArchivalRunActivityParams) SetProtectionGroupNameNil() {
	o.ProtectionGroupName.Set(nil)
}

// UnsetProtectionGroupName ensures that no value is present for ProtectionGroupName, not even an explicit nil
func (o *McmObjectArchivalRunActivityParams) UnsetProtectionGroupName() {
	o.ProtectionGroupName.Unset()
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *McmObjectArchivalRunActivityParams) GetSnapshotId() string {
	if o == nil || o.SnapshotId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SnapshotId.Get()
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *McmObjectArchivalRunActivityParams) GetSnapshotIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SnapshotId.Get(), o.SnapshotId.IsSet()
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *McmObjectArchivalRunActivityParams) HasSnapshotId() bool {
	if o != nil && o.SnapshotId.IsSet() {
		return true
	}

	return false
}

// SetSnapshotId gets a reference to the given NullableString and assigns it to the SnapshotId field.
func (o *McmObjectArchivalRunActivityParams) SetSnapshotId(v string) {
	o.SnapshotId.Set(&v)
}
// SetSnapshotIdNil sets the value for SnapshotId to be an explicit nil
func (o *McmObjectArchivalRunActivityParams) SetSnapshotIdNil() {
	o.SnapshotId.Set(nil)
}

// UnsetSnapshotId ensures that no value is present for SnapshotId, not even an explicit nil
func (o *McmObjectArchivalRunActivityParams) UnsetSnapshotId() {
	o.SnapshotId.Unset()
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *McmObjectArchivalRunActivityParams) GetPolicyId() string {
	if o == nil || o.PolicyId.Get() == nil {
		var ret string
		return ret
	}
	return *o.PolicyId.Get()
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *McmObjectArchivalRunActivityParams) GetPolicyIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PolicyId.Get(), o.PolicyId.IsSet()
}

// HasPolicyId returns a boolean if a field has been set.
func (o *McmObjectArchivalRunActivityParams) HasPolicyId() bool {
	if o != nil && o.PolicyId.IsSet() {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given NullableString and assigns it to the PolicyId field.
func (o *McmObjectArchivalRunActivityParams) SetPolicyId(v string) {
	o.PolicyId.Set(&v)
}
// SetPolicyIdNil sets the value for PolicyId to be an explicit nil
func (o *McmObjectArchivalRunActivityParams) SetPolicyIdNil() {
	o.PolicyId.Set(nil)
}

// UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
func (o *McmObjectArchivalRunActivityParams) UnsetPolicyId() {
	o.PolicyId.Unset()
}

// GetPolicyName returns the PolicyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *McmObjectArchivalRunActivityParams) GetPolicyName() string {
	if o == nil || o.PolicyName.Get() == nil {
		var ret string
		return ret
	}
	return *o.PolicyName.Get()
}

// GetPolicyNameOk returns a tuple with the PolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *McmObjectArchivalRunActivityParams) GetPolicyNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PolicyName.Get(), o.PolicyName.IsSet()
}

// HasPolicyName returns a boolean if a field has been set.
func (o *McmObjectArchivalRunActivityParams) HasPolicyName() bool {
	if o != nil && o.PolicyName.IsSet() {
		return true
	}

	return false
}

// SetPolicyName gets a reference to the given NullableString and assigns it to the PolicyName field.
func (o *McmObjectArchivalRunActivityParams) SetPolicyName(v string) {
	o.PolicyName.Set(&v)
}
// SetPolicyNameNil sets the value for PolicyName to be an explicit nil
func (o *McmObjectArchivalRunActivityParams) SetPolicyNameNil() {
	o.PolicyName.Set(nil)
}

// UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
func (o *McmObjectArchivalRunActivityParams) UnsetPolicyName() {
	o.PolicyName.Unset()
}

// GetStartTimeUsecs returns the StartTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *McmObjectArchivalRunActivityParams) GetStartTimeUsecs() int64 {
	if o == nil || o.StartTimeUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.StartTimeUsecs.Get()
}

// GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *McmObjectArchivalRunActivityParams) GetStartTimeUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartTimeUsecs.Get(), o.StartTimeUsecs.IsSet()
}

// HasStartTimeUsecs returns a boolean if a field has been set.
func (o *McmObjectArchivalRunActivityParams) HasStartTimeUsecs() bool {
	if o != nil && o.StartTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetStartTimeUsecs gets a reference to the given NullableInt64 and assigns it to the StartTimeUsecs field.
func (o *McmObjectArchivalRunActivityParams) SetStartTimeUsecs(v int64) {
	o.StartTimeUsecs.Set(&v)
}
// SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil
func (o *McmObjectArchivalRunActivityParams) SetStartTimeUsecsNil() {
	o.StartTimeUsecs.Set(nil)
}

// UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
func (o *McmObjectArchivalRunActivityParams) UnsetStartTimeUsecs() {
	o.StartTimeUsecs.Unset()
}

// GetEndTimeUsecs returns the EndTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *McmObjectArchivalRunActivityParams) GetEndTimeUsecs() int64 {
	if o == nil || o.EndTimeUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.EndTimeUsecs.Get()
}

// GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *McmObjectArchivalRunActivityParams) GetEndTimeUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndTimeUsecs.Get(), o.EndTimeUsecs.IsSet()
}

// HasEndTimeUsecs returns a boolean if a field has been set.
func (o *McmObjectArchivalRunActivityParams) HasEndTimeUsecs() bool {
	if o != nil && o.EndTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetEndTimeUsecs gets a reference to the given NullableInt64 and assigns it to the EndTimeUsecs field.
func (o *McmObjectArchivalRunActivityParams) SetEndTimeUsecs(v int64) {
	o.EndTimeUsecs.Set(&v)
}
// SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil
func (o *McmObjectArchivalRunActivityParams) SetEndTimeUsecsNil() {
	o.EndTimeUsecs.Set(nil)
}

// UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
func (o *McmObjectArchivalRunActivityParams) UnsetEndTimeUsecs() {
	o.EndTimeUsecs.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *McmObjectArchivalRunActivityParams) GetStatus() string {
	if o == nil || o.Status.Get() == nil {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *McmObjectArchivalRunActivityParams) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *McmObjectArchivalRunActivityParams) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *McmObjectArchivalRunActivityParams) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *McmObjectArchivalRunActivityParams) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *McmObjectArchivalRunActivityParams) UnsetStatus() {
	o.Status.Unset()
}

// GetProgressTaskId returns the ProgressTaskId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *McmObjectArchivalRunActivityParams) GetProgressTaskId() string {
	if o == nil || o.ProgressTaskId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProgressTaskId.Get()
}

// GetProgressTaskIdOk returns a tuple with the ProgressTaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *McmObjectArchivalRunActivityParams) GetProgressTaskIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProgressTaskId.Get(), o.ProgressTaskId.IsSet()
}

// HasProgressTaskId returns a boolean if a field has been set.
func (o *McmObjectArchivalRunActivityParams) HasProgressTaskId() bool {
	if o != nil && o.ProgressTaskId.IsSet() {
		return true
	}

	return false
}

// SetProgressTaskId gets a reference to the given NullableString and assigns it to the ProgressTaskId field.
func (o *McmObjectArchivalRunActivityParams) SetProgressTaskId(v string) {
	o.ProgressTaskId.Set(&v)
}
// SetProgressTaskIdNil sets the value for ProgressTaskId to be an explicit nil
func (o *McmObjectArchivalRunActivityParams) SetProgressTaskIdNil() {
	o.ProgressTaskId.Set(nil)
}

// UnsetProgressTaskId ensures that no value is present for ProgressTaskId, not even an explicit nil
func (o *McmObjectArchivalRunActivityParams) UnsetProgressTaskId() {
	o.ProgressTaskId.Unset()
}

// GetArchivalTargetId returns the ArchivalTargetId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *McmObjectArchivalRunActivityParams) GetArchivalTargetId() int64 {
	if o == nil || o.ArchivalTargetId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ArchivalTargetId.Get()
}

// GetArchivalTargetIdOk returns a tuple with the ArchivalTargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *McmObjectArchivalRunActivityParams) GetArchivalTargetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ArchivalTargetId.Get(), o.ArchivalTargetId.IsSet()
}

// HasArchivalTargetId returns a boolean if a field has been set.
func (o *McmObjectArchivalRunActivityParams) HasArchivalTargetId() bool {
	if o != nil && o.ArchivalTargetId.IsSet() {
		return true
	}

	return false
}

// SetArchivalTargetId gets a reference to the given NullableInt64 and assigns it to the ArchivalTargetId field.
func (o *McmObjectArchivalRunActivityParams) SetArchivalTargetId(v int64) {
	o.ArchivalTargetId.Set(&v)
}
// SetArchivalTargetIdNil sets the value for ArchivalTargetId to be an explicit nil
func (o *McmObjectArchivalRunActivityParams) SetArchivalTargetIdNil() {
	o.ArchivalTargetId.Set(nil)
}

// UnsetArchivalTargetId ensures that no value is present for ArchivalTargetId, not even an explicit nil
func (o *McmObjectArchivalRunActivityParams) UnsetArchivalTargetId() {
	o.ArchivalTargetId.Unset()
}

// GetArchivalTargetName returns the ArchivalTargetName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *McmObjectArchivalRunActivityParams) GetArchivalTargetName() string {
	if o == nil || o.ArchivalTargetName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ArchivalTargetName.Get()
}

// GetArchivalTargetNameOk returns a tuple with the ArchivalTargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *McmObjectArchivalRunActivityParams) GetArchivalTargetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ArchivalTargetName.Get(), o.ArchivalTargetName.IsSet()
}

// HasArchivalTargetName returns a boolean if a field has been set.
func (o *McmObjectArchivalRunActivityParams) HasArchivalTargetName() bool {
	if o != nil && o.ArchivalTargetName.IsSet() {
		return true
	}

	return false
}

// SetArchivalTargetName gets a reference to the given NullableString and assigns it to the ArchivalTargetName field.
func (o *McmObjectArchivalRunActivityParams) SetArchivalTargetName(v string) {
	o.ArchivalTargetName.Set(&v)
}
// SetArchivalTargetNameNil sets the value for ArchivalTargetName to be an explicit nil
func (o *McmObjectArchivalRunActivityParams) SetArchivalTargetNameNil() {
	o.ArchivalTargetName.Set(nil)
}

// UnsetArchivalTargetName ensures that no value is present for ArchivalTargetName, not even an explicit nil
func (o *McmObjectArchivalRunActivityParams) UnsetArchivalTargetName() {
	o.ArchivalTargetName.Unset()
}

// GetProtectionEnvironmentType returns the ProtectionEnvironmentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *McmObjectArchivalRunActivityParams) GetProtectionEnvironmentType() string {
	if o == nil || o.ProtectionEnvironmentType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProtectionEnvironmentType.Get()
}

// GetProtectionEnvironmentTypeOk returns a tuple with the ProtectionEnvironmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *McmObjectArchivalRunActivityParams) GetProtectionEnvironmentTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProtectionEnvironmentType.Get(), o.ProtectionEnvironmentType.IsSet()
}

// HasProtectionEnvironmentType returns a boolean if a field has been set.
func (o *McmObjectArchivalRunActivityParams) HasProtectionEnvironmentType() bool {
	if o != nil && o.ProtectionEnvironmentType.IsSet() {
		return true
	}

	return false
}

// SetProtectionEnvironmentType gets a reference to the given NullableString and assigns it to the ProtectionEnvironmentType field.
func (o *McmObjectArchivalRunActivityParams) SetProtectionEnvironmentType(v string) {
	o.ProtectionEnvironmentType.Set(&v)
}
// SetProtectionEnvironmentTypeNil sets the value for ProtectionEnvironmentType to be an explicit nil
func (o *McmObjectArchivalRunActivityParams) SetProtectionEnvironmentTypeNil() {
	o.ProtectionEnvironmentType.Set(nil)
}

// UnsetProtectionEnvironmentType ensures that no value is present for ProtectionEnvironmentType, not even an explicit nil
func (o *McmObjectArchivalRunActivityParams) UnsetProtectionEnvironmentType() {
	o.ProtectionEnvironmentType.Unset()
}

func (o McmObjectArchivalRunActivityParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RunId.IsSet() {
		toSerialize["runId"] = o.RunId.Get()
	}
	if o.RunType.IsSet() {
		toSerialize["runType"] = o.RunType.Get()
	}
	if o.ProtectionGroupId.IsSet() {
		toSerialize["protectionGroupId"] = o.ProtectionGroupId.Get()
	}
	if o.ProtectionGroupName.IsSet() {
		toSerialize["protectionGroupName"] = o.ProtectionGroupName.Get()
	}
	if o.SnapshotId.IsSet() {
		toSerialize["snapshotId"] = o.SnapshotId.Get()
	}
	if o.PolicyId.IsSet() {
		toSerialize["policyId"] = o.PolicyId.Get()
	}
	if o.PolicyName.IsSet() {
		toSerialize["policyName"] = o.PolicyName.Get()
	}
	if o.StartTimeUsecs.IsSet() {
		toSerialize["startTimeUsecs"] = o.StartTimeUsecs.Get()
	}
	if o.EndTimeUsecs.IsSet() {
		toSerialize["endTimeUsecs"] = o.EndTimeUsecs.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.ProgressTaskId.IsSet() {
		toSerialize["progressTaskId"] = o.ProgressTaskId.Get()
	}
	if o.ArchivalTargetId.IsSet() {
		toSerialize["archivalTargetId"] = o.ArchivalTargetId.Get()
	}
	if o.ArchivalTargetName.IsSet() {
		toSerialize["archivalTargetName"] = o.ArchivalTargetName.Get()
	}
	if o.ProtectionEnvironmentType.IsSet() {
		toSerialize["protectionEnvironmentType"] = o.ProtectionEnvironmentType.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMcmObjectArchivalRunActivityParams struct {
	value *McmObjectArchivalRunActivityParams
	isSet bool
}

func (v NullableMcmObjectArchivalRunActivityParams) Get() *McmObjectArchivalRunActivityParams {
	return v.value
}

func (v *NullableMcmObjectArchivalRunActivityParams) Set(val *McmObjectArchivalRunActivityParams) {
	v.value = val
	v.isSet = true
}

func (v NullableMcmObjectArchivalRunActivityParams) IsSet() bool {
	return v.isSet
}

func (v *NullableMcmObjectArchivalRunActivityParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMcmObjectArchivalRunActivityParams(val *McmObjectArchivalRunActivityParams) *NullableMcmObjectArchivalRunActivityParams {
	return &NullableMcmObjectArchivalRunActivityParams{value: val, isSet: true}
}

func (v NullableMcmObjectArchivalRunActivityParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMcmObjectArchivalRunActivityParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


