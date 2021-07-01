# DestroyClonedTaskStateProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionExecutorTargetType** | Pointer to **NullableInt32** | Denotes the target for action executor(Bridge/Bridge_Proxy) on which task on slave should execute actions. | [optional] 
**CloneTaskName** | Pointer to **NullableString** | The name of the clone task. | [optional] 
**DatastoreEntity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**DeployVmsToCloudTaskState** | Pointer to [**DeployVMsToCloudTaskStateProto**](DeployVMsToCloudTaskStateProto.md) |  | [optional] 
**DestroyCloneAppTaskInfo** | Pointer to [**DestroyCloneAppTaskInfoProto**](DestroyCloneAppTaskInfoProto.md) |  | [optional] 
**DestroyCloneVmTaskInfo** | Pointer to [**DestroyClonedVMTaskInfoProto**](DestroyClonedVMTaskInfoProto.md) |  | [optional] 
**DestroyMountVolumesTaskInfo** | Pointer to [**DestroyMountVolumesTaskInfoProto**](DestroyMountVolumesTaskInfoProto.md) |  | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | If the destroy clone task has finished, this field contains the end time of the task. | [optional] 
**Error** | Pointer to [**ErrorProto**](ErrorProto.md) |  | [optional] 
**FolderEntity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**ForceDelete** | Pointer to **NullableBool** | flag used to perform force delete, ignore error on delete steps | [optional] 
**FullViewName** | Pointer to **NullableString** | The full external view name where cloned objects are placed. | [optional] 
**ParentSourceConnectionParams** | Pointer to [**ConnectorParams**](ConnectorParams.md) |  | [optional] 
**ParentTaskId** | Pointer to **NullableInt64** | The id of the task that triggered the destroy task. This will be used by refresh task to mark the destroy task as internal sub-task. | [optional] 
**PerformCloneTaskId** | Pointer to **NullableInt64** | The unique id of the task that performed the clone operation. | [optional] 
**RestoreType** | Pointer to **NullableInt32** | The type of the restore/clone operation that is being destroyed. | [optional] 
**ScheduledConstituentId** | Pointer to **NullableInt64** | Constituent id (and the gandalf session id) where this task has been scheduled. If -1, the task is not running at any slave. It&#39;s possible that the task was previously scheduled, but is now being re-scheduled. | [optional] 
**ScheduledGandalfSessionId** | Pointer to **NullableInt64** |  | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | The start time of this destroy clone task. | [optional] 
**Status** | Pointer to **NullableInt32** | Status of the destroy clone task. | [optional] 
**TaskId** | Pointer to **NullableInt64** | A globally unique id of this destroy clone task. | [optional] 
**Type** | Pointer to **NullableInt32** | The type of environment that is being operated on. | [optional] 
**User** | Pointer to **NullableString** | The user who requested this destroy clone task. | [optional] 
**UserInfo** | Pointer to [**UserInformation**](UserInformation.md) |  | [optional] 
**VcdConfig** | Pointer to [**RestoredObjectVCDConfigProto**](RestoredObjectVCDConfigProto.md) |  | [optional] 
**ViewBoxId** | Pointer to **NullableInt64** | The view box id to which &#39;view_name&#39; belongs to. | [optional] 
**ViewNameDEPRECATED** | Pointer to **NullableString** | The view name as provided by the user for the clone operation. | [optional] 

## Methods

### NewDestroyClonedTaskStateProto

`func NewDestroyClonedTaskStateProto() *DestroyClonedTaskStateProto`

NewDestroyClonedTaskStateProto instantiates a new DestroyClonedTaskStateProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestroyClonedTaskStateProtoWithDefaults

`func NewDestroyClonedTaskStateProtoWithDefaults() *DestroyClonedTaskStateProto`

NewDestroyClonedTaskStateProtoWithDefaults instantiates a new DestroyClonedTaskStateProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionExecutorTargetType

`func (o *DestroyClonedTaskStateProto) GetActionExecutorTargetType() int32`

