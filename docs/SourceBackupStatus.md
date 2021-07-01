# SourceBackupStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppsBackupStatus** | Pointer to [**[]AppEntityBackupStatusInfo**](AppEntityBackupStatusInfo.md) | Specifies the backup status at app/DB level. | [optional] 
**CurrentSnapshotInfo** | Pointer to [**SnapshotInfo**](SnapshotInfo.md) |  | [optional] 
**Error** | Pointer to **NullableString** | Specifies if an error occurred (if any) while running this task. This field is populated when the status is equal to &#39;kFailure&#39;. | [optional] 
**IsFullBackup** | Pointer to **NullableBool** | Specifies whether this is a &#39;kFull&#39; or &#39;kRegular&#39; backup of the Run. This may be true even if the scheduled backup type is &#39;kRegular&#39;. This will happen when this run corresponds to the first backup run of the Job or if no previous snapshot information is found. | [optional] 
**NumRestarts** | Pointer to **NullableInt32** | Specifies the number of times the task was restarted because of the changes on the backup source host. | [optional] 
**ParentSourceId** | Pointer to **NullableInt64** | Specifies the id of the registered Protection Source that is the parent of the Objects that are protected by this Job Run. | [optional] 
**ProgressMonitorTaskPath** | Pointer to **NullableString** | Specifies the yoda progress monitor task path which is used to get pulse information about the source that is being backed up. | [optional] 
**Quiesced** | Pointer to **NullableBool** | Specifies if app-consistent snapshot was captured. This field is set to true, if an app-consistent snapshot was taken by quiescing applications and the file system before taking a backup. | [optional] 
**SlaViolated** | Pointer to **NullableBool** | Specifies if the SLA was violated for the Job Run. This field is set to true, if time to complete the Job Run is longer than the SLA specified. This field is populated when the status is set to &#39;kSuccess&#39; or &#39;kFailure&#39;. | [optional] 
**Source** | Pointer to [**ProtectionSource**](ProtectionSource.md) |  | [optional] 
**Stats** | Pointer to [**BackupSourceStats**](BackupSourceStats.md) |  | [optional] 
**Status** | Pointer to **NullableString** | Specifies the status of the source object being protected. &#39;kAccepted&#39; indicates the task is queued to run but not yet running. &#39;kRunning&#39; indicates the task is running. &#39;kCanceling&#39; indicates a request to cancel the task has occurred but the task is not yet canceled. &#39;kCanceled&#39; indicates the task has been canceled. &#39;kSuccess&#39; indicates the task was successful. &#39;kFailure&#39; indicates the task failed. &#39;kWarning&#39; indicates the task has finished with warning. &#39;kOnHold&#39; indicates the task is kept onHold. &#39;kMissed&#39; indicates the task is missed. | [optional] 
**Warnings** | Pointer to **[]string** | Array of Warnings.  Specifies the warnings that occurred (if any) while running this task. | [optional] 

## Methods

### NewSourceBackupStatus

`func NewSourceBackupStatus() *SourceBackupStatus`

NewSourceBackupStatus instantiates a new SourceBackupStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceBackupStatusWithDefaults

`func NewSourceBackupStatusWithDefaults() *SourceBackupStatus`

NewSourceBackupStatusWithDefaults instantiates a new SourceBackupStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppsBackupStatus

`func (o *SourceBackupStatus) GetAppsBackupStatus() []AppEntityBackupStatusInfo`

GetAppsBackupStatus returns the AppsBackupStatus field if non-nil, zero value otherwise.

### GetAppsBackupStatusOk

`func (o *SourceBackupStatus) GetAppsBackupStatusOk() (*[]AppEntityBackupStatusInfo, bool)`

GetAppsBackupStatusOk returns a tuple with the AppsBackupStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppsBackupStatus

`func (o *SourceBackupStatus) SetAppsBackupStatus(v []AppEntityBackupStatusInfo)`

SetAppsBackupStatus sets AppsBackupStatus field to given value.

### HasAppsBackupStatus

`func (o *SourceBackupStatus) HasAppsBackupStatus() bool`

HasAppsBackupStatus returns a boolean if a field has been set.

### SetAppsBackupStatusNil

`func (o *SourceBackupStatus) SetAppsBackupStatusNil(b bool)`

 SetAppsBackupStatusNil sets the value for AppsBackupStatus to be an explicit nil

