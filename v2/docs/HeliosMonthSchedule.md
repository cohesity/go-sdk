# HeliosMonthSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfWeek** | Pointer to **[]string** | Specifies a list of days of the week when to start Protection Group Runs. &lt;br&gt; Example: To run a Protection Group on every Monday and Tuesday, set the schedule with following values: &lt;br&gt;  unit: &#39;Weeks&#39; &lt;br&gt;  dayOfWeek: [&#39;Monday&#39;,&#39;Tuesday&#39;] | [optional] 
**WeekOfMonth** | Pointer to **NullableString** | Specifies the week of the month (such as &#39;Third&#39;) in a Monthly Schedule specified by unit field as &#39;Months&#39;. &lt;br&gt;This field is used in combination with &#39;dayOfWeek&#39; to define the day in the month to start the Protection Group Run. &lt;br&gt; Example: if &#39;weekOfMonth&#39; is set to &#39;Third&#39; and day is set to &#39;Monday&#39;, a backup is performed on the third Monday of every month. | [optional] 

## Methods

### NewHeliosMonthSchedule

`func NewHeliosMonthSchedule() *HeliosMonthSchedule`

NewHeliosMonthSchedule instantiates a new HeliosMonthSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosMonthScheduleWithDefaults

`func NewHeliosMonthScheduleWithDefaults() *HeliosMonthSchedule`

NewHeliosMonthScheduleWithDefaults instantiates a new HeliosMonthSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfWeek

`func (o *HeliosMonthSchedule) GetDayOfWeek() []string`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *HeliosMonthSchedule) GetDayOfWeekOk() (*[]string, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *HeliosMonthSchedule) SetDayOfWeek(v []string)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *HeliosMonthSchedule) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### SetDayOfWeekNil

`func (o *HeliosMonthSchedule) SetDayOfWeekNil(b bool)`

 SetDayOfWeekNil sets the value for DayOfWeek to be an explicit nil

### UnsetDayOfWeek
`func (o *HeliosMonthSchedule) UnsetDayOfWeek()`

UnsetDayOfWeek ensures that no value is present for DayOfWeek, not even an explicit nil
### GetWeekOfMonth

`func (o *HeliosMonthSchedule) GetWeekOfMonth() string`

GetWeekOfMonth returns the WeekOfMonth field if non-nil, zero value otherwise.

### GetWeekOfMonthOk

`func (o *HeliosMonthSchedule) GetWeekOfMonthOk() (*string, bool)`

GetWeekOfMonthOk returns a tuple with the WeekOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekOfMonth

`func (o *HeliosMonthSchedule) SetWeekOfMonth(v string)`

SetWeekOfMonth sets WeekOfMonth field to given value.

### HasWeekOfMonth

`func (o *HeliosMonthSchedule) HasWeekOfMonth() bool`

HasWeekOfMonth returns a boolean if a field has been set.

### SetWeekOfMonthNil

`func (o *HeliosMonthSchedule) SetWeekOfMonthNil(b bool)`

 SetWeekOfMonthNil sets the value for WeekOfMonth to be an explicit nil

### UnsetWeekOfMonth
`func (o *HeliosMonthSchedule) UnsetWeekOfMonth()`

UnsetWeekOfMonth ensures that no value is present for WeekOfMonth, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


