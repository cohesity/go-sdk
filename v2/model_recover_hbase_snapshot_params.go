/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RecoverHbaseSnapshotParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoverHbaseSnapshotParams{}

// RecoverHbaseSnapshotParams Specifies the snapshot parameters for a protected Hbase object.
type RecoverHbaseSnapshotParams struct {
	ArchivalTargetInfo NullableCommonRecoverObjectSnapshotParamsArchivalTargetInfo `json:"archivalTargetInfo,omitempty"`
	// Specify the total bytes restored.
	BytesRestored NullableInt64 `json:"bytesRestored,omitempty"`
	// Specifies the end time of the Recovery in Unix timestamp epoch in microseconds. This field will be populated only after Recovery is finished.
	EndTimeUsecs NullableInt64 `json:"endTimeUsecs,omitempty"`
	// Specify error messages about the object.
	Messages []string `json:"messages,omitempty"`
	ObjectInfo NullableCommonRecoverObjectSnapshotParamsObjectInfo `json:"objectInfo,omitempty"`
	// Specifies the timestamp (in microseconds. from epoch) for recovering to a point-in-time in the past.
	PointInTimeUsecs NullableInt64 `json:"pointInTimeUsecs,omitempty"`
	// Progress monitor task id for Recovery of VM.
	ProgressTaskId NullableString `json:"progressTaskId,omitempty"`
	// Specifies the protection group id of the object snapshot.
	ProtectionGroupId NullableString `json:"protectionGroupId,omitempty"`
	// Specifies the protection group name of the object snapshot.
	ProtectionGroupName NullableString `json:"protectionGroupName,omitempty"`
	// Specifies that user wants to perform standby restore if it is enabled for this object.
	RecoverFromStandby NullableBool `json:"recoverFromStandby,omitempty"`
	// Specifies the time when the snapshot is created in Unix timestamp epoch in microseconds.
	SnapshotCreationTimeUsecs NullableInt64 `json:"snapshotCreationTimeUsecs,omitempty"`
	// Specifies the snapshot id.
	SnapshotId string `json:"snapshotId"`
	// Specifies the snapshot target type.
	SnapshotTargetType NullableString `json:"snapshotTargetType,omitempty"`
	// Specifies the start time of the Recovery in Unix timestamp epoch in microseconds.
	StartTimeUsecs NullableInt64 `json:"startTimeUsecs,omitempty"`
	// Status of the Recovery. 'Running' indicates that the Recovery is still running. 'Canceled' indicates that the Recovery has been cancelled. 'Canceling' indicates that the Recovery is in the process of being cancelled. 'Failed' indicates that the Recovery has failed. 'Succeeded' indicates that the Recovery has finished successfully. 'SucceededWithWarning' indicates that the Recovery finished successfully, but there were some warning messages. 'Skipped' indicates that the Recovery task was skipped.
	Status NullableString `json:"status,omitempty"`
	// Specifies the ID of the Storage Domain where this snapshot is stored.
	StorageDomainId NullableInt64 `json:"storageDomainId,omitempty"`
	// Specifies details of objects to be recovered.
	Objects []RecoverHbaseNoSqlObjectParams `json:"objects,omitempty"`
}

type _RecoverHbaseSnapshotParams RecoverHbaseSnapshotParams

// NewRecoverHbaseSnapshotParams instantiates a new RecoverHbaseSnapshotParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverHbaseSnapshotParams(snapshotId string) *RecoverHbaseSnapshotParams {
	this := RecoverHbaseSnapshotParams{}
	this.SnapshotId = snapshotId
	return &this
}

// NewRecoverHbaseSnapshotParamsWithDefaults instantiates a new RecoverHbaseSnapshotParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverHbaseSnapshotParamsWithDefaults() *RecoverHbaseSnapshotParams {
	this := RecoverHbaseSnapshotParams{}
	return &this
}

