# FilePartitionBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskFileName** | Pointer to **NullableString** | Specifies the disk file name where the logical partition is. | [optional] 
**LengthBytes** | Pointer to **NullableInt64** | Specifies the length of the block in bytes. | [optional] 
**Number** | Pointer to **NullableInt64** | Specifies a unique number of the partition within the linear disk file. | [optional] 
**OffsetBytes** | Pointer to **NullableInt64** | Specifies the offset of the block (in bytes) from the beginning of the containing object such as a physical disk or a virtual disk file. | [optional] 

## Methods

### NewFilePartitionBlock

`func NewFilePartitionBlock() *FilePartitionBlock`

NewFilePartitionBlock instantiates a new FilePartitionBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilePartitionBlockWithDefaults

`func NewFilePartitionBlockWithDefaults() *FilePartitionBlock`

NewFilePartitionBlockWithDefaults instantiates a new FilePartitionBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskFileName

`func (o *FilePartitionBlock) GetDiskFileName() string`

GetDiskFileName returns the DiskFileName field if non-nil, zero value otherwise.

### GetDiskFileNameOk

`func (o *FilePartitionBlock) GetDiskFileNameOk() (*string, bool)`

GetDiskFileNameOk returns a tuple with the DiskFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskFileName

`func (o *FilePartitionBlock) SetDiskFileName(v string)`

SetDiskFileName sets DiskFileName field to given value.

### HasDiskFileName

`func (o *FilePartitionBlock) HasDiskFileName() bool`

HasDiskFileName returns a boolean if a field has been set.

### SetDiskFileNameNil

`func (o *FilePartitionBlock) SetDiskFileNameNil(b bool)`

 SetDiskFileNameNil sets the value for DiskFileName to be an explicit nil

### UnsetDiskFileName
`func (o *FilePartitionBlock) UnsetDiskFileName()`

UnsetDiskFileName ensures that no value is present for DiskFileName, not even an explicit nil
### GetLengthBytes

`func (o *FilePartitionBlock) GetLengthBytes() int64`

GetLengthBytes returns the LengthBytes field if non-nil, zero value otherwise.

### GetLengthBytesOk

`func (o *FilePartitionBlock) GetLengthBytesOk() (*int64, bool)`

GetLengthBytesOk returns a tuple with the LengthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthBytes

`func (o *FilePartitionBlock) SetLengthBytes(v int64)`

SetLengthBytes sets LengthBytes field to given value.

### HasLengthBytes

`func (o *FilePartitionBlock) HasLengthBytes() bool`

HasLengthBytes returns a boolean if a field has been set.

### SetLengthBytesNil

`func (o *FilePartitionBlock) SetLengthBytesNil(b bool)`

 SetLengthBytesNil sets the value for LengthBytes to be an explicit nil

### UnsetLengthBytes
`func (o *FilePartitionBlock) UnsetLengthBytes()`

UnsetLengthBytes ensures that no value is present for LengthBytes, not even an explicit nil
### GetNumber

`func (o *FilePartitionBlock) GetNumber() int64`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *FilePartitionBlock) GetNumberOk() (*int64, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *FilePartitionBlock) SetNumber(v int64)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *FilePartitionBlock) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *FilePartitionBlock) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *FilePartitionBlock) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetOffsetBytes

`func (o *FilePartitionBlock) GetOffsetBytes() int64`

GetOffsetBytes returns the OffsetBytes field if non-nil, zero value otherwise.

### GetOffsetBytesOk

`func (o *FilePartitionBlock) GetOffsetBytesOk() (*int64, bool)`

GetOffsetBytesOk returns a tuple with the OffsetBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetBytes

`func (o *FilePartitionBlock) SetOffsetBytes(v int64)`

SetOffsetBytes sets OffsetBytes field to given value.

### HasOffsetBytes

`func (o *FilePartitionBlock) HasOffsetBytes() bool`

HasOffsetBytes returns a boolean if a field has been set.

### SetOffsetBytesNil

`func (o *FilePartitionBlock) SetOffsetBytesNil(b bool)`

 SetOffsetBytesNil sets the value for OffsetBytes to be an explicit nil

### UnsetOffsetBytes
`func (o *FilePartitionBlock) UnsetOffsetBytes()`

UnsetOffsetBytes ensures that no value is present for OffsetBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


