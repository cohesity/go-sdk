# VmwareTargetParamsForRecoverVAppRecoveryTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewSourceConfig** | Pointer to [**NullableVmwareVAppRecoveryTargetConfigNewSourceConfig**](VmwareVAppRecoveryTargetConfigNewSourceConfig.md) |  | [optional] 
**OriginalSourceConfig** | Pointer to [**NullableVmwareVAppRecoveryTargetConfigOriginalSourceConfig**](VmwareVAppRecoveryTargetConfigOriginalSourceConfig.md) |  | [optional] 
**RecoverToNewSource** | **bool** | Specifies the parameter whether the recovery should be performed to a new or an existing Source Target. | 

## Methods

### NewVmwareTargetParamsForRecoverVAppRecoveryTargetConfig

`func NewVmwareTargetParamsForRecoverVAppRecoveryTargetConfig(recoverToNewSource bool, ) *VmwareTargetParamsForRecoverVAppRecoveryTargetConfig`

NewVmwareTargetParamsForRecoverVAppRecoveryTargetConfig instantiates a new VmwareTargetParamsForRecoverVAppRecoveryTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareTargetParamsForRecoverVAppRecoveryTargetConfigWithDefaults

`func NewVmwareTargetParamsForRecoverVAppRecoveryTargetConfigWithDefaults() *VmwareTargetParamsForRecoverVAppRecoveryTargetConfig`

NewVmwareTargetParamsForRecoverVAppRecoveryTargetConfigWithDefaults instantiates a new VmwareTargetParamsForRecoverVAppRecoveryTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewSourceConfig

`func (o *VmwareTargetParamsForRecoverVAppRecoveryTargetConfig) GetNewSourceConfig() VmwareVAppRecoveryTargetConfigNewSourceConfig`

GetNewSourceConfig returns the NewSourceConfig field if non-nil, zero value otherwise.

### GetNewSourceConfigOk

`func (o *VmwareTargetParamsForRecoverVAppRecoveryTargetConfig) GetNewSourceConfigOk() (*VmwareVAppRecoveryTargetConfigNewSourceConfig, bool)`

GetNewSourceConfigOk returns a tuple with the NewSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSourceConfig

`func (o *VmwareTargetParamsForRecoverVAppRecoveryTargetConfig) SetNewSourceConfig(v VmwareVAppRecoveryTargetConfigNewSourceConfig)`

SetNewSourceConfig sets NewSourceConfig field to given value.

### HasNewSourceConfig

`func (o *VmwareTargetParamsForRecoverVAppRecoveryTargetConfig) HasNewSourceConfig() bool`

HasNewSourceConfig returns a boolean if a field has been set.

### SetNewSourceConfigNil

`func (o *VmwareTargetParamsForRecoverVAppRecoveryTargetConfig) SetNewSourceConfigNil(b bool)`

 SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil

### UnsetNewSourceConfig
`func (o *VmwareTargetParamsForRecoverVAppRecoveryTargetConfig) UnsetNewSourceConfig()`

UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
### GetOriginalSourceConfig

`func (o *VmwareTargetParamsForRecoverVAppRecoveryTargetConfig) GetOriginalSourceConfig() VmwareVAppRecoveryTargetConfigOriginalSourceConfig`

GetOriginalSourceConfig returns the OriginalSourceConfig field if non-nil, zero value otherwise.

### GetOriginalSourceConfigOk

`func (o *VmwareTargetParamsForRecoverVAppRecoveryTargetConfig) GetOriginalSourceConfigOk() (*VmwareVAppRecoveryTargetConfigOriginalSourceConfig, bool)`

GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSourceConfig

`func (o *VmwareTargetParamsForRecoverVAppRecoveryTargetConfig) SetOriginalSourceConfig(v VmwareVAppRecoveryTargetConfigOriginalSourceConfig)`

SetOriginalSourceConfig sets OriginalSourceConfig field to given value.

### HasOriginalSourceConfig

`func (o *VmwareTargetParamsForRecoverVAppRecoveryTargetConfig) HasOriginalSourceConfig() bool`

HasOriginalSourceConfig returns a boolean if a field has been set.

### SetOriginalSourceConfigNil

`func (o *VmwareTargetParamsForRecoverVAppRecoveryTargetConfig) SetOriginalSourceConfigNil(b bool)`

 SetOriginalSourceConfigNil sets the value for OriginalSourceConfig to be an explicit nil

### UnsetOriginalSourceConfig
`func (o *VmwareTargetParamsForRecoverVAppRecoveryTargetConfig) UnsetOriginalSourceConfig()`

UnsetOriginalSourceConfig ensures that no value is present for OriginalSourceConfig, not even an explicit nil
### GetRecoverToNewSource

`func (o *VmwareTargetParamsForRecoverVAppRecoveryTargetConfig) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *VmwareTargetParamsForRecoverVAppRecoveryTargetConfig) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *VmwareTargetParamsForRecoverVAppRecoveryTargetConfig) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


