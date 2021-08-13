# TimeOfDay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hour** | **NullableInt32** | Specifies the hour of the day (0-23). | 
**Minute** | **NullableInt32** | Specifies the minute of the hour (0-59). | 
**TimeZone** | Pointer to **NullableString** | Specifies the time zone of the user. If not specified, default value is assumed as America/Los_Angeles. | [optional] [default to "America/Los_Angeles"]

## Methods

### NewTimeOfDay

`func NewTimeOfDay(hour NullableInt32, minute NullableInt32, ) *TimeOfDay`

NewTimeOfDay instantiates a new TimeOfDay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeOfDayWithDefaults

`func NewTimeOfDayWithDefaults() *TimeOfDay`

NewTimeOfDayWithDefaults instantiates a new TimeOfDay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHour

`func (o *TimeOfDay) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *TimeOfDay) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *TimeOfDay) SetHour(v int32)`

SetHour sets Hour field to given value.


### SetHourNil

`func (o *TimeOfDay) SetHourNil(b bool)`

 SetHourNil sets the value for Hour to be an explicit nil

### UnsetHour
`func (o *TimeOfDay) UnsetHour()`

UnsetHour ensures that no value is present for Hour, not even an explicit nil
### GetMinute

`func (o *TimeOfDay) GetMinute() int32`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *TimeOfDay) GetMinuteOk() (*int32, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *TimeOfDay) SetMinute(v int32)`

SetMinute sets Minute field to given value.


### SetMinuteNil

`func (o *TimeOfDay) SetMinuteNil(b bool)`

 SetMinuteNil sets the value for Minute to be an explicit nil

### UnsetMinute
`func (o *TimeOfDay) UnsetMinute()`

UnsetMinute ensures that no value is present for Minute, not even an explicit nil
### GetTimeZone

`func (o *TimeOfDay) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *TimeOfDay) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *TimeOfDay) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *TimeOfDay) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### SetTimeZoneNil

`func (o *TimeOfDay) SetTimeZoneNil(b bool)`

 SetTimeZoneNil sets the value for TimeZone to be an explicit nil

### UnsetTimeZone
`func (o *TimeOfDay) UnsetTimeZone()`

UnsetTimeZone ensures that no value is present for TimeZone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


