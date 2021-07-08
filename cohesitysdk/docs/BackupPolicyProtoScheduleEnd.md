# BackupPolicyProtoScheduleEnd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndAfterNumBackups** | Pointer to **NullableInt64** | If specified, the backup job will no longer be run after it has been run these many times. | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | If specified, the backup job will no longer be run after this time. | [optional] 

## Methods

### NewBackupPolicyProtoScheduleEnd

`func NewBackupPolicyProtoScheduleEnd() *BackupPolicyProtoScheduleEnd`

NewBackupPolicyProtoScheduleEnd instantiates a new BackupPolicyProtoScheduleEnd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupPolicyProtoScheduleEndWithDefaults

`func NewBackupPolicyProtoScheduleEndWithDefaults() *BackupPolicyProtoScheduleEnd`

NewBackupPolicyProtoScheduleEndWithDefaults instantiates a new BackupPolicyProtoScheduleEnd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndAfterNumBackups

`func (o *BackupPolicyProtoScheduleEnd) GetEndAfterNumBackups() int64`

GetEndAfterNumBackups returns the EndAfterNumBackups field if non-nil, zero value otherwise.

### GetEndAfterNumBackupsOk

`func (o *BackupPolicyProtoScheduleEnd) GetEndAfterNumBackupsOk() (*int64, bool)`

GetEndAfterNumBackupsOk returns a tuple with the EndAfterNumBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAfterNumBackups

`func (o *BackupPolicyProtoScheduleEnd) SetEndAfterNumBackups(v int64)`

SetEndAfterNumBackups sets EndAfterNumBackups field to given value.

### HasEndAfterNumBackups

`func (o *BackupPolicyProtoScheduleEnd) HasEndAfterNumBackups() bool`

HasEndAfterNumBackups returns a boolean if a field has been set.

### SetEndAfterNumBackupsNil

`func (o *BackupPolicyProtoScheduleEnd) SetEndAfterNumBackupsNil(b bool)`

 SetEndAfterNumBackupsNil sets the value for EndAfterNumBackups to be an explicit nil

### UnsetEndAfterNumBackups
`func (o *BackupPolicyProtoScheduleEnd) UnsetEndAfterNumBackups()`

UnsetEndAfterNumBackups ensures that no value is present for EndAfterNumBackups, not even an explicit nil
### GetEndTimeUsecs

`func (o *BackupPolicyProtoScheduleEnd) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *BackupPolicyProtoScheduleEnd) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *BackupPolicyProtoScheduleEnd) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *BackupPolicyProtoScheduleEnd) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *BackupPolicyProtoScheduleEnd) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *BackupPolicyProtoScheduleEnd) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


