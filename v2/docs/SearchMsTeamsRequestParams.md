# SearchMsTeamsRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryTypes** | Pointer to **[]string** | Specifies a list of teams files types. Only items within the given types will be returned. | [optional] 
**ChannelNames** | Pointer to **[]string** | Specifies the list of channel names to filter while doing search for files. | [optional] 
**ChannelParams** | Pointer to [**O365TeamsChannelsSearchRequestParams**](O365TeamsChannelsSearchRequestParams.md) |  | [optional] 
**CreationEndTimeSecs** | Pointer to **NullableInt64** | Specifies the end time in Unix timestamp epoch in seconds when the item is created. | [optional] 
**CreationStartTimeSecs** | Pointer to **NullableInt64** | Specifies the start time in Unix timestamp epoch in seconds when the item is created. | [optional] 
**O365Params** | Pointer to [**O365SearchRequestParams**](O365SearchRequestParams.md) |  | [optional] 
**OwnerNames** | Pointer to **[]string** | Specifies the list of owner email ids to filter on owner of the item. | [optional] 
**SearchString** | Pointer to **NullableString** | Specifies the search string to filter the items. User can specify a wildcard character &#39;*&#39; as a suffix to a string where all item names are matched with the prefix string. | [optional] 
**SizeBytesLowerLimit** | Pointer to **NullableInt64** | Specifies the minimum size of the item in bytes. | [optional] 
**SizeBytesUpperLimit** | Pointer to **NullableInt64** | Specifies the maximum size of the item in bytes. | [optional] 
**Types** | Pointer to **[]string** | Specifies a list of Teams item types. Only items within the given types will be returned. | [optional] 

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

### GetCategoryTypes

`func (o *SearchMsTeamsRequestParams) GetCategoryTypes() []string`

GetCategoryTypes returns the CategoryTypes field if non-nil, zero value otherwise.

### GetCategoryTypesOk

`func (o *SearchMsTeamsRequestParams) GetCategoryTypesOk() (*[]string, bool)`

GetCategoryTypesOk returns a tuple with the CategoryTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryTypes

`func (o *SearchMsTeamsRequestParams) SetCategoryTypes(v []string)`

SetCategoryTypes sets CategoryTypes field to given value.

### HasCategoryTypes

`func (o *SearchMsTeamsRequestParams) HasCategoryTypes() bool`

HasCategoryTypes returns a boolean if a field has been set.

### SetCategoryTypesNil

`func (o *SearchMsTeamsRequestParams) SetCategoryTypesNil(b bool)`

 SetCategoryTypesNil sets the value for CategoryTypes to be an explicit nil

### UnsetCategoryTypes
`func (o *SearchMsTeamsRequestParams) UnsetCategoryTypes()`

UnsetCategoryTypes ensures that no value is present for CategoryTypes, not even an explicit nil
### GetChannelNames

`func (o *SearchMsTeamsRequestParams) GetChannelNames() []string`

GetChannelNames returns the ChannelNames field if non-nil, zero value otherwise.

### GetChannelNamesOk

`func (o *SearchMsTeamsRequestParams) GetChannelNamesOk() (*[]string, bool)`

GetChannelNamesOk returns a tuple with the ChannelNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelNames

`func (o *SearchMsTeamsRequestParams) SetChannelNames(v []string)`

SetChannelNames sets ChannelNames field to given value.

### HasChannelNames

`func (o *SearchMsTeamsRequestParams) HasChannelNames() bool`

HasChannelNames returns a boolean if a field has been set.

### SetChannelNamesNil

`func (o *SearchMsTeamsRequestParams) SetChannelNamesNil(b bool)`

 SetChannelNamesNil sets the value for ChannelNames to be an explicit nil

### UnsetChannelNames
`func (o *SearchMsTeamsRequestParams) UnsetChannelNames()`

UnsetChannelNames ensures that no value is present for ChannelNames, not even an explicit nil
### GetChannelParams

`func (o *SearchMsTeamsRequestParams) GetChannelParams() O365TeamsChannelsSearchRequestParams`

