# Time

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hour** | Pointer to **NullableInt32** | The hour when backup should be performed (0 - 23). | [optional] 
**Minute** | Pointer to **NullableInt32** | The minute when backup should be performed (0 - 59). | [optional] 

## Methods

### NewTime

`func NewTime() *Time`

NewTime instantiates a new Time object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeWithDefaults

`func NewTimeWithDefaults() *Time`

NewTimeWithDefaults instantiates a new Time object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHour

`func (o *Time) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *Time) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *Time) SetHour(v int32)`

SetHour sets Hour field to given value.

### HasHour

`func (o *Time) HasHour() bool`

HasHour returns a boolean if a field has been set.

### SetHourNil

`func (o *Time) SetHourNil(b bool)`

 SetHourNil sets the value for Hour to be an explicit nil

### UnsetHour
`func (o *Time) UnsetHour()`

UnsetHour ensures that no value is present for Hour, not even an explicit nil
### GetMinute

`func (o *Time) GetMinute() int32`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *Time) GetMinuteOk() (*int32, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *Time) SetMinute(v int32)`

SetMinute sets Minute field to given value.

### HasMinute

`func (o *Time) HasMinute() bool`

HasMinute returns a boolean if a field has been set.

### SetMinuteNil

`func (o *Time) SetMinuteNil(b bool)`

 SetMinuteNil sets the value for Minute to be an explicit nil

### UnsetMinute
`func (o *Time) UnsetMinute()`

UnsetMinute ensures that no value is present for Minute, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


