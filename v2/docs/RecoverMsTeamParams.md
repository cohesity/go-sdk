# RecoverMsTeamParams

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

### NewRecoverMsTeamParams

`func NewRecoverMsTeamParams(objects []ObjectMsTeamParam, ) *RecoverMsTeamParams`

NewRecoverMsTeamParams instantiates a new RecoverMsTeamParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverMsTeamParamsWithDefaults

`func NewRecoverMsTeamParamsWithDefaults() *RecoverMsTeamParams`

NewRecoverMsTeamParamsWithDefaults instantiates a new RecoverMsTeamParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *RecoverMsTeamParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *RecoverMsTeamParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *RecoverMsTeamParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *RecoverMsTeamParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *RecoverMsTeamParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *RecoverMsTeamParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetCreateNewTeam

`func (o *RecoverMsTeamParams) GetCreateNewTeam() bool`

GetCreateNewTeam returns the CreateNewTeam field if non-nil, zero value otherwise.

### GetCreateNewTeamOk

`func (o *RecoverMsTeamParams) GetCreateNewTeamOk() (*bool, bool)`

GetCreateNewTeamOk returns a tuple with the CreateNewTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateNewTeam

`func (o *RecoverMsTeamParams) SetCreateNewTeam(v bool)`

SetCreateNewTeam sets CreateNewTeam field to given value.

### HasCreateNewTeam

`func (o *RecoverMsTeamParams) HasCreateNewTeam() bool`

HasCreateNewTeam returns a boolean if a field has been set.

### SetCreateNewTeamNil

`func (o *RecoverMsTeamParams) SetCreateNewTeamNil(b bool)`

 SetCreateNewTeamNil sets the value for CreateNewTeam to be an explicit nil

### UnsetCreateNewTeam
`func (o *RecoverMsTeamParams) UnsetCreateNewTeam()`

UnsetCreateNewTeam ensures that no value is present for CreateNewTeam, not even an explicit nil
### GetObjects

`func (o *RecoverMsTeamParams) GetObjects() []ObjectMsTeamParam`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *RecoverMsTeamParams) GetObjectsOk() (*[]ObjectMsTeamParam, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *RecoverMsTeamParams) SetObjects(v []ObjectMsTeamParam)`

SetObjects sets Objects field to given value.


### SetObjectsNil

`func (o *RecoverMsTeamParams) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *RecoverMsTeamParams) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetRestoreOriginalOwners

`func (o *RecoverMsTeamParams) GetRestoreOriginalOwners() bool`

GetRestoreOriginalOwners returns the RestoreOriginalOwners field if non-nil, zero value otherwise.

### GetRestoreOriginalOwnersOk

`func (o *RecoverMsTeamParams) GetRestoreOriginalOwnersOk() (*bool, bool)`

GetRestoreOriginalOwnersOk returns a tuple with the RestoreOriginalOwners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreOriginalOwners

`func (o *RecoverMsTeamParams) SetRestoreOriginalOwners(v bool)`

SetRestoreOriginalOwners sets RestoreOriginalOwners field to given value.

### HasRestoreOriginalOwners

`func (o *RecoverMsTeamParams) HasRestoreOriginalOwners() bool`

HasRestoreOriginalOwners returns a boolean if a field has been set.

### SetRestoreOriginalOwnersNil

`func (o *RecoverMsTeamParams) SetRestoreOriginalOwnersNil(b bool)`

 SetRestoreOriginalOwnersNil sets the value for RestoreOriginalOwners to be an explicit nil

### UnsetRestoreOriginalOwners
`func (o *RecoverMsTeamParams) UnsetRestoreOriginalOwners()`

UnsetRestoreOriginalOwners ensures that no value is present for RestoreOriginalOwners, not even an explicit nil
### GetRestoreToOriginal

`func (o *RecoverMsTeamParams) GetRestoreToOriginal() bool`

GetRestoreToOriginal returns the RestoreToOriginal field if non-nil, zero value otherwise.

