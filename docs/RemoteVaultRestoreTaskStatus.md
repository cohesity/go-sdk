# RemoteVaultRestoreTaskStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentIndexingStatus** | Pointer to [**NullableRemoteRestoreIndexingStatus**](RemoteRestoreIndexingStatus.md) | Specifies the status of an indexing task that builds an index from the Protection Job metadata retrieved from the remote Vault. The index contains information about Job Runs (Snapshots) for a Protection Job and is required to restore Snapshots to this local Cluster. | [optional] 
**CurrentSnapshotStatus** | Pointer to [**NullableRemoteRestoreSnapshotStatus**](RemoteRestoreSnapshotStatus.md) | Specifies the status of the Snapshot restore task. The Snapshot restore task restores the specified archived Snapshots from a remote Vault to this Cluster. | [optional] 
**LocalProtectionJobUid** | Pointer to [**NullableUniversalId**](UniversalId.md) | Specifies the globally unique id of the new inactive Protection Job created on the local Cluster as part of the restoration of archived data. | [optional] 
**ParentJobUid** | Pointer to [**NullableUniversalId**](UniversalId.md) | Specifies the unique id of the parent Job/task that spawned the indexing and Snapshot restore tasks. | [optional] 
**RemoteProtectionJobInformation** | Pointer to [**RemoteProtectionJobInformation**](RemoteProtectionJobInformation.md) |  | [optional] 
**SearchJobUid** | Pointer to [**NullableUniversalId**](UniversalId.md) | Specifies the unique id of the search Job that searched the remote Vault. | [optional] 

## Methods

### NewRemoteVaultRestoreTaskStatus

`func NewRemoteVaultRestoreTaskStatus() *RemoteVaultRestoreTaskStatus`

NewRemoteVaultRestoreTaskStatus instantiates a new RemoteVaultRestoreTaskStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteVaultRestoreTaskStatusWithDefaults

`func NewRemoteVaultRestoreTaskStatusWithDefaults() *RemoteVaultRestoreTaskStatus`

NewRemoteVaultRestoreTaskStatusWithDefaults instantiates a new RemoteVaultRestoreTaskStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentIndexingStatus

`func (o *RemoteVaultRestoreTaskStatus) GetCurrentIndexingStatus() RemoteRestoreIndexingStatus`

GetCurrentIndexingStatus returns the CurrentIndexingStatus field if non-nil, zero value otherwise.

### GetCurrentIndexingStatusOk

`func (o *RemoteVaultRestoreTaskStatus) GetCurrentIndexingStatusOk() (*RemoteRestoreIndexingStatus, bool)`

GetCurrentIndexingStatusOk returns a tuple with the CurrentIndexingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentIndexingStatus

`func (o *RemoteVaultRestoreTaskStatus) SetCurrentIndexingStatus(v RemoteRestoreIndexingStatus)`

SetCurrentIndexingStatus sets CurrentIndexingStatus field to given value.

### HasCurrentIndexingStatus

`func (o *RemoteVaultRestoreTaskStatus) HasCurrentIndexingStatus() bool`

HasCurrentIndexingStatus returns a boolean if a field has been set.

### SetCurrentIndexingStatusNil

`func (o *RemoteVaultRestoreTaskStatus) SetCurrentIndexingStatusNil(b bool)`

 SetCurrentIndexingStatusNil sets the value for CurrentIndexingStatus to be an explicit nil

### UnsetCurrentIndexingStatus
`func (o *RemoteVaultRestoreTaskStatus) UnsetCurrentIndexingStatus()`

UnsetCurrentIndexingStatus ensures that no value is present for CurrentIndexingStatus, not even an explicit nil
### GetCurrentSnapshotStatus

`func (o *RemoteVaultRestoreTaskStatus) GetCurrentSnapshotStatus() RemoteRestoreSnapshotStatus`

GetCurrentSnapshotStatus returns the CurrentSnapshotStatus field if non-nil, zero value otherwise.

### GetCurrentSnapshotStatusOk

`func (o *RemoteVaultRestoreTaskStatus) GetCurrentSnapshotStatusOk() (*RemoteRestoreSnapshotStatus, bool)`

GetCurrentSnapshotStatusOk returns a tuple with the CurrentSnapshotStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSnapshotStatus

`func (o *RemoteVaultRestoreTaskStatus) SetCurrentSnapshotStatus(v RemoteRestoreSnapshotStatus)`

SetCurrentSnapshotStatus sets CurrentSnapshotStatus field to given value.

### HasCurrentSnapshotStatus

`func (o *RemoteVaultRestoreTaskStatus) HasCurrentSnapshotStatus() bool`

HasCurrentSnapshotStatus returns a boolean if a field has been set.

### SetCurrentSnapshotStatusNil

`func (o *RemoteVaultRestoreTaskStatus) SetCurrentSnapshotStatusNil(b bool)`

 SetCurrentSnapshotStatusNil sets the value for CurrentSnapshotStatus to be an explicit nil

### UnsetCurrentSnapshotStatus
`func (o *RemoteVaultRestoreTaskStatus) UnsetCurrentSnapshotStatus()`

