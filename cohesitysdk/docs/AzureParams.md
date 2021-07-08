# AzureParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataDiskType** | Pointer to **NullableString** | Specifies the disk type used by the data. &#39;kPremiumSSD&#39; is disk type backed by SSDs, delivers high performance, low latency disk support for VMs running I/O intensive workloads. &#39;kStandardSSD&#39; implies disk type that offers more consistent performance and reliability than HDD. &#39;kStandardHDD&#39; implies disk type backed by HDDs, delivers cost effective storage. | [optional] 
**InstanceId** | Pointer to **NullableInt64** | Specifies Type of VM (e.g. small, medium, large) when cloning the VM in Azure. | [optional] 
**NetworkResourceGroupId** | Pointer to **NullableInt64** | Specifies id of the resource group for the selected virtual network. | [optional] 
**OsDiskType** | Pointer to **NullableString** | Specifies the disk type used by the OS. &#39;kPremiumSSD&#39; is disk type backed by SSDs, delivers high performance, low latency disk support for VMs running I/O intensive workloads. &#39;kStandardSSD&#39; implies disk type that offers more consistent performance and reliability than HDD. &#39;kStandardHDD&#39; implies disk type backed by HDDs, delivers cost effective storage. | [optional] 
**ResourceGroup** | Pointer to **NullableInt64** | Specifies id of the Azure resource group. Its value is globally unique within Azure. | [optional] 
**StorageAccount** | Pointer to **NullableInt64** | Specifies id of the storage account that will contain the storage container within which we will create the blob that will become the VHD disk for the cloned VM. | [optional] 
**StorageContainer** | Pointer to **NullableInt64** | Specifies id of the storage container within the above storage account. | [optional] 
**StorageResourceGroupId** | Pointer to **NullableInt64** | Specifies id of the resource group for the selected storage account. | [optional] 
**SubnetId** | Pointer to **NullableInt64** | Specifies Id of the subnet within the above virtual network. | [optional] 
**TempVmResourceGroupId** | Pointer to **NullableInt64** | Specifies the resource group where temporary VM needs to be created. | [optional] 
**TempVmStorageAccountId** | Pointer to **NullableInt64** | Specifies the Storage account where temporary VM needs to be created. | [optional] 
**TempVmStorageContainerId** | Pointer to **NullableInt64** | Specifies the Storage container where temporary VM needs to be created. | [optional] 
**TempVmSubnetId** | Pointer to **NullableInt64** | Specifies the Subnet where temporary VM needs to be created. | [optional] 
**TempVmVirtualNetworkId** | Pointer to **NullableInt64** | Specifies the Virtual network where temporary VM needs to be created. | [optional] 
**VirtualNetworkId** | Pointer to **NullableInt64** | Specifies Id of the Virtual Network. | [optional] 

## Methods

### NewAzureParams

`func NewAzureParams() *AzureParams`

NewAzureParams instantiates a new AzureParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureParamsWithDefaults

`func NewAzureParamsWithDefaults() *AzureParams`

NewAzureParamsWithDefaults instantiates a new AzureParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataDiskType

`func (o *AzureParams) GetDataDiskType() string`

GetDataDiskType returns the DataDiskType field if non-nil, zero value otherwise.

### GetDataDiskTypeOk

`func (o *AzureParams) GetDataDiskTypeOk() (*string, bool)`

GetDataDiskTypeOk returns a tuple with the DataDiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDiskType

`func (o *AzureParams) SetDataDiskType(v string)`

SetDataDiskType sets DataDiskType field to given value.

### HasDataDiskType

`func (o *AzureParams) HasDataDiskType() bool`

HasDataDiskType returns a boolean if a field has been set.

### SetDataDiskTypeNil

`func (o *AzureParams) SetDataDiskTypeNil(b bool)`

 SetDataDiskTypeNil sets the value for DataDiskType to be an explicit nil

### UnsetDataDiskType
`func (o *AzureParams) UnsetDataDiskType()`

UnsetDataDiskType ensures that no value is present for DataDiskType, not even an explicit nil
### GetInstanceId