// GetArchivalTargetInfo returns the ArchivalTargetInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverHbaseSnapshotParams) GetArchivalTargetInfo() CommonRecoverObjectSnapshotParamsArchivalTargetInfo {
	if o == nil || IsNil(o.ArchivalTargetInfo.Get()) {
		var ret CommonRecoverObjectSnapshotParamsArchivalTargetInfo
		return ret
	}
	return *o.ArchivalTargetInfo.Get()
}

// GetArchivalTargetInfoOk returns a tuple with the ArchivalTargetInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverHbaseSnapshotParams) GetArchivalTargetInfoOk() (*CommonRecoverObjectSnapshotParamsArchivalTargetInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArchivalTargetInfo.Get(), o.ArchivalTargetInfo.IsSet()
}

// HasArchivalTargetInfo returns a boolean if a field has been set.
func (o *RecoverHbaseSnapshotParams) HasArchivalTargetInfo() bool {
	if o != nil && o.ArchivalTargetInfo.IsSet() {
		return true
	}

	return false
}

// SetArchivalTargetInfo gets a reference to the given NullableCommonRecoverObjectSnapshotParamsArchivalTargetInfo and assigns it to the ArchivalTargetInfo field.
func (o *RecoverHbaseSnapshotParams) SetArchivalTargetInfo(v CommonRecoverObjectSnapshotParamsArchivalTargetInfo) {
	o.ArchivalTargetInfo.Set(&v)
}
// SetArchivalTargetInfoNil sets the value for ArchivalTargetInfo to be an explicit nil
func (o *RecoverHbaseSnapshotParams) SetArchivalTargetInfoNil() {
	o.ArchivalTargetInfo.Set(nil)
}

// UnsetArchivalTargetInfo ensures that no value is present for ArchivalTargetInfo, not even an explicit nil
func (o *RecoverHbaseSnapshotParams) UnsetArchivalTargetInfo() {
	o.ArchivalTargetInfo.Unset()
}

// GetBytesRestored returns the BytesRestored field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverHbaseSnapshotParams) GetBytesRestored() int64 {
	if o == nil || IsNil(o.BytesRestored.Get()) {
		var ret int64
		return ret
	}
	return *o.BytesRestored.Get()
}

// GetBytesRestoredOk returns a tuple with the BytesRestored field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverHbaseSnapshotParams) GetBytesRestoredOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BytesRestored.Get(), o.BytesRestored.IsSet()
}

// HasBytesRestored returns a boolean if a field has been set.
func (o *RecoverHbaseSnapshotParams) HasBytesRestored() bool {
	if o != nil && o.BytesRestored.IsSet() {
		return true
	}

	return false
}

// SetBytesRestored gets a reference to the given NullableInt64 and assigns it to the BytesRestored field.
func (o *RecoverHbaseSnapshotParams) SetBytesRestored(v int64) {
	o.BytesRestored.Set(&v)
}
// SetBytesRestoredNil sets the value for BytesRestored to be an explicit nil
func (o *RecoverHbaseSnapshotParams) SetBytesRestoredNil() {
	o.BytesRestored.Set(nil)
}

// UnsetBytesRestored ensures that no value is present for BytesRestored, not even an explicit nil
func (o *RecoverHbaseSnapshotParams) UnsetBytesRestored() {
	o.BytesRestored.Unset()
}

// GetEndTimeUsecs returns the EndTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverHbaseSnapshotParams) GetEndTimeUsecs() int64 {
	if o == nil || IsNil(o.EndTimeUsecs.Get()) {
		var ret int64
		return ret
	}
	return *o.EndTimeUsecs.Get()
}

// GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverHbaseSnapshotParams) GetEndTimeUsecsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndTimeUsecs.Get(), o.EndTimeUsecs.IsSet()
}

