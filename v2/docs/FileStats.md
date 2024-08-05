# FileStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilesCount** | Pointer to **NullableInt64** | Specifies the number of files. | [optional] 
**FilesSizeBytes** | Pointer to **NullableInt64** | Specifies the size of all the files in bytes. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the file type. | [optional] 

## Methods

### NewFileStats

`func NewFileStats() *FileStats`

NewFileStats instantiates a new FileStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileStatsWithDefaults

`func NewFileStatsWithDefaults() *FileStats`

NewFileStatsWithDefaults instantiates a new FileStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilesCount

`func (o *FileStats) GetFilesCount() int64`

GetFilesCount returns the FilesCount field if non-nil, zero value otherwise.

### GetFilesCountOk

`func (o *FileStats) GetFilesCountOk() (*int64, bool)`

GetFilesCountOk returns a tuple with the FilesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesCount

`func (o *FileStats) SetFilesCount(v int64)`

SetFilesCount sets FilesCount field to given value.

### HasFilesCount

`func (o *FileStats) HasFilesCount() bool`

HasFilesCount returns a boolean if a field has been set.

### SetFilesCountNil

`func (o *FileStats) SetFilesCountNil(b bool)`

 SetFilesCountNil sets the value for FilesCount to be an explicit nil

### UnsetFilesCount
`func (o *FileStats) UnsetFilesCount()`

UnsetFilesCount ensures that no value is present for FilesCount, not even an explicit nil
### GetFilesSizeBytes

`func (o *FileStats) GetFilesSizeBytes() int64`

GetFilesSizeBytes returns the FilesSizeBytes field if non-nil, zero value otherwise.

### GetFilesSizeBytesOk

`func (o *FileStats) GetFilesSizeBytesOk() (*int64, bool)`

GetFilesSizeBytesOk returns a tuple with the FilesSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesSizeBytes

`func (o *FileStats) SetFilesSizeBytes(v int64)`

SetFilesSizeBytes sets FilesSizeBytes field to given value.

### HasFilesSizeBytes

`func (o *FileStats) HasFilesSizeBytes() bool`

HasFilesSizeBytes returns a boolean if a field has been set.

### SetFilesSizeBytesNil

`func (o *FileStats) SetFilesSizeBytesNil(b bool)`

 SetFilesSizeBytesNil sets the value for FilesSizeBytes to be an explicit nil

### UnsetFilesSizeBytes
`func (o *FileStats) UnsetFilesSizeBytes()`

UnsetFilesSizeBytes ensures that no value is present for FilesSizeBytes, not even an explicit nil
### GetType

`func (o *FileStats) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FileStats) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FileStats) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FileStats) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *FileStats) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *FileStats) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


