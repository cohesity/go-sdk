# Rack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChassisIds** | Pointer to **[]int64** | List of chassis ids that are part of the rack. | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies unique id of the rack. | [optional] 
**Location** | Pointer to **NullableString** | Specifies location of the rack. | [optional] 
**Name** | Pointer to **NullableString** | Specifies name of the rack | [optional] 

## Methods

### NewRack

`func NewRack() *Rack`

NewRack instantiates a new Rack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRackWithDefaults

`func NewRackWithDefaults() *Rack`

NewRackWithDefaults instantiates a new Rack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChassisIds

`func (o *Rack) GetChassisIds() []int64`

GetChassisIds returns the ChassisIds field if non-nil, zero value otherwise.

### GetChassisIdsOk

`func (o *Rack) GetChassisIdsOk() (*[]int64, bool)`

GetChassisIdsOk returns a tuple with the ChassisIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisIds

`func (o *Rack) SetChassisIds(v []int64)`

SetChassisIds sets ChassisIds field to given value.

### HasChassisIds

`func (o *Rack) HasChassisIds() bool`

HasChassisIds returns a boolean if a field has been set.

### GetId

`func (o *Rack) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Rack) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Rack) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Rack) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Rack) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Rack) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLocation

`func (o *Rack) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Rack) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Rack) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Rack) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *Rack) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *Rack) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetName

`func (o *Rack) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Rack) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Rack) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Rack) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Rack) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Rack) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


