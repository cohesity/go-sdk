# DeviceTreePartitionSlice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskFileName** | Pointer to **NullableString** | The disk to use. | [optional] 
**Length** | Pointer to **NullableInt64** | The length of data for the LVM volume (for which this device tree is being built) in bytes. It does not include size of the LVM meta data. | [optional] 
**LvmDataOffset** | Pointer to **NullableInt64** | Each LVM partition starts with LVM meta data. After the meta data there can be data for one or more LVM volumes.  This field indicates the offset in bytes (relative to partition) where data for various LVM volumes starts on the partition. NOTE: If this device tree represents first LVM volume on the  partition, &#39;lvm_data_offset&#39; is equal to &#39;offset&#39;. | [optional] 
**Offset** | Pointer to **NullableInt64** | This is the offset (in bytes) where data for the LVM volume (for which this device tree is being build) starts relative to the start of the partition above. | [optional] 
**PartitionNumber** | Pointer to **NullableInt32** | The partition to use in the disk above. | [optional] 

## Methods

### NewDeviceTreePartitionSlice

`func NewDeviceTreePartitionSlice() *DeviceTreePartitionSlice`

NewDeviceTreePartitionSlice instantiates a new DeviceTreePartitionSlice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceTreePartitionSliceWithDefaults

`func NewDeviceTreePartitionSliceWithDefaults() *DeviceTreePartitionSlice`

NewDeviceTreePartitionSliceWithDefaults instantiates a new DeviceTreePartitionSlice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskFileName

`func (o *DeviceTreePartitionSlice) GetDiskFileName() string`

GetDiskFileName returns the DiskFileName field if non-nil, zero value otherwise.

### GetDiskFileNameOk

`func (o *DeviceTreePartitionSlice) GetDiskFileNameOk() (*string, bool)`

GetDiskFileNameOk returns a tuple with the DiskFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskFileName

`func (o *DeviceTreePartitionSlice) SetDiskFileName(v string)`

SetDiskFileName sets DiskFileName field to given value.

### HasDiskFileName

`func (o *DeviceTreePartitionSlice) HasDiskFileName() bool`

HasDiskFileName returns a boolean if a field has been set.

### SetDiskFileNameNil

`func (o *DeviceTreePartitionSlice) SetDiskFileNameNil(b bool)`

 SetDiskFileNameNil sets the value for DiskFileName to be an explicit nil

### UnsetDiskFileName
`func (o *DeviceTreePartitionSlice) UnsetDiskFileName()`

UnsetDiskFileName ensures that no value is present for DiskFileName, not even an explicit nil
### GetLength

`func (o *DeviceTreePartitionSlice) GetLength() int64`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *DeviceTreePartitionSlice) GetLengthOk() (*int64, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *DeviceTreePartitionSlice) SetLength(v int64)`

SetLength sets Length field to given value.

### HasLength

`func (o *DeviceTreePartitionSlice) HasLength() bool`

HasLength returns a boolean if a field has been set.

### SetLengthNil

`func (o *DeviceTreePartitionSlice) SetLengthNil(b bool)`

 SetLengthNil sets the value for Length to be an explicit nil

### UnsetLength
`func (o *DeviceTreePartitionSlice) UnsetLength()`

UnsetLength ensures that no value is present for Length, not even an explicit nil
### GetLvmDataOffset

`func (o *DeviceTreePartitionSlice) GetLvmDataOffset() int64`

GetLvmDataOffset returns the LvmDataOffset field if non-nil, zero value otherwise.

### GetLvmDataOffsetOk

`func (o *DeviceTreePartitionSlice) GetLvmDataOffsetOk() (*int64, bool)`

GetLvmDataOffsetOk returns a tuple with the LvmDataOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLvmDataOffset

`func (o *DeviceTreePartitionSlice) SetLvmDataOffset(v int64)`

SetLvmDataOffset sets LvmDataOffset field to given value.

### HasLvmDataOffset

`func (o *DeviceTreePartitionSlice) HasLvmDataOffset() bool`

HasLvmDataOffset returns a boolean if a field has been set.

### SetLvmDataOffsetNil

`func (o *DeviceTreePartitionSlice) SetLvmDataOffsetNil(b bool)`

 SetLvmDataOffsetNil sets the value for LvmDataOffset to be an explicit nil

### UnsetLvmDataOffset
`func (o *DeviceTreePartitionSlice) UnsetLvmDataOffset()`

UnsetLvmDataOffset ensures that no value is present for LvmDataOffset, not even an explicit nil
### GetOffset

`func (o *DeviceTreePartitionSlice) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *DeviceTreePartitionSlice) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *DeviceTreePartitionSlice) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *DeviceTreePartitionSlice) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *DeviceTreePartitionSlice) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *DeviceTreePartitionSlice) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetPartitionNumber

`func (o *DeviceTreePartitionSlice) GetPartitionNumber() int32`

GetPartitionNumber returns the PartitionNumber field if non-nil, zero value otherwise.

### GetPartitionNumberOk

`func (o *DeviceTreePartitionSlice) GetPartitionNumberOk() (*int32, bool)`

GetPartitionNumberOk returns a tuple with the PartitionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionNumber

`func (o *DeviceTreePartitionSlice) SetPartitionNumber(v int32)`

SetPartitionNumber sets PartitionNumber field to given value.

### HasPartitionNumber

`func (o *DeviceTreePartitionSlice) HasPartitionNumber() bool`

HasPartitionNumber returns a boolean if a field has been set.

### SetPartitionNumberNil

`func (o *DeviceTreePartitionSlice) SetPartitionNumberNil(b bool)`

 SetPartitionNumberNil sets the value for PartitionNumber to be an explicit nil

### UnsetPartitionNumber
`func (o *DeviceTreePartitionSlice) UnsetPartitionNumber()`

UnsetPartitionNumber ensures that no value is present for PartitionNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


