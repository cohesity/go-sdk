# ObjectProtectionRunSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to **NullableString** | Specifies the environment of the object. | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies object id. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the object. | [optional] 
**SourceId** | Pointer to **NullableInt64** | Specifies registered source id to which object belongs. | [optional] 
**SourceName** | Pointer to **NullableString** | Specifies registered source name to which object belongs. | [optional] 
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

### NewObjectProtectionRunSummary

`func NewObjectProtectionRunSummary() *ObjectProtectionRunSummary`

NewObjectProtectionRunSummary instantiates a new ObjectProtectionRunSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectProtectionRunSummaryWithDefaults

`func NewObjectProtectionRunSummaryWithDefaults() *ObjectProtectionRunSummary`

NewObjectProtectionRunSummaryWithDefaults instantiates a new ObjectProtectionRunSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *ObjectProtectionRunSummary) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ObjectProtectionRunSummary) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ObjectProtectionRunSummary) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ObjectProtectionRunSummary) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *ObjectProtectionRunSummary) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *ObjectProtectionRunSummary) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetId

`func (o *ObjectProtectionRunSummary) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectProtectionRunSummary) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectProtectionRunSummary) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ObjectProtectionRunSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ObjectProtectionRunSummary) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ObjectProtectionRunSummary) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ObjectProtectionRunSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectProtectionRunSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectProtectionRunSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ObjectProtectionRunSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ObjectProtectionRunSummary) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ObjectProtectionRunSummary) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSourceId

`func (o *ObjectProtectionRunSummary) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ObjectProtectionRunSummary) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ObjectProtectionRunSummary) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *ObjectProtectionRunSummary) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *ObjectProtectionRunSummary) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *ObjectProtectionRunSummary) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *ObjectProtectionRunSummary) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *ObjectProtectionRunSummary) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *ObjectProtectionRunSummary) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *ObjectProtectionRunSummary) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *ObjectProtectionRunSummary) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *ObjectProtectionRunSummary) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetArchivalInfo

`func (o *ObjectProtectionRunSummary) GetArchivalInfo() ArchivalRun`

GetArchivalInfo returns the ArchivalInfo field if non-nil, zero value otherwise.

### GetArchivalInfoOk

`func (o *ObjectProtectionRunSummary) GetArchivalInfoOk() (*ArchivalRun, bool)`

GetArchivalInfoOk returns a tuple with the ArchivalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalInfo

`func (o *ObjectProtectionRunSummary) SetArchivalInfo(v ArchivalRun)`

SetArchivalInfo sets ArchivalInfo field to given value.

### HasArchivalInfo

`func (o *ObjectProtectionRunSummary) HasArchivalInfo() bool`

HasArchivalInfo returns a boolean if a field has been set.

### GetCloudSpinInfo

`func (o *ObjectProtectionRunSummary) GetCloudSpinInfo() CloudSpinRun`

GetCloudSpinInfo returns the CloudSpinInfo field if non-nil, zero value otherwise.

### GetCloudSpinInfoOk

`func (o *ObjectProtectionRunSummary) GetCloudSpinInfoOk() (*CloudSpinRun, bool)`

GetCloudSpinInfoOk returns a tuple with the CloudSpinInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudSpinInfo

`func (o *ObjectProtectionRunSummary) SetCloudSpinInfo(v CloudSpinRun)`

SetCloudSpinInfo sets CloudSpinInfo field to given value.

### HasCloudSpinInfo

`func (o *ObjectProtectionRunSummary) HasCloudSpinInfo() bool`

HasCloudSpinInfo returns a boolean if a field has been set.

### GetDataLock

`func (o *ObjectProtectionRunSummary) GetDataLock() string`

GetDataLock returns the DataLock field if non-nil, zero value otherwise.

### GetDataLockOk

`func (o *ObjectProtectionRunSummary) GetDataLockOk() (*string, bool)`

GetDataLockOk returns a tuple with the DataLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLock

`func (o *ObjectProtectionRunSummary) SetDataLock(v string)`

SetDataLock sets DataLock field to given value.

### HasDataLock

`func (o *ObjectProtectionRunSummary) HasDataLock() bool`

HasDataLock returns a boolean if a field has been set.

### SetDataLockNil

`func (o *ObjectProtectionRunSummary) SetDataLockNil(b bool)`

 SetDataLockNil sets the value for DataLock to be an explicit nil

### UnsetDataLock
`func (o *ObjectProtectionRunSummary) UnsetDataLock()`

