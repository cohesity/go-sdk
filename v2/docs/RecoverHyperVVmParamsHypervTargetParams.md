# RecoverHyperVVmParamsHypervTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering other vms if one of vms failed to recover. Default value is false. | [optional] 
**InstantRecovery** | Pointer to **NullableBool** | Specifies whether to perform an instant recovery. By instant recovery, the recovered VM is available before files are completely copied to the recovered VM. Default is true. | [optional] 
**OverwriteExistingVms** | Pointer to **NullableBool** | Specifies whether to overwrite existing VMs while performing recovery of a VM. Default value is false. | [optional] 
**PowerOnVms** | Pointer to **NullableBool** | Specifies whether to power on vms after recovery. If not specified, or false, recovered vms will be in powered off state. | [optional] 
**PreserveUuids** | Pointer to **NullableBool** | Specifies whether to preserve uuids of recovered VMs. Default is false. | [optional] 
**RecoverExcludedDisk** | Pointer to **NullableBool** | Specifies whether to recover excluded disk while performing recovery of a VM by creating empty disks for them. Default value is false. | [optional] 
**RecoveryTargetConfig** | Pointer to [**NullableHyperVTargetParamsForRecoverVmRecoveryTargetConfig**](HyperVTargetParamsForRecoverVmRecoveryTargetConfig.md) |  | [optional] 
**RenameRecoveredVmsParams** | Pointer to [**NullableAcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams**](AcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams.md) |  | [optional] 
**UseSmbService** | Pointer to **NullableBool** | Specifies if the HyperV recovery is using the SMB service to perform the restore. On-prem, this is the case by default. However, as of today, DMaaS does not support SMB, and HyperV VM VM restores will employ an alternative restore method in this case. | [optional] 
**VlanConfig** | Pointer to [**NullableHyperVTargetParamsForRecoverVmVlanConfig**](HyperVTargetParamsForRecoverVmVlanConfig.md) |  | [optional] 

## Methods

### NewRecoverHyperVVmParamsHypervTargetParams

`func NewRecoverHyperVVmParamsHypervTargetParams() *RecoverHyperVVmParamsHypervTargetParams`

NewRecoverHyperVVmParamsHypervTargetParams instantiates a new RecoverHyperVVmParamsHypervTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverHyperVVmParamsHypervTargetParamsWithDefaults

`func NewRecoverHyperVVmParamsHypervTargetParamsWithDefaults() *RecoverHyperVVmParamsHypervTargetParams`

NewRecoverHyperVVmParamsHypervTargetParamsWithDefaults instantiates a new RecoverHyperVVmParamsHypervTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *RecoverHyperVVmParamsHypervTargetParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *RecoverHyperVVmParamsHypervTargetParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *RecoverHyperVVmParamsHypervTargetParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *RecoverHyperVVmParamsHypervTargetParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *RecoverHyperVVmParamsHypervTargetParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *RecoverHyperVVmParamsHypervTargetParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetInstantRecovery

`func (o *RecoverHyperVVmParamsHypervTargetParams) GetInstantRecovery() bool`

GetInstantRecovery returns the InstantRecovery field if non-nil, zero value otherwise.

### GetInstantRecoveryOk

`func (o *RecoverHyperVVmParamsHypervTargetParams) GetInstantRecoveryOk() (*bool, bool)`

GetInstantRecoveryOk returns a tuple with the InstantRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantRecovery

`func (o *RecoverHyperVVmParamsHypervTargetParams) SetInstantRecovery(v bool)`

SetInstantRecovery sets InstantRecovery field to given value.

### HasInstantRecovery

`func (o *RecoverHyperVVmParamsHypervTargetParams) HasInstantRecovery() bool`

HasInstantRecovery returns a boolean if a field has been set.

### SetInstantRecoveryNil

`func (o *RecoverHyperVVmParamsHypervTargetParams) SetInstantRecoveryNil(b bool)`

 SetInstantRecoveryNil sets the value for InstantRecovery to be an explicit nil

### UnsetInstantRecovery
`func (o *RecoverHyperVVmParamsHypervTargetParams) UnsetInstantRecovery()`

UnsetInstantRecovery ensures that no value is present for InstantRecovery, not even an explicit nil
### GetOverwriteExistingVms

`func (o *RecoverHyperVVmParamsHypervTargetParams) GetOverwriteExistingVms() bool`

GetOverwriteExistingVms returns the OverwriteExistingVms field if non-nil, zero value otherwise.

### GetOverwriteExistingVmsOk

`func (o *RecoverHyperVVmParamsHypervTargetParams) GetOverwriteExistingVmsOk() (*bool, bool)`

GetOverwriteExistingVmsOk returns a tuple with the OverwriteExistingVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteExistingVms

`func (o *RecoverHyperVVmParamsHypervTargetParams) SetOverwriteExistingVms(v bool)`

SetOverwriteExistingVms sets OverwriteExistingVms field to given value.

### HasOverwriteExistingVms

`func (o *RecoverHyperVVmParamsHypervTargetParams) HasOverwriteExistingVms() bool`

