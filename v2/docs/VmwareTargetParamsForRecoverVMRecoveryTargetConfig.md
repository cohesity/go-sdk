# VmwareTargetParamsForRecoverVMRecoveryTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewSourceConfig** | Pointer to [**NullableVmwareVmRecoveryTargetConfigNewSourceConfig**](VmwareVmRecoveryTargetConfigNewSourceConfig.md) |  | [optional] 
**OriginalSourceConfig** | Pointer to [**NullableVmwareVmRecoveryTargetConfigOriginalSourceConfig**](VmwareVmRecoveryTargetConfigOriginalSourceConfig.md) |  | [optional] 
**RecoverToNewSource** | **bool** | Specifies the parameter whether the recovery should be performed to a new or an existing Source Target. | 

## Methods

### NewVmwareTargetParamsForRecoverVMRecoveryTargetConfig

`func NewVmwareTargetParamsForRecoverVMRecoveryTargetConfig(recoverToNewSource bool, ) *VmwareTargetParamsForRecoverVMRecoveryTargetConfig`

NewVmwareTargetParamsForRecoverVMRecoveryTargetConfig instantiates a new VmwareTargetParamsForRecoverVMRecoveryTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareTargetParamsForRecoverVMRecoveryTargetConfigWithDefaults

`func NewVmwareTargetParamsForRecoverVMRecoveryTargetConfigWithDefaults() *VmwareTargetParamsForRecoverVMRecoveryTargetConfig`

NewVmwareTargetParamsForRecoverVMRecoveryTargetConfigWithDefaults instantiates a new VmwareTargetParamsForRecoverVMRecoveryTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewSourceConfig

`func (o *VmwareTargetParamsForRecoverVMRecoveryTargetConfig) GetNewSourceConfig() VmwareVmRecoveryTargetConfigNewSourceConfig`

GetNewSourceConfig returns the NewSourceConfig field if non-nil, zero value otherwise.

### GetNewSourceConfigOk

`func (o *VmwareTargetParamsForRecoverVMRecoveryTargetConfig) GetNewSourceConfigOk() (*VmwareVmRecoveryTargetConfigNewSourceConfig, bool)`

GetNewSourceConfigOk returns a tuple with the NewSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSourceConfig

`func (o *VmwareTargetParamsForRecoverVMRecoveryTargetConfig) SetNewSourceConfig(v VmwareVmRecoveryTargetConfigNewSourceConfig)`

SetNewSourceConfig sets NewSourceConfig field to given value.

### HasNewSourceConfig

`func (o *VmwareTargetParamsForRecoverVMRecoveryTargetConfig) HasNewSourceConfig() bool`

HasNewSourceConfig returns a boolean if a field has been set.

### SetNewSourceConfigNil

`func (o *VmwareTargetParamsForRecoverVMRecoveryTargetConfig) SetNewSourceConfigNil(b bool)`

 SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil

### UnsetNewSourceConfig
`func (o *VmwareTargetParamsForRecoverVMRecoveryTargetConfig) UnsetNewSourceConfig()`

UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
### GetOriginalSourceConfig

`func (o *VmwareTargetParamsForRecoverVMRecoveryTargetConfig) GetOriginalSourceConfig() VmwareVmRecoveryTargetConfigOriginalSourceConfig`

GetOriginalSourceConfig returns the OriginalSourceConfig field if non-nil, zero value otherwise.

### GetOriginalSourceConfigOk

`func (o *VmwareTargetParamsForRecoverVMRecoveryTargetConfig) GetOriginalSourceConfigOk() (*VmwareVmRecoveryTargetConfigOriginalSourceConfig, bool)`

GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSourceConfig

`func (o *VmwareTargetParamsForRecoverVMRecoveryTargetConfig) SetOriginalSourceConfig(v VmwareVmRecoveryTargetConfigOriginalSourceConfig)`

SetOriginalSourceConfig sets OriginalSourceConfig field to given value.

### HasOriginalSourceConfig

`func (o *VmwareTargetParamsForRecoverVMRecoveryTargetConfig) HasOriginalSourceConfig() bool`

HasOriginalSourceConfig returns a boolean if a field has been set.

### SetOriginalSourceConfigNil

`func (o *VmwareTargetParamsForRecoverVMRecoveryTargetConfig) SetOriginalSourceConfigNil(b bool)`

 SetOriginalSourceConfigNil sets the value for OriginalSourceConfig to be an explicit nil

### UnsetOriginalSourceConfig
`func (o *VmwareTargetParamsForRecoverVMRecoveryTargetConfig) UnsetOriginalSourceConfig()`

UnsetOriginalSourceConfig ensures that no value is present for OriginalSourceConfig, not even an explicit nil
### GetRecoverToNewSource

`func (o *VmwareTargetParamsForRecoverVMRecoveryTargetConfig) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *VmwareTargetParamsForRecoverVMRecoveryTargetConfig) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *VmwareTargetParamsForRecoverVMRecoveryTargetConfig) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


