# KvmTargetParamsForRecoverVm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering other vms if one of vms failed to recover. Default value is false. | [optional] 
**PowerOnVms** | Pointer to **NullableBool** | Specifies whether to power on vms after recovery. If not specified, or false, recovered vms will be in powered off state. | [optional] 
**RecoveryTargetConfig** | Pointer to [**KvmTargetParamsForRecoverVmRecoveryTargetConfig**](KvmTargetParamsForRecoverVmRecoveryTargetConfig.md) |  | [optional] 
**RenameRecoveredVmsParams** | Pointer to [**KvmTargetParamsForRecoverVmRenameRecoveredVmsParams**](KvmTargetParamsForRecoverVmRenameRecoveredVmsParams.md) |  | [optional] 
**VlanConfig** | Pointer to [**KvmTargetParamsForRecoverVmVlanConfig**](KvmTargetParamsForRecoverVmVlanConfig.md) |  | [optional] 

## Methods

### NewKvmTargetParamsForRecoverVm

`func NewKvmTargetParamsForRecoverVm() *KvmTargetParamsForRecoverVm`

NewKvmTargetParamsForRecoverVm instantiates a new KvmTargetParamsForRecoverVm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvmTargetParamsForRecoverVmWithDefaults

`func NewKvmTargetParamsForRecoverVmWithDefaults() *KvmTargetParamsForRecoverVm`

NewKvmTargetParamsForRecoverVmWithDefaults instantiates a new KvmTargetParamsForRecoverVm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *KvmTargetParamsForRecoverVm) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *KvmTargetParamsForRecoverVm) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *KvmTargetParamsForRecoverVm) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *KvmTargetParamsForRecoverVm) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *KvmTargetParamsForRecoverVm) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *KvmTargetParamsForRecoverVm) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetPowerOnVms

`func (o *KvmTargetParamsForRecoverVm) GetPowerOnVms() bool`

GetPowerOnVms returns the PowerOnVms field if non-nil, zero value otherwise.

### GetPowerOnVmsOk

`func (o *KvmTargetParamsForRecoverVm) GetPowerOnVmsOk() (*bool, bool)`

GetPowerOnVmsOk returns a tuple with the PowerOnVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnVms

`func (o *KvmTargetParamsForRecoverVm) SetPowerOnVms(v bool)`

SetPowerOnVms sets PowerOnVms field to given value.

### HasPowerOnVms

`func (o *KvmTargetParamsForRecoverVm) HasPowerOnVms() bool`

HasPowerOnVms returns a boolean if a field has been set.

### SetPowerOnVmsNil

`func (o *KvmTargetParamsForRecoverVm) SetPowerOnVmsNil(b bool)`

 SetPowerOnVmsNil sets the value for PowerOnVms to be an explicit nil

### UnsetPowerOnVms
`func (o *KvmTargetParamsForRecoverVm) UnsetPowerOnVms()`

UnsetPowerOnVms ensures that no value is present for PowerOnVms, not even an explicit nil
### GetRecoveryTargetConfig

`func (o *KvmTargetParamsForRecoverVm) GetRecoveryTargetConfig() KvmTargetParamsForRecoverVmRecoveryTargetConfig`

GetRecoveryTargetConfig returns the RecoveryTargetConfig field if non-nil, zero value otherwise.

### GetRecoveryTargetConfigOk

`func (o *KvmTargetParamsForRecoverVm) GetRecoveryTargetConfigOk() (*KvmTargetParamsForRecoverVmRecoveryTargetConfig, bool)`

GetRecoveryTargetConfigOk returns a tuple with the RecoveryTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTargetConfig

`func (o *KvmTargetParamsForRecoverVm) SetRecoveryTargetConfig(v KvmTargetParamsForRecoverVmRecoveryTargetConfig)`

SetRecoveryTargetConfig sets RecoveryTargetConfig field to given value.

### HasRecoveryTargetConfig

`func (o *KvmTargetParamsForRecoverVm) HasRecoveryTargetConfig() bool`

HasRecoveryTargetConfig returns a boolean if a field has been set.

### GetRenameRecoveredVmsParams

`func (o *KvmTargetParamsForRecoverVm) GetRenameRecoveredVmsParams() KvmTargetParamsForRecoverVmRenameRecoveredVmsParams`

GetRenameRecoveredVmsParams returns the RenameRecoveredVmsParams field if non-nil, zero value otherwise.

### GetRenameRecoveredVmsParamsOk

`func (o *KvmTargetParamsForRecoverVm) GetRenameRecoveredVmsParamsOk() (*KvmTargetParamsForRecoverVmRenameRecoveredVmsParams, bool)`

GetRenameRecoveredVmsParamsOk returns a tuple with the RenameRecoveredVmsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameRecoveredVmsParams

`func (o *KvmTargetParamsForRecoverVm) SetRenameRecoveredVmsParams(v KvmTargetParamsForRecoverVmRenameRecoveredVmsParams)`

SetRenameRecoveredVmsParams sets RenameRecoveredVmsParams field to given value.

### HasRenameRecoveredVmsParams

`func (o *KvmTargetParamsForRecoverVm) HasRenameRecoveredVmsParams() bool`

HasRenameRecoveredVmsParams returns a boolean if a field has been set.

### GetVlanConfig

`func (o *KvmTargetParamsForRecoverVm) GetVlanConfig() KvmTargetParamsForRecoverVmVlanConfig`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *KvmTargetParamsForRecoverVm) GetVlanConfigOk() (*KvmTargetParamsForRecoverVmVlanConfig, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *KvmTargetParamsForRecoverVm) SetVlanConfig(v KvmTargetParamsForRecoverVmVlanConfig)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *KvmTargetParamsForRecoverVm) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


