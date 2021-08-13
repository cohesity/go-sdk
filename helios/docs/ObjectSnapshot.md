# ObjectSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Specifies the id of the snapshot. | [optional] 
**SnapshotTargetType** | Pointer to **NullableString** | Specifies the target type where the Object&#39;s snapshot resides. | [optional] 
**IndexingStatus** | Pointer to **NullableString** | Specifies the indexing status of objects in this snapshot.&lt;br&gt; &#39;InProgress&#39; indicates the indexing is in progress.&lt;br&gt; &#39;Done&#39; indicates indexing is done.&lt;br&gt; &#39;NoIndex&#39; indicates indexing is not applicable.&lt;br&gt; &#39;Error&#39; indicates indexing failed with error. | [optional] 
**ProtectionGroupId** | Pointer to **NullableString** | Specifies id of the Protection Group. | [optional] 
**ProtectionGroupName** | Pointer to **NullableString** | Specifies name of the Protection Group. | [optional] 
**ProtectionGroupRunId** | Pointer to **NullableString** | Specifies id of the Protection Group Run. | [optional] 
**RunInstanceId** | Pointer to **NullableInt64** | Specifies the instance id of the protection run which create the snapshot. | [optional] 
**RunStartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of the run in micro seconds. | [optional] 
**SourceGroupId** | Pointer to **NullableString** | Specifies the source protection group id in case of replication. | [optional] 
**RunType** | Pointer to **NullableString** | Specifies the type of protection run created this snapshot. | [optional] 
**Environment** | Pointer to **NullableString** | Specifies the snapshot environment. | [optional] 
**SnapshotTimestampUsecs** | Pointer to **NullableInt64** | Specifies the timestamp in Unix time epoch in microseconds when the snapshot is taken for the specified Object. | [optional] 
**ExpiryTimeUsecs** | Pointer to **NullableInt64** | Specifies the expiry time of the snapshot in Unix timestamp epoch in microseconds. If the snapshot has no expiry, this property will not be set. | [optional] 
**ExternalTargetInfo** | Pointer to [**NullableArchivalTargetSummaryInfo**](ArchivalTargetSummaryInfo.md) | Specifies the external target information if this is an archival snapshot. | [optional] 
**StorageDomainId** | Pointer to **NullableInt64** | Specifies the Storage Domain id where the snapshot of object is present. | [optional] 
**HasDataLock** | Pointer to **NullableBool** | Specifies if this snapshot has datalock. | [optional] 
**OnLegalHold** | Pointer to **NullableBool** | Specifies if this snapshot is on legalhold. | [optional] 
**ObjectId** | Pointer to **NullableInt64** | Specifies the object id which the snapshot is taken from. | [optional] 
**ObjectName** | Pointer to **NullableString** | Specifies the object name which the snapshot is taken from. | [optional] 
**SourceId** | Pointer to **NullableInt64** | Specifies the object source id which the snapshot is taken from. | [optional] 
**RegionId** | Pointer to **NullableString** | Specifies the region id where this snapshot belongs to. | [optional] 
**ClusterId** | Pointer to **NullableInt64** | Specifies the cluster id where this snapshot belongs to. | [optional] 
**ClusterIncarnationId** | Pointer to **NullableInt64** | Specifies the cluster incarnation id where this snapshot belongs to. | [optional] 
**PhysicalParams** | Pointer to [**PhysicalSnapshotParams**](PhysicalSnapshotParams.md) | Specifies the parameters specific to Physical type snapshot. | [optional] 
**HypervParams** | Pointer to [**HypervSnapshotParams**](HypervSnapshotParams.md) | Specifies the parameters specific to HyperV type snapshot. | [optional] 
**AwsParams** | Pointer to [**AwsSnapshotParams**](AwsSnapshotParams.md) | Specifies the parameters specific to AWS type snapshot. | [optional] 
**AzureParams** | Pointer to [**AzureSnapshotParams**](AzureSnapshotParams.md) | Specifies the parameters specific to Azure type snapshot. | [optional] 
**NetappParams** | Pointer to [**NetappObjectParams**](NetappObjectParams.md) | Specifies the parameters specific to NetApp type snapshot. | [optional] 
**IsilonParams** | Pointer to [**IsilonObjectParams**](IsilonObjectParams.md) | Specifies the parameters specific to Isilon type snapshot. | [optional] 
**GpfsParams** | Pointer to [**CommonNasObjectParams**](CommonNasObjectParams.md) | Specifies the parameters specific to GPFS type snapshot. | [optional] 
**FlashbladeParams** | Pointer to [**FlashbladeObjectParams**](FlashbladeObjectParams.md) | Specifies the parameters specific to Flashblade type snapshot. | [optional] 
**GenericNasParams** | Pointer to [**CommonNasObjectParams**](CommonNasObjectParams.md) | Specifies the parameters specific to Generic NAS type snapshot. | [optional] 
**ElastifileParams** | Pointer to [**CommonNasObjectParams**](CommonNasObjectParams.md) | Specifies the parameters specific to NetApp type snapshot. | [optional] 

