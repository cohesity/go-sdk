# HyperVObjectProtectionUpdateRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppConsistentSnapshot** | Pointer to **NullableBool** | Specifies whether or not to quiesce apps and the file system in order to take app consistent snapshots. If not specified or false then snapshots will not be app consistent. | [optional] 
**FallbackToCrashConsistentSnapshot** | Pointer to **NullableBool** | Specifies whether or not to fallback to a crash consistent snapshot in the event that an app consistent snapshot fails. | [optional] 
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 
**BackupsCopyOnly** | Pointer to **NullableBool** | Specifies whether the backups should be copy-only. | [optional] 
**ExcludeObjectIds** | Pointer to **[]int64** | Specifies the list of IDs of the objects to not be protected by this Protection Group. This can be used to ignore specific objects under a parent object which has been included for protection. | [optional] 

## Methods

### NewHyperVObjectProtectionUpdateRequestParams

`func NewHyperVObjectProtectionUpdateRequestParams() *HyperVObjectProtectionUpdateRequestParams`

NewHyperVObjectProtectionUpdateRequestParams instantiates a new HyperVObjectProtectionUpdateRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperVObjectProtectionUpdateRequestParamsWithDefaults

`func NewHyperVObjectProtectionUpdateRequestParamsWithDefaults() *HyperVObjectProtectionUpdateRequestParams`

NewHyperVObjectProtectionUpdateRequestParamsWithDefaults instantiates a new HyperVObjectProtectionUpdateRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppConsistentSnapshot

`func (o *HyperVObjectProtectionUpdateRequestParams) GetAppConsistentSnapshot() bool`

GetAppConsistentSnapshot returns the AppConsistentSnapshot field if non-nil, zero value otherwise.

### GetAppConsistentSnapshotOk

`func (o *HyperVObjectProtectionUpdateRequestParams) GetAppConsistentSnapshotOk() (*bool, bool)`

GetAppConsistentSnapshotOk returns a tuple with the AppConsistentSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppConsistentSnapshot

`func (o *HyperVObjectProtectionUpdateRequestParams) SetAppConsistentSnapshot(v bool)`

SetAppConsistentSnapshot sets AppConsistentSnapshot field to given value.

### HasAppConsistentSnapshot

`func (o *HyperVObjectProtectionUpdateRequestParams) HasAppConsistentSnapshot() bool`

HasAppConsistentSnapshot returns a boolean if a field has been set.

### SetAppConsistentSnapshotNil

`func (o *HyperVObjectProtectionUpdateRequestParams) SetAppConsistentSnapshotNil(b bool)`

 SetAppConsistentSnapshotNil sets the value for AppConsistentSnapshot to be an explicit nil

### UnsetAppConsistentSnapshot
`func (o *HyperVObjectProtectionUpdateRequestParams) UnsetAppConsistentSnapshot()`

UnsetAppConsistentSnapshot ensures that no value is present for AppConsistentSnapshot, not even an explicit nil
### GetFallbackToCrashConsistentSnapshot

`func (o *HyperVObjectProtectionUpdateRequestParams) GetFallbackToCrashConsistentSnapshot() bool`

GetFallbackToCrashConsistentSnapshot returns the FallbackToCrashConsistentSnapshot field if non-nil, zero value otherwise.

### GetFallbackToCrashConsistentSnapshotOk

`func (o *HyperVObjectProtectionUpdateRequestParams) GetFallbackToCrashConsistentSnapshotOk() (*bool, bool)`

GetFallbackToCrashConsistentSnapshotOk returns a tuple with the FallbackToCrashConsistentSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackToCrashConsistentSnapshot

`func (o *HyperVObjectProtectionUpdateRequestParams) SetFallbackToCrashConsistentSnapshot(v bool)`

SetFallbackToCrashConsistentSnapshot sets FallbackToCrashConsistentSnapshot field to given value.

### HasFallbackToCrashConsistentSnapshot

`func (o *HyperVObjectProtectionUpdateRequestParams) HasFallbackToCrashConsistentSnapshot() bool`

HasFallbackToCrashConsistentSnapshot returns a boolean if a field has been set.

### SetFallbackToCrashConsistentSnapshotNil

`func (o *HyperVObjectProtectionUpdateRequestParams) SetFallbackToCrashConsistentSnapshotNil(b bool)`

 SetFallbackToCrashConsistentSnapshotNil sets the value for FallbackToCrashConsistentSnapshot to be an explicit nil

### UnsetFallbackToCrashConsistentSnapshot
`func (o *HyperVObjectProtectionUpdateRequestParams) UnsetFallbackToCrashConsistentSnapshot()`

UnsetFallbackToCrashConsistentSnapshot ensures that no value is present for FallbackToCrashConsistentSnapshot, not even an explicit nil
### GetIndexingPolicy

`func (o *HyperVObjectProtectionUpdateRequestParams) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *HyperVObjectProtectionUpdateRequestParams) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *HyperVObjectProtectionUpdateRequestParams) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *HyperVObjectProtectionUpdateRequestParams) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetBackupsCopyOnly

`func (o *HyperVObjectProtectionUpdateRequestParams) GetBackupsCopyOnly() bool`

GetBackupsCopyOnly returns the BackupsCopyOnly field if non-nil, zero value otherwise.

### GetBackupsCopyOnlyOk

`func (o *HyperVObjectProtectionUpdateRequestParams) GetBackupsCopyOnlyOk() (*bool, bool)`

GetBackupsCopyOnlyOk returns a tuple with the BackupsCopyOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupsCopyOnly

`func (o *HyperVObjectProtectionUpdateRequestParams) SetBackupsCopyOnly(v bool)`

SetBackupsCopyOnly sets BackupsCopyOnly field to given value.

### HasBackupsCopyOnly

`func (o *HyperVObjectProtectionUpdateRequestParams) HasBackupsCopyOnly() bool`

HasBackupsCopyOnly returns a boolean if a field has been set.

### SetBackupsCopyOnlyNil

`func (o *HyperVObjectProtectionUpdateRequestParams) SetBackupsCopyOnlyNil(b bool)`

 SetBackupsCopyOnlyNil sets the value for BackupsCopyOnly to be an explicit nil

### UnsetBackupsCopyOnly
`func (o *HyperVObjectProtectionUpdateRequestParams) UnsetBackupsCopyOnly()`

UnsetBackupsCopyOnly ensures that no value is present for BackupsCopyOnly, not even an explicit nil
### GetExcludeObjectIds

`func (o *HyperVObjectProtectionUpdateRequestParams) GetExcludeObjectIds() []*int64`

GetExcludeObjectIds returns the ExcludeObjectIds field if non-nil, zero value otherwise.

### GetExcludeObjectIdsOk

`func (o *HyperVObjectProtectionUpdateRequestParams) GetExcludeObjectIdsOk() (*[]*int64, bool)`

GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeObjectIds

`func (o *HyperVObjectProtectionUpdateRequestParams) SetExcludeObjectIds(v []*int64)`

SetExcludeObjectIds sets ExcludeObjectIds field to given value.

### HasExcludeObjectIds

`func (o *HyperVObjectProtectionUpdateRequestParams) HasExcludeObjectIds() bool`

HasExcludeObjectIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


