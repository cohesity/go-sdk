# IndexAndSnapshots

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveTaskUid** | Pointer to [**NullableUniversalId**](UniversalId.md) | Specifies a unique id of the Archive task that originally archived the object to the Vault. | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time as a Unix epoch Timestamp (in microseconds). If set, only index and Snapshots for Protection Job Runs that started before the specified end time are restored. | [optional] 
**RemoteProtectionJobUid** | Pointer to [**NullableUniversalId**](UniversalId.md) | Specifies a unique id assigned to the original Protection Job by the original Cluster that archived data to the remote Vault. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time as a Unix epoch Timestamp (in microseconds). If set, only the index and Snapshots for Protection Job Runs that started after the specified start time are restored. | [optional] 
**ViewBoxId** | Pointer to **NullableInt64** | Specifies the id of the local Storage Domain (View Box) where the index and the Snapshot will be restored to. | [optional] 

## Methods

### NewIndexAndSnapshots

`func NewIndexAndSnapshots() *IndexAndSnapshots`

NewIndexAndSnapshots instantiates a new IndexAndSnapshots object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexAndSnapshotsWithDefaults

`func NewIndexAndSnapshotsWithDefaults() *IndexAndSnapshots`

NewIndexAndSnapshotsWithDefaults instantiates a new IndexAndSnapshots object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveTaskUid

`func (o *IndexAndSnapshots) GetArchiveTaskUid() UniversalId`

GetArchiveTaskUid returns the ArchiveTaskUid field if non-nil, zero value otherwise.

### GetArchiveTaskUidOk

`func (o *IndexAndSnapshots) GetArchiveTaskUidOk() (*UniversalId, bool)`

GetArchiveTaskUidOk returns a tuple with the ArchiveTaskUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveTaskUid

`func (o *IndexAndSnapshots) SetArchiveTaskUid(v UniversalId)`

SetArchiveTaskUid sets ArchiveTaskUid field to given value.

### HasArchiveTaskUid

`func (o *IndexAndSnapshots) HasArchiveTaskUid() bool`

HasArchiveTaskUid returns a boolean if a field has been set.

### SetArchiveTaskUidNil

`func (o *IndexAndSnapshots) SetArchiveTaskUidNil(b bool)`

 SetArchiveTaskUidNil sets the value for ArchiveTaskUid to be an explicit nil

### UnsetArchiveTaskUid
`func (o *IndexAndSnapshots) UnsetArchiveTaskUid()`

UnsetArchiveTaskUid ensures that no value is present for ArchiveTaskUid, not even an explicit nil
### GetEndTimeUsecs

`func (o *IndexAndSnapshots) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *IndexAndSnapshots) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *IndexAndSnapshots) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *IndexAndSnapshots) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *IndexAndSnapshots) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *IndexAndSnapshots) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetRemoteProtectionJobUid

`func (o *IndexAndSnapshots) GetRemoteProtectionJobUid() UniversalId`

GetRemoteProtectionJobUid returns the RemoteProtectionJobUid field if non-nil, zero value otherwise.

### GetRemoteProtectionJobUidOk

`func (o *IndexAndSnapshots) GetRemoteProtectionJobUidOk() (*UniversalId, bool)`

GetRemoteProtectionJobUidOk returns a tuple with the RemoteProtectionJobUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteProtectionJobUid

`func (o *IndexAndSnapshots) SetRemoteProtectionJobUid(v UniversalId)`

SetRemoteProtectionJobUid sets RemoteProtectionJobUid field to given value.

### HasRemoteProtectionJobUid

`func (o *IndexAndSnapshots) HasRemoteProtectionJobUid() bool`

HasRemoteProtectionJobUid returns a boolean if a field has been set.

### SetRemoteProtectionJobUidNil

`func (o *IndexAndSnapshots) SetRemoteProtectionJobUidNil(b bool)`

 SetRemoteProtectionJobUidNil sets the value for RemoteProtectionJobUid to be an explicit nil

### UnsetRemoteProtectionJobUid
`func (o *IndexAndSnapshots) UnsetRemoteProtectionJobUid()`

UnsetRemoteProtectionJobUid ensures that no value is present for RemoteProtectionJobUid, not even an explicit nil
### GetStartTimeUsecs

`func (o *IndexAndSnapshots) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *IndexAndSnapshots) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *IndexAndSnapshots) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *IndexAndSnapshots) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *IndexAndSnapshots) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *IndexAndSnapshots) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetViewBoxId

`func (o *IndexAndSnapshots) GetViewBoxId() int64`

GetViewBoxId returns the ViewBoxId field if non-nil, zero value otherwise.

### GetViewBoxIdOk

`func (o *IndexAndSnapshots) GetViewBoxIdOk() (*int64, bool)`

GetViewBoxIdOk returns a tuple with the ViewBoxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewBoxId

`func (o *IndexAndSnapshots) SetViewBoxId(v int64)`

SetViewBoxId sets ViewBoxId field to given value.

### HasViewBoxId

`func (o *IndexAndSnapshots) HasViewBoxId() bool`

HasViewBoxId returns a boolean if a field has been set.

### SetViewBoxIdNil

`func (o *IndexAndSnapshots) SetViewBoxIdNil(b bool)`

 SetViewBoxIdNil sets the value for ViewBoxId to be an explicit nil

### UnsetViewBoxId
`func (o *IndexAndSnapshots) UnsetViewBoxId()`

UnsetViewBoxId ensures that no value is present for ViewBoxId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


