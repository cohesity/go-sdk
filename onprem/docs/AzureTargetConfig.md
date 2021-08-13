# AzureTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceId** | **NullableInt64** | Specifies the source id of the Azure protection source registered on Cohesity cluster. | 
**Name** | Pointer to **NullableString** | Specifies the name of the Azure Replication target. | [optional] [readonly] 
**ResourceGroup** | Pointer to **NullableInt64** | Specifies id of the Azure resource group used to filter regions in UI. | [optional] 
**ResourceGroupName** | Pointer to **NullableString** | Specifies name of the Azure resource group used to filter regions in UI. | [optional] [readonly] 
**StorageAccount** | Pointer to **NullableInt32** | Specifies id of the storage account of Azure replication target which will contain storage container. | [optional] [readonly] 
**StorageAccountName** | Pointer to **NullableString** | Specifies name of the storage account of Azure replication target which will contain storage container. | [optional] [readonly] 
**StorageContainer** | Pointer to **NullableInt32** | Specifies id of the storage container of Azure Replication target. | [optional] [readonly] 
**StorageContainerName** | Pointer to **NullableString** | Specifies name of the storage container of Azure Replication target. | [optional] [readonly] 
**StorageResourceGroup** | Pointer to **NullableInt32** | Specifies id of the storage resource group of Azure Replication target. | [optional] [readonly] 
**StorageResourceGroupName** | Pointer to **NullableString** | Specifies name of the storage resource group of Azure Replication target. | [optional] [readonly] 

## Methods

### NewAzureTargetConfig

`func NewAzureTargetConfig(sourceId NullableInt64, ) *AzureTargetConfig`

NewAzureTargetConfig instantiates a new AzureTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureTargetConfigWithDefaults

`func NewAzureTargetConfigWithDefaults() *AzureTargetConfig`

NewAzureTargetConfigWithDefaults instantiates a new AzureTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceId

`func (o *AzureTargetConfig) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *AzureTargetConfig) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *AzureTargetConfig) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.


### SetSourceIdNil

`func (o *AzureTargetConfig) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *AzureTargetConfig) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetName

`func (o *AzureTargetConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureTargetConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureTargetConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AzureTargetConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AzureTargetConfig) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AzureTargetConfig) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetResourceGroup

`func (o *AzureTargetConfig) GetResourceGroup() int64`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *AzureTargetConfig) GetResourceGroupOk() (*int64, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *AzureTargetConfig) SetResourceGroup(v int64)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *AzureTargetConfig) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### SetResourceGroupNil

`func (o *AzureTargetConfig) SetResourceGroupNil(b bool)`

 SetResourceGroupNil sets the value for ResourceGroup to be an explicit nil

### UnsetResourceGroup
`func (o *AzureTargetConfig) UnsetResourceGroup()`

UnsetResourceGroup ensures that no value is present for ResourceGroup, not even an explicit nil
### GetResourceGroupName

`func (o *AzureTargetConfig) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *AzureTargetConfig) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *AzureTargetConfig) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.

### HasResourceGroupName

`func (o *AzureTargetConfig) HasResourceGroupName() bool`

HasResourceGroupName returns a boolean if a field has been set.

### SetResourceGroupNameNil

`func (o *AzureTargetConfig) SetResourceGroupNameNil(b bool)`

 SetResourceGroupNameNil sets the value for ResourceGroupName to be an explicit nil

### UnsetResourceGroupName
`func (o *AzureTargetConfig) UnsetResourceGroupName()`

UnsetResourceGroupName ensures that no value is present for ResourceGroupName, not even an explicit nil
### GetStorageAccount

`func (o *AzureTargetConfig) GetStorageAccount() int32`

GetStorageAccount returns the StorageAccount field if non-nil, zero value otherwise.

### GetStorageAccountOk

`func (o *AzureTargetConfig) GetStorageAccountOk() (*int32, bool)`

GetStorageAccountOk returns a tuple with the StorageAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccount

`func (o *AzureTargetConfig) SetStorageAccount(v int32)`

SetStorageAccount sets StorageAccount field to given value.

### HasStorageAccount

`func (o *AzureTargetConfig) HasStorageAccount() bool`

HasStorageAccount returns a boolean if a field has been set.

### SetStorageAccountNil

`func (o *AzureTargetConfig) SetStorageAccountNil(b bool)`

 SetStorageAccountNil sets the value for StorageAccount to be an explicit nil

### UnsetStorageAccount
`func (o *AzureTargetConfig) UnsetStorageAccount()`

UnsetStorageAccount ensures that no value is present for StorageAccount, not even an explicit nil
### GetStorageAccountName

`func (o *AzureTargetConfig) GetStorageAccountName() string`

GetStorageAccountName returns the StorageAccountName field if non-nil, zero value otherwise.

### GetStorageAccountNameOk

`func (o *AzureTargetConfig) GetStorageAccountNameOk() (*string, bool)`

GetStorageAccountNameOk returns a tuple with the StorageAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccountName

`func (o *AzureTargetConfig) SetStorageAccountName(v string)`

SetStorageAccountName sets StorageAccountName field to given value.

