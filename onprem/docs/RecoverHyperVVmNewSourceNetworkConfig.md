# RecoverHyperVVmNewSourceNetworkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetachNetwork** | Pointer to **NullableBool** | If this is set to true, then the network will be detached from the recovered VMs. All the other networking parameters set will be ignored if set to true. Default value is false. | [optional] 
**VirtualSwitch** | Pointer to [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the virtual switch that will attached to all the network interfaces of the VMs being recovered. This parameter is mandatory if detach network is specified as false. | [optional] 

## Methods

### NewRecoverHyperVVmNewSourceNetworkConfig

`func NewRecoverHyperVVmNewSourceNetworkConfig() *RecoverHyperVVmNewSourceNetworkConfig`

NewRecoverHyperVVmNewSourceNetworkConfig instantiates a new RecoverHyperVVmNewSourceNetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverHyperVVmNewSourceNetworkConfigWithDefaults

`func NewRecoverHyperVVmNewSourceNetworkConfigWithDefaults() *RecoverHyperVVmNewSourceNetworkConfig`

NewRecoverHyperVVmNewSourceNetworkConfigWithDefaults instantiates a new RecoverHyperVVmNewSourceNetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetachNetwork

`func (o *RecoverHyperVVmNewSourceNetworkConfig) GetDetachNetwork() bool`

GetDetachNetwork returns the DetachNetwork field if non-nil, zero value otherwise.

### GetDetachNetworkOk

`func (o *RecoverHyperVVmNewSourceNetworkConfig) GetDetachNetworkOk() (*bool, bool)`

GetDetachNetworkOk returns a tuple with the DetachNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetachNetwork

`func (o *RecoverHyperVVmNewSourceNetworkConfig) SetDetachNetwork(v bool)`

SetDetachNetwork sets DetachNetwork field to given value.

### HasDetachNetwork

`func (o *RecoverHyperVVmNewSourceNetworkConfig) HasDetachNetwork() bool`

HasDetachNetwork returns a boolean if a field has been set.

### SetDetachNetworkNil

`func (o *RecoverHyperVVmNewSourceNetworkConfig) SetDetachNetworkNil(b bool)`

 SetDetachNetworkNil sets the value for DetachNetwork to be an explicit nil

### UnsetDetachNetwork
`func (o *RecoverHyperVVmNewSourceNetworkConfig) UnsetDetachNetwork()`

UnsetDetachNetwork ensures that no value is present for DetachNetwork, not even an explicit nil
### GetVirtualSwitch

`func (o *RecoverHyperVVmNewSourceNetworkConfig) GetVirtualSwitch() RecoveryObjectIdentifier`

GetVirtualSwitch returns the VirtualSwitch field if non-nil, zero value otherwise.

### GetVirtualSwitchOk

`func (o *RecoverHyperVVmNewSourceNetworkConfig) GetVirtualSwitchOk() (*RecoveryObjectIdentifier, bool)`

GetVirtualSwitchOk returns a tuple with the VirtualSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualSwitch

`func (o *RecoverHyperVVmNewSourceNetworkConfig) SetVirtualSwitch(v RecoveryObjectIdentifier)`

SetVirtualSwitch sets VirtualSwitch field to given value.

### HasVirtualSwitch

`func (o *RecoverHyperVVmNewSourceNetworkConfig) HasVirtualSwitch() bool`

HasVirtualSwitch returns a boolean if a field has been set.

### SetVirtualSwitchNil

`func (o *RecoverHyperVVmNewSourceNetworkConfig) SetVirtualSwitchNil(b bool)`

 SetVirtualSwitchNil sets the value for VirtualSwitch to be an explicit nil

### UnsetVirtualSwitch
`func (o *RecoverHyperVVmNewSourceNetworkConfig) UnsetVirtualSwitch()`

UnsetVirtualSwitch ensures that no value is present for VirtualSwitch, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


