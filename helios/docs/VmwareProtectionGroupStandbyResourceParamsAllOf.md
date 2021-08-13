# VmwareProtectionGroupStandbyResourceParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RenameRestoredObjectParams** | Pointer to [**NullableRecoveredOrClonedVmsRenameConfig**](RecoveredOrClonedVmsRenameConfig.md) | Specifies params to rename the standby resource. | [optional] 
**ParentObjectId** | Pointer to **NullableInt64** | Specifies the object id for parent vCenter source where this standby resource should be created. | [optional] 
**TargetFolderObjectId** | Pointer to **NullableInt64** | Specifies the object id for target vm folder where this standby resource should be created. | [optional] 
**TargetDatastoreFolderObjectId** | Pointer to **NullableInt64** | Specifies the object id for target datastore folder where disks for this standby resource should be placed. | [optional] 
**ResourcePoolObjectId** | Pointer to **NullableInt64** | Specifies the object id for resource pool where this standby resource should be created. | [optional] 
**DatastoreObjectIds** | Pointer to **[]int64** | Specifies the list of IDs of the datastore objects where this standby resource should be created. | [optional] 
**NetworkConfig** | Pointer to [**NullableRecoverVmwareVmNewSourceNetworkConfig**](RecoverVmwareVmNewSourceNetworkConfig.md) | Specifies the networking configuration to be applied to this standby resource. | [optional] 

## Methods

### NewVmwareProtectionGroupStandbyResourceParamsAllOf

`func NewVmwareProtectionGroupStandbyResourceParamsAllOf() *VmwareProtectionGroupStandbyResourceParamsAllOf`

NewVmwareProtectionGroupStandbyResourceParamsAllOf instantiates a new VmwareProtectionGroupStandbyResourceParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareProtectionGroupStandbyResourceParamsAllOfWithDefaults

`func NewVmwareProtectionGroupStandbyResourceParamsAllOfWithDefaults() *VmwareProtectionGroupStandbyResourceParamsAllOf`

NewVmwareProtectionGroupStandbyResourceParamsAllOfWithDefaults instantiates a new VmwareProtectionGroupStandbyResourceParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRenameRestoredObjectParams

`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) GetRenameRestoredObjectParams() RecoveredOrClonedVmsRenameConfig`

GetRenameRestoredObjectParams returns the RenameRestoredObjectParams field if non-nil, zero value otherwise.

### GetRenameRestoredObjectParamsOk

`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) GetRenameRestoredObjectParamsOk() (*RecoveredOrClonedVmsRenameConfig, bool)`

GetRenameRestoredObjectParamsOk returns a tuple with the RenameRestoredObjectParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameRestoredObjectParams

`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) SetRenameRestoredObjectParams(v RecoveredOrClonedVmsRenameConfig)`

SetRenameRestoredObjectParams sets RenameRestoredObjectParams field to given value.

### HasRenameRestoredObjectParams

`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) HasRenameRestoredObjectParams() bool`

HasRenameRestoredObjectParams returns a boolean if a field has been set.

### SetRenameRestoredObjectParamsNil

`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) SetRenameRestoredObjectParamsNil(b bool)`

 SetRenameRestoredObjectParamsNil sets the value for RenameRestoredObjectParams to be an explicit nil

### UnsetRenameRestoredObjectParams
`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) UnsetRenameRestoredObjectParams()`

UnsetRenameRestoredObjectParams ensures that no value is present for RenameRestoredObjectParams, not even an explicit nil
### GetParentObjectId

`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) GetParentObjectId() int64`

GetParentObjectId returns the ParentObjectId field if non-nil, zero value otherwise.

### GetParentObjectIdOk

`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) GetParentObjectIdOk() (*int64, bool)`

GetParentObjectIdOk returns a tuple with the ParentObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentObjectId

`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) SetParentObjectId(v int64)`

SetParentObjectId sets ParentObjectId field to given value.

### HasParentObjectId

`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) HasParentObjectId() bool`

HasParentObjectId returns a boolean if a field has been set.

### SetParentObjectIdNil

`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) SetParentObjectIdNil(b bool)`

 SetParentObjectIdNil sets the value for ParentObjectId to be an explicit nil

### UnsetParentObjectId
`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) UnsetParentObjectId()`

UnsetParentObjectId ensures that no value is present for ParentObjectId, not even an explicit nil
### GetTargetFolderObjectId

`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) GetTargetFolderObjectId() int64`

GetTargetFolderObjectId returns the TargetFolderObjectId field if non-nil, zero value otherwise.

### GetTargetFolderObjectIdOk

`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) GetTargetFolderObjectIdOk() (*int64, bool)`

GetTargetFolderObjectIdOk returns a tuple with the TargetFolderObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFolderObjectId

`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) SetTargetFolderObjectId(v int64)`

SetTargetFolderObjectId sets TargetFolderObjectId field to given value.

