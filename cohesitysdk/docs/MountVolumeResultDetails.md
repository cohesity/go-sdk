# MountVolumeResultDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MountError** | Pointer to [**NullableRequestError**](RequestError.md) | Specifies the cause of the mount failure if the mounting of a volume failed. | [optional] 
**MountPoint** | Pointer to **NullableString** | Specifies the mount point where the volume is mounted. NOTE: This field may not be populated for VM environments if the onlining of disks is not requested or there was any issue during onlining. | [optional] 
**VolumeName** | Pointer to **NullableString** | Specifies the name of the original volume. | [optional] 

## Methods

### NewMountVolumeResultDetails

`func NewMountVolumeResultDetails() *MountVolumeResultDetails`

NewMountVolumeResultDetails instantiates a new MountVolumeResultDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMountVolumeResultDetailsWithDefaults

`func NewMountVolumeResultDetailsWithDefaults() *MountVolumeResultDetails`

NewMountVolumeResultDetailsWithDefaults instantiates a new MountVolumeResultDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMountError

`func (o *MountVolumeResultDetails) GetMountError() RequestError`

GetMountError returns the MountError field if non-nil, zero value otherwise.

### GetMountErrorOk

`func (o *MountVolumeResultDetails) GetMountErrorOk() (*RequestError, bool)`

GetMountErrorOk returns a tuple with the MountError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountError

`func (o *MountVolumeResultDetails) SetMountError(v RequestError)`

SetMountError sets MountError field to given value.

### HasMountError

`func (o *MountVolumeResultDetails) HasMountError() bool`

HasMountError returns a boolean if a field has been set.

### SetMountErrorNil

`func (o *MountVolumeResultDetails) SetMountErrorNil(b bool)`

 SetMountErrorNil sets the value for MountError to be an explicit nil

### UnsetMountError
`func (o *MountVolumeResultDetails) UnsetMountError()`

UnsetMountError ensures that no value is present for MountError, not even an explicit nil
### GetMountPoint

`func (o *MountVolumeResultDetails) GetMountPoint() string`

GetMountPoint returns the MountPoint field if non-nil, zero value otherwise.

### GetMountPointOk

`func (o *MountVolumeResultDetails) GetMountPointOk() (*string, bool)`

GetMountPointOk returns a tuple with the MountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoint

`func (o *MountVolumeResultDetails) SetMountPoint(v string)`

SetMountPoint sets MountPoint field to given value.

### HasMountPoint

`func (o *MountVolumeResultDetails) HasMountPoint() bool`

HasMountPoint returns a boolean if a field has been set.

### SetMountPointNil

`func (o *MountVolumeResultDetails) SetMountPointNil(b bool)`

 SetMountPointNil sets the value for MountPoint to be an explicit nil

### UnsetMountPoint
`func (o *MountVolumeResultDetails) UnsetMountPoint()`

UnsetMountPoint ensures that no value is present for MountPoint, not even an explicit nil
### GetVolumeName

`func (o *MountVolumeResultDetails) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *MountVolumeResultDetails) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *MountVolumeResultDetails) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *MountVolumeResultDetails) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.

### SetVolumeNameNil

`func (o *MountVolumeResultDetails) SetVolumeNameNil(b bool)`

 SetVolumeNameNil sets the value for VolumeName to be an explicit nil

### UnsetVolumeName
`func (o *MountVolumeResultDetails) UnsetVolumeName()`

UnsetVolumeName ensures that no value is present for VolumeName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


