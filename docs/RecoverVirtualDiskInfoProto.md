# RecoverVirtualDiskInfoProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CleanupError** | Pointer to [**ErrorProto**](ErrorProto.md) |  | [optional] 
**DataMigrationError** | Pointer to [**ErrorProto**](ErrorProto.md) |  | [optional] 
**Error** | Pointer to [**ErrorProto**](ErrorProto.md) |  | [optional] 
**Finished** | Pointer to **NullableBool** | This will be set to true if the task is complete on the slave. | [optional] 
**InstantRecoveryFinished** | Pointer to **NullableBool** | This will be set to true once the instant recovery of the virtual disk is complete. | [optional] 
**MigrateTaskMoref** | Pointer to [**MORef**](MORef.md) |  | [optional] 
**RestoreDisksTaskInfoProto** | Pointer to [**SetupRestoreDiskTaskInfoProto**](SetupRestoreDiskTaskInfoProto.md) |  | [optional] 
**SlaveTaskStartTimeUsecs** | Pointer to **NullableInt64** | This is the timestamp at which the slave task started. | [optional] 
**TaskState** | Pointer to **NullableInt32** | The state of the task. | [optional] 
**Type** | Pointer to **NullableInt32** | The type of environment this recover virtual disk info pertains to. | [optional] 

## Methods

### NewRecoverVirtualDiskInfoProto

`func NewRecoverVirtualDiskInfoProto() *RecoverVirtualDiskInfoProto`

NewRecoverVirtualDiskInfoProto instantiates a new RecoverVirtualDiskInfoProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverVirtualDiskInfoProtoWithDefaults

`func NewRecoverVirtualDiskInfoProtoWithDefaults() *RecoverVirtualDiskInfoProto`

NewRecoverVirtualDiskInfoProtoWithDefaults instantiates a new RecoverVirtualDiskInfoProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCleanupError

`func (o *RecoverVirtualDiskInfoProto) GetCleanupError() ErrorProto`

GetCleanupError returns the CleanupError field if non-nil, zero value otherwise.

### GetCleanupErrorOk

`func (o *RecoverVirtualDiskInfoProto) GetCleanupErrorOk() (*ErrorProto, bool)`

GetCleanupErrorOk returns a tuple with the CleanupError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupError

`func (o *RecoverVirtualDiskInfoProto) SetCleanupError(v ErrorProto)`

SetCleanupError sets CleanupError field to given value.

### HasCleanupError

`func (o *RecoverVirtualDiskInfoProto) HasCleanupError() bool`

HasCleanupError returns a boolean if a field has been set.

### GetDataMigrationError

`func (o *RecoverVirtualDiskInfoProto) GetDataMigrationError() ErrorProto`

GetDataMigrationError returns the DataMigrationError field if non-nil, zero value otherwise.

### GetDataMigrationErrorOk

`func (o *RecoverVirtualDiskInfoProto) GetDataMigrationErrorOk() (*ErrorProto, bool)`

GetDataMigrationErrorOk returns a tuple with the DataMigrationError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataMigrationError

`func (o *RecoverVirtualDiskInfoProto) SetDataMigrationError(v ErrorProto)`

SetDataMigrationError sets DataMigrationError field to given value.

### HasDataMigrationError

`func (o *RecoverVirtualDiskInfoProto) HasDataMigrationError() bool`

HasDataMigrationError returns a boolean if a field has been set.

### GetError

