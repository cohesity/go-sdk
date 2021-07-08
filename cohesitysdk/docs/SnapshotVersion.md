# SnapshotVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttemptNumber** | Pointer to **NullableInt64** | Specifies the number of the attempts made by the Job Run to capture a snapshot of the object. For example, if an snapshot is successfully captured after three attempts, this field equals 3. | [optional] 
**DeltaSizeBytes** | Pointer to **NullableInt64** | Specifies the size of the data captured from the source object. For a full backup (where Change Block Tracking is not utilized) this field is equal to logicalSizeBytes. For an incremental backup (where Change Block Tracking is utilized), this field specifies the size of the data that has changed since the last backup. | [optional] 
**IndexingStatus** | Pointer to **NullableString** | Specifies the indexing status of the snapshot. Specifies the indexing status of the snapshot. &#39;kStarted&#39; indicates that indexing has started. &#39;kDone&#39; indicates that indexing has been completed according to the type of object. &#39;kNoIndex&#39; indicates that the snapshot cannot be indexed. This is the case during archival restore. &#39;kIceboxRestoreStarted&#39; indicates that indexing is started from an archive. &#39;kIceboxRestoreError&#39; indicates that an error occurred during restore from archiveand there is no index present. &#39;kSkipped&#39; indicates that indexing is skipped due to indexing backlog. | [optional] 
**IsAppConsistent** | Pointer to **NullableBool** | Specifies if an app-consistent snapshot was captured. For example, was the VM was quiesced before the snapshot was captured. | [optional] 
**IsFullBackup** | Pointer to **NullableBool** | Specifies if the snapshot is a full backup. For example, all blocks of the VM is captured and Change Block Tracking is not utilized. | [optional] 
**JobRunId** | Pointer to **NullableInt64** | Specifies the id of the Job Run that captured the snapshot. | [optional] 
**LocalMountPath** | Pointer to **NullableString** | Specifies the local path relative to the View, without the ViewBox/View prefix. | [optional] 
**LogicalSizeBytes** | Pointer to **NullableInt64** | Specifies the size of the snapshot if the data is fully hydrated or expanded and not reduced by change-block tracking, compression and deduplication. For example if a VMDK of size 100GB is created with thin provisioning and the disk size to store the VMDK is 20GB. The logical size of this object is 100GB and the physical size is 20GB. | [optional] 
**PhysicalSizeBytes** | Pointer to **NullableInt64** | Specifies the amount of data actually used on the disk to store this object after being reduced by change-block tracking, compression and deduplication. | [optional] 
**PrimaryPhysicalSizeBytes** | Pointer to **NullableInt64** | Specifies the total amount of disk space used to store this object on the primary storage. For example the total amount of disk space used to store the VM files (such as the VMDK files) on the primary datastore. | [optional] 
**ReplicaInfoList** | Pointer to [**[]ReplicaInfo**](ReplicaInfo.md) | Specifies the list of replication information about the current snapshot. | [optional] 
**StartedTimeUsecs** | Pointer to **NullableInt64** | Specifies the time when the Job Run starts capturing a snapshot. Specified as a Unix epoch Timestamp (in microseconds). | [optional] 

## Methods

### NewSnapshotVersion

`func NewSnapshotVersion() *SnapshotVersion`

NewSnapshotVersion instantiates a new SnapshotVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotVersionWithDefaults

`func NewSnapshotVersionWithDefaults() *SnapshotVersion`

NewSnapshotVersionWithDefaults instantiates a new SnapshotVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttemptNumber

`func (o *SnapshotVersion) GetAttemptNumber() int64`

GetAttemptNumber returns the AttemptNumber field if non-nil, zero value otherwise.

### GetAttemptNumberOk

`func (o *SnapshotVersion) GetAttemptNumberOk() (*int64, bool)`

GetAttemptNumberOk returns a tuple with the AttemptNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptNumber

`func (o *SnapshotVersion) SetAttemptNumber(v int64)`

SetAttemptNumber sets AttemptNumber field to given value.

### HasAttemptNumber

`func (o *SnapshotVersion) HasAttemptNumber() bool`

HasAttemptNumber returns a boolean if a field has been set.

### SetAttemptNumberNil

`func (o *SnapshotVersion) SetAttemptNumberNil(b bool)`

 SetAttemptNumberNil sets the value for AttemptNumber to be an explicit nil

