# HyperVTargetParamsForRecoverVm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RenameRecoveredVmsParams** | Pointer to [**NullableRecoveredOrClonedVmsRenameConfig**](RecoveredOrClonedVmsRenameConfig.md) | Specifies params to rename the VMs that are recovered. If not specified, the original names of the VMs are preserved. | [optional] 
**RecoveryTargetConfig** | Pointer to [**NullableHyperVVmRecoveryTargetConfig**](HyperVVmRecoveryTargetConfig.md) | Specifies the recovery target configuration if recovery has to be done to a different location which is different from original source or to original Source with different configuration. If not specified, then the recovery of the vms will be performed to original location with all configuration parameters retained. | [optional] 
**PowerOnVms** | Pointer to **NullableBool** | Specifies whether to power on vms after recovery. If not specified, or false, recovered vms will be in powered off state. | [optional] 
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering other vms if one of vms failed to recover. Default value is false. | [optional] 
**InstantRecovery** | Pointer to **NullableBool** | Specifies whether to perform an instant recovery. By instant recovery, the recovered VM is available before files are completely copied to the recovered VM. Default is true. | [optional] 
**UseSmbService** | Pointer to **NullableBool** | Specifies if the HyperV recovery is using the SMB service to perform the restore. On-prem, this is the case by default. However, as of today, DMaaS does not support SMB, and HyperV VM VM restores will employ an alternative restore method in this case. | [optional] 
**PreserveUuids** | Pointer to **NullableBool** | Specifies whether to preserve uuids of recovered VMs. Default is false. | [optional] 
**VlanConfig** | Pointer to [**NullableRecoveryVlanConfig**](RecoveryVlanConfig.md) | Specifies VLAN Params associated with the recovered. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client&#39;s (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery. | [optional] 

## Methods

### NewHyperVTargetParamsForRecoverVm

`func NewHyperVTargetParamsForRecoverVm() *HyperVTargetParamsForRecoverVm`

NewHyperVTargetParamsForRecoverVm instantiates a new HyperVTargetParamsForRecoverVm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperVTargetParamsForRecoverVmWithDefaults

`func NewHyperVTargetParamsForRecoverVmWithDefaults() *HyperVTargetParamsForRecoverVm`

NewHyperVTargetParamsForRecoverVmWithDefaults instantiates a new HyperVTargetParamsForRecoverVm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRenameRecoveredVmsParams

`func (o *HyperVTargetParamsForRecoverVm) GetRenameRecoveredVmsParams() RecoveredOrClonedVmsRenameConfig`

GetRenameRecoveredVmsParams returns the RenameRecoveredVmsParams field if non-nil, zero value otherwise.

### GetRenameRecoveredVmsParamsOk

`func (o *HyperVTargetParamsForRecoverVm) GetRenameRecoveredVmsParamsOk() (*RecoveredOrClonedVmsRenameConfig, bool)`

GetRenameRecoveredVmsParamsOk returns a tuple with the RenameRecoveredVmsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameRecoveredVmsParams

`func (o *HyperVTargetParamsForRecoverVm) SetRenameRecoveredVmsParams(v RecoveredOrClonedVmsRenameConfig)`

SetRenameRecoveredVmsParams sets RenameRecoveredVmsParams field to given value.

### HasRenameRecoveredVmsParams

`func (o *HyperVTargetParamsForRecoverVm) HasRenameRecoveredVmsParams() bool`

HasRenameRecoveredVmsParams returns a boolean if a field has been set.

### SetRenameRecoveredVmsParamsNil

`func (o *HyperVTargetParamsForRecoverVm) SetRenameRecoveredVmsParamsNil(b bool)`

 SetRenameRecoveredVmsParamsNil sets the value for RenameRecoveredVmsParams to be an explicit nil

### UnsetRenameRecoveredVmsParams
`func (o *HyperVTargetParamsForRecoverVm) UnsetRenameRecoveredVmsParams()`

UnsetRenameRecoveredVmsParams ensures that no value is present for RenameRecoveredVmsParams, not even an explicit nil
### GetRecoveryTargetConfig

`func (o *HyperVTargetParamsForRecoverVm) GetRecoveryTargetConfig() HyperVVmRecoveryTargetConfig`

GetRecoveryTargetConfig returns the RecoveryTargetConfig field if non-nil, zero value otherwise.

### GetRecoveryTargetConfigOk

`func (o *HyperVTargetParamsForRecoverVm) GetRecoveryTargetConfigOk() (*HyperVVmRecoveryTargetConfig, bool)`

GetRecoveryTargetConfigOk returns a tuple with the RecoveryTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTargetConfig

`func (o *HyperVTargetParamsForRecoverVm) SetRecoveryTargetConfig(v HyperVVmRecoveryTargetConfig)`

SetRecoveryTargetConfig sets RecoveryTargetConfig field to given value.

### HasRecoveryTargetConfig

`func (o *HyperVTargetParamsForRecoverVm) HasRecoveryTargetConfig() bool`

HasRecoveryTargetConfig returns a boolean if a field has been set.

### SetRecoveryTargetConfigNil

`func (o *HyperVTargetParamsForRecoverVm) SetRecoveryTargetConfigNil(b bool)`

 SetRecoveryTargetConfigNil sets the value for RecoveryTargetConfig to be an explicit nil

### UnsetRecoveryTargetConfig
`func (o *HyperVTargetParamsForRecoverVm) UnsetRecoveryTargetConfig()`

UnsetRecoveryTargetConfig ensures that no value is present for RecoveryTargetConfig, not even an explicit nil
### GetPowerOnVms

`func (o *HyperVTargetParamsForRecoverVm) GetPowerOnVms() bool`

GetPowerOnVms returns the PowerOnVms field if non-nil, zero value otherwise.

