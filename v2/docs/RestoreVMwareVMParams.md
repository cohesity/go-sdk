# RestoreVMwareVMParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttemptDifferentialRestore** | Pointer to **NullableBool** | Specifies whether to attempt differential restore. | [optional] 
**DatastoreIds** | Pointer to **[]int64** | Specifies Datastore Ids, if the restore is to alternate location. | [optional] 
**EnableCopyRecovery** | Pointer to **NullableBool** | Specifies whether to perform copy recovery or not. | [optional] 
**IsOnPremDeploy** | Pointer to **NullableBool** | Specifies whether a task in on prem deploy or not. | [optional] 
**OverwriteExistingVm** | Pointer to **NullableBool** | Specifies whether to overwrite the VM at the target location. | [optional] 
**PowerOffAndRenameExistingVm** | Pointer to **NullableBool** | Specifies whether to power off and mark the VM at the target location as deprecated. | [optional] 
**ResourcePoolId** | Pointer to **NullableInt64** | Specifies if the restore is to alternate location. | [optional] 
**TargetDataStoreId** | Pointer to **NullableInt64** | Specifies the folder where the restore datastore should be created. | [optional] 
**TargetVMFolderId** | Pointer to **NullableInt64** | Specifies the folder ID where the VMs should be created. | [optional] 

## Methods

### NewRestoreVMwareVMParams

`func NewRestoreVMwareVMParams() *RestoreVMwareVMParams`

NewRestoreVMwareVMParams instantiates a new RestoreVMwareVMParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreVMwareVMParamsWithDefaults

`func NewRestoreVMwareVMParamsWithDefaults() *RestoreVMwareVMParams`

NewRestoreVMwareVMParamsWithDefaults instantiates a new RestoreVMwareVMParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttemptDifferentialRestore

`func (o *RestoreVMwareVMParams) GetAttemptDifferentialRestore() bool`

GetAttemptDifferentialRestore returns the AttemptDifferentialRestore field if non-nil, zero value otherwise.

### GetAttemptDifferentialRestoreOk

`func (o *RestoreVMwareVMParams) GetAttemptDifferentialRestoreOk() (*bool, bool)`

GetAttemptDifferentialRestoreOk returns a tuple with the AttemptDifferentialRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptDifferentialRestore

`func (o *RestoreVMwareVMParams) SetAttemptDifferentialRestore(v bool)`

SetAttemptDifferentialRestore sets AttemptDifferentialRestore field to given value.

### HasAttemptDifferentialRestore

`func (o *RestoreVMwareVMParams) HasAttemptDifferentialRestore() bool`

HasAttemptDifferentialRestore returns a boolean if a field has been set.

### SetAttemptDifferentialRestoreNil

`func (o *RestoreVMwareVMParams) SetAttemptDifferentialRestoreNil(b bool)`

 SetAttemptDifferentialRestoreNil sets the value for AttemptDifferentialRestore to be an explicit nil

### UnsetAttemptDifferentialRestore
`func (o *RestoreVMwareVMParams) UnsetAttemptDifferentialRestore()`

UnsetAttemptDifferentialRestore ensures that no value is present for AttemptDifferentialRestore, not even an explicit nil
### GetDatastoreIds

`func (o *RestoreVMwareVMParams) GetDatastoreIds() []int64`

GetDatastoreIds returns the DatastoreIds field if non-nil, zero value otherwise.

### GetDatastoreIdsOk

`func (o *RestoreVMwareVMParams) GetDatastoreIdsOk() (*[]int64, bool)`

GetDatastoreIdsOk returns a tuple with the DatastoreIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreIds

`func (o *RestoreVMwareVMParams) SetDatastoreIds(v []int64)`

SetDatastoreIds sets DatastoreIds field to given value.

### HasDatastoreIds

`func (o *RestoreVMwareVMParams) HasDatastoreIds() bool`

HasDatastoreIds returns a boolean if a field has been set.

### SetDatastoreIdsNil

`func (o *RestoreVMwareVMParams) SetDatastoreIdsNil(b bool)`

 SetDatastoreIdsNil sets the value for DatastoreIds to be an explicit nil

### UnsetDatastoreIds
`func (o *RestoreVMwareVMParams) UnsetDatastoreIds()`

UnsetDatastoreIds ensures that no value is present for DatastoreIds, not even an explicit nil
### GetEnableCopyRecovery

