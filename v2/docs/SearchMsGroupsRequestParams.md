# SearchMsGroupsRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MailboxParams** | Pointer to [**SearchEmailRequestParamsBase**](SearchEmailRequestParamsBase.md) |  | [optional] 
**O365Params** | Pointer to [**O365SearchRequestParams**](O365SearchRequestParams.md) |  | [optional] 
**SiteParams** | Pointer to [**SearchDocumentLibraryRequestParams**](SearchDocumentLibraryRequestParams.md) |  | [optional] 

## Methods

### NewSearchMsGroupsRequestParams

`func NewSearchMsGroupsRequestParams() *SearchMsGroupsRequestParams`

NewSearchMsGroupsRequestParams instantiates a new SearchMsGroupsRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchMsGroupsRequestParamsWithDefaults

`func NewSearchMsGroupsRequestParamsWithDefaults() *SearchMsGroupsRequestParams`

NewSearchMsGroupsRequestParamsWithDefaults instantiates a new SearchMsGroupsRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMailboxParams

`func (o *SearchMsGroupsRequestParams) GetMailboxParams() SearchEmailRequestParamsBase`

GetMailboxParams returns the MailboxParams field if non-nil, zero value otherwise.

### GetMailboxParamsOk

`func (o *SearchMsGroupsRequestParams) GetMailboxParamsOk() (*SearchEmailRequestParamsBase, bool)`

GetMailboxParamsOk returns a tuple with the MailboxParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxParams

`func (o *SearchMsGroupsRequestParams) SetMailboxParams(v SearchEmailRequestParamsBase)`

SetMailboxParams sets MailboxParams field to given value.

### HasMailboxParams

`func (o *SearchMsGroupsRequestParams) HasMailboxParams() bool`

HasMailboxParams returns a boolean if a field has been set.

### GetO365Params

`func (o *SearchMsGroupsRequestParams) GetO365Params() O365SearchRequestParams`

GetO365Params returns the O365Params field if non-nil, zero value otherwise.

### GetO365ParamsOk

`func (o *SearchMsGroupsRequestParams) GetO365ParamsOk() (*O365SearchRequestParams, bool)`

GetO365ParamsOk returns a tuple with the O365Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetO365Params

`func (o *SearchMsGroupsRequestParams) SetO365Params(v O365SearchRequestParams)`

SetO365Params sets O365Params field to given value.

### HasO365Params

`func (o *SearchMsGroupsRequestParams) HasO365Params() bool`

HasO365Params returns a boolean if a field has been set.

### GetSiteParams

`func (o *SearchMsGroupsRequestParams) GetSiteParams() SearchDocumentLibraryRequestParams`

GetSiteParams returns the SiteParams field if non-nil, zero value otherwise.

### GetSiteParamsOk

`func (o *SearchMsGroupsRequestParams) GetSiteParamsOk() (*SearchDocumentLibraryRequestParams, bool)`

GetSiteParamsOk returns a tuple with the SiteParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteParams

`func (o *SearchMsGroupsRequestParams) SetSiteParams(v SearchDocumentLibraryRequestParams)`

SetSiteParams sets SiteParams field to given value.

### HasSiteParams

`func (o *SearchMsGroupsRequestParams) HasSiteParams() bool`

HasSiteParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


