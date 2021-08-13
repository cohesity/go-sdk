# ObjectLastRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** | Specifies object id. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the object. | [optional] 
**SourceId** | Pointer to **NullableInt64** | Specifies registered source id to which object belongs. | [optional] 
**SourceName** | Pointer to **NullableString** | Specifies registered source name to which object belongs. | [optional] 
**Environment** | Pointer to **NullableString** | Specifies the environment of the object. | [optional] 
**ObjectHash** | Pointer to **NullableString** | Specifies the hash identifier of the object. | [optional] 
**ObjectType** | Pointer to **NullableString** | Specifies the type of the object. | [optional] 
**LogicalSizeBytes** | Pointer to **NullableInt64** | Specifies the logical size of object in bytes. | [optional] 
**Uuid** | Pointer to **NullableString** | Specifies the uuid which is a unique identifier of the object. | [optional] 
**GlobalId** | Pointer to **NullableString** | Specifies the global id which is a unique identifier of the object. | [optional] 
**ProtectionType** | Pointer to **NullableString** | Specifies the protection type of the object if any. | [optional] 
**OsType** | Pointer to **NullableString** | Specifies the operating system type of the object. | [optional] 
**VCenterSummary** | Pointer to [**ObjectTypeVCenterParams**](ObjectTypeVCenterParams.md) |  | [optional] 
**SharepointSiteSummary** | Pointer to [**SharepointObjectParams**](SharepointObjectParams.md) |  | [optional] 
**RunId** | Pointer to **string** | Specifies the last run id. | [optional] 
**ProtectionGroupName** | Pointer to **NullableString** | Specifies the protection group name of last run. | [optional] 
**ProtectionGroupId** | Pointer to **NullableString** | Specifies the protection group id of last run. | [optional] 
**PolicyName** | Pointer to **NullableString** | Specifies the policy name of last run. | [optional] 
**PolicyId** | Pointer to **NullableString** | Specifies the policy id of last run. | [optional] 
**BackupRunStatus** | Pointer to **NullableString** | Specifies the status of last local back up run. | [optional] 
**ArchivalRunStatus** | Pointer to **NullableString** | Specifies the status of last archival run. | [optional] 
**ReplicationRunStatus** | Pointer to **NullableString** | Specifies the status of last replication run. | [optional] 
**IsSlaViolated** | Pointer to **NullableBool** | Specifies if the sla is violated in last run. | [optional] 

## Methods

### NewObjectLastRun

`func NewObjectLastRun() *ObjectLastRun`

NewObjectLastRun instantiates a new ObjectLastRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectLastRunWithDefaults

`func NewObjectLastRunWithDefaults() *ObjectLastRun`

NewObjectLastRunWithDefaults instantiates a new ObjectLastRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ObjectLastRun) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectLastRun) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectLastRun) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ObjectLastRun) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ObjectLastRun) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ObjectLastRun) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ObjectLastRun) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectLastRun) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectLastRun) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ObjectLastRun) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ObjectLastRun) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ObjectLastRun) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSourceId

`func (o *ObjectLastRun) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ObjectLastRun) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ObjectLastRun) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *ObjectLastRun) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *ObjectLastRun) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *ObjectLastRun) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *ObjectLastRun) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *ObjectLastRun) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *ObjectLastRun) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *ObjectLastRun) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *ObjectLastRun) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *ObjectLastRun) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetEnvironment

`func (o *ObjectLastRun) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ObjectLastRun) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ObjectLastRun) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ObjectLastRun) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *ObjectLastRun) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *ObjectLastRun) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetObjectHash

`func (o *ObjectLastRun) GetObjectHash() string`

GetObjectHash returns the ObjectHash field if non-nil, zero value otherwise.

### GetObjectHashOk

`func (o *ObjectLastRun) GetObjectHashOk() (*string, bool)`

GetObjectHashOk returns a tuple with the ObjectHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectHash

`func (o *ObjectLastRun) SetObjectHash(v string)`

SetObjectHash sets ObjectHash field to given value.

### HasObjectHash

`func (o *ObjectLastRun) HasObjectHash() bool`

HasObjectHash returns a boolean if a field has been set.

### SetObjectHashNil

