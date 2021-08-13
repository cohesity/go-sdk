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

// SnapshotInfo Snapshot info for an object.
type SnapshotInfo struct {
	// Snapshot id for a successful snapshot. This field will not be set if the Protection Group Run has no successful attempt.
	SnapshotId NullableString `json:"snapshotId,omitempty"`
	// Status of snapshot.
	Status NullableString `json:"status,omitempty"`
	// Specifies the start time of attempt in Unix epoch Timestamp(in microseconds) for an object.
	StartTimeUsecs NullableInt64 `json:"startTimeUsecs,omitempty"`
	// Specifies the end time of attempt in Unix epoch Timestamp(in microseconds) for an object.
	EndTimeUsecs NullableInt64 `json:"endTimeUsecs,omitempty"`
	// Specifies the time at which the backup task was admitted to run in Unix epoch Timestamp(in microseconds) for an object.
	AdmittedTimeUsecs NullableInt64 `json:"admittedTimeUsecs,omitempty"`
	// Specifies the time when gatekeeper permit is granted to the backup task. If the backup task is rescheduled due to errors, the field is updated to the time when permit is granted again.
	PermitGrantTimeUsecs NullableInt64 `json:"permitGrantTimeUsecs,omitempty"`
	// Specifies the duration between the startTime and when gatekeeper permit is granted to the backup task. If the backup task is rescheduled due to errors, the field is updated considering the time when permit is granted again. Queue duration = PermitGrantTimeUsecs - StartTimeUsecs
	QueueDurationUsecs NullableInt64 `json:"queueDurationUsecs,omitempty"`
	// Specifies the time at which the source snapshot was taken in Unix epoch Timestamp(in microseconds) for an object.
	SnapshotCreationTimeUsecs NullableInt64 `json:"snapshotCreationTimeUsecs,omitempty"`
	Stats *BackupDataStats `json:"stats,omitempty"`
	// Progress monitor task for an object.
	ProgressTaskId NullableString `json:"progressTaskId,omitempty"`
	// Specifies a list of warning messages.
	Warnings []string `json:"warnings,omitempty"`
	// Specifies whether the snapshot is deleted manually.
	IsManuallyDeleted NullableBool `json:"isManuallyDeleted,omitempty"`
	// Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds) for an object.
	ExpiryTimeUsecs NullableInt64 `json:"expiryTimeUsecs,omitempty"`
	// The total number of file and directory entities visited in this backup. Only applicable to file based backups.
	TotalFileCount NullableInt64 `json:"totalFileCount,omitempty"`
	// The total number of file and directory entities that are backed up in this run. Only applicable to file based backups.
	BackupFileCount NullableInt64 `json:"backupFileCount,omitempty"`
	DataLockConstraints *DataLockConstraints `json:"dataLockConstraints,omitempty"`
}

// NewSnapshotInfo instantiates a new SnapshotInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotInfo() *SnapshotInfo {
	this := SnapshotInfo{}
	return &this
}

// NewSnapshotInfoWithDefaults instantiates a new SnapshotInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotInfoWithDefaults() *SnapshotInfo {
	this := SnapshotInfo{}
	return &this
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnapshotInfo) GetSnapshotId() string {
	if o == nil || o.SnapshotId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SnapshotId.Get()
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotInfo) GetSnapshotIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SnapshotId.Get(), o.SnapshotId.IsSet()
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *SnapshotInfo) HasSnapshotId() bool {
	if o != nil && o.SnapshotId.IsSet() {
		return true
	}

	return false
}

// SetSnapshotId gets a reference to the given NullableString and assigns it to the SnapshotId field.
func (o *SnapshotInfo) SetSnapshotId(v string) {
	o.SnapshotId.Set(&v)
}
// SetSnapshotIdNil sets the value for SnapshotId to be an explicit nil
func (o *SnapshotInfo) SetSnapshotIdNil() {
	o.SnapshotId.Set(nil)
}

