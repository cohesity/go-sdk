# FileSnapshotInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasArchivalCopy** | Pointer to **NullableBool** | If true, this snapshot is located on an archival target (such as a tape or AWS). | [optional] 
**HasLocalCopy** | Pointer to **NullableBool** | If true, this snapshot is located on a local Cohesity Cluster. | [optional] 
**HasRemoteCopy** | Pointer to **NullableBool** | If true, this snapshot is located on a Remote Cohesity Cluster. | [optional] 
**ModifiedTimeUsecs** | Pointer to **NullableInt64** | Specifies the time when the file or folder was last modified. Specified as a Unix epoch Timestamp (in microseconds). | [optional] 
**ReplicaInfoList** | Pointer to [**[]ReplicaInfo**](ReplicaInfo.md) | Specifies the list of replication information about the current snapshot. | [optional] 
**SizeBytes** | Pointer to **NullableInt64** | Specifies the size of the file or folder in bytes. | [optional] 
**Snapshot** | Pointer to [**SnapshotAttempt**](SnapshotAttempt.md) |  | [optional] 

## Methods

### NewFileSnapshotInformation

`func NewFileSnapshotInformation() *FileSnapshotInformation`

NewFileSnapshotInformation instantiates a new FileSnapshotInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileSnapshotInformationWithDefaults

`func NewFileSnapshotInformationWithDefaults() *FileSnapshotInformation`

NewFileSnapshotInformationWithDefaults instantiates a new FileSnapshotInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasArchivalCopy

`func (o *FileSnapshotInformation) GetHasArchivalCopy() bool`

GetHasArchivalCopy returns the HasArchivalCopy field if non-nil, zero value otherwise.

### GetHasArchivalCopyOk

`func (o *FileSnapshotInformation) GetHasArchivalCopyOk() (*bool, bool)`

GetHasArchivalCopyOk returns a tuple with the HasArchivalCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasArchivalCopy

`func (o *FileSnapshotInformation) SetHasArchivalCopy(v bool)`

SetHasArchivalCopy sets HasArchivalCopy field to given value.

### HasHasArchivalCopy

`func (o *FileSnapshotInformation) HasHasArchivalCopy() bool`

HasHasArchivalCopy returns a boolean if a field has been set.

### SetHasArchivalCopyNil

`func (o *FileSnapshotInformation) SetHasArchivalCopyNil(b bool)`

 SetHasArchivalCopyNil sets the value for HasArchivalCopy to be an explicit nil

### UnsetHasArchivalCopy
`func (o *FileSnapshotInformation) UnsetHasArchivalCopy()`

UnsetHasArchivalCopy ensures that no value is present for HasArchivalCopy, not even an explicit nil
### GetHasLocalCopy

`func (o *FileSnapshotInformation) GetHasLocalCopy() bool`

GetHasLocalCopy returns the HasLocalCopy field if non-nil, zero value otherwise.

### GetHasLocalCopyOk

`func (o *FileSnapshotInformation) GetHasLocalCopyOk() (*bool, bool)`

GetHasLocalCopyOk returns a tuple with the HasLocalCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLocalCopy

`func (o *FileSnapshotInformation) SetHasLocalCopy(v bool)`

SetHasLocalCopy sets HasLocalCopy field to given value.

### HasHasLocalCopy

`func (o *FileSnapshotInformation) HasHasLocalCopy() bool`

HasHasLocalCopy returns a boolean if a field has been set.

### SetHasLocalCopyNil

`func (o *FileSnapshotInformation) SetHasLocalCopyNil(b bool)`

 SetHasLocalCopyNil sets the value for HasLocalCopy to be an explicit nil

### UnsetHasLocalCopy
`func (o *FileSnapshotInformation) UnsetHasLocalCopy()`

UnsetHasLocalCopy ensures that no value is present for HasLocalCopy, not even an explicit nil
### GetHasRemoteCopy

`func (o *FileSnapshotInformation) GetHasRemoteCopy() bool`

GetHasRemoteCopy returns the HasRemoteCopy field if non-nil, zero value otherwise.

### GetHasRemoteCopyOk

`func (o *FileSnapshotInformation) GetHasRemoteCopyOk() (*bool, bool)`