GetActionExecutorTargetType returns the ActionExecutorTargetType field if non-nil, zero value otherwise.

### GetActionExecutorTargetTypeOk

`func (o *DestroyClonedTaskStateProto) GetActionExecutorTargetTypeOk() (*int32, bool)`

GetActionExecutorTargetTypeOk returns a tuple with the ActionExecutorTargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionExecutorTargetType

`func (o *DestroyClonedTaskStateProto) SetActionExecutorTargetType(v int32)`

SetActionExecutorTargetType sets ActionExecutorTargetType field to given value.

### HasActionExecutorTargetType

`func (o *DestroyClonedTaskStateProto) HasActionExecutorTargetType() bool`

HasActionExecutorTargetType returns a boolean if a field has been set.

### SetActionExecutorTargetTypeNil

`func (o *DestroyClonedTaskStateProto) SetActionExecutorTargetTypeNil(b bool)`

 SetActionExecutorTargetTypeNil sets the value for ActionExecutorTargetType to be an explicit nil

### UnsetActionExecutorTargetType
`func (o *DestroyClonedTaskStateProto) UnsetActionExecutorTargetType()`

UnsetActionExecutorTargetType ensures that no value is present for ActionExecutorTargetType, not even an explicit nil
### GetCloneTaskName

`func (o *DestroyClonedTaskStateProto) GetCloneTaskName() string`

GetCloneTaskName returns the CloneTaskName field if non-nil, zero value otherwise.

### GetCloneTaskNameOk

`func (o *DestroyClonedTaskStateProto) GetCloneTaskNameOk() (*string, bool)`

GetCloneTaskNameOk returns a tuple with the CloneTaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneTaskName

`func (o *DestroyClonedTaskStateProto) SetCloneTaskName(v string)`

SetCloneTaskName sets CloneTaskName field to given value.

### HasCloneTaskName

`func (o *DestroyClonedTaskStateProto) HasCloneTaskName() bool`

HasCloneTaskName returns a boolean if a field has been set.

### SetCloneTaskNameNil

`func (o *DestroyClonedTaskStateProto) SetCloneTaskNameNil(b bool)`

 SetCloneTaskNameNil sets the value for CloneTaskName to be an explicit nil

### UnsetCloneTaskName
`func (o *DestroyClonedTaskStateProto) UnsetCloneTaskName()`

UnsetCloneTaskName ensures that no value is present for CloneTaskName, not even an explicit nil
### GetDatastoreEntity

`func (o *DestroyClonedTaskStateProto) GetDatastoreEntity() EntityProto`

GetDatastoreEntity returns the DatastoreEntity field if non-nil, zero value otherwise.

### GetDatastoreEntityOk

`func (o *DestroyClonedTaskStateProto) GetDatastoreEntityOk() (*EntityProto, bool)`

GetDatastoreEntityOk returns a tuple with the DatastoreEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreEntity

`func (o *DestroyClonedTaskStateProto) SetDatastoreEntity(v EntityProto)`

SetDatastoreEntity sets DatastoreEntity field to given value.

### HasDatastoreEntity

`func (o *DestroyClonedTaskStateProto) HasDatastoreEntity() bool`

HasDatastoreEntity returns a boolean if a field has been set.

### GetDeployVmsToCloudTaskState

`func (o *DestroyClonedTaskStateProto) GetDeployVmsToCloudTaskState() DeployVMsToCloudTaskStateProto`

GetDeployVmsToCloudTaskState returns the DeployVmsToCloudTaskState field if non-nil, zero value otherwise.

### GetDeployVmsToCloudTaskStateOk

`func (o *DestroyClonedTaskStateProto) GetDeployVmsToCloudTaskStateOk() (*DeployVMsToCloudTaskStateProto, bool)`

GetDeployVmsToCloudTaskStateOk returns a tuple with the DeployVmsToCloudTaskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployVmsToCloudTaskState

`func (o *DestroyClonedTaskStateProto) SetDeployVmsToCloudTaskState(v DeployVMsToCloudTaskStateProto)`

SetDeployVmsToCloudTaskState sets DeployVmsToCloudTaskState field to given value.

