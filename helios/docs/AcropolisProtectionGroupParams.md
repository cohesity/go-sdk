# AcropolisProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | [**[]AcropolisProtectionGroupObjectParams**](AcropolisProtectionGroupObjectParams.md) | Specifies the objects included in the Protection Group. | 
**SourceId** | Pointer to **NullableInt64** | Specifies the id of the parent of the objects. | [optional] [readonly] 
**SourceName** | Pointer to **NullableString** | Specifies the name of the parent of the objects. | [optional] [readonly] 
**ExcludeObjectIds** | Pointer to **[]int64** | Specifies the object ids to be excluded in the Protection Group. | [optional] 
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 
**AppConsistentSnapshot** | Pointer to **NullableBool** | Specifies whether or not to quiesce apps and the file system in order to take app consistent snapshots. If not specified or false then snapshots will not be app consistent. | [optional] 
**ContinueOnQuiesceFailure** | Pointer to **NullableBool** | Specifies whether to continue backing up on quiesce failure | [optional] 

## Methods

### NewAcropolisProtectionGroupParams

`func NewAcropolisProtectionGroupParams(objects []AcropolisProtectionGroupObjectParams, ) *AcropolisProtectionGroupParams`

NewAcropolisProtectionGroupParams instantiates a new AcropolisProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcropolisProtectionGroupParamsWithDefaults

`func NewAcropolisProtectionGroupParamsWithDefaults() *AcropolisProtectionGroupParams`

NewAcropolisProtectionGroupParamsWithDefaults instantiates a new AcropolisProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *AcropolisProtectionGroupParams) GetObjects() []AcropolisProtectionGroupObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *AcropolisProtectionGroupParams) GetObjectsOk() (*[]AcropolisProtectionGroupObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *AcropolisProtectionGroupParams) SetObjects(v []AcropolisProtectionGroupObjectParams)`

SetObjects sets Objects field to given value.


### GetSourceId

`func (o *AcropolisProtectionGroupParams) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *AcropolisProtectionGroupParams) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *AcropolisProtectionGroupParams) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *AcropolisProtectionGroupParams) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *AcropolisProtectionGroupParams) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *AcropolisProtectionGroupParams) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *AcropolisProtectionGroupParams) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *AcropolisProtectionGroupParams) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *AcropolisProtectionGroupParams) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *AcropolisProtectionGroupParams) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *AcropolisProtectionGroupParams) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *AcropolisProtectionGroupParams) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetExcludeObjectIds

`func (o *AcropolisProtectionGroupParams) GetExcludeObjectIds() []int64`

GetExcludeObjectIds returns the ExcludeObjectIds field if non-nil, zero value otherwise.

### GetExcludeObjectIdsOk

`func (o *AcropolisProtectionGroupParams) GetExcludeObjectIdsOk() (*[]int64, bool)`

GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeObjectIds

`func (o *AcropolisProtectionGroupParams) SetExcludeObjectIds(v []int64)`

SetExcludeObjectIds sets ExcludeObjectIds field to given value.

### HasExcludeObjectIds

`func (o *AcropolisProtectionGroupParams) HasExcludeObjectIds() bool`

HasExcludeObjectIds returns a boolean if a field has been set.

### SetExcludeObjectIdsNil

`func (o *AcropolisProtectionGroupParams) SetExcludeObjectIdsNil(b bool)`

 SetExcludeObjectIdsNil sets the value for ExcludeObjectIds to be an explicit nil

### UnsetExcludeObjectIds
`func (o *AcropolisProtectionGroupParams) UnsetExcludeObjectIds()`

UnsetExcludeObjectIds ensures that no value is present for ExcludeObjectIds, not even an explicit nil
### GetIndexingPolicy

`func (o *AcropolisProtectionGroupParams) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *AcropolisProtectionGroupParams) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *AcropolisProtectionGroupParams) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *AcropolisProtectionGroupParams) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetAppConsistentSnapshot

`func (o *AcropolisProtectionGroupParams) GetAppConsistentSnapshot() bool`

GetAppConsistentSnapshot returns the AppConsistentSnapshot field if non-nil, zero value otherwise.

### GetAppConsistentSnapshotOk

`func (o *AcropolisProtectionGroupParams) GetAppConsistentSnapshotOk() (*bool, bool)`

GetAppConsistentSnapshotOk returns a tuple with the AppConsistentSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppConsistentSnapshot

`func (o *AcropolisProtectionGroupParams) SetAppConsistentSnapshot(v bool)`

SetAppConsistentSnapshot sets AppConsistentSnapshot field to given value.

### HasAppConsistentSnapshot

`func (o *AcropolisProtectionGroupParams) HasAppConsistentSnapshot() bool`

HasAppConsistentSnapshot returns a boolean if a field has been set.

### SetAppConsistentSnapshotNil

`func (o *AcropolisProtectionGroupParams) SetAppConsistentSnapshotNil(b bool)`

 SetAppConsistentSnapshotNil sets the value for AppConsistentSnapshot to be an explicit nil

### UnsetAppConsistentSnapshot
`func (o *AcropolisProtectionGroupParams) UnsetAppConsistentSnapshot()`

UnsetAppConsistentSnapshot ensures that no value is present for AppConsistentSnapshot, not even an explicit nil
### GetContinueOnQuiesceFailure

`func (o *AcropolisProtectionGroupParams) GetContinueOnQuiesceFailure() bool`

GetContinueOnQuiesceFailure returns the ContinueOnQuiesceFailure field if non-nil, zero value otherwise.

### GetContinueOnQuiesceFailureOk

`func (o *AcropolisProtectionGroupParams) GetContinueOnQuiesceFailureOk() (*bool, bool)`

GetContinueOnQuiesceFailureOk returns a tuple with the ContinueOnQuiesceFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnQuiesceFailure

`func (o *AcropolisProtectionGroupParams) SetContinueOnQuiesceFailure(v bool)`

SetContinueOnQuiesceFailure sets ContinueOnQuiesceFailure field to given value.

### HasContinueOnQuiesceFailure

`func (o *AcropolisProtectionGroupParams) HasContinueOnQuiesceFailure() bool`

HasContinueOnQuiesceFailure returns a boolean if a field has been set.

### SetContinueOnQuiesceFailureNil

`func (o *AcropolisProtectionGroupParams) SetContinueOnQuiesceFailureNil(b bool)`

 SetContinueOnQuiesceFailureNil sets the value for ContinueOnQuiesceFailure to be an explicit nil

### UnsetContinueOnQuiesceFailure
`func (o *AcropolisProtectionGroupParams) UnsetContinueOnQuiesceFailure()`

UnsetContinueOnQuiesceFailure ensures that no value is present for ContinueOnQuiesceFailure, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


