# VmwareRecoverDisksOriginalSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disks** | [**[]VmwareRecoverOriginalSourceDiskParams**](VmwareRecoverOriginalSourceDiskParams.md) | Specifies the disks to be recovered and the location to which they will be recovered. | 

## Methods

### NewVmwareRecoverDisksOriginalSourceConfig

`func NewVmwareRecoverDisksOriginalSourceConfig(disks []VmwareRecoverOriginalSourceDiskParams, ) *VmwareRecoverDisksOriginalSourceConfig`

NewVmwareRecoverDisksOriginalSourceConfig instantiates a new VmwareRecoverDisksOriginalSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareRecoverDisksOriginalSourceConfigWithDefaults

`func NewVmwareRecoverDisksOriginalSourceConfigWithDefaults() *VmwareRecoverDisksOriginalSourceConfig`

NewVmwareRecoverDisksOriginalSourceConfigWithDefaults instantiates a new VmwareRecoverDisksOriginalSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisks

`func (o *VmwareRecoverDisksOriginalSourceConfig) GetDisks() []VmwareRecoverOriginalSourceDiskParams`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *VmwareRecoverDisksOriginalSourceConfig) GetDisksOk() (*[]VmwareRecoverOriginalSourceDiskParams, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *VmwareRecoverDisksOriginalSourceConfig) SetDisks(v []VmwareRecoverOriginalSourceDiskParams)`

SetDisks sets Disks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


