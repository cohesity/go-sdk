# VmwareVAppTemplateRecoveryTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoverToNewSource** | **bool** | Specifies the parameter whether the recovery should be performed to a new or an existing Source Target. | 
**NewSourceConfig** | Pointer to [**NullableRecoverVmwareVAppTemplateNewSourceConfig**](RecoverVmwareVAppTemplateNewSourceConfig.md) | Specifies the new destination Source configuration parameters where the vApp templates will be recovered. This is mandatory if recoverToNewSource is set to true. | [optional] 
**OriginalSourceConfig** | Pointer to [**NullableRecoverVmwareVAppOriginalSourceConfig**](RecoverVmwareVAppOriginalSourceConfig.md) | Specifies the Source configuration if vApp templates are being recovered to Original Source. If not specified, all the configuration parameters will be retained. | [optional] 

## Methods

### NewVmwareVAppTemplateRecoveryTargetConfig

`func NewVmwareVAppTemplateRecoveryTargetConfig(recoverToNewSource bool, ) *VmwareVAppTemplateRecoveryTargetConfig`

NewVmwareVAppTemplateRecoveryTargetConfig instantiates a new VmwareVAppTemplateRecoveryTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareVAppTemplateRecoveryTargetConfigWithDefaults

`func NewVmwareVAppTemplateRecoveryTargetConfigWithDefaults() *VmwareVAppTemplateRecoveryTargetConfig`

NewVmwareVAppTemplateRecoveryTargetConfigWithDefaults instantiates a new VmwareVAppTemplateRecoveryTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecoverToNewSource

`func (o *VmwareVAppTemplateRecoveryTargetConfig) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *VmwareVAppTemplateRecoveryTargetConfig) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *VmwareVAppTemplateRecoveryTargetConfig) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.


### GetNewSourceConfig

`func (o *VmwareVAppTemplateRecoveryTargetConfig) GetNewSourceConfig() RecoverVmwareVAppTemplateNewSourceConfig`

GetNewSourceConfig returns the NewSourceConfig field if non-nil, zero value otherwise.

### GetNewSourceConfigOk

`func (o *VmwareVAppTemplateRecoveryTargetConfig) GetNewSourceConfigOk() (*RecoverVmwareVAppTemplateNewSourceConfig, bool)`

GetNewSourceConfigOk returns a tuple with the NewSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSourceConfig

`func (o *VmwareVAppTemplateRecoveryTargetConfig) SetNewSourceConfig(v RecoverVmwareVAppTemplateNewSourceConfig)`

SetNewSourceConfig sets NewSourceConfig field to given value.

### HasNewSourceConfig

`func (o *VmwareVAppTemplateRecoveryTargetConfig) HasNewSourceConfig() bool`

HasNewSourceConfig returns a boolean if a field has been set.

### SetNewSourceConfigNil

`func (o *VmwareVAppTemplateRecoveryTargetConfig) SetNewSourceConfigNil(b bool)`

 SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil

### UnsetNewSourceConfig
`func (o *VmwareVAppTemplateRecoveryTargetConfig) UnsetNewSourceConfig()`

UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
### GetOriginalSourceConfig

`func (o *VmwareVAppTemplateRecoveryTargetConfig) GetOriginalSourceConfig() RecoverVmwareVAppOriginalSourceConfig`

GetOriginalSourceConfig returns the OriginalSourceConfig field if non-nil, zero value otherwise.

### GetOriginalSourceConfigOk

`func (o *VmwareVAppTemplateRecoveryTargetConfig) GetOriginalSourceConfigOk() (*RecoverVmwareVAppOriginalSourceConfig, bool)`

GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSourceConfig

`func (o *VmwareVAppTemplateRecoveryTargetConfig) SetOriginalSourceConfig(v RecoverVmwareVAppOriginalSourceConfig)`

SetOriginalSourceConfig sets OriginalSourceConfig field to given value.

### HasOriginalSourceConfig

`func (o *VmwareVAppTemplateRecoveryTargetConfig) HasOriginalSourceConfig() bool`

HasOriginalSourceConfig returns a boolean if a field has been set.

### SetOriginalSourceConfigNil

`func (o *VmwareVAppTemplateRecoveryTargetConfig) SetOriginalSourceConfigNil(b bool)`

 SetOriginalSourceConfigNil sets the value for OriginalSourceConfig to be an explicit nil

### UnsetOriginalSourceConfig
`func (o *VmwareVAppTemplateRecoveryTargetConfig) UnsetOriginalSourceConfig()`

UnsetOriginalSourceConfig ensures that no value is present for OriginalSourceConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


