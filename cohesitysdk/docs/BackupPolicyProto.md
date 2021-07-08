# BackupPolicyProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinuousSchedule** | Pointer to [**BackupPolicyProtoContinuousSchedule**](BackupPolicyProtoContinuousSchedule.md) |  | [optional] 
**DailySchedule** | Pointer to [**BackupPolicyProtoDailySchedule**](BackupPolicyProtoDailySchedule.md) |  | [optional] 
**MonthlySchedule** | Pointer to [**BackupPolicyProtoMonthlySchedule**](BackupPolicyProtoMonthlySchedule.md) |  | [optional] 
**Name** | Pointer to **NullableString** | A backup schedule can have an optional name. | [optional] 
**NumDaysToKeep** | Pointer to **NullableInt64** | Specifies how to determine the expiration time for snapshots created by a backup run. The snapshots will be marked as expiring (i.e., eligible to be garbage collected) in &#39;num_days_to_keep&#39; days from when the snapshots were created. | [optional] 
**NumRetries** | Pointer to **NullableInt32** | The number of retries to perform (for retryable errors) before giving up. | [optional] 
**OneOffSchedule** | Pointer to [**BackupPolicyProtoOneOffSchedule**](BackupPolicyProtoOneOffSchedule.md) |  | [optional] 
**Periodicity** | Pointer to **NullableInt32** | Determines how often the job should be run. | [optional] 
**RetryDelayMins** | Pointer to **NullableInt32** | The number of minutes to wait before retrying a failed job. | [optional] 
**ScheduleEnd** | Pointer to [**BackupPolicyProtoScheduleEnd**](BackupPolicyProtoScheduleEnd.md) |  | [optional] 
**StartWindowIntervalMins** | Pointer to **NullableInt32** | This field determines the amount of time (in minutes) after which a scheduled job will not be started. For example, if a job is scheduled to be run every Sunday at 5am, and this field is set to 30 minutes, but the job was unable to start by 5:30am on a Sunday due to other conflicts (say too many other jobs were already running), Magneto will not attempt to start the job until the next scheduled time (on the following Sunday). If this field is not set, the interval will be determined by the Magneto flag --magneto_master_default_start_window_interval_mins. | [optional] 
**TruncateLogs** | Pointer to **NullableBool** | Whether to truncate logs after a backup run. This is currently only relevant for full or incremental backups in a SQL environment. | [optional] 

## Methods

### NewBackupPolicyProto

`func NewBackupPolicyProto() *BackupPolicyProto`

NewBackupPolicyProto instantiates a new BackupPolicyProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupPolicyProtoWithDefaults

`func NewBackupPolicyProtoWithDefaults() *BackupPolicyProto`

NewBackupPolicyProtoWithDefaults instantiates a new BackupPolicyProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinuousSchedule

`func (o *BackupPolicyProto) GetContinuousSchedule() BackupPolicyProtoContinuousSchedule`

GetContinuousSchedule returns the ContinuousSchedule field if non-nil, zero value otherwise.

### GetContinuousScheduleOk

`func (o *BackupPolicyProto) GetContinuousScheduleOk() (*BackupPolicyProtoContinuousSchedule, bool)`

GetContinuousScheduleOk returns a tuple with the ContinuousSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuousSchedule

`func (o *BackupPolicyProto) SetContinuousSchedule(v BackupPolicyProtoContinuousSchedule)`

SetContinuousSchedule sets ContinuousSchedule field to given value.

### HasContinuousSchedule

`func (o *BackupPolicyProto) HasContinuousSchedule() bool`

HasContinuousSchedule returns a boolean if a field has been set.

### GetDailySchedule

`func (o *BackupPolicyProto) GetDailySchedule() BackupPolicyProtoDailySchedule`

GetDailySchedule returns the DailySchedule field if non-nil, zero value otherwise.

### GetDailyScheduleOk

`func (o *BackupPolicyProto) GetDailyScheduleOk() (*BackupPolicyProtoDailySchedule, bool)`

GetDailyScheduleOk returns a tuple with the DailySchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailySchedule

`func (o *BackupPolicyProto) SetDailySchedule(v BackupPolicyProtoDailySchedule)`

SetDailySchedule sets DailySchedule field to given value.

### HasDailySchedule

`func (o *BackupPolicyProto) HasDailySchedule() bool`

HasDailySchedule returns a boolean if a field has been set.

### GetMonthlySchedule

`func (o *BackupPolicyProto) GetMonthlySchedule() BackupPolicyProtoMonthlySchedule`

GetMonthlySchedule returns the MonthlySchedule field if non-nil, zero value otherwise.

