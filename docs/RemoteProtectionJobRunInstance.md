# RemoteProtectionJobRunInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveTaskUid** | Pointer to [**NullableUniversalId**](UniversalId.md) | Specifies the globally unique id of the archival task that archived the Snapshot to the Vault. | [optional] 
**ArchiveVersion** | Pointer to **NullableInt32** | Specifies the version of the archive. | [optional] 
**ExpiryTimeUsecs** | Pointer to **NullableInt64** | Specifies the time when the archive expires. This time is recorded as a Unix epoch Timestamp (in microseconds). | [optional] 
**IndexSizeBytes** | Pointer to **NullableInt64** | Specifies the size of the index for the archive. | [optional] 
**JobRunId** | Pointer to **NullableInt64** | Specifies the instance id of the Job Run task capturing the Snapshot. | [optional] 
**MetadataComplete** | Pointer to **NullableBool** | Specifies whether a full set of metadata is available now. | [optional] 
**SnapshotTimeUsecs** | Pointer to **NullableInt64** | Specify the time the Snapshot was captured as a Unix epoch Timestamp (in microseconds). | [optional] 

## Methods

### NewRemoteProtectionJobRunInstance

`func NewRemoteProtectionJobRunInstance() *RemoteProtectionJobRunInstance`

NewRemoteProtectionJobRunInstance instantiates a new RemoteProtectionJobRunInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteProtectionJobRunInstanceWithDefaults

`func NewRemoteProtectionJobRunInstanceWithDefaults() *RemoteProtectionJobRunInstance`

NewRemoteProtectionJobRunInstanceWithDefaults instantiates a new RemoteProtectionJobRunInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveTaskUid

`func (o *RemoteProtectionJobRunInstance) GetArchiveTaskUid() UniversalId`

GetArchiveTaskUid returns the ArchiveTaskUid field if non-nil, zero value otherwise.

### GetArchiveTaskUidOk

`func (o *RemoteProtectionJobRunInstance) GetArchiveTaskUidOk() (*UniversalId, bool)`

GetArchiveTaskUidOk returns a tuple with the ArchiveTaskUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveTaskUid

`func (o *RemoteProtectionJobRunInstance) SetArchiveTaskUid(v UniversalId)`

SetArchiveTaskUid sets ArchiveTaskUid field to given value.

### HasArchiveTaskUid

`func (o *RemoteProtectionJobRunInstance) HasArchiveTaskUid() bool`

HasArchiveTaskUid returns a boolean if a field has been set.

### SetArchiveTaskUidNil

`func (o *RemoteProtectionJobRunInstance) SetArchiveTaskUidNil(b bool)`

 SetArchiveTaskUidNil sets the value for ArchiveTaskUid to be an explicit nil

### UnsetArchiveTaskUid
`func (o *RemoteProtectionJobRunInstance) UnsetArchiveTaskUid()`

UnsetArchiveTaskUid ensures that no value is present for ArchiveTaskUid, not even an explicit nil
### GetArchiveVersion

`func (o *RemoteProtectionJobRunInstance) GetArchiveVersion() int32`

GetArchiveVersion returns the ArchiveVersion field if non-nil, zero value otherwise.

### GetArchiveVersionOk

`func (o *RemoteProtectionJobRunInstance) GetArchiveVersionOk() (*int32, bool)`

GetArchiveVersionOk returns a tuple with the ArchiveVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveVersion

`func (o *RemoteProtectionJobRunInstance) SetArchiveVersion(v int32)`

SetArchiveVersion sets ArchiveVersion field to given value.

### HasArchiveVersion

`func (o *RemoteProtectionJobRunInstance) HasArchiveVersion() bool`

HasArchiveVersion returns a boolean if a field has been set.

### SetArchiveVersionNil

`func (o *RemoteProtectionJobRunInstance) SetArchiveVersionNil(b bool)`

 SetArchiveVersionNil sets the value for ArchiveVersion to be an explicit nil

### UnsetArchiveVersion
`func (o *RemoteProtectionJobRunInstance) UnsetArchiveVersion()`

UnsetArchiveVersion ensures that no value is present for ArchiveVersion, not even an explicit nil
### GetExpiryTimeUsecs

`func (o *RemoteProtectionJobRunInstance) GetExpiryTimeUsecs() int64`

GetExpiryTimeUsecs returns the ExpiryTimeUsecs field if non-nil, zero value otherwise.

### GetExpiryTimeUsecsOk

`func (o *RemoteProtectionJobRunInstance) GetExpiryTimeUsecsOk() (*int64, bool)`

GetExpiryTimeUsecsOk returns a tuple with the ExpiryTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimeUsecs

`func (o *RemoteProtectionJobRunInstance) SetExpiryTimeUsecs(v int64)`

SetExpiryTimeUsecs sets ExpiryTimeUsecs field to given value.

### HasExpiryTimeUsecs

`func (o *RemoteProtectionJobRunInstance) HasExpiryTimeUsecs() bool`

HasExpiryTimeUsecs returns a boolean if a field has been set.

### SetExpiryTimeUsecsNil

`func (o *RemoteProtectionJobRunInstance) SetExpiryTimeUsecsNil(b bool)`

 SetExpiryTimeUsecsNil sets the value for ExpiryTimeUsecs to be an explicit nil

