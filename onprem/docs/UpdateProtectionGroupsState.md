# UpdateProtectionGroupsState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedProtectionGroups** | Pointer to [**[]FailedProtectionGroupDetails**](FailedProtectionGroupDetails.md) | Specifies a list of Protection Group ids along with details for which updation of state was failed. | [optional] 
**SuccessfulProtectionGroupIds** | Pointer to **[]string** | Specifies a list of Protection Group ids for which updation of state was successful. | [optional] 

## Methods

### NewUpdateProtectionGroupsState

`func NewUpdateProtectionGroupsState() *UpdateProtectionGroupsState`

NewUpdateProtectionGroupsState instantiates a new UpdateProtectionGroupsState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProtectionGroupsStateWithDefaults

`func NewUpdateProtectionGroupsStateWithDefaults() *UpdateProtectionGroupsState`

NewUpdateProtectionGroupsStateWithDefaults instantiates a new UpdateProtectionGroupsState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailedProtectionGroups

`func (o *UpdateProtectionGroupsState) GetFailedProtectionGroups() []FailedProtectionGroupDetails`

GetFailedProtectionGroups returns the FailedProtectionGroups field if non-nil, zero value otherwise.

### GetFailedProtectionGroupsOk

`func (o *UpdateProtectionGroupsState) GetFailedProtectionGroupsOk() (*[]FailedProtectionGroupDetails, bool)`

GetFailedProtectionGroupsOk returns a tuple with the FailedProtectionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedProtectionGroups

`func (o *UpdateProtectionGroupsState) SetFailedProtectionGroups(v []FailedProtectionGroupDetails)`

SetFailedProtectionGroups sets FailedProtectionGroups field to given value.

### HasFailedProtectionGroups

`func (o *UpdateProtectionGroupsState) HasFailedProtectionGroups() bool`

HasFailedProtectionGroups returns a boolean if a field has been set.

### SetFailedProtectionGroupsNil

`func (o *UpdateProtectionGroupsState) SetFailedProtectionGroupsNil(b bool)`

 SetFailedProtectionGroupsNil sets the value for FailedProtectionGroups to be an explicit nil

### UnsetFailedProtectionGroups
`func (o *UpdateProtectionGroupsState) UnsetFailedProtectionGroups()`

UnsetFailedProtectionGroups ensures that no value is present for FailedProtectionGroups, not even an explicit nil
### GetSuccessfulProtectionGroupIds

`func (o *UpdateProtectionGroupsState) GetSuccessfulProtectionGroupIds() []string`

GetSuccessfulProtectionGroupIds returns the SuccessfulProtectionGroupIds field if non-nil, zero value otherwise.

### GetSuccessfulProtectionGroupIdsOk

`func (o *UpdateProtectionGroupsState) GetSuccessfulProtectionGroupIdsOk() (*[]string, bool)`

GetSuccessfulProtectionGroupIdsOk returns a tuple with the SuccessfulProtectionGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulProtectionGroupIds

`func (o *UpdateProtectionGroupsState) SetSuccessfulProtectionGroupIds(v []string)`

SetSuccessfulProtectionGroupIds sets SuccessfulProtectionGroupIds field to given value.

### HasSuccessfulProtectionGroupIds

`func (o *UpdateProtectionGroupsState) HasSuccessfulProtectionGroupIds() bool`

HasSuccessfulProtectionGroupIds returns a boolean if a field has been set.

### SetSuccessfulProtectionGroupIdsNil

`func (o *UpdateProtectionGroupsState) SetSuccessfulProtectionGroupIdsNil(b bool)`

 SetSuccessfulProtectionGroupIdsNil sets the value for SuccessfulProtectionGroupIds to be an explicit nil

### UnsetSuccessfulProtectionGroupIds
`func (o *UpdateProtectionGroupsState) UnsetSuccessfulProtectionGroupIds()`

UnsetSuccessfulProtectionGroupIds ensures that no value is present for SuccessfulProtectionGroupIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