## Methods

### NewObjectSnapshot

`func NewObjectSnapshot() *ObjectSnapshot`

NewObjectSnapshot instantiates a new ObjectSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectSnapshotWithDefaults

`func NewObjectSnapshotWithDefaults() *ObjectSnapshot`

NewObjectSnapshotWithDefaults instantiates a new ObjectSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ObjectSnapshot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectSnapshot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectSnapshot) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ObjectSnapshot) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ObjectSnapshot) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ObjectSnapshot) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetSnapshotTargetType

`func (o *ObjectSnapshot) GetSnapshotTargetType() string`

GetSnapshotTargetType returns the SnapshotTargetType field if non-nil, zero value otherwise.

### GetSnapshotTargetTypeOk

`func (o *ObjectSnapshot) GetSnapshotTargetTypeOk() (*string, bool)`

GetSnapshotTargetTypeOk returns a tuple with the SnapshotTargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTargetType

`func (o *ObjectSnapshot) SetSnapshotTargetType(v string)`

SetSnapshotTargetType sets SnapshotTargetType field to given value.

### HasSnapshotTargetType

`func (o *ObjectSnapshot) HasSnapshotTargetType() bool`

HasSnapshotTargetType returns a boolean if a field has been set.

### SetSnapshotTargetTypeNil

`func (o *ObjectSnapshot) SetSnapshotTargetTypeNil(b bool)`

 SetSnapshotTargetTypeNil sets the value for SnapshotTargetType to be an explicit nil

### UnsetSnapshotTargetType
`func (o *ObjectSnapshot) UnsetSnapshotTargetType()`

UnsetSnapshotTargetType ensures that no value is present for SnapshotTargetType, not even an explicit nil
### GetIndexingStatus

`func (o *ObjectSnapshot) GetIndexingStatus() string`

GetIndexingStatus returns the IndexingStatus field if non-nil, zero value otherwise.

### GetIndexingStatusOk

`func (o *ObjectSnapshot) GetIndexingStatusOk() (*string, bool)`

GetIndexingStatusOk returns a tuple with the IndexingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingStatus

`func (o *ObjectSnapshot) SetIndexingStatus(v string)`

SetIndexingStatus sets IndexingStatus field to given value.

### HasIndexingStatus

`func (o *ObjectSnapshot) HasIndexingStatus() bool`

HasIndexingStatus returns a boolean if a field has been set.

### SetIndexingStatusNil

`func (o *ObjectSnapshot) SetIndexingStatusNil(b bool)`

 SetIndexingStatusNil sets the value for IndexingStatus to be an explicit nil

### UnsetIndexingStatus
`func (o *ObjectSnapshot) UnsetIndexingStatus()`

UnsetIndexingStatus ensures that no value is present for IndexingStatus, not even an explicit nil
### GetProtectionGroupId

