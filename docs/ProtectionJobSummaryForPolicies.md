# ProtectionJobSummaryForPolicies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupRun** | Pointer to [**BackupRun**](BackupRun.md) |  | [optional] 
**CopyRuns** | Pointer to [**[]CopyRun**](CopyRun.md) | Specifies details about the Copy tasks of the Job Run. A Copy task copies the captured snapshots to an external target or a Remote Cohesity Cluster. | [optional] 
**ProtectionJob** | Pointer to [**ProtectionJob**](ProtectionJob.md) |  | [optional] 

## Methods

### NewProtectionJobSummaryForPolicies

`func NewProtectionJobSummaryForPolicies() *ProtectionJobSummaryForPolicies`

NewProtectionJobSummaryForPolicies instantiates a new ProtectionJobSummaryForPolicies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionJobSummaryForPoliciesWithDefaults

`func NewProtectionJobSummaryForPoliciesWithDefaults() *ProtectionJobSummaryForPolicies`

NewProtectionJobSummaryForPoliciesWithDefaults instantiates a new ProtectionJobSummaryForPolicies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupRun

`func (o *ProtectionJobSummaryForPolicies) GetBackupRun() BackupRun`

GetBackupRun returns the BackupRun field if non-nil, zero value otherwise.

### GetBackupRunOk

`func (o *ProtectionJobSummaryForPolicies) GetBackupRunOk() (*BackupRun, bool)`

GetBackupRunOk returns a tuple with the BackupRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRun

`func (o *ProtectionJobSummaryForPolicies) SetBackupRun(v BackupRun)`

SetBackupRun sets BackupRun field to given value.

### HasBackupRun

`func (o *ProtectionJobSummaryForPolicies) HasBackupRun() bool`

HasBackupRun returns a boolean if a field has been set.

### GetCopyRuns

`func (o *ProtectionJobSummaryForPolicies) GetCopyRuns() []CopyRun`

GetCopyRuns returns the CopyRuns field if non-nil, zero value otherwise.

### GetCopyRunsOk

`func (o *ProtectionJobSummaryForPolicies) GetCopyRunsOk() (*[]CopyRun, bool)`

GetCopyRunsOk returns a tuple with the CopyRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyRuns

`func (o *ProtectionJobSummaryForPolicies) SetCopyRuns(v []CopyRun)`

SetCopyRuns sets CopyRuns field to given value.

### HasCopyRuns

`func (o *ProtectionJobSummaryForPolicies) HasCopyRuns() bool`

HasCopyRuns returns a boolean if a field has been set.

### SetCopyRunsNil

`func (o *ProtectionJobSummaryForPolicies) SetCopyRunsNil(b bool)`

 SetCopyRunsNil sets the value for CopyRuns to be an explicit nil

### UnsetCopyRuns
`func (o *ProtectionJobSummaryForPolicies) UnsetCopyRuns()`

UnsetCopyRuns ensures that no value is present for CopyRuns, not even an explicit nil
### GetProtectionJob

`func (o *ProtectionJobSummaryForPolicies) GetProtectionJob() ProtectionJob`

GetProtectionJob returns the ProtectionJob field if non-nil, zero value otherwise.

### GetProtectionJobOk

`func (o *ProtectionJobSummaryForPolicies) GetProtectionJobOk() (*ProtectionJob, bool)`

GetProtectionJobOk returns a tuple with the ProtectionJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionJob

`func (o *ProtectionJobSummaryForPolicies) SetProtectionJob(v ProtectionJob)`

SetProtectionJob sets ProtectionJob field to given value.

### HasProtectionJob

`func (o *ProtectionJobSummaryForPolicies) HasProtectionJob() bool`

HasProtectionJob returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


