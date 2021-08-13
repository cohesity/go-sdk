# SearchMsTeamsRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchString** | Pointer to **NullableString** | Specifies the search string to filter the items. User can specify a wildcard character &#39;*&#39; as a suffix to a string where all item names are matched with the prefix string. | [optional] 
**Types** | Pointer to **[]string** | Specifies a list of Teams item types. Only items within the given types will be returned. | [optional] 
**ChannelName** | Pointer to **NullableString** | Specifies the name of the channel. Only items within the specified channel will be returned. | [optional] 

## Methods

### NewSearchMsTeamsRequestParams

`func NewSearchMsTeamsRequestParams() *SearchMsTeamsRequestParams`

NewSearchMsTeamsRequestParams instantiates a new SearchMsTeamsRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchMsTeamsRequestParamsWithDefaults

`func NewSearchMsTeamsRequestParamsWithDefaults() *SearchMsTeamsRequestParams`

NewSearchMsTeamsRequestParamsWithDefaults instantiates a new SearchMsTeamsRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchString

`func (o *SearchMsTeamsRequestParams) GetSearchString() string`

GetSearchString returns the SearchString field if non-nil, zero value otherwise.

### GetSearchStringOk

`func (o *SearchMsTeamsRequestParams) GetSearchStringOk() (*string, bool)`

GetSearchStringOk returns a tuple with the SearchString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchString

`func (o *SearchMsTeamsRequestParams) SetSearchString(v string)`

SetSearchString sets SearchString field to given value.

### HasSearchString

`func (o *SearchMsTeamsRequestParams) HasSearchString() bool`

HasSearchString returns a boolean if a field has been set.

### SetSearchStringNil

`func (o *SearchMsTeamsRequestParams) SetSearchStringNil(b bool)`

 SetSearchStringNil sets the value for SearchString to be an explicit nil

### UnsetSearchString
`func (o *SearchMsTeamsRequestParams) UnsetSearchString()`

UnsetSearchString ensures that no value is present for SearchString, not even an explicit nil
### GetTypes

`func (o *SearchMsTeamsRequestParams) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *SearchMsTeamsRequestParams) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *SearchMsTeamsRequestParams) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *SearchMsTeamsRequestParams) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### SetTypesNil

`func (o *SearchMsTeamsRequestParams) SetTypesNil(b bool)`

 SetTypesNil sets the value for Types to be an explicit nil

### UnsetTypes
`func (o *SearchMsTeamsRequestParams) UnsetTypes()`

UnsetTypes ensures that no value is present for Types, not even an explicit nil
### GetChannelName

`func (o *SearchMsTeamsRequestParams) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *SearchMsTeamsRequestParams) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *SearchMsTeamsRequestParams) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.

### HasChannelName

`func (o *SearchMsTeamsRequestParams) HasChannelName() bool`

HasChannelName returns a boolean if a field has been set.

### SetChannelNameNil

`func (o *SearchMsTeamsRequestParams) SetChannelNameNil(b bool)`

 SetChannelNameNil sets the value for ChannelName to be an explicit nil

### UnsetChannelName
`func (o *SearchMsTeamsRequestParams) UnsetChannelName()`

UnsetChannelName ensures that no value is present for ChannelName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


