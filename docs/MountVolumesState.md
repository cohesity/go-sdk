# MountVolumesState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BringDisksOnline** | Pointer to **NullableBool** | Optional setting that determines if the volumes are brought online on the mount target after attaching the disks. This option is only significant for VMs. | [optional] 
**MountVolumeResults** | Pointer to [**[]MountVolumeResultDetails**](MountVolumeResultDetails.md) | Array of Mount Volume Results.  Specifies the results of mounting each specified volume. | [optional] 
**OtherError** | Pointer to [**NullableRequestError**](RequestError.md) | Specifies an error that did not occur during the mount operation. | [optional] 
**TargetSourceId** | Pointer to **NullableInt64** | Specifies the target Protection Source Id where the volumes will be mounted. NOTE: The source that was backed up and the mount target must be the same type, for example if the source is a VMware VM, then the mount target must also be a VMware VM. The mount target must be registered on the Cohesity Cluster. | [optional] 
**Username** | Pointer to **NullableString** | Specifies the username to access the mount target. | [optional] 

## Methods

### NewMountVolumesState

`func NewMountVolumesState() *MountVolumesState`

NewMountVolumesState instantiates a new MountVolumesState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMountVolumesStateWithDefaults

`func NewMountVolumesStateWithDefaults() *MountVolumesState`

NewMountVolumesStateWithDefaults instantiates a new MountVolumesState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBringDisksOnline

`func (o *MountVolumesState) GetBringDisksOnline() bool`

GetBringDisksOnline returns the BringDisksOnline field if non-nil, zero value otherwise.

### GetBringDisksOnlineOk

`func (o *MountVolumesState) GetBringDisksOnlineOk() (*bool, bool)`

GetBringDisksOnlineOk returns a tuple with the BringDisksOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBringDisksOnline

`func (o *MountVolumesState) SetBringDisksOnline(v bool)`

SetBringDisksOnline sets BringDisksOnline field to given value.

### HasBringDisksOnline

`func (o *MountVolumesState) HasBringDisksOnline() bool`

HasBringDisksOnline returns a boolean if a field has been set.

### SetBringDisksOnlineNil

`func (o *MountVolumesState) SetBringDisksOnlineNil(b bool)`

 SetBringDisksOnlineNil sets the value for BringDisksOnline to be an explicit nil

### UnsetBringDisksOnline
`func (o *MountVolumesState) UnsetBringDisksOnline()`

UnsetBringDisksOnline ensures that no value is present for BringDisksOnline, not even an explicit nil
### GetMountVolumeResults

`func (o *MountVolumesState) GetMountVolumeResults() []MountVolumeResultDetails`

GetMountVolumeResults returns the MountVolumeResults field if non-nil, zero value otherwise.

### GetMountVolumeResultsOk

`func (o *MountVolumesState) GetMountVolumeResultsOk() (*[]MountVolumeResultDetails, bool)`

GetMountVolumeResultsOk returns a tuple with the MountVolumeResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountVolumeResults

`func (o *MountVolumesState) SetMountVolumeResults(v []MountVolumeResultDetails)`

SetMountVolumeResults sets MountVolumeResults field to given value.

### HasMountVolumeResults

`func (o *MountVolumesState) HasMountVolumeResults() bool`

HasMountVolumeResults returns a boolean if a field has been set.

### SetMountVolumeResultsNil

`func (o *MountVolumesState) SetMountVolumeResultsNil(b bool)`

 SetMountVolumeResultsNil sets the value for MountVolumeResults to be an explicit nil

### UnsetMountVolumeResults
`func (o *MountVolumesState) UnsetMountVolumeResults()`

UnsetMountVolumeResults ensures that no value is present for MountVolumeResults, not even an explicit nil
### GetOtherError

`func (o *MountVolumesState) GetOtherError() RequestError`

GetOtherError returns the OtherError field if non-nil, zero value otherwise.

### GetOtherErrorOk

`func (o *MountVolumesState) GetOtherErrorOk() (*RequestError, bool)`

GetOtherErrorOk returns a tuple with the OtherError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherError

`func (o *MountVolumesState) SetOtherError(v RequestError)`

SetOtherError sets OtherError field to given value.

### HasOtherError

`func (o *MountVolumesState) HasOtherError() bool`

HasOtherError returns a boolean if a field has been set.

### SetOtherErrorNil

`func (o *MountVolumesState) SetOtherErrorNil(b bool)`

 SetOtherErrorNil sets the value for OtherError to be an explicit nil

### UnsetOtherError
`func (o *MountVolumesState) UnsetOtherError()`

UnsetOtherError ensures that no value is present for OtherError, not even an explicit nil
### GetTargetSourceId

`func (o *MountVolumesState) GetTargetSourceId() int64`

GetTargetSourceId returns the TargetSourceId field if non-nil, zero value otherwise.

### GetTargetSourceIdOk

`func (o *MountVolumesState) GetTargetSourceIdOk() (*int64, bool)`

GetTargetSourceIdOk returns a tuple with the TargetSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSourceId

`func (o *MountVolumesState) SetTargetSourceId(v int64)`

SetTargetSourceId sets TargetSourceId field to given value.

### HasTargetSourceId

`func (o *MountVolumesState) HasTargetSourceId() bool`

HasTargetSourceId returns a boolean if a field has been set.

### SetTargetSourceIdNil

`func (o *MountVolumesState) SetTargetSourceIdNil(b bool)`

 SetTargetSourceIdNil sets the value for TargetSourceId to be an explicit nil

### UnsetTargetSourceId
`func (o *MountVolumesState) UnsetTargetSourceId()`

UnsetTargetSourceId ensures that no value is present for TargetSourceId, not even an explicit nil
### GetUsername

`func (o *MountVolumesState) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *MountVolumesState) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *MountVolumesState) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *MountVolumesState) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *MountVolumesState) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *MountVolumesState) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


