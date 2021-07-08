# VmVolumesInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilesystemVolumes** | Pointer to [**[]FilesystemVolume**](FilesystemVolume.md) | Array of Filesystem Volumes.  Specifies information about the filesystem volumes found in a logical volume. | [optional] 

## Methods

### NewVmVolumesInformation

`func NewVmVolumesInformation() *VmVolumesInformation`

NewVmVolumesInformation instantiates a new VmVolumesInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmVolumesInformationWithDefaults

`func NewVmVolumesInformationWithDefaults() *VmVolumesInformation`

NewVmVolumesInformationWithDefaults instantiates a new VmVolumesInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilesystemVolumes

`func (o *VmVolumesInformation) GetFilesystemVolumes() []FilesystemVolume`

GetFilesystemVolumes returns the FilesystemVolumes field if non-nil, zero value otherwise.

### GetFilesystemVolumesOk

`func (o *VmVolumesInformation) GetFilesystemVolumesOk() (*[]FilesystemVolume, bool)`

GetFilesystemVolumesOk returns a tuple with the FilesystemVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystemVolumes

`func (o *VmVolumesInformation) SetFilesystemVolumes(v []FilesystemVolume)`

SetFilesystemVolumes sets FilesystemVolumes field to given value.

### HasFilesystemVolumes

`func (o *VmVolumesInformation) HasFilesystemVolumes() bool`

HasFilesystemVolumes returns a boolean if a field has been set.

### SetFilesystemVolumesNil

`func (o *VmVolumesInformation) SetFilesystemVolumesNil(b bool)`

 SetFilesystemVolumesNil sets the value for FilesystemVolumes to be an explicit nil

### UnsetFilesystemVolumes
`func (o *VmVolumesInformation) UnsetFilesystemVolumes()`

UnsetFilesystemVolumes ensures that no value is present for FilesystemVolumes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


