# RecoverAzureVmNewSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilitySet** | Pointer to [**NullableRecoverAzureVmNewSourceConfigAvailabilitySet**](RecoverAzureVmNewSourceConfigAvailabilitySet.md) |  | [optional] 
**ComputeOption** | Pointer to [**NullableRecoverAzureVmNewSourceConfigComputeOption**](RecoverAzureVmNewSourceConfigComputeOption.md) |  | [optional] 
**DataTransferInfo** | Pointer to [**DataTransferInfo**](DataTransferInfo.md) |  | [optional] 
**NetworkConfig** | [**NullableRecoverAzureVmNewSourceConfigNetworkConfig**](RecoverAzureVmNewSourceConfigNetworkConfig.md) |  | 
**Region** | Pointer to [**NullableRecoverAzureVmNewSourceConfigRegion**](RecoverAzureVmNewSourceConfigRegion.md) |  | [optional] 
**ResourceGroup** | [**NullableRecoverAzureVmNewSourceConfigResourceGroup**](RecoverAzureVmNewSourceConfigResourceGroup.md) |  | 
**Source** | [**NullableRecoverAcropolisVmNewSourceConfigSource**](RecoverAcropolisVmNewSourceConfigSource.md) |  | 
**StorageAccount** | Pointer to [**NullableRecoverAzureVmNewSourceConfigStorageAccount**](RecoverAzureVmNewSourceConfigStorageAccount.md) |  | [optional] 
**StorageContainer** | Pointer to [**NullableRecoverAzureVmNewSourceConfigStorageContainer**](RecoverAzureVmNewSourceConfigStorageContainer.md) |  | [optional] 
**StorageResourceGroup** | Pointer to [**NullableRecoverAzureVmNewSourceConfigStorageResourceGroup**](RecoverAzureVmNewSourceConfigStorageResourceGroup.md) |  | [optional] 
**Subscription** | Pointer to [**NullableRecoverAzureVmNewSourceConfigSubscription**](RecoverAzureVmNewSourceConfigSubscription.md) |  | [optional] 

## Methods

### NewRecoverAzureVmNewSourceConfig

`func NewRecoverAzureVmNewSourceConfig(networkConfig NullableRecoverAzureVmNewSourceConfigNetworkConfig, resourceGroup NullableRecoverAzureVmNewSourceConfigResourceGroup, source NullableRecoverAcropolisVmNewSourceConfigSource, ) *RecoverAzureVmNewSourceConfig`

NewRecoverAzureVmNewSourceConfig instantiates a new RecoverAzureVmNewSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverAzureVmNewSourceConfigWithDefaults

`func NewRecoverAzureVmNewSourceConfigWithDefaults() *RecoverAzureVmNewSourceConfig`

NewRecoverAzureVmNewSourceConfigWithDefaults instantiates a new RecoverAzureVmNewSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilitySet

`func (o *RecoverAzureVmNewSourceConfig) GetAvailabilitySet() RecoverAzureVmNewSourceConfigAvailabilitySet`

GetAvailabilitySet returns the AvailabilitySet field if non-nil, zero value otherwise.

### GetAvailabilitySetOk

`func (o *RecoverAzureVmNewSourceConfig) GetAvailabilitySetOk() (*RecoverAzureVmNewSourceConfigAvailabilitySet, bool)`

GetAvailabilitySetOk returns a tuple with the AvailabilitySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilitySet

`func (o *RecoverAzureVmNewSourceConfig) SetAvailabilitySet(v RecoverAzureVmNewSourceConfigAvailabilitySet)`

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
### GetComputeOption

`func (o *RecoverAzureVmNewSourceConfig) GetComputeOption() RecoverAzureVmNewSourceConfigComputeOption`

GetComputeOption returns the ComputeOption field if non-nil, zero value otherwise.

### GetComputeOptionOk

`func (o *RecoverAzureVmNewSourceConfig) GetComputeOptionOk() (*RecoverAzureVmNewSourceConfigComputeOption, bool)`

GetComputeOptionOk returns a tuple with the ComputeOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOption

`func (o *RecoverAzureVmNewSourceConfig) SetComputeOption(v RecoverAzureVmNewSourceConfigComputeOption)`

SetComputeOption sets ComputeOption field to given value.

### HasComputeOption

`func (o *RecoverAzureVmNewSourceConfig) HasComputeOption() bool`

HasComputeOption returns a boolean if a field has been set.

### SetComputeOptionNil

`func (o *RecoverAzureVmNewSourceConfig) SetComputeOptionNil(b bool)`

 SetComputeOptionNil sets the value for ComputeOption to be an explicit nil

