# ObjectSnapshotVolumeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Volumes** | Pointer to [**[]VolumeInfo**](VolumeInfo.md) | Specifies a list of volume info. | [optional] 

## Methods

### NewObjectSnapshotVolumeInfo

`func NewObjectSnapshotVolumeInfo() *ObjectSnapshotVolumeInfo`

NewObjectSnapshotVolumeInfo instantiates a new ObjectSnapshotVolumeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectSnapshotVolumeInfoWithDefaults

`func NewObjectSnapshotVolumeInfoWithDefaults() *ObjectSnapshotVolumeInfo`

NewObjectSnapshotVolumeInfoWithDefaults instantiates a new ObjectSnapshotVolumeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumes

`func (o *ObjectSnapshotVolumeInfo) GetVolumes() []VolumeInfo`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *ObjectSnapshotVolumeInfo) GetVolumesOk() (*[]VolumeInfo, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *ObjectSnapshotVolumeInfo) SetVolumes(v []VolumeInfo)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *ObjectSnapshotVolumeInfo) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### SetVolumesNil

`func (o *ObjectSnapshotVolumeInfo) SetVolumesNil(b bool)`

 SetVolumesNil sets the value for Volumes to be an explicit nil

### UnsetVolumes
`func (o *ObjectSnapshotVolumeInfo) UnsetVolumes()`

UnsetVolumes ensures that no value is present for Volumes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


