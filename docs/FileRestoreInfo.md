# FileRestoreInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**RequestError**](RequestError.md) |  | [optional] 
**Filename** | Pointer to **NullableString** | Specifies the path of the file/directory. | [optional] 
**FilesystemVolume** | Pointer to [**FilesystemVolume**](FilesystemVolume.md) |  | [optional] 
**IsFolder** | Pointer to **NullableBool** | Specifies whether the file path is a folder. | [optional] 

## Methods

### NewFileRestoreInfo

`func NewFileRestoreInfo() *FileRestoreInfo`

NewFileRestoreInfo instantiates a new FileRestoreInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileRestoreInfoWithDefaults

`func NewFileRestoreInfoWithDefaults() *FileRestoreInfo`

NewFileRestoreInfoWithDefaults instantiates a new FileRestoreInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *FileRestoreInfo) GetError() RequestError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *FileRestoreInfo) GetErrorOk() (*RequestError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *FileRestoreInfo) SetError(v RequestError)`

SetError sets Error field to given value.

### HasError

`func (o *FileRestoreInfo) HasError() bool`

HasError returns a boolean if a field has been set.

### GetFilename

`func (o *FileRestoreInfo) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *FileRestoreInfo) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *FileRestoreInfo) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *FileRestoreInfo) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### SetFilenameNil

`func (o *FileRestoreInfo) SetFilenameNil(b bool)`

 SetFilenameNil sets the value for Filename to be an explicit nil

### UnsetFilename
`func (o *FileRestoreInfo) UnsetFilename()`

UnsetFilename ensures that no value is present for Filename, not even an explicit nil
### GetFilesystemVolume

`func (o *FileRestoreInfo) GetFilesystemVolume() FilesystemVolume`

GetFilesystemVolume returns the FilesystemVolume field if non-nil, zero value otherwise.

### GetFilesystemVolumeOk

`func (o *FileRestoreInfo) GetFilesystemVolumeOk() (*FilesystemVolume, bool)`

GetFilesystemVolumeOk returns a tuple with the FilesystemVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystemVolume

`func (o *FileRestoreInfo) SetFilesystemVolume(v FilesystemVolume)`

SetFilesystemVolume sets FilesystemVolume field to given value.

### HasFilesystemVolume

`func (o *FileRestoreInfo) HasFilesystemVolume() bool`

HasFilesystemVolume returns a boolean if a field has been set.

### GetIsFolder

`func (o *FileRestoreInfo) GetIsFolder() bool`

GetIsFolder returns the IsFolder field if non-nil, zero value otherwise.

### GetIsFolderOk

`func (o *FileRestoreInfo) GetIsFolderOk() (*bool, bool)`

GetIsFolderOk returns a tuple with the IsFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFolder

`func (o *FileRestoreInfo) SetIsFolder(v bool)`

SetIsFolder sets IsFolder field to given value.

### HasIsFolder

`func (o *FileRestoreInfo) HasIsFolder() bool`

HasIsFolder returns a boolean if a field has been set.

### SetIsFolderNil

`func (o *FileRestoreInfo) SetIsFolderNil(b bool)`

 SetIsFolderNil sets the value for IsFolder to be an explicit nil

### UnsetIsFolder
`func (o *FileRestoreInfo) UnsetIsFolder()`

UnsetIsFolder ensures that no value is present for IsFolder, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


