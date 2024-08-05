# ObjectProtectionRunInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchivalInfo** | Pointer to [**ArchivalRun**](ArchivalRun.md) |  | [optional] 
**CloudSpinInfo** | Pointer to [**CloudSpinRun**](CloudSpinRun.md) |  | [optional] 
**DataLock** | Pointer to **NullableString** | Specifies WORM retention type for the local backeup. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept until the maximum of the snapshot retention time. During that time, the snapshots cannot be deleted. &lt;br&gt;&#39;Compliance&#39; implies WORM retention is set for compliance reason. &lt;br&gt;&#39;Administrative&#39; implies WORM retention is set for administrative purposes. | [optional] 
**IsCloudArchivalDirect** | Pointer to **NullableBool** | Specifies whether the run is a CAD run if cloud archive direct feature is enabled. If this field is true, the primary backup copy will only be available at the given archived location. | [optional] 
**IsLocalSnapshotsDeleted** | Pointer to **NullableBool** | Specifies if snapshots for this run has been deleted. | [optional] 
**IsReplicationRun** | Pointer to **NullableBool** | Specifies if this protection run is a replication run. | [optional] 
**IsSlaViolated** | Pointer to **NullableBool** | Indicated if SLA has been violated for this run. | [optional] 
**LocalSnapshotInfo** | Pointer to [**BackupRun**](BackupRun.md) |  | [optional] 
**OnLegalHold** | Pointer to **NullableBool** | Specifies if object&#39;s snapshot is on legal hold. | [optional] 
**OnPremDeployInfo** | Pointer to [**OnPremDeployRun**](OnPremDeployRun.md) |  | [optional] 
**OriginClusterIdentifier** | Pointer to [**ClusterIdentifier**](ClusterIdentifier.md) |  | [optional] 
**OriginProtectionGroupId** | Pointer to **NullableString** | ProtectionGroupId to which this run belongs on the primary cluster if this run is a replication run. | [optional] 
**OriginalBackupInfo** | Pointer to [**BackupRun**](BackupRun.md) |  | [optional] 
**Permissions** | Pointer to [**[]Tenant**](Tenant.md) | Specifies the list of tenants that have permissions for this protection group run. | [optional] 
**PolicyId** | Pointer to **NullableString** | Specifies the unique id of the Protection Policy associated with the Protection Run. The Policy provides retry settings Protection Schedules, Priority, SLA, etc. | [optional] 
**PolicyName** | Pointer to **NullableString** | Specifies Specifies the name of the Protection Policy. | [optional] 
**ProtectionGroupId** | Pointer to **NullableString** | ProtectionGroupId to which this run belongs. This will only be populated if the object is protected by a protection group. | [optional] 
**ProtectionGroupName** | Pointer to **NullableString** | Name of the Protection Group to which this run belongs. This will only be populated if the object is protected by a protection group. | [optional] 
**ReplicationInfo** | Pointer to [**ReplicationRun**](ReplicationRun.md) |  | [optional] 
**RunId** | Pointer to **NullableString** | Specifies the ID of the protection run. | [optional] 
**RunLabel** | Pointer to **NullableString** | Specifies a label with which this run is created. Only applicable for user triggered protect now action. | [optional] 
**RunType** | Pointer to **NullableString** | Type of Protection run. &#39;kRegular&#39; indicates an incremental (CBT) backup. Incremental backups utilizing CBT (if supported) are captured of the target protection objects. The first run of a kRegular schedule captures all the blocks. &#39;kFull&#39; indicates a full (no CBT) backup. A complete backup (all blocks) of the target protection objects are always captured and Change Block Tracking (CBT) is not utilized. &#39;kLog&#39; indicates a Database Log backup. Capture the database transaction logs to allow rolling back to a specific point in time. &#39;kSystem&#39; indicates system volume backup. It produces an image for bare metal recovery. | [optional] 
**StorageDomainId** | Pointer to **NullableInt64** | Specifies the Storage Domain (View Box) ID where this Protection Run writes data. | [optional] 

## Methods

### NewObjectProtectionRunInfo

`func NewObjectProtectionRunInfo() *ObjectProtectionRunInfo`

