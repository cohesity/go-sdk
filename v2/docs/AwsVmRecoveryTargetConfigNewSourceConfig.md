# AwsVmRecoveryTargetConfigNewSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptionConfig** | Pointer to [**NullableRecoverAwsVmNewSourceConfigEncryptionConfig**](RecoverAwsVmNewSourceConfigEncryptionConfig.md) |  | [optional] 
**KeyPair** | Pointer to [**NullableRecoverAwsVmNewSourceConfigKeyPair**](RecoverAwsVmNewSourceConfigKeyPair.md) |  | [optional] 
**NetworkConfig** | [**NullableRecoverAwsVmNewSourceConfigNetworkConfig**](RecoverAwsVmNewSourceConfigNetworkConfig.md) |  | 
**Region** | [**NullableRecoverAwsVmNewSourceConfigRegion**](RecoverAwsVmNewSourceConfigRegion.md) |  | 
**Source** | [**NullableRecoverAcropolisVmNewSourceConfigSource**](RecoverAcropolisVmNewSourceConfigSource.md) |  | 

## Methods

### NewAwsVmRecoveryTargetConfigNewSourceConfig

`func NewAwsVmRecoveryTargetConfigNewSourceConfig(networkConfig NullableRecoverAwsVmNewSourceConfigNetworkConfig, region NullableRecoverAwsVmNewSourceConfigRegion, source NullableRecoverAcropolisVmNewSourceConfigSource, ) *AwsVmRecoveryTargetConfigNewSourceConfig`

NewAwsVmRecoveryTargetConfigNewSourceConfig instantiates a new AwsVmRecoveryTargetConfigNewSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsVmRecoveryTargetConfigNewSourceConfigWithDefaults

`func NewAwsVmRecoveryTargetConfigNewSourceConfigWithDefaults() *AwsVmRecoveryTargetConfigNewSourceConfig`

NewAwsVmRecoveryTargetConfigNewSourceConfigWithDefaults instantiates a new AwsVmRecoveryTargetConfigNewSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptionConfig

`func (o *AwsVmRecoveryTargetConfigNewSourceConfig) GetEncryptionConfig() RecoverAwsVmNewSourceConfigEncryptionConfig`

GetEncryptionConfig returns the EncryptionConfig field if non-nil, zero value otherwise.

### GetEncryptionConfigOk

`func (o *AwsVmRecoveryTargetConfigNewSourceConfig) GetEncryptionConfigOk() (*RecoverAwsVmNewSourceConfigEncryptionConfig, bool)`

GetEncryptionConfigOk returns a tuple with the EncryptionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionConfig

`func (o *AwsVmRecoveryTargetConfigNewSourceConfig) SetEncryptionConfig(v RecoverAwsVmNewSourceConfigEncryptionConfig)`

SetEncryptionConfig sets EncryptionConfig field to given value.

### HasEncryptionConfig

`func (o *AwsVmRecoveryTargetConfigNewSourceConfig) HasEncryptionConfig() bool`

HasEncryptionConfig returns a boolean if a field has been set.

### SetEncryptionConfigNil

`func (o *AwsVmRecoveryTargetConfigNewSourceConfig) SetEncryptionConfigNil(b bool)`

 SetEncryptionConfigNil sets the value for EncryptionConfig to be an explicit nil

### UnsetEncryptionConfig
`func (o *AwsVmRecoveryTargetConfigNewSourceConfig) UnsetEncryptionConfig()`

UnsetEncryptionConfig ensures that no value is present for EncryptionConfig, not even an explicit nil
### GetKeyPair

`func (o *AwsVmRecoveryTargetConfigNewSourceConfig) GetKeyPair() RecoverAwsVmNewSourceConfigKeyPair`

GetKeyPair returns the KeyPair field if non-nil, zero value otherwise.

### GetKeyPairOk

`func (o *AwsVmRecoveryTargetConfigNewSourceConfig) GetKeyPairOk() (*RecoverAwsVmNewSourceConfigKeyPair, bool)`

GetKeyPairOk returns a tuple with the KeyPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPair

`func (o *AwsVmRecoveryTargetConfigNewSourceConfig) SetKeyPair(v RecoverAwsVmNewSourceConfigKeyPair)`

SetKeyPair sets KeyPair field to given value.

### HasKeyPair

`func (o *AwsVmRecoveryTargetConfigNewSourceConfig) HasKeyPair() bool`

HasKeyPair returns a boolean if a field has been set.

### SetKeyPairNil

`func (o *AwsVmRecoveryTargetConfigNewSourceConfig) SetKeyPairNil(b bool)`

 SetKeyPairNil sets the value for KeyPair to be an explicit nil

### UnsetKeyPair
`func (o *AwsVmRecoveryTargetConfigNewSourceConfig) UnsetKeyPair()`

UnsetKeyPair ensures that no value is present for KeyPair, not even an explicit nil
### GetNetworkConfig

`func (o *AwsVmRecoveryTargetConfigNewSourceConfig) GetNetworkConfig() RecoverAwsVmNewSourceConfigNetworkConfig`

GetNetworkConfig returns the NetworkConfig field if non-nil, zero value otherwise.

### GetNetworkConfigOk

`func (o *AwsVmRecoveryTargetConfigNewSourceConfig) GetNetworkConfigOk() (*RecoverAwsVmNewSourceConfigNetworkConfig, bool)`

GetNetworkConfigOk returns a tuple with the NetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfig

`func (o *AwsVmRecoveryTargetConfigNewSourceConfig) SetNetworkConfig(v RecoverAwsVmNewSourceConfigNetworkConfig)`

SetNetworkConfig sets NetworkConfig field to given value.


### SetNetworkConfigNil

`func (o *AwsVmRecoveryTargetConfigNewSourceConfig) SetNetworkConfigNil(b bool)`

 SetNetworkConfigNil sets the value for NetworkConfig to be an explicit nil

### UnsetNetworkConfig
`func (o *AwsVmRecoveryTargetConfigNewSourceConfig) UnsetNetworkConfig()`

UnsetNetworkConfig ensures that no value is present for NetworkConfig, not even an explicit nil
### GetRegion

`func (o *AwsVmRecoveryTargetConfigNewSourceConfig) GetRegion() RecoverAwsVmNewSourceConfigRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AwsVmRecoveryTargetConfigNewSourceConfig) GetRegionOk() (*RecoverAwsVmNewSourceConfigRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AwsVmRecoveryTargetConfigNewSourceConfig) SetRegion(v RecoverAwsVmNewSourceConfigRegion)`

SetRegion sets Region field to given value.


### SetRegionNil

`func (o *AwsVmRecoveryTargetConfigNewSourceConfig) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *AwsVmRecoveryTargetConfigNewSourceConfig) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetSource

`func (o *AwsVmRecoveryTargetConfigNewSourceConfig) GetSource() RecoverAcropolisVmNewSourceConfigSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AwsVmRecoveryTargetConfigNewSourceConfig) GetSourceOk() (*RecoverAcropolisVmNewSourceConfigSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AwsVmRecoveryTargetConfigNewSourceConfig) SetSource(v RecoverAcropolisVmNewSourceConfigSource)`

SetSource sets Source field to given value.


### SetSourceNil

`func (o *AwsVmRecoveryTargetConfigNewSourceConfig) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *AwsVmRecoveryTargetConfigNewSourceConfig) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


