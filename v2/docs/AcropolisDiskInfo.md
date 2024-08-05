# AcropolisDiskInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ControllerType** | **NullableString** | Specifies the disk controller type. | 
**UnitNumber** | **NullableInt32** | Specifies the disk index number. | 

## Methods

### NewAcropolisDiskInfo

`func NewAcropolisDiskInfo(controllerType NullableString, unitNumber NullableInt32, ) *AcropolisDiskInfo`

NewAcropolisDiskInfo instantiates a new AcropolisDiskInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcropolisDiskInfoWithDefaults

`func NewAcropolisDiskInfoWithDefaults() *AcropolisDiskInfo`

NewAcropolisDiskInfoWithDefaults instantiates a new AcropolisDiskInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControllerType

`func (o *AcropolisDiskInfo) GetControllerType() string`

GetControllerType returns the ControllerType field if non-nil, zero value otherwise.

### GetControllerTypeOk

`func (o *AcropolisDiskInfo) GetControllerTypeOk() (*string, bool)`

GetControllerTypeOk returns a tuple with the ControllerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerType

`func (o *AcropolisDiskInfo) SetControllerType(v string)`

SetControllerType sets ControllerType field to given value.


### SetControllerTypeNil

`func (o *AcropolisDiskInfo) SetControllerTypeNil(b bool)`

 SetControllerTypeNil sets the value for ControllerType to be an explicit nil

### UnsetControllerType
`func (o *AcropolisDiskInfo) UnsetControllerType()`

UnsetControllerType ensures that no value is present for ControllerType, not even an explicit nil
### GetUnitNumber

`func (o *AcropolisDiskInfo) GetUnitNumber() int32`

GetUnitNumber returns the UnitNumber field if non-nil, zero value otherwise.

### GetUnitNumberOk

`func (o *AcropolisDiskInfo) GetUnitNumberOk() (*int32, bool)`

GetUnitNumberOk returns a tuple with the UnitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitNumber

`func (o *AcropolisDiskInfo) SetUnitNumber(v int32)`

SetUnitNumber sets UnitNumber field to given value.


### SetUnitNumberNil

`func (o *AcropolisDiskInfo) SetUnitNumberNil(b bool)`

 SetUnitNumberNil sets the value for UnitNumber to be an explicit nil

### UnsetUnitNumber
`func (o *AcropolisDiskInfo) UnsetUnitNumber()`

UnsetUnitNumber ensures that no value is present for UnitNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


