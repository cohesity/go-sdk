# FilesAndFoldersObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbsolutePath** | **NullableString** | Specifies the absolute path of the file or folder. | 
**IsDirectory** | Pointer to **NullableBool** | Specifies whether the file or folder object is a directory. | [optional] 

## Methods

### NewFilesAndFoldersObject

`func NewFilesAndFoldersObject(absolutePath NullableString, ) *FilesAndFoldersObject`

NewFilesAndFoldersObject instantiates a new FilesAndFoldersObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesAndFoldersObjectWithDefaults

`func NewFilesAndFoldersObjectWithDefaults() *FilesAndFoldersObject`

NewFilesAndFoldersObjectWithDefaults instantiates a new FilesAndFoldersObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsolutePath

`func (o *FilesAndFoldersObject) GetAbsolutePath() string`

GetAbsolutePath returns the AbsolutePath field if non-nil, zero value otherwise.

### GetAbsolutePathOk

`func (o *FilesAndFoldersObject) GetAbsolutePathOk() (*string, bool)`

GetAbsolutePathOk returns a tuple with the AbsolutePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsolutePath

`func (o *FilesAndFoldersObject) SetAbsolutePath(v string)`

SetAbsolutePath sets AbsolutePath field to given value.


### SetAbsolutePathNil

`func (o *FilesAndFoldersObject) SetAbsolutePathNil(b bool)`

 SetAbsolutePathNil sets the value for AbsolutePath to be an explicit nil

### UnsetAbsolutePath
`func (o *FilesAndFoldersObject) UnsetAbsolutePath()`

UnsetAbsolutePath ensures that no value is present for AbsolutePath, not even an explicit nil
### GetIsDirectory

`func (o *FilesAndFoldersObject) GetIsDirectory() bool`

GetIsDirectory returns the IsDirectory field if non-nil, zero value otherwise.

### GetIsDirectoryOk

`func (o *FilesAndFoldersObject) GetIsDirectoryOk() (*bool, bool)`

GetIsDirectoryOk returns a tuple with the IsDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirectory

`func (o *FilesAndFoldersObject) SetIsDirectory(v bool)`

SetIsDirectory sets IsDirectory field to given value.

### HasIsDirectory

`func (o *FilesAndFoldersObject) HasIsDirectory() bool`

HasIsDirectory returns a boolean if a field has been set.

### SetIsDirectoryNil

`func (o *FilesAndFoldersObject) SetIsDirectoryNil(b bool)`

 SetIsDirectoryNil sets the value for IsDirectory to be an explicit nil

### UnsetIsDirectory
`func (o *FilesAndFoldersObject) UnsetIsDirectory()`

UnsetIsDirectory ensures that no value is present for IsDirectory, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


