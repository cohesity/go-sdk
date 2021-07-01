# VirtualDiskInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusNumber** | Pointer to **NullableInt64** | Specifies the Id of the controller bus that controls the disk. | [optional] 
**ControllerType** | Pointer to **NullableString** | Specifies the controller type like SCSI, or IDE etc. | [optional] 
**DiskId** | Pointer to **NullableString** | Specifies original disk id. This is sufficient to identify the disk information, but in some scenarios, users may specify the controller option instead. | [optional] 
**DiskLocation** | Pointer to [**ProtectionSource**](ProtectionSource.md) |  | [optional] 
**DiskSizeInBytes** | Pointer to **NullableInt64** | Specifies size of the virtual disk in bytes. | [optional] 
**FilePath** | Pointer to **NullableString** | Specifies the original file path if applicable. | [optional] 
**MountPoints** | Pointer to **[]string** | Specifies the list of mount points. | [optional] 
**UnitNumber** | Pointer to **NullableInt64** | Specifies the disk file name. This is the VMDK name and not the flat file name. | [optional] 

## Methods

### NewVirtualDiskInformation

`func NewVirtualDiskInformation() *VirtualDiskInformation`

NewVirtualDiskInformation instantiates a new VirtualDiskInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualDiskInformationWithDefaults

`func NewVirtualDiskInformationWithDefaults() *VirtualDiskInformation`

NewVirtualDiskInformationWithDefaults instantiates a new VirtualDiskInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusNumber

`func (o *VirtualDiskInformation) GetBusNumber() int64`

GetBusNumber returns the BusNumber field if non-nil, zero value otherwise.

### GetBusNumberOk

`func (o *VirtualDiskInformation) GetBusNumberOk() (*int64, bool)`

GetBusNumberOk returns a tuple with the BusNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusNumber

`func (o *VirtualDiskInformation) SetBusNumber(v int64)`

SetBusNumber sets BusNumber field to given value.

### HasBusNumber

`func (o *VirtualDiskInformation) HasBusNumber() bool`

HasBusNumber returns a boolean if a field has been set.

### SetBusNumberNil

`func (o *VirtualDiskInformation) SetBusNumberNil(b bool)`

 SetBusNumberNil sets the value for BusNumber to be an explicit nil

### UnsetBusNumber
`func (o *VirtualDiskInformation) UnsetBusNumber()`

UnsetBusNumber ensures that no value is present for BusNumber, not even an explicit nil
### GetControllerType

`func (o *VirtualDiskInformation) GetControllerType() string`

GetControllerType returns the ControllerType field if non-nil, zero value otherwise.

### GetControllerTypeOk

`func (o *VirtualDiskInformation) GetControllerTypeOk() (*string, bool)`

GetControllerTypeOk returns a tuple with the ControllerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerType

`func (o *VirtualDiskInformation) SetControllerType(v string)`

SetControllerType sets ControllerType field to given value.

### HasControllerType

`func (o *VirtualDiskInformation) HasControllerType() bool`

HasControllerType returns a boolean if a field has been set.

### SetControllerTypeNil

`func (o *VirtualDiskInformation) SetControllerTypeNil(b bool)`

 SetControllerTypeNil sets the value for ControllerType to be an explicit nil

### UnsetControllerType
`func (o *VirtualDiskInformation) UnsetControllerType()`

UnsetControllerType ensures that no value is present for ControllerType, not even an explicit nil
### GetDiskId

`func (o *VirtualDiskInformation) GetDiskId() string`

GetDiskId returns the DiskId field if non-nil, zero value otherwise.

### GetDiskIdOk

`func (o *VirtualDiskInformation) GetDiskIdOk() (*string, bool)`

GetDiskIdOk returns a tuple with the DiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskId

`func (o *VirtualDiskInformation) SetDiskId(v string)`

SetDiskId sets DiskId field to given value.

### HasDiskId

`func (o *VirtualDiskInformation) HasDiskId() bool`

HasDiskId returns a boolean if a field has been set.

### SetDiskIdNil

`func (o *VirtualDiskInformation) SetDiskIdNil(b bool)`

 SetDiskIdNil sets the value for DiskId to be an explicit nil

### UnsetDiskId
`func (o *VirtualDiskInformation) UnsetDiskId()`

UnsetDiskId ensures that no value is present for DiskId, not even an explicit nil
### GetDiskLocation