### HasTargetFolderObjectId

`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) HasTargetFolderObjectId() bool`

HasTargetFolderObjectId returns a boolean if a field has been set.

### SetTargetFolderObjectIdNil

`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) SetTargetFolderObjectIdNil(b bool)`

 SetTargetFolderObjectIdNil sets the value for TargetFolderObjectId to be an explicit nil

### UnsetTargetFolderObjectId
`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) UnsetTargetFolderObjectId()`

UnsetTargetFolderObjectId ensures that no value is present for TargetFolderObjectId, not even an explicit nil
### GetTargetDatastoreFolderObjectId

`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) GetTargetDatastoreFolderObjectId() int64`

GetTargetDatastoreFolderObjectId returns the TargetDatastoreFolderObjectId field if non-nil, zero value otherwise.

### GetTargetDatastoreFolderObjectIdOk

`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) GetTargetDatastoreFolderObjectIdOk() (*int64, bool)`

GetTargetDatastoreFolderObjectIdOk returns a tuple with the TargetDatastoreFolderObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDatastoreFolderObjectId

`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) SetTargetDatastoreFolderObjectId(v int64)`

SetTargetDatastoreFolderObjectId sets TargetDatastoreFolderObjectId field to given value.

### HasTargetDatastoreFolderObjectId

`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) HasTargetDatastoreFolderObjectId() bool`

HasTargetDatastoreFolderObjectId returns a boolean if a field has been set.

### SetTargetDatastoreFolderObjectIdNil

`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) SetTargetDatastoreFolderObjectIdNil(b bool)`

 SetTargetDatastoreFolderObjectIdNil sets the value for TargetDatastoreFolderObjectId to be an explicit nil

### UnsetTargetDatastoreFolderObjectId
`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) UnsetTargetDatastoreFolderObjectId()`

UnsetTargetDatastoreFolderObjectId ensures that no value is present for TargetDatastoreFolderObjectId, not even an explicit nil
### GetResourcePoolObjectId

`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) GetResourcePoolObjectId() int64`

GetResourcePoolObjectId returns the ResourcePoolObjectId field if non-nil, zero value otherwise.

### GetResourcePoolObjectIdOk

`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) GetResourcePoolObjectIdOk() (*int64, bool)`

GetResourcePoolObjectIdOk returns a tuple with the ResourcePoolObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolObjectId

`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) SetResourcePoolObjectId(v int64)`

SetResourcePoolObjectId sets ResourcePoolObjectId field to given value.

### HasResourcePoolObjectId

`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) HasResourcePoolObjectId() bool`

HasResourcePoolObjectId returns a boolean if a field has been set.

### SetResourcePoolObjectIdNil

`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) SetResourcePoolObjectIdNil(b bool)`

 SetResourcePoolObjectIdNil sets the value for ResourcePoolObjectId to be an explicit nil

### UnsetResourcePoolObjectId
`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) UnsetResourcePoolObjectId()`

UnsetResourcePoolObjectId ensures that no value is present for ResourcePoolObjectId, not even an explicit nil
### GetDatastoreObjectIds

`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) GetDatastoreObjectIds() []int64`

GetDatastoreObjectIds returns the DatastoreObjectIds field if non-nil, zero value otherwise.

### GetDatastoreObjectIdsOk

`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) GetDatastoreObjectIdsOk() (*[]int64, bool)`

GetDatastoreObjectIdsOk returns a tuple with the DatastoreObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreObjectIds

`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) SetDatastoreObjectIds(v []int64)`

SetDatastoreObjectIds sets DatastoreObjectIds field to given value.

### HasDatastoreObjectIds

`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) HasDatastoreObjectIds() bool`

HasDatastoreObjectIds returns a boolean if a field has been set.

### GetNetworkConfig

`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) GetNetworkConfig() RecoverVmwareVmNewSourceNetworkConfig`

GetNetworkConfig returns the NetworkConfig field if non-nil, zero value otherwise.

### GetNetworkConfigOk

`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) GetNetworkConfigOk() (*RecoverVmwareVmNewSourceNetworkConfig, bool)`

GetNetworkConfigOk returns a tuple with the NetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfig

`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) SetNetworkConfig(v RecoverVmwareVmNewSourceNetworkConfig)`

SetNetworkConfig sets NetworkConfig field to given value.

### HasNetworkConfig

`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) HasNetworkConfig() bool`

HasNetworkConfig returns a boolean if a field has been set.

### SetNetworkConfigNil

`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) SetNetworkConfigNil(b bool)`

 SetNetworkConfigNil sets the value for NetworkConfig to be an explicit nil

### UnsetNetworkConfig
`func (o *VmwareProtectionGroupStandbyResourceParamsAllOf) UnsetNetworkConfig()`

UnsetNetworkConfig ensures that no value is present for NetworkConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


