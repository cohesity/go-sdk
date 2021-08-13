# VmwareTargetParamsForRecoverVApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RenameRecoveredVmsParams** | Pointer to [**NullableRecoveredOrClonedVmsRenameConfig**](RecoveredOrClonedVmsRenameConfig.md) | Specifies params to rename the VMs that are recovered. If not specified, the original names of the VMs are preserved. | [optional] 
**RenameRecoveredVAppsParams** | Pointer to [**NullableRecoveredOrClonedVmsRenameConfig**](RecoveredOrClonedVmsRenameConfig.md) | Specifies params to rename the vApps that are recovered. If not specified, the original names of the vApps are preserved. | [optional] 
**RecoveryTargetConfig** | Pointer to [**NullableVmwareVAppRecoveryTargetConfig**](VmwareVAppRecoveryTargetConfig.md) | Specifies the recovery target configuration if recovery has to be done to a different location which is different from original source or to original Source with different configuration. If not specified, then the recovery of the vms will be performed to original location with all configuration parameters retained. | [optional] 
**VlanConfig** | Pointer to [**NullableRecoveryVlanConfig**](RecoveryVlanConfig.md) | Specifies VLAN Params associated with the recovered. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client&#39;s (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery. | [optional] 
**PowerOnVms** | Pointer to **NullableBool** | Specifies whether to power on vms after recovery. If not specified, or false, recovered vms will be in powered off state. | [optional] 
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering other vms if one of vms failed to recover. Default value is false. | [optional] 
**RecoveryProcessType** | Pointer to **string** | Specifies type of Recovery Process to be used. InstantRecovery/CopyRecovery etc... Default value is InstantRecovery. | [optional] 
**AttemptDifferentialRestore** | Pointer to **NullableBool** | Specifies whether to attempt differential restore. | [optional] 

## Methods

### NewVmwareTargetParamsForRecoverVApp

`func NewVmwareTargetParamsForRecoverVApp() *VmwareTargetParamsForRecoverVApp`

NewVmwareTargetParamsForRecoverVApp instantiates a new VmwareTargetParamsForRecoverVApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareTargetParamsForRecoverVAppWithDefaults

`func NewVmwareTargetParamsForRecoverVAppWithDefaults() *VmwareTargetParamsForRecoverVApp`

NewVmwareTargetParamsForRecoverVAppWithDefaults instantiates a new VmwareTargetParamsForRecoverVApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRenameRecoveredVmsParams

`func (o *VmwareTargetParamsForRecoverVApp) GetRenameRecoveredVmsParams() RecoveredOrClonedVmsRenameConfig`

GetRenameRecoveredVmsParams returns the RenameRecoveredVmsParams field if non-nil, zero value otherwise.

### GetRenameRecoveredVmsParamsOk

`func (o *VmwareTargetParamsForRecoverVApp) GetRenameRecoveredVmsParamsOk() (*RecoveredOrClonedVmsRenameConfig, bool)`

GetRenameRecoveredVmsParamsOk returns a tuple with the RenameRecoveredVmsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameRecoveredVmsParams

`func (o *VmwareTargetParamsForRecoverVApp) SetRenameRecoveredVmsParams(v RecoveredOrClonedVmsRenameConfig)`

SetRenameRecoveredVmsParams sets RenameRecoveredVmsParams field to given value.

### HasRenameRecoveredVmsParams

`func (o *VmwareTargetParamsForRecoverVApp) HasRenameRecoveredVmsParams() bool`

HasRenameRecoveredVmsParams returns a boolean if a field has been set.

### SetRenameRecoveredVmsParamsNil

`func (o *VmwareTargetParamsForRecoverVApp) SetRenameRecoveredVmsParamsNil(b bool)`

 SetRenameRecoveredVmsParamsNil sets the value for RenameRecoveredVmsParams to be an explicit nil

### UnsetRenameRecoveredVmsParams
`func (o *VmwareTargetParamsForRecoverVApp) UnsetRenameRecoveredVmsParams()`

UnsetRenameRecoveredVmsParams ensures that no value is present for RenameRecoveredVmsParams, not even an explicit nil
### GetRenameRecoveredVAppsParams

