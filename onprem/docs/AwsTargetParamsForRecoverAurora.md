# AwsTargetParamsForRecoverAurora

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoveryTargetConfig** | Pointer to [**NullableAwsAuroraRecoveryTargetConfig**](AwsAuroraRecoveryTargetConfig.md) | Specifies the recovery target configuration if recovery has to be done to a different location which is different from original source or to original Source with different configuration. If not specified, then the recovery of the vms will be performed to original location with all configuration parameters retained. | [optional] 
**AuroraConfig** | Pointer to [**NullableAuroraConfig**](AuroraConfig.md) | Specifies the Aurora params. | [optional] 

## Methods

### NewAwsTargetParamsForRecoverAurora

`func NewAwsTargetParamsForRecoverAurora() *AwsTargetParamsForRecoverAurora`

NewAwsTargetParamsForRecoverAurora instantiates a new AwsTargetParamsForRecoverAurora object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsTargetParamsForRecoverAuroraWithDefaults

`func NewAwsTargetParamsForRecoverAuroraWithDefaults() *AwsTargetParamsForRecoverAurora`

NewAwsTargetParamsForRecoverAuroraWithDefaults instantiates a new AwsTargetParamsForRecoverAurora object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecoveryTargetConfig

`func (o *AwsTargetParamsForRecoverAurora) GetRecoveryTargetConfig() AwsAuroraRecoveryTargetConfig`

GetRecoveryTargetConfig returns the RecoveryTargetConfig field if non-nil, zero value otherwise.

### GetRecoveryTargetConfigOk

`func (o *AwsTargetParamsForRecoverAurora) GetRecoveryTargetConfigOk() (*AwsAuroraRecoveryTargetConfig, bool)`

GetRecoveryTargetConfigOk returns a tuple with the RecoveryTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTargetConfig

`func (o *AwsTargetParamsForRecoverAurora) SetRecoveryTargetConfig(v AwsAuroraRecoveryTargetConfig)`

SetRecoveryTargetConfig sets RecoveryTargetConfig field to given value.

### HasRecoveryTargetConfig

`func (o *AwsTargetParamsForRecoverAurora) HasRecoveryTargetConfig() bool`

HasRecoveryTargetConfig returns a boolean if a field has been set.

### SetRecoveryTargetConfigNil

`func (o *AwsTargetParamsForRecoverAurora) SetRecoveryTargetConfigNil(b bool)`

 SetRecoveryTargetConfigNil sets the value for RecoveryTargetConfig to be an explicit nil

### UnsetRecoveryTargetConfig
`func (o *AwsTargetParamsForRecoverAurora) UnsetRecoveryTargetConfig()`

UnsetRecoveryTargetConfig ensures that no value is present for RecoveryTargetConfig, not even an explicit nil
### GetAuroraConfig

`func (o *AwsTargetParamsForRecoverAurora) GetAuroraConfig() AuroraConfig`

GetAuroraConfig returns the AuroraConfig field if non-nil, zero value otherwise.

### GetAuroraConfigOk

`func (o *AwsTargetParamsForRecoverAurora) GetAuroraConfigOk() (*AuroraConfig, bool)`

GetAuroraConfigOk returns a tuple with the AuroraConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuroraConfig

`func (o *AwsTargetParamsForRecoverAurora) SetAuroraConfig(v AuroraConfig)`

SetAuroraConfig sets AuroraConfig field to given value.

### HasAuroraConfig

`func (o *AwsTargetParamsForRecoverAurora) HasAuroraConfig() bool`

HasAuroraConfig returns a boolean if a field has been set.

### SetAuroraConfigNil

`func (o *AwsTargetParamsForRecoverAurora) SetAuroraConfigNil(b bool)`

 SetAuroraConfigNil sets the value for AuroraConfig to be an explicit nil

### UnsetAuroraConfig
`func (o *AwsTargetParamsForRecoverAurora) UnsetAuroraConfig()`

UnsetAuroraConfig ensures that no value is present for AuroraConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