// UnsetSnapshotId ensures that no value is present for SnapshotId, not even an explicit nil
func (o *SnapshotInfo) UnsetSnapshotId() {
	o.SnapshotId.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnapshotInfo) GetStatus() string {
	if o == nil || o.Status.Get() == nil {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotInfo) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *SnapshotInfo) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *SnapshotInfo) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *SnapshotInfo) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *SnapshotInfo) UnsetStatus() {
	o.Status.Unset()
}

// GetStartTimeUsecs returns the StartTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnapshotInfo) GetStartTimeUsecs() int64 {
	if o == nil || o.StartTimeUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.StartTimeUsecs.Get()
}

// GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotInfo) GetStartTimeUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartTimeUsecs.Get(), o.StartTimeUsecs.IsSet()
}

// HasStartTimeUsecs returns a boolean if a field has been set.
func (o *SnapshotInfo) HasStartTimeUsecs() bool {
	if o != nil && o.StartTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetStartTimeUsecs gets a reference to the given NullableInt64 and assigns it to the StartTimeUsecs field.
func (o *SnapshotInfo) SetStartTimeUsecs(v int64) {
	o.StartTimeUsecs.Set(&v)
}
// SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil
func (o *SnapshotInfo) SetStartTimeUsecsNil() {
	o.StartTimeUsecs.Set(nil)
}

// UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
func (o *SnapshotInfo) UnsetStartTimeUsecs() {
	o.StartTimeUsecs.Unset()
}

// GetEndTimeUsecs returns the EndTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnapshotInfo) GetEndTimeUsecs() int64 {
	if o == nil || o.EndTimeUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.EndTimeUsecs.Get()
}

// GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotInfo) GetEndTimeUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndTimeUsecs.Get(), o.EndTimeUsecs.IsSet()
}

// HasEndTimeUsecs returns a boolean if a field has been set.
func (o *SnapshotInfo) HasEndTimeUsecs() bool {
	if o != nil && o.EndTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetEndTimeUsecs gets a reference to the given NullableInt64 and assigns it to the EndTimeUsecs field.
func (o *SnapshotInfo) SetEndTimeUsecs(v int64) {
	o.EndTimeUsecs.Set(&v)
}
// SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil
func (o *SnapshotInfo) SetEndTimeUsecsNil() {
	o.EndTimeUsecs.Set(nil)
}

// UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
func (o *SnapshotInfo) UnsetEndTimeUsecs() {
	o.EndTimeUsecs.Unset()
}

// GetAdmittedTimeUsecs returns the AdmittedTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnapshotInfo) GetAdmittedTimeUsecs() int64 {
	if o == nil || o.AdmittedTimeUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AdmittedTimeUsecs.Get()
}

// GetAdmittedTimeUsecsOk returns a tuple with the AdmittedTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotInfo) GetAdmittedTimeUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AdmittedTimeUsecs.Get(), o.AdmittedTimeUsecs.IsSet()
}

// HasAdmittedTimeUsecs returns a boolean if a field has been set.
func (o *SnapshotInfo) HasAdmittedTimeUsecs() bool {
	if o != nil && o.AdmittedTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetAdmittedTimeUsecs gets a reference to the given NullableInt64 and assigns it to the AdmittedTimeUsecs field.
func (o *SnapshotInfo) SetAdmittedTimeUsecs(v int64) {
	o.AdmittedTimeUsecs.Set(&v)
}
// SetAdmittedTimeUsecsNil sets the value for AdmittedTimeUsecs to be an explicit nil
func (o *SnapshotInfo) SetAdmittedTimeUsecsNil() {
	o.AdmittedTimeUsecs.Set(nil)
}

// UnsetAdmittedTimeUsecs ensures that no value is present for AdmittedTimeUsecs, not even an explicit nil
func (o *SnapshotInfo) UnsetAdmittedTimeUsecs() {
	o.AdmittedTimeUsecs.Unset()
}

// GetPermitGrantTimeUsecs returns the PermitGrantTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnapshotInfo) GetPermitGrantTimeUsecs() int64 {
	if o == nil || o.PermitGrantTimeUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.PermitGrantTimeUsecs.Get()
}

