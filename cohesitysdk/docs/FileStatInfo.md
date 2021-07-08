# FileStatInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupSourceInodeId** | Pointer to **NullableInt64** | Source inode id metadata for certain adapters e.g. Netapp. | [optional] 
**MtimeUsecs** | Pointer to **NullableInt64** | If this is a file, the mtime as returned by stat. | [optional] 
**Size** | Pointer to **NullableInt64** | If this is a file, the size of the file as returned by stat. | [optional] 
**Type** | Pointer to **NullableInt32** | The type of this entity. This field will not be populated for ReadDir results, since the DirEntry already contains the type information. | [optional] 

## Methods

### NewFileStatInfo

`func NewFileStatInfo() *FileStatInfo`

NewFileStatInfo instantiates a new FileStatInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileStatInfoWithDefaults

`func NewFileStatInfoWithDefaults() *FileStatInfo`

NewFileStatInfoWithDefaults instantiates a new FileStatInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupSourceInodeId

`func (o *FileStatInfo) GetBackupSourceInodeId() int64`

GetBackupSourceInodeId returns the BackupSourceInodeId field if non-nil, zero value otherwise.

### GetBackupSourceInodeIdOk

`func (o *FileStatInfo) GetBackupSourceInodeIdOk() (*int64, bool)`

GetBackupSourceInodeIdOk returns a tuple with the BackupSourceInodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSourceInodeId

`func (o *FileStatInfo) SetBackupSourceInodeId(v int64)`

SetBackupSourceInodeId sets BackupSourceInodeId field to given value.

### HasBackupSourceInodeId

`func (o *FileStatInfo) HasBackupSourceInodeId() bool`

HasBackupSourceInodeId returns a boolean if a field has been set.

### SetBackupSourceInodeIdNil

`func (o *FileStatInfo) SetBackupSourceInodeIdNil(b bool)`

 SetBackupSourceInodeIdNil sets the value for BackupSourceInodeId to be an explicit nil

### UnsetBackupSourceInodeId
`func (o *FileStatInfo) UnsetBackupSourceInodeId()`

UnsetBackupSourceInodeId ensures that no value is present for BackupSourceInodeId, not even an explicit nil
### GetMtimeUsecs

`func (o *FileStatInfo) GetMtimeUsecs() int64`

GetMtimeUsecs returns the MtimeUsecs field if non-nil, zero value otherwise.

### GetMtimeUsecsOk

`func (o *FileStatInfo) GetMtimeUsecsOk() (*int64, bool)`

GetMtimeUsecsOk returns a tuple with the MtimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtimeUsecs

`func (o *FileStatInfo) SetMtimeUsecs(v int64)`

SetMtimeUsecs sets MtimeUsecs field to given value.

### HasMtimeUsecs

`func (o *FileStatInfo) HasMtimeUsecs() bool`

HasMtimeUsecs returns a boolean if a field has been set.

### SetMtimeUsecsNil

`func (o *FileStatInfo) SetMtimeUsecsNil(b bool)`

 SetMtimeUsecsNil sets the value for MtimeUsecs to be an explicit nil

### UnsetMtimeUsecs
`func (o *FileStatInfo) UnsetMtimeUsecs()`

UnsetMtimeUsecs ensures that no value is present for MtimeUsecs, not even an explicit nil
### GetSize

`func (o *FileStatInfo) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FileStatInfo) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FileStatInfo) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *FileStatInfo) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *FileStatInfo) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *FileStatInfo) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetType

`func (o *FileStatInfo) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FileStatInfo) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FileStatInfo) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *FileStatInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *FileStatInfo) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *FileStatInfo) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


