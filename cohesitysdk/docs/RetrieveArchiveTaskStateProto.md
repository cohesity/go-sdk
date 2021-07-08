# RetrieveArchiveTaskStateProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchivalTarget** | Pointer to [**ArchivalTarget**](ArchivalTarget.md) |  | [optional] 
**ArchiveTaskUid** | Pointer to [**UniversalIdProto**](UniversalIdProto.md) |  | [optional] 
**BackupRunStartTimeUsecs** | Pointer to **NullableInt64** | The start time of the backup run whose corresponding archive is being retrieved. This field is just used for logging purposes. | [optional] 
**CancellationRequested** | Pointer to **NullableBool** | Whether this retrieval task has a pending cancellation request. | [optional] 
**DownloadFilesInfo** | Pointer to [**RetrieveArchiveTaskStateProtoDownloadFilesInfo**](RetrieveArchiveTaskStateProtoDownloadFilesInfo.md) |  | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | If the retrieval task has finished, this field contains the end time for the task. | [optional] 
**EntityVec** | Pointer to [**[]EntityProto**](EntityProto.md) | Information on the exact set of objects to retrieve from archive. Even if the user wanted to retrieve all objects from the archive, this field will contain all individual leaf-level objects. | [optional] 
**Error** | Pointer to [**ErrorProto**](ErrorProto.md) |  | [optional] 
**FullViewNameDEPRECATED** | Pointer to **NullableString** | The full view name (external). This is composed of a Cohesity specific prefix and the user provided view name. | [optional] 
**JobUid** | Pointer to [**UniversalIdProto**](UniversalIdProto.md) |  | [optional] 
**Name** | Pointer to **NullableString** | The name of the retrieval task. | [optional] 
**ProgressMonitorTaskPath** | Pointer to **NullableString** | The path of the progress monitor for this task. | [optional] 
**RestoreArchiveFilesInfo** | Pointer to [**RetrieveArchiveTaskStateProtoDownloadFilesInfo**](RetrieveArchiveTaskStateProtoDownloadFilesInfo.md) |  | [optional] 
**RestoreTaskId** | Pointer to **NullableInt64** | For retrieve tasks created after the 2.8 release, this will contain the id of the restore task that created this retrieve task. | [optional] 
**RetrievalInfo** | Pointer to [**RetrieveArchiveInfo**](RetrieveArchiveInfo.md) |  | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | The start time for this retrieval task. | [optional] 
**Status** | Pointer to **NullableInt32** | The status of this task. | [optional] 
**TaskUid** | Pointer to [**UniversalIdProto**](UniversalIdProto.md) |  | [optional] 
**User** | Pointer to **NullableString** | The user who requested this retrieval task. | [optional] 
**VaultRestoreParams** | Pointer to [**VaultParamsRestoreParams**](VaultParamsRestoreParams.md) |  | [optional] 
**ViewBoxId** | Pointer to **NullableInt64** | The view box id to which &#39;view_name&#39; belongs to. | [optional] 
**ViewNameDEPRECATED** | Pointer to **NullableString** | The view name as provided by the user for this retrieval task. Retrieved snapshots of the entities will be placed in this view. | [optional] 

## Methods

### NewRetrieveArchiveTaskStateProto

`func NewRetrieveArchiveTaskStateProto() *RetrieveArchiveTaskStateProto`

NewRetrieveArchiveTaskStateProto instantiates a new RetrieveArchiveTaskStateProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetrieveArchiveTaskStateProtoWithDefaults

`func NewRetrieveArchiveTaskStateProtoWithDefaults() *RetrieveArchiveTaskStateProto`

NewRetrieveArchiveTaskStateProtoWithDefaults instantiates a new RetrieveArchiveTaskStateProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchivalTarget

`func (o *RetrieveArchiveTaskStateProto) GetArchivalTarget() ArchivalTarget`

GetArchivalTarget returns the ArchivalTarget field if non-nil, zero value otherwise.