### UnsetExpiryTimeUsecs
`func (o *RemoteProtectionJobRunInstance) UnsetExpiryTimeUsecs()`

UnsetExpiryTimeUsecs ensures that no value is present for ExpiryTimeUsecs, not even an explicit nil
### GetIndexSizeBytes

`func (o *RemoteProtectionJobRunInstance) GetIndexSizeBytes() int64`

GetIndexSizeBytes returns the IndexSizeBytes field if non-nil, zero value otherwise.

### GetIndexSizeBytesOk

`func (o *RemoteProtectionJobRunInstance) GetIndexSizeBytesOk() (*int64, bool)`

GetIndexSizeBytesOk returns a tuple with the IndexSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexSizeBytes

`func (o *RemoteProtectionJobRunInstance) SetIndexSizeBytes(v int64)`

SetIndexSizeBytes sets IndexSizeBytes field to given value.

### HasIndexSizeBytes

`func (o *RemoteProtectionJobRunInstance) HasIndexSizeBytes() bool`

HasIndexSizeBytes returns a boolean if a field has been set.

### SetIndexSizeBytesNil

`func (o *RemoteProtectionJobRunInstance) SetIndexSizeBytesNil(b bool)`

 SetIndexSizeBytesNil sets the value for IndexSizeBytes to be an explicit nil

### UnsetIndexSizeBytes
`func (o *RemoteProtectionJobRunInstance) UnsetIndexSizeBytes()`

UnsetIndexSizeBytes ensures that no value is present for IndexSizeBytes, not even an explicit nil
### GetJobRunId

`func (o *RemoteProtectionJobRunInstance) GetJobRunId() int64`

GetJobRunId returns the JobRunId field if non-nil, zero value otherwise.

### GetJobRunIdOk

`func (o *RemoteProtectionJobRunInstance) GetJobRunIdOk() (*int64, bool)`

GetJobRunIdOk returns a tuple with the JobRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobRunId

`func (o *RemoteProtectionJobRunInstance) SetJobRunId(v int64)`

SetJobRunId sets JobRunId field to given value.

### HasJobRunId

`func (o *RemoteProtectionJobRunInstance) HasJobRunId() bool`

HasJobRunId returns a boolean if a field has been set.

### SetJobRunIdNil

`func (o *RemoteProtectionJobRunInstance) SetJobRunIdNil(b bool)`

 SetJobRunIdNil sets the value for JobRunId to be an explicit nil

### UnsetJobRunId
`func (o *RemoteProtectionJobRunInstance) UnsetJobRunId()`

UnsetJobRunId ensures that no value is present for JobRunId, not even an explicit nil
### GetMetadataComplete

`func (o *RemoteProtectionJobRunInstance) GetMetadataComplete() bool`

GetMetadataComplete returns the MetadataComplete field if non-nil, zero value otherwise.

### GetMetadataCompleteOk

`func (o *RemoteProtectionJobRunInstance) GetMetadataCompleteOk() (*bool, bool)`

GetMetadataCompleteOk returns a tuple with the MetadataComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataComplete

`func (o *RemoteProtectionJobRunInstance) SetMetadataComplete(v bool)`

SetMetadataComplete sets MetadataComplete field to given value.

### HasMetadataComplete

`func (o *RemoteProtectionJobRunInstance) HasMetadataComplete() bool`

HasMetadataComplete returns a boolean if a field has been set.

### SetMetadataCompleteNil

`func (o *RemoteProtectionJobRunInstance) SetMetadataCompleteNil(b bool)`

 SetMetadataCompleteNil sets the value for MetadataComplete to be an explicit nil

### UnsetMetadataComplete
`func (o *RemoteProtectionJobRunInstance) UnsetMetadataComplete()`

UnsetMetadataComplete ensures that no value is present for MetadataComplete, not even an explicit nil
### GetSnapshotTimeUsecs

`func (o *RemoteProtectionJobRunInstance) GetSnapshotTimeUsecs() int64`

GetSnapshotTimeUsecs returns the SnapshotTimeUsecs field if non-nil, zero value otherwise.

### GetSnapshotTimeUsecsOk

`func (o *RemoteProtectionJobRunInstance) GetSnapshotTimeUsecsOk() (*int64, bool)`

GetSnapshotTimeUsecsOk returns a tuple with the SnapshotTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTimeUsecs

`func (o *RemoteProtectionJobRunInstance) SetSnapshotTimeUsecs(v int64)`

SetSnapshotTimeUsecs sets SnapshotTimeUsecs field to given value.

### HasSnapshotTimeUsecs

`func (o *RemoteProtectionJobRunInstance) HasSnapshotTimeUsecs() bool`

HasSnapshotTimeUsecs returns a boolean if a field has been set.

### SetSnapshotTimeUsecsNil

`func (o *RemoteProtectionJobRunInstance) SetSnapshotTimeUsecsNil(b bool)`

 SetSnapshotTimeUsecsNil sets the value for SnapshotTimeUsecs to be an explicit nil

### UnsetSnapshotTimeUsecs
`func (o *RemoteProtectionJobRunInstance) UnsetSnapshotTimeUsecs()`

UnsetSnapshotTimeUsecs ensures that no value is present for SnapshotTimeUsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