GetChannelParams returns the ChannelParams field if non-nil, zero value otherwise.

### GetChannelParamsOk

`func (o *SearchMsTeamsRequestParams) GetChannelParamsOk() (*O365TeamsChannelsSearchRequestParams, bool)`

GetChannelParamsOk returns a tuple with the ChannelParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelParams

`func (o *SearchMsTeamsRequestParams) SetChannelParams(v O365TeamsChannelsSearchRequestParams)`

SetChannelParams sets ChannelParams field to given value.

### HasChannelParams

`func (o *SearchMsTeamsRequestParams) HasChannelParams() bool`

HasChannelParams returns a boolean if a field has been set.

### GetCreationEndTimeSecs

`func (o *SearchMsTeamsRequestParams) GetCreationEndTimeSecs() int64`

GetCreationEndTimeSecs returns the CreationEndTimeSecs field if non-nil, zero value otherwise.

### GetCreationEndTimeSecsOk

`func (o *SearchMsTeamsRequestParams) GetCreationEndTimeSecsOk() (*int64, bool)`

GetCreationEndTimeSecsOk returns a tuple with the CreationEndTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationEndTimeSecs

`func (o *SearchMsTeamsRequestParams) SetCreationEndTimeSecs(v int64)`

SetCreationEndTimeSecs sets CreationEndTimeSecs field to given value.

### HasCreationEndTimeSecs

`func (o *SearchMsTeamsRequestParams) HasCreationEndTimeSecs() bool`

HasCreationEndTimeSecs returns a boolean if a field has been set.

### SetCreationEndTimeSecsNil

`func (o *SearchMsTeamsRequestParams) SetCreationEndTimeSecsNil(b bool)`

 SetCreationEndTimeSecsNil sets the value for CreationEndTimeSecs to be an explicit nil

### UnsetCreationEndTimeSecs
`func (o *SearchMsTeamsRequestParams) UnsetCreationEndTimeSecs()`

UnsetCreationEndTimeSecs ensures that no value is present for CreationEndTimeSecs, not even an explicit nil
### GetCreationStartTimeSecs

`func (o *SearchMsTeamsRequestParams) GetCreationStartTimeSecs() int64`

GetCreationStartTimeSecs returns the CreationStartTimeSecs field if non-nil, zero value otherwise.

### GetCreationStartTimeSecsOk

`func (o *SearchMsTeamsRequestParams) GetCreationStartTimeSecsOk() (*int64, bool)`

GetCreationStartTimeSecsOk returns a tuple with the CreationStartTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationStartTimeSecs

`func (o *SearchMsTeamsRequestParams) SetCreationStartTimeSecs(v int64)`

SetCreationStartTimeSecs sets CreationStartTimeSecs field to given value.

### HasCreationStartTimeSecs

`func (o *SearchMsTeamsRequestParams) HasCreationStartTimeSecs() bool`

HasCreationStartTimeSecs returns a boolean if a field has been set.

### SetCreationStartTimeSecsNil

`func (o *SearchMsTeamsRequestParams) SetCreationStartTimeSecsNil(b bool)`

 SetCreationStartTimeSecsNil sets the value for CreationStartTimeSecs to be an explicit nil

### UnsetCreationStartTimeSecs
`func (o *SearchMsTeamsRequestParams) UnsetCreationStartTimeSecs()`

UnsetCreationStartTimeSecs ensures that no value is present for CreationStartTimeSecs, not even an explicit nil
### GetO365Params

`func (o *SearchMsTeamsRequestParams) GetO365Params() O365SearchRequestParams`

GetO365Params returns the O365Params field if non-nil, zero value otherwise.

### GetO365ParamsOk

`func (o *SearchMsTeamsRequestParams) GetO365ParamsOk() (*O365SearchRequestParams, bool)`

GetO365ParamsOk returns a tuple with the O365Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetO365Params

`func (o *SearchMsTeamsRequestParams) SetO365Params(v O365SearchRequestParams)`

SetO365Params sets O365Params field to given value.

### HasO365Params

`func (o *SearchMsTeamsRequestParams) HasO365Params() bool`

