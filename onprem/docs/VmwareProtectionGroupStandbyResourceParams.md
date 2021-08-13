# VmwareProtectionGroupStandbyResourceParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoveryPointObjectiveSecs** | Pointer to **NullableInt64** | Specifies the recovery point objective time user expects for this standby resource. | [optional] 
**RenameRestoredObjectParams** | Pointer to [**NullableRecoveredOrClonedVmsRenameConfig**](RecoveredOrClonedVmsRenameConfig.md) | Specifies params to rename the standby resource. | [optional] 
**ParentObjectId** | Pointer to **NullableInt64** | Specifies the object id for parent vCenter source where this standby resource should be created. | [optional] 
**TargetFolderObjectId** | Pointer to **NullableInt64** | Specifies the object id for target vm folder where this standby resource should be created. | [optional] 
**TargetDatastoreFolderObjectId** | Pointer to **NullableInt64** | Specifies the object id for target datastore folder where disks for this standby resource should be placed. | [optional] 
**ResourcePoolObjectId** | Pointer to **NullableInt64** | Specifies the object id for resource pool where this standby resource should be created. | [optional] 
**DatastoreObjectIds** | Pointer to **[]int64** | Specifies the list of IDs of the datastore objects where this standby resource should be created. | [optional] 
**NetworkConfig** | Pointer to [**NullableRecoverVmwareVmNewSourceNetworkConfig**](RecoverVmwareVmNewSourceNetworkConfig.md) | Specifies the networking configuration to be applied to this standby resource. | [optional] 

## Methods

### NewVmwareProtectionGroupStandbyResourceParams

`func NewVmwareProtectionGroupStandbyResourceParams() *VmwareProtectionGroupStandbyResourceParams`

NewVmwareProtectionGroupStandbyResourceParams instantiates a new VmwareProtectionGroupStandbyResourceParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareProtectionGroupStandbyResourceParamsWithDefaults

`func NewVmwareProtectionGroupStandbyResourceParamsWithDefaults() *VmwareProtectionGroupStandbyResourceParams`

NewVmwareProtectionGroupStandbyResourceParamsWithDefaults instantiates a new VmwareProtectionGroupStandbyResourceParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecoveryPointObjectiveSecs

`func (o *VmwareProtectionGroupStandbyResourceParams) GetRecoveryPointObjectiveSecs() int64`

GetRecoveryPointObjectiveSecs returns the RecoveryPointObjectiveSecs field if non-nil, zero value otherwise.

### GetRecoveryPointObjectiveSecsOk

`func (o *VmwareProtectionGroupStandbyResourceParams) GetRecoveryPointObjectiveSecsOk() (*int64, bool)`

GetRecoveryPointObjectiveSecsOk returns a tuple with the RecoveryPointObjectiveSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryPointObjectiveSecs

`func (o *VmwareProtectionGroupStandbyResourceParams) SetRecoveryPointObjectiveSecs(v int64)`

SetRecoveryPointObjectiveSecs sets RecoveryPointObjectiveSecs field to given value.

### HasRecoveryPointObjectiveSecs

`func (o *VmwareProtectionGroupStandbyResourceParams) HasRecoveryPointObjectiveSecs() bool`

HasRecoveryPointObjectiveSecs returns a boolean if a field has been set.

### SetRecoveryPointObjectiveSecsNil

`func (o *VmwareProtectionGroupStandbyResourceParams) SetRecoveryPointObjectiveSecsNil(b bool)`

 SetRecoveryPointObjectiveSecsNil sets the value for RecoveryPointObjectiveSecs to be an explicit nil

### UnsetRecoveryPointObjectiveSecs
`func (o *VmwareProtectionGroupStandbyResourceParams) UnsetRecoveryPointObjectiveSecs()`

UnsetRecoveryPointObjectiveSecs ensures that no value is present for RecoveryPointObjectiveSecs, not even an explicit nil
### GetRenameRestoredObjectParams

`func (o *VmwareProtectionGroupStandbyResourceParams) GetRenameRestoredObjectParams() RecoveredOrClonedVmsRenameConfig`

GetRenameRestoredObjectParams returns the RenameRestoredObjectParams field if non-nil, zero value otherwise.

### GetRenameRestoredObjectParamsOk

`func (o *VmwareProtectionGroupStandbyResourceParams) GetRenameRestoredObjectParamsOk() (*RecoveredOrClonedVmsRenameConfig, bool)`

GetRenameRestoredObjectParamsOk returns a tuple with the RenameRestoredObjectParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameRestoredObjectParams

`func (o *VmwareProtectionGroupStandbyResourceParams) SetRenameRestoredObjectParams(v RecoveredOrClonedVmsRenameConfig)`

SetRenameRestoredObjectParams sets RenameRestoredObjectParams field to given value.

### HasRenameRestoredObjectParams

`func (o *VmwareProtectionGroupStandbyResourceParams) HasRenameRestoredObjectParams() bool`

HasRenameRestoredObjectParams returns a boolean if a field has been set.

### SetRenameRestoredObjectParamsNil

`func (o *VmwareProtectionGroupStandbyResourceParams) SetRenameRestoredObjectParamsNil(b bool)`

 SetRenameRestoredObjectParamsNil sets the value for RenameRestoredObjectParams to be an explicit nil

### UnsetRenameRestoredObjectParams
`func (o *VmwareProtectionGroupStandbyResourceParams) UnsetRenameRestoredObjectParams()`

UnsetRenameRestoredObjectParams ensures that no value is present for RenameRestoredObjectParams, not even an explicit nil
### GetParentObjectId

`func (o *VmwareProtectionGroupStandbyResourceParams) GetParentObjectId() int64`

GetParentObjectId returns the ParentObjectId field if non-nil, zero value otherwise.

### GetParentObjectIdOk

`func (o *VmwareProtectionGroupStandbyResourceParams) GetParentObjectIdOk() (*int64, bool)`

GetParentObjectIdOk returns a tuple with the ParentObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentObjectId

`func (o *VmwareProtectionGroupStandbyResourceParams) SetParentObjectId(v int64)`

SetParentObjectId sets ParentObjectId field to given value.

### HasParentObjectId

`func (o *VmwareProtectionGroupStandbyResourceParams) HasParentObjectId() bool`

HasParentObjectId returns a boolean if a field has been set.

### SetParentObjectIdNil

`func (o *VmwareProtectionGroupStandbyResourceParams) SetParentObjectIdNil(b bool)`

 SetParentObjectIdNil sets the value for ParentObjectId to be an explicit nil

### UnsetParentObjectId
`func (o *VmwareProtectionGroupStandbyResourceParams) UnsetParentObjectId()`

UnsetParentObjectId ensures that no value is present for ParentObjectId, not even an explicit nil
### GetTargetFolderObjectId

`func (o *VmwareProtectionGroupStandbyResourceParams) GetTargetFolderObjectId() int64`

GetTargetFolderObjectId returns the TargetFolderObjectId field if non-nil, zero value otherwise.

### GetTargetFolderObjectIdOk

`func (o *VmwareProtectionGroupStandbyResourceParams) GetTargetFolderObjectIdOk() (*int64, bool)`

GetTargetFolderObjectIdOk returns a tuple with the TargetFolderObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFolderObjectId

`func (o *VmwareProtectionGroupStandbyResourceParams) SetTargetFolderObjectId(v int64)`

SetTargetFolderObjectId sets TargetFolderObjectId field to given value.

### HasTargetFolderObjectId

`func (o *VmwareProtectionGroupStandbyResourceParams) HasTargetFolderObjectId() bool`

HasTargetFolderObjectId returns a boolean if a field has been set.

### SetTargetFolderObjectIdNil

`func (o *VmwareProtectionGroupStandbyResourceParams) SetTargetFolderObjectIdNil(b bool)`

 SetTargetFolderObjectIdNil sets the value for TargetFolderObjectId to be an explicit nil

### UnsetTargetFolderObjectId
`func (o *VmwareProtectionGroupStandbyResourceParams) UnsetTargetFolderObjectId()`

UnsetTargetFolderObjectId ensures that no value is present for TargetFolderObjectId, not even an explicit nil
### GetTargetDatastoreFolderObjectId

`func (o *VmwareProtectionGroupStandbyResourceParams) GetTargetDatastoreFolderObjectId() int64`

GetTargetDatastoreFolderObjectId returns the TargetDatastoreFolderObjectId field if non-nil, zero value otherwise.

### GetTargetDatastoreFolderObjectIdOk

`func (o *VmwareProtectionGroupStandbyResourceParams) GetTargetDatastoreFolderObjectIdOk() (*int64, bool)`

GetTargetDatastoreFolderObjectIdOk returns a tuple with the TargetDatastoreFolderObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDatastoreFolderObjectId

`func (o *VmwareProtectionGroupStandbyResourceParams) SetTargetDatastoreFolderObjectId(v int64)`

SetTargetDatastoreFolderObjectId sets TargetDatastoreFolderObjectId field to given value.

### HasTargetDatastoreFolderObjectId