### HasDeployVmsToCloudTaskState

`func (o *DestroyClonedTaskStateProto) HasDeployVmsToCloudTaskState() bool`

HasDeployVmsToCloudTaskState returns a boolean if a field has been set.

### GetDestroyCloneAppTaskInfo

`func (o *DestroyClonedTaskStateProto) GetDestroyCloneAppTaskInfo() DestroyCloneAppTaskInfoProto`

GetDestroyCloneAppTaskInfo returns the DestroyCloneAppTaskInfo field if non-nil, zero value otherwise.

### GetDestroyCloneAppTaskInfoOk

`func (o *DestroyClonedTaskStateProto) GetDestroyCloneAppTaskInfoOk() (*DestroyCloneAppTaskInfoProto, bool)`

GetDestroyCloneAppTaskInfoOk returns a tuple with the DestroyCloneAppTaskInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyCloneAppTaskInfo

`func (o *DestroyClonedTaskStateProto) SetDestroyCloneAppTaskInfo(v DestroyCloneAppTaskInfoProto)`

SetDestroyCloneAppTaskInfo sets DestroyCloneAppTaskInfo field to given value.

### HasDestroyCloneAppTaskInfo

`func (o *DestroyClonedTaskStateProto) HasDestroyCloneAppTaskInfo() bool`

HasDestroyCloneAppTaskInfo returns a boolean if a field has been set.

### GetDestroyCloneVmTaskInfo

`func (o *DestroyClonedTaskStateProto) GetDestroyCloneVmTaskInfo() DestroyClonedVMTaskInfoProto`

GetDestroyCloneVmTaskInfo returns the DestroyCloneVmTaskInfo field if non-nil, zero value otherwise.

### GetDestroyCloneVmTaskInfoOk

`func (o *DestroyClonedTaskStateProto) GetDestroyCloneVmTaskInfoOk() (*DestroyClonedVMTaskInfoProto, bool)`

GetDestroyCloneVmTaskInfoOk returns a tuple with the DestroyCloneVmTaskInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyCloneVmTaskInfo

`func (o *DestroyClonedTaskStateProto) SetDestroyCloneVmTaskInfo(v DestroyClonedVMTaskInfoProto)`

SetDestroyCloneVmTaskInfo sets DestroyCloneVmTaskInfo field to given value.

### HasDestroyCloneVmTaskInfo

`func (o *DestroyClonedTaskStateProto) HasDestroyCloneVmTaskInfo() bool`

HasDestroyCloneVmTaskInfo returns a boolean if a field has been set.

### GetDestroyMountVolumesTaskInfo

`func (o *DestroyClonedTaskStateProto) GetDestroyMountVolumesTaskInfo() DestroyMountVolumesTaskInfoProto`

GetDestroyMountVolumesTaskInfo returns the DestroyMountVolumesTaskInfo field if non-nil, zero value otherwise.

### GetDestroyMountVolumesTaskInfoOk

`func (o *DestroyClonedTaskStateProto) GetDestroyMountVolumesTaskInfoOk() (*DestroyMountVolumesTaskInfoProto, bool)`

GetDestroyMountVolumesTaskInfoOk returns a tuple with the DestroyMountVolumesTaskInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyMountVolumesTaskInfo

`func (o *DestroyClonedTaskStateProto) SetDestroyMountVolumesTaskInfo(v DestroyMountVolumesTaskInfoProto)`

SetDestroyMountVolumesTaskInfo sets DestroyMountVolumesTaskInfo field to given value.

### HasDestroyMountVolumesTaskInfo

`func (o *DestroyClonedTaskStateProto) HasDestroyMountVolumesTaskInfo() bool`

HasDestroyMountVolumesTaskInfo returns a boolean if a field has been set.

### GetEndTimeUsecs

