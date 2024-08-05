# ObjectMsTeamParamMsTeamParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelParams** | Pointer to [**[]ChannelParam**](ChannelParam.md) | Specifies the list of Channels to recover. These are applicable iff recoverEntireMsTeam is false. | [optional] 
**RecoverEntireMsTeam** | **NullableBool** | Specifies whether to recover the whole Microsoft 365 Team. | 

## Methods

### NewObjectMsTeamParamMsTeamParam

`func NewObjectMsTeamParamMsTeamParam(recoverEntireMsTeam NullableBool, ) *ObjectMsTeamParamMsTeamParam`

NewObjectMsTeamParamMsTeamParam instantiates a new ObjectMsTeamParamMsTeamParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectMsTeamParamMsTeamParamWithDefaults

`func NewObjectMsTeamParamMsTeamParamWithDefaults() *ObjectMsTeamParamMsTeamParam`

NewObjectMsTeamParamMsTeamParamWithDefaults instantiates a new ObjectMsTeamParamMsTeamParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelParams

`func (o *ObjectMsTeamParamMsTeamParam) GetChannelParams() []ChannelParam`

GetChannelParams returns the ChannelParams field if non-nil, zero value otherwise.

### GetChannelParamsOk

`func (o *ObjectMsTeamParamMsTeamParam) GetChannelParamsOk() (*[]ChannelParam, bool)`

GetChannelParamsOk returns a tuple with the ChannelParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelParams

`func (o *ObjectMsTeamParamMsTeamParam) SetChannelParams(v []ChannelParam)`

SetChannelParams sets ChannelParams field to given value.

### HasChannelParams

`func (o *ObjectMsTeamParamMsTeamParam) HasChannelParams() bool`

HasChannelParams returns a boolean if a field has been set.

### SetChannelParamsNil

`func (o *ObjectMsTeamParamMsTeamParam) SetChannelParamsNil(b bool)`

 SetChannelParamsNil sets the value for ChannelParams to be an explicit nil

### UnsetChannelParams
`func (o *ObjectMsTeamParamMsTeamParam) UnsetChannelParams()`

UnsetChannelParams ensures that no value is present for ChannelParams, not even an explicit nil
### GetRecoverEntireMsTeam

`func (o *ObjectMsTeamParamMsTeamParam) GetRecoverEntireMsTeam() bool`

GetRecoverEntireMsTeam returns the RecoverEntireMsTeam field if non-nil, zero value otherwise.

### GetRecoverEntireMsTeamOk

`func (o *ObjectMsTeamParamMsTeamParam) GetRecoverEntireMsTeamOk() (*bool, bool)`

GetRecoverEntireMsTeamOk returns a tuple with the RecoverEntireMsTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverEntireMsTeam

`func (o *ObjectMsTeamParamMsTeamParam) SetRecoverEntireMsTeam(v bool)`

SetRecoverEntireMsTeam sets RecoverEntireMsTeam field to given value.


### SetRecoverEntireMsTeamNil

`func (o *ObjectMsTeamParamMsTeamParam) SetRecoverEntireMsTeamNil(b bool)`

 SetRecoverEntireMsTeamNil sets the value for RecoverEntireMsTeam to be an explicit nil

### UnsetRecoverEntireMsTeam
`func (o *ObjectMsTeamParamMsTeamParam) UnsetRecoverEntireMsTeam()`

UnsetRecoverEntireMsTeam ensures that no value is present for RecoverEntireMsTeam, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


