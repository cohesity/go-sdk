# AcropolisTargetParamsForRecoverVmRecoveryTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewSourceConfig** | Pointer to [**NullableRecoverAcropolisVmTargetConfigNewSourceConfig**](RecoverAcropolisVmTargetConfigNewSourceConfig.md) |  | [optional] 
**OriginalSourceConfig** | Pointer to [**NullableRecoverAcropolisVmTargetConfigOriginalSourceConfig**](RecoverAcropolisVmTargetConfigOriginalSourceConfig.md) |  | [optional] 
**RecoverToNewSource** | **bool** | Specifies the parameter whether the recovery should be performed to a new or an existing Source Target. | 

## Methods

### NewAcropolisTargetParamsForRecoverVmRecoveryTargetConfig

`func NewAcropolisTargetParamsForRecoverVmRecoveryTargetConfig(recoverToNewSource bool, ) *AcropolisTargetParamsForRecoverVmRecoveryTargetConfig`

NewAcropolisTargetParamsForRecoverVmRecoveryTargetConfig instantiates a new AcropolisTargetParamsForRecoverVmRecoveryTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcropolisTargetParamsForRecoverVmRecoveryTargetConfigWithDefaults

`func NewAcropolisTargetParamsForRecoverVmRecoveryTargetConfigWithDefaults() *AcropolisTargetParamsForRecoverVmRecoveryTargetConfig`

NewAcropolisTargetParamsForRecoverVmRecoveryTargetConfigWithDefaults instantiates a new AcropolisTargetParamsForRecoverVmRecoveryTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewSourceConfig

`func (o *AcropolisTargetParamsForRecoverVmRecoveryTargetConfig) GetNewSourceConfig() RecoverAcropolisVmTargetConfigNewSourceConfig`

GetNewSourceConfig returns the NewSourceConfig field if non-nil, zero value otherwise.

### GetNewSourceConfigOk

`func (o *AcropolisTargetParamsForRecoverVmRecoveryTargetConfig) GetNewSourceConfigOk() (*RecoverAcropolisVmTargetConfigNewSourceConfig, bool)`

GetNewSourceConfigOk returns a tuple with the NewSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSourceConfig

`func (o *AcropolisTargetParamsForRecoverVmRecoveryTargetConfig) SetNewSourceConfig(v RecoverAcropolisVmTargetConfigNewSourceConfig)`

SetNewSourceConfig sets NewSourceConfig field to given value.

### HasNewSourceConfig

`func (o *AcropolisTargetParamsForRecoverVmRecoveryTargetConfig) HasNewSourceConfig() bool`

HasNewSourceConfig returns a boolean if a field has been set.

### SetNewSourceConfigNil

`func (o *AcropolisTargetParamsForRecoverVmRecoveryTargetConfig) SetNewSourceConfigNil(b bool)`

 SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil

### UnsetNewSourceConfig
`func (o *AcropolisTargetParamsForRecoverVmRecoveryTargetConfig) UnsetNewSourceConfig()`

UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
### GetOriginalSourceConfig

`func (o *AcropolisTargetParamsForRecoverVmRecoveryTargetConfig) GetOriginalSourceConfig() RecoverAcropolisVmTargetConfigOriginalSourceConfig`

GetOriginalSourceConfig returns the OriginalSourceConfig field if non-nil, zero value otherwise.

### GetOriginalSourceConfigOk

`func (o *AcropolisTargetParamsForRecoverVmRecoveryTargetConfig) GetOriginalSourceConfigOk() (*RecoverAcropolisVmTargetConfigOriginalSourceConfig, bool)`

GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSourceConfig

`func (o *AcropolisTargetParamsForRecoverVmRecoveryTargetConfig) SetOriginalSourceConfig(v RecoverAcropolisVmTargetConfigOriginalSourceConfig)`

SetOriginalSourceConfig sets OriginalSourceConfig field to given value.

### HasOriginalSourceConfig

`func (o *AcropolisTargetParamsForRecoverVmRecoveryTargetConfig) HasOriginalSourceConfig() bool`

HasOriginalSourceConfig returns a boolean if a field has been set.

### SetOriginalSourceConfigNil

`func (o *AcropolisTargetParamsForRecoverVmRecoveryTargetConfig) SetOriginalSourceConfigNil(b bool)`

 SetOriginalSourceConfigNil sets the value for OriginalSourceConfig to be an explicit nil

### UnsetOriginalSourceConfig
`func (o *AcropolisTargetParamsForRecoverVmRecoveryTargetConfig) UnsetOriginalSourceConfig()`

UnsetOriginalSourceConfig ensures that no value is present for OriginalSourceConfig, not even an explicit nil
### GetRecoverToNewSource

`func (o *AcropolisTargetParamsForRecoverVmRecoveryTargetConfig) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *AcropolisTargetParamsForRecoverVmRecoveryTargetConfig) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *AcropolisTargetParamsForRecoverVmRecoveryTargetConfig) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