`func (o *ObjectSnapshot) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *ObjectSnapshot) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *ObjectSnapshot) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *ObjectSnapshot) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *ObjectSnapshot) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *ObjectSnapshot) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetProtectionGroupName

`func (o *ObjectSnapshot) GetProtectionGroupName() string`

GetProtectionGroupName returns the ProtectionGroupName field if non-nil, zero value otherwise.

### GetProtectionGroupNameOk

`func (o *ObjectSnapshot) GetProtectionGroupNameOk() (*string, bool)`

GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupName

`func (o *ObjectSnapshot) SetProtectionGroupName(v string)`

SetProtectionGroupName sets ProtectionGroupName field to given value.

### HasProtectionGroupName

`func (o *ObjectSnapshot) HasProtectionGroupName() bool`

HasProtectionGroupName returns a boolean if a field has been set.

### SetProtectionGroupNameNil

`func (o *ObjectSnapshot) SetProtectionGroupNameNil(b bool)`

 SetProtectionGroupNameNil sets the value for ProtectionGroupName to be an explicit nil

### UnsetProtectionGroupName
`func (o *ObjectSnapshot) UnsetProtectionGroupName()`

UnsetProtectionGroupName ensures that no value is present for ProtectionGroupName, not even an explicit nil
### GetProtectionGroupRunId

`func (o *ObjectSnapshot) GetProtectionGroupRunId() string`

GetProtectionGroupRunId returns the ProtectionGroupRunId field if non-nil, zero value otherwise.

### GetProtectionGroupRunIdOk

`func (o *ObjectSnapshot) GetProtectionGroupRunIdOk() (*string, bool)`

GetProtectionGroupRunIdOk returns a tuple with the ProtectionGroupRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupRunId

`func (o *ObjectSnapshot) SetProtectionGroupRunId(v string)`

SetProtectionGroupRunId sets ProtectionGroupRunId field to given value.

### HasProtectionGroupRunId

`func (o *ObjectSnapshot) HasProtectionGroupRunId() bool`

HasProtectionGroupRunId returns a boolean if a field has been set.

### SetProtectionGroupRunIdNil

`func (o *ObjectSnapshot) SetProtectionGroupRunIdNil(b bool)`

 SetProtectionGroupRunIdNil sets the value for ProtectionGroupRunId to be an explicit nil

### UnsetProtectionGroupRunId
`func (o *ObjectSnapshot) UnsetProtectionGroupRunId()`

UnsetProtectionGroupRunId ensures that no value is present for ProtectionGroupRunId, not even an explicit nil
### GetRunInstanceId

`func (o *ObjectSnapshot) GetRunInstanceId() int64`

GetRunInstanceId returns the RunInstanceId field if non-nil, zero value otherwise.

### GetRunInstanceIdOk

`func (o *ObjectSnapshot) GetRunInstanceIdOk() (*int64, bool)`

GetRunInstanceIdOk returns a tuple with the RunInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunInstanceId

`func (o *ObjectSnapshot) SetRunInstanceId(v int64)`

SetRunInstanceId sets RunInstanceId field to given value.

### HasRunInstanceId

`func (o *ObjectSnapshot) HasRunInstanceId() bool`

HasRunInstanceId returns a boolean if a field has been set.

### SetRunInstanceIdNil

`func (o *ObjectSnapshot) SetRunInstanceIdNil(b bool)`

 SetRunInstanceIdNil sets the value for RunInstanceId to be an explicit nil

### UnsetRunInstanceId
`func (o *ObjectSnapshot) UnsetRunInstanceId()`

UnsetRunInstanceId ensures that no value is present for RunInstanceId, not even an explicit nil
### GetRunStartTimeUsecs

`func (o *ObjectSnapshot) GetRunStartTimeUsecs() int64`

GetRunStartTimeUsecs returns the RunStartTimeUsecs field if non-nil, zero value otherwise.

### GetRunStartTimeUsecsOk

`func (o *ObjectSnapshot) GetRunStartTimeUsecsOk() (*int64, bool)`

GetRunStartTimeUsecsOk returns a tuple with the RunStartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunStartTimeUsecs

`func (o *ObjectSnapshot) SetRunStartTimeUsecs(v int64)`

SetRunStartTimeUsecs sets RunStartTimeUsecs field to given value.

### HasRunStartTimeUsecs

`func (o *ObjectSnapshot) HasRunStartTimeUsecs() bool`

HasRunStartTimeUsecs returns a boolean if a field has been set.

### SetRunStartTimeUsecsNil

`func (o *ObjectSnapshot) SetRunStartTimeUsecsNil(b bool)`

 SetRunStartTimeUsecsNil sets the value for RunStartTimeUsecs to be an explicit nil

### UnsetRunStartTimeUsecs
`func (o *ObjectSnapshot) UnsetRunStartTimeUsecs()`

UnsetRunStartTimeUsecs ensures that no value is present for RunStartTimeUsecs, not even an explicit nil
### GetSourceGroupId

`func (o *ObjectSnapshot) GetSourceGroupId() string`

GetSourceGroupId returns the SourceGroupId field if non-nil, zero value otherwise.

### GetSourceGroupIdOk

`func (o *ObjectSnapshot) GetSourceGroupIdOk() (*string, bool)`

GetSourceGroupIdOk returns a tuple with the SourceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroupId

`func (o *ObjectSnapshot) SetSourceGroupId(v string)`

SetSourceGroupId sets SourceGroupId field to given value.

### HasSourceGroupId

`func (o *ObjectSnapshot) HasSourceGroupId() bool`

HasSourceGroupId returns a boolean if a field has been set.

### SetSourceGroupIdNil

`func (o *ObjectSnapshot) SetSourceGroupIdNil(b bool)`

 SetSourceGroupIdNil sets the value for SourceGroupId to be an explicit nil

### UnsetSourceGroupId
`func (o *ObjectSnapshot) UnsetSourceGroupId()`

UnsetSourceGroupId ensures that no value is present for SourceGroupId, not even an explicit nil
### GetRunType

`func (o *ObjectSnapshot) GetRunType() string`

GetRunType returns the RunType field if non-nil, zero value otherwise.

### GetRunTypeOk

`func (o *ObjectSnapshot) GetRunTypeOk() (*string, bool)`

GetRunTypeOk returns a tuple with the RunType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunType

`func (o *ObjectSnapshot) SetRunType(v string)`

SetRunType sets RunType field to given value.

### HasRunType

`func (o *ObjectSnapshot) HasRunType() bool`

HasRunType returns a boolean if a field has been set.

### SetRunTypeNil

`func (o *ObjectSnapshot) SetRunTypeNil(b bool)`

 SetRunTypeNil sets the value for RunType to be an explicit nil

### UnsetRunType
`func (o *ObjectSnapshot) UnsetRunType()`

UnsetRunType ensures that no value is present for RunType, not even an explicit nil
### GetEnvironment

`func (o *ObjectSnapshot) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ObjectSnapshot) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ObjectSnapshot) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ObjectSnapshot) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *ObjectSnapshot) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *ObjectSnapshot) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetSnapshotTimestampUsecs

