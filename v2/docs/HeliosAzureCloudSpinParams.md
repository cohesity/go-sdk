# HeliosAzureCloudSpinParams

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

### NewHeliosAzureCloudSpinParams

`func NewHeliosAzureCloudSpinParams() *HeliosAzureCloudSpinParams`

NewHeliosAzureCloudSpinParams instantiates a new HeliosAzureCloudSpinParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosAzureCloudSpinParamsWithDefaults

`func NewHeliosAzureCloudSpinParamsWithDefaults() *HeliosAzureCloudSpinParams`

NewHeliosAzureCloudSpinParamsWithDefaults instantiates a new HeliosAzureCloudSpinParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilitySetId

`func (o *HeliosAzureCloudSpinParams) GetAvailabilitySetId() int64`

GetAvailabilitySetId returns the AvailabilitySetId field if non-nil, zero value otherwise.

### GetAvailabilitySetIdOk

`func (o *HeliosAzureCloudSpinParams) GetAvailabilitySetIdOk() (*int64, bool)`

GetAvailabilitySetIdOk returns a tuple with the AvailabilitySetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilitySetId

`func (o *HeliosAzureCloudSpinParams) SetAvailabilitySetId(v int64)`

SetAvailabilitySetId sets AvailabilitySetId field to given value.

### HasAvailabilitySetId

`func (o *HeliosAzureCloudSpinParams) HasAvailabilitySetId() bool`

HasAvailabilitySetId returns a boolean if a field has been set.

### SetAvailabilitySetIdNil

`func (o *HeliosAzureCloudSpinParams) SetAvailabilitySetIdNil(b bool)`

 SetAvailabilitySetIdNil sets the value for AvailabilitySetId to be an explicit nil

### UnsetAvailabilitySetId
`func (o *HeliosAzureCloudSpinParams) UnsetAvailabilitySetId()`

UnsetAvailabilitySetId ensures that no value is present for AvailabilitySetId, not even an explicit nil
### GetNetworkResourceGroupId

`func (o *HeliosAzureCloudSpinParams) GetNetworkResourceGroupId() int64`

GetNetworkResourceGroupId returns the NetworkResourceGroupId field if non-nil, zero value otherwise.

### GetNetworkResourceGroupIdOk

`func (o *HeliosAzureCloudSpinParams) GetNetworkResourceGroupIdOk() (*int64, bool)`

GetNetworkResourceGroupIdOk returns a tuple with the NetworkResourceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkResourceGroupId

`func (o *HeliosAzureCloudSpinParams) SetNetworkResourceGroupId(v int64)`

SetNetworkResourceGroupId sets NetworkResourceGroupId field to given value.

### HasNetworkResourceGroupId

`func (o *HeliosAzureCloudSpinParams) HasNetworkResourceGroupId() bool`

HasNetworkResourceGroupId returns a boolean if a field has been set.

### SetNetworkResourceGroupIdNil

`func (o *HeliosAzureCloudSpinParams) SetNetworkResourceGroupIdNil(b bool)`

 SetNetworkResourceGroupIdNil sets the value for NetworkResourceGroupId to be an explicit nil

### UnsetNetworkResourceGroupId
`func (o *HeliosAzureCloudSpinParams) UnsetNetworkResourceGroupId()`

UnsetNetworkResourceGroupId ensures that no value is present for NetworkResourceGroupId, not even an explicit nil
### GetResourceGroupId

`func (o *HeliosAzureCloudSpinParams) GetResourceGroupId() int64`

GetResourceGroupId returns the ResourceGroupId field if non-nil, zero value otherwise.

### GetResourceGroupIdOk

`func (o *HeliosAzureCloudSpinParams) GetResourceGroupIdOk() (*int64, bool)`

GetResourceGroupIdOk returns a tuple with the ResourceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupId

`func (o *HeliosAzureCloudSpinParams) SetResourceGroupId(v int64)`

SetResourceGroupId sets ResourceGroupId field to given value.

### HasResourceGroupId

`func (o *HeliosAzureCloudSpinParams) HasResourceGroupId() bool`

HasResourceGroupId returns a boolean if a field has been set.

### SetResourceGroupIdNil

`func (o *HeliosAzureCloudSpinParams) SetResourceGroupIdNil(b bool)`

 SetResourceGroupIdNil sets the value for ResourceGroupId to be an explicit nil

### UnsetResourceGroupId
`func (o *HeliosAzureCloudSpinParams) UnsetResourceGroupId()`

UnsetResourceGroupId ensures that no value is present for ResourceGroupId, not even an explicit nil
### GetStorageAccountId

`func (o *HeliosAzureCloudSpinParams) GetStorageAccountId() int64`

GetStorageAccountId returns the StorageAccountId field if non-nil, zero value otherwise.

### GetStorageAccountIdOk

`func (o *HeliosAzureCloudSpinParams) GetStorageAccountIdOk() (*int64, bool)`

GetStorageAccountIdOk returns a tuple with the StorageAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccountId

`func (o *HeliosAzureCloudSpinParams) SetStorageAccountId(v int64)`

SetStorageAccountId sets StorageAccountId field to given value.

