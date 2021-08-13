# CommonHyperVProtectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppConsistentSnapshot** | Pointer to **NullableBool** | Specifies whether or not to quiesce apps and the file system in order to take app consistent snapshots. If not specified or false then snapshots will not be app consistent. | [optional] 
**FallbackToCrashConsistentSnapshot** | Pointer to **NullableBool** | Specifies whether or not to fallback to a crash consistent snapshot in the event that an app consistent snapshot fails. | [optional] 
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 

## Methods

### NewCommonHyperVProtectionParams

`func NewCommonHyperVProtectionParams() *CommonHyperVProtectionParams`

NewCommonHyperVProtectionParams instantiates a new CommonHyperVProtectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonHyperVProtectionParamsWithDefaults

`func NewCommonHyperVProtectionParamsWithDefaults() *CommonHyperVProtectionParams`

NewCommonHyperVProtectionParamsWithDefaults instantiates a new CommonHyperVProtectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppConsistentSnapshot

`func (o *CommonHyperVProtectionParams) GetAppConsistentSnapshot() bool`

GetAppConsistentSnapshot returns the AppConsistentSnapshot field if non-nil, zero value otherwise.

### GetAppConsistentSnapshotOk

`func (o *CommonHyperVProtectionParams) GetAppConsistentSnapshotOk() (*bool, bool)`

GetAppConsistentSnapshotOk returns a tuple with the AppConsistentSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppConsistentSnapshot

`func (o *CommonHyperVProtectionParams) SetAppConsistentSnapshot(v bool)`

SetAppConsistentSnapshot sets AppConsistentSnapshot field to given value.

### HasAppConsistentSnapshot

`func (o *CommonHyperVProtectionParams) HasAppConsistentSnapshot() bool`

HasAppConsistentSnapshot returns a boolean if a field has been set.

### SetAppConsistentSnapshotNil

`func (o *CommonHyperVProtectionParams) SetAppConsistentSnapshotNil(b bool)`

 SetAppConsistentSnapshotNil sets the value for AppConsistentSnapshot to be an explicit nil

### UnsetAppConsistentSnapshot
`func (o *CommonHyperVProtectionParams) UnsetAppConsistentSnapshot()`

UnsetAppConsistentSnapshot ensures that no value is present for AppConsistentSnapshot, not even an explicit nil
### GetFallbackToCrashConsistentSnapshot

`func (o *CommonHyperVProtectionParams) GetFallbackToCrashConsistentSnapshot() bool`

GetFallbackToCrashConsistentSnapshot returns the FallbackToCrashConsistentSnapshot field if non-nil, zero value otherwise.

### GetFallbackToCrashConsistentSnapshotOk

`func (o *CommonHyperVProtectionParams) GetFallbackToCrashConsistentSnapshotOk() (*bool, bool)`

GetFallbackToCrashConsistentSnapshotOk returns a tuple with the FallbackToCrashConsistentSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackToCrashConsistentSnapshot

`func (o *CommonHyperVProtectionParams) SetFallbackToCrashConsistentSnapshot(v bool)`

SetFallbackToCrashConsistentSnapshot sets FallbackToCrashConsistentSnapshot field to given value.

### HasFallbackToCrashConsistentSnapshot

`func (o *CommonHyperVProtectionParams) HasFallbackToCrashConsistentSnapshot() bool`

HasFallbackToCrashConsistentSnapshot returns a boolean if a field has been set.

### SetFallbackToCrashConsistentSnapshotNil

`func (o *CommonHyperVProtectionParams) SetFallbackToCrashConsistentSnapshotNil(b bool)`

 SetFallbackToCrashConsistentSnapshotNil sets the value for FallbackToCrashConsistentSnapshot to be an explicit nil

### UnsetFallbackToCrashConsistentSnapshot
`func (o *CommonHyperVProtectionParams) UnsetFallbackToCrashConsistentSnapshot()`

UnsetFallbackToCrashConsistentSnapshot ensures that no value is present for FallbackToCrashConsistentSnapshot, not even an explicit nil
### GetIndexingPolicy

`func (o *CommonHyperVProtectionParams) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *CommonHyperVProtectionParams) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *CommonHyperVProtectionParams) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *CommonHyperVProtectionParams) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


