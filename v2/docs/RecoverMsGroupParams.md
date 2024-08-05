# RecoverMsGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering other MS groups if one of MS groups failed to recover. Default value is false. | [optional] 
**MsGroups** | [**[]MsGroupParam**](MsGroupParam.md) | Specifies a list of groups getting restored. | 
**RestoreToOriginal** | Pointer to **NullableBool** | Specifies whether or not all groups are restored to original location. | [optional] 
**TargetGroup** | Pointer to **NullableString** | Specifies target group nickname in case restoreToOriginal is false. This needs to be specifid when restoreToOriginal is false. | [optional] 
**TargetGroupName** | Pointer to **NullableString** | Specifies target group name in case restore_to_original is false. This needs to be specifid when restoreToOriginal is false. However, this will be ignored if restoring to alternate existing group (i.e. to a group the nickname of which is same as the one supplied by the end user). | [optional] 

## Methods

### NewRecoverMsGroupParams

`func NewRecoverMsGroupParams(msGroups []MsGroupParam, ) *RecoverMsGroupParams`

NewRecoverMsGroupParams instantiates a new RecoverMsGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverMsGroupParamsWithDefaults

`func NewRecoverMsGroupParamsWithDefaults() *RecoverMsGroupParams`

NewRecoverMsGroupParamsWithDefaults instantiates a new RecoverMsGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *RecoverMsGroupParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *RecoverMsGroupParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *RecoverMsGroupParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *RecoverMsGroupParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *RecoverMsGroupParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *RecoverMsGroupParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetMsGroups

`func (o *RecoverMsGroupParams) GetMsGroups() []MsGroupParam`

GetMsGroups returns the MsGroups field if non-nil, zero value otherwise.

### GetMsGroupsOk

`func (o *RecoverMsGroupParams) GetMsGroupsOk() (*[]MsGroupParam, bool)`

GetMsGroupsOk returns a tuple with the MsGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsGroups

`func (o *RecoverMsGroupParams) SetMsGroups(v []MsGroupParam)`

SetMsGroups sets MsGroups field to given value.


### SetMsGroupsNil

`func (o *RecoverMsGroupParams) SetMsGroupsNil(b bool)`

 SetMsGroupsNil sets the value for MsGroups to be an explicit nil

### UnsetMsGroups
`func (o *RecoverMsGroupParams) UnsetMsGroups()`

UnsetMsGroups ensures that no value is present for MsGroups, not even an explicit nil
### GetRestoreToOriginal

`func (o *RecoverMsGroupParams) GetRestoreToOriginal() bool`

GetRestoreToOriginal returns the RestoreToOriginal field if non-nil, zero value otherwise.

### GetRestoreToOriginalOk

`func (o *RecoverMsGroupParams) GetRestoreToOriginalOk() (*bool, bool)`

GetRestoreToOriginalOk returns a tuple with the RestoreToOriginal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreToOriginal

`func (o *RecoverMsGroupParams) SetRestoreToOriginal(v bool)`

SetRestoreToOriginal sets RestoreToOriginal field to given value.

### HasRestoreToOriginal

`func (o *RecoverMsGroupParams) HasRestoreToOriginal() bool`

HasRestoreToOriginal returns a boolean if a field has been set.

### SetRestoreToOriginalNil

`func (o *RecoverMsGroupParams) SetRestoreToOriginalNil(b bool)`

 SetRestoreToOriginalNil sets the value for RestoreToOriginal to be an explicit nil

### UnsetRestoreToOriginal
`func (o *RecoverMsGroupParams) UnsetRestoreToOriginal()`

UnsetRestoreToOriginal ensures that no value is present for RestoreToOriginal, not even an explicit nil
### GetTargetGroup

`func (o *RecoverMsGroupParams) GetTargetGroup() string`

GetTargetGroup returns the TargetGroup field if non-nil, zero value otherwise.

### GetTargetGroupOk

`func (o *RecoverMsGroupParams) GetTargetGroupOk() (*string, bool)`

GetTargetGroupOk returns a tuple with the TargetGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroup

`func (o *RecoverMsGroupParams) SetTargetGroup(v string)`

SetTargetGroup sets TargetGroup field to given value.

### HasTargetGroup

`func (o *RecoverMsGroupParams) HasTargetGroup() bool`

HasTargetGroup returns a boolean if a field has been set.

### SetTargetGroupNil

`func (o *RecoverMsGroupParams) SetTargetGroupNil(b bool)`

 SetTargetGroupNil sets the value for TargetGroup to be an explicit nil

### UnsetTargetGroup
`func (o *RecoverMsGroupParams) UnsetTargetGroup()`

UnsetTargetGroup ensures that no value is present for TargetGroup, not even an explicit nil
### GetTargetGroupName

`func (o *RecoverMsGroupParams) GetTargetGroupName() string`

GetTargetGroupName returns the TargetGroupName field if non-nil, zero value otherwise.

### GetTargetGroupNameOk

`func (o *RecoverMsGroupParams) GetTargetGroupNameOk() (*string, bool)`

GetTargetGroupNameOk returns a tuple with the TargetGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupName

`func (o *RecoverMsGroupParams) SetTargetGroupName(v string)`

SetTargetGroupName sets TargetGroupName field to given value.

### HasTargetGroupName

`func (o *RecoverMsGroupParams) HasTargetGroupName() bool`

HasTargetGroupName returns a boolean if a field has been set.

### SetTargetGroupNameNil

`func (o *RecoverMsGroupParams) SetTargetGroupNameNil(b bool)`

 SetTargetGroupNameNil sets the value for TargetGroupName to be an explicit nil

### UnsetTargetGroupName
`func (o *RecoverMsGroupParams) UnsetTargetGroupName()`

UnsetTargetGroupName ensures that no value is present for TargetGroupName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