`func (o *DestroyClonedTaskStateProto) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *DestroyClonedTaskStateProto) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *DestroyClonedTaskStateProto) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *DestroyClonedTaskStateProto) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *DestroyClonedTaskStateProto) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *DestroyClonedTaskStateProto) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetError

`func (o *DestroyClonedTaskStateProto) GetError() ErrorProto`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DestroyClonedTaskStateProto) GetErrorOk() (*ErrorProto, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DestroyClonedTaskStateProto) SetError(v ErrorProto)`

SetError sets Error field to given value.

### HasError

`func (o *DestroyClonedTaskStateProto) HasError() bool`

HasError returns a boolean if a field has been set.

### GetFolderEntity

`func (o *DestroyClonedTaskStateProto) GetFolderEntity() EntityProto`

GetFolderEntity returns the FolderEntity field if non-nil, zero value otherwise.

### GetFolderEntityOk

`func (o *DestroyClonedTaskStateProto) GetFolderEntityOk() (*EntityProto, bool)`

GetFolderEntityOk returns a tuple with the FolderEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderEntity

`func (o *DestroyClonedTaskStateProto) SetFolderEntity(v EntityProto)`

SetFolderEntity sets FolderEntity field to given value.

### HasFolderEntity

`func (o *DestroyClonedTaskStateProto) HasFolderEntity() bool`

HasFolderEntity returns a boolean if a field has been set.

### GetForceDelete

`func (o *DestroyClonedTaskStateProto) GetForceDelete() bool`

GetForceDelete returns the ForceDelete field if non-nil, zero value otherwise.

### GetForceDeleteOk

`func (o *DestroyClonedTaskStateProto) GetForceDeleteOk() (*bool, bool)`

GetForceDeleteOk returns a tuple with the ForceDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceDelete

`func (o *DestroyClonedTaskStateProto) SetForceDelete(v bool)`

SetForceDelete sets ForceDelete field to given value.

### HasForceDelete

`func (o *DestroyClonedTaskStateProto) HasForceDelete() bool`

HasForceDelete returns a boolean if a field has been set.

### SetForceDeleteNil

`func (o *DestroyClonedTaskStateProto) SetForceDeleteNil(b bool)`

 SetForceDeleteNil sets the value for ForceDelete to be an explicit nil

### UnsetForceDelete
`func (o *DestroyClonedTaskStateProto) UnsetForceDelete()`

UnsetForceDelete ensures that no value is present for ForceDelete, not even an explicit nil
### GetFullViewName

`func (o *DestroyClonedTaskStateProto) GetFullViewName() string`

GetFullViewName returns the FullViewName field if non-nil, zero value otherwise.

### GetFullViewNameOk

`func (o *DestroyClonedTaskStateProto) GetFullViewNameOk() (*string, bool)`

GetFullViewNameOk returns a tuple with the FullViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullViewName

`func (o *DestroyClonedTaskStateProto) SetFullViewName(v string)`

SetFullViewName sets FullViewName field to given value.

### HasFullViewName

`func (o *DestroyClonedTaskStateProto) HasFullViewName() bool`

HasFullViewName returns a boolean if a field has been set.

### SetFullViewNameNil

`func (o *DestroyClonedTaskStateProto) SetFullViewNameNil(b bool)`

 SetFullViewNameNil sets the value for FullViewName to be an explicit nil

### UnsetFullViewName
`func (o *DestroyClonedTaskStateProto) UnsetFullViewName()`

UnsetFullViewName ensures that no value is present for FullViewName, not even an explicit nil
### GetParentSourceConnectionParams

`func (o *DestroyClonedTaskStateProto) GetParentSourceConnectionParams() ConnectorParams`

GetParentSourceConnectionParams returns the ParentSourceConnectionParams field if non-nil, zero value otherwise.

### GetParentSourceConnectionParamsOk

`func (o *DestroyClonedTaskStateProto) GetParentSourceConnectionParamsOk() (*ConnectorParams, bool)`

GetParentSourceConnectionParamsOk returns a tuple with the ParentSourceConnectionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSourceConnectionParams

`func (o *DestroyClonedTaskStateProto) SetParentSourceConnectionParams(v ConnectorParams)`

SetParentSourceConnectionParams sets ParentSourceConnectionParams field to given value.