UnsetCurrentSnapshotStatus ensures that no value is present for CurrentSnapshotStatus, not even an explicit nil
### GetLocalProtectionJobUid

`func (o *RemoteVaultRestoreTaskStatus) GetLocalProtectionJobUid() UniversalId`

GetLocalProtectionJobUid returns the LocalProtectionJobUid field if non-nil, zero value otherwise.

### GetLocalProtectionJobUidOk

`func (o *RemoteVaultRestoreTaskStatus) GetLocalProtectionJobUidOk() (*UniversalId, bool)`

GetLocalProtectionJobUidOk returns a tuple with the LocalProtectionJobUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalProtectionJobUid

`func (o *RemoteVaultRestoreTaskStatus) SetLocalProtectionJobUid(v UniversalId)`

SetLocalProtectionJobUid sets LocalProtectionJobUid field to given value.

### HasLocalProtectionJobUid

`func (o *RemoteVaultRestoreTaskStatus) HasLocalProtectionJobUid() bool`

HasLocalProtectionJobUid returns a boolean if a field has been set.

### SetLocalProtectionJobUidNil

`func (o *RemoteVaultRestoreTaskStatus) SetLocalProtectionJobUidNil(b bool)`

 SetLocalProtectionJobUidNil sets the value for LocalProtectionJobUid to be an explicit nil

### UnsetLocalProtectionJobUid
`func (o *RemoteVaultRestoreTaskStatus) UnsetLocalProtectionJobUid()`

UnsetLocalProtectionJobUid ensures that no value is present for LocalProtectionJobUid, not even an explicit nil
### GetParentJobUid

`func (o *RemoteVaultRestoreTaskStatus) GetParentJobUid() UniversalId`

GetParentJobUid returns the ParentJobUid field if non-nil, zero value otherwise.

### GetParentJobUidOk

`func (o *RemoteVaultRestoreTaskStatus) GetParentJobUidOk() (*UniversalId, bool)`

GetParentJobUidOk returns a tuple with the ParentJobUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentJobUid

`func (o *RemoteVaultRestoreTaskStatus) SetParentJobUid(v UniversalId)`

SetParentJobUid sets ParentJobUid field to given value.

### HasParentJobUid

`func (o *RemoteVaultRestoreTaskStatus) HasParentJobUid() bool`

HasParentJobUid returns a boolean if a field has been set.

### SetParentJobUidNil

`func (o *RemoteVaultRestoreTaskStatus) SetParentJobUidNil(b bool)`

 SetParentJobUidNil sets the value for ParentJobUid to be an explicit nil

### UnsetParentJobUid
`func (o *RemoteVaultRestoreTaskStatus) UnsetParentJobUid()`

UnsetParentJobUid ensures that no value is present for ParentJobUid, not even an explicit nil
### GetRemoteProtectionJobInformation

`func (o *RemoteVaultRestoreTaskStatus) GetRemoteProtectionJobInformation() RemoteProtectionJobInformation`

GetRemoteProtectionJobInformation returns the RemoteProtectionJobInformation field if non-nil, zero value otherwise.

### GetRemoteProtectionJobInformationOk

`func (o *RemoteVaultRestoreTaskStatus) GetRemoteProtectionJobInformationOk() (*RemoteProtectionJobInformation, bool)`

GetRemoteProtectionJobInformationOk returns a tuple with the RemoteProtectionJobInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteProtectionJobInformation

`func (o *RemoteVaultRestoreTaskStatus) SetRemoteProtectionJobInformation(v RemoteProtectionJobInformation)`

SetRemoteProtectionJobInformation sets RemoteProtectionJobInformation field to given value.

### HasRemoteProtectionJobInformation

`func (o *RemoteVaultRestoreTaskStatus) HasRemoteProtectionJobInformation() bool`

HasRemoteProtectionJobInformation returns a boolean if a field has been set.

### GetSearchJobUid

`func (o *RemoteVaultRestoreTaskStatus) GetSearchJobUid() UniversalId`

GetSearchJobUid returns the SearchJobUid field if non-nil, zero value otherwise.

### GetSearchJobUidOk

`func (o *RemoteVaultRestoreTaskStatus) GetSearchJobUidOk() (*UniversalId, bool)`

GetSearchJobUidOk returns a tuple with the SearchJobUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchJobUid

`func (o *RemoteVaultRestoreTaskStatus) SetSearchJobUid(v UniversalId)`

SetSearchJobUid sets SearchJobUid field to given value.

### HasSearchJobUid

`func (o *RemoteVaultRestoreTaskStatus) HasSearchJobUid() bool`

HasSearchJobUid returns a boolean if a field has been set.

### SetSearchJobUidNil

`func (o *RemoteVaultRestoreTaskStatus) SetSearchJobUidNil(b bool)`

 SetSearchJobUidNil sets the value for SearchJobUid to be an explicit nil

### UnsetSearchJobUid
`func (o *RemoteVaultRestoreTaskStatus) UnsetSearchJobUid()`

UnsetSearchJobUid ensures that no value is present for SearchJobUid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


