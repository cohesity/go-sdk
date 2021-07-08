# ProtectionJobRunStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdmittedTimeUsecs** | Pointer to **NullableInt64** | Specifies the time the task was unqueued from the queue to start running. This field can be used to determine the following times: initial-wait-time &#x3D; admittedTimeUsecs - startTimeUsecs run-time &#x3D; endTimeUsecs - admittedTimeUsecs If the task ends up waiting in other queues, then actual run-time will be smaller than the run-time computed this way. This field is only populated for Backup tasks currently. | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time of the Protection Run. The end time is specified as a Unix epoch Timestamp (in microseconds). | [optional] 
**NumAppInstances** | Pointer to **NullableInt32** | Specifies the number of application instances backed up by this Run. For example if the environment type is kSQL, this field contains the number of SQL Server instances. | [optional] 
**NumCanceledTasks** | Pointer to **NullableInt64** | Specifies the number of backup tasks that were canceled. | [optional] 
**NumCancelledAppObjects** | Pointer to **NullableInt32** | Specifies the number of application objects that were cancelled in this Run. | [optional] 
**NumFailedAppObjects** | Pointer to **NullableInt32** | Specifies the number of application objects that failed in this Run. | [optional] 
**NumFailedTasks** | Pointer to **NullableInt64** | Specifies the number of backup tasks that failed. | [optional] 
**NumSuccessfulAppObjects** | Pointer to **NullableInt32** | Specifies the number of application objects successfully backed up by this Run. For example, if the environment type is kSQL, this number is for all of the SQL server databases. | [optional] 
**NumSuccessfulTasks** | Pointer to **NullableInt64** | Specifies the number of backup tasks that completed successfully. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of the Protection Run. The start time is specified as a Unix epoch Timestamp (in microseconds). This time is when the task is queued to an internal queue where tasks are waiting to run. | [optional] 
**TimeTakenUsecs** | Pointer to **NullableInt64** | Specifies the actual execution time for the protection run to complete the backup task and the copy tasks. This time will not include the time waited in various internal queues. This field is only populated for Backup tasks currently. | [optional] 
**TotalBytesReadFromSource** | Pointer to **NullableInt64** | Specifies the total amount of data read from the source (so far). | [optional] 
**TotalBytesToReadFromSource** | Pointer to **NullableInt64** | Specifies the total amount of data expected to be read from the source. | [optional] 
**TotalLogicalBackupSizeBytes** | Pointer to **NullableInt64** | Specifies the size of the source object (such as a VM) protected by this task on the primary storage after the snapshot is taken. The logical size of the data on the source if the data is fully hydrated or expanded and not reduced by change-block tracking, compression and deduplication. | [optional] 
**TotalPhysicalBackupSizeBytes** | Pointer to **NullableInt64** | Specifies the total amount of physical space used on the Cohesity Cluster to store the protected object after being reduced by change-block tracking, compression and deduplication. For example, if the logical backup size is 1GB, but only 1MB was used on the Cohesity Cluster to store it, this field be equal to 1MB. | [optional] 
**TotalSourceSizeBytes** | Pointer to **NullableInt64** | Specifies the size of the source object (such as a VM) protected by this task on the primary storage before the snapshot is taken. The logical size of the data on the source if the data is fully hydrated or expanded and not reduced by change-block tracking, compression and deduplication. | [optional] 

## Methods

### NewProtectionJobRunStats

`func NewProtectionJobRunStats() *ProtectionJobRunStats`

NewProtectionJobRunStats instantiates a new ProtectionJobRunStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionJobRunStatsWithDefaults

`func NewProtectionJobRunStatsWithDefaults() *ProtectionJobRunStats`

NewProtectionJobRunStatsWithDefaults instantiates a new ProtectionJobRunStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmittedTimeUsecs

`func (o *ProtectionJobRunStats) GetAdmittedTimeUsecs() int64`

