# AzureCloudSpinParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilitySetId** | Pointer to **NullableInt64** | Specifies the availability set. | [optional] 
**NetworkResourceGroupId** | Pointer to **NullableInt64** | Specifies id of the resource group for the selected virtual network. | [optional] 
**ResourceGroupId** | Pointer to **NullableInt64** | Specifies id of the Azure resource group. Its value is globally unique within Azure. | [optional] 
**StorageAccountId** | Pointer to **NullableInt64** | Specifies id of the storage account that will contain the storage container within which we will create the blob that will become the VHD disk for the cloned VM. | [optional] 
**StorageContainerId** | Pointer to **NullableInt64** | Specifies id of the storage container within the above storage account. | [optional] 
**StorageResourceGroupId** | Pointer to **NullableInt64** | Specifies id of the resource group for the selected storage account. | [optional] 
**TempVmResourceGroupId** | Pointer to **NullableInt64** | Specifies id of the temporary Azure resource group. | [optional] 
**TempVmStorageAccountId** | Pointer to **NullableInt64** | Specifies id of the temporary VM storage account that will contain the storage container within which we will create the blob that will become the VHD disk for the cloned VM. | [optional] 
**TempVmStorageContainerId** | Pointer to **NullableInt64** | Specifies id of the temporary VM storage container within the above storage account. | [optional] 
**TempVmSubnetId** | Pointer to **NullableInt64** | Specifies Id of the temporary VM subnet within the above virtual network. | [optional] 
**TempVmVirtualNetworkId** | Pointer to **NullableInt64** | Specifies Id of the temporary VM Virtual Network. | [optional] 

## Methods

### NewAzureCloudSpinParams

`func NewAzureCloudSpinParams() *AzureCloudSpinParams`

NewAzureCloudSpinParams instantiates a new AzureCloudSpinParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureCloudSpinParamsWithDefaults

`func NewAzureCloudSpinParamsWithDefaults() *AzureCloudSpinParams`

NewAzureCloudSpinParamsWithDefaults instantiates a new AzureCloudSpinParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilitySetId

`func (o *AzureCloudSpinParams) GetAvailabilitySetId() int64`

GetAvailabilitySetId returns the AvailabilitySetId field if non-nil, zero value otherwise.

### GetAvailabilitySetIdOk

`func (o *AzureCloudSpinParams) GetAvailabilitySetIdOk() (*int64, bool)`

GetAvailabilitySetIdOk returns a tuple with the AvailabilitySetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilitySetId

`func (o *AzureCloudSpinParams) SetAvailabilitySetId(v int64)`

SetAvailabilitySetId sets AvailabilitySetId field to given value.

### HasAvailabilitySetId

`func (o *AzureCloudSpinParams) HasAvailabilitySetId() bool`

HasAvailabilitySetId returns a boolean if a field has been set.

### SetAvailabilitySetIdNil

`func (o *AzureCloudSpinParams) SetAvailabilitySetIdNil(b bool)`

 SetAvailabilitySetIdNil sets the value for AvailabilitySetId to be an explicit nil

### UnsetAvailabilitySetId
`func (o *AzureCloudSpinParams) UnsetAvailabilitySetId()`

UnsetAvailabilitySetId ensures that no value is present for AvailabilitySetId, not even an explicit nil
### GetNetworkResourceGroupId

`func (o *AzureCloudSpinParams) GetNetworkResourceGroupId() int64`

GetNetworkResourceGroupId returns the NetworkResourceGroupId field if non-nil, zero value otherwise.

### GetNetworkResourceGroupIdOk

`func (o *AzureCloudSpinParams) GetNetworkResourceGroupIdOk() (*int64, bool)`

GetNetworkResourceGroupIdOk returns a tuple with the NetworkResourceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkResourceGroupId

`func (o *AzureCloudSpinParams) SetNetworkResourceGroupId(v int64)`

SetNetworkResourceGroupId sets NetworkResourceGroupId field to given value.

### HasNetworkResourceGroupId