// GetPermitGrantTimeUsecsOk returns a tuple with the PermitGrantTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotInfo) GetPermitGrantTimeUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PermitGrantTimeUsecs.Get(), o.PermitGrantTimeUsecs.IsSet()
}

// HasPermitGrantTimeUsecs returns a boolean if a field has been set.
func (o *SnapshotInfo) HasPermitGrantTimeUsecs() bool {
	if o != nil && o.PermitGrantTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetPermitGrantTimeUsecs gets a reference to the given NullableInt64 and assigns it to the PermitGrantTimeUsecs field.
func (o *SnapshotInfo) SetPermitGrantTimeUsecs(v int64) {
	o.PermitGrantTimeUsecs.Set(&v)
}
// SetPermitGrantTimeUsecsNil sets the value for PermitGrantTimeUsecs to be an explicit nil
func (o *SnapshotInfo) SetPermitGrantTimeUsecsNil() {
	o.PermitGrantTimeUsecs.Set(nil)
}

// UnsetPermitGrantTimeUsecs ensures that no value is present for PermitGrantTimeUsecs, not even an explicit nil
func (o *SnapshotInfo) UnsetPermitGrantTimeUsecs() {
	o.PermitGrantTimeUsecs.Unset()
}

// GetQueueDurationUsecs returns the QueueDurationUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnapshotInfo) GetQueueDurationUsecs() int64 {
	if o == nil || o.QueueDurationUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.QueueDurationUsecs.Get()
}

// GetQueueDurationUsecsOk returns a tuple with the QueueDurationUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotInfo) GetQueueDurationUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.QueueDurationUsecs.Get(), o.QueueDurationUsecs.IsSet()
}

// HasQueueDurationUsecs returns a boolean if a field has been set.
func (o *SnapshotInfo) HasQueueDurationUsecs() bool {
	if o != nil && o.QueueDurationUsecs.IsSet() {
		return true
	}

	return false
}

// SetQueueDurationUsecs gets a reference to the given NullableInt64 and assigns it to the QueueDurationUsecs field.
func (o *SnapshotInfo) SetQueueDurationUsecs(v int64) {
	o.QueueDurationUsecs.Set(&v)
}
// SetQueueDurationUsecsNil sets the value for QueueDurationUsecs to be an explicit nil
func (o *SnapshotInfo) SetQueueDurationUsecsNil() {
	o.QueueDurationUsecs.Set(nil)
}

// UnsetQueueDurationUsecs ensures that no value is present for QueueDurationUsecs, not even an explicit nil
func (o *SnapshotInfo) UnsetQueueDurationUsecs() {
	o.QueueDurationUsecs.Unset()
}

// GetSnapshotCreationTimeUsecs returns the SnapshotCreationTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnapshotInfo) GetSnapshotCreationTimeUsecs() int64 {
	if o == nil || o.SnapshotCreationTimeUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SnapshotCreationTimeUsecs.Get()
}

// GetSnapshotCreationTimeUsecsOk returns a tuple with the SnapshotCreationTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotInfo) GetSnapshotCreationTimeUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SnapshotCreationTimeUsecs.Get(), o.SnapshotCreationTimeUsecs.IsSet()
}

