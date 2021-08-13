# MonthScheduleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WeekOfMonth** | **NullableString** | Specifies the week of the month (such as &#39;Third&#39;) in a Monthly Schedule specified by unit field as &#39;Months&#39;. &lt;br&gt;This field is used in combination with &#39;dayOfWeek&#39; to define the day in the month to start the Protection Group Run. &lt;br&gt; Example: if &#39;weekOfMonth&#39; is set to &#39;Third&#39; and day is set to &#39;Monday&#39;, a backup is performed on the third Monday of every month. | 

## Methods

### NewMonthScheduleAllOf

`func NewMonthScheduleAllOf(weekOfMonth NullableString, ) *MonthScheduleAllOf`

NewMonthScheduleAllOf instantiates a new MonthScheduleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonthScheduleAllOfWithDefaults

`func NewMonthScheduleAllOfWithDefaults() *MonthScheduleAllOf`

NewMonthScheduleAllOfWithDefaults instantiates a new MonthScheduleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWeekOfMonth

`func (o *MonthScheduleAllOf) GetWeekOfMonth() string`

GetWeekOfMonth returns the WeekOfMonth field if non-nil, zero value otherwise.

### GetWeekOfMonthOk

`func (o *MonthScheduleAllOf) GetWeekOfMonthOk() (*string, bool)`

GetWeekOfMonthOk returns a tuple with the WeekOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekOfMonth

`func (o *MonthScheduleAllOf) SetWeekOfMonth(v string)`

SetWeekOfMonth sets WeekOfMonth field to given value.


### SetWeekOfMonthNil

`func (o *MonthScheduleAllOf) SetWeekOfMonthNil(b bool)`

 SetWeekOfMonthNil sets the value for WeekOfMonth to be an explicit nil

### UnsetWeekOfMonth
`func (o *MonthScheduleAllOf) UnsetWeekOfMonth()`

UnsetWeekOfMonth ensures that no value is present for WeekOfMonth, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


