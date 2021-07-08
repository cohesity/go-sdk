# BackupSourceStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdmittedTimeUsecs** | Pointer to **NullableInt64** | Specifies the time the task was unqueued from the queue to start running. This field can be used to determine the following times: initial-wait-time &#x3D; admittedTimeUsecs - startTimeUsecs run-time &#x3D; endTimeUsecs - admittedTimeUsecs If the task ends up waiting in other queues, then actual run-time will be smaller than the run-time computed this way. This field is only populated for Backup tasks currently. | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time of the Protection Run. The end time is specified as a Unix epoch Timestamp (in microseconds). | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of the Protection Run. The start time is specified as a Unix epoch Timestamp (in microseconds). This time is when the task is queued to an internal queue where tasks are waiting to run. | [optional] 
**TimeTakenUsecs** | Pointer to **NullableInt64** | Specifies the actual execution time for the protection run to complete the backup task and the copy tasks. This time will not include the time waited in various internal queues. This field is only populated for Backup tasks currently. | [optional] 
**TotalBytesReadFromSource** | Pointer to **NullableInt64** | Specifies the total amount of data read from the source (so far). | [optional] 
**TotalBytesToReadFromSource** | Pointer to **NullableInt64** | Specifies the total amount of data expected to be read from the source. | [optional] 
**TotalLogicalBackupSizeBytes** | Pointer to **NullableInt64** | Specifies the size of the source object (such as a VM) protected by this task on the primary storage after the snapshot is taken. The logical size of the data on the source if the data is fully hydrated or expanded and not reduced by change-block tracking, compression and deduplication. | [optional] 
**TotalPhysicalBackupSizeBytes** | Pointer to **NullableInt64** | Specifies the total amount of physical space used on the Cohesity Cluster to store the protected object after being reduced by change-block tracking, compression and deduplication. For example, if the logical backup size is 1GB, but only 1MB was used on the Cohesity Cluster to store it, this field be equal to 1MB. | [optional] 
**TotalSourceSizeBytes** | Pointer to **NullableInt64** | Specifies the size of the source object (such as a VM) protected by this task on the primary storage before the snapshot is taken. The logical size of the data on the source if the data is fully hydrated or expanded and not reduced by change-block tracking, compression and deduplication. | [optional] 

## Methods

### NewBackupSourceStats

`func NewBackupSourceStats() *BackupSourceStats`

NewBackupSourceStats instantiates a new BackupSourceStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupSourceStatsWithDefaults

`func NewBackupSourceStatsWithDefaults() *BackupSourceStats`

NewBackupSourceStatsWithDefaults instantiates a new BackupSourceStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmittedTimeUsecs

`func (o *BackupSourceStats) GetAdmittedTimeUsecs() int64`

GetAdmittedTimeUsecs returns the AdmittedTimeUsecs field if non-nil, zero value otherwise.

### GetAdmittedTimeUsecsOk

`func (o *BackupSourceStats) GetAdmittedTimeUsecsOk() (*int64, bool)`

GetAdmittedTimeUsecsOk returns a tuple with the AdmittedTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmittedTimeUsecs

`func (o *BackupSourceStats) SetAdmittedTimeUsecs(v int64)`

SetAdmittedTimeUsecs sets AdmittedTimeUsecs field to given value.

### HasAdmittedTimeUsecs

`func (o *BackupSourceStats) HasAdmittedTimeUsecs() bool`

HasAdmittedTimeUsecs returns a boolean if a field has been set.

### SetAdmittedTimeUsecsNil

`func (o *BackupSourceStats) SetAdmittedTimeUsecsNil(b bool)`

 SetAdmittedTimeUsecsNil sets the value for AdmittedTimeUsecs to be an explicit nil

### UnsetAdmittedTimeUsecs
`func (o *BackupSourceStats) UnsetAdmittedTimeUsecs()`

UnsetAdmittedTimeUsecs ensures that no value is present for AdmittedTimeUsecs, not even an explicit nil
### GetEndTimeUsecs

`func (o *BackupSourceStats) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *BackupSourceStats) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *BackupSourceStats) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *BackupSourceStats) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *BackupSourceStats) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *BackupSourceStats) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetStartTimeUsecs

`func (o *BackupSourceStats) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *BackupSourceStats) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *BackupSourceStats) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *BackupSourceStats) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *BackupSourceStats) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *BackupSourceStats) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetTimeTakenUsecs

`func (o *BackupSourceStats) GetTimeTakenUsecs() int64`

GetTimeTakenUsecs returns the TimeTakenUsecs field if non-nil, zero value otherwise.

### GetTimeTakenUsecsOk

`func (o *BackupSourceStats) GetTimeTakenUsecsOk() (*int64, bool)`

GetTimeTakenUsecsOk returns a tuple with the TimeTakenUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeTakenUsecs

`func (o *BackupSourceStats) SetTimeTakenUsecs(v int64)`

SetTimeTakenUsecs sets TimeTakenUsecs field to given value.

### HasTimeTakenUsecs

`func (o *BackupSourceStats) HasTimeTakenUsecs() bool`

HasTimeTakenUsecs returns a boolean if a field has been set.

### SetTimeTakenUsecsNil

`func (o *BackupSourceStats) SetTimeTakenUsecsNil(b bool)`

 SetTimeTakenUsecsNil sets the value for TimeTakenUsecs to be an explicit nil

### UnsetTimeTakenUsecs
`func (o *BackupSourceStats) UnsetTimeTakenUsecs()`

UnsetTimeTakenUsecs ensures that no value is present for TimeTakenUsecs, not even an explicit nil
### GetTotalBytesReadFromSource

`func (o *BackupSourceStats) GetTotalBytesReadFromSource() int64`

GetTotalBytesReadFromSource returns the TotalBytesReadFromSource field if non-nil, zero value otherwise.

### GetTotalBytesReadFromSourceOk

`func (o *BackupSourceStats) GetTotalBytesReadFromSourceOk() (*int64, bool)`

GetTotalBytesReadFromSourceOk returns a tuple with the TotalBytesReadFromSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytesReadFromSource

`func (o *BackupSourceStats) SetTotalBytesReadFromSource(v int64)`

SetTotalBytesReadFromSource sets TotalBytesReadFromSource field to given value.

### HasTotalBytesReadFromSource

`func (o *BackupSourceStats) HasTotalBytesReadFromSource() bool`

HasTotalBytesReadFromSource returns a boolean if a field has been set.

### SetTotalBytesReadFromSourceNil

`func (o *BackupSourceStats) SetTotalBytesReadFromSourceNil(b bool)`

 SetTotalBytesReadFromSourceNil sets the value for TotalBytesReadFromSource to be an explicit nil

### UnsetTotalBytesReadFromSource
`func (o *BackupSourceStats) UnsetTotalBytesReadFromSource()`

UnsetTotalBytesReadFromSource ensures that no value is present for TotalBytesReadFromSource, not even an explicit nil
### GetTotalBytesToReadFromSource

`func (o *BackupSourceStats) GetTotalBytesToReadFromSource() int64`

GetTotalBytesToReadFromSource returns the TotalBytesToReadFromSource field if non-nil, zero value otherwise.

### GetTotalBytesToReadFromSourceOk

`func (o *BackupSourceStats) GetTotalBytesToReadFromSourceOk() (*int64, bool)`

GetTotalBytesToReadFromSourceOk returns a tuple with the TotalBytesToReadFromSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytesToReadFromSource

`func (o *BackupSourceStats) SetTotalBytesToReadFromSource(v int64)`

SetTotalBytesToReadFromSource sets TotalBytesToReadFromSource field to given value.

### HasTotalBytesToReadFromSource

`func (o *BackupSourceStats) HasTotalBytesToReadFromSource() bool`

HasTotalBytesToReadFromSource returns a boolean if a field has been set.

### SetTotalBytesToReadFromSourceNil

`func (o *BackupSourceStats) SetTotalBytesToReadFromSourceNil(b bool)`

 SetTotalBytesToReadFromSourceNil sets the value for TotalBytesToReadFromSource to be an explicit nil

### UnsetTotalBytesToReadFromSource
`func (o *BackupSourceStats) UnsetTotalBytesToReadFromSource()`

UnsetTotalBytesToReadFromSource ensures that no value is present for TotalBytesToReadFromSource, not even an explicit nil
### GetTotalLogicalBackupSizeBytes

`func (o *BackupSourceStats) GetTotalLogicalBackupSizeBytes() int64`

GetTotalLogicalBackupSizeBytes returns the TotalLogicalBackupSizeBytes field if non-nil, zero value otherwise.