### GetRestoreToOriginalOk

`func (o *RecoverMsTeamParams) GetRestoreToOriginalOk() (*bool, bool)`

GetRestoreToOriginalOk returns a tuple with the RestoreToOriginal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreToOriginal

`func (o *RecoverMsTeamParams) SetRestoreToOriginal(v bool)`

SetRestoreToOriginal sets RestoreToOriginal field to given value.

### HasRestoreToOriginal

`func (o *RecoverMsTeamParams) HasRestoreToOriginal() bool`

HasRestoreToOriginal returns a boolean if a field has been set.

### SetRestoreToOriginalNil

`func (o *RecoverMsTeamParams) SetRestoreToOriginalNil(b bool)`

 SetRestoreToOriginalNil sets the value for RestoreToOriginal to be an explicit nil

### UnsetRestoreToOriginal
`func (o *RecoverMsTeamParams) UnsetRestoreToOriginal()`

UnsetRestoreToOriginal ensures that no value is present for RestoreToOriginal, not even an explicit nil
### GetTargetMsTeam

`func (o *RecoverMsTeamParams) GetTargetMsTeam() RecoverMsTeamParamsTargetMsTeam`

GetTargetMsTeam returns the TargetMsTeam field if non-nil, zero value otherwise.

### GetTargetMsTeamOk

`func (o *RecoverMsTeamParams) GetTargetMsTeamOk() (*RecoverMsTeamParamsTargetMsTeam, bool)`

GetTargetMsTeamOk returns a tuple with the TargetMsTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMsTeam

`func (o *RecoverMsTeamParams) SetTargetMsTeam(v RecoverMsTeamParamsTargetMsTeam)`

SetTargetMsTeam sets TargetMsTeam field to given value.

### HasTargetMsTeam

`func (o *RecoverMsTeamParams) HasTargetMsTeam() bool`

HasTargetMsTeam returns a boolean if a field has been set.

### GetTargetMsTeamParam

`func (o *RecoverMsTeamParams) GetTargetMsTeamParam() RecoverMsTeamParamsTargetMsTeamParam`

GetTargetMsTeamParam returns the TargetMsTeamParam field if non-nil, zero value otherwise.

### GetTargetMsTeamParamOk

`func (o *RecoverMsTeamParams) GetTargetMsTeamParamOk() (*RecoverMsTeamParamsTargetMsTeamParam, bool)`

GetTargetMsTeamParamOk returns a tuple with the TargetMsTeamParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMsTeamParam

`func (o *RecoverMsTeamParams) SetTargetMsTeamParam(v RecoverMsTeamParamsTargetMsTeamParam)`

SetTargetMsTeamParam sets TargetMsTeamParam field to given value.

### HasTargetMsTeamParam

`func (o *RecoverMsTeamParams) HasTargetMsTeamParam() bool`

HasTargetMsTeamParam returns a boolean if a field has been set.

### GetTargetTeamFullName

`func (o *RecoverMsTeamParams) GetTargetTeamFullName() string`

GetTargetTeamFullName returns the TargetTeamFullName field if non-nil, zero value otherwise.

### GetTargetTeamFullNameOk

`func (o *RecoverMsTeamParams) GetTargetTeamFullNameOk() (*string, bool)`

GetTargetTeamFullNameOk returns a tuple with the TargetTeamFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTeamFullName

`func (o *RecoverMsTeamParams) SetTargetTeamFullName(v string)`

SetTargetTeamFullName sets TargetTeamFullName field to given value.

### HasTargetTeamFullName

`func (o *RecoverMsTeamParams) HasTargetTeamFullName() bool`

HasTargetTeamFullName returns a boolean if a field has been set.

### SetTargetTeamFullNameNil

`func (o *RecoverMsTeamParams) SetTargetTeamFullNameNil(b bool)`

 SetTargetTeamFullNameNil sets the value for TargetTeamFullName to be an explicit nil

