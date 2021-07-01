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

// RestoreInfo Specifies the info regarding a full SQL snapshot.
type RestoreInfo struct {
	ArchivalTarget *ArchivalExternalTarget `json:"archivalTarget,omitempty"`
	// Specifies the attempt number.
	AttemptNumber NullableInt32 `json:"attemptNumber,omitempty"`
	CloudDeployTarget *CloudDeployTargetDetails `json:"cloudDeployTarget,omitempty"`
	// Specifies the id of the job run.
	JobRunId NullableInt64 `json:"jobRunId,omitempty"`
	JobUid *UniversalId `json:"jobUid,omitempty"`
	ParentSource *ProtectionSource `json:"parentSource,omitempty"`
	// This field specifies the time in to which the object needs to be restored. This filed is only applicable when object is being backeup using CDP feature.
	RestoreTimeUsecs NullableInt64 `json:"restoreTimeUsecs,omitempty"`
	// Specifies the relative path of the snapshot directory.
	SnapshotRelativeDirPath NullableString `json:"snapshotRelativeDirPath,omitempty"`
	Source *ProtectionSource `json:"source,omitempty"`
	// Specifies the start time specified as a Unix epoch Timestamp (in microseconds).
	StartTimeUsecs NullableInt64 `json:"startTimeUsecs,omitempty"`
	// Specifies the name of the view.
	ViewName NullableString `json:"viewName,omitempty"`
	// Specifies if the VM had independent disks.
	VmHadIndependentDisks NullableBool `json:"vmHadIndependentDisks,omitempty"`
}

// NewRestoreInfo instantiates a new RestoreInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestoreInfo() *RestoreInfo {
	this := RestoreInfo{}
	return &this
}

// NewRestoreInfoWithDefaults instantiates a new RestoreInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestoreInfoWithDefaults() *RestoreInfo {
	this := RestoreInfo{}
	return &this
}

// GetArchivalTarget returns the ArchivalTarget field value if set, zero value otherwise.
func (o *RestoreInfo) GetArchivalTarget() ArchivalExternalTarget {
	if o == nil || o.ArchivalTarget == nil {
		var ret ArchivalExternalTarget
		return ret
	}
	return *o.ArchivalTarget
}

// GetArchivalTargetOk returns a tuple with the ArchivalTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreInfo) GetArchivalTargetOk() (*ArchivalExternalTarget, bool) {
	if o == nil || o.ArchivalTarget == nil {
		return nil, false
	}
	return o.ArchivalTarget, true
}

// HasArchivalTarget returns a boolean if a field has been set.
func (o *RestoreInfo) HasArchivalTarget() bool {
	if o != nil && o.ArchivalTarget != nil {
		return true
	}

	return false
}

// SetArchivalTarget gets a reference to the given ArchivalExternalTarget and assigns it to the ArchivalTarget field.
func (o *RestoreInfo) SetArchivalTarget(v ArchivalExternalTarget) {
	o.ArchivalTarget = &v
}

// GetAttemptNumber returns the AttemptNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestoreInfo) GetAttemptNumber() int32 {
	if o == nil || o.AttemptNumber.Get() == nil {
		var ret int32
		return ret
	}
	return *o.AttemptNumber.Get()
}

// GetAttemptNumberOk returns a tuple with the AttemptNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestoreInfo) GetAttemptNumberOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AttemptNumber.Get(), o.AttemptNumber.IsSet()
}

// HasAttemptNumber returns a boolean if a field has been set.
func (o *RestoreInfo) HasAttemptNumber() bool {
	if o != nil && o.AttemptNumber.IsSet() {
		return true
	}

	return false
}

// SetAttemptNumber gets a reference to the given NullableInt32 and assigns it to the AttemptNumber field.
func (o *RestoreInfo) SetAttemptNumber(v int32) {
	o.AttemptNumber.Set(&v)
}
// SetAttemptNumberNil sets the value for AttemptNumber to be an explicit nil
func (o *RestoreInfo) SetAttemptNumberNil() {
	o.AttemptNumber.Set(nil)
}

// UnsetAttemptNumber ensures that no value is present for AttemptNumber, not even an explicit nil
func (o *RestoreInfo) UnsetAttemptNumber() {
	o.AttemptNumber.Unset()
}