`func (o *ObjectLastRun) SetObjectHashNil(b bool)`

 SetObjectHashNil sets the value for ObjectHash to be an explicit nil

### UnsetObjectHash
`func (o *ObjectLastRun) UnsetObjectHash()`

UnsetObjectHash ensures that no value is present for ObjectHash, not even an explicit nil
### GetObjectType

`func (o *ObjectLastRun) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ObjectLastRun) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ObjectLastRun) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *ObjectLastRun) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### SetObjectTypeNil

`func (o *ObjectLastRun) SetObjectTypeNil(b bool)`

 SetObjectTypeNil sets the value for ObjectType to be an explicit nil

### UnsetObjectType
`func (o *ObjectLastRun) UnsetObjectType()`

UnsetObjectType ensures that no value is present for ObjectType, not even an explicit nil
### GetLogicalSizeBytes

`func (o *ObjectLastRun) GetLogicalSizeBytes() int64`

GetLogicalSizeBytes returns the LogicalSizeBytes field if non-nil, zero value otherwise.

### GetLogicalSizeBytesOk

`func (o *ObjectLastRun) GetLogicalSizeBytesOk() (*int64, bool)`

GetLogicalSizeBytesOk returns a tuple with the LogicalSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalSizeBytes

`func (o *ObjectLastRun) SetLogicalSizeBytes(v int64)`

SetLogicalSizeBytes sets LogicalSizeBytes field to given value.

### HasLogicalSizeBytes

`func (o *ObjectLastRun) HasLogicalSizeBytes() bool`

HasLogicalSizeBytes returns a boolean if a field has been set.

### SetLogicalSizeBytesNil

`func (o *ObjectLastRun) SetLogicalSizeBytesNil(b bool)`

 SetLogicalSizeBytesNil sets the value for LogicalSizeBytes to be an explicit nil

### UnsetLogicalSizeBytes
`func (o *ObjectLastRun) UnsetLogicalSizeBytes()`

UnsetLogicalSizeBytes ensures that no value is present for LogicalSizeBytes, not even an explicit nil
### GetUuid

`func (o *ObjectLastRun) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ObjectLastRun) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ObjectLastRun) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ObjectLastRun) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *ObjectLastRun) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *ObjectLastRun) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetGlobalId

`func (o *ObjectLastRun) GetGlobalId() string`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *ObjectLastRun) GetGlobalIdOk() (*string, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *ObjectLastRun) SetGlobalId(v string)`

SetGlobalId sets GlobalId field to given value.

### HasGlobalId

`func (o *ObjectLastRun) HasGlobalId() bool`

HasGlobalId returns a boolean if a field has been set.

### SetGlobalIdNil

`func (o *ObjectLastRun) SetGlobalIdNil(b bool)`

 SetGlobalIdNil sets the value for GlobalId to be an explicit nil

### UnsetGlobalId
`func (o *ObjectLastRun) UnsetGlobalId()`

UnsetGlobalId ensures that no value is present for GlobalId, not even an explicit nil
### GetProtectionType

`func (o *ObjectLastRun) GetProtectionType() string`

GetProtectionType returns the ProtectionType field if non-nil, zero value otherwise.

### GetProtectionTypeOk

`func (o *ObjectLastRun) GetProtectionTypeOk() (*string, bool)`

GetProtectionTypeOk returns a tuple with the ProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionType

`func (o *ObjectLastRun) SetProtectionType(v string)`

SetProtectionType sets ProtectionType field to given value.

### HasProtectionType

`func (o *ObjectLastRun) HasProtectionType() bool`

HasProtectionType returns a boolean if a field has been set.

### SetProtectionTypeNil

`func (o *ObjectLastRun) SetProtectionTypeNil(b bool)`

 SetProtectionTypeNil sets the value for ProtectionType to be an explicit nil

### UnsetProtectionType
`func (o *ObjectLastRun) UnsetProtectionType()`

UnsetProtectionType ensures that no value is present for ProtectionType, not even an explicit nil
### GetOsType

`func (o *ObjectLastRun) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *ObjectLastRun) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *ObjectLastRun) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *ObjectLastRun) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### SetOsTypeNil

