# MountedVolumeMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginalVolume** | Pointer to **NullableString** | Specifies the name of the original volume. | [optional] 
**MountedVolume** | Pointer to **NullableString** | Specifies the name of the point where the volume is mounted. | [optional] 
**FileSystemType** | Pointer to **NullableString** | Specifies the type of the file system of the volume. | [optional] 

## Methods

### NewMountedVolumeMapping

`func NewMountedVolumeMapping() *MountedVolumeMapping`

NewMountedVolumeMapping instantiates a new MountedVolumeMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMountedVolumeMappingWithDefaults

`func NewMountedVolumeMappingWithDefaults() *MountedVolumeMapping`

NewMountedVolumeMappingWithDefaults instantiates a new MountedVolumeMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginalVolume

`func (o *MountedVolumeMapping) GetOriginalVolume() string`

GetOriginalVolume returns the OriginalVolume field if non-nil, zero value otherwise.

### GetOriginalVolumeOk

`func (o *MountedVolumeMapping) GetOriginalVolumeOk() (*string, bool)`

GetOriginalVolumeOk returns a tuple with the OriginalVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalVolume

`func (o *MountedVolumeMapping) SetOriginalVolume(v string)`

SetOriginalVolume sets OriginalVolume field to given value.

### HasOriginalVolume

`func (o *MountedVolumeMapping) HasOriginalVolume() bool`

HasOriginalVolume returns a boolean if a field has been set.

### SetOriginalVolumeNil

`func (o *MountedVolumeMapping) SetOriginalVolumeNil(b bool)`

 SetOriginalVolumeNil sets the value for OriginalVolume to be an explicit nil

### UnsetOriginalVolume
`func (o *MountedVolumeMapping) UnsetOriginalVolume()`

UnsetOriginalVolume ensures that no value is present for OriginalVolume, not even an explicit nil
### GetMountedVolume

`func (o *MountedVolumeMapping) GetMountedVolume() string`

GetMountedVolume returns the MountedVolume field if non-nil, zero value otherwise.

### GetMountedVolumeOk

`func (o *MountedVolumeMapping) GetMountedVolumeOk() (*string, bool)`

GetMountedVolumeOk returns a tuple with the MountedVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountedVolume

`func (o *MountedVolumeMapping) SetMountedVolume(v string)`

SetMountedVolume sets MountedVolume field to given value.

### HasMountedVolume

`func (o *MountedVolumeMapping) HasMountedVolume() bool`

HasMountedVolume returns a boolean if a field has been set.

### SetMountedVolumeNil

`func (o *MountedVolumeMapping) SetMountedVolumeNil(b bool)`

 SetMountedVolumeNil sets the value for MountedVolume to be an explicit nil

### UnsetMountedVolume
`func (o *MountedVolumeMapping) UnsetMountedVolume()`

UnsetMountedVolume ensures that no value is present for MountedVolume, not even an explicit nil
### GetFileSystemType

`func (o *MountedVolumeMapping) GetFileSystemType() string`

GetFileSystemType returns the FileSystemType field if non-nil, zero value otherwise.

### GetFileSystemTypeOk

`func (o *MountedVolumeMapping) GetFileSystemTypeOk() (*string, bool)`

GetFileSystemTypeOk returns a tuple with the FileSystemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSystemType

`func (o *MountedVolumeMapping) SetFileSystemType(v string)`

SetFileSystemType sets FileSystemType field to given value.

### HasFileSystemType

`func (o *MountedVolumeMapping) HasFileSystemType() bool`

HasFileSystemType returns a boolean if a field has been set.

### SetFileSystemTypeNil

`func (o *MountedVolumeMapping) SetFileSystemTypeNil(b bool)`

 SetFileSystemTypeNil sets the value for FileSystemType to be an explicit nil

### UnsetFileSystemType
`func (o *MountedVolumeMapping) UnsetFileSystemType()`

UnsetFileSystemType ensures that no value is present for FileSystemType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


