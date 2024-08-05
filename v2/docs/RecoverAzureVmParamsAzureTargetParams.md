# RecoverAzureVmParamsAzureTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering other vms if one of vms failed to recover. Default value is false. | [optional] 
**PowerOnVms** | Pointer to **NullableBool** | Specifies whether to power on vms after recovery. If not specified, or false, recovered vms will be in powered off state. | [optional] 
**RecoveryTargetConfig** | Pointer to [**NullableAzureTargetParamsForRecoverVmRecoveryTargetConfig**](AzureTargetParamsForRecoverVmRecoveryTargetConfig.md) |  | [optional] 
**RenameRecoveredVmsParams** | Pointer to [**NullableAcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams**](AcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams.md) |  | [optional] 

## Methods

### NewRecoverAzureVmParamsAzureTargetParams

`func NewRecoverAzureVmParamsAzureTargetParams() *RecoverAzureVmParamsAzureTargetParams`

NewRecoverAzureVmParamsAzureTargetParams instantiates a new RecoverAzureVmParamsAzureTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverAzureVmParamsAzureTargetParamsWithDefaults

`func NewRecoverAzureVmParamsAzureTargetParamsWithDefaults() *RecoverAzureVmParamsAzureTargetParams`

NewRecoverAzureVmParamsAzureTargetParamsWithDefaults instantiates a new RecoverAzureVmParamsAzureTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *RecoverAzureVmParamsAzureTargetParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *RecoverAzureVmParamsAzureTargetParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *RecoverAzureVmParamsAzureTargetParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *RecoverAzureVmParamsAzureTargetParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *RecoverAzureVmParamsAzureTargetParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *RecoverAzureVmParamsAzureTargetParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetPowerOnVms

`func (o *RecoverAzureVmParamsAzureTargetParams) GetPowerOnVms() bool`

GetPowerOnVms returns the PowerOnVms field if non-nil, zero value otherwise.

### GetPowerOnVmsOk

`func (o *RecoverAzureVmParamsAzureTargetParams) GetPowerOnVmsOk() (*bool, bool)`

GetPowerOnVmsOk returns a tuple with the PowerOnVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnVms

`func (o *RecoverAzureVmParamsAzureTargetParams) SetPowerOnVms(v bool)`

SetPowerOnVms sets PowerOnVms field to given value.

### HasPowerOnVms

`func (o *RecoverAzureVmParamsAzureTargetParams) HasPowerOnVms() bool`

HasPowerOnVms returns a boolean if a field has been set.

### SetPowerOnVmsNil

`func (o *RecoverAzureVmParamsAzureTargetParams) SetPowerOnVmsNil(b bool)`

 SetPowerOnVmsNil sets the value for PowerOnVms to be an explicit nil

### UnsetPowerOnVms
`func (o *RecoverAzureVmParamsAzureTargetParams) UnsetPowerOnVms()`

UnsetPowerOnVms ensures that no value is present for PowerOnVms, not even an explicit nil
### GetRecoveryTargetConfig

`func (o *RecoverAzureVmParamsAzureTargetParams) GetRecoveryTargetConfig() AzureTargetParamsForRecoverVmRecoveryTargetConfig`

GetRecoveryTargetConfig returns the RecoveryTargetConfig field if non-nil, zero value otherwise.

### GetRecoveryTargetConfigOk

`func (o *RecoverAzureVmParamsAzureTargetParams) GetRecoveryTargetConfigOk() (*AzureTargetParamsForRecoverVmRecoveryTargetConfig, bool)`

GetRecoveryTargetConfigOk returns a tuple with the RecoveryTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTargetConfig

`func (o *RecoverAzureVmParamsAzureTargetParams) SetRecoveryTargetConfig(v AzureTargetParamsForRecoverVmRecoveryTargetConfig)`

SetRecoveryTargetConfig sets RecoveryTargetConfig field to given value.

### HasRecoveryTargetConfig

`func (o *RecoverAzureVmParamsAzureTargetParams) HasRecoveryTargetConfig() bool`

HasRecoveryTargetConfig returns a boolean if a field has been set.

### SetRecoveryTargetConfigNil

`func (o *RecoverAzureVmParamsAzureTargetParams) SetRecoveryTargetConfigNil(b bool)`

 SetRecoveryTargetConfigNil sets the value for RecoveryTargetConfig to be an explicit nil

### UnsetRecoveryTargetConfig
`func (o *RecoverAzureVmParamsAzureTargetParams) UnsetRecoveryTargetConfig()`

UnsetRecoveryTargetConfig ensures that no value is present for RecoveryTargetConfig, not even an explicit nil
### GetRenameRecoveredVmsParams

`func (o *RecoverAzureVmParamsAzureTargetParams) GetRenameRecoveredVmsParams() AcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams`

GetRenameRecoveredVmsParams returns the RenameRecoveredVmsParams field if non-nil, zero value otherwise.

### GetRenameRecoveredVmsParamsOk

`func (o *RecoverAzureVmParamsAzureTargetParams) GetRenameRecoveredVmsParamsOk() (*AcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams, bool)`

GetRenameRecoveredVmsParamsOk returns a tuple with the RenameRecoveredVmsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameRecoveredVmsParams

`func (o *RecoverAzureVmParamsAzureTargetParams) SetRenameRecoveredVmsParams(v AcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams)`

SetRenameRecoveredVmsParams sets RenameRecoveredVmsParams field to given value.

### HasRenameRecoveredVmsParams

`func (o *RecoverAzureVmParamsAzureTargetParams) HasRenameRecoveredVmsParams() bool`

HasRenameRecoveredVmsParams returns a boolean if a field has been set.

### SetRenameRecoveredVmsParamsNil

`func (o *RecoverAzureVmParamsAzureTargetParams) SetRenameRecoveredVmsParamsNil(b bool)`

 SetRenameRecoveredVmsParamsNil sets the value for RenameRecoveredVmsParams to be an explicit nil

### UnsetRenameRecoveredVmsParams
`func (o *RecoverAzureVmParamsAzureTargetParams) UnsetRenameRecoveredVmsParams()`

UnsetRenameRecoveredVmsParams ensures that no value is present for RenameRecoveredVmsParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


