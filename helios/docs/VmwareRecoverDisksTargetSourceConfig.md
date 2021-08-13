# VmwareRecoverDisksTargetSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceId** | **NullableInt64** | Specifies the source ID of the VM to which the disks will be restored. | 
**Disks** | [**[]VmwareRecoverTargetSourceDiskParams**](VmwareRecoverTargetSourceDiskParams.md) | Specifies the disks to be recovered and the location to which they will be recovered. | 

## Methods

### NewVmwareRecoverDisksTargetSourceConfig

`func NewVmwareRecoverDisksTargetSourceConfig(sourceId NullableInt64, disks []VmwareRecoverTargetSourceDiskParams, ) *VmwareRecoverDisksTargetSourceConfig`

NewVmwareRecoverDisksTargetSourceConfig instantiates a new VmwareRecoverDisksTargetSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareRecoverDisksTargetSourceConfigWithDefaults

`func NewVmwareRecoverDisksTargetSourceConfigWithDefaults() *VmwareRecoverDisksTargetSourceConfig`

NewVmwareRecoverDisksTargetSourceConfigWithDefaults instantiates a new VmwareRecoverDisksTargetSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceId

`func (o *VmwareRecoverDisksTargetSourceConfig) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *VmwareRecoverDisksTargetSourceConfig) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *VmwareRecoverDisksTargetSourceConfig) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.


### SetSourceIdNil

`func (o *VmwareRecoverDisksTargetSourceConfig) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *VmwareRecoverDisksTargetSourceConfig) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetDisks

`func (o *VmwareRecoverDisksTargetSourceConfig) GetDisks() []VmwareRecoverTargetSourceDiskParams`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *VmwareRecoverDisksTargetSourceConfig) GetDisksOk() (*[]VmwareRecoverTargetSourceDiskParams, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *VmwareRecoverDisksTargetSourceConfig) SetDisks(v []VmwareRecoverTargetSourceDiskParams)`

SetDisks sets Disks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