// GetCloudDeployTarget returns the CloudDeployTarget field value if set, zero value otherwise.
func (o *RestoreInfo) GetCloudDeployTarget() CloudDeployTargetDetails {
	if o == nil || o.CloudDeployTarget == nil {
		var ret CloudDeployTargetDetails
		return ret
	}
	return *o.CloudDeployTarget
}

// GetCloudDeployTargetOk returns a tuple with the CloudDeployTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreInfo) GetCloudDeployTargetOk() (*CloudDeployTargetDetails, bool) {
	if o == nil || o.CloudDeployTarget == nil {
		return nil, false
	}
	return o.CloudDeployTarget, true
}

// HasCloudDeployTarget returns a boolean if a field has been set.
func (o *RestoreInfo) HasCloudDeployTarget() bool {
	if o != nil && o.CloudDeployTarget != nil {
		return true
	}

	return false
}

// SetCloudDeployTarget gets a reference to the given CloudDeployTargetDetails and assigns it to the CloudDeployTarget field.
func (o *RestoreInfo) SetCloudDeployTarget(v CloudDeployTargetDetails) {
	o.CloudDeployTarget = &v
}

// GetJobRunId returns the JobRunId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestoreInfo) GetJobRunId() int64 {
	if o == nil || o.JobRunId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.JobRunId.Get()
}

// GetJobRunIdOk returns a tuple with the JobRunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestoreInfo) GetJobRunIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.JobRunId.Get(), o.JobRunId.IsSet()
}

// HasJobRunId returns a boolean if a field has been set.
func (o *RestoreInfo) HasJobRunId() bool {
	if o != nil && o.JobRunId.IsSet() {
		return true
	}

	return false
}

// SetJobRunId gets a reference to the given NullableInt64 and assigns it to the JobRunId field.
func (o *RestoreInfo) SetJobRunId(v int64) {
	o.JobRunId.Set(&v)
}
// SetJobRunIdNil sets the value for JobRunId to be an explicit nil
func (o *RestoreInfo) SetJobRunIdNil() {
	o.JobRunId.Set(nil)
}

// UnsetJobRunId ensures that no value is present for JobRunId, not even an explicit nil
func (o *RestoreInfo) UnsetJobRunId() {
	o.JobRunId.Unset()
}

// GetJobUid returns the JobUid field value if set, zero value otherwise.
func (o *RestoreInfo) GetJobUid() UniversalId {
	if o == nil || o.JobUid == nil {
		var ret UniversalId
		return ret
	}
	return *o.JobUid
}

// GetJobUidOk returns a tuple with the JobUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreInfo) GetJobUidOk() (*UniversalId, bool) {
	if o == nil || o.JobUid == nil {
		return nil, false
	}
	return o.JobUid, true
}

// HasJobUid returns a boolean if a field has been set.
func (o *RestoreInfo) HasJobUid() bool {
	if o != nil && o.JobUid != nil {
		return true
	}

	return false
}

// SetJobUid gets a reference to the given UniversalId and assigns it to the JobUid field.
func (o *RestoreInfo) SetJobUid(v UniversalId) {
	o.JobUid = &v
}

// GetParentSource returns the ParentSource field value if set, zero value otherwise.
func (o *RestoreInfo) GetParentSource() ProtectionSource {
	if o == nil || o.ParentSource == nil {
		var ret ProtectionSource
		return ret
	}
	return *o.ParentSource
}

// GetParentSourceOk returns a tuple with the ParentSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreInfo) GetParentSourceOk() (*ProtectionSource, bool) {
	if o == nil || o.ParentSource == nil {
		return nil, false
	}
	return o.ParentSource, true
}

// HasParentSource returns a boolean if a field has been set.
func (o *RestoreInfo) HasParentSource() bool {
	if o != nil && o.ParentSource != nil {
		return true
	}

	return false
}

// SetParentSource gets a reference to the given ProtectionSource and assigns it to the ParentSource field.
func (o *RestoreInfo) SetParentSource(v ProtectionSource) {
	o.ParentSource = &v
}

// GetRestoreTimeUsecs returns the RestoreTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestoreInfo) GetRestoreTimeUsecs() int64 {
	if o == nil || o.RestoreTimeUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RestoreTimeUsecs.Get()
}

