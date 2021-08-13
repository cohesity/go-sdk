# RecoverKvmVmNewSourceNetworkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetachNetwork** | Pointer to **NullableBool** | If this is set to true, then the network will be detached from the recovered VMs. All the other networking parameters set will be ignored if set to true. Default value is false. | [optional] 
**NewNetworkConfig** | Pointer to [**RecoverKvmVmNewNetworkConfig**](RecoverKvmVmNewNetworkConfig.md) | Specifies the new network configuration of the Kvm recovery. | [optional] 

## Methods

### NewRecoverKvmVmNewSourceNetworkConfig

`func NewRecoverKvmVmNewSourceNetworkConfig() *RecoverKvmVmNewSourceNetworkConfig`

NewRecoverKvmVmNewSourceNetworkConfig instantiates a new RecoverKvmVmNewSourceNetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverKvmVmNewSourceNetworkConfigWithDefaults

`func NewRecoverKvmVmNewSourceNetworkConfigWithDefaults() *RecoverKvmVmNewSourceNetworkConfig`

NewRecoverKvmVmNewSourceNetworkConfigWithDefaults instantiates a new RecoverKvmVmNewSourceNetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetachNetwork

`func (o *RecoverKvmVmNewSourceNetworkConfig) GetDetachNetwork() bool`

GetDetachNetwork returns the DetachNetwork field if non-nil, zero value otherwise.

### GetDetachNetworkOk

`func (o *RecoverKvmVmNewSourceNetworkConfig) GetDetachNetworkOk() (*bool, bool)`

GetDetachNetworkOk returns a tuple with the DetachNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetachNetwork

`func (o *RecoverKvmVmNewSourceNetworkConfig) SetDetachNetwork(v bool)`

SetDetachNetwork sets DetachNetwork field to given value.

### HasDetachNetwork

`func (o *RecoverKvmVmNewSourceNetworkConfig) HasDetachNetwork() bool`

HasDetachNetwork returns a boolean if a field has been set.

### SetDetachNetworkNil

`func (o *RecoverKvmVmNewSourceNetworkConfig) SetDetachNetworkNil(b bool)`

 SetDetachNetworkNil sets the value for DetachNetwork to be an explicit nil

### UnsetDetachNetwork
`func (o *RecoverKvmVmNewSourceNetworkConfig) UnsetDetachNetwork()`

UnsetDetachNetwork ensures that no value is present for DetachNetwork, not even an explicit nil
### GetNewNetworkConfig

`func (o *RecoverKvmVmNewSourceNetworkConfig) GetNewNetworkConfig() RecoverKvmVmNewNetworkConfig`

GetNewNetworkConfig returns the NewNetworkConfig field if non-nil, zero value otherwise.

### GetNewNetworkConfigOk

`func (o *RecoverKvmVmNewSourceNetworkConfig) GetNewNetworkConfigOk() (*RecoverKvmVmNewNetworkConfig, bool)`

GetNewNetworkConfigOk returns a tuple with the NewNetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewNetworkConfig

`func (o *RecoverKvmVmNewSourceNetworkConfig) SetNewNetworkConfig(v RecoverKvmVmNewNetworkConfig)`

SetNewNetworkConfig sets NewNetworkConfig field to given value.

### HasNewNetworkConfig

`func (o *RecoverKvmVmNewSourceNetworkConfig) HasNewNetworkConfig() bool`

HasNewNetworkConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


