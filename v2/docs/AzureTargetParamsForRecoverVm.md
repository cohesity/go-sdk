# AzureTargetParamsForRecoverVm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering other vms if one of vms failed to recover. Default value is false. | [optional] 
**PowerOnVms** | Pointer to **NullableBool** | Specifies whether to power on vms after recovery. If not specified, or false, recovered vms will be in powered off state. | [optional] 
**RecoveryTargetConfig** | Pointer to [**NullableAzureTargetParamsForRecoverVmRecoveryTargetConfig**](AzureTargetParamsForRecoverVmRecoveryTargetConfig.md) |  | [optional] 
**RenameRecoveredVmsParams** | Pointer to [**NullableAcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams**](AcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams.md) |  | [optional] 

## Methods

### NewAzureTargetParamsForRecoverVm

`func NewAzureTargetParamsForRecoverVm() *AzureTargetParamsForRecoverVm`

NewAzureTargetParamsForRecoverVm instantiates a new AzureTargetParamsForRecoverVm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureTargetParamsForRecoverVmWithDefaults

`func NewAzureTargetParamsForRecoverVmWithDefaults() *AzureTargetParamsForRecoverVm`

NewAzureTargetParamsForRecoverVmWithDefaults instantiates a new AzureTargetParamsForRecoverVm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *AzureTargetParamsForRecoverVm) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *AzureTargetParamsForRecoverVm) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *AzureTargetParamsForRecoverVm) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *AzureTargetParamsForRecoverVm) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *AzureTargetParamsForRecoverVm) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *AzureTargetParamsForRecoverVm) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetPowerOnVms

`func (o *AzureTargetParamsForRecoverVm) GetPowerOnVms() bool`

GetPowerOnVms returns the PowerOnVms field if non-nil, zero value otherwise.

### GetPowerOnVmsOk

`func (o *AzureTargetParamsForRecoverVm) GetPowerOnVmsOk() (*bool, bool)`

GetPowerOnVmsOk returns a tuple with the PowerOnVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnVms

`func (o *AzureTargetParamsForRecoverVm) SetPowerOnVms(v bool)`

SetPowerOnVms sets PowerOnVms field to given value.

### HasPowerOnVms

`func (o *AzureTargetParamsForRecoverVm) HasPowerOnVms() bool`

HasPowerOnVms returns a boolean if a field has been set.

### SetPowerOnVmsNil

`func (o *AzureTargetParamsForRecoverVm) SetPowerOnVmsNil(b bool)`

 SetPowerOnVmsNil sets the value for PowerOnVms to be an explicit nil

### UnsetPowerOnVms
`func (o *AzureTargetParamsForRecoverVm) UnsetPowerOnVms()`

UnsetPowerOnVms ensures that no value is present for PowerOnVms, not even an explicit nil
### GetRecoveryTargetConfig

`func (o *AzureTargetParamsForRecoverVm) GetRecoveryTargetConfig() AzureTargetParamsForRecoverVmRecoveryTargetConfig`

GetRecoveryTargetConfig returns the RecoveryTargetConfig field if non-nil, zero value otherwise.

### GetRecoveryTargetConfigOk

`func (o *AzureTargetParamsForRecoverVm) GetRecoveryTargetConfigOk() (*AzureTargetParamsForRecoverVmRecoveryTargetConfig, bool)`

GetRecoveryTargetConfigOk returns a tuple with the RecoveryTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTargetConfig

`func (o *AzureTargetParamsForRecoverVm) SetRecoveryTargetConfig(v AzureTargetParamsForRecoverVmRecoveryTargetConfig)`

SetRecoveryTargetConfig sets RecoveryTargetConfig field to given value.

### HasRecoveryTargetConfig

`func (o *AzureTargetParamsForRecoverVm) HasRecoveryTargetConfig() bool`

HasRecoveryTargetConfig returns a boolean if a field has been set.

### SetRecoveryTargetConfigNil

`func (o *AzureTargetParamsForRecoverVm) SetRecoveryTargetConfigNil(b bool)`

 SetRecoveryTargetConfigNil sets the value for RecoveryTargetConfig to be an explicit nil

### UnsetRecoveryTargetConfig
`func (o *AzureTargetParamsForRecoverVm) UnsetRecoveryTargetConfig()`

UnsetRecoveryTargetConfig ensures that no value is present for RecoveryTargetConfig, not even an explicit nil
### GetRenameRecoveredVmsParams

`func (o *AzureTargetParamsForRecoverVm) GetRenameRecoveredVmsParams() AcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams`

GetRenameRecoveredVmsParams returns the RenameRecoveredVmsParams field if non-nil, zero value otherwise.

### GetRenameRecoveredVmsParamsOk

`func (o *AzureTargetParamsForRecoverVm) GetRenameRecoveredVmsParamsOk() (*AcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams, bool)`

GetRenameRecoveredVmsParamsOk returns a tuple with the RenameRecoveredVmsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameRecoveredVmsParams

`func (o *AzureTargetParamsForRecoverVm) SetRenameRecoveredVmsParams(v AcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams)`

SetRenameRecoveredVmsParams sets RenameRecoveredVmsParams field to given value.

### HasRenameRecoveredVmsParams

`func (o *AzureTargetParamsForRecoverVm) HasRenameRecoveredVmsParams() bool`

HasRenameRecoveredVmsParams returns a boolean if a field has been set.

### SetRenameRecoveredVmsParamsNil

`func (o *AzureTargetParamsForRecoverVm) SetRenameRecoveredVmsParamsNil(b bool)`

 SetRenameRecoveredVmsParamsNil sets the value for RenameRecoveredVmsParams to be an explicit nil

### UnsetRenameRecoveredVmsParams
`func (o *AzureTargetParamsForRecoverVm) UnsetRenameRecoveredVmsParams()`

UnsetRenameRecoveredVmsParams ensures that no value is present for RenameRecoveredVmsParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