HasOverwriteExistingVms returns a boolean if a field has been set.

### SetOverwriteExistingVmsNil

`func (o *RecoverHyperVVmParamsHypervTargetParams) SetOverwriteExistingVmsNil(b bool)`

 SetOverwriteExistingVmsNil sets the value for OverwriteExistingVms to be an explicit nil

### UnsetOverwriteExistingVms
`func (o *RecoverHyperVVmParamsHypervTargetParams) UnsetOverwriteExistingVms()`

UnsetOverwriteExistingVms ensures that no value is present for OverwriteExistingVms, not even an explicit nil
### GetPowerOnVms

`func (o *RecoverHyperVVmParamsHypervTargetParams) GetPowerOnVms() bool`

GetPowerOnVms returns the PowerOnVms field if non-nil, zero value otherwise.

### GetPowerOnVmsOk

`func (o *RecoverHyperVVmParamsHypervTargetParams) GetPowerOnVmsOk() (*bool, bool)`

GetPowerOnVmsOk returns a tuple with the PowerOnVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnVms

`func (o *RecoverHyperVVmParamsHypervTargetParams) SetPowerOnVms(v bool)`

SetPowerOnVms sets PowerOnVms field to given value.

### HasPowerOnVms

`func (o *RecoverHyperVVmParamsHypervTargetParams) HasPowerOnVms() bool`

HasPowerOnVms returns a boolean if a field has been set.

### SetPowerOnVmsNil

`func (o *RecoverHyperVVmParamsHypervTargetParams) SetPowerOnVmsNil(b bool)`

 SetPowerOnVmsNil sets the value for PowerOnVms to be an explicit nil

### UnsetPowerOnVms
`func (o *RecoverHyperVVmParamsHypervTargetParams) UnsetPowerOnVms()`

UnsetPowerOnVms ensures that no value is present for PowerOnVms, not even an explicit nil
### GetPreserveUuids

`func (o *RecoverHyperVVmParamsHypervTargetParams) GetPreserveUuids() bool`

GetPreserveUuids returns the PreserveUuids field if non-nil, zero value otherwise.

### GetPreserveUuidsOk

`func (o *RecoverHyperVVmParamsHypervTargetParams) GetPreserveUuidsOk() (*bool, bool)`

GetPreserveUuidsOk returns a tuple with the PreserveUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveUuids

`func (o *RecoverHyperVVmParamsHypervTargetParams) SetPreserveUuids(v bool)`

SetPreserveUuids sets PreserveUuids field to given value.

### HasPreserveUuids

`func (o *RecoverHyperVVmParamsHypervTargetParams) HasPreserveUuids() bool`

HasPreserveUuids returns a boolean if a field has been set.

### SetPreserveUuidsNil

`func (o *RecoverHyperVVmParamsHypervTargetParams) SetPreserveUuidsNil(b bool)`

 SetPreserveUuidsNil sets the value for PreserveUuids to be an explicit nil

### UnsetPreserveUuids
`func (o *RecoverHyperVVmParamsHypervTargetParams) UnsetPreserveUuids()`

UnsetPreserveUuids ensures that no value is present for PreserveUuids, not even an explicit nil
### GetRecoverExcludedDisk

`func (o *RecoverHyperVVmParamsHypervTargetParams) GetRecoverExcludedDisk() bool`

GetRecoverExcludedDisk returns the RecoverExcludedDisk field if non-nil, zero value otherwise.

### GetRecoverExcludedDiskOk

`func (o *RecoverHyperVVmParamsHypervTargetParams) GetRecoverExcludedDiskOk() (*bool, bool)`

GetRecoverExcludedDiskOk returns a tuple with the RecoverExcludedDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverExcludedDisk

`func (o *RecoverHyperVVmParamsHypervTargetParams) SetRecoverExcludedDisk(v bool)`

SetRecoverExcludedDisk sets RecoverExcludedDisk field to given value.

### HasRecoverExcludedDisk

`func (o *RecoverHyperVVmParamsHypervTargetParams) HasRecoverExcludedDisk() bool`

HasRecoverExcludedDisk returns a boolean if a field has been set.

### SetRecoverExcludedDiskNil

`func (o *RecoverHyperVVmParamsHypervTargetParams) SetRecoverExcludedDiskNil(b bool)`

 SetRecoverExcludedDiskNil sets the value for RecoverExcludedDisk to be an explicit nil

### UnsetRecoverExcludedDisk
`func (o *RecoverHyperVVmParamsHypervTargetParams) UnsetRecoverExcludedDisk()`

UnsetRecoverExcludedDisk ensures that no value is present for RecoverExcludedDisk, not even an explicit nil
### GetRecoveryTargetConfig

`func (o *RecoverHyperVVmParamsHypervTargetParams) GetRecoveryTargetConfig() HyperVTargetParamsForRecoverVmRecoveryTargetConfig`

GetRecoveryTargetConfig returns the RecoveryTargetConfig field if non-nil, zero value otherwise.

### GetRecoveryTargetConfigOk

