# CommonTdmTaskResponseParams

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

## Methods

### NewCommonTdmTaskResponseParams

`func NewCommonTdmTaskResponseParams(action NullableString, id NullableString, ) *CommonTdmTaskResponseParams`

NewCommonTdmTaskResponseParams instantiates a new CommonTdmTaskResponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonTdmTaskResponseParamsWithDefaults

`func NewCommonTdmTaskResponseParamsWithDefaults() *CommonTdmTaskResponseParams`

NewCommonTdmTaskResponseParamsWithDefaults instantiates a new CommonTdmTaskResponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *CommonTdmTaskResponseParams) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CommonTdmTaskResponseParams) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CommonTdmTaskResponseParams) SetAction(v string)`

SetAction sets Action field to given value.


### SetActionNil

`func (o *CommonTdmTaskResponseParams) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *CommonTdmTaskResponseParams) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetCreatedByUser

`func (o *CommonTdmTaskResponseParams) GetCreatedByUser() User`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *CommonTdmTaskResponseParams) GetCreatedByUserOk() (*User, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *CommonTdmTaskResponseParams) SetCreatedByUser(v User)`

SetCreatedByUser sets CreatedByUser field to given value.

### HasCreatedByUser

`func (o *CommonTdmTaskResponseParams) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### GetEndTimeUsecs

`func (o *CommonTdmTaskResponseParams) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *CommonTdmTaskResponseParams) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *CommonTdmTaskResponseParams) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *CommonTdmTaskResponseParams) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *CommonTdmTaskResponseParams) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *CommonTdmTaskResponseParams) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetId

`func (o *CommonTdmTaskResponseParams) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommonTdmTaskResponseParams) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommonTdmTaskResponseParams) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *CommonTdmTaskResponseParams) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CommonTdmTaskResponseParams) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *CommonTdmTaskResponseParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommonTdmTaskResponseParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommonTdmTaskResponseParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CommonTdmTaskResponseParams) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CommonTdmTaskResponseParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CommonTdmTaskResponseParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProgressTaskId

`func (o *CommonTdmTaskResponseParams) GetProgressTaskId() string`

GetProgressTaskId returns the ProgressTaskId field if non-nil, zero value otherwise.

### GetProgressTaskIdOk

`func (o *CommonTdmTaskResponseParams) GetProgressTaskIdOk() (*string, bool)`

GetProgressTaskIdOk returns a tuple with the ProgressTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressTaskId

`func (o *CommonTdmTaskResponseParams) SetProgressTaskId(v string)`

SetProgressTaskId sets ProgressTaskId field to given value.

### HasProgressTaskId

`func (o *CommonTdmTaskResponseParams) HasProgressTaskId() bool`

HasProgressTaskId returns a boolean if a field has been set.

### SetProgressTaskIdNil

`func (o *CommonTdmTaskResponseParams) SetProgressTaskIdNil(b bool)`

 SetProgressTaskIdNil sets the value for ProgressTaskId to be an explicit nil

### UnsetProgressTaskId
`func (o *CommonTdmTaskResponseParams) UnsetProgressTaskId()`

UnsetProgressTaskId ensures that no value is present for ProgressTaskId, not even an explicit nil
### GetStartTimeUsecs

`func (o *CommonTdmTaskResponseParams) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *CommonTdmTaskResponseParams) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *CommonTdmTaskResponseParams) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *CommonTdmTaskResponseParams) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *CommonTdmTaskResponseParams) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *CommonTdmTaskResponseParams) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetStatus

`func (o *CommonTdmTaskResponseParams) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CommonTdmTaskResponseParams) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CommonTdmTaskResponseParams) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CommonTdmTaskResponseParams) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *CommonTdmTaskResponseParams) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *CommonTdmTaskResponseParams) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


