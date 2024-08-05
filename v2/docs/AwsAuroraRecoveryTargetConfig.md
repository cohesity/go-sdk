# AwsAuroraRecoveryTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewSourceConfig** | Pointer to [**NullableAwsAuroraRecoveryTargetConfigNewSourceConfig**](AwsAuroraRecoveryTargetConfigNewSourceConfig.md) |  | [optional] 
**RecoverToNewSource** | **bool** | Specifies the parameter whether the recovery should be performed to a new or an existing Source Target. | 

## Methods

### NewAwsAuroraRecoveryTargetConfig

`func NewAwsAuroraRecoveryTargetConfig(recoverToNewSource bool, ) *AwsAuroraRecoveryTargetConfig`

NewAwsAuroraRecoveryTargetConfig instantiates a new AwsAuroraRecoveryTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsAuroraRecoveryTargetConfigWithDefaults

`func NewAwsAuroraRecoveryTargetConfigWithDefaults() *AwsAuroraRecoveryTargetConfig`

NewAwsAuroraRecoveryTargetConfigWithDefaults instantiates a new AwsAuroraRecoveryTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewSourceConfig

`func (o *AwsAuroraRecoveryTargetConfig) GetNewSourceConfig() AwsAuroraRecoveryTargetConfigNewSourceConfig`

GetNewSourceConfig returns the NewSourceConfig field if non-nil, zero value otherwise.

### GetNewSourceConfigOk

`func (o *AwsAuroraRecoveryTargetConfig) GetNewSourceConfigOk() (*AwsAuroraRecoveryTargetConfigNewSourceConfig, bool)`

GetNewSourceConfigOk returns a tuple with the NewSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSourceConfig

`func (o *AwsAuroraRecoveryTargetConfig) SetNewSourceConfig(v AwsAuroraRecoveryTargetConfigNewSourceConfig)`

SetNewSourceConfig sets NewSourceConfig field to given value.

### HasNewSourceConfig

`func (o *AwsAuroraRecoveryTargetConfig) HasNewSourceConfig() bool`

HasNewSourceConfig returns a boolean if a field has been set.

### SetNewSourceConfigNil

`func (o *AwsAuroraRecoveryTargetConfig) SetNewSourceConfigNil(b bool)`

 SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil

### UnsetNewSourceConfig
`func (o *AwsAuroraRecoveryTargetConfig) UnsetNewSourceConfig()`

UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
### GetRecoverToNewSource

`func (o *AwsAuroraRecoveryTargetConfig) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *AwsAuroraRecoveryTargetConfig) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *AwsAuroraRecoveryTargetConfig) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