`func (o *RestoreVMwareVMParams) GetEnableCopyRecovery() bool`

GetEnableCopyRecovery returns the EnableCopyRecovery field if non-nil, zero value otherwise.

### GetEnableCopyRecoveryOk

`func (o *RestoreVMwareVMParams) GetEnableCopyRecoveryOk() (*bool, bool)`

GetEnableCopyRecoveryOk returns a tuple with the EnableCopyRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCopyRecovery

`func (o *RestoreVMwareVMParams) SetEnableCopyRecovery(v bool)`

SetEnableCopyRecovery sets EnableCopyRecovery field to given value.

### HasEnableCopyRecovery

`func (o *RestoreVMwareVMParams) HasEnableCopyRecovery() bool`

HasEnableCopyRecovery returns a boolean if a field has been set.

### SetEnableCopyRecoveryNil

`func (o *RestoreVMwareVMParams) SetEnableCopyRecoveryNil(b bool)`

 SetEnableCopyRecoveryNil sets the value for EnableCopyRecovery to be an explicit nil

### UnsetEnableCopyRecovery
`func (o *RestoreVMwareVMParams) UnsetEnableCopyRecovery()`

UnsetEnableCopyRecovery ensures that no value is present for EnableCopyRecovery, not even an explicit nil
### GetIsOnPremDeploy

`func (o *RestoreVMwareVMParams) GetIsOnPremDeploy() bool`

GetIsOnPremDeploy returns the IsOnPremDeploy field if non-nil, zero value otherwise.

### GetIsOnPremDeployOk

`func (o *RestoreVMwareVMParams) GetIsOnPremDeployOk() (*bool, bool)`

GetIsOnPremDeployOk returns a tuple with the IsOnPremDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnPremDeploy

`func (o *RestoreVMwareVMParams) SetIsOnPremDeploy(v bool)`

SetIsOnPremDeploy sets IsOnPremDeploy field to given value.

### HasIsOnPremDeploy

`func (o *RestoreVMwareVMParams) HasIsOnPremDeploy() bool`

HasIsOnPremDeploy returns a boolean if a field has been set.

### SetIsOnPremDeployNil

`func (o *RestoreVMwareVMParams) SetIsOnPremDeployNil(b bool)`

 SetIsOnPremDeployNil sets the value for IsOnPremDeploy to be an explicit nil

### UnsetIsOnPremDeploy
`func (o *RestoreVMwareVMParams) UnsetIsOnPremDeploy()`

UnsetIsOnPremDeploy ensures that no value is present for IsOnPremDeploy, not even an explicit nil
### GetOverwriteExistingVm

`func (o *RestoreVMwareVMParams) GetOverwriteExistingVm() bool`

GetOverwriteExistingVm returns the OverwriteExistingVm field if non-nil, zero value otherwise.

### GetOverwriteExistingVmOk

`func (o *RestoreVMwareVMParams) GetOverwriteExistingVmOk() (*bool, bool)`

GetOverwriteExistingVmOk returns a tuple with the OverwriteExistingVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteExistingVm

`func (o *RestoreVMwareVMParams) SetOverwriteExistingVm(v bool)`

SetOverwriteExistingVm sets OverwriteExistingVm field to given value.

### HasOverwriteExistingVm

`func (o *RestoreVMwareVMParams) HasOverwriteExistingVm() bool`

HasOverwriteExistingVm returns a boolean if a field has been set.

### SetOverwriteExistingVmNil

`func (o *RestoreVMwareVMParams) SetOverwriteExistingVmNil(b bool)`

 SetOverwriteExistingVmNil sets the value for OverwriteExistingVm to be an explicit nil

### UnsetOverwriteExistingVm
`func (o *RestoreVMwareVMParams) UnsetOverwriteExistingVm()`

UnsetOverwriteExistingVm ensures that no value is present for OverwriteExistingVm, not even an explicit nil
### GetPowerOffAndRenameExistingVm

`func (o *RestoreVMwareVMParams) GetPowerOffAndRenameExistingVm() bool`

GetPowerOffAndRenameExistingVm returns the PowerOffAndRenameExistingVm field if non-nil, zero value otherwise.

### GetPowerOffAndRenameExistingVmOk

`func (o *RestoreVMwareVMParams) GetPowerOffAndRenameExistingVmOk() (*bool, bool)`

