# O365TeamsChannelsSearchRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelEmail** | Pointer to **NullableString** | Specifies the email id of the channel. | [optional] 
**ChannelId** | Pointer to **NullableString** | Specifies the unique id of the channel. | [optional] 
**ChannelName** | Pointer to **NullableString** | Specifies the name of the channel. Only items within the specified channel will be returned. | [optional] 
**IncludePrivateChannels** | Pointer to **NullableBool** | Specifies whether to include private channels in the response. Default is true. | [optional] [default to true]
**IncludePublicChannels** | Pointer to **NullableBool** | Specifies whether to include public channels in the response. Default is true. | [optional] [default to true]

## Methods

### NewO365TeamsChannelsSearchRequestParams

`func NewO365TeamsChannelsSearchRequestParams() *O365TeamsChannelsSearchRequestParams`

NewO365TeamsChannelsSearchRequestParams instantiates a new O365TeamsChannelsSearchRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO365TeamsChannelsSearchRequestParamsWithDefaults

`func NewO365TeamsChannelsSearchRequestParamsWithDefaults() *O365TeamsChannelsSearchRequestParams`

NewO365TeamsChannelsSearchRequestParamsWithDefaults instantiates a new O365TeamsChannelsSearchRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelEmail

`func (o *O365TeamsChannelsSearchRequestParams) GetChannelEmail() string`

GetChannelEmail returns the ChannelEmail field if non-nil, zero value otherwise.

### GetChannelEmailOk

`func (o *O365TeamsChannelsSearchRequestParams) GetChannelEmailOk() (*string, bool)`

GetChannelEmailOk returns a tuple with the ChannelEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelEmail

`func (o *O365TeamsChannelsSearchRequestParams) SetChannelEmail(v string)`

SetChannelEmail sets ChannelEmail field to given value.

### HasChannelEmail

`func (o *O365TeamsChannelsSearchRequestParams) HasChannelEmail() bool`

HasChannelEmail returns a boolean if a field has been set.

### SetChannelEmailNil

`func (o *O365TeamsChannelsSearchRequestParams) SetChannelEmailNil(b bool)`

 SetChannelEmailNil sets the value for ChannelEmail to be an explicit nil

### UnsetChannelEmail
`func (o *O365TeamsChannelsSearchRequestParams) UnsetChannelEmail()`

UnsetChannelEmail ensures that no value is present for ChannelEmail, not even an explicit nil
### GetChannelId

`func (o *O365TeamsChannelsSearchRequestParams) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *O365TeamsChannelsSearchRequestParams) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *O365TeamsChannelsSearchRequestParams) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *O365TeamsChannelsSearchRequestParams) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### SetChannelIdNil

`func (o *O365TeamsChannelsSearchRequestParams) SetChannelIdNil(b bool)`

 SetChannelIdNil sets the value for ChannelId to be an explicit nil

### UnsetChannelId
`func (o *O365TeamsChannelsSearchRequestParams) UnsetChannelId()`

UnsetChannelId ensures that no value is present for ChannelId, not even an explicit nil
### GetChannelName

`func (o *O365TeamsChannelsSearchRequestParams) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *O365TeamsChannelsSearchRequestParams) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *O365TeamsChannelsSearchRequestParams) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.

### HasChannelName

`func (o *O365TeamsChannelsSearchRequestParams) HasChannelName() bool`

HasChannelName returns a boolean if a field has been set.

### SetChannelNameNil

`func (o *O365TeamsChannelsSearchRequestParams) SetChannelNameNil(b bool)`

 SetChannelNameNil sets the value for ChannelName to be an explicit nil

### UnsetChannelName
`func (o *O365TeamsChannelsSearchRequestParams) UnsetChannelName()`

UnsetChannelName ensures that no value is present for ChannelName, not even an explicit nil
### GetIncludePrivateChannels

`func (o *O365TeamsChannelsSearchRequestParams) GetIncludePrivateChannels() bool`

GetIncludePrivateChannels returns the IncludePrivateChannels field if non-nil, zero value otherwise.

### GetIncludePrivateChannelsOk

`func (o *O365TeamsChannelsSearchRequestParams) GetIncludePrivateChannelsOk() (*bool, bool)`

GetIncludePrivateChannelsOk returns a tuple with the IncludePrivateChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePrivateChannels

`func (o *O365TeamsChannelsSearchRequestParams) SetIncludePrivateChannels(v bool)`

SetIncludePrivateChannels sets IncludePrivateChannels field to given value.

### HasIncludePrivateChannels

`func (o *O365TeamsChannelsSearchRequestParams) HasIncludePrivateChannels() bool`

HasIncludePrivateChannels returns a boolean if a field has been set.

### SetIncludePrivateChannelsNil

`func (o *O365TeamsChannelsSearchRequestParams) SetIncludePrivateChannelsNil(b bool)`

 SetIncludePrivateChannelsNil sets the value for IncludePrivateChannels to be an explicit nil

### UnsetIncludePrivateChannels
`func (o *O365TeamsChannelsSearchRequestParams) UnsetIncludePrivateChannels()`

UnsetIncludePrivateChannels ensures that no value is present for IncludePrivateChannels, not even an explicit nil
### GetIncludePublicChannels

`func (o *O365TeamsChannelsSearchRequestParams) GetIncludePublicChannels() bool`

GetIncludePublicChannels returns the IncludePublicChannels field if non-nil, zero value otherwise.

### GetIncludePublicChannelsOk

`func (o *O365TeamsChannelsSearchRequestParams) GetIncludePublicChannelsOk() (*bool, bool)`

GetIncludePublicChannelsOk returns a tuple with the IncludePublicChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePublicChannels

`func (o *O365TeamsChannelsSearchRequestParams) SetIncludePublicChannels(v bool)`

SetIncludePublicChannels sets IncludePublicChannels field to given value.

### HasIncludePublicChannels

`func (o *O365TeamsChannelsSearchRequestParams) HasIncludePublicChannels() bool`

HasIncludePublicChannels returns a boolean if a field has been set.

### SetIncludePublicChannelsNil

`func (o *O365TeamsChannelsSearchRequestParams) SetIncludePublicChannelsNil(b bool)`

 SetIncludePublicChannelsNil sets the value for IncludePublicChannels to be an explicit nil

### UnsetIncludePublicChannels
`func (o *O365TeamsChannelsSearchRequestParams) UnsetIncludePublicChannels()`

UnsetIncludePublicChannels ensures that no value is present for IncludePublicChannels, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