### HasParentSourceConnectionParams

`func (o *DestroyClonedTaskStateProto) HasParentSourceConnectionParams() bool`

HasParentSourceConnectionParams returns a boolean if a field has been set.

### GetParentTaskId

`func (o *DestroyClonedTaskStateProto) GetParentTaskId() int64`

GetParentTaskId returns the ParentTaskId field if non-nil, zero value otherwise.

### GetParentTaskIdOk

`func (o *DestroyClonedTaskStateProto) GetParentTaskIdOk() (*int64, bool)`

GetParentTaskIdOk returns a tuple with the ParentTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTaskId

`func (o *DestroyClonedTaskStateProto) SetParentTaskId(v int64)`

SetParentTaskId sets ParentTaskId field to given value.

### HasParentTaskId

`func (o *DestroyClonedTaskStateProto) HasParentTaskId() bool`

HasParentTaskId returns a boolean if a field has been set.

### SetParentTaskIdNil

`func (o *DestroyClonedTaskStateProto) SetParentTaskIdNil(b bool)`

 SetParentTaskIdNil sets the value for ParentTaskId to be an explicit nil

### UnsetParentTaskId
`func (o *DestroyClonedTaskStateProto) UnsetParentTaskId()`

UnsetParentTaskId ensures that no value is present for ParentTaskId, not even an explicit nil
### GetPerformCloneTaskId

`func (o *DestroyClonedTaskStateProto) GetPerformCloneTaskId() int64`

GetPerformCloneTaskId returns the PerformCloneTaskId field if non-nil, zero value otherwise.

### GetPerformCloneTaskIdOk

`func (o *DestroyClonedTaskStateProto) GetPerformCloneTaskIdOk() (*int64, bool)`

GetPerformCloneTaskIdOk returns a tuple with the PerformCloneTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformCloneTaskId

`func (o *DestroyClonedTaskStateProto) SetPerformCloneTaskId(v int64)`

SetPerformCloneTaskId sets PerformCloneTaskId field to given value.

### HasPerformCloneTaskId

`func (o *DestroyClonedTaskStateProto) HasPerformCloneTaskId() bool`

HasPerformCloneTaskId returns a boolean if a field has been set.

### SetPerformCloneTaskIdNil

`func (o *DestroyClonedTaskStateProto) SetPerformCloneTaskIdNil(b bool)`

 SetPerformCloneTaskIdNil sets the value for PerformCloneTaskId to be an explicit nil

### UnsetPerformCloneTaskId
`func (o *DestroyClonedTaskStateProto) UnsetPerformCloneTaskId()`

UnsetPerformCloneTaskId ensures that no value is present for PerformCloneTaskId, not even an explicit nil
### GetRestoreType

`func (o *DestroyClonedTaskStateProto) GetRestoreType() int32`

GetRestoreType returns the RestoreType field if non-nil, zero value otherwise.

### GetRestoreTypeOk

`func (o *DestroyClonedTaskStateProto) GetRestoreTypeOk() (*int32, bool)`

GetRestoreTypeOk returns a tuple with the RestoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreType

`func (o *DestroyClonedTaskStateProto) SetRestoreType(v int32)`

SetRestoreType sets RestoreType field to given value.

### HasRestoreType

`func (o *DestroyClonedTaskStateProto) HasRestoreType() bool`

HasRestoreType returns a boolean if a field has been set.

### SetRestoreTypeNil

`func (o *DestroyClonedTaskStateProto) SetRestoreTypeNil(b bool)`

 SetRestoreTypeNil sets the value for RestoreType to be an explicit nil

### UnsetRestoreType
`func (o *DestroyClonedTaskStateProto) UnsetRestoreType()`

UnsetRestoreType ensures that no value is present for RestoreType, not even an explicit nil
### GetScheduledConstituentId

`func (o *DestroyClonedTaskStateProto) GetScheduledConstituentId() int64`

GetScheduledConstituentId returns the ScheduledConstituentId field if non-nil, zero value otherwise.

### GetScheduledConstituentIdOk

