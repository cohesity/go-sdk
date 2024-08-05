# AwsTargetParamsForRecoverRdsRecoveryTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewSourceConfig** | Pointer to [**NullableAwsRdsRecoveryTargetConfigNewSourceConfig**](AwsRdsRecoveryTargetConfigNewSourceConfig.md) |  | [optional] 
**RecoverToNewSource** | **bool** | Specifies the parameter whether the recovery should be performed to a new or an existing Source Target. | 

## Methods

### NewAwsTargetParamsForRecoverRdsRecoveryTargetConfig

`func NewAwsTargetParamsForRecoverRdsRecoveryTargetConfig(recoverToNewSource bool, ) *AwsTargetParamsForRecoverRdsRecoveryTargetConfig`

NewAwsTargetParamsForRecoverRdsRecoveryTargetConfig instantiates a new AwsTargetParamsForRecoverRdsRecoveryTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsTargetParamsForRecoverRdsRecoveryTargetConfigWithDefaults

`func NewAwsTargetParamsForRecoverRdsRecoveryTargetConfigWithDefaults() *AwsTargetParamsForRecoverRdsRecoveryTargetConfig`

NewAwsTargetParamsForRecoverRdsRecoveryTargetConfigWithDefaults instantiates a new AwsTargetParamsForRecoverRdsRecoveryTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewSourceConfig

`func (o *AwsTargetParamsForRecoverRdsRecoveryTargetConfig) GetNewSourceConfig() AwsRdsRecoveryTargetConfigNewSourceConfig`

GetNewSourceConfig returns the NewSourceConfig field if non-nil, zero value otherwise.

### GetNewSourceConfigOk

`func (o *AwsTargetParamsForRecoverRdsRecoveryTargetConfig) GetNewSourceConfigOk() (*AwsRdsRecoveryTargetConfigNewSourceConfig, bool)`

GetNewSourceConfigOk returns a tuple with the NewSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSourceConfig

`func (o *AwsTargetParamsForRecoverRdsRecoveryTargetConfig) SetNewSourceConfig(v AwsRdsRecoveryTargetConfigNewSourceConfig)`

SetNewSourceConfig sets NewSourceConfig field to given value.

### HasNewSourceConfig

`func (o *AwsTargetParamsForRecoverRdsRecoveryTargetConfig) HasNewSourceConfig() bool`

HasNewSourceConfig returns a boolean if a field has been set.

### SetNewSourceConfigNil

`func (o *AwsTargetParamsForRecoverRdsRecoveryTargetConfig) SetNewSourceConfigNil(b bool)`

 SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil

### UnsetNewSourceConfig
`func (o *AwsTargetParamsForRecoverRdsRecoveryTargetConfig) UnsetNewSourceConfig()`

UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
### GetRecoverToNewSource

`func (o *AwsTargetParamsForRecoverRdsRecoveryTargetConfig) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *AwsTargetParamsForRecoverRdsRecoveryTargetConfig) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *AwsTargetParamsForRecoverRdsRecoveryTargetConfig) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


