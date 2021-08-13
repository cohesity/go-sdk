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

// RecoverExchangeAppSnapshotParams Specifies the snapshot parameters to recover Exchange databases.
type RecoverExchangeAppSnapshotParams struct {
	// Specifies the snapshot id.
	SnapshotId string `json:"snapshotId"`
	// Specifies the timestamp (in microseconds. from epoch) for recovering to a point-in-time in the past.
	PointInTimeUsecs NullableInt64 `json:"pointInTimeUsecs,omitempty"`
	// Specifies the protection group id of the object snapshot.
	ProtectionGroupId NullableString `json:"protectionGroupId,omitempty"`
	// Specifies the protection group name of the object snapshot.
	ProtectionGroupName NullableString `json:"protectionGroupName,omitempty"`
	// Specifies the time when the snapshot is created in Unix timestamp epoch in microseconds.
	SnapshotCreationTimeUsecs NullableInt64 `json:"snapshotCreationTimeUsecs,omitempty"`
	// Specifies the information about the object for which the snapshot is taken.
	ObjectInfo NullableObjectSummary `json:"objectInfo,omitempty"`
	// Specifies the snapshot target type.
	SnapshotTargetType NullableString `json:"snapshotTargetType,omitempty"`
	// Specifies the ID of the Storage Domain where this snapshot is stored.
	StorageDomainId NullableInt64 `json:"storageDomainId,omitempty"`
	// Specifies the archival target information if the snapshot is an archival snapshot.
	ArchivalTargetInfo NullableArchivalTargetSummaryInfo `json:"archivalTargetInfo,omitempty"`
	// Progress monitor task id for Recovery of VM.
	ProgressTaskId NullableString `json:"progressTaskId,omitempty"`
	// Specifies that user wants to perform standby restore if it is enabled for this object.
	RecoverFromStandby NullableBool `json:"recoverFromStandby,omitempty"`
	// Status of the Recovery. 'Running' indicates that the Recovery is still running. 'Canceled' indicates that the Recovery has been cancelled. 'Canceling' indicates that the Recovery is in the process of being cancelled. 'Failed' indicates that the Recovery has failed. 'Succeeded' indicates that the Recovery has finished successfully. 'SucceededWithWarning' indicates that the Recovery finished successfully, but there were some warning messages.
	Status NullableString `json:"status,omitempty"`
	// Specifies the start time of the Recovery in Unix timestamp epoch in microseconds.
	StartTimeUsecs NullableInt64 `json:"startTimeUsecs,omitempty"`
	// Specifies the end time of the Recovery in Unix timestamp epoch in microseconds. This field will be populated only after Recovery is finished.
	EndTimeUsecs NullableInt64 `json:"endTimeUsecs,omitempty"`
	// Specify error messages about the object.
	Messages []string `json:"messages,omitempty"`
	// Specifies the list of params for all the databases which have to be recovered.
	AppObjects []ExchangeRecoverDatabaseParams `json:"appObjects,omitempty"`
}

// NewRecoverExchangeAppSnapshotParams instantiates a new RecoverExchangeAppSnapshotParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverExchangeAppSnapshotParams(snapshotId string) *RecoverExchangeAppSnapshotParams {
	this := RecoverExchangeAppSnapshotParams{}
	this.SnapshotId = snapshotId
	return &this
}

// NewRecoverExchangeAppSnapshotParamsWithDefaults instantiates a new RecoverExchangeAppSnapshotParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverExchangeAppSnapshotParamsWithDefaults() *RecoverExchangeAppSnapshotParams {
	this := RecoverExchangeAppSnapshotParams{}
	return &this
}

// GetSnapshotId returns the SnapshotId field value
func (o *RecoverExchangeAppSnapshotParams) GetSnapshotId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value
// and a boolean to check if the value has been set.
func (o *RecoverExchangeAppSnapshotParams) GetSnapshotIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SnapshotId, true
}

// SetSnapshotId sets field value
func (o *RecoverExchangeAppSnapshotParams) SetSnapshotId(v string) {
	o.SnapshotId = v
}

// GetPointInTimeUsecs returns the PointInTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverExchangeAppSnapshotParams) GetPointInTimeUsecs() int64 {
	if o == nil || o.PointInTimeUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.PointInTimeUsecs.Get()
}