`func (o *ObjectLastRun) SetOsTypeNil(b bool)`

 SetOsTypeNil sets the value for OsType to be an explicit nil

### UnsetOsType
`func (o *ObjectLastRun) UnsetOsType()`

UnsetOsType ensures that no value is present for OsType, not even an explicit nil
### GetVCenterSummary

`func (o *ObjectLastRun) GetVCenterSummary() ObjectTypeVCenterParams`

GetVCenterSummary returns the VCenterSummary field if non-nil, zero value otherwise.

### GetVCenterSummaryOk

`func (o *ObjectLastRun) GetVCenterSummaryOk() (*ObjectTypeVCenterParams, bool)`

GetVCenterSummaryOk returns a tuple with the VCenterSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCenterSummary

`func (o *ObjectLastRun) SetVCenterSummary(v ObjectTypeVCenterParams)`

SetVCenterSummary sets VCenterSummary field to given value.

### HasVCenterSummary

`func (o *ObjectLastRun) HasVCenterSummary() bool`

HasVCenterSummary returns a boolean if a field has been set.

### GetSharepointSiteSummary

`func (o *ObjectLastRun) GetSharepointSiteSummary() SharepointObjectParams`

GetSharepointSiteSummary returns the SharepointSiteSummary field if non-nil, zero value otherwise.

### GetSharepointSiteSummaryOk

`func (o *ObjectLastRun) GetSharepointSiteSummaryOk() (*SharepointObjectParams, bool)`

GetSharepointSiteSummaryOk returns a tuple with the SharepointSiteSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharepointSiteSummary

`func (o *ObjectLastRun) SetSharepointSiteSummary(v SharepointObjectParams)`

SetSharepointSiteSummary sets SharepointSiteSummary field to given value.

### HasSharepointSiteSummary

`func (o *ObjectLastRun) HasSharepointSiteSummary() bool`

HasSharepointSiteSummary returns a boolean if a field has been set.

### GetRunId

`func (o *ObjectLastRun) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *ObjectLastRun) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *ObjectLastRun) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *ObjectLastRun) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### GetProtectionGroupName

`func (o *ObjectLastRun) GetProtectionGroupName() string`

GetProtectionGroupName returns the ProtectionGroupName field if non-nil, zero value otherwise.

### GetProtectionGroupNameOk

`func (o *ObjectLastRun) GetProtectionGroupNameOk() (*string, bool)`

GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupName

`func (o *ObjectLastRun) SetProtectionGroupName(v string)`

SetProtectionGroupName sets ProtectionGroupName field to given value.

### HasProtectionGroupName

`func (o *ObjectLastRun) HasProtectionGroupName() bool`

HasProtectionGroupName returns a boolean if a field has been set.

### SetProtectionGroupNameNil

`func (o *ObjectLastRun) SetProtectionGroupNameNil(b bool)`

 SetProtectionGroupNameNil sets the value for ProtectionGroupName to be an explicit nil

### UnsetProtectionGroupName
`func (o *ObjectLastRun) UnsetProtectionGroupName()`

UnsetProtectionGroupName ensures that no value is present for ProtectionGroupName, not even an explicit nil
### GetProtectionGroupId

`func (o *ObjectLastRun) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *ObjectLastRun) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *ObjectLastRun) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *ObjectLastRun) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *ObjectLastRun) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *ObjectLastRun) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetPolicyName

`func (o *ObjectLastRun) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *ObjectLastRun) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *ObjectLastRun) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *ObjectLastRun) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### SetPolicyNameNil

`func (o *ObjectLastRun) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *ObjectLastRun) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
### GetPolicyId