GetAdmittedTimeUsecs returns the AdmittedTimeUsecs field if non-nil, zero value otherwise.

### GetAdmittedTimeUsecsOk

`func (o *ProtectionJobRunStats) GetAdmittedTimeUsecsOk() (*int64, bool)`

GetAdmittedTimeUsecsOk returns a tuple with the AdmittedTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmittedTimeUsecs

`func (o *ProtectionJobRunStats) SetAdmittedTimeUsecs(v int64)`

SetAdmittedTimeUsecs sets AdmittedTimeUsecs field to given value.

### HasAdmittedTimeUsecs

`func (o *ProtectionJobRunStats) HasAdmittedTimeUsecs() bool`

HasAdmittedTimeUsecs returns a boolean if a field has been set.

### SetAdmittedTimeUsecsNil

`func (o *ProtectionJobRunStats) SetAdmittedTimeUsecsNil(b bool)`

 SetAdmittedTimeUsecsNil sets the value for AdmittedTimeUsecs to be an explicit nil

### UnsetAdmittedTimeUsecs
`func (o *ProtectionJobRunStats) UnsetAdmittedTimeUsecs()`

UnsetAdmittedTimeUsecs ensures that no value is present for AdmittedTimeUsecs, not even an explicit nil
### GetEndTimeUsecs