`func (o *AzureCloudSpinParams) HasNetworkResourceGroupId() bool`

HasNetworkResourceGroupId returns a boolean if a field has been set.

### SetNetworkResourceGroupIdNil

`func (o *AzureCloudSpinParams) SetNetworkResourceGroupIdNil(b bool)`

 SetNetworkResourceGroupIdNil sets the value for NetworkResourceGroupId to be an explicit nil

### UnsetNetworkResourceGroupId
`func (o *AzureCloudSpinParams) UnsetNetworkResourceGroupId()`

UnsetNetworkResourceGroupId ensures that no value is present for NetworkResourceGroupId, not even an explicit nil
### GetResourceGroupId

`func (o *AzureCloudSpinParams) GetResourceGroupId() int64`

GetResourceGroupId returns the ResourceGroupId field if non-nil, zero value otherwise.

### GetResourceGroupIdOk

`func (o *AzureCloudSpinParams) GetResourceGroupIdOk() (*int64, bool)`

GetResourceGroupIdOk returns a tuple with the ResourceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupId

`func (o *AzureCloudSpinParams) SetResourceGroupId(v int64)`

SetResourceGroupId sets ResourceGroupId field to given value.

### HasResourceGroupId

`func (o *AzureCloudSpinParams) HasResourceGroupId() bool`

HasResourceGroupId returns a boolean if a field has been set.

### SetResourceGroupIdNil

`func (o *AzureCloudSpinParams) SetResourceGroupIdNil(b bool)`

 SetResourceGroupIdNil sets the value for ResourceGroupId to be an explicit nil

### UnsetResourceGroupId
`func (o *AzureCloudSpinParams) UnsetResourceGroupId()`

UnsetResourceGroupId ensures that no value is present for ResourceGroupId, not even an explicit nil
### GetStorageAccountId

`func (o *AzureCloudSpinParams) GetStorageAccountId() int64`

GetStorageAccountId returns the StorageAccountId field if non-nil, zero value otherwise.

### GetStorageAccountIdOk

`func (o *AzureCloudSpinParams) GetStorageAccountIdOk() (*int64, bool)`

GetStorageAccountIdOk returns a tuple with the StorageAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccountId

`func (o *AzureCloudSpinParams) SetStorageAccountId(v int64)`

SetStorageAccountId sets StorageAccountId field to given value.

### HasStorageAccountId

`func (o *AzureCloudSpinParams) HasStorageAccountId() bool`

HasStorageAccountId returns a boolean if a field has been set.

### SetStorageAccountIdNil

`func (o *AzureCloudSpinParams) SetStorageAccountIdNil(b bool)`

 SetStorageAccountIdNil sets the value for StorageAccountId to be an explicit nil

### UnsetStorageAccountId
`func (o *AzureCloudSpinParams) UnsetStorageAccountId()`

UnsetStorageAccountId ensures that no value is present for StorageAccountId, not even an explicit nil
### GetStorageContainerId

`func (o *AzureCloudSpinParams) GetStorageContainerId() int64`

GetStorageContainerId returns the StorageContainerId field if non-nil, zero value otherwise.

### GetStorageContainerIdOk

`func (o *AzureCloudSpinParams) GetStorageContainerIdOk() (*int64, bool)`

GetStorageContainerIdOk returns a tuple with the StorageContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageContainerId

`func (o *AzureCloudSpinParams) SetStorageContainerId(v int64)`

SetStorageContainerId sets StorageContainerId field to given value.

### HasStorageContainerId

`func (o *AzureCloudSpinParams) HasStorageContainerId() bool`

HasStorageContainerId returns a boolean if a field has been set.

### SetStorageContainerIdNil

`func (o *AzureCloudSpinParams) SetStorageContainerIdNil(b bool)`

 SetStorageContainerIdNil sets the value for StorageContainerId to be an explicit nil

### UnsetStorageContainerId
`func (o *AzureCloudSpinParams) UnsetStorageContainerId()`

UnsetStorageContainerId ensures that no value is present for StorageContainerId, not even an explicit nil
### GetStorageResourceGroupId

