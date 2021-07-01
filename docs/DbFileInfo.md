# DbFileInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileType** | Pointer to **NullableString** | Specifies the format type of the file that SQL database stores the data. Specifies the format type of the file that SQL database stores the data. &#39;kRows&#39; refers to a data file &#39;kLog&#39; refers to a log file &#39;kFileStream&#39; refers to a directory containing FILESTREAM data &#39;kNotSupportedType&#39; is for information purposes only. Not supported. &#39;kFullText&#39; refers to a full-text catalog. | [optional] 
**FullPath** | Pointer to **NullableString** | Specifies the full path of the database file on the SQL host machine. | [optional] 
**SizeBytes** | Pointer to **NullableInt64** | Specifies the last known size of the database file. | [optional] 

## Methods

### NewDbFileInfo

`func NewDbFileInfo() *DbFileInfo`

NewDbFileInfo instantiates a new DbFileInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbFileInfoWithDefaults

`func NewDbFileInfoWithDefaults() *DbFileInfo`

NewDbFileInfoWithDefaults instantiates a new DbFileInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileType

`func (o *DbFileInfo) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *DbFileInfo) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *DbFileInfo) SetFileType(v string)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *DbFileInfo) HasFileType() bool`

HasFileType returns a boolean if a field has been set.

### SetFileTypeNil

`func (o *DbFileInfo) SetFileTypeNil(b bool)`

 SetFileTypeNil sets the value for FileType to be an explicit nil

### UnsetFileType
`func (o *DbFileInfo) UnsetFileType()`

UnsetFileType ensures that no value is present for FileType, not even an explicit nil
### GetFullPath

`func (o *DbFileInfo) GetFullPath() string`

GetFullPath returns the FullPath field if non-nil, zero value otherwise.

### GetFullPathOk

`func (o *DbFileInfo) GetFullPathOk() (*string, bool)`

GetFullPathOk returns a tuple with the FullPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullPath

`func (o *DbFileInfo) SetFullPath(v string)`

SetFullPath sets FullPath field to given value.

### HasFullPath

`func (o *DbFileInfo) HasFullPath() bool`

HasFullPath returns a boolean if a field has been set.

### SetFullPathNil

`func (o *DbFileInfo) SetFullPathNil(b bool)`

 SetFullPathNil sets the value for FullPath to be an explicit nil

### UnsetFullPath
`func (o *DbFileInfo) UnsetFullPath()`

UnsetFullPath ensures that no value is present for FullPath, not even an explicit nil
### GetSizeBytes

`func (o *DbFileInfo) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *DbFileInfo) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *DbFileInfo) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *DbFileInfo) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### SetSizeBytesNil

`func (o *DbFileInfo) SetSizeBytesNil(b bool)`

 SetSizeBytesNil sets the value for SizeBytes to be an explicit nil

### UnsetSizeBytes
`func (o *DbFileInfo) UnsetSizeBytes()`

UnsetSizeBytes ensures that no value is present for SizeBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


