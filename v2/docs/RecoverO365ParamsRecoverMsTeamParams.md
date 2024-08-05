# RecoverO365ParamsRecoverMsTeamParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering other teams, if some of the teams fail to recover. Default value is false. | [optional] 
**CreateNewTeam** | Pointer to **NullableBool** | Specifies to create new team in case the target team doesn&#39;t exists in case restoreToOriginal is false. | [optional] 
**Objects** | [**[]ObjectMsTeamParam**](ObjectMsTeamParam.md) | Specifies a list of Microsoft 365 Teams params associated with objects to recover. | 
**RestoreOriginalOwners** | Pointer to **NullableBool** | Specifies if the original members/owners should be part of the newly created target team or not. | [optional] 
**RestoreToOriginal** | Pointer to **NullableBool** | Specifies whether or not all Microsoft 365 Teams are restored to original location. | [optional] 
**TargetMsTeam** | Pointer to [**RecoverMsTeamParamsTargetMsTeam**](RecoverMsTeamParamsTargetMsTeam.md) |  | [optional] 
**TargetMsTeamParam** | Pointer to [**RecoverMsTeamParamsTargetMsTeamParam**](RecoverMsTeamParamsTargetMsTeamParam.md) |  | [optional] 
**TargetTeamFullName** | Pointer to **NullableString** | This field is deprecated. Specifies target team name in case restoreToOriginal is false. This will be ignored if restoring to alternate existing team (i.e. to a team the nickname of which is same as the one supplied by the end user). | [optional] 
**TargetTeamName** | Pointer to **NullableString** | Specifies the target team name in case restoreToOriginal is false. | [optional] 
**TargetTeamNickName** | Pointer to **NullableString** | This field is deprecated. Specifies target team nickname in case restoreToOriginal is false. | [optional] 
**TargetTeamOwner** | Pointer to [**NullableRecoverMsTeamParamsTargetTeamOwner**](RecoverMsTeamParamsTargetTeamOwner.md) |  | [optional] 

## Methods

### NewRecoverO365ParamsRecoverMsTeamParams

`func NewRecoverO365ParamsRecoverMsTeamParams(objects []ObjectMsTeamParam, ) *RecoverO365ParamsRecoverMsTeamParams`

NewRecoverO365ParamsRecoverMsTeamParams instantiates a new RecoverO365ParamsRecoverMsTeamParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverO365ParamsRecoverMsTeamParamsWithDefaults

`func NewRecoverO365ParamsRecoverMsTeamParamsWithDefaults() *RecoverO365ParamsRecoverMsTeamParams`

NewRecoverO365ParamsRecoverMsTeamParamsWithDefaults instantiates a new RecoverO365ParamsRecoverMsTeamParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *RecoverO365ParamsRecoverMsTeamParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *RecoverO365ParamsRecoverMsTeamParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *RecoverO365ParamsRecoverMsTeamParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *RecoverO365ParamsRecoverMsTeamParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *RecoverO365ParamsRecoverMsTeamParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *RecoverO365ParamsRecoverMsTeamParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetCreateNewTeam

`func (o *RecoverO365ParamsRecoverMsTeamParams) GetCreateNewTeam() bool`

GetCreateNewTeam returns the CreateNewTeam field if non-nil, zero value otherwise.

### GetCreateNewTeamOk

`func (o *RecoverO365ParamsRecoverMsTeamParams) GetCreateNewTeamOk() (*bool, bool)`

GetCreateNewTeamOk returns a tuple with the CreateNewTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateNewTeam

`func (o *RecoverO365ParamsRecoverMsTeamParams) SetCreateNewTeam(v bool)`

SetCreateNewTeam sets CreateNewTeam field to given value.

### HasCreateNewTeam

`func (o *RecoverO365ParamsRecoverMsTeamParams) HasCreateNewTeam() bool`

