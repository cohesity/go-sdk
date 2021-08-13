# RecoverHyperVVmOriginalSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkConfig** | Pointer to [**NullableRecoverHyperVVmOriginalSourceNetworkConfig**](RecoverHyperVVmOriginalSourceNetworkConfig.md) | Specifies the networking configuration to be applied to the recovered VMs. | [optional] 

## Methods

### NewRecoverHyperVVmOriginalSourceConfig

`func NewRecoverHyperVVmOriginalSourceConfig() *RecoverHyperVVmOriginalSourceConfig`

NewRecoverHyperVVmOriginalSourceConfig instantiates a new RecoverHyperVVmOriginalSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverHyperVVmOriginalSourceConfigWithDefaults

`func NewRecoverHyperVVmOriginalSourceConfigWithDefaults() *RecoverHyperVVmOriginalSourceConfig`

NewRecoverHyperVVmOriginalSourceConfigWithDefaults instantiates a new RecoverHyperVVmOriginalSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkConfig

`func (o *RecoverHyperVVmOriginalSourceConfig) GetNetworkConfig() RecoverHyperVVmOriginalSourceNetworkConfig`

GetNetworkConfig returns the NetworkConfig field if non-nil, zero value otherwise.

### GetNetworkConfigOk

`func (o *RecoverHyperVVmOriginalSourceConfig) GetNetworkConfigOk() (*RecoverHyperVVmOriginalSourceNetworkConfig, bool)`

GetNetworkConfigOk returns a tuple with the NetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfig

`func (o *RecoverHyperVVmOriginalSourceConfig) SetNetworkConfig(v RecoverHyperVVmOriginalSourceNetworkConfig)`

SetNetworkConfig sets NetworkConfig field to given value.

### HasNetworkConfig

`func (o *RecoverHyperVVmOriginalSourceConfig) HasNetworkConfig() bool`

HasNetworkConfig returns a boolean if a field has been set.

### SetNetworkConfigNil

`func (o *RecoverHyperVVmOriginalSourceConfig) SetNetworkConfigNil(b bool)`

 SetNetworkConfigNil sets the value for NetworkConfig to be an explicit nil

### UnsetNetworkConfig
`func (o *RecoverHyperVVmOriginalSourceConfig) UnsetNetworkConfig()`

UnsetNetworkConfig ensures that no value is present for NetworkConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


