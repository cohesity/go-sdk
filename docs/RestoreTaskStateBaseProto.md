# RestoreTaskStateBaseProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancellationRequested** | Pointer to **NullableBool** | Whether this task has a pending cancellation request. | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | If the restore task has finished, this field contains the end time for the task. | [optional] 
**Error** | Pointer to [**ErrorProto**](ErrorProto.md) |  | [optional] 
**Name** | Pointer to **NullableString** | The name of the restore task. | [optional] 
**ParentSourceConnectionParams** | Pointer to [**ConnectorParams**](ConnectorParams.md) |  | [optional] 
**PublicStatus** | Pointer to **NullableInt32** | Iris-facing task state. This field is stamped during the export. | [optional] 
**RefreshStatus** | Pointer to **NullableInt32** | Status of the refresh task. | [optional] 
**RestoreVlanParams** | Pointer to [**VlanParams**](VlanParams.md) |  | [optional] 
**ScheduledConstituentId** | Pointer to **NullableInt64** | Constituent id (and the gandalf session id) where this task has been scheduled. If -1, the task is not running at any slave. It&#39;s possible that the task was previously scheduled, but is now being re-scheduled. | [optional] 
**ScheduledGandalfSessionId** | Pointer to **NullableInt64** |  | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | The start time for this restore task. | [optional] 
**Status** | Pointer to **NullableInt32** | Status of the restore task. | [optional] 
**TaskId** | Pointer to **NullableInt64** | A globally unique id for this task. | [optional] 
**TotalLogicalSizeBytes** | Pointer to **NullableInt64** | Logical size of this restore task. This is the amount of data that needs to be transferred to restore the entity. | [optional] 
**TotalPhysicalSizeBytes** | Pointer to **NullableInt64** | Physical size of this restore task. This is the amount of data that was actually transferred to restore the entity. | [optional] 
**Type** | Pointer to **NullableInt32** | The type of restore being performed. | [optional] 
**User** | Pointer to **NullableString** | The user who requested this restore task. | [optional] 
**UserInfo** | Pointer to [**UserInformation**](UserInformation.md) |  | [optional] 
**UserMessages** | Pointer to **[]string** | Messages displayed to the user for this task (if any). Only valid if the status of the task is kFinished. This is used for informing the user with additional details when there is not an error. | [optional] 
**Warnings** | Pointer to [**[]ErrorProto**](ErrorProto.md) | The warnings encountered by this task (if any) during its execution. | [optional] 

## Methods

### NewRestoreTaskStateBaseProto

`func NewRestoreTaskStateBaseProto() *RestoreTaskStateBaseProto`

NewRestoreTaskStateBaseProto instantiates a new RestoreTaskStateBaseProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreTaskStateBaseProtoWithDefaults

`func NewRestoreTaskStateBaseProtoWithDefaults() *RestoreTaskStateBaseProto`

NewRestoreTaskStateBaseProtoWithDefaults instantiates a new RestoreTaskStateBaseProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancellationRequested

`func (o *RestoreTaskStateBaseProto) GetCancellationRequested() bool`

GetCancellationRequested returns the CancellationRequested field if non-nil, zero value otherwise.

### GetCancellationRequestedOk

`func (o *RestoreTaskStateBaseProto) GetCancellationRequestedOk() (*bool, bool)`

GetCancellationRequestedOk returns a tuple with the CancellationRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationRequested

`func (o *RestoreTaskStateBaseProto) SetCancellationRequested(v bool)`

SetCancellationRequested sets CancellationRequested field to given value.

### HasCancellationRequested

`func (o *RestoreTaskStateBaseProto) HasCancellationRequested() bool`

HasCancellationRequested returns a boolean if a field has been set.

### SetCancellationRequestedNil

`func (o *RestoreTaskStateBaseProto) SetCancellationRequestedNil(b bool)`

 SetCancellationRequestedNil sets the value for CancellationRequested to be an explicit nil

### UnsetCancellationRequested
`func (o *RestoreTaskStateBaseProto) UnsetCancellationRequested()`

UnsetCancellationRequested ensures that no value is present for CancellationRequested, not even an explicit nil
### GetEndTimeUsecs

`func (o *RestoreTaskStateBaseProto) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *RestoreTaskStateBaseProto) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *RestoreTaskStateBaseProto) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *RestoreTaskStateBaseProto) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *RestoreTaskStateBaseProto) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *RestoreTaskStateBaseProto) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetError

`func (o *RestoreTaskStateBaseProto) GetError() ErrorProto`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RestoreTaskStateBaseProto) GetErrorOk() (*ErrorProto, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RestoreTaskStateBaseProto) SetError(v ErrorProto)`