`func (o *VmwareTargetParamsForRecoverVApp) GetRenameRecoveredVAppsParams() RecoveredOrClonedVmsRenameConfig`

GetRenameRecoveredVAppsParams returns the RenameRecoveredVAppsParams field if non-nil, zero value otherwise.

### GetRenameRecoveredVAppsParamsOk

`func (o *VmwareTargetParamsForRecoverVApp) GetRenameRecoveredVAppsParamsOk() (*RecoveredOrClonedVmsRenameConfig, bool)`

GetRenameRecoveredVAppsParamsOk returns a tuple with the RenameRecoveredVAppsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameRecoveredVAppsParams

`func (o *VmwareTargetParamsForRecoverVApp) SetRenameRecoveredVAppsParams(v RecoveredOrClonedVmsRenameConfig)`

SetRenameRecoveredVAppsParams sets RenameRecoveredVAppsParams field to given value.

### HasRenameRecoveredVAppsParams

`func (o *VmwareTargetParamsForRecoverVApp) HasRenameRecoveredVAppsParams() bool`

HasRenameRecoveredVAppsParams returns a boolean if a field has been set.

### SetRenameRecoveredVAppsParamsNil

`func (o *VmwareTargetParamsForRecoverVApp) SetRenameRecoveredVAppsParamsNil(b bool)`

 SetRenameRecoveredVAppsParamsNil sets the value for RenameRecoveredVAppsParams to be an explicit nil

### UnsetRenameRecoveredVAppsParams
`func (o *VmwareTargetParamsForRecoverVApp) UnsetRenameRecoveredVAppsParams()`

UnsetRenameRecoveredVAppsParams ensures that no value is present for RenameRecoveredVAppsParams, not even an explicit nil
### GetRecoveryTargetConfig

`func (o *VmwareTargetParamsForRecoverVApp) GetRecoveryTargetConfig() VmwareVAppRecoveryTargetConfig`

GetRecoveryTargetConfig returns the RecoveryTargetConfig field if non-nil, zero value otherwise.

### GetRecoveryTargetConfigOk

`func (o *VmwareTargetParamsForRecoverVApp) GetRecoveryTargetConfigOk() (*VmwareVAppRecoveryTargetConfig, bool)`

GetRecoveryTargetConfigOk returns a tuple with the RecoveryTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTargetConfig

`func (o *VmwareTargetParamsForRecoverVApp) SetRecoveryTargetConfig(v VmwareVAppRecoveryTargetConfig)`

SetRecoveryTargetConfig sets RecoveryTargetConfig field to given value.

### HasRecoveryTargetConfig

`func (o *VmwareTargetParamsForRecoverVApp) HasRecoveryTargetConfig() bool`

HasRecoveryTargetConfig returns a boolean if a field has been set.

### SetRecoveryTargetConfigNil

`func (o *VmwareTargetParamsForRecoverVApp) SetRecoveryTargetConfigNil(b bool)`

 SetRecoveryTargetConfigNil sets the value for RecoveryTargetConfig to be an explicit nil

### UnsetRecoveryTargetConfig
`func (o *VmwareTargetParamsForRecoverVApp) UnsetRecoveryTargetConfig()`

UnsetRecoveryTargetConfig ensures that no value is present for RecoveryTargetConfig, not even an explicit nil
### GetVlanConfig

`func (o *VmwareTargetParamsForRecoverVApp) GetVlanConfig() RecoveryVlanConfig`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *VmwareTargetParamsForRecoverVApp) GetVlanConfigOk() (*RecoveryVlanConfig, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *VmwareTargetParamsForRecoverVApp) SetVlanConfig(v RecoveryVlanConfig)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *VmwareTargetParamsForRecoverVApp) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.

### SetVlanConfigNil

`func (o *VmwareTargetParamsForRecoverVApp) SetVlanConfigNil(b bool)`

 SetVlanConfigNil sets the value for VlanConfig to be an explicit nil

### UnsetVlanConfig
`func (o *VmwareTargetParamsForRecoverVApp) UnsetVlanConfig()`