### GetArchivalTargetOk

`func (o *RetrieveArchiveTaskStateProto) GetArchivalTargetOk() (*ArchivalTarget, bool)`

GetArchivalTargetOk returns a tuple with the ArchivalTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalTarget

`func (o *RetrieveArchiveTaskStateProto) SetArchivalTarget(v ArchivalTarget)`

SetArchivalTarget sets ArchivalTarget field to given value.

### HasArchivalTarget

`func (o *RetrieveArchiveTaskStateProto) HasArchivalTarget() bool`

HasArchivalTarget returns a boolean if a field has been set.

### GetArchiveTaskUid

`func (o *RetrieveArchiveTaskStateProto) GetArchiveTaskUid() UniversalIdProto`

GetArchiveTaskUid returns the ArchiveTaskUid field if non-nil, zero value otherwise.

### GetArchiveTaskUidOk

`func (o *RetrieveArchiveTaskStateProto) GetArchiveTaskUidOk() (*UniversalIdProto, bool)`

GetArchiveTaskUidOk returns a tuple with the ArchiveTaskUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveTaskUid

`func (o *RetrieveArchiveTaskStateProto) SetArchiveTaskUid(v UniversalIdProto)`

SetArchiveTaskUid sets ArchiveTaskUid field to given value.

### HasArchiveTaskUid

`func (o *RetrieveArchiveTaskStateProto) HasArchiveTaskUid() bool`

HasArchiveTaskUid returns a boolean if a field has been set.

### GetBackupRunStartTimeUsecs

`func (o *RetrieveArchiveTaskStateProto) GetBackupRunStartTimeUsecs() int64`

GetBackupRunStartTimeUsecs returns the BackupRunStartTimeUsecs field if non-nil, zero value otherwise.

### GetBackupRunStartTimeUsecsOk

`func (o *RetrieveArchiveTaskStateProto) GetBackupRunStartTimeUsecsOk() (*int64, bool)`

GetBackupRunStartTimeUsecsOk returns a tuple with the BackupRunStartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRunStartTimeUsecs

`func (o *RetrieveArchiveTaskStateProto) SetBackupRunStartTimeUsecs(v int64)`

SetBackupRunStartTimeUsecs sets BackupRunStartTimeUsecs field to given value.

### HasBackupRunStartTimeUsecs

`func (o *RetrieveArchiveTaskStateProto) HasBackupRunStartTimeUsecs() bool`

HasBackupRunStartTimeUsecs returns a boolean if a field has been set.

### SetBackupRunStartTimeUsecsNil

`func (o *RetrieveArchiveTaskStateProto) SetBackupRunStartTimeUsecsNil(b bool)`

 SetBackupRunStartTimeUsecsNil sets the value for BackupRunStartTimeUsecs to be an explicit nil

### UnsetBackupRunStartTimeUsecs
`func (o *RetrieveArchiveTaskStateProto) UnsetBackupRunStartTimeUsecs()`

UnsetBackupRunStartTimeUsecs ensures that no value is present for BackupRunStartTimeUsecs, not even an explicit nil
### GetCancellationRequested

`func (o *RetrieveArchiveTaskStateProto) GetCancellationRequested() bool`

GetCancellationRequested returns the CancellationRequested field if non-nil, zero value otherwise.

### GetCancellationRequestedOk

`func (o *RetrieveArchiveTaskStateProto) GetCancellationRequestedOk() (*bool, bool)`

GetCancellationRequestedOk returns a tuple with the CancellationRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationRequested

`func (o *RetrieveArchiveTaskStateProto) SetCancellationRequested(v bool)`

SetCancellationRequested sets CancellationRequested field to given value.

### HasCancellationRequested

`func (o *RetrieveArchiveTaskStateProto) HasCancellationRequested() bool`

HasCancellationRequested returns a boolean if a field has been set.

### SetCancellationRequestedNil

`func (o *RetrieveArchiveTaskStateProto) SetCancellationRequestedNil(b bool)`

 SetCancellationRequestedNil sets the value for CancellationRequested to be an explicit nil