UnsetDataLock ensures that no value is present for DataLock, not even an explicit nil
### GetIsCloudArchivalDirect

`func (o *ObjectProtectionRunSummary) GetIsCloudArchivalDirect() bool`

GetIsCloudArchivalDirect returns the IsCloudArchivalDirect field if non-nil, zero value otherwise.

### GetIsCloudArchivalDirectOk

`func (o *ObjectProtectionRunSummary) GetIsCloudArchivalDirectOk() (*bool, bool)`

GetIsCloudArchivalDirectOk returns a tuple with the IsCloudArchivalDirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloudArchivalDirect

`func (o *ObjectProtectionRunSummary) SetIsCloudArchivalDirect(v bool)`

SetIsCloudArchivalDirect sets IsCloudArchivalDirect field to given value.

### HasIsCloudArchivalDirect

`func (o *ObjectProtectionRunSummary) HasIsCloudArchivalDirect() bool`

HasIsCloudArchivalDirect returns a boolean if a field has been set.

### SetIsCloudArchivalDirectNil

`func (o *ObjectProtectionRunSummary) SetIsCloudArchivalDirectNil(b bool)`

 SetIsCloudArchivalDirectNil sets the value for IsCloudArchivalDirect to be an explicit nil

### UnsetIsCloudArchivalDirect
`func (o *ObjectProtectionRunSummary) UnsetIsCloudArchivalDirect()`

UnsetIsCloudArchivalDirect ensures that no value is present for IsCloudArchivalDirect, not even an explicit nil
### GetIsLocalSnapshotsDeleted

`func (o *ObjectProtectionRunSummary) GetIsLocalSnapshotsDeleted() bool`

GetIsLocalSnapshotsDeleted returns the IsLocalSnapshotsDeleted field if non-nil, zero value otherwise.

### GetIsLocalSnapshotsDeletedOk

`func (o *ObjectProtectionRunSummary) GetIsLocalSnapshotsDeletedOk() (*bool, bool)`

GetIsLocalSnapshotsDeletedOk returns a tuple with the IsLocalSnapshotsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocalSnapshotsDeleted

`func (o *ObjectProtectionRunSummary) SetIsLocalSnapshotsDeleted(v bool)`

SetIsLocalSnapshotsDeleted sets IsLocalSnapshotsDeleted field to given value.

### HasIsLocalSnapshotsDeleted

`func (o *ObjectProtectionRunSummary) HasIsLocalSnapshotsDeleted() bool`

HasIsLocalSnapshotsDeleted returns a boolean if a field has been set.

### SetIsLocalSnapshotsDeletedNil

`func (o *ObjectProtectionRunSummary) SetIsLocalSnapshotsDeletedNil(b bool)`

 SetIsLocalSnapshotsDeletedNil sets the value for IsLocalSnapshotsDeleted to be an explicit nil

### UnsetIsLocalSnapshotsDeleted
`func (o *ObjectProtectionRunSummary) UnsetIsLocalSnapshotsDeleted()`

UnsetIsLocalSnapshotsDeleted ensures that no value is present for IsLocalSnapshotsDeleted, not even an explicit nil
### GetIsReplicationRun

`func (o *ObjectProtectionRunSummary) GetIsReplicationRun() bool`

GetIsReplicationRun returns the IsReplicationRun field if non-nil, zero value otherwise.

### GetIsReplicationRunOk

`func (o *ObjectProtectionRunSummary) GetIsReplicationRunOk() (*bool, bool)`

GetIsReplicationRunOk returns a tuple with the IsReplicationRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReplicationRun

`func (o *ObjectProtectionRunSummary) SetIsReplicationRun(v bool)`

SetIsReplicationRun sets IsReplicationRun field to given value.

### HasIsReplicationRun

`func (o *ObjectProtectionRunSummary) HasIsReplicationRun() bool`

HasIsReplicationRun returns a boolean if a field has been set.

### SetIsReplicationRunNil

`func (o *ObjectProtectionRunSummary) SetIsReplicationRunNil(b bool)`

 SetIsReplicationRunNil sets the value for IsReplicationRun to be an explicit nil

### UnsetIsReplicationRun
`func (o *ObjectProtectionRunSummary) UnsetIsReplicationRun()`

UnsetIsReplicationRun ensures that no value is present for IsReplicationRun, not even an explicit nil
### GetIsSlaViolated

`func (o *ObjectProtectionRunSummary) GetIsSlaViolated() bool`

