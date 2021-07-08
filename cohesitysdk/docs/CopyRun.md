# CopyRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CopySnapshotTasks** | Pointer to [**[]CopySnapshotTaskStatus**](CopySnapshotTaskStatus.md) | Specifies the status information of each task that copies the snapshot taken for a Protection Source. | [optional] 
**DataLockConstraints** | Pointer to [**DataLockConstraints**](DataLockConstraints.md) |  | [optional] 
**Error** | Pointer to **NullableString** | Specifies if an error occurred (if any) while running this task. This field is populated when the status is equal to &#39;kFailure&#39;. | [optional] 
**ExpiryTimeUsecs** | Pointer to **NullableInt64** | Specifies expiry time of the copies of the snapshots in this Protection Run. | [optional] 
**HoldForLegalPurpose** | Pointer to **NullableBool** | Specifies whether legal hold is enabled on this run. It is true if the run is put on legal hold. Independent of this flag, some of the entities may be on legal hold. | [optional] 
**LegalHoldings** | Pointer to [**[]LegalHoldings**](LegalHoldings.md) | Specifies the list of Protection Source Ids and the legal hold status. | [optional] 
**RunStartTimeUsecs** | Pointer to **NullableInt64** | Specifies start time of the copy run. | [optional] 
**Stats** | Pointer to [**CopyRunStats**](CopyRunStats.md) |  | [optional] 
**Status** | Pointer to **NullableString** | Specifies the aggregated status of copy tasks such as &#39;kRunning&#39;, &#39;kSuccess&#39;, &#39;kFailure&#39; etc. &#39;kAccepted&#39; indicates the task is queued to run but not yet running. &#39;kRunning&#39; indicates the task is running. &#39;kCanceling&#39; indicates a request to cancel the task has occurred but the task is not yet canceled. &#39;kCanceled&#39; indicates the task has been canceled. &#39;kSuccess&#39; indicates the task was successful. &#39;kFailure&#39; indicates the task failed. &#39;kWarning&#39; indicates the task has finished with warning. &#39;kOnHold&#39; indicates the task is kept onHold. &#39;kMissed&#39; indicates the task is missed. | [optional] 
**Target** | Pointer to [**SnapshotTargetSettings**](SnapshotTargetSettings.md) |  | [optional] 
**TaskUid** | Pointer to [**NullableUniversalId**](UniversalId.md) | Specifies a globally unique id of the copy task. | [optional] 
**UserActionMessage** | Pointer to **NullableString** | Specifies a message to the user if any manual intervention is needed to make forward progress for the archival task. This message is mainly relevant for tape based archival tasks where a backup admin might be asked to load a new media when the tape library does not have any more media to use. | [optional] 

## Methods

### NewCopyRun

`func NewCopyRun() *CopyRun`

NewCopyRun instantiates a new CopyRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCopyRunWithDefaults

`func NewCopyRunWithDefaults() *CopyRun`

NewCopyRunWithDefaults instantiates a new CopyRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCopySnapshotTasks

`func (o *CopyRun) GetCopySnapshotTasks() []CopySnapshotTaskStatus`

GetCopySnapshotTasks returns the CopySnapshotTasks field if non-nil, zero value otherwise.

### GetCopySnapshotTasksOk

`func (o *CopyRun) GetCopySnapshotTasksOk() (*[]CopySnapshotTaskStatus, bool)`

GetCopySnapshotTasksOk returns a tuple with the CopySnapshotTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopySnapshotTasks

`func (o *CopyRun) SetCopySnapshotTasks(v []CopySnapshotTaskStatus)`

SetCopySnapshotTasks sets CopySnapshotTasks field to given value.

### HasCopySnapshotTasks

`func (o *CopyRun) HasCopySnapshotTasks() bool`

HasCopySnapshotTasks returns a boolean if a field has been set.

### SetCopySnapshotTasksNil

`func (o *CopyRun) SetCopySnapshotTasksNil(b bool)`

 SetCopySnapshotTasksNil sets the value for CopySnapshotTasks to be an explicit nil

### UnsetCopySnapshotTasks
`func (o *CopyRun) UnsetCopySnapshotTasks()`

