# HyperVDiskInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ControllerNumber** | **NullableInt64** | Specifies the disk controller number. | 
**ControllerType** | **NullableString** | Specifies the disk controller type. | 
**UnitNumber** | **NullableInt64** | Specifies the disk index number. | 

## Methods

### NewHyperVDiskInfo

`func NewHyperVDiskInfo(controllerNumber NullableInt64, controllerType NullableString, unitNumber NullableInt64, ) *HyperVDiskInfo`

NewHyperVDiskInfo instantiates a new HyperVDiskInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperVDiskInfoWithDefaults

`func NewHyperVDiskInfoWithDefaults() *HyperVDiskInfo`

NewHyperVDiskInfoWithDefaults instantiates a new HyperVDiskInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControllerNumber

`func (o *HyperVDiskInfo) GetControllerNumber() int64`

GetControllerNumber returns the ControllerNumber field if non-nil, zero value otherwise.

### GetControllerNumberOk

`func (o *HyperVDiskInfo) GetControllerNumberOk() (*int64, bool)`

GetControllerNumberOk returns a tuple with the ControllerNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerNumber

`func (o *HyperVDiskInfo) SetControllerNumber(v int64)`

SetControllerNumber sets ControllerNumber field to given value.


### SetControllerNumberNil

`func (o *HyperVDiskInfo) SetControllerNumberNil(b bool)`

 SetControllerNumberNil sets the value for ControllerNumber to be an explicit nil

### UnsetControllerNumber
`func (o *HyperVDiskInfo) UnsetControllerNumber()`

UnsetControllerNumber ensures that no value is present for ControllerNumber, not even an explicit nil
### GetControllerType

`func (o *HyperVDiskInfo) GetControllerType() string`

GetControllerType returns the ControllerType field if non-nil, zero value otherwise.

### GetControllerTypeOk

`func (o *HyperVDiskInfo) GetControllerTypeOk() (*string, bool)`

GetControllerTypeOk returns a tuple with the ControllerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerType

`func (o *HyperVDiskInfo) SetControllerType(v string)`

SetControllerType sets ControllerType field to given value.


### SetControllerTypeNil

`func (o *HyperVDiskInfo) SetControllerTypeNil(b bool)`

 SetControllerTypeNil sets the value for ControllerType to be an explicit nil

### UnsetControllerType
`func (o *HyperVDiskInfo) UnsetControllerType()`

UnsetControllerType ensures that no value is present for ControllerType, not even an explicit nil
### GetUnitNumber

`func (o *HyperVDiskInfo) GetUnitNumber() int64`

GetUnitNumber returns the UnitNumber field if non-nil, zero value otherwise.

### GetUnitNumberOk

`func (o *HyperVDiskInfo) GetUnitNumberOk() (*int64, bool)`

GetUnitNumberOk returns a tuple with the UnitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitNumber

`func (o *HyperVDiskInfo) SetUnitNumber(v int64)`

SetUnitNumber sets UnitNumber field to given value.


### SetUnitNumberNil

`func (o *HyperVDiskInfo) SetUnitNumberNil(b bool)`

 SetUnitNumberNil sets the value for UnitNumber to be an explicit nil

### UnsetUnitNumber
`func (o *HyperVDiskInfo) UnsetUnitNumber()`

UnsetUnitNumber ensures that no value is present for UnitNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


