# TargetMsTeamParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentSourceId** | Pointer to **NullableInt64** | Specifies the id of the domain during alternate domain recovery. | [optional] 
**TargetTeam** | Pointer to [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the selected existing target team info. | [optional] 
**TargetTeamsChannelParam** | Pointer to [**TargetMsTeamParamTargetTeamsChannelParam**](TargetMsTeamParamTargetTeamsChannelParam.md) |  | [optional] 

## Methods

### NewTargetMsTeamParam

`func NewTargetMsTeamParam() *TargetMsTeamParam`

NewTargetMsTeamParam instantiates a new TargetMsTeamParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetMsTeamParamWithDefaults

`func NewTargetMsTeamParamWithDefaults() *TargetMsTeamParam`

NewTargetMsTeamParamWithDefaults instantiates a new TargetMsTeamParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentSourceId

`func (o *TargetMsTeamParam) GetParentSourceId() int64`

GetParentSourceId returns the ParentSourceId field if non-nil, zero value otherwise.

### GetParentSourceIdOk

`func (o *TargetMsTeamParam) GetParentSourceIdOk() (*int64, bool)`

GetParentSourceIdOk returns a tuple with the ParentSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSourceId

`func (o *TargetMsTeamParam) SetParentSourceId(v int64)`

SetParentSourceId sets ParentSourceId field to given value.

### HasParentSourceId

`func (o *TargetMsTeamParam) HasParentSourceId() bool`

HasParentSourceId returns a boolean if a field has been set.

### SetParentSourceIdNil

`func (o *TargetMsTeamParam) SetParentSourceIdNil(b bool)`

 SetParentSourceIdNil sets the value for ParentSourceId to be an explicit nil

### UnsetParentSourceId
`func (o *TargetMsTeamParam) UnsetParentSourceId()`

UnsetParentSourceId ensures that no value is present for ParentSourceId, not even an explicit nil
### GetTargetTeam

`func (o *TargetMsTeamParam) GetTargetTeam() RecoveryObjectIdentifier`

GetTargetTeam returns the TargetTeam field if non-nil, zero value otherwise.

### GetTargetTeamOk

`func (o *TargetMsTeamParam) GetTargetTeamOk() (*RecoveryObjectIdentifier, bool)`

GetTargetTeamOk returns a tuple with the TargetTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTeam

`func (o *TargetMsTeamParam) SetTargetTeam(v RecoveryObjectIdentifier)`

SetTargetTeam sets TargetTeam field to given value.

### HasTargetTeam

`func (o *TargetMsTeamParam) HasTargetTeam() bool`

HasTargetTeam returns a boolean if a field has been set.

### SetTargetTeamNil

`func (o *TargetMsTeamParam) SetTargetTeamNil(b bool)`

 SetTargetTeamNil sets the value for TargetTeam to be an explicit nil

### UnsetTargetTeam
`func (o *TargetMsTeamParam) UnsetTargetTeam()`

UnsetTargetTeam ensures that no value is present for TargetTeam, not even an explicit nil
### GetTargetTeamsChannelParam

`func (o *TargetMsTeamParam) GetTargetTeamsChannelParam() TargetMsTeamParamTargetTeamsChannelParam`

GetTargetTeamsChannelParam returns the TargetTeamsChannelParam field if non-nil, zero value otherwise.

### GetTargetTeamsChannelParamOk

`func (o *TargetMsTeamParam) GetTargetTeamsChannelParamOk() (*TargetMsTeamParamTargetTeamsChannelParam, bool)`

GetTargetTeamsChannelParamOk returns a tuple with the TargetTeamsChannelParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTeamsChannelParam

`func (o *TargetMsTeamParam) SetTargetTeamsChannelParam(v TargetMsTeamParamTargetTeamsChannelParam)`

SetTargetTeamsChannelParam sets TargetTeamsChannelParam field to given value.

### HasTargetTeamsChannelParam

`func (o *TargetMsTeamParam) HasTargetTeamsChannelParam() bool`

HasTargetTeamsChannelParam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


