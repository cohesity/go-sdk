# RecoverO365ParamsRecoverMsGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering other MS groups if one of MS groups failed to recover. Default value is false. | [optional] 
**MsGroups** | [**[]MsGroupParam**](MsGroupParam.md) | Specifies a list of groups getting restored. | 
**RestoreToOriginal** | Pointer to **NullableBool** | Specifies whether or not all groups are restored to original location. | [optional] 
**TargetGroup** | Pointer to **NullableString** | Specifies target group nickname in case restoreToOriginal is false. This needs to be specifid when restoreToOriginal is false. | [optional] 
**TargetGroupName** | Pointer to **NullableString** | Specifies target group name in case restore_to_original is false. This needs to be specifid when restoreToOriginal is false. However, this will be ignored if restoring to alternate existing group (i.e. to a group the nickname of which is same as the one supplied by the end user). | [optional] 

## Methods

### NewRecoverO365ParamsRecoverMsGroupParams

`func NewRecoverO365ParamsRecoverMsGroupParams(msGroups []MsGroupParam, ) *RecoverO365ParamsRecoverMsGroupParams`

NewRecoverO365ParamsRecoverMsGroupParams instantiates a new RecoverO365ParamsRecoverMsGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverO365ParamsRecoverMsGroupParamsWithDefaults

`func NewRecoverO365ParamsRecoverMsGroupParamsWithDefaults() *RecoverO365ParamsRecoverMsGroupParams`

NewRecoverO365ParamsRecoverMsGroupParamsWithDefaults instantiates a new RecoverO365ParamsRecoverMsGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *RecoverO365ParamsRecoverMsGroupParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *RecoverO365ParamsRecoverMsGroupParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *RecoverO365ParamsRecoverMsGroupParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *RecoverO365ParamsRecoverMsGroupParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *RecoverO365ParamsRecoverMsGroupParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *RecoverO365ParamsRecoverMsGroupParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetMsGroups

`func (o *RecoverO365ParamsRecoverMsGroupParams) GetMsGroups() []MsGroupParam`

GetMsGroups returns the MsGroups field if non-nil, zero value otherwise.

### GetMsGroupsOk

`func (o *RecoverO365ParamsRecoverMsGroupParams) GetMsGroupsOk() (*[]MsGroupParam, bool)`

GetMsGroupsOk returns a tuple with the MsGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsGroups

`func (o *RecoverO365ParamsRecoverMsGroupParams) SetMsGroups(v []MsGroupParam)`

SetMsGroups sets MsGroups field to given value.


### SetMsGroupsNil

`func (o *RecoverO365ParamsRecoverMsGroupParams) SetMsGroupsNil(b bool)`

 SetMsGroupsNil sets the value for MsGroups to be an explicit nil

### UnsetMsGroups
`func (o *RecoverO365ParamsRecoverMsGroupParams) UnsetMsGroups()`

UnsetMsGroups ensures that no value is present for MsGroups, not even an explicit nil
### GetRestoreToOriginal

`func (o *RecoverO365ParamsRecoverMsGroupParams) GetRestoreToOriginal() bool`

GetRestoreToOriginal returns the RestoreToOriginal field if non-nil, zero value otherwise.

### GetRestoreToOriginalOk

`func (o *RecoverO365ParamsRecoverMsGroupParams) GetRestoreToOriginalOk() (*bool, bool)`

GetRestoreToOriginalOk returns a tuple with the RestoreToOriginal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreToOriginal

`func (o *RecoverO365ParamsRecoverMsGroupParams) SetRestoreToOriginal(v bool)`

SetRestoreToOriginal sets RestoreToOriginal field to given value.

### HasRestoreToOriginal

`func (o *RecoverO365ParamsRecoverMsGroupParams) HasRestoreToOriginal() bool`

HasRestoreToOriginal returns a boolean if a field has been set.

### SetRestoreToOriginalNil

`func (o *RecoverO365ParamsRecoverMsGroupParams) SetRestoreToOriginalNil(b bool)`

 SetRestoreToOriginalNil sets the value for RestoreToOriginal to be an explicit nil

### UnsetRestoreToOriginal
`func (o *RecoverO365ParamsRecoverMsGroupParams) UnsetRestoreToOriginal()`

UnsetRestoreToOriginal ensures that no value is present for RestoreToOriginal, not even an explicit nil
### GetTargetGroup

`func (o *RecoverO365ParamsRecoverMsGroupParams) GetTargetGroup() string`

GetTargetGroup returns the TargetGroup field if non-nil, zero value otherwise.

### GetTargetGroupOk

`func (o *RecoverO365ParamsRecoverMsGroupParams) GetTargetGroupOk() (*string, bool)`

GetTargetGroupOk returns a tuple with the TargetGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroup

`func (o *RecoverO365ParamsRecoverMsGroupParams) SetTargetGroup(v string)`

SetTargetGroup sets TargetGroup field to given value.

### HasTargetGroup

`func (o *RecoverO365ParamsRecoverMsGroupParams) HasTargetGroup() bool`

HasTargetGroup returns a boolean if a field has been set.

### SetTargetGroupNil

`func (o *RecoverO365ParamsRecoverMsGroupParams) SetTargetGroupNil(b bool)`

 SetTargetGroupNil sets the value for TargetGroup to be an explicit nil

### UnsetTargetGroup
`func (o *RecoverO365ParamsRecoverMsGroupParams) UnsetTargetGroup()`

UnsetTargetGroup ensures that no value is present for TargetGroup, not even an explicit nil
### GetTargetGroupName

`func (o *RecoverO365ParamsRecoverMsGroupParams) GetTargetGroupName() string`

GetTargetGroupName returns the TargetGroupName field if non-nil, zero value otherwise.

### GetTargetGroupNameOk

`func (o *RecoverO365ParamsRecoverMsGroupParams) GetTargetGroupNameOk() (*string, bool)`

GetTargetGroupNameOk returns a tuple with the TargetGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupName

`func (o *RecoverO365ParamsRecoverMsGroupParams) SetTargetGroupName(v string)`

SetTargetGroupName sets TargetGroupName field to given value.

### HasTargetGroupName

`func (o *RecoverO365ParamsRecoverMsGroupParams) HasTargetGroupName() bool`

HasTargetGroupName returns a boolean if a field has been set.

### SetTargetGroupNameNil

`func (o *RecoverO365ParamsRecoverMsGroupParams) SetTargetGroupNameNil(b bool)`

 SetTargetGroupNameNil sets the value for TargetGroupName to be an explicit nil

### UnsetTargetGroupName
`func (o *RecoverO365ParamsRecoverMsGroupParams) UnsetTargetGroupName()`

UnsetTargetGroupName ensures that no value is present for TargetGroupName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


