# ChannelItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelEmail** | Pointer to **NullableString** | Specifies the email of this channel. | [optional] 
**ChannelId** | Pointer to **NullableString** | Specifies the id of this channel. | [optional] 
**ChannelName** | Pointer to **NullableString** | Specifies the channel name. | [optional] 
**ChannelType** | Pointer to **NullableString** | Specifies the channel type. | [optional] 
**CreationTimeSecs** | Pointer to **NullableInt64** | Specifies the Unix timestamp epoch in seconds at which this channel is created. | [optional] 
**OwnerNames** | Pointer to **[]string** | Specifies the names of owners of this channel. | [optional] 

## Methods

### NewChannelItem

`func NewChannelItem() *ChannelItem`

NewChannelItem instantiates a new ChannelItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelItemWithDefaults

`func NewChannelItemWithDefaults() *ChannelItem`

NewChannelItemWithDefaults instantiates a new ChannelItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelEmail

`func (o *ChannelItem) GetChannelEmail() string`

GetChannelEmail returns the ChannelEmail field if non-nil, zero value otherwise.

### GetChannelEmailOk

`func (o *ChannelItem) GetChannelEmailOk() (*string, bool)`

GetChannelEmailOk returns a tuple with the ChannelEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelEmail

`func (o *ChannelItem) SetChannelEmail(v string)`

SetChannelEmail sets ChannelEmail field to given value.

### HasChannelEmail

`func (o *ChannelItem) HasChannelEmail() bool`

HasChannelEmail returns a boolean if a field has been set.

### SetChannelEmailNil

`func (o *ChannelItem) SetChannelEmailNil(b bool)`

 SetChannelEmailNil sets the value for ChannelEmail to be an explicit nil

### UnsetChannelEmail
`func (o *ChannelItem) UnsetChannelEmail()`

UnsetChannelEmail ensures that no value is present for ChannelEmail, not even an explicit nil
### GetChannelId

`func (o *ChannelItem) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *ChannelItem) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *ChannelItem) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *ChannelItem) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### SetChannelIdNil

`func (o *ChannelItem) SetChannelIdNil(b bool)`

 SetChannelIdNil sets the value for ChannelId to be an explicit nil

### UnsetChannelId
`func (o *ChannelItem) UnsetChannelId()`

UnsetChannelId ensures that no value is present for ChannelId, not even an explicit nil
### GetChannelName

`func (o *ChannelItem) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *ChannelItem) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *ChannelItem) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.

### HasChannelName

`func (o *ChannelItem) HasChannelName() bool`

HasChannelName returns a boolean if a field has been set.

### SetChannelNameNil

`func (o *ChannelItem) SetChannelNameNil(b bool)`

 SetChannelNameNil sets the value for ChannelName to be an explicit nil

### UnsetChannelName
`func (o *ChannelItem) UnsetChannelName()`

UnsetChannelName ensures that no value is present for ChannelName, not even an explicit nil
### GetChannelType

`func (o *ChannelItem) GetChannelType() string`

GetChannelType returns the ChannelType field if non-nil, zero value otherwise.

### GetChannelTypeOk

`func (o *ChannelItem) GetChannelTypeOk() (*string, bool)`

GetChannelTypeOk returns a tuple with the ChannelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelType

`func (o *ChannelItem) SetChannelType(v string)`

SetChannelType sets ChannelType field to given value.

### HasChannelType

`func (o *ChannelItem) HasChannelType() bool`

HasChannelType returns a boolean if a field has been set.

### SetChannelTypeNil

`func (o *ChannelItem) SetChannelTypeNil(b bool)`

 SetChannelTypeNil sets the value for ChannelType to be an explicit nil

### UnsetChannelType
`func (o *ChannelItem) UnsetChannelType()`

UnsetChannelType ensures that no value is present for ChannelType, not even an explicit nil
### GetCreationTimeSecs

`func (o *ChannelItem) GetCreationTimeSecs() int64`

GetCreationTimeSecs returns the CreationTimeSecs field if non-nil, zero value otherwise.

### GetCreationTimeSecsOk

`func (o *ChannelItem) GetCreationTimeSecsOk() (*int64, bool)`

GetCreationTimeSecsOk returns a tuple with the CreationTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimeSecs

`func (o *ChannelItem) SetCreationTimeSecs(v int64)`

SetCreationTimeSecs sets CreationTimeSecs field to given value.

### HasCreationTimeSecs

`func (o *ChannelItem) HasCreationTimeSecs() bool`

HasCreationTimeSecs returns a boolean if a field has been set.

### SetCreationTimeSecsNil

`func (o *ChannelItem) SetCreationTimeSecsNil(b bool)`

 SetCreationTimeSecsNil sets the value for CreationTimeSecs to be an explicit nil

### UnsetCreationTimeSecs
`func (o *ChannelItem) UnsetCreationTimeSecs()`

UnsetCreationTimeSecs ensures that no value is present for CreationTimeSecs, not even an explicit nil
### GetOwnerNames

`func (o *ChannelItem) GetOwnerNames() []string`

GetOwnerNames returns the OwnerNames field if non-nil, zero value otherwise.

### GetOwnerNamesOk

`func (o *ChannelItem) GetOwnerNamesOk() (*[]string, bool)`

GetOwnerNamesOk returns a tuple with the OwnerNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerNames

`func (o *ChannelItem) SetOwnerNames(v []string)`

SetOwnerNames sets OwnerNames field to given value.

### HasOwnerNames

`func (o *ChannelItem) HasOwnerNames() bool`

HasOwnerNames returns a boolean if a field has been set.

### SetOwnerNamesNil

`func (o *ChannelItem) SetOwnerNamesNil(b bool)`

 SetOwnerNamesNil sets the value for OwnerNames to be an explicit nil

### UnsetOwnerNames
`func (o *ChannelItem) UnsetOwnerNames()`

UnsetOwnerNames ensures that no value is present for OwnerNames, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


