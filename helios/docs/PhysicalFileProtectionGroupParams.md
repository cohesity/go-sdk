# PhysicalFileProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | [**[]PhysicalFileProtectionGroupObjectParams**](PhysicalFileProtectionGroupObjectParams.md) | Specifies the list of objects protected by this Protection Group. | 
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 
**PerformSourceSideDeduplication** | Pointer to **NullableBool** | Specifies whether or not to perform source side deduplication on this Protection Group. | [optional] 
**Quiesce** | Pointer to **NullableBool** | Specifies Whether to take app-consistent snapshots by quiescing apps and the filesystem before taking a backup. | [optional] 
**ContinueOnQuiesceFailure** | Pointer to **NullableBool** | Specifies whether to continue backing up on quiesce failure. | [optional] 
**PrePostScript** | Pointer to [**PrePostScriptParams**](PrePostScriptParams.md) |  | [optional] 
**DedupExclusionSourceIds** | Pointer to **[]int64** | Specifies ids of sources for which deduplication has to be disabled. | [optional] 
**GlobalExcludePaths** | Pointer to **[]string** | Specifies global exclude filters which are applied to all sources in a job. | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