UnsetCopySnapshotTasks ensures that no value is present for CopySnapshotTasks, not even an explicit nil
### GetDataLockConstraints

`func (o *CopyRun) GetDataLockConstraints() DataLockConstraints`

GetDataLockConstraints returns the DataLockConstraints field if non-nil, zero value otherwise.

### GetDataLockConstraintsOk

`func (o *CopyRun) GetDataLockConstraintsOk() (*DataLockConstraints, bool)`

GetDataLockConstraintsOk returns a tuple with the DataLockConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLockConstraints

`func (o *CopyRun) SetDataLockConstraints(v DataLockConstraints)`

SetDataLockConstraints sets DataLockConstraints field to given value.

### HasDataLockConstraints

`func (o *CopyRun) HasDataLockConstraints() bool`

HasDataLockConstraints returns a boolean if a field has been set.

### GetError

`func (o *CopyRun) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CopyRun) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CopyRun) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CopyRun) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *CopyRun) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *CopyRun) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetExpiryTimeUsecs

`func (o *CopyRun) GetExpiryTimeUsecs() int64`

GetExpiryTimeUsecs returns the ExpiryTimeUsecs field if non-nil, zero value otherwise.

### GetExpiryTimeUsecsOk

`func (o *CopyRun) GetExpiryTimeUsecsOk() (*int64, bool)`

GetExpiryTimeUsecsOk returns a tuple with the ExpiryTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimeUsecs

`func (o *CopyRun) SetExpiryTimeUsecs(v int64)`

SetExpiryTimeUsecs sets ExpiryTimeUsecs field to given value.

### HasExpiryTimeUsecs

`func (o *CopyRun) HasExpiryTimeUsecs() bool`

HasExpiryTimeUsecs returns a boolean if a field has been set.

### SetExpiryTimeUsecsNil

`func (o *CopyRun) SetExpiryTimeUsecsNil(b bool)`

 SetExpiryTimeUsecsNil sets the value for ExpiryTimeUsecs to be an explicit nil

### UnsetExpiryTimeUsecs
`func (o *CopyRun) UnsetExpiryTimeUsecs()`

UnsetExpiryTimeUsecs ensures that no value is present for ExpiryTimeUsecs, not even an explicit nil
### GetHoldForLegalPurpose

`func (o *CopyRun) GetHoldForLegalPurpose() bool`

GetHoldForLegalPurpose returns the HoldForLegalPurpose field if non-nil, zero value otherwise.

### GetHoldForLegalPurposeOk

`func (o *CopyRun) GetHoldForLegalPurposeOk() (*bool, bool)`

GetHoldForLegalPurposeOk returns a tuple with the HoldForLegalPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldForLegalPurpose

`func (o *CopyRun) SetHoldForLegalPurpose(v bool)`

SetHoldForLegalPurpose sets HoldForLegalPurpose field to given value.

### HasHoldForLegalPurpose

`func (o *CopyRun) HasHoldForLegalPurpose() bool`

HasHoldForLegalPurpose returns a boolean if a field has been set.

### SetHoldForLegalPurposeNil

`func (o *CopyRun) SetHoldForLegalPurposeNil(b bool)`

 SetHoldForLegalPurposeNil sets the value for HoldForLegalPurpose to be an explicit nil

### UnsetHoldForLegalPurpose
`func (o *CopyRun) UnsetHoldForLegalPurpose()`

UnsetHoldForLegalPurpose ensures that no value is present for HoldForLegalPurpose, not even an explicit nil
### GetLegalHoldings

`func (o *CopyRun) GetLegalHoldings() []LegalHoldings`

GetLegalHoldings returns the LegalHoldings field if non-nil, zero value otherwise.

### GetLegalHoldingsOk

`func (o *CopyRun) GetLegalHoldingsOk() (*[]LegalHoldings, bool)`

GetLegalHoldingsOk returns a tuple with the LegalHoldings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalHoldings

`func (o *CopyRun) SetLegalHoldings(v []LegalHoldings)`

SetLegalHoldings sets LegalHoldings field to given value.

### HasLegalHoldings

`func (o *CopyRun) HasLegalHoldings() bool`

HasLegalHoldings returns a boolean if a field has been set.

### SetLegalHoldingsNil

`func (o *CopyRun) SetLegalHoldingsNil(b bool)`

 SetLegalHoldingsNil sets the value for LegalHoldings to be an explicit nil

### UnsetLegalHoldings
`func (o *CopyRun) UnsetLegalHoldings()`

UnsetLegalHoldings ensures that no value is present for LegalHoldings, not even an explicit nil
### GetRunStartTimeUsecs

`func (o *CopyRun) GetRunStartTimeUsecs() int64`

GetRunStartTimeUsecs returns the RunStartTimeUsecs field if non-nil, zero value otherwise.

### GetRunStartTimeUsecsOk

`func (o *CopyRun) GetRunStartTimeUsecsOk() (*int64, bool)`

GetRunStartTimeUsecsOk returns a tuple with the RunStartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunStartTimeUsecs

`func (o *CopyRun) SetRunStartTimeUsecs(v int64)`

SetRunStartTimeUsecs sets RunStartTimeUsecs field to given value.

### HasRunStartTimeUsecs

`func (o *CopyRun) HasRunStartTimeUsecs() bool`

HasRunStartTimeUsecs returns a boolean if a field has been set.

### SetRunStartTimeUsecsNil

`func (o *CopyRun) SetRunStartTimeUsecsNil(b bool)`

 SetRunStartTimeUsecsNil sets the value for RunStartTimeUsecs to be an explicit nil

### UnsetRunStartTimeUsecs
`func (o *CopyRun) UnsetRunStartTimeUsecs()`

UnsetRunStartTimeUsecs ensures that no value is present for RunStartTimeUsecs, not even an explicit nil
### GetStats

`func (o *CopyRun) GetStats() CopyRunStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *CopyRun) GetStatsOk() (*CopyRunStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *CopyRun) SetStats(v CopyRunStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *CopyRun) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetStatus

`func (o *CopyRun) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CopyRun) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CopyRun) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CopyRun) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *CopyRun) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *CopyRun) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTarget

`func (o *CopyRun) GetTarget() SnapshotTargetSettings`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CopyRun) GetTargetOk() (*SnapshotTargetSettings, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CopyRun) SetTarget(v SnapshotTargetSettings)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *CopyRun) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTaskUid

`func (o *CopyRun) GetTaskUid() UniversalId`

GetTaskUid returns the TaskUid field if non-nil, zero value otherwise.

### GetTaskUidOk

`func (o *CopyRun) GetTaskUidOk() (*UniversalId, bool)`

GetTaskUidOk returns a tuple with the TaskUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskUid

`func (o *CopyRun) SetTaskUid(v UniversalId)`

SetTaskUid sets TaskUid field to given value.

### HasTaskUid

`func (o *CopyRun) HasTaskUid() bool`

HasTaskUid returns a boolean if a field has been set.

### SetTaskUidNil

`func (o *CopyRun) SetTaskUidNil(b bool)`

 SetTaskUidNil sets the value for TaskUid to be an explicit nil

### UnsetTaskUid
`func (o *CopyRun) UnsetTaskUid()`

UnsetTaskUid ensures that no value is present for TaskUid, not even an explicit nil
### GetUserActionMessage

`func (o *CopyRun) GetUserActionMessage() string`

GetUserActionMessage returns the UserActionMessage field if non-nil, zero value otherwise.

### GetUserActionMessageOk

`func (o *CopyRun) GetUserActionMessageOk() (*string, bool)`

GetUserActionMessageOk returns a tuple with the UserActionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserActionMessage

`func (o *CopyRun) SetUserActionMessage(v string)`

SetUserActionMessage sets UserActionMessage field to given value.

### HasUserActionMessage

`func (o *CopyRun) HasUserActionMessage() bool`

HasUserActionMessage returns a boolean if a field has been set.

### SetUserActionMessageNil

`func (o *CopyRun) SetUserActionMessageNil(b bool)`

 SetUserActionMessageNil sets the value for UserActionMessage to be an explicit nil

### UnsetUserActionMessage
`func (o *CopyRun) UnsetUserActionMessage()`

UnsetUserActionMessage ensures that no value is present for UserActionMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