HasCreateNewTeam returns a boolean if a field has been set.

### SetCreateNewTeamNil

`func (o *RecoverO365ParamsRecoverMsTeamParams) SetCreateNewTeamNil(b bool)`

 SetCreateNewTeamNil sets the value for CreateNewTeam to be an explicit nil

### UnsetCreateNewTeam
`func (o *RecoverO365ParamsRecoverMsTeamParams) UnsetCreateNewTeam()`

UnsetCreateNewTeam ensures that no value is present for CreateNewTeam, not even an explicit nil
### GetObjects

`func (o *RecoverO365ParamsRecoverMsTeamParams) GetObjects() []ObjectMsTeamParam`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *RecoverO365ParamsRecoverMsTeamParams) GetObjectsOk() (*[]ObjectMsTeamParam, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *RecoverO365ParamsRecoverMsTeamParams) SetObjects(v []ObjectMsTeamParam)`

SetObjects sets Objects field to given value.


### SetObjectsNil

`func (o *RecoverO365ParamsRecoverMsTeamParams) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *RecoverO365ParamsRecoverMsTeamParams) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetRestoreOriginalOwners

`func (o *RecoverO365ParamsRecoverMsTeamParams) GetRestoreOriginalOwners() bool`

GetRestoreOriginalOwners returns the RestoreOriginalOwners field if non-nil, zero value otherwise.

### GetRestoreOriginalOwnersOk

`func (o *RecoverO365ParamsRecoverMsTeamParams) GetRestoreOriginalOwnersOk() (*bool, bool)`

GetRestoreOriginalOwnersOk returns a tuple with the RestoreOriginalOwners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreOriginalOwners

`func (o *RecoverO365ParamsRecoverMsTeamParams) SetRestoreOriginalOwners(v bool)`

SetRestoreOriginalOwners sets RestoreOriginalOwners field to given value.

### HasRestoreOriginalOwners

`func (o *RecoverO365ParamsRecoverMsTeamParams) HasRestoreOriginalOwners() bool`

HasRestoreOriginalOwners returns a boolean if a field has been set.

### SetRestoreOriginalOwnersNil

`func (o *RecoverO365ParamsRecoverMsTeamParams) SetRestoreOriginalOwnersNil(b bool)`

 SetRestoreOriginalOwnersNil sets the value for RestoreOriginalOwners to be an explicit nil

### UnsetRestoreOriginalOwners
`func (o *RecoverO365ParamsRecoverMsTeamParams) UnsetRestoreOriginalOwners()`

UnsetRestoreOriginalOwners ensures that no value is present for RestoreOriginalOwners, not even an explicit nil
### GetRestoreToOriginal

`func (o *RecoverO365ParamsRecoverMsTeamParams) GetRestoreToOriginal() bool`

GetRestoreToOriginal returns the RestoreToOriginal field if non-nil, zero value otherwise.

### GetRestoreToOriginalOk

`func (o *RecoverO365ParamsRecoverMsTeamParams) GetRestoreToOriginalOk() (*bool, bool)`

GetRestoreToOriginalOk returns a tuple with the RestoreToOriginal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreToOriginal

`func (o *RecoverO365ParamsRecoverMsTeamParams) SetRestoreToOriginal(v bool)`

SetRestoreToOriginal sets RestoreToOriginal field to given value.

### HasRestoreToOriginal

`func (o *RecoverO365ParamsRecoverMsTeamParams) HasRestoreToOriginal() bool`

HasRestoreToOriginal returns a boolean if a field has been set.

### SetRestoreToOriginalNil

`func (o *RecoverO365ParamsRecoverMsTeamParams) SetRestoreToOriginalNil(b bool)`

 SetRestoreToOriginalNil sets the value for RestoreToOriginal to be an explicit nil

### UnsetRestoreToOriginal
`func (o *RecoverO365ParamsRecoverMsTeamParams) UnsetRestoreToOriginal()`