NewObjectProtectionRunInfo instantiates a new ObjectProtectionRunInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectProtectionRunInfoWithDefaults

`func NewObjectProtectionRunInfoWithDefaults() *ObjectProtectionRunInfo`

NewObjectProtectionRunInfoWithDefaults instantiates a new ObjectProtectionRunInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchivalInfo

`func (o *ObjectProtectionRunInfo) GetArchivalInfo() ArchivalRun`

GetArchivalInfo returns the ArchivalInfo field if non-nil, zero value otherwise.

### GetArchivalInfoOk

`func (o *ObjectProtectionRunInfo) GetArchivalInfoOk() (*ArchivalRun, bool)`

GetArchivalInfoOk returns a tuple with the ArchivalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalInfo

`func (o *ObjectProtectionRunInfo) SetArchivalInfo(v ArchivalRun)`

SetArchivalInfo sets ArchivalInfo field to given value.

### HasArchivalInfo

`func (o *ObjectProtectionRunInfo) HasArchivalInfo() bool`

HasArchivalInfo returns a boolean if a field has been set.

### GetCloudSpinInfo

`func (o *ObjectProtectionRunInfo) GetCloudSpinInfo() CloudSpinRun`

GetCloudSpinInfo returns the CloudSpinInfo field if non-nil, zero value otherwise.

### GetCloudSpinInfoOk

`func (o *ObjectProtectionRunInfo) GetCloudSpinInfoOk() (*CloudSpinRun, bool)`

GetCloudSpinInfoOk returns a tuple with the CloudSpinInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudSpinInfo

`func (o *ObjectProtectionRunInfo) SetCloudSpinInfo(v CloudSpinRun)`

SetCloudSpinInfo sets CloudSpinInfo field to given value.

### HasCloudSpinInfo

`func (o *ObjectProtectionRunInfo) HasCloudSpinInfo() bool`

HasCloudSpinInfo returns a boolean if a field has been set.

### GetDataLock

`func (o *ObjectProtectionRunInfo) GetDataLock() string`

GetDataLock returns the DataLock field if non-nil, zero value otherwise.

### GetDataLockOk

`func (o *ObjectProtectionRunInfo) GetDataLockOk() (*string, bool)`

GetDataLockOk returns a tuple with the DataLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLock

`func (o *ObjectProtectionRunInfo) SetDataLock(v string)`

SetDataLock sets DataLock field to given value.

### HasDataLock

`func (o *ObjectProtectionRunInfo) HasDataLock() bool`

HasDataLock returns a boolean if a field has been set.

### SetDataLockNil

`func (o *ObjectProtectionRunInfo) SetDataLockNil(b bool)`

 SetDataLockNil sets the value for DataLock to be an explicit nil

### UnsetDataLock
`func (o *ObjectProtectionRunInfo) UnsetDataLock()`

UnsetDataLock ensures that no value is present for DataLock, not even an explicit nil
### GetIsCloudArchivalDirect

`func (o *ObjectProtectionRunInfo) GetIsCloudArchivalDirect() bool`

GetIsCloudArchivalDirect returns the IsCloudArchivalDirect field if non-nil, zero value otherwise.

### GetIsCloudArchivalDirectOk

`func (o *ObjectProtectionRunInfo) GetIsCloudArchivalDirectOk() (*bool, bool)`

GetIsCloudArchivalDirectOk returns a tuple with the IsCloudArchivalDirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloudArchivalDirect

`func (o *ObjectProtectionRunInfo) SetIsCloudArchivalDirect(v bool)`

SetIsCloudArchivalDirect sets IsCloudArchivalDirect field to given value.

### HasIsCloudArchivalDirect

`func (o *ObjectProtectionRunInfo) HasIsCloudArchivalDirect() bool`

HasIsCloudArchivalDirect returns a boolean if a field has been set.

### SetIsCloudArchivalDirectNil

`func (o *ObjectProtectionRunInfo) SetIsCloudArchivalDirectNil(b bool)`

 SetIsCloudArchivalDirectNil sets the value for IsCloudArchivalDirect to be an explicit nil

