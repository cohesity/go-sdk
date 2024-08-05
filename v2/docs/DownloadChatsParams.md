# DownloadChatsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelIds** | Pointer to **[]string** | Specifies channel IDs whose posts needs to be downloaded. If channelIds is nil or empty then full teams&#39; posts will be downloaded. | [optional] 
**DownloadFileType** | **string** | Specifies the file type for the downloaded content. | 
**HtmlTemplate** | Pointer to **NullableString** | Specifies the html template for the downloaded chats. | [optional] 

## Methods

### NewDownloadChatsParams

`func NewDownloadChatsParams(downloadFileType string, ) *DownloadChatsParams`

NewDownloadChatsParams instantiates a new DownloadChatsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDownloadChatsParamsWithDefaults

`func NewDownloadChatsParamsWithDefaults() *DownloadChatsParams`

NewDownloadChatsParamsWithDefaults instantiates a new DownloadChatsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelIds

`func (o *DownloadChatsParams) GetChannelIds() []string`

GetChannelIds returns the ChannelIds field if non-nil, zero value otherwise.

### GetChannelIdsOk

`func (o *DownloadChatsParams) GetChannelIdsOk() (*[]string, bool)`

GetChannelIdsOk returns a tuple with the ChannelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelIds

`func (o *DownloadChatsParams) SetChannelIds(v []string)`

SetChannelIds sets ChannelIds field to given value.

### HasChannelIds

`func (o *DownloadChatsParams) HasChannelIds() bool`

HasChannelIds returns a boolean if a field has been set.

### SetChannelIdsNil

`func (o *DownloadChatsParams) SetChannelIdsNil(b bool)`

 SetChannelIdsNil sets the value for ChannelIds to be an explicit nil

### UnsetChannelIds
`func (o *DownloadChatsParams) UnsetChannelIds()`

UnsetChannelIds ensures that no value is present for ChannelIds, not even an explicit nil
### GetDownloadFileType

`func (o *DownloadChatsParams) GetDownloadFileType() string`

GetDownloadFileType returns the DownloadFileType field if non-nil, zero value otherwise.

### GetDownloadFileTypeOk

`func (o *DownloadChatsParams) GetDownloadFileTypeOk() (*string, bool)`

GetDownloadFileTypeOk returns a tuple with the DownloadFileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadFileType

`func (o *DownloadChatsParams) SetDownloadFileType(v string)`

SetDownloadFileType sets DownloadFileType field to given value.


### GetHtmlTemplate

`func (o *DownloadChatsParams) GetHtmlTemplate() string`

GetHtmlTemplate returns the HtmlTemplate field if non-nil, zero value otherwise.

### GetHtmlTemplateOk

`func (o *DownloadChatsParams) GetHtmlTemplateOk() (*string, bool)`

GetHtmlTemplateOk returns a tuple with the HtmlTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlTemplate

`func (o *DownloadChatsParams) SetHtmlTemplate(v string)`

SetHtmlTemplate sets HtmlTemplate field to given value.

### HasHtmlTemplate

`func (o *DownloadChatsParams) HasHtmlTemplate() bool`

HasHtmlTemplate returns a boolean if a field has been set.

### SetHtmlTemplateNil

`func (o *DownloadChatsParams) SetHtmlTemplateNil(b bool)`

 SetHtmlTemplateNil sets the value for HtmlTemplate to be an explicit nil

### UnsetHtmlTemplate
`func (o *DownloadChatsParams) UnsetHtmlTemplate()`

UnsetHtmlTemplate ensures that no value is present for HtmlTemplate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