### UnsetAttemptNumber
`func (o *SnapshotVersion) UnsetAttemptNumber()`

UnsetAttemptNumber ensures that no value is present for AttemptNumber, not even an explicit nil
### GetDeltaSizeBytes

`func (o *SnapshotVersion) GetDeltaSizeBytes() int64`

GetDeltaSizeBytes returns the DeltaSizeBytes field if non-nil, zero value otherwise.

### GetDeltaSizeBytesOk

`func (o *SnapshotVersion) GetDeltaSizeBytesOk() (*int64, bool)`

GetDeltaSizeBytesOk returns a tuple with the DeltaSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaSizeBytes

`func (o *SnapshotVersion) SetDeltaSizeBytes(v int64)`

SetDeltaSizeBytes sets DeltaSizeBytes field to given value.

### HasDeltaSizeBytes

`func (o *SnapshotVersion) HasDeltaSizeBytes() bool`

HasDeltaSizeBytes returns a boolean if a field has been set.

### SetDeltaSizeBytesNil

`func (o *SnapshotVersion) SetDeltaSizeBytesNil(b bool)`

 SetDeltaSizeBytesNil sets the value for DeltaSizeBytes to be an explicit nil

### UnsetDeltaSizeBytes
`func (o *SnapshotVersion) UnsetDeltaSizeBytes()`

UnsetDeltaSizeBytes ensures that no value is present for DeltaSizeBytes, not even an explicit nil
### GetIndexingStatus

`func (o *SnapshotVersion) GetIndexingStatus() string`

GetIndexingStatus returns the IndexingStatus field if non-nil, zero value otherwise.

### GetIndexingStatusOk

`func (o *SnapshotVersion) GetIndexingStatusOk() (*string, bool)`

GetIndexingStatusOk returns a tuple with the IndexingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingStatus

`func (o *SnapshotVersion) SetIndexingStatus(v string)`

SetIndexingStatus sets IndexingStatus field to given value.

### HasIndexingStatus

`func (o *SnapshotVersion) HasIndexingStatus() bool`

HasIndexingStatus returns a boolean if a field has been set.

### SetIndexingStatusNil

`func (o *SnapshotVersion) SetIndexingStatusNil(b bool)`

 SetIndexingStatusNil sets the value for IndexingStatus to be an explicit nil

### UnsetIndexingStatus
`func (o *SnapshotVersion) UnsetIndexingStatus()`

UnsetIndexingStatus ensures that no value is present for IndexingStatus, not even an explicit nil
### GetIsAppConsistent

`func (o *SnapshotVersion) GetIsAppConsistent() bool`

GetIsAppConsistent returns the IsAppConsistent field if non-nil, zero value otherwise.

### GetIsAppConsistentOk

`func (o *SnapshotVersion) GetIsAppConsistentOk() (*bool, bool)`

GetIsAppConsistentOk returns a tuple with the IsAppConsistent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAppConsistent

`func (o *SnapshotVersion) SetIsAppConsistent(v bool)`

SetIsAppConsistent sets IsAppConsistent field to given value.

### HasIsAppConsistent

`func (o *SnapshotVersion) HasIsAppConsistent() bool`

HasIsAppConsistent returns a boolean if a field has been set.

### SetIsAppConsistentNil

`func (o *SnapshotVersion) SetIsAppConsistentNil(b bool)`

 SetIsAppConsistentNil sets the value for IsAppConsistent to be an explicit nil

### UnsetIsAppConsistent
`func (o *SnapshotVersion) UnsetIsAppConsistent()`

UnsetIsAppConsistent ensures that no value is present for IsAppConsistent, not even an explicit nil
### GetIsFullBackup

`func (o *SnapshotVersion) GetIsFullBackup() bool`

GetIsFullBackup returns the IsFullBackup field if non-nil, zero value otherwise.

### GetIsFullBackupOk

`func (o *SnapshotVersion) GetIsFullBackupOk() (*bool, bool)`

GetIsFullBackupOk returns a tuple with the IsFullBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFullBackup

`func (o *SnapshotVersion) SetIsFullBackup(v bool)`

SetIsFullBackup sets IsFullBackup field to given value.

### HasIsFullBackup

`func (o *SnapshotVersion) HasIsFullBackup() bool`

HasIsFullBackup returns a boolean if a field has been set.

### SetIsFullBackupNil

`func (o *SnapshotVersion) SetIsFullBackupNil(b bool)`

 SetIsFullBackupNil sets the value for IsFullBackup to be an explicit nil

