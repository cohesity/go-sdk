# VirtualDiskId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ControllerBusNumber** | Pointer to **NullableInt64** | Controller&#39;s bus-id controlling the virtual disk in question. | [optional] 
**ControllerType** | Pointer to **NullableString** | Controller&#39;s type (SCSI, IDE etc). | [optional] 
**DiskId** | Pointer to **NullableString** | Original disk id. This is sufficient to identify the disk information, but in some scenarios, user&#39;s may specify the controller option instead. | [optional] 
**UnitNumber** | Pointer to **NullableInt64** | Disk unit number to identify the virtual disk within a controller. | [optional] 

## Methods

### NewVirtualDiskId

`func NewVirtualDiskId() *VirtualDiskId`

NewVirtualDiskId instantiates a new VirtualDiskId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualDiskIdWithDefaults

`func NewVirtualDiskIdWithDefaults() *VirtualDiskId`

NewVirtualDiskIdWithDefaults instantiates a new VirtualDiskId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControllerBusNumber

`func (o *VirtualDiskId) GetControllerBusNumber() int64`

GetControllerBusNumber returns the ControllerBusNumber field if non-nil, zero value otherwise.

### GetControllerBusNumberOk

`func (o *VirtualDiskId) GetControllerBusNumberOk() (*int64, bool)`

GetControllerBusNumberOk returns a tuple with the ControllerBusNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerBusNumber

`func (o *VirtualDiskId) SetControllerBusNumber(v int64)`

SetControllerBusNumber sets ControllerBusNumber field to given value.

### HasControllerBusNumber

`func (o *VirtualDiskId) HasControllerBusNumber() bool`

HasControllerBusNumber returns a boolean if a field has been set.

### SetControllerBusNumberNil

`func (o *VirtualDiskId) SetControllerBusNumberNil(b bool)`

 SetControllerBusNumberNil sets the value for ControllerBusNumber to be an explicit nil

### UnsetControllerBusNumber
`func (o *VirtualDiskId) UnsetControllerBusNumber()`

UnsetControllerBusNumber ensures that no value is present for ControllerBusNumber, not even an explicit nil
### GetControllerType

`func (o *VirtualDiskId) GetControllerType() string`

GetControllerType returns the ControllerType field if non-nil, zero value otherwise.

### GetControllerTypeOk

`func (o *VirtualDiskId) GetControllerTypeOk() (*string, bool)`

GetControllerTypeOk returns a tuple with the ControllerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerType

`func (o *VirtualDiskId) SetControllerType(v string)`

SetControllerType sets ControllerType field to given value.

### HasControllerType

`func (o *VirtualDiskId) HasControllerType() bool`

HasControllerType returns a boolean if a field has been set.

### SetControllerTypeNil

`func (o *VirtualDiskId) SetControllerTypeNil(b bool)`

 SetControllerTypeNil sets the value for ControllerType to be an explicit nil

### UnsetControllerType
`func (o *VirtualDiskId) UnsetControllerType()`

UnsetControllerType ensures that no value is present for ControllerType, not even an explicit nil
### GetDiskId

`func (o *VirtualDiskId) GetDiskId() string`

GetDiskId returns the DiskId field if non-nil, zero value otherwise.

### GetDiskIdOk

`func (o *VirtualDiskId) GetDiskIdOk() (*string, bool)`

GetDiskIdOk returns a tuple with the DiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskId

`func (o *VirtualDiskId) SetDiskId(v string)`

SetDiskId sets DiskId field to given value.

### HasDiskId

`func (o *VirtualDiskId) HasDiskId() bool`

HasDiskId returns a boolean if a field has been set.

### SetDiskIdNil

`func (o *VirtualDiskId) SetDiskIdNil(b bool)`

 SetDiskIdNil sets the value for DiskId to be an explicit nil

### UnsetDiskId
`func (o *VirtualDiskId) UnsetDiskId()`

UnsetDiskId ensures that no value is present for DiskId, not even an explicit nil
### GetUnitNumber

`func (o *VirtualDiskId) GetUnitNumber() int64`

GetUnitNumber returns the UnitNumber field if non-nil, zero value otherwise.

### GetUnitNumberOk

`func (o *VirtualDiskId) GetUnitNumberOk() (*int64, bool)`

GetUnitNumberOk returns a tuple with the UnitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitNumber

`func (o *VirtualDiskId) SetUnitNumber(v int64)`

SetUnitNumber sets UnitNumber field to given value.

### HasUnitNumber

`func (o *VirtualDiskId) HasUnitNumber() bool`

HasUnitNumber returns a boolean if a field has been set.

### SetUnitNumberNil

`func (o *VirtualDiskId) SetUnitNumberNil(b bool)`

 SetUnitNumberNil sets the value for UnitNumber to be an explicit nil

### UnsetUnitNumber
`func (o *VirtualDiskId) UnsetUnitNumber()`

UnsetUnitNumber ensures that no value is present for UnitNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