`func (o *ProtectionJobRunStats) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *ProtectionJobRunStats) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *ProtectionJobRunStats) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *ProtectionJobRunStats) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *ProtectionJobRunStats) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *ProtectionJobRunStats) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetNumAppInstances

`func (o *ProtectionJobRunStats) GetNumAppInstances() int32`

GetNumAppInstances returns the NumAppInstances field if non-nil, zero value otherwise.

### GetNumAppInstancesOk

`func (o *ProtectionJobRunStats) GetNumAppInstancesOk() (*int32, bool)`

GetNumAppInstancesOk returns a tuple with the NumAppInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAppInstances

`func (o *ProtectionJobRunStats) SetNumAppInstances(v int32)`

SetNumAppInstances sets NumAppInstances field to given value.

### HasNumAppInstances

`func (o *ProtectionJobRunStats) HasNumAppInstances() bool`

HasNumAppInstances returns a boolean if a field has been set.

### SetNumAppInstancesNil

`func (o *ProtectionJobRunStats) SetNumAppInstancesNil(b bool)`

 SetNumAppInstancesNil sets the value for NumAppInstances to be an explicit nil

### UnsetNumAppInstances
`func (o *ProtectionJobRunStats) UnsetNumAppInstances()`

UnsetNumAppInstances ensures that no value is present for NumAppInstances, not even an explicit nil
### GetNumCanceledTasks

`func (o *ProtectionJobRunStats) GetNumCanceledTasks() int64`

GetNumCanceledTasks returns the NumCanceledTasks field if non-nil, zero value otherwise.

### GetNumCanceledTasksOk

`func (o *ProtectionJobRunStats) GetNumCanceledTasksOk() (*int64, bool)`

GetNumCanceledTasksOk returns a tuple with the NumCanceledTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCanceledTasks

`func (o *ProtectionJobRunStats) SetNumCanceledTasks(v int64)`

SetNumCanceledTasks sets NumCanceledTasks field to given value.

### HasNumCanceledTasks

`func (o *ProtectionJobRunStats) HasNumCanceledTasks() bool`

HasNumCanceledTasks returns a boolean if a field has been set.

### SetNumCanceledTasksNil

`func (o *ProtectionJobRunStats) SetNumCanceledTasksNil(b bool)`

 SetNumCanceledTasksNil sets the value for NumCanceledTasks to be an explicit nil

### UnsetNumCanceledTasks
`func (o *ProtectionJobRunStats) UnsetNumCanceledTasks()`

UnsetNumCanceledTasks ensures that no value is present for NumCanceledTasks, not even an explicit nil
### GetNumCancelledAppObjects

`func (o *ProtectionJobRunStats) GetNumCancelledAppObjects() int32`

GetNumCancelledAppObjects returns the NumCancelledAppObjects field if non-nil, zero value otherwise.

### GetNumCancelledAppObjectsOk

`func (o *ProtectionJobRunStats) GetNumCancelledAppObjectsOk() (*int32, bool)`

GetNumCancelledAppObjectsOk returns a tuple with the NumCancelledAppObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCancelledAppObjects

`func (o *ProtectionJobRunStats) SetNumCancelledAppObjects(v int32)`

SetNumCancelledAppObjects sets NumCancelledAppObjects field to given value.

### HasNumCancelledAppObjects

`func (o *ProtectionJobRunStats) HasNumCancelledAppObjects() bool`

HasNumCancelledAppObjects returns a boolean if a field has been set.

### SetNumCancelledAppObjectsNil

`func (o *ProtectionJobRunStats) SetNumCancelledAppObjectsNil(b bool)`

 SetNumCancelledAppObjectsNil sets the value for NumCancelledAppObjects to be an explicit nil

### UnsetNumCancelledAppObjects
`func (o *ProtectionJobRunStats) UnsetNumCancelledAppObjects()`

UnsetNumCancelledAppObjects ensures that no value is present for NumCancelledAppObjects, not even an explicit nil
### GetNumFailedAppObjects

`func (o *ProtectionJobRunStats) GetNumFailedAppObjects() int32`

GetNumFailedAppObjects returns the NumFailedAppObjects field if non-nil, zero value otherwise.

### GetNumFailedAppObjectsOk

`func (o *ProtectionJobRunStats) GetNumFailedAppObjectsOk() (*int32, bool)`

GetNumFailedAppObjectsOk returns a tuple with the NumFailedAppObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFailedAppObjects

`func (o *ProtectionJobRunStats) SetNumFailedAppObjects(v int32)`

SetNumFailedAppObjects sets NumFailedAppObjects field to given value.

### HasNumFailedAppObjects

`func (o *ProtectionJobRunStats) HasNumFailedAppObjects() bool`

HasNumFailedAppObjects returns a boolean if a field has been set.

### SetNumFailedAppObjectsNil

`func (o *ProtectionJobRunStats) SetNumFailedAppObjectsNil(b bool)`

 SetNumFailedAppObjectsNil sets the value for NumFailedAppObjects to be an explicit nil

### UnsetNumFailedAppObjects
`func (o *ProtectionJobRunStats) UnsetNumFailedAppObjects()`

UnsetNumFailedAppObjects ensures that no value is present for NumFailedAppObjects, not even an explicit nil
### GetNumFailedTasks

`func (o *ProtectionJobRunStats) GetNumFailedTasks() int64`

GetNumFailedTasks returns the NumFailedTasks field if non-nil, zero value otherwise.

### GetNumFailedTasksOk

`func (o *ProtectionJobRunStats) GetNumFailedTasksOk() (*int64, bool)`

GetNumFailedTasksOk returns a tuple with the NumFailedTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFailedTasks

`func (o *ProtectionJobRunStats) SetNumFailedTasks(v int64)`

SetNumFailedTasks sets NumFailedTasks field to given value.

### HasNumFailedTasks

`func (o *ProtectionJobRunStats) HasNumFailedTasks() bool`

HasNumFailedTasks returns a boolean if a field has been set.

### SetNumFailedTasksNil

`func (o *ProtectionJobRunStats) SetNumFailedTasksNil(b bool)`

 SetNumFailedTasksNil sets the value for NumFailedTasks to be an explicit nil

### UnsetNumFailedTasks
`func (o *ProtectionJobRunStats) UnsetNumFailedTasks()`

UnsetNumFailedTasks ensures that no value is present for NumFailedTasks, not even an explicit nil
### GetNumSuccessfulAppObjects

`func (o *ProtectionJobRunStats) GetNumSuccessfulAppObjects() int32`

GetNumSuccessfulAppObjects returns the NumSuccessfulAppObjects field if non-nil, zero value otherwise.

### GetNumSuccessfulAppObjectsOk

`func (o *ProtectionJobRunStats) GetNumSuccessfulAppObjectsOk() (*int32, bool)`

GetNumSuccessfulAppObjectsOk returns a tuple with the NumSuccessfulAppObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSuccessfulAppObjects

`func (o *ProtectionJobRunStats) SetNumSuccessfulAppObjects(v int32)`

SetNumSuccessfulAppObjects sets NumSuccessfulAppObjects field to given value.

### HasNumSuccessfulAppObjects

`func (o *ProtectionJobRunStats) HasNumSuccessfulAppObjects() bool`

HasNumSuccessfulAppObjects returns a boolean if a field has been set.

### SetNumSuccessfulAppObjectsNil

`func (o *ProtectionJobRunStats) SetNumSuccessfulAppObjectsNil(b bool)`

 SetNumSuccessfulAppObjectsNil sets the value for NumSuccessfulAppObjects to be an explicit nil

### UnsetNumSuccessfulAppObjects
`func (o *ProtectionJobRunStats) UnsetNumSuccessfulAppObjects()`

UnsetNumSuccessfulAppObjects ensures that no value is present for NumSuccessfulAppObjects, not even an explicit nil
### GetNumSuccessfulTasks

`func (o *ProtectionJobRunStats) GetNumSuccessfulTasks() int64`

GetNumSuccessfulTasks returns the NumSuccessfulTasks field if non-nil, zero value otherwise.

### GetNumSuccessfulTasksOk

`func (o *ProtectionJobRunStats) GetNumSuccessfulTasksOk() (*int64, bool)`

GetNumSuccessfulTasksOk returns a tuple with the NumSuccessfulTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSuccessfulTasks

`func (o *ProtectionJobRunStats) SetNumSuccessfulTasks(v int64)`

SetNumSuccessfulTasks sets NumSuccessfulTasks field to given value.

### HasNumSuccessfulTasks

`func (o *ProtectionJobRunStats) HasNumSuccessfulTasks() bool`

HasNumSuccessfulTasks returns a boolean if a field has been set.

### SetNumSuccessfulTasksNil

`func (o *ProtectionJobRunStats) SetNumSuccessfulTasksNil(b bool)`

 SetNumSuccessfulTasksNil sets the value for NumSuccessfulTasks to be an explicit nil

### UnsetNumSuccessfulTasks
`func (o *ProtectionJobRunStats) UnsetNumSuccessfulTasks()`

UnsetNumSuccessfulTasks ensures that no value is present for NumSuccessfulTasks, not even an explicit nil
### GetStartTimeUsecs

`func (o *ProtectionJobRunStats) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *ProtectionJobRunStats) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *ProtectionJobRunStats) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *ProtectionJobRunStats) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *ProtectionJobRunStats) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *ProtectionJobRunStats) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetTimeTakenUsecs

`func (o *ProtectionJobRunStats) GetTimeTakenUsecs() int64`

GetTimeTakenUsecs returns the TimeTakenUsecs field if non-nil, zero value otherwise.

### GetTimeTakenUsecsOk

`func (o *ProtectionJobRunStats) GetTimeTakenUsecsOk() (*int64, bool)`

GetTimeTakenUsecsOk returns a tuple with the TimeTakenUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeTakenUsecs

`func (o *ProtectionJobRunStats) SetTimeTakenUsecs(v int64)`

SetTimeTakenUsecs sets TimeTakenUsecs field to given value.

### HasTimeTakenUsecs

`func (o *ProtectionJobRunStats) HasTimeTakenUsecs() bool`

HasTimeTakenUsecs returns a boolean if a field has been set.

### SetTimeTakenUsecsNil

`func (o *ProtectionJobRunStats) SetTimeTakenUsecsNil(b bool)`

 SetTimeTakenUsecsNil sets the value for TimeTakenUsecs to be an explicit nil

### UnsetTimeTakenUsecs
`func (o *ProtectionJobRunStats) UnsetTimeTakenUsecs()`

UnsetTimeTakenUsecs ensures that no value is present for TimeTakenUsecs, not even an explicit nil
### GetTotalBytesReadFromSource

