# RecoverAzureVmNewSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the id of the parent source to recover the VMs. | 
**ResourceGroup** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the Azure resource group. | 
**StorageAccount** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the storage account that will contain the storage container | 
**StorageContainer** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the storage container within the above storage account. | 
**NetworkConfig** | [**NullableRecoverAzureVmNewSourceNetworkConfig**](RecoverAzureVmNewSourceNetworkConfig.md) | Specifies the networking configuration to be applied to the recovered VMs. | 
**StorageResourceGroup** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies id of the resource group for the selected storage account. | 
**ComputeOption** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the type of VM (e.g. small, medium, large) when cloning/restoring the VM in Azure. | 
**AvailabilitySet** | Pointer to [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the availability set. | [optional] 

## Methods

### NewRecoverAzureVmNewSourceConfig

`func NewRecoverAzureVmNewSourceConfig(source NullableRecoveryObjectIdentifier, resourceGroup NullableRecoveryObjectIdentifier, storageAccount NullableRecoveryObjectIdentifier, storageContainer NullableRecoveryObjectIdentifier, networkConfig NullableRecoverAzureVmNewSourceNetworkConfig, storageResourceGroup NullableRecoveryObjectIdentifier, computeOption NullableRecoveryObjectIdentifier, ) *RecoverAzureVmNewSourceConfig`

NewRecoverAzureVmNewSourceConfig instantiates a new RecoverAzureVmNewSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverAzureVmNewSourceConfigWithDefaults

`func NewRecoverAzureVmNewSourceConfigWithDefaults() *RecoverAzureVmNewSourceConfig`

NewRecoverAzureVmNewSourceConfigWithDefaults instantiates a new RecoverAzureVmNewSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *RecoverAzureVmNewSourceConfig) GetSource() RecoveryObjectIdentifier`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RecoverAzureVmNewSourceConfig) GetSourceOk() (*RecoveryObjectIdentifier, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RecoverAzureVmNewSourceConfig) SetSource(v RecoveryObjectIdentifier)`

SetSource sets Source field to given value.


### SetSourceNil

`func (o *RecoverAzureVmNewSourceConfig) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *RecoverAzureVmNewSourceConfig) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetResourceGroup

`func (o *RecoverAzureVmNewSourceConfig) GetResourceGroup() RecoveryObjectIdentifier`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *RecoverAzureVmNewSourceConfig) GetResourceGroupOk() (*RecoveryObjectIdentifier, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *RecoverAzureVmNewSourceConfig) SetResourceGroup(v RecoveryObjectIdentifier)`

SetResourceGroup sets ResourceGroup field to given value.


### SetResourceGroupNil

`func (o *RecoverAzureVmNewSourceConfig) SetResourceGroupNil(b bool)`

 SetResourceGroupNil sets the value for ResourceGroup to be an explicit nil

### UnsetResourceGroup
`func (o *RecoverAzureVmNewSourceConfig) UnsetResourceGroup()`

UnsetResourceGroup ensures that no value is present for ResourceGroup, not even an explicit nil
### GetStorageAccount

`func (o *RecoverAzureVmNewSourceConfig) GetStorageAccount() RecoveryObjectIdentifier`

GetStorageAccount returns the StorageAccount field if non-nil, zero value otherwise.

### GetStorageAccountOk

`func (o *RecoverAzureVmNewSourceConfig) GetStorageAccountOk() (*RecoveryObjectIdentifier, bool)`

GetStorageAccountOk returns a tuple with the StorageAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccount

`func (o *RecoverAzureVmNewSourceConfig) SetStorageAccount(v RecoveryObjectIdentifier)`

SetStorageAccount sets StorageAccount field to given value.


### SetStorageAccountNil

`func (o *RecoverAzureVmNewSourceConfig) SetStorageAccountNil(b bool)`

 SetStorageAccountNil sets the value for StorageAccount to be an explicit nil

### UnsetStorageAccount
`func (o *RecoverAzureVmNewSourceConfig) UnsetStorageAccount()`

UnsetStorageAccount ensures that no value is present for StorageAccount, not even an explicit nil
### GetStorageContainer

`func (o *RecoverAzureVmNewSourceConfig) GetStorageContainer() RecoveryObjectIdentifier`

GetStorageContainer returns the StorageContainer field if non-nil, zero value otherwise.

### GetStorageContainerOk

`func (o *RecoverAzureVmNewSourceConfig) GetStorageContainerOk() (*RecoveryObjectIdentifier, bool)`

GetStorageContainerOk returns a tuple with the StorageContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageContainer

`func (o *RecoverAzureVmNewSourceConfig) SetStorageContainer(v RecoveryObjectIdentifier)`

SetStorageContainer sets StorageContainer field to given value.


### SetStorageContainerNil

`func (o *RecoverAzureVmNewSourceConfig) SetStorageContainerNil(b bool)`

 SetStorageContainerNil sets the value for StorageContainer to be an explicit nil

