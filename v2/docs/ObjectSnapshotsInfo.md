# ObjectSnapshotsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchivalSnapshotsInfo** | Pointer to [**[]ObjectArchivalSnapshotInfo**](ObjectArchivalSnapshotInfo.md) | Specifies the archival snapshots information. | [optional] 
**IndexingStatus** | Pointer to **NullableString** | Specifies the indexing status of objects in this snapshot.&lt;br&gt; &#39;InProgress&#39; indicates the indexing is in progress.&lt;br&gt; &#39;Done&#39; indicates indexing is done.&lt;br&gt; &#39;NoIndex&#39; indicates indexing is not applicable.&lt;br&gt; &#39;Error&#39; indicates indexing failed with error. | [optional] 
**LocalSnapshotInfo** | Pointer to [**NullableObjectSnapshotsInfoLocalSnapshotInfo**](ObjectSnapshotsInfoLocalSnapshotInfo.md) |  | [optional] 
**ProtectionGroupId** | Pointer to **NullableString** | Specifies id of the Protection Group. | [optional] 
**ProtectionGroupName** | Pointer to **NullableString** | Specifies name of the Protection Group. | [optional] 
**ProtectionRunEndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time of Protection Group Run in Unix timestamp epoch in microseconds. | [optional] 
**ProtectionRunId** | Pointer to **NullableString** | Specifies the id of Protection Group Run. | [optional] 
**ProtectionRunStartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of Protection Group Run in Unix timestamp epoch in microseconds. | [optional] 
**RunInstanceId** | Pointer to **NullableInt64** | Specifies the instance id of the protection run which create the snapshot. | [optional] 
**RunType** | Pointer to **NullableString** | Specifies the type of protection run created this snapshot. | [optional] 
**SourceGroupId** | Pointer to **NullableString** | Specifies the source protection group id in case of replication. | [optional] 
**StorageDomainId** | Pointer to **NullableInt64** | Specifies the Storage Domain id where the backup data of Object is present. | [optional] 
**StorageDomainName** | Pointer to **NullableString** | Specifies the name of Storage Domain id where the backup data of Object is present | [optional] 

## Methods

### NewObjectSnapshotsInfo

`func NewObjectSnapshotsInfo() *ObjectSnapshotsInfo`

NewObjectSnapshotsInfo instantiates a new ObjectSnapshotsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectSnapshotsInfoWithDefaults

`func NewObjectSnapshotsInfoWithDefaults() *ObjectSnapshotsInfo`

NewObjectSnapshotsInfoWithDefaults instantiates a new ObjectSnapshotsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchivalSnapshotsInfo

`func (o *ObjectSnapshotsInfo) GetArchivalSnapshotsInfo() []ObjectArchivalSnapshotInfo`

GetArchivalSnapshotsInfo returns the ArchivalSnapshotsInfo field if non-nil, zero value otherwise.

### GetArchivalSnapshotsInfoOk

`func (o *ObjectSnapshotsInfo) GetArchivalSnapshotsInfoOk() (*[]ObjectArchivalSnapshotInfo, bool)`

GetArchivalSnapshotsInfoOk returns a tuple with the ArchivalSnapshotsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalSnapshotsInfo

`func (o *ObjectSnapshotsInfo) SetArchivalSnapshotsInfo(v []ObjectArchivalSnapshotInfo)`

SetArchivalSnapshotsInfo sets ArchivalSnapshotsInfo field to given value.

### HasArchivalSnapshotsInfo

`func (o *ObjectSnapshotsInfo) HasArchivalSnapshotsInfo() bool`

HasArchivalSnapshotsInfo returns a boolean if a field has been set.

### SetArchivalSnapshotsInfoNil

`func (o *ObjectSnapshotsInfo) SetArchivalSnapshotsInfoNil(b bool)`

 SetArchivalSnapshotsInfoNil sets the value for ArchivalSnapshotsInfo to be an explicit nil

### UnsetArchivalSnapshotsInfo
`func (o *ObjectSnapshotsInfo) UnsetArchivalSnapshotsInfo()`

UnsetArchivalSnapshotsInfo ensures that no value is present for ArchivalSnapshotsInfo, not even an explicit nil
### GetIndexingStatus

`func (o *ObjectSnapshotsInfo) GetIndexingStatus() string`

GetIndexingStatus returns the IndexingStatus field if non-nil, zero value otherwise.

