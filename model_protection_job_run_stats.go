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

// ProtectionJobRunStats Specifies statistics about a Protection Job Run. This contains the Job Run level statistics.
type ProtectionJobRunStats struct {
	// Specifies the time the task was unqueued from the queue to start running. This field can be used to determine the following times: initial-wait-time = admittedTimeUsecs - startTimeUsecs run-time = endTimeUsecs - admittedTimeUsecs If the task ends up waiting in other queues, then actual run-time will be smaller than the run-time computed this way. This field is only populated for Backup tasks currently.
	AdmittedTimeUsecs NullableInt64 `json:"admittedTimeUsecs,omitempty"`
	// Specifies the end time of the Protection Run. The end time is specified as a Unix epoch Timestamp (in microseconds).
	EndTimeUsecs NullableInt64 `json:"endTimeUsecs,omitempty"`
	// Specifies the number of application instances backed up by this Run. For example if the environment type is kSQL, this field contains the number of SQL Server instances.
	NumAppInstances NullableInt32 `json:"numAppInstances,omitempty"`
	// Specifies the number of backup tasks that were canceled.
	NumCanceledTasks NullableInt64 `json:"numCanceledTasks,omitempty"`
	// Specifies the number of application objects that were cancelled in this Run.
	NumCancelledAppObjects NullableInt32 `json:"numCancelledAppObjects,omitempty"`
	// Specifies the number of application objects that failed in this Run.
	NumFailedAppObjects NullableInt32 `json:"numFailedAppObjects,omitempty"`
	// Specifies the number of backup tasks that failed.
	NumFailedTasks NullableInt64 `json:"numFailedTasks,omitempty"`
	// Specifies the number of application objects successfully backed up by this Run. For example, if the environment type is kSQL, this number is for all of the SQL server databases.
	NumSuccessfulAppObjects NullableInt32 `json:"numSuccessfulAppObjects,omitempty"`
	// Specifies the number of backup tasks that completed successfully.
	NumSuccessfulTasks NullableInt64 `json:"numSuccessfulTasks,omitempty"`
	// Specifies the start time of the Protection Run. The start time is specified as a Unix epoch Timestamp (in microseconds). This time is when the task is queued to an internal queue where tasks are waiting to run.
	StartTimeUsecs NullableInt64 `json:"startTimeUsecs,omitempty"`
	// Specifies the actual execution time for the protection run to complete the backup task and the copy tasks. This time will not include the time waited in various internal queues. This field is only populated for Backup tasks currently.
	TimeTakenUsecs NullableInt64 `json:"timeTakenUsecs,omitempty"`
	// Specifies the total amount of data read from the source (so far).
	TotalBytesReadFromSource NullableInt64 `json:"totalBytesReadFromSource,omitempty"`
	// Specifies the total amount of data expected to be read from the source.
	TotalBytesToReadFromSource NullableInt64 `json:"totalBytesToReadFromSource,omitempty"`
	// Specifies the size of the source object (such as a VM) protected by this task on the primary storage after the snapshot is taken. The logical size of the data on the source if the data is fully hydrated or expanded and not reduced by change-block tracking, compression and deduplication.
	TotalLogicalBackupSizeBytes NullableInt64 `json:"totalLogicalBackupSizeBytes,omitempty"`
	// Specifies the total amount of physical space used on the Cohesity Cluster to store the protected object after being reduced by change-block tracking, compression and deduplication. For example, if the logical backup size is 1GB, but only 1MB was used on the Cohesity Cluster to store it, this field be equal to 1MB.
	TotalPhysicalBackupSizeBytes NullableInt64 `json:"totalPhysicalBackupSizeBytes,omitempty"`
	// Specifies the size of the source object (such as a VM) protected by this task on the primary storage before the snapshot is taken. The logical size of the data on the source if the data is fully hydrated or expanded and not reduced by change-block tracking, compression and deduplication.
	TotalSourceSizeBytes NullableInt64 `json:"totalSourceSizeBytes,omitempty"`
}

