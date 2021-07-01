# VolumeInfoDiskInfoPartitionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Length** | Pointer to **NullableInt64** | Length of partition in bytes. | [optional] 
**PartitionNumber** | Pointer to **NullableInt64** | Partition number. | [optional] 
**PartitionTypeUuid** | Pointer to **NullableString** | Partition type uuid. If disk is unpartitioned, this field will not be set. If disk is MBR partitioned, this field will be set to partition type. Example: 83 (from below fdisk output) [This value is in hex] bash$ fdisk -l foobar.vmdk Device        Boot Start   End    Sectors  Size Id Type foobar.vmdk1       2048  1050623  1048576  512M 83 Linux If disk is GPT partitioned, this field will be set to partition type GUID. Example: fc63daf-8483-4772-8e793d69d8477de4 (Linux filesystem data) | [optional] 
**PartitionUuid** | Pointer to **NullableString** | Partition uuid. If disk is unpartitioned, this field will not be set. If disk is MBR partitioned, this field will not be set. If disk is GPT partitioned, this field will be set to partition GUID. | [optional] 
**StartOffset** | Pointer to **NullableInt64** | Start offset of partition in bytes. | [optional] 

## Methods

### NewVolumeInfoDiskInfoPartitionInfo

`func NewVolumeInfoDiskInfoPartitionInfo() *VolumeInfoDiskInfoPartitionInfo`

NewVolumeInfoDiskInfoPartitionInfo instantiates a new VolumeInfoDiskInfoPartitionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeInfoDiskInfoPartitionInfoWithDefaults

`func NewVolumeInfoDiskInfoPartitionInfoWithDefaults() *VolumeInfoDiskInfoPartitionInfo`

NewVolumeInfoDiskInfoPartitionInfoWithDefaults instantiates a new VolumeInfoDiskInfoPartitionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLength

`func (o *VolumeInfoDiskInfoPartitionInfo) GetLength() int64`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *VolumeInfoDiskInfoPartitionInfo) GetLengthOk() (*int64, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *VolumeInfoDiskInfoPartitionInfo) SetLength(v int64)`

SetLength sets Length field to given value.

### HasLength

`func (o *VolumeInfoDiskInfoPartitionInfo) HasLength() bool`

HasLength returns a boolean if a field has been set.

### SetLengthNil

`func (o *VolumeInfoDiskInfoPartitionInfo) SetLengthNil(b bool)`

 SetLengthNil sets the value for Length to be an explicit nil

### UnsetLength
`func (o *VolumeInfoDiskInfoPartitionInfo) UnsetLength()`

UnsetLength ensures that no value is present for Length, not even an explicit nil
### GetPartitionNumber

`func (o *VolumeInfoDiskInfoPartitionInfo) GetPartitionNumber() int64`

GetPartitionNumber returns the PartitionNumber field if non-nil, zero value otherwise.

### GetPartitionNumberOk

`func (o *VolumeInfoDiskInfoPartitionInfo) GetPartitionNumberOk() (*int64, bool)`

GetPartitionNumberOk returns a tuple with the PartitionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionNumber

`func (o *VolumeInfoDiskInfoPartitionInfo) SetPartitionNumber(v int64)`

SetPartitionNumber sets PartitionNumber field to given value.

### HasPartitionNumber

`func (o *VolumeInfoDiskInfoPartitionInfo) HasPartitionNumber() bool`

HasPartitionNumber returns a boolean if a field has been set.

### SetPartitionNumberNil

`func (o *VolumeInfoDiskInfoPartitionInfo) SetPartitionNumberNil(b bool)`

 SetPartitionNumberNil sets the value for PartitionNumber to be an explicit nil

### UnsetPartitionNumber
`func (o *VolumeInfoDiskInfoPartitionInfo) UnsetPartitionNumber()`

UnsetPartitionNumber ensures that no value is present for PartitionNumber, not even an explicit nil
### GetPartitionTypeUuid

`func (o *VolumeInfoDiskInfoPartitionInfo) GetPartitionTypeUuid() string`

GetPartitionTypeUuid returns the PartitionTypeUuid field if non-nil, zero value otherwise.

### GetPartitionTypeUuidOk

`func (o *VolumeInfoDiskInfoPartitionInfo) GetPartitionTypeUuidOk() (*string, bool)`

GetPartitionTypeUuidOk returns a tuple with the PartitionTypeUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionTypeUuid

`func (o *VolumeInfoDiskInfoPartitionInfo) SetPartitionTypeUuid(v string)`

SetPartitionTypeUuid sets PartitionTypeUuid field to given value.

### HasPartitionTypeUuid

`func (o *VolumeInfoDiskInfoPartitionInfo) HasPartitionTypeUuid() bool`

HasPartitionTypeUuid returns a boolean if a field has been set.

### SetPartitionTypeUuidNil

`func (o *VolumeInfoDiskInfoPartitionInfo) SetPartitionTypeUuidNil(b bool)`

 SetPartitionTypeUuidNil sets the value for PartitionTypeUuid to be an explicit nil

### UnsetPartitionTypeUuid
`func (o *VolumeInfoDiskInfoPartitionInfo) UnsetPartitionTypeUuid()`

UnsetPartitionTypeUuid ensures that no value is present for PartitionTypeUuid, not even an explicit nil
### GetPartitionUuid

`func (o *VolumeInfoDiskInfoPartitionInfo) GetPartitionUuid() string`

GetPartitionUuid returns the PartitionUuid field if non-nil, zero value otherwise.

### GetPartitionUuidOk

`func (o *VolumeInfoDiskInfoPartitionInfo) GetPartitionUuidOk() (*string, bool)`

GetPartitionUuidOk returns a tuple with the PartitionUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionUuid

`func (o *VolumeInfoDiskInfoPartitionInfo) SetPartitionUuid(v string)`

SetPartitionUuid sets PartitionUuid field to given value.

### HasPartitionUuid

`func (o *VolumeInfoDiskInfoPartitionInfo) HasPartitionUuid() bool`

HasPartitionUuid returns a boolean if a field has been set.

### SetPartitionUuidNil

`func (o *VolumeInfoDiskInfoPartitionInfo) SetPartitionUuidNil(b bool)`

 SetPartitionUuidNil sets the value for PartitionUuid to be an explicit nil

### UnsetPartitionUuid
`func (o *VolumeInfoDiskInfoPartitionInfo) UnsetPartitionUuid()`

UnsetPartitionUuid ensures that no value is present for PartitionUuid, not even an explicit nil
### GetStartOffset

`func (o *VolumeInfoDiskInfoPartitionInfo) GetStartOffset() int64`

GetStartOffset returns the StartOffset field if non-nil, zero value otherwise.

### GetStartOffsetOk

`func (o *VolumeInfoDiskInfoPartitionInfo) GetStartOffsetOk() (*int64, bool)`

GetStartOffsetOk returns a tuple with the StartOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOffset

`func (o *VolumeInfoDiskInfoPartitionInfo) SetStartOffset(v int64)`

SetStartOffset sets StartOffset field to given value.

### HasStartOffset

`func (o *VolumeInfoDiskInfoPartitionInfo) HasStartOffset() bool`

HasStartOffset returns a boolean if a field has been set.

### SetStartOffsetNil

`func (o *VolumeInfoDiskInfoPartitionInfo) SetStartOffsetNil(b bool)`

 SetStartOffsetNil sets the value for StartOffset to be an explicit nil

### UnsetStartOffset
`func (o *VolumeInfoDiskInfoPartitionInfo) UnsetStartOffset()`

UnsetStartOffset ensures that no value is present for StartOffset, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


