# DataTieringShareStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessTimeTag** | Pointer to **NullableString** | Specifies the access time bucket. | [optional] 
**FileCount** | Pointer to **NullableInt32** | Specifies the file count. | [optional] 
**FileSizeTag** | Pointer to **NullableString** | Specifies the file size bucket. | [optional] 
**FileTypeTag** | Pointer to **NullableString** | Specifies the file type bucket. | [optional] 
**Id** | Pointer to **NullableInt32** | Specifies the unique identifer for stat. | [optional] 
**ModTimeTag** | Pointer to **NullableString** | Specifies the modification time bucket. | [optional] 
**TotalSize** | Pointer to **NullableInt32** | Specifies the total count. | [optional] 

## Methods

### NewDataTieringShareStats

`func NewDataTieringShareStats() *DataTieringShareStats`

NewDataTieringShareStats instantiates a new DataTieringShareStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTieringShareStatsWithDefaults

`func NewDataTieringShareStatsWithDefaults() *DataTieringShareStats`

NewDataTieringShareStatsWithDefaults instantiates a new DataTieringShareStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessTimeTag

`func (o *DataTieringShareStats) GetAccessTimeTag() string`

GetAccessTimeTag returns the AccessTimeTag field if non-nil, zero value otherwise.

### GetAccessTimeTagOk

`func (o *DataTieringShareStats) GetAccessTimeTagOk() (*string, bool)`

GetAccessTimeTagOk returns a tuple with the AccessTimeTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTimeTag

`func (o *DataTieringShareStats) SetAccessTimeTag(v string)`

SetAccessTimeTag sets AccessTimeTag field to given value.

### HasAccessTimeTag

`func (o *DataTieringShareStats) HasAccessTimeTag() bool`

HasAccessTimeTag returns a boolean if a field has been set.

### SetAccessTimeTagNil

`func (o *DataTieringShareStats) SetAccessTimeTagNil(b bool)`

 SetAccessTimeTagNil sets the value for AccessTimeTag to be an explicit nil

### UnsetAccessTimeTag
`func (o *DataTieringShareStats) UnsetAccessTimeTag()`

UnsetAccessTimeTag ensures that no value is present for AccessTimeTag, not even an explicit nil
### GetFileCount

`func (o *DataTieringShareStats) GetFileCount() int32`

GetFileCount returns the FileCount field if non-nil, zero value otherwise.

### GetFileCountOk

`func (o *DataTieringShareStats) GetFileCountOk() (*int32, bool)`

GetFileCountOk returns a tuple with the FileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileCount

`func (o *DataTieringShareStats) SetFileCount(v int32)`

SetFileCount sets FileCount field to given value.

### HasFileCount

`func (o *DataTieringShareStats) HasFileCount() bool`

HasFileCount returns a boolean if a field has been set.

### SetFileCountNil

`func (o *DataTieringShareStats) SetFileCountNil(b bool)`

 SetFileCountNil sets the value for FileCount to be an explicit nil

### UnsetFileCount
`func (o *DataTieringShareStats) UnsetFileCount()`

UnsetFileCount ensures that no value is present for FileCount, not even an explicit nil
### GetFileSizeTag

`func (o *DataTieringShareStats) GetFileSizeTag() string`

GetFileSizeTag returns the FileSizeTag field if non-nil, zero value otherwise.

### GetFileSizeTagOk

`func (o *DataTieringShareStats) GetFileSizeTagOk() (*string, bool)`

GetFileSizeTagOk returns a tuple with the FileSizeTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSizeTag

`func (o *DataTieringShareStats) SetFileSizeTag(v string)`

SetFileSizeTag sets FileSizeTag field to given value.

### HasFileSizeTag

`func (o *DataTieringShareStats) HasFileSizeTag() bool`

HasFileSizeTag returns a boolean if a field has been set.

### SetFileSizeTagNil

`func (o *DataTieringShareStats) SetFileSizeTagNil(b bool)`

 SetFileSizeTagNil sets the value for FileSizeTag to be an explicit nil

### UnsetFileSizeTag
`func (o *DataTieringShareStats) UnsetFileSizeTag()`

UnsetFileSizeTag ensures that no value is present for FileSizeTag, not even an explicit nil
### GetFileTypeTag

`func (o *DataTieringShareStats) GetFileTypeTag() string`

GetFileTypeTag returns the FileTypeTag field if non-nil, zero value otherwise.

### GetFileTypeTagOk

`func (o *DataTieringShareStats) GetFileTypeTagOk() (*string, bool)`

GetFileTypeTagOk returns a tuple with the FileTypeTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileTypeTag

`func (o *DataTieringShareStats) SetFileTypeTag(v string)`

SetFileTypeTag sets FileTypeTag field to given value.

### HasFileTypeTag

`func (o *DataTieringShareStats) HasFileTypeTag() bool`

HasFileTypeTag returns a boolean if a field has been set.

### SetFileTypeTagNil

`func (o *DataTieringShareStats) SetFileTypeTagNil(b bool)`

 SetFileTypeTagNil sets the value for FileTypeTag to be an explicit nil

### UnsetFileTypeTag
`func (o *DataTieringShareStats) UnsetFileTypeTag()`

UnsetFileTypeTag ensures that no value is present for FileTypeTag, not even an explicit nil
### GetId

`func (o *DataTieringShareStats) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataTieringShareStats) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataTieringShareStats) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DataTieringShareStats) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DataTieringShareStats) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DataTieringShareStats) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetModTimeTag

`func (o *DataTieringShareStats) GetModTimeTag() string`

GetModTimeTag returns the ModTimeTag field if non-nil, zero value otherwise.

### GetModTimeTagOk

`func (o *DataTieringShareStats) GetModTimeTagOk() (*string, bool)`

GetModTimeTagOk returns a tuple with the ModTimeTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTimeTag

`func (o *DataTieringShareStats) SetModTimeTag(v string)`

SetModTimeTag sets ModTimeTag field to given value.

### HasModTimeTag

`func (o *DataTieringShareStats) HasModTimeTag() bool`

HasModTimeTag returns a boolean if a field has been set.

### SetModTimeTagNil

`func (o *DataTieringShareStats) SetModTimeTagNil(b bool)`

 SetModTimeTagNil sets the value for ModTimeTag to be an explicit nil

### UnsetModTimeTag
`func (o *DataTieringShareStats) UnsetModTimeTag()`

UnsetModTimeTag ensures that no value is present for ModTimeTag, not even an explicit nil
### GetTotalSize

`func (o *DataTieringShareStats) GetTotalSize() int32`

GetTotalSize returns the TotalSize field if non-nil, zero value otherwise.

### GetTotalSizeOk

`func (o *DataTieringShareStats) GetTotalSizeOk() (*int32, bool)`

GetTotalSizeOk returns a tuple with the TotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSize

`func (o *DataTieringShareStats) SetTotalSize(v int32)`

SetTotalSize sets TotalSize field to given value.

### HasTotalSize

`func (o *DataTieringShareStats) HasTotalSize() bool`

HasTotalSize returns a boolean if a field has been set.

### SetTotalSizeNil

`func (o *DataTieringShareStats) SetTotalSizeNil(b bool)`

 SetTotalSizeNil sets the value for TotalSize to be an explicit nil

### UnsetTotalSize
`func (o *DataTieringShareStats) UnsetTotalSize()`

UnsetTotalSize ensures that no value is present for TotalSize, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