`func (o *RecoverHyperVVmParamsHypervTargetParams) GetRecoveryTargetConfigOk() (*HyperVTargetParamsForRecoverVmRecoveryTargetConfig, bool)`

GetRecoveryTargetConfigOk returns a tuple with the RecoveryTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTargetConfig

`func (o *RecoverHyperVVmParamsHypervTargetParams) SetRecoveryTargetConfig(v HyperVTargetParamsForRecoverVmRecoveryTargetConfig)`

SetRecoveryTargetConfig sets RecoveryTargetConfig field to given value.

### HasRecoveryTargetConfig

`func (o *RecoverHyperVVmParamsHypervTargetParams) HasRecoveryTargetConfig() bool`

HasRecoveryTargetConfig returns a boolean if a field has been set.

### SetRecoveryTargetConfigNil

`func (o *RecoverHyperVVmParamsHypervTargetParams) SetRecoveryTargetConfigNil(b bool)`

 SetRecoveryTargetConfigNil sets the value for RecoveryTargetConfig to be an explicit nil

### UnsetRecoveryTargetConfig
`func (o *RecoverHyperVVmParamsHypervTargetParams) UnsetRecoveryTargetConfig()`

UnsetRecoveryTargetConfig ensures that no value is present for RecoveryTargetConfig, not even an explicit nil
### GetRenameRecoveredVmsParams

`func (o *RecoverHyperVVmParamsHypervTargetParams) GetRenameRecoveredVmsParams() AcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams`

GetRenameRecoveredVmsParams returns the RenameRecoveredVmsParams field if non-nil, zero value otherwise.

### GetRenameRecoveredVmsParamsOk

`func (o *RecoverHyperVVmParamsHypervTargetParams) GetRenameRecoveredVmsParamsOk() (*AcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams, bool)`

GetRenameRecoveredVmsParamsOk returns a tuple with the RenameRecoveredVmsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameRecoveredVmsParams

`func (o *RecoverHyperVVmParamsHypervTargetParams) SetRenameRecoveredVmsParams(v AcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams)`

SetRenameRecoveredVmsParams sets RenameRecoveredVmsParams field to given value.

### HasRenameRecoveredVmsParams

`func (o *RecoverHyperVVmParamsHypervTargetParams) HasRenameRecoveredVmsParams() bool`

HasRenameRecoveredVmsParams returns a boolean if a field has been set.

### SetRenameRecoveredVmsParamsNil

`func (o *RecoverHyperVVmParamsHypervTargetParams) SetRenameRecoveredVmsParamsNil(b bool)`

 SetRenameRecoveredVmsParamsNil sets the value for RenameRecoveredVmsParams to be an explicit nil

### UnsetRenameRecoveredVmsParams
`func (o *RecoverHyperVVmParamsHypervTargetParams) UnsetRenameRecoveredVmsParams()`

UnsetRenameRecoveredVmsParams ensures that no value is present for RenameRecoveredVmsParams, not even an explicit nil
### GetUseSmbService

`func (o *RecoverHyperVVmParamsHypervTargetParams) GetUseSmbService() bool`

GetUseSmbService returns the UseSmbService field if non-nil, zero value otherwise.

### GetUseSmbServiceOk

`func (o *RecoverHyperVVmParamsHypervTargetParams) GetUseSmbServiceOk() (*bool, bool)`

GetUseSmbServiceOk returns a tuple with the UseSmbService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSmbService

`func (o *RecoverHyperVVmParamsHypervTargetParams) SetUseSmbService(v bool)`

SetUseSmbService sets UseSmbService field to given value.

### HasUseSmbService

`func (o *RecoverHyperVVmParamsHypervTargetParams) HasUseSmbService() bool`

HasUseSmbService returns a boolean if a field has been set.

### SetUseSmbServiceNil

`func (o *RecoverHyperVVmParamsHypervTargetParams) SetUseSmbServiceNil(b bool)`

 SetUseSmbServiceNil sets the value for UseSmbService to be an explicit nil

### UnsetUseSmbService
`func (o *RecoverHyperVVmParamsHypervTargetParams) UnsetUseSmbService()`

UnsetUseSmbService ensures that no value is present for UseSmbService, not even an explicit nil
### GetVlanConfig

`func (o *RecoverHyperVVmParamsHypervTargetParams) GetVlanConfig() HyperVTargetParamsForRecoverVmVlanConfig`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *RecoverHyperVVmParamsHypervTargetParams) GetVlanConfigOk() (*HyperVTargetParamsForRecoverVmVlanConfig, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *RecoverHyperVVmParamsHypervTargetParams) SetVlanConfig(v HyperVTargetParamsForRecoverVmVlanConfig)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *RecoverHyperVVmParamsHypervTargetParams) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.

### SetVlanConfigNil

`func (o *RecoverHyperVVmParamsHypervTargetParams) SetVlanConfigNil(b bool)`

 SetVlanConfigNil sets the value for VlanConfig to be an explicit nil

### UnsetVlanConfig
`func (o *RecoverHyperVVmParamsHypervTargetParams) UnsetVlanConfig()`

UnsetVlanConfig ensures that no value is present for VlanConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