### UnsetComputeOption
`func (o *RecoverAzureVmNewSourceConfig) UnsetComputeOption()`

UnsetComputeOption ensures that no value is present for ComputeOption, not even an explicit nil
### GetDataTransferInfo

`func (o *RecoverAzureVmNewSourceConfig) GetDataTransferInfo() DataTransferInfo`

GetDataTransferInfo returns the DataTransferInfo field if non-nil, zero value otherwise.

### GetDataTransferInfoOk

`func (o *RecoverAzureVmNewSourceConfig) GetDataTransferInfoOk() (*DataTransferInfo, bool)`

GetDataTransferInfoOk returns a tuple with the DataTransferInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTransferInfo

`func (o *RecoverAzureVmNewSourceConfig) SetDataTransferInfo(v DataTransferInfo)`

SetDataTransferInfo sets DataTransferInfo field to given value.

### HasDataTransferInfo

`func (o *RecoverAzureVmNewSourceConfig) HasDataTransferInfo() bool`

HasDataTransferInfo returns a boolean if a field has been set.

### GetNetworkConfig

`func (o *RecoverAzureVmNewSourceConfig) GetNetworkConfig() RecoverAzureVmNewSourceConfigNetworkConfig`

GetNetworkConfig returns the NetworkConfig field if non-nil, zero value otherwise.

### GetNetworkConfigOk

`func (o *RecoverAzureVmNewSourceConfig) GetNetworkConfigOk() (*RecoverAzureVmNewSourceConfigNetworkConfig, bool)`

GetNetworkConfigOk returns a tuple with the NetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfig

`func (o *RecoverAzureVmNewSourceConfig) SetNetworkConfig(v RecoverAzureVmNewSourceConfigNetworkConfig)`

SetNetworkConfig sets NetworkConfig field to given value.


### SetNetworkConfigNil

`func (o *RecoverAzureVmNewSourceConfig) SetNetworkConfigNil(b bool)`

 SetNetworkConfigNil sets the value for NetworkConfig to be an explicit nil

### UnsetNetworkConfig
`func (o *RecoverAzureVmNewSourceConfig) UnsetNetworkConfig()`

UnsetNetworkConfig ensures that no value is present for NetworkConfig, not even an explicit nil
### GetRegion

`func (o *RecoverAzureVmNewSourceConfig) GetRegion() RecoverAzureVmNewSourceConfigRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *RecoverAzureVmNewSourceConfig) GetRegionOk() (*RecoverAzureVmNewSourceConfigRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *RecoverAzureVmNewSourceConfig) SetRegion(v RecoverAzureVmNewSourceConfigRegion)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *RecoverAzureVmNewSourceConfig) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *RecoverAzureVmNewSourceConfig) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *RecoverAzureVmNewSourceConfig) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetResourceGroup

`func (o *RecoverAzureVmNewSourceConfig) GetResourceGroup() RecoverAzureVmNewSourceConfigResourceGroup`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *RecoverAzureVmNewSourceConfig) GetResourceGroupOk() (*RecoverAzureVmNewSourceConfigResourceGroup, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *RecoverAzureVmNewSourceConfig) SetResourceGroup(v RecoverAzureVmNewSourceConfigResourceGroup)`

SetResourceGroup sets ResourceGroup field to given value.


### SetResourceGroupNil

`func (o *RecoverAzureVmNewSourceConfig) SetResourceGroupNil(b bool)`

 SetResourceGroupNil sets the value for ResourceGroup to be an explicit nil

### UnsetResourceGroup
`func (o *RecoverAzureVmNewSourceConfig) UnsetResourceGroup()`

UnsetResourceGroup ensures that no value is present for ResourceGroup, not even an explicit nil
### GetSource

`func (o *RecoverAzureVmNewSourceConfig) GetSource() RecoverAcropolisVmNewSourceConfigSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RecoverAzureVmNewSourceConfig) GetSourceOk() (*RecoverAcropolisVmNewSourceConfigSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RecoverAzureVmNewSourceConfig) SetSource(v RecoverAcropolisVmNewSourceConfigSource)`

SetSource sets Source field to given value.


### SetSourceNil