### GetMonthlyScheduleOk

`func (o *BackupPolicyProto) GetMonthlyScheduleOk() (*BackupPolicyProtoMonthlySchedule, bool)`

GetMonthlyScheduleOk returns a tuple with the MonthlySchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlySchedule

`func (o *BackupPolicyProto) SetMonthlySchedule(v BackupPolicyProtoMonthlySchedule)`

SetMonthlySchedule sets MonthlySchedule field to given value.

### HasMonthlySchedule

`func (o *BackupPolicyProto) HasMonthlySchedule() bool`

HasMonthlySchedule returns a boolean if a field has been set.

### GetName

`func (o *BackupPolicyProto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BackupPolicyProto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BackupPolicyProto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BackupPolicyProto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BackupPolicyProto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BackupPolicyProto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNumDaysToKeep

`func (o *BackupPolicyProto) GetNumDaysToKeep() int64`

GetNumDaysToKeep returns the NumDaysToKeep field if non-nil, zero value otherwise.

### GetNumDaysToKeepOk

`func (o *BackupPolicyProto) GetNumDaysToKeepOk() (*int64, bool)`

GetNumDaysToKeepOk returns a tuple with the NumDaysToKeep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDaysToKeep

`func (o *BackupPolicyProto) SetNumDaysToKeep(v int64)`

SetNumDaysToKeep sets NumDaysToKeep field to given value.

### HasNumDaysToKeep

`func (o *BackupPolicyProto) HasNumDaysToKeep() bool`

HasNumDaysToKeep returns a boolean if a field has been set.

### SetNumDaysToKeepNil

`func (o *BackupPolicyProto) SetNumDaysToKeepNil(b bool)`

 SetNumDaysToKeepNil sets the value for NumDaysToKeep to be an explicit nil

### UnsetNumDaysToKeep
`func (o *BackupPolicyProto) UnsetNumDaysToKeep()`

UnsetNumDaysToKeep ensures that no value is present for NumDaysToKeep, not even an explicit nil
### GetNumRetries

`func (o *BackupPolicyProto) GetNumRetries() int32`

GetNumRetries returns the NumRetries field if non-nil, zero value otherwise.

### GetNumRetriesOk

`func (o *BackupPolicyProto) GetNumRetriesOk() (*int32, bool)`

GetNumRetriesOk returns a tuple with the NumRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRetries

`func (o *BackupPolicyProto) SetNumRetries(v int32)`

SetNumRetries sets NumRetries field to given value.

### HasNumRetries

`func (o *BackupPolicyProto) HasNumRetries() bool`

HasNumRetries returns a boolean if a field has been set.

### SetNumRetriesNil

`func (o *BackupPolicyProto) SetNumRetriesNil(b bool)`

 SetNumRetriesNil sets the value for NumRetries to be an explicit nil

### UnsetNumRetries
`func (o *BackupPolicyProto) UnsetNumRetries()`

UnsetNumRetries ensures that no value is present for NumRetries, not even an explicit nil
### GetOneOffSchedule

`func (o *BackupPolicyProto) GetOneOffSchedule() BackupPolicyProtoOneOffSchedule`

GetOneOffSchedule returns the OneOffSchedule field if non-nil, zero value otherwise.

### GetOneOffScheduleOk

`func (o *BackupPolicyProto) GetOneOffScheduleOk() (*BackupPolicyProtoOneOffSchedule, bool)`

GetOneOffScheduleOk returns a tuple with the OneOffSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneOffSchedule

`func (o *BackupPolicyProto) SetOneOffSchedule(v BackupPolicyProtoOneOffSchedule)`

SetOneOffSchedule sets OneOffSchedule field to given value.

### HasOneOffSchedule

`func (o *BackupPolicyProto) HasOneOffSchedule() bool`

HasOneOffSchedule returns a boolean if a field has been set.

### GetPeriodicity

`func (o *BackupPolicyProto) GetPeriodicity() int32`

GetPeriodicity returns the Periodicity field if non-nil, zero value otherwise.

### GetPeriodicityOk

`func (o *BackupPolicyProto) GetPeriodicityOk() (*int32, bool)`

GetPeriodicityOk returns a tuple with the Periodicity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicity

`func (o *BackupPolicyProto) SetPeriodicity(v int32)`

SetPeriodicity sets Periodicity field to given value.

### HasPeriodicity

`func (o *BackupPolicyProto) HasPeriodicity() bool`

HasPeriodicity returns a boolean if a field has been set.

### SetPeriodicityNil

`func (o *BackupPolicyProto) SetPeriodicityNil(b bool)`

 SetPeriodicityNil sets the value for Periodicity to be an explicit nil

### UnsetPeriodicity
`func (o *BackupPolicyProto) UnsetPeriodicity()`

UnsetPeriodicity ensures that no value is present for Periodicity, not even an explicit nil
### GetRetryDelayMins

`func (o *BackupPolicyProto) GetRetryDelayMins() int32`

GetRetryDelayMins returns the RetryDelayMins field if non-nil, zero value otherwise.

### GetRetryDelayMinsOk

`func (o *BackupPolicyProto) GetRetryDelayMinsOk() (*int32, bool)`

GetRetryDelayMinsOk returns a tuple with the RetryDelayMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDelayMins

`func (o *BackupPolicyProto) SetRetryDelayMins(v int32)`

SetRetryDelayMins sets RetryDelayMins field to given value.

### HasRetryDelayMins

`func (o *BackupPolicyProto) HasRetryDelayMins() bool`

HasRetryDelayMins returns a boolean if a field has been set.

### SetRetryDelayMinsNil

`func (o *BackupPolicyProto) SetRetryDelayMinsNil(b bool)`

 SetRetryDelayMinsNil sets the value for RetryDelayMins to be an explicit nil

### UnsetRetryDelayMins
`func (o *BackupPolicyProto) UnsetRetryDelayMins()`

UnsetRetryDelayMins ensures that no value is present for RetryDelayMins, not even an explicit nil
### GetScheduleEnd

`func (o *BackupPolicyProto) GetScheduleEnd() BackupPolicyProtoScheduleEnd`

GetScheduleEnd returns the ScheduleEnd field if non-nil, zero value otherwise.

### GetScheduleEndOk

`func (o *BackupPolicyProto) GetScheduleEndOk() (*BackupPolicyProtoScheduleEnd, bool)`

GetScheduleEndOk returns a tuple with the ScheduleEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleEnd

`func (o *BackupPolicyProto) SetScheduleEnd(v BackupPolicyProtoScheduleEnd)`

SetScheduleEnd sets ScheduleEnd field to given value.

### HasScheduleEnd

`func (o *BackupPolicyProto) HasScheduleEnd() bool`

HasScheduleEnd returns a boolean if a field has been set.

### GetStartWindowIntervalMins

`func (o *BackupPolicyProto) GetStartWindowIntervalMins() int32`

GetStartWindowIntervalMins returns the StartWindowIntervalMins field if non-nil, zero value otherwise.

### GetStartWindowIntervalMinsOk

`func (o *BackupPolicyProto) GetStartWindowIntervalMinsOk() (*int32, bool)`

GetStartWindowIntervalMinsOk returns a tuple with the StartWindowIntervalMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartWindowIntervalMins

`func (o *BackupPolicyProto) SetStartWindowIntervalMins(v int32)`

SetStartWindowIntervalMins sets StartWindowIntervalMins field to given value.

### HasStartWindowIntervalMins

`func (o *BackupPolicyProto) HasStartWindowIntervalMins() bool`

HasStartWindowIntervalMins returns a boolean if a field has been set.

### SetStartWindowIntervalMinsNil

`func (o *BackupPolicyProto) SetStartWindowIntervalMinsNil(b bool)`

 SetStartWindowIntervalMinsNil sets the value for StartWindowIntervalMins to be an explicit nil

### UnsetStartWindowIntervalMins
`func (o *BackupPolicyProto) UnsetStartWindowIntervalMins()`

UnsetStartWindowIntervalMins ensures that no value is present for StartWindowIntervalMins, not even an explicit nil
### GetTruncateLogs

`func (o *BackupPolicyProto) GetTruncateLogs() bool`

GetTruncateLogs returns the TruncateLogs field if non-nil, zero value otherwise.

### GetTruncateLogsOk

`func (o *BackupPolicyProto) GetTruncateLogsOk() (*bool, bool)`

GetTruncateLogsOk returns a tuple with the TruncateLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncateLogs

`func (o *BackupPolicyProto) SetTruncateLogs(v bool)`

SetTruncateLogs sets TruncateLogs field to given value.

### HasTruncateLogs

`func (o *BackupPolicyProto) HasTruncateLogs() bool`

HasTruncateLogs returns a boolean if a field has been set.

### SetTruncateLogsNil

`func (o *BackupPolicyProto) SetTruncateLogsNil(b bool)`

 SetTruncateLogsNil sets the value for TruncateLogs to be an explicit nil

### UnsetTruncateLogs
`func (o *BackupPolicyProto) UnsetTruncateLogs()`

UnsetTruncateLogs ensures that no value is present for TruncateLogs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


