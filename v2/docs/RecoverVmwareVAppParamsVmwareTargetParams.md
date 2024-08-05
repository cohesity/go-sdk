# RecoverVmwareVAppParamsVmwareTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttemptDifferentialRestore** | Pointer to **NullableBool** | Specifies whether to attempt differential restore. | [optional] 
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering other vms if one of vms failed to recover. Default value is false. | [optional] 
**DiskProvisionType** | Pointer to **NullableString** | Specifies the Virtual Disk Provisioning Policies for Vmware VM | [optional] 
**EnableNBDSSLFallback** | Pointer to **NullableBool** | If this field is set to true and SAN transport recovery fails, then recovery will fallback to use NBDSSL transport. This field only applies if &#39;leverageSanTransport&#39; is set to true. | [optional] 
**LeverageSanTransport** | Pointer to **NullableBool** | Specifies whether to enable SAN transport for copy recovery or not | [optional] 
**PowerOnVms** | Pointer to **NullableBool** | Specifies whether to power on vms after recovery. If not specified, or false, recovered vms will be in powered off state. | [optional] 
**RecoveryProcessType** | Pointer to **string** | Specifies type of Recovery Process to be used. InstantRecovery/CopyRecovery etc... Default value is InstantRecovery. | [optional] 
**RecoveryTargetConfig** | Pointer to [**NullableVmwareTargetParamsForRecoverVAppRecoveryTargetConfig**](VmwareTargetParamsForRecoverVAppRecoveryTargetConfig.md) |  | [optional] 
**RenameRecoveredVAppsParams** | Pointer to [**NullableVmwareTargetParamsForRecoverVAppRenameRecoveredVAppsParams**](VmwareTargetParamsForRecoverVAppRenameRecoveredVAppsParams.md) |  | [optional] 
**RenameRecoveredVmsParams** | Pointer to [**NullableAcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams**](AcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams.md) |  | [optional] 
**VlanConfig** | Pointer to [**NullableAcropolisTargetParamsForRecoverVmVlanConfig**](AcropolisTargetParamsForRecoverVmVlanConfig.md) |  | [optional] 

## Methods

### NewRecoverVmwareVAppParamsVmwareTargetParams

`func NewRecoverVmwareVAppParamsVmwareTargetParams() *RecoverVmwareVAppParamsVmwareTargetParams`

NewRecoverVmwareVAppParamsVmwareTargetParams instantiates a new RecoverVmwareVAppParamsVmwareTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverVmwareVAppParamsVmwareTargetParamsWithDefaults

`func NewRecoverVmwareVAppParamsVmwareTargetParamsWithDefaults() *RecoverVmwareVAppParamsVmwareTargetParams`

NewRecoverVmwareVAppParamsVmwareTargetParamsWithDefaults instantiates a new RecoverVmwareVAppParamsVmwareTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttemptDifferentialRestore

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) GetAttemptDifferentialRestore() bool`

GetAttemptDifferentialRestore returns the AttemptDifferentialRestore field if non-nil, zero value otherwise.

### GetAttemptDifferentialRestoreOk

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) GetAttemptDifferentialRestoreOk() (*bool, bool)`

GetAttemptDifferentialRestoreOk returns a tuple with the AttemptDifferentialRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptDifferentialRestore

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) SetAttemptDifferentialRestore(v bool)`

SetAttemptDifferentialRestore sets AttemptDifferentialRestore field to given value.

### HasAttemptDifferentialRestore

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) HasAttemptDifferentialRestore() bool`

HasAttemptDifferentialRestore returns a boolean if a field has been set.

### SetAttemptDifferentialRestoreNil

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) SetAttemptDifferentialRestoreNil(b bool)`

 SetAttemptDifferentialRestoreNil sets the value for AttemptDifferentialRestore to be an explicit nil

### UnsetAttemptDifferentialRestore
`func (o *RecoverVmwareVAppParamsVmwareTargetParams) UnsetAttemptDifferentialRestore()`

UnsetAttemptDifferentialRestore ensures that no value is present for AttemptDifferentialRestore, not even an explicit nil
### GetContinueOnError

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *RecoverVmwareVAppParamsVmwareTargetParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetDiskProvisionType

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) GetDiskProvisionType() string`

GetDiskProvisionType returns the DiskProvisionType field if non-nil, zero value otherwise.

### GetDiskProvisionTypeOk

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) GetDiskProvisionTypeOk() (*string, bool)`

GetDiskProvisionTypeOk returns a tuple with the DiskProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskProvisionType

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) SetDiskProvisionType(v string)`

SetDiskProvisionType sets DiskProvisionType field to given value.

