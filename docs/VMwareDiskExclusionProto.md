# VMwareDiskExclusionProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ControllerBusNumber** | Pointer to **NullableInt64** | Controller&#39;s bus-id controlling the virtual disk in question. | [optional] 
**ControllerType** | Pointer to **NullableString** | Controller&#39;s type (SCSI, IDE etc). | [optional] 
**UnitNumber** | Pointer to **NullableInt64** | Disk unit number to identify the virtual disk within a controller. | [optional] 

## Methods

### NewVMwareDiskExclusionProto

`func NewVMwareDiskExclusionProto() *VMwareDiskExclusionProto`

NewVMwareDiskExclusionProto instantiates a new VMwareDiskExclusionProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMwareDiskExclusionProtoWithDefaults

`func NewVMwareDiskExclusionProtoWithDefaults() *VMwareDiskExclusionProto`

NewVMwareDiskExclusionProtoWithDefaults instantiates a new VMwareDiskExclusionProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControllerBusNumber

`func (o *VMwareDiskExclusionProto) GetControllerBusNumber() int64`

GetControllerBusNumber returns the ControllerBusNumber field if non-nil, zero value otherwise.

### GetControllerBusNumberOk

`func (o *VMwareDiskExclusionProto) GetControllerBusNumberOk() (*int64, bool)`

GetControllerBusNumberOk returns a tuple with the ControllerBusNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerBusNumber

`func (o *VMwareDiskExclusionProto) SetControllerBusNumber(v int64)`

SetControllerBusNumber sets ControllerBusNumber field to given value.

### HasControllerBusNumber

`func (o *VMwareDiskExclusionProto) HasControllerBusNumber() bool`

HasControllerBusNumber returns a boolean if a field has been set.

### SetControllerBusNumberNil

`func (o *VMwareDiskExclusionProto) SetControllerBusNumberNil(b bool)`

 SetControllerBusNumberNil sets the value for ControllerBusNumber to be an explicit nil

### UnsetControllerBusNumber
`func (o *VMwareDiskExclusionProto) UnsetControllerBusNumber()`

UnsetControllerBusNumber ensures that no value is present for ControllerBusNumber, not even an explicit nil
### GetControllerType

`func (o *VMwareDiskExclusionProto) GetControllerType() string`

GetControllerType returns the ControllerType field if non-nil, zero value otherwise.

### GetControllerTypeOk

`func (o *VMwareDiskExclusionProto) GetControllerTypeOk() (*string, bool)`

GetControllerTypeOk returns a tuple with the ControllerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerType

`func (o *VMwareDiskExclusionProto) SetControllerType(v string)`

SetControllerType sets ControllerType field to given value.

### HasControllerType

`func (o *VMwareDiskExclusionProto) HasControllerType() bool`

HasControllerType returns a boolean if a field has been set.

### SetControllerTypeNil

`func (o *VMwareDiskExclusionProto) SetControllerTypeNil(b bool)`

 SetControllerTypeNil sets the value for ControllerType to be an explicit nil

### UnsetControllerType
`func (o *VMwareDiskExclusionProto) UnsetControllerType()`

UnsetControllerType ensures that no value is present for ControllerType, not even an explicit nil
### GetUnitNumber

`func (o *VMwareDiskExclusionProto) GetUnitNumber() int64`

GetUnitNumber returns the UnitNumber field if non-nil, zero value otherwise.

### GetUnitNumberOk

`func (o *VMwareDiskExclusionProto) GetUnitNumberOk() (*int64, bool)`

GetUnitNumberOk returns a tuple with the UnitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitNumber

`func (o *VMwareDiskExclusionProto) SetUnitNumber(v int64)`

SetUnitNumber sets UnitNumber field to given value.

### HasUnitNumber

`func (o *VMwareDiskExclusionProto) HasUnitNumber() bool`

HasUnitNumber returns a boolean if a field has been set.

### SetUnitNumberNil

`func (o *VMwareDiskExclusionProto) SetUnitNumberNil(b bool)`

 SetUnitNumberNil sets the value for UnitNumber to be an explicit nil

### UnsetUnitNumber
`func (o *VMwareDiskExclusionProto) UnsetUnitNumber()`

UnsetUnitNumber ensures that no value is present for UnitNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


