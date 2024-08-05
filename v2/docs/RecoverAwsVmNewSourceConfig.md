# RecoverAwsVmNewSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptionConfig** | Pointer to [**NullableRecoverAwsVmNewSourceConfigEncryptionConfig**](RecoverAwsVmNewSourceConfigEncryptionConfig.md) |  | [optional] 
**KeyPair** | Pointer to [**NullableRecoverAwsVmNewSourceConfigKeyPair**](RecoverAwsVmNewSourceConfigKeyPair.md) |  | [optional] 
**NetworkConfig** | [**NullableRecoverAwsVmNewSourceConfigNetworkConfig**](RecoverAwsVmNewSourceConfigNetworkConfig.md) |  | 
**Region** | [**NullableRecoverAwsVmNewSourceConfigRegion**](RecoverAwsVmNewSourceConfigRegion.md) |  | 
**Source** | [**NullableRecoverAcropolisVmNewSourceConfigSource**](RecoverAcropolisVmNewSourceConfigSource.md) |  | 

## Methods

### NewRecoverAwsVmNewSourceConfig

`func NewRecoverAwsVmNewSourceConfig(networkConfig NullableRecoverAwsVmNewSourceConfigNetworkConfig, region NullableRecoverAwsVmNewSourceConfigRegion, source NullableRecoverAcropolisVmNewSourceConfigSource, ) *RecoverAwsVmNewSourceConfig`

NewRecoverAwsVmNewSourceConfig instantiates a new RecoverAwsVmNewSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverAwsVmNewSourceConfigWithDefaults

`func NewRecoverAwsVmNewSourceConfigWithDefaults() *RecoverAwsVmNewSourceConfig`

NewRecoverAwsVmNewSourceConfigWithDefaults instantiates a new RecoverAwsVmNewSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptionConfig

`func (o *RecoverAwsVmNewSourceConfig) GetEncryptionConfig() RecoverAwsVmNewSourceConfigEncryptionConfig`

GetEncryptionConfig returns the EncryptionConfig field if non-nil, zero value otherwise.

### GetEncryptionConfigOk

`func (o *RecoverAwsVmNewSourceConfig) GetEncryptionConfigOk() (*RecoverAwsVmNewSourceConfigEncryptionConfig, bool)`

GetEncryptionConfigOk returns a tuple with the EncryptionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionConfig

`func (o *RecoverAwsVmNewSourceConfig) SetEncryptionConfig(v RecoverAwsVmNewSourceConfigEncryptionConfig)`

SetEncryptionConfig sets EncryptionConfig field to given value.

### HasEncryptionConfig

`func (o *RecoverAwsVmNewSourceConfig) HasEncryptionConfig() bool`

HasEncryptionConfig returns a boolean if a field has been set.

### SetEncryptionConfigNil

`func (o *RecoverAwsVmNewSourceConfig) SetEncryptionConfigNil(b bool)`

 SetEncryptionConfigNil sets the value for EncryptionConfig to be an explicit nil

### UnsetEncryptionConfig
`func (o *RecoverAwsVmNewSourceConfig) UnsetEncryptionConfig()`

UnsetEncryptionConfig ensures that no value is present for EncryptionConfig, not even an explicit nil
### GetKeyPair

`func (o *RecoverAwsVmNewSourceConfig) GetKeyPair() RecoverAwsVmNewSourceConfigKeyPair`

GetKeyPair returns the KeyPair field if non-nil, zero value otherwise.

### GetKeyPairOk

`func (o *RecoverAwsVmNewSourceConfig) GetKeyPairOk() (*RecoverAwsVmNewSourceConfigKeyPair, bool)`

GetKeyPairOk returns a tuple with the KeyPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPair

`func (o *RecoverAwsVmNewSourceConfig) SetKeyPair(v RecoverAwsVmNewSourceConfigKeyPair)`

SetKeyPair sets KeyPair field to given value.

### HasKeyPair

`func (o *RecoverAwsVmNewSourceConfig) HasKeyPair() bool`

HasKeyPair returns a boolean if a field has been set.

### SetKeyPairNil

`func (o *RecoverAwsVmNewSourceConfig) SetKeyPairNil(b bool)`

 SetKeyPairNil sets the value for KeyPair to be an explicit nil

### UnsetKeyPair
`func (o *RecoverAwsVmNewSourceConfig) UnsetKeyPair()`

UnsetKeyPair ensures that no value is present for KeyPair, not even an explicit nil
### GetNetworkConfig

`func (o *RecoverAwsVmNewSourceConfig) GetNetworkConfig() RecoverAwsVmNewSourceConfigNetworkConfig`

GetNetworkConfig returns the NetworkConfig field if non-nil, zero value otherwise.

### GetNetworkConfigOk

`func (o *RecoverAwsVmNewSourceConfig) GetNetworkConfigOk() (*RecoverAwsVmNewSourceConfigNetworkConfig, bool)`

GetNetworkConfigOk returns a tuple with the NetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfig

`func (o *RecoverAwsVmNewSourceConfig) SetNetworkConfig(v RecoverAwsVmNewSourceConfigNetworkConfig)`

SetNetworkConfig sets NetworkConfig field to given value.


### SetNetworkConfigNil

`func (o *RecoverAwsVmNewSourceConfig) SetNetworkConfigNil(b bool)`

 SetNetworkConfigNil sets the value for NetworkConfig to be an explicit nil

### UnsetNetworkConfig
`func (o *RecoverAwsVmNewSourceConfig) UnsetNetworkConfig()`

UnsetNetworkConfig ensures that no value is present for NetworkConfig, not even an explicit nil
### GetRegion

`func (o *RecoverAwsVmNewSourceConfig) GetRegion() RecoverAwsVmNewSourceConfigRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *RecoverAwsVmNewSourceConfig) GetRegionOk() (*RecoverAwsVmNewSourceConfigRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *RecoverAwsVmNewSourceConfig) SetRegion(v RecoverAwsVmNewSourceConfigRegion)`

SetRegion sets Region field to given value.


### SetRegionNil

`func (o *RecoverAwsVmNewSourceConfig) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *RecoverAwsVmNewSourceConfig) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetSource

`func (o *RecoverAwsVmNewSourceConfig) GetSource() RecoverAcropolisVmNewSourceConfigSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RecoverAwsVmNewSourceConfig) GetSourceOk() (*RecoverAcropolisVmNewSourceConfigSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RecoverAwsVmNewSourceConfig) SetSource(v RecoverAcropolisVmNewSourceConfigSource)`

SetSource sets Source field to given value.


### SetSourceNil

`func (o *RecoverAwsVmNewSourceConfig) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *RecoverAwsVmNewSourceConfig) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


