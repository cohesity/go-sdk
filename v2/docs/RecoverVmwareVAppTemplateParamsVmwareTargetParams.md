# RecoverVmwareVAppTemplateParamsVmwareTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttemptDifferentialRestore** | Pointer to **NullableBool** | Specifies whether to attempt differential restore. | [optional] 
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering other vms if one of vms failed to recover. Default value is false. | [optional] 
**DiskProvisionType** | Pointer to **NullableString** | Specifies the Virtual Disk Provisioning Policies for Vmware VM | [optional] 
**EnableNBDSSLFallback** | Pointer to **NullableBool** | If this field is set to true and SAN transport recovery fails, then recovery will fallback to use NBDSSL transport. This field only applies if &#39;leverageSanTransport&#39; is set to true. | [optional] 
**LeverageSanTransport** | Pointer to **NullableBool** | Specifies whether to enable SAN transport for copy recovery or not | [optional] 
**RecoveryProcessType** | Pointer to **string** | Specifies type of Recovery Process to be used. InstantRecovery/CopyRecovery etc... Default value is InstantRecovery. | [optional] 
**RecoveryTargetConfig** | Pointer to [**NullableVmwareTargetParamsForRecoverVAppTemplateRecoveryTargetConfig**](VmwareTargetParamsForRecoverVAppTemplateRecoveryTargetConfig.md) |  | [optional] 
**RenameRecoveredVAppTemplateParams** | Pointer to [**NullableVmwareTargetParamsForRecoverVAppTemplateRenameRecoveredVAppTemplateParams**](VmwareTargetParamsForRecoverVAppTemplateRenameRecoveredVAppTemplateParams.md) |  | [optional] 
**RenameRecoveredVmsParams** | Pointer to [**NullableAcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams**](AcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams.md) |  | [optional] 
**VlanConfig** | Pointer to [**NullableAcropolisTargetParamsForRecoverVmVlanConfig**](AcropolisTargetParamsForRecoverVmVlanConfig.md) |  | [optional] 

## Methods

### NewRecoverVmwareVAppTemplateParamsVmwareTargetParams

`func NewRecoverVmwareVAppTemplateParamsVmwareTargetParams() *RecoverVmwareVAppTemplateParamsVmwareTargetParams`

NewRecoverVmwareVAppTemplateParamsVmwareTargetParams instantiates a new RecoverVmwareVAppTemplateParamsVmwareTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverVmwareVAppTemplateParamsVmwareTargetParamsWithDefaults

`func NewRecoverVmwareVAppTemplateParamsVmwareTargetParamsWithDefaults() *RecoverVmwareVAppTemplateParamsVmwareTargetParams`

NewRecoverVmwareVAppTemplateParamsVmwareTargetParamsWithDefaults instantiates a new RecoverVmwareVAppTemplateParamsVmwareTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttemptDifferentialRestore

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) GetAttemptDifferentialRestore() bool`

GetAttemptDifferentialRestore returns the AttemptDifferentialRestore field if non-nil, zero value otherwise.

### GetAttemptDifferentialRestoreOk

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) GetAttemptDifferentialRestoreOk() (*bool, bool)`

GetAttemptDifferentialRestoreOk returns a tuple with the AttemptDifferentialRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptDifferentialRestore

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) SetAttemptDifferentialRestore(v bool)`

SetAttemptDifferentialRestore sets AttemptDifferentialRestore field to given value.

### HasAttemptDifferentialRestore

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) HasAttemptDifferentialRestore() bool`

HasAttemptDifferentialRestore returns a boolean if a field has been set.

### SetAttemptDifferentialRestoreNil

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) SetAttemptDifferentialRestoreNil(b bool)`

 SetAttemptDifferentialRestoreNil sets the value for AttemptDifferentialRestore to be an explicit nil

### UnsetAttemptDifferentialRestore
`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) UnsetAttemptDifferentialRestore()`

UnsetAttemptDifferentialRestore ensures that no value is present for AttemptDifferentialRestore, not even an explicit nil
### GetContinueOnError

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetDiskProvisionType

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) GetDiskProvisionType() string`

GetDiskProvisionType returns the DiskProvisionType field if non-nil, zero value otherwise.