`func (o *RecoverAzureVmNewSourceConfig) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *RecoverAzureVmNewSourceConfig) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetStorageAccount

`func (o *RecoverAzureVmNewSourceConfig) GetStorageAccount() RecoverAzureVmNewSourceConfigStorageAccount`

GetStorageAccount returns the StorageAccount field if non-nil, zero value otherwise.

### GetStorageAccountOk

`func (o *RecoverAzureVmNewSourceConfig) GetStorageAccountOk() (*RecoverAzureVmNewSourceConfigStorageAccount, bool)`

GetStorageAccountOk returns a tuple with the StorageAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccount

`func (o *RecoverAzureVmNewSourceConfig) SetStorageAccount(v RecoverAzureVmNewSourceConfigStorageAccount)`

SetStorageAccount sets StorageAccount field to given value.

### HasStorageAccount

`func (o *RecoverAzureVmNewSourceConfig) HasStorageAccount() bool`

HasStorageAccount returns a boolean if a field has been set.

### SetStorageAccountNil

`func (o *RecoverAzureVmNewSourceConfig) SetStorageAccountNil(b bool)`

 SetStorageAccountNil sets the value for StorageAccount to be an explicit nil

### UnsetStorageAccount
`func (o *RecoverAzureVmNewSourceConfig) UnsetStorageAccount()`

UnsetStorageAccount ensures that no value is present for StorageAccount, not even an explicit nil
### GetStorageContainer

`func (o *RecoverAzureVmNewSourceConfig) GetStorageContainer() RecoverAzureVmNewSourceConfigStorageContainer`

GetStorageContainer returns the StorageContainer field if non-nil, zero value otherwise.

### GetStorageContainerOk

`func (o *RecoverAzureVmNewSourceConfig) GetStorageContainerOk() (*RecoverAzureVmNewSourceConfigStorageContainer, bool)`

GetStorageContainerOk returns a tuple with the StorageContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageContainer

`func (o *RecoverAzureVmNewSourceConfig) SetStorageContainer(v RecoverAzureVmNewSourceConfigStorageContainer)`

SetStorageContainer sets StorageContainer field to given value.

### HasStorageContainer

`func (o *RecoverAzureVmNewSourceConfig) HasStorageContainer() bool`

HasStorageContainer returns a boolean if a field has been set.

### SetStorageContainerNil

`func (o *RecoverAzureVmNewSourceConfig) SetStorageContainerNil(b bool)`

 SetStorageContainerNil sets the value for StorageContainer to be an explicit nil

### UnsetStorageContainer
`func (o *RecoverAzureVmNewSourceConfig) UnsetStorageContainer()`

UnsetStorageContainer ensures that no value is present for StorageContainer, not even an explicit nil
### GetStorageResourceGroup

`func (o *RecoverAzureVmNewSourceConfig) GetStorageResourceGroup() RecoverAzureVmNewSourceConfigStorageResourceGroup`

GetStorageResourceGroup returns the StorageResourceGroup field if non-nil, zero value otherwise.

### GetStorageResourceGroupOk

`func (o *RecoverAzureVmNewSourceConfig) GetStorageResourceGroupOk() (*RecoverAzureVmNewSourceConfigStorageResourceGroup, bool)`

GetStorageResourceGroupOk returns a tuple with the StorageResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageResourceGroup

`func (o *RecoverAzureVmNewSourceConfig) SetStorageResourceGroup(v RecoverAzureVmNewSourceConfigStorageResourceGroup)`

SetStorageResourceGroup sets StorageResourceGroup field to given value.

### HasStorageResourceGroup

`func (o *RecoverAzureVmNewSourceConfig) HasStorageResourceGroup() bool`

HasStorageResourceGroup returns a boolean if a field has been set.

### SetStorageResourceGroupNil

`func (o *RecoverAzureVmNewSourceConfig) SetStorageResourceGroupNil(b bool)`

 SetStorageResourceGroupNil sets the value for StorageResourceGroup to be an explicit nil

### UnsetStorageResourceGroup
`func (o *RecoverAzureVmNewSourceConfig) UnsetStorageResourceGroup()`

UnsetStorageResourceGroup ensures that no value is present for StorageResourceGroup, not even an explicit nil
### GetSubscription

`func (o *RecoverAzureVmNewSourceConfig) GetSubscription() RecoverAzureVmNewSourceConfigSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *RecoverAzureVmNewSourceConfig) GetSubscriptionOk() (*RecoverAzureVmNewSourceConfigSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *RecoverAzureVmNewSourceConfig) SetSubscription(v RecoverAzureVmNewSourceConfigSubscription)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *RecoverAzureVmNewSourceConfig) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### SetSubscriptionNil

`func (o *RecoverAzureVmNewSourceConfig) SetSubscriptionNil(b bool)`

 SetSubscriptionNil sets the value for Subscription to be an explicit nil

### UnsetSubscription
`func (o *RecoverAzureVmNewSourceConfig) UnsetSubscription()`

UnsetSubscription ensures that no value is present for Subscription, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