### HasDiskProvisionType

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) HasDiskProvisionType() bool`

HasDiskProvisionType returns a boolean if a field has been set.

### SetDiskProvisionTypeNil

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) SetDiskProvisionTypeNil(b bool)`

 SetDiskProvisionTypeNil sets the value for DiskProvisionType to be an explicit nil

### UnsetDiskProvisionType
`func (o *RecoverVmwareVAppParamsVmwareTargetParams) UnsetDiskProvisionType()`

UnsetDiskProvisionType ensures that no value is present for DiskProvisionType, not even an explicit nil
### GetEnableNBDSSLFallback

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) GetEnableNBDSSLFallback() bool`

GetEnableNBDSSLFallback returns the EnableNBDSSLFallback field if non-nil, zero value otherwise.

### GetEnableNBDSSLFallbackOk

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) GetEnableNBDSSLFallbackOk() (*bool, bool)`

GetEnableNBDSSLFallbackOk returns a tuple with the EnableNBDSSLFallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNBDSSLFallback

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) SetEnableNBDSSLFallback(v bool)`

SetEnableNBDSSLFallback sets EnableNBDSSLFallback field to given value.

### HasEnableNBDSSLFallback

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) HasEnableNBDSSLFallback() bool`

HasEnableNBDSSLFallback returns a boolean if a field has been set.

### SetEnableNBDSSLFallbackNil

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) SetEnableNBDSSLFallbackNil(b bool)`

 SetEnableNBDSSLFallbackNil sets the value for EnableNBDSSLFallback to be an explicit nil

### UnsetEnableNBDSSLFallback
`func (o *RecoverVmwareVAppParamsVmwareTargetParams) UnsetEnableNBDSSLFallback()`

UnsetEnableNBDSSLFallback ensures that no value is present for EnableNBDSSLFallback, not even an explicit nil
### GetLeverageSanTransport

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) GetLeverageSanTransport() bool`

GetLeverageSanTransport returns the LeverageSanTransport field if non-nil, zero value otherwise.

### GetLeverageSanTransportOk

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) GetLeverageSanTransportOk() (*bool, bool)`

GetLeverageSanTransportOk returns a tuple with the LeverageSanTransport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeverageSanTransport

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) SetLeverageSanTransport(v bool)`

SetLeverageSanTransport sets LeverageSanTransport field to given value.

### HasLeverageSanTransport

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) HasLeverageSanTransport() bool`

HasLeverageSanTransport returns a boolean if a field has been set.

### SetLeverageSanTransportNil

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) SetLeverageSanTransportNil(b bool)`

 SetLeverageSanTransportNil sets the value for LeverageSanTransport to be an explicit nil

### UnsetLeverageSanTransport
`func (o *RecoverVmwareVAppParamsVmwareTargetParams) UnsetLeverageSanTransport()`

UnsetLeverageSanTransport ensures that no value is present for LeverageSanTransport, not even an explicit nil
### GetPowerOnVms

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) GetPowerOnVms() bool`

GetPowerOnVms returns the PowerOnVms field if non-nil, zero value otherwise.

### GetPowerOnVmsOk

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) GetPowerOnVmsOk() (*bool, bool)`

GetPowerOnVmsOk returns a tuple with the PowerOnVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnVms

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) SetPowerOnVms(v bool)`

SetPowerOnVms sets PowerOnVms field to given value.

### HasPowerOnVms

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) HasPowerOnVms() bool`

HasPowerOnVms returns a boolean if a field has been set.

### SetPowerOnVmsNil

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) SetPowerOnVmsNil(b bool)`

 SetPowerOnVmsNil sets the value for PowerOnVms to be an explicit nil

### UnsetPowerOnVms
`func (o *RecoverVmwareVAppParamsVmwareTargetParams) UnsetPowerOnVms()`

UnsetPowerOnVms ensures that no value is present for PowerOnVms, not even an explicit nil
### GetRecoveryProcessType

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) GetRecoveryProcessType() string`

GetRecoveryProcessType returns the RecoveryProcessType field if non-nil, zero value otherwise.

### GetRecoveryProcessTypeOk

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) GetRecoveryProcessTypeOk() (*string, bool)`

GetRecoveryProcessTypeOk returns a tuple with the RecoveryProcessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryProcessType

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) SetRecoveryProcessType(v string)`

SetRecoveryProcessType sets RecoveryProcessType field to given value.

### HasRecoveryProcessType

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) HasRecoveryProcessType() bool`

HasRecoveryProcessType returns a boolean if a field has been set.

### GetRecoveryTargetConfig

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) GetRecoveryTargetConfig() VmwareTargetParamsForRecoverVAppRecoveryTargetConfig`

GetRecoveryTargetConfig returns the RecoveryTargetConfig field if non-nil, zero value otherwise.

