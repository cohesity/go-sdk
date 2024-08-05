# RecoverAcropolisVmNewSourceConfigNetworkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetachNetwork** | Pointer to **NullableBool** | If this is set to true, then the network will be detached from the recovered VMs. All the other networking parameters set will be ignored if set to true. Default value is false. | [optional] 
**NetworkPortGroup** | Pointer to [**NullableRecoverAcropolisVmNewSourceNetworkConfigNetworkPortGroup**](RecoverAcropolisVmNewSourceNetworkConfigNetworkPortGroup.md) |  | [optional] 

## Methods

### NewRecoverAcropolisVmNewSourceConfigNetworkConfig

`func NewRecoverAcropolisVmNewSourceConfigNetworkConfig() *RecoverAcropolisVmNewSourceConfigNetworkConfig`

NewRecoverAcropolisVmNewSourceConfigNetworkConfig instantiates a new RecoverAcropolisVmNewSourceConfigNetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverAcropolisVmNewSourceConfigNetworkConfigWithDefaults

`func NewRecoverAcropolisVmNewSourceConfigNetworkConfigWithDefaults() *RecoverAcropolisVmNewSourceConfigNetworkConfig`

NewRecoverAcropolisVmNewSourceConfigNetworkConfigWithDefaults instantiates a new RecoverAcropolisVmNewSourceConfigNetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetachNetwork

`func (o *RecoverAcropolisVmNewSourceConfigNetworkConfig) GetDetachNetwork() bool`

GetDetachNetwork returns the DetachNetwork field if non-nil, zero value otherwise.

### GetDetachNetworkOk

`func (o *RecoverAcropolisVmNewSourceConfigNetworkConfig) GetDetachNetworkOk() (*bool, bool)`

GetDetachNetworkOk returns a tuple with the DetachNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetachNetwork

`func (o *RecoverAcropolisVmNewSourceConfigNetworkConfig) SetDetachNetwork(v bool)`

SetDetachNetwork sets DetachNetwork field to given value.

### HasDetachNetwork

`func (o *RecoverAcropolisVmNewSourceConfigNetworkConfig) HasDetachNetwork() bool`

HasDetachNetwork returns a boolean if a field has been set.

### SetDetachNetworkNil

`func (o *RecoverAcropolisVmNewSourceConfigNetworkConfig) SetDetachNetworkNil(b bool)`

 SetDetachNetworkNil sets the value for DetachNetwork to be an explicit nil

### UnsetDetachNetwork
`func (o *RecoverAcropolisVmNewSourceConfigNetworkConfig) UnsetDetachNetwork()`

UnsetDetachNetwork ensures that no value is present for DetachNetwork, not even an explicit nil
### GetNetworkPortGroup

`func (o *RecoverAcropolisVmNewSourceConfigNetworkConfig) GetNetworkPortGroup() RecoverAcropolisVmNewSourceNetworkConfigNetworkPortGroup`

GetNetworkPortGroup returns the NetworkPortGroup field if non-nil, zero value otherwise.

### GetNetworkPortGroupOk

`func (o *RecoverAcropolisVmNewSourceConfigNetworkConfig) GetNetworkPortGroupOk() (*RecoverAcropolisVmNewSourceNetworkConfigNetworkPortGroup, bool)`

GetNetworkPortGroupOk returns a tuple with the NetworkPortGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPortGroup

`func (o *RecoverAcropolisVmNewSourceConfigNetworkConfig) SetNetworkPortGroup(v RecoverAcropolisVmNewSourceNetworkConfigNetworkPortGroup)`

SetNetworkPortGroup sets NetworkPortGroup field to given value.

### HasNetworkPortGroup

`func (o *RecoverAcropolisVmNewSourceConfigNetworkConfig) HasNetworkPortGroup() bool`

HasNetworkPortGroup returns a boolean if a field has been set.

### SetNetworkPortGroupNil

`func (o *RecoverAcropolisVmNewSourceConfigNetworkConfig) SetNetworkPortGroupNil(b bool)`

 SetNetworkPortGroupNil sets the value for NetworkPortGroup to be an explicit nil

### UnsetNetworkPortGroup
`func (o *RecoverAcropolisVmNewSourceConfigNetworkConfig) UnsetNetworkPortGroup()`

UnsetNetworkPortGroup ensures that no value is present for NetworkPortGroup, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


