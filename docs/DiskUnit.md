# DiskUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusNumber** | Pointer to **NullableInt64** | Specifies the Id of the controller bus that controls the disk. | [optional] 
**ControllerType** | Pointer to **NullableString** | Specifies the controller type like SCSI, or IDE etc. | [optional] 
**UnitNumber** | Pointer to **NullableInt64** | Specifies the disk file name. This is the VMDK name and not the flat file name. | [optional] 

## Methods

### NewDiskUnit

`func NewDiskUnit() *DiskUnit`

NewDiskUnit instantiates a new DiskUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskUnitWithDefaults

`func NewDiskUnitWithDefaults() *DiskUnit`

NewDiskUnitWithDefaults instantiates a new DiskUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusNumber

`func (o *DiskUnit) GetBusNumber() int64`

GetBusNumber returns the BusNumber field if non-nil, zero value otherwise.

### GetBusNumberOk

`func (o *DiskUnit) GetBusNumberOk() (*int64, bool)`

GetBusNumberOk returns a tuple with the BusNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusNumber

`func (o *DiskUnit) SetBusNumber(v int64)`

SetBusNumber sets BusNumber field to given value.

### HasBusNumber

`func (o *DiskUnit) HasBusNumber() bool`

HasBusNumber returns a boolean if a field has been set.

### SetBusNumberNil

`func (o *DiskUnit) SetBusNumberNil(b bool)`

 SetBusNumberNil sets the value for BusNumber to be an explicit nil

### UnsetBusNumber
`func (o *DiskUnit) UnsetBusNumber()`

UnsetBusNumber ensures that no value is present for BusNumber, not even an explicit nil
### GetControllerType

`func (o *DiskUnit) GetControllerType() string`

GetControllerType returns the ControllerType field if non-nil, zero value otherwise.

### GetControllerTypeOk

`func (o *DiskUnit) GetControllerTypeOk() (*string, bool)`

GetControllerTypeOk returns a tuple with the ControllerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerType

`func (o *DiskUnit) SetControllerType(v string)`

SetControllerType sets ControllerType field to given value.

### HasControllerType

`func (o *DiskUnit) HasControllerType() bool`

HasControllerType returns a boolean if a field has been set.

### SetControllerTypeNil

`func (o *DiskUnit) SetControllerTypeNil(b bool)`

 SetControllerTypeNil sets the value for ControllerType to be an explicit nil

### UnsetControllerType
`func (o *DiskUnit) UnsetControllerType()`

UnsetControllerType ensures that no value is present for ControllerType, not even an explicit nil
### GetUnitNumber

`func (o *DiskUnit) GetUnitNumber() int64`

GetUnitNumber returns the UnitNumber field if non-nil, zero value otherwise.

### GetUnitNumberOk

`func (o *DiskUnit) GetUnitNumberOk() (*int64, bool)`

GetUnitNumberOk returns a tuple with the UnitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitNumber

`func (o *DiskUnit) SetUnitNumber(v int64)`

SetUnitNumber sets UnitNumber field to given value.

### HasUnitNumber

`func (o *DiskUnit) HasUnitNumber() bool`

HasUnitNumber returns a boolean if a field has been set.

### SetUnitNumberNil

`func (o *DiskUnit) SetUnitNumberNil(b bool)`

 SetUnitNumberNil sets the value for UnitNumber to be an explicit nil

### UnsetUnitNumber
`func (o *DiskUnit) UnsetUnitNumber()`

UnsetUnitNumber ensures that no value is present for UnitNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


