# PhysicalVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DevicePath** | Pointer to **NullableString** | Specifies the path to the device that hosts the volume locally. | [optional] 
**Guid** | Pointer to **NullableString** | Specifies an id for the Physical Volume. | [optional] 
**IsBootVolume** | Pointer to **NullableBool** | Specifies whether the volume is boot volume. | [optional] 
**IsExtendedAttributesSupported** | Pointer to **NullableBool** | Specifies whether this volume supports extended attributes (like ACLs) when performing file backups. | [optional] 
**IsProtected** | Pointer to **NullableBool** | Specifies if a volume is protected by a Job. | [optional] 
**IsSharedVolume** | Pointer to **NullableBool** | Specifies whether the volume is shared volume. | [optional] 
**Label** | Pointer to **NullableString** | Specifies a volume label that can be used for displaying additional identifying information about a volume. | [optional] 
**LogicalSizeBytes** | Pointer to **NullableInt32** | Specifies the logical size of the volume in bytes that is not reduced by change-block tracking, compression and deduplication. | [optional] 
**MountPoints** | Pointer to **[]string** | Array of Mount Points.  Specifies the mount points where the volume is mounted, for example: &#39;C:\\&#39;, &#39;/mnt/foo&#39; etc. | [optional] 
**MountType** | Pointer to **NullableString** | Specifies mount type of volume e.g. nfs, autofs, ext4 etc. | [optional] 
**NetworkPath** | Pointer to **NullableString** | Specifies the full path to connect to the network attached volume. For example, (IP or hostname):/path/to/share for NFS volumes). | [optional] 
**UsedSizeBytes** | Pointer to **NullableInt32** | Specifies the size used by the volume in bytes. | [optional] 

## Methods

### NewPhysicalVolume

`func NewPhysicalVolume() *PhysicalVolume`

NewPhysicalVolume instantiates a new PhysicalVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalVolumeWithDefaults

`func NewPhysicalVolumeWithDefaults() *PhysicalVolume`

NewPhysicalVolumeWithDefaults instantiates a new PhysicalVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevicePath

`func (o *PhysicalVolume) GetDevicePath() string`

GetDevicePath returns the DevicePath field if non-nil, zero value otherwise.

### GetDevicePathOk

`func (o *PhysicalVolume) GetDevicePathOk() (*string, bool)`

GetDevicePathOk returns a tuple with the DevicePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePath

`func (o *PhysicalVolume) SetDevicePath(v string)`

SetDevicePath sets DevicePath field to given value.

### HasDevicePath

`func (o *PhysicalVolume) HasDevicePath() bool`

HasDevicePath returns a boolean if a field has been set.

### SetDevicePathNil

`func (o *PhysicalVolume) SetDevicePathNil(b bool)`

 SetDevicePathNil sets the value for DevicePath to be an explicit nil

### UnsetDevicePath
`func (o *PhysicalVolume) UnsetDevicePath()`

UnsetDevicePath ensures that no value is present for DevicePath, not even an explicit nil
### GetGuid