### UnsetStorageContainer
`func (o *RecoverAzureVmNewSourceConfig) UnsetStorageContainer()`

UnsetStorageContainer ensures that no value is present for StorageContainer, not even an explicit nil
### GetNetworkConfig

`func (o *RecoverAzureVmNewSourceConfig) GetNetworkConfig() RecoverAzureVmNewSourceNetworkConfig`

GetNetworkConfig returns the NetworkConfig field if non-nil, zero value otherwise.

### GetNetworkConfigOk

`func (o *RecoverAzureVmNewSourceConfig) GetNetworkConfigOk() (*RecoverAzureVmNewSourceNetworkConfig, bool)`

GetNetworkConfigOk returns a tuple with the NetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfig

`func (o *RecoverAzureVmNewSourceConfig) SetNetworkConfig(v RecoverAzureVmNewSourceNetworkConfig)`

SetNetworkConfig sets NetworkConfig field to given value.


### SetNetworkConfigNil

`func (o *RecoverAzureVmNewSourceConfig) SetNetworkConfigNil(b bool)`

 SetNetworkConfigNil sets the value for NetworkConfig to be an explicit nil

### UnsetNetworkConfig
`func (o *RecoverAzureVmNewSourceConfig) UnsetNetworkConfig()`

UnsetNetworkConfig ensures that no value is present for NetworkConfig, not even an explicit nil
### GetStorageResourceGroup

`func (o *RecoverAzureVmNewSourceConfig) GetStorageResourceGroup() RecoveryObjectIdentifier`

GetStorageResourceGroup returns the StorageResourceGroup field if non-nil, zero value otherwise.

### GetStorageResourceGroupOk

`func (o *RecoverAzureVmNewSourceConfig) GetStorageResourceGroupOk() (*RecoveryObjectIdentifier, bool)`

GetStorageResourceGroupOk returns a tuple with the StorageResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageResourceGroup

`func (o *RecoverAzureVmNewSourceConfig) SetStorageResourceGroup(v RecoveryObjectIdentifier)`

SetStorageResourceGroup sets StorageResourceGroup field to given value.


### SetStorageResourceGroupNil

`func (o *RecoverAzureVmNewSourceConfig) SetStorageResourceGroupNil(b bool)`

 SetStorageResourceGroupNil sets the value for StorageResourceGroup to be an explicit nil

### UnsetStorageResourceGroup
`func (o *RecoverAzureVmNewSourceConfig) UnsetStorageResourceGroup()`

UnsetStorageResourceGroup ensures that no value is present for StorageResourceGroup, not even an explicit nil
### GetComputeOption

`func (o *RecoverAzureVmNewSourceConfig) GetComputeOption() RecoveryObjectIdentifier`

GetComputeOption returns the ComputeOption field if non-nil, zero value otherwise.

### GetComputeOptionOk

`func (o *RecoverAzureVmNewSourceConfig) GetComputeOptionOk() (*RecoveryObjectIdentifier, bool)`

GetComputeOptionOk returns a tuple with the ComputeOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOption

`func (o *RecoverAzureVmNewSourceConfig) SetComputeOption(v RecoveryObjectIdentifier)`

SetComputeOption sets ComputeOption field to given value.


### SetComputeOptionNil

`func (o *RecoverAzureVmNewSourceConfig) SetComputeOptionNil(b bool)`

 SetComputeOptionNil sets the value for ComputeOption to be an explicit nil

### UnsetComputeOption
`func (o *RecoverAzureVmNewSourceConfig) UnsetComputeOption()`

UnsetComputeOption ensures that no value is present for ComputeOption, not even an explicit nil
### GetAvailabilitySet

`func (o *RecoverAzureVmNewSourceConfig) GetAvailabilitySet() RecoveryObjectIdentifier`

GetAvailabilitySet returns the AvailabilitySet field if non-nil, zero value otherwise.

### GetAvailabilitySetOk

`func (o *RecoverAzureVmNewSourceConfig) GetAvailabilitySetOk() (*RecoveryObjectIdentifier, bool)`

GetAvailabilitySetOk returns a tuple with the AvailabilitySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilitySet

`func (o *RecoverAzureVmNewSourceConfig) SetAvailabilitySet(v RecoveryObjectIdentifier)`

SetAvailabilitySet sets AvailabilitySet field to given value.

### HasAvailabilitySet

`func (o *RecoverAzureVmNewSourceConfig) HasAvailabilitySet() bool`

HasAvailabilitySet returns a boolean if a field has been set.

### SetAvailabilitySetNil

`func (o *RecoverAzureVmNewSourceConfig) SetAvailabilitySetNil(b bool)`

 SetAvailabilitySetNil sets the value for AvailabilitySet to be an explicit nil

### UnsetAvailabilitySet
`func (o *RecoverAzureVmNewSourceConfig) UnsetAvailabilitySet()`

UnsetAvailabilitySet ensures that no value is present for AvailabilitySet, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