UnsetRestoreToOriginal ensures that no value is present for RestoreToOriginal, not even an explicit nil
### GetTargetMsTeam

`func (o *RecoverO365ParamsRecoverMsTeamParams) GetTargetMsTeam() RecoverMsTeamParamsTargetMsTeam`

GetTargetMsTeam returns the TargetMsTeam field if non-nil, zero value otherwise.

### GetTargetMsTeamOk

`func (o *RecoverO365ParamsRecoverMsTeamParams) GetTargetMsTeamOk() (*RecoverMsTeamParamsTargetMsTeam, bool)`

GetTargetMsTeamOk returns a tuple with the TargetMsTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMsTeam

`func (o *RecoverO365ParamsRecoverMsTeamParams) SetTargetMsTeam(v RecoverMsTeamParamsTargetMsTeam)`

SetTargetMsTeam sets TargetMsTeam field to given value.

### HasTargetMsTeam

`func (o *RecoverO365ParamsRecoverMsTeamParams) HasTargetMsTeam() bool`

HasTargetMsTeam returns a boolean if a field has been set.

### GetTargetMsTeamParam

`func (o *RecoverO365ParamsRecoverMsTeamParams) GetTargetMsTeamParam() RecoverMsTeamParamsTargetMsTeamParam`

GetTargetMsTeamParam returns the TargetMsTeamParam field if non-nil, zero value otherwise.

### GetTargetMsTeamParamOk

`func (o *RecoverO365ParamsRecoverMsTeamParams) GetTargetMsTeamParamOk() (*RecoverMsTeamParamsTargetMsTeamParam, bool)`

GetTargetMsTeamParamOk returns a tuple with the TargetMsTeamParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMsTeamParam

`func (o *RecoverO365ParamsRecoverMsTeamParams) SetTargetMsTeamParam(v RecoverMsTeamParamsTargetMsTeamParam)`

SetTargetMsTeamParam sets TargetMsTeamParam field to given value.

### HasTargetMsTeamParam

`func (o *RecoverO365ParamsRecoverMsTeamParams) HasTargetMsTeamParam() bool`

HasTargetMsTeamParam returns a boolean if a field has been set.

### GetTargetTeamFullName

`func (o *RecoverO365ParamsRecoverMsTeamParams) GetTargetTeamFullName() string`

GetTargetTeamFullName returns the TargetTeamFullName field if non-nil, zero value otherwise.

### GetTargetTeamFullNameOk

`func (o *RecoverO365ParamsRecoverMsTeamParams) GetTargetTeamFullNameOk() (*string, bool)`

GetTargetTeamFullNameOk returns a tuple with the TargetTeamFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTeamFullName

`func (o *RecoverO365ParamsRecoverMsTeamParams) SetTargetTeamFullName(v string)`

SetTargetTeamFullName sets TargetTeamFullName field to given value.

### HasTargetTeamFullName

`func (o *RecoverO365ParamsRecoverMsTeamParams) HasTargetTeamFullName() bool`

HasTargetTeamFullName returns a boolean if a field has been set.

### SetTargetTeamFullNameNil

`func (o *RecoverO365ParamsRecoverMsTeamParams) SetTargetTeamFullNameNil(b bool)`

 SetTargetTeamFullNameNil sets the value for TargetTeamFullName to be an explicit nil

### UnsetTargetTeamFullName
`func (o *RecoverO365ParamsRecoverMsTeamParams) UnsetTargetTeamFullName()`

UnsetTargetTeamFullName ensures that no value is present for TargetTeamFullName, not even an explicit nil
### GetTargetTeamName

`func (o *RecoverO365ParamsRecoverMsTeamParams) GetTargetTeamName() string`

GetTargetTeamName returns the TargetTeamName field if non-nil, zero value otherwise.

### GetTargetTeamNameOk

`func (o *RecoverO365ParamsRecoverMsTeamParams) GetTargetTeamNameOk() (*string, bool)`

