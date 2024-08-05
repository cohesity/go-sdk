# VmwareTargetParamsForRecoverVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttemptDifferentialRestore** | Pointer to **NullableBool** | Specifies whether to attempt differential restore. | [optional] 
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering other vms if one of vms failed to recover. Default value is false. | [optional] 
**DiskProvisionType** | Pointer to **NullableString** | Specifies the Virtual Disk Provisioning Policies for Vmware VM | [optional] 
**EnableNBDSSLFallback** | Pointer to **NullableBool** | If this field is set to true and SAN transport recovery fails, then recovery will fallback to use NBDSSL transport. This field only applies if &#39;leverageSanTransport&#39; is set to true. | [optional] 
**IsMultiStageRestore** | Pointer to **NullableBool** | Specifies whether this is a multistage restore which is used for migration/hot-standby purpose. | [optional] 
**LeverageSanTransport** | Pointer to **NullableBool** | Specifies whether to enable SAN transport for copy recovery or not | [optional] 
**OverwriteExistingVm** | Pointer to **NullableBool** | Specifies whether to overwrite the VM at the target location. This is a data destructive operation and if this is selected, the original VM may no longer be accessible. This option is only applicable if renameRecoveredVmParams is null and powerOffAndRenameExistingVm is false. This option is not supported for vApp or vApp template recoveries. Default value is false. | [optional] 
**PowerOffAndRenameExistingVm** | Pointer to **NullableBool** | Specifies whether to power off and mark the VM at the target location as deprecated. As an example, &lt;vm_name&gt; will be renamed to deprecated::&lt;vm_name&gt;, and a new VM with the name &lt;vm_name&gt; in place of the now deprecated VM. Both deprecated::&lt;vm_name&gt; and &lt;vm_name&gt; will exist on the primary, but the corresponding protection job will only backup &lt;vm_name&gt; on its next run. Only applicable if renameRecoveredVmParams is null and overwriteExistingVm is false. This option is not supported for vApp or vApp template recoveries. Default value is false. | [optional] 
**PowerOnVms** | Pointer to **NullableBool** | Specifies whether to power on vms after recovery. If not specified, or false, recovered vms will be in powered off state. | [optional] 
**RecoveryProcessType** | Pointer to **string** | Specifies type of Recovery Process to be used. InstantRecovery/CopyRecovery etc... Default value is InstantRecovery. | [optional] 
**RecoveryTargetConfig** | Pointer to [**NullableVmwareTargetParamsForRecoverVMRecoveryTargetConfig**](VmwareTargetParamsForRecoverVMRecoveryTargetConfig.md) |  | [optional] 
**RenameRecoveredVmsParams** | Pointer to [**NullableAcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams**](AcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams.md) |  | [optional] 
**VlanConfig** | Pointer to [**NullableRecoverKubernetesNamespaceParamsVlanConfig**](RecoverKubernetesNamespaceParamsVlanConfig.md) |  | [optional] 

## Methods

### NewVmwareTargetParamsForRecoverVM

`func NewVmwareTargetParamsForRecoverVM() *VmwareTargetParamsForRecoverVM`

NewVmwareTargetParamsForRecoverVM instantiates a new VmwareTargetParamsForRecoverVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareTargetParamsForRecoverVMWithDefaults

`func NewVmwareTargetParamsForRecoverVMWithDefaults() *VmwareTargetParamsForRecoverVM`

NewVmwareTargetParamsForRecoverVMWithDefaults instantiates a new VmwareTargetParamsForRecoverVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttemptDifferentialRestore

`func (o *VmwareTargetParamsForRecoverVM) GetAttemptDifferentialRestore() bool`

GetAttemptDifferentialRestore returns the AttemptDifferentialRestore field if non-nil, zero value otherwise.

### GetAttemptDifferentialRestoreOk

`func (o *VmwareTargetParamsForRecoverVM) GetAttemptDifferentialRestoreOk() (*bool, bool)`

GetAttemptDifferentialRestoreOk returns a tuple with the AttemptDifferentialRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptDifferentialRestore

`func (o *VmwareTargetParamsForRecoverVM) SetAttemptDifferentialRestore(v bool)`

SetAttemptDifferentialRestore sets AttemptDifferentialRestore field to given value.

### HasAttemptDifferentialRestore

`func (o *VmwareTargetParamsForRecoverVM) HasAttemptDifferentialRestore() bool`

HasAttemptDifferentialRestore returns a boolean if a field has been set.

### SetAttemptDifferentialRestoreNil

`func (o *VmwareTargetParamsForRecoverVM) SetAttemptDifferentialRestoreNil(b bool)`

 SetAttemptDifferentialRestoreNil sets the value for AttemptDifferentialRestore to be an explicit nil

### UnsetAttemptDifferentialRestore
`func (o *VmwareTargetParamsForRecoverVM) UnsetAttemptDifferentialRestore()`

UnsetAttemptDifferentialRestore ensures that no value is present for AttemptDifferentialRestore, not even an explicit nil
### GetContinueOnError

`func (o *VmwareTargetParamsForRecoverVM) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *VmwareTargetParamsForRecoverVM) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *VmwareTargetParamsForRecoverVM) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *VmwareTargetParamsForRecoverVM) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *VmwareTargetParamsForRecoverVM) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *VmwareTargetParamsForRecoverVM) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetDiskProvisionType

`func (o *VmwareTargetParamsForRecoverVM) GetDiskProvisionType() string`

GetDiskProvisionType returns the DiskProvisionType field if non-nil, zero value otherwise.

### GetDiskProvisionTypeOk

`func (o *VmwareTargetParamsForRecoverVM) GetDiskProvisionTypeOk() (*string, bool)`

GetDiskProvisionTypeOk returns a tuple with the DiskProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskProvisionType

`func (o *VmwareTargetParamsForRecoverVM) SetDiskProvisionType(v string)`

SetDiskProvisionType sets DiskProvisionType field to given value.

### HasDiskProvisionType

`func (o *VmwareTargetParamsForRecoverVM) HasDiskProvisionType() bool`

HasDiskProvisionType returns a boolean if a field has been set.

### SetDiskProvisionTypeNil

`func (o *VmwareTargetParamsForRecoverVM) SetDiskProvisionTypeNil(b bool)`

 SetDiskProvisionTypeNil sets the value for DiskProvisionType to be an explicit nil

### UnsetDiskProvisionType
`func (o *VmwareTargetParamsForRecoverVM) UnsetDiskProvisionType()`

UnsetDiskProvisionType ensures that no value is present for DiskProvisionType, not even an explicit nil
### GetEnableNBDSSLFallback

`func (o *VmwareTargetParamsForRecoverVM) GetEnableNBDSSLFallback() bool`

GetEnableNBDSSLFallback returns the EnableNBDSSLFallback field if non-nil, zero value otherwise.

### GetEnableNBDSSLFallbackOk

`func (o *VmwareTargetParamsForRecoverVM) GetEnableNBDSSLFallbackOk() (*bool, bool)`

GetEnableNBDSSLFallbackOk returns a tuple with the EnableNBDSSLFallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNBDSSLFallback

`func (o *VmwareTargetParamsForRecoverVM) SetEnableNBDSSLFallback(v bool)`

SetEnableNBDSSLFallback sets EnableNBDSSLFallback field to given value.

### HasEnableNBDSSLFallback

`func (o *VmwareTargetParamsForRecoverVM) HasEnableNBDSSLFallback() bool`

HasEnableNBDSSLFallback returns a boolean if a field has been set.

### SetEnableNBDSSLFallbackNil

`func (o *VmwareTargetParamsForRecoverVM) SetEnableNBDSSLFallbackNil(b bool)`

 SetEnableNBDSSLFallbackNil sets the value for EnableNBDSSLFallback to be an explicit nil

### UnsetEnableNBDSSLFallback
`func (o *VmwareTargetParamsForRecoverVM) UnsetEnableNBDSSLFallback()`

UnsetEnableNBDSSLFallback ensures that no value is present for EnableNBDSSLFallback, not even an explicit nil
### GetIsMultiStageRestore

`func (o *VmwareTargetParamsForRecoverVM) GetIsMultiStageRestore() bool`

GetIsMultiStageRestore returns the IsMultiStageRestore field if non-nil, zero value otherwise.

### GetIsMultiStageRestoreOk

`func (o *VmwareTargetParamsForRecoverVM) GetIsMultiStageRestoreOk() (*bool, bool)`

GetIsMultiStageRestoreOk returns a tuple with the IsMultiStageRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiStageRestore

`func (o *VmwareTargetParamsForRecoverVM) SetIsMultiStageRestore(v bool)`

SetIsMultiStageRestore sets IsMultiStageRestore field to given value.

### HasIsMultiStageRestore

`func (o *VmwareTargetParamsForRecoverVM) HasIsMultiStageRestore() bool`

HasIsMultiStageRestore returns a boolean if a field has been set.

### SetIsMultiStageRestoreNil

`func (o *VmwareTargetParamsForRecoverVM) SetIsMultiStageRestoreNil(b bool)`

 SetIsMultiStageRestoreNil sets the value for IsMultiStageRestore to be an explicit nil

### UnsetIsMultiStageRestore
`func (o *VmwareTargetParamsForRecoverVM) UnsetIsMultiStageRestore()`

UnsetIsMultiStageRestore ensures that no value is present for IsMultiStageRestore, not even an explicit nil
### GetLeverageSanTransport

`func (o *VmwareTargetParamsForRecoverVM) GetLeverageSanTransport() bool`

GetLeverageSanTransport returns the LeverageSanTransport field if non-nil, zero value otherwise.

### GetLeverageSanTransportOk

`func (o *VmwareTargetParamsForRecoverVM) GetLeverageSanTransportOk() (*bool, bool)`

GetLeverageSanTransportOk returns a tuple with the LeverageSanTransport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeverageSanTransport

`func (o *VmwareTargetParamsForRecoverVM) SetLeverageSanTransport(v bool)`

SetLeverageSanTransport sets LeverageSanTransport field to given value.

### HasLeverageSanTransport

`func (o *VmwareTargetParamsForRecoverVM) HasLeverageSanTransport() bool`

HasLeverageSanTransport returns a boolean if a field has been set.

### SetLeverageSanTransportNil

`func (o *VmwareTargetParamsForRecoverVM) SetLeverageSanTransportNil(b bool)`

 SetLeverageSanTransportNil sets the value for LeverageSanTransport to be an explicit nil

### UnsetLeverageSanTransport
`func (o *VmwareTargetParamsForRecoverVM) UnsetLeverageSanTransport()`

UnsetLeverageSanTransport ensures that no value is present for LeverageSanTransport, not even an explicit nil
### GetOverwriteExistingVm

`func (o *VmwareTargetParamsForRecoverVM) GetOverwriteExistingVm() bool`

GetOverwriteExistingVm returns the OverwriteExistingVm field if non-nil, zero value otherwise.

### GetOverwriteExistingVmOk

`func (o *VmwareTargetParamsForRecoverVM) GetOverwriteExistingVmOk() (*bool, bool)`

GetOverwriteExistingVmOk returns a tuple with the OverwriteExistingVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteExistingVm

`func (o *VmwareTargetParamsForRecoverVM) SetOverwriteExistingVm(v bool)`

SetOverwriteExistingVm sets OverwriteExistingVm field to given value.

### HasOverwriteExistingVm

`func (o *VmwareTargetParamsForRecoverVM) HasOverwriteExistingVm() bool`

HasOverwriteExistingVm returns a boolean if a field has been set.

### SetOverwriteExistingVmNil

`func (o *VmwareTargetParamsForRecoverVM) SetOverwriteExistingVmNil(b bool)`

 SetOverwriteExistingVmNil sets the value for OverwriteExistingVm to be an explicit nil

### UnsetOverwriteExistingVm
`func (o *VmwareTargetParamsForRecoverVM) UnsetOverwriteExistingVm()`

UnsetOverwriteExistingVm ensures that no value is present for OverwriteExistingVm, not even an explicit nil
### GetPowerOffAndRenameExistingVm

`func (o *VmwareTargetParamsForRecoverVM) GetPowerOffAndRenameExistingVm() bool`

GetPowerOffAndRenameExistingVm returns the PowerOffAndRenameExistingVm field if non-nil, zero value otherwise.

### GetPowerOffAndRenameExistingVmOk

`func (o *VmwareTargetParamsForRecoverVM) GetPowerOffAndRenameExistingVmOk() (*bool, bool)`

GetPowerOffAndRenameExistingVmOk returns a tuple with the PowerOffAndRenameExistingVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOffAndRenameExistingVm

`func (o *VmwareTargetParamsForRecoverVM) SetPowerOffAndRenameExistingVm(v bool)`

SetPowerOffAndRenameExistingVm sets PowerOffAndRenameExistingVm field to given value.

### HasPowerOffAndRenameExistingVm

`func (o *VmwareTargetParamsForRecoverVM) HasPowerOffAndRenameExistingVm() bool`

HasPowerOffAndRenameExistingVm returns a boolean if a field has been set.

### SetPowerOffAndRenameExistingVmNil

`func (o *VmwareTargetParamsForRecoverVM) SetPowerOffAndRenameExistingVmNil(b bool)`

 SetPowerOffAndRenameExistingVmNil sets the value for PowerOffAndRenameExistingVm to be an explicit nil

### UnsetPowerOffAndRenameExistingVm
`func (o *VmwareTargetParamsForRecoverVM) UnsetPowerOffAndRenameExistingVm()`

UnsetPowerOffAndRenameExistingVm ensures that no value is present for PowerOffAndRenameExistingVm, not even an explicit nil
### GetPowerOnVms

`func (o *VmwareTargetParamsForRecoverVM) GetPowerOnVms() bool`

GetPowerOnVms returns the PowerOnVms field if non-nil, zero value otherwise.

### GetPowerOnVmsOk

`func (o *VmwareTargetParamsForRecoverVM) GetPowerOnVmsOk() (*bool, bool)`

GetPowerOnVmsOk returns a tuple with the PowerOnVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnVms

`func (o *VmwareTargetParamsForRecoverVM) SetPowerOnVms(v bool)`

SetPowerOnVms sets PowerOnVms field to given value.

### HasPowerOnVms

`func (o *VmwareTargetParamsForRecoverVM) HasPowerOnVms() bool`

HasPowerOnVms returns a boolean if a field has been set.

### SetPowerOnVmsNil

`func (o *VmwareTargetParamsForRecoverVM) SetPowerOnVmsNil(b bool)`

 SetPowerOnVmsNil sets the value for PowerOnVms to be an explicit nil

### UnsetPowerOnVms
`func (o *VmwareTargetParamsForRecoverVM) UnsetPowerOnVms()`

UnsetPowerOnVms ensures that no value is present for PowerOnVms, not even an explicit nil
### GetRecoveryProcessType

`func (o *VmwareTargetParamsForRecoverVM) GetRecoveryProcessType() string`

GetRecoveryProcessType returns the RecoveryProcessType field if non-nil, zero value otherwise.

### GetRecoveryProcessTypeOk

`func (o *VmwareTargetParamsForRecoverVM) GetRecoveryProcessTypeOk() (*string, bool)`

GetRecoveryProcessTypeOk returns a tuple with the RecoveryProcessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryProcessType

`func (o *VmwareTargetParamsForRecoverVM) SetRecoveryProcessType(v string)`

SetRecoveryProcessType sets RecoveryProcessType field to given value.

### HasRecoveryProcessType

`func (o *VmwareTargetParamsForRecoverVM) HasRecoveryProcessType() bool`

HasRecoveryProcessType returns a boolean if a field has been set.

### GetRecoveryTargetConfig

`func (o *VmwareTargetParamsForRecoverVM) GetRecoveryTargetConfig() VmwareTargetParamsForRecoverVMRecoveryTargetConfig`

GetRecoveryTargetConfig returns the RecoveryTargetConfig field if non-nil, zero value otherwise.

### GetRecoveryTargetConfigOk

`func (o *VmwareTargetParamsForRecoverVM) GetRecoveryTargetConfigOk() (*VmwareTargetParamsForRecoverVMRecoveryTargetConfig, bool)`

GetRecoveryTargetConfigOk returns a tuple with the RecoveryTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTargetConfig

`func (o *VmwareTargetParamsForRecoverVM) SetRecoveryTargetConfig(v VmwareTargetParamsForRecoverVMRecoveryTargetConfig)`

SetRecoveryTargetConfig sets RecoveryTargetConfig field to given value.

### HasRecoveryTargetConfig

`func (o *VmwareTargetParamsForRecoverVM) HasRecoveryTargetConfig() bool`

HasRecoveryTargetConfig returns a boolean if a field has been set.

### SetRecoveryTargetConfigNil

`func (o *VmwareTargetParamsForRecoverVM) SetRecoveryTargetConfigNil(b bool)`

 SetRecoveryTargetConfigNil sets the value for RecoveryTargetConfig to be an explicit nil

### UnsetRecoveryTargetConfig
`func (o *VmwareTargetParamsForRecoverVM) UnsetRecoveryTargetConfig()`

UnsetRecoveryTargetConfig ensures that no value is present for RecoveryTargetConfig, not even an explicit nil
### GetRenameRecoveredVmsParams

`func (o *VmwareTargetParamsForRecoverVM) GetRenameRecoveredVmsParams() AcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams`

GetRenameRecoveredVmsParams returns the RenameRecoveredVmsParams field if non-nil, zero value otherwise.

### GetRenameRecoveredVmsParamsOk

`func (o *VmwareTargetParamsForRecoverVM) GetRenameRecoveredVmsParamsOk() (*AcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams, bool)`

GetRenameRecoveredVmsParamsOk returns a tuple with the RenameRecoveredVmsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameRecoveredVmsParams

`func (o *VmwareTargetParamsForRecoverVM) SetRenameRecoveredVmsParams(v AcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams)`

SetRenameRecoveredVmsParams sets RenameRecoveredVmsParams field to given value.

### HasRenameRecoveredVmsParams

`func (o *VmwareTargetParamsForRecoverVM) HasRenameRecoveredVmsParams() bool`

HasRenameRecoveredVmsParams returns a boolean if a field has been set.

### SetRenameRecoveredVmsParamsNil

`func (o *VmwareTargetParamsForRecoverVM) SetRenameRecoveredVmsParamsNil(b bool)`

 SetRenameRecoveredVmsParamsNil sets the value for RenameRecoveredVmsParams to be an explicit nil

### UnsetRenameRecoveredVmsParams
`func (o *VmwareTargetParamsForRecoverVM) UnsetRenameRecoveredVmsParams()`

UnsetRenameRecoveredVmsParams ensures that no value is present for RenameRecoveredVmsParams, not even an explicit nil
### GetVlanConfig

`func (o *VmwareTargetParamsForRecoverVM) GetVlanConfig() RecoverKubernetesNamespaceParamsVlanConfig`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *VmwareTargetParamsForRecoverVM) GetVlanConfigOk() (*RecoverKubernetesNamespaceParamsVlanConfig, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *VmwareTargetParamsForRecoverVM) SetVlanConfig(v RecoverKubernetesNamespaceParamsVlanConfig)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *VmwareTargetParamsForRecoverVM) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.

### SetVlanConfigNil

`func (o *VmwareTargetParamsForRecoverVM) SetVlanConfigNil(b bool)`

 SetVlanConfigNil sets the value for VlanConfig to be an explicit nil

### UnsetVlanConfig
`func (o *VmwareTargetParamsForRecoverVM) UnsetVlanConfig()`

UnsetVlanConfig ensures that no value is present for VlanConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


