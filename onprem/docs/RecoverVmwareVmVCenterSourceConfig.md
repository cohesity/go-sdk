# RecoverVmwareVmVCenterSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the id of the parent source to recover the VMs. | 
**ResourcePool** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the resource pool object where the recovered objects will be attached. | 
**Datastores** | [**[]RecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the datastore objects where the object&#39;s files should be recovered to. | 
**VmFolder** | Pointer to [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Folder where the VMs should be created. | [optional] 
**NetworkConfig** | Pointer to [**NullableRecoverVmwareVmNewSourceNetworkConfig**](RecoverVmwareVmNewSourceNetworkConfig.md) | Specifies the networking configuration to be applied to the recovered VMs. | [optional] 

## Methods

### NewRecoverVmwareVmVCenterSourceConfig

`func NewRecoverVmwareVmVCenterSourceConfig(source NullableRecoveryObjectIdentifier, resourcePool NullableRecoveryObjectIdentifier, datastores []RecoveryObjectIdentifier, ) *RecoverVmwareVmVCenterSourceConfig`

NewRecoverVmwareVmVCenterSourceConfig instantiates a new RecoverVmwareVmVCenterSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverVmwareVmVCenterSourceConfigWithDefaults

`func NewRecoverVmwareVmVCenterSourceConfigWithDefaults() *RecoverVmwareVmVCenterSourceConfig`

NewRecoverVmwareVmVCenterSourceConfigWithDefaults instantiates a new RecoverVmwareVmVCenterSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *RecoverVmwareVmVCenterSourceConfig) GetSource() RecoveryObjectIdentifier`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RecoverVmwareVmVCenterSourceConfig) GetSourceOk() (*RecoveryObjectIdentifier, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RecoverVmwareVmVCenterSourceConfig) SetSource(v RecoveryObjectIdentifier)`

SetSource sets Source field to given value.


### SetSourceNil

`func (o *RecoverVmwareVmVCenterSourceConfig) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *RecoverVmwareVmVCenterSourceConfig) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetResourcePool

`func (o *RecoverVmwareVmVCenterSourceConfig) GetResourcePool() RecoveryObjectIdentifier`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *RecoverVmwareVmVCenterSourceConfig) GetResourcePoolOk() (*RecoveryObjectIdentifier, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *RecoverVmwareVmVCenterSourceConfig) SetResourcePool(v RecoveryObjectIdentifier)`

SetResourcePool sets ResourcePool field to given value.


### SetResourcePoolNil

`func (o *RecoverVmwareVmVCenterSourceConfig) SetResourcePoolNil(b bool)`

 SetResourcePoolNil sets the value for ResourcePool to be an explicit nil

### UnsetResourcePool
`func (o *RecoverVmwareVmVCenterSourceConfig) UnsetResourcePool()`

UnsetResourcePool ensures that no value is present for ResourcePool, not even an explicit nil
### GetDatastores

`func (o *RecoverVmwareVmVCenterSourceConfig) GetDatastores() []RecoveryObjectIdentifier`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *RecoverVmwareVmVCenterSourceConfig) GetDatastoresOk() (*[]RecoveryObjectIdentifier, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *RecoverVmwareVmVCenterSourceConfig) SetDatastores(v []RecoveryObjectIdentifier)`

SetDatastores sets Datastores field to given value.


### SetDatastoresNil

`func (o *RecoverVmwareVmVCenterSourceConfig) SetDatastoresNil(b bool)`

 SetDatastoresNil sets the value for Datastores to be an explicit nil

### UnsetDatastores
`func (o *RecoverVmwareVmVCenterSourceConfig) UnsetDatastores()`

UnsetDatastores ensures that no value is present for Datastores, not even an explicit nil
### GetVmFolder

`func (o *RecoverVmwareVmVCenterSourceConfig) GetVmFolder() RecoveryObjectIdentifier`

GetVmFolder returns the VmFolder field if non-nil, zero value otherwise.

### GetVmFolderOk

`func (o *RecoverVmwareVmVCenterSourceConfig) GetVmFolderOk() (*RecoveryObjectIdentifier, bool)`

GetVmFolderOk returns a tuple with the VmFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmFolder

`func (o *RecoverVmwareVmVCenterSourceConfig) SetVmFolder(v RecoveryObjectIdentifier)`

SetVmFolder sets VmFolder field to given value.

### HasVmFolder

`func (o *RecoverVmwareVmVCenterSourceConfig) HasVmFolder() bool`

HasVmFolder returns a boolean if a field has been set.

### SetVmFolderNil

`func (o *RecoverVmwareVmVCenterSourceConfig) SetVmFolderNil(b bool)`

 SetVmFolderNil sets the value for VmFolder to be an explicit nil

### UnsetVmFolder
`func (o *RecoverVmwareVmVCenterSourceConfig) UnsetVmFolder()`

UnsetVmFolder ensures that no value is present for VmFolder, not even an explicit nil
### GetNetworkConfig

`func (o *RecoverVmwareVmVCenterSourceConfig) GetNetworkConfig() RecoverVmwareVmNewSourceNetworkConfig`

GetNetworkConfig returns the NetworkConfig field if non-nil, zero value otherwise.

### GetNetworkConfigOk

`func (o *RecoverVmwareVmVCenterSourceConfig) GetNetworkConfigOk() (*RecoverVmwareVmNewSourceNetworkConfig, bool)`

GetNetworkConfigOk returns a tuple with the NetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfig

`func (o *RecoverVmwareVmVCenterSourceConfig) SetNetworkConfig(v RecoverVmwareVmNewSourceNetworkConfig)`

SetNetworkConfig sets NetworkConfig field to given value.

### HasNetworkConfig

`func (o *RecoverVmwareVmVCenterSourceConfig) HasNetworkConfig() bool`

HasNetworkConfig returns a boolean if a field has been set.

### SetNetworkConfigNil

`func (o *RecoverVmwareVmVCenterSourceConfig) SetNetworkConfigNil(b bool)`

 SetNetworkConfigNil sets the value for NetworkConfig to be an explicit nil

### UnsetNetworkConfig
`func (o *RecoverVmwareVmVCenterSourceConfig) UnsetNetworkConfig()`

UnsetNetworkConfig ensures that no value is present for NetworkConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


