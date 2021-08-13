# RecoverVmwareVAppOriginalSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkConfig** | Pointer to [**NullableRecoverVmwareVmOriginalSourceNetworkConfig**](RecoverVmwareVmOriginalSourceNetworkConfig.md) | Specifies the networking configuration to be applied to the recovered vApps. | [optional] 

## Methods

### NewRecoverVmwareVAppOriginalSourceConfig

`func NewRecoverVmwareVAppOriginalSourceConfig() *RecoverVmwareVAppOriginalSourceConfig`

NewRecoverVmwareVAppOriginalSourceConfig instantiates a new RecoverVmwareVAppOriginalSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverVmwareVAppOriginalSourceConfigWithDefaults

`func NewRecoverVmwareVAppOriginalSourceConfigWithDefaults() *RecoverVmwareVAppOriginalSourceConfig`

NewRecoverVmwareVAppOriginalSourceConfigWithDefaults instantiates a new RecoverVmwareVAppOriginalSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkConfig

`func (o *RecoverVmwareVAppOriginalSourceConfig) GetNetworkConfig() RecoverVmwareVmOriginalSourceNetworkConfig`

GetNetworkConfig returns the NetworkConfig field if non-nil, zero value otherwise.

### GetNetworkConfigOk

`func (o *RecoverVmwareVAppOriginalSourceConfig) GetNetworkConfigOk() (*RecoverVmwareVmOriginalSourceNetworkConfig, bool)`

GetNetworkConfigOk returns a tuple with the NetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfig

`func (o *RecoverVmwareVAppOriginalSourceConfig) SetNetworkConfig(v RecoverVmwareVmOriginalSourceNetworkConfig)`

SetNetworkConfig sets NetworkConfig field to given value.

### HasNetworkConfig

`func (o *RecoverVmwareVAppOriginalSourceConfig) HasNetworkConfig() bool`

HasNetworkConfig returns a boolean if a field has been set.

### SetNetworkConfigNil

`func (o *RecoverVmwareVAppOriginalSourceConfig) SetNetworkConfigNil(b bool)`

 SetNetworkConfigNil sets the value for NetworkConfig to be an explicit nil

### UnsetNetworkConfig
`func (o *RecoverVmwareVAppOriginalSourceConfig) UnsetNetworkConfig()`

UnsetNetworkConfig ensures that no value is present for NetworkConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