`func (o *DestroyClonedTaskStateProto) GetScheduledConstituentIdOk() (*int64, bool)`

GetScheduledConstituentIdOk returns a tuple with the ScheduledConstituentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledConstituentId

`func (o *DestroyClonedTaskStateProto) SetScheduledConstituentId(v int64)`

SetScheduledConstituentId sets ScheduledConstituentId field to given value.

### HasScheduledConstituentId

`func (o *DestroyClonedTaskStateProto) HasScheduledConstituentId() bool`

HasScheduledConstituentId returns a boolean if a field has been set.

### SetScheduledConstituentIdNil

`func (o *DestroyClonedTaskStateProto) SetScheduledConstituentIdNil(b bool)`

 SetScheduledConstituentIdNil sets the value for ScheduledConstituentId to be an explicit nil

### UnsetScheduledConstituentId
`func (o *DestroyClonedTaskStateProto) UnsetScheduledConstituentId()`

UnsetScheduledConstituentId ensures that no value is present for ScheduledConstituentId, not even an explicit nil
### GetScheduledGandalfSessionId

`func (o *DestroyClonedTaskStateProto) GetScheduledGandalfSessionId() int64`

GetScheduledGandalfSessionId returns the ScheduledGandalfSessionId field if non-nil, zero value otherwise.

### GetScheduledGandalfSessionIdOk

`func (o *DestroyClonedTaskStateProto) GetScheduledGandalfSessionIdOk() (*int64, bool)`

GetScheduledGandalfSessionIdOk returns a tuple with the ScheduledGandalfSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledGandalfSessionId

`func (o *DestroyClonedTaskStateProto) SetScheduledGandalfSessionId(v int64)`

SetScheduledGandalfSessionId sets ScheduledGandalfSessionId field to given value.

### HasScheduledGandalfSessionId

`func (o *DestroyClonedTaskStateProto) HasScheduledGandalfSessionId() bool`

HasScheduledGandalfSessionId returns a boolean if a field has been set.

### SetScheduledGandalfSessionIdNil

`func (o *DestroyClonedTaskStateProto) SetScheduledGandalfSessionIdNil(b bool)`

 SetScheduledGandalfSessionIdNil sets the value for ScheduledGandalfSessionId to be an explicit nil

### UnsetScheduledGandalfSessionId
`func (o *DestroyClonedTaskStateProto) UnsetScheduledGandalfSessionId()`

UnsetScheduledGandalfSessionId ensures that no value is present for ScheduledGandalfSessionId, not even an explicit nil
### GetStartTimeUsecs

`func (o *DestroyClonedTaskStateProto) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *DestroyClonedTaskStateProto) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *DestroyClonedTaskStateProto) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *DestroyClonedTaskStateProto) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *DestroyClonedTaskStateProto) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *DestroyClonedTaskStateProto) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetStatus

`func (o *DestroyClonedTaskStateProto) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DestroyClonedTaskStateProto) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DestroyClonedTaskStateProto) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DestroyClonedTaskStateProto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *DestroyClonedTaskStateProto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *DestroyClonedTaskStateProto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTaskId

`func (o *DestroyClonedTaskStateProto) GetTaskId() int64`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *DestroyClonedTaskStateProto) GetTaskIdOk() (*int64, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *DestroyClonedTaskStateProto) SetTaskId(v int64)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *DestroyClonedTaskStateProto) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskIdNil

`func (o *DestroyClonedTaskStateProto) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *DestroyClonedTaskStateProto) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil
### GetType

`func (o *DestroyClonedTaskStateProto) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DestroyClonedTaskStateProto) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DestroyClonedTaskStateProto) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *DestroyClonedTaskStateProto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *DestroyClonedTaskStateProto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *DestroyClonedTaskStateProto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUser

`func (o *DestroyClonedTaskStateProto) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *DestroyClonedTaskStateProto) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *DestroyClonedTaskStateProto) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *DestroyClonedTaskStateProto) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *DestroyClonedTaskStateProto) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *DestroyClonedTaskStateProto) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetUserInfo

`func (o *DestroyClonedTaskStateProto) GetUserInfo() UserInformation`

GetUserInfo returns the UserInfo field if non-nil, zero value otherwise.

### GetUserInfoOk

`func (o *DestroyClonedTaskStateProto) GetUserInfoOk() (*UserInformation, bool)`

GetUserInfoOk returns a tuple with the UserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfo

`func (o *DestroyClonedTaskStateProto) SetUserInfo(v UserInformation)`

SetUserInfo sets UserInfo field to given value.

### HasUserInfo

`func (o *DestroyClonedTaskStateProto) HasUserInfo() bool`

HasUserInfo returns a boolean if a field has been set.

### GetVcdConfig

`func (o *DestroyClonedTaskStateProto) GetVcdConfig() RestoredObjectVCDConfigProto`

GetVcdConfig returns the VcdConfig field if non-nil, zero value otherwise.

### GetVcdConfigOk

`func (o *DestroyClonedTaskStateProto) GetVcdConfigOk() (*RestoredObjectVCDConfigProto, bool)`

GetVcdConfigOk returns a tuple with the VcdConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcdConfig

`func (o *DestroyClonedTaskStateProto) SetVcdConfig(v RestoredObjectVCDConfigProto)`

SetVcdConfig sets VcdConfig field to given value.

### HasVcdConfig

`func (o *DestroyClonedTaskStateProto) HasVcdConfig() bool`

HasVcdConfig returns a boolean if a field has been set.

### GetViewBoxId

`func (o *DestroyClonedTaskStateProto) GetViewBoxId() int64`

GetViewBoxId returns the ViewBoxId field if non-nil, zero value otherwise.

### GetViewBoxIdOk

`func (o *DestroyClonedTaskStateProto) GetViewBoxIdOk() (*int64, bool)`

GetViewBoxIdOk returns a tuple with the ViewBoxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewBoxId

`func (o *DestroyClonedTaskStateProto) SetViewBoxId(v int64)`

SetViewBoxId sets ViewBoxId field to given value.

### HasViewBoxId

`func (o *DestroyClonedTaskStateProto) HasViewBoxId() bool`

HasViewBoxId returns a boolean if a field has been set.

### SetViewBoxIdNil

`func (o *DestroyClonedTaskStateProto) SetViewBoxIdNil(b bool)`

 SetViewBoxIdNil sets the value for ViewBoxId to be an explicit nil

### UnsetViewBoxId
`func (o *DestroyClonedTaskStateProto) UnsetViewBoxId()`

UnsetViewBoxId ensures that no value is present for ViewBoxId, not even an explicit nil
### GetViewNameDEPRECATED

`func (o *DestroyClonedTaskStateProto) GetViewNameDEPRECATED() string`

GetViewNameDEPRECATED returns the ViewNameDEPRECATED field if non-nil, zero value otherwise.

### GetViewNameDEPRECATEDOk

`func (o *DestroyClonedTaskStateProto) GetViewNameDEPRECATEDOk() (*string, bool)`

GetViewNameDEPRECATEDOk returns a tuple with the ViewNameDEPRECATED field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewNameDEPRECATED

`func (o *DestroyClonedTaskStateProto) SetViewNameDEPRECATED(v string)`

SetViewNameDEPRECATED sets ViewNameDEPRECATED field to given value.

### HasViewNameDEPRECATED

`func (o *DestroyClonedTaskStateProto) HasViewNameDEPRECATED() bool`

HasViewNameDEPRECATED returns a boolean if a field has been set.

### SetViewNameDEPRECATEDNil

`func (o *DestroyClonedTaskStateProto) SetViewNameDEPRECATEDNil(b bool)`

 SetViewNameDEPRECATEDNil sets the value for ViewNameDEPRECATED to be an explicit nil

### UnsetViewNameDEPRECATED
`func (o *DestroyClonedTaskStateProto) UnsetViewNameDEPRECATED()`

UnsetViewNameDEPRECATED ensures that no value is present for ViewNameDEPRECATED, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


