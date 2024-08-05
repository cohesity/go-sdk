# IncrementalSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DaySchedule** | Pointer to [**DaySchedule**](DaySchedule.md) |  | [optional] 
**HourSchedule** | Pointer to [**HourSchedule**](HourSchedule.md) |  | [optional] 
**MinuteSchedule** | Pointer to [**MinuteSchedule**](MinuteSchedule.md) |  | [optional] 
**MonthSchedule** | Pointer to [**MonthSchedule**](MonthSchedule.md) |  | [optional] 
**Unit** | **NullableString** | Specifies how often to start new runs of a Protection Group. &lt;br&gt;&#39;Minutes&#39; specifies that Protection Group run starts periodically after certain number of minutes specified in &#39;frequency&#39; field. &lt;br&gt;&#39;Hours&#39; specifies that Protection Group run starts periodically after certain number of hours specified in &#39;frequency&#39; field. &lt;br&gt;&#39;Days&#39; specifies that Protection Group run starts periodically after certain number of days specified in &#39;frequency&#39; field. &lt;br&gt;&#39;Week&#39; specifies that new Protection Group runs start weekly on certain days specified using &#39;dayOfWeek&#39; field. &lt;br&gt;&#39;Month&#39; specifies that new Protection Group runs start monthly on certain day of specific week. This schedule needs &#39;weekOfMonth&#39; and &#39;dayOfWeek&#39; fields to be set. &lt;br&gt; Example: To run the Protection Group on Second Sunday of Every Month, following schedule need to be set: &lt;br&gt; unit: &#39;Month&#39; &lt;br&gt; dayOfWeek: &#39;Sunday&#39; &lt;br&gt; weekOfMonth: &#39;Second&#39; | 
**WeekSchedule** | Pointer to [**WeekSchedule**](WeekSchedule.md) |  | [optional] 
**YearSchedule** | Pointer to [**YearSchedule**](YearSchedule.md) |  | [optional] 

## Methods

### NewIncrementalSchedule

`func NewIncrementalSchedule(unit NullableString, ) *IncrementalSchedule`

NewIncrementalSchedule instantiates a new IncrementalSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncrementalScheduleWithDefaults

`func NewIncrementalScheduleWithDefaults() *IncrementalSchedule`

NewIncrementalScheduleWithDefaults instantiates a new IncrementalSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDaySchedule

`func (o *IncrementalSchedule) GetDaySchedule() DaySchedule`

GetDaySchedule returns the DaySchedule field if non-nil, zero value otherwise.

### GetDayScheduleOk

`func (o *IncrementalSchedule) GetDayScheduleOk() (*DaySchedule, bool)`

GetDayScheduleOk returns a tuple with the DaySchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaySchedule

`func (o *IncrementalSchedule) SetDaySchedule(v DaySchedule)`

SetDaySchedule sets DaySchedule field to given value.

### HasDaySchedule

`func (o *IncrementalSchedule) HasDaySchedule() bool`

HasDaySchedule returns a boolean if a field has been set.

### GetHourSchedule

`func (o *IncrementalSchedule) GetHourSchedule() HourSchedule`

GetHourSchedule returns the HourSchedule field if non-nil, zero value otherwise.

### GetHourScheduleOk

`func (o *IncrementalSchedule) GetHourScheduleOk() (*HourSchedule, bool)`

GetHourScheduleOk returns a tuple with the HourSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourSchedule

`func (o *IncrementalSchedule) SetHourSchedule(v HourSchedule)`

SetHourSchedule sets HourSchedule field to given value.

### HasHourSchedule

`func (o *IncrementalSchedule) HasHourSchedule() bool`

HasHourSchedule returns a boolean if a field has been set.

### GetMinuteSchedule

`func (o *IncrementalSchedule) GetMinuteSchedule() MinuteSchedule`

GetMinuteSchedule returns the MinuteSchedule field if non-nil, zero value otherwise.

### GetMinuteScheduleOk

`func (o *IncrementalSchedule) GetMinuteScheduleOk() (*MinuteSchedule, bool)`

GetMinuteScheduleOk returns a tuple with the MinuteSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinuteSchedule

`func (o *IncrementalSchedule) SetMinuteSchedule(v MinuteSchedule)`

SetMinuteSchedule sets MinuteSchedule field to given value.

### HasMinuteSchedule

`func (o *IncrementalSchedule) HasMinuteSchedule() bool`

HasMinuteSchedule returns a boolean if a field has been set.

### GetMonthSchedule

`func (o *IncrementalSchedule) GetMonthSchedule() MonthSchedule`

GetMonthSchedule returns the MonthSchedule field if non-nil, zero value otherwise.

### GetMonthScheduleOk

`func (o *IncrementalSchedule) GetMonthScheduleOk() (*MonthSchedule, bool)`

GetMonthScheduleOk returns a tuple with the MonthSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthSchedule

`func (o *IncrementalSchedule) SetMonthSchedule(v MonthSchedule)`

SetMonthSchedule sets MonthSchedule field to given value.

### HasMonthSchedule

`func (o *IncrementalSchedule) HasMonthSchedule() bool`

HasMonthSchedule returns a boolean if a field has been set.

### GetUnit

`func (o *IncrementalSchedule) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *IncrementalSchedule) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *IncrementalSchedule) SetUnit(v string)`

SetUnit sets Unit field to given value.


### SetUnitNil

`func (o *IncrementalSchedule) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *IncrementalSchedule) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetWeekSchedule

`func (o *IncrementalSchedule) GetWeekSchedule() WeekSchedule`

GetWeekSchedule returns the WeekSchedule field if non-nil, zero value otherwise.

### GetWeekScheduleOk

`func (o *IncrementalSchedule) GetWeekScheduleOk() (*WeekSchedule, bool)`

GetWeekScheduleOk returns a tuple with the WeekSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekSchedule

`func (o *IncrementalSchedule) SetWeekSchedule(v WeekSchedule)`

SetWeekSchedule sets WeekSchedule field to given value.

### HasWeekSchedule

`func (o *IncrementalSchedule) HasWeekSchedule() bool`

HasWeekSchedule returns a boolean if a field has been set.

### GetYearSchedule

`func (o *IncrementalSchedule) GetYearSchedule() YearSchedule`

GetYearSchedule returns the YearSchedule field if non-nil, zero value otherwise.

### GetYearScheduleOk

`func (o *IncrementalSchedule) GetYearScheduleOk() (*YearSchedule, bool)`

GetYearScheduleOk returns a tuple with the YearSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearSchedule

`func (o *IncrementalSchedule) SetYearSchedule(v YearSchedule)`

SetYearSchedule sets YearSchedule field to given value.

### HasYearSchedule

`func (o *IncrementalSchedule) HasYearSchedule() bool`

HasYearSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