`func (o *ObjectLastRun) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ObjectLastRun) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ObjectLastRun) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *ObjectLastRun) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *ObjectLastRun) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *ObjectLastRun) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetBackupRunStatus

`func (o *ObjectLastRun) GetBackupRunStatus() string`

GetBackupRunStatus returns the BackupRunStatus field if non-nil, zero value otherwise.

### GetBackupRunStatusOk

`func (o *ObjectLastRun) GetBackupRunStatusOk() (*string, bool)`

GetBackupRunStatusOk returns a tuple with the BackupRunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRunStatus

`func (o *ObjectLastRun) SetBackupRunStatus(v string)`

SetBackupRunStatus sets BackupRunStatus field to given value.

### HasBackupRunStatus

`func (o *ObjectLastRun) HasBackupRunStatus() bool`

HasBackupRunStatus returns a boolean if a field has been set.

### SetBackupRunStatusNil

`func (o *ObjectLastRun) SetBackupRunStatusNil(b bool)`

 SetBackupRunStatusNil sets the value for BackupRunStatus to be an explicit nil

### UnsetBackupRunStatus
`func (o *ObjectLastRun) UnsetBackupRunStatus()`

UnsetBackupRunStatus ensures that no value is present for BackupRunStatus, not even an explicit nil
### GetArchivalRunStatus

`func (o *ObjectLastRun) GetArchivalRunStatus() string`

GetArchivalRunStatus returns the ArchivalRunStatus field if non-nil, zero value otherwise.

### GetArchivalRunStatusOk

`func (o *ObjectLastRun) GetArchivalRunStatusOk() (*string, bool)`

GetArchivalRunStatusOk returns a tuple with the ArchivalRunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalRunStatus

`func (o *ObjectLastRun) SetArchivalRunStatus(v string)`

SetArchivalRunStatus sets ArchivalRunStatus field to given value.

### HasArchivalRunStatus

`func (o *ObjectLastRun) HasArchivalRunStatus() bool`

HasArchivalRunStatus returns a boolean if a field has been set.

### SetArchivalRunStatusNil

`func (o *ObjectLastRun) SetArchivalRunStatusNil(b bool)`

 SetArchivalRunStatusNil sets the value for ArchivalRunStatus to be an explicit nil

### UnsetArchivalRunStatus
`func (o *ObjectLastRun) UnsetArchivalRunStatus()`

UnsetArchivalRunStatus ensures that no value is present for ArchivalRunStatus, not even an explicit nil
### GetReplicationRunStatus

`func (o *ObjectLastRun) GetReplicationRunStatus() string`

GetReplicationRunStatus returns the ReplicationRunStatus field if non-nil, zero value otherwise.

### GetReplicationRunStatusOk

`func (o *ObjectLastRun) GetReplicationRunStatusOk() (*string, bool)`

GetReplicationRunStatusOk returns a tuple with the ReplicationRunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationRunStatus

`func (o *ObjectLastRun) SetReplicationRunStatus(v string)`

SetReplicationRunStatus sets ReplicationRunStatus field to given value.

### HasReplicationRunStatus

`func (o *ObjectLastRun) HasReplicationRunStatus() bool`

HasReplicationRunStatus returns a boolean if a field has been set.

### SetReplicationRunStatusNil

`func (o *ObjectLastRun) SetReplicationRunStatusNil(b bool)`

 SetReplicationRunStatusNil sets the value for ReplicationRunStatus to be an explicit nil

### UnsetReplicationRunStatus
`func (o *ObjectLastRun) UnsetReplicationRunStatus()`

UnsetReplicationRunStatus ensures that no value is present for ReplicationRunStatus, not even an explicit nil
### GetIsSlaViolated

`func (o *ObjectLastRun) GetIsSlaViolated() bool`

GetIsSlaViolated returns the IsSlaViolated field if non-nil, zero value otherwise.

### GetIsSlaViolatedOk

`func (o *ObjectLastRun) GetIsSlaViolatedOk() (*bool, bool)`

GetIsSlaViolatedOk returns a tuple with the IsSlaViolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSlaViolated

`func (o *ObjectLastRun) SetIsSlaViolated(v bool)`

SetIsSlaViolated sets IsSlaViolated field to given value.

### HasIsSlaViolated

`func (o *ObjectLastRun) HasIsSlaViolated() bool`

HasIsSlaViolated returns a boolean if a field has been set.

### SetIsSlaViolatedNil

`func (o *ObjectLastRun) SetIsSlaViolatedNil(b bool)`

 SetIsSlaViolatedNil sets the value for IsSlaViolated to be an explicit nil

### UnsetIsSlaViolated
`func (o *ObjectLastRun) UnsetIsSlaViolated()`

UnsetIsSlaViolated ensures that no value is present for IsSlaViolated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


