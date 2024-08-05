# RecoverVmwareDiskParamsVmwareTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether or not to continue performing the recovery in the event that an error is encountered. | [optional] 
**OriginalSourceConfig** | Pointer to [**VmwareRecoverDisksOriginalSourceConfig**](VmwareRecoverDisksOriginalSourceConfig.md) |  | [optional] 
**PowerOffVms** | Pointer to **NullableBool** | Specifies whether or not to power off VMs before performing the recovery. | [optional] 
**PowerOnVms** | Pointer to **NullableBool** | Specifies whether or not to power on VMs after performing the recovery. | [optional] 
**TargetSourceConfig** | Pointer to [**VmwareRecoverDisksTargetSourceConfig**](VmwareRecoverDisksTargetSourceConfig.md) |  | [optional] 
**VlanConfig** | Pointer to [**NullableRecoverKubernetesNamespaceParamsVlanConfig**](RecoverKubernetesNamespaceParamsVlanConfig.md) |  | [optional] 

## Methods

### NewRecoverVmwareDiskParamsVmwareTargetParams

`func NewRecoverVmwareDiskParamsVmwareTargetParams() *RecoverVmwareDiskParamsVmwareTargetParams`

NewRecoverVmwareDiskParamsVmwareTargetParams instantiates a new RecoverVmwareDiskParamsVmwareTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverVmwareDiskParamsVmwareTargetParamsWithDefaults

`func NewRecoverVmwareDiskParamsVmwareTargetParamsWithDefaults() *RecoverVmwareDiskParamsVmwareTargetParams`

NewRecoverVmwareDiskParamsVmwareTargetParamsWithDefaults instantiates a new RecoverVmwareDiskParamsVmwareTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *RecoverVmwareDiskParamsVmwareTargetParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *RecoverVmwareDiskParamsVmwareTargetParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *RecoverVmwareDiskParamsVmwareTargetParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *RecoverVmwareDiskParamsVmwareTargetParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *RecoverVmwareDiskParamsVmwareTargetParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *RecoverVmwareDiskParamsVmwareTargetParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetOriginalSourceConfig

`func (o *RecoverVmwareDiskParamsVmwareTargetParams) GetOriginalSourceConfig() VmwareRecoverDisksOriginalSourceConfig`

GetOriginalSourceConfig returns the OriginalSourceConfig field if non-nil, zero value otherwise.

### GetOriginalSourceConfigOk

`func (o *RecoverVmwareDiskParamsVmwareTargetParams) GetOriginalSourceConfigOk() (*VmwareRecoverDisksOriginalSourceConfig, bool)`

GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSourceConfig

`func (o *RecoverVmwareDiskParamsVmwareTargetParams) SetOriginalSourceConfig(v VmwareRecoverDisksOriginalSourceConfig)`

SetOriginalSourceConfig sets OriginalSourceConfig field to given value.

### HasOriginalSourceConfig

`func (o *RecoverVmwareDiskParamsVmwareTargetParams) HasOriginalSourceConfig() bool`

HasOriginalSourceConfig returns a boolean if a field has been set.

### GetPowerOffVms

`func (o *RecoverVmwareDiskParamsVmwareTargetParams) GetPowerOffVms() bool`

GetPowerOffVms returns the PowerOffVms field if non-nil, zero value otherwise.

### GetPowerOffVmsOk

`func (o *RecoverVmwareDiskParamsVmwareTargetParams) GetPowerOffVmsOk() (*bool, bool)`

GetPowerOffVmsOk returns a tuple with the PowerOffVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOffVms

`func (o *RecoverVmwareDiskParamsVmwareTargetParams) SetPowerOffVms(v bool)`

SetPowerOffVms sets PowerOffVms field to given value.

### HasPowerOffVms

`func (o *RecoverVmwareDiskParamsVmwareTargetParams) HasPowerOffVms() bool`

HasPowerOffVms returns a boolean if a field has been set.

### SetPowerOffVmsNil

`func (o *RecoverVmwareDiskParamsVmwareTargetParams) SetPowerOffVmsNil(b bool)`

 SetPowerOffVmsNil sets the value for PowerOffVms to be an explicit nil

### UnsetPowerOffVms
`func (o *RecoverVmwareDiskParamsVmwareTargetParams) UnsetPowerOffVms()`

UnsetPowerOffVms ensures that no value is present for PowerOffVms, not even an explicit nil
### GetPowerOnVms

`func (o *RecoverVmwareDiskParamsVmwareTargetParams) GetPowerOnVms() bool`

GetPowerOnVms returns the PowerOnVms field if non-nil, zero value otherwise.

### GetPowerOnVmsOk

`func (o *RecoverVmwareDiskParamsVmwareTargetParams) GetPowerOnVmsOk() (*bool, bool)`

GetPowerOnVmsOk returns a tuple with the PowerOnVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnVms

`func (o *RecoverVmwareDiskParamsVmwareTargetParams) SetPowerOnVms(v bool)`

SetPowerOnVms sets PowerOnVms field to given value.

### HasPowerOnVms

`func (o *RecoverVmwareDiskParamsVmwareTargetParams) HasPowerOnVms() bool`

HasPowerOnVms returns a boolean if a field has been set.

### SetPowerOnVmsNil

`func (o *RecoverVmwareDiskParamsVmwareTargetParams) SetPowerOnVmsNil(b bool)`

 SetPowerOnVmsNil sets the value for PowerOnVms to be an explicit nil

### UnsetPowerOnVms
`func (o *RecoverVmwareDiskParamsVmwareTargetParams) UnsetPowerOnVms()`

UnsetPowerOnVms ensures that no value is present for PowerOnVms, not even an explicit nil
### GetTargetSourceConfig

`func (o *RecoverVmwareDiskParamsVmwareTargetParams) GetTargetSourceConfig() VmwareRecoverDisksTargetSourceConfig`

GetTargetSourceConfig returns the TargetSourceConfig field if non-nil, zero value otherwise.

### GetTargetSourceConfigOk

`func (o *RecoverVmwareDiskParamsVmwareTargetParams) GetTargetSourceConfigOk() (*VmwareRecoverDisksTargetSourceConfig, bool)`

GetTargetSourceConfigOk returns a tuple with the TargetSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSourceConfig

`func (o *RecoverVmwareDiskParamsVmwareTargetParams) SetTargetSourceConfig(v VmwareRecoverDisksTargetSourceConfig)`

SetTargetSourceConfig sets TargetSourceConfig field to given value.

### HasTargetSourceConfig

`func (o *RecoverVmwareDiskParamsVmwareTargetParams) HasTargetSourceConfig() bool`

HasTargetSourceConfig returns a boolean if a field has been set.

### GetVlanConfig

`func (o *RecoverVmwareDiskParamsVmwareTargetParams) GetVlanConfig() RecoverKubernetesNamespaceParamsVlanConfig`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *RecoverVmwareDiskParamsVmwareTargetParams) GetVlanConfigOk() (*RecoverKubernetesNamespaceParamsVlanConfig, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *RecoverVmwareDiskParamsVmwareTargetParams) SetVlanConfig(v RecoverKubernetesNamespaceParamsVlanConfig)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *RecoverVmwareDiskParamsVmwareTargetParams) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.

### SetVlanConfigNil

`func (o *RecoverVmwareDiskParamsVmwareTargetParams) SetVlanConfigNil(b bool)`

 SetVlanConfigNil sets the value for VlanConfig to be an explicit nil

### UnsetVlanConfig
`func (o *RecoverVmwareDiskParamsVmwareTargetParams) UnsetVlanConfig()`

UnsetVlanConfig ensures that no value is present for VlanConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


