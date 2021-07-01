# MountVolumeResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**ErrorProto**](ErrorProto.md) |  | [optional] 
**FilesystemType** | Pointer to **NullableString** | Filesystem on this volume. | [optional] 
**MountPoint** | Pointer to **NullableString** | This is populated with the mount point where the volume corresponding to the newly attached volume is mounted. NOTE: This may not be present in the VM environments if onlining of disks is not requested or if the there was any issue during onlining. | [optional] 
**OriginalVolumeName** | Pointer to **NullableString** | This is the name or mount point of the original volume. | [optional] 

## Methods

### NewMountVolumeResult

`func NewMountVolumeResult() *MountVolumeResult`

NewMountVolumeResult instantiates a new MountVolumeResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMountVolumeResultWithDefaults

`func NewMountVolumeResultWithDefaults() *MountVolumeResult`

NewMountVolumeResultWithDefaults instantiates a new MountVolumeResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *MountVolumeResult) GetError() ErrorProto`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MountVolumeResult) GetErrorOk() (*ErrorProto, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MountVolumeResult) SetError(v ErrorProto)`

SetError sets Error field to given value.

### HasError

`func (o *MountVolumeResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetFilesystemType

`func (o *MountVolumeResult) GetFilesystemType() string`

GetFilesystemType returns the FilesystemType field if non-nil, zero value otherwise.

### GetFilesystemTypeOk

`func (o *MountVolumeResult) GetFilesystemTypeOk() (*string, bool)`

GetFilesystemTypeOk returns a tuple with the FilesystemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystemType

`func (o *MountVolumeResult) SetFilesystemType(v string)`

SetFilesystemType sets FilesystemType field to given value.

### HasFilesystemType

`func (o *MountVolumeResult) HasFilesystemType() bool`

HasFilesystemType returns a boolean if a field has been set.

### SetFilesystemTypeNil

`func (o *MountVolumeResult) SetFilesystemTypeNil(b bool)`

 SetFilesystemTypeNil sets the value for FilesystemType to be an explicit nil

### UnsetFilesystemType
`func (o *MountVolumeResult) UnsetFilesystemType()`

UnsetFilesystemType ensures that no value is present for FilesystemType, not even an explicit nil
### GetMountPoint

`func (o *MountVolumeResult) GetMountPoint() string`

GetMountPoint returns the MountPoint field if non-nil, zero value otherwise.

### GetMountPointOk

`func (o *MountVolumeResult) GetMountPointOk() (*string, bool)`

GetMountPointOk returns a tuple with the MountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoint

`func (o *MountVolumeResult) SetMountPoint(v string)`

SetMountPoint sets MountPoint field to given value.

### HasMountPoint

`func (o *MountVolumeResult) HasMountPoint() bool`

HasMountPoint returns a boolean if a field has been set.

### SetMountPointNil

`func (o *MountVolumeResult) SetMountPointNil(b bool)`

 SetMountPointNil sets the value for MountPoint to be an explicit nil

### UnsetMountPoint
`func (o *MountVolumeResult) UnsetMountPoint()`

UnsetMountPoint ensures that no value is present for MountPoint, not even an explicit nil
### GetOriginalVolumeName

`func (o *MountVolumeResult) GetOriginalVolumeName() string`

GetOriginalVolumeName returns the OriginalVolumeName field if non-nil, zero value otherwise.

### GetOriginalVolumeNameOk

`func (o *MountVolumeResult) GetOriginalVolumeNameOk() (*string, bool)`

GetOriginalVolumeNameOk returns a tuple with the OriginalVolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalVolumeName

`func (o *MountVolumeResult) SetOriginalVolumeName(v string)`

SetOriginalVolumeName sets OriginalVolumeName field to given value.

### HasOriginalVolumeName

`func (o *MountVolumeResult) HasOriginalVolumeName() bool`

HasOriginalVolumeName returns a boolean if a field has been set.

### SetOriginalVolumeNameNil

`func (o *MountVolumeResult) SetOriginalVolumeNameNil(b bool)`

 SetOriginalVolumeNameNil sets the value for OriginalVolumeName to be an explicit nil

### UnsetOriginalVolumeName
`func (o *MountVolumeResult) UnsetOriginalVolumeName()`

UnsetOriginalVolumeName ensures that no value is present for OriginalVolumeName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


