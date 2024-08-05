# TimeWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfTheWeek** | Pointer to **NullableString** | Specifies the week day. | [optional] 
**EndTime** | Pointer to [**Time**](Time.md) |  | [optional] 
**StartTime** | Pointer to [**Time**](Time.md) |  | [optional] 

## Methods

### NewTimeWindow

`func NewTimeWindow() *TimeWindow`

NewTimeWindow instantiates a new TimeWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeWindowWithDefaults

`func NewTimeWindowWithDefaults() *TimeWindow`

NewTimeWindowWithDefaults instantiates a new TimeWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfTheWeek

`func (o *TimeWindow) GetDayOfTheWeek() string`

GetDayOfTheWeek returns the DayOfTheWeek field if non-nil, zero value otherwise.

### GetDayOfTheWeekOk

`func (o *TimeWindow) GetDayOfTheWeekOk() (*string, bool)`

GetDayOfTheWeekOk returns a tuple with the DayOfTheWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfTheWeek

`func (o *TimeWindow) SetDayOfTheWeek(v string)`

SetDayOfTheWeek sets DayOfTheWeek field to given value.

### HasDayOfTheWeek

`func (o *TimeWindow) HasDayOfTheWeek() bool`

HasDayOfTheWeek returns a boolean if a field has been set.

### SetDayOfTheWeekNil

`func (o *TimeWindow) SetDayOfTheWeekNil(b bool)`

 SetDayOfTheWeekNil sets the value for DayOfTheWeek to be an explicit nil

### UnsetDayOfTheWeek
`func (o *TimeWindow) UnsetDayOfTheWeek()`

UnsetDayOfTheWeek ensures that no value is present for DayOfTheWeek, not even an explicit nil
### GetEndTime

`func (o *TimeWindow) GetEndTime() Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TimeWindow) GetEndTimeOk() (*Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TimeWindow) SetEndTime(v Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TimeWindow) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetStartTime

`func (o *TimeWindow) GetStartTime() Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TimeWindow) GetStartTimeOk() (*Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TimeWindow) SetStartTime(v Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TimeWindow) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


