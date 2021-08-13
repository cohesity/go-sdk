# Office365OutlookProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeFolders** | Pointer to **[]string** | Array of Excluded Outlook folders. Specifies filters to match Outlook folders which should be excluded when backing up Office 365 source. Two kinds of filters are supported. a) prefix which always starts with &#39;/&#39;. b) posix which always starts with empty quotes(&#39;&#39;). Regular expressions are not supported. If not specified, all the mailboxes will be protected. | [optional] 

## Methods

### NewOffice365OutlookProtectionGroupParams

`func NewOffice365OutlookProtectionGroupParams() *Office365OutlookProtectionGroupParams`

NewOffice365OutlookProtectionGroupParams instantiates a new Office365OutlookProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOffice365OutlookProtectionGroupParamsWithDefaults

`func NewOffice365OutlookProtectionGroupParamsWithDefaults() *Office365OutlookProtectionGroupParams`

NewOffice365OutlookProtectionGroupParamsWithDefaults instantiates a new Office365OutlookProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeFolders

`func (o *Office365OutlookProtectionGroupParams) GetExcludeFolders() []string`

GetExcludeFolders returns the ExcludeFolders field if non-nil, zero value otherwise.

### GetExcludeFoldersOk

`func (o *Office365OutlookProtectionGroupParams) GetExcludeFoldersOk() (*[]string, bool)`

GetExcludeFoldersOk returns a tuple with the ExcludeFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFolders

`func (o *Office365OutlookProtectionGroupParams) SetExcludeFolders(v []string)`

SetExcludeFolders sets ExcludeFolders field to given value.

### HasExcludeFolders

`func (o *Office365OutlookProtectionGroupParams) HasExcludeFolders() bool`

HasExcludeFolders returns a boolean if a field has been set.

### SetExcludeFoldersNil

`func (o *Office365OutlookProtectionGroupParams) SetExcludeFoldersNil(b bool)`

 SetExcludeFoldersNil sets the value for ExcludeFolders to be an explicit nil

### UnsetExcludeFolders
`func (o *Office365OutlookProtectionGroupParams) UnsetExcludeFolders()`

UnsetExcludeFolders ensures that no value is present for ExcludeFolders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


