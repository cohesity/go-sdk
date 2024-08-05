# RecoverAwsParamsDownloadFileAndFolderParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownloadFilePath** | Pointer to **NullableString** | Specifies the path location to download the files and folders. | [optional] 
**ExpiryTimeUsecs** | Pointer to **NullableInt64** | Specifies the time upto which the download link is available. | [optional] 
**FilesAndFolders** | Pointer to [**[]CommonRecoverFileAndFolderInfo**](CommonRecoverFileAndFolderInfo.md) | Specifies the info about the files and folders to be recovered. | [optional] 

## Methods

### NewRecoverAwsParamsDownloadFileAndFolderParams

`func NewRecoverAwsParamsDownloadFileAndFolderParams() *RecoverAwsParamsDownloadFileAndFolderParams`

NewRecoverAwsParamsDownloadFileAndFolderParams instantiates a new RecoverAwsParamsDownloadFileAndFolderParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverAwsParamsDownloadFileAndFolderParamsWithDefaults

`func NewRecoverAwsParamsDownloadFileAndFolderParamsWithDefaults() *RecoverAwsParamsDownloadFileAndFolderParams`

NewRecoverAwsParamsDownloadFileAndFolderParamsWithDefaults instantiates a new RecoverAwsParamsDownloadFileAndFolderParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadFilePath

`func (o *RecoverAwsParamsDownloadFileAndFolderParams) GetDownloadFilePath() string`

GetDownloadFilePath returns the DownloadFilePath field if non-nil, zero value otherwise.

### GetDownloadFilePathOk

`func (o *RecoverAwsParamsDownloadFileAndFolderParams) GetDownloadFilePathOk() (*string, bool)`

GetDownloadFilePathOk returns a tuple with the DownloadFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadFilePath

`func (o *RecoverAwsParamsDownloadFileAndFolderParams) SetDownloadFilePath(v string)`

SetDownloadFilePath sets DownloadFilePath field to given value.

### HasDownloadFilePath

`func (o *RecoverAwsParamsDownloadFileAndFolderParams) HasDownloadFilePath() bool`

HasDownloadFilePath returns a boolean if a field has been set.

### SetDownloadFilePathNil

`func (o *RecoverAwsParamsDownloadFileAndFolderParams) SetDownloadFilePathNil(b bool)`

 SetDownloadFilePathNil sets the value for DownloadFilePath to be an explicit nil

### UnsetDownloadFilePath
`func (o *RecoverAwsParamsDownloadFileAndFolderParams) UnsetDownloadFilePath()`

UnsetDownloadFilePath ensures that no value is present for DownloadFilePath, not even an explicit nil
### GetExpiryTimeUsecs

`func (o *RecoverAwsParamsDownloadFileAndFolderParams) GetExpiryTimeUsecs() int64`

GetExpiryTimeUsecs returns the ExpiryTimeUsecs field if non-nil, zero value otherwise.

### GetExpiryTimeUsecsOk

`func (o *RecoverAwsParamsDownloadFileAndFolderParams) GetExpiryTimeUsecsOk() (*int64, bool)`

GetExpiryTimeUsecsOk returns a tuple with the ExpiryTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimeUsecs

`func (o *RecoverAwsParamsDownloadFileAndFolderParams) SetExpiryTimeUsecs(v int64)`

SetExpiryTimeUsecs sets ExpiryTimeUsecs field to given value.

### HasExpiryTimeUsecs

`func (o *RecoverAwsParamsDownloadFileAndFolderParams) HasExpiryTimeUsecs() bool`

HasExpiryTimeUsecs returns a boolean if a field has been set.

### SetExpiryTimeUsecsNil

`func (o *RecoverAwsParamsDownloadFileAndFolderParams) SetExpiryTimeUsecsNil(b bool)`

 SetExpiryTimeUsecsNil sets the value for ExpiryTimeUsecs to be an explicit nil

### UnsetExpiryTimeUsecs
`func (o *RecoverAwsParamsDownloadFileAndFolderParams) UnsetExpiryTimeUsecs()`

UnsetExpiryTimeUsecs ensures that no value is present for ExpiryTimeUsecs, not even an explicit nil
### GetFilesAndFolders

`func (o *RecoverAwsParamsDownloadFileAndFolderParams) GetFilesAndFolders() []CommonRecoverFileAndFolderInfo`

GetFilesAndFolders returns the FilesAndFolders field if non-nil, zero value otherwise.

### GetFilesAndFoldersOk

`func (o *RecoverAwsParamsDownloadFileAndFolderParams) GetFilesAndFoldersOk() (*[]CommonRecoverFileAndFolderInfo, bool)`

GetFilesAndFoldersOk returns a tuple with the FilesAndFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesAndFolders

`func (o *RecoverAwsParamsDownloadFileAndFolderParams) SetFilesAndFolders(v []CommonRecoverFileAndFolderInfo)`

SetFilesAndFolders sets FilesAndFolders field to given value.

### HasFilesAndFolders

`func (o *RecoverAwsParamsDownloadFileAndFolderParams) HasFilesAndFolders() bool`

HasFilesAndFolders returns a boolean if a field has been set.

### SetFilesAndFoldersNil

`func (o *RecoverAwsParamsDownloadFileAndFolderParams) SetFilesAndFoldersNil(b bool)`

 SetFilesAndFoldersNil sets the value for FilesAndFolders to be an explicit nil

### UnsetFilesAndFolders
`func (o *RecoverAwsParamsDownloadFileAndFolderParams) UnsetFilesAndFolders()`

UnsetFilesAndFolders ensures that no value is present for FilesAndFolders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


