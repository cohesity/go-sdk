# DataAccessSnapshotInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectId** | **int64** | Specifies the entity id of the object. | 
**RestoreTimeUsecs** | Pointer to **NullableInt64** | Specifies the time to which the object needs to be restored. If this is not specified the object is restore from the base snapshot identified by the run_start_time_usecs. | [optional] 
**SnapshotId** | **string** | Specifies the id of the object snapshot. | 
**VaultId** | Pointer to **NullableInt64** | Specifies Id of the Vault from which the object is restored. This field must be set if the object is to be restored/retrieved from an archive. | [optional] 

## Methods

### NewDataAccessSnapshotInfo

`func NewDataAccessSnapshotInfo(objectId int64, snapshotId string, ) *DataAccessSnapshotInfo`

NewDataAccessSnapshotInfo instantiates a new DataAccessSnapshotInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataAccessSnapshotInfoWithDefaults

`func NewDataAccessSnapshotInfoWithDefaults() *DataAccessSnapshotInfo`

NewDataAccessSnapshotInfoWithDefaults instantiates a new DataAccessSnapshotInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectId

`func (o *DataAccessSnapshotInfo) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *DataAccessSnapshotInfo) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *DataAccessSnapshotInfo) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.


### GetRestoreTimeUsecs

`func (o *DataAccessSnapshotInfo) GetRestoreTimeUsecs() int64`

GetRestoreTimeUsecs returns the RestoreTimeUsecs field if non-nil, zero value otherwise.

### GetRestoreTimeUsecsOk

`func (o *DataAccessSnapshotInfo) GetRestoreTimeUsecsOk() (*int64, bool)`

GetRestoreTimeUsecsOk returns a tuple with the RestoreTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreTimeUsecs

`func (o *DataAccessSnapshotInfo) SetRestoreTimeUsecs(v int64)`

SetRestoreTimeUsecs sets RestoreTimeUsecs field to given value.

### HasRestoreTimeUsecs

`func (o *DataAccessSnapshotInfo) HasRestoreTimeUsecs() bool`

HasRestoreTimeUsecs returns a boolean if a field has been set.

### SetRestoreTimeUsecsNil

`func (o *DataAccessSnapshotInfo) SetRestoreTimeUsecsNil(b bool)`

 SetRestoreTimeUsecsNil sets the value for RestoreTimeUsecs to be an explicit nil

### UnsetRestoreTimeUsecs
`func (o *DataAccessSnapshotInfo) UnsetRestoreTimeUsecs()`

UnsetRestoreTimeUsecs ensures that no value is present for RestoreTimeUsecs, not even an explicit nil
### GetSnapshotId

`func (o *DataAccessSnapshotInfo) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *DataAccessSnapshotInfo) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *DataAccessSnapshotInfo) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.


### GetVaultId

`func (o *DataAccessSnapshotInfo) GetVaultId() int64`

GetVaultId returns the VaultId field if non-nil, zero value otherwise.

### GetVaultIdOk

`func (o *DataAccessSnapshotInfo) GetVaultIdOk() (*int64, bool)`

GetVaultIdOk returns a tuple with the VaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultId

`func (o *DataAccessSnapshotInfo) SetVaultId(v int64)`

SetVaultId sets VaultId field to given value.

### HasVaultId

`func (o *DataAccessSnapshotInfo) HasVaultId() bool`

HasVaultId returns a boolean if a field has been set.

### SetVaultIdNil

`func (o *DataAccessSnapshotInfo) SetVaultIdNil(b bool)`

 SetVaultIdNil sets the value for VaultId to be an explicit nil

### UnsetVaultId
`func (o *DataAccessSnapshotInfo) UnsetVaultId()`

UnsetVaultId ensures that no value is present for VaultId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


