# HyperVObjectProtectionRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppConsistentSnapshot** | Pointer to **NullableBool** | Specifies whether or not to quiesce apps and the file system in order to take app consistent snapshots. If not specified or false then snapshots will not be app consistent. | [optional] 
**FallbackToCrashConsistentSnapshot** | Pointer to **NullableBool** | Specifies whether or not to fallback to a crash consistent snapshot in the event that an app consistent snapshot fails. | [optional] 
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 
**Objects** | [**[]HyperVObjectProtectionRequest**](HyperVObjectProtectionRequest.md) | Specifies the objects to include in the backup. | 

## Methods

### NewHyperVObjectProtectionRequestParams

`func NewHyperVObjectProtectionRequestParams(objects []HyperVObjectProtectionRequest, ) *HyperVObjectProtectionRequestParams`

NewHyperVObjectProtectionRequestParams instantiates a new HyperVObjectProtectionRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperVObjectProtectionRequestParamsWithDefaults

`func NewHyperVObjectProtectionRequestParamsWithDefaults() *HyperVObjectProtectionRequestParams`

NewHyperVObjectProtectionRequestParamsWithDefaults instantiates a new HyperVObjectProtectionRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppConsistentSnapshot

`func (o *HyperVObjectProtectionRequestParams) GetAppConsistentSnapshot() bool`

GetAppConsistentSnapshot returns the AppConsistentSnapshot field if non-nil, zero value otherwise.

### GetAppConsistentSnapshotOk

`func (o *HyperVObjectProtectionRequestParams) GetAppConsistentSnapshotOk() (*bool, bool)`

GetAppConsistentSnapshotOk returns a tuple with the AppConsistentSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppConsistentSnapshot

`func (o *HyperVObjectProtectionRequestParams) SetAppConsistentSnapshot(v bool)`

SetAppConsistentSnapshot sets AppConsistentSnapshot field to given value.

### HasAppConsistentSnapshot

`func (o *HyperVObjectProtectionRequestParams) HasAppConsistentSnapshot() bool`

HasAppConsistentSnapshot returns a boolean if a field has been set.

### SetAppConsistentSnapshotNil

`func (o *HyperVObjectProtectionRequestParams) SetAppConsistentSnapshotNil(b bool)`

 SetAppConsistentSnapshotNil sets the value for AppConsistentSnapshot to be an explicit nil

### UnsetAppConsistentSnapshot
`func (o *HyperVObjectProtectionRequestParams) UnsetAppConsistentSnapshot()`

UnsetAppConsistentSnapshot ensures that no value is present for AppConsistentSnapshot, not even an explicit nil
### GetFallbackToCrashConsistentSnapshot

`func (o *HyperVObjectProtectionRequestParams) GetFallbackToCrashConsistentSnapshot() bool`

GetFallbackToCrashConsistentSnapshot returns the FallbackToCrashConsistentSnapshot field if non-nil, zero value otherwise.

### GetFallbackToCrashConsistentSnapshotOk

`func (o *HyperVObjectProtectionRequestParams) GetFallbackToCrashConsistentSnapshotOk() (*bool, bool)`

GetFallbackToCrashConsistentSnapshotOk returns a tuple with the FallbackToCrashConsistentSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackToCrashConsistentSnapshot

`func (o *HyperVObjectProtectionRequestParams) SetFallbackToCrashConsistentSnapshot(v bool)`

SetFallbackToCrashConsistentSnapshot sets FallbackToCrashConsistentSnapshot field to given value.

### HasFallbackToCrashConsistentSnapshot

`func (o *HyperVObjectProtectionRequestParams) HasFallbackToCrashConsistentSnapshot() bool`

HasFallbackToCrashConsistentSnapshot returns a boolean if a field has been set.

### SetFallbackToCrashConsistentSnapshotNil

`func (o *HyperVObjectProtectionRequestParams) SetFallbackToCrashConsistentSnapshotNil(b bool)`

 SetFallbackToCrashConsistentSnapshotNil sets the value for FallbackToCrashConsistentSnapshot to be an explicit nil

### UnsetFallbackToCrashConsistentSnapshot
`func (o *HyperVObjectProtectionRequestParams) UnsetFallbackToCrashConsistentSnapshot()`

UnsetFallbackToCrashConsistentSnapshot ensures that no value is present for FallbackToCrashConsistentSnapshot, not even an explicit nil
### GetIndexingPolicy

`func (o *HyperVObjectProtectionRequestParams) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *HyperVObjectProtectionRequestParams) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *HyperVObjectProtectionRequestParams) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *HyperVObjectProtectionRequestParams) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetObjects

`func (o *HyperVObjectProtectionRequestParams) GetObjects() []HyperVObjectProtectionRequest`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *HyperVObjectProtectionRequestParams) GetObjectsOk() (*[]HyperVObjectProtectionRequest, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *HyperVObjectProtectionRequestParams) SetObjects(v []HyperVObjectProtectionRequest)`

SetObjects sets Objects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


