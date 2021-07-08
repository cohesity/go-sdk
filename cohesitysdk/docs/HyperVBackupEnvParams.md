# HyperVBackupEnvParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowCrashConsistentSnapshot** | Pointer to **NullableBool** | Whether to fallback to take a crash-consistent snapshot incase taking an app-consistent snapshot fails. | [optional] 
**BackupJobType** | Pointer to **NullableInt32** | The type of backup job to use. Default is to auto-detect the best type to use based on the VMs to backup. End user may select RCT or VSS also. | [optional] 

## Methods

### NewHyperVBackupEnvParams

`func NewHyperVBackupEnvParams() *HyperVBackupEnvParams`

NewHyperVBackupEnvParams instantiates a new HyperVBackupEnvParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperVBackupEnvParamsWithDefaults

`func NewHyperVBackupEnvParamsWithDefaults() *HyperVBackupEnvParams`

NewHyperVBackupEnvParamsWithDefaults instantiates a new HyperVBackupEnvParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowCrashConsistentSnapshot

`func (o *HyperVBackupEnvParams) GetAllowCrashConsistentSnapshot() bool`

GetAllowCrashConsistentSnapshot returns the AllowCrashConsistentSnapshot field if non-nil, zero value otherwise.

### GetAllowCrashConsistentSnapshotOk

`func (o *HyperVBackupEnvParams) GetAllowCrashConsistentSnapshotOk() (*bool, bool)`

GetAllowCrashConsistentSnapshotOk returns a tuple with the AllowCrashConsistentSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCrashConsistentSnapshot

`func (o *HyperVBackupEnvParams) SetAllowCrashConsistentSnapshot(v bool)`

SetAllowCrashConsistentSnapshot sets AllowCrashConsistentSnapshot field to given value.

### HasAllowCrashConsistentSnapshot

`func (o *HyperVBackupEnvParams) HasAllowCrashConsistentSnapshot() bool`

HasAllowCrashConsistentSnapshot returns a boolean if a field has been set.

### SetAllowCrashConsistentSnapshotNil

`func (o *HyperVBackupEnvParams) SetAllowCrashConsistentSnapshotNil(b bool)`

 SetAllowCrashConsistentSnapshotNil sets the value for AllowCrashConsistentSnapshot to be an explicit nil

### UnsetAllowCrashConsistentSnapshot
`func (o *HyperVBackupEnvParams) UnsetAllowCrashConsistentSnapshot()`

UnsetAllowCrashConsistentSnapshot ensures that no value is present for AllowCrashConsistentSnapshot, not even an explicit nil
### GetBackupJobType

`func (o *HyperVBackupEnvParams) GetBackupJobType() int32`

GetBackupJobType returns the BackupJobType field if non-nil, zero value otherwise.

### GetBackupJobTypeOk

`func (o *HyperVBackupEnvParams) GetBackupJobTypeOk() (*int32, bool)`

GetBackupJobTypeOk returns a tuple with the BackupJobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupJobType

`func (o *HyperVBackupEnvParams) SetBackupJobType(v int32)`

SetBackupJobType sets BackupJobType field to given value.

### HasBackupJobType

`func (o *HyperVBackupEnvParams) HasBackupJobType() bool`

HasBackupJobType returns a boolean if a field has been set.

### SetBackupJobTypeNil

`func (o *HyperVBackupEnvParams) SetBackupJobTypeNil(b bool)`

 SetBackupJobTypeNil sets the value for BackupJobType to be an explicit nil

### UnsetBackupJobType
`func (o *HyperVBackupEnvParams) UnsetBackupJobType()`

UnsetBackupJobType ensures that no value is present for BackupJobType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