### UnsetCancellationRequested
`func (o *RetrieveArchiveTaskStateProto) UnsetCancellationRequested()`

UnsetCancellationRequested ensures that no value is present for CancellationRequested, not even an explicit nil
### GetDownloadFilesInfo

`func (o *RetrieveArchiveTaskStateProto) GetDownloadFilesInfo() RetrieveArchiveTaskStateProtoDownloadFilesInfo`

GetDownloadFilesInfo returns the DownloadFilesInfo field if non-nil, zero value otherwise.

### GetDownloadFilesInfoOk

`func (o *RetrieveArchiveTaskStateProto) GetDownloadFilesInfoOk() (*RetrieveArchiveTaskStateProtoDownloadFilesInfo, bool)`

GetDownloadFilesInfoOk returns a tuple with the DownloadFilesInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadFilesInfo

`func (o *RetrieveArchiveTaskStateProto) SetDownloadFilesInfo(v RetrieveArchiveTaskStateProtoDownloadFilesInfo)`

SetDownloadFilesInfo sets DownloadFilesInfo field to given value.

### HasDownloadFilesInfo

`func (o *RetrieveArchiveTaskStateProto) HasDownloadFilesInfo() bool`

HasDownloadFilesInfo returns a boolean if a field has been set.

### GetEndTimeUsecs

`func (o *RetrieveArchiveTaskStateProto) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *RetrieveArchiveTaskStateProto) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *RetrieveArchiveTaskStateProto) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *RetrieveArchiveTaskStateProto) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *RetrieveArchiveTaskStateProto) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *RetrieveArchiveTaskStateProto) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetEntityVec

`func (o *RetrieveArchiveTaskStateProto) GetEntityVec() []EntityProto`

GetEntityVec returns the EntityVec field if non-nil, zero value otherwise.

### GetEntityVecOk

`func (o *RetrieveArchiveTaskStateProto) GetEntityVecOk() (*[]EntityProto, bool)`

GetEntityVecOk returns a tuple with the EntityVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityVec

`func (o *RetrieveArchiveTaskStateProto) SetEntityVec(v []EntityProto)`

SetEntityVec sets EntityVec field to given value.

### HasEntityVec

`func (o *RetrieveArchiveTaskStateProto) HasEntityVec() bool`

HasEntityVec returns a boolean if a field has been set.

### SetEntityVecNil

`func (o *RetrieveArchiveTaskStateProto) SetEntityVecNil(b bool)`

 SetEntityVecNil sets the value for EntityVec to be an explicit nil

### UnsetEntityVec
`func (o *RetrieveArchiveTaskStateProto) UnsetEntityVec()`

UnsetEntityVec ensures that no value is present for EntityVec, not even an explicit nil
### GetError

`func (o *RetrieveArchiveTaskStateProto) GetError() ErrorProto`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RetrieveArchiveTaskStateProto) GetErrorOk() (*ErrorProto, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RetrieveArchiveTaskStateProto) SetError(v ErrorProto)`

SetError sets Error field to given value.

### HasError

`func (o *RetrieveArchiveTaskStateProto) HasError() bool`

HasError returns a boolean if a field has been set.

### GetFullViewNameDEPRECATED

`func (o *RetrieveArchiveTaskStateProto) GetFullViewNameDEPRECATED() string`

GetFullViewNameDEPRECATED returns the FullViewNameDEPRECATED field if non-nil, zero value otherwise.

### GetFullViewNameDEPRECATEDOk

`func (o *RetrieveArchiveTaskStateProto) GetFullViewNameDEPRECATEDOk() (*string, bool)`

GetFullViewNameDEPRECATEDOk returns a tuple with the FullViewNameDEPRECATED field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullViewNameDEPRECATED

`func (o *RetrieveArchiveTaskStateProto) SetFullViewNameDEPRECATED(v string)`

