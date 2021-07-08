# DeployVMsToAzureParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureManagedDiskParams** | Pointer to [**AzureManagedDiskParams**](AzureManagedDiskParams.md) |  | [optional] 
**ComputeOptions** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**NetworkResourceGroup** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**NetworkSecurityGroup** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**ResourceGroup** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**StorageAccount** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**StorageContainer** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**StorageKey** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**StorageResourceGroup** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**Subnet** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**TempVmResourceGroup** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**TempVmStorageAccount** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**TempVmStorageContainer** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**TempVmSubnet** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**TempVmVirtualNetwork** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**VirtualNetwork** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 

## Methods

### NewDeployVMsToAzureParams

`func NewDeployVMsToAzureParams() *DeployVMsToAzureParams`

NewDeployVMsToAzureParams instantiates a new DeployVMsToAzureParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployVMsToAzureParamsWithDefaults

`func NewDeployVMsToAzureParamsWithDefaults() *DeployVMsToAzureParams`

NewDeployVMsToAzureParamsWithDefaults instantiates a new DeployVMsToAzureParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzureManagedDiskParams

`func (o *DeployVMsToAzureParams) GetAzureManagedDiskParams() AzureManagedDiskParams`

GetAzureManagedDiskParams returns the AzureManagedDiskParams field if non-nil, zero value otherwise.

### GetAzureManagedDiskParamsOk

`func (o *DeployVMsToAzureParams) GetAzureManagedDiskParamsOk() (*AzureManagedDiskParams, bool)`

GetAzureManagedDiskParamsOk returns a tuple with the AzureManagedDiskParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureManagedDiskParams

`func (o *DeployVMsToAzureParams) SetAzureManagedDiskParams(v AzureManagedDiskParams)`

SetAzureManagedDiskParams sets AzureManagedDiskParams field to given value.

### HasAzureManagedDiskParams

`func (o *DeployVMsToAzureParams) HasAzureManagedDiskParams() bool`

HasAzureManagedDiskParams returns a boolean if a field has been set.

### GetComputeOptions

`func (o *DeployVMsToAzureParams) GetComputeOptions() EntityProto`

GetComputeOptions returns the ComputeOptions field if non-nil, zero value otherwise.

### GetComputeOptionsOk

`func (o *DeployVMsToAzureParams) GetComputeOptionsOk() (*EntityProto, bool)`

GetComputeOptionsOk returns a tuple with the ComputeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOptions

`func (o *DeployVMsToAzureParams) SetComputeOptions(v EntityProto)`

SetComputeOptions sets ComputeOptions field to given value.

### HasComputeOptions

`func (o *DeployVMsToAzureParams) HasComputeOptions() bool`

HasComputeOptions returns a boolean if a field has been set.

### GetNetworkResourceGroup

`func (o *DeployVMsToAzureParams) GetNetworkResourceGroup() EntityProto`

GetNetworkResourceGroup returns the NetworkResourceGroup field if non-nil, zero value otherwise.

### GetNetworkResourceGroupOk

`func (o *DeployVMsToAzureParams) GetNetworkResourceGroupOk() (*EntityProto, bool)`

GetNetworkResourceGroupOk returns a tuple with the NetworkResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkResourceGroup

`func (o *DeployVMsToAzureParams) SetNetworkResourceGroup(v EntityProto)`

SetNetworkResourceGroup sets NetworkResourceGroup field to given value.

### HasNetworkResourceGroup

`func (o *DeployVMsToAzureParams) HasNetworkResourceGroup() bool`

HasNetworkResourceGroup returns a boolean if a field has been set.

### GetNetworkSecurityGroup

`func (o *DeployVMsToAzureParams) GetNetworkSecurityGroup() EntityProto`

GetNetworkSecurityGroup returns the NetworkSecurityGroup field if non-nil, zero value otherwise.

### GetNetworkSecurityGroupOk

`func (o *DeployVMsToAzureParams) GetNetworkSecurityGroupOk() (*EntityProto, bool)`

GetNetworkSecurityGroupOk returns a tuple with the NetworkSecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSecurityGroup

`func (o *DeployVMsToAzureParams) SetNetworkSecurityGroup(v EntityProto)`

SetNetworkSecurityGroup sets NetworkSecurityGroup field to given value.

### HasNetworkSecurityGroup