### UnsetIsCloudArchivalDirect
`func (o *ObjectProtectionRunInfo) UnsetIsCloudArchivalDirect()`

UnsetIsCloudArchivalDirect ensures that no value is present for IsCloudArchivalDirect, not even an explicit nil
### GetIsLocalSnapshotsDeleted

`func (o *ObjectProtectionRunInfo) GetIsLocalSnapshotsDeleted() bool`

GetIsLocalSnapshotsDeleted returns the IsLocalSnapshotsDeleted field if non-nil, zero value otherwise.

### GetIsLocalSnapshotsDeletedOk

`func (o *ObjectProtectionRunInfo) GetIsLocalSnapshotsDeletedOk() (*bool, bool)`

GetIsLocalSnapshotsDeletedOk returns a tuple with the IsLocalSnapshotsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocalSnapshotsDeleted

`func (o *ObjectProtectionRunInfo) SetIsLocalSnapshotsDeleted(v bool)`

SetIsLocalSnapshotsDeleted sets IsLocalSnapshotsDeleted field to given value.

### HasIsLocalSnapshotsDeleted

`func (o *ObjectProtectionRunInfo) HasIsLocalSnapshotsDeleted() bool`

HasIsLocalSnapshotsDeleted returns a boolean if a field has been set.

### SetIsLocalSnapshotsDeletedNil

`func (o *ObjectProtectionRunInfo) SetIsLocalSnapshotsDeletedNil(b bool)`

 SetIsLocalSnapshotsDeletedNil sets the value for IsLocalSnapshotsDeleted to be an explicit nil

### UnsetIsLocalSnapshotsDeleted
`func (o *ObjectProtectionRunInfo) UnsetIsLocalSnapshotsDeleted()`

UnsetIsLocalSnapshotsDeleted ensures that no value is present for IsLocalSnapshotsDeleted, not even an explicit nil
### GetIsReplicationRun

`func (o *ObjectProtectionRunInfo) GetIsReplicationRun() bool`

GetIsReplicationRun returns the IsReplicationRun field if non-nil, zero value otherwise.

### GetIsReplicationRunOk

`func (o *ObjectProtectionRunInfo) GetIsReplicationRunOk() (*bool, bool)`

GetIsReplicationRunOk returns a tuple with the IsReplicationRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReplicationRun

`func (o *ObjectProtectionRunInfo) SetIsReplicationRun(v bool)`

SetIsReplicationRun sets IsReplicationRun field to given value.

### HasIsReplicationRun

`func (o *ObjectProtectionRunInfo) HasIsReplicationRun() bool`

HasIsReplicationRun returns a boolean if a field has been set.

### SetIsReplicationRunNil

`func (o *ObjectProtectionRunInfo) SetIsReplicationRunNil(b bool)`

 SetIsReplicationRunNil sets the value for IsReplicationRun to be an explicit nil

### UnsetIsReplicationRun
`func (o *ObjectProtectionRunInfo) UnsetIsReplicationRun()`

UnsetIsReplicationRun ensures that no value is present for IsReplicationRun, not even an explicit nil
### GetIsSlaViolated

`func (o *ObjectProtectionRunInfo) GetIsSlaViolated() bool`

GetIsSlaViolated returns the IsSlaViolated field if non-nil, zero value otherwise.

### GetIsSlaViolatedOk

`func (o *ObjectProtectionRunInfo) GetIsSlaViolatedOk() (*bool, bool)`

GetIsSlaViolatedOk returns a tuple with the IsSlaViolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSlaViolated

`func (o *ObjectProtectionRunInfo) SetIsSlaViolated(v bool)`

SetIsSlaViolated sets IsSlaViolated field to given value.

### HasIsSlaViolated

`func (o *ObjectProtectionRunInfo) HasIsSlaViolated() bool`

HasIsSlaViolated returns a boolean if a field has been set.

### SetIsSlaViolatedNil

`func (o *ObjectProtectionRunInfo) SetIsSlaViolatedNil(b bool)`

 SetIsSlaViolatedNil sets the value for IsSlaViolated to be an explicit nil

### UnsetIsSlaViolated
`func (o *ObjectProtectionRunInfo) UnsetIsSlaViolated()`