### UnsetIsFullBackup
`func (o *SnapshotVersion) UnsetIsFullBackup()`

UnsetIsFullBackup ensures that no value is present for IsFullBackup, not even an explicit nil
### GetJobRunId

`func (o *SnapshotVersion) GetJobRunId() int64`

GetJobRunId returns the JobRunId field if non-nil, zero value otherwise.

### GetJobRunIdOk

`func (o *SnapshotVersion) GetJobRunIdOk() (*int64, bool)`

GetJobRunIdOk returns a tuple with the JobRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobRunId

`func (o *SnapshotVersion) SetJobRunId(v int64)`

SetJobRunId sets JobRunId field to given value.

### HasJobRunId

`func (o *SnapshotVersion) HasJobRunId() bool`

HasJobRunId returns a boolean if a field has been set.

### SetJobRunIdNil

`func (o *SnapshotVersion) SetJobRunIdNil(b bool)`

 SetJobRunIdNil sets the value for JobRunId to be an explicit nil

### UnsetJobRunId
`func (o *SnapshotVersion) UnsetJobRunId()`

UnsetJobRunId ensures that no value is present for JobRunId, not even an explicit nil
### GetLocalMountPath

`func (o *SnapshotVersion) GetLocalMountPath() string`

GetLocalMountPath returns the LocalMountPath field if non-nil, zero value otherwise.

### GetLocalMountPathOk

`func (o *SnapshotVersion) GetLocalMountPathOk() (*string, bool)`

GetLocalMountPathOk returns a tuple with the LocalMountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalMountPath

`func (o *SnapshotVersion) SetLocalMountPath(v string)`

SetLocalMountPath sets LocalMountPath field to given value.

### HasLocalMountPath

`func (o *SnapshotVersion) HasLocalMountPath() bool`

HasLocalMountPath returns a boolean if a field has been set.

### SetLocalMountPathNil

`func (o *SnapshotVersion) SetLocalMountPathNil(b bool)`

 SetLocalMountPathNil sets the value for LocalMountPath to be an explicit nil

### UnsetLocalMountPath
`func (o *SnapshotVersion) UnsetLocalMountPath()`

UnsetLocalMountPath ensures that no value is present for LocalMountPath, not even an explicit nil
### GetLogicalSizeBytes

`func (o *SnapshotVersion) GetLogicalSizeBytes() int64`

GetLogicalSizeBytes returns the LogicalSizeBytes field if non-nil, zero value otherwise.

### GetLogicalSizeBytesOk

`func (o *SnapshotVersion) GetLogicalSizeBytesOk() (*int64, bool)`

GetLogicalSizeBytesOk returns a tuple with the LogicalSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalSizeBytes

`func (o *SnapshotVersion) SetLogicalSizeBytes(v int64)`

SetLogicalSizeBytes sets LogicalSizeBytes field to given value.

### HasLogicalSizeBytes

`func (o *SnapshotVersion) HasLogicalSizeBytes() bool`

HasLogicalSizeBytes returns a boolean if a field has been set.

### SetLogicalSizeBytesNil

`func (o *SnapshotVersion) SetLogicalSizeBytesNil(b bool)`

 SetLogicalSizeBytesNil sets the value for LogicalSizeBytes to be an explicit nil

### UnsetLogicalSizeBytes
`func (o *SnapshotVersion) UnsetLogicalSizeBytes()`

UnsetLogicalSizeBytes ensures that no value is present for LogicalSizeBytes, not even an explicit nil
### GetPhysicalSizeBytes

`func (o *SnapshotVersion) GetPhysicalSizeBytes() int64`

GetPhysicalSizeBytes returns the PhysicalSizeBytes field if non-nil, zero value otherwise.

### GetPhysicalSizeBytesOk

`func (o *SnapshotVersion) GetPhysicalSizeBytesOk() (*int64, bool)`

GetPhysicalSizeBytesOk returns a tuple with the PhysicalSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalSizeBytes

`func (o *SnapshotVersion) SetPhysicalSizeBytes(v int64)`

SetPhysicalSizeBytes sets PhysicalSizeBytes field to given value.

### HasPhysicalSizeBytes

`func (o *SnapshotVersion) HasPhysicalSizeBytes() bool`

HasPhysicalSizeBytes returns a boolean if a field has been set.

### SetPhysicalSizeBytesNil