// HasSnapshotCreationTimeUsecs returns a boolean if a field has been set.
func (o *SnapshotInfo) HasSnapshotCreationTimeUsecs() bool {
	if o != nil && o.SnapshotCreationTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetSnapshotCreationTimeUsecs gets a reference to the given NullableInt64 and assigns it to the SnapshotCreationTimeUsecs field.
func (o *SnapshotInfo) SetSnapshotCreationTimeUsecs(v int64) {
	o.SnapshotCreationTimeUsecs.Set(&v)
}
// SetSnapshotCreationTimeUsecsNil sets the value for SnapshotCreationTimeUsecs to be an explicit nil
func (o *SnapshotInfo) SetSnapshotCreationTimeUsecsNil() {
	o.SnapshotCreationTimeUsecs.Set(nil)
}

// UnsetSnapshotCreationTimeUsecs ensures that no value is present for SnapshotCreationTimeUsecs, not even an explicit nil
func (o *SnapshotInfo) UnsetSnapshotCreationTimeUsecs() {
	o.SnapshotCreationTimeUsecs.Unset()
}

// GetStats returns the Stats field value if set, zero value otherwise.
func (o *SnapshotInfo) GetStats() BackupDataStats {
	if o == nil || o.Stats == nil {
		var ret BackupDataStats
		return ret
	}
	return *o.Stats
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotInfo) GetStatsOk() (*BackupDataStats, bool) {
	if o == nil || o.Stats == nil {
		return nil, false
	}
	return o.Stats, true
}

// HasStats returns a boolean if a field has been set.
func (o *SnapshotInfo) HasStats() bool {
	if o != nil && o.Stats != nil {
		return true
	}

	return false
}

// SetStats gets a reference to the given BackupDataStats and assigns it to the Stats field.
func (o *SnapshotInfo) SetStats(v BackupDataStats) {
	o.Stats = &v
}

// GetProgressTaskId returns the ProgressTaskId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnapshotInfo) GetProgressTaskId() string {
	if o == nil || o.ProgressTaskId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProgressTaskId.Get()
}

// GetProgressTaskIdOk returns a tuple with the ProgressTaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotInfo) GetProgressTaskIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProgressTaskId.Get(), o.ProgressTaskId.IsSet()
}

// HasProgressTaskId returns a boolean if a field has been set.
func (o *SnapshotInfo) HasProgressTaskId() bool {
	if o != nil && o.ProgressTaskId.IsSet() {
		return true
	}

	return false
}

// SetProgressTaskId gets a reference to the given NullableString and assigns it to the ProgressTaskId field.
func (o *SnapshotInfo) SetProgressTaskId(v string) {
	o.ProgressTaskId.Set(&v)
}
// SetProgressTaskIdNil sets the value for ProgressTaskId to be an explicit nil
func (o *SnapshotInfo) SetProgressTaskIdNil() {
	o.ProgressTaskId.Set(nil)
}

// UnsetProgressTaskId ensures that no value is present for ProgressTaskId, not even an explicit nil
func (o *SnapshotInfo) UnsetProgressTaskId() {
	o.ProgressTaskId.Unset()
}

// GetWarnings returns the Warnings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnapshotInfo) GetWarnings() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotInfo) GetWarningsOk() (*[]string, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return &o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *SnapshotInfo) HasWarnings() bool {
	if o != nil && o.Warnings != nil {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *SnapshotInfo) SetWarnings(v []string) {
	o.Warnings = v
}

// GetIsManuallyDeleted returns the IsManuallyDeleted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnapshotInfo) GetIsManuallyDeleted() bool {
	if o == nil || o.IsManuallyDeleted.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsManuallyDeleted.Get()
}

// GetIsManuallyDeletedOk returns a tuple with the IsManuallyDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotInfo) GetIsManuallyDeletedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsManuallyDeleted.Get(), o.IsManuallyDeleted.IsSet()
}

// HasIsManuallyDeleted returns a boolean if a field has been set.
func (o *SnapshotInfo) HasIsManuallyDeleted() bool {
	if o != nil && o.IsManuallyDeleted.IsSet() {
		return true
	}

	return false
}

// SetIsManuallyDeleted gets a reference to the given NullableBool and assigns it to the IsManuallyDeleted field.
func (o *SnapshotInfo) SetIsManuallyDeleted(v bool) {
	o.IsManuallyDeleted.Set(&v)
}
// SetIsManuallyDeletedNil sets the value for IsManuallyDeleted to be an explicit nil
func (o *SnapshotInfo) SetIsManuallyDeletedNil() {
	o.IsManuallyDeleted.Set(nil)
}