// GetPointInTimeUsecsOk returns a tuple with the PointInTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverExchangeAppSnapshotParams) GetPointInTimeUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PointInTimeUsecs.Get(), o.PointInTimeUsecs.IsSet()
}

// HasPointInTimeUsecs returns a boolean if a field has been set.
func (o *RecoverExchangeAppSnapshotParams) HasPointInTimeUsecs() bool {
	if o != nil && o.PointInTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetPointInTimeUsecs gets a reference to the given NullableInt64 and assigns it to the PointInTimeUsecs field.
func (o *RecoverExchangeAppSnapshotParams) SetPointInTimeUsecs(v int64) {
	o.PointInTimeUsecs.Set(&v)
}
// SetPointInTimeUsecsNil sets the value for PointInTimeUsecs to be an explicit nil
func (o *RecoverExchangeAppSnapshotParams) SetPointInTimeUsecsNil() {
	o.PointInTimeUsecs.Set(nil)
}

// UnsetPointInTimeUsecs ensures that no value is present for PointInTimeUsecs, not even an explicit nil
func (o *RecoverExchangeAppSnapshotParams) UnsetPointInTimeUsecs() {
	o.PointInTimeUsecs.Unset()
}

// GetProtectionGroupId returns the ProtectionGroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverExchangeAppSnapshotParams) GetProtectionGroupId() string {
	if o == nil || o.ProtectionGroupId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProtectionGroupId.Get()
}

// GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverExchangeAppSnapshotParams) GetProtectionGroupIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProtectionGroupId.Get(), o.ProtectionGroupId.IsSet()
}

// HasProtectionGroupId returns a boolean if a field has been set.
func (o *RecoverExchangeAppSnapshotParams) HasProtectionGroupId() bool {
	if o != nil && o.ProtectionGroupId.IsSet() {
		return true
	}

	return false
}

// SetProtectionGroupId gets a reference to the given NullableString and assigns it to the ProtectionGroupId field.
func (o *RecoverExchangeAppSnapshotParams) SetProtectionGroupId(v string) {
	o.ProtectionGroupId.Set(&v)
}
// SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil
func (o *RecoverExchangeAppSnapshotParams) SetProtectionGroupIdNil() {
	o.ProtectionGroupId.Set(nil)
}

// UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
func (o *RecoverExchangeAppSnapshotParams) UnsetProtectionGroupId() {
	o.ProtectionGroupId.Unset()
}

// GetProtectionGroupName returns the ProtectionGroupName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverExchangeAppSnapshotParams) GetProtectionGroupName() string {
	if o == nil || o.ProtectionGroupName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProtectionGroupName.Get()
}

// GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverExchangeAppSnapshotParams) GetProtectionGroupNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProtectionGroupName.Get(), o.ProtectionGroupName.IsSet()
}

// HasProtectionGroupName returns a boolean if a field has been set.
func (o *RecoverExchangeAppSnapshotParams) HasProtectionGroupName() bool {
	if o != nil && o.ProtectionGroupName.IsSet() {
		return true
	}

	return false
}

// SetProtectionGroupName gets a reference to the given NullableString and assigns it to the ProtectionGroupName field.
func (o *RecoverExchangeAppSnapshotParams) SetProtectionGroupName(v string) {
	o.ProtectionGroupName.Set(&v)
}
// SetProtectionGroupNameNil sets the value for ProtectionGroupName to be an explicit nil
func (o *RecoverExchangeAppSnapshotParams) SetProtectionGroupNameNil() {
	o.ProtectionGroupName.Set(nil)
}

// UnsetProtectionGroupName ensures that no value is present for ProtectionGroupName, not even an explicit nil
func (o *RecoverExchangeAppSnapshotParams) UnsetProtectionGroupName() {
	o.ProtectionGroupName.Unset()
}

// GetSnapshotCreationTimeUsecs returns the SnapshotCreationTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverExchangeAppSnapshotParams) GetSnapshotCreationTimeUsecs() int64 {
	if o == nil || o.SnapshotCreationTimeUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SnapshotCreationTimeUsecs.Get()
}

// GetSnapshotCreationTimeUsecsOk returns a tuple with the SnapshotCreationTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverExchangeAppSnapshotParams) GetSnapshotCreationTimeUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SnapshotCreationTimeUsecs.Get(), o.SnapshotCreationTimeUsecs.IsSet()
}

