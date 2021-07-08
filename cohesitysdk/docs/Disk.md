# Disk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskBlocks** | Pointer to [**[]DiskBlock**](DiskBlock.md) | Array of Disk Blocks.  Specifies a set of disk blocks by defining the location and offset of disk blocks in a disk. | [optional] 
**DiskFormat** | Pointer to **NullableString** | Specifies the format of the virtual disk. &#39;kVMDK&#39; indicates VMware&#39;s Virtual Disk format. &#39;kVHD&#39; indicates Microsoft&#39;s Virtual Hard Drive format. &#39;kVHDx&#39; indicates Microsoft&#39;s Hyper-V Virtual Hard Drive format. &#39;kRaw&#39; indicates Raw disk format used by KVM, Acropolis. &#39;kUnknow&#39; indicates Unknown disk format. | [optional] 
**DiskPartitions** | Pointer to [**[]DiskPartition**](DiskPartition.md) | Array of Partitions.  Specifies information about all the partitions in this disk. | [optional] 
**PartitionTableFormat** | Pointer to **NullableString** | Specifies partition table format on a disk. &#39;kNoPartition&#39; indicates missing partition table. &#39;kMBRPartition&#39; indicates partition table is in Master Boot Record format. &#39;kGPTPartition&#39; indicates partition table is in Guid Partition Table format. &#39;kSGIPartition&#39; indicates partition table uses SGI scheme. &#39;kSUNPartition&#39; indicates partition table uses SUN scheme. | [optional] 
**SectorSizeBytes** | Pointer to **NullableInt64** | Specifies the sector size of hard disk. It is used for mapping the disk blocks of the disk file into a linear list of sectors. | [optional] 
**Uuid** | Pointer to **NullableString** | Specifies the disk uuid. | [optional] 
**VmdkFileName** | Pointer to **NullableString** | Specifies the disk file name. This is the VMDK name and not the flat file name. | [optional] 
**VmdkSizeBytes** | Pointer to **NullableInt64** | Specifies the disk size in bytes. | [optional] 

## Methods

### NewDisk

`func NewDisk() *Disk`

NewDisk instantiates a new Disk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskWithDefaults

`func NewDiskWithDefaults() *Disk`

NewDiskWithDefaults instantiates a new Disk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskBlocks

`func (o *Disk) GetDiskBlocks() []DiskBlock`

GetDiskBlocks returns the DiskBlocks field if non-nil, zero value otherwise.

### GetDiskBlocksOk

`func (o *Disk) GetDiskBlocksOk() (*[]DiskBlock, bool)`

GetDiskBlocksOk returns a tuple with the DiskBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskBlocks

`func (o *Disk) SetDiskBlocks(v []DiskBlock)`

SetDiskBlocks sets DiskBlocks field to given value.

### HasDiskBlocks

`func (o *Disk) HasDiskBlocks() bool`

HasDiskBlocks returns a boolean if a field has been set.

### SetDiskBlocksNil

`func (o *Disk) SetDiskBlocksNil(b bool)`

 SetDiskBlocksNil sets the value for DiskBlocks to be an explicit nil

### UnsetDiskBlocks
`func (o *Disk) UnsetDiskBlocks()`

UnsetDiskBlocks ensures that no value is present for DiskBlocks, not even an explicit nil
### GetDiskFormat

`func (o *Disk) GetDiskFormat() string`

GetDiskFormat returns the DiskFormat field if non-nil, zero value otherwise.

### GetDiskFormatOk

`func (o *Disk) GetDiskFormatOk() (*string, bool)`

GetDiskFormatOk returns a tuple with the DiskFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskFormat

`func (o *Disk) SetDiskFormat(v string)`

SetDiskFormat sets DiskFormat field to given value.

### HasDiskFormat

`func (o *Disk) HasDiskFormat() bool`

HasDiskFormat returns a boolean if a field has been set.

### SetDiskFormatNil

`func (o *Disk) SetDiskFormatNil(b bool)`

 SetDiskFormatNil sets the value for DiskFormat to be an explicit nil

### UnsetDiskFormat
`func (o *Disk) UnsetDiskFormat()`

UnsetDiskFormat ensures that no value is present for DiskFormat, not even an explicit nil
### GetDiskPartitions

`func (o *Disk) GetDiskPartitions() []DiskPartition`

GetDiskPartitions returns the DiskPartitions field if non-nil, zero value otherwise.

### GetDiskPartitionsOk

`func (o *Disk) GetDiskPartitionsOk() (*[]DiskPartition, bool)`

GetDiskPartitionsOk returns a tuple with the DiskPartitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskPartitions

`func (o *Disk) SetDiskPartitions(v []DiskPartition)`

SetDiskPartitions sets DiskPartitions field to given value.

### HasDiskPartitions

`func (o *Disk) HasDiskPartitions() bool`

HasDiskPartitions returns a boolean if a field has been set.

