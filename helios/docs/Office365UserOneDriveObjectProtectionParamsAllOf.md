# Office365UserOneDriveObjectProtectionParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeFolders** | Pointer to **[]string** | Specifies filters to match OneDrive folders which should be excluded when backing up office365 onedrive source. Two kinds of filters are supported. a) prefix which always starts with &#39;/&#39;. b) posix which always starts with empty quotes(&#39;&#39;). Regular expressions are not supported. If not specified, all the OneDrive will be protected. | [optional] 

## Methods

### NewOffice365UserOneDriveObjectProtectionParamsAllOf

`func NewOffice365UserOneDriveObjectProtectionParamsAllOf() *Office365UserOneDriveObjectProtectionParamsAllOf`

NewOffice365UserOneDriveObjectProtectionParamsAllOf instantiates a new Office365UserOneDriveObjectProtectionParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOffice365UserOneDriveObjectProtectionParamsAllOfWithDefaults

`func NewOffice365UserOneDriveObjectProtectionParamsAllOfWithDefaults() *Office365UserOneDriveObjectProtectionParamsAllOf`

NewOffice365UserOneDriveObjectProtectionParamsAllOfWithDefaults instantiates a new Office365UserOneDriveObjectProtectionParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeFolders

`func (o *Office365UserOneDriveObjectProtectionParamsAllOf) GetExcludeFolders() []string`

GetExcludeFolders returns the ExcludeFolders field if non-nil, zero value otherwise.

### GetExcludeFoldersOk

`func (o *Office365UserOneDriveObjectProtectionParamsAllOf) GetExcludeFoldersOk() (*[]string, bool)`

GetExcludeFoldersOk returns a tuple with the ExcludeFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFolders

`func (o *Office365UserOneDriveObjectProtectionParamsAllOf) SetExcludeFolders(v []string)`

SetExcludeFolders sets ExcludeFolders field to given value.

### HasExcludeFolders

`func (o *Office365UserOneDriveObjectProtectionParamsAllOf) HasExcludeFolders() bool`

HasExcludeFolders returns a boolean if a field has been set.

### SetExcludeFoldersNil

`func (o *Office365UserOneDriveObjectProtectionParamsAllOf) SetExcludeFoldersNil(b bool)`

 SetExcludeFoldersNil sets the value for ExcludeFolders to be an explicit nil

### UnsetExcludeFolders
`func (o *Office365UserOneDriveObjectProtectionParamsAllOf) UnsetExcludeFolders()`

UnsetExcludeFolders ensures that no value is present for ExcludeFolders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


