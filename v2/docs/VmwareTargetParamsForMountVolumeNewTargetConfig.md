# VmwareTargetParamsForMountVolumeNewTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BringDisksOnline** | **NullableBool** | Specifies whether the volumes need to be online within the target environment after attaching the disks. For linux VMs, this should always be set to true. For Windows, this is optional. If this is set to true, VMware tools must be installed on the VM. If this is set to false, useExistingAgent and targetCredentials are not needed. | 
**MountTarget** | [**RecoverTarget**](RecoverTarget.md) |  | 
**TargetVmCredentials** | Pointer to [**NullableVMwareMountVolumesNewTargetConfigTargetVmCredentials**](VMwareMountVolumesNewTargetConfigTargetVmCredentials.md) |  | [optional] 
**UseExistingAgent** | Pointer to **NullableBool** | Specifies whether this will use an existing agent on the target vm or will deploy a new agent. This is required if bringDisksOnline is set to true. | [optional] 

## Methods

### NewVmwareTargetParamsForMountVolumeNewTargetConfig

`func NewVmwareTargetParamsForMountVolumeNewTargetConfig(bringDisksOnline NullableBool, mountTarget RecoverTarget, ) *VmwareTargetParamsForMountVolumeNewTargetConfig`

NewVmwareTargetParamsForMountVolumeNewTargetConfig instantiates a new VmwareTargetParamsForMountVolumeNewTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareTargetParamsForMountVolumeNewTargetConfigWithDefaults

`func NewVmwareTargetParamsForMountVolumeNewTargetConfigWithDefaults() *VmwareTargetParamsForMountVolumeNewTargetConfig`

NewVmwareTargetParamsForMountVolumeNewTargetConfigWithDefaults instantiates a new VmwareTargetParamsForMountVolumeNewTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBringDisksOnline

`func (o *VmwareTargetParamsForMountVolumeNewTargetConfig) GetBringDisksOnline() bool`

GetBringDisksOnline returns the BringDisksOnline field if non-nil, zero value otherwise.

### GetBringDisksOnlineOk

`func (o *VmwareTargetParamsForMountVolumeNewTargetConfig) GetBringDisksOnlineOk() (*bool, bool)`

GetBringDisksOnlineOk returns a tuple with the BringDisksOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBringDisksOnline

`func (o *VmwareTargetParamsForMountVolumeNewTargetConfig) SetBringDisksOnline(v bool)`

SetBringDisksOnline sets BringDisksOnline field to given value.


### SetBringDisksOnlineNil

`func (o *VmwareTargetParamsForMountVolumeNewTargetConfig) SetBringDisksOnlineNil(b bool)`

 SetBringDisksOnlineNil sets the value for BringDisksOnline to be an explicit nil

### UnsetBringDisksOnline
`func (o *VmwareTargetParamsForMountVolumeNewTargetConfig) UnsetBringDisksOnline()`

UnsetBringDisksOnline ensures that no value is present for BringDisksOnline, not even an explicit nil
### GetMountTarget

`func (o *VmwareTargetParamsForMountVolumeNewTargetConfig) GetMountTarget() RecoverTarget`

GetMountTarget returns the MountTarget field if non-nil, zero value otherwise.

### GetMountTargetOk

`func (o *VmwareTargetParamsForMountVolumeNewTargetConfig) GetMountTargetOk() (*RecoverTarget, bool)`

GetMountTargetOk returns a tuple with the MountTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountTarget

`func (o *VmwareTargetParamsForMountVolumeNewTargetConfig) SetMountTarget(v RecoverTarget)`

SetMountTarget sets MountTarget field to given value.


### GetTargetVmCredentials

`func (o *VmwareTargetParamsForMountVolumeNewTargetConfig) GetTargetVmCredentials() VMwareMountVolumesNewTargetConfigTargetVmCredentials`

GetTargetVmCredentials returns the TargetVmCredentials field if non-nil, zero value otherwise.

### GetTargetVmCredentialsOk

`func (o *VmwareTargetParamsForMountVolumeNewTargetConfig) GetTargetVmCredentialsOk() (*VMwareMountVolumesNewTargetConfigTargetVmCredentials, bool)`

GetTargetVmCredentialsOk returns a tuple with the TargetVmCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVmCredentials

`func (o *VmwareTargetParamsForMountVolumeNewTargetConfig) SetTargetVmCredentials(v VMwareMountVolumesNewTargetConfigTargetVmCredentials)`

SetTargetVmCredentials sets TargetVmCredentials field to given value.

### HasTargetVmCredentials

`func (o *VmwareTargetParamsForMountVolumeNewTargetConfig) HasTargetVmCredentials() bool`

HasTargetVmCredentials returns a boolean if a field has been set.

### SetTargetVmCredentialsNil

`func (o *VmwareTargetParamsForMountVolumeNewTargetConfig) SetTargetVmCredentialsNil(b bool)`

 SetTargetVmCredentialsNil sets the value for TargetVmCredentials to be an explicit nil

### UnsetTargetVmCredentials
`func (o *VmwareTargetParamsForMountVolumeNewTargetConfig) UnsetTargetVmCredentials()`

UnsetTargetVmCredentials ensures that no value is present for TargetVmCredentials, not even an explicit nil
### GetUseExistingAgent

`func (o *VmwareTargetParamsForMountVolumeNewTargetConfig) GetUseExistingAgent() bool`

GetUseExistingAgent returns the UseExistingAgent field if non-nil, zero value otherwise.

### GetUseExistingAgentOk

`func (o *VmwareTargetParamsForMountVolumeNewTargetConfig) GetUseExistingAgentOk() (*bool, bool)`

GetUseExistingAgentOk returns a tuple with the UseExistingAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseExistingAgent

`func (o *VmwareTargetParamsForMountVolumeNewTargetConfig) SetUseExistingAgent(v bool)`

SetUseExistingAgent sets UseExistingAgent field to given value.

### HasUseExistingAgent

`func (o *VmwareTargetParamsForMountVolumeNewTargetConfig) HasUseExistingAgent() bool`

HasUseExistingAgent returns a boolean if a field has been set.

### SetUseExistingAgentNil

`func (o *VmwareTargetParamsForMountVolumeNewTargetConfig) SetUseExistingAgentNil(b bool)`

 SetUseExistingAgentNil sets the value for UseExistingAgent to be an explicit nil

### UnsetUseExistingAgent
`func (o *VmwareTargetParamsForMountVolumeNewTargetConfig) UnsetUseExistingAgent()`

UnsetUseExistingAgent ensures that no value is present for UseExistingAgent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


