# CommonRecoveryResponseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Specifies the id of the Recovery. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the Recovery. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of the Recovery in Unix timestamp epoch in microseconds. | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time of the Recovery in Unix timestamp epoch in microseconds. This field will be populated only after Recovery is finished. | [optional] 
**Status** | Pointer to **NullableString** | Status of the Recovery. &#39;Running&#39; indicates that the Recovery is still running. &#39;Canceled&#39; indicates that the Recovery has been cancelled. &#39;Canceling&#39; indicates that the Recovery is in the process of being cancelled. &#39;Failed&#39; indicates that the Recovery has failed. &#39;Succeeded&#39; indicates that the Recovery has finished successfully. &#39;SucceededWithWarning&#39; indicates that the Recovery finished successfully, but there were some warning messages. | [optional] 
**ProgressTaskId** | Pointer to **NullableString** | Progress monitor task id for Recovery. | [optional] 
**SnapshotEnvironment** | Pointer to **string** | Specifies the type of snapshot environment for which the Recovery was performed. | [optional] 
**RecoveryAction** | Pointer to **string** | Specifies the type of recover action. | [optional] 
**Permissions** | Pointer to [**[]Tenant**](Tenant.md) | Specifies the list of tenants that have permissions for this recovery. | [optional] 
**CreationInfo** | Pointer to [**CreationInfo**](CreationInfo.md) |  | [optional] 
**CanTearDown** | Pointer to **NullableBool** | Specifies whether it&#39;s possible to tear down the objects created by the recovery. | [optional] 
**TearDownStatus** | Pointer to **NullableString** | Specifies the status of the tear down operation. This is only set when the canTearDown is set to true. &#39;DestroyScheduled&#39; indicates that the tear down is ready to schedule. &#39;Destroying&#39; indicates that the tear down is still running. &#39;Destroyed&#39; indicates that the tear down succeeded. &#39;DestroyError&#39; indicates that the tear down failed. | [optional] 
**TearDownMessage** | Pointer to **NullableString** | Specifies the error message about the tear down operation if it fails. | [optional] 
**Messages** | Pointer to **[]string** | Specifies messages about the recovery. | [optional] 
**IsParentRecovery** | Pointer to **NullableBool** | Specifies whether the current recovery operation has created child recoveries. This is currently used in SQL recovery where multiple child recoveries can be tracked under a common/parent recovery. | [optional] 
**ParentRecoveryId** | Pointer to **NullableString** | If current recovery is child recovery triggered by another parent recovery operation, then this field willt specify the id of the parent recovery. | [optional] 
**RetrieveArchiveTasks** | Pointer to [**[]RetrieveArchiveTask**](RetrieveArchiveTask.md) | Specifies the list of persistent state of a retrieve of an archive task. | [optional] 

## Methods

### NewCommonRecoveryResponseParams

`func NewCommonRecoveryResponseParams() *CommonRecoveryResponseParams`

NewCommonRecoveryResponseParams instantiates a new CommonRecoveryResponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonRecoveryResponseParamsWithDefaults

`func NewCommonRecoveryResponseParamsWithDefaults() *CommonRecoveryResponseParams`

NewCommonRecoveryResponseParamsWithDefaults instantiates a new CommonRecoveryResponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommonRecoveryResponseParams) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommonRecoveryResponseParams) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommonRecoveryResponseParams) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CommonRecoveryResponseParams) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CommonRecoveryResponseParams) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CommonRecoveryResponseParams) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *CommonRecoveryResponseParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommonRecoveryResponseParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommonRecoveryResponseParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CommonRecoveryResponseParams) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CommonRecoveryResponseParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CommonRecoveryResponseParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStartTimeUsecs

`func (o *CommonRecoveryResponseParams) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *CommonRecoveryResponseParams) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *CommonRecoveryResponseParams) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *CommonRecoveryResponseParams) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *CommonRecoveryResponseParams) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *CommonRecoveryResponseParams) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetEndTimeUsecs

`func (o *CommonRecoveryResponseParams) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *CommonRecoveryResponseParams) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *CommonRecoveryResponseParams) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *CommonRecoveryResponseParams) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *CommonRecoveryResponseParams) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *CommonRecoveryResponseParams) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetStatus

`func (o *CommonRecoveryResponseParams) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CommonRecoveryResponseParams) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CommonRecoveryResponseParams) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CommonRecoveryResponseParams) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *CommonRecoveryResponseParams) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *CommonRecoveryResponseParams) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetProgressTaskId

`func (o *CommonRecoveryResponseParams) GetProgressTaskId() string`

GetProgressTaskId returns the ProgressTaskId field if non-nil, zero value otherwise.

### GetProgressTaskIdOk

`func (o *CommonRecoveryResponseParams) GetProgressTaskIdOk() (*string, bool)`

GetProgressTaskIdOk returns a tuple with the ProgressTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressTaskId

`func (o *CommonRecoveryResponseParams) SetProgressTaskId(v string)`

SetProgressTaskId sets ProgressTaskId field to given value.

### HasProgressTaskId

`func (o *CommonRecoveryResponseParams) HasProgressTaskId() bool`

HasProgressTaskId returns a boolean if a field has been set.

### SetProgressTaskIdNil

`func (o *CommonRecoveryResponseParams) SetProgressTaskIdNil(b bool)`

 SetProgressTaskIdNil sets the value for ProgressTaskId to be an explicit nil

### UnsetProgressTaskId
`func (o *CommonRecoveryResponseParams) UnsetProgressTaskId()`

UnsetProgressTaskId ensures that no value is present for ProgressTaskId, not even an explicit nil
### GetSnapshotEnvironment

`func (o *CommonRecoveryResponseParams) GetSnapshotEnvironment() string`

GetSnapshotEnvironment returns the SnapshotEnvironment field if non-nil, zero value otherwise.

### GetSnapshotEnvironmentOk

`func (o *CommonRecoveryResponseParams) GetSnapshotEnvironmentOk() (*string, bool)`

GetSnapshotEnvironmentOk returns a tuple with the SnapshotEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotEnvironment

`func (o *CommonRecoveryResponseParams) SetSnapshotEnvironment(v string)`

SetSnapshotEnvironment sets SnapshotEnvironment field to given value.

### HasSnapshotEnvironment

`func (o *CommonRecoveryResponseParams) HasSnapshotEnvironment() bool`

HasSnapshotEnvironment returns a boolean if a field has been set.

### GetRecoveryAction

`func (o *CommonRecoveryResponseParams) GetRecoveryAction() string`

GetRecoveryAction returns the RecoveryAction field if non-nil, zero value otherwise.

### GetRecoveryActionOk

`func (o *CommonRecoveryResponseParams) GetRecoveryActionOk() (*string, bool)`

GetRecoveryActionOk returns a tuple with the RecoveryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryAction

`func (o *CommonRecoveryResponseParams) SetRecoveryAction(v string)`

SetRecoveryAction sets RecoveryAction field to given value.

### HasRecoveryAction

`func (o *CommonRecoveryResponseParams) HasRecoveryAction() bool`

HasRecoveryAction returns a boolean if a field has been set.

### GetPermissions

`func (o *CommonRecoveryResponseParams) GetPermissions() []Tenant`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CommonRecoveryResponseParams) GetPermissionsOk() (*[]Tenant, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CommonRecoveryResponseParams) SetPermissions(v []Tenant)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CommonRecoveryResponseParams) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *CommonRecoveryResponseParams) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *CommonRecoveryResponseParams) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetCreationInfo

`func (o *CommonRecoveryResponseParams) GetCreationInfo() CreationInfo`

GetCreationInfo returns the CreationInfo field if non-nil, zero value otherwise.

### GetCreationInfoOk

`func (o *CommonRecoveryResponseParams) GetCreationInfoOk() (*CreationInfo, bool)`

GetCreationInfoOk returns a tuple with the CreationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationInfo

`func (o *CommonRecoveryResponseParams) SetCreationInfo(v CreationInfo)`

SetCreationInfo sets CreationInfo field to given value.

### HasCreationInfo

`func (o *CommonRecoveryResponseParams) HasCreationInfo() bool`

HasCreationInfo returns a boolean if a field has been set.

### GetCanTearDown

`func (o *CommonRecoveryResponseParams) GetCanTearDown() bool`

GetCanTearDown returns the CanTearDown field if non-nil, zero value otherwise.

### GetCanTearDownOk

`func (o *CommonRecoveryResponseParams) GetCanTearDownOk() (*bool, bool)`

GetCanTearDownOk returns a tuple with the CanTearDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTearDown

`func (o *CommonRecoveryResponseParams) SetCanTearDown(v bool)`

SetCanTearDown sets CanTearDown field to given value.

### HasCanTearDown

`func (o *CommonRecoveryResponseParams) HasCanTearDown() bool`

HasCanTearDown returns a boolean if a field has been set.

### SetCanTearDownNil

`func (o *CommonRecoveryResponseParams) SetCanTearDownNil(b bool)`

 SetCanTearDownNil sets the value for CanTearDown to be an explicit nil

### UnsetCanTearDown
`func (o *CommonRecoveryResponseParams) UnsetCanTearDown()`

UnsetCanTearDown ensures that no value is present for CanTearDown, not even an explicit nil
### GetTearDownStatus

`func (o *CommonRecoveryResponseParams) GetTearDownStatus() string`

GetTearDownStatus returns the TearDownStatus field if non-nil, zero value otherwise.

### GetTearDownStatusOk

`func (o *CommonRecoveryResponseParams) GetTearDownStatusOk() (*string, bool)`

GetTearDownStatusOk returns a tuple with the TearDownStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTearDownStatus

`func (o *CommonRecoveryResponseParams) SetTearDownStatus(v string)`

SetTearDownStatus sets TearDownStatus field to given value.

### HasTearDownStatus

`func (o *CommonRecoveryResponseParams) HasTearDownStatus() bool`

HasTearDownStatus returns a boolean if a field has been set.

### SetTearDownStatusNil

`func (o *CommonRecoveryResponseParams) SetTearDownStatusNil(b bool)`

 SetTearDownStatusNil sets the value for TearDownStatus to be an explicit nil

### UnsetTearDownStatus
`func (o *CommonRecoveryResponseParams) UnsetTearDownStatus()`

UnsetTearDownStatus ensures that no value is present for TearDownStatus, not even an explicit nil
### GetTearDownMessage

`func (o *CommonRecoveryResponseParams) GetTearDownMessage() string`

GetTearDownMessage returns the TearDownMessage field if non-nil, zero value otherwise.

### GetTearDownMessageOk

`func (o *CommonRecoveryResponseParams) GetTearDownMessageOk() (*string, bool)`

GetTearDownMessageOk returns a tuple with the TearDownMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTearDownMessage

`func (o *CommonRecoveryResponseParams) SetTearDownMessage(v string)`

SetTearDownMessage sets TearDownMessage field to given value.

### HasTearDownMessage

`func (o *CommonRecoveryResponseParams) HasTearDownMessage() bool`

HasTearDownMessage returns a boolean if a field has been set.

### SetTearDownMessageNil

`func (o *CommonRecoveryResponseParams) SetTearDownMessageNil(b bool)`

 SetTearDownMessageNil sets the value for TearDownMessage to be an explicit nil

### UnsetTearDownMessage
`func (o *CommonRecoveryResponseParams) UnsetTearDownMessage()`

UnsetTearDownMessage ensures that no value is present for TearDownMessage, not even an explicit nil
### GetMessages

`func (o *CommonRecoveryResponseParams) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *CommonRecoveryResponseParams) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *CommonRecoveryResponseParams) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *CommonRecoveryResponseParams) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### SetMessagesNil

