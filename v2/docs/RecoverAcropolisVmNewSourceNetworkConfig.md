# RecoverAcropolisVmNewSourceNetworkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetachNetwork** | Pointer to **NullableBool** | If this is set to true, then the network will be detached from the recovered VMs. All the other networking parameters set will be ignored if set to true. Default value is false. | [optional] 
**NetworkPortGroup** | Pointer to [**NullableRecoverAcropolisVmNewSourceNetworkConfigNetworkPortGroup**](RecoverAcropolisVmNewSourceNetworkConfigNetworkPortGroup.md) |  | [optional] 

## Methods

### NewRecoverAcropolisVmNewSourceNetworkConfig

`func NewRecoverAcropolisVmNewSourceNetworkConfig() *RecoverAcropolisVmNewSourceNetworkConfig`

NewRecoverAcropolisVmNewSourceNetworkConfig instantiates a new RecoverAcropolisVmNewSourceNetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverAcropolisVmNewSourceNetworkConfigWithDefaults

`func NewRecoverAcropolisVmNewSourceNetworkConfigWithDefaults() *RecoverAcropolisVmNewSourceNetworkConfig`

NewRecoverAcropolisVmNewSourceNetworkConfigWithDefaults instantiates a new RecoverAcropolisVmNewSourceNetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetachNetwork

`func (o *RecoverAcropolisVmNewSourceNetworkConfig) GetDetachNetwork() bool`

GetDetachNetwork returns the DetachNetwork field if non-nil, zero value otherwise.

### GetDetachNetworkOk

`func (o *RecoverAcropolisVmNewSourceNetworkConfig) GetDetachNetworkOk() (*bool, bool)`

GetDetachNetworkOk returns a tuple with the DetachNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetachNetwork

`func (o *RecoverAcropolisVmNewSourceNetworkConfig) SetDetachNetwork(v bool)`

SetDetachNetwork sets DetachNetwork field to given value.

### HasDetachNetwork

`func (o *RecoverAcropolisVmNewSourceNetworkConfig) HasDetachNetwork() bool`

HasDetachNetwork returns a boolean if a field has been set.

### SetDetachNetworkNil

`func (o *RecoverAcropolisVmNewSourceNetworkConfig) SetDetachNetworkNil(b bool)`

 SetDetachNetworkNil sets the value for DetachNetwork to be an explicit nil

### UnsetDetachNetwork
`func (o *RecoverAcropolisVmNewSourceNetworkConfig) UnsetDetachNetwork()`

UnsetDetachNetwork ensures that no value is present for DetachNetwork, not even an explicit nil
### GetNetworkPortGroup

`func (o *RecoverAcropolisVmNewSourceNetworkConfig) GetNetworkPortGroup() RecoverAcropolisVmNewSourceNetworkConfigNetworkPortGroup`

GetNetworkPortGroup returns the NetworkPortGroup field if non-nil, zero value otherwise.

### GetNetworkPortGroupOk

`func (o *RecoverAcropolisVmNewSourceNetworkConfig) GetNetworkPortGroupOk() (*RecoverAcropolisVmNewSourceNetworkConfigNetworkPortGroup, bool)`

GetNetworkPortGroupOk returns a tuple with the NetworkPortGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPortGroup

`func (o *RecoverAcropolisVmNewSourceNetworkConfig) SetNetworkPortGroup(v RecoverAcropolisVmNewSourceNetworkConfigNetworkPortGroup)`

SetNetworkPortGroup sets NetworkPortGroup field to given value.

### HasNetworkPortGroup

`func (o *RecoverAcropolisVmNewSourceNetworkConfig) HasNetworkPortGroup() bool`

HasNetworkPortGroup returns a boolean if a field has been set.

### SetNetworkPortGroupNil

`func (o *RecoverAcropolisVmNewSourceNetworkConfig) SetNetworkPortGroupNil(b bool)`

 SetNetworkPortGroupNil sets the value for NetworkPortGroup to be an explicit nil

### UnsetNetworkPortGroup
`func (o *RecoverAcropolisVmNewSourceNetworkConfig) UnsetNetworkPortGroup()`

UnsetNetworkPortGroup ensures that no value is present for NetworkPortGroup, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