SetError sets Error field to given value.

### HasError

`func (o *RestoreTaskStateBaseProto) HasError() bool`

HasError returns a boolean if a field has been set.

### GetName

`func (o *RestoreTaskStateBaseProto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RestoreTaskStateBaseProto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RestoreTaskStateBaseProto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RestoreTaskStateBaseProto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RestoreTaskStateBaseProto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RestoreTaskStateBaseProto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetParentSourceConnectionParams

`func (o *RestoreTaskStateBaseProto) GetParentSourceConnectionParams() ConnectorParams`

GetParentSourceConnectionParams returns the ParentSourceConnectionParams field if non-nil, zero value otherwise.

### GetParentSourceConnectionParamsOk

`func (o *RestoreTaskStateBaseProto) GetParentSourceConnectionParamsOk() (*ConnectorParams, bool)`

GetParentSourceConnectionParamsOk returns a tuple with the ParentSourceConnectionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSourceConnectionParams

`func (o *RestoreTaskStateBaseProto) SetParentSourceConnectionParams(v ConnectorParams)`

SetParentSourceConnectionParams sets ParentSourceConnectionParams field to given value.

### HasParentSourceConnectionParams

`func (o *RestoreTaskStateBaseProto) HasParentSourceConnectionParams() bool`

HasParentSourceConnectionParams returns a boolean if a field has been set.

### GetPublicStatus

`func (o *RestoreTaskStateBaseProto) GetPublicStatus() int32`

GetPublicStatus returns the PublicStatus field if non-nil, zero value otherwise.

### GetPublicStatusOk

`func (o *RestoreTaskStateBaseProto) GetPublicStatusOk() (*int32, bool)`

GetPublicStatusOk returns a tuple with the PublicStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicStatus

`func (o *RestoreTaskStateBaseProto) SetPublicStatus(v int32)`

SetPublicStatus sets PublicStatus field to given value.

### HasPublicStatus

`func (o *RestoreTaskStateBaseProto) HasPublicStatus() bool`

HasPublicStatus returns a boolean if a field has been set.

### SetPublicStatusNil

`func (o *RestoreTaskStateBaseProto) SetPublicStatusNil(b bool)`

 SetPublicStatusNil sets the value for PublicStatus to be an explicit nil

### UnsetPublicStatus
`func (o *RestoreTaskStateBaseProto) UnsetPublicStatus()`

UnsetPublicStatus ensures that no value is present for PublicStatus, not even an explicit nil
### GetRefreshStatus

`func (o *RestoreTaskStateBaseProto) GetRefreshStatus() int32`

GetRefreshStatus returns the RefreshStatus field if non-nil, zero value otherwise.

### GetRefreshStatusOk

`func (o *RestoreTaskStateBaseProto) GetRefreshStatusOk() (*int32, bool)`

GetRefreshStatusOk returns a tuple with the RefreshStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshStatus

`func (o *RestoreTaskStateBaseProto) SetRefreshStatus(v int32)`

SetRefreshStatus sets RefreshStatus field to given value.

### HasRefreshStatus

`func (o *RestoreTaskStateBaseProto) HasRefreshStatus() bool`

HasRefreshStatus returns a boolean if a field has been set.

### SetRefreshStatusNil

`func (o *RestoreTaskStateBaseProto) SetRefreshStatusNil(b bool)`

 SetRefreshStatusNil sets the value for RefreshStatus to be an explicit nil

### UnsetRefreshStatus
`func (o *RestoreTaskStateBaseProto) UnsetRefreshStatus()`

UnsetRefreshStatus ensures that no value is present for RefreshStatus, not even an explicit nil
### GetRestoreVlanParams

`func (o *RestoreTaskStateBaseProto) GetRestoreVlanParams() VlanParams`

GetRestoreVlanParams returns the RestoreVlanParams field if non-nil, zero value otherwise.

### GetRestoreVlanParamsOk

`func (o *RestoreTaskStateBaseProto) GetRestoreVlanParamsOk() (*VlanParams, bool)`

GetRestoreVlanParamsOk returns a tuple with the RestoreVlanParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreVlanParams

`func (o *RestoreTaskStateBaseProto) SetRestoreVlanParams(v VlanParams)`

SetRestoreVlanParams sets RestoreVlanParams field to given value.

### HasRestoreVlanParams

`func (o *RestoreTaskStateBaseProto) HasRestoreVlanParams() bool`

HasRestoreVlanParams returns a boolean if a field has been set.

### GetScheduledConstituentId

`func (o *RestoreTaskStateBaseProto) GetScheduledConstituentId() int64`

GetScheduledConstituentId returns the ScheduledConstituentId field if non-nil, zero value otherwise.

### GetScheduledConstituentIdOk

`func (o *RestoreTaskStateBaseProto) GetScheduledConstituentIdOk() (*int64, bool)`

GetScheduledConstituentIdOk returns a tuple with the ScheduledConstituentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledConstituentId

`func (o *RestoreTaskStateBaseProto) SetScheduledConstituentId(v int64)`

SetScheduledConstituentId sets ScheduledConstituentId field to given value.

### HasScheduledConstituentId

`func (o *RestoreTaskStateBaseProto) HasScheduledConstituentId() bool`

HasScheduledConstituentId returns a boolean if a field has been set.

### SetScheduledConstituentIdNil

`func (o *RestoreTaskStateBaseProto) SetScheduledConstituentIdNil(b bool)`

 SetScheduledConstituentIdNil sets the value for ScheduledConstituentId to be an explicit nil

### UnsetScheduledConstituentId
`func (o *RestoreTaskStateBaseProto) UnsetScheduledConstituentId()`

UnsetScheduledConstituentId ensures that no value is present for ScheduledConstituentId, not even an explicit nil
### GetScheduledGandalfSessionId

`func (o *RestoreTaskStateBaseProto) GetScheduledGandalfSessionId() int64`

GetScheduledGandalfSessionId returns the ScheduledGandalfSessionId field if non-nil, zero value otherwise.

### GetScheduledGandalfSessionIdOk

`func (o *RestoreTaskStateBaseProto) GetScheduledGandalfSessionIdOk() (*int64, bool)`

GetScheduledGandalfSessionIdOk returns a tuple with the ScheduledGandalfSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledGandalfSessionId

`func (o *RestoreTaskStateBaseProto) SetScheduledGandalfSessionId(v int64)`

SetScheduledGandalfSessionId sets ScheduledGandalfSessionId field to given value.

### HasScheduledGandalfSessionId

`func (o *RestoreTaskStateBaseProto) HasScheduledGandalfSessionId() bool`

HasScheduledGandalfSessionId returns a boolean if a field has been set.

### SetScheduledGandalfSessionIdNil

`func (o *RestoreTaskStateBaseProto) SetScheduledGandalfSessionIdNil(b bool)`

 SetScheduledGandalfSessionIdNil sets the value for ScheduledGandalfSessionId to be an explicit nil

### UnsetScheduledGandalfSessionId
`func (o *RestoreTaskStateBaseProto) UnsetScheduledGandalfSessionId()`

UnsetScheduledGandalfSessionId ensures that no value is present for ScheduledGandalfSessionId, not even an explicit nil
### GetStartTimeUsecs

`func (o *RestoreTaskStateBaseProto) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *RestoreTaskStateBaseProto) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *RestoreTaskStateBaseProto) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *RestoreTaskStateBaseProto) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *RestoreTaskStateBaseProto) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *RestoreTaskStateBaseProto) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetStatus

`func (o *RestoreTaskStateBaseProto) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RestoreTaskStateBaseProto) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RestoreTaskStateBaseProto) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RestoreTaskStateBaseProto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *RestoreTaskStateBaseProto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *RestoreTaskStateBaseProto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTaskId

`func (o *RestoreTaskStateBaseProto) GetTaskId() int64`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *RestoreTaskStateBaseProto) GetTaskIdOk() (*int64, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *RestoreTaskStateBaseProto) SetTaskId(v int64)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *RestoreTaskStateBaseProto) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskIdNil

`func (o *RestoreTaskStateBaseProto) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *RestoreTaskStateBaseProto) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil
### GetTotalLogicalSizeBytes

`func (o *RestoreTaskStateBaseProto) GetTotalLogicalSizeBytes() int64`

GetTotalLogicalSizeBytes returns the TotalLogicalSizeBytes field if non-nil, zero value otherwise.

### GetTotalLogicalSizeBytesOk

`func (o *RestoreTaskStateBaseProto) GetTotalLogicalSizeBytesOk() (*int64, bool)`

GetTotalLogicalSizeBytesOk returns a tuple with the TotalLogicalSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLogicalSizeBytes

`func (o *RestoreTaskStateBaseProto) SetTotalLogicalSizeBytes(v int64)`

SetTotalLogicalSizeBytes sets TotalLogicalSizeBytes field to given value.

### HasTotalLogicalSizeBytes

`func (o *RestoreTaskStateBaseProto) HasTotalLogicalSizeBytes() bool`

HasTotalLogicalSizeBytes returns a boolean if a field has been set.

### SetTotalLogicalSizeBytesNil