`func (o *CommonRecoveryResponseParams) SetMessagesNil(b bool)`

 SetMessagesNil sets the value for Messages to be an explicit nil

### UnsetMessages
`func (o *CommonRecoveryResponseParams) UnsetMessages()`

UnsetMessages ensures that no value is present for Messages, not even an explicit nil
### GetIsParentRecovery

`func (o *CommonRecoveryResponseParams) GetIsParentRecovery() bool`

GetIsParentRecovery returns the IsParentRecovery field if non-nil, zero value otherwise.

### GetIsParentRecoveryOk

`func (o *CommonRecoveryResponseParams) GetIsParentRecoveryOk() (*bool, bool)`

GetIsParentRecoveryOk returns a tuple with the IsParentRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsParentRecovery

`func (o *CommonRecoveryResponseParams) SetIsParentRecovery(v bool)`

SetIsParentRecovery sets IsParentRecovery field to given value.

### HasIsParentRecovery

`func (o *CommonRecoveryResponseParams) HasIsParentRecovery() bool`

HasIsParentRecovery returns a boolean if a field has been set.

### SetIsParentRecoveryNil

`func (o *CommonRecoveryResponseParams) SetIsParentRecoveryNil(b bool)`

 SetIsParentRecoveryNil sets the value for IsParentRecovery to be an explicit nil

