# RecoverAzureVmNewSourceNetworkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkResourceGroup** | Pointer to [**NullableRecoverAzureVmNewSourceNetworkConfigNetworkResourceGroup**](RecoverAzureVmNewSourceNetworkConfigNetworkResourceGroup.md) |  | [optional] 
**Subnet** | [**NullableRecoverAzureVmNewSourceNetworkConfigSubnet**](RecoverAzureVmNewSourceNetworkConfigSubnet.md) |  | 
**VirtualNetwork** | [**NullableRecoverAzureVmNewSourceNetworkConfigVirtualNetwork**](RecoverAzureVmNewSourceNetworkConfigVirtualNetwork.md) |  | 

## Methods

### NewRecoverAzureVmNewSourceNetworkConfig

`func NewRecoverAzureVmNewSourceNetworkConfig(subnet NullableRecoverAzureVmNewSourceNetworkConfigSubnet, virtualNetwork NullableRecoverAzureVmNewSourceNetworkConfigVirtualNetwork, ) *RecoverAzureVmNewSourceNetworkConfig`

NewRecoverAzureVmNewSourceNetworkConfig instantiates a new RecoverAzureVmNewSourceNetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverAzureVmNewSourceNetworkConfigWithDefaults

`func NewRecoverAzureVmNewSourceNetworkConfigWithDefaults() *RecoverAzureVmNewSourceNetworkConfig`

NewRecoverAzureVmNewSourceNetworkConfigWithDefaults instantiates a new RecoverAzureVmNewSourceNetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkResourceGroup

`func (o *RecoverAzureVmNewSourceNetworkConfig) GetNetworkResourceGroup() RecoverAzureVmNewSourceNetworkConfigNetworkResourceGroup`

GetNetworkResourceGroup returns the NetworkResourceGroup field if non-nil, zero value otherwise.

### GetNetworkResourceGroupOk

`func (o *RecoverAzureVmNewSourceNetworkConfig) GetNetworkResourceGroupOk() (*RecoverAzureVmNewSourceNetworkConfigNetworkResourceGroup, bool)`

GetNetworkResourceGroupOk returns a tuple with the NetworkResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkResourceGroup

`func (o *RecoverAzureVmNewSourceNetworkConfig) SetNetworkResourceGroup(v RecoverAzureVmNewSourceNetworkConfigNetworkResourceGroup)`

SetNetworkResourceGroup sets NetworkResourceGroup field to given value.

### HasNetworkResourceGroup

`func (o *RecoverAzureVmNewSourceNetworkConfig) HasNetworkResourceGroup() bool`

HasNetworkResourceGroup returns a boolean if a field has been set.

### SetNetworkResourceGroupNil

`func (o *RecoverAzureVmNewSourceNetworkConfig) SetNetworkResourceGroupNil(b bool)`

 SetNetworkResourceGroupNil sets the value for NetworkResourceGroup to be an explicit nil

### UnsetNetworkResourceGroup
`func (o *RecoverAzureVmNewSourceNetworkConfig) UnsetNetworkResourceGroup()`

UnsetNetworkResourceGroup ensures that no value is present for NetworkResourceGroup, not even an explicit nil
### GetSubnet

`func (o *RecoverAzureVmNewSourceNetworkConfig) GetSubnet() RecoverAzureVmNewSourceNetworkConfigSubnet`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *RecoverAzureVmNewSourceNetworkConfig) GetSubnetOk() (*RecoverAzureVmNewSourceNetworkConfigSubnet, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *RecoverAzureVmNewSourceNetworkConfig) SetSubnet(v RecoverAzureVmNewSourceNetworkConfigSubnet)`

SetSubnet sets Subnet field to given value.


### SetSubnetNil

`func (o *RecoverAzureVmNewSourceNetworkConfig) SetSubnetNil(b bool)`

 SetSubnetNil sets the value for Subnet to be an explicit nil

### UnsetSubnet
`func (o *RecoverAzureVmNewSourceNetworkConfig) UnsetSubnet()`

UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
### GetVirtualNetwork

`func (o *RecoverAzureVmNewSourceNetworkConfig) GetVirtualNetwork() RecoverAzureVmNewSourceNetworkConfigVirtualNetwork`

GetVirtualNetwork returns the VirtualNetwork field if non-nil, zero value otherwise.

### GetVirtualNetworkOk

`func (o *RecoverAzureVmNewSourceNetworkConfig) GetVirtualNetworkOk() (*RecoverAzureVmNewSourceNetworkConfigVirtualNetwork, bool)`

GetVirtualNetworkOk returns a tuple with the VirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetwork

`func (o *RecoverAzureVmNewSourceNetworkConfig) SetVirtualNetwork(v RecoverAzureVmNewSourceNetworkConfigVirtualNetwork)`

SetVirtualNetwork sets VirtualNetwork field to given value.


### SetVirtualNetworkNil

`func (o *RecoverAzureVmNewSourceNetworkConfig) SetVirtualNetworkNil(b bool)`

 SetVirtualNetworkNil sets the value for VirtualNetwork to be an explicit nil

### UnsetVirtualNetwork
`func (o *RecoverAzureVmNewSourceNetworkConfig) UnsetVirtualNetwork()`

UnsetVirtualNetwork ensures that no value is present for VirtualNetwork, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