SetFullViewNameDEPRECATED sets FullViewNameDEPRECATED field to given value.

### HasFullViewNameDEPRECATED

`func (o *RetrieveArchiveTaskStateProto) HasFullViewNameDEPRECATED() bool`

HasFullViewNameDEPRECATED returns a boolean if a field has been set.

### SetFullViewNameDEPRECATEDNil

`func (o *RetrieveArchiveTaskStateProto) SetFullViewNameDEPRECATEDNil(b bool)`

 SetFullViewNameDEPRECATEDNil sets the value for FullViewNameDEPRECATED to be an explicit nil

### UnsetFullViewNameDEPRECATED
`func (o *RetrieveArchiveTaskStateProto) UnsetFullViewNameDEPRECATED()`

UnsetFullViewNameDEPRECATED ensures that no value is present for FullViewNameDEPRECATED, not even an explicit nil
### GetJobUid

`func (o *RetrieveArchiveTaskStateProto) GetJobUid() UniversalIdProto`

GetJobUid returns the JobUid field if non-nil, zero value otherwise.

### GetJobUidOk

`func (o *RetrieveArchiveTaskStateProto) GetJobUidOk() (*UniversalIdProto, bool)`

GetJobUidOk returns a tuple with the JobUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobUid

`func (o *RetrieveArchiveTaskStateProto) SetJobUid(v UniversalIdProto)`

SetJobUid sets JobUid field to given value.

### HasJobUid

`func (o *RetrieveArchiveTaskStateProto) HasJobUid() bool`

HasJobUid returns a boolean if a field has been set.

### GetName