`func (o *DeployVMsToAzureParams) HasNetworkSecurityGroup() bool`

HasNetworkSecurityGroup returns a boolean if a field has been set.

### GetResourceGroup

`func (o *DeployVMsToAzureParams) GetResourceGroup() EntityProto`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *DeployVMsToAzureParams) GetResourceGroupOk() (*EntityProto, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *DeployVMsToAzureParams) SetResourceGroup(v EntityProto)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *DeployVMsToAzureParams) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### GetStorageAccount

`func (o *DeployVMsToAzureParams) GetStorageAccount() EntityProto`

GetStorageAccount returns the StorageAccount field if non-nil, zero value otherwise.

### GetStorageAccountOk

`func (o *DeployVMsToAzureParams) GetStorageAccountOk() (*EntityProto, bool)`

GetStorageAccountOk returns a tuple with the StorageAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccount

`func (o *DeployVMsToAzureParams) SetStorageAccount(v EntityProto)`

SetStorageAccount sets StorageAccount field to given value.

### HasStorageAccount

`func (o *DeployVMsToAzureParams) HasStorageAccount() bool`

HasStorageAccount returns a boolean if a field has been set.

### GetStorageContainer

`func (o *DeployVMsToAzureParams) GetStorageContainer() EntityProto`

GetStorageContainer returns the StorageContainer field if non-nil, zero value otherwise.

### GetStorageContainerOk

`func (o *DeployVMsToAzureParams) GetStorageContainerOk() (*EntityProto, bool)`

GetStorageContainerOk returns a tuple with the StorageContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageContainer

`func (o *DeployVMsToAzureParams) SetStorageContainer(v EntityProto)`

SetStorageContainer sets StorageContainer field to given value.

### HasStorageContainer

`func (o *DeployVMsToAzureParams) HasStorageContainer() bool`

HasStorageContainer returns a boolean if a field has been set.

### GetStorageKey

`func (o *DeployVMsToAzureParams) GetStorageKey() EntityProto`

GetStorageKey returns the StorageKey field if non-nil, zero value otherwise.

### GetStorageKeyOk

`func (o *DeployVMsToAzureParams) GetStorageKeyOk() (*EntityProto, bool)`

GetStorageKeyOk returns a tuple with the StorageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageKey

`func (o *DeployVMsToAzureParams) SetStorageKey(v EntityProto)`

SetStorageKey sets StorageKey field to given value.

### HasStorageKey

`func (o *DeployVMsToAzureParams) HasStorageKey() bool`

HasStorageKey returns a boolean if a field has been set.

### GetStorageResourceGroup

`func (o *DeployVMsToAzureParams) GetStorageResourceGroup() EntityProto`

GetStorageResourceGroup returns the StorageResourceGroup field if non-nil, zero value otherwise.

### GetStorageResourceGroupOk

`func (o *DeployVMsToAzureParams) GetStorageResourceGroupOk() (*EntityProto, bool)`

GetStorageResourceGroupOk returns a tuple with the StorageResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageResourceGroup

`func (o *DeployVMsToAzureParams) SetStorageResourceGroup(v EntityProto)`

SetStorageResourceGroup sets StorageResourceGroup field to given value.

### HasStorageResourceGroup

`func (o *DeployVMsToAzureParams) HasStorageResourceGroup() bool`

HasStorageResourceGroup returns a boolean if a field has been set.

### GetSubnet

`func (o *DeployVMsToAzureParams) GetSubnet() EntityProto`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *DeployVMsToAzureParams) GetSubnetOk() (*EntityProto, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *DeployVMsToAzureParams) SetSubnet(v EntityProto)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *DeployVMsToAzureParams) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetTempVmResourceGroup

`func (o *DeployVMsToAzureParams) GetTempVmResourceGroup() EntityProto`

GetTempVmResourceGroup returns the TempVmResourceGroup field if non-nil, zero value otherwise.

### GetTempVmResourceGroupOk

`func (o *DeployVMsToAzureParams) GetTempVmResourceGroupOk() (*EntityProto, bool)`

GetTempVmResourceGroupOk returns a tuple with the TempVmResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempVmResourceGroup

`func (o *DeployVMsToAzureParams) SetTempVmResourceGroup(v EntityProto)`

SetTempVmResourceGroup sets TempVmResourceGroup field to given value.

