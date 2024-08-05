# UpdateSnapshotPolicyParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DaysOfMonth** | Pointer to **[]int32** | Days of the month. | [optional] 
**DaysOfWeek** | Pointer to **[]string** | Days of the week. | [optional] 
**NumOfDaysToKeep** | Pointer to **NullableInt64** | Number of days to keep the snapshot. | [optional] 
**NumOfVersionsToKeep** | Pointer to **NullableInt64** | Number of snapshot versions to keep. | [optional] 
**SuspendRetentionPolicy** | Pointer to **NullableBool** | Suspend snapshot retention policy. | [optional] 
**SuspendSchedulePolicy** | Pointer to **NullableBool** | Suspend snapshot schedule policy. | [optional] 
**Time** | Pointer to **NullableString** | Time of the day. | [optional] 
**TimeZone** | Pointer to **NullableString** | Time zone. | [optional] 

## Methods

### NewUpdateSnapshotPolicyParams

`func NewUpdateSnapshotPolicyParams() *UpdateSnapshotPolicyParams`

NewUpdateSnapshotPolicyParams instantiates a new UpdateSnapshotPolicyParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSnapshotPolicyParamsWithDefaults

`func NewUpdateSnapshotPolicyParamsWithDefaults() *UpdateSnapshotPolicyParams`

NewUpdateSnapshotPolicyParamsWithDefaults instantiates a new UpdateSnapshotPolicyParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDaysOfMonth

`func (o *UpdateSnapshotPolicyParams) GetDaysOfMonth() []int32`

GetDaysOfMonth returns the DaysOfMonth field if non-nil, zero value otherwise.

### GetDaysOfMonthOk

`func (o *UpdateSnapshotPolicyParams) GetDaysOfMonthOk() (*[]int32, bool)`

GetDaysOfMonthOk returns a tuple with the DaysOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysOfMonth

`func (o *UpdateSnapshotPolicyParams) SetDaysOfMonth(v []int32)`

SetDaysOfMonth sets DaysOfMonth field to given value.

### HasDaysOfMonth

`func (o *UpdateSnapshotPolicyParams) HasDaysOfMonth() bool`

HasDaysOfMonth returns a boolean if a field has been set.

### SetDaysOfMonthNil

`func (o *UpdateSnapshotPolicyParams) SetDaysOfMonthNil(b bool)`

 SetDaysOfMonthNil sets the value for DaysOfMonth to be an explicit nil

### UnsetDaysOfMonth
`func (o *UpdateSnapshotPolicyParams) UnsetDaysOfMonth()`

UnsetDaysOfMonth ensures that no value is present for DaysOfMonth, not even an explicit nil
### GetDaysOfWeek

`func (o *UpdateSnapshotPolicyParams) GetDaysOfWeek() []string`

GetDaysOfWeek returns the DaysOfWeek field if non-nil, zero value otherwise.

### GetDaysOfWeekOk

`func (o *UpdateSnapshotPolicyParams) GetDaysOfWeekOk() (*[]string, bool)`

GetDaysOfWeekOk returns a tuple with the DaysOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysOfWeek

`func (o *UpdateSnapshotPolicyParams) SetDaysOfWeek(v []string)`

SetDaysOfWeek sets DaysOfWeek field to given value.

### HasDaysOfWeek

`func (o *UpdateSnapshotPolicyParams) HasDaysOfWeek() bool`

HasDaysOfWeek returns a boolean if a field has been set.

### SetDaysOfWeekNil

`func (o *UpdateSnapshotPolicyParams) SetDaysOfWeekNil(b bool)`

 SetDaysOfWeekNil sets the value for DaysOfWeek to be an explicit nil

### UnsetDaysOfWeek
`func (o *UpdateSnapshotPolicyParams) UnsetDaysOfWeek()`

UnsetDaysOfWeek ensures that no value is present for DaysOfWeek, not even an explicit nil
### GetNumOfDaysToKeep

`func (o *UpdateSnapshotPolicyParams) GetNumOfDaysToKeep() int64`

GetNumOfDaysToKeep returns the NumOfDaysToKeep field if non-nil, zero value otherwise.

### GetNumOfDaysToKeepOk

`func (o *UpdateSnapshotPolicyParams) GetNumOfDaysToKeepOk() (*int64, bool)`

GetNumOfDaysToKeepOk returns a tuple with the NumOfDaysToKeep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfDaysToKeep

`func (o *UpdateSnapshotPolicyParams) SetNumOfDaysToKeep(v int64)`

SetNumOfDaysToKeep sets NumOfDaysToKeep field to given value.

### HasNumOfDaysToKeep

`func (o *UpdateSnapshotPolicyParams) HasNumOfDaysToKeep() bool`

HasNumOfDaysToKeep returns a boolean if a field has been set.

### SetNumOfDaysToKeepNil

`func (o *UpdateSnapshotPolicyParams) SetNumOfDaysToKeepNil(b bool)`

 SetNumOfDaysToKeepNil sets the value for NumOfDaysToKeep to be an explicit nil

### UnsetNumOfDaysToKeep
`func (o *UpdateSnapshotPolicyParams) UnsetNumOfDaysToKeep()`

UnsetNumOfDaysToKeep ensures that no value is present for NumOfDaysToKeep, not even an explicit nil
### GetNumOfVersionsToKeep

`func (o *UpdateSnapshotPolicyParams) GetNumOfVersionsToKeep() int64`

