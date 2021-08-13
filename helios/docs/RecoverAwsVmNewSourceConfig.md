# RecoverAwsVmNewSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the id of the parent source to recover the VMs. | 
**Region** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the AWS region in which to deploy the VM. | 
**KeyPair** | Pointer to [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the pair of public and private key used to login into the VM | [optional] 
**NetworkConfig** | [**NullableRecoverAwsVmNewSourceNetworkConfig**](RecoverAwsVmNewSourceNetworkConfig.md) | Specifies the networking configuration to be applied to the recovered VMs. | 

## Methods

### NewRecoverAwsVmNewSourceConfig

`func NewRecoverAwsVmNewSourceConfig(source NullableRecoveryObjectIdentifier, region NullableRecoveryObjectIdentifier, networkConfig NullableRecoverAwsVmNewSourceNetworkConfig, ) *RecoverAwsVmNewSourceConfig`

NewRecoverAwsVmNewSourceConfig instantiates a new RecoverAwsVmNewSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverAwsVmNewSourceConfigWithDefaults

`func NewRecoverAwsVmNewSourceConfigWithDefaults() *RecoverAwsVmNewSourceConfig`

NewRecoverAwsVmNewSourceConfigWithDefaults instantiates a new RecoverAwsVmNewSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *RecoverAwsVmNewSourceConfig) GetSource() RecoveryObjectIdentifier`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RecoverAwsVmNewSourceConfig) GetSourceOk() (*RecoveryObjectIdentifier, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RecoverAwsVmNewSourceConfig) SetSource(v RecoveryObjectIdentifier)`

SetSource sets Source field to given value.


### SetSourceNil

`func (o *RecoverAwsVmNewSourceConfig) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *RecoverAwsVmNewSourceConfig) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetRegion

`func (o *RecoverAwsVmNewSourceConfig) GetRegion() RecoveryObjectIdentifier`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *RecoverAwsVmNewSourceConfig) GetRegionOk() (*RecoveryObjectIdentifier, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *RecoverAwsVmNewSourceConfig) SetRegion(v RecoveryObjectIdentifier)`

SetRegion sets Region field to given value.


### SetRegionNil

`func (o *RecoverAwsVmNewSourceConfig) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *RecoverAwsVmNewSourceConfig) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetKeyPair

`func (o *RecoverAwsVmNewSourceConfig) GetKeyPair() RecoveryObjectIdentifier`

GetKeyPair returns the KeyPair field if non-nil, zero value otherwise.

### GetKeyPairOk

`func (o *RecoverAwsVmNewSourceConfig) GetKeyPairOk() (*RecoveryObjectIdentifier, bool)`

GetKeyPairOk returns a tuple with the KeyPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPair

`func (o *RecoverAwsVmNewSourceConfig) SetKeyPair(v RecoveryObjectIdentifier)`

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

`func (o *RecoverAwsVmNewSourceConfig) GetNetworkConfig() RecoverAwsVmNewSourceNetworkConfig`

GetNetworkConfig returns the NetworkConfig field if non-nil, zero value otherwise.

### GetNetworkConfigOk

`func (o *RecoverAwsVmNewSourceConfig) GetNetworkConfigOk() (*RecoverAwsVmNewSourceNetworkConfig, bool)`

GetNetworkConfigOk returns a tuple with the NetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfig

`func (o *RecoverAwsVmNewSourceConfig) SetNetworkConfig(v RecoverAwsVmNewSourceNetworkConfig)`

SetNetworkConfig sets NetworkConfig field to given value.


### SetNetworkConfigNil

`func (o *RecoverAwsVmNewSourceConfig) SetNetworkConfigNil(b bool)`

 SetNetworkConfigNil sets the value for NetworkConfig to be an explicit nil

### UnsetNetworkConfig
`func (o *RecoverAwsVmNewSourceConfig) UnsetNetworkConfig()`

UnsetNetworkConfig ensures that no value is present for NetworkConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


