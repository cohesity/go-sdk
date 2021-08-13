# Office365OneDriveProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeFolders** | Pointer to **[]string** | Array of Excluded OneDrive folders. Specifies filters to match OneDrive folders which should be excluded when backing up Office 365 source. Two kinds of filters are supported. a) prefix which always starts with &#39;/&#39;. b) posix which always starts with empty quotes(&#39;&#39;). Regular expressions are not supported. If not specified, all the mailboxes will be protected. | [optional] 

## Methods

### NewOffice365OneDriveProtectionGroupParams

`func NewOffice365OneDriveProtectionGroupParams() *Office365OneDriveProtectionGroupParams`

NewOffice365OneDriveProtectionGroupParams instantiates a new Office365OneDriveProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOffice365OneDriveProtectionGroupParamsWithDefaults

`func NewOffice365OneDriveProtectionGroupParamsWithDefaults() *Office365OneDriveProtectionGroupParams`

NewOffice365OneDriveProtectionGroupParamsWithDefaults instantiates a new Office365OneDriveProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeFolders

`func (o *Office365OneDriveProtectionGroupParams) GetExcludeFolders() []string`

GetExcludeFolders returns the ExcludeFolders field if non-nil, zero value otherwise.

### GetExcludeFoldersOk

`func (o *Office365OneDriveProtectionGroupParams) GetExcludeFoldersOk() (*[]string, bool)`

GetExcludeFoldersOk returns a tuple with the ExcludeFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFolders

`func (o *Office365OneDriveProtectionGroupParams) SetExcludeFolders(v []string)`

SetExcludeFolders sets ExcludeFolders field to given value.

### HasExcludeFolders

`func (o *Office365OneDriveProtectionGroupParams) HasExcludeFolders() bool`

HasExcludeFolders returns a boolean if a field has been set.

### SetExcludeFoldersNil

`func (o *Office365OneDriveProtectionGroupParams) SetExcludeFoldersNil(b bool)`

 SetExcludeFoldersNil sets the value for ExcludeFolders to be an explicit nil

### UnsetExcludeFolders
`func (o *Office365OneDriveProtectionGroupParams) UnsetExcludeFolders()`

UnsetExcludeFolders ensures that no value is present for ExcludeFolders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


