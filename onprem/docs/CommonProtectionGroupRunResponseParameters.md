# CommonProtectionGroupRunResponseParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Specifies the ID of the Protection Group run. | [optional] 
**ProtectionGroupInstanceId** | Pointer to **NullableInt64** | Protection Group instance Id. This field will be removed later. | [optional] 
**ProtectionGroupId** | Pointer to **NullableString** | ProtectionGroupId to which this run belongs. | [optional] 
**IsReplicationRun** | Pointer to **NullableBool** | Specifies if this protection run is a replication run. | [optional] 
**OriginClusterIdentifier** | Pointer to [**ClusterIdentifier**](ClusterIdentifier.md) |  | [optional] 
**OriginProtectionGroupId** | Pointer to **NullableString** | ProtectionGroupId to which this run belongs on the primary cluster if this run is a replication run. | [optional] 
**ProtectionGroupName** | Pointer to **NullableString** | Name of the Protection Group to which this run belongs. | [optional] 
**IsLocalSnapshotsDeleted** | Pointer to **NullableBool** | Specifies if snapshots for this run has been deleted. | [optional] 
**Objects** | Pointer to [**[]ObjectRunResult**](ObjectRunResult.md) | Snapahot, replication, archival results for each object. | [optional] 
**LocalBackupInfo** | Pointer to [**BackupRunSummary**](BackupRunSummary.md) |  | [optional] 
**OriginalBackupInfo** | Pointer to [**BackupRunSummary**](BackupRunSummary.md) |  | [optional] 
**ReplicationInfo** | Pointer to [**ReplicationRunSummary**](ReplicationRunSummary.md) |  | [optional] 
**ArchivalInfo** | Pointer to [**ArchivalRunSummary**](ArchivalRunSummary.md) |  | [optional] 
**CloudSpinInfo** | Pointer to [**CloudSpinRunSummary**](CloudSpinRunSummary.md) |  | [optional] 
**OnLegalHold** | Pointer to **NullableBool** | Specifies if the Protection Run is on legal hold. | [optional] 
**Permissions** | Pointer to [**[]Tenant**](Tenant.md) | Specifies the list of tenants that have permissions for this protection group run. | [optional] 
**IsCloudArchivalDirect** | Pointer to **NullableBool** | Specifies whether the run is a CAD run if cloud archive direct feature is enabled. If this field is true, the primary backup copy will only be available at the given archived location. | [optional] 
**HasLocalSnapshot** | Pointer to **NullableBool** | Specifies whether the run has a local snapshot. For cloud retrieved runs there may not be local snapshots. | [optional] 
**Environment** | Pointer to **NullableString** | Specifies the environment of the Protection Group. | [optional] 
**ExternallyTriggeredBackupTag** | Pointer to **NullableString** | The tag of externally triggered backup job. | [optional] 

## Methods

### NewCommonProtectionGroupRunResponseParameters

`func NewCommonProtectionGroupRunResponseParameters() *CommonProtectionGroupRunResponseParameters`

NewCommonProtectionGroupRunResponseParameters instantiates a new CommonProtectionGroupRunResponseParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonProtectionGroupRunResponseParametersWithDefaults

`func NewCommonProtectionGroupRunResponseParametersWithDefaults() *CommonProtectionGroupRunResponseParameters`

NewCommonProtectionGroupRunResponseParametersWithDefaults instantiates a new CommonProtectionGroupRunResponseParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommonProtectionGroupRunResponseParameters) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommonProtectionGroupRunResponseParameters) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommonProtectionGroupRunResponseParameters) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CommonProtectionGroupRunResponseParameters) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CommonProtectionGroupRunResponseParameters) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CommonProtectionGroupRunResponseParameters) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetProtectionGroupInstanceId

`func (o *CommonProtectionGroupRunResponseParameters) GetProtectionGroupInstanceId() int64`

GetProtectionGroupInstanceId returns the ProtectionGroupInstanceId field if non-nil, zero value otherwise.

