# TdmObjectTimelineEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** | Specifies the unique ID of the event. | 
**CreatedAt** | Pointer to **NullableInt64** | Specifies the time (in usecs from epoch) at which the event was created. | [optional] 
**CreatedByUser** | Pointer to [**NullableUser**](User.md) | Specifies the user, who triggered the event. | [optional] 
**Status** | Pointer to **NullableString** | Specifies the current status of the event. | [optional] 
**ErrorMessage** | Pointer to **NullableString** | Specifies the error message if the event is in failed state. | [optional] 
**Description** | Pointer to **NullableString** | Specifies the description of the event, as provided by the user. | [optional] 
**EventGroupId** | Pointer to **NullableString** | Specifies the ID of the group this event belongs to. Events with same group ID are considered to be a single timeline for the TDM object. Different group IDs mean different timelines for the TDM object. | [optional] 
**Action** | **NullableString** | Specifies the action for the event. | 
**CloneParams** | Pointer to [**TdmCloneTaskResponseParams**](TdmCloneTaskResponseParams.md) |  | [optional] 
**SnapshotParams** | Pointer to [**TdmSnapshot**](TdmSnapshot.md) |  | [optional] 

## Methods

### NewTdmObjectTimelineEvent

`func NewTdmObjectTimelineEvent(id NullableString, action NullableString, ) *TdmObjectTimelineEvent`

NewTdmObjectTimelineEvent instantiates a new TdmObjectTimelineEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTdmObjectTimelineEventWithDefaults

`func NewTdmObjectTimelineEventWithDefaults() *TdmObjectTimelineEvent`

NewTdmObjectTimelineEventWithDefaults instantiates a new TdmObjectTimelineEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TdmObjectTimelineEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TdmObjectTimelineEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TdmObjectTimelineEvent) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *TdmObjectTimelineEvent) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TdmObjectTimelineEvent) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetCreatedAt

`func (o *TdmObjectTimelineEvent) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TdmObjectTimelineEvent) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TdmObjectTimelineEvent) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TdmObjectTimelineEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *TdmObjectTimelineEvent) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *TdmObjectTimelineEvent) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCreatedByUser

`func (o *TdmObjectTimelineEvent) GetCreatedByUser() User`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *TdmObjectTimelineEvent) GetCreatedByUserOk() (*User, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *TdmObjectTimelineEvent) SetCreatedByUser(v User)`

SetCreatedByUser sets CreatedByUser field to given value.

### HasCreatedByUser

`func (o *TdmObjectTimelineEvent) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### SetCreatedByUserNil

`func (o *TdmObjectTimelineEvent) SetCreatedByUserNil(b bool)`

 SetCreatedByUserNil sets the value for CreatedByUser to be an explicit nil

### UnsetCreatedByUser
`func (o *TdmObjectTimelineEvent) UnsetCreatedByUser()`

UnsetCreatedByUser ensures that no value is present for CreatedByUser, not even an explicit nil
### GetStatus

`func (o *TdmObjectTimelineEvent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TdmObjectTimelineEvent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TdmObjectTimelineEvent) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TdmObjectTimelineEvent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TdmObjectTimelineEvent) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TdmObjectTimelineEvent) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetErrorMessage

`func (o *TdmObjectTimelineEvent) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *TdmObjectTimelineEvent) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *TdmObjectTimelineEvent) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *TdmObjectTimelineEvent) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *TdmObjectTimelineEvent) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *TdmObjectTimelineEvent) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetDescription

`func (o *TdmObjectTimelineEvent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TdmObjectTimelineEvent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TdmObjectTimelineEvent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TdmObjectTimelineEvent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TdmObjectTimelineEvent) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TdmObjectTimelineEvent) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEventGroupId

`func (o *TdmObjectTimelineEvent) GetEventGroupId() string`

GetEventGroupId returns the EventGroupId field if non-nil, zero value otherwise.

### GetEventGroupIdOk

`func (o *TdmObjectTimelineEvent) GetEventGroupIdOk() (*string, bool)`

GetEventGroupIdOk returns a tuple with the EventGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventGroupId

`func (o *TdmObjectTimelineEvent) SetEventGroupId(v string)`

SetEventGroupId sets EventGroupId field to given value.

### HasEventGroupId

`func (o *TdmObjectTimelineEvent) HasEventGroupId() bool`

HasEventGroupId returns a boolean if a field has been set.

### SetEventGroupIdNil

`func (o *TdmObjectTimelineEvent) SetEventGroupIdNil(b bool)`

 SetEventGroupIdNil sets the value for EventGroupId to be an explicit nil

### UnsetEventGroupId
`func (o *TdmObjectTimelineEvent) UnsetEventGroupId()`

UnsetEventGroupId ensures that no value is present for EventGroupId, not even an explicit nil
### GetAction

`func (o *TdmObjectTimelineEvent) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TdmObjectTimelineEvent) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TdmObjectTimelineEvent) SetAction(v string)`

SetAction sets Action field to given value.


### SetActionNil

`func (o *TdmObjectTimelineEvent) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *TdmObjectTimelineEvent) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetCloneParams

`func (o *TdmObjectTimelineEvent) GetCloneParams() TdmCloneTaskResponseParams`

GetCloneParams returns the CloneParams field if non-nil, zero value otherwise.

### GetCloneParamsOk

`func (o *TdmObjectTimelineEvent) GetCloneParamsOk() (*TdmCloneTaskResponseParams, bool)`

GetCloneParamsOk returns a tuple with the CloneParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneParams

`func (o *TdmObjectTimelineEvent) SetCloneParams(v TdmCloneTaskResponseParams)`

SetCloneParams sets CloneParams field to given value.

### HasCloneParams

`func (o *TdmObjectTimelineEvent) HasCloneParams() bool`

HasCloneParams returns a boolean if a field has been set.

### GetSnapshotParams

`func (o *TdmObjectTimelineEvent) GetSnapshotParams() TdmSnapshot`

GetSnapshotParams returns the SnapshotParams field if non-nil, zero value otherwise.

### GetSnapshotParamsOk

`func (o *TdmObjectTimelineEvent) GetSnapshotParamsOk() (*TdmSnapshot, bool)`

GetSnapshotParamsOk returns a tuple with the SnapshotParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotParams

`func (o *TdmObjectTimelineEvent) SetSnapshotParams(v TdmSnapshot)`

SetSnapshotParams sets SnapshotParams field to given value.

### HasSnapshotParams

`func (o *TdmObjectTimelineEvent) HasSnapshotParams() bool`

HasSnapshotParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