### UnsetTargetTeamFullName
`func (o *RecoverMsTeamParams) UnsetTargetTeamFullName()`

UnsetTargetTeamFullName ensures that no value is present for TargetTeamFullName, not even an explicit nil
### GetTargetTeamName

`func (o *RecoverMsTeamParams) GetTargetTeamName() string`

GetTargetTeamName returns the TargetTeamName field if non-nil, zero value otherwise.

### GetTargetTeamNameOk

`func (o *RecoverMsTeamParams) GetTargetTeamNameOk() (*string, bool)`

GetTargetTeamNameOk returns a tuple with the TargetTeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTeamName

`func (o *RecoverMsTeamParams) SetTargetTeamName(v string)`

SetTargetTeamName sets TargetTeamName field to given value.

### HasTargetTeamName

`func (o *RecoverMsTeamParams) HasTargetTeamName() bool`

HasTargetTeamName returns a boolean if a field has been set.

### SetTargetTeamNameNil

`func (o *RecoverMsTeamParams) SetTargetTeamNameNil(b bool)`

 SetTargetTeamNameNil sets the value for TargetTeamName to be an explicit nil

### UnsetTargetTeamName
`func (o *RecoverMsTeamParams) UnsetTargetTeamName()`

UnsetTargetTeamName ensures that no value is present for TargetTeamName, not even an explicit nil
### GetTargetTeamNickName

`func (o *RecoverMsTeamParams) GetTargetTeamNickName() string`

GetTargetTeamNickName returns the TargetTeamNickName field if non-nil, zero value otherwise.

### GetTargetTeamNickNameOk

`func (o *RecoverMsTeamParams) GetTargetTeamNickNameOk() (*string, bool)`

GetTargetTeamNickNameOk returns a tuple with the TargetTeamNickName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTeamNickName

`func (o *RecoverMsTeamParams) SetTargetTeamNickName(v string)`

SetTargetTeamNickName sets TargetTeamNickName field to given value.

### HasTargetTeamNickName

`func (o *RecoverMsTeamParams) HasTargetTeamNickName() bool`

HasTargetTeamNickName returns a boolean if a field has been set.

### SetTargetTeamNickNameNil

`func (o *RecoverMsTeamParams) SetTargetTeamNickNameNil(b bool)`

 SetTargetTeamNickNameNil sets the value for TargetTeamNickName to be an explicit nil

### UnsetTargetTeamNickName
`func (o *RecoverMsTeamParams) UnsetTargetTeamNickName()`

UnsetTargetTeamNickName ensures that no value is present for TargetTeamNickName, not even an explicit nil
### GetTargetTeamOwner

`func (o *RecoverMsTeamParams) GetTargetTeamOwner() RecoverMsTeamParamsTargetTeamOwner`

GetTargetTeamOwner returns the TargetTeamOwner field if non-nil, zero value otherwise.

### GetTargetTeamOwnerOk

`func (o *RecoverMsTeamParams) GetTargetTeamOwnerOk() (*RecoverMsTeamParamsTargetTeamOwner, bool)`

GetTargetTeamOwnerOk returns a tuple with the TargetTeamOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTeamOwner

`func (o *RecoverMsTeamParams) SetTargetTeamOwner(v RecoverMsTeamParamsTargetTeamOwner)`

SetTargetTeamOwner sets TargetTeamOwner field to given value.

### HasTargetTeamOwner

`func (o *RecoverMsTeamParams) HasTargetTeamOwner() bool`

HasTargetTeamOwner returns a boolean if a field has been set.

### SetTargetTeamOwnerNil

`func (o *RecoverMsTeamParams) SetTargetTeamOwnerNil(b bool)`

 SetTargetTeamOwnerNil sets the value for TargetTeamOwner to be an explicit nil

### UnsetTargetTeamOwner
`func (o *RecoverMsTeamParams) UnsetTargetTeamOwner()`

UnsetTargetTeamOwner ensures that no value is present for TargetTeamOwner, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