### GetProtectionGroupInstanceIdOk

`func (o *CommonProtectionGroupRunResponseParameters) GetProtectionGroupInstanceIdOk() (*int64, bool)`

GetProtectionGroupInstanceIdOk returns a tuple with the ProtectionGroupInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupInstanceId

`func (o *CommonProtectionGroupRunResponseParameters) SetProtectionGroupInstanceId(v int64)`

SetProtectionGroupInstanceId sets ProtectionGroupInstanceId field to given value.

### HasProtectionGroupInstanceId

`func (o *CommonProtectionGroupRunResponseParameters) HasProtectionGroupInstanceId() bool`

HasProtectionGroupInstanceId returns a boolean if a field has been set.

### SetProtectionGroupInstanceIdNil

`func (o *CommonProtectionGroupRunResponseParameters) SetProtectionGroupInstanceIdNil(b bool)`

 SetProtectionGroupInstanceIdNil sets the value for ProtectionGroupInstanceId to be an explicit nil

### UnsetProtectionGroupInstanceId
`func (o *CommonProtectionGroupRunResponseParameters) UnsetProtectionGroupInstanceId()`

UnsetProtectionGroupInstanceId ensures that no value is present for ProtectionGroupInstanceId, not even an explicit nil
### GetProtectionGroupId

