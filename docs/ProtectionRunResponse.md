# ProtectionRunResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchivalRuns** | Pointer to [**[]LatestProtectionJobRunInfo**](LatestProtectionJobRunInfo.md) | Specifies the list of archival job information. | [optional] 
**BackupRuns** | Pointer to [**[]LatestProtectionJobRunInfo**](LatestProtectionJobRunInfo.md) | Specifies the list of local backup job information. | [optional] 
**ReplicationRuns** | Pointer to [**[]LatestProtectionJobRunInfo**](LatestProtectionJobRunInfo.md) | Specifies the list of replication job information. | [optional] 

## Methods

### NewProtectionRunResponse

`func NewProtectionRunResponse() *ProtectionRunResponse`

NewProtectionRunResponse instantiates a new ProtectionRunResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionRunResponseWithDefaults

`func NewProtectionRunResponseWithDefaults() *ProtectionRunResponse`

NewProtectionRunResponseWithDefaults instantiates a new ProtectionRunResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchivalRuns

`func (o *ProtectionRunResponse) GetArchivalRuns() []LatestProtectionJobRunInfo`

GetArchivalRuns returns the ArchivalRuns field if non-nil, zero value otherwise.

### GetArchivalRunsOk

`func (o *ProtectionRunResponse) GetArchivalRunsOk() (*[]LatestProtectionJobRunInfo, bool)`

GetArchivalRunsOk returns a tuple with the ArchivalRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalRuns

`func (o *ProtectionRunResponse) SetArchivalRuns(v []LatestProtectionJobRunInfo)`

SetArchivalRuns sets ArchivalRuns field to given value.

### HasArchivalRuns

`func (o *ProtectionRunResponse) HasArchivalRuns() bool`

HasArchivalRuns returns a boolean if a field has been set.

### SetArchivalRunsNil

`func (o *ProtectionRunResponse) SetArchivalRunsNil(b bool)`

 SetArchivalRunsNil sets the value for ArchivalRuns to be an explicit nil

### UnsetArchivalRuns
`func (o *ProtectionRunResponse) UnsetArchivalRuns()`

UnsetArchivalRuns ensures that no value is present for ArchivalRuns, not even an explicit nil
### GetBackupRuns

`func (o *ProtectionRunResponse) GetBackupRuns() []LatestProtectionJobRunInfo`

GetBackupRuns returns the BackupRuns field if non-nil, zero value otherwise.

### GetBackupRunsOk

`func (o *ProtectionRunResponse) GetBackupRunsOk() (*[]LatestProtectionJobRunInfo, bool)`

GetBackupRunsOk returns a tuple with the BackupRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRuns

`func (o *ProtectionRunResponse) SetBackupRuns(v []LatestProtectionJobRunInfo)`

SetBackupRuns sets BackupRuns field to given value.

### HasBackupRuns

`func (o *ProtectionRunResponse) HasBackupRuns() bool`

HasBackupRuns returns a boolean if a field has been set.

### SetBackupRunsNil

`func (o *ProtectionRunResponse) SetBackupRunsNil(b bool)`

 SetBackupRunsNil sets the value for BackupRuns to be an explicit nil

### UnsetBackupRuns
`func (o *ProtectionRunResponse) UnsetBackupRuns()`

UnsetBackupRuns ensures that no value is present for BackupRuns, not even an explicit nil
### GetReplicationRuns

`func (o *ProtectionRunResponse) GetReplicationRuns() []LatestProtectionJobRunInfo`

GetReplicationRuns returns the ReplicationRuns field if non-nil, zero value otherwise.

### GetReplicationRunsOk

`func (o *ProtectionRunResponse) GetReplicationRunsOk() (*[]LatestProtectionJobRunInfo, bool)`

GetReplicationRunsOk returns a tuple with the ReplicationRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationRuns

`func (o *ProtectionRunResponse) SetReplicationRuns(v []LatestProtectionJobRunInfo)`

SetReplicationRuns sets ReplicationRuns field to given value.

### HasReplicationRuns

`func (o *ProtectionRunResponse) HasReplicationRuns() bool`

HasReplicationRuns returns a boolean if a field has been set.

### SetReplicationRunsNil

`func (o *ProtectionRunResponse) SetReplicationRunsNil(b bool)`

 SetReplicationRunsNil sets the value for ReplicationRuns to be an explicit nil

### UnsetReplicationRuns
`func (o *ProtectionRunResponse) UnsetReplicationRuns()`

UnsetReplicationRuns ensures that no value is present for ReplicationRuns, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