`func (o *PhysicalVolume) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *PhysicalVolume) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *PhysicalVolume) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *PhysicalVolume) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### SetGuidNil

`func (o *PhysicalVolume) SetGuidNil(b bool)`

 SetGuidNil sets the value for Guid to be an explicit nil

### UnsetGuid
`func (o *PhysicalVolume) UnsetGuid()`

UnsetGuid ensures that no value is present for Guid, not even an explicit nil
### GetIsBootVolume

`func (o *PhysicalVolume) GetIsBootVolume() bool`

GetIsBootVolume returns the IsBootVolume field if non-nil, zero value otherwise.

### GetIsBootVolumeOk

`func (o *PhysicalVolume) GetIsBootVolumeOk() (*bool, bool)`

GetIsBootVolumeOk returns a tuple with the IsBootVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBootVolume

`func (o *PhysicalVolume) SetIsBootVolume(v bool)`

SetIsBootVolume sets IsBootVolume field to given value.

### HasIsBootVolume

`func (o *PhysicalVolume) HasIsBootVolume() bool`

HasIsBootVolume returns a boolean if a field has been set.

### SetIsBootVolumeNil

`func (o *PhysicalVolume) SetIsBootVolumeNil(b bool)`

 SetIsBootVolumeNil sets the value for IsBootVolume to be an explicit nil

### UnsetIsBootVolume
`func (o *PhysicalVolume) UnsetIsBootVolume()`

UnsetIsBootVolume ensures that no value is present for IsBootVolume, not even an explicit nil
### GetIsExtendedAttributesSupported

`func (o *PhysicalVolume) GetIsExtendedAttributesSupported() bool`

GetIsExtendedAttributesSupported returns the IsExtendedAttributesSupported field if non-nil, zero value otherwise.

### GetIsExtendedAttributesSupportedOk

`func (o *PhysicalVolume) GetIsExtendedAttributesSupportedOk() (*bool, bool)`

GetIsExtendedAttributesSupportedOk returns a tuple with the IsExtendedAttributesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExtendedAttributesSupported

`func (o *PhysicalVolume) SetIsExtendedAttributesSupported(v bool)`

SetIsExtendedAttributesSupported sets IsExtendedAttributesSupported field to given value.

### HasIsExtendedAttributesSupported

`func (o *PhysicalVolume) HasIsExtendedAttributesSupported() bool`

HasIsExtendedAttributesSupported returns a boolean if a field has been set.

### SetIsExtendedAttributesSupportedNil

`func (o *PhysicalVolume) SetIsExtendedAttributesSupportedNil(b bool)`

 SetIsExtendedAttributesSupportedNil sets the value for IsExtendedAttributesSupported to be an explicit nil

### UnsetIsExtendedAttributesSupported
`func (o *PhysicalVolume) UnsetIsExtendedAttributesSupported()`

UnsetIsExtendedAttributesSupported ensures that no value is present for IsExtendedAttributesSupported, not even an explicit nil
### GetIsProtected

`func (o *PhysicalVolume) GetIsProtected() bool`

GetIsProtected returns the IsProtected field if non-nil, zero value otherwise.

### GetIsProtectedOk

`func (o *PhysicalVolume) GetIsProtectedOk() (*bool, bool)`

GetIsProtectedOk returns a tuple with the IsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProtected

`func (o *PhysicalVolume) SetIsProtected(v bool)`

SetIsProtected sets IsProtected field to given value.

### HasIsProtected

`func (o *PhysicalVolume) HasIsProtected() bool`

HasIsProtected returns a boolean if a field has been set.

### SetIsProtectedNil

`func (o *PhysicalVolume) SetIsProtectedNil(b bool)`

 SetIsProtectedNil sets the value for IsProtected to be an explicit nil

### UnsetIsProtected
`func (o *PhysicalVolume) UnsetIsProtected()`

UnsetIsProtected ensures that no value is present for IsProtected, not even an explicit nil
### GetIsSharedVolume

`func (o *PhysicalVolume) GetIsSharedVolume() bool`

GetIsSharedVolume returns the IsSharedVolume field if non-nil, zero value otherwise.

### GetIsSharedVolumeOk

`func (o *PhysicalVolume) GetIsSharedVolumeOk() (*bool, bool)`

GetIsSharedVolumeOk returns a tuple with the IsSharedVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSharedVolume

`func (o *PhysicalVolume) SetIsSharedVolume(v bool)`

SetIsSharedVolume sets IsSharedVolume field to given value.

### HasIsSharedVolume

`func (o *PhysicalVolume) HasIsSharedVolume() bool`

HasIsSharedVolume returns a boolean if a field has been set.

### SetIsSharedVolumeNil

`func (o *PhysicalVolume) SetIsSharedVolumeNil(b bool)`

 SetIsSharedVolumeNil sets the value for IsSharedVolume to be an explicit nil

### UnsetIsSharedVolume
`func (o *PhysicalVolume) UnsetIsSharedVolume()`

UnsetIsSharedVolume ensures that no value is present for IsSharedVolume, not even an explicit nil
### GetLabel

`func (o *PhysicalVolume) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PhysicalVolume) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PhysicalVolume) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PhysicalVolume) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *PhysicalVolume) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *PhysicalVolume) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetLogicalSizeBytes

`func (o *PhysicalVolume) GetLogicalSizeBytes() int32`

GetLogicalSizeBytes returns the LogicalSizeBytes field if non-nil, zero value otherwise.

### GetLogicalSizeBytesOk

`func (o *PhysicalVolume) GetLogicalSizeBytesOk() (*int32, bool)`

GetLogicalSizeBytesOk returns a tuple with the LogicalSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalSizeBytes

`func (o *PhysicalVolume) SetLogicalSizeBytes(v int32)`

SetLogicalSizeBytes sets LogicalSizeBytes field to given value.

### HasLogicalSizeBytes

`func (o *PhysicalVolume) HasLogicalSizeBytes() bool`

