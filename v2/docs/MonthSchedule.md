# MonthSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfMonth** | Pointer to **NullableInt32** | Specifies the exact date of the month (such as 18) in a Monthly Schedule specified by unit field as &#39;Years&#39;. &lt;br&gt; Example: if &#39;dayOfMonth&#39; is set to &#39;18&#39;, a backup is performed on the 18th of every month. | [optional] 
**DayOfWeek** | Pointer to **[]string** | Specifies a list of days of the week when to start Protection Group Runs. &lt;br&gt; Example: To run a Protection Group on every Monday and Tuesday, set the schedule with following values: &lt;br&gt;  unit: &#39;Weeks&#39; &lt;br&gt;  dayOfWeek: [&#39;Monday&#39;,&#39;Tuesday&#39;] | [optional] 
**WeekOfMonth** | Pointer to **NullableString** | Specifies the week of the month (such as &#39;Third&#39;) or nth day of month (such as &#39;First&#39; or &#39;Last&#39;) in a Monthly Schedule specified by unit field as &#39;Months&#39;. &lt;br&gt;This field can be used in combination with &#39;dayOfWeek&#39; to define the day in the month to start the Protection Group Run. &lt;br&gt; Example: if &#39;weekOfMonth&#39; is set to &#39;Third&#39; and day is set to &#39;Monday&#39;, a backup is performed on the third Monday of every month. &lt;br&gt; Example: if &#39;weekOfMonth&#39; is set to &#39;Last&#39; and dayOfWeek is not set, a backup is performed on the last day of every month. | [optional] 

## Methods

### NewMonthSchedule

`func NewMonthSchedule() *MonthSchedule`

NewMonthSchedule instantiates a new MonthSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonthScheduleWithDefaults

`func NewMonthScheduleWithDefaults() *MonthSchedule`

NewMonthScheduleWithDefaults instantiates a new MonthSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfMonth

`func (o *MonthSchedule) GetDayOfMonth() int32`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *MonthSchedule) GetDayOfMonthOk() (*int32, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *MonthSchedule) SetDayOfMonth(v int32)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *MonthSchedule) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.

### SetDayOfMonthNil

`func (o *MonthSchedule) SetDayOfMonthNil(b bool)`

 SetDayOfMonthNil sets the value for DayOfMonth to be an explicit nil

### UnsetDayOfMonth
`func (o *MonthSchedule) UnsetDayOfMonth()`

UnsetDayOfMonth ensures that no value is present for DayOfMonth, not even an explicit nil
### GetDayOfWeek

`func (o *MonthSchedule) GetDayOfWeek() []string`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *MonthSchedule) GetDayOfWeekOk() (*[]string, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *MonthSchedule) SetDayOfWeek(v []string)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *MonthSchedule) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### SetDayOfWeekNil

`func (o *MonthSchedule) SetDayOfWeekNil(b bool)`

 SetDayOfWeekNil sets the value for DayOfWeek to be an explicit nil

### UnsetDayOfWeek
`func (o *MonthSchedule) UnsetDayOfWeek()`

UnsetDayOfWeek ensures that no value is present for DayOfWeek, not even an explicit nil
### GetWeekOfMonth

`func (o *MonthSchedule) GetWeekOfMonth() string`

GetWeekOfMonth returns the WeekOfMonth field if non-nil, zero value otherwise.

### GetWeekOfMonthOk

`func (o *MonthSchedule) GetWeekOfMonthOk() (*string, bool)`

GetWeekOfMonthOk returns a tuple with the WeekOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekOfMonth

`func (o *MonthSchedule) SetWeekOfMonth(v string)`

SetWeekOfMonth sets WeekOfMonth field to given value.

### HasWeekOfMonth

`func (o *MonthSchedule) HasWeekOfMonth() bool`

HasWeekOfMonth returns a boolean if a field has been set.

### SetWeekOfMonthNil

`func (o *MonthSchedule) SetWeekOfMonthNil(b bool)`

 SetWeekOfMonthNil sets the value for WeekOfMonth to be an explicit nil

### UnsetWeekOfMonth
`func (o *MonthSchedule) UnsetWeekOfMonth()`

UnsetWeekOfMonth ensures that no value is present for WeekOfMonth, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