`func (o *ProtectionJobRunStats) GetTotalBytesReadFromSource() int64`

GetTotalBytesReadFromSource returns the TotalBytesReadFromSource field if non-nil, zero value otherwise.

### GetTotalBytesReadFromSourceOk

`func (o *ProtectionJobRunStats) GetTotalBytesReadFromSourceOk() (*int64, bool)`

GetTotalBytesReadFromSourceOk returns a tuple with the TotalBytesReadFromSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytesReadFromSource

`func (o *ProtectionJobRunStats) SetTotalBytesReadFromSource(v int64)`

SetTotalBytesReadFromSource sets TotalBytesReadFromSource field to given value.

### HasTotalBytesReadFromSource

`func (o *ProtectionJobRunStats) HasTotalBytesReadFromSource() bool`

HasTotalBytesReadFromSource returns a boolean if a field has been set.

### SetTotalBytesReadFromSourceNil

`func (o *ProtectionJobRunStats) SetTotalBytesReadFromSourceNil(b bool)`

 SetTotalBytesReadFromSourceNil sets the value for TotalBytesReadFromSource to be an explicit nil

### UnsetTotalBytesReadFromSource
`func (o *ProtectionJobRunStats) UnsetTotalBytesReadFromSource()`

UnsetTotalBytesReadFromSource ensures that no value is present for TotalBytesReadFromSource, not even an explicit nil
### GetTotalBytesToReadFromSource

`func (o *ProtectionJobRunStats) GetTotalBytesToReadFromSource() int64`

GetTotalBytesToReadFromSource returns the TotalBytesToReadFromSource field if non-nil, zero value otherwise.

### GetTotalBytesToReadFromSourceOk

`func (o *ProtectionJobRunStats) GetTotalBytesToReadFromSourceOk() (*int64, bool)`

GetTotalBytesToReadFromSourceOk returns a tuple with the TotalBytesToReadFromSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytesToReadFromSource

`func (o *ProtectionJobRunStats) SetTotalBytesToReadFromSource(v int64)`

SetTotalBytesToReadFromSource sets TotalBytesToReadFromSource field to given value.

### HasTotalBytesToReadFromSource

`func (o *ProtectionJobRunStats) HasTotalBytesToReadFromSource() bool`

HasTotalBytesToReadFromSource returns a boolean if a field has been set.

### SetTotalBytesToReadFromSourceNil

`func (o *ProtectionJobRunStats) SetTotalBytesToReadFromSourceNil(b bool)`

 SetTotalBytesToReadFromSourceNil sets the value for TotalBytesToReadFromSource to be an explicit nil

### UnsetTotalBytesToReadFromSource
`func (o *ProtectionJobRunStats) UnsetTotalBytesToReadFromSource()`

UnsetTotalBytesToReadFromSource ensures that no value is present for TotalBytesToReadFromSource, not even an explicit nil
### GetTotalLogicalBackupSizeBytes

`func (o *ProtectionJobRunStats) GetTotalLogicalBackupSizeBytes() int64`

GetTotalLogicalBackupSizeBytes returns the TotalLogicalBackupSizeBytes field if non-nil, zero value otherwise.

### GetTotalLogicalBackupSizeBytesOk

`func (o *ProtectionJobRunStats) GetTotalLogicalBackupSizeBytesOk() (*int64, bool)`

GetTotalLogicalBackupSizeBytesOk returns a tuple with the TotalLogicalBackupSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLogicalBackupSizeBytes

`func (o *ProtectionJobRunStats) SetTotalLogicalBackupSizeBytes(v int64)`

SetTotalLogicalBackupSizeBytes sets TotalLogicalBackupSizeBytes field to given value.

