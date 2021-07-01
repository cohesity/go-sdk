# BackupPolicyProtoContinuousSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupIntervalMins** | Pointer to **NullableInt64** | If this field is set, backups will be performed periodically every &#39;interval_mins&#39; number of minutes. NOTE: This is the interval between the start time of two successive backups. | [optional] 
**ExclusionRanges** | Pointer to [**[]BackupPolicyProtoExclusionTimeRange**](BackupPolicyProtoExclusionTimeRange.md) | Do not start backups in these time-ranges. It&#39;s possible for a previously started backup to spill over into an exclusion range.  NOTE: This field has been deprecated. Use the exclusion_ranges field defined within BackupPolicyProto instead. | [optional] 

## Methods

### NewBackupPolicyProtoContinuousSchedule

`func NewBackupPolicyProtoContinuousSchedule() *BackupPolicyProtoContinuousSchedule`

NewBackupPolicyProtoContinuousSchedule instantiates a new BackupPolicyProtoContinuousSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupPolicyProtoContinuousScheduleWithDefaults

`func NewBackupPolicyProtoContinuousScheduleWithDefaults() *BackupPolicyProtoContinuousSchedule`

NewBackupPolicyProtoContinuousScheduleWithDefaults instantiates a new BackupPolicyProtoContinuousSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupIntervalMins

`func (o *BackupPolicyProtoContinuousSchedule) GetBackupIntervalMins() int64`

GetBackupIntervalMins returns the BackupIntervalMins field if non-nil, zero value otherwise.

### GetBackupIntervalMinsOk

`func (o *BackupPolicyProtoContinuousSchedule) GetBackupIntervalMinsOk() (*int64, bool)`

GetBackupIntervalMinsOk returns a tuple with the BackupIntervalMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupIntervalMins

`func (o *BackupPolicyProtoContinuousSchedule) SetBackupIntervalMins(v int64)`

SetBackupIntervalMins sets BackupIntervalMins field to given value.

### HasBackupIntervalMins

`func (o *BackupPolicyProtoContinuousSchedule) HasBackupIntervalMins() bool`

HasBackupIntervalMins returns a boolean if a field has been set.

### SetBackupIntervalMinsNil

`func (o *BackupPolicyProtoContinuousSchedule) SetBackupIntervalMinsNil(b bool)`

 SetBackupIntervalMinsNil sets the value for BackupIntervalMins to be an explicit nil

### UnsetBackupIntervalMins
`func (o *BackupPolicyProtoContinuousSchedule) UnsetBackupIntervalMins()`

UnsetBackupIntervalMins ensures that no value is present for BackupIntervalMins, not even an explicit nil
### GetExclusionRanges

`func (o *BackupPolicyProtoContinuousSchedule) GetExclusionRanges() []BackupPolicyProtoExclusionTimeRange`

GetExclusionRanges returns the ExclusionRanges field if non-nil, zero value otherwise.

### GetExclusionRangesOk

`func (o *BackupPolicyProtoContinuousSchedule) GetExclusionRangesOk() (*[]BackupPolicyProtoExclusionTimeRange, bool)`

GetExclusionRangesOk returns a tuple with the ExclusionRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusionRanges

`func (o *BackupPolicyProtoContinuousSchedule) SetExclusionRanges(v []BackupPolicyProtoExclusionTimeRange)`

SetExclusionRanges sets ExclusionRanges field to given value.

### HasExclusionRanges

`func (o *BackupPolicyProtoContinuousSchedule) HasExclusionRanges() bool`

HasExclusionRanges returns a boolean if a field has been set.

### SetExclusionRangesNil

`func (o *BackupPolicyProtoContinuousSchedule) SetExclusionRangesNil(b bool)`

 SetExclusionRangesNil sets the value for ExclusionRanges to be an explicit nil

### UnsetExclusionRanges
`func (o *BackupPolicyProtoContinuousSchedule) UnsetExclusionRanges()`

UnsetExclusionRanges ensures that no value is present for ExclusionRanges, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