`func (o *RestoreTaskStateBaseProto) SetTotalLogicalSizeBytesNil(b bool)`

 SetTotalLogicalSizeBytesNil sets the value for TotalLogicalSizeBytes to be an explicit nil

### UnsetTotalLogicalSizeBytes
`func (o *RestoreTaskStateBaseProto) UnsetTotalLogicalSizeBytes()`

UnsetTotalLogicalSizeBytes ensures that no value is present for TotalLogicalSizeBytes, not even an explicit nil
### GetTotalPhysicalSizeBytes

`func (o *RestoreTaskStateBaseProto) GetTotalPhysicalSizeBytes() int64`

GetTotalPhysicalSizeBytes returns the TotalPhysicalSizeBytes field if non-nil, zero value otherwise.

### GetTotalPhysicalSizeBytesOk

`func (o *RestoreTaskStateBaseProto) GetTotalPhysicalSizeBytesOk() (*int64, bool)`

GetTotalPhysicalSizeBytesOk returns a tuple with the TotalPhysicalSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPhysicalSizeBytes

`func (o *RestoreTaskStateBaseProto) SetTotalPhysicalSizeBytes(v int64)`

SetTotalPhysicalSizeBytes sets TotalPhysicalSizeBytes field to given value.

### HasTotalPhysicalSizeBytes

`func (o *RestoreTaskStateBaseProto) HasTotalPhysicalSizeBytes() bool`

HasTotalPhysicalSizeBytes returns a boolean if a field has been set.

### SetTotalPhysicalSizeBytesNil

`func (o *RestoreTaskStateBaseProto) SetTotalPhysicalSizeBytesNil(b bool)`

 SetTotalPhysicalSizeBytesNil sets the value for TotalPhysicalSizeBytes to be an explicit nil

### UnsetTotalPhysicalSizeBytes
`func (o *RestoreTaskStateBaseProto) UnsetTotalPhysicalSizeBytes()`

UnsetTotalPhysicalSizeBytes ensures that no value is present for TotalPhysicalSizeBytes, not even an explicit nil
### GetType

`func (o *RestoreTaskStateBaseProto) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RestoreTaskStateBaseProto) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RestoreTaskStateBaseProto) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *RestoreTaskStateBaseProto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *RestoreTaskStateBaseProto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *RestoreTaskStateBaseProto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUser

`func (o *RestoreTaskStateBaseProto) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RestoreTaskStateBaseProto) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RestoreTaskStateBaseProto) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *RestoreTaskStateBaseProto) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *RestoreTaskStateBaseProto) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *RestoreTaskStateBaseProto) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetUserInfo

`func (o *RestoreTaskStateBaseProto) GetUserInfo() UserInformation`

GetUserInfo returns the UserInfo field if non-nil, zero value otherwise.

### GetUserInfoOk

`func (o *RestoreTaskStateBaseProto) GetUserInfoOk() (*UserInformation, bool)`

GetUserInfoOk returns a tuple with the UserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfo

`func (o *RestoreTaskStateBaseProto) SetUserInfo(v UserInformation)`

SetUserInfo sets UserInfo field to given value.

### HasUserInfo

`func (o *RestoreTaskStateBaseProto) HasUserInfo() bool`

HasUserInfo returns a boolean if a field has been set.

### GetUserMessages

`func (o *RestoreTaskStateBaseProto) GetUserMessages() []string`

GetUserMessages returns the UserMessages field if non-nil, zero value otherwise.

### GetUserMessagesOk

`func (o *RestoreTaskStateBaseProto) GetUserMessagesOk() (*[]string, bool)`

GetUserMessagesOk returns a tuple with the UserMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMessages

`func (o *RestoreTaskStateBaseProto) SetUserMessages(v []string)`

SetUserMessages sets UserMessages field to given value.

### HasUserMessages

`func (o *RestoreTaskStateBaseProto) HasUserMessages() bool`

HasUserMessages returns a boolean if a field has been set.

### SetUserMessagesNil

`func (o *RestoreTaskStateBaseProto) SetUserMessagesNil(b bool)`

 SetUserMessagesNil sets the value for UserMessages to be an explicit nil

### UnsetUserMessages
`func (o *RestoreTaskStateBaseProto) UnsetUserMessages()`

UnsetUserMessages ensures that no value is present for UserMessages, not even an explicit nil
### GetWarnings

`func (o *RestoreTaskStateBaseProto) GetWarnings() []ErrorProto`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *RestoreTaskStateBaseProto) GetWarningsOk() (*[]ErrorProto, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *RestoreTaskStateBaseProto) SetWarnings(v []ErrorProto)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *RestoreTaskStateBaseProto) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarningsNil

`func (o *RestoreTaskStateBaseProto) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *RestoreTaskStateBaseProto) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