HasO365Params returns a boolean if a field has been set.

### GetOwnerNames

`func (o *SearchMsTeamsRequestParams) GetOwnerNames() []string`

GetOwnerNames returns the OwnerNames field if non-nil, zero value otherwise.

### GetOwnerNamesOk

`func (o *SearchMsTeamsRequestParams) GetOwnerNamesOk() (*[]string, bool)`

GetOwnerNamesOk returns a tuple with the OwnerNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerNames

`func (o *SearchMsTeamsRequestParams) SetOwnerNames(v []string)`

SetOwnerNames sets OwnerNames field to given value.

### HasOwnerNames

`func (o *SearchMsTeamsRequestParams) HasOwnerNames() bool`

HasOwnerNames returns a boolean if a field has been set.

### SetOwnerNamesNil

`func (o *SearchMsTeamsRequestParams) SetOwnerNamesNil(b bool)`

 SetOwnerNamesNil sets the value for OwnerNames to be an explicit nil

### UnsetOwnerNames
`func (o *SearchMsTeamsRequestParams) UnsetOwnerNames()`

UnsetOwnerNames ensures that no value is present for OwnerNames, not even an explicit nil
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
### GetSizeBytesLowerLimit

`func (o *SearchMsTeamsRequestParams) GetSizeBytesLowerLimit() int64`

GetSizeBytesLowerLimit returns the SizeBytesLowerLimit field if non-nil, zero value otherwise.

### GetSizeBytesLowerLimitOk

`func (o *SearchMsTeamsRequestParams) GetSizeBytesLowerLimitOk() (*int64, bool)`

GetSizeBytesLowerLimitOk returns a tuple with the SizeBytesLowerLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytesLowerLimit

`func (o *SearchMsTeamsRequestParams) SetSizeBytesLowerLimit(v int64)`

SetSizeBytesLowerLimit sets SizeBytesLowerLimit field to given value.

### HasSizeBytesLowerLimit

`func (o *SearchMsTeamsRequestParams) HasSizeBytesLowerLimit() bool`

HasSizeBytesLowerLimit returns a boolean if a field has been set.

### SetSizeBytesLowerLimitNil

`func (o *SearchMsTeamsRequestParams) SetSizeBytesLowerLimitNil(b bool)`

 SetSizeBytesLowerLimitNil sets the value for SizeBytesLowerLimit to be an explicit nil

### UnsetSizeBytesLowerLimit
`func (o *SearchMsTeamsRequestParams) UnsetSizeBytesLowerLimit()`

UnsetSizeBytesLowerLimit ensures that no value is present for SizeBytesLowerLimit, not even an explicit nil
### GetSizeBytesUpperLimit

`func (o *SearchMsTeamsRequestParams) GetSizeBytesUpperLimit() int64`

GetSizeBytesUpperLimit returns the SizeBytesUpperLimit field if non-nil, zero value otherwise.

### GetSizeBytesUpperLimitOk

`func (o *SearchMsTeamsRequestParams) GetSizeBytesUpperLimitOk() (*int64, bool)`

GetSizeBytesUpperLimitOk returns a tuple with the SizeBytesUpperLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytesUpperLimit

`func (o *SearchMsTeamsRequestParams) SetSizeBytesUpperLimit(v int64)`

SetSizeBytesUpperLimit sets SizeBytesUpperLimit field to given value.

### HasSizeBytesUpperLimit

`func (o *SearchMsTeamsRequestParams) HasSizeBytesUpperLimit() bool`

HasSizeBytesUpperLimit returns a boolean if a field has been set.

### SetSizeBytesUpperLimitNil

`func (o *SearchMsTeamsRequestParams) SetSizeBytesUpperLimitNil(b bool)`

 SetSizeBytesUpperLimitNil sets the value for SizeBytesUpperLimit to be an explicit nil

### UnsetSizeBytesUpperLimit
`func (o *SearchMsTeamsRequestParams) UnsetSizeBytesUpperLimit()`

UnsetSizeBytesUpperLimit ensures that no value is present for SizeBytesUpperLimit, not even an explicit nil
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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


