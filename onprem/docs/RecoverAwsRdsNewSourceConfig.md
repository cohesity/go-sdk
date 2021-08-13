# RecoverAwsRdsNewSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the id of the parent source to recover the RDS. | 
**Region** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the AWS region in which to deploy the RDS instance. | 
**NetworkConfig** | [**NullableRecoverAwsRdsNewSourceNetworkConfig**](RecoverAwsRdsNewSourceNetworkConfig.md) | Specifies the networking configuration to be applied to the recovered VMs. | 

## Methods

### NewRecoverAwsRdsNewSourceConfig

`func NewRecoverAwsRdsNewSourceConfig(source NullableRecoveryObjectIdentifier, region NullableRecoveryObjectIdentifier, networkConfig NullableRecoverAwsRdsNewSourceNetworkConfig, ) *RecoverAwsRdsNewSourceConfig`

NewRecoverAwsRdsNewSourceConfig instantiates a new RecoverAwsRdsNewSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverAwsRdsNewSourceConfigWithDefaults

`func NewRecoverAwsRdsNewSourceConfigWithDefaults() *RecoverAwsRdsNewSourceConfig`

NewRecoverAwsRdsNewSourceConfigWithDefaults instantiates a new RecoverAwsRdsNewSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *RecoverAwsRdsNewSourceConfig) GetSource() RecoveryObjectIdentifier`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RecoverAwsRdsNewSourceConfig) GetSourceOk() (*RecoveryObjectIdentifier, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RecoverAwsRdsNewSourceConfig) SetSource(v RecoveryObjectIdentifier)`

SetSource sets Source field to given value.


### SetSourceNil

`func (o *RecoverAwsRdsNewSourceConfig) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *RecoverAwsRdsNewSourceConfig) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetRegion

`func (o *RecoverAwsRdsNewSourceConfig) GetRegion() RecoveryObjectIdentifier`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *RecoverAwsRdsNewSourceConfig) GetRegionOk() (*RecoveryObjectIdentifier, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *RecoverAwsRdsNewSourceConfig) SetRegion(v RecoveryObjectIdentifier)`

SetRegion sets Region field to given value.


### SetRegionNil

`func (o *RecoverAwsRdsNewSourceConfig) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *RecoverAwsRdsNewSourceConfig) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetNetworkConfig

`func (o *RecoverAwsRdsNewSourceConfig) GetNetworkConfig() RecoverAwsRdsNewSourceNetworkConfig`

GetNetworkConfig returns the NetworkConfig field if non-nil, zero value otherwise.

### GetNetworkConfigOk

`func (o *RecoverAwsRdsNewSourceConfig) GetNetworkConfigOk() (*RecoverAwsRdsNewSourceNetworkConfig, bool)`

GetNetworkConfigOk returns a tuple with the NetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfig

`func (o *RecoverAwsRdsNewSourceConfig) SetNetworkConfig(v RecoverAwsRdsNewSourceNetworkConfig)`

SetNetworkConfig sets NetworkConfig field to given value.


### SetNetworkConfigNil

`func (o *RecoverAwsRdsNewSourceConfig) SetNetworkConfigNil(b bool)`

 SetNetworkConfigNil sets the value for NetworkConfig to be an explicit nil

### UnsetNetworkConfig
`func (o *RecoverAwsRdsNewSourceConfig) UnsetNetworkConfig()`

UnsetNetworkConfig ensures that no value is present for NetworkConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