// HasEndTimeUsecs returns a boolean if a field has been set.
func (o *RecoverHbaseSnapshotParams) HasEndTimeUsecs() bool {
	if o != nil && o.EndTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetEndTimeUsecs gets a reference to the given NullableInt64 and assigns it to the EndTimeUsecs field.
func (o *RecoverHbaseSnapshotParams) SetEndTimeUsecs(v int64) {
	o.EndTimeUsecs.Set(&v)
}
// SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil
func (o *RecoverHbaseSnapshotParams) SetEndTimeUsecsNil() {
	o.EndTimeUsecs.Set(nil)
}

// UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
func (o *RecoverHbaseSnapshotParams) UnsetEndTimeUsecs() {
	o.EndTimeUsecs.Unset()
}

// GetMessages returns the Messages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverHbaseSnapshotParams) GetMessages() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverHbaseSnapshotParams) GetMessagesOk() ([]string, bool) {
	if o == nil || IsNil(o.Messages) {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *RecoverHbaseSnapshotParams) HasMessages() bool {
	if o != nil && !IsNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []string and assigns it to the Messages field.
func (o *RecoverHbaseSnapshotParams) SetMessages(v []string) {
	o.Messages = v
}

// GetObjectInfo returns the ObjectInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverHbaseSnapshotParams) GetObjectInfo() CommonRecoverObjectSnapshotParamsObjectInfo {
	if o == nil || IsNil(o.ObjectInfo.Get()) {
		var ret CommonRecoverObjectSnapshotParamsObjectInfo
		return ret
	}
	return *o.ObjectInfo.Get()
}

// GetObjectInfoOk returns a tuple with the ObjectInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverHbaseSnapshotParams) GetObjectInfoOk() (*CommonRecoverObjectSnapshotParamsObjectInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObjectInfo.Get(), o.ObjectInfo.IsSet()
}

// HasObjectInfo returns a boolean if a field has been set.
func (o *RecoverHbaseSnapshotParams) HasObjectInfo() bool {
	if o != nil && o.ObjectInfo.IsSet() {
		return true
	}

	return false
}

// SetObjectInfo gets a reference to the given NullableCommonRecoverObjectSnapshotParamsObjectInfo and assigns it to the ObjectInfo field.
func (o *RecoverHbaseSnapshotParams) SetObjectInfo(v CommonRecoverObjectSnapshotParamsObjectInfo) {
	o.ObjectInfo.Set(&v)
}
// SetObjectInfoNil sets the value for ObjectInfo to be an explicit nil
func (o *RecoverHbaseSnapshotParams) SetObjectInfoNil() {
	o.ObjectInfo.Set(nil)
}

// UnsetObjectInfo ensures that no value is present for ObjectInfo, not even an explicit nil
func (o *RecoverHbaseSnapshotParams) UnsetObjectInfo() {
	o.ObjectInfo.Unset()
}

// GetPointInTimeUsecs returns the PointInTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverHbaseSnapshotParams) GetPointInTimeUsecs() int64 {
	if o == nil || IsNil(o.PointInTimeUsecs.Get()) {
		var ret int64
		return ret
	}
	return *o.PointInTimeUsecs.Get()
}

// GetPointInTimeUsecsOk returns a tuple with the PointInTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverHbaseSnapshotParams) GetPointInTimeUsecsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PointInTimeUsecs.Get(), o.PointInTimeUsecs.IsSet()
}

// HasPointInTimeUsecs returns a boolean if a field has been set.
func (o *RecoverHbaseSnapshotParams) HasPointInTimeUsecs() bool {
	if o != nil && o.PointInTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetPointInTimeUsecs gets a reference to the given NullableInt64 and assigns it to the PointInTimeUsecs field.
func (o *RecoverHbaseSnapshotParams) SetPointInTimeUsecs(v int64) {
	o.PointInTimeUsecs.Set(&v)
}
// SetPointInTimeUsecsNil sets the value for PointInTimeUsecs to be an explicit nil
func (o *RecoverHbaseSnapshotParams) SetPointInTimeUsecsNil() {
	o.PointInTimeUsecs.Set(nil)
}

// UnsetPointInTimeUsecs ensures that no value is present for PointInTimeUsecs, not even an explicit nil
func (o *RecoverHbaseSnapshotParams) UnsetPointInTimeUsecs() {
	o.PointInTimeUsecs.Unset()
}

// GetProgressTaskId returns the ProgressTaskId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverHbaseSnapshotParams) GetProgressTaskId() string {
	if o == nil || IsNil(o.ProgressTaskId.Get()) {
		var ret string
		return ret
	}
	return *o.ProgressTaskId.Get()
}

