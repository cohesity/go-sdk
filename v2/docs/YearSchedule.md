# YearSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfYear** | **NullableString** | Specifies the day of the Year (such as &#39;First&#39; or &#39;Last&#39;) in a Yearly Schedule. &lt;br&gt;This field is used to define the day in the year to start the Protection Group Run. &lt;br&gt; Example: if &#39;dayOfYear&#39; is set to &#39;First&#39;, a backup is performed on the first day of every year. &lt;br&gt; Example: if &#39;dayOfYear&#39; is set to &#39;Last&#39;, a backup is performed on the last day of every year. | 

## Methods

### NewYearSchedule

`func NewYearSchedule(dayOfYear NullableString, ) *YearSchedule`

NewYearSchedule instantiates a new YearSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYearScheduleWithDefaults

`func NewYearScheduleWithDefaults() *YearSchedule`

NewYearScheduleWithDefaults instantiates a new YearSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfYear

`func (o *YearSchedule) GetDayOfYear() string`

GetDayOfYear returns the DayOfYear field if non-nil, zero value otherwise.

### GetDayOfYearOk

`func (o *YearSchedule) GetDayOfYearOk() (*string, bool)`

GetDayOfYearOk returns a tuple with the DayOfYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfYear

`func (o *YearSchedule) SetDayOfYear(v string)`

SetDayOfYear sets DayOfYear field to given value.


### SetDayOfYearNil

`func (o *YearSchedule) SetDayOfYearNil(b bool)`

 SetDayOfYearNil sets the value for DayOfYear to be an explicit nil

### UnsetDayOfYear
`func (o *YearSchedule) UnsetDayOfYear()`

UnsetDayOfYear ensures that no value is present for DayOfYear, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