`func (o *ObjectSnapshot) GetSnapshotTimestampUsecs() int64`

GetSnapshotTimestampUsecs returns the SnapshotTimestampUsecs field if non-nil, zero value otherwise.

### GetSnapshotTimestampUsecsOk

`func (o *ObjectSnapshot) GetSnapshotTimestampUsecsOk() (*int64, bool)`

GetSnapshotTimestampUsecsOk returns a tuple with the SnapshotTimestampUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTimestampUsecs

`func (o *ObjectSnapshot) SetSnapshotTimestampUsecs(v int64)`

SetSnapshotTimestampUsecs sets SnapshotTimestampUsecs field to given value.

### HasSnapshotTimestampUsecs

`func (o *ObjectSnapshot) HasSnapshotTimestampUsecs() bool`

HasSnapshotTimestampUsecs returns a boolean if a field has been set.

### SetSnapshotTimestampUsecsNil

`func (o *ObjectSnapshot) SetSnapshotTimestampUsecsNil(b bool)`

 SetSnapshotTimestampUsecsNil sets the value for SnapshotTimestampUsecs to be an explicit nil

### UnsetSnapshotTimestampUsecs
`func (o *ObjectSnapshot) UnsetSnapshotTimestampUsecs()`

UnsetSnapshotTimestampUsecs ensures that no value is present for SnapshotTimestampUsecs, not even an explicit nil
### GetExpiryTimeUsecs