### UnsetAppsBackupStatus
`func (o *SourceBackupStatus) UnsetAppsBackupStatus()`

UnsetAppsBackupStatus ensures that no value is present for AppsBackupStatus, not even an explicit nil
### GetCurrentSnapshotInfo

`func (o *SourceBackupStatus) GetCurrentSnapshotInfo() SnapshotInfo`

GetCurrentSnapshotInfo returns the CurrentSnapshotInfo field if non-nil, zero value otherwise.

### GetCurrentSnapshotInfoOk

`func (o *SourceBackupStatus) GetCurrentSnapshotInfoOk() (*SnapshotInfo, bool)`

GetCurrentSnapshotInfoOk returns a tuple with the CurrentSnapshotInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSnapshotInfo

`func (o *SourceBackupStatus) SetCurrentSnapshotInfo(v SnapshotInfo)`

SetCurrentSnapshotInfo sets CurrentSnapshotInfo field to given value.

### HasCurrentSnapshotInfo

`func (o *SourceBackupStatus) HasCurrentSnapshotInfo() bool`

HasCurrentSnapshotInfo returns a boolean if a field has been set.

### GetError

`func (o *SourceBackupStatus) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SourceBackupStatus) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SourceBackupStatus) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *SourceBackupStatus) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *SourceBackupStatus) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *SourceBackupStatus) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetIsFullBackup

`func (o *SourceBackupStatus) GetIsFullBackup() bool`

GetIsFullBackup returns the IsFullBackup field if non-nil, zero value otherwise.

### GetIsFullBackupOk

`func (o *SourceBackupStatus) GetIsFullBackupOk() (*bool, bool)`

GetIsFullBackupOk returns a tuple with the IsFullBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFullBackup

`func (o *SourceBackupStatus) SetIsFullBackup(v bool)`

SetIsFullBackup sets IsFullBackup field to given value.

### HasIsFullBackup

`func (o *SourceBackupStatus) HasIsFullBackup() bool`

HasIsFullBackup returns a boolean if a field has been set.

### SetIsFullBackupNil

`func (o *SourceBackupStatus) SetIsFullBackupNil(b bool)`

 SetIsFullBackupNil sets the value for IsFullBackup to be an explicit nil

### UnsetIsFullBackup
`func (o *SourceBackupStatus) UnsetIsFullBackup()`

UnsetIsFullBackup ensures that no value is present for IsFullBackup, not even an explicit nil
### GetNumRestarts

`func (o *SourceBackupStatus) GetNumRestarts() int32`

GetNumRestarts returns the NumRestarts field if non-nil, zero value otherwise.

### GetNumRestartsOk

`func (o *SourceBackupStatus) GetNumRestartsOk() (*int32, bool)`

GetNumRestartsOk returns a tuple with the NumRestarts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRestarts

`func (o *SourceBackupStatus) SetNumRestarts(v int32)`

SetNumRestarts sets NumRestarts field to given value.

### HasNumRestarts

`func (o *SourceBackupStatus) HasNumRestarts() bool`

HasNumRestarts returns a boolean if a field has been set.

### SetNumRestartsNil

`func (o *SourceBackupStatus) SetNumRestartsNil(b bool)`

 SetNumRestartsNil sets the value for NumRestarts to be an explicit nil

### UnsetNumRestarts
`func (o *SourceBackupStatus) UnsetNumRestarts()`

UnsetNumRestarts ensures that no value is present for NumRestarts, not even an explicit nil
### GetParentSourceId

`func (o *SourceBackupStatus) GetParentSourceId() int64`

GetParentSourceId returns the ParentSourceId field if non-nil, zero value otherwise.

### GetParentSourceIdOk

`func (o *SourceBackupStatus) GetParentSourceIdOk() (*int64, bool)`

GetParentSourceIdOk returns a tuple with the ParentSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSourceId

`func (o *SourceBackupStatus) SetParentSourceId(v int64)`

SetParentSourceId sets ParentSourceId field to given value.

### HasParentSourceId

`func (o *SourceBackupStatus) HasParentSourceId() bool`

HasParentSourceId returns a boolean if a field has been set.

### SetParentSourceIdNil

`func (o *SourceBackupStatus) SetParentSourceIdNil(b bool)`

 SetParentSourceIdNil sets the value for ParentSourceId to be an explicit nil

### UnsetParentSourceId
`func (o *SourceBackupStatus) UnsetParentSourceId()`

UnsetParentSourceId ensures that no value is present for ParentSourceId, not even an explicit nil
### GetProgressMonitorTaskPath

`func (o *SourceBackupStatus) GetProgressMonitorTaskPath() string`

GetProgressMonitorTaskPath returns the ProgressMonitorTaskPath field if non-nil, zero value otherwise.

### GetProgressMonitorTaskPathOk

`func (o *SourceBackupStatus) GetProgressMonitorTaskPathOk() (*string, bool)`

GetProgressMonitorTaskPathOk returns a tuple with the ProgressMonitorTaskPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressMonitorTaskPath

`func (o *SourceBackupStatus) SetProgressMonitorTaskPath(v string)`

SetProgressMonitorTaskPath sets ProgressMonitorTaskPath field to given value.

### HasProgressMonitorTaskPath

`func (o *SourceBackupStatus) HasProgressMonitorTaskPath() bool`

HasProgressMonitorTaskPath returns a boolean if a field has been set.

### SetProgressMonitorTaskPathNil

`func (o *SourceBackupStatus) SetProgressMonitorTaskPathNil(b bool)`

 SetProgressMonitorTaskPathNil sets the value for ProgressMonitorTaskPath to be an explicit nil

### UnsetProgressMonitorTaskPath
`func (o *SourceBackupStatus) UnsetProgressMonitorTaskPath()`

UnsetProgressMonitorTaskPath ensures that no value is present for ProgressMonitorTaskPath, not even an explicit nil
### GetQuiesced

`func (o *SourceBackupStatus) GetQuiesced() bool`

GetQuiesced returns the Quiesced field if non-nil, zero value otherwise.

### GetQuiescedOk

`func (o *SourceBackupStatus) GetQuiescedOk() (*bool, bool)`

GetQuiescedOk returns a tuple with the Quiesced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuiesced

`func (o *SourceBackupStatus) SetQuiesced(v bool)`

SetQuiesced sets Quiesced field to given value.

### HasQuiesced

`func (o *SourceBackupStatus) HasQuiesced() bool`

HasQuiesced returns a boolean if a field has been set.

### SetQuiescedNil

`func (o *SourceBackupStatus) SetQuiescedNil(b bool)`

 SetQuiescedNil sets the value for Quiesced to be an explicit nil

### UnsetQuiesced
`func (o *SourceBackupStatus) UnsetQuiesced()`

UnsetQuiesced ensures that no value is present for Quiesced, not even an explicit nil
### GetSlaViolated

`func (o *SourceBackupStatus) GetSlaViolated() bool`

GetSlaViolated returns the SlaViolated field if non-nil, zero value otherwise.

### GetSlaViolatedOk

`func (o *SourceBackupStatus) GetSlaViolatedOk() (*bool, bool)`

GetSlaViolatedOk returns a tuple with the SlaViolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaViolated

`func (o *SourceBackupStatus) SetSlaViolated(v bool)`

SetSlaViolated sets SlaViolated field to given value.

### HasSlaViolated

`func (o *SourceBackupStatus) HasSlaViolated() bool`

HasSlaViolated returns a boolean if a field has been set.

### SetSlaViolatedNil

`func (o *SourceBackupStatus) SetSlaViolatedNil(b bool)`

 SetSlaViolatedNil sets the value for SlaViolated to be an explicit nil

### UnsetSlaViolated
`func (o *SourceBackupStatus) UnsetSlaViolated()`

UnsetSlaViolated ensures that no value is present for SlaViolated, not even an explicit nil
### GetSource

`func (o *SourceBackupStatus) GetSource() ProtectionSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SourceBackupStatus) GetSourceOk() (*ProtectionSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SourceBackupStatus) SetSource(v ProtectionSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *SourceBackupStatus) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStats

`func (o *SourceBackupStatus) GetStats() BackupSourceStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *SourceBackupStatus) GetStatsOk() (*BackupSourceStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *SourceBackupStatus) SetStats(v BackupSourceStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *SourceBackupStatus) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetStatus

`func (o *SourceBackupStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SourceBackupStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SourceBackupStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SourceBackupStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *SourceBackupStatus) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *SourceBackupStatus) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetWarnings

`func (o *SourceBackupStatus) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *SourceBackupStatus) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *SourceBackupStatus) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *SourceBackupStatus) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarningsNil

`func (o *SourceBackupStatus) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *SourceBackupStatus) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