// NewProtectionJobRunStats instantiates a new ProtectionJobRunStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtectionJobRunStats() *ProtectionJobRunStats {
	this := ProtectionJobRunStats{}
	return &this
}

// NewProtectionJobRunStatsWithDefaults instantiates a new ProtectionJobRunStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtectionJobRunStatsWithDefaults() *ProtectionJobRunStats {
	this := ProtectionJobRunStats{}
	return &this
}

// GetAdmittedTimeUsecs returns the AdmittedTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionJobRunStats) GetAdmittedTimeUsecs() int64 {
	if o == nil || o.AdmittedTimeUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AdmittedTimeUsecs.Get()
}

// GetAdmittedTimeUsecsOk returns a tuple with the AdmittedTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionJobRunStats) GetAdmittedTimeUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AdmittedTimeUsecs.Get(), o.AdmittedTimeUsecs.IsSet()
}

// HasAdmittedTimeUsecs returns a boolean if a field has been set.
func (o *ProtectionJobRunStats) HasAdmittedTimeUsecs() bool {
	if o != nil && o.AdmittedTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetAdmittedTimeUsecs gets a reference to the given NullableInt64 and assigns it to the AdmittedTimeUsecs field.
func (o *ProtectionJobRunStats) SetAdmittedTimeUsecs(v int64) {
	o.AdmittedTimeUsecs.Set(&v)
}
// SetAdmittedTimeUsecsNil sets the value for AdmittedTimeUsecs to be an explicit nil
func (o *ProtectionJobRunStats) SetAdmittedTimeUsecsNil() {
	o.AdmittedTimeUsecs.Set(nil)
}

// UnsetAdmittedTimeUsecs ensures that no value is present for AdmittedTimeUsecs, not even an explicit nil
func (o *ProtectionJobRunStats) UnsetAdmittedTimeUsecs() {
	o.AdmittedTimeUsecs.Unset()
}

// GetEndTimeUsecs returns the EndTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionJobRunStats) GetEndTimeUsecs() int64 {
	if o == nil || o.EndTimeUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.EndTimeUsecs.Get()
}

// GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionJobRunStats) GetEndTimeUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndTimeUsecs.Get(), o.EndTimeUsecs.IsSet()
}

// HasEndTimeUsecs returns a boolean if a field has been set.
func (o *ProtectionJobRunStats) HasEndTimeUsecs() bool {
	if o != nil && o.EndTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetEndTimeUsecs gets a reference to the given NullableInt64 and assigns it to the EndTimeUsecs field.
func (o *ProtectionJobRunStats) SetEndTimeUsecs(v int64) {
	o.EndTimeUsecs.Set(&v)
}
// SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil
func (o *ProtectionJobRunStats) SetEndTimeUsecsNil() {
	o.EndTimeUsecs.Set(nil)
}

// UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
func (o *ProtectionJobRunStats) UnsetEndTimeUsecs() {
	o.EndTimeUsecs.Unset()
}

// GetNumAppInstances returns the NumAppInstances field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionJobRunStats) GetNumAppInstances() int32 {
	if o == nil || o.NumAppInstances.Get() == nil {
		var ret int32
		return ret
	}
	return *o.NumAppInstances.Get()
}

// GetNumAppInstancesOk returns a tuple with the NumAppInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionJobRunStats) GetNumAppInstancesOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NumAppInstances.Get(), o.NumAppInstances.IsSet()
}

// HasNumAppInstances returns a boolean if a field has been set.
func (o *ProtectionJobRunStats) HasNumAppInstances() bool {
	if o != nil && o.NumAppInstances.IsSet() {
		return true
	}

	return false
}

// SetNumAppInstances gets a reference to the given NullableInt32 and assigns it to the NumAppInstances field.
func (o *ProtectionJobRunStats) SetNumAppInstances(v int32) {
	o.NumAppInstances.Set(&v)
}
// SetNumAppInstancesNil sets the value for NumAppInstances to be an explicit nil
func (o *ProtectionJobRunStats) SetNumAppInstancesNil() {
	o.NumAppInstances.Set(nil)
}