GetTargetTeamNameOk returns a tuple with the TargetTeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTeamName

`func (o *RecoverO365ParamsRecoverMsTeamParams) SetTargetTeamName(v string)`

SetTargetTeamName sets TargetTeamName field to given value.

### HasTargetTeamName

`func (o *RecoverO365ParamsRecoverMsTeamParams) HasTargetTeamName() bool`

HasTargetTeamName returns a boolean if a field has been set.

### SetTargetTeamNameNil

`func (o *RecoverO365ParamsRecoverMsTeamParams) SetTargetTeamNameNil(b bool)`

 SetTargetTeamNameNil sets the value for TargetTeamName to be an explicit nil

### UnsetTargetTeamName
`func (o *RecoverO365ParamsRecoverMsTeamParams) UnsetTargetTeamName()`

UnsetTargetTeamName ensures that no value is present for TargetTeamName, not even an explicit nil
### GetTargetTeamNickName

`func (o *RecoverO365ParamsRecoverMsTeamParams) GetTargetTeamNickName() string`

GetTargetTeamNickName returns the TargetTeamNickName field if non-nil, zero value otherwise.

### GetTargetTeamNickNameOk

`func (o *RecoverO365ParamsRecoverMsTeamParams) GetTargetTeamNickNameOk() (*string, bool)`

GetTargetTeamNickNameOk returns a tuple with the TargetTeamNickName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTeamNickName

`func (o *RecoverO365ParamsRecoverMsTeamParams) SetTargetTeamNickName(v string)`

SetTargetTeamNickName sets TargetTeamNickName field to given value.

### HasTargetTeamNickName

`func (o *RecoverO365ParamsRecoverMsTeamParams) HasTargetTeamNickName() bool`

HasTargetTeamNickName returns a boolean if a field has been set.

### SetTargetTeamNickNameNil

`func (o *RecoverO365ParamsRecoverMsTeamParams) SetTargetTeamNickNameNil(b bool)`

 SetTargetTeamNickNameNil sets the value for TargetTeamNickName to be an explicit nil

### UnsetTargetTeamNickName
`func (o *RecoverO365ParamsRecoverMsTeamParams) UnsetTargetTeamNickName()`

UnsetTargetTeamNickName ensures that no value is present for TargetTeamNickName, not even an explicit nil
### GetTargetTeamOwner

`func (o *RecoverO365ParamsRecoverMsTeamParams) GetTargetTeamOwner() RecoverMsTeamParamsTargetTeamOwner`

GetTargetTeamOwner returns the TargetTeamOwner field if non-nil, zero value otherwise.

### GetTargetTeamOwnerOk

`func (o *RecoverO365ParamsRecoverMsTeamParams) GetTargetTeamOwnerOk() (*RecoverMsTeamParamsTargetTeamOwner, bool)`

GetTargetTeamOwnerOk returns a tuple with the TargetTeamOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTeamOwner

`func (o *RecoverO365ParamsRecoverMsTeamParams) SetTargetTeamOwner(v RecoverMsTeamParamsTargetTeamOwner)`

SetTargetTeamOwner sets TargetTeamOwner field to given value.

### HasTargetTeamOwner

`func (o *RecoverO365ParamsRecoverMsTeamParams) HasTargetTeamOwner() bool`

HasTargetTeamOwner returns a boolean if a field has been set.

### SetTargetTeamOwnerNil

`func (o *RecoverO365ParamsRecoverMsTeamParams) SetTargetTeamOwnerNil(b bool)`

 SetTargetTeamOwnerNil sets the value for TargetTeamOwner to be an explicit nil

### UnsetTargetTeamOwner
`func (o *RecoverO365ParamsRecoverMsTeamParams) UnsetTargetTeamOwner()`

UnsetTargetTeamOwner ensures that no value is present for TargetTeamOwner, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


