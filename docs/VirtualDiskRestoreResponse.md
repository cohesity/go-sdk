# VirtualDiskRestoreResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PowerOffVmBeforeRecovery** | Pointer to **NullableBool** | Specifies whether to power off the VM before recovering virtual disks. | [optional] 
**PowerOnVmAfterRecovery** | Pointer to **NullableBool** | Specifies whether to power on the VM after recovering virtual disks. | [optional] 
**TargetSource** | Pointer to [**ProtectionSource**](ProtectionSource.md) |  | [optional] 
**VirtualDiskMappings** | Pointer to [**[]VirtualDiskMappingResponse**](VirtualDiskMappingResponse.md) | Specifies the list of virtual disks mappings. | [optional] 

## Methods

### NewVirtualDiskRestoreResponse

`func NewVirtualDiskRestoreResponse() *VirtualDiskRestoreResponse`

NewVirtualDiskRestoreResponse instantiates a new VirtualDiskRestoreResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualDiskRestoreResponseWithDefaults

`func NewVirtualDiskRestoreResponseWithDefaults() *VirtualDiskRestoreResponse`

NewVirtualDiskRestoreResponseWithDefaults instantiates a new VirtualDiskRestoreResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPowerOffVmBeforeRecovery

`func (o *VirtualDiskRestoreResponse) GetPowerOffVmBeforeRecovery() bool`

GetPowerOffVmBeforeRecovery returns the PowerOffVmBeforeRecovery field if non-nil, zero value otherwise.

### GetPowerOffVmBeforeRecoveryOk

`func (o *VirtualDiskRestoreResponse) GetPowerOffVmBeforeRecoveryOk() (*bool, bool)`

GetPowerOffVmBeforeRecoveryOk returns a tuple with the PowerOffVmBeforeRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOffVmBeforeRecovery

`func (o *VirtualDiskRestoreResponse) SetPowerOffVmBeforeRecovery(v bool)`

SetPowerOffVmBeforeRecovery sets PowerOffVmBeforeRecovery field to given value.

### HasPowerOffVmBeforeRecovery

`func (o *VirtualDiskRestoreResponse) HasPowerOffVmBeforeRecovery() bool`

HasPowerOffVmBeforeRecovery returns a boolean if a field has been set.

### SetPowerOffVmBeforeRecoveryNil

`func (o *VirtualDiskRestoreResponse) SetPowerOffVmBeforeRecoveryNil(b bool)`

 SetPowerOffVmBeforeRecoveryNil sets the value for PowerOffVmBeforeRecovery to be an explicit nil

### UnsetPowerOffVmBeforeRecovery
`func (o *VirtualDiskRestoreResponse) UnsetPowerOffVmBeforeRecovery()`

UnsetPowerOffVmBeforeRecovery ensures that no value is present for PowerOffVmBeforeRecovery, not even an explicit nil
### GetPowerOnVmAfterRecovery

`func (o *VirtualDiskRestoreResponse) GetPowerOnVmAfterRecovery() bool`

GetPowerOnVmAfterRecovery returns the PowerOnVmAfterRecovery field if non-nil, zero value otherwise.

### GetPowerOnVmAfterRecoveryOk

`func (o *VirtualDiskRestoreResponse) GetPowerOnVmAfterRecoveryOk() (*bool, bool)`

GetPowerOnVmAfterRecoveryOk returns a tuple with the PowerOnVmAfterRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnVmAfterRecovery

`func (o *VirtualDiskRestoreResponse) SetPowerOnVmAfterRecovery(v bool)`

SetPowerOnVmAfterRecovery sets PowerOnVmAfterRecovery field to given value.

### HasPowerOnVmAfterRecovery

`func (o *VirtualDiskRestoreResponse) HasPowerOnVmAfterRecovery() bool`

HasPowerOnVmAfterRecovery returns a boolean if a field has been set.

### SetPowerOnVmAfterRecoveryNil

`func (o *VirtualDiskRestoreResponse) SetPowerOnVmAfterRecoveryNil(b bool)`

 SetPowerOnVmAfterRecoveryNil sets the value for PowerOnVmAfterRecovery to be an explicit nil

### UnsetPowerOnVmAfterRecovery
`func (o *VirtualDiskRestoreResponse) UnsetPowerOnVmAfterRecovery()`

UnsetPowerOnVmAfterRecovery ensures that no value is present for PowerOnVmAfterRecovery, not even an explicit nil
### GetTargetSource

`func (o *VirtualDiskRestoreResponse) GetTargetSource() ProtectionSource`

GetTargetSource returns the TargetSource field if non-nil, zero value otherwise.

### GetTargetSourceOk

`func (o *VirtualDiskRestoreResponse) GetTargetSourceOk() (*ProtectionSource, bool)`

GetTargetSourceOk returns a tuple with the TargetSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSource

`func (o *VirtualDiskRestoreResponse) SetTargetSource(v ProtectionSource)`

SetTargetSource sets TargetSource field to given value.

### HasTargetSource

`func (o *VirtualDiskRestoreResponse) HasTargetSource() bool`

HasTargetSource returns a boolean if a field has been set.

### GetVirtualDiskMappings

`func (o *VirtualDiskRestoreResponse) GetVirtualDiskMappings() []VirtualDiskMappingResponse`

GetVirtualDiskMappings returns the VirtualDiskMappings field if non-nil, zero value otherwise.

### GetVirtualDiskMappingsOk

`func (o *VirtualDiskRestoreResponse) GetVirtualDiskMappingsOk() (*[]VirtualDiskMappingResponse, bool)`

GetVirtualDiskMappingsOk returns a tuple with the VirtualDiskMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDiskMappings

`func (o *VirtualDiskRestoreResponse) SetVirtualDiskMappings(v []VirtualDiskMappingResponse)`

SetVirtualDiskMappings sets VirtualDiskMappings field to given value.

### HasVirtualDiskMappings

`func (o *VirtualDiskRestoreResponse) HasVirtualDiskMappings() bool`

HasVirtualDiskMappings returns a boolean if a field has been set.

### SetVirtualDiskMappingsNil

`func (o *VirtualDiskRestoreResponse) SetVirtualDiskMappingsNil(b bool)`

 SetVirtualDiskMappingsNil sets the value for VirtualDiskMappings to be an explicit nil

### UnsetVirtualDiskMappings
`func (o *VirtualDiskRestoreResponse) UnsetVirtualDiskMappings()`

UnsetVirtualDiskMappings ensures that no value is present for VirtualDiskMappings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