// UnsetNumAppInstances ensures that no value is present for NumAppInstances, not even an explicit nil
func (o *ProtectionJobRunStats) UnsetNumAppInstances() {
	o.NumAppInstances.Unset()
}

// GetNumCanceledTasks returns the NumCanceledTasks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionJobRunStats) GetNumCanceledTasks() int64 {
	if o == nil || o.NumCanceledTasks.Get() == nil {
		var ret int64
		return ret
	}
	return *o.NumCanceledTasks.Get()
}

// GetNumCanceledTasksOk returns a tuple with the NumCanceledTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionJobRunStats) GetNumCanceledTasksOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NumCanceledTasks.Get(), o.NumCanceledTasks.IsSet()
}

// HasNumCanceledTasks returns a boolean if a field has been set.
func (o *ProtectionJobRunStats) HasNumCanceledTasks() bool {
	if o != nil && o.NumCanceledTasks.IsSet() {
		return true
	}

	return false
}

// SetNumCanceledTasks gets a reference to the given NullableInt64 and assigns it to the NumCanceledTasks field.
func (o *ProtectionJobRunStats) SetNumCanceledTasks(v int64) {
	o.NumCanceledTasks.Set(&v)
}
// SetNumCanceledTasksNil sets the value for NumCanceledTasks to be an explicit nil
func (o *ProtectionJobRunStats) SetNumCanceledTasksNil() {
	o.NumCanceledTasks.Set(nil)
}

// UnsetNumCanceledTasks ensures that no value is present for NumCanceledTasks, not even an explicit nil
func (o *ProtectionJobRunStats) UnsetNumCanceledTasks() {
	o.NumCanceledTasks.Unset()
}

// GetNumCancelledAppObjects returns the NumCancelledAppObjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionJobRunStats) GetNumCancelledAppObjects() int32 {
	if o == nil || o.NumCancelledAppObjects.Get() == nil {
		var ret int32
		return ret
	}
	return *o.NumCancelledAppObjects.Get()
}

// GetNumCancelledAppObjectsOk returns a tuple with the NumCancelledAppObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionJobRunStats) GetNumCancelledAppObjectsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NumCancelledAppObjects.Get(), o.NumCancelledAppObjects.IsSet()
}

// HasNumCancelledAppObjects returns a boolean if a field has been set.
func (o *ProtectionJobRunStats) HasNumCancelledAppObjects() bool {
	if o != nil && o.NumCancelledAppObjects.IsSet() {
		return true
	}

	return false
}

// SetNumCancelledAppObjects gets a reference to the given NullableInt32 and assigns it to the NumCancelledAppObjects field.
func (o *ProtectionJobRunStats) SetNumCancelledAppObjects(v int32) {
	o.NumCancelledAppObjects.Set(&v)
}
// SetNumCancelledAppObjectsNil sets the value for NumCancelledAppObjects to be an explicit nil
func (o *ProtectionJobRunStats) SetNumCancelledAppObjectsNil() {
	o.NumCancelledAppObjects.Set(nil)
}

// UnsetNumCancelledAppObjects ensures that no value is present for NumCancelledAppObjects, not even an explicit nil
func (o *ProtectionJobRunStats) UnsetNumCancelledAppObjects() {
	o.NumCancelledAppObjects.Unset()
}

// GetNumFailedAppObjects returns the NumFailedAppObjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionJobRunStats) GetNumFailedAppObjects() int32 {
	if o == nil || o.NumFailedAppObjects.Get() == nil {
		var ret int32
		return ret
	}
	return *o.NumFailedAppObjects.Get()
}

// GetNumFailedAppObjectsOk returns a tuple with the NumFailedAppObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionJobRunStats) GetNumFailedAppObjectsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NumFailedAppObjects.Get(), o.NumFailedAppObjects.IsSet()
}

// HasNumFailedAppObjects returns a boolean if a field has been set.
func (o *ProtectionJobRunStats) HasNumFailedAppObjects() bool {
	if o != nil && o.NumFailedAppObjects.IsSet() {
		return true
	}

	return false
}

