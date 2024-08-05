# HeliosIncrementalSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DaySchedule** | Pointer to [**HeliosDaySchedule**](HeliosDaySchedule.md) |  | [optional] 
**HourSchedule** | Pointer to [**HeliosHourSchedule**](HeliosHourSchedule.md) |  | [optional] 
**MinuteSchedule** | Pointer to [**HeliosMinuteSchedule**](HeliosMinuteSchedule.md) |  | [optional] 
**MonthSchedule** | Pointer to [**HeliosMonthSchedule**](HeliosMonthSchedule.md) |  | [optional] 
**Unit** | **NullableString** | Specifies how often to start new runs of a Protection Group. &lt;br&gt;&#39;Minutes&#39; specifies that Protection Group run starts periodically after certain number of minutes specified in &#39;frequency&#39; field. &lt;br&gt;&#39;Hours&#39; specifies that Protection Group run starts periodically after certain number of hours specified in &#39;frequency&#39; field. &lt;br&gt;&#39;Days&#39; specifies that Protection Group run starts periodically after certain number of days specified in &#39;frequency&#39; field. &lt;br&gt;&#39;Week&#39; specifies that new Protection Group runs start weekly on certain days specified using &#39;dayOfWeek&#39; field. &lt;br&gt;&#39;Month&#39; specifies that new Protection Group runs start monthly on certain day of specific week. This schedule needs &#39;weekOfMonth&#39; and &#39;dayOfWeek&#39; fields to be set. &lt;br&gt; Example: To run the Protection Group on Second Sunday of Every Month, following schedule need to be set: &lt;br&gt; unit: &#39;Month&#39; &lt;br&gt; dayOfWeek: &#39;Sunday&#39; &lt;br&gt; weekOfMonth: &#39;Second&#39; | 
**WeekSchedule** | Pointer to [**HeliosWeekSchedule**](HeliosWeekSchedule.md) |  | [optional] 

## Methods

### NewHeliosIncrementalSchedule

`func NewHeliosIncrementalSchedule(unit NullableString, ) *HeliosIncrementalSchedule`

NewHeliosIncrementalSchedule instantiates a new HeliosIncrementalSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosIncrementalScheduleWithDefaults

`func NewHeliosIncrementalScheduleWithDefaults() *HeliosIncrementalSchedule`

NewHeliosIncrementalScheduleWithDefaults instantiates a new HeliosIncrementalSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDaySchedule

`func (o *HeliosIncrementalSchedule) GetDaySchedule() HeliosDaySchedule`

GetDaySchedule returns the DaySchedule field if non-nil, zero value otherwise.

### GetDayScheduleOk

`func (o *HeliosIncrementalSchedule) GetDayScheduleOk() (*HeliosDaySchedule, bool)`

GetDayScheduleOk returns a tuple with the DaySchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaySchedule

`func (o *HeliosIncrementalSchedule) SetDaySchedule(v HeliosDaySchedule)`

SetDaySchedule sets DaySchedule field to given value.

### HasDaySchedule

`func (o *HeliosIncrementalSchedule) HasDaySchedule() bool`

HasDaySchedule returns a boolean if a field has been set.

### GetHourSchedule

`func (o *HeliosIncrementalSchedule) GetHourSchedule() HeliosHourSchedule`

GetHourSchedule returns the HourSchedule field if non-nil, zero value otherwise.

### GetHourScheduleOk

`func (o *HeliosIncrementalSchedule) GetHourScheduleOk() (*HeliosHourSchedule, bool)`

GetHourScheduleOk returns a tuple with the HourSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourSchedule

`func (o *HeliosIncrementalSchedule) SetHourSchedule(v HeliosHourSchedule)`

SetHourSchedule sets HourSchedule field to given value.

### HasHourSchedule

`func (o *HeliosIncrementalSchedule) HasHourSchedule() bool`

HasHourSchedule returns a boolean if a field has been set.

### GetMinuteSchedule

`func (o *HeliosIncrementalSchedule) GetMinuteSchedule() HeliosMinuteSchedule`

GetMinuteSchedule returns the MinuteSchedule field if non-nil, zero value otherwise.

### GetMinuteScheduleOk

`func (o *HeliosIncrementalSchedule) GetMinuteScheduleOk() (*HeliosMinuteSchedule, bool)`

GetMinuteScheduleOk returns a tuple with the MinuteSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinuteSchedule

`func (o *HeliosIncrementalSchedule) SetMinuteSchedule(v HeliosMinuteSchedule)`

SetMinuteSchedule sets MinuteSchedule field to given value.

### HasMinuteSchedule

`func (o *HeliosIncrementalSchedule) HasMinuteSchedule() bool`

HasMinuteSchedule returns a boolean if a field has been set.

### GetMonthSchedule

`func (o *HeliosIncrementalSchedule) GetMonthSchedule() HeliosMonthSchedule`

GetMonthSchedule returns the MonthSchedule field if non-nil, zero value otherwise.

### GetMonthScheduleOk

`func (o *HeliosIncrementalSchedule) GetMonthScheduleOk() (*HeliosMonthSchedule, bool)`

GetMonthScheduleOk returns a tuple with the MonthSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthSchedule

`func (o *HeliosIncrementalSchedule) SetMonthSchedule(v HeliosMonthSchedule)`

SetMonthSchedule sets MonthSchedule field to given value.

### HasMonthSchedule

`func (o *HeliosIncrementalSchedule) HasMonthSchedule() bool`

HasMonthSchedule returns a boolean if a field has been set.

### GetUnit

`func (o *HeliosIncrementalSchedule) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *HeliosIncrementalSchedule) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *HeliosIncrementalSchedule) SetUnit(v string)`

SetUnit sets Unit field to given value.


### SetUnitNil

`func (o *HeliosIncrementalSchedule) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *HeliosIncrementalSchedule) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetWeekSchedule

`func (o *HeliosIncrementalSchedule) GetWeekSchedule() HeliosWeekSchedule`

GetWeekSchedule returns the WeekSchedule field if non-nil, zero value otherwise.

### GetWeekScheduleOk

`func (o *HeliosIncrementalSchedule) GetWeekScheduleOk() (*HeliosWeekSchedule, bool)`

GetWeekScheduleOk returns a tuple with the WeekSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekSchedule

`func (o *HeliosIncrementalSchedule) SetWeekSchedule(v HeliosWeekSchedule)`

SetWeekSchedule sets WeekSchedule field to given value.

### HasWeekSchedule

`func (o *HeliosIncrementalSchedule) HasWeekSchedule() bool`

HasWeekSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


