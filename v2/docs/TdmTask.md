# TdmTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **NullableString** | Specifies the TDM Task action. | 
**CreatedByUser** | Pointer to [**User**](User.md) | Specifies the user, who created this task. | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the time (in usecs from epoch) when the task was completed. | [optional] 
**Id** | **NullableString** | Specifies the unique ID of the task. | 
**Name** | Pointer to **NullableString** | Specifies the name of the task. | [optional] 
**ProgressTaskId** | Pointer to **NullableString** | Specifies the ID for tracking progress of this task. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the time (in usecs from epoch) when the task was started. | [optional] 
**Status** | Pointer to **NullableString** | Specifies the current status of the task. | [optional] 
**CloneParams** | Pointer to [**TdmCloneTaskResponseParams**](TdmCloneTaskResponseParams.md) |  | [optional] 
**RefreshParams** | Pointer to [**TdmRefreshTaskResponseParams**](TdmRefreshTaskResponseParams.md) |  | [optional] 

## Methods

### NewTdmTask

`func NewTdmTask(action NullableString, id NullableString, ) *TdmTask`

NewTdmTask instantiates a new TdmTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTdmTaskWithDefaults

`func NewTdmTaskWithDefaults() *TdmTask`

NewTdmTaskWithDefaults instantiates a new TdmTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *TdmTask) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TdmTask) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TdmTask) SetAction(v string)`

SetAction sets Action field to given value.


### SetActionNil

`func (o *TdmTask) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *TdmTask) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetCreatedByUser

`func (o *TdmTask) GetCreatedByUser() User`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *TdmTask) GetCreatedByUserOk() (*User, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *TdmTask) SetCreatedByUser(v User)`

SetCreatedByUser sets CreatedByUser field to given value.

### HasCreatedByUser

`func (o *TdmTask) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### GetEndTimeUsecs

`func (o *TdmTask) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *TdmTask) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *TdmTask) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *TdmTask) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *TdmTask) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *TdmTask) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetId

`func (o *TdmTask) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TdmTask) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TdmTask) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *TdmTask) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TdmTask) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *TdmTask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TdmTask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TdmTask) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TdmTask) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TdmTask) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TdmTask) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProgressTaskId

`func (o *TdmTask) GetProgressTaskId() string`

GetProgressTaskId returns the ProgressTaskId field if non-nil, zero value otherwise.

### GetProgressTaskIdOk

`func (o *TdmTask) GetProgressTaskIdOk() (*string, bool)`

GetProgressTaskIdOk returns a tuple with the ProgressTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressTaskId

`func (o *TdmTask) SetProgressTaskId(v string)`

SetProgressTaskId sets ProgressTaskId field to given value.

### HasProgressTaskId

`func (o *TdmTask) HasProgressTaskId() bool`

HasProgressTaskId returns a boolean if a field has been set.

### SetProgressTaskIdNil

`func (o *TdmTask) SetProgressTaskIdNil(b bool)`

 SetProgressTaskIdNil sets the value for ProgressTaskId to be an explicit nil

### UnsetProgressTaskId
`func (o *TdmTask) UnsetProgressTaskId()`

UnsetProgressTaskId ensures that no value is present for ProgressTaskId, not even an explicit nil
### GetStartTimeUsecs

`func (o *TdmTask) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *TdmTask) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *TdmTask) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *TdmTask) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *TdmTask) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *TdmTask) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetStatus

`func (o *TdmTask) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TdmTask) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TdmTask) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TdmTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TdmTask) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TdmTask) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetCloneParams

`func (o *TdmTask) GetCloneParams() TdmCloneTaskResponseParams`

GetCloneParams returns the CloneParams field if non-nil, zero value otherwise.

### GetCloneParamsOk

`func (o *TdmTask) GetCloneParamsOk() (*TdmCloneTaskResponseParams, bool)`

GetCloneParamsOk returns a tuple with the CloneParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneParams

`func (o *TdmTask) SetCloneParams(v TdmCloneTaskResponseParams)`

SetCloneParams sets CloneParams field to given value.

### HasCloneParams

`func (o *TdmTask) HasCloneParams() bool`

HasCloneParams returns a boolean if a field has been set.

### GetRefreshParams

`func (o *TdmTask) GetRefreshParams() TdmRefreshTaskResponseParams`

GetRefreshParams returns the RefreshParams field if non-nil, zero value otherwise.

### GetRefreshParamsOk

`func (o *TdmTask) GetRefreshParamsOk() (*TdmRefreshTaskResponseParams, bool)`

GetRefreshParamsOk returns a tuple with the RefreshParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshParams

`func (o *TdmTask) SetRefreshParams(v TdmRefreshTaskResponseParams)`

SetRefreshParams sets RefreshParams field to given value.

### HasRefreshParams

`func (o *TdmTask) HasRefreshParams() bool`

HasRefreshParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