UnsetIsSlaViolated ensures that no value is present for IsSlaViolated, not even an explicit nil
### GetLocalSnapshotInfo

`func (o *ObjectProtectionRunInfo) GetLocalSnapshotInfo() BackupRun`

GetLocalSnapshotInfo returns the LocalSnapshotInfo field if non-nil, zero value otherwise.

### GetLocalSnapshotInfoOk

`func (o *ObjectProtectionRunInfo) GetLocalSnapshotInfoOk() (*BackupRun, bool)`

GetLocalSnapshotInfoOk returns a tuple with the LocalSnapshotInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSnapshotInfo

`func (o *ObjectProtectionRunInfo) SetLocalSnapshotInfo(v BackupRun)`

SetLocalSnapshotInfo sets LocalSnapshotInfo field to given value.

### HasLocalSnapshotInfo

`func (o *ObjectProtectionRunInfo) HasLocalSnapshotInfo() bool`

HasLocalSnapshotInfo returns a boolean if a field has been set.

### GetOnLegalHold

`func (o *ObjectProtectionRunInfo) GetOnLegalHold() bool`

GetOnLegalHold returns the OnLegalHold field if non-nil, zero value otherwise.

### GetOnLegalHoldOk

`func (o *ObjectProtectionRunInfo) GetOnLegalHoldOk() (*bool, bool)`

GetOnLegalHoldOk returns a tuple with the OnLegalHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnLegalHold

`func (o *ObjectProtectionRunInfo) SetOnLegalHold(v bool)`

SetOnLegalHold sets OnLegalHold field to given value.

### HasOnLegalHold

`func (o *ObjectProtectionRunInfo) HasOnLegalHold() bool`

HasOnLegalHold returns a boolean if a field has been set.

### SetOnLegalHoldNil

`func (o *ObjectProtectionRunInfo) SetOnLegalHoldNil(b bool)`

 SetOnLegalHoldNil sets the value for OnLegalHold to be an explicit nil

### UnsetOnLegalHold
`func (o *ObjectProtectionRunInfo) UnsetOnLegalHold()`

UnsetOnLegalHold ensures that no value is present for OnLegalHold, not even an explicit nil
### GetOnPremDeployInfo

`func (o *ObjectProtectionRunInfo) GetOnPremDeployInfo() OnPremDeployRun`

GetOnPremDeployInfo returns the OnPremDeployInfo field if non-nil, zero value otherwise.

### GetOnPremDeployInfoOk

`func (o *ObjectProtectionRunInfo) GetOnPremDeployInfoOk() (*OnPremDeployRun, bool)`

GetOnPremDeployInfoOk returns a tuple with the OnPremDeployInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremDeployInfo

`func (o *ObjectProtectionRunInfo) SetOnPremDeployInfo(v OnPremDeployRun)`

SetOnPremDeployInfo sets OnPremDeployInfo field to given value.

### HasOnPremDeployInfo

`func (o *ObjectProtectionRunInfo) HasOnPremDeployInfo() bool`

HasOnPremDeployInfo returns a boolean if a field has been set.

### GetOriginClusterIdentifier

`func (o *ObjectProtectionRunInfo) GetOriginClusterIdentifier() ClusterIdentifier`

GetOriginClusterIdentifier returns the OriginClusterIdentifier field if non-nil, zero value otherwise.

### GetOriginClusterIdentifierOk

`func (o *ObjectProtectionRunInfo) GetOriginClusterIdentifierOk() (*ClusterIdentifier, bool)`

GetOriginClusterIdentifierOk returns a tuple with the OriginClusterIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginClusterIdentifier

`func (o *ObjectProtectionRunInfo) SetOriginClusterIdentifier(v ClusterIdentifier)`

SetOriginClusterIdentifier sets OriginClusterIdentifier field to given value.

### HasOriginClusterIdentifier

`func (o *ObjectProtectionRunInfo) HasOriginClusterIdentifier() bool`

HasOriginClusterIdentifier returns a boolean if a field has been set.

### GetOriginProtectionGroupId

`func (o *ObjectProtectionRunInfo) GetOriginProtectionGroupId() string`

