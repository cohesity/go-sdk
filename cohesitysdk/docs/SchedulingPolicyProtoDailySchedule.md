# SchedulingPolicyProtoDailySchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | Pointer to **[]int32** | The list of weekdays for scheduling a backup. This is populated only for selected weekday schedules. | [optional] 
**Frequency** | Pointer to **NullableInt64** | This is set only for every-n-day schedules. | [optional] 

## Methods

### NewSchedulingPolicyProtoDailySchedule

`func NewSchedulingPolicyProtoDailySchedule() *SchedulingPolicyProtoDailySchedule`

NewSchedulingPolicyProtoDailySchedule instantiates a new SchedulingPolicyProtoDailySchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulingPolicyProtoDailyScheduleWithDefaults

`func NewSchedulingPolicyProtoDailyScheduleWithDefaults() *SchedulingPolicyProtoDailySchedule`

NewSchedulingPolicyProtoDailyScheduleWithDefaults instantiates a new SchedulingPolicyProtoDailySchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *SchedulingPolicyProtoDailySchedule) GetDays() []int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *SchedulingPolicyProtoDailySchedule) GetDaysOk() (*[]int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *SchedulingPolicyProtoDailySchedule) SetDays(v []int32)`

SetDays sets Days field to given value.

### HasDays

`func (o *SchedulingPolicyProtoDailySchedule) HasDays() bool`

HasDays returns a boolean if a field has been set.

### SetDaysNil

`func (o *SchedulingPolicyProtoDailySchedule) SetDaysNil(b bool)`

 SetDaysNil sets the value for Days to be an explicit nil

### UnsetDays
`func (o *SchedulingPolicyProtoDailySchedule) UnsetDays()`

UnsetDays ensures that no value is present for Days, not even an explicit nil
### GetFrequency

`func (o *SchedulingPolicyProtoDailySchedule) GetFrequency() int64`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *SchedulingPolicyProtoDailySchedule) GetFrequencyOk() (*int64, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *SchedulingPolicyProtoDailySchedule) SetFrequency(v int64)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *SchedulingPolicyProtoDailySchedule) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### SetFrequencyNil

`func (o *SchedulingPolicyProtoDailySchedule) SetFrequencyNil(b bool)`

 SetFrequencyNil sets the value for Frequency to be an explicit nil

### UnsetFrequency
`func (o *SchedulingPolicyProtoDailySchedule) UnsetFrequency()`

UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