`func (o *VmwareProtectionGroupStandbyResourceParams) HasTargetDatastoreFolderObjectId() bool`

HasTargetDatastoreFolderObjectId returns a boolean if a field has been set.

### SetTargetDatastoreFolderObjectIdNil

`func (o *VmwareProtectionGroupStandbyResourceParams) SetTargetDatastoreFolderObjectIdNil(b bool)`

 SetTargetDatastoreFolderObjectIdNil sets the value for TargetDatastoreFolderObjectId to be an explicit nil

### UnsetTargetDatastoreFolderObjectId
`func (o *VmwareProtectionGroupStandbyResourceParams) UnsetTargetDatastoreFolderObjectId()`

UnsetTargetDatastoreFolderObjectId ensures that no value is present for TargetDatastoreFolderObjectId, not even an explicit nil
### GetResourcePoolObjectId

`func (o *VmwareProtectionGroupStandbyResourceParams) GetResourcePoolObjectId() int64`

GetResourcePoolObjectId returns the ResourcePoolObjectId field if non-nil, zero value otherwise.

### GetResourcePoolObjectIdOk

`func (o *VmwareProtectionGroupStandbyResourceParams) GetResourcePoolObjectIdOk() (*int64, bool)`

GetResourcePoolObjectIdOk returns a tuple with the ResourcePoolObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolObjectId

`func (o *VmwareProtectionGroupStandbyResourceParams) SetResourcePoolObjectId(v int64)`

SetResourcePoolObjectId sets ResourcePoolObjectId field to given value.

### HasResourcePoolObjectId

`func (o *VmwareProtectionGroupStandbyResourceParams) HasResourcePoolObjectId() bool`

HasResourcePoolObjectId returns a boolean if a field has been set.

### SetResourcePoolObjectIdNil

`func (o *VmwareProtectionGroupStandbyResourceParams) SetResourcePoolObjectIdNil(b bool)`

 SetResourcePoolObjectIdNil sets the value for ResourcePoolObjectId to be an explicit nil

### UnsetResourcePoolObjectId
`func (o *VmwareProtectionGroupStandbyResourceParams) UnsetResourcePoolObjectId()`

UnsetResourcePoolObjectId ensures that no value is present for ResourcePoolObjectId, not even an explicit nil
### GetDatastoreObjectIds

`func (o *VmwareProtectionGroupStandbyResourceParams) GetDatastoreObjectIds() []int64`

GetDatastoreObjectIds returns the DatastoreObjectIds field if non-nil, zero value otherwise.

### GetDatastoreObjectIdsOk

`func (o *VmwareProtectionGroupStandbyResourceParams) GetDatastoreObjectIdsOk() (*[]int64, bool)`

GetDatastoreObjectIdsOk returns a tuple with the DatastoreObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreObjectIds

`func (o *VmwareProtectionGroupStandbyResourceParams) SetDatastoreObjectIds(v []int64)`

SetDatastoreObjectIds sets DatastoreObjectIds field to given value.

### HasDatastoreObjectIds

`func (o *VmwareProtectionGroupStandbyResourceParams) HasDatastoreObjectIds() bool`

HasDatastoreObjectIds returns a boolean if a field has been set.

### GetNetworkConfig

`func (o *VmwareProtectionGroupStandbyResourceParams) GetNetworkConfig() RecoverVmwareVmNewSourceNetworkConfig`

GetNetworkConfig returns the NetworkConfig field if non-nil, zero value otherwise.

### GetNetworkConfigOk

`func (o *VmwareProtectionGroupStandbyResourceParams) GetNetworkConfigOk() (*RecoverVmwareVmNewSourceNetworkConfig, bool)`

GetNetworkConfigOk returns a tuple with the NetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfig

`func (o *VmwareProtectionGroupStandbyResourceParams) SetNetworkConfig(v RecoverVmwareVmNewSourceNetworkConfig)`

SetNetworkConfig sets NetworkConfig field to given value.

### HasNetworkConfig

`func (o *VmwareProtectionGroupStandbyResourceParams) HasNetworkConfig() bool`

HasNetworkConfig returns a boolean if a field has been set.

### SetNetworkConfigNil

`func (o *VmwareProtectionGroupStandbyResourceParams) SetNetworkConfigNil(b bool)`

 SetNetworkConfigNil sets the value for NetworkConfig to be an explicit nil

### UnsetNetworkConfig
`func (o *VmwareProtectionGroupStandbyResourceParams) UnsetNetworkConfig()`

UnsetNetworkConfig ensures that no value is present for NetworkConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


