# RecoverAcropolisVmParamsAcropolisTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering other vms if one of vms failed to recover. Default value is false. | [optional] 
**PowerOnVms** | Pointer to **NullableBool** | Specifies whether to power on vms after recovery. If not specified, or false, recovered vms will be in powered off state. | [optional] 
**RecoverExcludedDisk** | Pointer to **NullableBool** | Specifies whether to recover excluded disk while performing recovery of a VM by creating empty disks for them. Default value is false. | [optional] 
**RecoveryProcessType** | Pointer to **string** | Specifies type of Recovery Process to be used. InstantRecovery/CopyRecovery etc... Default value is InstantRecovery. | [optional] 
**RecoveryTargetConfig** | Pointer to [**NullableAcropolisTargetParamsForRecoverVmRecoveryTargetConfig**](AcropolisTargetParamsForRecoverVmRecoveryTargetConfig.md) |  | [optional] 
**RenameRecoveredVmsParams** | Pointer to [**NullableAcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams**](AcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams.md) |  | [optional] 
**VlanConfig** | Pointer to [**NullableAcropolisTargetParamsForRecoverVmVlanConfig**](AcropolisTargetParamsForRecoverVmVlanConfig.md) |  | [optional] 

## Methods

### NewRecoverAcropolisVmParamsAcropolisTargetParams

`func NewRecoverAcropolisVmParamsAcropolisTargetParams() *RecoverAcropolisVmParamsAcropolisTargetParams`

NewRecoverAcropolisVmParamsAcropolisTargetParams instantiates a new RecoverAcropolisVmParamsAcropolisTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverAcropolisVmParamsAcropolisTargetParamsWithDefaults

`func NewRecoverAcropolisVmParamsAcropolisTargetParamsWithDefaults() *RecoverAcropolisVmParamsAcropolisTargetParams`

NewRecoverAcropolisVmParamsAcropolisTargetParamsWithDefaults instantiates a new RecoverAcropolisVmParamsAcropolisTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetPowerOnVms

`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) GetPowerOnVms() bool`

GetPowerOnVms returns the PowerOnVms field if non-nil, zero value otherwise.

### GetPowerOnVmsOk

`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) GetPowerOnVmsOk() (*bool, bool)`

GetPowerOnVmsOk returns a tuple with the PowerOnVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnVms

`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) SetPowerOnVms(v bool)`

SetPowerOnVms sets PowerOnVms field to given value.

### HasPowerOnVms

`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) HasPowerOnVms() bool`

HasPowerOnVms returns a boolean if a field has been set.

### SetPowerOnVmsNil

`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) SetPowerOnVmsNil(b bool)`

 SetPowerOnVmsNil sets the value for PowerOnVms to be an explicit nil

### UnsetPowerOnVms
`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) UnsetPowerOnVms()`

UnsetPowerOnVms ensures that no value is present for PowerOnVms, not even an explicit nil
### GetRecoverExcludedDisk

`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) GetRecoverExcludedDisk() bool`

GetRecoverExcludedDisk returns the RecoverExcludedDisk field if non-nil, zero value otherwise.

### GetRecoverExcludedDiskOk

`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) GetRecoverExcludedDiskOk() (*bool, bool)`

GetRecoverExcludedDiskOk returns a tuple with the RecoverExcludedDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverExcludedDisk

`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) SetRecoverExcludedDisk(v bool)`

SetRecoverExcludedDisk sets RecoverExcludedDisk field to given value.

### HasRecoverExcludedDisk

`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) HasRecoverExcludedDisk() bool`

HasRecoverExcludedDisk returns a boolean if a field has been set.

### SetRecoverExcludedDiskNil

`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) SetRecoverExcludedDiskNil(b bool)`

 SetRecoverExcludedDiskNil sets the value for RecoverExcludedDisk to be an explicit nil

### UnsetRecoverExcludedDisk
`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) UnsetRecoverExcludedDisk()`

UnsetRecoverExcludedDisk ensures that no value is present for RecoverExcludedDisk, not even an explicit nil
### GetRecoveryProcessType