`func (o *CommonProtectionGroupRunResponseParameters) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *CommonProtectionGroupRunResponseParameters) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *CommonProtectionGroupRunResponseParameters) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *CommonProtectionGroupRunResponseParameters) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *CommonProtectionGroupRunResponseParameters) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *CommonProtectionGroupRunResponseParameters) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetIsReplicationRun

`func (o *CommonProtectionGroupRunResponseParameters) GetIsReplicationRun() bool`

GetIsReplicationRun returns the IsReplicationRun field if non-nil, zero value otherwise.

### GetIsReplicationRunOk

`func (o *CommonProtectionGroupRunResponseParameters) GetIsReplicationRunOk() (*bool, bool)`

GetIsReplicationRunOk returns a tuple with the IsReplicationRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReplicationRun

`func (o *CommonProtectionGroupRunResponseParameters) SetIsReplicationRun(v bool)`

SetIsReplicationRun sets IsReplicationRun field to given value.

### HasIsReplicationRun

`func (o *CommonProtectionGroupRunResponseParameters) HasIsReplicationRun() bool`

HasIsReplicationRun returns a boolean if a field has been set.

### SetIsReplicationRunNil

`func (o *CommonProtectionGroupRunResponseParameters) SetIsReplicationRunNil(b bool)`

 SetIsReplicationRunNil sets the value for IsReplicationRun to be an explicit nil

### UnsetIsReplicationRun
`func (o *CommonProtectionGroupRunResponseParameters) UnsetIsReplicationRun()`

UnsetIsReplicationRun ensures that no value is present for IsReplicationRun, not even an explicit nil
### GetOriginClusterIdentifier

`func (o *CommonProtectionGroupRunResponseParameters) GetOriginClusterIdentifier() ClusterIdentifier`

GetOriginClusterIdentifier returns the OriginClusterIdentifier field if non-nil, zero value otherwise.

### GetOriginClusterIdentifierOk

`func (o *CommonProtectionGroupRunResponseParameters) GetOriginClusterIdentifierOk() (*ClusterIdentifier, bool)`

GetOriginClusterIdentifierOk returns a tuple with the OriginClusterIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginClusterIdentifier

`func (o *CommonProtectionGroupRunResponseParameters) SetOriginClusterIdentifier(v ClusterIdentifier)`

SetOriginClusterIdentifier sets OriginClusterIdentifier field to given value.

### HasOriginClusterIdentifier

`func (o *CommonProtectionGroupRunResponseParameters) HasOriginClusterIdentifier() bool`

HasOriginClusterIdentifier returns a boolean if a field has been set.

### GetOriginProtectionGroupId

`func (o *CommonProtectionGroupRunResponseParameters) GetOriginProtectionGroupId() string`

GetOriginProtectionGroupId returns the OriginProtectionGroupId field if non-nil, zero value otherwise.

### GetOriginProtectionGroupIdOk

`func (o *CommonProtectionGroupRunResponseParameters) GetOriginProtectionGroupIdOk() (*string, bool)`

GetOriginProtectionGroupIdOk returns a tuple with the OriginProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginProtectionGroupId

`func (o *CommonProtectionGroupRunResponseParameters) SetOriginProtectionGroupId(v string)`

SetOriginProtectionGroupId sets OriginProtectionGroupId field to given value.

### HasOriginProtectionGroupId

`func (o *CommonProtectionGroupRunResponseParameters) HasOriginProtectionGroupId() bool`

HasOriginProtectionGroupId returns a boolean if a field has been set.

### SetOriginProtectionGroupIdNil

`func (o *CommonProtectionGroupRunResponseParameters) SetOriginProtectionGroupIdNil(b bool)`

 SetOriginProtectionGroupIdNil sets the value for OriginProtectionGroupId to be an explicit nil

### UnsetOriginProtectionGroupId
`func (o *CommonProtectionGroupRunResponseParameters) UnsetOriginProtectionGroupId()`

UnsetOriginProtectionGroupId ensures that no value is present for OriginProtectionGroupId, not even an explicit nil
### GetProtectionGroupName

`func (o *CommonProtectionGroupRunResponseParameters) GetProtectionGroupName() string`

GetProtectionGroupName returns the ProtectionGroupName field if non-nil, zero value otherwise.

### GetProtectionGroupNameOk

`func (o *CommonProtectionGroupRunResponseParameters) GetProtectionGroupNameOk() (*string, bool)`

GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupName

`func (o *CommonProtectionGroupRunResponseParameters) SetProtectionGroupName(v string)`

SetProtectionGroupName sets ProtectionGroupName field to given value.

### HasProtectionGroupName

`func (o *CommonProtectionGroupRunResponseParameters) HasProtectionGroupName() bool`

HasProtectionGroupName returns a boolean if a field has been set.

### SetProtectionGroupNameNil

`func (o *CommonProtectionGroupRunResponseParameters) SetProtectionGroupNameNil(b bool)`

 SetProtectionGroupNameNil sets the value for ProtectionGroupName to be an explicit nil

### UnsetProtectionGroupName
`func (o *CommonProtectionGroupRunResponseParameters) UnsetProtectionGroupName()`

UnsetProtectionGroupName ensures that no value is present for ProtectionGroupName, not even an explicit nil
### GetIsLocalSnapshotsDeleted

`func (o *CommonProtectionGroupRunResponseParameters) GetIsLocalSnapshotsDeleted() bool`

GetIsLocalSnapshotsDeleted returns the IsLocalSnapshotsDeleted field if non-nil, zero value otherwise.

### GetIsLocalSnapshotsDeletedOk

`func (o *CommonProtectionGroupRunResponseParameters) GetIsLocalSnapshotsDeletedOk() (*bool, bool)`

GetIsLocalSnapshotsDeletedOk returns a tuple with the IsLocalSnapshotsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocalSnapshotsDeleted

`func (o *CommonProtectionGroupRunResponseParameters) SetIsLocalSnapshotsDeleted(v bool)`

SetIsLocalSnapshotsDeleted sets IsLocalSnapshotsDeleted field to given value.

### HasIsLocalSnapshotsDeleted

`func (o *CommonProtectionGroupRunResponseParameters) HasIsLocalSnapshotsDeleted() bool`

HasIsLocalSnapshotsDeleted returns a boolean if a field has been set.

### SetIsLocalSnapshotsDeletedNil

`func (o *CommonProtectionGroupRunResponseParameters) SetIsLocalSnapshotsDeletedNil(b bool)`

 SetIsLocalSnapshotsDeletedNil sets the value for IsLocalSnapshotsDeleted to be an explicit nil

### UnsetIsLocalSnapshotsDeleted
`func (o *CommonProtectionGroupRunResponseParameters) UnsetIsLocalSnapshotsDeleted()`

UnsetIsLocalSnapshotsDeleted ensures that no value is present for IsLocalSnapshotsDeleted, not even an explicit nil
### GetObjects

`func (o *CommonProtectionGroupRunResponseParameters) GetObjects() []ObjectRunResult`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *CommonProtectionGroupRunResponseParameters) GetObjectsOk() (*[]ObjectRunResult, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *CommonProtectionGroupRunResponseParameters) SetObjects(v []ObjectRunResult)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *CommonProtectionGroupRunResponseParameters) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetLocalBackupInfo

`func (o *CommonProtectionGroupRunResponseParameters) GetLocalBackupInfo() BackupRunSummary`

GetLocalBackupInfo returns the LocalBackupInfo field if non-nil, zero value otherwise.

### GetLocalBackupInfoOk

`func (o *CommonProtectionGroupRunResponseParameters) GetLocalBackupInfoOk() (*BackupRunSummary, bool)`

GetLocalBackupInfoOk returns a tuple with the LocalBackupInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalBackupInfo

`func (o *CommonProtectionGroupRunResponseParameters) SetLocalBackupInfo(v BackupRunSummary)`

SetLocalBackupInfo sets LocalBackupInfo field to given value.

### HasLocalBackupInfo

`func (o *CommonProtectionGroupRunResponseParameters) HasLocalBackupInfo() bool`

HasLocalBackupInfo returns a boolean if a field has been set.

### GetOriginalBackupInfo

`func (o *CommonProtectionGroupRunResponseParameters) GetOriginalBackupInfo() BackupRunSummary`

GetOriginalBackupInfo returns the OriginalBackupInfo field if non-nil, zero value otherwise.

### GetOriginalBackupInfoOk

`func (o *CommonProtectionGroupRunResponseParameters) GetOriginalBackupInfoOk() (*BackupRunSummary, bool)`

GetOriginalBackupInfoOk returns a tuple with the OriginalBackupInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalBackupInfo

`func (o *CommonProtectionGroupRunResponseParameters) SetOriginalBackupInfo(v BackupRunSummary)`

SetOriginalBackupInfo sets OriginalBackupInfo field to given value.

### HasOriginalBackupInfo

`func (o *CommonProtectionGroupRunResponseParameters) HasOriginalBackupInfo() bool`

HasOriginalBackupInfo returns a boolean if a field has been set.

### GetReplicationInfo

`func (o *CommonProtectionGroupRunResponseParameters) GetReplicationInfo() ReplicationRunSummary`

GetReplicationInfo returns the ReplicationInfo field if non-nil, zero value otherwise.

### GetReplicationInfoOk

`func (o *CommonProtectionGroupRunResponseParameters) GetReplicationInfoOk() (*ReplicationRunSummary, bool)`

GetReplicationInfoOk returns a tuple with the ReplicationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationInfo

`func (o *CommonProtectionGroupRunResponseParameters) SetReplicationInfo(v ReplicationRunSummary)`

SetReplicationInfo sets ReplicationInfo field to given value.

### HasReplicationInfo

`func (o *CommonProtectionGroupRunResponseParameters) HasReplicationInfo() bool`

HasReplicationInfo returns a boolean if a field has been set.

### GetArchivalInfo

`func (o *CommonProtectionGroupRunResponseParameters) GetArchivalInfo() ArchivalRunSummary`

GetArchivalInfo returns the ArchivalInfo field if non-nil, zero value otherwise.

### GetArchivalInfoOk

`func (o *CommonProtectionGroupRunResponseParameters) GetArchivalInfoOk() (*ArchivalRunSummary, bool)`

GetArchivalInfoOk returns a tuple with the ArchivalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalInfo

`func (o *CommonProtectionGroupRunResponseParameters) SetArchivalInfo(v ArchivalRunSummary)`

SetArchivalInfo sets ArchivalInfo field to given value.

### HasArchivalInfo

`func (o *CommonProtectionGroupRunResponseParameters) HasArchivalInfo() bool`

HasArchivalInfo returns a boolean if a field has been set.

### GetCloudSpinInfo

`func (o *CommonProtectionGroupRunResponseParameters) GetCloudSpinInfo() CloudSpinRunSummary`

GetCloudSpinInfo returns the CloudSpinInfo field if non-nil, zero value otherwise.

### GetCloudSpinInfoOk

`func (o *CommonProtectionGroupRunResponseParameters) GetCloudSpinInfoOk() (*CloudSpinRunSummary, bool)`

GetCloudSpinInfoOk returns a tuple with the CloudSpinInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudSpinInfo

`func (o *CommonProtectionGroupRunResponseParameters) SetCloudSpinInfo(v CloudSpinRunSummary)`

SetCloudSpinInfo sets CloudSpinInfo field to given value.

### HasCloudSpinInfo

`func (o *CommonProtectionGroupRunResponseParameters) HasCloudSpinInfo() bool`

HasCloudSpinInfo returns a boolean if a field has been set.

### GetOnLegalHold

`func (o *CommonProtectionGroupRunResponseParameters) GetOnLegalHold() bool`

GetOnLegalHold returns the OnLegalHold field if non-nil, zero value otherwise.

### GetOnLegalHoldOk

`func (o *CommonProtectionGroupRunResponseParameters) GetOnLegalHoldOk() (*bool, bool)`

GetOnLegalHoldOk returns a tuple with the OnLegalHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnLegalHold

`func (o *CommonProtectionGroupRunResponseParameters) SetOnLegalHold(v bool)`

SetOnLegalHold sets OnLegalHold field to given value.

### HasOnLegalHold

`func (o *CommonProtectionGroupRunResponseParameters) HasOnLegalHold() bool`

HasOnLegalHold returns a boolean if a field has been set.

### SetOnLegalHoldNil

`func (o *CommonProtectionGroupRunResponseParameters) SetOnLegalHoldNil(b bool)`

 SetOnLegalHoldNil sets the value for OnLegalHold to be an explicit nil

### UnsetOnLegalHold
`func (o *CommonProtectionGroupRunResponseParameters) UnsetOnLegalHold()`

UnsetOnLegalHold ensures that no value is present for OnLegalHold, not even an explicit nil
### GetPermissions

`func (o *CommonProtectionGroupRunResponseParameters) GetPermissions() []Tenant`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CommonProtectionGroupRunResponseParameters) GetPermissionsOk() (*[]Tenant, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CommonProtectionGroupRunResponseParameters) SetPermissions(v []Tenant)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CommonProtectionGroupRunResponseParameters) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *CommonProtectionGroupRunResponseParameters) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *CommonProtectionGroupRunResponseParameters) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetIsCloudArchivalDirect

`func (o *CommonProtectionGroupRunResponseParameters) GetIsCloudArchivalDirect() bool`

GetIsCloudArchivalDirect returns the IsCloudArchivalDirect field if non-nil, zero value otherwise.

### GetIsCloudArchivalDirectOk

`func (o *CommonProtectionGroupRunResponseParameters) GetIsCloudArchivalDirectOk() (*bool, bool)`

GetIsCloudArchivalDirectOk returns a tuple with the IsCloudArchivalDirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloudArchivalDirect

`func (o *CommonProtectionGroupRunResponseParameters) SetIsCloudArchivalDirect(v bool)`

SetIsCloudArchivalDirect sets IsCloudArchivalDirect field to given value.

### HasIsCloudArchivalDirect

`func (o *CommonProtectionGroupRunResponseParameters) HasIsCloudArchivalDirect() bool`

HasIsCloudArchivalDirect returns a boolean if a field has been set.

### SetIsCloudArchivalDirectNil

`func (o *CommonProtectionGroupRunResponseParameters) SetIsCloudArchivalDirectNil(b bool)`

 SetIsCloudArchivalDirectNil sets the value for IsCloudArchivalDirect to be an explicit nil

### UnsetIsCloudArchivalDirect
`func (o *CommonProtectionGroupRunResponseParameters) UnsetIsCloudArchivalDirect()`

UnsetIsCloudArchivalDirect ensures that no value is present for IsCloudArchivalDirect, not even an explicit nil
### GetHasLocalSnapshot

`func (o *CommonProtectionGroupRunResponseParameters) GetHasLocalSnapshot() bool`

GetHasLocalSnapshot returns the HasLocalSnapshot field if non-nil, zero value otherwise.

### GetHasLocalSnapshotOk

`func (o *CommonProtectionGroupRunResponseParameters) GetHasLocalSnapshotOk() (*bool, bool)`

GetHasLocalSnapshotOk returns a tuple with the HasLocalSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLocalSnapshot

`func (o *CommonProtectionGroupRunResponseParameters) SetHasLocalSnapshot(v bool)`

SetHasLocalSnapshot sets HasLocalSnapshot field to given value.

### HasHasLocalSnapshot

`func (o *CommonProtectionGroupRunResponseParameters) HasHasLocalSnapshot() bool`

HasHasLocalSnapshot returns a boolean if a field has been set.

### SetHasLocalSnapshotNil

`func (o *CommonProtectionGroupRunResponseParameters) SetHasLocalSnapshotNil(b bool)`

 SetHasLocalSnapshotNil sets the value for HasLocalSnapshot to be an explicit nil

### UnsetHasLocalSnapshot
`func (o *CommonProtectionGroupRunResponseParameters) UnsetHasLocalSnapshot()`

UnsetHasLocalSnapshot ensures that no value is present for HasLocalSnapshot, not even an explicit nil
### GetEnvironment

`func (o *CommonProtectionGroupRunResponseParameters) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CommonProtectionGroupRunResponseParameters) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CommonProtectionGroupRunResponseParameters) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CommonProtectionGroupRunResponseParameters) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *CommonProtectionGroupRunResponseParameters) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *CommonProtectionGroupRunResponseParameters) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetExternallyTriggeredBackupTag

`func (o *CommonProtectionGroupRunResponseParameters) GetExternallyTriggeredBackupTag() string`

GetExternallyTriggeredBackupTag returns the ExternallyTriggeredBackupTag field if non-nil, zero value otherwise.

### GetExternallyTriggeredBackupTagOk

`func (o *CommonProtectionGroupRunResponseParameters) GetExternallyTriggeredBackupTagOk() (*string, bool)`

GetExternallyTriggeredBackupTagOk returns a tuple with the ExternallyTriggeredBackupTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternallyTriggeredBackupTag

`func (o *CommonProtectionGroupRunResponseParameters) SetExternallyTriggeredBackupTag(v string)`

SetExternallyTriggeredBackupTag sets ExternallyTriggeredBackupTag field to given value.

### HasExternallyTriggeredBackupTag

`func (o *CommonProtectionGroupRunResponseParameters) HasExternallyTriggeredBackupTag() bool`

HasExternallyTriggeredBackupTag returns a boolean if a field has been set.

### SetExternallyTriggeredBackupTagNil

`func (o *CommonProtectionGroupRunResponseParameters) SetExternallyTriggeredBackupTagNil(b bool)`

 SetExternallyTriggeredBackupTagNil sets the value for ExternallyTriggeredBackupTag to be an explicit nil

### UnsetExternallyTriggeredBackupTag
`func (o *CommonProtectionGroupRunResponseParameters) UnsetExternallyTriggeredBackupTag()`

UnsetExternallyTriggeredBackupTag ensures that no value is present for ExternallyTriggeredBackupTag, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


