# IndexedObjectSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IndexedObjectName** | Pointer to **NullableString** | Specifies the indexed object name. | [optional] 
**ObjectSnapshotid** | Pointer to **NullableString** | Specifies snapshot id of the object containing this indexed object. | [optional] 
**SnapshotTimestampUsecs** | Pointer to **NullableInt64** | Specifies a unix timestamp when the object snapshot was taken in micro seconds. | [optional] 
**RunType** | Pointer to **NullableString** | Specifies the type of protection run created this snapshot. | [optional] 
**ProtectionGroupId** | Pointer to **NullableString** | Specifies the protection group id which contains this snapshot. | [optional] 
**ProtectionGroupName** | Pointer to **NullableString** | Specifies the protection group name which contains this snapshot. | [optional] 
**StorageDomainId** | Pointer to **NullableInt64** | Specifies the storage domain id containing this snapshot. | [optional] 
**Attempts** | Pointer to **NullableInt64** | Specifies the number of runs have been executed before the run completed successfully. | [optional] 
**SizeBytes** | Pointer to **NullableInt64** | Specifies the indexed object size in bytes. | [optional] 
**InodeId** | Pointer to **NullableInt64** | Specifies the source inode number of the file being recovered. | [optional] [readonly] 
**LastModifiedTimeUsecs** | Pointer to **NullableInt64** | Specifies the last time file was modified in unix timestamp. | [optional] 
**ExternalTargetInfo** | Pointer to [**NullableArchivalTargetSummaryInfo**](ArchivalTargetSummaryInfo.md) | Specifies the external target information if this is an archival snapshot. | [optional] 

## Methods

### NewIndexedObjectSnapshot

`func NewIndexedObjectSnapshot() *IndexedObjectSnapshot`

NewIndexedObjectSnapshot instantiates a new IndexedObjectSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexedObjectSnapshotWithDefaults

`func NewIndexedObjectSnapshotWithDefaults() *IndexedObjectSnapshot`

NewIndexedObjectSnapshotWithDefaults instantiates a new IndexedObjectSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexedObjectName

`func (o *IndexedObjectSnapshot) GetIndexedObjectName() string`

GetIndexedObjectName returns the IndexedObjectName field if non-nil, zero value otherwise.

### GetIndexedObjectNameOk

`func (o *IndexedObjectSnapshot) GetIndexedObjectNameOk() (*string, bool)`

GetIndexedObjectNameOk returns a tuple with the IndexedObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexedObjectName

`func (o *IndexedObjectSnapshot) SetIndexedObjectName(v string)`

SetIndexedObjectName sets IndexedObjectName field to given value.

### HasIndexedObjectName

`func (o *IndexedObjectSnapshot) HasIndexedObjectName() bool`

HasIndexedObjectName returns a boolean if a field has been set.

### SetIndexedObjectNameNil

`func (o *IndexedObjectSnapshot) SetIndexedObjectNameNil(b bool)`

 SetIndexedObjectNameNil sets the value for IndexedObjectName to be an explicit nil

### UnsetIndexedObjectName
`func (o *IndexedObjectSnapshot) UnsetIndexedObjectName()`

UnsetIndexedObjectName ensures that no value is present for IndexedObjectName, not even an explicit nil
### GetObjectSnapshotid

`func (o *IndexedObjectSnapshot) GetObjectSnapshotid() string`

GetObjectSnapshotid returns the ObjectSnapshotid field if non-nil, zero value otherwise.

### GetObjectSnapshotidOk

`func (o *IndexedObjectSnapshot) GetObjectSnapshotidOk() (*string, bool)`

GetObjectSnapshotidOk returns a tuple with the ObjectSnapshotid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSnapshotid

`func (o *IndexedObjectSnapshot) SetObjectSnapshotid(v string)`

SetObjectSnapshotid sets ObjectSnapshotid field to given value.

### HasObjectSnapshotid

`func (o *IndexedObjectSnapshot) HasObjectSnapshotid() bool`

HasObjectSnapshotid returns a boolean if a field has been set.

### SetObjectSnapshotidNil

`func (o *IndexedObjectSnapshot) SetObjectSnapshotidNil(b bool)`

 SetObjectSnapshotidNil sets the value for ObjectSnapshotid to be an explicit nil

### UnsetObjectSnapshotid
`func (o *IndexedObjectSnapshot) UnsetObjectSnapshotid()`

UnsetObjectSnapshotid ensures that no value is present for ObjectSnapshotid, not even an explicit nil
### GetSnapshotTimestampUsecs

`func (o *IndexedObjectSnapshot) GetSnapshotTimestampUsecs() int64`

