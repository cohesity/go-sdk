# BackupPolicyProtoDailySchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | Pointer to **[]int32** | The days of the week backup must be performed. If no days are specified, then the backup will be performed on all days. | [optional] 
**Time** | Pointer to [**Time**](Time.md) |  | [optional] 

## Methods

### NewBackupPolicyProtoDailySchedule

`func NewBackupPolicyProtoDailySchedule() *BackupPolicyProtoDailySchedule`

NewBackupPolicyProtoDailySchedule instantiates a new BackupPolicyProtoDailySchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupPolicyProtoDailyScheduleWithDefaults

`func NewBackupPolicyProtoDailyScheduleWithDefaults() *BackupPolicyProtoDailySchedule`

NewBackupPolicyProtoDailyScheduleWithDefaults instantiates a new BackupPolicyProtoDailySchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *BackupPolicyProtoDailySchedule) GetDays() []int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *BackupPolicyProtoDailySchedule) GetDaysOk() (*[]int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *BackupPolicyProtoDailySchedule) SetDays(v []int32)`

SetDays sets Days field to given value.

### HasDays

`func (o *BackupPolicyProtoDailySchedule) HasDays() bool`

HasDays returns a boolean if a field has been set.

### SetDaysNil

`func (o *BackupPolicyProtoDailySchedule) SetDaysNil(b bool)`

 SetDaysNil sets the value for Days to be an explicit nil

### UnsetDays
`func (o *BackupPolicyProtoDailySchedule) UnsetDays()`

UnsetDays ensures that no value is present for Days, not even an explicit nil
### GetTime

`func (o *BackupPolicyProtoDailySchedule) GetTime() Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *BackupPolicyProtoDailySchedule) GetTimeOk() (*Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *BackupPolicyProtoDailySchedule) SetTime(v Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *BackupPolicyProtoDailySchedule) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