UnsetVlanConfig ensures that no value is present for VlanConfig, not even an explicit nil
### GetPowerOnVms

`func (o *VmwareTargetParamsForRecoverVApp) GetPowerOnVms() bool`

GetPowerOnVms returns the PowerOnVms field if non-nil, zero value otherwise.

### GetPowerOnVmsOk

`func (o *VmwareTargetParamsForRecoverVApp) GetPowerOnVmsOk() (*bool, bool)`

GetPowerOnVmsOk returns a tuple with the PowerOnVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnVms

`func (o *VmwareTargetParamsForRecoverVApp) SetPowerOnVms(v bool)`

SetPowerOnVms sets PowerOnVms field to given value.

### HasPowerOnVms

`func (o *VmwareTargetParamsForRecoverVApp) HasPowerOnVms() bool`

HasPowerOnVms returns a boolean if a field has been set.

### SetPowerOnVmsNil

`func (o *VmwareTargetParamsForRecoverVApp) SetPowerOnVmsNil(b bool)`

 SetPowerOnVmsNil sets the value for PowerOnVms to be an explicit nil

### UnsetPowerOnVms
`func (o *VmwareTargetParamsForRecoverVApp) UnsetPowerOnVms()`

UnsetPowerOnVms ensures that no value is present for PowerOnVms, not even an explicit nil
### GetContinueOnError

`func (o *VmwareTargetParamsForRecoverVApp) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *VmwareTargetParamsForRecoverVApp) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *VmwareTargetParamsForRecoverVApp) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *VmwareTargetParamsForRecoverVApp) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *VmwareTargetParamsForRecoverVApp) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *VmwareTargetParamsForRecoverVApp) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetRecoveryProcessType

`func (o *VmwareTargetParamsForRecoverVApp) GetRecoveryProcessType() string`

GetRecoveryProcessType returns the RecoveryProcessType field if non-nil, zero value otherwise.

### GetRecoveryProcessTypeOk

`func (o *VmwareTargetParamsForRecoverVApp) GetRecoveryProcessTypeOk() (*string, bool)`

GetRecoveryProcessTypeOk returns a tuple with the RecoveryProcessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryProcessType

`func (o *VmwareTargetParamsForRecoverVApp) SetRecoveryProcessType(v string)`

SetRecoveryProcessType sets RecoveryProcessType field to given value.

### HasRecoveryProcessType

`func (o *VmwareTargetParamsForRecoverVApp) HasRecoveryProcessType() bool`

HasRecoveryProcessType returns a boolean if a field has been set.

### GetAttemptDifferentialRestore

`func (o *VmwareTargetParamsForRecoverVApp) GetAttemptDifferentialRestore() bool`

GetAttemptDifferentialRestore returns the AttemptDifferentialRestore field if non-nil, zero value otherwise.

### GetAttemptDifferentialRestoreOk

`func (o *VmwareTargetParamsForRecoverVApp) GetAttemptDifferentialRestoreOk() (*bool, bool)`

GetAttemptDifferentialRestoreOk returns a tuple with the AttemptDifferentialRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptDifferentialRestore

`func (o *VmwareTargetParamsForRecoverVApp) SetAttemptDifferentialRestore(v bool)`

SetAttemptDifferentialRestore sets AttemptDifferentialRestore field to given value.

### HasAttemptDifferentialRestore

`func (o *VmwareTargetParamsForRecoverVApp) HasAttemptDifferentialRestore() bool`

HasAttemptDifferentialRestore returns a boolean if a field has been set.

### SetAttemptDifferentialRestoreNil

`func (o *VmwareTargetParamsForRecoverVApp) SetAttemptDifferentialRestoreNil(b bool)`

 SetAttemptDifferentialRestoreNil sets the value for AttemptDifferentialRestore to be an explicit nil

### UnsetAttemptDifferentialRestore
`func (o *VmwareTargetParamsForRecoverVApp) UnsetAttemptDifferentialRestore()`

UnsetAttemptDifferentialRestore ensures that no value is present for AttemptDifferentialRestore, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