### UnsetIsParentRecovery
`func (o *CommonRecoveryResponseParams) UnsetIsParentRecovery()`

UnsetIsParentRecovery ensures that no value is present for IsParentRecovery, not even an explicit nil
### GetParentRecoveryId

`func (o *CommonRecoveryResponseParams) GetParentRecoveryId() string`

GetParentRecoveryId returns the ParentRecoveryId field if non-nil, zero value otherwise.

### GetParentRecoveryIdOk

`func (o *CommonRecoveryResponseParams) GetParentRecoveryIdOk() (*string, bool)`

GetParentRecoveryIdOk returns a tuple with the ParentRecoveryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRecoveryId

`func (o *CommonRecoveryResponseParams) SetParentRecoveryId(v string)`

SetParentRecoveryId sets ParentRecoveryId field to given value.

### HasParentRecoveryId

`func (o *CommonRecoveryResponseParams) HasParentRecoveryId() bool`

HasParentRecoveryId returns a boolean if a field has been set.

### SetParentRecoveryIdNil

`func (o *CommonRecoveryResponseParams) SetParentRecoveryIdNil(b bool)`

 SetParentRecoveryIdNil sets the value for ParentRecoveryId to be an explicit nil

### UnsetParentRecoveryId
`func (o *CommonRecoveryResponseParams) UnsetParentRecoveryId()`

UnsetParentRecoveryId ensures that no value is present for ParentRecoveryId, not even an explicit nil
### GetRetrieveArchiveTasks

`func (o *CommonRecoveryResponseParams) GetRetrieveArchiveTasks() []RetrieveArchiveTask`

GetRetrieveArchiveTasks returns the RetrieveArchiveTasks field if non-nil, zero value otherwise.

### GetRetrieveArchiveTasksOk

`func (o *CommonRecoveryResponseParams) GetRetrieveArchiveTasksOk() (*[]RetrieveArchiveTask, bool)`

GetRetrieveArchiveTasksOk returns a tuple with the RetrieveArchiveTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrieveArchiveTasks

`func (o *CommonRecoveryResponseParams) SetRetrieveArchiveTasks(v []RetrieveArchiveTask)`

SetRetrieveArchiveTasks sets RetrieveArchiveTasks field to given value.

### HasRetrieveArchiveTasks

`func (o *CommonRecoveryResponseParams) HasRetrieveArchiveTasks() bool`

HasRetrieveArchiveTasks returns a boolean if a field has been set.

### SetRetrieveArchiveTasksNil

`func (o *CommonRecoveryResponseParams) SetRetrieveArchiveTasksNil(b bool)`

 SetRetrieveArchiveTasksNil sets the value for RetrieveArchiveTasks to be an explicit nil

### UnsetRetrieveArchiveTasks
`func (o *CommonRecoveryResponseParams) UnsetRetrieveArchiveTasks()`

UnsetRetrieveArchiveTasks ensures that no value is present for RetrieveArchiveTasks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