GetOriginProtectionGroupId returns the OriginProtectionGroupId field if non-nil, zero value otherwise.

### GetOriginProtectionGroupIdOk

`func (o *ObjectProtectionRunInfo) GetOriginProtectionGroupIdOk() (*string, bool)`

GetOriginProtectionGroupIdOk returns a tuple with the OriginProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginProtectionGroupId

`func (o *ObjectProtectionRunInfo) SetOriginProtectionGroupId(v string)`

SetOriginProtectionGroupId sets OriginProtectionGroupId field to given value.

### HasOriginProtectionGroupId

`func (o *ObjectProtectionRunInfo) HasOriginProtectionGroupId() bool`

HasOriginProtectionGroupId returns a boolean if a field has been set.

### SetOriginProtectionGroupIdNil

`func (o *ObjectProtectionRunInfo) SetOriginProtectionGroupIdNil(b bool)`

 SetOriginProtectionGroupIdNil sets the value for OriginProtectionGroupId to be an explicit nil

### UnsetOriginProtectionGroupId
`func (o *ObjectProtectionRunInfo) UnsetOriginProtectionGroupId()`

UnsetOriginProtectionGroupId ensures that no value is present for OriginProtectionGroupId, not even an explicit nil
### GetOriginalBackupInfo

`func (o *ObjectProtectionRunInfo) GetOriginalBackupInfo() BackupRun`

GetOriginalBackupInfo returns the OriginalBackupInfo field if non-nil, zero value otherwise.

### GetOriginalBackupInfoOk

`func (o *ObjectProtectionRunInfo) GetOriginalBackupInfoOk() (*BackupRun, bool)`

GetOriginalBackupInfoOk returns a tuple with the OriginalBackupInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalBackupInfo

`func (o *ObjectProtectionRunInfo) SetOriginalBackupInfo(v BackupRun)`

SetOriginalBackupInfo sets OriginalBackupInfo field to given value.

### HasOriginalBackupInfo

`func (o *ObjectProtectionRunInfo) HasOriginalBackupInfo() bool`

HasOriginalBackupInfo returns a boolean if a field has been set.

### GetPermissions

`func (o *ObjectProtectionRunInfo) GetPermissions() []Tenant`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ObjectProtectionRunInfo) GetPermissionsOk() (*[]Tenant, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ObjectProtectionRunInfo) SetPermissions(v []Tenant)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ObjectProtectionRunInfo) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *ObjectProtectionRunInfo) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *ObjectProtectionRunInfo) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetPolicyId

`func (o *ObjectProtectionRunInfo) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ObjectProtectionRunInfo) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ObjectProtectionRunInfo) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *ObjectProtectionRunInfo) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *ObjectProtectionRunInfo) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *ObjectProtectionRunInfo) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetPolicyName

`func (o *ObjectProtectionRunInfo) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *ObjectProtectionRunInfo) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *ObjectProtectionRunInfo) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *ObjectProtectionRunInfo) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### SetPolicyNameNil

`func (o *ObjectProtectionRunInfo) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *ObjectProtectionRunInfo) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
### GetProtectionGroupId

`func (o *ObjectProtectionRunInfo) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *ObjectProtectionRunInfo) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *ObjectProtectionRunInfo) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *ObjectProtectionRunInfo) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *ObjectProtectionRunInfo) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *ObjectProtectionRunInfo) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetProtectionGroupName

`func (o *ObjectProtectionRunInfo) GetProtectionGroupName() string`

GetProtectionGroupName returns the ProtectionGroupName field if non-nil, zero value otherwise.

### GetProtectionGroupNameOk

`func (o *ObjectProtectionRunInfo) GetProtectionGroupNameOk() (*string, bool)`

GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupName

`func (o *ObjectProtectionRunInfo) SetProtectionGroupName(v string)`

SetProtectionGroupName sets ProtectionGroupName field to given value.

### HasProtectionGroupName

`func (o *ObjectProtectionRunInfo) HasProtectionGroupName() bool`

HasProtectionGroupName returns a boolean if a field has been set.

### SetProtectionGroupNameNil

