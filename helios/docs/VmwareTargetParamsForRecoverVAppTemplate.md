# VmwareTargetParamsForRecoverVAppTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RenameRecoveredVmsParams** | Pointer to [**NullableRecoveredOrClonedVmsRenameConfig**](RecoveredOrClonedVmsRenameConfig.md) | Specifies params to rename the VMs that are recovered. If not specified, the original names of the VMs are preserved. | [optional] 
**RenameRecoveredVAppTemplateParams** | Pointer to [**NullableRecoveredOrClonedVmsRenameConfig**](RecoveredOrClonedVmsRenameConfig.md) | Specifies params to rename the vApps templates that are recovered. If not specified, the original names of the vApp templates are preserved. | [optional] 
**RecoveryTargetConfig** | Pointer to [**NullableVmwareVAppTemplateRecoveryTargetConfig**](VmwareVAppTemplateRecoveryTargetConfig.md) | Specifies the recovery target configuration if recovery has to be done to a different location which is different from original source or to original Source with different configuration. If not specified, then the recovery of the vApp templates will be performed to original location with all configuration parameters retained. | [optional] 
**VlanConfig** | Pointer to [**NullableRecoveryVlanConfig**](RecoveryVlanConfig.md) | Specifies VLAN Params associated with the recovered. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client&#39;s (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery. | [optional] 
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering other vms if one of vms failed to recover. Default value is false. | [optional] 
**RecoveryProcessType** | Pointer to **string** | Specifies type of Recovery Process to be used. InstantRecovery/CopyRecovery etc... Default value is InstantRecovery. | [optional] 
**AttemptDifferentialRestore** | Pointer to **NullableBool** | Specifies whether to attempt differential restore. | [optional] 

## Methods

### NewVmwareTargetParamsForRecoverVAppTemplate

`func NewVmwareTargetParamsForRecoverVAppTemplate() *VmwareTargetParamsForRecoverVAppTemplate`

NewVmwareTargetParamsForRecoverVAppTemplate instantiates a new VmwareTargetParamsForRecoverVAppTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareTargetParamsForRecoverVAppTemplateWithDefaults

`func NewVmwareTargetParamsForRecoverVAppTemplateWithDefaults() *VmwareTargetParamsForRecoverVAppTemplate`

NewVmwareTargetParamsForRecoverVAppTemplateWithDefaults instantiates a new VmwareTargetParamsForRecoverVAppTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRenameRecoveredVmsParams

`func (o *VmwareTargetParamsForRecoverVAppTemplate) GetRenameRecoveredVmsParams() RecoveredOrClonedVmsRenameConfig`

GetRenameRecoveredVmsParams returns the RenameRecoveredVmsParams field if non-nil, zero value otherwise.

### GetRenameRecoveredVmsParamsOk

`func (o *VmwareTargetParamsForRecoverVAppTemplate) GetRenameRecoveredVmsParamsOk() (*RecoveredOrClonedVmsRenameConfig, bool)`

GetRenameRecoveredVmsParamsOk returns a tuple with the RenameRecoveredVmsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameRecoveredVmsParams

`func (o *VmwareTargetParamsForRecoverVAppTemplate) SetRenameRecoveredVmsParams(v RecoveredOrClonedVmsRenameConfig)`

SetRenameRecoveredVmsParams sets RenameRecoveredVmsParams field to given value.

### HasRenameRecoveredVmsParams

`func (o *VmwareTargetParamsForRecoverVAppTemplate) HasRenameRecoveredVmsParams() bool`

HasRenameRecoveredVmsParams returns a boolean if a field has been set.

### SetRenameRecoveredVmsParamsNil

`func (o *VmwareTargetParamsForRecoverVAppTemplate) SetRenameRecoveredVmsParamsNil(b bool)`

 SetRenameRecoveredVmsParamsNil sets the value for RenameRecoveredVmsParams to be an explicit nil