// UnsetIsManuallyDeleted ensures that no value is present for IsManuallyDeleted, not even an explicit nil
func (o *SnapshotInfo) UnsetIsManuallyDeleted() {
	o.IsManuallyDeleted.Unset()
}

// GetExpiryTimeUsecs returns the ExpiryTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnapshotInfo) GetExpiryTimeUsecs() int64 {
	if o == nil || o.ExpiryTimeUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ExpiryTimeUsecs.Get()
}

// GetExpiryTimeUsecsOk returns a tuple with the ExpiryTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotInfo) GetExpiryTimeUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExpiryTimeUsecs.Get(), o.ExpiryTimeUsecs.IsSet()
}

// HasExpiryTimeUsecs returns a boolean if a field has been set.
func (o *SnapshotInfo) HasExpiryTimeUsecs() bool {
	if o != nil && o.ExpiryTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetExpiryTimeUsecs gets a reference to the given NullableInt64 and assigns it to the ExpiryTimeUsecs field.
func (o *SnapshotInfo) SetExpiryTimeUsecs(v int64) {
	o.ExpiryTimeUsecs.Set(&v)
}
// SetExpiryTimeUsecsNil sets the value for ExpiryTimeUsecs to be an explicit nil
func (o *SnapshotInfo) SetExpiryTimeUsecsNil() {
	o.ExpiryTimeUsecs.Set(nil)
}

// UnsetExpiryTimeUsecs ensures that no value is present for ExpiryTimeUsecs, not even an explicit nil
func (o *SnapshotInfo) UnsetExpiryTimeUsecs() {
	o.ExpiryTimeUsecs.Unset()
}

// GetTotalFileCount returns the TotalFileCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnapshotInfo) GetTotalFileCount() int64 {
	if o == nil || o.TotalFileCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.TotalFileCount.Get()
}

// GetTotalFileCountOk returns a tuple with the TotalFileCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotInfo) GetTotalFileCountOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TotalFileCount.Get(), o.TotalFileCount.IsSet()
}

// HasTotalFileCount returns a boolean if a field has been set.
func (o *SnapshotInfo) HasTotalFileCount() bool {
	if o != nil && o.TotalFileCount.IsSet() {
		return true
	}

	return false
}

// SetTotalFileCount gets a reference to the given NullableInt64 and assigns it to the TotalFileCount field.
func (o *SnapshotInfo) SetTotalFileCount(v int64) {
	o.TotalFileCount.Set(&v)
}
// SetTotalFileCountNil sets the value for TotalFileCount to be an explicit nil
func (o *SnapshotInfo) SetTotalFileCountNil() {
	o.TotalFileCount.Set(nil)
}

// UnsetTotalFileCount ensures that no value is present for TotalFileCount, not even an explicit nil
func (o *SnapshotInfo) UnsetTotalFileCount() {
	o.TotalFileCount.Unset()
}

// GetBackupFileCount returns the BackupFileCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnapshotInfo) GetBackupFileCount() int64 {
	if o == nil || o.BackupFileCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.BackupFileCount.Get()
}

// GetBackupFileCountOk returns a tuple with the BackupFileCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotInfo) GetBackupFileCountOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BackupFileCount.Get(), o.BackupFileCount.IsSet()
}

// HasBackupFileCount returns a boolean if a field has been set.
func (o *SnapshotInfo) HasBackupFileCount() bool {
	if o != nil && o.BackupFileCount.IsSet() {
		return true
	}

	return false
}

// SetBackupFileCount gets a reference to the given NullableInt64 and assigns it to the BackupFileCount field.
func (o *SnapshotInfo) SetBackupFileCount(v int64) {
	o.BackupFileCount.Set(&v)
}
// SetBackupFileCountNil sets the value for BackupFileCount to be an explicit nil
func (o *SnapshotInfo) SetBackupFileCountNil() {
	o.BackupFileCount.Set(nil)
}