// GetRestoreTimeUsecsOk returns a tuple with the RestoreTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestoreInfo) GetRestoreTimeUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RestoreTimeUsecs.Get(), o.RestoreTimeUsecs.IsSet()
}

// HasRestoreTimeUsecs returns a boolean if a field has been set.
func (o *RestoreInfo) HasRestoreTimeUsecs() bool {
	if o != nil && o.RestoreTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetRestoreTimeUsecs gets a reference to the given NullableInt64 and assigns it to the RestoreTimeUsecs field.
func (o *RestoreInfo) SetRestoreTimeUsecs(v int64) {
	o.RestoreTimeUsecs.Set(&v)
}
// SetRestoreTimeUsecsNil sets the value for RestoreTimeUsecs to be an explicit nil
func (o *RestoreInfo) SetRestoreTimeUsecsNil() {
	o.RestoreTimeUsecs.Set(nil)
}

// UnsetRestoreTimeUsecs ensures that no value is present for RestoreTimeUsecs, not even an explicit nil
func (o *RestoreInfo) UnsetRestoreTimeUsecs() {
	o.RestoreTimeUsecs.Unset()
}

// GetSnapshotRelativeDirPath returns the SnapshotRelativeDirPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestoreInfo) GetSnapshotRelativeDirPath() string {
	if o == nil || o.SnapshotRelativeDirPath.Get() == nil {
		var ret string
		return ret
	}
	return *o.SnapshotRelativeDirPath.Get()
}

// GetSnapshotRelativeDirPathOk returns a tuple with the SnapshotRelativeDirPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestoreInfo) GetSnapshotRelativeDirPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SnapshotRelativeDirPath.Get(), o.SnapshotRelativeDirPath.IsSet()
}

// HasSnapshotRelativeDirPath returns a boolean if a field has been set.
func (o *RestoreInfo) HasSnapshotRelativeDirPath() bool {
	if o != nil && o.SnapshotRelativeDirPath.IsSet() {
		return true
	}

	return false
}

// SetSnapshotRelativeDirPath gets a reference to the given NullableString and assigns it to the SnapshotRelativeDirPath field.
func (o *RestoreInfo) SetSnapshotRelativeDirPath(v string) {
	o.SnapshotRelativeDirPath.Set(&v)
}
// SetSnapshotRelativeDirPathNil sets the value for SnapshotRelativeDirPath to be an explicit nil
func (o *RestoreInfo) SetSnapshotRelativeDirPathNil() {
	o.SnapshotRelativeDirPath.Set(nil)
}

// UnsetSnapshotRelativeDirPath ensures that no value is present for SnapshotRelativeDirPath, not even an explicit nil
func (o *RestoreInfo) UnsetSnapshotRelativeDirPath() {
	o.SnapshotRelativeDirPath.Unset()
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *RestoreInfo) GetSource() ProtectionSource {
	if o == nil || o.Source == nil {
		var ret ProtectionSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreInfo) GetSourceOk() (*ProtectionSource, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *RestoreInfo) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given ProtectionSource and assigns it to the Source field.
func (o *RestoreInfo) SetSource(v ProtectionSource) {
	o.Source = &v
}

// GetStartTimeUsecs returns the StartTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestoreInfo) GetStartTimeUsecs() int64 {
	if o == nil || o.StartTimeUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.StartTimeUsecs.Get()
}

// GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestoreInfo) GetStartTimeUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartTimeUsecs.Get(), o.StartTimeUsecs.IsSet()
}

// HasStartTimeUsecs returns a boolean if a field has been set.
func (o *RestoreInfo) HasStartTimeUsecs() bool {
	if o != nil && o.StartTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetStartTimeUsecs gets a reference to the given NullableInt64 and assigns it to the StartTimeUsecs field.
func (o *RestoreInfo) SetStartTimeUsecs(v int64) {
	o.StartTimeUsecs.Set(&v)
}
// SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil
func (o *RestoreInfo) SetStartTimeUsecsNil() {
	o.StartTimeUsecs.Set(nil)
}

// UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
func (o *RestoreInfo) UnsetStartTimeUsecs() {
	o.StartTimeUsecs.Unset()
}

// GetViewName returns the ViewName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestoreInfo) GetViewName() string {
	if o == nil || o.ViewName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ViewName.Get()
}