`func (o *RecoverVirtualDiskInfoProto) GetError() ErrorProto`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RecoverVirtualDiskInfoProto) GetErrorOk() (*ErrorProto, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RecoverVirtualDiskInfoProto) SetError(v ErrorProto)`

SetError sets Error field to given value.

### HasError

`func (o *RecoverVirtualDiskInfoProto) HasError() bool`

HasError returns a boolean if a field has been set.

### GetFinished

`func (o *RecoverVirtualDiskInfoProto) GetFinished() bool`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *RecoverVirtualDiskInfoProto) GetFinishedOk() (*bool, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *RecoverVirtualDiskInfoProto) SetFinished(v bool)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *RecoverVirtualDiskInfoProto) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### SetFinishedNil

`func (o *RecoverVirtualDiskInfoProto) SetFinishedNil(b bool)`

 SetFinishedNil sets the value for Finished to be an explicit nil

### UnsetFinished
`func (o *RecoverVirtualDiskInfoProto) UnsetFinished()`

UnsetFinished ensures that no value is present for Finished, not even an explicit nil
### GetInstantRecoveryFinished

`func (o *RecoverVirtualDiskInfoProto) GetInstantRecoveryFinished() bool`

GetInstantRecoveryFinished returns the InstantRecoveryFinished field if non-nil, zero value otherwise.

### GetInstantRecoveryFinishedOk

`func (o *RecoverVirtualDiskInfoProto) GetInstantRecoveryFinishedOk() (*bool, bool)`

GetInstantRecoveryFinishedOk returns a tuple with the InstantRecoveryFinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantRecoveryFinished

`func (o *RecoverVirtualDiskInfoProto) SetInstantRecoveryFinished(v bool)`

SetInstantRecoveryFinished sets InstantRecoveryFinished field to given value.

### HasInstantRecoveryFinished

`func (o *RecoverVirtualDiskInfoProto) HasInstantRecoveryFinished() bool`

HasInstantRecoveryFinished returns a boolean if a field has been set.

### SetInstantRecoveryFinishedNil

`func (o *RecoverVirtualDiskInfoProto) SetInstantRecoveryFinishedNil(b bool)`

 SetInstantRecoveryFinishedNil sets the value for InstantRecoveryFinished to be an explicit nil

### UnsetInstantRecoveryFinished
`func (o *RecoverVirtualDiskInfoProto) UnsetInstantRecoveryFinished()`

UnsetInstantRecoveryFinished ensures that no value is present for InstantRecoveryFinished, not even an explicit nil
### GetMigrateTaskMoref

`func (o *RecoverVirtualDiskInfoProto) GetMigrateTaskMoref() MORef`

GetMigrateTaskMoref returns the MigrateTaskMoref field if non-nil, zero value otherwise.

### GetMigrateTaskMorefOk

`func (o *RecoverVirtualDiskInfoProto) GetMigrateTaskMorefOk() (*MORef, bool)`

GetMigrateTaskMorefOk returns a tuple with the MigrateTaskMoref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrateTaskMoref

`func (o *RecoverVirtualDiskInfoProto) SetMigrateTaskMoref(v MORef)`

SetMigrateTaskMoref sets MigrateTaskMoref field to given value.

### HasMigrateTaskMoref

`func (o *RecoverVirtualDiskInfoProto) HasMigrateTaskMoref() bool`

HasMigrateTaskMoref returns a boolean if a field has been set.

### GetRestoreDisksTaskInfoProto

`func (o *RecoverVirtualDiskInfoProto) GetRestoreDisksTaskInfoProto() SetupRestoreDiskTaskInfoProto`

GetRestoreDisksTaskInfoProto returns the RestoreDisksTaskInfoProto field if non-nil, zero value otherwise.

### GetRestoreDisksTaskInfoProtoOk

`func (o *RecoverVirtualDiskInfoProto) GetRestoreDisksTaskInfoProtoOk() (*SetupRestoreDiskTaskInfoProto, bool)`

GetRestoreDisksTaskInfoProtoOk returns a tuple with the RestoreDisksTaskInfoProto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreDisksTaskInfoProto

`func (o *RecoverVirtualDiskInfoProto) SetRestoreDisksTaskInfoProto(v SetupRestoreDiskTaskInfoProto)`

SetRestoreDisksTaskInfoProto sets RestoreDisksTaskInfoProto field to given value.

### HasRestoreDisksTaskInfoProto

`func (o *RecoverVirtualDiskInfoProto) HasRestoreDisksTaskInfoProto() bool`

HasRestoreDisksTaskInfoProto returns a boolean if a field has been set.

### GetSlaveTaskStartTimeUsecs

`func (o *RecoverVirtualDiskInfoProto) GetSlaveTaskStartTimeUsecs() int64`

GetSlaveTaskStartTimeUsecs returns the SlaveTaskStartTimeUsecs field if non-nil, zero value otherwise.

### GetSlaveTaskStartTimeUsecsOk

`func (o *RecoverVirtualDiskInfoProto) GetSlaveTaskStartTimeUsecsOk() (*int64, bool)`

GetSlaveTaskStartTimeUsecsOk returns a tuple with the SlaveTaskStartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaveTaskStartTimeUsecs

`func (o *RecoverVirtualDiskInfoProto) SetSlaveTaskStartTimeUsecs(v int64)`

SetSlaveTaskStartTimeUsecs sets SlaveTaskStartTimeUsecs field to given value.

### HasSlaveTaskStartTimeUsecs

`func (o *RecoverVirtualDiskInfoProto) HasSlaveTaskStartTimeUsecs() bool`

HasSlaveTaskStartTimeUsecs returns a boolean if a field has been set.

### SetSlaveTaskStartTimeUsecsNil

`func (o *RecoverVirtualDiskInfoProto) SetSlaveTaskStartTimeUsecsNil(b bool)`

 SetSlaveTaskStartTimeUsecsNil sets the value for SlaveTaskStartTimeUsecs to be an explicit nil

### UnsetSlaveTaskStartTimeUsecs
`func (o *RecoverVirtualDiskInfoProto) UnsetSlaveTaskStartTimeUsecs()`

UnsetSlaveTaskStartTimeUsecs ensures that no value is present for SlaveTaskStartTimeUsecs, not even an explicit nil
### GetTaskState

`func (o *RecoverVirtualDiskInfoProto) GetTaskState() int32`

GetTaskState returns the TaskState field if non-nil, zero value otherwise.

### GetTaskStateOk

`func (o *RecoverVirtualDiskInfoProto) GetTaskStateOk() (*int32, bool)`

GetTaskStateOk returns a tuple with the TaskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskState

`func (o *RecoverVirtualDiskInfoProto) SetTaskState(v int32)`

SetTaskState sets TaskState field to given value.

### HasTaskState

`func (o *RecoverVirtualDiskInfoProto) HasTaskState() bool`

HasTaskState returns a boolean if a field has been set.

### SetTaskStateNil

`func (o *RecoverVirtualDiskInfoProto) SetTaskStateNil(b bool)`

 SetTaskStateNil sets the value for TaskState to be an explicit nil

### UnsetTaskState
`func (o *RecoverVirtualDiskInfoProto) UnsetTaskState()`

UnsetTaskState ensures that no value is present for TaskState, not even an explicit nil
### GetType

`func (o *RecoverVirtualDiskInfoProto) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RecoverVirtualDiskInfoProto) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RecoverVirtualDiskInfoProto) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *RecoverVirtualDiskInfoProto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *RecoverVirtualDiskInfoProto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *RecoverVirtualDiskInfoProto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