// UnsetBackupFileCount ensures that no value is present for BackupFileCount, not even an explicit nil
func (o *SnapshotInfo) UnsetBackupFileCount() {
	o.BackupFileCount.Unset()
}

// GetDataLockConstraints returns the DataLockConstraints field value if set, zero value otherwise.
func (o *SnapshotInfo) GetDataLockConstraints() DataLockConstraints {
	if o == nil || o.DataLockConstraints == nil {
		var ret DataLockConstraints
		return ret
	}
	return *o.DataLockConstraints
}

// GetDataLockConstraintsOk returns a tuple with the DataLockConstraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotInfo) GetDataLockConstraintsOk() (*DataLockConstraints, bool) {
	if o == nil || o.DataLockConstraints == nil {
		return nil, false
	}
	return o.DataLockConstraints, true
}

// HasDataLockConstraints returns a boolean if a field has been set.
func (o *SnapshotInfo) HasDataLockConstraints() bool {
	if o != nil && o.DataLockConstraints != nil {
		return true
	}

	return false
}

// SetDataLockConstraints gets a reference to the given DataLockConstraints and assigns it to the DataLockConstraints field.
func (o *SnapshotInfo) SetDataLockConstraints(v DataLockConstraints) {
	o.DataLockConstraints = &v
}

func (o SnapshotInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SnapshotId.IsSet() {
		toSerialize["snapshotId"] = o.SnapshotId.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.StartTimeUsecs.IsSet() {
		toSerialize["startTimeUsecs"] = o.StartTimeUsecs.Get()
	}
	if o.EndTimeUsecs.IsSet() {
		toSerialize["endTimeUsecs"] = o.EndTimeUsecs.Get()
	}
	if o.AdmittedTimeUsecs.IsSet() {
		toSerialize["admittedTimeUsecs"] = o.AdmittedTimeUsecs.Get()
	}
	if o.PermitGrantTimeUsecs.IsSet() {
		toSerialize["permitGrantTimeUsecs"] = o.PermitGrantTimeUsecs.Get()
	}
	if o.QueueDurationUsecs.IsSet() {
		toSerialize["queueDurationUsecs"] = o.QueueDurationUsecs.Get()
	}
	if o.SnapshotCreationTimeUsecs.IsSet() {
		toSerialize["snapshotCreationTimeUsecs"] = o.SnapshotCreationTimeUsecs.Get()
	}
	if o.Stats != nil {
		toSerialize["stats"] = o.Stats
	}
	if o.ProgressTaskId.IsSet() {
		toSerialize["progressTaskId"] = o.ProgressTaskId.Get()
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	if o.IsManuallyDeleted.IsSet() {
		toSerialize["isManuallyDeleted"] = o.IsManuallyDeleted.Get()
	}
	if o.ExpiryTimeUsecs.IsSet() {
		toSerialize["expiryTimeUsecs"] = o.ExpiryTimeUsecs.Get()
	}
	if o.TotalFileCount.IsSet() {
		toSerialize["totalFileCount"] = o.TotalFileCount.Get()
	}
	if o.BackupFileCount.IsSet() {
		toSerialize["backupFileCount"] = o.BackupFileCount.Get()
	}
	if o.DataLockConstraints != nil {
		toSerialize["dataLockConstraints"] = o.DataLockConstraints
	}
	return json.Marshal(toSerialize)
}

type NullableSnapshotInfo struct {
	value *SnapshotInfo
	isSet bool
}

func (v NullableSnapshotInfo) Get() *SnapshotInfo {
	return v.value
}

func (v *NullableSnapshotInfo) Set(val *SnapshotInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotInfo(val *SnapshotInfo) *NullableSnapshotInfo {
	return &NullableSnapshotInfo{value: val, isSet: true}
}

func (v NullableSnapshotInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


