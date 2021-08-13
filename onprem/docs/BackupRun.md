# BackupRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapshotInfo** | Pointer to [**SnapshotInfo**](SnapshotInfo.md) |  | [optional] 
**FailedAttempts** | Pointer to [**[]BackupAttempt**](BackupAttempt.md) | Failed backup attempts for an object. | [optional] 

## Methods

### NewBackupRun

`func NewBackupRun() *BackupRun`

NewBackupRun instantiates a new BackupRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRunWithDefaults

`func NewBackupRunWithDefaults() *BackupRun`

NewBackupRunWithDefaults instantiates a new BackupRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshotInfo

`func (o *BackupRun) GetSnapshotInfo() SnapshotInfo`

GetSnapshotInfo returns the SnapshotInfo field if non-nil, zero value otherwise.

### GetSnapshotInfoOk

`func (o *BackupRun) GetSnapshotInfoOk() (*SnapshotInfo, bool)`

GetSnapshotInfoOk returns a tuple with the SnapshotInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotInfo

`func (o *BackupRun) SetSnapshotInfo(v SnapshotInfo)`

SetSnapshotInfo sets SnapshotInfo field to given value.

### HasSnapshotInfo

`func (o *BackupRun) HasSnapshotInfo() bool`

HasSnapshotInfo returns a boolean if a field has been set.

### GetFailedAttempts

`func (o *BackupRun) GetFailedAttempts() []BackupAttempt`

GetFailedAttempts returns the FailedAttempts field if non-nil, zero value otherwise.

### GetFailedAttemptsOk

`func (o *BackupRun) GetFailedAttemptsOk() (*[]BackupAttempt, bool)`

GetFailedAttemptsOk returns a tuple with the FailedAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAttempts

`func (o *BackupRun) SetFailedAttempts(v []BackupAttempt)`

SetFailedAttempts sets FailedAttempts field to given value.

### HasFailedAttempts

`func (o *BackupRun) HasFailedAttempts() bool`

HasFailedAttempts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