### HasTotalLogicalBackupSizeBytes

`func (o *ProtectionJobRunStats) HasTotalLogicalBackupSizeBytes() bool`

HasTotalLogicalBackupSizeBytes returns a boolean if a field has been set.

### SetTotalLogicalBackupSizeBytesNil

`func (o *ProtectionJobRunStats) SetTotalLogicalBackupSizeBytesNil(b bool)`

 SetTotalLogicalBackupSizeBytesNil sets the value for TotalLogicalBackupSizeBytes to be an explicit nil

### UnsetTotalLogicalBackupSizeBytes
`func (o *ProtectionJobRunStats) UnsetTotalLogicalBackupSizeBytes()`

UnsetTotalLogicalBackupSizeBytes ensures that no value is present for TotalLogicalBackupSizeBytes, not even an explicit nil
### GetTotalPhysicalBackupSizeBytes

`func (o *ProtectionJobRunStats) GetTotalPhysicalBackupSizeBytes() int64`

GetTotalPhysicalBackupSizeBytes returns the TotalPhysicalBackupSizeBytes field if non-nil, zero value otherwise.

### GetTotalPhysicalBackupSizeBytesOk

`func (o *ProtectionJobRunStats) GetTotalPhysicalBackupSizeBytesOk() (*int64, bool)`

GetTotalPhysicalBackupSizeBytesOk returns a tuple with the TotalPhysicalBackupSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPhysicalBackupSizeBytes

`func (o *ProtectionJobRunStats) SetTotalPhysicalBackupSizeBytes(v int64)`

SetTotalPhysicalBackupSizeBytes sets TotalPhysicalBackupSizeBytes field to given value.

### HasTotalPhysicalBackupSizeBytes

`func (o *ProtectionJobRunStats) HasTotalPhysicalBackupSizeBytes() bool`

HasTotalPhysicalBackupSizeBytes returns a boolean if a field has been set.

### SetTotalPhysicalBackupSizeBytesNil

`func (o *ProtectionJobRunStats) SetTotalPhysicalBackupSizeBytesNil(b bool)`

 SetTotalPhysicalBackupSizeBytesNil sets the value for TotalPhysicalBackupSizeBytes to be an explicit nil

### UnsetTotalPhysicalBackupSizeBytes
`func (o *ProtectionJobRunStats) UnsetTotalPhysicalBackupSizeBytes()`

UnsetTotalPhysicalBackupSizeBytes ensures that no value is present for TotalPhysicalBackupSizeBytes, not even an explicit nil
### GetTotalSourceSizeBytes

`func (o *ProtectionJobRunStats) GetTotalSourceSizeBytes() int64`

GetTotalSourceSizeBytes returns the TotalSourceSizeBytes field if non-nil, zero value otherwise.

### GetTotalSourceSizeBytesOk

`func (o *ProtectionJobRunStats) GetTotalSourceSizeBytesOk() (*int64, bool)`

GetTotalSourceSizeBytesOk returns a tuple with the TotalSourceSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSourceSizeBytes

`func (o *ProtectionJobRunStats) SetTotalSourceSizeBytes(v int64)`

SetTotalSourceSizeBytes sets TotalSourceSizeBytes field to given value.

### HasTotalSourceSizeBytes

`func (o *ProtectionJobRunStats) HasTotalSourceSizeBytes() bool`

HasTotalSourceSizeBytes returns a boolean if a field has been set.

### SetTotalSourceSizeBytesNil

`func (o *ProtectionJobRunStats) SetTotalSourceSizeBytesNil(b bool)`

 SetTotalSourceSizeBytesNil sets the value for TotalSourceSizeBytes to be an explicit nil

### UnsetTotalSourceSizeBytes
`func (o *ProtectionJobRunStats) UnsetTotalSourceSizeBytes()`

UnsetTotalSourceSizeBytes ensures that no value is present for TotalSourceSizeBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