HasLogicalSizeBytes returns a boolean if a field has been set.

### SetLogicalSizeBytesNil

`func (o *PhysicalVolume) SetLogicalSizeBytesNil(b bool)`

 SetLogicalSizeBytesNil sets the value for LogicalSizeBytes to be an explicit nil

### UnsetLogicalSizeBytes
`func (o *PhysicalVolume) UnsetLogicalSizeBytes()`

UnsetLogicalSizeBytes ensures that no value is present for LogicalSizeBytes, not even an explicit nil
### GetMountPoints

`func (o *PhysicalVolume) GetMountPoints() []string`

GetMountPoints returns the MountPoints field if non-nil, zero value otherwise.

### GetMountPointsOk

`func (o *PhysicalVolume) GetMountPointsOk() (*[]string, bool)`

GetMountPointsOk returns a tuple with the MountPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoints

`func (o *PhysicalVolume) SetMountPoints(v []string)`

SetMountPoints sets MountPoints field to given value.

### HasMountPoints

`func (o *PhysicalVolume) HasMountPoints() bool`

HasMountPoints returns a boolean if a field has been set.

### SetMountPointsNil

`func (o *PhysicalVolume) SetMountPointsNil(b bool)`

 SetMountPointsNil sets the value for MountPoints to be an explicit nil

### UnsetMountPoints
`func (o *PhysicalVolume) UnsetMountPoints()`

UnsetMountPoints ensures that no value is present for MountPoints, not even an explicit nil
### GetMountType

`func (o *PhysicalVolume) GetMountType() string`

GetMountType returns the MountType field if non-nil, zero value otherwise.

### GetMountTypeOk

`func (o *PhysicalVolume) GetMountTypeOk() (*string, bool)`

GetMountTypeOk returns a tuple with the MountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountType

`func (o *PhysicalVolume) SetMountType(v string)`

SetMountType sets MountType field to given value.

### HasMountType

`func (o *PhysicalVolume) HasMountType() bool`

HasMountType returns a boolean if a field has been set.

### SetMountTypeNil

`func (o *PhysicalVolume) SetMountTypeNil(b bool)`

 SetMountTypeNil sets the value for MountType to be an explicit nil

### UnsetMountType
`func (o *PhysicalVolume) UnsetMountType()`

UnsetMountType ensures that no value is present for MountType, not even an explicit nil
### GetNetworkPath

`func (o *PhysicalVolume) GetNetworkPath() string`

GetNetworkPath returns the NetworkPath field if non-nil, zero value otherwise.

### GetNetworkPathOk

`func (o *PhysicalVolume) GetNetworkPathOk() (*string, bool)`

GetNetworkPathOk returns a tuple with the NetworkPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPath

`func (o *PhysicalVolume) SetNetworkPath(v string)`

SetNetworkPath sets NetworkPath field to given value.

### HasNetworkPath

`func (o *PhysicalVolume) HasNetworkPath() bool`

HasNetworkPath returns a boolean if a field has been set.

### SetNetworkPathNil

`func (o *PhysicalVolume) SetNetworkPathNil(b bool)`

 SetNetworkPathNil sets the value for NetworkPath to be an explicit nil

### UnsetNetworkPath
`func (o *PhysicalVolume) UnsetNetworkPath()`

UnsetNetworkPath ensures that no value is present for NetworkPath, not even an explicit nil
### GetUsedSizeBytes

`func (o *PhysicalVolume) GetUsedSizeBytes() int32`

GetUsedSizeBytes returns the UsedSizeBytes field if non-nil, zero value otherwise.

### GetUsedSizeBytesOk

`func (o *PhysicalVolume) GetUsedSizeBytesOk() (*int32, bool)`

GetUsedSizeBytesOk returns a tuple with the UsedSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSizeBytes

`func (o *PhysicalVolume) SetUsedSizeBytes(v int32)`

SetUsedSizeBytes sets UsedSizeBytes field to given value.

### HasUsedSizeBytes

`func (o *PhysicalVolume) HasUsedSizeBytes() bool`

HasUsedSizeBytes returns a boolean if a field has been set.

### SetUsedSizeBytesNil

`func (o *PhysicalVolume) SetUsedSizeBytesNil(b bool)`

 SetUsedSizeBytesNil sets the value for UsedSizeBytes to be an explicit nil

### UnsetUsedSizeBytes
`func (o *PhysicalVolume) UnsetUsedSizeBytes()`

UnsetUsedSizeBytes ensures that no value is present for UsedSizeBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


