# CommonDownloadFileAndFolderParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilesAndFolders** | Pointer to [**[]CommonRecoverFileAndFolderInfo**](CommonRecoverFileAndFolderInfo.md) | Specifies the info about the files and folders to be recovered. | [optional] 
**DownloadFilePath** | Pointer to **NullableString** | Specifies the path location to download the files and folders. | [optional] 

## Methods

### NewCommonDownloadFileAndFolderParams

`func NewCommonDownloadFileAndFolderParams() *CommonDownloadFileAndFolderParams`

NewCommonDownloadFileAndFolderParams instantiates a new CommonDownloadFileAndFolderParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonDownloadFileAndFolderParamsWithDefaults

`func NewCommonDownloadFileAndFolderParamsWithDefaults() *CommonDownloadFileAndFolderParams`

NewCommonDownloadFileAndFolderParamsWithDefaults instantiates a new CommonDownloadFileAndFolderParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilesAndFolders

`func (o *CommonDownloadFileAndFolderParams) GetFilesAndFolders() []CommonRecoverFileAndFolderInfo`

GetFilesAndFolders returns the FilesAndFolders field if non-nil, zero value otherwise.

### GetFilesAndFoldersOk

`func (o *CommonDownloadFileAndFolderParams) GetFilesAndFoldersOk() (*[]CommonRecoverFileAndFolderInfo, bool)`

GetFilesAndFoldersOk returns a tuple with the FilesAndFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesAndFolders

`func (o *CommonDownloadFileAndFolderParams) SetFilesAndFolders(v []CommonRecoverFileAndFolderInfo)`

SetFilesAndFolders sets FilesAndFolders field to given value.

### HasFilesAndFolders

`func (o *CommonDownloadFileAndFolderParams) HasFilesAndFolders() bool`

HasFilesAndFolders returns a boolean if a field has been set.

### SetFilesAndFoldersNil

`func (o *CommonDownloadFileAndFolderParams) SetFilesAndFoldersNil(b bool)`

 SetFilesAndFoldersNil sets the value for FilesAndFolders to be an explicit nil

### UnsetFilesAndFolders
`func (o *CommonDownloadFileAndFolderParams) UnsetFilesAndFolders()`

UnsetFilesAndFolders ensures that no value is present for FilesAndFolders, not even an explicit nil
### GetDownloadFilePath

`func (o *CommonDownloadFileAndFolderParams) GetDownloadFilePath() string`

GetDownloadFilePath returns the DownloadFilePath field if non-nil, zero value otherwise.

### GetDownloadFilePathOk

`func (o *CommonDownloadFileAndFolderParams) GetDownloadFilePathOk() (*string, bool)`

GetDownloadFilePathOk returns a tuple with the DownloadFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadFilePath

`func (o *CommonDownloadFileAndFolderParams) SetDownloadFilePath(v string)`

SetDownloadFilePath sets DownloadFilePath field to given value.

### HasDownloadFilePath

`func (o *CommonDownloadFileAndFolderParams) HasDownloadFilePath() bool`

HasDownloadFilePath returns a boolean if a field has been set.

### SetDownloadFilePathNil

`func (o *CommonDownloadFileAndFolderParams) SetDownloadFilePathNil(b bool)`

 SetDownloadFilePathNil sets the value for DownloadFilePath to be an explicit nil

### UnsetDownloadFilePath
`func (o *CommonDownloadFileAndFolderParams) UnsetDownloadFilePath()`

UnsetDownloadFilePath ensures that no value is present for DownloadFilePath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