// SetNumFailedAppObjects gets a reference to the given NullableInt32 and assigns it to the NumFailedAppObjects field.
func (o *ProtectionJobRunStats) SetNumFailedAppObjects(v int32) {
	o.NumFailedAppObjects.Set(&v)
}
// SetNumFailedAppObjectsNil sets the value for NumFailedAppObjects to be an explicit nil
func (o *ProtectionJobRunStats) SetNumFailedAppObjectsNil() {
	o.NumFailedAppObjects.Set(nil)
}

// UnsetNumFailedAppObjects ensures that no value is present for NumFailedAppObjects, not even an explicit nil
func (o *ProtectionJobRunStats) UnsetNumFailedAppObjects() {
	o.NumFailedAppObjects.Unset()
}

// GetNumFailedTasks returns the NumFailedTasks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionJobRunStats) GetNumFailedTasks() int64 {
	if o == nil || o.NumFailedTasks.Get() == nil {
		var ret int64
		return ret
	}
	return *o.NumFailedTasks.Get()
}

// GetNumFailedTasksOk returns a tuple with the NumFailedTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionJobRunStats) GetNumFailedTasksOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NumFailedTasks.Get(), o.NumFailedTasks.IsSet()
}

// HasNumFailedTasks returns a boolean if a field has been set.
func (o *ProtectionJobRunStats) HasNumFailedTasks() bool {
	if o != nil && o.NumFailedTasks.IsSet() {
		return true
	}

	return false
}

// SetNumFailedTasks gets a reference to the given NullableInt64 and assigns it to the NumFailedTasks field.
func (o *ProtectionJobRunStats) SetNumFailedTasks(v int64) {
	o.NumFailedTasks.Set(&v)
}
// SetNumFailedTasksNil sets the value for NumFailedTasks to be an explicit nil
func (o *ProtectionJobRunStats) SetNumFailedTasksNil() {
	o.NumFailedTasks.Set(nil)
}

// UnsetNumFailedTasks ensures that no value is present for NumFailedTasks, not even an explicit nil
func (o *ProtectionJobRunStats) UnsetNumFailedTasks() {
	o.NumFailedTasks.Unset()
}

// GetNumSuccessfulAppObjects returns the NumSuccessfulAppObjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionJobRunStats) GetNumSuccessfulAppObjects() int32 {
	if o == nil || o.NumSuccessfulAppObjects.Get() == nil {
		var ret int32
		return ret
	}
	return *o.NumSuccessfulAppObjects.Get()
}

// GetNumSuccessfulAppObjectsOk returns a tuple with the NumSuccessfulAppObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionJobRunStats) GetNumSuccessfulAppObjectsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NumSuccessfulAppObjects.Get(), o.NumSuccessfulAppObjects.IsSet()
}

// HasNumSuccessfulAppObjects returns a boolean if a field has been set.
func (o *ProtectionJobRunStats) HasNumSuccessfulAppObjects() bool {
	if o != nil && o.NumSuccessfulAppObjects.IsSet() {
		return true
	}

	return false
}

// SetNumSuccessfulAppObjects gets a reference to the given NullableInt32 and assigns it to the NumSuccessfulAppObjects field.
func (o *ProtectionJobRunStats) SetNumSuccessfulAppObjects(v int32) {
	o.NumSuccessfulAppObjects.Set(&v)
}
// SetNumSuccessfulAppObjectsNil sets the value for NumSuccessfulAppObjects to be an explicit nil
func (o *ProtectionJobRunStats) SetNumSuccessfulAppObjectsNil() {
	o.NumSuccessfulAppObjects.Set(nil)
}

// UnsetNumSuccessfulAppObjects ensures that no value is present for NumSuccessfulAppObjects, not even an explicit nil
func (o *ProtectionJobRunStats) UnsetNumSuccessfulAppObjects() {
	o.NumSuccessfulAppObjects.Unset()
}

