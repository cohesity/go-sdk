# RecoverMsTeamParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | [**[]ObjectMsTeamParam**](ObjectMsTeamParam.md) | Specifies a list of Microsoft 365 Teams params associated with objects to recover. | 
**TargetMsTeam** | Pointer to [**TargetMsTeamParam**](TargetMsTeamParam.md) | This field is deprecated. Use targetTeamNickName and targetTeamFullName instead. | [optional] 
**TargetTeamNickName** | Pointer to **NullableString** | Specifies target team nickname in case restoreToOriginal is false. | [optional] 
**TargetTeamFullName** | Pointer to **NullableString** | Specifies target team name in case restoreToOriginal is false. This will be ignored if restoring to alternate existing team (i.e. to a team the nickname of which is same as the one supplied by the end user). | [optional] 
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering other teams, if some of the teams fail to recover. Default value is false. | [optional] 

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
### GetTargetMsTeam

`func (o *RecoverMsTeamParams) GetTargetMsTeam() TargetMsTeamParam`

GetTargetMsTeam returns the TargetMsTeam field if non-nil, zero value otherwise.

### GetTargetMsTeamOk

`func (o *RecoverMsTeamParams) GetTargetMsTeamOk() (*TargetMsTeamParam, bool)`

GetTargetMsTeamOk returns a tuple with the TargetMsTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMsTeam

`func (o *RecoverMsTeamParams) SetTargetMsTeam(v TargetMsTeamParam)`

SetTargetMsTeam sets TargetMsTeam field to given value.

### HasTargetMsTeam

`func (o *RecoverMsTeamParams) HasTargetMsTeam() bool`

HasTargetMsTeam returns a boolean if a field has been set.

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


