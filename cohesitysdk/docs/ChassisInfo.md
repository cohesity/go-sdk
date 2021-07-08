# ChassisInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChassisId** | Pointer to **NullableInt64** | ChassisId is a unique id assigned to the chassis. | [optional] 
**ChassisName** | Pointer to **NullableString** | ChassisName is the name of the chassis. This could be the chassis serial number by default. | [optional] 
**ChassisSerial** | Pointer to **NullableString** | Chassis serial. | [optional] 
**Location** | Pointer to **NullableString** | Location is the location of the chassis within the rack. | [optional] 
**RackId** | Pointer to **NullableInt64** | Rack is the rack within which this chassis lives. | [optional] 

## Methods

### NewChassisInfo

`func NewChassisInfo() *ChassisInfo`

NewChassisInfo instantiates a new ChassisInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChassisInfoWithDefaults

`func NewChassisInfoWithDefaults() *ChassisInfo`

NewChassisInfoWithDefaults instantiates a new ChassisInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChassisId

`func (o *ChassisInfo) GetChassisId() int64`

GetChassisId returns the ChassisId field if non-nil, zero value otherwise.

### GetChassisIdOk

`func (o *ChassisInfo) GetChassisIdOk() (*int64, bool)`

GetChassisIdOk returns a tuple with the ChassisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisId

`func (o *ChassisInfo) SetChassisId(v int64)`

SetChassisId sets ChassisId field to given value.

### HasChassisId

`func (o *ChassisInfo) HasChassisId() bool`

HasChassisId returns a boolean if a field has been set.

### SetChassisIdNil

`func (o *ChassisInfo) SetChassisIdNil(b bool)`

 SetChassisIdNil sets the value for ChassisId to be an explicit nil

### UnsetChassisId
`func (o *ChassisInfo) UnsetChassisId()`

UnsetChassisId ensures that no value is present for ChassisId, not even an explicit nil
### GetChassisName

`func (o *ChassisInfo) GetChassisName() string`

GetChassisName returns the ChassisName field if non-nil, zero value otherwise.

### GetChassisNameOk

`func (o *ChassisInfo) GetChassisNameOk() (*string, bool)`

GetChassisNameOk returns a tuple with the ChassisName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisName

`func (o *ChassisInfo) SetChassisName(v string)`

SetChassisName sets ChassisName field to given value.

### HasChassisName

`func (o *ChassisInfo) HasChassisName() bool`

HasChassisName returns a boolean if a field has been set.

### SetChassisNameNil

`func (o *ChassisInfo) SetChassisNameNil(b bool)`

 SetChassisNameNil sets the value for ChassisName to be an explicit nil

### UnsetChassisName
`func (o *ChassisInfo) UnsetChassisName()`

UnsetChassisName ensures that no value is present for ChassisName, not even an explicit nil
### GetChassisSerial

`func (o *ChassisInfo) GetChassisSerial() string`

GetChassisSerial returns the ChassisSerial field if non-nil, zero value otherwise.

### GetChassisSerialOk

`func (o *ChassisInfo) GetChassisSerialOk() (*string, bool)`

GetChassisSerialOk returns a tuple with the ChassisSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisSerial

`func (o *ChassisInfo) SetChassisSerial(v string)`

SetChassisSerial sets ChassisSerial field to given value.

### HasChassisSerial

`func (o *ChassisInfo) HasChassisSerial() bool`

HasChassisSerial returns a boolean if a field has been set.

### SetChassisSerialNil

`func (o *ChassisInfo) SetChassisSerialNil(b bool)`

 SetChassisSerialNil sets the value for ChassisSerial to be an explicit nil

### UnsetChassisSerial
`func (o *ChassisInfo) UnsetChassisSerial()`

UnsetChassisSerial ensures that no value is present for ChassisSerial, not even an explicit nil
### GetLocation

`func (o *ChassisInfo) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ChassisInfo) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ChassisInfo) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ChassisInfo) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *ChassisInfo) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *ChassisInfo) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetRackId

`func (o *ChassisInfo) GetRackId() int64`

GetRackId returns the RackId field if non-nil, zero value otherwise.

### GetRackIdOk

`func (o *ChassisInfo) GetRackIdOk() (*int64, bool)`

GetRackIdOk returns a tuple with the RackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackId

`func (o *ChassisInfo) SetRackId(v int64)`

SetRackId sets RackId field to given value.

### HasRackId

`func (o *ChassisInfo) HasRackId() bool`

HasRackId returns a boolean if a field has been set.

### SetRackIdNil

`func (o *ChassisInfo) SetRackIdNil(b bool)`

 SetRackIdNil sets the value for RackId to be an explicit nil

### UnsetRackId
`func (o *ChassisInfo) UnsetRackId()`

UnsetRackId ensures that no value is present for RackId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