// GetNumSuccessfulTasks returns the NumSuccessfulTasks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionJobRunStats) GetNumSuccessfulTasks() int64 {
	if o == nil || o.NumSuccessfulTasks.Get() == nil {
		var ret int64
		return ret
	}
	return *o.NumSuccessfulTasks.Get()
}

// GetNumSuccessfulTasksOk returns a tuple with the NumSuccessfulTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionJobRunStats) GetNumSuccessfulTasksOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NumSuccessfulTasks.Get(), o.NumSuccessfulTasks.IsSet()
}

// HasNumSuccessfulTasks returns a boolean if a field has been set.
func (o *ProtectionJobRunStats) HasNumSuccessfulTasks() bool {
	if o != nil && o.NumSuccessfulTasks.IsSet() {
		return true
	}

	return false
}

// SetNumSuccessfulTasks gets a reference to the given NullableInt64 and assigns it to the NumSuccessfulTasks field.
func (o *ProtectionJobRunStats) SetNumSuccessfulTasks(v int64) {
	o.NumSuccessfulTasks.Set(&v)
}
// SetNumSuccessfulTasksNil sets the value for NumSuccessfulTasks to be an explicit nil
func (o *ProtectionJobRunStats) SetNumSuccessfulTasksNil() {
	o.NumSuccessfulTasks.Set(nil)
}

// UnsetNumSuccessfulTasks ensures that no value is present for NumSuccessfulTasks, not even an explicit nil
func (o *ProtectionJobRunStats) UnsetNumSuccessfulTasks() {
	o.NumSuccessfulTasks.Unset()
}

// GetStartTimeUsecs returns the StartTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionJobRunStats) GetStartTimeUsecs() int64 {
	if o == nil || o.StartTimeUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.StartTimeUsecs.Get()
}

// GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionJobRunStats) GetStartTimeUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartTimeUsecs.Get(), o.StartTimeUsecs.IsSet()
}

// HasStartTimeUsecs returns a boolean if a field has been set.
func (o *ProtectionJobRunStats) HasStartTimeUsecs() bool {
	if o != nil && o.StartTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetStartTimeUsecs gets a reference to the given NullableInt64 and assigns it to the StartTimeUsecs field.
func (o *ProtectionJobRunStats) SetStartTimeUsecs(v int64) {
	o.StartTimeUsecs.Set(&v)
}
// SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil
func (o *ProtectionJobRunStats) SetStartTimeUsecsNil() {
	o.StartTimeUsecs.Set(nil)
}

// UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
func (o *ProtectionJobRunStats) UnsetStartTimeUsecs() {
	o.StartTimeUsecs.Unset()
}

// GetTimeTakenUsecs returns the TimeTakenUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionJobRunStats) GetTimeTakenUsecs() int64 {
	if o == nil || o.TimeTakenUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.TimeTakenUsecs.Get()
}

// GetTimeTakenUsecsOk returns a tuple with the TimeTakenUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionJobRunStats) GetTimeTakenUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TimeTakenUsecs.Get(), o.TimeTakenUsecs.IsSet()
}

// HasTimeTakenUsecs returns a boolean if a field has been set.
func (o *ProtectionJobRunStats) HasTimeTakenUsecs() bool {
	if o != nil && o.TimeTakenUsecs.IsSet() {
		return true
	}

	return false
}

// SetTimeTakenUsecs gets a reference to the given NullableInt64 and assigns it to the TimeTakenUsecs field.
func (o *ProtectionJobRunStats) SetTimeTakenUsecs(v int64) {
	o.TimeTakenUsecs.Set(&v)
}
// SetTimeTakenUsecsNil sets the value for TimeTakenUsecs to be an explicit nil
func (o *ProtectionJobRunStats) SetTimeTakenUsecsNil() {
	o.TimeTakenUsecs.Set(nil)
}

// UnsetTimeTakenUsecs ensures that no value is present for TimeTakenUsecs, not even an explicit nil
func (o *ProtectionJobRunStats) UnsetTimeTakenUsecs() {
	o.TimeTakenUsecs.Unset()
}

