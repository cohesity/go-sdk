# RecoverHyperVVmStandaloneHostSourceConfigNetworkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetachNetwork** | Pointer to **NullableBool** | If this is set to true, then the network will be detached from the recovered VMs. All the other networking parameters set will be ignored if set to true. Default value is false. | [optional] 
**VirtualSwitch** | Pointer to [**NullableRecoverHyperVVmNewSourceNetworkConfigVirtualSwitch**](RecoverHyperVVmNewSourceNetworkConfigVirtualSwitch.md) |  | [optional] 

## Methods

### NewRecoverHyperVVmStandaloneHostSourceConfigNetworkConfig

`func NewRecoverHyperVVmStandaloneHostSourceConfigNetworkConfig() *RecoverHyperVVmStandaloneHostSourceConfigNetworkConfig`

NewRecoverHyperVVmStandaloneHostSourceConfigNetworkConfig instantiates a new RecoverHyperVVmStandaloneHostSourceConfigNetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverHyperVVmStandaloneHostSourceConfigNetworkConfigWithDefaults

`func NewRecoverHyperVVmStandaloneHostSourceConfigNetworkConfigWithDefaults() *RecoverHyperVVmStandaloneHostSourceConfigNetworkConfig`

NewRecoverHyperVVmStandaloneHostSourceConfigNetworkConfigWithDefaults instantiates a new RecoverHyperVVmStandaloneHostSourceConfigNetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetachNetwork

`func (o *RecoverHyperVVmStandaloneHostSourceConfigNetworkConfig) GetDetachNetwork() bool`

GetDetachNetwork returns the DetachNetwork field if non-nil, zero value otherwise.

### GetDetachNetworkOk

`func (o *RecoverHyperVVmStandaloneHostSourceConfigNetworkConfig) GetDetachNetworkOk() (*bool, bool)`

GetDetachNetworkOk returns a tuple with the DetachNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetachNetwork

`func (o *RecoverHyperVVmStandaloneHostSourceConfigNetworkConfig) SetDetachNetwork(v bool)`

SetDetachNetwork sets DetachNetwork field to given value.

### HasDetachNetwork

`func (o *RecoverHyperVVmStandaloneHostSourceConfigNetworkConfig) HasDetachNetwork() bool`

HasDetachNetwork returns a boolean if a field has been set.

### SetDetachNetworkNil

`func (o *RecoverHyperVVmStandaloneHostSourceConfigNetworkConfig) SetDetachNetworkNil(b bool)`

 SetDetachNetworkNil sets the value for DetachNetwork to be an explicit nil

### UnsetDetachNetwork
`func (o *RecoverHyperVVmStandaloneHostSourceConfigNetworkConfig) UnsetDetachNetwork()`

UnsetDetachNetwork ensures that no value is present for DetachNetwork, not even an explicit nil
### GetVirtualSwitch

`func (o *RecoverHyperVVmStandaloneHostSourceConfigNetworkConfig) GetVirtualSwitch() RecoverHyperVVmNewSourceNetworkConfigVirtualSwitch`

GetVirtualSwitch returns the VirtualSwitch field if non-nil, zero value otherwise.

### GetVirtualSwitchOk

`func (o *RecoverHyperVVmStandaloneHostSourceConfigNetworkConfig) GetVirtualSwitchOk() (*RecoverHyperVVmNewSourceNetworkConfigVirtualSwitch, bool)`

GetVirtualSwitchOk returns a tuple with the VirtualSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualSwitch

`func (o *RecoverHyperVVmStandaloneHostSourceConfigNetworkConfig) SetVirtualSwitch(v RecoverHyperVVmNewSourceNetworkConfigVirtualSwitch)`

SetVirtualSwitch sets VirtualSwitch field to given value.

### HasVirtualSwitch

`func (o *RecoverHyperVVmStandaloneHostSourceConfigNetworkConfig) HasVirtualSwitch() bool`

HasVirtualSwitch returns a boolean if a field has been set.

### SetVirtualSwitchNil

`func (o *RecoverHyperVVmStandaloneHostSourceConfigNetworkConfig) SetVirtualSwitchNil(b bool)`

 SetVirtualSwitchNil sets the value for VirtualSwitch to be an explicit nil

### UnsetVirtualSwitch
`func (o *RecoverHyperVVmStandaloneHostSourceConfigNetworkConfig) UnsetVirtualSwitch()`

UnsetVirtualSwitch ensures that no value is present for VirtualSwitch, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


