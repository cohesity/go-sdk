# UpdateProtectionJobRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CopyRunTargets** | Pointer to [**[]RunJobSnapshotTarget**](RunJobSnapshotTarget.md) | Specifies the retention for archival, replication or extended local retention. | [optional] 
**JobUid** | Pointer to [**NullableUniversalId**](UniversalId.md) | Specifies a unique universal id for the Job. | [optional] 
**RunStartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of the Job Run to update. The start time is specified as a Unix epoch Timestamp (in microseconds). This uniquely identifies a snapshot. This parameter is required. | [optional] 
**SourceIds** | Pointer to **[]int64** | Ids of the Protection Sources. If this is specified, retention time will only be updated for the sources specified. | [optional] 

## Methods

### NewUpdateProtectionJobRun

`func NewUpdateProtectionJobRun() *UpdateProtectionJobRun`

NewUpdateProtectionJobRun instantiates a new UpdateProtectionJobRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProtectionJobRunWithDefaults

`func NewUpdateProtectionJobRunWithDefaults() *UpdateProtectionJobRun`

NewUpdateProtectionJobRunWithDefaults instantiates a new UpdateProtectionJobRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCopyRunTargets

`func (o *UpdateProtectionJobRun) GetCopyRunTargets() []RunJobSnapshotTarget`

GetCopyRunTargets returns the CopyRunTargets field if non-nil, zero value otherwise.

### GetCopyRunTargetsOk

`func (o *UpdateProtectionJobRun) GetCopyRunTargetsOk() (*[]RunJobSnapshotTarget, bool)`

GetCopyRunTargetsOk returns a tuple with the CopyRunTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyRunTargets

`func (o *UpdateProtectionJobRun) SetCopyRunTargets(v []RunJobSnapshotTarget)`

SetCopyRunTargets sets CopyRunTargets field to given value.

### HasCopyRunTargets

`func (o *UpdateProtectionJobRun) HasCopyRunTargets() bool`

HasCopyRunTargets returns a boolean if a field has been set.

### SetCopyRunTargetsNil

`func (o *UpdateProtectionJobRun) SetCopyRunTargetsNil(b bool)`

 SetCopyRunTargetsNil sets the value for CopyRunTargets to be an explicit nil

### UnsetCopyRunTargets
`func (o *UpdateProtectionJobRun) UnsetCopyRunTargets()`

UnsetCopyRunTargets ensures that no value is present for CopyRunTargets, not even an explicit nil
### GetJobUid

`func (o *UpdateProtectionJobRun) GetJobUid() UniversalId`

GetJobUid returns the JobUid field if non-nil, zero value otherwise.

### GetJobUidOk

`func (o *UpdateProtectionJobRun) GetJobUidOk() (*UniversalId, bool)`

GetJobUidOk returns a tuple with the JobUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobUid

`func (o *UpdateProtectionJobRun) SetJobUid(v UniversalId)`

SetJobUid sets JobUid field to given value.

### HasJobUid

`func (o *UpdateProtectionJobRun) HasJobUid() bool`

HasJobUid returns a boolean if a field has been set.

### SetJobUidNil

`func (o *UpdateProtectionJobRun) SetJobUidNil(b bool)`

 SetJobUidNil sets the value for JobUid to be an explicit nil

### UnsetJobUid
`func (o *UpdateProtectionJobRun) UnsetJobUid()`

UnsetJobUid ensures that no value is present for JobUid, not even an explicit nil
### GetRunStartTimeUsecs

`func (o *UpdateProtectionJobRun) GetRunStartTimeUsecs() int64`

GetRunStartTimeUsecs returns the RunStartTimeUsecs field if non-nil, zero value otherwise.

### GetRunStartTimeUsecsOk

`func (o *UpdateProtectionJobRun) GetRunStartTimeUsecsOk() (*int64, bool)`

GetRunStartTimeUsecsOk returns a tuple with the RunStartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunStartTimeUsecs

`func (o *UpdateProtectionJobRun) SetRunStartTimeUsecs(v int64)`

SetRunStartTimeUsecs sets RunStartTimeUsecs field to given value.

### HasRunStartTimeUsecs

`func (o *UpdateProtectionJobRun) HasRunStartTimeUsecs() bool`

HasRunStartTimeUsecs returns a boolean if a field has been set.

### SetRunStartTimeUsecsNil

`func (o *UpdateProtectionJobRun) SetRunStartTimeUsecsNil(b bool)`

 SetRunStartTimeUsecsNil sets the value for RunStartTimeUsecs to be an explicit nil

### UnsetRunStartTimeUsecs
`func (o *UpdateProtectionJobRun) UnsetRunStartTimeUsecs()`

UnsetRunStartTimeUsecs ensures that no value is present for RunStartTimeUsecs, not even an explicit nil
### GetSourceIds

`func (o *UpdateProtectionJobRun) GetSourceIds() []int64`

GetSourceIds returns the SourceIds field if non-nil, zero value otherwise.

### GetSourceIdsOk

`func (o *UpdateProtectionJobRun) GetSourceIdsOk() (*[]int64, bool)`

GetSourceIdsOk returns a tuple with the SourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIds

`func (o *UpdateProtectionJobRun) SetSourceIds(v []int64)`

SetSourceIds sets SourceIds field to given value.

### HasSourceIds

`func (o *UpdateProtectionJobRun) HasSourceIds() bool`

HasSourceIds returns a boolean if a field has been set.

### SetSourceIdsNil

`func (o *UpdateProtectionJobRun) SetSourceIdsNil(b bool)`

 SetSourceIdsNil sets the value for SourceIds to be an explicit nil

### UnsetSourceIds
`func (o *UpdateProtectionJobRun) UnsetSourceIds()`

UnsetSourceIds ensures that no value is present for SourceIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