// GetViewNameOk returns a tuple with the ViewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestoreInfo) GetViewNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ViewName.Get(), o.ViewName.IsSet()
}

// HasViewName returns a boolean if a field has been set.
func (o *RestoreInfo) HasViewName() bool {
	if o != nil && o.ViewName.IsSet() {
		return true
	}

	return false
}

// SetViewName gets a reference to the given NullableString and assigns it to the ViewName field.
func (o *RestoreInfo) SetViewName(v string) {
	o.ViewName.Set(&v)
}
// SetViewNameNil sets the value for ViewName to be an explicit nil
func (o *RestoreInfo) SetViewNameNil() {
	o.ViewName.Set(nil)
}

// UnsetViewName ensures that no value is present for ViewName, not even an explicit nil
func (o *RestoreInfo) UnsetViewName() {
	o.ViewName.Unset()
}

// GetVmHadIndependentDisks returns the VmHadIndependentDisks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestoreInfo) GetVmHadIndependentDisks() bool {
	if o == nil || o.VmHadIndependentDisks.Get() == nil {
		var ret bool
		return ret
	}
	return *o.VmHadIndependentDisks.Get()
}

// GetVmHadIndependentDisksOk returns a tuple with the VmHadIndependentDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestoreInfo) GetVmHadIndependentDisksOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VmHadIndependentDisks.Get(), o.VmHadIndependentDisks.IsSet()
}

// HasVmHadIndependentDisks returns a boolean if a field has been set.
func (o *RestoreInfo) HasVmHadIndependentDisks() bool {
	if o != nil && o.VmHadIndependentDisks.IsSet() {
		return true
	}

	return false
}

// SetVmHadIndependentDisks gets a reference to the given NullableBool and assigns it to the VmHadIndependentDisks field.
func (o *RestoreInfo) SetVmHadIndependentDisks(v bool) {
	o.VmHadIndependentDisks.Set(&v)
}
// SetVmHadIndependentDisksNil sets the value for VmHadIndependentDisks to be an explicit nil
func (o *RestoreInfo) SetVmHadIndependentDisksNil() {
	o.VmHadIndependentDisks.Set(nil)
}

// UnsetVmHadIndependentDisks ensures that no value is present for VmHadIndependentDisks, not even an explicit nil
func (o *RestoreInfo) UnsetVmHadIndependentDisks() {
	o.VmHadIndependentDisks.Unset()
}

func (o RestoreInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ArchivalTarget != nil {
		toSerialize["archivalTarget"] = o.ArchivalTarget
	}
	if o.AttemptNumber.IsSet() {
		toSerialize["attemptNumber"] = o.AttemptNumber.Get()
	}
	if o.CloudDeployTarget != nil {
		toSerialize["cloudDeployTarget"] = o.CloudDeployTarget
	}
	if o.JobRunId.IsSet() {
		toSerialize["jobRunId"] = o.JobRunId.Get()
	}
	if o.JobUid != nil {
		toSerialize["jobUid"] = o.JobUid
	}
	if o.ParentSource != nil {
		toSerialize["parentSource"] = o.ParentSource
	}
	if o.RestoreTimeUsecs.IsSet() {
		toSerialize["restoreTimeUsecs"] = o.RestoreTimeUsecs.Get()
	}
	if o.SnapshotRelativeDirPath.IsSet() {
		toSerialize["snapshotRelativeDirPath"] = o.SnapshotRelativeDirPath.Get()
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.StartTimeUsecs.IsSet() {
		toSerialize["startTimeUsecs"] = o.StartTimeUsecs.Get()
	}
	if o.ViewName.IsSet() {
		toSerialize["viewName"] = o.ViewName.Get()
	}
	if o.VmHadIndependentDisks.IsSet() {
		toSerialize["vmHadIndependentDisks"] = o.VmHadIndependentDisks.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRestoreInfo struct {
	value *RestoreInfo
	isSet bool
}

func (v NullableRestoreInfo) Get() *RestoreInfo {
	return v.value
}

func (v *NullableRestoreInfo) Set(val *RestoreInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRestoreInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRestoreInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestoreInfo(val *RestoreInfo) *NullableRestoreInfo {
	return &NullableRestoreInfo{value: val, isSet: true}
}

func (v NullableRestoreInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestoreInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