// GetProgressTaskIdOk returns a tuple with the ProgressTaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverHbaseSnapshotParams) GetProgressTaskIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProgressTaskId.Get(), o.ProgressTaskId.IsSet()
}

// HasProgressTaskId returns a boolean if a field has been set.
func (o *RecoverHbaseSnapshotParams) HasProgressTaskId() bool {
	if o != nil && o.ProgressTaskId.IsSet() {
		return true
	}

	return false
}

// SetProgressTaskId gets a reference to the given NullableString and assigns it to the ProgressTaskId field.
func (o *RecoverHbaseSnapshotParams) SetProgressTaskId(v string) {
	o.ProgressTaskId.Set(&v)
}
// SetProgressTaskIdNil sets the value for ProgressTaskId to be an explicit nil
func (o *RecoverHbaseSnapshotParams) SetProgressTaskIdNil() {
	o.ProgressTaskId.Set(nil)
}

// UnsetProgressTaskId ensures that no value is present for ProgressTaskId, not even an explicit nil
func (o *RecoverHbaseSnapshotParams) UnsetProgressTaskId() {
	o.ProgressTaskId.Unset()
}

// GetProtectionGroupId returns the ProtectionGroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverHbaseSnapshotParams) GetProtectionGroupId() string {
	if o == nil || IsNil(o.ProtectionGroupId.Get()) {
		var ret string
		return ret
	}
	return *o.ProtectionGroupId.Get()
}

// GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverHbaseSnapshotParams) GetProtectionGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProtectionGroupId.Get(), o.ProtectionGroupId.IsSet()
}

// HasProtectionGroupId returns a boolean if a field has been set.
func (o *RecoverHbaseSnapshotParams) HasProtectionGroupId() bool {
	if o != nil && o.ProtectionGroupId.IsSet() {
		return true
	}

	return false
}

// SetProtectionGroupId gets a reference to the given NullableString and assigns it to the ProtectionGroupId field.
func (o *RecoverHbaseSnapshotParams) SetProtectionGroupId(v string) {
	o.ProtectionGroupId.Set(&v)
}
// SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil
func (o *RecoverHbaseSnapshotParams) SetProtectionGroupIdNil() {
	o.ProtectionGroupId.Set(nil)
}

// UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
func (o *RecoverHbaseSnapshotParams) UnsetProtectionGroupId() {
	o.ProtectionGroupId.Unset()
}

// GetProtectionGroupName returns the ProtectionGroupName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverHbaseSnapshotParams) GetProtectionGroupName() string {
	if o == nil || IsNil(o.ProtectionGroupName.Get()) {
		var ret string
		return ret
	}
	return *o.ProtectionGroupName.Get()
}

// GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverHbaseSnapshotParams) GetProtectionGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProtectionGroupName.Get(), o.ProtectionGroupName.IsSet()
}

// HasProtectionGroupName returns a boolean if a field has been set.
func (o *RecoverHbaseSnapshotParams) HasProtectionGroupName() bool {
	if o != nil && o.ProtectionGroupName.IsSet() {
		return true
	}

	return false
}

// SetProtectionGroupName gets a reference to the given NullableString and assigns it to the ProtectionGroupName field.
func (o *RecoverHbaseSnapshotParams) SetProtectionGroupName(v string) {
	o.ProtectionGroupName.Set(&v)
}
// SetProtectionGroupNameNil sets the value for ProtectionGroupName to be an explicit nil
func (o *RecoverHbaseSnapshotParams) SetProtectionGroupNameNil() {
	o.ProtectionGroupName.Set(nil)
}

// UnsetProtectionGroupName ensures that no value is present for ProtectionGroupName, not even an explicit nil
func (o *RecoverHbaseSnapshotParams) UnsetProtectionGroupName() {
	o.ProtectionGroupName.Unset()
}