`func (o *SnapshotVersion) SetPhysicalSizeBytesNil(b bool)`

 SetPhysicalSizeBytesNil sets the value for PhysicalSizeBytes to be an explicit nil

### UnsetPhysicalSizeBytes
`func (o *SnapshotVersion) UnsetPhysicalSizeBytes()`

UnsetPhysicalSizeBytes ensures that no value is present for PhysicalSizeBytes, not even an explicit nil
### GetPrimaryPhysicalSizeBytes

`func (o *SnapshotVersion) GetPrimaryPhysicalSizeBytes() int64`

GetPrimaryPhysicalSizeBytes returns the PrimaryPhysicalSizeBytes field if non-nil, zero value otherwise.

### GetPrimaryPhysicalSizeBytesOk

`func (o *SnapshotVersion) GetPrimaryPhysicalSizeBytesOk() (*int64, bool)`

GetPrimaryPhysicalSizeBytesOk returns a tuple with the PrimaryPhysicalSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryPhysicalSizeBytes

`func (o *SnapshotVersion) SetPrimaryPhysicalSizeBytes(v int64)`

SetPrimaryPhysicalSizeBytes sets PrimaryPhysicalSizeBytes field to given value.

### HasPrimaryPhysicalSizeBytes

`func (o *SnapshotVersion) HasPrimaryPhysicalSizeBytes() bool`

HasPrimaryPhysicalSizeBytes returns a boolean if a field has been set.

### SetPrimaryPhysicalSizeBytesNil

`func (o *SnapshotVersion) SetPrimaryPhysicalSizeBytesNil(b bool)`

 SetPrimaryPhysicalSizeBytesNil sets the value for PrimaryPhysicalSizeBytes to be an explicit nil

### UnsetPrimaryPhysicalSizeBytes
`func (o *SnapshotVersion) UnsetPrimaryPhysicalSizeBytes()`

UnsetPrimaryPhysicalSizeBytes ensures that no value is present for PrimaryPhysicalSizeBytes, not even an explicit nil
### GetReplicaInfoList

`func (o *SnapshotVersion) GetReplicaInfoList() []ReplicaInfo`

GetReplicaInfoList returns the ReplicaInfoList field if non-nil, zero value otherwise.

### GetReplicaInfoListOk

`func (o *SnapshotVersion) GetReplicaInfoListOk() (*[]ReplicaInfo, bool)`

GetReplicaInfoListOk returns a tuple with the ReplicaInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaInfoList

`func (o *SnapshotVersion) SetReplicaInfoList(v []ReplicaInfo)`

SetReplicaInfoList sets ReplicaInfoList field to given value.

### HasReplicaInfoList

`func (o *SnapshotVersion) HasReplicaInfoList() bool`

HasReplicaInfoList returns a boolean if a field has been set.

### SetReplicaInfoListNil

`func (o *SnapshotVersion) SetReplicaInfoListNil(b bool)`

 SetReplicaInfoListNil sets the value for ReplicaInfoList to be an explicit nil

### UnsetReplicaInfoList
`func (o *SnapshotVersion) UnsetReplicaInfoList()`

UnsetReplicaInfoList ensures that no value is present for ReplicaInfoList, not even an explicit nil
### GetStartedTimeUsecs

`func (o *SnapshotVersion) GetStartedTimeUsecs() int64`

GetStartedTimeUsecs returns the StartedTimeUsecs field if non-nil, zero value otherwise.

### GetStartedTimeUsecsOk

`func (o *SnapshotVersion) GetStartedTimeUsecsOk() (*int64, bool)`

GetStartedTimeUsecsOk returns a tuple with the StartedTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedTimeUsecs

`func (o *SnapshotVersion) SetStartedTimeUsecs(v int64)`

SetStartedTimeUsecs sets StartedTimeUsecs field to given value.

### HasStartedTimeUsecs

`func (o *SnapshotVersion) HasStartedTimeUsecs() bool`

HasStartedTimeUsecs returns a boolean if a field has been set.

### SetStartedTimeUsecsNil

`func (o *SnapshotVersion) SetStartedTimeUsecsNil(b bool)`

 SetStartedTimeUsecsNil sets the value for StartedTimeUsecs to be an explicit nil

### UnsetStartedTimeUsecs
`func (o *SnapshotVersion) UnsetStartedTimeUsecs()`

UnsetStartedTimeUsecs ensures that no value is present for StartedTimeUsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


