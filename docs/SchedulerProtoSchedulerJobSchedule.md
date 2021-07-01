# SchedulerProtoSchedulerJobSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Day** | Pointer to **NullableInt32** | The day of the week when schedule should be executed (0-6). | [optional] 
**Hour** | Pointer to **NullableInt32** | The hour of the day when schedule should be executed (0-23). | [optional] 
**Timezone** | Pointer to **NullableString** | The timezone for the execution of the schedule. | [optional] 

## Methods

### NewSchedulerProtoSchedulerJobSchedule

`func NewSchedulerProtoSchedulerJobSchedule() *SchedulerProtoSchedulerJobSchedule`

NewSchedulerProtoSchedulerJobSchedule instantiates a new SchedulerProtoSchedulerJobSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulerProtoSchedulerJobScheduleWithDefaults

`func NewSchedulerProtoSchedulerJobScheduleWithDefaults() *SchedulerProtoSchedulerJobSchedule`

NewSchedulerProtoSchedulerJobScheduleWithDefaults instantiates a new SchedulerProtoSchedulerJobSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDay

`func (o *SchedulerProtoSchedulerJobSchedule) GetDay() int32`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *SchedulerProtoSchedulerJobSchedule) GetDayOk() (*int32, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *SchedulerProtoSchedulerJobSchedule) SetDay(v int32)`

SetDay sets Day field to given value.

### HasDay

`func (o *SchedulerProtoSchedulerJobSchedule) HasDay() bool`

HasDay returns a boolean if a field has been set.

### SetDayNil

`func (o *SchedulerProtoSchedulerJobSchedule) SetDayNil(b bool)`

 SetDayNil sets the value for Day to be an explicit nil

### UnsetDay
`func (o *SchedulerProtoSchedulerJobSchedule) UnsetDay()`

UnsetDay ensures that no value is present for Day, not even an explicit nil
### GetHour

`func (o *SchedulerProtoSchedulerJobSchedule) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *SchedulerProtoSchedulerJobSchedule) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *SchedulerProtoSchedulerJobSchedule) SetHour(v int32)`

SetHour sets Hour field to given value.

### HasHour

`func (o *SchedulerProtoSchedulerJobSchedule) HasHour() bool`

HasHour returns a boolean if a field has been set.

### SetHourNil

`func (o *SchedulerProtoSchedulerJobSchedule) SetHourNil(b bool)`

 SetHourNil sets the value for Hour to be an explicit nil

### UnsetHour
`func (o *SchedulerProtoSchedulerJobSchedule) UnsetHour()`

UnsetHour ensures that no value is present for Hour, not even an explicit nil
### GetTimezone

`func (o *SchedulerProtoSchedulerJobSchedule) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *SchedulerProtoSchedulerJobSchedule) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *SchedulerProtoSchedulerJobSchedule) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *SchedulerProtoSchedulerJobSchedule) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *SchedulerProtoSchedulerJobSchedule) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *SchedulerProtoSchedulerJobSchedule) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