// GetRecoverFromStandby returns the RecoverFromStandby field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverHbaseSnapshotParams) GetRecoverFromStandby() bool {
	if o == nil || IsNil(o.RecoverFromStandby.Get()) {
		var ret bool
		return ret
	}
	return *o.RecoverFromStandby.Get()
}

// GetRecoverFromStandbyOk returns a tuple with the RecoverFromStandby field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverHbaseSnapshotParams) GetRecoverFromStandbyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecoverFromStandby.Get(), o.RecoverFromStandby.IsSet()
}

// HasRecoverFromStandby returns a boolean if a field has been set.
func (o *RecoverHbaseSnapshotParams) HasRecoverFromStandby() bool {
	if o != nil && o.RecoverFromStandby.IsSet() {
		return true
	}

	return false
}

// SetRecoverFromStandby gets a reference to the given NullableBool and assigns it to the RecoverFromStandby field.
func (o *RecoverHbaseSnapshotParams) SetRecoverFromStandby(v bool) {
	o.RecoverFromStandby.Set(&v)
}
// SetRecoverFromStandbyNil sets the value for RecoverFromStandby to be an explicit nil
func (o *RecoverHbaseSnapshotParams) SetRecoverFromStandbyNil() {
	o.RecoverFromStandby.Set(nil)
}

// UnsetRecoverFromStandby ensures that no value is present for RecoverFromStandby, not even an explicit nil
func (o *RecoverHbaseSnapshotParams) UnsetRecoverFromStandby() {
	o.RecoverFromStandby.Unset()
}

// GetSnapshotCreationTimeUsecs returns the SnapshotCreationTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverHbaseSnapshotParams) GetSnapshotCreationTimeUsecs() int64 {
	if o == nil || IsNil(o.SnapshotCreationTimeUsecs.Get()) {
		var ret int64
		return ret
	}
	return *o.SnapshotCreationTimeUsecs.Get()
}

// GetSnapshotCreationTimeUsecsOk returns a tuple with the SnapshotCreationTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverHbaseSnapshotParams) GetSnapshotCreationTimeUsecsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SnapshotCreationTimeUsecs.Get(), o.SnapshotCreationTimeUsecs.IsSet()
}

// HasSnapshotCreationTimeUsecs returns a boolean if a field has been set.
func (o *RecoverHbaseSnapshotParams) HasSnapshotCreationTimeUsecs() bool {
	if o != nil && o.SnapshotCreationTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetSnapshotCreationTimeUsecs gets a reference to the given NullableInt64 and assigns it to the SnapshotCreationTimeUsecs field.
func (o *RecoverHbaseSnapshotParams) SetSnapshotCreationTimeUsecs(v int64) {
	o.SnapshotCreationTimeUsecs.Set(&v)
}
// SetSnapshotCreationTimeUsecsNil sets the value for SnapshotCreationTimeUsecs to be an explicit nil
func (o *RecoverHbaseSnapshotParams) SetSnapshotCreationTimeUsecsNil() {
	o.SnapshotCreationTimeUsecs.Set(nil)
}

// UnsetSnapshotCreationTimeUsecs ensures that no value is present for SnapshotCreationTimeUsecs, not even an explicit nil
func (o *RecoverHbaseSnapshotParams) UnsetSnapshotCreationTimeUsecs() {
	o.SnapshotCreationTimeUsecs.Unset()
}

// GetSnapshotId returns the SnapshotId field value
func (o *RecoverHbaseSnapshotParams) GetSnapshotId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value
// and a boolean to check if the value has been set.
func (o *RecoverHbaseSnapshotParams) GetSnapshotIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SnapshotId, true
}

// SetSnapshotId sets field value
func (o *RecoverHbaseSnapshotParams) SetSnapshotId(v string) {
	o.SnapshotId = v
}

// GetSnapshotTargetType returns the SnapshotTargetType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverHbaseSnapshotParams) GetSnapshotTargetType() string {
	if o == nil || IsNil(o.SnapshotTargetType.Get()) {
		var ret string
		return ret
	}
	return *o.SnapshotTargetType.Get()
}