GetIsSlaViolated returns the IsSlaViolated field if non-nil, zero value otherwise.

### GetIsSlaViolatedOk

`func (o *ObjectProtectionRunSummary) GetIsSlaViolatedOk() (*bool, bool)`

GetIsSlaViolatedOk returns a tuple with the IsSlaViolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSlaViolated

`func (o *ObjectProtectionRunSummary) SetIsSlaViolated(v bool)`

SetIsSlaViolated sets IsSlaViolated field to given value.

### HasIsSlaViolated

`func (o *ObjectProtectionRunSummary) HasIsSlaViolated() bool`

HasIsSlaViolated returns a boolean if a field has been set.

### SetIsSlaViolatedNil

`func (o *ObjectProtectionRunSummary) SetIsSlaViolatedNil(b bool)`

 SetIsSlaViolatedNil sets the value for IsSlaViolated to be an explicit nil

### UnsetIsSlaViolated
`func (o *ObjectProtectionRunSummary) UnsetIsSlaViolated()`

UnsetIsSlaViolated ensures that no value is present for IsSlaViolated, not even an explicit nil
### GetLocalSnapshotInfo

`func (o *ObjectProtectionRunSummary) GetLocalSnapshotInfo() BackupRun`

GetLocalSnapshotInfo returns the LocalSnapshotInfo field if non-nil, zero value otherwise.

### GetLocalSnapshotInfoOk

`func (o *ObjectProtectionRunSummary) GetLocalSnapshotInfoOk() (*BackupRun, bool)`

GetLocalSnapshotInfoOk returns a tuple with the LocalSnapshotInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSnapshotInfo

`func (o *ObjectProtectionRunSummary) SetLocalSnapshotInfo(v BackupRun)`

SetLocalSnapshotInfo sets LocalSnapshotInfo field to given value.

### HasLocalSnapshotInfo

`func (o *ObjectProtectionRunSummary) HasLocalSnapshotInfo() bool`

HasLocalSnapshotInfo returns a boolean if a field has been set.

### GetOnLegalHold

`func (o *ObjectProtectionRunSummary) GetOnLegalHold() bool`

GetOnLegalHold returns the OnLegalHold field if non-nil, zero value otherwise.

### GetOnLegalHoldOk

`func (o *ObjectProtectionRunSummary) GetOnLegalHoldOk() (*bool, bool)`

GetOnLegalHoldOk returns a tuple with the OnLegalHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnLegalHold

`func (o *ObjectProtectionRunSummary) SetOnLegalHold(v bool)`

SetOnLegalHold sets OnLegalHold field to given value.

### HasOnLegalHold

`func (o *ObjectProtectionRunSummary) HasOnLegalHold() bool`

HasOnLegalHold returns a boolean if a field has been set.

### SetOnLegalHoldNil

`func (o *ObjectProtectionRunSummary) SetOnLegalHoldNil(b bool)`

 SetOnLegalHoldNil sets the value for OnLegalHold to be an explicit nil

### UnsetOnLegalHold
`func (o *ObjectProtectionRunSummary) UnsetOnLegalHold()`

UnsetOnLegalHold ensures that no value is present for OnLegalHold, not even an explicit nil
### GetOnPremDeployInfo

`func (o *ObjectProtectionRunSummary) GetOnPremDeployInfo() OnPremDeployRun`

GetOnPremDeployInfo returns the OnPremDeployInfo field if non-nil, zero value otherwise.

### GetOnPremDeployInfoOk

`func (o *ObjectProtectionRunSummary) GetOnPremDeployInfoOk() (*OnPremDeployRun, bool)`

GetOnPremDeployInfoOk returns a tuple with the OnPremDeployInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremDeployInfo

`func (o *ObjectProtectionRunSummary) SetOnPremDeployInfo(v OnPremDeployRun)`

SetOnPremDeployInfo sets OnPremDeployInfo field to given value.

### HasOnPremDeployInfo

`func (o *ObjectProtectionRunSummary) HasOnPremDeployInfo() bool`

HasOnPremDeployInfo returns a boolean if a field has been set.

### GetOriginClusterIdentifier

`func (o *ObjectProtectionRunSummary) GetOriginClusterIdentifier() ClusterIdentifier`

GetOriginClusterIdentifier returns the OriginClusterIdentifier field if non-nil, zero value otherwise.

### GetOriginClusterIdentifierOk

`func (o *ObjectProtectionRunSummary) GetOriginClusterIdentifierOk() (*ClusterIdentifier, bool)`