// HasSnapshotCreationTimeUsecs returns a boolean if a field has been set.
func (o *RecoverExchangeAppSnapshotParams) HasSnapshotCreationTimeUsecs() bool {
	if o != nil && o.SnapshotCreationTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetSnapshotCreationTimeUsecs gets a reference to the given NullableInt64 and assigns it to the SnapshotCreationTimeUsecs field.
func (o *RecoverExchangeAppSnapshotParams) SetSnapshotCreationTimeUsecs(v int64) {
	o.SnapshotCreationTimeUsecs.Set(&v)
}
// SetSnapshotCreationTimeUsecsNil sets the value for SnapshotCreationTimeUsecs to be an explicit nil
func (o *RecoverExchangeAppSnapshotParams) SetSnapshotCreationTimeUsecsNil() {
	o.SnapshotCreationTimeUsecs.Set(nil)
}

// UnsetSnapshotCreationTimeUsecs ensures that no value is present for SnapshotCreationTimeUsecs, not even an explicit nil
func (o *RecoverExchangeAppSnapshotParams) UnsetSnapshotCreationTimeUsecs() {
	o.SnapshotCreationTimeUsecs.Unset()
}

// GetObjectInfo returns the ObjectInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverExchangeAppSnapshotParams) GetObjectInfo() ObjectSummary {
	if o == nil || o.ObjectInfo.Get() == nil {
		var ret ObjectSummary
		return ret
	}
	return *o.ObjectInfo.Get()
}

// GetObjectInfoOk returns a tuple with the ObjectInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverExchangeAppSnapshotParams) GetObjectInfoOk() (*ObjectSummary, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ObjectInfo.Get(), o.ObjectInfo.IsSet()
}

// HasObjectInfo returns a boolean if a field has been set.
func (o *RecoverExchangeAppSnapshotParams) HasObjectInfo() bool {
	if o != nil && o.ObjectInfo.IsSet() {
		return true
	}

	return false
}

// SetObjectInfo gets a reference to the given NullableObjectSummary and assigns it to the ObjectInfo field.
func (o *RecoverExchangeAppSnapshotParams) SetObjectInfo(v ObjectSummary) {
	o.ObjectInfo.Set(&v)
}
// SetObjectInfoNil sets the value for ObjectInfo to be an explicit nil
func (o *RecoverExchangeAppSnapshotParams) SetObjectInfoNil() {
	o.ObjectInfo.Set(nil)
}

// UnsetObjectInfo ensures that no value is present for ObjectInfo, not even an explicit nil
func (o *RecoverExchangeAppSnapshotParams) UnsetObjectInfo() {
	o.ObjectInfo.Unset()
}

// GetSnapshotTargetType returns the SnapshotTargetType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverExchangeAppSnapshotParams) GetSnapshotTargetType() string {
	if o == nil || o.SnapshotTargetType.Get() == nil {
		var ret string
		return ret
	}
	return *o.SnapshotTargetType.Get()
}

// GetSnapshotTargetTypeOk returns a tuple with the SnapshotTargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverExchangeAppSnapshotParams) GetSnapshotTargetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SnapshotTargetType.Get(), o.SnapshotTargetType.IsSet()
}

// HasSnapshotTargetType returns a boolean if a field has been set.
func (o *RecoverExchangeAppSnapshotParams) HasSnapshotTargetType() bool {
	if o != nil && o.SnapshotTargetType.IsSet() {
		return true
	}

	return false
}

// SetSnapshotTargetType gets a reference to the given NullableString and assigns it to the SnapshotTargetType field.
func (o *RecoverExchangeAppSnapshotParams) SetSnapshotTargetType(v string) {
	o.SnapshotTargetType.Set(&v)
}
// SetSnapshotTargetTypeNil sets the value for SnapshotTargetType to be an explicit nil
func (o *RecoverExchangeAppSnapshotParams) SetSnapshotTargetTypeNil() {
	o.SnapshotTargetType.Set(nil)
}

// UnsetSnapshotTargetType ensures that no value is present for SnapshotTargetType, not even an explicit nil
func (o *RecoverExchangeAppSnapshotParams) UnsetSnapshotTargetType() {
	o.SnapshotTargetType.Unset()
}