// GetSnapshotTargetTypeOk returns a tuple with the SnapshotTargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverHbaseSnapshotParams) GetSnapshotTargetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SnapshotTargetType.Get(), o.SnapshotTargetType.IsSet()
}

// HasSnapshotTargetType returns a boolean if a field has been set.
func (o *RecoverHbaseSnapshotParams) HasSnapshotTargetType() bool {
	if o != nil && o.SnapshotTargetType.IsSet() {
		return true
	}

	return false
}

// SetSnapshotTargetType gets a reference to the given NullableString and assigns it to the SnapshotTargetType field.
func (o *RecoverHbaseSnapshotParams) SetSnapshotTargetType(v string) {
	o.SnapshotTargetType.Set(&v)
}
// SetSnapshotTargetTypeNil sets the value for SnapshotTargetType to be an explicit nil
func (o *RecoverHbaseSnapshotParams) SetSnapshotTargetTypeNil() {
	o.SnapshotTargetType.Set(nil)
}

// UnsetSnapshotTargetType ensures that no value is present for SnapshotTargetType, not even an explicit nil
func (o *RecoverHbaseSnapshotParams) UnsetSnapshotTargetType() {
	o.SnapshotTargetType.Unset()
}

// GetStartTimeUsecs returns the StartTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverHbaseSnapshotParams) GetStartTimeUsecs() int64 {
	if o == nil || IsNil(o.StartTimeUsecs.Get()) {
		var ret int64
		return ret
	}
	return *o.StartTimeUsecs.Get()
}

// GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverHbaseSnapshotParams) GetStartTimeUsecsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartTimeUsecs.Get(), o.StartTimeUsecs.IsSet()
}

// HasStartTimeUsecs returns a boolean if a field has been set.
func (o *RecoverHbaseSnapshotParams) HasStartTimeUsecs() bool {
	if o != nil && o.StartTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetStartTimeUsecs gets a reference to the given NullableInt64 and assigns it to the StartTimeUsecs field.
func (o *RecoverHbaseSnapshotParams) SetStartTimeUsecs(v int64) {
	o.StartTimeUsecs.Set(&v)
}
// SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil
func (o *RecoverHbaseSnapshotParams) SetStartTimeUsecsNil() {
	o.StartTimeUsecs.Set(nil)
}

// UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
func (o *RecoverHbaseSnapshotParams) UnsetStartTimeUsecs() {
	o.StartTimeUsecs.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverHbaseSnapshotParams) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverHbaseSnapshotParams) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *RecoverHbaseSnapshotParams) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *RecoverHbaseSnapshotParams) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *RecoverHbaseSnapshotParams) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *RecoverHbaseSnapshotParams) UnsetStatus() {
	o.Status.Unset()
}

// GetStorageDomainId returns the StorageDomainId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverHbaseSnapshotParams) GetStorageDomainId() int64 {
	if o == nil || IsNil(o.StorageDomainId.Get()) {
		var ret int64
		return ret
	}
	return *o.StorageDomainId.Get()
}

// GetStorageDomainIdOk returns a tuple with the StorageDomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverHbaseSnapshotParams) GetStorageDomainIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageDomainId.Get(), o.StorageDomainId.IsSet()
}

// HasStorageDomainId returns a boolean if a field has been set.
func (o *RecoverHbaseSnapshotParams) HasStorageDomainId() bool {
	if o != nil && o.StorageDomainId.IsSet() {
		return true
	}

	return false
}

// SetStorageDomainId gets a reference to the given NullableInt64 and assigns it to the StorageDomainId field.
func (o *RecoverHbaseSnapshotParams) SetStorageDomainId(v int64) {
	o.StorageDomainId.Set(&v)
}
// SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil
func (o *RecoverHbaseSnapshotParams) SetStorageDomainIdNil() {
	o.StorageDomainId.Set(nil)
}

// UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
func (o *RecoverHbaseSnapshotParams) UnsetStorageDomainId() {
	o.StorageDomainId.Unset()
}

