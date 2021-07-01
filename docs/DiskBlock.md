# DiskBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LengthBytes** | Pointer to **NullableInt64** | Specifies the length of the block in bytes. | [optional] 
**OffsetBytes** | Pointer to **NullableInt64** | Specifies the offset of the block (in bytes) from the beginning of the containing object such as a physical disk or a virtual disk file. | [optional] 

## Methods

### NewDiskBlock

`func NewDiskBlock() *DiskBlock`

NewDiskBlock instantiates a new DiskBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskBlockWithDefaults

`func NewDiskBlockWithDefaults() *DiskBlock`

NewDiskBlockWithDefaults instantiates a new DiskBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLengthBytes

`func (o *DiskBlock) GetLengthBytes() int64`

GetLengthBytes returns the LengthBytes field if non-nil, zero value otherwise.

### GetLengthBytesOk

`func (o *DiskBlock) GetLengthBytesOk() (*int64, bool)`

GetLengthBytesOk returns a tuple with the LengthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthBytes

`func (o *DiskBlock) SetLengthBytes(v int64)`

SetLengthBytes sets LengthBytes field to given value.

### HasLengthBytes

`func (o *DiskBlock) HasLengthBytes() bool`

HasLengthBytes returns a boolean if a field has been set.

### SetLengthBytesNil

`func (o *DiskBlock) SetLengthBytesNil(b bool)`

 SetLengthBytesNil sets the value for LengthBytes to be an explicit nil

### UnsetLengthBytes
`func (o *DiskBlock) UnsetLengthBytes()`

UnsetLengthBytes ensures that no value is present for LengthBytes, not even an explicit nil
### GetOffsetBytes

`func (o *DiskBlock) GetOffsetBytes() int64`

GetOffsetBytes returns the OffsetBytes field if non-nil, zero value otherwise.

### GetOffsetBytesOk

`func (o *DiskBlock) GetOffsetBytesOk() (*int64, bool)`

GetOffsetBytesOk returns a tuple with the OffsetBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetBytes

`func (o *DiskBlock) SetOffsetBytes(v int64)`

SetOffsetBytes sets OffsetBytes field to given value.

### HasOffsetBytes

`func (o *DiskBlock) HasOffsetBytes() bool`

HasOffsetBytes returns a boolean if a field has been set.

### SetOffsetBytesNil

`func (o *DiskBlock) SetOffsetBytesNil(b bool)`

 SetOffsetBytesNil sets the value for OffsetBytes to be an explicit nil

### UnsetOffsetBytes
`func (o *DiskBlock) UnsetOffsetBytes()`

UnsetOffsetBytes ensures that no value is present for OffsetBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