GetNumOfVersionsToKeep returns the NumOfVersionsToKeep field if non-nil, zero value otherwise.

### GetNumOfVersionsToKeepOk

`func (o *UpdateSnapshotPolicyParams) GetNumOfVersionsToKeepOk() (*int64, bool)`

GetNumOfVersionsToKeepOk returns a tuple with the NumOfVersionsToKeep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfVersionsToKeep

`func (o *UpdateSnapshotPolicyParams) SetNumOfVersionsToKeep(v int64)`

SetNumOfVersionsToKeep sets NumOfVersionsToKeep field to given value.

### HasNumOfVersionsToKeep

`func (o *UpdateSnapshotPolicyParams) HasNumOfVersionsToKeep() bool`

HasNumOfVersionsToKeep returns a boolean if a field has been set.

### SetNumOfVersionsToKeepNil

`func (o *UpdateSnapshotPolicyParams) SetNumOfVersionsToKeepNil(b bool)`

 SetNumOfVersionsToKeepNil sets the value for NumOfVersionsToKeep to be an explicit nil

### UnsetNumOfVersionsToKeep
`func (o *UpdateSnapshotPolicyParams) UnsetNumOfVersionsToKeep()`

UnsetNumOfVersionsToKeep ensures that no value is present for NumOfVersionsToKeep, not even an explicit nil
### GetSuspendRetentionPolicy

`func (o *UpdateSnapshotPolicyParams) GetSuspendRetentionPolicy() bool`

GetSuspendRetentionPolicy returns the SuspendRetentionPolicy field if non-nil, zero value otherwise.

### GetSuspendRetentionPolicyOk

`func (o *UpdateSnapshotPolicyParams) GetSuspendRetentionPolicyOk() (*bool, bool)`

GetSuspendRetentionPolicyOk returns a tuple with the SuspendRetentionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendRetentionPolicy

`func (o *UpdateSnapshotPolicyParams) SetSuspendRetentionPolicy(v bool)`

SetSuspendRetentionPolicy sets SuspendRetentionPolicy field to given value.

### HasSuspendRetentionPolicy

`func (o *UpdateSnapshotPolicyParams) HasSuspendRetentionPolicy() bool`

HasSuspendRetentionPolicy returns a boolean if a field has been set.

### SetSuspendRetentionPolicyNil

`func (o *UpdateSnapshotPolicyParams) SetSuspendRetentionPolicyNil(b bool)`

 SetSuspendRetentionPolicyNil sets the value for SuspendRetentionPolicy to be an explicit nil

### UnsetSuspendRetentionPolicy
`func (o *UpdateSnapshotPolicyParams) UnsetSuspendRetentionPolicy()`

UnsetSuspendRetentionPolicy ensures that no value is present for SuspendRetentionPolicy, not even an explicit nil
### GetSuspendSchedulePolicy

`func (o *UpdateSnapshotPolicyParams) GetSuspendSchedulePolicy() bool`

GetSuspendSchedulePolicy returns the SuspendSchedulePolicy field if non-nil, zero value otherwise.

### GetSuspendSchedulePolicyOk

`func (o *UpdateSnapshotPolicyParams) GetSuspendSchedulePolicyOk() (*bool, bool)`

GetSuspendSchedulePolicyOk returns a tuple with the SuspendSchedulePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendSchedulePolicy

`func (o *UpdateSnapshotPolicyParams) SetSuspendSchedulePolicy(v bool)`

SetSuspendSchedulePolicy sets SuspendSchedulePolicy field to given value.

### HasSuspendSchedulePolicy

`func (o *UpdateSnapshotPolicyParams) HasSuspendSchedulePolicy() bool`

HasSuspendSchedulePolicy returns a boolean if a field has been set.

### SetSuspendSchedulePolicyNil

`func (o *UpdateSnapshotPolicyParams) SetSuspendSchedulePolicyNil(b bool)`

 SetSuspendSchedulePolicyNil sets the value for SuspendSchedulePolicy to be an explicit nil

### UnsetSuspendSchedulePolicy
`func (o *UpdateSnapshotPolicyParams) UnsetSuspendSchedulePolicy()`

UnsetSuspendSchedulePolicy ensures that no value is present for SuspendSchedulePolicy, not even an explicit nil
### GetTime

`func (o *UpdateSnapshotPolicyParams) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *UpdateSnapshotPolicyParams) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *UpdateSnapshotPolicyParams) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *UpdateSnapshotPolicyParams) HasTime() bool`

HasTime returns a boolean if a field has been set.

### SetTimeNil

`func (o *UpdateSnapshotPolicyParams) SetTimeNil(b bool)`

 SetTimeNil sets the value for Time to be an explicit nil

### UnsetTime
`func (o *UpdateSnapshotPolicyParams) UnsetTime()`

UnsetTime ensures that no value is present for Time, not even an explicit nil
### GetTimeZone

`func (o *UpdateSnapshotPolicyParams) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *UpdateSnapshotPolicyParams) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *UpdateSnapshotPolicyParams) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *UpdateSnapshotPolicyParams) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### SetTimeZoneNil

`func (o *UpdateSnapshotPolicyParams) SetTimeZoneNil(b bool)`

 SetTimeZoneNil sets the value for TimeZone to be an explicit nil

### UnsetTimeZone
`func (o *UpdateSnapshotPolicyParams) UnsetTimeZone()`

UnsetTimeZone ensures that no value is present for TimeZone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


