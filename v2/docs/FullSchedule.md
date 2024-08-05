# FullSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DaySchedule** | Pointer to [**DaySchedule**](DaySchedule.md) |  | [optional] 
**MonthSchedule** | Pointer to [**MonthSchedule**](MonthSchedule.md) |  | [optional] 
**Unit** | **NullableString** | Specifies how often to start new runs of a Protection Group. &lt;br&gt;&#39;Days&#39; specifies that Protection Group run starts periodically on every day. For full backup schedule, currently we only support frequecny of 1 which indicates that full backup will be performed daily. &lt;br&gt;&#39;Weeks&#39; specifies that new Protection Group runs start weekly on certain days specified using &#39;dayOfWeek&#39; field. &lt;br&gt;&#39;Months&#39; specifies that new Protection Group runs start monthly on certain day of specific week. This schedule needs &#39;weekOfMonth&#39; and &#39;dayOfWeek&#39; fields to be set. &lt;br&gt;&#39;ProtectOnce&#39; specifies that groups using this policy option will run only once and after that group will permanently be disabled. &lt;br&gt; Example: To run the Protection Group on Second Sunday of Every Month, following schedule need to be set: &lt;br&gt; unit: &#39;Month&#39; &lt;br&gt; dayOfWeek: &#39;Sunday&#39; &lt;br&gt; weekOfMonth: &#39;Second&#39; | 
**WeekSchedule** | Pointer to [**WeekSchedule**](WeekSchedule.md) |  | [optional] 
**YearSchedule** | Pointer to [**YearSchedule**](YearSchedule.md) |  | [optional] 

## Methods

### NewFullSchedule

`func NewFullSchedule(unit NullableString, ) *FullSchedule`

NewFullSchedule instantiates a new FullSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullScheduleWithDefaults

`func NewFullScheduleWithDefaults() *FullSchedule`

NewFullScheduleWithDefaults instantiates a new FullSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDaySchedule

`func (o *FullSchedule) GetDaySchedule() DaySchedule`

GetDaySchedule returns the DaySchedule field if non-nil, zero value otherwise.

### GetDayScheduleOk

`func (o *FullSchedule) GetDayScheduleOk() (*DaySchedule, bool)`

GetDayScheduleOk returns a tuple with the DaySchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaySchedule

`func (o *FullSchedule) SetDaySchedule(v DaySchedule)`

SetDaySchedule sets DaySchedule field to given value.

### HasDaySchedule

`func (o *FullSchedule) HasDaySchedule() bool`

HasDaySchedule returns a boolean if a field has been set.

### GetMonthSchedule

`func (o *FullSchedule) GetMonthSchedule() MonthSchedule`

GetMonthSchedule returns the MonthSchedule field if non-nil, zero value otherwise.

### GetMonthScheduleOk

`func (o *FullSchedule) GetMonthScheduleOk() (*MonthSchedule, bool)`

GetMonthScheduleOk returns a tuple with the MonthSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthSchedule

`func (o *FullSchedule) SetMonthSchedule(v MonthSchedule)`

SetMonthSchedule sets MonthSchedule field to given value.

### HasMonthSchedule

`func (o *FullSchedule) HasMonthSchedule() bool`

HasMonthSchedule returns a boolean if a field has been set.

### GetUnit

`func (o *FullSchedule) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *FullSchedule) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *FullSchedule) SetUnit(v string)`

SetUnit sets Unit field to given value.


### SetUnitNil

`func (o *FullSchedule) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *FullSchedule) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetWeekSchedule

`func (o *FullSchedule) GetWeekSchedule() WeekSchedule`

GetWeekSchedule returns the WeekSchedule field if non-nil, zero value otherwise.

### GetWeekScheduleOk

`func (o *FullSchedule) GetWeekScheduleOk() (*WeekSchedule, bool)`

GetWeekScheduleOk returns a tuple with the WeekSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekSchedule

`func (o *FullSchedule) SetWeekSchedule(v WeekSchedule)`

SetWeekSchedule sets WeekSchedule field to given value.

### HasWeekSchedule

`func (o *FullSchedule) HasWeekSchedule() bool`

HasWeekSchedule returns a boolean if a field has been set.

### GetYearSchedule

`func (o *FullSchedule) GetYearSchedule() YearSchedule`

GetYearSchedule returns the YearSchedule field if non-nil, zero value otherwise.

### GetYearScheduleOk

`func (o *FullSchedule) GetYearScheduleOk() (*YearSchedule, bool)`

GetYearScheduleOk returns a tuple with the YearSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearSchedule

`func (o *FullSchedule) SetYearSchedule(v YearSchedule)`

SetYearSchedule sets YearSchedule field to given value.

### HasYearSchedule

`func (o *FullSchedule) HasYearSchedule() bool`

HasYearSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