GetOriginClusterIdentifierOk returns a tuple with the OriginClusterIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginClusterIdentifier

`func (o *ObjectProtectionRunSummary) SetOriginClusterIdentifier(v ClusterIdentifier)`

SetOriginClusterIdentifier sets OriginClusterIdentifier field to given value.

### HasOriginClusterIdentifier

`func (o *ObjectProtectionRunSummary) HasOriginClusterIdentifier() bool`

HasOriginClusterIdentifier returns a boolean if a field has been set.

### GetOriginProtectionGroupId

`func (o *ObjectProtectionRunSummary) GetOriginProtectionGroupId() string`

GetOriginProtectionGroupId returns the OriginProtectionGroupId field if non-nil, zero value otherwise.

### GetOriginProtectionGroupIdOk

`func (o *ObjectProtectionRunSummary) GetOriginProtectionGroupIdOk() (*string, bool)`

GetOriginProtectionGroupIdOk returns a tuple with the OriginProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginProtectionGroupId

`func (o *ObjectProtectionRunSummary) SetOriginProtectionGroupId(v string)`

SetOriginProtectionGroupId sets OriginProtectionGroupId field to given value.

### HasOriginProtectionGroupId

`func (o *ObjectProtectionRunSummary) HasOriginProtectionGroupId() bool`

HasOriginProtectionGroupId returns a boolean if a field has been set.

### SetOriginProtectionGroupIdNil

`func (o *ObjectProtectionRunSummary) SetOriginProtectionGroupIdNil(b bool)`

 SetOriginProtectionGroupIdNil sets the value for OriginProtectionGroupId to be an explicit nil

### UnsetOriginProtectionGroupId
`func (o *ObjectProtectionRunSummary) UnsetOriginProtectionGroupId()`

UnsetOriginProtectionGroupId ensures that no value is present for OriginProtectionGroupId, not even an explicit nil
### GetOriginalBackupInfo

`func (o *ObjectProtectionRunSummary) GetOriginalBackupInfo() BackupRun`

GetOriginalBackupInfo returns the OriginalBackupInfo field if non-nil, zero value otherwise.

### GetOriginalBackupInfoOk

`func (o *ObjectProtectionRunSummary) GetOriginalBackupInfoOk() (*BackupRun, bool)`

GetOriginalBackupInfoOk returns a tuple with the OriginalBackupInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalBackupInfo

`func (o *ObjectProtectionRunSummary) SetOriginalBackupInfo(v BackupRun)`

SetOriginalBackupInfo sets OriginalBackupInfo field to given value.

### HasOriginalBackupInfo

`func (o *ObjectProtectionRunSummary) HasOriginalBackupInfo() bool`

HasOriginalBackupInfo returns a boolean if a field has been set.

### GetPermissions

`func (o *ObjectProtectionRunSummary) GetPermissions() []Tenant`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ObjectProtectionRunSummary) GetPermissionsOk() (*[]Tenant, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ObjectProtectionRunSummary) SetPermissions(v []Tenant)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ObjectProtectionRunSummary) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *ObjectProtectionRunSummary) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *ObjectProtectionRunSummary) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetPolicyId

`func (o *ObjectProtectionRunSummary) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ObjectProtectionRunSummary) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ObjectProtectionRunSummary) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *ObjectProtectionRunSummary) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *ObjectProtectionRunSummary) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *ObjectProtectionRunSummary) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetPolicyName

`func (o *ObjectProtectionRunSummary) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *ObjectProtectionRunSummary) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *ObjectProtectionRunSummary) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *ObjectProtectionRunSummary) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### SetPolicyNameNil

`func (o *ObjectProtectionRunSummary) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *ObjectProtectionRunSummary) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
### GetProtectionGroupId

`func (o *ObjectProtectionRunSummary) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *ObjectProtectionRunSummary) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *ObjectProtectionRunSummary) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *ObjectProtectionRunSummary) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *ObjectProtectionRunSummary) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *ObjectProtectionRunSummary) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetProtectionGroupName

`func (o *ObjectProtectionRunSummary) GetProtectionGroupName() string`

GetProtectionGroupName returns the ProtectionGroupName field if non-nil, zero value otherwise.

### GetProtectionGroupNameOk

`func (o *ObjectProtectionRunSummary) GetProtectionGroupNameOk() (*string, bool)`

GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupName

`func (o *ObjectProtectionRunSummary) SetProtectionGroupName(v string)`

