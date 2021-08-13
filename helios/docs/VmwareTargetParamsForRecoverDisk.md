# VmwareTargetParamsForRecoverDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginalSourceConfig** | Pointer to [**VmwareRecoverDisksOriginalSourceConfig**](VmwareRecoverDisksOriginalSourceConfig.md) |  | [optional] 
**TargetSourceConfig** | Pointer to [**VmwareRecoverDisksTargetSourceConfig**](VmwareRecoverDisksTargetSourceConfig.md) |  | [optional] 
**VlanConfig** | Pointer to [**NullableRecoveryVlanConfig**](RecoveryVlanConfig.md) | Specifies VLAN Params associated with the recovered. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client&#39;s (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery. | [optional] 
**PowerOffVms** | Pointer to **NullableBool** | Specifies whether or not to power off VMs before performing the recovery. | [optional] 
**PowerOnVms** | Pointer to **NullableBool** | Specifies whether or not to power on VMs after performing the recovery. | [optional] 
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether or not to continue performing the recovery in the event that an error is encountered. | [optional] 

## Methods

### NewVmwareTargetParamsForRecoverDisk

`func NewVmwareTargetParamsForRecoverDisk() *VmwareTargetParamsForRecoverDisk`

NewVmwareTargetParamsForRecoverDisk instantiates a new VmwareTargetParamsForRecoverDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareTargetParamsForRecoverDiskWithDefaults

`func NewVmwareTargetParamsForRecoverDiskWithDefaults() *VmwareTargetParamsForRecoverDisk`

NewVmwareTargetParamsForRecoverDiskWithDefaults instantiates a new VmwareTargetParamsForRecoverDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginalSourceConfig

`func (o *VmwareTargetParamsForRecoverDisk) GetOriginalSourceConfig() VmwareRecoverDisksOriginalSourceConfig`

GetOriginalSourceConfig returns the OriginalSourceConfig field if non-nil, zero value otherwise.

### GetOriginalSourceConfigOk

`func (o *VmwareTargetParamsForRecoverDisk) GetOriginalSourceConfigOk() (*VmwareRecoverDisksOriginalSourceConfig, bool)`

GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSourceConfig

`func (o *VmwareTargetParamsForRecoverDisk) SetOriginalSourceConfig(v VmwareRecoverDisksOriginalSourceConfig)`

SetOriginalSourceConfig sets OriginalSourceConfig field to given value.

### HasOriginalSourceConfig

`func (o *VmwareTargetParamsForRecoverDisk) HasOriginalSourceConfig() bool`

HasOriginalSourceConfig returns a boolean if a field has been set.

### GetTargetSourceConfig

`func (o *VmwareTargetParamsForRecoverDisk) GetTargetSourceConfig() VmwareRecoverDisksTargetSourceConfig`

GetTargetSourceConfig returns the TargetSourceConfig field if non-nil, zero value otherwise.

### GetTargetSourceConfigOk

`func (o *VmwareTargetParamsForRecoverDisk) GetTargetSourceConfigOk() (*VmwareRecoverDisksTargetSourceConfig, bool)`

GetTargetSourceConfigOk returns a tuple with the TargetSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSourceConfig

`func (o *VmwareTargetParamsForRecoverDisk) SetTargetSourceConfig(v VmwareRecoverDisksTargetSourceConfig)`

SetTargetSourceConfig sets TargetSourceConfig field to given value.

### HasTargetSourceConfig

`func (o *VmwareTargetParamsForRecoverDisk) HasTargetSourceConfig() bool`

HasTargetSourceConfig returns a boolean if a field has been set.

### GetVlanConfig

`func (o *VmwareTargetParamsForRecoverDisk) GetVlanConfig() RecoveryVlanConfig`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *VmwareTargetParamsForRecoverDisk) GetVlanConfigOk() (*RecoveryVlanConfig, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *VmwareTargetParamsForRecoverDisk) SetVlanConfig(v RecoveryVlanConfig)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *VmwareTargetParamsForRecoverDisk) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.

### SetVlanConfigNil

`func (o *VmwareTargetParamsForRecoverDisk) SetVlanConfigNil(b bool)`

 SetVlanConfigNil sets the value for VlanConfig to be an explicit nil

### UnsetVlanConfig
`func (o *VmwareTargetParamsForRecoverDisk) UnsetVlanConfig()`

UnsetVlanConfig ensures that no value is present for VlanConfig, not even an explicit nil
### GetPowerOffVms

`func (o *VmwareTargetParamsForRecoverDisk) GetPowerOffVms() bool`

GetPowerOffVms returns the PowerOffVms field if non-nil, zero value otherwise.

### GetPowerOffVmsOk

`func (o *VmwareTargetParamsForRecoverDisk) GetPowerOffVmsOk() (*bool, bool)`

GetPowerOffVmsOk returns a tuple with the PowerOffVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOffVms

`func (o *VmwareTargetParamsForRecoverDisk) SetPowerOffVms(v bool)`

SetPowerOffVms sets PowerOffVms field to given value.

### HasPowerOffVms

`func (o *VmwareTargetParamsForRecoverDisk) HasPowerOffVms() bool`

HasPowerOffVms returns a boolean if a field has been set.

### SetPowerOffVmsNil

`func (o *VmwareTargetParamsForRecoverDisk) SetPowerOffVmsNil(b bool)`

 SetPowerOffVmsNil sets the value for PowerOffVms to be an explicit nil

### UnsetPowerOffVms
`func (o *VmwareTargetParamsForRecoverDisk) UnsetPowerOffVms()`

UnsetPowerOffVms ensures that no value is present for PowerOffVms, not even an explicit nil
### GetPowerOnVms

`func (o *VmwareTargetParamsForRecoverDisk) GetPowerOnVms() bool`

GetPowerOnVms returns the PowerOnVms field if non-nil, zero value otherwise.

### GetPowerOnVmsOk

`func (o *VmwareTargetParamsForRecoverDisk) GetPowerOnVmsOk() (*bool, bool)`

GetPowerOnVmsOk returns a tuple with the PowerOnVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnVms

`func (o *VmwareTargetParamsForRecoverDisk) SetPowerOnVms(v bool)`

SetPowerOnVms sets PowerOnVms field to given value.

### HasPowerOnVms

`func (o *VmwareTargetParamsForRecoverDisk) HasPowerOnVms() bool`

HasPowerOnVms returns a boolean if a field has been set.

### SetPowerOnVmsNil

`func (o *VmwareTargetParamsForRecoverDisk) SetPowerOnVmsNil(b bool)`

 SetPowerOnVmsNil sets the value for PowerOnVms to be an explicit nil

### UnsetPowerOnVms
`func (o *VmwareTargetParamsForRecoverDisk) UnsetPowerOnVms()`

UnsetPowerOnVms ensures that no value is present for PowerOnVms, not even an explicit nil
### GetContinueOnError

`func (o *VmwareTargetParamsForRecoverDisk) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *VmwareTargetParamsForRecoverDisk) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *VmwareTargetParamsForRecoverDisk) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *VmwareTargetParamsForRecoverDisk) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *VmwareTargetParamsForRecoverDisk) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *VmwareTargetParamsForRecoverDisk) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


