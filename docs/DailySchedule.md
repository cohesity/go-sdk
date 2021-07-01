# DailySchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | Pointer to **[]string** | Array of Days.  Specifies a list of days of the week when to start Job Runs. If no days are specified, the Jobs Runs will run every day of the week. Specifies a day in a week such as &#39;kSunday&#39;, &#39;kMonday&#39;, etc. | [optional] 

## Methods

### NewDailySchedule

`func NewDailySchedule() *DailySchedule`

NewDailySchedule instantiates a new DailySchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDailyScheduleWithDefaults

`func NewDailyScheduleWithDefaults() *DailySchedule`

NewDailyScheduleWithDefaults instantiates a new DailySchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *DailySchedule) GetDays() []string`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *DailySchedule) GetDaysOk() (*[]string, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *DailySchedule) SetDays(v []string)`

SetDays sets Days field to given value.

### HasDays

`func (o *DailySchedule) HasDays() bool`

HasDays returns a boolean if a field has been set.

### SetDaysNil

`func (o *DailySchedule) SetDaysNil(b bool)`

 SetDaysNil sets the value for Days to be an explicit nil

### UnsetDays
`func (o *DailySchedule) UnsetDays()`

UnsetDays ensures that no value is present for Days, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


