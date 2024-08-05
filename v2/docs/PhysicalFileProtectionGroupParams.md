# PhysicalFileProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowParallelRuns** | Pointer to **NullableBool** | Specifies whether or not this job can have parallel runs. | [optional] 
**CobmrBackup** | Pointer to **NullableBool** | Specifies whether to take CoBMR backup. | [optional] 
**ContinueOnQuiesceFailure** | Pointer to **NullableBool** | Specifies whether to continue backing up on quiesce failure. | [optional] 
**DedupExclusionSourceIds** | Pointer to **[]int64** | Specifies ids of sources for which deduplication has to be disabled. | [optional] 
**ExcludedVssWriters** | Pointer to **[]string** | Specifies writer names which should be excluded from physical file based backups. | [optional] 
**GlobalExcludeFS** | Pointer to **[]string** | Specifies global exclude filesystems which are applied to all sources in a job. | [optional] 
**GlobalExcludePaths** | Pointer to **[]string** | Specifies global exclude filters which are applied to all sources in a job. | [optional] 
**IgnorableErrors** | Pointer to **[]string** | Specifies the Errors to be ignored in error db. | [optional] 
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 
**Objects** | [**[]PhysicalFileProtectionGroupObjectParams**](PhysicalFileProtectionGroupObjectParams.md) | Specifies the list of objects protected by this Protection Group. | 
**PerformBrickBasedDeduplication** | Pointer to **NullableBool** | Specifies whether or not to perform brick based deduplication on this Protection Group. | [optional] 
**PerformSourceSideDeduplication** | Pointer to **NullableBool** | Specifies whether or not to perform source side deduplication on this Protection Group. | [optional] 
**PrePostScript** | Pointer to [**PrePostScriptParams**](PrePostScriptParams.md) |  | [optional] 
**Quiesce** | Pointer to **NullableBool** | Specifies Whether to take app-consistent snapshots by quiescing apps and the filesystem before taking a backup. | [optional] 
**TaskTimeouts** | Pointer to [**[]CancellationTimeoutParams**](CancellationTimeoutParams.md) | Specifies the timeouts for all the objects inside this Protection Group, for both full and incremental backups. | [optional] 

## Methods

### NewPhysicalFileProtectionGroupParams

`func NewPhysicalFileProtectionGroupParams(objects []PhysicalFileProtectionGroupObjectParams, ) *PhysicalFileProtectionGroupParams`

NewPhysicalFileProtectionGroupParams instantiates a new PhysicalFileProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalFileProtectionGroupParamsWithDefaults

`func NewPhysicalFileProtectionGroupParamsWithDefaults() *PhysicalFileProtectionGroupParams`

NewPhysicalFileProtectionGroupParamsWithDefaults instantiates a new PhysicalFileProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowParallelRuns

`func (o *PhysicalFileProtectionGroupParams) GetAllowParallelRuns() bool`

GetAllowParallelRuns returns the AllowParallelRuns field if non-nil, zero value otherwise.

### GetAllowParallelRunsOk

`func (o *PhysicalFileProtectionGroupParams) GetAllowParallelRunsOk() (*bool, bool)`

GetAllowParallelRunsOk returns a tuple with the AllowParallelRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowParallelRuns

`func (o *PhysicalFileProtectionGroupParams) SetAllowParallelRuns(v bool)`

SetAllowParallelRuns sets AllowParallelRuns field to given value.

### HasAllowParallelRuns

`func (o *PhysicalFileProtectionGroupParams) HasAllowParallelRuns() bool`

HasAllowParallelRuns returns a boolean if a field has been set.

### SetAllowParallelRunsNil

`func (o *PhysicalFileProtectionGroupParams) SetAllowParallelRunsNil(b bool)`

 SetAllowParallelRunsNil sets the value for AllowParallelRuns to be an explicit nil

### UnsetAllowParallelRuns
`func (o *PhysicalFileProtectionGroupParams) UnsetAllowParallelRuns()`

UnsetAllowParallelRuns ensures that no value is present for AllowParallelRuns, not even an explicit nil
### GetCobmrBackup

`func (o *PhysicalFileProtectionGroupParams) GetCobmrBackup() bool`

GetCobmrBackup returns the CobmrBackup field if non-nil, zero value otherwise.

### GetCobmrBackupOk

`func (o *PhysicalFileProtectionGroupParams) GetCobmrBackupOk() (*bool, bool)`

GetCobmrBackupOk returns a tuple with the CobmrBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCobmrBackup

`func (o *PhysicalFileProtectionGroupParams) SetCobmrBackup(v bool)`

SetCobmrBackup sets CobmrBackup field to given value.

### HasCobmrBackup

`func (o *PhysicalFileProtectionGroupParams) HasCobmrBackup() bool`

