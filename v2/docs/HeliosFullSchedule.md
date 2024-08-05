# HeliosFullSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DaySchedule** | Pointer to [**HeliosDaySchedule**](HeliosDaySchedule.md) |  | [optional] 
**MonthSchedule** | Pointer to [**HeliosMonthSchedule**](HeliosMonthSchedule.md) |  | [optional] 
**Unit** | Pointer to **NullableString** | Specifies how often to start new runs of a Protection Group. &lt;br&gt;&#39;Days&#39; specifies that Protection Group run starts periodically on every day. For full backup schedule, currently we only support frequecny of 1 which indicates that full backup will be performed daily. &lt;br&gt;&#39;Weeks&#39; specifies that new Protection Group runs start weekly on certain days specified using &#39;dayOfWeek&#39; field. &lt;br&gt;&#39;Months&#39; specifies that new Protection Group runs start monthly on certain day of specific week. This schedule needs &#39;weekOfMonth&#39; and &#39;dayOfWeek&#39; fields to be set. &lt;br&gt;&#39;ProtectOnce&#39; specifies that groups using this policy option will run only once and after that group will permanently be disabled. &lt;br&gt; Example: To run the Protection Group on Second Sunday of Every Month, following schedule need to be set: &lt;br&gt; unit: &#39;Month&#39; &lt;br&gt; dayOfWeek: &#39;Sunday&#39; &lt;br&gt; weekOfMonth: &#39;Second&#39; | [optional] 
**WeekSchedule** | Pointer to [**HeliosWeekSchedule**](HeliosWeekSchedule.md) |  | [optional] 

## Methods

### NewHeliosFullSchedule

`func NewHeliosFullSchedule() *HeliosFullSchedule`

NewHeliosFullSchedule instantiates a new HeliosFullSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosFullScheduleWithDefaults

`func NewHeliosFullScheduleWithDefaults() *HeliosFullSchedule`

NewHeliosFullScheduleWithDefaults instantiates a new HeliosFullSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDaySchedule

`func (o *HeliosFullSchedule) GetDaySchedule() HeliosDaySchedule`

GetDaySchedule returns the DaySchedule field if non-nil, zero value otherwise.

### GetDayScheduleOk

`func (o *HeliosFullSchedule) GetDayScheduleOk() (*HeliosDaySchedule, bool)`

GetDayScheduleOk returns a tuple with the DaySchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaySchedule

`func (o *HeliosFullSchedule) SetDaySchedule(v HeliosDaySchedule)`

SetDaySchedule sets DaySchedule field to given value.

### HasDaySchedule

`func (o *HeliosFullSchedule) HasDaySchedule() bool`

HasDaySchedule returns a boolean if a field has been set.

### GetMonthSchedule

`func (o *HeliosFullSchedule) GetMonthSchedule() HeliosMonthSchedule`

GetMonthSchedule returns the MonthSchedule field if non-nil, zero value otherwise.

### GetMonthScheduleOk

`func (o *HeliosFullSchedule) GetMonthScheduleOk() (*HeliosMonthSchedule, bool)`

GetMonthScheduleOk returns a tuple with the MonthSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthSchedule

`func (o *HeliosFullSchedule) SetMonthSchedule(v HeliosMonthSchedule)`

SetMonthSchedule sets MonthSchedule field to given value.

### HasMonthSchedule

`func (o *HeliosFullSchedule) HasMonthSchedule() bool`

HasMonthSchedule returns a boolean if a field has been set.

### GetUnit

`func (o *HeliosFullSchedule) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *HeliosFullSchedule) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *HeliosFullSchedule) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *HeliosFullSchedule) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *HeliosFullSchedule) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *HeliosFullSchedule) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetWeekSchedule

`func (o *HeliosFullSchedule) GetWeekSchedule() HeliosWeekSchedule`

GetWeekSchedule returns the WeekSchedule field if non-nil, zero value otherwise.

### GetWeekScheduleOk

`func (o *HeliosFullSchedule) GetWeekScheduleOk() (*HeliosWeekSchedule, bool)`

GetWeekScheduleOk returns a tuple with the WeekSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekSchedule

`func (o *HeliosFullSchedule) SetWeekSchedule(v HeliosWeekSchedule)`

SetWeekSchedule sets WeekSchedule field to given value.

### HasWeekSchedule

`func (o *HeliosFullSchedule) HasWeekSchedule() bool`

HasWeekSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


