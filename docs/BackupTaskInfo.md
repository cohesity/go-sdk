# BackupTaskInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | Pointer to **NullableString** | Id of that particular instance | [optional] 
**Name** | Pointer to **NullableString** | Name of the recovery task. | [optional] 
**StartTimeUsecs** | Pointer to **NullableString** | Denotes the start time of the backuptask, needed for deeplinking. | [optional] 
**TaskId** | Pointer to **NullableString** | Id of the recovery task. | [optional] 

## Methods

### NewBackupTaskInfo

`func NewBackupTaskInfo() *BackupTaskInfo`

NewBackupTaskInfo instantiates a new BackupTaskInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupTaskInfoWithDefaults

`func NewBackupTaskInfoWithDefaults() *BackupTaskInfo`

NewBackupTaskInfoWithDefaults instantiates a new BackupTaskInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceId

`func (o *BackupTaskInfo) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *BackupTaskInfo) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *BackupTaskInfo) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *BackupTaskInfo) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### SetInstanceIdNil

`func (o *BackupTaskInfo) SetInstanceIdNil(b bool)`

 SetInstanceIdNil sets the value for InstanceId to be an explicit nil

### UnsetInstanceId
`func (o *BackupTaskInfo) UnsetInstanceId()`

UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil
### GetName

`func (o *BackupTaskInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BackupTaskInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BackupTaskInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BackupTaskInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BackupTaskInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BackupTaskInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStartTimeUsecs

`func (o *BackupTaskInfo) GetStartTimeUsecs() string`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *BackupTaskInfo) GetStartTimeUsecsOk() (*string, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *BackupTaskInfo) SetStartTimeUsecs(v string)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *BackupTaskInfo) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *BackupTaskInfo) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *BackupTaskInfo) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetTaskId

`func (o *BackupTaskInfo) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *BackupTaskInfo) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *BackupTaskInfo) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *BackupTaskInfo) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskIdNil

`func (o *BackupTaskInfo) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *BackupTaskInfo) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


