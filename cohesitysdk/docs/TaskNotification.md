# TaskNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupTask** | Pointer to [**BackupTaskInfo**](BackupTaskInfo.md) |  | [optional] 
**BulkInstallAppTask** | Pointer to [**BulkInstallAppTaskInfo**](BulkInstallAppTaskInfo.md) |  | [optional] 
**CloneTask** | Pointer to [**CloneTaskInfo**](CloneTaskInfo.md) |  | [optional] 
**CreatedTimeSecs** | Pointer to **NullableInt64** | Timestamp at which the notification was created. | [optional] 
**Description** | Pointer to **NullableString** | Description holds the actual notification text generated for the event. | [optional] 
**Dismissed** | Pointer to **NullableBool** | Dismissed keeps track of whether a notification has been seen or not. User may choose to dismiss individual event or all notifications at once. Nil or 0 value represents false. | [optional] 
**DismissedTimeSecs** | Pointer to **NullableInt64** | Timestamp at which user dismissed this notification event. | [optional] 
**FieldMessageTask** | Pointer to [**BasicTaskInfo**](BasicTaskInfo.md) |  | [optional] 
**Id** | Pointer to **NullableString** | id identifies a user notification event uniquely. This can also be used to dismiss individual notifications. | [optional] 
**RecoveryTask** | Pointer to [**RecoveryTaskInfo**](RecoveryTaskInfo.md) |  | [optional] 
**Status** | Pointer to **NullableString** | Status of the task. Status of the task. &#39;kSuccess&#39; indicates that task completed successfully. &#39;kError&#39; indicates that task encountered errors. | [optional] 
**TaskType** | Pointer to **NullableString** | Task type denotes which type of task this notification is for. This param is used to reflect the taskType. &#39;Restore&#39; notification type is generated upon completion of Restore tasks. &#39;Clone&#39; notification type is generated upon completion of Clone tasks. &#39;BackupNow&#39; notification type is generated upon completion of Backup tasks. &#39;FieldMessage&#39; notification type is generated when field message from Cohesity support is created. &#39;bulkInstallApp&#39; notification type is generated from bulk install app | [optional] 
**Visited** | Pointer to **NullableBool** | Visited keeps track of whether a notification has been seen or not. | [optional] 
**VisitedTimeSecs** | Pointer to **NullableInt64** | Timestamp at which user visited this notification event. | [optional] 

## Methods

### NewTaskNotification

`func NewTaskNotification() *TaskNotification`

NewTaskNotification instantiates a new TaskNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskNotificationWithDefaults

`func NewTaskNotificationWithDefaults() *TaskNotification`

NewTaskNotificationWithDefaults instantiates a new TaskNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupTask

`func (o *TaskNotification) GetBackupTask() BackupTaskInfo`

GetBackupTask returns the BackupTask field if non-nil, zero value otherwise.

### GetBackupTaskOk

`func (o *TaskNotification) GetBackupTaskOk() (*BackupTaskInfo, bool)`

GetBackupTaskOk returns a tuple with the BackupTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupTask

`func (o *TaskNotification) SetBackupTask(v BackupTaskInfo)`

SetBackupTask sets BackupTask field to given value.

### HasBackupTask

`func (o *TaskNotification) HasBackupTask() bool`

HasBackupTask returns a boolean if a field has been set.

### GetBulkInstallAppTask

`func (o *TaskNotification) GetBulkInstallAppTask() BulkInstallAppTaskInfo`

GetBulkInstallAppTask returns the BulkInstallAppTask field if non-nil, zero value otherwise.

### GetBulkInstallAppTaskOk

`func (o *TaskNotification) GetBulkInstallAppTaskOk() (*BulkInstallAppTaskInfo, bool)`

GetBulkInstallAppTaskOk returns a tuple with the BulkInstallAppTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkInstallAppTask

`func (o *TaskNotification) SetBulkInstallAppTask(v BulkInstallAppTaskInfo)`

SetBulkInstallAppTask sets BulkInstallAppTask field to given value.

### HasBulkInstallAppTask

`func (o *TaskNotification) HasBulkInstallAppTask() bool`

HasBulkInstallAppTask returns a boolean if a field has been set.

### GetCloneTask

