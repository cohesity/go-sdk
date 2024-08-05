# HyperVMountVolumesOriginalTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BringDisksOnline** | **NullableBool** | Specifies whether the volumes need to be online within the target environment after attaching the disks. For linux VMs, this should always be set to false because bring disks online is only supported for Windows VM. For Windows, this is optional. If this is set to true, HyperV Integration Services must be installed on the VM. | 
**TargetVmCredentials** | Pointer to [**NullableHyperVMountVolumesOriginalTargetConfigTargetVmCredentials**](HyperVMountVolumesOriginalTargetConfigTargetVmCredentials.md) |  | [optional] 

## Methods

### NewHyperVMountVolumesOriginalTargetConfig

`func NewHyperVMountVolumesOriginalTargetConfig(bringDisksOnline NullableBool, ) *HyperVMountVolumesOriginalTargetConfig`

NewHyperVMountVolumesOriginalTargetConfig instantiates a new HyperVMountVolumesOriginalTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperVMountVolumesOriginalTargetConfigWithDefaults

`func NewHyperVMountVolumesOriginalTargetConfigWithDefaults() *HyperVMountVolumesOriginalTargetConfig`

NewHyperVMountVolumesOriginalTargetConfigWithDefaults instantiates a new HyperVMountVolumesOriginalTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBringDisksOnline

`func (o *HyperVMountVolumesOriginalTargetConfig) GetBringDisksOnline() bool`

GetBringDisksOnline returns the BringDisksOnline field if non-nil, zero value otherwise.

### GetBringDisksOnlineOk

`func (o *HyperVMountVolumesOriginalTargetConfig) GetBringDisksOnlineOk() (*bool, bool)`

GetBringDisksOnlineOk returns a tuple with the BringDisksOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBringDisksOnline

`func (o *HyperVMountVolumesOriginalTargetConfig) SetBringDisksOnline(v bool)`

SetBringDisksOnline sets BringDisksOnline field to given value.


### SetBringDisksOnlineNil

`func (o *HyperVMountVolumesOriginalTargetConfig) SetBringDisksOnlineNil(b bool)`

 SetBringDisksOnlineNil sets the value for BringDisksOnline to be an explicit nil

### UnsetBringDisksOnline
`func (o *HyperVMountVolumesOriginalTargetConfig) UnsetBringDisksOnline()`

UnsetBringDisksOnline ensures that no value is present for BringDisksOnline, not even an explicit nil
### GetTargetVmCredentials

`func (o *HyperVMountVolumesOriginalTargetConfig) GetTargetVmCredentials() HyperVMountVolumesOriginalTargetConfigTargetVmCredentials`

GetTargetVmCredentials returns the TargetVmCredentials field if non-nil, zero value otherwise.

### GetTargetVmCredentialsOk

`func (o *HyperVMountVolumesOriginalTargetConfig) GetTargetVmCredentialsOk() (*HyperVMountVolumesOriginalTargetConfigTargetVmCredentials, bool)`

GetTargetVmCredentialsOk returns a tuple with the TargetVmCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVmCredentials

`func (o *HyperVMountVolumesOriginalTargetConfig) SetTargetVmCredentials(v HyperVMountVolumesOriginalTargetConfigTargetVmCredentials)`

SetTargetVmCredentials sets TargetVmCredentials field to given value.

### HasTargetVmCredentials

`func (o *HyperVMountVolumesOriginalTargetConfig) HasTargetVmCredentials() bool`

HasTargetVmCredentials returns a boolean if a field has been set.

### SetTargetVmCredentialsNil

`func (o *HyperVMountVolumesOriginalTargetConfig) SetTargetVmCredentialsNil(b bool)`

 SetTargetVmCredentialsNil sets the value for TargetVmCredentials to be an explicit nil

### UnsetTargetVmCredentials
`func (o *HyperVMountVolumesOriginalTargetConfig) UnsetTargetVmCredentials()`

UnsetTargetVmCredentials ensures that no value is present for TargetVmCredentials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


