# DeviceTreeLeafNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskName** | Pointer to **NullableString** | Specifies the disk name. | [optional] 
**LengthBytes** | Pointer to **NullableInt64** | Specifies The length of data in bytes for the LVM volume (for which this device tree is being built). It does not include size of the LVM meta data. | [optional] 
**OffsetBytes** | Pointer to **NullableInt64** | Specifies the offset in bytes where data for the LVM volume (for which this device tree is being build) starts relative to the start of the partition. | [optional] 
**PartitionNumber** | Pointer to **NullableInt32** | Specifies the paritition number. | [optional] 

## Methods

### NewDeviceTreeLeafNode

`func NewDeviceTreeLeafNode() *DeviceTreeLeafNode`

NewDeviceTreeLeafNode instantiates a new DeviceTreeLeafNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceTreeLeafNodeWithDefaults

`func NewDeviceTreeLeafNodeWithDefaults() *DeviceTreeLeafNode`

NewDeviceTreeLeafNodeWithDefaults instantiates a new DeviceTreeLeafNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskName

`func (o *DeviceTreeLeafNode) GetDiskName() string`

GetDiskName returns the DiskName field if non-nil, zero value otherwise.

### GetDiskNameOk

`func (o *DeviceTreeLeafNode) GetDiskNameOk() (*string, bool)`

GetDiskNameOk returns a tuple with the DiskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskName

`func (o *DeviceTreeLeafNode) SetDiskName(v string)`

SetDiskName sets DiskName field to given value.

### HasDiskName

`func (o *DeviceTreeLeafNode) HasDiskName() bool`

HasDiskName returns a boolean if a field has been set.

### SetDiskNameNil

`func (o *DeviceTreeLeafNode) SetDiskNameNil(b bool)`

 SetDiskNameNil sets the value for DiskName to be an explicit nil

### UnsetDiskName
`func (o *DeviceTreeLeafNode) UnsetDiskName()`

UnsetDiskName ensures that no value is present for DiskName, not even an explicit nil
### GetLengthBytes

`func (o *DeviceTreeLeafNode) GetLengthBytes() int64`

GetLengthBytes returns the LengthBytes field if non-nil, zero value otherwise.

### GetLengthBytesOk

`func (o *DeviceTreeLeafNode) GetLengthBytesOk() (*int64, bool)`

GetLengthBytesOk returns a tuple with the LengthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthBytes

`func (o *DeviceTreeLeafNode) SetLengthBytes(v int64)`

SetLengthBytes sets LengthBytes field to given value.

### HasLengthBytes

`func (o *DeviceTreeLeafNode) HasLengthBytes() bool`

HasLengthBytes returns a boolean if a field has been set.

### SetLengthBytesNil

`func (o *DeviceTreeLeafNode) SetLengthBytesNil(b bool)`

 SetLengthBytesNil sets the value for LengthBytes to be an explicit nil

### UnsetLengthBytes
`func (o *DeviceTreeLeafNode) UnsetLengthBytes()`

UnsetLengthBytes ensures that no value is present for LengthBytes, not even an explicit nil
### GetOffsetBytes

`func (o *DeviceTreeLeafNode) GetOffsetBytes() int64`

GetOffsetBytes returns the OffsetBytes field if non-nil, zero value otherwise.

### GetOffsetBytesOk

`func (o *DeviceTreeLeafNode) GetOffsetBytesOk() (*int64, bool)`

GetOffsetBytesOk returns a tuple with the OffsetBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetBytes

`func (o *DeviceTreeLeafNode) SetOffsetBytes(v int64)`

SetOffsetBytes sets OffsetBytes field to given value.

### HasOffsetBytes

`func (o *DeviceTreeLeafNode) HasOffsetBytes() bool`

HasOffsetBytes returns a boolean if a field has been set.

### SetOffsetBytesNil

`func (o *DeviceTreeLeafNode) SetOffsetBytesNil(b bool)`

 SetOffsetBytesNil sets the value for OffsetBytes to be an explicit nil

### UnsetOffsetBytes
`func (o *DeviceTreeLeafNode) UnsetOffsetBytes()`

UnsetOffsetBytes ensures that no value is present for OffsetBytes, not even an explicit nil
### GetPartitionNumber

`func (o *DeviceTreeLeafNode) GetPartitionNumber() int32`

GetPartitionNumber returns the PartitionNumber field if non-nil, zero value otherwise.

### GetPartitionNumberOk

`func (o *DeviceTreeLeafNode) GetPartitionNumberOk() (*int32, bool)`

GetPartitionNumberOk returns a tuple with the PartitionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionNumber

`func (o *DeviceTreeLeafNode) SetPartitionNumber(v int32)`

SetPartitionNumber sets PartitionNumber field to given value.

### HasPartitionNumber

`func (o *DeviceTreeLeafNode) HasPartitionNumber() bool`

HasPartitionNumber returns a boolean if a field has been set.

### SetPartitionNumberNil

`func (o *DeviceTreeLeafNode) SetPartitionNumberNil(b bool)`

 SetPartitionNumberNil sets the value for PartitionNumber to be an explicit nil

### UnsetPartitionNumber
`func (o *DeviceTreeLeafNode) UnsetPartitionNumber()`

UnsetPartitionNumber ensures that no value is present for PartitionNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