### GetIndexingStatusOk

`func (o *ObjectSnapshotsInfo) GetIndexingStatusOk() (*string, bool)`

GetIndexingStatusOk returns a tuple with the IndexingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingStatus

`func (o *ObjectSnapshotsInfo) SetIndexingStatus(v string)`

SetIndexingStatus sets IndexingStatus field to given value.

### HasIndexingStatus

`func (o *ObjectSnapshotsInfo) HasIndexingStatus() bool`

HasIndexingStatus returns a boolean if a field has been set.

### SetIndexingStatusNil

`func (o *ObjectSnapshotsInfo) SetIndexingStatusNil(b bool)`

 SetIndexingStatusNil sets the value for IndexingStatus to be an explicit nil

### UnsetIndexingStatus
`func (o *ObjectSnapshotsInfo) UnsetIndexingStatus()`

UnsetIndexingStatus ensures that no value is present for IndexingStatus, not even an explicit nil
### GetLocalSnapshotInfo

`func (o *ObjectSnapshotsInfo) GetLocalSnapshotInfo() ObjectSnapshotsInfoLocalSnapshotInfo`

GetLocalSnapshotInfo returns the LocalSnapshotInfo field if non-nil, zero value otherwise.

### GetLocalSnapshotInfoOk

`func (o *ObjectSnapshotsInfo) GetLocalSnapshotInfoOk() (*ObjectSnapshotsInfoLocalSnapshotInfo, bool)`

GetLocalSnapshotInfoOk returns a tuple with the LocalSnapshotInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSnapshotInfo

`func (o *ObjectSnapshotsInfo) SetLocalSnapshotInfo(v ObjectSnapshotsInfoLocalSnapshotInfo)`

SetLocalSnapshotInfo sets LocalSnapshotInfo field to given value.

### HasLocalSnapshotInfo

`func (o *ObjectSnapshotsInfo) HasLocalSnapshotInfo() bool`

HasLocalSnapshotInfo returns a boolean if a field has been set.

### SetLocalSnapshotInfoNil

`func (o *ObjectSnapshotsInfo) SetLocalSnapshotInfoNil(b bool)`

 SetLocalSnapshotInfoNil sets the value for LocalSnapshotInfo to be an explicit nil

### UnsetLocalSnapshotInfo
`func (o *ObjectSnapshotsInfo) UnsetLocalSnapshotInfo()`

UnsetLocalSnapshotInfo ensures that no value is present for LocalSnapshotInfo, not even an explicit nil
### GetProtectionGroupId

