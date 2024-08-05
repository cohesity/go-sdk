# RecoverHyperVVmSCVMMSourceConfigNetworkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetachNetwork** | Pointer to **NullableBool** | If this is set to true, then the network will be detached from the recovered VMs. All the other networking parameters set will be ignored if set to true. Default value is false. | [optional] 
**VirtualSwitch** | Pointer to [**NullableRecoverHyperVVmNewSourceNetworkConfigVirtualSwitch**](RecoverHyperVVmNewSourceNetworkConfigVirtualSwitch.md) |  | [optional] 

## Methods

### NewRecoverHyperVVmSCVMMSourceConfigNetworkConfig

`func NewRecoverHyperVVmSCVMMSourceConfigNetworkConfig() *RecoverHyperVVmSCVMMSourceConfigNetworkConfig`

NewRecoverHyperVVmSCVMMSourceConfigNetworkConfig instantiates a new RecoverHyperVVmSCVMMSourceConfigNetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverHyperVVmSCVMMSourceConfigNetworkConfigWithDefaults

`func NewRecoverHyperVVmSCVMMSourceConfigNetworkConfigWithDefaults() *RecoverHyperVVmSCVMMSourceConfigNetworkConfig`

NewRecoverHyperVVmSCVMMSourceConfigNetworkConfigWithDefaults instantiates a new RecoverHyperVVmSCVMMSourceConfigNetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetachNetwork

`func (o *RecoverHyperVVmSCVMMSourceConfigNetworkConfig) GetDetachNetwork() bool`

GetDetachNetwork returns the DetachNetwork field if non-nil, zero value otherwise.

### GetDetachNetworkOk

`func (o *RecoverHyperVVmSCVMMSourceConfigNetworkConfig) GetDetachNetworkOk() (*bool, bool)`

GetDetachNetworkOk returns a tuple with the DetachNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetachNetwork

`func (o *RecoverHyperVVmSCVMMSourceConfigNetworkConfig) SetDetachNetwork(v bool)`

SetDetachNetwork sets DetachNetwork field to given value.

### HasDetachNetwork

`func (o *RecoverHyperVVmSCVMMSourceConfigNetworkConfig) HasDetachNetwork() bool`

HasDetachNetwork returns a boolean if a field has been set.

### SetDetachNetworkNil

`func (o *RecoverHyperVVmSCVMMSourceConfigNetworkConfig) SetDetachNetworkNil(b bool)`

 SetDetachNetworkNil sets the value for DetachNetwork to be an explicit nil

### UnsetDetachNetwork
`func (o *RecoverHyperVVmSCVMMSourceConfigNetworkConfig) UnsetDetachNetwork()`

UnsetDetachNetwork ensures that no value is present for DetachNetwork, not even an explicit nil
### GetVirtualSwitch

`func (o *RecoverHyperVVmSCVMMSourceConfigNetworkConfig) GetVirtualSwitch() RecoverHyperVVmNewSourceNetworkConfigVirtualSwitch`

GetVirtualSwitch returns the VirtualSwitch field if non-nil, zero value otherwise.

### GetVirtualSwitchOk

`func (o *RecoverHyperVVmSCVMMSourceConfigNetworkConfig) GetVirtualSwitchOk() (*RecoverHyperVVmNewSourceNetworkConfigVirtualSwitch, bool)`

GetVirtualSwitchOk returns a tuple with the VirtualSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualSwitch

`func (o *RecoverHyperVVmSCVMMSourceConfigNetworkConfig) SetVirtualSwitch(v RecoverHyperVVmNewSourceNetworkConfigVirtualSwitch)`

SetVirtualSwitch sets VirtualSwitch field to given value.

### HasVirtualSwitch

`func (o *RecoverHyperVVmSCVMMSourceConfigNetworkConfig) HasVirtualSwitch() bool`

HasVirtualSwitch returns a boolean if a field has been set.

### SetVirtualSwitchNil

`func (o *RecoverHyperVVmSCVMMSourceConfigNetworkConfig) SetVirtualSwitchNil(b bool)`

 SetVirtualSwitchNil sets the value for VirtualSwitch to be an explicit nil

### UnsetVirtualSwitch
`func (o *RecoverHyperVVmSCVMMSourceConfigNetworkConfig) UnsetVirtualSwitch()`

UnsetVirtualSwitch ensures that no value is present for VirtualSwitch, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


