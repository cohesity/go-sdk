# TargetTeamsChannelParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelOwners** | Pointer to [**[]RecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | List of owners for the private channel. At least one owner is needed to create a private channel | [optional] 
**ChannelType** | Pointer to **string** | Specifies whether to create a public or private channel | [optional] 
**CreateNewChannel** | Pointer to **NullableBool** | Specifies whether we should create a new channel. If this is true name must not be empty | [optional] 
**Id** | Pointer to **NullableString** | Specifies the id of the target channel. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the target channel. | [optional] 

## Methods

### NewTargetTeamsChannelParam

`func NewTargetTeamsChannelParam() *TargetTeamsChannelParam`

NewTargetTeamsChannelParam instantiates a new TargetTeamsChannelParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetTeamsChannelParamWithDefaults

`func NewTargetTeamsChannelParamWithDefaults() *TargetTeamsChannelParam`

NewTargetTeamsChannelParamWithDefaults instantiates a new TargetTeamsChannelParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelOwners

`func (o *TargetTeamsChannelParam) GetChannelOwners() []RecoveryObjectIdentifier`

GetChannelOwners returns the ChannelOwners field if non-nil, zero value otherwise.

### GetChannelOwnersOk

`func (o *TargetTeamsChannelParam) GetChannelOwnersOk() (*[]RecoveryObjectIdentifier, bool)`

GetChannelOwnersOk returns a tuple with the ChannelOwners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelOwners

`func (o *TargetTeamsChannelParam) SetChannelOwners(v []RecoveryObjectIdentifier)`

SetChannelOwners sets ChannelOwners field to given value.

### HasChannelOwners

`func (o *TargetTeamsChannelParam) HasChannelOwners() bool`

HasChannelOwners returns a boolean if a field has been set.

### SetChannelOwnersNil

`func (o *TargetTeamsChannelParam) SetChannelOwnersNil(b bool)`

 SetChannelOwnersNil sets the value for ChannelOwners to be an explicit nil

### UnsetChannelOwners
`func (o *TargetTeamsChannelParam) UnsetChannelOwners()`

UnsetChannelOwners ensures that no value is present for ChannelOwners, not even an explicit nil
### GetChannelType

`func (o *TargetTeamsChannelParam) GetChannelType() string`

GetChannelType returns the ChannelType field if non-nil, zero value otherwise.

### GetChannelTypeOk

`func (o *TargetTeamsChannelParam) GetChannelTypeOk() (*string, bool)`

GetChannelTypeOk returns a tuple with the ChannelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelType

`func (o *TargetTeamsChannelParam) SetChannelType(v string)`

SetChannelType sets ChannelType field to given value.

### HasChannelType

`func (o *TargetTeamsChannelParam) HasChannelType() bool`

HasChannelType returns a boolean if a field has been set.

### GetCreateNewChannel

`func (o *TargetTeamsChannelParam) GetCreateNewChannel() bool`

GetCreateNewChannel returns the CreateNewChannel field if non-nil, zero value otherwise.

### GetCreateNewChannelOk

`func (o *TargetTeamsChannelParam) GetCreateNewChannelOk() (*bool, bool)`

GetCreateNewChannelOk returns a tuple with the CreateNewChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateNewChannel

`func (o *TargetTeamsChannelParam) SetCreateNewChannel(v bool)`

SetCreateNewChannel sets CreateNewChannel field to given value.

### HasCreateNewChannel

`func (o *TargetTeamsChannelParam) HasCreateNewChannel() bool`

HasCreateNewChannel returns a boolean if a field has been set.

### SetCreateNewChannelNil

`func (o *TargetTeamsChannelParam) SetCreateNewChannelNil(b bool)`

 SetCreateNewChannelNil sets the value for CreateNewChannel to be an explicit nil

### UnsetCreateNewChannel
`func (o *TargetTeamsChannelParam) UnsetCreateNewChannel()`

UnsetCreateNewChannel ensures that no value is present for CreateNewChannel, not even an explicit nil
### GetId

`func (o *TargetTeamsChannelParam) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TargetTeamsChannelParam) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TargetTeamsChannelParam) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TargetTeamsChannelParam) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TargetTeamsChannelParam) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TargetTeamsChannelParam) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *TargetTeamsChannelParam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetTeamsChannelParam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetTeamsChannelParam) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TargetTeamsChannelParam) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TargetTeamsChannelParam) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TargetTeamsChannelParam) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


