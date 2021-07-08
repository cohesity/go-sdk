# MonthlySchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Day** | Pointer to **NullableString** | Specifies the day of the week (such as &#39;kMonday&#39;) to start the Job Run. Used with day count to define the day in the month to start the Job Run. Specifies a day in a week such as &#39;kSunday&#39;, &#39;kMonday&#39;, etc. | [optional] 
**DayCount** | Pointer to **NullableString** | Specifies the day count in the month (such as &#39;kThird&#39;) to start the Job Run. Used in combination with day to define the day in the month to start the Job Run. Specifies the day count in the month to start the backup. For example if day count is set to &#39;kThird&#39; and day is set to &#39;kMonday&#39;, a backup is performed on the third Monday of every month. &#39;kFirst&#39; indicates that the first week should be chosen for specified day of every month. &#39;kSecond&#39; indicates that the second week should be chosen for specified day of every month. &#39;kThird&#39; indicates that the third week should be chosen for specified day of every month. &#39;kFourth&#39; indicates that the fourth week should be chosen for specified day of every month. &#39;kLast&#39; indicates that the last week should be chosen for specified day of every month. | [optional] 

## Methods

### NewMonthlySchedule

`func NewMonthlySchedule() *MonthlySchedule`

NewMonthlySchedule instantiates a new MonthlySchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonthlyScheduleWithDefaults

`func NewMonthlyScheduleWithDefaults() *MonthlySchedule`

NewMonthlyScheduleWithDefaults instantiates a new MonthlySchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDay

`func (o *MonthlySchedule) GetDay() string`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *MonthlySchedule) GetDayOk() (*string, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *MonthlySchedule) SetDay(v string)`

SetDay sets Day field to given value.

### HasDay

`func (o *MonthlySchedule) HasDay() bool`

HasDay returns a boolean if a field has been set.

### SetDayNil

`func (o *MonthlySchedule) SetDayNil(b bool)`

 SetDayNil sets the value for Day to be an explicit nil

### UnsetDay
`func (o *MonthlySchedule) UnsetDay()`

UnsetDay ensures that no value is present for Day, not even an explicit nil
### GetDayCount

`func (o *MonthlySchedule) GetDayCount() string`

GetDayCount returns the DayCount field if non-nil, zero value otherwise.

### GetDayCountOk

`func (o *MonthlySchedule) GetDayCountOk() (*string, bool)`

GetDayCountOk returns a tuple with the DayCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayCount

`func (o *MonthlySchedule) SetDayCount(v string)`

SetDayCount sets DayCount field to given value.

### HasDayCount

`func (o *MonthlySchedule) HasDayCount() bool`

HasDayCount returns a boolean if a field has been set.

### SetDayCountNil

`func (o *MonthlySchedule) SetDayCountNil(b bool)`

 SetDayCountNil sets the value for DayCount to be an explicit nil

### UnsetDayCount
`func (o *MonthlySchedule) UnsetDayCount()`

UnsetDayCount ensures that no value is present for DayCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


