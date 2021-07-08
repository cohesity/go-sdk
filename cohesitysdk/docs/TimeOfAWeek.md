# TimeOfAWeek

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | Pointer to **[]string** | Array of Week Days.  Specifies a list of days of a week (such as &#39;kSunday&#39;) when the time period should be applied. If not set, the time range applies to all days of the week. Specifies a day in a week such as &#39;kSunday&#39;, &#39;kMonday&#39;, etc. | [optional] 
**EndTime** | Pointer to [**NullableTimeOfDay**](TimeOfDay.md) | Specifies the end time for the daily time period. | [optional] 
**IsAllDay** | Pointer to **NullableBool** | All Day.  Specifies that time range is applied for entire day. | [optional] 
**StartTime** | Pointer to [**NullableTimeOfDay**](TimeOfDay.md) | Specifies the start time for the daily time period. | [optional] 

## Methods

### NewTimeOfAWeek

`func NewTimeOfAWeek() *TimeOfAWeek`

NewTimeOfAWeek instantiates a new TimeOfAWeek object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeOfAWeekWithDefaults

`func NewTimeOfAWeekWithDefaults() *TimeOfAWeek`

NewTimeOfAWeekWithDefaults instantiates a new TimeOfAWeek object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *TimeOfAWeek) GetDays() []string`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *TimeOfAWeek) GetDaysOk() (*[]string, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *TimeOfAWeek) SetDays(v []string)`

SetDays sets Days field to given value.

### HasDays

`func (o *TimeOfAWeek) HasDays() bool`

HasDays returns a boolean if a field has been set.

### SetDaysNil

`func (o *TimeOfAWeek) SetDaysNil(b bool)`

 SetDaysNil sets the value for Days to be an explicit nil

### UnsetDays
`func (o *TimeOfAWeek) UnsetDays()`

UnsetDays ensures that no value is present for Days, not even an explicit nil
### GetEndTime

`func (o *TimeOfAWeek) GetEndTime() TimeOfDay`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TimeOfAWeek) GetEndTimeOk() (*TimeOfDay, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TimeOfAWeek) SetEndTime(v TimeOfDay)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TimeOfAWeek) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *TimeOfAWeek) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *TimeOfAWeek) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetIsAllDay

`func (o *TimeOfAWeek) GetIsAllDay() bool`

GetIsAllDay returns the IsAllDay field if non-nil, zero value otherwise.

### GetIsAllDayOk

`func (o *TimeOfAWeek) GetIsAllDayOk() (*bool, bool)`

GetIsAllDayOk returns a tuple with the IsAllDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllDay

`func (o *TimeOfAWeek) SetIsAllDay(v bool)`

SetIsAllDay sets IsAllDay field to given value.

### HasIsAllDay

`func (o *TimeOfAWeek) HasIsAllDay() bool`

HasIsAllDay returns a boolean if a field has been set.

### SetIsAllDayNil

`func (o *TimeOfAWeek) SetIsAllDayNil(b bool)`

 SetIsAllDayNil sets the value for IsAllDay to be an explicit nil

### UnsetIsAllDay
`func (o *TimeOfAWeek) UnsetIsAllDay()`

UnsetIsAllDay ensures that no value is present for IsAllDay, not even an explicit nil
### GetStartTime

`func (o *TimeOfAWeek) GetStartTime() TimeOfDay`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TimeOfAWeek) GetStartTimeOk() (*TimeOfDay, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TimeOfAWeek) SetStartTime(v TimeOfDay)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TimeOfAWeek) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *TimeOfAWeek) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *TimeOfAWeek) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