GetSnapshotTimestampUsecs returns the SnapshotTimestampUsecs field if non-nil, zero value otherwise.

### GetSnapshotTimestampUsecsOk

`func (o *IndexedObjectSnapshot) GetSnapshotTimestampUsecsOk() (*int64, bool)`

GetSnapshotTimestampUsecsOk returns a tuple with the SnapshotTimestampUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTimestampUsecs

`func (o *IndexedObjectSnapshot) SetSnapshotTimestampUsecs(v int64)`

SetSnapshotTimestampUsecs sets SnapshotTimestampUsecs field to given value.

### HasSnapshotTimestampUsecs

`func (o *IndexedObjectSnapshot) HasSnapshotTimestampUsecs() bool`

HasSnapshotTimestampUsecs returns a boolean if a field has been set.

### SetSnapshotTimestampUsecsNil

`func (o *IndexedObjectSnapshot) SetSnapshotTimestampUsecsNil(b bool)`

 SetSnapshotTimestampUsecsNil sets the value for SnapshotTimestampUsecs to be an explicit nil

### UnsetSnapshotTimestampUsecs
`func (o *IndexedObjectSnapshot) UnsetSnapshotTimestampUsecs()`

UnsetSnapshotTimestampUsecs ensures that no value is present for SnapshotTimestampUsecs, not even an explicit nil
### GetRunType

`func (o *IndexedObjectSnapshot) GetRunType() string`

GetRunType returns the RunType field if non-nil, zero value otherwise.

### GetRunTypeOk

`func (o *IndexedObjectSnapshot) GetRunTypeOk() (*string, bool)`

GetRunTypeOk returns a tuple with the RunType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunType

`func (o *IndexedObjectSnapshot) SetRunType(v string)`

SetRunType sets RunType field to given value.

### HasRunType

`func (o *IndexedObjectSnapshot) HasRunType() bool`

HasRunType returns a boolean if a field has been set.

### SetRunTypeNil

`func (o *IndexedObjectSnapshot) SetRunTypeNil(b bool)`

 SetRunTypeNil sets the value for RunType to be an explicit nil

### UnsetRunType
`func (o *IndexedObjectSnapshot) UnsetRunType()`

UnsetRunType ensures that no value is present for RunType, not even an explicit nil
### GetProtectionGroupId

`func (o *IndexedObjectSnapshot) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *IndexedObjectSnapshot) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *IndexedObjectSnapshot) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *IndexedObjectSnapshot) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *IndexedObjectSnapshot) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *IndexedObjectSnapshot) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetProtectionGroupName

`func (o *IndexedObjectSnapshot) GetProtectionGroupName() string`

GetProtectionGroupName returns the ProtectionGroupName field if non-nil, zero value otherwise.

### GetProtectionGroupNameOk

`func (o *IndexedObjectSnapshot) GetProtectionGroupNameOk() (*string, bool)`

GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupName

`func (o *IndexedObjectSnapshot) SetProtectionGroupName(v string)`

SetProtectionGroupName sets ProtectionGroupName field to given value.

### HasProtectionGroupName

`func (o *IndexedObjectSnapshot) HasProtectionGroupName() bool`

HasProtectionGroupName returns a boolean if a field has been set.

### SetProtectionGroupNameNil

`func (o *IndexedObjectSnapshot) SetProtectionGroupNameNil(b bool)`

 SetProtectionGroupNameNil sets the value for ProtectionGroupName to be an explicit nil

### UnsetProtectionGroupName
`func (o *IndexedObjectSnapshot) UnsetProtectionGroupName()`

UnsetProtectionGroupName ensures that no value is present for ProtectionGroupName, not even an explicit nil
### GetStorageDomainId

`func (o *IndexedObjectSnapshot) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *IndexedObjectSnapshot) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *IndexedObjectSnapshot) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.

### HasStorageDomainId

`func (o *IndexedObjectSnapshot) HasStorageDomainId() bool`

HasStorageDomainId returns a boolean if a field has been set.

### SetStorageDomainIdNil

`func (o *IndexedObjectSnapshot) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *IndexedObjectSnapshot) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
### GetAttempts

`func (o *IndexedObjectSnapshot) GetAttempts() int64`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *IndexedObjectSnapshot) GetAttemptsOk() (*int64, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *IndexedObjectSnapshot) SetAttempts(v int64)`

SetAttempts sets Attempts field to given value.

### HasAttempts

`func (o *IndexedObjectSnapshot) HasAttempts() bool`

HasAttempts returns a boolean if a field has been set.

### SetAttemptsNil