`func (o *ObjectProtectionRunInfo) SetProtectionGroupNameNil(b bool)`

 SetProtectionGroupNameNil sets the value for ProtectionGroupName to be an explicit nil

### UnsetProtectionGroupName
`func (o *ObjectProtectionRunInfo) UnsetProtectionGroupName()`

UnsetProtectionGroupName ensures that no value is present for ProtectionGroupName, not even an explicit nil
### GetReplicationInfo

`func (o *ObjectProtectionRunInfo) GetReplicationInfo() ReplicationRun`

GetReplicationInfo returns the ReplicationInfo field if non-nil, zero value otherwise.

### GetReplicationInfoOk

`func (o *ObjectProtectionRunInfo) GetReplicationInfoOk() (*ReplicationRun, bool)`

GetReplicationInfoOk returns a tuple with the ReplicationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationInfo

`func (o *ObjectProtectionRunInfo) SetReplicationInfo(v ReplicationRun)`

SetReplicationInfo sets ReplicationInfo field to given value.

### HasReplicationInfo

`func (o *ObjectProtectionRunInfo) HasReplicationInfo() bool`

HasReplicationInfo returns a boolean if a field has been set.

### GetRunId

`func (o *ObjectProtectionRunInfo) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *ObjectProtectionRunInfo) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *ObjectProtectionRunInfo) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *ObjectProtectionRunInfo) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### SetRunIdNil

`func (o *ObjectProtectionRunInfo) SetRunIdNil(b bool)`

 SetRunIdNil sets the value for RunId to be an explicit nil

### UnsetRunId
`func (o *ObjectProtectionRunInfo) UnsetRunId()`

UnsetRunId ensures that no value is present for RunId, not even an explicit nil
### GetRunLabel

`func (o *ObjectProtectionRunInfo) GetRunLabel() string`

GetRunLabel returns the RunLabel field if non-nil, zero value otherwise.

### GetRunLabelOk

`func (o *ObjectProtectionRunInfo) GetRunLabelOk() (*string, bool)`

GetRunLabelOk returns a tuple with the RunLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunLabel

`func (o *ObjectProtectionRunInfo) SetRunLabel(v string)`

SetRunLabel sets RunLabel field to given value.

### HasRunLabel

`func (o *ObjectProtectionRunInfo) HasRunLabel() bool`

HasRunLabel returns a boolean if a field has been set.

### SetRunLabelNil

`func (o *ObjectProtectionRunInfo) SetRunLabelNil(b bool)`

 SetRunLabelNil sets the value for RunLabel to be an explicit nil

### UnsetRunLabel
`func (o *ObjectProtectionRunInfo) UnsetRunLabel()`

UnsetRunLabel ensures that no value is present for RunLabel, not even an explicit nil
### GetRunType

`func (o *ObjectProtectionRunInfo) GetRunType() string`

GetRunType returns the RunType field if non-nil, zero value otherwise.

### GetRunTypeOk

`func (o *ObjectProtectionRunInfo) GetRunTypeOk() (*string, bool)`

GetRunTypeOk returns a tuple with the RunType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunType

`func (o *ObjectProtectionRunInfo) SetRunType(v string)`

SetRunType sets RunType field to given value.

### HasRunType

`func (o *ObjectProtectionRunInfo) HasRunType() bool`

HasRunType returns a boolean if a field has been set.

### SetRunTypeNil

`func (o *ObjectProtectionRunInfo) SetRunTypeNil(b bool)`

 SetRunTypeNil sets the value for RunType to be an explicit nil

### UnsetRunType
`func (o *ObjectProtectionRunInfo) UnsetRunType()`

UnsetRunType ensures that no value is present for RunType, not even an explicit nil
### GetStorageDomainId

`func (o *ObjectProtectionRunInfo) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *ObjectProtectionRunInfo) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *ObjectProtectionRunInfo) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.

### HasStorageDomainId

`func (o *ObjectProtectionRunInfo) HasStorageDomainId() bool`

HasStorageDomainId returns a boolean if a field has been set.

### SetStorageDomainIdNil

`func (o *ObjectProtectionRunInfo) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *ObjectProtectionRunInfo) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


