# Office365UserMailboxObjectProtectionParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeFolders** | Pointer to **[]string** | Specifies filters to match Outlook folders which should be excluded when backing up office365 mailbox source. Two kinds of filters are supported. a) prefix which always starts with &#39;/&#39;. b) posix which always starts with empty quotes(&#39;&#39;). Regular expressions are not supported. If not specified, all the mailbox folders will be protected. | [optional] 

## Methods

### NewOffice365UserMailboxObjectProtectionParamsAllOf

`func NewOffice365UserMailboxObjectProtectionParamsAllOf() *Office365UserMailboxObjectProtectionParamsAllOf`

NewOffice365UserMailboxObjectProtectionParamsAllOf instantiates a new Office365UserMailboxObjectProtectionParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOffice365UserMailboxObjectProtectionParamsAllOfWithDefaults

`func NewOffice365UserMailboxObjectProtectionParamsAllOfWithDefaults() *Office365UserMailboxObjectProtectionParamsAllOf`

NewOffice365UserMailboxObjectProtectionParamsAllOfWithDefaults instantiates a new Office365UserMailboxObjectProtectionParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeFolders

`func (o *Office365UserMailboxObjectProtectionParamsAllOf) GetExcludeFolders() []string`

GetExcludeFolders returns the ExcludeFolders field if non-nil, zero value otherwise.

### GetExcludeFoldersOk

`func (o *Office365UserMailboxObjectProtectionParamsAllOf) GetExcludeFoldersOk() (*[]string, bool)`

GetExcludeFoldersOk returns a tuple with the ExcludeFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFolders

`func (o *Office365UserMailboxObjectProtectionParamsAllOf) SetExcludeFolders(v []string)`

SetExcludeFolders sets ExcludeFolders field to given value.

### HasExcludeFolders

`func (o *Office365UserMailboxObjectProtectionParamsAllOf) HasExcludeFolders() bool`

HasExcludeFolders returns a boolean if a field has been set.

### SetExcludeFoldersNil

`func (o *Office365UserMailboxObjectProtectionParamsAllOf) SetExcludeFoldersNil(b bool)`

 SetExcludeFoldersNil sets the value for ExcludeFolders to be an explicit nil

### UnsetExcludeFolders
`func (o *Office365UserMailboxObjectProtectionParamsAllOf) UnsetExcludeFolders()`

UnsetExcludeFolders ensures that no value is present for ExcludeFolders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


