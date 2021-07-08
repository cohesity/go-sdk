# DiskPartition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LengthBytes** | Pointer to **NullableInt64** | Specifies the length of the block in bytes. | [optional] 
**Number** | Pointer to **NullableInt64** | Specifies a unique number of the partition within the linear disk file. | [optional] 
**OffsetBytes** | Pointer to **NullableInt64** | Specifies the offset of the block (in bytes) from the beginning of the containing object such as a physical disk or a virtual disk file. | [optional] 
**TypeUuid** | Pointer to **NullableString** | Specifies the partition type uuid. If disk is unpartitioned, this field is not set. If disk is MBR partitioned, this field is set to a partition type. If disk is GPT partitioned, this field is set to a partition type GUID. | [optional] 
**Uuid** | Pointer to **NullableString** | Specifies the partition uuid. If disk is unpartitioned, this field is not set. If disk is MBR partitioned, this field is not set. If disk is GPT partitioned, this field is set to a partition GUID. | [optional] 

## Methods

### NewDiskPartition

`func NewDiskPartition() *DiskPartition`

NewDiskPartition instantiates a new DiskPartition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskPartitionWithDefaults

`func NewDiskPartitionWithDefaults() *DiskPartition`

NewDiskPartitionWithDefaults instantiates a new DiskPartition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLengthBytes

`func (o *DiskPartition) GetLengthBytes() int64`

GetLengthBytes returns the LengthBytes field if non-nil, zero value otherwise.

### GetLengthBytesOk

`func (o *DiskPartition) GetLengthBytesOk() (*int64, bool)`

GetLengthBytesOk returns a tuple with the LengthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthBytes

`func (o *DiskPartition) SetLengthBytes(v int64)`

SetLengthBytes sets LengthBytes field to given value.

### HasLengthBytes

`func (o *DiskPartition) HasLengthBytes() bool`

HasLengthBytes returns a boolean if a field has been set.

### SetLengthBytesNil

`func (o *DiskPartition) SetLengthBytesNil(b bool)`

 SetLengthBytesNil sets the value for LengthBytes to be an explicit nil

### UnsetLengthBytes
`func (o *DiskPartition) UnsetLengthBytes()`

UnsetLengthBytes ensures that no value is present for LengthBytes, not even an explicit nil
### GetNumber

`func (o *DiskPartition) GetNumber() int64`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *DiskPartition) GetNumberOk() (*int64, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *DiskPartition) SetNumber(v int64)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *DiskPartition) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *DiskPartition) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *DiskPartition) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetOffsetBytes

`func (o *DiskPartition) GetOffsetBytes() int64`

GetOffsetBytes returns the OffsetBytes field if non-nil, zero value otherwise.

### GetOffsetBytesOk

`func (o *DiskPartition) GetOffsetBytesOk() (*int64, bool)`

GetOffsetBytesOk returns a tuple with the OffsetBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetBytes

`func (o *DiskPartition) SetOffsetBytes(v int64)`

SetOffsetBytes sets OffsetBytes field to given value.

### HasOffsetBytes

`func (o *DiskPartition) HasOffsetBytes() bool`

HasOffsetBytes returns a boolean if a field has been set.

### SetOffsetBytesNil

`func (o *DiskPartition) SetOffsetBytesNil(b bool)`

 SetOffsetBytesNil sets the value for OffsetBytes to be an explicit nil

### UnsetOffsetBytes
`func (o *DiskPartition) UnsetOffsetBytes()`

UnsetOffsetBytes ensures that no value is present for OffsetBytes, not even an explicit nil
### GetTypeUuid

`func (o *DiskPartition) GetTypeUuid() string`

GetTypeUuid returns the TypeUuid field if non-nil, zero value otherwise.

### GetTypeUuidOk

`func (o *DiskPartition) GetTypeUuidOk() (*string, bool)`

GetTypeUuidOk returns a tuple with the TypeUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeUuid

`func (o *DiskPartition) SetTypeUuid(v string)`

SetTypeUuid sets TypeUuid field to given value.

### HasTypeUuid

`func (o *DiskPartition) HasTypeUuid() bool`

HasTypeUuid returns a boolean if a field has been set.

### SetTypeUuidNil

`func (o *DiskPartition) SetTypeUuidNil(b bool)`

 SetTypeUuidNil sets the value for TypeUuid to be an explicit nil

### UnsetTypeUuid
`func (o *DiskPartition) UnsetTypeUuid()`

UnsetTypeUuid ensures that no value is present for TypeUuid, not even an explicit nil
### GetUuid

`func (o *DiskPartition) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *DiskPartition) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *DiskPartition) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *DiskPartition) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *DiskPartition) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *DiskPartition) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


