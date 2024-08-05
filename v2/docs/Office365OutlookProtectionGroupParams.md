# Office365OutlookProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeFolders** | Pointer to **[]string** | Array of prefixes used to exclude folders which are by default included. Two kinds of filters are supported. a) prefix which always starts with &#39;/&#39;. b) posix which always starts with empty quotes(&#39;&#39;). Regular expressions are not supported. If not specified, all folders which are included by default will be included. These prefixes have no effect on folders that are excluded by default. The only folders excluded by default are documented with includeFolders. | [optional] 
**IncludeFolders** | Pointer to **[]string** | Array of prefixes used to include folders which are by default excluded. Two kinds of filters are supported. a) prefix which always starts with &#39;/&#39;. b) posix which always starts with empty quotes(&#39;&#39;). Regular expressions are not supported. If not specified, all folders which are excluded by default will be excluded. These prefixes have no effect on folders that are included by default. All folders are included by default except for the Recoverable Items folder. | [optional] 

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
### GetIncludeFolders

`func (o *Office365OutlookProtectionGroupParams) GetIncludeFolders() []string`

GetIncludeFolders returns the IncludeFolders field if non-nil, zero value otherwise.

### GetIncludeFoldersOk

`func (o *Office365OutlookProtectionGroupParams) GetIncludeFoldersOk() (*[]string, bool)`

GetIncludeFoldersOk returns a tuple with the IncludeFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFolders

`func (o *Office365OutlookProtectionGroupParams) SetIncludeFolders(v []string)`

SetIncludeFolders sets IncludeFolders field to given value.

### HasIncludeFolders

`func (o *Office365OutlookProtectionGroupParams) HasIncludeFolders() bool`

HasIncludeFolders returns a boolean if a field has been set.

### SetIncludeFoldersNil

`func (o *Office365OutlookProtectionGroupParams) SetIncludeFoldersNil(b bool)`

 SetIncludeFoldersNil sets the value for IncludeFolders to be an explicit nil

### UnsetIncludeFolders
`func (o *Office365OutlookProtectionGroupParams) UnsetIncludeFolders()`

UnsetIncludeFolders ensures that no value is present for IncludeFolders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