### GetDiskProvisionTypeOk

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) GetDiskProvisionTypeOk() (*string, bool)`

GetDiskProvisionTypeOk returns a tuple with the DiskProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskProvisionType

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) SetDiskProvisionType(v string)`

SetDiskProvisionType sets DiskProvisionType field to given value.

### HasDiskProvisionType

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) HasDiskProvisionType() bool`

HasDiskProvisionType returns a boolean if a field has been set.

### SetDiskProvisionTypeNil

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) SetDiskProvisionTypeNil(b bool)`

 SetDiskProvisionTypeNil sets the value for DiskProvisionType to be an explicit nil

### UnsetDiskProvisionType
`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) UnsetDiskProvisionType()`

UnsetDiskProvisionType ensures that no value is present for DiskProvisionType, not even an explicit nil
### GetEnableNBDSSLFallback

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) GetEnableNBDSSLFallback() bool`

GetEnableNBDSSLFallback returns the EnableNBDSSLFallback field if non-nil, zero value otherwise.

### GetEnableNBDSSLFallbackOk

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) GetEnableNBDSSLFallbackOk() (*bool, bool)`

GetEnableNBDSSLFallbackOk returns a tuple with the EnableNBDSSLFallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNBDSSLFallback

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) SetEnableNBDSSLFallback(v bool)`

SetEnableNBDSSLFallback sets EnableNBDSSLFallback field to given value.

### HasEnableNBDSSLFallback

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) HasEnableNBDSSLFallback() bool`

HasEnableNBDSSLFallback returns a boolean if a field has been set.

### SetEnableNBDSSLFallbackNil

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) SetEnableNBDSSLFallbackNil(b bool)`

 SetEnableNBDSSLFallbackNil sets the value for EnableNBDSSLFallback to be an explicit nil

### UnsetEnableNBDSSLFallback
`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) UnsetEnableNBDSSLFallback()`

UnsetEnableNBDSSLFallback ensures that no value is present for EnableNBDSSLFallback, not even an explicit nil
### GetLeverageSanTransport

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) GetLeverageSanTransport() bool`

GetLeverageSanTransport returns the LeverageSanTransport field if non-nil, zero value otherwise.

### GetLeverageSanTransportOk

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) GetLeverageSanTransportOk() (*bool, bool)`

GetLeverageSanTransportOk returns a tuple with the LeverageSanTransport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeverageSanTransport

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) SetLeverageSanTransport(v bool)`

SetLeverageSanTransport sets LeverageSanTransport field to given value.

### HasLeverageSanTransport

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) HasLeverageSanTransport() bool`

HasLeverageSanTransport returns a boolean if a field has been set.

### SetLeverageSanTransportNil

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) SetLeverageSanTransportNil(b bool)`

 SetLeverageSanTransportNil sets the value for LeverageSanTransport to be an explicit nil

### UnsetLeverageSanTransport
`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) UnsetLeverageSanTransport()`

UnsetLeverageSanTransport ensures that no value is present for LeverageSanTransport, not even an explicit nil
### GetRecoveryProcessType

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) GetRecoveryProcessType() string`

GetRecoveryProcessType returns the RecoveryProcessType field if non-nil, zero value otherwise.

### GetRecoveryProcessTypeOk

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) GetRecoveryProcessTypeOk() (*string, bool)`

GetRecoveryProcessTypeOk returns a tuple with the RecoveryProcessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryProcessType

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) SetRecoveryProcessType(v string)`

SetRecoveryProcessType sets RecoveryProcessType field to given value.

### HasRecoveryProcessType

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) HasRecoveryProcessType() bool`

HasRecoveryProcessType returns a boolean if a field has been set.

### GetRecoveryTargetConfig

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) GetRecoveryTargetConfig() VmwareTargetParamsForRecoverVAppTemplateRecoveryTargetConfig`

GetRecoveryTargetConfig returns the RecoveryTargetConfig field if non-nil, zero value otherwise.

### GetRecoveryTargetConfigOk

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) GetRecoveryTargetConfigOk() (*VmwareTargetParamsForRecoverVAppTemplateRecoveryTargetConfig, bool)`

GetRecoveryTargetConfigOk returns a tuple with the RecoveryTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTargetConfig

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) SetRecoveryTargetConfig(v VmwareTargetParamsForRecoverVAppTemplateRecoveryTargetConfig)`

SetRecoveryTargetConfig sets RecoveryTargetConfig field to given value.

### HasRecoveryTargetConfig

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) HasRecoveryTargetConfig() bool`

