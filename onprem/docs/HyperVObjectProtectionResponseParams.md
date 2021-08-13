# HyperVObjectProtectionResponseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupsCopyOnly** | Pointer to **NullableBool** | Specifies whether the backups should be copy-only. | [optional] 
**ExcludeObjectIds** | Pointer to **[]int64** | Specifies the list of IDs of the objects to not be protected by this Protection Group. This can be used to ignore specific objects under a parent object which has been included for protection. | [optional] 
**AppConsistentSnapshot** | Pointer to **NullableBool** | Specifies whether or not to quiesce apps and the file system in order to take app consistent snapshots. If not specified or false then snapshots will not be app consistent. | [optional] 
**FallbackToCrashConsistentSnapshot** | Pointer to **NullableBool** | Specifies whether or not to fallback to a crash consistent snapshot in the event that an app consistent snapshot fails. | [optional] 
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 

## Methods

### NewHyperVObjectProtectionResponseParams

`func NewHyperVObjectProtectionResponseParams() *HyperVObjectProtectionResponseParams`

NewHyperVObjectProtectionResponseParams instantiates a new HyperVObjectProtectionResponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperVObjectProtectionResponseParamsWithDefaults

`func NewHyperVObjectProtectionResponseParamsWithDefaults() *HyperVObjectProtectionResponseParams`

NewHyperVObjectProtectionResponseParamsWithDefaults instantiates a new HyperVObjectProtectionResponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupsCopyOnly

`func (o *HyperVObjectProtectionResponseParams) GetBackupsCopyOnly() bool`

GetBackupsCopyOnly returns the BackupsCopyOnly field if non-nil, zero value otherwise.

### GetBackupsCopyOnlyOk

`func (o *HyperVObjectProtectionResponseParams) GetBackupsCopyOnlyOk() (*bool, bool)`

GetBackupsCopyOnlyOk returns a tuple with the BackupsCopyOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupsCopyOnly

`func (o *HyperVObjectProtectionResponseParams) SetBackupsCopyOnly(v bool)`

SetBackupsCopyOnly sets BackupsCopyOnly field to given value.

### HasBackupsCopyOnly

`func (o *HyperVObjectProtectionResponseParams) HasBackupsCopyOnly() bool`

HasBackupsCopyOnly returns a boolean if a field has been set.

### SetBackupsCopyOnlyNil

`func (o *HyperVObjectProtectionResponseParams) SetBackupsCopyOnlyNil(b bool)`

 SetBackupsCopyOnlyNil sets the value for BackupsCopyOnly to be an explicit nil

### UnsetBackupsCopyOnly
`func (o *HyperVObjectProtectionResponseParams) UnsetBackupsCopyOnly()`

UnsetBackupsCopyOnly ensures that no value is present for BackupsCopyOnly, not even an explicit nil
### GetExcludeObjectIds

`func (o *HyperVObjectProtectionResponseParams) GetExcludeObjectIds() []int64`

GetExcludeObjectIds returns the ExcludeObjectIds field if non-nil, zero value otherwise.

### GetExcludeObjectIdsOk

`func (o *HyperVObjectProtectionResponseParams) GetExcludeObjectIdsOk() (*[]int64, bool)`

GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeObjectIds

`func (o *HyperVObjectProtectionResponseParams) SetExcludeObjectIds(v []int64)`

SetExcludeObjectIds sets ExcludeObjectIds field to given value.

### HasExcludeObjectIds

`func (o *HyperVObjectProtectionResponseParams) HasExcludeObjectIds() bool`

HasExcludeObjectIds returns a boolean if a field has been set.

### GetAppConsistentSnapshot

`func (o *HyperVObjectProtectionResponseParams) GetAppConsistentSnapshot() bool`

GetAppConsistentSnapshot returns the AppConsistentSnapshot field if non-nil, zero value otherwise.

### GetAppConsistentSnapshotOk

`func (o *HyperVObjectProtectionResponseParams) GetAppConsistentSnapshotOk() (*bool, bool)`

GetAppConsistentSnapshotOk returns a tuple with the AppConsistentSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppConsistentSnapshot

`func (o *HyperVObjectProtectionResponseParams) SetAppConsistentSnapshot(v bool)`

SetAppConsistentSnapshot sets AppConsistentSnapshot field to given value.

### HasAppConsistentSnapshot

`func (o *HyperVObjectProtectionResponseParams) HasAppConsistentSnapshot() bool`

HasAppConsistentSnapshot returns a boolean if a field has been set.

### SetAppConsistentSnapshotNil

`func (o *HyperVObjectProtectionResponseParams) SetAppConsistentSnapshotNil(b bool)`

 SetAppConsistentSnapshotNil sets the value for AppConsistentSnapshot to be an explicit nil

### UnsetAppConsistentSnapshot
`func (o *HyperVObjectProtectionResponseParams) UnsetAppConsistentSnapshot()`

UnsetAppConsistentSnapshot ensures that no value is present for AppConsistentSnapshot, not even an explicit nil
### GetFallbackToCrashConsistentSnapshot

`func (o *HyperVObjectProtectionResponseParams) GetFallbackToCrashConsistentSnapshot() bool`

GetFallbackToCrashConsistentSnapshot returns the FallbackToCrashConsistentSnapshot field if non-nil, zero value otherwise.

### GetFallbackToCrashConsistentSnapshotOk

`func (o *HyperVObjectProtectionResponseParams) GetFallbackToCrashConsistentSnapshotOk() (*bool, bool)`

GetFallbackToCrashConsistentSnapshotOk returns a tuple with the FallbackToCrashConsistentSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackToCrashConsistentSnapshot

`func (o *HyperVObjectProtectionResponseParams) SetFallbackToCrashConsistentSnapshot(v bool)`

SetFallbackToCrashConsistentSnapshot sets FallbackToCrashConsistentSnapshot field to given value.

### HasFallbackToCrashConsistentSnapshot

`func (o *HyperVObjectProtectionResponseParams) HasFallbackToCrashConsistentSnapshot() bool`

HasFallbackToCrashConsistentSnapshot returns a boolean if a field has been set.

### SetFallbackToCrashConsistentSnapshotNil

`func (o *HyperVObjectProtectionResponseParams) SetFallbackToCrashConsistentSnapshotNil(b bool)`

 SetFallbackToCrashConsistentSnapshotNil sets the value for FallbackToCrashConsistentSnapshot to be an explicit nil

### UnsetFallbackToCrashConsistentSnapshot
`func (o *HyperVObjectProtectionResponseParams) UnsetFallbackToCrashConsistentSnapshot()`

UnsetFallbackToCrashConsistentSnapshot ensures that no value is present for FallbackToCrashConsistentSnapshot, not even an explicit nil
### GetIndexingPolicy

`func (o *HyperVObjectProtectionResponseParams) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *HyperVObjectProtectionResponseParams) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *HyperVObjectProtectionResponseParams) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *HyperVObjectProtectionResponseParams) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