HasCobmrBackup returns a boolean if a field has been set.

### SetCobmrBackupNil

`func (o *PhysicalFileProtectionGroupParams) SetCobmrBackupNil(b bool)`

 SetCobmrBackupNil sets the value for CobmrBackup to be an explicit nil

### UnsetCobmrBackup
`func (o *PhysicalFileProtectionGroupParams) UnsetCobmrBackup()`

UnsetCobmrBackup ensures that no value is present for CobmrBackup, not even an explicit nil
### GetContinueOnQuiesceFailure

`func (o *PhysicalFileProtectionGroupParams) GetContinueOnQuiesceFailure() bool`

GetContinueOnQuiesceFailure returns the ContinueOnQuiesceFailure field if non-nil, zero value otherwise.

### GetContinueOnQuiesceFailureOk

`func (o *PhysicalFileProtectionGroupParams) GetContinueOnQuiesceFailureOk() (*bool, bool)`

GetContinueOnQuiesceFailureOk returns a tuple with the ContinueOnQuiesceFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnQuiesceFailure

`func (o *PhysicalFileProtectionGroupParams) SetContinueOnQuiesceFailure(v bool)`

SetContinueOnQuiesceFailure sets ContinueOnQuiesceFailure field to given value.

### HasContinueOnQuiesceFailure

`func (o *PhysicalFileProtectionGroupParams) HasContinueOnQuiesceFailure() bool`

HasContinueOnQuiesceFailure returns a boolean if a field has been set.

### SetContinueOnQuiesceFailureNil

`func (o *PhysicalFileProtectionGroupParams) SetContinueOnQuiesceFailureNil(b bool)`

 SetContinueOnQuiesceFailureNil sets the value for ContinueOnQuiesceFailure to be an explicit nil

### UnsetContinueOnQuiesceFailure
`func (o *PhysicalFileProtectionGroupParams) UnsetContinueOnQuiesceFailure()`

UnsetContinueOnQuiesceFailure ensures that no value is present for ContinueOnQuiesceFailure, not even an explicit nil
### GetDedupExclusionSourceIds

`func (o *PhysicalFileProtectionGroupParams) GetDedupExclusionSourceIds() []int64`

GetDedupExclusionSourceIds returns the DedupExclusionSourceIds field if non-nil, zero value otherwise.

### GetDedupExclusionSourceIdsOk

`func (o *PhysicalFileProtectionGroupParams) GetDedupExclusionSourceIdsOk() (*[]int64, bool)`

GetDedupExclusionSourceIdsOk returns a tuple with the DedupExclusionSourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedupExclusionSourceIds

`func (o *PhysicalFileProtectionGroupParams) SetDedupExclusionSourceIds(v []int64)`

SetDedupExclusionSourceIds sets DedupExclusionSourceIds field to given value.

### HasDedupExclusionSourceIds

`func (o *PhysicalFileProtectionGroupParams) HasDedupExclusionSourceIds() bool`

HasDedupExclusionSourceIds returns a boolean if a field has been set.

### GetExcludedVssWriters

`func (o *PhysicalFileProtectionGroupParams) GetExcludedVssWriters() []string`

GetExcludedVssWriters returns the ExcludedVssWriters field if non-nil, zero value otherwise.

### GetExcludedVssWritersOk

`func (o *PhysicalFileProtectionGroupParams) GetExcludedVssWritersOk() (*[]string, bool)`

GetExcludedVssWritersOk returns a tuple with the ExcludedVssWriters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedVssWriters

`func (o *PhysicalFileProtectionGroupParams) SetExcludedVssWriters(v []string)`

SetExcludedVssWriters sets ExcludedVssWriters field to given value.

### HasExcludedVssWriters

`func (o *PhysicalFileProtectionGroupParams) HasExcludedVssWriters() bool`

HasExcludedVssWriters returns a boolean if a field has been set.

### SetExcludedVssWritersNil

`func (o *PhysicalFileProtectionGroupParams) SetExcludedVssWritersNil(b bool)`

 SetExcludedVssWritersNil sets the value for ExcludedVssWriters to be an explicit nil

### UnsetExcludedVssWriters
`func (o *PhysicalFileProtectionGroupParams) UnsetExcludedVssWriters()`

UnsetExcludedVssWriters ensures that no value is present for ExcludedVssWriters, not even an explicit nil
### GetGlobalExcludeFS

`func (o *PhysicalFileProtectionGroupParams) GetGlobalExcludeFS() []string`

GetGlobalExcludeFS returns the GlobalExcludeFS field if non-nil, zero value otherwise.

### GetGlobalExcludeFSOk

`func (o *PhysicalFileProtectionGroupParams) GetGlobalExcludeFSOk() (*[]string, bool)`

GetGlobalExcludeFSOk returns a tuple with the GlobalExcludeFS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalExcludeFS

`func (o *PhysicalFileProtectionGroupParams) SetGlobalExcludeFS(v []string)`

SetGlobalExcludeFS sets GlobalExcludeFS field to given value.

### HasGlobalExcludeFS

`func (o *PhysicalFileProtectionGroupParams) HasGlobalExcludeFS() bool`

HasGlobalExcludeFS returns a boolean if a field has been set.

### GetGlobalExcludePaths

`func (o *PhysicalFileProtectionGroupParams) GetGlobalExcludePaths() []string`

GetGlobalExcludePaths returns the GlobalExcludePaths field if non-nil, zero value otherwise.

### GetGlobalExcludePathsOk

`func (o *PhysicalFileProtectionGroupParams) GetGlobalExcludePathsOk() (*[]string, bool)`

GetGlobalExcludePathsOk returns a tuple with the GlobalExcludePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalExcludePaths

`func (o *PhysicalFileProtectionGroupParams) SetGlobalExcludePaths(v []string)`

SetGlobalExcludePaths sets GlobalExcludePaths field to given value.

### HasGlobalExcludePaths

`func (o *PhysicalFileProtectionGroupParams) HasGlobalExcludePaths() bool`

HasGlobalExcludePaths returns a boolean if a field has been set.

### GetIgnorableErrors

`func (o *PhysicalFileProtectionGroupParams) GetIgnorableErrors() []string`

GetIgnorableErrors returns the IgnorableErrors field if non-nil, zero value otherwise.

### GetIgnorableErrorsOk

`func (o *PhysicalFileProtectionGroupParams) GetIgnorableErrorsOk() (*[]string, bool)`

GetIgnorableErrorsOk returns a tuple with the IgnorableErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnorableErrors

`func (o *PhysicalFileProtectionGroupParams) SetIgnorableErrors(v []string)`

SetIgnorableErrors sets IgnorableErrors field to given value.

### HasIgnorableErrors

`func (o *PhysicalFileProtectionGroupParams) HasIgnorableErrors() bool`

HasIgnorableErrors returns a boolean if a field has been set.

### SetIgnorableErrorsNil

`func (o *PhysicalFileProtectionGroupParams) SetIgnorableErrorsNil(b bool)`

 SetIgnorableErrorsNil sets the value for IgnorableErrors to be an explicit nil

### UnsetIgnorableErrors
`func (o *PhysicalFileProtectionGroupParams) UnsetIgnorableErrors()`

UnsetIgnorableErrors ensures that no value is present for IgnorableErrors, not even an explicit nil
### GetIndexingPolicy