// GetStorageDomainId returns the StorageDomainId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverExchangeAppSnapshotParams) GetStorageDomainId() int64 {
	if o == nil || o.StorageDomainId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.StorageDomainId.Get()
}

// GetStorageDomainIdOk returns a tuple with the StorageDomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverExchangeAppSnapshotParams) GetStorageDomainIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StorageDomainId.Get(), o.StorageDomainId.IsSet()
}

// HasStorageDomainId returns a boolean if a field has been set.
func (o *RecoverExchangeAppSnapshotParams) HasStorageDomainId() bool {
	if o != nil && o.StorageDomainId.IsSet() {
		return true
	}

	return false
}

// SetStorageDomainId gets a reference to the given NullableInt64 and assigns it to the StorageDomainId field.
func (o *RecoverExchangeAppSnapshotParams) SetStorageDomainId(v int64) {
	o.StorageDomainId.Set(&v)
}
// SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil
func (o *RecoverExchangeAppSnapshotParams) SetStorageDomainIdNil() {
	o.StorageDomainId.Set(nil)
}

// UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
func (o *RecoverExchangeAppSnapshotParams) UnsetStorageDomainId() {
	o.StorageDomainId.Unset()
}

// GetArchivalTargetInfo returns the ArchivalTargetInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverExchangeAppSnapshotParams) GetArchivalTargetInfo() ArchivalTargetSummaryInfo {
	if o == nil || o.ArchivalTargetInfo.Get() == nil {
		var ret ArchivalTargetSummaryInfo
		return ret
	}
	return *o.ArchivalTargetInfo.Get()
}

// GetArchivalTargetInfoOk returns a tuple with the ArchivalTargetInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverExchangeAppSnapshotParams) GetArchivalTargetInfoOk() (*ArchivalTargetSummaryInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ArchivalTargetInfo.Get(), o.ArchivalTargetInfo.IsSet()
}

// HasArchivalTargetInfo returns a boolean if a field has been set.
func (o *RecoverExchangeAppSnapshotParams) HasArchivalTargetInfo() bool {
	if o != nil && o.ArchivalTargetInfo.IsSet() {
		return true
	}

	return false
}

// SetArchivalTargetInfo gets a reference to the given NullableArchivalTargetSummaryInfo and assigns it to the ArchivalTargetInfo field.
func (o *RecoverExchangeAppSnapshotParams) SetArchivalTargetInfo(v ArchivalTargetSummaryInfo) {
	o.ArchivalTargetInfo.Set(&v)
}
// SetArchivalTargetInfoNil sets the value for ArchivalTargetInfo to be an explicit nil
func (o *RecoverExchangeAppSnapshotParams) SetArchivalTargetInfoNil() {
	o.ArchivalTargetInfo.Set(nil)
}

// UnsetArchivalTargetInfo ensures that no value is present for ArchivalTargetInfo, not even an explicit nil
func (o *RecoverExchangeAppSnapshotParams) UnsetArchivalTargetInfo() {
	o.ArchivalTargetInfo.Unset()
}

// GetProgressTaskId returns the ProgressTaskId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverExchangeAppSnapshotParams) GetProgressTaskId() string {
	if o == nil || o.ProgressTaskId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProgressTaskId.Get()
}

// GetProgressTaskIdOk returns a tuple with the ProgressTaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverExchangeAppSnapshotParams) GetProgressTaskIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProgressTaskId.Get(), o.ProgressTaskId.IsSet()
}

// HasProgressTaskId returns a boolean if a field has been set.
func (o *RecoverExchangeAppSnapshotParams) HasProgressTaskId() bool {
	if o != nil && o.ProgressTaskId.IsSet() {
		return true
	}

	return false
}

// SetProgressTaskId gets a reference to the given NullableString and assigns it to the ProgressTaskId field.
func (o *RecoverExchangeAppSnapshotParams) SetProgressTaskId(v string) {
	o.ProgressTaskId.Set(&v)
}
// SetProgressTaskIdNil sets the value for ProgressTaskId to be an explicit nil
func (o *RecoverExchangeAppSnapshotParams) SetProgressTaskIdNil() {
	o.ProgressTaskId.Set(nil)
}