// GetTotalBytesReadFromSource returns the TotalBytesReadFromSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionJobRunStats) GetTotalBytesReadFromSource() int64 {
	if o == nil || o.TotalBytesReadFromSource.Get() == nil {
		var ret int64
		return ret
	}
	return *o.TotalBytesReadFromSource.Get()
}

// GetTotalBytesReadFromSourceOk returns a tuple with the TotalBytesReadFromSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionJobRunStats) GetTotalBytesReadFromSourceOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TotalBytesReadFromSource.Get(), o.TotalBytesReadFromSource.IsSet()
}

// HasTotalBytesReadFromSource returns a boolean if a field has been set.
func (o *ProtectionJobRunStats) HasTotalBytesReadFromSource() bool {
	if o != nil && o.TotalBytesReadFromSource.IsSet() {
		return true
	}

	return false
}

// SetTotalBytesReadFromSource gets a reference to the given NullableInt64 and assigns it to the TotalBytesReadFromSource field.
func (o *ProtectionJobRunStats) SetTotalBytesReadFromSource(v int64) {
	o.TotalBytesReadFromSource.Set(&v)
}
// SetTotalBytesReadFromSourceNil sets the value for TotalBytesReadFromSource to be an explicit nil
func (o *ProtectionJobRunStats) SetTotalBytesReadFromSourceNil() {
	o.TotalBytesReadFromSource.Set(nil)
}

// UnsetTotalBytesReadFromSource ensures that no value is present for TotalBytesReadFromSource, not even an explicit nil
func (o *ProtectionJobRunStats) UnsetTotalBytesReadFromSource() {
	o.TotalBytesReadFromSource.Unset()
}

// GetTotalBytesToReadFromSource returns the TotalBytesToReadFromSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionJobRunStats) GetTotalBytesToReadFromSource() int64 {
	if o == nil || o.TotalBytesToReadFromSource.Get() == nil {
		var ret int64
		return ret
	}
	return *o.TotalBytesToReadFromSource.Get()
}

// GetTotalBytesToReadFromSourceOk returns a tuple with the TotalBytesToReadFromSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionJobRunStats) GetTotalBytesToReadFromSourceOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TotalBytesToReadFromSource.Get(), o.TotalBytesToReadFromSource.IsSet()
}

// HasTotalBytesToReadFromSource returns a boolean if a field has been set.
func (o *ProtectionJobRunStats) HasTotalBytesToReadFromSource() bool {
	if o != nil && o.TotalBytesToReadFromSource.IsSet() {
		return true
	}

	return false
}

// SetTotalBytesToReadFromSource gets a reference to the given NullableInt64 and assigns it to the TotalBytesToReadFromSource field.
func (o *ProtectionJobRunStats) SetTotalBytesToReadFromSource(v int64) {
	o.TotalBytesToReadFromSource.Set(&v)
}
// SetTotalBytesToReadFromSourceNil sets the value for TotalBytesToReadFromSource to be an explicit nil
func (o *ProtectionJobRunStats) SetTotalBytesToReadFromSourceNil() {
	o.TotalBytesToReadFromSource.Set(nil)
}

// UnsetTotalBytesToReadFromSource ensures that no value is present for TotalBytesToReadFromSource, not even an explicit nil
func (o *ProtectionJobRunStats) UnsetTotalBytesToReadFromSource() {
	o.TotalBytesToReadFromSource.Unset()
}

// GetTotalLogicalBackupSizeBytes returns the TotalLogicalBackupSizeBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionJobRunStats) GetTotalLogicalBackupSizeBytes() int64 {
	if o == nil || o.TotalLogicalBackupSizeBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.TotalLogicalBackupSizeBytes.Get()
}

// GetTotalLogicalBackupSizeBytesOk returns a tuple with the TotalLogicalBackupSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionJobRunStats) GetTotalLogicalBackupSizeBytesOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TotalLogicalBackupSizeBytes.Get(), o.TotalLogicalBackupSizeBytes.IsSet()
}