// GetObjects returns the Objects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverHbaseSnapshotParams) GetObjects() []RecoverHbaseNoSqlObjectParams {
	if o == nil {
		var ret []RecoverHbaseNoSqlObjectParams
		return ret
	}
	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverHbaseSnapshotParams) GetObjectsOk() ([]RecoverHbaseNoSqlObjectParams, bool) {
	if o == nil || IsNil(o.Objects) {
		return nil, false
	}
	return o.Objects, true
}

// HasObjects returns a boolean if a field has been set.
func (o *RecoverHbaseSnapshotParams) HasObjects() bool {
	if o != nil && !IsNil(o.Objects) {
		return true
	}

	return false
}

// SetObjects gets a reference to the given []RecoverHbaseNoSqlObjectParams and assigns it to the Objects field.
func (o *RecoverHbaseSnapshotParams) SetObjects(v []RecoverHbaseNoSqlObjectParams) {
	o.Objects = v
}

func (o RecoverHbaseSnapshotParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoverHbaseSnapshotParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ArchivalTargetInfo.IsSet() {
		toSerialize["archivalTargetInfo"] = o.ArchivalTargetInfo.Get()
	}
	if o.BytesRestored.IsSet() {
		toSerialize["bytesRestored"] = o.BytesRestored.Get()
	}
	if o.EndTimeUsecs.IsSet() {
		toSerialize["endTimeUsecs"] = o.EndTimeUsecs.Get()
	}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}
	if o.ObjectInfo.IsSet() {
		toSerialize["objectInfo"] = o.ObjectInfo.Get()
	}
	if o.PointInTimeUsecs.IsSet() {
		toSerialize["pointInTimeUsecs"] = o.PointInTimeUsecs.Get()
	}
	if o.ProgressTaskId.IsSet() {
		toSerialize["progressTaskId"] = o.ProgressTaskId.Get()
	}
	if o.ProtectionGroupId.IsSet() {
		toSerialize["protectionGroupId"] = o.ProtectionGroupId.Get()
	}
	if o.ProtectionGroupName.IsSet() {
		toSerialize["protectionGroupName"] = o.ProtectionGroupName.Get()
	}
	if o.RecoverFromStandby.IsSet() {
		toSerialize["recoverFromStandby"] = o.RecoverFromStandby.Get()
	}
	if o.SnapshotCreationTimeUsecs.IsSet() {
		toSerialize["snapshotCreationTimeUsecs"] = o.SnapshotCreationTimeUsecs.Get()
	}
	toSerialize["snapshotId"] = o.SnapshotId
	if o.SnapshotTargetType.IsSet() {
		toSerialize["snapshotTargetType"] = o.SnapshotTargetType.Get()
	}
	if o.StartTimeUsecs.IsSet() {
		toSerialize["startTimeUsecs"] = o.StartTimeUsecs.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.StorageDomainId.IsSet() {
		toSerialize["storageDomainId"] = o.StorageDomainId.Get()
	}
	if o.Objects != nil {
		toSerialize["objects"] = o.Objects
	}
	return toSerialize, nil
}

func (o *RecoverHbaseSnapshotParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"snapshotId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRecoverHbaseSnapshotParams := _RecoverHbaseSnapshotParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRecoverHbaseSnapshotParams)

	if err != nil {
		return err
	}

	*o = RecoverHbaseSnapshotParams(varRecoverHbaseSnapshotParams)

	return err
}

type NullableRecoverHbaseSnapshotParams struct {
	value *RecoverHbaseSnapshotParams
	isSet bool
}

func (v NullableRecoverHbaseSnapshotParams) Get() *RecoverHbaseSnapshotParams {
	return v.value
}

func (v *NullableRecoverHbaseSnapshotParams) Set(val *RecoverHbaseSnapshotParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverHbaseSnapshotParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverHbaseSnapshotParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverHbaseSnapshotParams(val *RecoverHbaseSnapshotParams) *NullableRecoverHbaseSnapshotParams {
	return &NullableRecoverHbaseSnapshotParams{value: val, isSet: true}
}

func (v NullableRecoverHbaseSnapshotParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverHbaseSnapshotParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


