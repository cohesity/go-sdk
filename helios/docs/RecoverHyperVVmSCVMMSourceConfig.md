# RecoverHyperVVmSCVMMSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the id of the parent source to recover the VMs. | 
**Host** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the HyperV host where the recovered VMs will be attached. For standalone host targets, the host must be the same as the source. | 
**Volume** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the datastore object where the VMs&#39; files should be recovered to. | 
**NetworkConfig** | Pointer to [**NullableRecoverHyperVVmNewSourceNetworkConfig**](RecoverHyperVVmNewSourceNetworkConfig.md) | Specifies the networking configuration to be applied to the recovered VMs. | [optional] 

## Methods

### NewRecoverHyperVVmSCVMMSourceConfig

`func NewRecoverHyperVVmSCVMMSourceConfig(source NullableRecoveryObjectIdentifier, host NullableRecoveryObjectIdentifier, volume NullableRecoveryObjectIdentifier, ) *RecoverHyperVVmSCVMMSourceConfig`

NewRecoverHyperVVmSCVMMSourceConfig instantiates a new RecoverHyperVVmSCVMMSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverHyperVVmSCVMMSourceConfigWithDefaults

`func NewRecoverHyperVVmSCVMMSourceConfigWithDefaults() *RecoverHyperVVmSCVMMSourceConfig`

NewRecoverHyperVVmSCVMMSourceConfigWithDefaults instantiates a new RecoverHyperVVmSCVMMSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *RecoverHyperVVmSCVMMSourceConfig) GetSource() RecoveryObjectIdentifier`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RecoverHyperVVmSCVMMSourceConfig) GetSourceOk() (*RecoveryObjectIdentifier, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RecoverHyperVVmSCVMMSourceConfig) SetSource(v RecoveryObjectIdentifier)`

SetSource sets Source field to given value.


### SetSourceNil

`func (o *RecoverHyperVVmSCVMMSourceConfig) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *RecoverHyperVVmSCVMMSourceConfig) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetHost

`func (o *RecoverHyperVVmSCVMMSourceConfig) GetHost() RecoveryObjectIdentifier`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RecoverHyperVVmSCVMMSourceConfig) GetHostOk() (*RecoveryObjectIdentifier, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RecoverHyperVVmSCVMMSourceConfig) SetHost(v RecoveryObjectIdentifier)`

SetHost sets Host field to given value.


### SetHostNil

`func (o *RecoverHyperVVmSCVMMSourceConfig) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *RecoverHyperVVmSCVMMSourceConfig) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetVolume

`func (o *RecoverHyperVVmSCVMMSourceConfig) GetVolume() RecoveryObjectIdentifier`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *RecoverHyperVVmSCVMMSourceConfig) GetVolumeOk() (*RecoveryObjectIdentifier, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *RecoverHyperVVmSCVMMSourceConfig) SetVolume(v RecoveryObjectIdentifier)`

SetVolume sets Volume field to given value.


### SetVolumeNil

`func (o *RecoverHyperVVmSCVMMSourceConfig) SetVolumeNil(b bool)`

 SetVolumeNil sets the value for Volume to be an explicit nil

### UnsetVolume
`func (o *RecoverHyperVVmSCVMMSourceConfig) UnsetVolume()`

UnsetVolume ensures that no value is present for Volume, not even an explicit nil
### GetNetworkConfig

`func (o *RecoverHyperVVmSCVMMSourceConfig) GetNetworkConfig() RecoverHyperVVmNewSourceNetworkConfig`

GetNetworkConfig returns the NetworkConfig field if non-nil, zero value otherwise.

### GetNetworkConfigOk

`func (o *RecoverHyperVVmSCVMMSourceConfig) GetNetworkConfigOk() (*RecoverHyperVVmNewSourceNetworkConfig, bool)`

GetNetworkConfigOk returns a tuple with the NetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfig

`func (o *RecoverHyperVVmSCVMMSourceConfig) SetNetworkConfig(v RecoverHyperVVmNewSourceNetworkConfig)`

SetNetworkConfig sets NetworkConfig field to given value.

### HasNetworkConfig

`func (o *RecoverHyperVVmSCVMMSourceConfig) HasNetworkConfig() bool`

HasNetworkConfig returns a boolean if a field has been set.

### SetNetworkConfigNil

`func (o *RecoverHyperVVmSCVMMSourceConfig) SetNetworkConfigNil(b bool)`

 SetNetworkConfigNil sets the value for NetworkConfig to be an explicit nil

### UnsetNetworkConfig
`func (o *RecoverHyperVVmSCVMMSourceConfig) UnsetNetworkConfig()`

UnsetNetworkConfig ensures that no value is present for NetworkConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


