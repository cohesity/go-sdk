# SchedulingPolicyProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinuousSchedule** | Pointer to [**SchedulingPolicyProtoContinuousSchedule**](SchedulingPolicyProtoContinuousSchedule.md) |  | [optional] 
**DailySchedule** | Pointer to [**SchedulingPolicyProtoDailySchedule**](SchedulingPolicyProtoDailySchedule.md) |  | [optional] 
**MonthlySchedule** | Pointer to [**SchedulingPolicyProtoMonthlySchedule**](SchedulingPolicyProtoMonthlySchedule.md) |  | [optional] 
**Periodicity** | Pointer to **NullableInt32** | Determines how often the job should be run. | [optional] 
**RpoSchedule** | Pointer to [**SchedulingPolicyProtoRPOSchedule**](SchedulingPolicyProtoRPOSchedule.md) |  | [optional] 

## Methods

### NewSchedulingPolicyProto

`func NewSchedulingPolicyProto() *SchedulingPolicyProto`

NewSchedulingPolicyProto instantiates a new SchedulingPolicyProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulingPolicyProtoWithDefaults

`func NewSchedulingPolicyProtoWithDefaults() *SchedulingPolicyProto`

NewSchedulingPolicyProtoWithDefaults instantiates a new SchedulingPolicyProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinuousSchedule

`func (o *SchedulingPolicyProto) GetContinuousSchedule() SchedulingPolicyProtoContinuousSchedule`

GetContinuousSchedule returns the ContinuousSchedule field if non-nil, zero value otherwise.

### GetContinuousScheduleOk

`func (o *SchedulingPolicyProto) GetContinuousScheduleOk() (*SchedulingPolicyProtoContinuousSchedule, bool)`

GetContinuousScheduleOk returns a tuple with the ContinuousSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuousSchedule

`func (o *SchedulingPolicyProto) SetContinuousSchedule(v SchedulingPolicyProtoContinuousSchedule)`

SetContinuousSchedule sets ContinuousSchedule field to given value.

### HasContinuousSchedule

`func (o *SchedulingPolicyProto) HasContinuousSchedule() bool`

HasContinuousSchedule returns a boolean if a field has been set.

### GetDailySchedule

`func (o *SchedulingPolicyProto) GetDailySchedule() SchedulingPolicyProtoDailySchedule`

GetDailySchedule returns the DailySchedule field if non-nil, zero value otherwise.

### GetDailyScheduleOk

`func (o *SchedulingPolicyProto) GetDailyScheduleOk() (*SchedulingPolicyProtoDailySchedule, bool)`

GetDailyScheduleOk returns a tuple with the DailySchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailySchedule

`func (o *SchedulingPolicyProto) SetDailySchedule(v SchedulingPolicyProtoDailySchedule)`

SetDailySchedule sets DailySchedule field to given value.

### HasDailySchedule

`func (o *SchedulingPolicyProto) HasDailySchedule() bool`

HasDailySchedule returns a boolean if a field has been set.

### GetMonthlySchedule

`func (o *SchedulingPolicyProto) GetMonthlySchedule() SchedulingPolicyProtoMonthlySchedule`

GetMonthlySchedule returns the MonthlySchedule field if non-nil, zero value otherwise.

### GetMonthlyScheduleOk

`func (o *SchedulingPolicyProto) GetMonthlyScheduleOk() (*SchedulingPolicyProtoMonthlySchedule, bool)`

GetMonthlyScheduleOk returns a tuple with the MonthlySchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlySchedule

`func (o *SchedulingPolicyProto) SetMonthlySchedule(v SchedulingPolicyProtoMonthlySchedule)`

SetMonthlySchedule sets MonthlySchedule field to given value.

### HasMonthlySchedule

`func (o *SchedulingPolicyProto) HasMonthlySchedule() bool`

HasMonthlySchedule returns a boolean if a field has been set.

### GetPeriodicity

`func (o *SchedulingPolicyProto) GetPeriodicity() int32`

GetPeriodicity returns the Periodicity field if non-nil, zero value otherwise.

### GetPeriodicityOk

`func (o *SchedulingPolicyProto) GetPeriodicityOk() (*int32, bool)`

GetPeriodicityOk returns a tuple with the Periodicity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicity

`func (o *SchedulingPolicyProto) SetPeriodicity(v int32)`

SetPeriodicity sets Periodicity field to given value.

### HasPeriodicity

`func (o *SchedulingPolicyProto) HasPeriodicity() bool`

HasPeriodicity returns a boolean if a field has been set.

### SetPeriodicityNil

`func (o *SchedulingPolicyProto) SetPeriodicityNil(b bool)`

 SetPeriodicityNil sets the value for Periodicity to be an explicit nil

### UnsetPeriodicity
`func (o *SchedulingPolicyProto) UnsetPeriodicity()`

UnsetPeriodicity ensures that no value is present for Periodicity, not even an explicit nil
### GetRpoSchedule

`func (o *SchedulingPolicyProto) GetRpoSchedule() SchedulingPolicyProtoRPOSchedule`

GetRpoSchedule returns the RpoSchedule field if non-nil, zero value otherwise.

### GetRpoScheduleOk

`func (o *SchedulingPolicyProto) GetRpoScheduleOk() (*SchedulingPolicyProtoRPOSchedule, bool)`

GetRpoScheduleOk returns a tuple with the RpoSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpoSchedule

`func (o *SchedulingPolicyProto) SetRpoSchedule(v SchedulingPolicyProtoRPOSchedule)`

SetRpoSchedule sets RpoSchedule field to given value.

### HasRpoSchedule

`func (o *SchedulingPolicyProto) HasRpoSchedule() bool`

HasRpoSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