`func (o *ObjectSnapshot) GetExpiryTimeUsecs() int64`

GetExpiryTimeUsecs returns the ExpiryTimeUsecs field if non-nil, zero value otherwise.

### GetExpiryTimeUsecsOk

`func (o *ObjectSnapshot) GetExpiryTimeUsecsOk() (*int64, bool)`

GetExpiryTimeUsecsOk returns a tuple with the ExpiryTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimeUsecs

`func (o *ObjectSnapshot) SetExpiryTimeUsecs(v int64)`

SetExpiryTimeUsecs sets ExpiryTimeUsecs field to given value.

### HasExpiryTimeUsecs

`func (o *ObjectSnapshot) HasExpiryTimeUsecs() bool`

HasExpiryTimeUsecs returns a boolean if a field has been set.

### SetExpiryTimeUsecsNil

`func (o *ObjectSnapshot) SetExpiryTimeUsecsNil(b bool)`

 SetExpiryTimeUsecsNil sets the value for ExpiryTimeUsecs to be an explicit nil

### UnsetExpiryTimeUsecs
`func (o *ObjectSnapshot) UnsetExpiryTimeUsecs()`

UnsetExpiryTimeUsecs ensures that no value is present for ExpiryTimeUsecs, not even an explicit nil
### GetExternalTargetInfo

`func (o *ObjectSnapshot) GetExternalTargetInfo() ArchivalTargetSummaryInfo`

GetExternalTargetInfo returns the ExternalTargetInfo field if non-nil, zero value otherwise.

### GetExternalTargetInfoOk

`func (o *ObjectSnapshot) GetExternalTargetInfoOk() (*ArchivalTargetSummaryInfo, bool)`

GetExternalTargetInfoOk returns a tuple with the ExternalTargetInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTargetInfo

`func (o *ObjectSnapshot) SetExternalTargetInfo(v ArchivalTargetSummaryInfo)`

SetExternalTargetInfo sets ExternalTargetInfo field to given value.

### HasExternalTargetInfo

`func (o *ObjectSnapshot) HasExternalTargetInfo() bool`

HasExternalTargetInfo returns a boolean if a field has been set.

### SetExternalTargetInfoNil

`func (o *ObjectSnapshot) SetExternalTargetInfoNil(b bool)`

 SetExternalTargetInfoNil sets the value for ExternalTargetInfo to be an explicit nil

### UnsetExternalTargetInfo
`func (o *ObjectSnapshot) UnsetExternalTargetInfo()`

UnsetExternalTargetInfo ensures that no value is present for ExternalTargetInfo, not even an explicit nil
### GetStorageDomainId

`func (o *ObjectSnapshot) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *ObjectSnapshot) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *ObjectSnapshot) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.

### HasStorageDomainId

`func (o *ObjectSnapshot) HasStorageDomainId() bool`

HasStorageDomainId returns a boolean if a field has been set.

### SetStorageDomainIdNil

`func (o *ObjectSnapshot) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *ObjectSnapshot) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
### GetHasDataLock

`func (o *ObjectSnapshot) GetHasDataLock() bool`

GetHasDataLock returns the HasDataLock field if non-nil, zero value otherwise.

### GetHasDataLockOk

