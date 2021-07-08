# FilesystemVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disks** | Pointer to [**[]Disk**](Disk.md) | Array of Disks and Partitions.  Specifies information about all the disks and partitions needed to mount this logical volume. | [optional] 
**DisplayName** | Pointer to **NullableString** | Specifies a description about the filesystem. | [optional] 
**FilesystemType** | Pointer to **NullableString** | Specifies type of the filesystem on this volume. | [optional] 
**FilesystemUuid** | Pointer to **NullableString** | Specifies the uuid of the filesystem. | [optional] 
**IsSupported** | Pointer to **NullableBool** | If true, this is a supported filesystem volume type. | [optional] 
**LogicalVolume** | Pointer to [**NullableLogicalVolume**](LogicalVolume.md) | Specify attributes for a kLMV (Linux) or kLDM (Windows) filesystem. This field is set only for kLVM and kLDM volume types. | [optional] 
**LogicalVolumeType** | Pointer to **NullableString** | Specifies the type of logical volume such as kSimpleVolume, kLVM or kLDM. &#39;kSimpleVolume&#39; indicates a simple volume. Data can be used by just mounting the only one partition present on the disk. &#39;kLVM&#39; indicates a logical volume on Linux managed by a Logical Volume Manager. In order to access the data, deviceTree must be created based on the specification in logicalVolume.deviceTree. &#39;kLDM&#39; indicates a logical volume on Windows managed by Logical Disk Manager. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the volume such as /C. | [optional] 
**VolumeGuid** | Pointer to **NullableString** | VolumeGuid is the Volume guid. This is populated for kPhysical environments. | [optional] 

## Methods

### NewFilesystemVolume

`func NewFilesystemVolume() *FilesystemVolume`

NewFilesystemVolume instantiates a new FilesystemVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesystemVolumeWithDefaults

`func NewFilesystemVolumeWithDefaults() *FilesystemVolume`

NewFilesystemVolumeWithDefaults instantiates a new FilesystemVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisks

`func (o *FilesystemVolume) GetDisks() []Disk`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *FilesystemVolume) GetDisksOk() (*[]Disk, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *FilesystemVolume) SetDisks(v []Disk)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *FilesystemVolume) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### SetDisksNil

`func (o *FilesystemVolume) SetDisksNil(b bool)`

 SetDisksNil sets the value for Disks to be an explicit nil

### UnsetDisks
`func (o *FilesystemVolume) UnsetDisks()`

UnsetDisks ensures that no value is present for Disks, not even an explicit nil
### GetDisplayName