`func (o *AzureCloudSpinParams) GetStorageResourceGroupId() int64`

GetStorageResourceGroupId returns the StorageResourceGroupId field if non-nil, zero value otherwise.

### GetStorageResourceGroupIdOk

`func (o *AzureCloudSpinParams) GetStorageResourceGroupIdOk() (*int64, bool)`

GetStorageResourceGroupIdOk returns a tuple with the StorageResourceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageResourceGroupId

`func (o *AzureCloudSpinParams) SetStorageResourceGroupId(v int64)`

SetStorageResourceGroupId sets StorageResourceGroupId field to given value.

### HasStorageResourceGroupId

`func (o *AzureCloudSpinParams) HasStorageResourceGroupId() bool`

HasStorageResourceGroupId returns a boolean if a field has been set.

### SetStorageResourceGroupIdNil

`func (o *AzureCloudSpinParams) SetStorageResourceGroupIdNil(b bool)`

 SetStorageResourceGroupIdNil sets the value for StorageResourceGroupId to be an explicit nil

### UnsetStorageResourceGroupId
`func (o *AzureCloudSpinParams) UnsetStorageResourceGroupId()`

UnsetStorageResourceGroupId ensures that no value is present for StorageResourceGroupId, not even an explicit nil
### GetTempVmResourceGroupId

`func (o *AzureCloudSpinParams) GetTempVmResourceGroupId() int64`

GetTempVmResourceGroupId returns the TempVmResourceGroupId field if non-nil, zero value otherwise.

### GetTempVmResourceGroupIdOk

`func (o *AzureCloudSpinParams) GetTempVmResourceGroupIdOk() (*int64, bool)`

GetTempVmResourceGroupIdOk returns a tuple with the TempVmResourceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempVmResourceGroupId

`func (o *AzureCloudSpinParams) SetTempVmResourceGroupId(v int64)`

SetTempVmResourceGroupId sets TempVmResourceGroupId field to given value.

### HasTempVmResourceGroupId

`func (o *AzureCloudSpinParams) HasTempVmResourceGroupId() bool`

HasTempVmResourceGroupId returns a boolean if a field has been set.

### SetTempVmResourceGroupIdNil

`func (o *AzureCloudSpinParams) SetTempVmResourceGroupIdNil(b bool)`

 SetTempVmResourceGroupIdNil sets the value for TempVmResourceGroupId to be an explicit nil

### UnsetTempVmResourceGroupId
`func (o *AzureCloudSpinParams) UnsetTempVmResourceGroupId()`

UnsetTempVmResourceGroupId ensures that no value is present for TempVmResourceGroupId, not even an explicit nil
### GetTempVmStorageAccountId

`func (o *AzureCloudSpinParams) GetTempVmStorageAccountId() int64`

GetTempVmStorageAccountId returns the TempVmStorageAccountId field if non-nil, zero value otherwise.

### GetTempVmStorageAccountIdOk

`func (o *AzureCloudSpinParams) GetTempVmStorageAccountIdOk() (*int64, bool)`

GetTempVmStorageAccountIdOk returns a tuple with the TempVmStorageAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempVmStorageAccountId

`func (o *AzureCloudSpinParams) SetTempVmStorageAccountId(v int64)`

SetTempVmStorageAccountId sets TempVmStorageAccountId field to given value.

### HasTempVmStorageAccountId

`func (o *AzureCloudSpinParams) HasTempVmStorageAccountId() bool`

HasTempVmStorageAccountId returns a boolean if a field has been set.

### SetTempVmStorageAccountIdNil

`func (o *AzureCloudSpinParams) SetTempVmStorageAccountIdNil(b bool)`

 SetTempVmStorageAccountIdNil sets the value for TempVmStorageAccountId to be an explicit nil

### UnsetTempVmStorageAccountId
`func (o *AzureCloudSpinParams) UnsetTempVmStorageAccountId()`

UnsetTempVmStorageAccountId ensures that no value is present for TempVmStorageAccountId, not even an explicit nil
### GetTempVmStorageContainerId