GetHasRemoteCopyOk returns a tuple with the HasRemoteCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRemoteCopy

`func (o *FileSnapshotInformation) SetHasRemoteCopy(v bool)`

SetHasRemoteCopy sets HasRemoteCopy field to given value.

### HasHasRemoteCopy

`func (o *FileSnapshotInformation) HasHasRemoteCopy() bool`

HasHasRemoteCopy returns a boolean if a field has been set.

### SetHasRemoteCopyNil

`func (o *FileSnapshotInformation) SetHasRemoteCopyNil(b bool)`

 SetHasRemoteCopyNil sets the value for HasRemoteCopy to be an explicit nil

### UnsetHasRemoteCopy
`func (o *FileSnapshotInformation) UnsetHasRemoteCopy()`

UnsetHasRemoteCopy ensures that no value is present for HasRemoteCopy, not even an explicit nil
### GetModifiedTimeUsecs

`func (o *FileSnapshotInformation) GetModifiedTimeUsecs() int64`

GetModifiedTimeUsecs returns the ModifiedTimeUsecs field if non-nil, zero value otherwise.

### GetModifiedTimeUsecsOk

`func (o *FileSnapshotInformation) GetModifiedTimeUsecsOk() (*int64, bool)`

GetModifiedTimeUsecsOk returns a tuple with the ModifiedTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTimeUsecs

`func (o *FileSnapshotInformation) SetModifiedTimeUsecs(v int64)`

SetModifiedTimeUsecs sets ModifiedTimeUsecs field to given value.

### HasModifiedTimeUsecs

`func (o *FileSnapshotInformation) HasModifiedTimeUsecs() bool`

HasModifiedTimeUsecs returns a boolean if a field has been set.

### SetModifiedTimeUsecsNil

`func (o *FileSnapshotInformation) SetModifiedTimeUsecsNil(b bool)`

 SetModifiedTimeUsecsNil sets the value for ModifiedTimeUsecs to be an explicit nil

### UnsetModifiedTimeUsecs
`func (o *FileSnapshotInformation) UnsetModifiedTimeUsecs()`

UnsetModifiedTimeUsecs ensures that no value is present for ModifiedTimeUsecs, not even an explicit nil
### GetReplicaInfoList

`func (o *FileSnapshotInformation) GetReplicaInfoList() []ReplicaInfo`

GetReplicaInfoList returns the ReplicaInfoList field if non-nil, zero value otherwise.

### GetReplicaInfoListOk

`func (o *FileSnapshotInformation) GetReplicaInfoListOk() (*[]ReplicaInfo, bool)`

GetReplicaInfoListOk returns a tuple with the ReplicaInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaInfoList

`func (o *FileSnapshotInformation) SetReplicaInfoList(v []ReplicaInfo)`

SetReplicaInfoList sets ReplicaInfoList field to given value.

### HasReplicaInfoList

`func (o *FileSnapshotInformation) HasReplicaInfoList() bool`

HasReplicaInfoList returns a boolean if a field has been set.

### SetReplicaInfoListNil

`func (o *FileSnapshotInformation) SetReplicaInfoListNil(b bool)`

 SetReplicaInfoListNil sets the value for ReplicaInfoList to be an explicit nil

### UnsetReplicaInfoList
`func (o *FileSnapshotInformation) UnsetReplicaInfoList()`

UnsetReplicaInfoList ensures that no value is present for ReplicaInfoList, not even an explicit nil
### GetSizeBytes

`func (o *FileSnapshotInformation) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *FileSnapshotInformation) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *FileSnapshotInformation) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *FileSnapshotInformation) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### SetSizeBytesNil

`func (o *FileSnapshotInformation) SetSizeBytesNil(b bool)`

 SetSizeBytesNil sets the value for SizeBytes to be an explicit nil

### UnsetSizeBytes
`func (o *FileSnapshotInformation) UnsetSizeBytes()`

UnsetSizeBytes ensures that no value is present for SizeBytes, not even an explicit nil
### GetSnapshot

`func (o *FileSnapshotInformation) GetSnapshot() SnapshotAttempt`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *FileSnapshotInformation) GetSnapshotOk() (*SnapshotAttempt, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *FileSnapshotInformation) SetSnapshot(v SnapshotAttempt)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *FileSnapshotInformation) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


