# VmwareTargetParamsForMountVolumeOriginalTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BringDisksOnline** | **NullableBool** | Specifies whether the volumes need to be online within the target environment after attaching the disks. For linux VMs, this should always be set to true. For Windows, this is optional. If this is set to true, VMware tools must be installed on the VM. If this is set to false, useExistingAgent and targetCredentials are not needed. | 
**TargetVmCredentials** | Pointer to [**NullableVMwareMountVolumesOriginalTargetConfigTargetVmCredentials**](VMwareMountVolumesOriginalTargetConfigTargetVmCredentials.md) |  | [optional] 
**UseExistingAgent** | Pointer to **NullableBool** | Specifies whether this will use an existing agent on the target vm or will deploy a new agent. This is required if bringDisksOnline is set to true. | [optional] 

## Methods

### NewVmwareTargetParamsForMountVolumeOriginalTargetConfig

`func NewVmwareTargetParamsForMountVolumeOriginalTargetConfig(bringDisksOnline NullableBool, ) *VmwareTargetParamsForMountVolumeOriginalTargetConfig`

NewVmwareTargetParamsForMountVolumeOriginalTargetConfig instantiates a new VmwareTargetParamsForMountVolumeOriginalTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareTargetParamsForMountVolumeOriginalTargetConfigWithDefaults

`func NewVmwareTargetParamsForMountVolumeOriginalTargetConfigWithDefaults() *VmwareTargetParamsForMountVolumeOriginalTargetConfig`

NewVmwareTargetParamsForMountVolumeOriginalTargetConfigWithDefaults instantiates a new VmwareTargetParamsForMountVolumeOriginalTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBringDisksOnline

`func (o *VmwareTargetParamsForMountVolumeOriginalTargetConfig) GetBringDisksOnline() bool`

GetBringDisksOnline returns the BringDisksOnline field if non-nil, zero value otherwise.

### GetBringDisksOnlineOk

`func (o *VmwareTargetParamsForMountVolumeOriginalTargetConfig) GetBringDisksOnlineOk() (*bool, bool)`

GetBringDisksOnlineOk returns a tuple with the BringDisksOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBringDisksOnline

`func (o *VmwareTargetParamsForMountVolumeOriginalTargetConfig) SetBringDisksOnline(v bool)`

SetBringDisksOnline sets BringDisksOnline field to given value.


### SetBringDisksOnlineNil

`func (o *VmwareTargetParamsForMountVolumeOriginalTargetConfig) SetBringDisksOnlineNil(b bool)`

 SetBringDisksOnlineNil sets the value for BringDisksOnline to be an explicit nil

### UnsetBringDisksOnline
`func (o *VmwareTargetParamsForMountVolumeOriginalTargetConfig) UnsetBringDisksOnline()`

UnsetBringDisksOnline ensures that no value is present for BringDisksOnline, not even an explicit nil
### GetTargetVmCredentials

`func (o *VmwareTargetParamsForMountVolumeOriginalTargetConfig) GetTargetVmCredentials() VMwareMountVolumesOriginalTargetConfigTargetVmCredentials`

GetTargetVmCredentials returns the TargetVmCredentials field if non-nil, zero value otherwise.

### GetTargetVmCredentialsOk

`func (o *VmwareTargetParamsForMountVolumeOriginalTargetConfig) GetTargetVmCredentialsOk() (*VMwareMountVolumesOriginalTargetConfigTargetVmCredentials, bool)`

GetTargetVmCredentialsOk returns a tuple with the TargetVmCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVmCredentials

`func (o *VmwareTargetParamsForMountVolumeOriginalTargetConfig) SetTargetVmCredentials(v VMwareMountVolumesOriginalTargetConfigTargetVmCredentials)`

SetTargetVmCredentials sets TargetVmCredentials field to given value.

### HasTargetVmCredentials

`func (o *VmwareTargetParamsForMountVolumeOriginalTargetConfig) HasTargetVmCredentials() bool`

HasTargetVmCredentials returns a boolean if a field has been set.

### SetTargetVmCredentialsNil

`func (o *VmwareTargetParamsForMountVolumeOriginalTargetConfig) SetTargetVmCredentialsNil(b bool)`

 SetTargetVmCredentialsNil sets the value for TargetVmCredentials to be an explicit nil

### UnsetTargetVmCredentials
`func (o *VmwareTargetParamsForMountVolumeOriginalTargetConfig) UnsetTargetVmCredentials()`

UnsetTargetVmCredentials ensures that no value is present for TargetVmCredentials, not even an explicit nil
### GetUseExistingAgent

`func (o *VmwareTargetParamsForMountVolumeOriginalTargetConfig) GetUseExistingAgent() bool`

GetUseExistingAgent returns the UseExistingAgent field if non-nil, zero value otherwise.

### GetUseExistingAgentOk

`func (o *VmwareTargetParamsForMountVolumeOriginalTargetConfig) GetUseExistingAgentOk() (*bool, bool)`

GetUseExistingAgentOk returns a tuple with the UseExistingAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseExistingAgent

`func (o *VmwareTargetParamsForMountVolumeOriginalTargetConfig) SetUseExistingAgent(v bool)`

SetUseExistingAgent sets UseExistingAgent field to given value.

### HasUseExistingAgent

`func (o *VmwareTargetParamsForMountVolumeOriginalTargetConfig) HasUseExistingAgent() bool`

HasUseExistingAgent returns a boolean if a field has been set.

### SetUseExistingAgentNil

`func (o *VmwareTargetParamsForMountVolumeOriginalTargetConfig) SetUseExistingAgentNil(b bool)`

 SetUseExistingAgentNil sets the value for UseExistingAgent to be an explicit nil

### UnsetUseExistingAgent
`func (o *VmwareTargetParamsForMountVolumeOriginalTargetConfig) UnsetUseExistingAgent()`

UnsetUseExistingAgent ensures that no value is present for UseExistingAgent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


