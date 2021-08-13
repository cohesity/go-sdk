# CommonTdmObjectTimelineEventsResponseParams

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

## Methods

### NewCommonTdmObjectTimelineEventsResponseParams

`func NewCommonTdmObjectTimelineEventsResponseParams(id NullableString, action NullableString, ) *CommonTdmObjectTimelineEventsResponseParams`

NewCommonTdmObjectTimelineEventsResponseParams instantiates a new CommonTdmObjectTimelineEventsResponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonTdmObjectTimelineEventsResponseParamsWithDefaults

`func NewCommonTdmObjectTimelineEventsResponseParamsWithDefaults() *CommonTdmObjectTimelineEventsResponseParams`

NewCommonTdmObjectTimelineEventsResponseParamsWithDefaults instantiates a new CommonTdmObjectTimelineEventsResponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommonTdmObjectTimelineEventsResponseParams) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommonTdmObjectTimelineEventsResponseParams) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommonTdmObjectTimelineEventsResponseParams) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *CommonTdmObjectTimelineEventsResponseParams) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CommonTdmObjectTimelineEventsResponseParams) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetCreatedAt

`func (o *CommonTdmObjectTimelineEventsResponseParams) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CommonTdmObjectTimelineEventsResponseParams) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CommonTdmObjectTimelineEventsResponseParams) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CommonTdmObjectTimelineEventsResponseParams) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *CommonTdmObjectTimelineEventsResponseParams) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *CommonTdmObjectTimelineEventsResponseParams) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCreatedByUser

`func (o *CommonTdmObjectTimelineEventsResponseParams) GetCreatedByUser() User`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *CommonTdmObjectTimelineEventsResponseParams) GetCreatedByUserOk() (*User, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *CommonTdmObjectTimelineEventsResponseParams) SetCreatedByUser(v User)`

SetCreatedByUser sets CreatedByUser field to given value.

### HasCreatedByUser

`func (o *CommonTdmObjectTimelineEventsResponseParams) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### SetCreatedByUserNil

`func (o *CommonTdmObjectTimelineEventsResponseParams) SetCreatedByUserNil(b bool)`

 SetCreatedByUserNil sets the value for CreatedByUser to be an explicit nil

### UnsetCreatedByUser
`func (o *CommonTdmObjectTimelineEventsResponseParams) UnsetCreatedByUser()`

UnsetCreatedByUser ensures that no value is present for CreatedByUser, not even an explicit nil
### GetStatus

`func (o *CommonTdmObjectTimelineEventsResponseParams) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CommonTdmObjectTimelineEventsResponseParams) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CommonTdmObjectTimelineEventsResponseParams) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CommonTdmObjectTimelineEventsResponseParams) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *CommonTdmObjectTimelineEventsResponseParams) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *CommonTdmObjectTimelineEventsResponseParams) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetErrorMessage

`func (o *CommonTdmObjectTimelineEventsResponseParams) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *CommonTdmObjectTimelineEventsResponseParams) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *CommonTdmObjectTimelineEventsResponseParams) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *CommonTdmObjectTimelineEventsResponseParams) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *CommonTdmObjectTimelineEventsResponseParams) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *CommonTdmObjectTimelineEventsResponseParams) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetDescription

`func (o *CommonTdmObjectTimelineEventsResponseParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CommonTdmObjectTimelineEventsResponseParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CommonTdmObjectTimelineEventsResponseParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CommonTdmObjectTimelineEventsResponseParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CommonTdmObjectTimelineEventsResponseParams) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CommonTdmObjectTimelineEventsResponseParams) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEventGroupId

`func (o *CommonTdmObjectTimelineEventsResponseParams) GetEventGroupId() string`

GetEventGroupId returns the EventGroupId field if non-nil, zero value otherwise.

### GetEventGroupIdOk

`func (o *CommonTdmObjectTimelineEventsResponseParams) GetEventGroupIdOk() (*string, bool)`

GetEventGroupIdOk returns a tuple with the EventGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventGroupId

`func (o *CommonTdmObjectTimelineEventsResponseParams) SetEventGroupId(v string)`

SetEventGroupId sets EventGroupId field to given value.

### HasEventGroupId

`func (o *CommonTdmObjectTimelineEventsResponseParams) HasEventGroupId() bool`

HasEventGroupId returns a boolean if a field has been set.

### SetEventGroupIdNil

`func (o *CommonTdmObjectTimelineEventsResponseParams) SetEventGroupIdNil(b bool)`

 SetEventGroupIdNil sets the value for EventGroupId to be an explicit nil

### UnsetEventGroupId
`func (o *CommonTdmObjectTimelineEventsResponseParams) UnsetEventGroupId()`

UnsetEventGroupId ensures that no value is present for EventGroupId, not even an explicit nil
### GetAction

`func (o *CommonTdmObjectTimelineEventsResponseParams) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CommonTdmObjectTimelineEventsResponseParams) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CommonTdmObjectTimelineEventsResponseParams) SetAction(v string)`

SetAction sets Action field to given value.


### SetActionNil

`func (o *CommonTdmObjectTimelineEventsResponseParams) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *CommonTdmObjectTimelineEventsResponseParams) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