`func (o *ObjectSnapshot) GetHasDataLockOk() (*bool, bool)`

GetHasDataLockOk returns a tuple with the HasDataLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDataLock

`func (o *ObjectSnapshot) SetHasDataLock(v bool)`

SetHasDataLock sets HasDataLock field to given value.

### HasHasDataLock

`func (o *ObjectSnapshot) HasHasDataLock() bool`

HasHasDataLock returns a boolean if a field has been set.

### SetHasDataLockNil

`func (o *ObjectSnapshot) SetHasDataLockNil(b bool)`

 SetHasDataLockNil sets the value for HasDataLock to be an explicit nil

### UnsetHasDataLock
`func (o *ObjectSnapshot) UnsetHasDataLock()`

UnsetHasDataLock ensures that no value is present for HasDataLock, not even an explicit nil
### GetOnLegalHold

`func (o *ObjectSnapshot) GetOnLegalHold() bool`

GetOnLegalHold returns the OnLegalHold field if non-nil, zero value otherwise.

### GetOnLegalHoldOk

`func (o *ObjectSnapshot) GetOnLegalHoldOk() (*bool, bool)`

GetOnLegalHoldOk returns a tuple with the OnLegalHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnLegalHold

`func (o *ObjectSnapshot) SetOnLegalHold(v bool)`

SetOnLegalHold sets OnLegalHold field to given value.

### HasOnLegalHold

`func (o *ObjectSnapshot) HasOnLegalHold() bool`

HasOnLegalHold returns a boolean if a field has been set.

### SetOnLegalHoldNil

`func (o *ObjectSnapshot) SetOnLegalHoldNil(b bool)`

 SetOnLegalHoldNil sets the value for OnLegalHold to be an explicit nil

### UnsetOnLegalHold
`func (o *ObjectSnapshot) UnsetOnLegalHold()`

UnsetOnLegalHold ensures that no value is present for OnLegalHold, not even an explicit nil
### GetObjectId

`func (o *ObjectSnapshot) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *ObjectSnapshot) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *ObjectSnapshot) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *ObjectSnapshot) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### SetObjectIdNil

`func (o *ObjectSnapshot) SetObjectIdNil(b bool)`

 SetObjectIdNil sets the value for ObjectId to be an explicit nil

### UnsetObjectId
`func (o *ObjectSnapshot) UnsetObjectId()`

UnsetObjectId ensures that no value is present for ObjectId, not even an explicit nil
### GetObjectName

`func (o *ObjectSnapshot) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ObjectSnapshot) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ObjectSnapshot) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ObjectSnapshot) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### SetObjectNameNil

`func (o *ObjectSnapshot) SetObjectNameNil(b bool)`

 SetObjectNameNil sets the value for ObjectName to be an explicit nil

### UnsetObjectName
`func (o *ObjectSnapshot) UnsetObjectName()`

UnsetObjectName ensures that no value is present for ObjectName, not even an explicit nil
### GetSourceId

`func (o *ObjectSnapshot) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ObjectSnapshot) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ObjectSnapshot) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *ObjectSnapshot) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *ObjectSnapshot) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *ObjectSnapshot) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetRegionId

`func (o *ObjectSnapshot) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *ObjectSnapshot) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *ObjectSnapshot) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *ObjectSnapshot) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### SetRegionIdNil

`func (o *ObjectSnapshot) SetRegionIdNil(b bool)`

 SetRegionIdNil sets the value for RegionId to be an explicit nil

### UnsetRegionId
`func (o *ObjectSnapshot) UnsetRegionId()`

UnsetRegionId ensures that no value is present for RegionId, not even an explicit nil
### GetClusterId

`func (o *ObjectSnapshot) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ObjectSnapshot) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ObjectSnapshot) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *ObjectSnapshot) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *ObjectSnapshot) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *ObjectSnapshot) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterIncarnationId

`func (o *ObjectSnapshot) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *ObjectSnapshot) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *ObjectSnapshot) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.