`func (o *TaskNotification) GetCloneTask() CloneTaskInfo`

GetCloneTask returns the CloneTask field if non-nil, zero value otherwise.

### GetCloneTaskOk

`func (o *TaskNotification) GetCloneTaskOk() (*CloneTaskInfo, bool)`

GetCloneTaskOk returns a tuple with the CloneTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneTask

`func (o *TaskNotification) SetCloneTask(v CloneTaskInfo)`

SetCloneTask sets CloneTask field to given value.

### HasCloneTask

`func (o *TaskNotification) HasCloneTask() bool`

HasCloneTask returns a boolean if a field has been set.

### GetCreatedTimeSecs

`func (o *TaskNotification) GetCreatedTimeSecs() int64`

GetCreatedTimeSecs returns the CreatedTimeSecs field if non-nil, zero value otherwise.

### GetCreatedTimeSecsOk

`func (o *TaskNotification) GetCreatedTimeSecsOk() (*int64, bool)`

GetCreatedTimeSecsOk returns a tuple with the CreatedTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimeSecs

`func (o *TaskNotification) SetCreatedTimeSecs(v int64)`

SetCreatedTimeSecs sets CreatedTimeSecs field to given value.

### HasCreatedTimeSecs

`func (o *TaskNotification) HasCreatedTimeSecs() bool`

HasCreatedTimeSecs returns a boolean if a field has been set.

### SetCreatedTimeSecsNil

`func (o *TaskNotification) SetCreatedTimeSecsNil(b bool)`

 SetCreatedTimeSecsNil sets the value for CreatedTimeSecs to be an explicit nil

### UnsetCreatedTimeSecs
`func (o *TaskNotification) UnsetCreatedTimeSecs()`

UnsetCreatedTimeSecs ensures that no value is present for CreatedTimeSecs, not even an explicit nil
### GetDescription

`func (o *TaskNotification) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaskNotification) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaskNotification) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TaskNotification) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TaskNotification) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TaskNotification) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDismissed

`func (o *TaskNotification) GetDismissed() bool`

GetDismissed returns the Dismissed field if non-nil, zero value otherwise.

### GetDismissedOk

`func (o *TaskNotification) GetDismissedOk() (*bool, bool)`

GetDismissedOk returns a tuple with the Dismissed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissed

`func (o *TaskNotification) SetDismissed(v bool)`

SetDismissed sets Dismissed field to given value.

### HasDismissed

`func (o *TaskNotification) HasDismissed() bool`

HasDismissed returns a boolean if a field has been set.

### SetDismissedNil

`func (o *TaskNotification) SetDismissedNil(b bool)`

 SetDismissedNil sets the value for Dismissed to be an explicit nil

### UnsetDismissed
`func (o *TaskNotification) UnsetDismissed()`

UnsetDismissed ensures that no value is present for Dismissed, not even an explicit nil
### GetDismissedTimeSecs

`func (o *TaskNotification) GetDismissedTimeSecs() int64`

GetDismissedTimeSecs returns the DismissedTimeSecs field if non-nil, zero value otherwise.

### GetDismissedTimeSecsOk

`func (o *TaskNotification) GetDismissedTimeSecsOk() (*int64, bool)`

GetDismissedTimeSecsOk returns a tuple with the DismissedTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissedTimeSecs

`func (o *TaskNotification) SetDismissedTimeSecs(v int64)`

SetDismissedTimeSecs sets DismissedTimeSecs field to given value.

### HasDismissedTimeSecs

`func (o *TaskNotification) HasDismissedTimeSecs() bool`

HasDismissedTimeSecs returns a boolean if a field has been set.

### SetDismissedTimeSecsNil

`func (o *TaskNotification) SetDismissedTimeSecsNil(b bool)`

 SetDismissedTimeSecsNil sets the value for DismissedTimeSecs to be an explicit nil

### UnsetDismissedTimeSecs
`func (o *TaskNotification) UnsetDismissedTimeSecs()`

UnsetDismissedTimeSecs ensures that no value is present for DismissedTimeSecs, not even an explicit nil
### GetFieldMessageTask

`func (o *TaskNotification) GetFieldMessageTask() BasicTaskInfo`

GetFieldMessageTask returns the FieldMessageTask field if non-nil, zero value otherwise.

### GetFieldMessageTaskOk

`func (o *TaskNotification) GetFieldMessageTaskOk() (*BasicTaskInfo, bool)`

GetFieldMessageTaskOk returns a tuple with the FieldMessageTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldMessageTask

`func (o *TaskNotification) SetFieldMessageTask(v BasicTaskInfo)`

SetFieldMessageTask sets FieldMessageTask field to given value.

### HasFieldMessageTask

`func (o *TaskNotification) HasFieldMessageTask() bool`

HasFieldMessageTask returns a boolean if a field has been set.

### GetId

`func (o *TaskNotification) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskNotification) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskNotification) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaskNotification) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TaskNotification) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TaskNotification) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetRecoveryTask

`func (o *TaskNotification) GetRecoveryTask() RecoveryTaskInfo`

GetRecoveryTask returns the RecoveryTask field if non-nil, zero value otherwise.

### GetRecoveryTaskOk

`func (o *TaskNotification) GetRecoveryTaskOk() (*RecoveryTaskInfo, bool)`

GetRecoveryTaskOk returns a tuple with the RecoveryTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTask

`func (o *TaskNotification) SetRecoveryTask(v RecoveryTaskInfo)`

SetRecoveryTask sets RecoveryTask field to given value.

### HasRecoveryTask

`func (o *TaskNotification) HasRecoveryTask() bool`

HasRecoveryTask returns a boolean if a field has been set.

### GetStatus

`func (o *TaskNotification) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskNotification) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskNotification) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TaskNotification) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TaskNotification) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TaskNotification) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTaskType

`func (o *TaskNotification) GetTaskType() string`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *TaskNotification) GetTaskTypeOk() (*string, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *TaskNotification) SetTaskType(v string)`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *TaskNotification) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.

### SetTaskTypeNil

`func (o *TaskNotification) SetTaskTypeNil(b bool)`

 SetTaskTypeNil sets the value for TaskType to be an explicit nil

### UnsetTaskType
`func (o *TaskNotification) UnsetTaskType()`

UnsetTaskType ensures that no value is present for TaskType, not even an explicit nil
### GetVisited

`func (o *TaskNotification) GetVisited() bool`

GetVisited returns the Visited field if non-nil, zero value otherwise.

### GetVisitedOk

`func (o *TaskNotification) GetVisitedOk() (*bool, bool)`

GetVisitedOk returns a tuple with the Visited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisited

`func (o *TaskNotification) SetVisited(v bool)`

SetVisited sets Visited field to given value.

### HasVisited

`func (o *TaskNotification) HasVisited() bool`

HasVisited returns a boolean if a field has been set.

### SetVisitedNil

`func (o *TaskNotification) SetVisitedNil(b bool)`

 SetVisitedNil sets the value for Visited to be an explicit nil

### UnsetVisited
`func (o *TaskNotification) UnsetVisited()`

UnsetVisited ensures that no value is present for Visited, not even an explicit nil
### GetVisitedTimeSecs

`func (o *TaskNotification) GetVisitedTimeSecs() int64`

GetVisitedTimeSecs returns the VisitedTimeSecs field if non-nil, zero value otherwise.

### GetVisitedTimeSecsOk

`func (o *TaskNotification) GetVisitedTimeSecsOk() (*int64, bool)`

GetVisitedTimeSecsOk returns a tuple with the VisitedTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitedTimeSecs

`func (o *TaskNotification) SetVisitedTimeSecs(v int64)`

SetVisitedTimeSecs sets VisitedTimeSecs field to given value.

### HasVisitedTimeSecs

`func (o *TaskNotification) HasVisitedTimeSecs() bool`

HasVisitedTimeSecs returns a boolean if a field has been set.

### SetVisitedTimeSecsNil

`func (o *TaskNotification) SetVisitedTimeSecsNil(b bool)`

 SetVisitedTimeSecsNil sets the value for VisitedTimeSecs to be an explicit nil

### UnsetVisitedTimeSecs
`func (o *TaskNotification) UnsetVisitedTimeSecs()`

UnsetVisitedTimeSecs ensures that no value is present for VisitedTimeSecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