// UnsetProgressTaskId ensures that no value is present for ProgressTaskId, not even an explicit nil
func (o *RecoverExchangeAppSnapshotParams) UnsetProgressTaskId() {
	o.ProgressTaskId.Unset()
}

// GetRecoverFromStandby returns the RecoverFromStandby field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverExchangeAppSnapshotParams) GetRecoverFromStandby() bool {
	if o == nil || o.RecoverFromStandby.Get() == nil {
		var ret bool
		return ret
	}
	return *o.RecoverFromStandby.Get()
}

// GetRecoverFromStandbyOk returns a tuple with the RecoverFromStandby field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverExchangeAppSnapshotParams) GetRecoverFromStandbyOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RecoverFromStandby.Get(), o.RecoverFromStandby.IsSet()
}

// HasRecoverFromStandby returns a boolean if a field has been set.
func (o *RecoverExchangeAppSnapshotParams) HasRecoverFromStandby() bool {
	if o != nil && o.RecoverFromStandby.IsSet() {
		return true
	}

	return false
}

// SetRecoverFromStandby gets a reference to the given NullableBool and assigns it to the RecoverFromStandby field.
func (o *RecoverExchangeAppSnapshotParams) SetRecoverFromStandby(v bool) {
	o.RecoverFromStandby.Set(&v)
}
// SetRecoverFromStandbyNil sets the value for RecoverFromStandby to be an explicit nil
func (o *RecoverExchangeAppSnapshotParams) SetRecoverFromStandbyNil() {
	o.RecoverFromStandby.Set(nil)
}

// UnsetRecoverFromStandby ensures that no value is present for RecoverFromStandby, not even an explicit nil
func (o *RecoverExchangeAppSnapshotParams) UnsetRecoverFromStandby() {
	o.RecoverFromStandby.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverExchangeAppSnapshotParams) GetStatus() string {
	if o == nil || o.Status.Get() == nil {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverExchangeAppSnapshotParams) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *RecoverExchangeAppSnapshotParams) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *RecoverExchangeAppSnapshotParams) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *RecoverExchangeAppSnapshotParams) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *RecoverExchangeAppSnapshotParams) UnsetStatus() {
	o.Status.Unset()
}

// GetStartTimeUsecs returns the StartTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverExchangeAppSnapshotParams) GetStartTimeUsecs() int64 {
	if o == nil || o.StartTimeUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.StartTimeUsecs.Get()
}

// GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverExchangeAppSnapshotParams) GetStartTimeUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartTimeUsecs.Get(), o.StartTimeUsecs.IsSet()
}

// HasStartTimeUsecs returns a boolean if a field has been set.
func (o *RecoverExchangeAppSnapshotParams) HasStartTimeUsecs() bool {
	if o != nil && o.StartTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetStartTimeUsecs gets a reference to the given NullableInt64 and assigns it to the StartTimeUsecs field.
func (o *RecoverExchangeAppSnapshotParams) SetStartTimeUsecs(v int64) {
	o.StartTimeUsecs.Set(&v)
}
// SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil
func (o *RecoverExchangeAppSnapshotParams) SetStartTimeUsecsNil() {
	o.StartTimeUsecs.Set(nil)
}

// UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
func (o *RecoverExchangeAppSnapshotParams) UnsetStartTimeUsecs() {
	o.StartTimeUsecs.Unset()
}

// GetEndTimeUsecs returns the EndTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverExchangeAppSnapshotParams) GetEndTimeUsecs() int64 {
	if o == nil || o.EndTimeUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.EndTimeUsecs.Get()
}

// GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverExchangeAppSnapshotParams) GetEndTimeUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndTimeUsecs.Get(), o.EndTimeUsecs.IsSet()
}