### HasClusterIncarnationId

`func (o *ObjectSnapshot) HasClusterIncarnationId() bool`

HasClusterIncarnationId returns a boolean if a field has been set.

### SetClusterIncarnationIdNil

`func (o *ObjectSnapshot) SetClusterIncarnationIdNil(b bool)`

 SetClusterIncarnationIdNil sets the value for ClusterIncarnationId to be an explicit nil

### UnsetClusterIncarnationId
`func (o *ObjectSnapshot) UnsetClusterIncarnationId()`

UnsetClusterIncarnationId ensures that no value is present for ClusterIncarnationId, not even an explicit nil
### GetPhysicalParams

`func (o *ObjectSnapshot) GetPhysicalParams() PhysicalSnapshotParams`

GetPhysicalParams returns the PhysicalParams field if non-nil, zero value otherwise.

### GetPhysicalParamsOk

`func (o *ObjectSnapshot) GetPhysicalParamsOk() (*PhysicalSnapshotParams, bool)`

GetPhysicalParamsOk returns a tuple with the PhysicalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalParams

`func (o *ObjectSnapshot) SetPhysicalParams(v PhysicalSnapshotParams)`

SetPhysicalParams sets PhysicalParams field to given value.

### HasPhysicalParams

`func (o *ObjectSnapshot) HasPhysicalParams() bool`

HasPhysicalParams returns a boolean if a field has been set.

### GetHypervParams

`func (o *ObjectSnapshot) GetHypervParams() HypervSnapshotParams`

GetHypervParams returns the HypervParams field if non-nil, zero value otherwise.

### GetHypervParamsOk

`func (o *ObjectSnapshot) GetHypervParamsOk() (*HypervSnapshotParams, bool)`

GetHypervParamsOk returns a tuple with the HypervParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervParams

`func (o *ObjectSnapshot) SetHypervParams(v HypervSnapshotParams)`

SetHypervParams sets HypervParams field to given value.

### HasHypervParams

`func (o *ObjectSnapshot) HasHypervParams() bool`

HasHypervParams returns a boolean if a field has been set.

### GetAwsParams

`func (o *ObjectSnapshot) GetAwsParams() AwsSnapshotParams`

GetAwsParams returns the AwsParams field if non-nil, zero value otherwise.

### GetAwsParamsOk

`func (o *ObjectSnapshot) GetAwsParamsOk() (*AwsSnapshotParams, bool)`

GetAwsParamsOk returns a tuple with the AwsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsParams

`func (o *ObjectSnapshot) SetAwsParams(v AwsSnapshotParams)`

SetAwsParams sets AwsParams field to given value.

### HasAwsParams

`func (o *ObjectSnapshot) HasAwsParams() bool`

HasAwsParams returns a boolean if a field has been set.

### GetAzureParams

`func (o *ObjectSnapshot) GetAzureParams() AzureSnapshotParams`

GetAzureParams returns the AzureParams field if non-nil, zero value otherwise.

### GetAzureParamsOk

`func (o *ObjectSnapshot) GetAzureParamsOk() (*AzureSnapshotParams, bool)`

GetAzureParamsOk returns a tuple with the AzureParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureParams

`func (o *ObjectSnapshot) SetAzureParams(v AzureSnapshotParams)`

SetAzureParams sets AzureParams field to given value.

### HasAzureParams

`func (o *ObjectSnapshot) HasAzureParams() bool`

HasAzureParams returns a boolean if a field has been set.

### GetNetappParams

`func (o *ObjectSnapshot) GetNetappParams() NetappObjectParams`

GetNetappParams returns the NetappParams field if non-nil, zero value otherwise.

### GetNetappParamsOk

`func (o *ObjectSnapshot) GetNetappParamsOk() (*NetappObjectParams, bool)`

GetNetappParamsOk returns a tuple with the NetappParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappParams

`func (o *ObjectSnapshot) SetNetappParams(v NetappObjectParams)`