`func (o *AzureParams) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *AzureParams) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *AzureParams) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *AzureParams) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### SetInstanceIdNil

`func (o *AzureParams) SetInstanceIdNil(b bool)`

 SetInstanceIdNil sets the value for InstanceId to be an explicit nil

### UnsetInstanceId
`func (o *AzureParams) UnsetInstanceId()`

UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil
### GetNetworkResourceGroupId

`func (o *AzureParams) GetNetworkResourceGroupId() int64`

GetNetworkResourceGroupId returns the NetworkResourceGroupId field if non-nil, zero value otherwise.

### GetNetworkResourceGroupIdOk

`func (o *AzureParams) GetNetworkResourceGroupIdOk() (*int64, bool)`

GetNetworkResourceGroupIdOk returns a tuple with the NetworkResourceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkResourceGroupId

`func (o *AzureParams) SetNetworkResourceGroupId(v int64)`

SetNetworkResourceGroupId sets NetworkResourceGroupId field to given value.

### HasNetworkResourceGroupId

`func (o *AzureParams) HasNetworkResourceGroupId() bool`

HasNetworkResourceGroupId returns a boolean if a field has been set.

### SetNetworkResourceGroupIdNil

`func (o *AzureParams) SetNetworkResourceGroupIdNil(b bool)`

 SetNetworkResourceGroupIdNil sets the value for NetworkResourceGroupId to be an explicit nil

### UnsetNetworkResourceGroupId
`func (o *AzureParams) UnsetNetworkResourceGroupId()`

UnsetNetworkResourceGroupId ensures that no value is present for NetworkResourceGroupId, not even an explicit nil
### GetOsDiskType

`func (o *AzureParams) GetOsDiskType() string`

GetOsDiskType returns the OsDiskType field if non-nil, zero value otherwise.

### GetOsDiskTypeOk

`func (o *AzureParams) GetOsDiskTypeOk() (*string, bool)`

GetOsDiskTypeOk returns a tuple with the OsDiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDiskType

`func (o *AzureParams) SetOsDiskType(v string)`

SetOsDiskType sets OsDiskType field to given value.

### HasOsDiskType

`func (o *AzureParams) HasOsDiskType() bool`

HasOsDiskType returns a boolean if a field has been set.

### SetOsDiskTypeNil

`func (o *AzureParams) SetOsDiskTypeNil(b bool)`

 SetOsDiskTypeNil sets the value for OsDiskType to be an explicit nil

### UnsetOsDiskType
`func (o *AzureParams) UnsetOsDiskType()`

UnsetOsDiskType ensures that no value is present for OsDiskType, not even an explicit nil
### GetResourceGroup

`func (o *AzureParams) GetResourceGroup() int64`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *AzureParams) GetResourceGroupOk() (*int64, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *AzureParams) SetResourceGroup(v int64)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *AzureParams) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### SetResourceGroupNil

`func (o *AzureParams) SetResourceGroupNil(b bool)`

 SetResourceGroupNil sets the value for ResourceGroup to be an explicit nil

### UnsetResourceGroup
`func (o *AzureParams) UnsetResourceGroup()`

UnsetResourceGroup ensures that no value is present for ResourceGroup, not even an explicit nil
### GetStorageAccount

`func (o *AzureParams) GetStorageAccount() int64`

GetStorageAccount returns the StorageAccount field if non-nil, zero value otherwise.

### GetStorageAccountOk

`func (o *AzureParams) GetStorageAccountOk() (*int64, bool)`

GetStorageAccountOk returns a tuple with the StorageAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccount

`func (o *AzureParams) SetStorageAccount(v int64)`

SetStorageAccount sets StorageAccount field to given value.

### HasStorageAccount

`func (o *AzureParams) HasStorageAccount() bool`

HasStorageAccount returns a boolean if a field has been set.

### SetStorageAccountNil

`func (o *AzureParams) SetStorageAccountNil(b bool)`

 SetStorageAccountNil sets the value for StorageAccount to be an explicit nil

### UnsetStorageAccount
`func (o *AzureParams) UnsetStorageAccount()`

UnsetStorageAccount ensures that no value is present for StorageAccount, not even an explicit nil
### GetStorageContainer

`func (o *AzureParams) GetStorageContainer() int64`

GetStorageContainer returns the StorageContainer field if non-nil, zero value otherwise.

### GetStorageContainerOk

`func (o *AzureParams) GetStorageContainerOk() (*int64, bool)`

GetStorageContainerOk returns a tuple with the StorageContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageContainer

`func (o *AzureParams) SetStorageContainer(v int64)`

SetStorageContainer sets StorageContainer field to given value.

### HasStorageContainer

`func (o *AzureParams) HasStorageContainer() bool`

HasStorageContainer returns a boolean if a field has been set.

### SetStorageContainerNil

`func (o *AzureParams) SetStorageContainerNil(b bool)`

 SetStorageContainerNil sets the value for StorageContainer to be an explicit nil

### UnsetStorageContainer
`func (o *AzureParams) UnsetStorageContainer()`

UnsetStorageContainer ensures that no value is present for StorageContainer, not even an explicit nil
### GetStorageResourceGroupId

`func (o *AzureParams) GetStorageResourceGroupId() int64`

GetStorageResourceGroupId returns the StorageResourceGroupId field if non-nil, zero value otherwise.

### GetStorageResourceGroupIdOk

`func (o *AzureParams) GetStorageResourceGroupIdOk() (*int64, bool)`

GetStorageResourceGroupIdOk returns a tuple with the StorageResourceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageResourceGroupId

`func (o *AzureParams) SetStorageResourceGroupId(v int64)`

SetStorageResourceGroupId sets StorageResourceGroupId field to given value.

### HasStorageResourceGroupId

`func (o *AzureParams) HasStorageResourceGroupId() bool`

HasStorageResourceGroupId returns a boolean if a field has been set.

### SetStorageResourceGroupIdNil

`func (o *AzureParams) SetStorageResourceGroupIdNil(b bool)`

 SetStorageResourceGroupIdNil sets the value for StorageResourceGroupId to be an explicit nil

### UnsetStorageResourceGroupId
`func (o *AzureParams) UnsetStorageResourceGroupId()`

UnsetStorageResourceGroupId ensures that no value is present for StorageResourceGroupId, not even an explicit nil
### GetSubnetId

`func (o *AzureParams) GetSubnetId() int64`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *AzureParams) GetSubnetIdOk() (*int64, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *AzureParams) SetSubnetId(v int64)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *AzureParams) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetIdNil

`func (o *AzureParams) SetSubnetIdNil(b bool)`

 SetSubnetIdNil sets the value for SubnetId to be an explicit nil

### UnsetSubnetId
`func (o *AzureParams) UnsetSubnetId()`

UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil
### GetTempVmResourceGroupId

`func (o *AzureParams) GetTempVmResourceGroupId() int64`

GetTempVmResourceGroupId returns the TempVmResourceGroupId field if non-nil, zero value otherwise.

### GetTempVmResourceGroupIdOk

`func (o *AzureParams) GetTempVmResourceGroupIdOk() (*int64, bool)`

GetTempVmResourceGroupIdOk returns a tuple with the TempVmResourceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempVmResourceGroupId

`func (o *AzureParams) SetTempVmResourceGroupId(v int64)`

SetTempVmResourceGroupId sets TempVmResourceGroupId field to given value.

### HasTempVmResourceGroupId

`func (o *AzureParams) HasTempVmResourceGroupId() bool`

HasTempVmResourceGroupId returns a boolean if a field has been set.

### SetTempVmResourceGroupIdNil

`func (o *AzureParams) SetTempVmResourceGroupIdNil(b bool)`

 SetTempVmResourceGroupIdNil sets the value for TempVmResourceGroupId to be an explicit nil

### UnsetTempVmResourceGroupId
`func (o *AzureParams) UnsetTempVmResourceGroupId()`

UnsetTempVmResourceGroupId ensures that no value is present for TempVmResourceGroupId, not even an explicit nil
### GetTempVmStorageAccountId

`func (o *AzureParams) GetTempVmStorageAccountId() int64`

GetTempVmStorageAccountId returns the TempVmStorageAccountId field if non-nil, zero value otherwise.

### GetTempVmStorageAccountIdOk

`func (o *AzureParams) GetTempVmStorageAccountIdOk() (*int64, bool)`

GetTempVmStorageAccountIdOk returns a tuple with the TempVmStorageAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempVmStorageAccountId

`func (o *AzureParams) SetTempVmStorageAccountId(v int64)`

SetTempVmStorageAccountId sets TempVmStorageAccountId field to given value.

### HasTempVmStorageAccountId

`func (o *AzureParams) HasTempVmStorageAccountId() bool`

HasTempVmStorageAccountId returns a boolean if a field has been set.

### SetTempVmStorageAccountIdNil

`func (o *AzureParams) SetTempVmStorageAccountIdNil(b bool)`

 SetTempVmStorageAccountIdNil sets the value for TempVmStorageAccountId to be an explicit nil

### UnsetTempVmStorageAccountId
`func (o *AzureParams) UnsetTempVmStorageAccountId()`

UnsetTempVmStorageAccountId ensures that no value is present for TempVmStorageAccountId, not even an explicit nil
### GetTempVmStorageContainerId

`func (o *AzureParams) GetTempVmStorageContainerId() int64`

GetTempVmStorageContainerId returns the TempVmStorageContainerId field if non-nil, zero value otherwise.

### GetTempVmStorageContainerIdOk

`func (o *AzureParams) GetTempVmStorageContainerIdOk() (*int64, bool)`

GetTempVmStorageContainerIdOk returns a tuple with the TempVmStorageContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempVmStorageContainerId

`func (o *AzureParams) SetTempVmStorageContainerId(v int64)`

SetTempVmStorageContainerId sets TempVmStorageContainerId field to given value.

### HasTempVmStorageContainerId

`func (o *AzureParams) HasTempVmStorageContainerId() bool`

HasTempVmStorageContainerId returns a boolean if a field has been set.

### SetTempVmStorageContainerIdNil

`func (o *AzureParams) SetTempVmStorageContainerIdNil(b bool)`

 SetTempVmStorageContainerIdNil sets the value for TempVmStorageContainerId to be an explicit nil

### UnsetTempVmStorageContainerId
`func (o *AzureParams) UnsetTempVmStorageContainerId()`

UnsetTempVmStorageContainerId ensures that no value is present for TempVmStorageContainerId, not even an explicit nil
### GetTempVmSubnetId

`func (o *AzureParams) GetTempVmSubnetId() int64`

GetTempVmSubnetId returns the TempVmSubnetId field if non-nil, zero value otherwise.

### GetTempVmSubnetIdOk

`func (o *AzureParams) GetTempVmSubnetIdOk() (*int64, bool)`

GetTempVmSubnetIdOk returns a tuple with the TempVmSubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempVmSubnetId

`func (o *AzureParams) SetTempVmSubnetId(v int64)`

SetTempVmSubnetId sets TempVmSubnetId field to given value.

### HasTempVmSubnetId

`func (o *AzureParams) HasTempVmSubnetId() bool`

HasTempVmSubnetId returns a boolean if a field has been set.

### SetTempVmSubnetIdNil

`func (o *AzureParams) SetTempVmSubnetIdNil(b bool)`

 SetTempVmSubnetIdNil sets the value for TempVmSubnetId to be an explicit nil

### UnsetTempVmSubnetId
`func (o *AzureParams) UnsetTempVmSubnetId()`

UnsetTempVmSubnetId ensures that no value is present for TempVmSubnetId, not even an explicit nil
### GetTempVmVirtualNetworkId

`func (o *AzureParams) GetTempVmVirtualNetworkId() int64`

GetTempVmVirtualNetworkId returns the TempVmVirtualNetworkId field if non-nil, zero value otherwise.

### GetTempVmVirtualNetworkIdOk

`func (o *AzureParams) GetTempVmVirtualNetworkIdOk() (*int64, bool)`

GetTempVmVirtualNetworkIdOk returns a tuple with the TempVmVirtualNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempVmVirtualNetworkId

`func (o *AzureParams) SetTempVmVirtualNetworkId(v int64)`

SetTempVmVirtualNetworkId sets TempVmVirtualNetworkId field to given value.

### HasTempVmVirtualNetworkId

`func (o *AzureParams) HasTempVmVirtualNetworkId() bool`

HasTempVmVirtualNetworkId returns a boolean if a field has been set.

### SetTempVmVirtualNetworkIdNil

`func (o *AzureParams) SetTempVmVirtualNetworkIdNil(b bool)`

 SetTempVmVirtualNetworkIdNil sets the value for TempVmVirtualNetworkId to be an explicit nil

### UnsetTempVmVirtualNetworkId
`func (o *AzureParams) UnsetTempVmVirtualNetworkId()`

UnsetTempVmVirtualNetworkId ensures that no value is present for TempVmVirtualNetworkId, not even an explicit nil
### GetVirtualNetworkId

`func (o *AzureParams) GetVirtualNetworkId() int64`

GetVirtualNetworkId returns the VirtualNetworkId field if non-nil, zero value otherwise.

### GetVirtualNetworkIdOk

`func (o *AzureParams) GetVirtualNetworkIdOk() (*int64, bool)`

GetVirtualNetworkIdOk returns a tuple with the VirtualNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetworkId

`func (o *AzureParams) SetVirtualNetworkId(v int64)`

SetVirtualNetworkId sets VirtualNetworkId field to given value.

### HasVirtualNetworkId

`func (o *AzureParams) HasVirtualNetworkId() bool`

HasVirtualNetworkId returns a boolean if a field has been set.

### SetVirtualNetworkIdNil

`func (o *AzureParams) SetVirtualNetworkIdNil(b bool)`

 SetVirtualNetworkIdNil sets the value for VirtualNetworkId to be an explicit nil

### UnsetVirtualNetworkId
`func (o *AzureParams) UnsetVirtualNetworkId()`

UnsetVirtualNetworkId ensures that no value is present for VirtualNetworkId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