`func (o *FilesystemVolume) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FilesystemVolume) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FilesystemVolume) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FilesystemVolume) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *FilesystemVolume) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *FilesystemVolume) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetFilesystemType

`func (o *FilesystemVolume) GetFilesystemType() string`

GetFilesystemType returns the FilesystemType field if non-nil, zero value otherwise.

### GetFilesystemTypeOk

`func (o *FilesystemVolume) GetFilesystemTypeOk() (*string, bool)`

GetFilesystemTypeOk returns a tuple with the FilesystemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystemType

`func (o *FilesystemVolume) SetFilesystemType(v string)`

SetFilesystemType sets FilesystemType field to given value.

### HasFilesystemType

`func (o *FilesystemVolume) HasFilesystemType() bool`

HasFilesystemType returns a boolean if a field has been set.

### SetFilesystemTypeNil

`func (o *FilesystemVolume) SetFilesystemTypeNil(b bool)`

 SetFilesystemTypeNil sets the value for FilesystemType to be an explicit nil

### UnsetFilesystemType
`func (o *FilesystemVolume) UnsetFilesystemType()`

UnsetFilesystemType ensures that no value is present for FilesystemType, not even an explicit nil
### GetFilesystemUuid

`func (o *FilesystemVolume) GetFilesystemUuid() string`

GetFilesystemUuid returns the FilesystemUuid field if non-nil, zero value otherwise.

### GetFilesystemUuidOk

`func (o *FilesystemVolume) GetFilesystemUuidOk() (*string, bool)`

GetFilesystemUuidOk returns a tuple with the FilesystemUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystemUuid

`func (o *FilesystemVolume) SetFilesystemUuid(v string)`

SetFilesystemUuid sets FilesystemUuid field to given value.

### HasFilesystemUuid

`func (o *FilesystemVolume) HasFilesystemUuid() bool`

HasFilesystemUuid returns a boolean if a field has been set.

### SetFilesystemUuidNil

`func (o *FilesystemVolume) SetFilesystemUuidNil(b bool)`

 SetFilesystemUuidNil sets the value for FilesystemUuid to be an explicit nil

### UnsetFilesystemUuid
`func (o *FilesystemVolume) UnsetFilesystemUuid()`

UnsetFilesystemUuid ensures that no value is present for FilesystemUuid, not even an explicit nil
### GetIsSupported

`func (o *FilesystemVolume) GetIsSupported() bool`

GetIsSupported returns the IsSupported field if non-nil, zero value otherwise.

### GetIsSupportedOk

`func (o *FilesystemVolume) GetIsSupportedOk() (*bool, bool)`

GetIsSupportedOk returns a tuple with the IsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSupported

`func (o *FilesystemVolume) SetIsSupported(v bool)`

SetIsSupported sets IsSupported field to given value.

### HasIsSupported

`func (o *FilesystemVolume) HasIsSupported() bool`

HasIsSupported returns a boolean if a field has been set.

### SetIsSupportedNil

`func (o *FilesystemVolume) SetIsSupportedNil(b bool)`

 SetIsSupportedNil sets the value for IsSupported to be an explicit nil

### UnsetIsSupported
`func (o *FilesystemVolume) UnsetIsSupported()`

UnsetIsSupported ensures that no value is present for IsSupported, not even an explicit nil
### GetLogicalVolume

`func (o *FilesystemVolume) GetLogicalVolume() LogicalVolume`

GetLogicalVolume returns the LogicalVolume field if non-nil, zero value otherwise.

### GetLogicalVolumeOk

`func (o *FilesystemVolume) GetLogicalVolumeOk() (*LogicalVolume, bool)`

GetLogicalVolumeOk returns a tuple with the LogicalVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalVolume

`func (o *FilesystemVolume) SetLogicalVolume(v LogicalVolume)`

SetLogicalVolume sets LogicalVolume field to given value.

### HasLogicalVolume

`func (o *FilesystemVolume) HasLogicalVolume() bool`

HasLogicalVolume returns a boolean if a field has been set.

### SetLogicalVolumeNil

`func (o *FilesystemVolume) SetLogicalVolumeNil(b bool)`

 SetLogicalVolumeNil sets the value for LogicalVolume to be an explicit nil

### UnsetLogicalVolume
`func (o *FilesystemVolume) UnsetLogicalVolume()`

UnsetLogicalVolume ensures that no value is present for LogicalVolume, not even an explicit nil
### GetLogicalVolumeType

`func (o *FilesystemVolume) GetLogicalVolumeType() string`

GetLogicalVolumeType returns the LogicalVolumeType field if non-nil, zero value otherwise.

### GetLogicalVolumeTypeOk

`func (o *FilesystemVolume) GetLogicalVolumeTypeOk() (*string, bool)`

GetLogicalVolumeTypeOk returns a tuple with the LogicalVolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalVolumeType

`func (o *FilesystemVolume) SetLogicalVolumeType(v string)`

SetLogicalVolumeType sets LogicalVolumeType field to given value.

### HasLogicalVolumeType

`func (o *FilesystemVolume) HasLogicalVolumeType() bool`

HasLogicalVolumeType returns a boolean if a field has been set.

### SetLogicalVolumeTypeNil

`func (o *FilesystemVolume) SetLogicalVolumeTypeNil(b bool)`

 SetLogicalVolumeTypeNil sets the value for LogicalVolumeType to be an explicit nil

### UnsetLogicalVolumeType
`func (o *FilesystemVolume) UnsetLogicalVolumeType()`

UnsetLogicalVolumeType ensures that no value is present for LogicalVolumeType, not even an explicit nil
### GetName

`func (o *FilesystemVolume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FilesystemVolume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FilesystemVolume) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FilesystemVolume) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FilesystemVolume) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FilesystemVolume) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetVolumeGuid

`func (o *FilesystemVolume) GetVolumeGuid() string`

GetVolumeGuid returns the VolumeGuid field if non-nil, zero value otherwise.

### GetVolumeGuidOk

`func (o *FilesystemVolume) GetVolumeGuidOk() (*string, bool)`

GetVolumeGuidOk returns a tuple with the VolumeGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGuid

`func (o *FilesystemVolume) SetVolumeGuid(v string)`

SetVolumeGuid sets VolumeGuid field to given value.

### HasVolumeGuid

`func (o *FilesystemVolume) HasVolumeGuid() bool`

HasVolumeGuid returns a boolean if a field has been set.

### SetVolumeGuidNil

`func (o *FilesystemVolume) SetVolumeGuidNil(b bool)`

 SetVolumeGuidNil sets the value for VolumeGuid to be an explicit nil

### UnsetVolumeGuid
`func (o *FilesystemVolume) UnsetVolumeGuid()`

UnsetVolumeGuid ensures that no value is present for VolumeGuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