HasRecoveryTargetConfig returns a boolean if a field has been set.

### SetRecoveryTargetConfigNil

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) SetRecoveryTargetConfigNil(b bool)`

 SetRecoveryTargetConfigNil sets the value for RecoveryTargetConfig to be an explicit nil

### UnsetRecoveryTargetConfig
`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) UnsetRecoveryTargetConfig()`

UnsetRecoveryTargetConfig ensures that no value is present for RecoveryTargetConfig, not even an explicit nil
### GetRenameRecoveredVAppTemplateParams

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) GetRenameRecoveredVAppTemplateParams() VmwareTargetParamsForRecoverVAppTemplateRenameRecoveredVAppTemplateParams`

GetRenameRecoveredVAppTemplateParams returns the RenameRecoveredVAppTemplateParams field if non-nil, zero value otherwise.

### GetRenameRecoveredVAppTemplateParamsOk

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) GetRenameRecoveredVAppTemplateParamsOk() (*VmwareTargetParamsForRecoverVAppTemplateRenameRecoveredVAppTemplateParams, bool)`

GetRenameRecoveredVAppTemplateParamsOk returns a tuple with the RenameRecoveredVAppTemplateParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameRecoveredVAppTemplateParams

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) SetRenameRecoveredVAppTemplateParams(v VmwareTargetParamsForRecoverVAppTemplateRenameRecoveredVAppTemplateParams)`

SetRenameRecoveredVAppTemplateParams sets RenameRecoveredVAppTemplateParams field to given value.

### HasRenameRecoveredVAppTemplateParams

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) HasRenameRecoveredVAppTemplateParams() bool`

HasRenameRecoveredVAppTemplateParams returns a boolean if a field has been set.

### SetRenameRecoveredVAppTemplateParamsNil

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) SetRenameRecoveredVAppTemplateParamsNil(b bool)`

 SetRenameRecoveredVAppTemplateParamsNil sets the value for RenameRecoveredVAppTemplateParams to be an explicit nil

### UnsetRenameRecoveredVAppTemplateParams
`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) UnsetRenameRecoveredVAppTemplateParams()`

UnsetRenameRecoveredVAppTemplateParams ensures that no value is present for RenameRecoveredVAppTemplateParams, not even an explicit nil
### GetRenameRecoveredVmsParams

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) GetRenameRecoveredVmsParams() AcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams`

GetRenameRecoveredVmsParams returns the RenameRecoveredVmsParams field if non-nil, zero value otherwise.

### GetRenameRecoveredVmsParamsOk

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) GetRenameRecoveredVmsParamsOk() (*AcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams, bool)`

GetRenameRecoveredVmsParamsOk returns a tuple with the RenameRecoveredVmsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameRecoveredVmsParams

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) SetRenameRecoveredVmsParams(v AcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams)`

SetRenameRecoveredVmsParams sets RenameRecoveredVmsParams field to given value.

### HasRenameRecoveredVmsParams

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) HasRenameRecoveredVmsParams() bool`

HasRenameRecoveredVmsParams returns a boolean if a field has been set.

### SetRenameRecoveredVmsParamsNil

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) SetRenameRecoveredVmsParamsNil(b bool)`

 SetRenameRecoveredVmsParamsNil sets the value for RenameRecoveredVmsParams to be an explicit nil

### UnsetRenameRecoveredVmsParams
`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) UnsetRenameRecoveredVmsParams()`

UnsetRenameRecoveredVmsParams ensures that no value is present for RenameRecoveredVmsParams, not even an explicit nil
### GetVlanConfig

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) GetVlanConfig() AcropolisTargetParamsForRecoverVmVlanConfig`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) GetVlanConfigOk() (*AcropolisTargetParamsForRecoverVmVlanConfig, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) SetVlanConfig(v AcropolisTargetParamsForRecoverVmVlanConfig)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.

### SetVlanConfigNil

`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) SetVlanConfigNil(b bool)`

 SetVlanConfigNil sets the value for VlanConfig to be an explicit nil

### UnsetVlanConfig
`func (o *RecoverVmwareVAppTemplateParamsVmwareTargetParams) UnsetVlanConfig()`

UnsetVlanConfig ensures that no value is present for VlanConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