`func (o *VirtualDiskInformation) GetDiskLocation() ProtectionSource`

GetDiskLocation returns the DiskLocation field if non-nil, zero value otherwise.

### GetDiskLocationOk

`func (o *VirtualDiskInformation) GetDiskLocationOk() (*ProtectionSource, bool)`

GetDiskLocationOk returns a tuple with the DiskLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskLocation

`func (o *VirtualDiskInformation) SetDiskLocation(v ProtectionSource)`

SetDiskLocation sets DiskLocation field to given value.

### HasDiskLocation

`func (o *VirtualDiskInformation) HasDiskLocation() bool`

HasDiskLocation returns a boolean if a field has been set.

### GetDiskSizeInBytes

`func (o *VirtualDiskInformation) GetDiskSizeInBytes() int64`

GetDiskSizeInBytes returns the DiskSizeInBytes field if non-nil, zero value otherwise.

### GetDiskSizeInBytesOk

`func (o *VirtualDiskInformation) GetDiskSizeInBytesOk() (*int64, bool)`

GetDiskSizeInBytesOk returns a tuple with the DiskSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeInBytes

`func (o *VirtualDiskInformation) SetDiskSizeInBytes(v int64)`

SetDiskSizeInBytes sets DiskSizeInBytes field to given value.

### HasDiskSizeInBytes

`func (o *VirtualDiskInformation) HasDiskSizeInBytes() bool`

HasDiskSizeInBytes returns a boolean if a field has been set.

### SetDiskSizeInBytesNil

`func (o *VirtualDiskInformation) SetDiskSizeInBytesNil(b bool)`

 SetDiskSizeInBytesNil sets the value for DiskSizeInBytes to be an explicit nil

### UnsetDiskSizeInBytes
`func (o *VirtualDiskInformation) UnsetDiskSizeInBytes()`

UnsetDiskSizeInBytes ensures that no value is present for DiskSizeInBytes, not even an explicit nil
### GetFilePath

`func (o *VirtualDiskInformation) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *VirtualDiskInformation) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *VirtualDiskInformation) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *VirtualDiskInformation) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### SetFilePathNil

`func (o *VirtualDiskInformation) SetFilePathNil(b bool)`

 SetFilePathNil sets the value for FilePath to be an explicit nil

### UnsetFilePath
`func (o *VirtualDiskInformation) UnsetFilePath()`

UnsetFilePath ensures that no value is present for FilePath, not even an explicit nil
### GetMountPoints

`func (o *VirtualDiskInformation) GetMountPoints() []string`

GetMountPoints returns the MountPoints field if non-nil, zero value otherwise.

### GetMountPointsOk

`func (o *VirtualDiskInformation) GetMountPointsOk() (*[]string, bool)`

GetMountPointsOk returns a tuple with the MountPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoints

`func (o *VirtualDiskInformation) SetMountPoints(v []string)`

SetMountPoints sets MountPoints field to given value.

### HasMountPoints

`func (o *VirtualDiskInformation) HasMountPoints() bool`

HasMountPoints returns a boolean if a field has been set.

### SetMountPointsNil

`func (o *VirtualDiskInformation) SetMountPointsNil(b bool)`

 SetMountPointsNil sets the value for MountPoints to be an explicit nil

### UnsetMountPoints
`func (o *VirtualDiskInformation) UnsetMountPoints()`

UnsetMountPoints ensures that no value is present for MountPoints, not even an explicit nil
### GetUnitNumber

`func (o *VirtualDiskInformation) GetUnitNumber() int64`

GetUnitNumber returns the UnitNumber field if non-nil, zero value otherwise.

### GetUnitNumberOk

`func (o *VirtualDiskInformation) GetUnitNumberOk() (*int64, bool)`

GetUnitNumberOk returns a tuple with the UnitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitNumber

`func (o *VirtualDiskInformation) SetUnitNumber(v int64)`

SetUnitNumber sets UnitNumber field to given value.

### HasUnitNumber

`func (o *VirtualDiskInformation) HasUnitNumber() bool`

HasUnitNumber returns a boolean if a field has been set.

### SetUnitNumberNil

`func (o *VirtualDiskInformation) SetUnitNumberNil(b bool)`

 SetUnitNumberNil sets the value for UnitNumber to be an explicit nil

### UnsetUnitNumber
`func (o *VirtualDiskInformation) UnsetUnitNumber()`

UnsetUnitNumber ensures that no value is present for UnitNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