### GetRecoveryTargetConfigOk

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) GetRecoveryTargetConfigOk() (*VmwareTargetParamsForRecoverVAppRecoveryTargetConfig, bool)`

GetRecoveryTargetConfigOk returns a tuple with the RecoveryTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTargetConfig

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) SetRecoveryTargetConfig(v VmwareTargetParamsForRecoverVAppRecoveryTargetConfig)`

SetRecoveryTargetConfig sets RecoveryTargetConfig field to given value.

### HasRecoveryTargetConfig

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) HasRecoveryTargetConfig() bool`

HasRecoveryTargetConfig returns a boolean if a field has been set.

### SetRecoveryTargetConfigNil

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) SetRecoveryTargetConfigNil(b bool)`

 SetRecoveryTargetConfigNil sets the value for RecoveryTargetConfig to be an explicit nil

### UnsetRecoveryTargetConfig
`func (o *RecoverVmwareVAppParamsVmwareTargetParams) UnsetRecoveryTargetConfig()`

UnsetRecoveryTargetConfig ensures that no value is present for RecoveryTargetConfig, not even an explicit nil
### GetRenameRecoveredVAppsParams

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) GetRenameRecoveredVAppsParams() VmwareTargetParamsForRecoverVAppRenameRecoveredVAppsParams`

GetRenameRecoveredVAppsParams returns the RenameRecoveredVAppsParams field if non-nil, zero value otherwise.

### GetRenameRecoveredVAppsParamsOk

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) GetRenameRecoveredVAppsParamsOk() (*VmwareTargetParamsForRecoverVAppRenameRecoveredVAppsParams, bool)`

GetRenameRecoveredVAppsParamsOk returns a tuple with the RenameRecoveredVAppsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameRecoveredVAppsParams

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) SetRenameRecoveredVAppsParams(v VmwareTargetParamsForRecoverVAppRenameRecoveredVAppsParams)`

SetRenameRecoveredVAppsParams sets RenameRecoveredVAppsParams field to given value.

### HasRenameRecoveredVAppsParams

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) HasRenameRecoveredVAppsParams() bool`

HasRenameRecoveredVAppsParams returns a boolean if a field has been set.

### SetRenameRecoveredVAppsParamsNil

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) SetRenameRecoveredVAppsParamsNil(b bool)`

 SetRenameRecoveredVAppsParamsNil sets the value for RenameRecoveredVAppsParams to be an explicit nil

### UnsetRenameRecoveredVAppsParams
`func (o *RecoverVmwareVAppParamsVmwareTargetParams) UnsetRenameRecoveredVAppsParams()`

UnsetRenameRecoveredVAppsParams ensures that no value is present for RenameRecoveredVAppsParams, not even an explicit nil
### GetRenameRecoveredVmsParams

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) GetRenameRecoveredVmsParams() AcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams`

GetRenameRecoveredVmsParams returns the RenameRecoveredVmsParams field if non-nil, zero value otherwise.

### GetRenameRecoveredVmsParamsOk

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) GetRenameRecoveredVmsParamsOk() (*AcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams, bool)`

GetRenameRecoveredVmsParamsOk returns a tuple with the RenameRecoveredVmsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameRecoveredVmsParams

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) SetRenameRecoveredVmsParams(v AcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams)`

SetRenameRecoveredVmsParams sets RenameRecoveredVmsParams field to given value.

### HasRenameRecoveredVmsParams

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) HasRenameRecoveredVmsParams() bool`

HasRenameRecoveredVmsParams returns a boolean if a field has been set.

### SetRenameRecoveredVmsParamsNil

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) SetRenameRecoveredVmsParamsNil(b bool)`

 SetRenameRecoveredVmsParamsNil sets the value for RenameRecoveredVmsParams to be an explicit nil

### UnsetRenameRecoveredVmsParams
`func (o *RecoverVmwareVAppParamsVmwareTargetParams) UnsetRenameRecoveredVmsParams()`

UnsetRenameRecoveredVmsParams ensures that no value is present for RenameRecoveredVmsParams, not even an explicit nil
### GetVlanConfig

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) GetVlanConfig() AcropolisTargetParamsForRecoverVmVlanConfig`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) GetVlanConfigOk() (*AcropolisTargetParamsForRecoverVmVlanConfig, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) SetVlanConfig(v AcropolisTargetParamsForRecoverVmVlanConfig)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.

### SetVlanConfigNil

`func (o *RecoverVmwareVAppParamsVmwareTargetParams) SetVlanConfigNil(b bool)`

 SetVlanConfigNil sets the value for VlanConfig to be an explicit nil

### UnsetVlanConfig
`func (o *RecoverVmwareVAppParamsVmwareTargetParams) UnsetVlanConfig()`

UnsetVlanConfig ensures that no value is present for VlanConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


