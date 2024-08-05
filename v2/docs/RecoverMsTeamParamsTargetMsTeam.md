# RecoverMsTeamParamsTargetMsTeam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentSourceId** | Pointer to **NullableInt64** | Specifies the id of the domain during alternate domain recovery. | [optional] 
**TargetTeam** | Pointer to [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the selected existing target team info. | [optional] 
**TargetTeamsChannelParam** | Pointer to [**TargetMsTeamParamTargetTeamsChannelParam**](TargetMsTeamParamTargetTeamsChannelParam.md) |  | [optional] 

## Methods

### NewRecoverMsTeamParamsTargetMsTeam

`func NewRecoverMsTeamParamsTargetMsTeam() *RecoverMsTeamParamsTargetMsTeam`

NewRecoverMsTeamParamsTargetMsTeam instantiates a new RecoverMsTeamParamsTargetMsTeam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverMsTeamParamsTargetMsTeamWithDefaults

`func NewRecoverMsTeamParamsTargetMsTeamWithDefaults() *RecoverMsTeamParamsTargetMsTeam`

NewRecoverMsTeamParamsTargetMsTeamWithDefaults instantiates a new RecoverMsTeamParamsTargetMsTeam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentSourceId

`func (o *RecoverMsTeamParamsTargetMsTeam) GetParentSourceId() int64`

GetParentSourceId returns the ParentSourceId field if non-nil, zero value otherwise.

### GetParentSourceIdOk

`func (o *RecoverMsTeamParamsTargetMsTeam) GetParentSourceIdOk() (*int64, bool)`

GetParentSourceIdOk returns a tuple with the ParentSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSourceId

`func (o *RecoverMsTeamParamsTargetMsTeam) SetParentSourceId(v int64)`

SetParentSourceId sets ParentSourceId field to given value.

### HasParentSourceId

`func (o *RecoverMsTeamParamsTargetMsTeam) HasParentSourceId() bool`

HasParentSourceId returns a boolean if a field has been set.

### SetParentSourceIdNil

`func (o *RecoverMsTeamParamsTargetMsTeam) SetParentSourceIdNil(b bool)`

 SetParentSourceIdNil sets the value for ParentSourceId to be an explicit nil

### UnsetParentSourceId
`func (o *RecoverMsTeamParamsTargetMsTeam) UnsetParentSourceId()`

UnsetParentSourceId ensures that no value is present for ParentSourceId, not even an explicit nil
### GetTargetTeam

`func (o *RecoverMsTeamParamsTargetMsTeam) GetTargetTeam() RecoveryObjectIdentifier`

GetTargetTeam returns the TargetTeam field if non-nil, zero value otherwise.

### GetTargetTeamOk

`func (o *RecoverMsTeamParamsTargetMsTeam) GetTargetTeamOk() (*RecoveryObjectIdentifier, bool)`

GetTargetTeamOk returns a tuple with the TargetTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTeam

`func (o *RecoverMsTeamParamsTargetMsTeam) SetTargetTeam(v RecoveryObjectIdentifier)`

SetTargetTeam sets TargetTeam field to given value.

### HasTargetTeam

`func (o *RecoverMsTeamParamsTargetMsTeam) HasTargetTeam() bool`

HasTargetTeam returns a boolean if a field has been set.

### SetTargetTeamNil

`func (o *RecoverMsTeamParamsTargetMsTeam) SetTargetTeamNil(b bool)`

 SetTargetTeamNil sets the value for TargetTeam to be an explicit nil

### UnsetTargetTeam
`func (o *RecoverMsTeamParamsTargetMsTeam) UnsetTargetTeam()`

UnsetTargetTeam ensures that no value is present for TargetTeam, not even an explicit nil
### GetTargetTeamsChannelParam

`func (o *RecoverMsTeamParamsTargetMsTeam) GetTargetTeamsChannelParam() TargetMsTeamParamTargetTeamsChannelParam`

GetTargetTeamsChannelParam returns the TargetTeamsChannelParam field if non-nil, zero value otherwise.

### GetTargetTeamsChannelParamOk

`func (o *RecoverMsTeamParamsTargetMsTeam) GetTargetTeamsChannelParamOk() (*TargetMsTeamParamTargetTeamsChannelParam, bool)`

GetTargetTeamsChannelParamOk returns a tuple with the TargetTeamsChannelParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTeamsChannelParam

`func (o *RecoverMsTeamParamsTargetMsTeam) SetTargetTeamsChannelParam(v TargetMsTeamParamTargetTeamsChannelParam)`

SetTargetTeamsChannelParam sets TargetTeamsChannelParam field to given value.

### HasTargetTeamsChannelParam

`func (o *RecoverMsTeamParamsTargetMsTeam) HasTargetTeamsChannelParam() bool`

HasTargetTeamsChannelParam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


