# HyperVVmRecoveryTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewSourceConfig** | Pointer to [**NullableHyperVVmRecoveryTargetConfigNewSourceConfig**](HyperVVmRecoveryTargetConfigNewSourceConfig.md) |  | [optional] 
**OriginalSourceConfig** | Pointer to [**NullableHyperVVmRecoveryTargetConfigOriginalSourceConfig**](HyperVVmRecoveryTargetConfigOriginalSourceConfig.md) |  | [optional] 
**RecoverToNewSource** | **bool** | Specifies the parameter whether the recovery should be performed to a new or an existing Source Target. | 

## Methods

### NewHyperVVmRecoveryTargetConfig

`func NewHyperVVmRecoveryTargetConfig(recoverToNewSource bool, ) *HyperVVmRecoveryTargetConfig`

NewHyperVVmRecoveryTargetConfig instantiates a new HyperVVmRecoveryTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperVVmRecoveryTargetConfigWithDefaults

`func NewHyperVVmRecoveryTargetConfigWithDefaults() *HyperVVmRecoveryTargetConfig`

NewHyperVVmRecoveryTargetConfigWithDefaults instantiates a new HyperVVmRecoveryTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewSourceConfig

`func (o *HyperVVmRecoveryTargetConfig) GetNewSourceConfig() HyperVVmRecoveryTargetConfigNewSourceConfig`

GetNewSourceConfig returns the NewSourceConfig field if non-nil, zero value otherwise.

### GetNewSourceConfigOk

`func (o *HyperVVmRecoveryTargetConfig) GetNewSourceConfigOk() (*HyperVVmRecoveryTargetConfigNewSourceConfig, bool)`

GetNewSourceConfigOk returns a tuple with the NewSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSourceConfig

`func (o *HyperVVmRecoveryTargetConfig) SetNewSourceConfig(v HyperVVmRecoveryTargetConfigNewSourceConfig)`

SetNewSourceConfig sets NewSourceConfig field to given value.

### HasNewSourceConfig

`func (o *HyperVVmRecoveryTargetConfig) HasNewSourceConfig() bool`

HasNewSourceConfig returns a boolean if a field has been set.

### SetNewSourceConfigNil

`func (o *HyperVVmRecoveryTargetConfig) SetNewSourceConfigNil(b bool)`

 SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil

### UnsetNewSourceConfig
`func (o *HyperVVmRecoveryTargetConfig) UnsetNewSourceConfig()`

UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
### GetOriginalSourceConfig

`func (o *HyperVVmRecoveryTargetConfig) GetOriginalSourceConfig() HyperVVmRecoveryTargetConfigOriginalSourceConfig`

GetOriginalSourceConfig returns the OriginalSourceConfig field if non-nil, zero value otherwise.

### GetOriginalSourceConfigOk

`func (o *HyperVVmRecoveryTargetConfig) GetOriginalSourceConfigOk() (*HyperVVmRecoveryTargetConfigOriginalSourceConfig, bool)`

GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSourceConfig

`func (o *HyperVVmRecoveryTargetConfig) SetOriginalSourceConfig(v HyperVVmRecoveryTargetConfigOriginalSourceConfig)`

SetOriginalSourceConfig sets OriginalSourceConfig field to given value.

### HasOriginalSourceConfig

`func (o *HyperVVmRecoveryTargetConfig) HasOriginalSourceConfig() bool`

HasOriginalSourceConfig returns a boolean if a field has been set.

### SetOriginalSourceConfigNil

`func (o *HyperVVmRecoveryTargetConfig) SetOriginalSourceConfigNil(b bool)`

 SetOriginalSourceConfigNil sets the value for OriginalSourceConfig to be an explicit nil

### UnsetOriginalSourceConfig
`func (o *HyperVVmRecoveryTargetConfig) UnsetOriginalSourceConfig()`

UnsetOriginalSourceConfig ensures that no value is present for OriginalSourceConfig, not even an explicit nil
### GetRecoverToNewSource

`func (o *HyperVVmRecoveryTargetConfig) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *HyperVVmRecoveryTargetConfig) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *HyperVVmRecoveryTargetConfig) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


