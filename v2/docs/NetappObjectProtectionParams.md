# NetappObjectProtectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupExistingSnapshot** | Pointer to **NullableBool** | Specifies that snapshot label is not set for Data-Protect Netapp Volumes backup. If field is set to true, existing oldest snapshot is used for backup and subsequent incremental will be selected in ascending order of snapshot create time on the source. If snapshot label is set, this field is set to false. | [optional] 
**ContinuousSnapshots** | Pointer to [**ContinuousSnapshotParams**](ContinuousSnapshotParams.md) |  | [optional] 
**ExcludeObjectIds** | Pointer to **[]int64** | Specifies the objects to be excluded in the Protection. | [optional] 
**Protocol** | Pointer to **NullableString** | Specifies the protocol of the NAS device being backed up. | [optional] 
**SnapshotLabel** | Pointer to [**SnapshotLabel**](SnapshotLabel.md) |  | [optional] 

## Methods

### NewNetappObjectProtectionParams

`func NewNetappObjectProtectionParams() *NetappObjectProtectionParams`

NewNetappObjectProtectionParams instantiates a new NetappObjectProtectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetappObjectProtectionParamsWithDefaults

`func NewNetappObjectProtectionParamsWithDefaults() *NetappObjectProtectionParams`

NewNetappObjectProtectionParamsWithDefaults instantiates a new NetappObjectProtectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupExistingSnapshot

`func (o *NetappObjectProtectionParams) GetBackupExistingSnapshot() bool`

GetBackupExistingSnapshot returns the BackupExistingSnapshot field if non-nil, zero value otherwise.

### GetBackupExistingSnapshotOk

`func (o *NetappObjectProtectionParams) GetBackupExistingSnapshotOk() (*bool, bool)`

GetBackupExistingSnapshotOk returns a tuple with the BackupExistingSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupExistingSnapshot

`func (o *NetappObjectProtectionParams) SetBackupExistingSnapshot(v bool)`

SetBackupExistingSnapshot sets BackupExistingSnapshot field to given value.

### HasBackupExistingSnapshot

`func (o *NetappObjectProtectionParams) HasBackupExistingSnapshot() bool`

HasBackupExistingSnapshot returns a boolean if a field has been set.

### SetBackupExistingSnapshotNil

`func (o *NetappObjectProtectionParams) SetBackupExistingSnapshotNil(b bool)`

 SetBackupExistingSnapshotNil sets the value for BackupExistingSnapshot to be an explicit nil

### UnsetBackupExistingSnapshot
`func (o *NetappObjectProtectionParams) UnsetBackupExistingSnapshot()`

UnsetBackupExistingSnapshot ensures that no value is present for BackupExistingSnapshot, not even an explicit nil
### GetContinuousSnapshots

`func (o *NetappObjectProtectionParams) GetContinuousSnapshots() ContinuousSnapshotParams`

GetContinuousSnapshots returns the ContinuousSnapshots field if non-nil, zero value otherwise.

### GetContinuousSnapshotsOk

`func (o *NetappObjectProtectionParams) GetContinuousSnapshotsOk() (*ContinuousSnapshotParams, bool)`

GetContinuousSnapshotsOk returns a tuple with the ContinuousSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuousSnapshots

`func (o *NetappObjectProtectionParams) SetContinuousSnapshots(v ContinuousSnapshotParams)`

SetContinuousSnapshots sets ContinuousSnapshots field to given value.

### HasContinuousSnapshots

`func (o *NetappObjectProtectionParams) HasContinuousSnapshots() bool`

HasContinuousSnapshots returns a boolean if a field has been set.

### GetExcludeObjectIds

`func (o *NetappObjectProtectionParams) GetExcludeObjectIds() []int64`

GetExcludeObjectIds returns the ExcludeObjectIds field if non-nil, zero value otherwise.

### GetExcludeObjectIdsOk

`func (o *NetappObjectProtectionParams) GetExcludeObjectIdsOk() (*[]int64, bool)`

GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeObjectIds

`func (o *NetappObjectProtectionParams) SetExcludeObjectIds(v []int64)`

SetExcludeObjectIds sets ExcludeObjectIds field to given value.

### HasExcludeObjectIds

`func (o *NetappObjectProtectionParams) HasExcludeObjectIds() bool`

HasExcludeObjectIds returns a boolean if a field has been set.

### SetExcludeObjectIdsNil

`func (o *NetappObjectProtectionParams) SetExcludeObjectIdsNil(b bool)`

 SetExcludeObjectIdsNil sets the value for ExcludeObjectIds to be an explicit nil

### UnsetExcludeObjectIds
`func (o *NetappObjectProtectionParams) UnsetExcludeObjectIds()`

UnsetExcludeObjectIds ensures that no value is present for ExcludeObjectIds, not even an explicit nil
### GetProtocol

`func (o *NetappObjectProtectionParams) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *NetappObjectProtectionParams) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *NetappObjectProtectionParams) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *NetappObjectProtectionParams) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *NetappObjectProtectionParams) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *NetappObjectProtectionParams) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetSnapshotLabel

`func (o *NetappObjectProtectionParams) GetSnapshotLabel() SnapshotLabel`

GetSnapshotLabel returns the SnapshotLabel field if non-nil, zero value otherwise.

### GetSnapshotLabelOk

`func (o *NetappObjectProtectionParams) GetSnapshotLabelOk() (*SnapshotLabel, bool)`

GetSnapshotLabelOk returns a tuple with the SnapshotLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotLabel

`func (o *NetappObjectProtectionParams) SetSnapshotLabel(v SnapshotLabel)`

SetSnapshotLabel sets SnapshotLabel field to given value.

### HasSnapshotLabel

`func (o *NetappObjectProtectionParams) HasSnapshotLabel() bool`

HasSnapshotLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


