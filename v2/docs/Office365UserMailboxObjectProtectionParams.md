# Office365UserMailboxObjectProtectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 
**Objects** | [**[]Office365ObjectProtectionObjectParams**](Office365ObjectProtectionObjectParams.md) | Specifies the objects to be included in the Object Protection. | 
**SourceId** | Pointer to **NullableInt64** | Specifies the id of the parent of the objects. | [optional] [readonly] 
**SourceName** | Pointer to **NullableString** | Specifies the name of the parent of the objects. | [optional] [readonly] 
**ExcludeFolders** | Pointer to **[]string** | Array of prefixes used to exclude folders which are by default included. Two kinds of filters are supported. a) prefix which always starts with &#39;/&#39;. b) posix which always starts with empty quotes(&#39;&#39;). Regular expressions are not supported. If not specified, all folders which are included by default will be included. These prefixes have no effect on folders that are excluded by default. The only folders excluded by default are documented with includeFolders. | [optional] 
**IncludeFolders** | Pointer to **[]string** | Array of prefixes used to include folders which are by default excluded. Two kinds of filters are supported. a) prefix which always starts with &#39;/&#39;. b) posix which always starts with empty quotes(&#39;&#39;). Regular expressions are not supported. If not specified, all folders which are excluded by default will be excluded. These prefixes have no effect on folders that are included by default. All folders are included by default except for the Recoverable Items folder. | [optional] 

## Methods

### NewOffice365UserMailboxObjectProtectionParams

`func NewOffice365UserMailboxObjectProtectionParams(objects []Office365ObjectProtectionObjectParams, ) *Office365UserMailboxObjectProtectionParams`

NewOffice365UserMailboxObjectProtectionParams instantiates a new Office365UserMailboxObjectProtectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOffice365UserMailboxObjectProtectionParamsWithDefaults

`func NewOffice365UserMailboxObjectProtectionParamsWithDefaults() *Office365UserMailboxObjectProtectionParams`

NewOffice365UserMailboxObjectProtectionParamsWithDefaults instantiates a new Office365UserMailboxObjectProtectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexingPolicy

`func (o *Office365UserMailboxObjectProtectionParams) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *Office365UserMailboxObjectProtectionParams) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *Office365UserMailboxObjectProtectionParams) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *Office365UserMailboxObjectProtectionParams) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetObjects

`func (o *Office365UserMailboxObjectProtectionParams) GetObjects() []Office365ObjectProtectionObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *Office365UserMailboxObjectProtectionParams) GetObjectsOk() (*[]Office365ObjectProtectionObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *Office365UserMailboxObjectProtectionParams) SetObjects(v []Office365ObjectProtectionObjectParams)`

SetObjects sets Objects field to given value.


### GetSourceId

`func (o *Office365UserMailboxObjectProtectionParams) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *Office365UserMailboxObjectProtectionParams) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *Office365UserMailboxObjectProtectionParams) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *Office365UserMailboxObjectProtectionParams) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *Office365UserMailboxObjectProtectionParams) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *Office365UserMailboxObjectProtectionParams) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *Office365UserMailboxObjectProtectionParams) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *Office365UserMailboxObjectProtectionParams) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *Office365UserMailboxObjectProtectionParams) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *Office365UserMailboxObjectProtectionParams) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *Office365UserMailboxObjectProtectionParams) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *Office365UserMailboxObjectProtectionParams) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetExcludeFolders

`func (o *Office365UserMailboxObjectProtectionParams) GetExcludeFolders() []string`

GetExcludeFolders returns the ExcludeFolders field if non-nil, zero value otherwise.

### GetExcludeFoldersOk

`func (o *Office365UserMailboxObjectProtectionParams) GetExcludeFoldersOk() (*[]string, bool)`

GetExcludeFoldersOk returns a tuple with the ExcludeFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFolders

`func (o *Office365UserMailboxObjectProtectionParams) SetExcludeFolders(v []string)`

SetExcludeFolders sets ExcludeFolders field to given value.

### HasExcludeFolders

`func (o *Office365UserMailboxObjectProtectionParams) HasExcludeFolders() bool`

HasExcludeFolders returns a boolean if a field has been set.

### SetExcludeFoldersNil

`func (o *Office365UserMailboxObjectProtectionParams) SetExcludeFoldersNil(b bool)`

 SetExcludeFoldersNil sets the value for ExcludeFolders to be an explicit nil

### UnsetExcludeFolders
`func (o *Office365UserMailboxObjectProtectionParams) UnsetExcludeFolders()`

UnsetExcludeFolders ensures that no value is present for ExcludeFolders, not even an explicit nil
### GetIncludeFolders

`func (o *Office365UserMailboxObjectProtectionParams) GetIncludeFolders() []string`

GetIncludeFolders returns the IncludeFolders field if non-nil, zero value otherwise.

### GetIncludeFoldersOk

`func (o *Office365UserMailboxObjectProtectionParams) GetIncludeFoldersOk() (*[]string, bool)`

GetIncludeFoldersOk returns a tuple with the IncludeFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFolders

`func (o *Office365UserMailboxObjectProtectionParams) SetIncludeFolders(v []string)`

SetIncludeFolders sets IncludeFolders field to given value.

### HasIncludeFolders

`func (o *Office365UserMailboxObjectProtectionParams) HasIncludeFolders() bool`

HasIncludeFolders returns a boolean if a field has been set.

### SetIncludeFoldersNil

`func (o *Office365UserMailboxObjectProtectionParams) SetIncludeFoldersNil(b bool)`

 SetIncludeFoldersNil sets the value for IncludeFolders to be an explicit nil

### UnsetIncludeFolders
`func (o *Office365UserMailboxObjectProtectionParams) UnsetIncludeFolders()`

UnsetIncludeFolders ensures that no value is present for IncludeFolders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


