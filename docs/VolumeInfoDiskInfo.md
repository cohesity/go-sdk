# VolumeInfoDiskInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskFileName** | Pointer to **NullableString** | Disk name. This is the vmdk names, and not the flat file name. | [optional] 
**DiskFormat** | Pointer to **NullableInt32** | Disk format type of this file. See util/disklib/base/enums.proto for available types. | [optional] 
**DiskUuid** | Pointer to **NullableString** | Disk uuid. | [optional] 
**PartitionType** | Pointer to **NullableInt32** | Disk partition type. | [optional] 
**PartitionVec** | Pointer to [**[]VolumeInfoDiskInfoPartitionInfo**](VolumeInfoDiskInfoPartitionInfo.md) | Information about all the partitions in this disk. | [optional] 
**PhysicalRangeVec** | Pointer to [**[]VolumeInfoDiskInfoPhysicalRange**](VolumeInfoDiskInfoPhysicalRange.md) | This disk is formed by following physical ranges. Ranges are arranged sequentially to form a disk. | [optional] 
**SectorSize** | Pointer to **NullableInt64** | Sector size of disk. This is sector size of disk which is formed by mapping the physical ranges of the disk into a linear device. | [optional] 
**VmdkSize** | Pointer to **NullableInt64** | Disk size in bytes. | [optional] 

## Methods

### NewVolumeInfoDiskInfo

`func NewVolumeInfoDiskInfo() *VolumeInfoDiskInfo`

NewVolumeInfoDiskInfo instantiates a new VolumeInfoDiskInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeInfoDiskInfoWithDefaults

`func NewVolumeInfoDiskInfoWithDefaults() *VolumeInfoDiskInfo`

NewVolumeInfoDiskInfoWithDefaults instantiates a new VolumeInfoDiskInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskFileName

`func (o *VolumeInfoDiskInfo) GetDiskFileName() string`

GetDiskFileName returns the DiskFileName field if non-nil, zero value otherwise.

### GetDiskFileNameOk

`func (o *VolumeInfoDiskInfo) GetDiskFileNameOk() (*string, bool)`

GetDiskFileNameOk returns a tuple with the DiskFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskFileName

`func (o *VolumeInfoDiskInfo) SetDiskFileName(v string)`

SetDiskFileName sets DiskFileName field to given value.

### HasDiskFileName

`func (o *VolumeInfoDiskInfo) HasDiskFileName() bool`

HasDiskFileName returns a boolean if a field has been set.

### SetDiskFileNameNil

`func (o *VolumeInfoDiskInfo) SetDiskFileNameNil(b bool)`

 SetDiskFileNameNil sets the value for DiskFileName to be an explicit nil

### UnsetDiskFileName
`func (o *VolumeInfoDiskInfo) UnsetDiskFileName()`

UnsetDiskFileName ensures that no value is present for DiskFileName, not even an explicit nil
### GetDiskFormat

`func (o *VolumeInfoDiskInfo) GetDiskFormat() int32`

GetDiskFormat returns the DiskFormat field if non-nil, zero value otherwise.

### GetDiskFormatOk

`func (o *VolumeInfoDiskInfo) GetDiskFormatOk() (*int32, bool)`

GetDiskFormatOk returns a tuple with the DiskFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskFormat

`func (o *VolumeInfoDiskInfo) SetDiskFormat(v int32)`

SetDiskFormat sets DiskFormat field to given value.

### HasDiskFormat

`func (o *VolumeInfoDiskInfo) HasDiskFormat() bool`

HasDiskFormat returns a boolean if a field has been set.

### SetDiskFormatNil

`func (o *VolumeInfoDiskInfo) SetDiskFormatNil(b bool)`

 SetDiskFormatNil sets the value for DiskFormat to be an explicit nil

### UnsetDiskFormat
`func (o *VolumeInfoDiskInfo) UnsetDiskFormat()`

UnsetDiskFormat ensures that no value is present for DiskFormat, not even an explicit nil
### GetDiskUuid

`func (o *VolumeInfoDiskInfo) GetDiskUuid() string`

GetDiskUuid returns the DiskUuid field if non-nil, zero value otherwise.

### GetDiskUuidOk

`func (o *VolumeInfoDiskInfo) GetDiskUuidOk() (*string, bool)`

GetDiskUuidOk returns a tuple with the DiskUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUuid

`func (o *VolumeInfoDiskInfo) SetDiskUuid(v string)`

SetDiskUuid sets DiskUuid field to given value.

### HasDiskUuid

`func (o *VolumeInfoDiskInfo) HasDiskUuid() bool`

HasDiskUuid returns a boolean if a field has been set.

### SetDiskUuidNil

`func (o *VolumeInfoDiskInfo) SetDiskUuidNil(b bool)`

 SetDiskUuidNil sets the value for DiskUuid to be an explicit nil

### UnsetDiskUuid
`func (o *VolumeInfoDiskInfo) UnsetDiskUuid()`

UnsetDiskUuid ensures that no value is present for DiskUuid, not even an explicit nil
### GetPartitionType

`func (o *VolumeInfoDiskInfo) GetPartitionType() int32`

GetPartitionType returns the PartitionType field if non-nil, zero value otherwise.

### GetPartitionTypeOk

`func (o *VolumeInfoDiskInfo) GetPartitionTypeOk() (*int32, bool)`

GetPartitionTypeOk returns a tuple with the PartitionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionType

`func (o *VolumeInfoDiskInfo) SetPartitionType(v int32)`

