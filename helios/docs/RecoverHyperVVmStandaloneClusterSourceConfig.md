# RecoverHyperVVmStandaloneClusterSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the id of the parent source to recover the VMs. | 
**Host** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the HyperV host where the recovered VMs will be attached. For standalone host targets, the host must be the same as the source. | 
**Volume** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the datastore object where the VMs&#39; files should be recovered to. | 
**NetworkConfig** | Pointer to [**NullableRecoverHyperVVmNewSourceNetworkConfig**](RecoverHyperVVmNewSourceNetworkConfig.md) | Specifies the networking configuration to be applied to the recovered VMs. | [optional] 

## Methods

### NewRecoverHyperVVmStandaloneClusterSourceConfig

`func NewRecoverHyperVVmStandaloneClusterSourceConfig(source NullableRecoveryObjectIdentifier, host NullableRecoveryObjectIdentifier, volume NullableRecoveryObjectIdentifier, ) *RecoverHyperVVmStandaloneClusterSourceConfig`

NewRecoverHyperVVmStandaloneClusterSourceConfig instantiates a new RecoverHyperVVmStandaloneClusterSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverHyperVVmStandaloneClusterSourceConfigWithDefaults

`func NewRecoverHyperVVmStandaloneClusterSourceConfigWithDefaults() *RecoverHyperVVmStandaloneClusterSourceConfig`

NewRecoverHyperVVmStandaloneClusterSourceConfigWithDefaults instantiates a new RecoverHyperVVmStandaloneClusterSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *RecoverHyperVVmStandaloneClusterSourceConfig) GetSource() RecoveryObjectIdentifier`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RecoverHyperVVmStandaloneClusterSourceConfig) GetSourceOk() (*RecoveryObjectIdentifier, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RecoverHyperVVmStandaloneClusterSourceConfig) SetSource(v RecoveryObjectIdentifier)`

SetSource sets Source field to given value.


### SetSourceNil

`func (o *RecoverHyperVVmStandaloneClusterSourceConfig) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *RecoverHyperVVmStandaloneClusterSourceConfig) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetHost

`func (o *RecoverHyperVVmStandaloneClusterSourceConfig) GetHost() RecoveryObjectIdentifier`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RecoverHyperVVmStandaloneClusterSourceConfig) GetHostOk() (*RecoveryObjectIdentifier, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RecoverHyperVVmStandaloneClusterSourceConfig) SetHost(v RecoveryObjectIdentifier)`

SetHost sets Host field to given value.


### SetHostNil

`func (o *RecoverHyperVVmStandaloneClusterSourceConfig) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *RecoverHyperVVmStandaloneClusterSourceConfig) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetVolume

`func (o *RecoverHyperVVmStandaloneClusterSourceConfig) GetVolume() RecoveryObjectIdentifier`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *RecoverHyperVVmStandaloneClusterSourceConfig) GetVolumeOk() (*RecoveryObjectIdentifier, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *RecoverHyperVVmStandaloneClusterSourceConfig) SetVolume(v RecoveryObjectIdentifier)`

SetVolume sets Volume field to given value.


### SetVolumeNil

`func (o *RecoverHyperVVmStandaloneClusterSourceConfig) SetVolumeNil(b bool)`

 SetVolumeNil sets the value for Volume to be an explicit nil

### UnsetVolume
`func (o *RecoverHyperVVmStandaloneClusterSourceConfig) UnsetVolume()`

UnsetVolume ensures that no value is present for Volume, not even an explicit nil
### GetNetworkConfig

`func (o *RecoverHyperVVmStandaloneClusterSourceConfig) GetNetworkConfig() RecoverHyperVVmNewSourceNetworkConfig`

GetNetworkConfig returns the NetworkConfig field if non-nil, zero value otherwise.

### GetNetworkConfigOk

`func (o *RecoverHyperVVmStandaloneClusterSourceConfig) GetNetworkConfigOk() (*RecoverHyperVVmNewSourceNetworkConfig, bool)`

GetNetworkConfigOk returns a tuple with the NetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfig

`func (o *RecoverHyperVVmStandaloneClusterSourceConfig) SetNetworkConfig(v RecoverHyperVVmNewSourceNetworkConfig)`

SetNetworkConfig sets NetworkConfig field to given value.

### HasNetworkConfig

`func (o *RecoverHyperVVmStandaloneClusterSourceConfig) HasNetworkConfig() bool`

HasNetworkConfig returns a boolean if a field has been set.

### SetNetworkConfigNil

`func (o *RecoverHyperVVmStandaloneClusterSourceConfig) SetNetworkConfigNil(b bool)`

 SetNetworkConfigNil sets the value for NetworkConfig to be an explicit nil

### UnsetNetworkConfig
`func (o *RecoverHyperVVmStandaloneClusterSourceConfig) UnsetNetworkConfig()`

UnsetNetworkConfig ensures that no value is present for NetworkConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