### SetDiskPartitionsNil

`func (o *Disk) SetDiskPartitionsNil(b bool)`

 SetDiskPartitionsNil sets the value for DiskPartitions to be an explicit nil

### UnsetDiskPartitions
`func (o *Disk) UnsetDiskPartitions()`

UnsetDiskPartitions ensures that no value is present for DiskPartitions, not even an explicit nil
### GetPartitionTableFormat

`func (o *Disk) GetPartitionTableFormat() string`

GetPartitionTableFormat returns the PartitionTableFormat field if non-nil, zero value otherwise.

### GetPartitionTableFormatOk

`func (o *Disk) GetPartitionTableFormatOk() (*string, bool)`

GetPartitionTableFormatOk returns a tuple with the PartitionTableFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionTableFormat

`func (o *Disk) SetPartitionTableFormat(v string)`

SetPartitionTableFormat sets PartitionTableFormat field to given value.

### HasPartitionTableFormat

`func (o *Disk) HasPartitionTableFormat() bool`

HasPartitionTableFormat returns a boolean if a field has been set.

### SetPartitionTableFormatNil

`func (o *Disk) SetPartitionTableFormatNil(b bool)`

 SetPartitionTableFormatNil sets the value for PartitionTableFormat to be an explicit nil

### UnsetPartitionTableFormat
`func (o *Disk) UnsetPartitionTableFormat()`

UnsetPartitionTableFormat ensures that no value is present for PartitionTableFormat, not even an explicit nil
### GetSectorSizeBytes

`func (o *Disk) GetSectorSizeBytes() int64`

GetSectorSizeBytes returns the SectorSizeBytes field if non-nil, zero value otherwise.

### GetSectorSizeBytesOk

`func (o *Disk) GetSectorSizeBytesOk() (*int64, bool)`

GetSectorSizeBytesOk returns a tuple with the SectorSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectorSizeBytes

`func (o *Disk) SetSectorSizeBytes(v int64)`

SetSectorSizeBytes sets SectorSizeBytes field to given value.

### HasSectorSizeBytes

`func (o *Disk) HasSectorSizeBytes() bool`

HasSectorSizeBytes returns a boolean if a field has been set.

### SetSectorSizeBytesNil

`func (o *Disk) SetSectorSizeBytesNil(b bool)`

 SetSectorSizeBytesNil sets the value for SectorSizeBytes to be an explicit nil

### UnsetSectorSizeBytes
`func (o *Disk) UnsetSectorSizeBytes()`

UnsetSectorSizeBytes ensures that no value is present for SectorSizeBytes, not even an explicit nil
### GetUuid

`func (o *Disk) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Disk) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Disk) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Disk) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *Disk) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *Disk) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetVmdkFileName

`func (o *Disk) GetVmdkFileName() string`

GetVmdkFileName returns the VmdkFileName field if non-nil, zero value otherwise.

### GetVmdkFileNameOk

`func (o *Disk) GetVmdkFileNameOk() (*string, bool)`

GetVmdkFileNameOk returns a tuple with the VmdkFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmdkFileName

`func (o *Disk) SetVmdkFileName(v string)`

SetVmdkFileName sets VmdkFileName field to given value.

### HasVmdkFileName

`func (o *Disk) HasVmdkFileName() bool`

HasVmdkFileName returns a boolean if a field has been set.

### SetVmdkFileNameNil

`func (o *Disk) SetVmdkFileNameNil(b bool)`

 SetVmdkFileNameNil sets the value for VmdkFileName to be an explicit nil

### UnsetVmdkFileName
`func (o *Disk) UnsetVmdkFileName()`

UnsetVmdkFileName ensures that no value is present for VmdkFileName, not even an explicit nil
### GetVmdkSizeBytes

`func (o *Disk) GetVmdkSizeBytes() int64`

GetVmdkSizeBytes returns the VmdkSizeBytes field if non-nil, zero value otherwise.

### GetVmdkSizeBytesOk

`func (o *Disk) GetVmdkSizeBytesOk() (*int64, bool)`

GetVmdkSizeBytesOk returns a tuple with the VmdkSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmdkSizeBytes

`func (o *Disk) SetVmdkSizeBytes(v int64)`

SetVmdkSizeBytes sets VmdkSizeBytes field to given value.

### HasVmdkSizeBytes

`func (o *Disk) HasVmdkSizeBytes() bool`

HasVmdkSizeBytes returns a boolean if a field has been set.

### SetVmdkSizeBytesNil

`func (o *Disk) SetVmdkSizeBytesNil(b bool)`

 SetVmdkSizeBytesNil sets the value for VmdkSizeBytes to be an explicit nil

### UnsetVmdkSizeBytes
`func (o *Disk) UnsetVmdkSizeBytes()`

UnsetVmdkSizeBytes ensures that no value is present for VmdkSizeBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