SetPartitionType sets PartitionType field to given value.

### HasPartitionType

`func (o *VolumeInfoDiskInfo) HasPartitionType() bool`

HasPartitionType returns a boolean if a field has been set.

### SetPartitionTypeNil

`func (o *VolumeInfoDiskInfo) SetPartitionTypeNil(b bool)`

 SetPartitionTypeNil sets the value for PartitionType to be an explicit nil

### UnsetPartitionType
`func (o *VolumeInfoDiskInfo) UnsetPartitionType()`

UnsetPartitionType ensures that no value is present for PartitionType, not even an explicit nil
### GetPartitionVec

`func (o *VolumeInfoDiskInfo) GetPartitionVec() []VolumeInfoDiskInfoPartitionInfo`

GetPartitionVec returns the PartitionVec field if non-nil, zero value otherwise.

### GetPartitionVecOk

`func (o *VolumeInfoDiskInfo) GetPartitionVecOk() (*[]VolumeInfoDiskInfoPartitionInfo, bool)`

GetPartitionVecOk returns a tuple with the PartitionVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionVec

`func (o *VolumeInfoDiskInfo) SetPartitionVec(v []VolumeInfoDiskInfoPartitionInfo)`

SetPartitionVec sets PartitionVec field to given value.

### HasPartitionVec

`func (o *VolumeInfoDiskInfo) HasPartitionVec() bool`

HasPartitionVec returns a boolean if a field has been set.

### SetPartitionVecNil

`func (o *VolumeInfoDiskInfo) SetPartitionVecNil(b bool)`

 SetPartitionVecNil sets the value for PartitionVec to be an explicit nil

### UnsetPartitionVec
`func (o *VolumeInfoDiskInfo) UnsetPartitionVec()`

UnsetPartitionVec ensures that no value is present for PartitionVec, not even an explicit nil
### GetPhysicalRangeVec

`func (o *VolumeInfoDiskInfo) GetPhysicalRangeVec() []VolumeInfoDiskInfoPhysicalRange`

GetPhysicalRangeVec returns the PhysicalRangeVec field if non-nil, zero value otherwise.

### GetPhysicalRangeVecOk

`func (o *VolumeInfoDiskInfo) GetPhysicalRangeVecOk() (*[]VolumeInfoDiskInfoPhysicalRange, bool)`

GetPhysicalRangeVecOk returns a tuple with the PhysicalRangeVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalRangeVec

`func (o *VolumeInfoDiskInfo) SetPhysicalRangeVec(v []VolumeInfoDiskInfoPhysicalRange)`

SetPhysicalRangeVec sets PhysicalRangeVec field to given value.

### HasPhysicalRangeVec

`func (o *VolumeInfoDiskInfo) HasPhysicalRangeVec() bool`

HasPhysicalRangeVec returns a boolean if a field has been set.

### SetPhysicalRangeVecNil

`func (o *VolumeInfoDiskInfo) SetPhysicalRangeVecNil(b bool)`

 SetPhysicalRangeVecNil sets the value for PhysicalRangeVec to be an explicit nil

### UnsetPhysicalRangeVec
`func (o *VolumeInfoDiskInfo) UnsetPhysicalRangeVec()`

UnsetPhysicalRangeVec ensures that no value is present for PhysicalRangeVec, not even an explicit nil
### GetSectorSize

`func (o *VolumeInfoDiskInfo) GetSectorSize() int64`

GetSectorSize returns the SectorSize field if non-nil, zero value otherwise.

### GetSectorSizeOk

`func (o *VolumeInfoDiskInfo) GetSectorSizeOk() (*int64, bool)`

GetSectorSizeOk returns a tuple with the SectorSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectorSize

`func (o *VolumeInfoDiskInfo) SetSectorSize(v int64)`

SetSectorSize sets SectorSize field to given value.

### HasSectorSize

`func (o *VolumeInfoDiskInfo) HasSectorSize() bool`

HasSectorSize returns a boolean if a field has been set.

### SetSectorSizeNil

`func (o *VolumeInfoDiskInfo) SetSectorSizeNil(b bool)`

 SetSectorSizeNil sets the value for SectorSize to be an explicit nil

### UnsetSectorSize
`func (o *VolumeInfoDiskInfo) UnsetSectorSize()`

UnsetSectorSize ensures that no value is present for SectorSize, not even an explicit nil
### GetVmdkSize

`func (o *VolumeInfoDiskInfo) GetVmdkSize() int64`

GetVmdkSize returns the VmdkSize field if non-nil, zero value otherwise.

### GetVmdkSizeOk

`func (o *VolumeInfoDiskInfo) GetVmdkSizeOk() (*int64, bool)`

GetVmdkSizeOk returns a tuple with the VmdkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmdkSize

`func (o *VolumeInfoDiskInfo) SetVmdkSize(v int64)`

SetVmdkSize sets VmdkSize field to given value.

### HasVmdkSize

`func (o *VolumeInfoDiskInfo) HasVmdkSize() bool`

HasVmdkSize returns a boolean if a field has been set.

### SetVmdkSizeNil

`func (o *VolumeInfoDiskInfo) SetVmdkSizeNil(b bool)`

 SetVmdkSizeNil sets the value for VmdkSize to be an explicit nil

### UnsetVmdkSize
`func (o *VolumeInfoDiskInfo) UnsetVmdkSize()`

UnsetVmdkSize ensures that no value is present for VmdkSize, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


