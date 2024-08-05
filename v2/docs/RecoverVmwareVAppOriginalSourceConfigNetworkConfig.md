# RecoverVmwareVAppOriginalSourceConfigNetworkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetachNetwork** | Pointer to **NullableBool** | If this is set to true, then the network will be detached from the recovered VMs. All the other networking parameters set will be ignored if set to true. Default value is false. | [optional] 
**DisableNetwork** | Pointer to **NullableBool** | Specifies whether the attached network should be left in disabled state. Default is false. | [optional] 

## Methods

### NewRecoverVmwareVAppOriginalSourceConfigNetworkConfig

`func NewRecoverVmwareVAppOriginalSourceConfigNetworkConfig() *RecoverVmwareVAppOriginalSourceConfigNetworkConfig`

NewRecoverVmwareVAppOriginalSourceConfigNetworkConfig instantiates a new RecoverVmwareVAppOriginalSourceConfigNetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverVmwareVAppOriginalSourceConfigNetworkConfigWithDefaults

`func NewRecoverVmwareVAppOriginalSourceConfigNetworkConfigWithDefaults() *RecoverVmwareVAppOriginalSourceConfigNetworkConfig`

NewRecoverVmwareVAppOriginalSourceConfigNetworkConfigWithDefaults instantiates a new RecoverVmwareVAppOriginalSourceConfigNetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetachNetwork

`func (o *RecoverVmwareVAppOriginalSourceConfigNetworkConfig) GetDetachNetwork() bool`

GetDetachNetwork returns the DetachNetwork field if non-nil, zero value otherwise.

### GetDetachNetworkOk

`func (o *RecoverVmwareVAppOriginalSourceConfigNetworkConfig) GetDetachNetworkOk() (*bool, bool)`

GetDetachNetworkOk returns a tuple with the DetachNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetachNetwork

`func (o *RecoverVmwareVAppOriginalSourceConfigNetworkConfig) SetDetachNetwork(v bool)`

SetDetachNetwork sets DetachNetwork field to given value.

### HasDetachNetwork

`func (o *RecoverVmwareVAppOriginalSourceConfigNetworkConfig) HasDetachNetwork() bool`

HasDetachNetwork returns a boolean if a field has been set.

### SetDetachNetworkNil

`func (o *RecoverVmwareVAppOriginalSourceConfigNetworkConfig) SetDetachNetworkNil(b bool)`

 SetDetachNetworkNil sets the value for DetachNetwork to be an explicit nil

### UnsetDetachNetwork
`func (o *RecoverVmwareVAppOriginalSourceConfigNetworkConfig) UnsetDetachNetwork()`

UnsetDetachNetwork ensures that no value is present for DetachNetwork, not even an explicit nil
### GetDisableNetwork

`func (o *RecoverVmwareVAppOriginalSourceConfigNetworkConfig) GetDisableNetwork() bool`

GetDisableNetwork returns the DisableNetwork field if non-nil, zero value otherwise.

### GetDisableNetworkOk

`func (o *RecoverVmwareVAppOriginalSourceConfigNetworkConfig) GetDisableNetworkOk() (*bool, bool)`

GetDisableNetworkOk returns a tuple with the DisableNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNetwork

`func (o *RecoverVmwareVAppOriginalSourceConfigNetworkConfig) SetDisableNetwork(v bool)`

SetDisableNetwork sets DisableNetwork field to given value.

### HasDisableNetwork

`func (o *RecoverVmwareVAppOriginalSourceConfigNetworkConfig) HasDisableNetwork() bool`

HasDisableNetwork returns a boolean if a field has been set.

### SetDisableNetworkNil

`func (o *RecoverVmwareVAppOriginalSourceConfigNetworkConfig) SetDisableNetworkNil(b bool)`

 SetDisableNetworkNil sets the value for DisableNetwork to be an explicit nil

### UnsetDisableNetwork
`func (o *RecoverVmwareVAppOriginalSourceConfigNetworkConfig) UnsetDisableNetwork()`

UnsetDisableNetwork ensures that no value is present for DisableNetwork, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


