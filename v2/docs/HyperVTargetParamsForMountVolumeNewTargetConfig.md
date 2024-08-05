# HyperVTargetParamsForMountVolumeNewTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BringDisksOnline** | **NullableBool** | Specifies whether the volumes need to be online within the target environment after attaching the disks. For linux VMs, this should always be set to false because bring disks online is only supported for Windows VM. If this is set to true, HyperV Integration Services must be installed on the VM. | 
**MountTarget** | [**RecoverTarget**](RecoverTarget.md) |  | 
**TargetVmCredentials** | Pointer to [**NullableHyperVMountVolumesNewTargetConfigTargetVmCredentials**](HyperVMountVolumesNewTargetConfigTargetVmCredentials.md) |  | [optional] 

## Methods

### NewHyperVTargetParamsForMountVolumeNewTargetConfig

`func NewHyperVTargetParamsForMountVolumeNewTargetConfig(bringDisksOnline NullableBool, mountTarget RecoverTarget, ) *HyperVTargetParamsForMountVolumeNewTargetConfig`

NewHyperVTargetParamsForMountVolumeNewTargetConfig instantiates a new HyperVTargetParamsForMountVolumeNewTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperVTargetParamsForMountVolumeNewTargetConfigWithDefaults

`func NewHyperVTargetParamsForMountVolumeNewTargetConfigWithDefaults() *HyperVTargetParamsForMountVolumeNewTargetConfig`

NewHyperVTargetParamsForMountVolumeNewTargetConfigWithDefaults instantiates a new HyperVTargetParamsForMountVolumeNewTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBringDisksOnline

`func (o *HyperVTargetParamsForMountVolumeNewTargetConfig) GetBringDisksOnline() bool`

GetBringDisksOnline returns the BringDisksOnline field if non-nil, zero value otherwise.

### GetBringDisksOnlineOk

`func (o *HyperVTargetParamsForMountVolumeNewTargetConfig) GetBringDisksOnlineOk() (*bool, bool)`

GetBringDisksOnlineOk returns a tuple with the BringDisksOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBringDisksOnline

`func (o *HyperVTargetParamsForMountVolumeNewTargetConfig) SetBringDisksOnline(v bool)`

SetBringDisksOnline sets BringDisksOnline field to given value.


### SetBringDisksOnlineNil

`func (o *HyperVTargetParamsForMountVolumeNewTargetConfig) SetBringDisksOnlineNil(b bool)`

 SetBringDisksOnlineNil sets the value for BringDisksOnline to be an explicit nil

### UnsetBringDisksOnline
`func (o *HyperVTargetParamsForMountVolumeNewTargetConfig) UnsetBringDisksOnline()`

UnsetBringDisksOnline ensures that no value is present for BringDisksOnline, not even an explicit nil
### GetMountTarget

`func (o *HyperVTargetParamsForMountVolumeNewTargetConfig) GetMountTarget() RecoverTarget`

GetMountTarget returns the MountTarget field if non-nil, zero value otherwise.

### GetMountTargetOk

`func (o *HyperVTargetParamsForMountVolumeNewTargetConfig) GetMountTargetOk() (*RecoverTarget, bool)`

GetMountTargetOk returns a tuple with the MountTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountTarget

`func (o *HyperVTargetParamsForMountVolumeNewTargetConfig) SetMountTarget(v RecoverTarget)`

SetMountTarget sets MountTarget field to given value.


### GetTargetVmCredentials

`func (o *HyperVTargetParamsForMountVolumeNewTargetConfig) GetTargetVmCredentials() HyperVMountVolumesNewTargetConfigTargetVmCredentials`

GetTargetVmCredentials returns the TargetVmCredentials field if non-nil, zero value otherwise.

### GetTargetVmCredentialsOk

`func (o *HyperVTargetParamsForMountVolumeNewTargetConfig) GetTargetVmCredentialsOk() (*HyperVMountVolumesNewTargetConfigTargetVmCredentials, bool)`

GetTargetVmCredentialsOk returns a tuple with the TargetVmCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVmCredentials

`func (o *HyperVTargetParamsForMountVolumeNewTargetConfig) SetTargetVmCredentials(v HyperVMountVolumesNewTargetConfigTargetVmCredentials)`

SetTargetVmCredentials sets TargetVmCredentials field to given value.

### HasTargetVmCredentials

`func (o *HyperVTargetParamsForMountVolumeNewTargetConfig) HasTargetVmCredentials() bool`

HasTargetVmCredentials returns a boolean if a field has been set.

### SetTargetVmCredentialsNil

`func (o *HyperVTargetParamsForMountVolumeNewTargetConfig) SetTargetVmCredentialsNil(b bool)`

 SetTargetVmCredentialsNil sets the value for TargetVmCredentials to be an explicit nil

### UnsetTargetVmCredentials
`func (o *HyperVTargetParamsForMountVolumeNewTargetConfig) UnsetTargetVmCredentials()`

UnsetTargetVmCredentials ensures that no value is present for TargetVmCredentials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