### HasStorageAccountId

`func (o *HeliosAzureCloudSpinParams) HasStorageAccountId() bool`

HasStorageAccountId returns a boolean if a field has been set.

### SetStorageAccountIdNil

`func (o *HeliosAzureCloudSpinParams) SetStorageAccountIdNil(b bool)`

 SetStorageAccountIdNil sets the value for StorageAccountId to be an explicit nil

### UnsetStorageAccountId
`func (o *HeliosAzureCloudSpinParams) UnsetStorageAccountId()`

UnsetStorageAccountId ensures that no value is present for StorageAccountId, not even an explicit nil
### GetStorageContainerId

`func (o *HeliosAzureCloudSpinParams) GetStorageContainerId() int64`

GetStorageContainerId returns the StorageContainerId field if non-nil, zero value otherwise.

### GetStorageContainerIdOk

`func (o *HeliosAzureCloudSpinParams) GetStorageContainerIdOk() (*int64, bool)`

GetStorageContainerIdOk returns a tuple with the StorageContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageContainerId

`func (o *HeliosAzureCloudSpinParams) SetStorageContainerId(v int64)`

SetStorageContainerId sets StorageContainerId field to given value.

### HasStorageContainerId

`func (o *HeliosAzureCloudSpinParams) HasStorageContainerId() bool`

HasStorageContainerId returns a boolean if a field has been set.

### SetStorageContainerIdNil

`func (o *HeliosAzureCloudSpinParams) SetStorageContainerIdNil(b bool)`

 SetStorageContainerIdNil sets the value for StorageContainerId to be an explicit nil

### UnsetStorageContainerId
`func (o *HeliosAzureCloudSpinParams) UnsetStorageContainerId()`

UnsetStorageContainerId ensures that no value is present for StorageContainerId, not even an explicit nil
### GetStorageResourceGroupId

`func (o *HeliosAzureCloudSpinParams) GetStorageResourceGroupId() int64`

GetStorageResourceGroupId returns the StorageResourceGroupId field if non-nil, zero value otherwise.

### GetStorageResourceGroupIdOk

`func (o *HeliosAzureCloudSpinParams) GetStorageResourceGroupIdOk() (*int64, bool)`

GetStorageResourceGroupIdOk returns a tuple with the StorageResourceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageResourceGroupId

`func (o *HeliosAzureCloudSpinParams) SetStorageResourceGroupId(v int64)`

SetStorageResourceGroupId sets StorageResourceGroupId field to given value.

### HasStorageResourceGroupId

`func (o *HeliosAzureCloudSpinParams) HasStorageResourceGroupId() bool`

HasStorageResourceGroupId returns a boolean if a field has been set.

### SetStorageResourceGroupIdNil

`func (o *HeliosAzureCloudSpinParams) SetStorageResourceGroupIdNil(b bool)`

 SetStorageResourceGroupIdNil sets the value for StorageResourceGroupId to be an explicit nil

### UnsetStorageResourceGroupId
`func (o *HeliosAzureCloudSpinParams) UnsetStorageResourceGroupId()`

UnsetStorageResourceGroupId ensures that no value is present for StorageResourceGroupId, not even an explicit nil
### GetTempVmResourceGroupId

`func (o *HeliosAzureCloudSpinParams) GetTempVmResourceGroupId() int64`

GetTempVmResourceGroupId returns the TempVmResourceGroupId field if non-nil, zero value otherwise.

### GetTempVmResourceGroupIdOk

`func (o *HeliosAzureCloudSpinParams) GetTempVmResourceGroupIdOk() (*int64, bool)`

GetTempVmResourceGroupIdOk returns a tuple with the TempVmResourceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempVmResourceGroupId

`func (o *HeliosAzureCloudSpinParams) SetTempVmResourceGroupId(v int64)`

SetTempVmResourceGroupId sets TempVmResourceGroupId field to given value.

### HasTempVmResourceGroupId

`func (o *HeliosAzureCloudSpinParams) HasTempVmResourceGroupId() bool`

HasTempVmResourceGroupId returns a boolean if a field has been set.

### SetTempVmResourceGroupIdNil

`func (o *HeliosAzureCloudSpinParams) SetTempVmResourceGroupIdNil(b bool)`

 SetTempVmResourceGroupIdNil sets the value for TempVmResourceGroupId to be an explicit nil

### UnsetTempVmResourceGroupId
`func (o *HeliosAzureCloudSpinParams) UnsetTempVmResourceGroupId()`

UnsetTempVmResourceGroupId ensures that no value is present for TempVmResourceGroupId, not even an explicit nil
### GetTempVmStorageAccountId

`func (o *HeliosAzureCloudSpinParams) GetTempVmStorageAccountId() int64`

GetTempVmStorageAccountId returns the TempVmStorageAccountId field if non-nil, zero value otherwise.

### GetTempVmStorageAccountIdOk

`func (o *HeliosAzureCloudSpinParams) GetTempVmStorageAccountIdOk() (*int64, bool)`

GetTempVmStorageAccountIdOk returns a tuple with the TempVmStorageAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempVmStorageAccountId

`func (o *HeliosAzureCloudSpinParams) SetTempVmStorageAccountId(v int64)`

SetTempVmStorageAccountId sets TempVmStorageAccountId field to given value.

### HasTempVmStorageAccountId

`func (o *HeliosAzureCloudSpinParams) HasTempVmStorageAccountId() bool`

HasTempVmStorageAccountId returns a boolean if a field has been set.

### SetTempVmStorageAccountIdNil

`func (o *HeliosAzureCloudSpinParams) SetTempVmStorageAccountIdNil(b bool)`

 SetTempVmStorageAccountIdNil sets the value for TempVmStorageAccountId to be an explicit nil

### UnsetTempVmStorageAccountId
`func (o *HeliosAzureCloudSpinParams) UnsetTempVmStorageAccountId()`

UnsetTempVmStorageAccountId ensures that no value is present for TempVmStorageAccountId, not even an explicit nil
### GetTempVmStorageContainerId

`func (o *HeliosAzureCloudSpinParams) GetTempVmStorageContainerId() int64`

GetTempVmStorageContainerId returns the TempVmStorageContainerId field if non-nil, zero value otherwise.

### GetTempVmStorageContainerIdOk

`func (o *HeliosAzureCloudSpinParams) GetTempVmStorageContainerIdOk() (*int64, bool)`

GetTempVmStorageContainerIdOk returns a tuple with the TempVmStorageContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempVmStorageContainerId

`func (o *HeliosAzureCloudSpinParams) SetTempVmStorageContainerId(v int64)`

SetTempVmStorageContainerId sets TempVmStorageContainerId field to given value.

### HasTempVmStorageContainerId

`func (o *HeliosAzureCloudSpinParams) HasTempVmStorageContainerId() bool`

HasTempVmStorageContainerId returns a boolean if a field has been set.

### SetTempVmStorageContainerIdNil

`func (o *HeliosAzureCloudSpinParams) SetTempVmStorageContainerIdNil(b bool)`

 SetTempVmStorageContainerIdNil sets the value for TempVmStorageContainerId to be an explicit nil

### UnsetTempVmStorageContainerId
`func (o *HeliosAzureCloudSpinParams) UnsetTempVmStorageContainerId()`

UnsetTempVmStorageContainerId ensures that no value is present for TempVmStorageContainerId, not even an explicit nil
### GetTempVmSubnetId

`func (o *HeliosAzureCloudSpinParams) GetTempVmSubnetId() int64`

GetTempVmSubnetId returns the TempVmSubnetId field if non-nil, zero value otherwise.

### GetTempVmSubnetIdOk

`func (o *HeliosAzureCloudSpinParams) GetTempVmSubnetIdOk() (*int64, bool)`

GetTempVmSubnetIdOk returns a tuple with the TempVmSubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempVmSubnetId

`func (o *HeliosAzureCloudSpinParams) SetTempVmSubnetId(v int64)`

SetTempVmSubnetId sets TempVmSubnetId field to given value.

### HasTempVmSubnetId

`func (o *HeliosAzureCloudSpinParams) HasTempVmSubnetId() bool`

HasTempVmSubnetId returns a boolean if a field has been set.

### SetTempVmSubnetIdNil

`func (o *HeliosAzureCloudSpinParams) SetTempVmSubnetIdNil(b bool)`

 SetTempVmSubnetIdNil sets the value for TempVmSubnetId to be an explicit nil

### UnsetTempVmSubnetId
`func (o *HeliosAzureCloudSpinParams) UnsetTempVmSubnetId()`

UnsetTempVmSubnetId ensures that no value is present for TempVmSubnetId, not even an explicit nil
### GetTempVmVirtualNetworkId

`func (o *HeliosAzureCloudSpinParams) GetTempVmVirtualNetworkId() int64`

GetTempVmVirtualNetworkId returns the TempVmVirtualNetworkId field if non-nil, zero value otherwise.

### GetTempVmVirtualNetworkIdOk

`func (o *HeliosAzureCloudSpinParams) GetTempVmVirtualNetworkIdOk() (*int64, bool)`

GetTempVmVirtualNetworkIdOk returns a tuple with the TempVmVirtualNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempVmVirtualNetworkId

`func (o *HeliosAzureCloudSpinParams) SetTempVmVirtualNetworkId(v int64)`

SetTempVmVirtualNetworkId sets TempVmVirtualNetworkId field to given value.

### HasTempVmVirtualNetworkId

`func (o *HeliosAzureCloudSpinParams) HasTempVmVirtualNetworkId() bool`

HasTempVmVirtualNetworkId returns a boolean if a field has been set.

### SetTempVmVirtualNetworkIdNil

`func (o *HeliosAzureCloudSpinParams) SetTempVmVirtualNetworkIdNil(b bool)`

 SetTempVmVirtualNetworkIdNil sets the value for TempVmVirtualNetworkId to be an explicit nil

### UnsetTempVmVirtualNetworkId
`func (o *HeliosAzureCloudSpinParams) UnsetTempVmVirtualNetworkId()`

UnsetTempVmVirtualNetworkId ensures that no value is present for TempVmVirtualNetworkId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


