# TargetMsTeamParamTargetTeamsChannelParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelOwners** | Pointer to [**[]RecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | List of owners for the private channel. At least one owner is needed to create a private channel | [optional] 
**ChannelType** | Pointer to **string** | Specifies whether to create a public or private channel | [optional] 
**CreateNewChannel** | Pointer to **NullableBool** | Specifies whether we should create a new channel. If this is true name must not be empty | [optional] 
**Id** | Pointer to **NullableString** | Specifies the id of the target channel. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the target channel. | [optional] 

## Methods

### NewTargetMsTeamParamTargetTeamsChannelParam

`func NewTargetMsTeamParamTargetTeamsChannelParam() *TargetMsTeamParamTargetTeamsChannelParam`

NewTargetMsTeamParamTargetTeamsChannelParam instantiates a new TargetMsTeamParamTargetTeamsChannelParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetMsTeamParamTargetTeamsChannelParamWithDefaults

`func NewTargetMsTeamParamTargetTeamsChannelParamWithDefaults() *TargetMsTeamParamTargetTeamsChannelParam`

NewTargetMsTeamParamTargetTeamsChannelParamWithDefaults instantiates a new TargetMsTeamParamTargetTeamsChannelParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelOwners

`func (o *TargetMsTeamParamTargetTeamsChannelParam) GetChannelOwners() []RecoveryObjectIdentifier`

GetChannelOwners returns the ChannelOwners field if non-nil, zero value otherwise.

### GetChannelOwnersOk

`func (o *TargetMsTeamParamTargetTeamsChannelParam) GetChannelOwnersOk() (*[]RecoveryObjectIdentifier, bool)`

GetChannelOwnersOk returns a tuple with the ChannelOwners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelOwners

`func (o *TargetMsTeamParamTargetTeamsChannelParam) SetChannelOwners(v []RecoveryObjectIdentifier)`

SetChannelOwners sets ChannelOwners field to given value.

### HasChannelOwners

`func (o *TargetMsTeamParamTargetTeamsChannelParam) HasChannelOwners() bool`

HasChannelOwners returns a boolean if a field has been set.

### SetChannelOwnersNil

`func (o *TargetMsTeamParamTargetTeamsChannelParam) SetChannelOwnersNil(b bool)`

 SetChannelOwnersNil sets the value for ChannelOwners to be an explicit nil

### UnsetChannelOwners
`func (o *TargetMsTeamParamTargetTeamsChannelParam) UnsetChannelOwners()`

UnsetChannelOwners ensures that no value is present for ChannelOwners, not even an explicit nil
### GetChannelType

`func (o *TargetMsTeamParamTargetTeamsChannelParam) GetChannelType() string`

GetChannelType returns the ChannelType field if non-nil, zero value otherwise.

### GetChannelTypeOk

`func (o *TargetMsTeamParamTargetTeamsChannelParam) GetChannelTypeOk() (*string, bool)`

GetChannelTypeOk returns a tuple with the ChannelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelType

`func (o *TargetMsTeamParamTargetTeamsChannelParam) SetChannelType(v string)`

SetChannelType sets ChannelType field to given value.

### HasChannelType

`func (o *TargetMsTeamParamTargetTeamsChannelParam) HasChannelType() bool`

HasChannelType returns a boolean if a field has been set.

### GetCreateNewChannel

`func (o *TargetMsTeamParamTargetTeamsChannelParam) GetCreateNewChannel() bool`

GetCreateNewChannel returns the CreateNewChannel field if non-nil, zero value otherwise.

### GetCreateNewChannelOk

`func (o *TargetMsTeamParamTargetTeamsChannelParam) GetCreateNewChannelOk() (*bool, bool)`

GetCreateNewChannelOk returns a tuple with the CreateNewChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateNewChannel

`func (o *TargetMsTeamParamTargetTeamsChannelParam) SetCreateNewChannel(v bool)`

SetCreateNewChannel sets CreateNewChannel field to given value.

### HasCreateNewChannel

`func (o *TargetMsTeamParamTargetTeamsChannelParam) HasCreateNewChannel() bool`

HasCreateNewChannel returns a boolean if a field has been set.

### SetCreateNewChannelNil

`func (o *TargetMsTeamParamTargetTeamsChannelParam) SetCreateNewChannelNil(b bool)`

 SetCreateNewChannelNil sets the value for CreateNewChannel to be an explicit nil

### UnsetCreateNewChannel
`func (o *TargetMsTeamParamTargetTeamsChannelParam) UnsetCreateNewChannel()`

UnsetCreateNewChannel ensures that no value is present for CreateNewChannel, not even an explicit nil
### GetId

`func (o *TargetMsTeamParamTargetTeamsChannelParam) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TargetMsTeamParamTargetTeamsChannelParam) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TargetMsTeamParamTargetTeamsChannelParam) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TargetMsTeamParamTargetTeamsChannelParam) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TargetMsTeamParamTargetTeamsChannelParam) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TargetMsTeamParamTargetTeamsChannelParam) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *TargetMsTeamParamTargetTeamsChannelParam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetMsTeamParamTargetTeamsChannelParam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetMsTeamParamTargetTeamsChannelParam) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TargetMsTeamParamTargetTeamsChannelParam) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TargetMsTeamParamTargetTeamsChannelParam) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TargetMsTeamParamTargetTeamsChannelParam) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


