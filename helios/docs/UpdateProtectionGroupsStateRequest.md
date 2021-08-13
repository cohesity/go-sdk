# UpdateProtectionGroupsStateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **NullableString** | Specifies the action to be performed on all the specfied Protection Groups. &#39;kActivate&#39; specifies that Protection Group should be activated. &#39;kDeactivate&#39; sepcifies that Protection Group should be deactivated. &#39;kPause&#39; specifies that Protection Group should be paused. &#39;kResume&#39; specifies that Protection Group should be resumed. | 
**Ids** | **[]string** | Specifies a list of Protection Group ids for which the state should change. | 

## Methods

### NewUpdateProtectionGroupsStateRequest

`func NewUpdateProtectionGroupsStateRequest(action NullableString, ids []string, ) *UpdateProtectionGroupsStateRequest`

NewUpdateProtectionGroupsStateRequest instantiates a new UpdateProtectionGroupsStateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProtectionGroupsStateRequestWithDefaults

`func NewUpdateProtectionGroupsStateRequestWithDefaults() *UpdateProtectionGroupsStateRequest`

NewUpdateProtectionGroupsStateRequestWithDefaults instantiates a new UpdateProtectionGroupsStateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *UpdateProtectionGroupsStateRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UpdateProtectionGroupsStateRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UpdateProtectionGroupsStateRequest) SetAction(v string)`

SetAction sets Action field to given value.


### SetActionNil

`func (o *UpdateProtectionGroupsStateRequest) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *UpdateProtectionGroupsStateRequest) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetIds

`func (o *UpdateProtectionGroupsStateRequest) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *UpdateProtectionGroupsStateRequest) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *UpdateProtectionGroupsStateRequest) SetIds(v []string)`

SetIds sets Ids field to given value.


### SetIdsNil

`func (o *UpdateProtectionGroupsStateRequest) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *UpdateProtectionGroupsStateRequest) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


