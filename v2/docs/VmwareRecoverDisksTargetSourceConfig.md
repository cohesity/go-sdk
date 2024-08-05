# VmwareRecoverDisksTargetSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disks** | [**[]VmwareRecoverTargetSourceDiskParams**](VmwareRecoverTargetSourceDiskParams.md) | Specifies the disks to be recovered and the location to which they will be recovered. | 
**SourceId** | **NullableInt64** | Specifies the source ID of the VM to which the disks will be restored. | 
**SourceName** | Pointer to **NullableString** | Specifies the source name of the VM to which the disks will be restored. | [optional] 
**Target** | Pointer to **NullableString** | Specifies the name of the vm to which the disks will be restored. | [optional] 

## Methods

### NewVmwareRecoverDisksTargetSourceConfig

`func NewVmwareRecoverDisksTargetSourceConfig(disks []VmwareRecoverTargetSourceDiskParams, sourceId NullableInt64, ) *VmwareRecoverDisksTargetSourceConfig`

NewVmwareRecoverDisksTargetSourceConfig instantiates a new VmwareRecoverDisksTargetSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareRecoverDisksTargetSourceConfigWithDefaults

`func NewVmwareRecoverDisksTargetSourceConfigWithDefaults() *VmwareRecoverDisksTargetSourceConfig`

NewVmwareRecoverDisksTargetSourceConfigWithDefaults instantiates a new VmwareRecoverDisksTargetSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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
### GetSourceName

`func (o *VmwareRecoverDisksTargetSourceConfig) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *VmwareRecoverDisksTargetSourceConfig) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *VmwareRecoverDisksTargetSourceConfig) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *VmwareRecoverDisksTargetSourceConfig) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *VmwareRecoverDisksTargetSourceConfig) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *VmwareRecoverDisksTargetSourceConfig) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetTarget

`func (o *VmwareRecoverDisksTargetSourceConfig) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *VmwareRecoverDisksTargetSourceConfig) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *VmwareRecoverDisksTargetSourceConfig) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *VmwareRecoverDisksTargetSourceConfig) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *VmwareRecoverDisksTargetSourceConfig) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *VmwareRecoverDisksTargetSourceConfig) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


