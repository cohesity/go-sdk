# RunProtectionJobParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CopyRunTargets** | Pointer to [**[]RunJobSnapshotTarget**](RunJobSnapshotTarget.md) | Optional parameter to be set if you want specific replication or archival associated with the policy to run. | [optional] 
**RunNowParameters** | Pointer to [**[]RunNowParameters**](RunNowParameters.md) | Optional parameters of a Run Now operation. | [optional] 
**RunType** | Pointer to **NullableString** | Specifies the type of backup. If not specified, &#39;kRegular&#39; is assumed. &#39;kRegular&#39; indicates a incremental (CBT) backup. Incremental backups utilizing CBT (if supported) are captured of the target protection objects. The first run of a kRegular schedule captures all the blocks. &#39;kFull&#39; indicates a full (no CBT) backup. A complete backup (all blocks) of the target protection objects are always captured and Change Block Tracking (CBT) is not utilized. &#39;kLog&#39; indicates a Database Log backup. Capture the database transaction logs to allow rolling back to a specific point in time. &#39;kSystem&#39; indicates a system backup. System backups are used to do bare metal recovery of the system to a specific point in time. | [optional] 
**SourceIds** | Pointer to **[]int64** | Optional parameter if you want to back up only a subset of sources that are protected by the job in this run. If a Run Now operation is to be performed then the source ids should only be provided in the runNowParameters along with the database Ids. | [optional] 
**UsePolicyDefaults** | Pointer to **NullableBool** | Specifies if default policy settings should be used interanally to copy snapshots to external targets already configured in policy. This field will only apply if \&quot;CopyRunTargets\&quot; is empty. | [optional] 

## Methods

### NewRunProtectionJobParam

`func NewRunProtectionJobParam() *RunProtectionJobParam`

NewRunProtectionJobParam instantiates a new RunProtectionJobParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunProtectionJobParamWithDefaults

`func NewRunProtectionJobParamWithDefaults() *RunProtectionJobParam`

NewRunProtectionJobParamWithDefaults instantiates a new RunProtectionJobParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCopyRunTargets

`func (o *RunProtectionJobParam) GetCopyRunTargets() []RunJobSnapshotTarget`

GetCopyRunTargets returns the CopyRunTargets field if non-nil, zero value otherwise.

### GetCopyRunTargetsOk

`func (o *RunProtectionJobParam) GetCopyRunTargetsOk() (*[]RunJobSnapshotTarget, bool)`

GetCopyRunTargetsOk returns a tuple with the CopyRunTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyRunTargets

`func (o *RunProtectionJobParam) SetCopyRunTargets(v []RunJobSnapshotTarget)`

SetCopyRunTargets sets CopyRunTargets field to given value.

### HasCopyRunTargets

`func (o *RunProtectionJobParam) HasCopyRunTargets() bool`

HasCopyRunTargets returns a boolean if a field has been set.

### SetCopyRunTargetsNil

`func (o *RunProtectionJobParam) SetCopyRunTargetsNil(b bool)`

 SetCopyRunTargetsNil sets the value for CopyRunTargets to be an explicit nil

### UnsetCopyRunTargets
`func (o *RunProtectionJobParam) UnsetCopyRunTargets()`

UnsetCopyRunTargets ensures that no value is present for CopyRunTargets, not even an explicit nil
### GetRunNowParameters

`func (o *RunProtectionJobParam) GetRunNowParameters() []RunNowParameters`

GetRunNowParameters returns the RunNowParameters field if non-nil, zero value otherwise.

### GetRunNowParametersOk

`func (o *RunProtectionJobParam) GetRunNowParametersOk() (*[]RunNowParameters, bool)`

GetRunNowParametersOk returns a tuple with the RunNowParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunNowParameters

`func (o *RunProtectionJobParam) SetRunNowParameters(v []RunNowParameters)`

SetRunNowParameters sets RunNowParameters field to given value.

### HasRunNowParameters

`func (o *RunProtectionJobParam) HasRunNowParameters() bool`

HasRunNowParameters returns a boolean if a field has been set.

### SetRunNowParametersNil

`func (o *RunProtectionJobParam) SetRunNowParametersNil(b bool)`

 SetRunNowParametersNil sets the value for RunNowParameters to be an explicit nil

### UnsetRunNowParameters
`func (o *RunProtectionJobParam) UnsetRunNowParameters()`

UnsetRunNowParameters ensures that no value is present for RunNowParameters, not even an explicit nil
### GetRunType

`func (o *RunProtectionJobParam) GetRunType() string`

GetRunType returns the RunType field if non-nil, zero value otherwise.

### GetRunTypeOk

`func (o *RunProtectionJobParam) GetRunTypeOk() (*string, bool)`

GetRunTypeOk returns a tuple with the RunType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunType

`func (o *RunProtectionJobParam) SetRunType(v string)`

SetRunType sets RunType field to given value.

### HasRunType

`func (o *RunProtectionJobParam) HasRunType() bool`

HasRunType returns a boolean if a field has been set.

### SetRunTypeNil

`func (o *RunProtectionJobParam) SetRunTypeNil(b bool)`

 SetRunTypeNil sets the value for RunType to be an explicit nil

### UnsetRunType
`func (o *RunProtectionJobParam) UnsetRunType()`

UnsetRunType ensures that no value is present for RunType, not even an explicit nil
### GetSourceIds

`func (o *RunProtectionJobParam) GetSourceIds() []int64`

GetSourceIds returns the SourceIds field if non-nil, zero value otherwise.

### GetSourceIdsOk

`func (o *RunProtectionJobParam) GetSourceIdsOk() (*[]int64, bool)`

GetSourceIdsOk returns a tuple with the SourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIds

`func (o *RunProtectionJobParam) SetSourceIds(v []int64)`

SetSourceIds sets SourceIds field to given value.

### HasSourceIds

`func (o *RunProtectionJobParam) HasSourceIds() bool`

HasSourceIds returns a boolean if a field has been set.

### SetSourceIdsNil

`func (o *RunProtectionJobParam) SetSourceIdsNil(b bool)`

 SetSourceIdsNil sets the value for SourceIds to be an explicit nil

### UnsetSourceIds
`func (o *RunProtectionJobParam) UnsetSourceIds()`

UnsetSourceIds ensures that no value is present for SourceIds, not even an explicit nil
### GetUsePolicyDefaults

`func (o *RunProtectionJobParam) GetUsePolicyDefaults() bool`

GetUsePolicyDefaults returns the UsePolicyDefaults field if non-nil, zero value otherwise.

### GetUsePolicyDefaultsOk

`func (o *RunProtectionJobParam) GetUsePolicyDefaultsOk() (*bool, bool)`

GetUsePolicyDefaultsOk returns a tuple with the UsePolicyDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePolicyDefaults

`func (o *RunProtectionJobParam) SetUsePolicyDefaults(v bool)`

SetUsePolicyDefaults sets UsePolicyDefaults field to given value.

### HasUsePolicyDefaults

`func (o *RunProtectionJobParam) HasUsePolicyDefaults() bool`

HasUsePolicyDefaults returns a boolean if a field has been set.

### SetUsePolicyDefaultsNil

`func (o *RunProtectionJobParam) SetUsePolicyDefaultsNil(b bool)`

 SetUsePolicyDefaultsNil sets the value for UsePolicyDefaults to be an explicit nil

### UnsetUsePolicyDefaults
`func (o *RunProtectionJobParam) UnsetUsePolicyDefaults()`

UnsetUsePolicyDefaults ensures that no value is present for UsePolicyDefaults, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