### UnsetRenameRecoveredVmsParams
`func (o *VmwareTargetParamsForRecoverVAppTemplate) UnsetRenameRecoveredVmsParams()`

UnsetRenameRecoveredVmsParams ensures that no value is present for RenameRecoveredVmsParams, not even an explicit nil
### GetRenameRecoveredVAppTemplateParams

`func (o *VmwareTargetParamsForRecoverVAppTemplate) GetRenameRecoveredVAppTemplateParams() RecoveredOrClonedVmsRenameConfig`

GetRenameRecoveredVAppTemplateParams returns the RenameRecoveredVAppTemplateParams field if non-nil, zero value otherwise.

### GetRenameRecoveredVAppTemplateParamsOk

`func (o *VmwareTargetParamsForRecoverVAppTemplate) GetRenameRecoveredVAppTemplateParamsOk() (*RecoveredOrClonedVmsRenameConfig, bool)`

GetRenameRecoveredVAppTemplateParamsOk returns a tuple with the RenameRecoveredVAppTemplateParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameRecoveredVAppTemplateParams

`func (o *VmwareTargetParamsForRecoverVAppTemplate) SetRenameRecoveredVAppTemplateParams(v RecoveredOrClonedVmsRenameConfig)`

SetRenameRecoveredVAppTemplateParams sets RenameRecoveredVAppTemplateParams field to given value.

### HasRenameRecoveredVAppTemplateParams

`func (o *VmwareTargetParamsForRecoverVAppTemplate) HasRenameRecoveredVAppTemplateParams() bool`

HasRenameRecoveredVAppTemplateParams returns a boolean if a field has been set.

### SetRenameRecoveredVAppTemplateParamsNil

`func (o *VmwareTargetParamsForRecoverVAppTemplate) SetRenameRecoveredVAppTemplateParamsNil(b bool)`

 SetRenameRecoveredVAppTemplateParamsNil sets the value for RenameRecoveredVAppTemplateParams to be an explicit nil

### UnsetRenameRecoveredVAppTemplateParams
`func (o *VmwareTargetParamsForRecoverVAppTemplate) UnsetRenameRecoveredVAppTemplateParams()`

UnsetRenameRecoveredVAppTemplateParams ensures that no value is present for RenameRecoveredVAppTemplateParams, not even an explicit nil
### GetRecoveryTargetConfig

`func (o *VmwareTargetParamsForRecoverVAppTemplate) GetRecoveryTargetConfig() VmwareVAppTemplateRecoveryTargetConfig`

GetRecoveryTargetConfig returns the RecoveryTargetConfig field if non-nil, zero value otherwise.

### GetRecoveryTargetConfigOk

`func (o *VmwareTargetParamsForRecoverVAppTemplate) GetRecoveryTargetConfigOk() (*VmwareVAppTemplateRecoveryTargetConfig, bool)`

GetRecoveryTargetConfigOk returns a tuple with the RecoveryTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTargetConfig

`func (o *VmwareTargetParamsForRecoverVAppTemplate) SetRecoveryTargetConfig(v VmwareVAppTemplateRecoveryTargetConfig)`

SetRecoveryTargetConfig sets RecoveryTargetConfig field to given value.

### HasRecoveryTargetConfig

`func (o *VmwareTargetParamsForRecoverVAppTemplate) HasRecoveryTargetConfig() bool`

HasRecoveryTargetConfig returns a boolean if a field has been set.

### SetRecoveryTargetConfigNil

`func (o *VmwareTargetParamsForRecoverVAppTemplate) SetRecoveryTargetConfigNil(b bool)`

 SetRecoveryTargetConfigNil sets the value for RecoveryTargetConfig to be an explicit nil

### UnsetRecoveryTargetConfig
`func (o *VmwareTargetParamsForRecoverVAppTemplate) UnsetRecoveryTargetConfig()`

