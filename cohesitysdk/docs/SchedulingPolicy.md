# SchedulingPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinuousSchedule** | Pointer to [**NullableContinuousSchedule**](ContinuousSchedule.md) | Specifies the time interval between two Job Runs of a continuous backup schedule and any blackout periods when new Job Runs should NOT be started. Set if periodicity is kContinuous. | [optional] 
**DailySchedule** | Pointer to [**NullableDailySchedule**](DailySchedule.md) | Specifies a daily or weekly backup schedule. Set if periodicity is kDaily. | [optional] 
**MonthlySchedule** | Pointer to [**NullableMonthlySchedule**](MonthlySchedule.md) | Specifies a monthly backup schedule. Set if periodicity is kMonthly. | [optional] 
**Periodicity** | Pointer to **NullableString** | Specifies how often to start new Job Runs of a Protection Job. &#39;kDaily&#39; means new Job Runs start daily. &#39;kMonthly&#39; means new Job Runs start monthly. &#39;kContinuous&#39; means new Job Runs repetitively start at the beginning of the specified time interval (in hours or minutes). &#39;kContinuousRPO&#39; means this is an RPO schedule. | [optional] 
**RpoSchedule** | Pointer to [**NullableRpoSchedule**](RpoSchedule.md) | Specifies an RPO backup schedule. Set if periodicity is kContinuousRPO. | [optional] 

## Methods

### NewSchedulingPolicy

`func NewSchedulingPolicy() *SchedulingPolicy`

NewSchedulingPolicy instantiates a new SchedulingPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulingPolicyWithDefaults

`func NewSchedulingPolicyWithDefaults() *SchedulingPolicy`

NewSchedulingPolicyWithDefaults instantiates a new SchedulingPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinuousSchedule

`func (o *SchedulingPolicy) GetContinuousSchedule() ContinuousSchedule`

GetContinuousSchedule returns the ContinuousSchedule field if non-nil, zero value otherwise.

### GetContinuousScheduleOk

`func (o *SchedulingPolicy) GetContinuousScheduleOk() (*ContinuousSchedule, bool)`

GetContinuousScheduleOk returns a tuple with the ContinuousSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuousSchedule

`func (o *SchedulingPolicy) SetContinuousSchedule(v ContinuousSchedule)`

SetContinuousSchedule sets ContinuousSchedule field to given value.

### HasContinuousSchedule

`func (o *SchedulingPolicy) HasContinuousSchedule() bool`

HasContinuousSchedule returns a boolean if a field has been set.

### SetContinuousScheduleNil

`func (o *SchedulingPolicy) SetContinuousScheduleNil(b bool)`

 SetContinuousScheduleNil sets the value for ContinuousSchedule to be an explicit nil

### UnsetContinuousSchedule
`func (o *SchedulingPolicy) UnsetContinuousSchedule()`

UnsetContinuousSchedule ensures that no value is present for ContinuousSchedule, not even an explicit nil
### GetDailySchedule

`func (o *SchedulingPolicy) GetDailySchedule() DailySchedule`

GetDailySchedule returns the DailySchedule field if non-nil, zero value otherwise.

### GetDailyScheduleOk

`func (o *SchedulingPolicy) GetDailyScheduleOk() (*DailySchedule, bool)`

GetDailyScheduleOk returns a tuple with the DailySchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailySchedule

`func (o *SchedulingPolicy) SetDailySchedule(v DailySchedule)`

SetDailySchedule sets DailySchedule field to given value.

### HasDailySchedule

`func (o *SchedulingPolicy) HasDailySchedule() bool`

HasDailySchedule returns a boolean if a field has been set.

### SetDailyScheduleNil

`func (o *SchedulingPolicy) SetDailyScheduleNil(b bool)`

 SetDailyScheduleNil sets the value for DailySchedule to be an explicit nil

### UnsetDailySchedule
`func (o *SchedulingPolicy) UnsetDailySchedule()`

UnsetDailySchedule ensures that no value is present for DailySchedule, not even an explicit nil
### GetMonthlySchedule

`func (o *SchedulingPolicy) GetMonthlySchedule() MonthlySchedule`

GetMonthlySchedule returns the MonthlySchedule field if non-nil, zero value otherwise.

### GetMonthlyScheduleOk

`func (o *SchedulingPolicy) GetMonthlyScheduleOk() (*MonthlySchedule, bool)`

GetMonthlyScheduleOk returns a tuple with the MonthlySchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlySchedule

`func (o *SchedulingPolicy) SetMonthlySchedule(v MonthlySchedule)`

SetMonthlySchedule sets MonthlySchedule field to given value.

### HasMonthlySchedule

`func (o *SchedulingPolicy) HasMonthlySchedule() bool`

HasMonthlySchedule returns a boolean if a field has been set.

### SetMonthlyScheduleNil

`func (o *SchedulingPolicy) SetMonthlyScheduleNil(b bool)`

 SetMonthlyScheduleNil sets the value for MonthlySchedule to be an explicit nil

### UnsetMonthlySchedule
`func (o *SchedulingPolicy) UnsetMonthlySchedule()`

UnsetMonthlySchedule ensures that no value is present for MonthlySchedule, not even an explicit nil
### GetPeriodicity

`func (o *SchedulingPolicy) GetPeriodicity() string`

GetPeriodicity returns the Periodicity field if non-nil, zero value otherwise.

### GetPeriodicityOk

`func (o *SchedulingPolicy) GetPeriodicityOk() (*string, bool)`

GetPeriodicityOk returns a tuple with the Periodicity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicity

`func (o *SchedulingPolicy) SetPeriodicity(v string)`

SetPeriodicity sets Periodicity field to given value.

### HasPeriodicity

`func (o *SchedulingPolicy) HasPeriodicity() bool`

HasPeriodicity returns a boolean if a field has been set.

### SetPeriodicityNil

`func (o *SchedulingPolicy) SetPeriodicityNil(b bool)`

 SetPeriodicityNil sets the value for Periodicity to be an explicit nil

### UnsetPeriodicity
`func (o *SchedulingPolicy) UnsetPeriodicity()`

UnsetPeriodicity ensures that no value is present for Periodicity, not even an explicit nil
### GetRpoSchedule

`func (o *SchedulingPolicy) GetRpoSchedule() RpoSchedule`

GetRpoSchedule returns the RpoSchedule field if non-nil, zero value otherwise.

### GetRpoScheduleOk

`func (o *SchedulingPolicy) GetRpoScheduleOk() (*RpoSchedule, bool)`

GetRpoScheduleOk returns a tuple with the RpoSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpoSchedule

`func (o *SchedulingPolicy) SetRpoSchedule(v RpoSchedule)`

SetRpoSchedule sets RpoSchedule field to given value.

### HasRpoSchedule

`func (o *SchedulingPolicy) HasRpoSchedule() bool`

HasRpoSchedule returns a boolean if a field has been set.

### SetRpoScheduleNil

`func (o *SchedulingPolicy) SetRpoScheduleNil(b bool)`

 SetRpoScheduleNil sets the value for RpoSchedule to be an explicit nil

### UnsetRpoSchedule
`func (o *SchedulingPolicy) UnsetRpoSchedule()`

UnsetRpoSchedule ensures that no value is present for RpoSchedule, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