`func (o *IndexedObjectSnapshot) SetAttemptsNil(b bool)`

 SetAttemptsNil sets the value for Attempts to be an explicit nil

### UnsetAttempts
`func (o *IndexedObjectSnapshot) UnsetAttempts()`

UnsetAttempts ensures that no value is present for Attempts, not even an explicit nil
### GetSizeBytes

`func (o *IndexedObjectSnapshot) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *IndexedObjectSnapshot) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *IndexedObjectSnapshot) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *IndexedObjectSnapshot) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### SetSizeBytesNil

`func (o *IndexedObjectSnapshot) SetSizeBytesNil(b bool)`

 SetSizeBytesNil sets the value for SizeBytes to be an explicit nil

### UnsetSizeBytes
`func (o *IndexedObjectSnapshot) UnsetSizeBytes()`

UnsetSizeBytes ensures that no value is present for SizeBytes, not even an explicit nil
### GetInodeId

`func (o *IndexedObjectSnapshot) GetInodeId() int64`

GetInodeId returns the InodeId field if non-nil, zero value otherwise.

### GetInodeIdOk

`func (o *IndexedObjectSnapshot) GetInodeIdOk() (*int64, bool)`

GetInodeIdOk returns a tuple with the InodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInodeId

`func (o *IndexedObjectSnapshot) SetInodeId(v int64)`

SetInodeId sets InodeId field to given value.

### HasInodeId

`func (o *IndexedObjectSnapshot) HasInodeId() bool`

HasInodeId returns a boolean if a field has been set.

### SetInodeIdNil

`func (o *IndexedObjectSnapshot) SetInodeIdNil(b bool)`

 SetInodeIdNil sets the value for InodeId to be an explicit nil

### UnsetInodeId
`func (o *IndexedObjectSnapshot) UnsetInodeId()`

UnsetInodeId ensures that no value is present for InodeId, not even an explicit nil
### GetLastModifiedTimeUsecs

`func (o *IndexedObjectSnapshot) GetLastModifiedTimeUsecs() int64`

GetLastModifiedTimeUsecs returns the LastModifiedTimeUsecs field if non-nil, zero value otherwise.

### GetLastModifiedTimeUsecsOk

`func (o *IndexedObjectSnapshot) GetLastModifiedTimeUsecsOk() (*int64, bool)`

GetLastModifiedTimeUsecsOk returns a tuple with the LastModifiedTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTimeUsecs

`func (o *IndexedObjectSnapshot) SetLastModifiedTimeUsecs(v int64)`

SetLastModifiedTimeUsecs sets LastModifiedTimeUsecs field to given value.

### HasLastModifiedTimeUsecs

`func (o *IndexedObjectSnapshot) HasLastModifiedTimeUsecs() bool`

HasLastModifiedTimeUsecs returns a boolean if a field has been set.

### SetLastModifiedTimeUsecsNil

`func (o *IndexedObjectSnapshot) SetLastModifiedTimeUsecsNil(b bool)`

 SetLastModifiedTimeUsecsNil sets the value for LastModifiedTimeUsecs to be an explicit nil

### UnsetLastModifiedTimeUsecs
`func (o *IndexedObjectSnapshot) UnsetLastModifiedTimeUsecs()`

UnsetLastModifiedTimeUsecs ensures that no value is present for LastModifiedTimeUsecs, not even an explicit nil
### GetExternalTargetInfo

`func (o *IndexedObjectSnapshot) GetExternalTargetInfo() ArchivalTargetSummaryInfo`

GetExternalTargetInfo returns the ExternalTargetInfo field if non-nil, zero value otherwise.

### GetExternalTargetInfoOk

`func (o *IndexedObjectSnapshot) GetExternalTargetInfoOk() (*ArchivalTargetSummaryInfo, bool)`

GetExternalTargetInfoOk returns a tuple with the ExternalTargetInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTargetInfo

`func (o *IndexedObjectSnapshot) SetExternalTargetInfo(v ArchivalTargetSummaryInfo)`

SetExternalTargetInfo sets ExternalTargetInfo field to given value.

### HasExternalTargetInfo

`func (o *IndexedObjectSnapshot) HasExternalTargetInfo() bool`

HasExternalTargetInfo returns a boolean if a field has been set.

### SetExternalTargetInfoNil

`func (o *IndexedObjectSnapshot) SetExternalTargetInfoNil(b bool)`

 SetExternalTargetInfoNil sets the value for ExternalTargetInfo to be an explicit nil

### UnsetExternalTargetInfo
`func (o *IndexedObjectSnapshot) UnsetExternalTargetInfo()`

UnsetExternalTargetInfo ensures that no value is present for ExternalTargetInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


