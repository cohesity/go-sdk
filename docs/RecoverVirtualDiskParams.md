# RecoverVirtualDiskParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PowerOffVmBeforeRecovery** | Pointer to **NullableBool** | Whether to power-off the VM before recovering virtual disks. | [optional] 
**PowerOnVmAfterRecovery** | Pointer to **NullableBool** | Whether to power-on the VM after recovering virtual disks. | [optional] 
**TargetEntity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**VirtualDiskMappings** | Pointer to [**[]RecoverVirtualDiskParamsVirtualDiskMapping**](RecoverVirtualDiskParamsVirtualDiskMapping.md) |  | [optional] 

## Methods

### NewRecoverVirtualDiskParams

`func NewRecoverVirtualDiskParams() *RecoverVirtualDiskParams`

NewRecoverVirtualDiskParams instantiates a new RecoverVirtualDiskParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverVirtualDiskParamsWithDefaults

`func NewRecoverVirtualDiskParamsWithDefaults() *RecoverVirtualDiskParams`

NewRecoverVirtualDiskParamsWithDefaults instantiates a new RecoverVirtualDiskParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPowerOffVmBeforeRecovery

`func (o *RecoverVirtualDiskParams) GetPowerOffVmBeforeRecovery() bool`

GetPowerOffVmBeforeRecovery returns the PowerOffVmBeforeRecovery field if non-nil, zero value otherwise.

### GetPowerOffVmBeforeRecoveryOk

`func (o *RecoverVirtualDiskParams) GetPowerOffVmBeforeRecoveryOk() (*bool, bool)`

GetPowerOffVmBeforeRecoveryOk returns a tuple with the PowerOffVmBeforeRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOffVmBeforeRecovery

`func (o *RecoverVirtualDiskParams) SetPowerOffVmBeforeRecovery(v bool)`

SetPowerOffVmBeforeRecovery sets PowerOffVmBeforeRecovery field to given value.

### HasPowerOffVmBeforeRecovery

`func (o *RecoverVirtualDiskParams) HasPowerOffVmBeforeRecovery() bool`

HasPowerOffVmBeforeRecovery returns a boolean if a field has been set.

### SetPowerOffVmBeforeRecoveryNil

`func (o *RecoverVirtualDiskParams) SetPowerOffVmBeforeRecoveryNil(b bool)`

 SetPowerOffVmBeforeRecoveryNil sets the value for PowerOffVmBeforeRecovery to be an explicit nil

### UnsetPowerOffVmBeforeRecovery
`func (o *RecoverVirtualDiskParams) UnsetPowerOffVmBeforeRecovery()`

UnsetPowerOffVmBeforeRecovery ensures that no value is present for PowerOffVmBeforeRecovery, not even an explicit nil
### GetPowerOnVmAfterRecovery

`func (o *RecoverVirtualDiskParams) GetPowerOnVmAfterRecovery() bool`

GetPowerOnVmAfterRecovery returns the PowerOnVmAfterRecovery field if non-nil, zero value otherwise.

### GetPowerOnVmAfterRecoveryOk

`func (o *RecoverVirtualDiskParams) GetPowerOnVmAfterRecoveryOk() (*bool, bool)`

GetPowerOnVmAfterRecoveryOk returns a tuple with the PowerOnVmAfterRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnVmAfterRecovery

`func (o *RecoverVirtualDiskParams) SetPowerOnVmAfterRecovery(v bool)`

SetPowerOnVmAfterRecovery sets PowerOnVmAfterRecovery field to given value.

### HasPowerOnVmAfterRecovery

`func (o *RecoverVirtualDiskParams) HasPowerOnVmAfterRecovery() bool`

HasPowerOnVmAfterRecovery returns a boolean if a field has been set.

### SetPowerOnVmAfterRecoveryNil

`func (o *RecoverVirtualDiskParams) SetPowerOnVmAfterRecoveryNil(b bool)`

 SetPowerOnVmAfterRecoveryNil sets the value for PowerOnVmAfterRecovery to be an explicit nil

### UnsetPowerOnVmAfterRecovery
`func (o *RecoverVirtualDiskParams) UnsetPowerOnVmAfterRecovery()`

UnsetPowerOnVmAfterRecovery ensures that no value is present for PowerOnVmAfterRecovery, not even an explicit nil
### GetTargetEntity

`func (o *RecoverVirtualDiskParams) GetTargetEntity() EntityProto`

GetTargetEntity returns the TargetEntity field if non-nil, zero value otherwise.

### GetTargetEntityOk

`func (o *RecoverVirtualDiskParams) GetTargetEntityOk() (*EntityProto, bool)`

GetTargetEntityOk returns a tuple with the TargetEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEntity

`func (o *RecoverVirtualDiskParams) SetTargetEntity(v EntityProto)`

SetTargetEntity sets TargetEntity field to given value.

### HasTargetEntity

`func (o *RecoverVirtualDiskParams) HasTargetEntity() bool`

HasTargetEntity returns a boolean if a field has been set.

### GetVirtualDiskMappings

`func (o *RecoverVirtualDiskParams) GetVirtualDiskMappings() []RecoverVirtualDiskParamsVirtualDiskMapping`

GetVirtualDiskMappings returns the VirtualDiskMappings field if non-nil, zero value otherwise.

### GetVirtualDiskMappingsOk

`func (o *RecoverVirtualDiskParams) GetVirtualDiskMappingsOk() (*[]RecoverVirtualDiskParamsVirtualDiskMapping, bool)`

GetVirtualDiskMappingsOk returns a tuple with the VirtualDiskMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDiskMappings

`func (o *RecoverVirtualDiskParams) SetVirtualDiskMappings(v []RecoverVirtualDiskParamsVirtualDiskMapping)`

SetVirtualDiskMappings sets VirtualDiskMappings field to given value.

### HasVirtualDiskMappings

`func (o *RecoverVirtualDiskParams) HasVirtualDiskMappings() bool`

HasVirtualDiskMappings returns a boolean if a field has been set.

### SetVirtualDiskMappingsNil

`func (o *RecoverVirtualDiskParams) SetVirtualDiskMappingsNil(b bool)`

 SetVirtualDiskMappingsNil sets the value for VirtualDiskMappings to be an explicit nil

### UnsetVirtualDiskMappings
`func (o *RecoverVirtualDiskParams) UnsetVirtualDiskMappings()`

UnsetVirtualDiskMappings ensures that no value is present for VirtualDiskMappings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