### GetPowerOnVmsOk

`func (o *HyperVTargetParamsForRecoverVm) GetPowerOnVmsOk() (*bool, bool)`

GetPowerOnVmsOk returns a tuple with the PowerOnVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnVms

`func (o *HyperVTargetParamsForRecoverVm) SetPowerOnVms(v bool)`

SetPowerOnVms sets PowerOnVms field to given value.

### HasPowerOnVms

`func (o *HyperVTargetParamsForRecoverVm) HasPowerOnVms() bool`

HasPowerOnVms returns a boolean if a field has been set.

### SetPowerOnVmsNil

`func (o *HyperVTargetParamsForRecoverVm) SetPowerOnVmsNil(b bool)`

 SetPowerOnVmsNil sets the value for PowerOnVms to be an explicit nil

### UnsetPowerOnVms
`func (o *HyperVTargetParamsForRecoverVm) UnsetPowerOnVms()`

UnsetPowerOnVms ensures that no value is present for PowerOnVms, not even an explicit nil
### GetContinueOnError

`func (o *HyperVTargetParamsForRecoverVm) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *HyperVTargetParamsForRecoverVm) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *HyperVTargetParamsForRecoverVm) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *HyperVTargetParamsForRecoverVm) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *HyperVTargetParamsForRecoverVm) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *HyperVTargetParamsForRecoverVm) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetInstantRecovery

`func (o *HyperVTargetParamsForRecoverVm) GetInstantRecovery() bool`

GetInstantRecovery returns the InstantRecovery field if non-nil, zero value otherwise.

### GetInstantRecoveryOk

`func (o *HyperVTargetParamsForRecoverVm) GetInstantRecoveryOk() (*bool, bool)`

GetInstantRecoveryOk returns a tuple with the InstantRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantRecovery

`func (o *HyperVTargetParamsForRecoverVm) SetInstantRecovery(v bool)`

SetInstantRecovery sets InstantRecovery field to given value.

### HasInstantRecovery

`func (o *HyperVTargetParamsForRecoverVm) HasInstantRecovery() bool`

HasInstantRecovery returns a boolean if a field has been set.

### SetInstantRecoveryNil

`func (o *HyperVTargetParamsForRecoverVm) SetInstantRecoveryNil(b bool)`

 SetInstantRecoveryNil sets the value for InstantRecovery to be an explicit nil

### UnsetInstantRecovery
`func (o *HyperVTargetParamsForRecoverVm) UnsetInstantRecovery()`

UnsetInstantRecovery ensures that no value is present for InstantRecovery, not even an explicit nil
### GetUseSmbService

`func (o *HyperVTargetParamsForRecoverVm) GetUseSmbService() bool`

GetUseSmbService returns the UseSmbService field if non-nil, zero value otherwise.

### GetUseSmbServiceOk

`func (o *HyperVTargetParamsForRecoverVm) GetUseSmbServiceOk() (*bool, bool)`

GetUseSmbServiceOk returns a tuple with the UseSmbService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSmbService

`func (o *HyperVTargetParamsForRecoverVm) SetUseSmbService(v bool)`

SetUseSmbService sets UseSmbService field to given value.

### HasUseSmbService

`func (o *HyperVTargetParamsForRecoverVm) HasUseSmbService() bool`

HasUseSmbService returns a boolean if a field has been set.

### SetUseSmbServiceNil

`func (o *HyperVTargetParamsForRecoverVm) SetUseSmbServiceNil(b bool)`

 SetUseSmbServiceNil sets the value for UseSmbService to be an explicit nil

### UnsetUseSmbService
`func (o *HyperVTargetParamsForRecoverVm) UnsetUseSmbService()`

UnsetUseSmbService ensures that no value is present for UseSmbService, not even an explicit nil
### GetPreserveUuids

`func (o *HyperVTargetParamsForRecoverVm) GetPreserveUuids() bool`

GetPreserveUuids returns the PreserveUuids field if non-nil, zero value otherwise.

### GetPreserveUuidsOk

`func (o *HyperVTargetParamsForRecoverVm) GetPreserveUuidsOk() (*bool, bool)`

GetPreserveUuidsOk returns a tuple with the PreserveUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveUuids

`func (o *HyperVTargetParamsForRecoverVm) SetPreserveUuids(v bool)`

SetPreserveUuids sets PreserveUuids field to given value.

### HasPreserveUuids

`func (o *HyperVTargetParamsForRecoverVm) HasPreserveUuids() bool`

HasPreserveUuids returns a boolean if a field has been set.

### SetPreserveUuidsNil

`func (o *HyperVTargetParamsForRecoverVm) SetPreserveUuidsNil(b bool)`

 SetPreserveUuidsNil sets the value for PreserveUuids to be an explicit nil

### UnsetPreserveUuids
`func (o *HyperVTargetParamsForRecoverVm) UnsetPreserveUuids()`

UnsetPreserveUuids ensures that no value is present for PreserveUuids, not even an explicit nil
### GetVlanConfig

`func (o *HyperVTargetParamsForRecoverVm) GetVlanConfig() RecoveryVlanConfig`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *HyperVTargetParamsForRecoverVm) GetVlanConfigOk() (*RecoveryVlanConfig, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *HyperVTargetParamsForRecoverVm) SetVlanConfig(v RecoveryVlanConfig)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *HyperVTargetParamsForRecoverVm) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.

### SetVlanConfigNil

`func (o *HyperVTargetParamsForRecoverVm) SetVlanConfigNil(b bool)`

 SetVlanConfigNil sets the value for VlanConfig to be an explicit nil

### UnsetVlanConfig
`func (o *HyperVTargetParamsForRecoverVm) UnsetVlanConfig()`

UnsetVlanConfig ensures that no value is present for VlanConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