`func (o *PhysicalFileProtectionGroupParams) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *PhysicalFileProtectionGroupParams) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *PhysicalFileProtectionGroupParams) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *PhysicalFileProtectionGroupParams) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetObjects

`func (o *PhysicalFileProtectionGroupParams) GetObjects() []PhysicalFileProtectionGroupObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *PhysicalFileProtectionGroupParams) GetObjectsOk() (*[]PhysicalFileProtectionGroupObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *PhysicalFileProtectionGroupParams) SetObjects(v []PhysicalFileProtectionGroupObjectParams)`

SetObjects sets Objects field to given value.


### GetPerformBrickBasedDeduplication

`func (o *PhysicalFileProtectionGroupParams) GetPerformBrickBasedDeduplication() bool`

GetPerformBrickBasedDeduplication returns the PerformBrickBasedDeduplication field if non-nil, zero value otherwise.

### GetPerformBrickBasedDeduplicationOk

`func (o *PhysicalFileProtectionGroupParams) GetPerformBrickBasedDeduplicationOk() (*bool, bool)`

GetPerformBrickBasedDeduplicationOk returns a tuple with the PerformBrickBasedDeduplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformBrickBasedDeduplication

`func (o *PhysicalFileProtectionGroupParams) SetPerformBrickBasedDeduplication(v bool)`

SetPerformBrickBasedDeduplication sets PerformBrickBasedDeduplication field to given value.

### HasPerformBrickBasedDeduplication

`func (o *PhysicalFileProtectionGroupParams) HasPerformBrickBasedDeduplication() bool`

HasPerformBrickBasedDeduplication returns a boolean if a field has been set.

### SetPerformBrickBasedDeduplicationNil

`func (o *PhysicalFileProtectionGroupParams) SetPerformBrickBasedDeduplicationNil(b bool)`

 SetPerformBrickBasedDeduplicationNil sets the value for PerformBrickBasedDeduplication to be an explicit nil

### UnsetPerformBrickBasedDeduplication
`func (o *PhysicalFileProtectionGroupParams) UnsetPerformBrickBasedDeduplication()`

UnsetPerformBrickBasedDeduplication ensures that no value is present for PerformBrickBasedDeduplication, not even an explicit nil
### GetPerformSourceSideDeduplication

`func (o *PhysicalFileProtectionGroupParams) GetPerformSourceSideDeduplication() bool`

GetPerformSourceSideDeduplication returns the PerformSourceSideDeduplication field if non-nil, zero value otherwise.

### GetPerformSourceSideDeduplicationOk

`func (o *PhysicalFileProtectionGroupParams) GetPerformSourceSideDeduplicationOk() (*bool, bool)`

GetPerformSourceSideDeduplicationOk returns a tuple with the PerformSourceSideDeduplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformSourceSideDeduplication

`func (o *PhysicalFileProtectionGroupParams) SetPerformSourceSideDeduplication(v bool)`

SetPerformSourceSideDeduplication sets PerformSourceSideDeduplication field to given value.

### HasPerformSourceSideDeduplication

`func (o *PhysicalFileProtectionGroupParams) HasPerformSourceSideDeduplication() bool`

HasPerformSourceSideDeduplication returns a boolean if a field has been set.

### SetPerformSourceSideDeduplicationNil

`func (o *PhysicalFileProtectionGroupParams) SetPerformSourceSideDeduplicationNil(b bool)`

 SetPerformSourceSideDeduplicationNil sets the value for PerformSourceSideDeduplication to be an explicit nil

### UnsetPerformSourceSideDeduplication
`func (o *PhysicalFileProtectionGroupParams) UnsetPerformSourceSideDeduplication()`

UnsetPerformSourceSideDeduplication ensures that no value is present for PerformSourceSideDeduplication, not even an explicit nil
### GetPrePostScript

`func (o *PhysicalFileProtectionGroupParams) GetPrePostScript() PrePostScriptParams`

GetPrePostScript returns the PrePostScript field if non-nil, zero value otherwise.

### GetPrePostScriptOk

`func (o *PhysicalFileProtectionGroupParams) GetPrePostScriptOk() (*PrePostScriptParams, bool)`

GetPrePostScriptOk returns a tuple with the PrePostScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePostScript

`func (o *PhysicalFileProtectionGroupParams) SetPrePostScript(v PrePostScriptParams)`

SetPrePostScript sets PrePostScript field to given value.

### HasPrePostScript

`func (o *PhysicalFileProtectionGroupParams) HasPrePostScript() bool`

HasPrePostScript returns a boolean if a field has been set.

### GetQuiesce

`func (o *PhysicalFileProtectionGroupParams) GetQuiesce() bool`

GetQuiesce returns the Quiesce field if non-nil, zero value otherwise.

### GetQuiesceOk

`func (o *PhysicalFileProtectionGroupParams) GetQuiesceOk() (*bool, bool)`

GetQuiesceOk returns a tuple with the Quiesce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuiesce

`func (o *PhysicalFileProtectionGroupParams) SetQuiesce(v bool)`

SetQuiesce sets Quiesce field to given value.

### HasQuiesce

`func (o *PhysicalFileProtectionGroupParams) HasQuiesce() bool`

HasQuiesce returns a boolean if a field has been set.

### SetQuiesceNil

`func (o *PhysicalFileProtectionGroupParams) SetQuiesceNil(b bool)`

 SetQuiesceNil sets the value for Quiesce to be an explicit nil

### UnsetQuiesce
`func (o *PhysicalFileProtectionGroupParams) UnsetQuiesce()`

UnsetQuiesce ensures that no value is present for Quiesce, not even an explicit nil
### GetTaskTimeouts

`func (o *PhysicalFileProtectionGroupParams) GetTaskTimeouts() []CancellationTimeoutParams`

GetTaskTimeouts returns the TaskTimeouts field if non-nil, zero value otherwise.

### GetTaskTimeoutsOk

`func (o *PhysicalFileProtectionGroupParams) GetTaskTimeoutsOk() (*[]CancellationTimeoutParams, bool)`

GetTaskTimeoutsOk returns a tuple with the TaskTimeouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskTimeouts

`func (o *PhysicalFileProtectionGroupParams) SetTaskTimeouts(v []CancellationTimeoutParams)`

SetTaskTimeouts sets TaskTimeouts field to given value.

### HasTaskTimeouts

`func (o *PhysicalFileProtectionGroupParams) HasTaskTimeouts() bool`

HasTaskTimeouts returns a boolean if a field has been set.

### SetTaskTimeoutsNil

`func (o *PhysicalFileProtectionGroupParams) SetTaskTimeoutsNil(b bool)`

 SetTaskTimeoutsNil sets the value for TaskTimeouts to be an explicit nil

### UnsetTaskTimeouts
`func (o *PhysicalFileProtectionGroupParams) UnsetTaskTimeouts()`

UnsetTaskTimeouts ensures that no value is present for TaskTimeouts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


