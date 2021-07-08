# FilesAndFoldersInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbsolutePath** | Pointer to **NullableString** | AbsolutePath specifies the absolute path of the specified file or folder. | [optional] 
**IsDirectory** | Pointer to **NullableBool** | IsDirectory specifies if specified object is a directory or not. | [optional] 

## Methods

### NewFilesAndFoldersInfo

`func NewFilesAndFoldersInfo() *FilesAndFoldersInfo`

NewFilesAndFoldersInfo instantiates a new FilesAndFoldersInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesAndFoldersInfoWithDefaults

`func NewFilesAndFoldersInfoWithDefaults() *FilesAndFoldersInfo`

NewFilesAndFoldersInfoWithDefaults instantiates a new FilesAndFoldersInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsolutePath

`func (o *FilesAndFoldersInfo) GetAbsolutePath() string`

GetAbsolutePath returns the AbsolutePath field if non-nil, zero value otherwise.

### GetAbsolutePathOk

`func (o *FilesAndFoldersInfo) GetAbsolutePathOk() (*string, bool)`

GetAbsolutePathOk returns a tuple with the AbsolutePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsolutePath

`func (o *FilesAndFoldersInfo) SetAbsolutePath(v string)`

SetAbsolutePath sets AbsolutePath field to given value.

### HasAbsolutePath

`func (o *FilesAndFoldersInfo) HasAbsolutePath() bool`

HasAbsolutePath returns a boolean if a field has been set.

### SetAbsolutePathNil

`func (o *FilesAndFoldersInfo) SetAbsolutePathNil(b bool)`

 SetAbsolutePathNil sets the value for AbsolutePath to be an explicit nil

### UnsetAbsolutePath
`func (o *FilesAndFoldersInfo) UnsetAbsolutePath()`

UnsetAbsolutePath ensures that no value is present for AbsolutePath, not even an explicit nil
### GetIsDirectory

`func (o *FilesAndFoldersInfo) GetIsDirectory() bool`

GetIsDirectory returns the IsDirectory field if non-nil, zero value otherwise.

### GetIsDirectoryOk

`func (o *FilesAndFoldersInfo) GetIsDirectoryOk() (*bool, bool)`

GetIsDirectoryOk returns a tuple with the IsDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirectory

`func (o *FilesAndFoldersInfo) SetIsDirectory(v bool)`

SetIsDirectory sets IsDirectory field to given value.

### HasIsDirectory

`func (o *FilesAndFoldersInfo) HasIsDirectory() bool`

HasIsDirectory returns a boolean if a field has been set.

### SetIsDirectoryNil

`func (o *FilesAndFoldersInfo) SetIsDirectoryNil(b bool)`

 SetIsDirectoryNil sets the value for IsDirectory to be an explicit nil

### UnsetIsDirectory
`func (o *FilesAndFoldersInfo) UnsetIsDirectory()`

UnsetIsDirectory ensures that no value is present for IsDirectory, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