`func (o *AzureCloudSpinParams) GetTempVmStorageContainerId() int64`

GetTempVmStorageContainerId returns the TempVmStorageContainerId field if non-nil, zero value otherwise.

### GetTempVmStorageContainerIdOk

`func (o *AzureCloudSpinParams) GetTempVmStorageContainerIdOk() (*int64, bool)`

GetTempVmStorageContainerIdOk returns a tuple with the TempVmStorageContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempVmStorageContainerId

`func (o *AzureCloudSpinParams) SetTempVmStorageContainerId(v int64)`

SetTempVmStorageContainerId sets TempVmStorageContainerId field to given value.

### HasTempVmStorageContainerId

`func (o *AzureCloudSpinParams) HasTempVmStorageContainerId() bool`

HasTempVmStorageContainerId returns a boolean if a field has been set.

### SetTempVmStorageContainerIdNil

`func (o *AzureCloudSpinParams) SetTempVmStorageContainerIdNil(b bool)`

 SetTempVmStorageContainerIdNil sets the value for TempVmStorageContainerId to be an explicit nil

### UnsetTempVmStorageContainerId
`func (o *AzureCloudSpinParams) UnsetTempVmStorageContainerId()`

UnsetTempVmStorageContainerId ensures that no value is present for TempVmStorageContainerId, not even an explicit nil
### GetTempVmSubnetId

`func (o *AzureCloudSpinParams) GetTempVmSubnetId() int64`

GetTempVmSubnetId returns the TempVmSubnetId field if non-nil, zero value otherwise.

### GetTempVmSubnetIdOk

`func (o *AzureCloudSpinParams) GetTempVmSubnetIdOk() (*int64, bool)`

GetTempVmSubnetIdOk returns a tuple with the TempVmSubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempVmSubnetId

`func (o *AzureCloudSpinParams) SetTempVmSubnetId(v int64)`

SetTempVmSubnetId sets TempVmSubnetId field to given value.

### HasTempVmSubnetId

`func (o *AzureCloudSpinParams) HasTempVmSubnetId() bool`

HasTempVmSubnetId returns a boolean if a field has been set.

### SetTempVmSubnetIdNil

`func (o *AzureCloudSpinParams) SetTempVmSubnetIdNil(b bool)`

 SetTempVmSubnetIdNil sets the value for TempVmSubnetId to be an explicit nil

### UnsetTempVmSubnetId
`func (o *AzureCloudSpinParams) UnsetTempVmSubnetId()`

UnsetTempVmSubnetId ensures that no value is present for TempVmSubnetId, not even an explicit nil
### GetTempVmVirtualNetworkId

`func (o *AzureCloudSpinParams) GetTempVmVirtualNetworkId() int64`

GetTempVmVirtualNetworkId returns the TempVmVirtualNetworkId field if non-nil, zero value otherwise.

### GetTempVmVirtualNetworkIdOk

`func (o *AzureCloudSpinParams) GetTempVmVirtualNetworkIdOk() (*int64, bool)`

GetTempVmVirtualNetworkIdOk returns a tuple with the TempVmVirtualNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempVmVirtualNetworkId

`func (o *AzureCloudSpinParams) SetTempVmVirtualNetworkId(v int64)`

SetTempVmVirtualNetworkId sets TempVmVirtualNetworkId field to given value.

### HasTempVmVirtualNetworkId

`func (o *AzureCloudSpinParams) HasTempVmVirtualNetworkId() bool`

HasTempVmVirtualNetworkId returns a boolean if a field has been set.

### SetTempVmVirtualNetworkIdNil

`func (o *AzureCloudSpinParams) SetTempVmVirtualNetworkIdNil(b bool)`

 SetTempVmVirtualNetworkIdNil sets the value for TempVmVirtualNetworkId to be an explicit nil

### UnsetTempVmVirtualNetworkId
`func (o *AzureCloudSpinParams) UnsetTempVmVirtualNetworkId()`

UnsetTempVmVirtualNetworkId ensures that no value is present for TempVmVirtualNetworkId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


