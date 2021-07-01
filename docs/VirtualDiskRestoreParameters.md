# VirtualDiskRestoreParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PowerOffVmBeforeRecovery** | Pointer to **NullableBool** | Specifies whether to power off the VM before recovering virtual disks. | [optional] 
**PowerOnVmAfterRecovery** | Pointer to **NullableBool** | Specifies whether to power on the VM after recovering virtual disks. | [optional] 
**TargetSourceId** | Pointer to **NullableInt64** | Specifies the target entity to which the disks should be attached. | [optional] 
**VirtualDiskMappings** | Pointer to [**[]VirtualDiskMapping**](VirtualDiskMapping.md) | Specifies the list of virtual disks mappings. | [optional] 

## Methods

### NewVirtualDiskRestoreParameters

`func NewVirtualDiskRestoreParameters() *VirtualDiskRestoreParameters`

NewVirtualDiskRestoreParameters instantiates a new VirtualDiskRestoreParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualDiskRestoreParametersWithDefaults

`func NewVirtualDiskRestoreParametersWithDefaults() *VirtualDiskRestoreParameters`

NewVirtualDiskRestoreParametersWithDefaults instantiates a new VirtualDiskRestoreParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPowerOffVmBeforeRecovery

`func (o *VirtualDiskRestoreParameters) GetPowerOffVmBeforeRecovery() bool`

GetPowerOffVmBeforeRecovery returns the PowerOffVmBeforeRecovery field if non-nil, zero value otherwise.

### GetPowerOffVmBeforeRecoveryOk

`func (o *VirtualDiskRestoreParameters) GetPowerOffVmBeforeRecoveryOk() (*bool, bool)`

GetPowerOffVmBeforeRecoveryOk returns a tuple with the PowerOffVmBeforeRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOffVmBeforeRecovery

`func (o *VirtualDiskRestoreParameters) SetPowerOffVmBeforeRecovery(v bool)`

SetPowerOffVmBeforeRecovery sets PowerOffVmBeforeRecovery field to given value.

### HasPowerOffVmBeforeRecovery

`func (o *VirtualDiskRestoreParameters) HasPowerOffVmBeforeRecovery() bool`

HasPowerOffVmBeforeRecovery returns a boolean if a field has been set.

### SetPowerOffVmBeforeRecoveryNil

`func (o *VirtualDiskRestoreParameters) SetPowerOffVmBeforeRecoveryNil(b bool)`

 SetPowerOffVmBeforeRecoveryNil sets the value for PowerOffVmBeforeRecovery to be an explicit nil

### UnsetPowerOffVmBeforeRecovery
`func (o *VirtualDiskRestoreParameters) UnsetPowerOffVmBeforeRecovery()`

UnsetPowerOffVmBeforeRecovery ensures that no value is present for PowerOffVmBeforeRecovery, not even an explicit nil
### GetPowerOnVmAfterRecovery

`func (o *VirtualDiskRestoreParameters) GetPowerOnVmAfterRecovery() bool`

GetPowerOnVmAfterRecovery returns the PowerOnVmAfterRecovery field if non-nil, zero value otherwise.

### GetPowerOnVmAfterRecoveryOk

`func (o *VirtualDiskRestoreParameters) GetPowerOnVmAfterRecoveryOk() (*bool, bool)`

GetPowerOnVmAfterRecoveryOk returns a tuple with the PowerOnVmAfterRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnVmAfterRecovery

`func (o *VirtualDiskRestoreParameters) SetPowerOnVmAfterRecovery(v bool)`

SetPowerOnVmAfterRecovery sets PowerOnVmAfterRecovery field to given value.

### HasPowerOnVmAfterRecovery

`func (o *VirtualDiskRestoreParameters) HasPowerOnVmAfterRecovery() bool`

HasPowerOnVmAfterRecovery returns a boolean if a field has been set.

### SetPowerOnVmAfterRecoveryNil

`func (o *VirtualDiskRestoreParameters) SetPowerOnVmAfterRecoveryNil(b bool)`

 SetPowerOnVmAfterRecoveryNil sets the value for PowerOnVmAfterRecovery to be an explicit nil

### UnsetPowerOnVmAfterRecovery
`func (o *VirtualDiskRestoreParameters) UnsetPowerOnVmAfterRecovery()`

UnsetPowerOnVmAfterRecovery ensures that no value is present for PowerOnVmAfterRecovery, not even an explicit nil
### GetTargetSourceId

`func (o *VirtualDiskRestoreParameters) GetTargetSourceId() int64`

GetTargetSourceId returns the TargetSourceId field if non-nil, zero value otherwise.

### GetTargetSourceIdOk

`func (o *VirtualDiskRestoreParameters) GetTargetSourceIdOk() (*int64, bool)`

GetTargetSourceIdOk returns a tuple with the TargetSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSourceId

`func (o *VirtualDiskRestoreParameters) SetTargetSourceId(v int64)`

SetTargetSourceId sets TargetSourceId field to given value.

### HasTargetSourceId

`func (o *VirtualDiskRestoreParameters) HasTargetSourceId() bool`

HasTargetSourceId returns a boolean if a field has been set.

### SetTargetSourceIdNil

`func (o *VirtualDiskRestoreParameters) SetTargetSourceIdNil(b bool)`

 SetTargetSourceIdNil sets the value for TargetSourceId to be an explicit nil

### UnsetTargetSourceId
`func (o *VirtualDiskRestoreParameters) UnsetTargetSourceId()`

UnsetTargetSourceId ensures that no value is present for TargetSourceId, not even an explicit nil
### GetVirtualDiskMappings

`func (o *VirtualDiskRestoreParameters) GetVirtualDiskMappings() []VirtualDiskMapping`

GetVirtualDiskMappings returns the VirtualDiskMappings field if non-nil, zero value otherwise.

### GetVirtualDiskMappingsOk

`func (o *VirtualDiskRestoreParameters) GetVirtualDiskMappingsOk() (*[]VirtualDiskMapping, bool)`

GetVirtualDiskMappingsOk returns a tuple with the VirtualDiskMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDiskMappings

`func (o *VirtualDiskRestoreParameters) SetVirtualDiskMappings(v []VirtualDiskMapping)`

SetVirtualDiskMappings sets VirtualDiskMappings field to given value.

### HasVirtualDiskMappings

`func (o *VirtualDiskRestoreParameters) HasVirtualDiskMappings() bool`

HasVirtualDiskMappings returns a boolean if a field has been set.

### SetVirtualDiskMappingsNil

`func (o *VirtualDiskRestoreParameters) SetVirtualDiskMappingsNil(b bool)`

 SetVirtualDiskMappingsNil sets the value for VirtualDiskMappings to be an explicit nil

### UnsetVirtualDiskMappings
`func (o *VirtualDiskRestoreParameters) UnsetVirtualDiskMappings()`

UnsetVirtualDiskMappings ensures that no value is present for VirtualDiskMappings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