UnsetRecoveryTargetConfig ensures that no value is present for RecoveryTargetConfig, not even an explicit nil
### GetVlanConfig

`func (o *VmwareTargetParamsForRecoverVAppTemplate) GetVlanConfig() RecoveryVlanConfig`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *VmwareTargetParamsForRecoverVAppTemplate) GetVlanConfigOk() (*RecoveryVlanConfig, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *VmwareTargetParamsForRecoverVAppTemplate) SetVlanConfig(v RecoveryVlanConfig)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *VmwareTargetParamsForRecoverVAppTemplate) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.

### SetVlanConfigNil

`func (o *VmwareTargetParamsForRecoverVAppTemplate) SetVlanConfigNil(b bool)`

 SetVlanConfigNil sets the value for VlanConfig to be an explicit nil

### UnsetVlanConfig
`func (o *VmwareTargetParamsForRecoverVAppTemplate) UnsetVlanConfig()`

UnsetVlanConfig ensures that no value is present for VlanConfig, not even an explicit nil
### GetContinueOnError

`func (o *VmwareTargetParamsForRecoverVAppTemplate) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *VmwareTargetParamsForRecoverVAppTemplate) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *VmwareTargetParamsForRecoverVAppTemplate) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *VmwareTargetParamsForRecoverVAppTemplate) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *VmwareTargetParamsForRecoverVAppTemplate) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *VmwareTargetParamsForRecoverVAppTemplate) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetRecoveryProcessType

`func (o *VmwareTargetParamsForRecoverVAppTemplate) GetRecoveryProcessType() string`

GetRecoveryProcessType returns the RecoveryProcessType field if non-nil, zero value otherwise.

### GetRecoveryProcessTypeOk

`func (o *VmwareTargetParamsForRecoverVAppTemplate) GetRecoveryProcessTypeOk() (*string, bool)`

GetRecoveryProcessTypeOk returns a tuple with the RecoveryProcessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryProcessType

`func (o *VmwareTargetParamsForRecoverVAppTemplate) SetRecoveryProcessType(v string)`

SetRecoveryProcessType sets RecoveryProcessType field to given value.

### HasRecoveryProcessType

`func (o *VmwareTargetParamsForRecoverVAppTemplate) HasRecoveryProcessType() bool`

HasRecoveryProcessType returns a boolean if a field has been set.

### GetAttemptDifferentialRestore

`func (o *VmwareTargetParamsForRecoverVAppTemplate) GetAttemptDifferentialRestore() bool`

GetAttemptDifferentialRestore returns the AttemptDifferentialRestore field if non-nil, zero value otherwise.

### GetAttemptDifferentialRestoreOk

`func (o *VmwareTargetParamsForRecoverVAppTemplate) GetAttemptDifferentialRestoreOk() (*bool, bool)`

GetAttemptDifferentialRestoreOk returns a tuple with the AttemptDifferentialRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptDifferentialRestore

`func (o *VmwareTargetParamsForRecoverVAppTemplate) SetAttemptDifferentialRestore(v bool)`

SetAttemptDifferentialRestore sets AttemptDifferentialRestore field to given value.

### HasAttemptDifferentialRestore

`func (o *VmwareTargetParamsForRecoverVAppTemplate) HasAttemptDifferentialRestore() bool`

HasAttemptDifferentialRestore returns a boolean if a field has been set.

### SetAttemptDifferentialRestoreNil

`func (o *VmwareTargetParamsForRecoverVAppTemplate) SetAttemptDifferentialRestoreNil(b bool)`

 SetAttemptDifferentialRestoreNil sets the value for AttemptDifferentialRestore to be an explicit nil

### UnsetAttemptDifferentialRestore
`func (o *VmwareTargetParamsForRecoverVAppTemplate) UnsetAttemptDifferentialRestore()`

UnsetAttemptDifferentialRestore ensures that no value is present for AttemptDifferentialRestore, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


