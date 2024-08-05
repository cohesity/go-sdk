# BackupPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bmr** | Pointer to [**BmrBackupPolicy**](BmrBackupPolicy.md) |  | [optional] 
**Cdp** | Pointer to [**CdpBackupPolicy**](CdpBackupPolicy.md) |  | [optional] 
**Log** | Pointer to [**LogBackupPolicy**](LogBackupPolicy.md) |  | [optional] 
**Regular** | [**RegularBackupPolicy**](RegularBackupPolicy.md) |  | 
**RunTimeouts** | Pointer to [**[]CancellationTimeoutParams**](CancellationTimeoutParams.md) | Specifies the backup timeouts for different type of runs(kFull, kRegular etc.). | [optional] 
**StorageArraySnapshot** | Pointer to [**StorageArraySnapshotBackupPolicy**](StorageArraySnapshotBackupPolicy.md) |  | [optional] 

## Methods

### NewBackupPolicy

`func NewBackupPolicy(regular RegularBackupPolicy, ) *BackupPolicy`

NewBackupPolicy instantiates a new BackupPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupPolicyWithDefaults

`func NewBackupPolicyWithDefaults() *BackupPolicy`

NewBackupPolicyWithDefaults instantiates a new BackupPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBmr

`func (o *BackupPolicy) GetBmr() BmrBackupPolicy`

GetBmr returns the Bmr field if non-nil, zero value otherwise.

### GetBmrOk

`func (o *BackupPolicy) GetBmrOk() (*BmrBackupPolicy, bool)`

GetBmrOk returns a tuple with the Bmr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmr

`func (o *BackupPolicy) SetBmr(v BmrBackupPolicy)`

SetBmr sets Bmr field to given value.

### HasBmr

`func (o *BackupPolicy) HasBmr() bool`

HasBmr returns a boolean if a field has been set.

### GetCdp

`func (o *BackupPolicy) GetCdp() CdpBackupPolicy`

GetCdp returns the Cdp field if non-nil, zero value otherwise.

### GetCdpOk

`func (o *BackupPolicy) GetCdpOk() (*CdpBackupPolicy, bool)`

GetCdpOk returns a tuple with the Cdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdp

`func (o *BackupPolicy) SetCdp(v CdpBackupPolicy)`

SetCdp sets Cdp field to given value.

### HasCdp

`func (o *BackupPolicy) HasCdp() bool`

HasCdp returns a boolean if a field has been set.

### GetLog

`func (o *BackupPolicy) GetLog() LogBackupPolicy`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *BackupPolicy) GetLogOk() (*LogBackupPolicy, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *BackupPolicy) SetLog(v LogBackupPolicy)`

SetLog sets Log field to given value.

### HasLog

`func (o *BackupPolicy) HasLog() bool`

HasLog returns a boolean if a field has been set.

### GetRegular

`func (o *BackupPolicy) GetRegular() RegularBackupPolicy`

GetRegular returns the Regular field if non-nil, zero value otherwise.

### GetRegularOk

`func (o *BackupPolicy) GetRegularOk() (*RegularBackupPolicy, bool)`

GetRegularOk returns a tuple with the Regular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegular

`func (o *BackupPolicy) SetRegular(v RegularBackupPolicy)`

SetRegular sets Regular field to given value.


### GetRunTimeouts

`func (o *BackupPolicy) GetRunTimeouts() []CancellationTimeoutParams`

GetRunTimeouts returns the RunTimeouts field if non-nil, zero value otherwise.

### GetRunTimeoutsOk

`func (o *BackupPolicy) GetRunTimeoutsOk() (*[]CancellationTimeoutParams, bool)`

GetRunTimeoutsOk returns a tuple with the RunTimeouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunTimeouts

`func (o *BackupPolicy) SetRunTimeouts(v []CancellationTimeoutParams)`

SetRunTimeouts sets RunTimeouts field to given value.

### HasRunTimeouts

`func (o *BackupPolicy) HasRunTimeouts() bool`

HasRunTimeouts returns a boolean if a field has been set.

### SetRunTimeoutsNil

`func (o *BackupPolicy) SetRunTimeoutsNil(b bool)`

 SetRunTimeoutsNil sets the value for RunTimeouts to be an explicit nil

### UnsetRunTimeouts
`func (o *BackupPolicy) UnsetRunTimeouts()`

UnsetRunTimeouts ensures that no value is present for RunTimeouts, not even an explicit nil
### GetStorageArraySnapshot

`func (o *BackupPolicy) GetStorageArraySnapshot() StorageArraySnapshotBackupPolicy`

GetStorageArraySnapshot returns the StorageArraySnapshot field if non-nil, zero value otherwise.

### GetStorageArraySnapshotOk

`func (o *BackupPolicy) GetStorageArraySnapshotOk() (*StorageArraySnapshotBackupPolicy, bool)`

GetStorageArraySnapshotOk returns a tuple with the StorageArraySnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageArraySnapshot

`func (o *BackupPolicy) SetStorageArraySnapshot(v StorageArraySnapshotBackupPolicy)`

SetStorageArraySnapshot sets StorageArraySnapshot field to given value.

### HasStorageArraySnapshot

`func (o *BackupPolicy) HasStorageArraySnapshot() bool`

HasStorageArraySnapshot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