`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) GetRecoveryProcessType() string`

GetRecoveryProcessType returns the RecoveryProcessType field if non-nil, zero value otherwise.

### GetRecoveryProcessTypeOk

`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) GetRecoveryProcessTypeOk() (*string, bool)`

GetRecoveryProcessTypeOk returns a tuple with the RecoveryProcessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryProcessType

`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) SetRecoveryProcessType(v string)`

SetRecoveryProcessType sets RecoveryProcessType field to given value.

### HasRecoveryProcessType

`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) HasRecoveryProcessType() bool`

HasRecoveryProcessType returns a boolean if a field has been set.

### GetRecoveryTargetConfig

`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) GetRecoveryTargetConfig() AcropolisTargetParamsForRecoverVmRecoveryTargetConfig`

GetRecoveryTargetConfig returns the RecoveryTargetConfig field if non-nil, zero value otherwise.

### GetRecoveryTargetConfigOk

`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) GetRecoveryTargetConfigOk() (*AcropolisTargetParamsForRecoverVmRecoveryTargetConfig, bool)`

GetRecoveryTargetConfigOk returns a tuple with the RecoveryTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTargetConfig

`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) SetRecoveryTargetConfig(v AcropolisTargetParamsForRecoverVmRecoveryTargetConfig)`

SetRecoveryTargetConfig sets RecoveryTargetConfig field to given value.

### HasRecoveryTargetConfig

`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) HasRecoveryTargetConfig() bool`

HasRecoveryTargetConfig returns a boolean if a field has been set.

### SetRecoveryTargetConfigNil

`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) SetRecoveryTargetConfigNil(b bool)`

 SetRecoveryTargetConfigNil sets the value for RecoveryTargetConfig to be an explicit nil

### UnsetRecoveryTargetConfig
`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) UnsetRecoveryTargetConfig()`

UnsetRecoveryTargetConfig ensures that no value is present for RecoveryTargetConfig, not even an explicit nil
### GetRenameRecoveredVmsParams

`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) GetRenameRecoveredVmsParams() AcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams`

GetRenameRecoveredVmsParams returns the RenameRecoveredVmsParams field if non-nil, zero value otherwise.

### GetRenameRecoveredVmsParamsOk

`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) GetRenameRecoveredVmsParamsOk() (*AcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams, bool)`

GetRenameRecoveredVmsParamsOk returns a tuple with the RenameRecoveredVmsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameRecoveredVmsParams

`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) SetRenameRecoveredVmsParams(v AcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams)`

SetRenameRecoveredVmsParams sets RenameRecoveredVmsParams field to given value.

### HasRenameRecoveredVmsParams

`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) HasRenameRecoveredVmsParams() bool`

HasRenameRecoveredVmsParams returns a boolean if a field has been set.

### SetRenameRecoveredVmsParamsNil

`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) SetRenameRecoveredVmsParamsNil(b bool)`

 SetRenameRecoveredVmsParamsNil sets the value for RenameRecoveredVmsParams to be an explicit nil

### UnsetRenameRecoveredVmsParams
`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) UnsetRenameRecoveredVmsParams()`

UnsetRenameRecoveredVmsParams ensures that no value is present for RenameRecoveredVmsParams, not even an explicit nil
### GetVlanConfig

`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) GetVlanConfig() AcropolisTargetParamsForRecoverVmVlanConfig`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) GetVlanConfigOk() (*AcropolisTargetParamsForRecoverVmVlanConfig, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) SetVlanConfig(v AcropolisTargetParamsForRecoverVmVlanConfig)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.

### SetVlanConfigNil

`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) SetVlanConfigNil(b bool)`

 SetVlanConfigNil sets the value for VlanConfig to be an explicit nil

### UnsetVlanConfig
`func (o *RecoverAcropolisVmParamsAcropolisTargetParams) UnsetVlanConfig()`

UnsetVlanConfig ensures that no value is present for VlanConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


