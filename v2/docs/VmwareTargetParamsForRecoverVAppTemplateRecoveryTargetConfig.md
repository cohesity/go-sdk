# VmwareTargetParamsForRecoverVAppTemplateRecoveryTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewSourceConfig** | Pointer to [**NullableVmwareVAppTemplateRecoveryTargetConfigNewSourceConfig**](VmwareVAppTemplateRecoveryTargetConfigNewSourceConfig.md) |  | [optional] 
**OriginalSourceConfig** | Pointer to [**NullableVmwareVAppTemplateRecoveryTargetConfigOriginalSourceConfig**](VmwareVAppTemplateRecoveryTargetConfigOriginalSourceConfig.md) |  | [optional] 
**RecoverToNewSource** | **bool** | Specifies the parameter whether the recovery should be performed to a new or an existing Source Target. | 

## Methods

### NewVmwareTargetParamsForRecoverVAppTemplateRecoveryTargetConfig

`func NewVmwareTargetParamsForRecoverVAppTemplateRecoveryTargetConfig(recoverToNewSource bool, ) *VmwareTargetParamsForRecoverVAppTemplateRecoveryTargetConfig`

NewVmwareTargetParamsForRecoverVAppTemplateRecoveryTargetConfig instantiates a new VmwareTargetParamsForRecoverVAppTemplateRecoveryTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareTargetParamsForRecoverVAppTemplateRecoveryTargetConfigWithDefaults

`func NewVmwareTargetParamsForRecoverVAppTemplateRecoveryTargetConfigWithDefaults() *VmwareTargetParamsForRecoverVAppTemplateRecoveryTargetConfig`

NewVmwareTargetParamsForRecoverVAppTemplateRecoveryTargetConfigWithDefaults instantiates a new VmwareTargetParamsForRecoverVAppTemplateRecoveryTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewSourceConfig

`func (o *VmwareTargetParamsForRecoverVAppTemplateRecoveryTargetConfig) GetNewSourceConfig() VmwareVAppTemplateRecoveryTargetConfigNewSourceConfig`

GetNewSourceConfig returns the NewSourceConfig field if non-nil, zero value otherwise.

### GetNewSourceConfigOk

`func (o *VmwareTargetParamsForRecoverVAppTemplateRecoveryTargetConfig) GetNewSourceConfigOk() (*VmwareVAppTemplateRecoveryTargetConfigNewSourceConfig, bool)`

GetNewSourceConfigOk returns a tuple with the NewSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSourceConfig

`func (o *VmwareTargetParamsForRecoverVAppTemplateRecoveryTargetConfig) SetNewSourceConfig(v VmwareVAppTemplateRecoveryTargetConfigNewSourceConfig)`

SetNewSourceConfig sets NewSourceConfig field to given value.

### HasNewSourceConfig

`func (o *VmwareTargetParamsForRecoverVAppTemplateRecoveryTargetConfig) HasNewSourceConfig() bool`

HasNewSourceConfig returns a boolean if a field has been set.

### SetNewSourceConfigNil

`func (o *VmwareTargetParamsForRecoverVAppTemplateRecoveryTargetConfig) SetNewSourceConfigNil(b bool)`

 SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil

### UnsetNewSourceConfig
`func (o *VmwareTargetParamsForRecoverVAppTemplateRecoveryTargetConfig) UnsetNewSourceConfig()`

UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
### GetOriginalSourceConfig

`func (o *VmwareTargetParamsForRecoverVAppTemplateRecoveryTargetConfig) GetOriginalSourceConfig() VmwareVAppTemplateRecoveryTargetConfigOriginalSourceConfig`

GetOriginalSourceConfig returns the OriginalSourceConfig field if non-nil, zero value otherwise.

### GetOriginalSourceConfigOk

`func (o *VmwareTargetParamsForRecoverVAppTemplateRecoveryTargetConfig) GetOriginalSourceConfigOk() (*VmwareVAppTemplateRecoveryTargetConfigOriginalSourceConfig, bool)`

GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSourceConfig

`func (o *VmwareTargetParamsForRecoverVAppTemplateRecoveryTargetConfig) SetOriginalSourceConfig(v VmwareVAppTemplateRecoveryTargetConfigOriginalSourceConfig)`

SetOriginalSourceConfig sets OriginalSourceConfig field to given value.

### HasOriginalSourceConfig

`func (o *VmwareTargetParamsForRecoverVAppTemplateRecoveryTargetConfig) HasOriginalSourceConfig() bool`

HasOriginalSourceConfig returns a boolean if a field has been set.

### SetOriginalSourceConfigNil

`func (o *VmwareTargetParamsForRecoverVAppTemplateRecoveryTargetConfig) SetOriginalSourceConfigNil(b bool)`

 SetOriginalSourceConfigNil sets the value for OriginalSourceConfig to be an explicit nil

### UnsetOriginalSourceConfig
`func (o *VmwareTargetParamsForRecoverVAppTemplateRecoveryTargetConfig) UnsetOriginalSourceConfig()`

UnsetOriginalSourceConfig ensures that no value is present for OriginalSourceConfig, not even an explicit nil
### GetRecoverToNewSource

`func (o *VmwareTargetParamsForRecoverVAppTemplateRecoveryTargetConfig) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *VmwareTargetParamsForRecoverVAppTemplateRecoveryTargetConfig) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *VmwareTargetParamsForRecoverVAppTemplateRecoveryTargetConfig) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