SetProtectionGroupName sets ProtectionGroupName field to given value.

### HasProtectionGroupName

`func (o *ObjectProtectionRunSummary) HasProtectionGroupName() bool`

HasProtectionGroupName returns a boolean if a field has been set.

### SetProtectionGroupNameNil

`func (o *ObjectProtectionRunSummary) SetProtectionGroupNameNil(b bool)`

 SetProtectionGroupNameNil sets the value for ProtectionGroupName to be an explicit nil

### UnsetProtectionGroupName
`func (o *ObjectProtectionRunSummary) UnsetProtectionGroupName()`

UnsetProtectionGroupName ensures that no value is present for ProtectionGroupName, not even an explicit nil
### GetReplicationInfo

`func (o *ObjectProtectionRunSummary) GetReplicationInfo() ReplicationRun`

GetReplicationInfo returns the ReplicationInfo field if non-nil, zero value otherwise.

### GetReplicationInfoOk

`func (o *ObjectProtectionRunSummary) GetReplicationInfoOk() (*ReplicationRun, bool)`

GetReplicationInfoOk returns a tuple with the ReplicationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationInfo

`func (o *ObjectProtectionRunSummary) SetReplicationInfo(v ReplicationRun)`

SetReplicationInfo sets ReplicationInfo field to given value.

### HasReplicationInfo

`func (o *ObjectProtectionRunSummary) HasReplicationInfo() bool`

HasReplicationInfo returns a boolean if a field has been set.

### GetRunId

`func (o *ObjectProtectionRunSummary) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *ObjectProtectionRunSummary) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *ObjectProtectionRunSummary) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *ObjectProtectionRunSummary) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### SetRunIdNil

`func (o *ObjectProtectionRunSummary) SetRunIdNil(b bool)`

 SetRunIdNil sets the value for RunId to be an explicit nil

### UnsetRunId
`func (o *ObjectProtectionRunSummary) UnsetRunId()`

UnsetRunId ensures that no value is present for RunId, not even an explicit nil
### GetRunLabel

`func (o *ObjectProtectionRunSummary) GetRunLabel() string`

GetRunLabel returns the RunLabel field if non-nil, zero value otherwise.

### GetRunLabelOk

`func (o *ObjectProtectionRunSummary) GetRunLabelOk() (*string, bool)`

GetRunLabelOk returns a tuple with the RunLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunLabel

`func (o *ObjectProtectionRunSummary) SetRunLabel(v string)`

SetRunLabel sets RunLabel field to given value.

### HasRunLabel

`func (o *ObjectProtectionRunSummary) HasRunLabel() bool`

HasRunLabel returns a boolean if a field has been set.

### SetRunLabelNil

`func (o *ObjectProtectionRunSummary) SetRunLabelNil(b bool)`

 SetRunLabelNil sets the value for RunLabel to be an explicit nil

### UnsetRunLabel
`func (o *ObjectProtectionRunSummary) UnsetRunLabel()`

UnsetRunLabel ensures that no value is present for RunLabel, not even an explicit nil
### GetRunType

`func (o *ObjectProtectionRunSummary) GetRunType() string`

GetRunType returns the RunType field if non-nil, zero value otherwise.

### GetRunTypeOk

`func (o *ObjectProtectionRunSummary) GetRunTypeOk() (*string, bool)`

GetRunTypeOk returns a tuple with the RunType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunType

`func (o *ObjectProtectionRunSummary) SetRunType(v string)`

SetRunType sets RunType field to given value.

### HasRunType

`func (o *ObjectProtectionRunSummary) HasRunType() bool`

HasRunType returns a boolean if a field has been set.

### SetRunTypeNil

`func (o *ObjectProtectionRunSummary) SetRunTypeNil(b bool)`

 SetRunTypeNil sets the value for RunType to be an explicit nil

### UnsetRunType
`func (o *ObjectProtectionRunSummary) UnsetRunType()`

UnsetRunType ensures that no value is present for RunType, not even an explicit nil
### GetStorageDomainId

`func (o *ObjectProtectionRunSummary) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *ObjectProtectionRunSummary) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *ObjectProtectionRunSummary) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.

### HasStorageDomainId

`func (o *ObjectProtectionRunSummary) HasStorageDomainId() bool`

HasStorageDomainId returns a boolean if a field has been set.

### SetStorageDomainIdNil

`func (o *ObjectProtectionRunSummary) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *ObjectProtectionRunSummary) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


