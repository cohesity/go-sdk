# BmrSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | **NullableString** | Specifies how often to start new runs of a Protection Group. &lt;br&gt;&#39;Weeks&#39; specifies that new Protection Group runs start weekly on certain days specified using &#39;dayOfWeek&#39; field. &lt;br&gt;&#39;Months&#39; specifies that new Protection Group runs start monthly on certain day of specific week. | 
**DaySchedule** | Pointer to [**DaySchedule**](DaySchedule.md) |  | [optional] 
**WeekSchedule** | Pointer to [**WeekSchedule**](WeekSchedule.md) |  | [optional] 
**MonthSchedule** | Pointer to [**MonthSchedule**](MonthSchedule.md) |  | [optional] 

## Methods

### NewBmrSchedule

`func NewBmrSchedule(unit NullableString, ) *BmrSchedule`

NewBmrSchedule instantiates a new BmrSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBmrScheduleWithDefaults

`func NewBmrScheduleWithDefaults() *BmrSchedule`

NewBmrScheduleWithDefaults instantiates a new BmrSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *BmrSchedule) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *BmrSchedule) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *BmrSchedule) SetUnit(v string)`

SetUnit sets Unit field to given value.


### SetUnitNil

`func (o *BmrSchedule) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *BmrSchedule) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetDaySchedule

`func (o *BmrSchedule) GetDaySchedule() DaySchedule`

GetDaySchedule returns the DaySchedule field if non-nil, zero value otherwise.

### GetDayScheduleOk

`func (o *BmrSchedule) GetDayScheduleOk() (*DaySchedule, bool)`

GetDayScheduleOk returns a tuple with the DaySchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaySchedule

`func (o *BmrSchedule) SetDaySchedule(v DaySchedule)`

SetDaySchedule sets DaySchedule field to given value.

### HasDaySchedule

`func (o *BmrSchedule) HasDaySchedule() bool`

HasDaySchedule returns a boolean if a field has been set.

### GetWeekSchedule

`func (o *BmrSchedule) GetWeekSchedule() WeekSchedule`

GetWeekSchedule returns the WeekSchedule field if non-nil, zero value otherwise.

### GetWeekScheduleOk

`func (o *BmrSchedule) GetWeekScheduleOk() (*WeekSchedule, bool)`

GetWeekScheduleOk returns a tuple with the WeekSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekSchedule

`func (o *BmrSchedule) SetWeekSchedule(v WeekSchedule)`

SetWeekSchedule sets WeekSchedule field to given value.

### HasWeekSchedule

`func (o *BmrSchedule) HasWeekSchedule() bool`

HasWeekSchedule returns a boolean if a field has been set.

### GetMonthSchedule

`func (o *BmrSchedule) GetMonthSchedule() MonthSchedule`

GetMonthSchedule returns the MonthSchedule field if non-nil, zero value otherwise.

### GetMonthScheduleOk

`func (o *BmrSchedule) GetMonthScheduleOk() (*MonthSchedule, bool)`

GetMonthScheduleOk returns a tuple with the MonthSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthSchedule

`func (o *BmrSchedule) SetMonthSchedule(v MonthSchedule)`

SetMonthSchedule sets MonthSchedule field to given value.

### HasMonthSchedule

`func (o *BmrSchedule) HasMonthSchedule() bool`

HasMonthSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


