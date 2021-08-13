# RecoverHyperVVmStandaloneHostSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the id of the parent source to recover the VMs. | 
**Volume** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the datastore object where the VMs&#39; files should be recovered to. | 
**NetworkConfig** | Pointer to [**NullableRecoverHyperVVmNewSourceNetworkConfig**](RecoverHyperVVmNewSourceNetworkConfig.md) | Specifies the networking configuration to be applied to the recovered VMs. | [optional] 

## Methods

### NewRecoverHyperVVmStandaloneHostSourceConfig

`func NewRecoverHyperVVmStandaloneHostSourceConfig(source NullableRecoveryObjectIdentifier, volume NullableRecoveryObjectIdentifier, ) *RecoverHyperVVmStandaloneHostSourceConfig`

NewRecoverHyperVVmStandaloneHostSourceConfig instantiates a new RecoverHyperVVmStandaloneHostSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverHyperVVmStandaloneHostSourceConfigWithDefaults

`func NewRecoverHyperVVmStandaloneHostSourceConfigWithDefaults() *RecoverHyperVVmStandaloneHostSourceConfig`

NewRecoverHyperVVmStandaloneHostSourceConfigWithDefaults instantiates a new RecoverHyperVVmStandaloneHostSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *RecoverHyperVVmStandaloneHostSourceConfig) GetSource() RecoveryObjectIdentifier`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RecoverHyperVVmStandaloneHostSourceConfig) GetSourceOk() (*RecoveryObjectIdentifier, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RecoverHyperVVmStandaloneHostSourceConfig) SetSource(v RecoveryObjectIdentifier)`

SetSource sets Source field to given value.


### SetSourceNil

`func (o *RecoverHyperVVmStandaloneHostSourceConfig) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *RecoverHyperVVmStandaloneHostSourceConfig) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetVolume

`func (o *RecoverHyperVVmStandaloneHostSourceConfig) GetVolume() RecoveryObjectIdentifier`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *RecoverHyperVVmStandaloneHostSourceConfig) GetVolumeOk() (*RecoveryObjectIdentifier, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *RecoverHyperVVmStandaloneHostSourceConfig) SetVolume(v RecoveryObjectIdentifier)`

SetVolume sets Volume field to given value.


### SetVolumeNil

`func (o *RecoverHyperVVmStandaloneHostSourceConfig) SetVolumeNil(b bool)`

 SetVolumeNil sets the value for Volume to be an explicit nil

### UnsetVolume
`func (o *RecoverHyperVVmStandaloneHostSourceConfig) UnsetVolume()`

UnsetVolume ensures that no value is present for Volume, not even an explicit nil
### GetNetworkConfig

`func (o *RecoverHyperVVmStandaloneHostSourceConfig) GetNetworkConfig() RecoverHyperVVmNewSourceNetworkConfig`

GetNetworkConfig returns the NetworkConfig field if non-nil, zero value otherwise.

### GetNetworkConfigOk

`func (o *RecoverHyperVVmStandaloneHostSourceConfig) GetNetworkConfigOk() (*RecoverHyperVVmNewSourceNetworkConfig, bool)`

GetNetworkConfigOk returns a tuple with the NetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfig

`func (o *RecoverHyperVVmStandaloneHostSourceConfig) SetNetworkConfig(v RecoverHyperVVmNewSourceNetworkConfig)`

SetNetworkConfig sets NetworkConfig field to given value.

### HasNetworkConfig

`func (o *RecoverHyperVVmStandaloneHostSourceConfig) HasNetworkConfig() bool`

HasNetworkConfig returns a boolean if a field has been set.

### SetNetworkConfigNil

`func (o *RecoverHyperVVmStandaloneHostSourceConfig) SetNetworkConfigNil(b bool)`

 SetNetworkConfigNil sets the value for NetworkConfig to be an explicit nil

### UnsetNetworkConfig
`func (o *RecoverHyperVVmStandaloneHostSourceConfig) UnsetNetworkConfig()`

UnsetNetworkConfig ensures that no value is present for NetworkConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


