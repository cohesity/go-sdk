# ContinuousSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupIntervalMins** | Pointer to **NullableInt64** | If specified, this field defines the time interval in minutes when new Job Runs are started. | [optional] 

## Methods

### NewContinuousSchedule

`func NewContinuousSchedule() *ContinuousSchedule`

NewContinuousSchedule instantiates a new ContinuousSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContinuousScheduleWithDefaults

`func NewContinuousScheduleWithDefaults() *ContinuousSchedule`

NewContinuousScheduleWithDefaults instantiates a new ContinuousSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupIntervalMins

`func (o *ContinuousSchedule) GetBackupIntervalMins() int64`

GetBackupIntervalMins returns the BackupIntervalMins field if non-nil, zero value otherwise.

### GetBackupIntervalMinsOk

`func (o *ContinuousSchedule) GetBackupIntervalMinsOk() (*int64, bool)`

GetBackupIntervalMinsOk returns a tuple with the BackupIntervalMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupIntervalMins

`func (o *ContinuousSchedule) SetBackupIntervalMins(v int64)`

SetBackupIntervalMins sets BackupIntervalMins field to given value.

### HasBackupIntervalMins

`func (o *ContinuousSchedule) HasBackupIntervalMins() bool`

HasBackupIntervalMins returns a boolean if a field has been set.

### SetBackupIntervalMinsNil

`func (o *ContinuousSchedule) SetBackupIntervalMinsNil(b bool)`

 SetBackupIntervalMinsNil sets the value for BackupIntervalMins to be an explicit nil

### UnsetBackupIntervalMins
`func (o *ContinuousSchedule) UnsetBackupIntervalMins()`

UnsetBackupIntervalMins ensures that no value is present for BackupIntervalMins, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


