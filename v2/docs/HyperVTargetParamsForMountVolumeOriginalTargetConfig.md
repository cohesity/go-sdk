# HyperVTargetParamsForMountVolumeOriginalTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BringDisksOnline** | **NullableBool** | Specifies whether the volumes need to be online within the target environment after attaching the disks. For linux VMs, this should always be set to false because bring disks online is only supported for Windows VM. For Windows, this is optional. If this is set to true, HyperV Integration Services must be installed on the VM. | 
**TargetVmCredentials** | Pointer to [**NullableHyperVMountVolumesOriginalTargetConfigTargetVmCredentials**](HyperVMountVolumesOriginalTargetConfigTargetVmCredentials.md) |  | [optional] 

## Methods

### NewHyperVTargetParamsForMountVolumeOriginalTargetConfig

`func NewHyperVTargetParamsForMountVolumeOriginalTargetConfig(bringDisksOnline NullableBool, ) *HyperVTargetParamsForMountVolumeOriginalTargetConfig`

NewHyperVTargetParamsForMountVolumeOriginalTargetConfig instantiates a new HyperVTargetParamsForMountVolumeOriginalTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperVTargetParamsForMountVolumeOriginalTargetConfigWithDefaults

`func NewHyperVTargetParamsForMountVolumeOriginalTargetConfigWithDefaults() *HyperVTargetParamsForMountVolumeOriginalTargetConfig`

NewHyperVTargetParamsForMountVolumeOriginalTargetConfigWithDefaults instantiates a new HyperVTargetParamsForMountVolumeOriginalTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBringDisksOnline

`func (o *HyperVTargetParamsForMountVolumeOriginalTargetConfig) GetBringDisksOnline() bool`

GetBringDisksOnline returns the BringDisksOnline field if non-nil, zero value otherwise.

### GetBringDisksOnlineOk

`func (o *HyperVTargetParamsForMountVolumeOriginalTargetConfig) GetBringDisksOnlineOk() (*bool, bool)`

GetBringDisksOnlineOk returns a tuple with the BringDisksOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBringDisksOnline

`func (o *HyperVTargetParamsForMountVolumeOriginalTargetConfig) SetBringDisksOnline(v bool)`

SetBringDisksOnline sets BringDisksOnline field to given value.


### SetBringDisksOnlineNil

`func (o *HyperVTargetParamsForMountVolumeOriginalTargetConfig) SetBringDisksOnlineNil(b bool)`

 SetBringDisksOnlineNil sets the value for BringDisksOnline to be an explicit nil

### UnsetBringDisksOnline
`func (o *HyperVTargetParamsForMountVolumeOriginalTargetConfig) UnsetBringDisksOnline()`

UnsetBringDisksOnline ensures that no value is present for BringDisksOnline, not even an explicit nil
### GetTargetVmCredentials

`func (o *HyperVTargetParamsForMountVolumeOriginalTargetConfig) GetTargetVmCredentials() HyperVMountVolumesOriginalTargetConfigTargetVmCredentials`

GetTargetVmCredentials returns the TargetVmCredentials field if non-nil, zero value otherwise.

### GetTargetVmCredentialsOk

`func (o *HyperVTargetParamsForMountVolumeOriginalTargetConfig) GetTargetVmCredentialsOk() (*HyperVMountVolumesOriginalTargetConfigTargetVmCredentials, bool)`

GetTargetVmCredentialsOk returns a tuple with the TargetVmCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVmCredentials

`func (o *HyperVTargetParamsForMountVolumeOriginalTargetConfig) SetTargetVmCredentials(v HyperVMountVolumesOriginalTargetConfigTargetVmCredentials)`

SetTargetVmCredentials sets TargetVmCredentials field to given value.

### HasTargetVmCredentials

`func (o *HyperVTargetParamsForMountVolumeOriginalTargetConfig) HasTargetVmCredentials() bool`

HasTargetVmCredentials returns a boolean if a field has been set.

### SetTargetVmCredentialsNil

`func (o *HyperVTargetParamsForMountVolumeOriginalTargetConfig) SetTargetVmCredentialsNil(b bool)`

 SetTargetVmCredentialsNil sets the value for TargetVmCredentials to be an explicit nil

### UnsetTargetVmCredentials
`func (o *HyperVTargetParamsForMountVolumeOriginalTargetConfig) UnsetTargetVmCredentials()`

UnsetTargetVmCredentials ensures that no value is present for TargetVmCredentials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