// HasTotalLogicalBackupSizeBytes returns a boolean if a field has been set.
func (o *ProtectionJobRunStats) HasTotalLogicalBackupSizeBytes() bool {
	if o != nil && o.TotalLogicalBackupSizeBytes.IsSet() {
		return true
	}

	return false
}

// SetTotalLogicalBackupSizeBytes gets a reference to the given NullableInt64 and assigns it to the TotalLogicalBackupSizeBytes field.
func (o *ProtectionJobRunStats) SetTotalLogicalBackupSizeBytes(v int64) {
	o.TotalLogicalBackupSizeBytes.Set(&v)
}
// SetTotalLogicalBackupSizeBytesNil sets the value for TotalLogicalBackupSizeBytes to be an explicit nil
func (o *ProtectionJobRunStats) SetTotalLogicalBackupSizeBytesNil() {
	o.TotalLogicalBackupSizeBytes.Set(nil)
}

// UnsetTotalLogicalBackupSizeBytes ensures that no value is present for TotalLogicalBackupSizeBytes, not even an explicit nil
func (o *ProtectionJobRunStats) UnsetTotalLogicalBackupSizeBytes() {
	o.TotalLogicalBackupSizeBytes.Unset()
}

// GetTotalPhysicalBackupSizeBytes returns the TotalPhysicalBackupSizeBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionJobRunStats) GetTotalPhysicalBackupSizeBytes() int64 {
	if o == nil || o.TotalPhysicalBackupSizeBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.TotalPhysicalBackupSizeBytes.Get()
}

// GetTotalPhysicalBackupSizeBytesOk returns a tuple with the TotalPhysicalBackupSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionJobRunStats) GetTotalPhysicalBackupSizeBytesOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TotalPhysicalBackupSizeBytes.Get(), o.TotalPhysicalBackupSizeBytes.IsSet()
}

// HasTotalPhysicalBackupSizeBytes returns a boolean if a field has been set.
func (o *ProtectionJobRunStats) HasTotalPhysicalBackupSizeBytes() bool {
	if o != nil && o.TotalPhysicalBackupSizeBytes.IsSet() {
		return true
	}

	return false
}

// SetTotalPhysicalBackupSizeBytes gets a reference to the given NullableInt64 and assigns it to the TotalPhysicalBackupSizeBytes field.
func (o *ProtectionJobRunStats) SetTotalPhysicalBackupSizeBytes(v int64) {
	o.TotalPhysicalBackupSizeBytes.Set(&v)
}
// SetTotalPhysicalBackupSizeBytesNil sets the value for TotalPhysicalBackupSizeBytes to be an explicit nil
func (o *ProtectionJobRunStats) SetTotalPhysicalBackupSizeBytesNil() {
	o.TotalPhysicalBackupSizeBytes.Set(nil)
}

// UnsetTotalPhysicalBackupSizeBytes ensures that no value is present for TotalPhysicalBackupSizeBytes, not even an explicit nil
func (o *ProtectionJobRunStats) UnsetTotalPhysicalBackupSizeBytes() {
	o.TotalPhysicalBackupSizeBytes.Unset()
}

// GetTotalSourceSizeBytes returns the TotalSourceSizeBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionJobRunStats) GetTotalSourceSizeBytes() int64 {
	if o == nil || o.TotalSourceSizeBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.TotalSourceSizeBytes.Get()
}

// GetTotalSourceSizeBytesOk returns a tuple with the TotalSourceSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionJobRunStats) GetTotalSourceSizeBytesOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TotalSourceSizeBytes.Get(), o.TotalSourceSizeBytes.IsSet()
}

// HasTotalSourceSizeBytes returns a boolean if a field has been set.
func (o *ProtectionJobRunStats) HasTotalSourceSizeBytes() bool {
	if o != nil && o.TotalSourceSizeBytes.IsSet() {
		return true
	}

	return false
}

