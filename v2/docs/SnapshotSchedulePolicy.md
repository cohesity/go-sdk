# SnapshotSchedulePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DaysOfMonth** | Pointer to **[]int32** | Days of the month. | [optional] 
**DaysOfWeek** | Pointer to **[]string** | Days of the week. | [optional] 
**Suspended** | Pointer to **NullableBool** | Bool to denote if the policy is suspended. | [optional] 
**Time** | Pointer to **NullableString** | Time of the day. | [optional] 
**TimeZone** | Pointer to **NullableString** | Time zone. | [optional] 

## Methods

### NewSnapshotSchedulePolicy

`func NewSnapshotSchedulePolicy() *SnapshotSchedulePolicy`

NewSnapshotSchedulePolicy instantiates a new SnapshotSchedulePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotSchedulePolicyWithDefaults

`func NewSnapshotSchedulePolicyWithDefaults() *SnapshotSchedulePolicy`

NewSnapshotSchedulePolicyWithDefaults instantiates a new SnapshotSchedulePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDaysOfMonth

`func (o *SnapshotSchedulePolicy) GetDaysOfMonth() []int32`

GetDaysOfMonth returns the DaysOfMonth field if non-nil, zero value otherwise.

### GetDaysOfMonthOk

`func (o *SnapshotSchedulePolicy) GetDaysOfMonthOk() (*[]int32, bool)`

GetDaysOfMonthOk returns a tuple with the DaysOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysOfMonth

`func (o *SnapshotSchedulePolicy) SetDaysOfMonth(v []int32)`

SetDaysOfMonth sets DaysOfMonth field to given value.

### HasDaysOfMonth

`func (o *SnapshotSchedulePolicy) HasDaysOfMonth() bool`

HasDaysOfMonth returns a boolean if a field has been set.

### SetDaysOfMonthNil

`func (o *SnapshotSchedulePolicy) SetDaysOfMonthNil(b bool)`

 SetDaysOfMonthNil sets the value for DaysOfMonth to be an explicit nil

### UnsetDaysOfMonth
`func (o *SnapshotSchedulePolicy) UnsetDaysOfMonth()`

UnsetDaysOfMonth ensures that no value is present for DaysOfMonth, not even an explicit nil
### GetDaysOfWeek

`func (o *SnapshotSchedulePolicy) GetDaysOfWeek() []string`

GetDaysOfWeek returns the DaysOfWeek field if non-nil, zero value otherwise.

### GetDaysOfWeekOk

`func (o *SnapshotSchedulePolicy) GetDaysOfWeekOk() (*[]string, bool)`

GetDaysOfWeekOk returns a tuple with the DaysOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysOfWeek

`func (o *SnapshotSchedulePolicy) SetDaysOfWeek(v []string)`

SetDaysOfWeek sets DaysOfWeek field to given value.

### HasDaysOfWeek

`func (o *SnapshotSchedulePolicy) HasDaysOfWeek() bool`

HasDaysOfWeek returns a boolean if a field has been set.

### SetDaysOfWeekNil

`func (o *SnapshotSchedulePolicy) SetDaysOfWeekNil(b bool)`

 SetDaysOfWeekNil sets the value for DaysOfWeek to be an explicit nil

### UnsetDaysOfWeek
`func (o *SnapshotSchedulePolicy) UnsetDaysOfWeek()`

UnsetDaysOfWeek ensures that no value is present for DaysOfWeek, not even an explicit nil
### GetSuspended

`func (o *SnapshotSchedulePolicy) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *SnapshotSchedulePolicy) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *SnapshotSchedulePolicy) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *SnapshotSchedulePolicy) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### SetSuspendedNil

`func (o *SnapshotSchedulePolicy) SetSuspendedNil(b bool)`

 SetSuspendedNil sets the value for Suspended to be an explicit nil

### UnsetSuspended
`func (o *SnapshotSchedulePolicy) UnsetSuspended()`

UnsetSuspended ensures that no value is present for Suspended, not even an explicit nil
### GetTime

`func (o *SnapshotSchedulePolicy) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *SnapshotSchedulePolicy) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *SnapshotSchedulePolicy) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *SnapshotSchedulePolicy) HasTime() bool`

HasTime returns a boolean if a field has been set.

### SetTimeNil

`func (o *SnapshotSchedulePolicy) SetTimeNil(b bool)`

 SetTimeNil sets the value for Time to be an explicit nil

### UnsetTime
`func (o *SnapshotSchedulePolicy) UnsetTime()`

UnsetTime ensures that no value is present for Time, not even an explicit nil
### GetTimeZone

`func (o *SnapshotSchedulePolicy) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *SnapshotSchedulePolicy) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *SnapshotSchedulePolicy) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *SnapshotSchedulePolicy) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### SetTimeZoneNil

`func (o *SnapshotSchedulePolicy) SetTimeZoneNil(b bool)`

 SetTimeZoneNil sets the value for TimeZone to be an explicit nil

### UnsetTimeZone
`func (o *SnapshotSchedulePolicy) UnsetTimeZone()`

UnsetTimeZone ensures that no value is present for TimeZone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


