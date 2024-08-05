# Office365UserOneDriveObjectProtectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 
**Objects** | [**[]Office365ObjectProtectionObjectParams**](Office365ObjectProtectionObjectParams.md) | Specifies the objects to be included in the Object Protection. | 
**SourceId** | Pointer to **NullableInt64** | Specifies the id of the parent of the objects. | [optional] [readonly] 
**SourceName** | Pointer to **NullableString** | Specifies the name of the parent of the objects. | [optional] [readonly] 
**ExcludeFolders** | Pointer to **[]string** | Specifies filters to match OneDrive folders which should be excluded when backing up office365 onedrive source. Two kinds of filters are supported. a) prefix which always starts with &#39;/&#39;. b) posix which always starts with empty quotes(&#39;&#39;). Regular expressions are not supported. If not specified, all the OneDrive will be protected. | [optional] 
**PreservationHoldLibraryParams** | Pointer to [**Office365PreservationHoldLibraryParams**](Office365PreservationHoldLibraryParams.md) |  | [optional] 

## Methods

### NewOffice365UserOneDriveObjectProtectionParams

`func NewOffice365UserOneDriveObjectProtectionParams(objects []Office365ObjectProtectionObjectParams, ) *Office365UserOneDriveObjectProtectionParams`

NewOffice365UserOneDriveObjectProtectionParams instantiates a new Office365UserOneDriveObjectProtectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOffice365UserOneDriveObjectProtectionParamsWithDefaults

`func NewOffice365UserOneDriveObjectProtectionParamsWithDefaults() *Office365UserOneDriveObjectProtectionParams`

NewOffice365UserOneDriveObjectProtectionParamsWithDefaults instantiates a new Office365UserOneDriveObjectProtectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexingPolicy

`func (o *Office365UserOneDriveObjectProtectionParams) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *Office365UserOneDriveObjectProtectionParams) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *Office365UserOneDriveObjectProtectionParams) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *Office365UserOneDriveObjectProtectionParams) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetObjects

`func (o *Office365UserOneDriveObjectProtectionParams) GetObjects() []Office365ObjectProtectionObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *Office365UserOneDriveObjectProtectionParams) GetObjectsOk() (*[]Office365ObjectProtectionObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *Office365UserOneDriveObjectProtectionParams) SetObjects(v []Office365ObjectProtectionObjectParams)`

SetObjects sets Objects field to given value.


### GetSourceId

`func (o *Office365UserOneDriveObjectProtectionParams) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *Office365UserOneDriveObjectProtectionParams) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *Office365UserOneDriveObjectProtectionParams) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *Office365UserOneDriveObjectProtectionParams) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *Office365UserOneDriveObjectProtectionParams) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *Office365UserOneDriveObjectProtectionParams) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *Office365UserOneDriveObjectProtectionParams) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *Office365UserOneDriveObjectProtectionParams) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *Office365UserOneDriveObjectProtectionParams) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *Office365UserOneDriveObjectProtectionParams) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *Office365UserOneDriveObjectProtectionParams) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *Office365UserOneDriveObjectProtectionParams) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetExcludeFolders

`func (o *Office365UserOneDriveObjectProtectionParams) GetExcludeFolders() []string`

GetExcludeFolders returns the ExcludeFolders field if non-nil, zero value otherwise.

### GetExcludeFoldersOk

`func (o *Office365UserOneDriveObjectProtectionParams) GetExcludeFoldersOk() (*[]string, bool)`

GetExcludeFoldersOk returns a tuple with the ExcludeFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFolders

`func (o *Office365UserOneDriveObjectProtectionParams) SetExcludeFolders(v []string)`

SetExcludeFolders sets ExcludeFolders field to given value.

### HasExcludeFolders

`func (o *Office365UserOneDriveObjectProtectionParams) HasExcludeFolders() bool`

HasExcludeFolders returns a boolean if a field has been set.

### SetExcludeFoldersNil

`func (o *Office365UserOneDriveObjectProtectionParams) SetExcludeFoldersNil(b bool)`

 SetExcludeFoldersNil sets the value for ExcludeFolders to be an explicit nil

### UnsetExcludeFolders
`func (o *Office365UserOneDriveObjectProtectionParams) UnsetExcludeFolders()`

UnsetExcludeFolders ensures that no value is present for ExcludeFolders, not even an explicit nil
### GetPreservationHoldLibraryParams

`func (o *Office365UserOneDriveObjectProtectionParams) GetPreservationHoldLibraryParams() Office365PreservationHoldLibraryParams`

GetPreservationHoldLibraryParams returns the PreservationHoldLibraryParams field if non-nil, zero value otherwise.

### GetPreservationHoldLibraryParamsOk

`func (o *Office365UserOneDriveObjectProtectionParams) GetPreservationHoldLibraryParamsOk() (*Office365PreservationHoldLibraryParams, bool)`

GetPreservationHoldLibraryParamsOk returns a tuple with the PreservationHoldLibraryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreservationHoldLibraryParams

`func (o *Office365UserOneDriveObjectProtectionParams) SetPreservationHoldLibraryParams(v Office365PreservationHoldLibraryParams)`

SetPreservationHoldLibraryParams sets PreservationHoldLibraryParams field to given value.

### HasPreservationHoldLibraryParams

`func (o *Office365UserOneDriveObjectProtectionParams) HasPreservationHoldLibraryParams() bool`

HasPreservationHoldLibraryParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