// SetTotalSourceSizeBytes gets a reference to the given NullableInt64 and assigns it to the TotalSourceSizeBytes field.
func (o *ProtectionJobRunStats) SetTotalSourceSizeBytes(v int64) {
	o.TotalSourceSizeBytes.Set(&v)
}
// SetTotalSourceSizeBytesNil sets the value for TotalSourceSizeBytes to be an explicit nil
func (o *ProtectionJobRunStats) SetTotalSourceSizeBytesNil() {
	o.TotalSourceSizeBytes.Set(nil)
}

// UnsetTotalSourceSizeBytes ensures that no value is present for TotalSourceSizeBytes, not even an explicit nil
func (o *ProtectionJobRunStats) UnsetTotalSourceSizeBytes() {
	o.TotalSourceSizeBytes.Unset()
}

func (o ProtectionJobRunStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdmittedTimeUsecs.IsSet() {
		toSerialize["admittedTimeUsecs"] = o.AdmittedTimeUsecs.Get()
	}
	if o.EndTimeUsecs.IsSet() {
		toSerialize["endTimeUsecs"] = o.EndTimeUsecs.Get()
	}
	if o.NumAppInstances.IsSet() {
		toSerialize["numAppInstances"] = o.NumAppInstances.Get()
	}
	if o.NumCanceledTasks.IsSet() {
		toSerialize["numCanceledTasks"] = o.NumCanceledTasks.Get()
	}
	if o.NumCancelledAppObjects.IsSet() {
		toSerialize["numCancelledAppObjects"] = o.NumCancelledAppObjects.Get()
	}
	if o.NumFailedAppObjects.IsSet() {
		toSerialize["numFailedAppObjects"] = o.NumFailedAppObjects.Get()
	}
	if o.NumFailedTasks.IsSet() {
		toSerialize["numFailedTasks"] = o.NumFailedTasks.Get()
	}
	if o.NumSuccessfulAppObjects.IsSet() {
		toSerialize["numSuccessfulAppObjects"] = o.NumSuccessfulAppObjects.Get()
	}
	if o.NumSuccessfulTasks.IsSet() {
		toSerialize["numSuccessfulTasks"] = o.NumSuccessfulTasks.Get()
	}
	if o.StartTimeUsecs.IsSet() {
		toSerialize["startTimeUsecs"] = o.StartTimeUsecs.Get()
	}
	if o.TimeTakenUsecs.IsSet() {
		toSerialize["timeTakenUsecs"] = o.TimeTakenUsecs.Get()
	}
	if o.TotalBytesReadFromSource.IsSet() {
		toSerialize["totalBytesReadFromSource"] = o.TotalBytesReadFromSource.Get()
	}
	if o.TotalBytesToReadFromSource.IsSet() {
		toSerialize["totalBytesToReadFromSource"] = o.TotalBytesToReadFromSource.Get()
	}
	if o.TotalLogicalBackupSizeBytes.IsSet() {
		toSerialize["totalLogicalBackupSizeBytes"] = o.TotalLogicalBackupSizeBytes.Get()
	}
	if o.TotalPhysicalBackupSizeBytes.IsSet() {
		toSerialize["totalPhysicalBackupSizeBytes"] = o.TotalPhysicalBackupSizeBytes.Get()
	}
	if o.TotalSourceSizeBytes.IsSet() {
		toSerialize["totalSourceSizeBytes"] = o.TotalSourceSizeBytes.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableProtectionJobRunStats struct {
	value *ProtectionJobRunStats
	isSet bool
}

func (v NullableProtectionJobRunStats) Get() *ProtectionJobRunStats {
	return v.value
}

func (v *NullableProtectionJobRunStats) Set(val *ProtectionJobRunStats) {
	v.value = val
	v.isSet = true
}

func (v NullableProtectionJobRunStats) IsSet() bool {
	return v.isSet
}

func (v *NullableProtectionJobRunStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtectionJobRunStats(val *ProtectionJobRunStats) *NullableProtectionJobRunStats {
	return &NullableProtectionJobRunStats{value: val, isSet: true}
}

func (v NullableProtectionJobRunStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtectionJobRunStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