`func (o *RetrieveArchiveTaskStateProto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RetrieveArchiveTaskStateProto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RetrieveArchiveTaskStateProto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RetrieveArchiveTaskStateProto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RetrieveArchiveTaskStateProto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RetrieveArchiveTaskStateProto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProgressMonitorTaskPath

`func (o *RetrieveArchiveTaskStateProto) GetProgressMonitorTaskPath() string`

GetProgressMonitorTaskPath returns the ProgressMonitorTaskPath field if non-nil, zero value otherwise.

### GetProgressMonitorTaskPathOk

`func (o *RetrieveArchiveTaskStateProto) GetProgressMonitorTaskPathOk() (*string, bool)`

GetProgressMonitorTaskPathOk returns a tuple with the ProgressMonitorTaskPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressMonitorTaskPath

`func (o *RetrieveArchiveTaskStateProto) SetProgressMonitorTaskPath(v string)`

SetProgressMonitorTaskPath sets ProgressMonitorTaskPath field to given value.

### HasProgressMonitorTaskPath

`func (o *RetrieveArchiveTaskStateProto) HasProgressMonitorTaskPath() bool`

HasProgressMonitorTaskPath returns a boolean if a field has been set.

### SetProgressMonitorTaskPathNil

`func (o *RetrieveArchiveTaskStateProto) SetProgressMonitorTaskPathNil(b bool)`

 SetProgressMonitorTaskPathNil sets the value for ProgressMonitorTaskPath to be an explicit nil

### UnsetProgressMonitorTaskPath
`func (o *RetrieveArchiveTaskStateProto) UnsetProgressMonitorTaskPath()`

UnsetProgressMonitorTaskPath ensures that no value is present for ProgressMonitorTaskPath, not even an explicit nil
### GetRestoreArchiveFilesInfo

`func (o *RetrieveArchiveTaskStateProto) GetRestoreArchiveFilesInfo() RetrieveArchiveTaskStateProtoDownloadFilesInfo`

GetRestoreArchiveFilesInfo returns the RestoreArchiveFilesInfo field if non-nil, zero value otherwise.

### GetRestoreArchiveFilesInfoOk

`func (o *RetrieveArchiveTaskStateProto) GetRestoreArchiveFilesInfoOk() (*RetrieveArchiveTaskStateProtoDownloadFilesInfo, bool)`

GetRestoreArchiveFilesInfoOk returns a tuple with the RestoreArchiveFilesInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreArchiveFilesInfo

`func (o *RetrieveArchiveTaskStateProto) SetRestoreArchiveFilesInfo(v RetrieveArchiveTaskStateProtoDownloadFilesInfo)`

SetRestoreArchiveFilesInfo sets RestoreArchiveFilesInfo field to given value.

### HasRestoreArchiveFilesInfo

`func (o *RetrieveArchiveTaskStateProto) HasRestoreArchiveFilesInfo() bool`

HasRestoreArchiveFilesInfo returns a boolean if a field has been set.

### GetRestoreTaskId

`func (o *RetrieveArchiveTaskStateProto) GetRestoreTaskId() int64`

GetRestoreTaskId returns the RestoreTaskId field if non-nil, zero value otherwise.

### GetRestoreTaskIdOk

`func (o *RetrieveArchiveTaskStateProto) GetRestoreTaskIdOk() (*int64, bool)`

GetRestoreTaskIdOk returns a tuple with the RestoreTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreTaskId

`func (o *RetrieveArchiveTaskStateProto) SetRestoreTaskId(v int64)`

SetRestoreTaskId sets RestoreTaskId field to given value.

### HasRestoreTaskId

`func (o *RetrieveArchiveTaskStateProto) HasRestoreTaskId() bool`

HasRestoreTaskId returns a boolean if a field has been set.

### SetRestoreTaskIdNil

`func (o *RetrieveArchiveTaskStateProto) SetRestoreTaskIdNil(b bool)`

 SetRestoreTaskIdNil sets the value for RestoreTaskId to be an explicit nil

### UnsetRestoreTaskId
`func (o *RetrieveArchiveTaskStateProto) UnsetRestoreTaskId()`

UnsetRestoreTaskId ensures that no value is present for RestoreTaskId, not even an explicit nil
### GetRetrievalInfo

`func (o *RetrieveArchiveTaskStateProto) GetRetrievalInfo() RetrieveArchiveInfo`

GetRetrievalInfo returns the RetrievalInfo field if non-nil, zero value otherwise.

### GetRetrievalInfoOk

`func (o *RetrieveArchiveTaskStateProto) GetRetrievalInfoOk() (*RetrieveArchiveInfo, bool)`

GetRetrievalInfoOk returns a tuple with the RetrievalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrievalInfo

`func (o *RetrieveArchiveTaskStateProto) SetRetrievalInfo(v RetrieveArchiveInfo)`

SetRetrievalInfo sets RetrievalInfo field to given value.

### HasRetrievalInfo

`func (o *RetrieveArchiveTaskStateProto) HasRetrievalInfo() bool`

HasRetrievalInfo returns a boolean if a field has been set.

### GetStartTimeUsecs

`func (o *RetrieveArchiveTaskStateProto) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *RetrieveArchiveTaskStateProto) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *RetrieveArchiveTaskStateProto) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *RetrieveArchiveTaskStateProto) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *RetrieveArchiveTaskStateProto) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *RetrieveArchiveTaskStateProto) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetStatus

`func (o *RetrieveArchiveTaskStateProto) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RetrieveArchiveTaskStateProto) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RetrieveArchiveTaskStateProto) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RetrieveArchiveTaskStateProto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *RetrieveArchiveTaskStateProto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *RetrieveArchiveTaskStateProto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTaskUid

`func (o *RetrieveArchiveTaskStateProto) GetTaskUid() UniversalIdProto`

GetTaskUid returns the TaskUid field if non-nil, zero value otherwise.

### GetTaskUidOk

`func (o *RetrieveArchiveTaskStateProto) GetTaskUidOk() (*UniversalIdProto, bool)`

GetTaskUidOk returns a tuple with the TaskUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskUid

`func (o *RetrieveArchiveTaskStateProto) SetTaskUid(v UniversalIdProto)`

SetTaskUid sets TaskUid field to given value.

### HasTaskUid

`func (o *RetrieveArchiveTaskStateProto) HasTaskUid() bool`

HasTaskUid returns a boolean if a field has been set.

### GetUser

`func (o *RetrieveArchiveTaskStateProto) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RetrieveArchiveTaskStateProto) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RetrieveArchiveTaskStateProto) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *RetrieveArchiveTaskStateProto) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *RetrieveArchiveTaskStateProto) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *RetrieveArchiveTaskStateProto) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetVaultRestoreParams

`func (o *RetrieveArchiveTaskStateProto) GetVaultRestoreParams() VaultParamsRestoreParams`

GetVaultRestoreParams returns the VaultRestoreParams field if non-nil, zero value otherwise.

### GetVaultRestoreParamsOk

`func (o *RetrieveArchiveTaskStateProto) GetVaultRestoreParamsOk() (*VaultParamsRestoreParams, bool)`

GetVaultRestoreParamsOk returns a tuple with the VaultRestoreParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultRestoreParams

`func (o *RetrieveArchiveTaskStateProto) SetVaultRestoreParams(v VaultParamsRestoreParams)`

SetVaultRestoreParams sets VaultRestoreParams field to given value.

### HasVaultRestoreParams

`func (o *RetrieveArchiveTaskStateProto) HasVaultRestoreParams() bool`

HasVaultRestoreParams returns a boolean if a field has been set.

### GetViewBoxId

`func (o *RetrieveArchiveTaskStateProto) GetViewBoxId() int64`

GetViewBoxId returns the ViewBoxId field if non-nil, zero value otherwise.

### GetViewBoxIdOk

`func (o *RetrieveArchiveTaskStateProto) GetViewBoxIdOk() (*int64, bool)`

GetViewBoxIdOk returns a tuple with the ViewBoxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewBoxId

`func (o *RetrieveArchiveTaskStateProto) SetViewBoxId(v int64)`

SetViewBoxId sets ViewBoxId field to given value.

### HasViewBoxId

`func (o *RetrieveArchiveTaskStateProto) HasViewBoxId() bool`

HasViewBoxId returns a boolean if a field has been set.

### SetViewBoxIdNil

`func (o *RetrieveArchiveTaskStateProto) SetViewBoxIdNil(b bool)`

 SetViewBoxIdNil sets the value for ViewBoxId to be an explicit nil

### UnsetViewBoxId
`func (o *RetrieveArchiveTaskStateProto) UnsetViewBoxId()`

UnsetViewBoxId ensures that no value is present for ViewBoxId, not even an explicit nil
### GetViewNameDEPRECATED

`func (o *RetrieveArchiveTaskStateProto) GetViewNameDEPRECATED() string`

GetViewNameDEPRECATED returns the ViewNameDEPRECATED field if non-nil, zero value otherwise.

### GetViewNameDEPRECATEDOk

`func (o *RetrieveArchiveTaskStateProto) GetViewNameDEPRECATEDOk() (*string, bool)`

GetViewNameDEPRECATEDOk returns a tuple with the ViewNameDEPRECATED field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewNameDEPRECATED

`func (o *RetrieveArchiveTaskStateProto) SetViewNameDEPRECATED(v string)`

SetViewNameDEPRECATED sets ViewNameDEPRECATED field to given value.

### HasViewNameDEPRECATED

`func (o *RetrieveArchiveTaskStateProto) HasViewNameDEPRECATED() bool`

HasViewNameDEPRECATED returns a boolean if a field has been set.

### SetViewNameDEPRECATEDNil

`func (o *RetrieveArchiveTaskStateProto) SetViewNameDEPRECATEDNil(b bool)`

 SetViewNameDEPRECATEDNil sets the value for ViewNameDEPRECATED to be an explicit nil

### UnsetViewNameDEPRECATED
`func (o *RetrieveArchiveTaskStateProto) UnsetViewNameDEPRECATED()`

UnsetViewNameDEPRECATED ensures that no value is present for ViewNameDEPRECATED, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