SetNetappParams sets NetappParams field to given value.

### HasNetappParams

`func (o *ObjectSnapshot) HasNetappParams() bool`

HasNetappParams returns a boolean if a field has been set.

### GetIsilonParams

`func (o *ObjectSnapshot) GetIsilonParams() IsilonObjectParams`

GetIsilonParams returns the IsilonParams field if non-nil, zero value otherwise.

### GetIsilonParamsOk

`func (o *ObjectSnapshot) GetIsilonParamsOk() (*IsilonObjectParams, bool)`

GetIsilonParamsOk returns a tuple with the IsilonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsilonParams

`func (o *ObjectSnapshot) SetIsilonParams(v IsilonObjectParams)`

SetIsilonParams sets IsilonParams field to given value.

### HasIsilonParams

`func (o *ObjectSnapshot) HasIsilonParams() bool`

HasIsilonParams returns a boolean if a field has been set.

### GetGpfsParams

`func (o *ObjectSnapshot) GetGpfsParams() CommonNasObjectParams`

GetGpfsParams returns the GpfsParams field if non-nil, zero value otherwise.

### GetGpfsParamsOk

`func (o *ObjectSnapshot) GetGpfsParamsOk() (*CommonNasObjectParams, bool)`

GetGpfsParamsOk returns a tuple with the GpfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpfsParams

`func (o *ObjectSnapshot) SetGpfsParams(v CommonNasObjectParams)`

SetGpfsParams sets GpfsParams field to given value.

### HasGpfsParams

`func (o *ObjectSnapshot) HasGpfsParams() bool`

HasGpfsParams returns a boolean if a field has been set.

### GetFlashbladeParams

`func (o *ObjectSnapshot) GetFlashbladeParams() FlashbladeObjectParams`

GetFlashbladeParams returns the FlashbladeParams field if non-nil, zero value otherwise.

### GetFlashbladeParamsOk

`func (o *ObjectSnapshot) GetFlashbladeParamsOk() (*FlashbladeObjectParams, bool)`

GetFlashbladeParamsOk returns a tuple with the FlashbladeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashbladeParams

`func (o *ObjectSnapshot) SetFlashbladeParams(v FlashbladeObjectParams)`

SetFlashbladeParams sets FlashbladeParams field to given value.

### HasFlashbladeParams

`func (o *ObjectSnapshot) HasFlashbladeParams() bool`

HasFlashbladeParams returns a boolean if a field has been set.

### GetGenericNasParams

`func (o *ObjectSnapshot) GetGenericNasParams() CommonNasObjectParams`

GetGenericNasParams returns the GenericNasParams field if non-nil, zero value otherwise.

### GetGenericNasParamsOk

`func (o *ObjectSnapshot) GetGenericNasParamsOk() (*CommonNasObjectParams, bool)`

GetGenericNasParamsOk returns a tuple with the GenericNasParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericNasParams

`func (o *ObjectSnapshot) SetGenericNasParams(v CommonNasObjectParams)`

SetGenericNasParams sets GenericNasParams field to given value.

### HasGenericNasParams

`func (o *ObjectSnapshot) HasGenericNasParams() bool`

HasGenericNasParams returns a boolean if a field has been set.

### GetElastifileParams

`func (o *ObjectSnapshot) GetElastifileParams() CommonNasObjectParams`

GetElastifileParams returns the ElastifileParams field if non-nil, zero value otherwise.

### GetElastifileParamsOk

`func (o *ObjectSnapshot) GetElastifileParamsOk() (*CommonNasObjectParams, bool)`

GetElastifileParamsOk returns a tuple with the ElastifileParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElastifileParams

`func (o *ObjectSnapshot) SetElastifileParams(v CommonNasObjectParams)`

SetElastifileParams sets ElastifileParams field to given value.

### HasElastifileParams

`func (o *ObjectSnapshot) HasElastifileParams() bool`

HasElastifileParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