`func (o *ObjectSnapshotsInfo) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *ObjectSnapshotsInfo) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *ObjectSnapshotsInfo) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *ObjectSnapshotsInfo) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *ObjectSnapshotsInfo) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *ObjectSnapshotsInfo) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetProtectionGroupName

`func (o *ObjectSnapshotsInfo) GetProtectionGroupName() string`

GetProtectionGroupName returns the ProtectionGroupName field if non-nil, zero value otherwise.

### GetProtectionGroupNameOk

`func (o *ObjectSnapshotsInfo) GetProtectionGroupNameOk() (*string, bool)`

GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupName

`func (o *ObjectSnapshotsInfo) SetProtectionGroupName(v string)`

SetProtectionGroupName sets ProtectionGroupName field to given value.

### HasProtectionGroupName

`func (o *ObjectSnapshotsInfo) HasProtectionGroupName() bool`

HasProtectionGroupName returns a boolean if a field has been set.

### SetProtectionGroupNameNil

`func (o *ObjectSnapshotsInfo) SetProtectionGroupNameNil(b bool)`

 SetProtectionGroupNameNil sets the value for ProtectionGroupName to be an explicit nil

### UnsetProtectionGroupName
`func (o *ObjectSnapshotsInfo) UnsetProtectionGroupName()`

UnsetProtectionGroupName ensures that no value is present for ProtectionGroupName, not even an explicit nil
### GetProtectionRunEndTimeUsecs

`func (o *ObjectSnapshotsInfo) GetProtectionRunEndTimeUsecs() int64`

GetProtectionRunEndTimeUsecs returns the ProtectionRunEndTimeUsecs field if non-nil, zero value otherwise.

### GetProtectionRunEndTimeUsecsOk

`func (o *ObjectSnapshotsInfo) GetProtectionRunEndTimeUsecsOk() (*int64, bool)`

GetProtectionRunEndTimeUsecsOk returns a tuple with the ProtectionRunEndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionRunEndTimeUsecs

`func (o *ObjectSnapshotsInfo) SetProtectionRunEndTimeUsecs(v int64)`

SetProtectionRunEndTimeUsecs sets ProtectionRunEndTimeUsecs field to given value.

### HasProtectionRunEndTimeUsecs

`func (o *ObjectSnapshotsInfo) HasProtectionRunEndTimeUsecs() bool`

HasProtectionRunEndTimeUsecs returns a boolean if a field has been set.

### SetProtectionRunEndTimeUsecsNil

`func (o *ObjectSnapshotsInfo) SetProtectionRunEndTimeUsecsNil(b bool)`

 SetProtectionRunEndTimeUsecsNil sets the value for ProtectionRunEndTimeUsecs to be an explicit nil

### UnsetProtectionRunEndTimeUsecs
`func (o *ObjectSnapshotsInfo) UnsetProtectionRunEndTimeUsecs()`

UnsetProtectionRunEndTimeUsecs ensures that no value is present for ProtectionRunEndTimeUsecs, not even an explicit nil
### GetProtectionRunId

`func (o *ObjectSnapshotsInfo) GetProtectionRunId() string`

GetProtectionRunId returns the ProtectionRunId field if non-nil, zero value otherwise.

### GetProtectionRunIdOk

`func (o *ObjectSnapshotsInfo) GetProtectionRunIdOk() (*string, bool)`

GetProtectionRunIdOk returns a tuple with the ProtectionRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionRunId

`func (o *ObjectSnapshotsInfo) SetProtectionRunId(v string)`

SetProtectionRunId sets ProtectionRunId field to given value.

### HasProtectionRunId

`func (o *ObjectSnapshotsInfo) HasProtectionRunId() bool`

HasProtectionRunId returns a boolean if a field has been set.

### SetProtectionRunIdNil

`func (o *ObjectSnapshotsInfo) SetProtectionRunIdNil(b bool)`

 SetProtectionRunIdNil sets the value for ProtectionRunId to be an explicit nil

### UnsetProtectionRunId
`func (o *ObjectSnapshotsInfo) UnsetProtectionRunId()`

UnsetProtectionRunId ensures that no value is present for ProtectionRunId, not even an explicit nil
### GetProtectionRunStartTimeUsecs

`func (o *ObjectSnapshotsInfo) GetProtectionRunStartTimeUsecs() int64`

GetProtectionRunStartTimeUsecs returns the ProtectionRunStartTimeUsecs field if non-nil, zero value otherwise.

### GetProtectionRunStartTimeUsecsOk

`func (o *ObjectSnapshotsInfo) GetProtectionRunStartTimeUsecsOk() (*int64, bool)`

GetProtectionRunStartTimeUsecsOk returns a tuple with the ProtectionRunStartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionRunStartTimeUsecs

`func (o *ObjectSnapshotsInfo) SetProtectionRunStartTimeUsecs(v int64)`

SetProtectionRunStartTimeUsecs sets ProtectionRunStartTimeUsecs field to given value.

### HasProtectionRunStartTimeUsecs

`func (o *ObjectSnapshotsInfo) HasProtectionRunStartTimeUsecs() bool`

HasProtectionRunStartTimeUsecs returns a boolean if a field has been set.

### SetProtectionRunStartTimeUsecsNil

`func (o *ObjectSnapshotsInfo) SetProtectionRunStartTimeUsecsNil(b bool)`

 SetProtectionRunStartTimeUsecsNil sets the value for ProtectionRunStartTimeUsecs to be an explicit nil

### UnsetProtectionRunStartTimeUsecs
`func (o *ObjectSnapshotsInfo) UnsetProtectionRunStartTimeUsecs()`

UnsetProtectionRunStartTimeUsecs ensures that no value is present for ProtectionRunStartTimeUsecs, not even an explicit nil
### GetRunInstanceId

`func (o *ObjectSnapshotsInfo) GetRunInstanceId() int64`

GetRunInstanceId returns the RunInstanceId field if non-nil, zero value otherwise.

### GetRunInstanceIdOk

`func (o *ObjectSnapshotsInfo) GetRunInstanceIdOk() (*int64, bool)`

GetRunInstanceIdOk returns a tuple with the RunInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunInstanceId

`func (o *ObjectSnapshotsInfo) SetRunInstanceId(v int64)`

SetRunInstanceId sets RunInstanceId field to given value.

### HasRunInstanceId

`func (o *ObjectSnapshotsInfo) HasRunInstanceId() bool`

HasRunInstanceId returns a boolean if a field has been set.

### SetRunInstanceIdNil

`func (o *ObjectSnapshotsInfo) SetRunInstanceIdNil(b bool)`

 SetRunInstanceIdNil sets the value for RunInstanceId to be an explicit nil

### UnsetRunInstanceId
`func (o *ObjectSnapshotsInfo) UnsetRunInstanceId()`

UnsetRunInstanceId ensures that no value is present for RunInstanceId, not even an explicit nil
### GetRunType

`func (o *ObjectSnapshotsInfo) GetRunType() string`

GetRunType returns the RunType field if non-nil, zero value otherwise.

### GetRunTypeOk

`func (o *ObjectSnapshotsInfo) GetRunTypeOk() (*string, bool)`

GetRunTypeOk returns a tuple with the RunType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunType

`func (o *ObjectSnapshotsInfo) SetRunType(v string)`

SetRunType sets RunType field to given value.

### HasRunType

`func (o *ObjectSnapshotsInfo) HasRunType() bool`

HasRunType returns a boolean if a field has been set.

### SetRunTypeNil

`func (o *ObjectSnapshotsInfo) SetRunTypeNil(b bool)`

 SetRunTypeNil sets the value for RunType to be an explicit nil

### UnsetRunType
`func (o *ObjectSnapshotsInfo) UnsetRunType()`

UnsetRunType ensures that no value is present for RunType, not even an explicit nil
### GetSourceGroupId

`func (o *ObjectSnapshotsInfo) GetSourceGroupId() string`

GetSourceGroupId returns the SourceGroupId field if non-nil, zero value otherwise.

### GetSourceGroupIdOk

`func (o *ObjectSnapshotsInfo) GetSourceGroupIdOk() (*string, bool)`

GetSourceGroupIdOk returns a tuple with the SourceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroupId

`func (o *ObjectSnapshotsInfo) SetSourceGroupId(v string)`

SetSourceGroupId sets SourceGroupId field to given value.

### HasSourceGroupId

`func (o *ObjectSnapshotsInfo) HasSourceGroupId() bool`

HasSourceGroupId returns a boolean if a field has been set.

### SetSourceGroupIdNil

`func (o *ObjectSnapshotsInfo) SetSourceGroupIdNil(b bool)`

 SetSourceGroupIdNil sets the value for SourceGroupId to be an explicit nil

### UnsetSourceGroupId
`func (o *ObjectSnapshotsInfo) UnsetSourceGroupId()`

UnsetSourceGroupId ensures that no value is present for SourceGroupId, not even an explicit nil
### GetStorageDomainId

`func (o *ObjectSnapshotsInfo) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *ObjectSnapshotsInfo) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *ObjectSnapshotsInfo) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.