### GetTotalLogicalBackupSizeBytesOk

`func (o *BackupSourceStats) GetTotalLogicalBackupSizeBytesOk() (*int64, bool)`

GetTotalLogicalBackupSizeBytesOk returns a tuple with the TotalLogicalBackupSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLogicalBackupSizeBytes

`func (o *BackupSourceStats) SetTotalLogicalBackupSizeBytes(v int64)`

SetTotalLogicalBackupSizeBytes sets TotalLogicalBackupSizeBytes field to given value.

### HasTotalLogicalBackupSizeBytes

`func (o *BackupSourceStats) HasTotalLogicalBackupSizeBytes() bool`

HasTotalLogicalBackupSizeBytes returns a boolean if a field has been set.

### SetTotalLogicalBackupSizeBytesNil

`func (o *BackupSourceStats) SetTotalLogicalBackupSizeBytesNil(b bool)`

 SetTotalLogicalBackupSizeBytesNil sets the value for TotalLogicalBackupSizeBytes to be an explicit nil

### UnsetTotalLogicalBackupSizeBytes
`func (o *BackupSourceStats) UnsetTotalLogicalBackupSizeBytes()`

UnsetTotalLogicalBackupSizeBytes ensures that no value is present for TotalLogicalBackupSizeBytes, not even an explicit nil
### GetTotalPhysicalBackupSizeBytes

`func (o *BackupSourceStats) GetTotalPhysicalBackupSizeBytes() int64`

GetTotalPhysicalBackupSizeBytes returns the TotalPhysicalBackupSizeBytes field if non-nil, zero value otherwise.

### GetTotalPhysicalBackupSizeBytesOk

`func (o *BackupSourceStats) GetTotalPhysicalBackupSizeBytesOk() (*int64, bool)`

GetTotalPhysicalBackupSizeBytesOk returns a tuple with the TotalPhysicalBackupSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPhysicalBackupSizeBytes

`func (o *BackupSourceStats) SetTotalPhysicalBackupSizeBytes(v int64)`

SetTotalPhysicalBackupSizeBytes sets TotalPhysicalBackupSizeBytes field to given value.

### HasTotalPhysicalBackupSizeBytes

`func (o *BackupSourceStats) HasTotalPhysicalBackupSizeBytes() bool`

HasTotalPhysicalBackupSizeBytes returns a boolean if a field has been set.

### SetTotalPhysicalBackupSizeBytesNil

`func (o *BackupSourceStats) SetTotalPhysicalBackupSizeBytesNil(b bool)`

 SetTotalPhysicalBackupSizeBytesNil sets the value for TotalPhysicalBackupSizeBytes to be an explicit nil

### UnsetTotalPhysicalBackupSizeBytes
`func (o *BackupSourceStats) UnsetTotalPhysicalBackupSizeBytes()`

UnsetTotalPhysicalBackupSizeBytes ensures that no value is present for TotalPhysicalBackupSizeBytes, not even an explicit nil
### GetTotalSourceSizeBytes

`func (o *BackupSourceStats) GetTotalSourceSizeBytes() int64`

GetTotalSourceSizeBytes returns the TotalSourceSizeBytes field if non-nil, zero value otherwise.

### GetTotalSourceSizeBytesOk

`func (o *BackupSourceStats) GetTotalSourceSizeBytesOk() (*int64, bool)`

GetTotalSourceSizeBytesOk returns a tuple with the TotalSourceSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSourceSizeBytes

`func (o *BackupSourceStats) SetTotalSourceSizeBytes(v int64)`

SetTotalSourceSizeBytes sets TotalSourceSizeBytes field to given value.

### HasTotalSourceSizeBytes

`func (o *BackupSourceStats) HasTotalSourceSizeBytes() bool`

HasTotalSourceSizeBytes returns a boolean if a field has been set.

### SetTotalSourceSizeBytesNil

`func (o *BackupSourceStats) SetTotalSourceSizeBytesNil(b bool)`

 SetTotalSourceSizeBytesNil sets the value for TotalSourceSizeBytes to be an explicit nil

### UnsetTotalSourceSizeBytes
`func (o *BackupSourceStats) UnsetTotalSourceSizeBytes()`

UnsetTotalSourceSizeBytes ensures that no value is present for TotalSourceSizeBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