GetPowerOffAndRenameExistingVmOk returns a tuple with the PowerOffAndRenameExistingVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOffAndRenameExistingVm

`func (o *RestoreVMwareVMParams) SetPowerOffAndRenameExistingVm(v bool)`

SetPowerOffAndRenameExistingVm sets PowerOffAndRenameExistingVm field to given value.

### HasPowerOffAndRenameExistingVm

`func (o *RestoreVMwareVMParams) HasPowerOffAndRenameExistingVm() bool`

HasPowerOffAndRenameExistingVm returns a boolean if a field has been set.

### SetPowerOffAndRenameExistingVmNil

`func (o *RestoreVMwareVMParams) SetPowerOffAndRenameExistingVmNil(b bool)`

 SetPowerOffAndRenameExistingVmNil sets the value for PowerOffAndRenameExistingVm to be an explicit nil

### UnsetPowerOffAndRenameExistingVm
`func (o *RestoreVMwareVMParams) UnsetPowerOffAndRenameExistingVm()`

UnsetPowerOffAndRenameExistingVm ensures that no value is present for PowerOffAndRenameExistingVm, not even an explicit nil
### GetResourcePoolId

`func (o *RestoreVMwareVMParams) GetResourcePoolId() int64`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *RestoreVMwareVMParams) GetResourcePoolIdOk() (*int64, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *RestoreVMwareVMParams) SetResourcePoolId(v int64)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *RestoreVMwareVMParams) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### SetResourcePoolIdNil

`func (o *RestoreVMwareVMParams) SetResourcePoolIdNil(b bool)`

 SetResourcePoolIdNil sets the value for ResourcePoolId to be an explicit nil

### UnsetResourcePoolId
`func (o *RestoreVMwareVMParams) UnsetResourcePoolId()`

UnsetResourcePoolId ensures that no value is present for ResourcePoolId, not even an explicit nil
### GetTargetDataStoreId

`func (o *RestoreVMwareVMParams) GetTargetDataStoreId() int64`

GetTargetDataStoreId returns the TargetDataStoreId field if non-nil, zero value otherwise.

### GetTargetDataStoreIdOk

`func (o *RestoreVMwareVMParams) GetTargetDataStoreIdOk() (*int64, bool)`

GetTargetDataStoreIdOk returns a tuple with the TargetDataStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDataStoreId

`func (o *RestoreVMwareVMParams) SetTargetDataStoreId(v int64)`

SetTargetDataStoreId sets TargetDataStoreId field to given value.

### HasTargetDataStoreId

`func (o *RestoreVMwareVMParams) HasTargetDataStoreId() bool`

HasTargetDataStoreId returns a boolean if a field has been set.

### SetTargetDataStoreIdNil

`func (o *RestoreVMwareVMParams) SetTargetDataStoreIdNil(b bool)`

 SetTargetDataStoreIdNil sets the value for TargetDataStoreId to be an explicit nil

### UnsetTargetDataStoreId
`func (o *RestoreVMwareVMParams) UnsetTargetDataStoreId()`

UnsetTargetDataStoreId ensures that no value is present for TargetDataStoreId, not even an explicit nil
### GetTargetVMFolderId

`func (o *RestoreVMwareVMParams) GetTargetVMFolderId() int64`

GetTargetVMFolderId returns the TargetVMFolderId field if non-nil, zero value otherwise.

### GetTargetVMFolderIdOk

`func (o *RestoreVMwareVMParams) GetTargetVMFolderIdOk() (*int64, bool)`

GetTargetVMFolderIdOk returns a tuple with the TargetVMFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVMFolderId

`func (o *RestoreVMwareVMParams) SetTargetVMFolderId(v int64)`

SetTargetVMFolderId sets TargetVMFolderId field to given value.

### HasTargetVMFolderId

`func (o *RestoreVMwareVMParams) HasTargetVMFolderId() bool`

HasTargetVMFolderId returns a boolean if a field has been set.

### SetTargetVMFolderIdNil

`func (o *RestoreVMwareVMParams) SetTargetVMFolderIdNil(b bool)`

 SetTargetVMFolderIdNil sets the value for TargetVMFolderId to be an explicit nil

### UnsetTargetVMFolderId
`func (o *RestoreVMwareVMParams) UnsetTargetVMFolderId()`

UnsetTargetVMFolderId ensures that no value is present for TargetVMFolderId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