// HasEndTimeUsecs returns a boolean if a field has been set.
func (o *RecoverExchangeAppSnapshotParams) HasEndTimeUsecs() bool {
	if o != nil && o.EndTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetEndTimeUsecs gets a reference to the given NullableInt64 and assigns it to the EndTimeUsecs field.
func (o *RecoverExchangeAppSnapshotParams) SetEndTimeUsecs(v int64) {
	o.EndTimeUsecs.Set(&v)
}
// SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil
func (o *RecoverExchangeAppSnapshotParams) SetEndTimeUsecsNil() {
	o.EndTimeUsecs.Set(nil)
}

// UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
func (o *RecoverExchangeAppSnapshotParams) UnsetEndTimeUsecs() {
	o.EndTimeUsecs.Unset()
}

// GetMessages returns the Messages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverExchangeAppSnapshotParams) GetMessages() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverExchangeAppSnapshotParams) GetMessagesOk() (*[]string, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return &o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *RecoverExchangeAppSnapshotParams) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []string and assigns it to the Messages field.
func (o *RecoverExchangeAppSnapshotParams) SetMessages(v []string) {
	o.Messages = v
}

// GetAppObjects returns the AppObjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverExchangeAppSnapshotParams) GetAppObjects() []ExchangeRecoverDatabaseParams {
	if o == nil  {
		var ret []ExchangeRecoverDatabaseParams
		return ret
	}
	return o.AppObjects
}

// GetAppObjectsOk returns a tuple with the AppObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverExchangeAppSnapshotParams) GetAppObjectsOk() (*[]ExchangeRecoverDatabaseParams, bool) {
	if o == nil || o.AppObjects == nil {
		return nil, false
	}
	return &o.AppObjects, true
}

// HasAppObjects returns a boolean if a field has been set.
func (o *RecoverExchangeAppSnapshotParams) HasAppObjects() bool {
	if o != nil && o.AppObjects != nil {
		return true
	}

	return false
}

// SetAppObjects gets a reference to the given []ExchangeRecoverDatabaseParams and assigns it to the AppObjects field.
func (o *RecoverExchangeAppSnapshotParams) SetAppObjects(v []ExchangeRecoverDatabaseParams) {
	o.AppObjects = v
}

func (o RecoverExchangeAppSnapshotParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["snapshotId"] = o.SnapshotId
	}
	if o.PointInTimeUsecs.IsSet() {
		toSerialize["pointInTimeUsecs"] = o.PointInTimeUsecs.Get()
	}
	if o.ProtectionGroupId.IsSet() {
		toSerialize["protectionGroupId"] = o.ProtectionGroupId.Get()
	}
	if o.ProtectionGroupName.IsSet() {
		toSerialize["protectionGroupName"] = o.ProtectionGroupName.Get()
	}
	if o.SnapshotCreationTimeUsecs.IsSet() {
		toSerialize["snapshotCreationTimeUsecs"] = o.SnapshotCreationTimeUsecs.Get()
	}
	if o.ObjectInfo.IsSet() {
		toSerialize["objectInfo"] = o.ObjectInfo.Get()
	}
	if o.SnapshotTargetType.IsSet() {
		toSerialize["snapshotTargetType"] = o.SnapshotTargetType.Get()
	}
	if o.StorageDomainId.IsSet() {
		toSerialize["storageDomainId"] = o.StorageDomainId.Get()
	}
	if o.ArchivalTargetInfo.IsSet() {
		toSerialize["archivalTargetInfo"] = o.ArchivalTargetInfo.Get()
	}
	if o.ProgressTaskId.IsSet() {
		toSerialize["progressTaskId"] = o.ProgressTaskId.Get()
	}
	if o.RecoverFromStandby.IsSet() {
		toSerialize["recoverFromStandby"] = o.RecoverFromStandby.Get()
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
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}
	if o.AppObjects != nil {
		toSerialize["appObjects"] = o.AppObjects
	}
	return json.Marshal(toSerialize)
}

type NullableRecoverExchangeAppSnapshotParams struct {
	value *RecoverExchangeAppSnapshotParams
	isSet bool
}

func (v NullableRecoverExchangeAppSnapshotParams) Get() *RecoverExchangeAppSnapshotParams {
	return v.value
}

func (v *NullableRecoverExchangeAppSnapshotParams) Set(val *RecoverExchangeAppSnapshotParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverExchangeAppSnapshotParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverExchangeAppSnapshotParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverExchangeAppSnapshotParams(val *RecoverExchangeAppSnapshotParams) *NullableRecoverExchangeAppSnapshotParams {
	return &NullableRecoverExchangeAppSnapshotParams{value: val, isSet: true}
}

func (v NullableRecoverExchangeAppSnapshotParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverExchangeAppSnapshotParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