### HasStorageAccountName

`func (o *AzureTargetConfig) HasStorageAccountName() bool`

HasStorageAccountName returns a boolean if a field has been set.

### SetStorageAccountNameNil

`func (o *AzureTargetConfig) SetStorageAccountNameNil(b bool)`

 SetStorageAccountNameNil sets the value for StorageAccountName to be an explicit nil

### UnsetStorageAccountName
`func (o *AzureTargetConfig) UnsetStorageAccountName()`

UnsetStorageAccountName ensures that no value is present for StorageAccountName, not even an explicit nil
### GetStorageContainer

`func (o *AzureTargetConfig) GetStorageContainer() int32`

GetStorageContainer returns the StorageContainer field if non-nil, zero value otherwise.

### GetStorageContainerOk

`func (o *AzureTargetConfig) GetStorageContainerOk() (*int32, bool)`

GetStorageContainerOk returns a tuple with the StorageContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageContainer

`func (o *AzureTargetConfig) SetStorageContainer(v int32)`

SetStorageContainer sets StorageContainer field to given value.

### HasStorageContainer

`func (o *AzureTargetConfig) HasStorageContainer() bool`

HasStorageContainer returns a boolean if a field has been set.

### SetStorageContainerNil

`func (o *AzureTargetConfig) SetStorageContainerNil(b bool)`

 SetStorageContainerNil sets the value for StorageContainer to be an explicit nil

### UnsetStorageContainer
`func (o *AzureTargetConfig) UnsetStorageContainer()`

UnsetStorageContainer ensures that no value is present for StorageContainer, not even an explicit nil
### GetStorageContainerName

`func (o *AzureTargetConfig) GetStorageContainerName() string`

GetStorageContainerName returns the StorageContainerName field if non-nil, zero value otherwise.

### GetStorageContainerNameOk

`func (o *AzureTargetConfig) GetStorageContainerNameOk() (*string, bool)`

GetStorageContainerNameOk returns a tuple with the StorageContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageContainerName

`func (o *AzureTargetConfig) SetStorageContainerName(v string)`

SetStorageContainerName sets StorageContainerName field to given value.

### HasStorageContainerName

`func (o *AzureTargetConfig) HasStorageContainerName() bool`

HasStorageContainerName returns a boolean if a field has been set.

### SetStorageContainerNameNil

`func (o *AzureTargetConfig) SetStorageContainerNameNil(b bool)`

 SetStorageContainerNameNil sets the value for StorageContainerName to be an explicit nil

### UnsetStorageContainerName
`func (o *AzureTargetConfig) UnsetStorageContainerName()`

UnsetStorageContainerName ensures that no value is present for StorageContainerName, not even an explicit nil
### GetStorageResourceGroup

`func (o *AzureTargetConfig) GetStorageResourceGroup() int32`

GetStorageResourceGroup returns the StorageResourceGroup field if non-nil, zero value otherwise.

### GetStorageResourceGroupOk

`func (o *AzureTargetConfig) GetStorageResourceGroupOk() (*int32, bool)`

GetStorageResourceGroupOk returns a tuple with the StorageResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageResourceGroup

`func (o *AzureTargetConfig) SetStorageResourceGroup(v int32)`

SetStorageResourceGroup sets StorageResourceGroup field to given value.

### HasStorageResourceGroup

`func (o *AzureTargetConfig) HasStorageResourceGroup() bool`

HasStorageResourceGroup returns a boolean if a field has been set.

### SetStorageResourceGroupNil

`func (o *AzureTargetConfig) SetStorageResourceGroupNil(b bool)`

 SetStorageResourceGroupNil sets the value for StorageResourceGroup to be an explicit nil

### UnsetStorageResourceGroup
`func (o *AzureTargetConfig) UnsetStorageResourceGroup()`

UnsetStorageResourceGroup ensures that no value is present for StorageResourceGroup, not even an explicit nil
### GetStorageResourceGroupName

`func (o *AzureTargetConfig) GetStorageResourceGroupName() string`

GetStorageResourceGroupName returns the StorageResourceGroupName field if non-nil, zero value otherwise.

### GetStorageResourceGroupNameOk

`func (o *AzureTargetConfig) GetStorageResourceGroupNameOk() (*string, bool)`

GetStorageResourceGroupNameOk returns a tuple with the StorageResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageResourceGroupName

`func (o *AzureTargetConfig) SetStorageResourceGroupName(v string)`

SetStorageResourceGroupName sets StorageResourceGroupName field to given value.

### HasStorageResourceGroupName

`func (o *AzureTargetConfig) HasStorageResourceGroupName() bool`

HasStorageResourceGroupName returns a boolean if a field has been set.

### SetStorageResourceGroupNameNil

`func (o *AzureTargetConfig) SetStorageResourceGroupNameNil(b bool)`

 SetStorageResourceGroupNameNil sets the value for StorageResourceGroupName to be an explicit nil

### UnsetStorageResourceGroupName
`func (o *AzureTargetConfig) UnsetStorageResourceGroupName()`

UnsetStorageResourceGroupName ensures that no value is present for StorageResourceGroupName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