### HasTempVmResourceGroup

`func (o *DeployVMsToAzureParams) HasTempVmResourceGroup() bool`

HasTempVmResourceGroup returns a boolean if a field has been set.

### GetTempVmStorageAccount

`func (o *DeployVMsToAzureParams) GetTempVmStorageAccount() EntityProto`

GetTempVmStorageAccount returns the TempVmStorageAccount field if non-nil, zero value otherwise.

### GetTempVmStorageAccountOk

`func (o *DeployVMsToAzureParams) GetTempVmStorageAccountOk() (*EntityProto, bool)`

GetTempVmStorageAccountOk returns a tuple with the TempVmStorageAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempVmStorageAccount

`func (o *DeployVMsToAzureParams) SetTempVmStorageAccount(v EntityProto)`

SetTempVmStorageAccount sets TempVmStorageAccount field to given value.

### HasTempVmStorageAccount

`func (o *DeployVMsToAzureParams) HasTempVmStorageAccount() bool`

HasTempVmStorageAccount returns a boolean if a field has been set.

### GetTempVmStorageContainer

`func (o *DeployVMsToAzureParams) GetTempVmStorageContainer() EntityProto`

GetTempVmStorageContainer returns the TempVmStorageContainer field if non-nil, zero value otherwise.

### GetTempVmStorageContainerOk

`func (o *DeployVMsToAzureParams) GetTempVmStorageContainerOk() (*EntityProto, bool)`

GetTempVmStorageContainerOk returns a tuple with the TempVmStorageContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempVmStorageContainer

`func (o *DeployVMsToAzureParams) SetTempVmStorageContainer(v EntityProto)`

SetTempVmStorageContainer sets TempVmStorageContainer field to given value.

### HasTempVmStorageContainer

`func (o *DeployVMsToAzureParams) HasTempVmStorageContainer() bool`

HasTempVmStorageContainer returns a boolean if a field has been set.

### GetTempVmSubnet

`func (o *DeployVMsToAzureParams) GetTempVmSubnet() EntityProto`

GetTempVmSubnet returns the TempVmSubnet field if non-nil, zero value otherwise.

### GetTempVmSubnetOk

`func (o *DeployVMsToAzureParams) GetTempVmSubnetOk() (*EntityProto, bool)`

GetTempVmSubnetOk returns a tuple with the TempVmSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempVmSubnet

`func (o *DeployVMsToAzureParams) SetTempVmSubnet(v EntityProto)`

SetTempVmSubnet sets TempVmSubnet field to given value.

### HasTempVmSubnet

`func (o *DeployVMsToAzureParams) HasTempVmSubnet() bool`

HasTempVmSubnet returns a boolean if a field has been set.

### GetTempVmVirtualNetwork

`func (o *DeployVMsToAzureParams) GetTempVmVirtualNetwork() EntityProto`

GetTempVmVirtualNetwork returns the TempVmVirtualNetwork field if non-nil, zero value otherwise.

### GetTempVmVirtualNetworkOk

`func (o *DeployVMsToAzureParams) GetTempVmVirtualNetworkOk() (*EntityProto, bool)`

GetTempVmVirtualNetworkOk returns a tuple with the TempVmVirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempVmVirtualNetwork

`func (o *DeployVMsToAzureParams) SetTempVmVirtualNetwork(v EntityProto)`

SetTempVmVirtualNetwork sets TempVmVirtualNetwork field to given value.

### HasTempVmVirtualNetwork

`func (o *DeployVMsToAzureParams) HasTempVmVirtualNetwork() bool`

HasTempVmVirtualNetwork returns a boolean if a field has been set.

### GetVirtualNetwork

`func (o *DeployVMsToAzureParams) GetVirtualNetwork() EntityProto`

GetVirtualNetwork returns the VirtualNetwork field if non-nil, zero value otherwise.

### GetVirtualNetworkOk

`func (o *DeployVMsToAzureParams) GetVirtualNetworkOk() (*EntityProto, bool)`

GetVirtualNetworkOk returns a tuple with the VirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetwork

`func (o *DeployVMsToAzureParams) SetVirtualNetwork(v EntityProto)`

SetVirtualNetwork sets VirtualNetwork field to given value.

### HasVirtualNetwork

`func (o *DeployVMsToAzureParams) HasVirtualNetwork() bool`

HasVirtualNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


