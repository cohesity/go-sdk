# PhysicalVolumeProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | [**[]PhysicalVolumeProtectionGroupObjectParams**](PhysicalVolumeProtectionGroupObjectParams.md) |  | 
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 
**PerformSourceSideDeduplication** | Pointer to **NullableBool** | Specifies whether or not to perform source side deduplication on this Protection Group. | [optional] 
**Quiesce** | Pointer to **NullableBool** | Specifies Whether to take app-consistent snapshots by quiescing apps and the filesystem before taking a backup | [optional] 
**ContinueOnQuiesceFailure** | Pointer to **NullableBool** | Specifies whether to continue backing up on quiesce failure | [optional] 
**IncrementalBackupAfterRestart** | Pointer to **NullableBool** | Specifies whether or not to perform an incremental backup after the server restarts. This is applicable to windows environments. | [optional] 
**PrePostScript** | Pointer to [**PrePostScriptParams**](PrePostScriptParams.md) |  | [optional] 
**DedupExclusionSourceIds** | Pointer to **[]int64** | Specifies ids of sources for which deduplication has to be disabled. | [optional] 
**ExcludedVssWriters** | Pointer to **[]string** | Specifies writer names which should be excluded from physical volume based backups. | [optional] 

## Methods

### NewPhysicalVolumeProtectionGroupParams

`func NewPhysicalVolumeProtectionGroupParams(objects []PhysicalVolumeProtectionGroupObjectParams, ) *PhysicalVolumeProtectionGroupParams`

NewPhysicalVolumeProtectionGroupParams instantiates a new PhysicalVolumeProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalVolumeProtectionGroupParamsWithDefaults

`func NewPhysicalVolumeProtectionGroupParamsWithDefaults() *PhysicalVolumeProtectionGroupParams`

NewPhysicalVolumeProtectionGroupParamsWithDefaults instantiates a new PhysicalVolumeProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *PhysicalVolumeProtectionGroupParams) GetObjects() []PhysicalVolumeProtectionGroupObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *PhysicalVolumeProtectionGroupParams) GetObjectsOk() (*[]PhysicalVolumeProtectionGroupObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *PhysicalVolumeProtectionGroupParams) SetObjects(v []PhysicalVolumeProtectionGroupObjectParams)`

SetObjects sets Objects field to given value.


### GetIndexingPolicy

`func (o *PhysicalVolumeProtectionGroupParams) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *PhysicalVolumeProtectionGroupParams) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *PhysicalVolumeProtectionGroupParams) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *PhysicalVolumeProtectionGroupParams) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetPerformSourceSideDeduplication

`func (o *PhysicalVolumeProtectionGroupParams) GetPerformSourceSideDeduplication() bool`

GetPerformSourceSideDeduplication returns the PerformSourceSideDeduplication field if non-nil, zero value otherwise.

### GetPerformSourceSideDeduplicationOk

`func (o *PhysicalVolumeProtectionGroupParams) GetPerformSourceSideDeduplicationOk() (*bool, bool)`

GetPerformSourceSideDeduplicationOk returns a tuple with the PerformSourceSideDeduplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformSourceSideDeduplication

`func (o *PhysicalVolumeProtectionGroupParams) SetPerformSourceSideDeduplication(v bool)`

SetPerformSourceSideDeduplication sets PerformSourceSideDeduplication field to given value.

### HasPerformSourceSideDeduplication

`func (o *PhysicalVolumeProtectionGroupParams) HasPerformSourceSideDeduplication() bool`

HasPerformSourceSideDeduplication returns a boolean if a field has been set.

### SetPerformSourceSideDeduplicationNil

`func (o *PhysicalVolumeProtectionGroupParams) SetPerformSourceSideDeduplicationNil(b bool)`

 SetPerformSourceSideDeduplicationNil sets the value for PerformSourceSideDeduplication to be an explicit nil

### UnsetPerformSourceSideDeduplication
`func (o *PhysicalVolumeProtectionGroupParams) UnsetPerformSourceSideDeduplication()`

UnsetPerformSourceSideDeduplication ensures that no value is present for PerformSourceSideDeduplication, not even an explicit nil
### GetQuiesce

`func (o *PhysicalVolumeProtectionGroupParams) GetQuiesce() bool`

GetQuiesce returns the Quiesce field if non-nil, zero value otherwise.

### GetQuiesceOk

`func (o *PhysicalVolumeProtectionGroupParams) GetQuiesceOk() (*bool, bool)`

GetQuiesceOk returns a tuple with the Quiesce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuiesce

`func (o *PhysicalVolumeProtectionGroupParams) SetQuiesce(v bool)`

SetQuiesce sets Quiesce field to given value.

### HasQuiesce

`func (o *PhysicalVolumeProtectionGroupParams) HasQuiesce() bool`

HasQuiesce returns a boolean if a field has been set.

### SetQuiesceNil

`func (o *PhysicalVolumeProtectionGroupParams) SetQuiesceNil(b bool)`

 SetQuiesceNil sets the value for Quiesce to be an explicit nil

### UnsetQuiesce
`func (o *PhysicalVolumeProtectionGroupParams) UnsetQuiesce()`

UnsetQuiesce ensures that no value is present for Quiesce, not even an explicit nil
### GetContinueOnQuiesceFailure

`func (o *PhysicalVolumeProtectionGroupParams) GetContinueOnQuiesceFailure() bool`

GetContinueOnQuiesceFailure returns the ContinueOnQuiesceFailure field if non-nil, zero value otherwise.

### GetContinueOnQuiesceFailureOk

`func (o *PhysicalVolumeProtectionGroupParams) GetContinueOnQuiesceFailureOk() (*bool, bool)`

GetContinueOnQuiesceFailureOk returns a tuple with the ContinueOnQuiesceFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnQuiesceFailure

`func (o *PhysicalVolumeProtectionGroupParams) SetContinueOnQuiesceFailure(v bool)`

SetContinueOnQuiesceFailure sets ContinueOnQuiesceFailure field to given value.

### HasContinueOnQuiesceFailure

`func (o *PhysicalVolumeProtectionGroupParams) HasContinueOnQuiesceFailure() bool`

HasContinueOnQuiesceFailure returns a boolean if a field has been set.

### SetContinueOnQuiesceFailureNil

`func (o *PhysicalVolumeProtectionGroupParams) SetContinueOnQuiesceFailureNil(b bool)`

 SetContinueOnQuiesceFailureNil sets the value for ContinueOnQuiesceFailure to be an explicit nil

### UnsetContinueOnQuiesceFailure
`func (o *PhysicalVolumeProtectionGroupParams) UnsetContinueOnQuiesceFailure()`

UnsetContinueOnQuiesceFailure ensures that no value is present for ContinueOnQuiesceFailure, not even an explicit nil
### GetIncrementalBackupAfterRestart

`func (o *PhysicalVolumeProtectionGroupParams) GetIncrementalBackupAfterRestart() bool`

GetIncrementalBackupAfterRestart returns the IncrementalBackupAfterRestart field if non-nil, zero value otherwise.

### GetIncrementalBackupAfterRestartOk

`func (o *PhysicalVolumeProtectionGroupParams) GetIncrementalBackupAfterRestartOk() (*bool, bool)`

GetIncrementalBackupAfterRestartOk returns a tuple with the IncrementalBackupAfterRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalBackupAfterRestart

`func (o *PhysicalVolumeProtectionGroupParams) SetIncrementalBackupAfterRestart(v bool)`

SetIncrementalBackupAfterRestart sets IncrementalBackupAfterRestart field to given value.

### HasIncrementalBackupAfterRestart

`func (o *PhysicalVolumeProtectionGroupParams) HasIncrementalBackupAfterRestart() bool`

HasIncrementalBackupAfterRestart returns a boolean if a field has been set.

### SetIncrementalBackupAfterRestartNil

`func (o *PhysicalVolumeProtectionGroupParams) SetIncrementalBackupAfterRestartNil(b bool)`

 SetIncrementalBackupAfterRestartNil sets the value for IncrementalBackupAfterRestart to be an explicit nil

### UnsetIncrementalBackupAfterRestart
`func (o *PhysicalVolumeProtectionGroupParams) UnsetIncrementalBackupAfterRestart()`

UnsetIncrementalBackupAfterRestart ensures that no value is present for IncrementalBackupAfterRestart, not even an explicit nil
### GetPrePostScript

`func (o *PhysicalVolumeProtectionGroupParams) GetPrePostScript() PrePostScriptParams`

GetPrePostScript returns the PrePostScript field if non-nil, zero value otherwise.

### GetPrePostScriptOk

`func (o *PhysicalVolumeProtectionGroupParams) GetPrePostScriptOk() (*PrePostScriptParams, bool)`

GetPrePostScriptOk returns a tuple with the PrePostScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePostScript

`func (o *PhysicalVolumeProtectionGroupParams) SetPrePostScript(v PrePostScriptParams)`

SetPrePostScript sets PrePostScript field to given value.

### HasPrePostScript

`func (o *PhysicalVolumeProtectionGroupParams) HasPrePostScript() bool`

HasPrePostScript returns a boolean if a field has been set.

### GetDedupExclusionSourceIds

`func (o *PhysicalVolumeProtectionGroupParams) GetDedupExclusionSourceIds() []int64`

GetDedupExclusionSourceIds returns the DedupExclusionSourceIds field if non-nil, zero value otherwise.

### GetDedupExclusionSourceIdsOk

`func (o *PhysicalVolumeProtectionGroupParams) GetDedupExclusionSourceIdsOk() (*[]int64, bool)`

GetDedupExclusionSourceIdsOk returns a tuple with the DedupExclusionSourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedupExclusionSourceIds

`func (o *PhysicalVolumeProtectionGroupParams) SetDedupExclusionSourceIds(v []int64)`

SetDedupExclusionSourceIds sets DedupExclusionSourceIds field to given value.

### HasDedupExclusionSourceIds

`func (o *PhysicalVolumeProtectionGroupParams) HasDedupExclusionSourceIds() bool`

HasDedupExclusionSourceIds returns a boolean if a field has been set.

### GetExcludedVssWriters

`func (o *PhysicalVolumeProtectionGroupParams) GetExcludedVssWriters() []string`

GetExcludedVssWriters returns the ExcludedVssWriters field if non-nil, zero value otherwise.

### GetExcludedVssWritersOk

`func (o *PhysicalVolumeProtectionGroupParams) GetExcludedVssWritersOk() (*[]string, bool)`

GetExcludedVssWritersOk returns a tuple with the ExcludedVssWriters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedVssWriters

`func (o *PhysicalVolumeProtectionGroupParams) SetExcludedVssWriters(v []string)`

SetExcludedVssWriters sets ExcludedVssWriters field to given value.

### HasExcludedVssWriters

`func (o *PhysicalVolumeProtectionGroupParams) HasExcludedVssWriters() bool`

HasExcludedVssWriters returns a boolean if a field has been set.

### SetExcludedVssWritersNil

`func (o *PhysicalVolumeProtectionGroupParams) SetExcludedVssWritersNil(b bool)`

 SetExcludedVssWritersNil sets the value for ExcludedVssWriters to be an explicit nil

### UnsetExcludedVssWriters
`func (o *PhysicalVolumeProtectionGroupParams) UnsetExcludedVssWriters()`

UnsetExcludedVssWriters ensures that no value is present for ExcludedVssWriters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