### HasStorageDomainId

`func (o *ObjectSnapshotsInfo) HasStorageDomainId() bool`

HasStorageDomainId returns a boolean if a field has been set.

### SetStorageDomainIdNil

`func (o *ObjectSnapshotsInfo) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *ObjectSnapshotsInfo) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
### GetStorageDomainName

`func (o *ObjectSnapshotsInfo) GetStorageDomainName() string`

GetStorageDomainName returns the StorageDomainName field if non-nil, zero value otherwise.

### GetStorageDomainNameOk

`func (o *ObjectSnapshotsInfo) GetStorageDomainNameOk() (*string, bool)`

GetStorageDomainNameOk returns a tuple with the StorageDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainName

`func (o *ObjectSnapshotsInfo) SetStorageDomainName(v string)`

SetStorageDomainName sets StorageDomainName field to given value.

### HasStorageDomainName

`func (o *ObjectSnapshotsInfo) HasStorageDomainName() bool`

HasStorageDomainName returns a boolean if a field has been set.

### SetStorageDomainNameNil

`func (o *ObjectSnapshotsInfo) SetStorageDomainNameNil(b bool)`

 SetStorageDomainNameNil sets the value for StorageDomainName to be an explicit nil

### UnsetStorageDomainName
`func (o *ObjectSnapshotsInfo) UnsetStorageDomainName()`

UnsetStorageDomainName ensures that no value is present for StorageDomainName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


