# MountVolumesParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BringDisksOnline** | Pointer to **NullableBool** | Optional setting that determines if the volumes are brought online on the mount target after attaching the disks. This field is only set for VMs. The Cohesity Cluster always attempts to mount Physical servers. If true and the mount target is a VM, to mount the volumes VMware Tools must be installed on the guest operating system of the VM and login credentials to the mount target must be specified. NOTE: If automount is configured for a Windows system, the volumes may be automatically brought online. | [optional] 
**Password** | Pointer to **NullableString** | Specifies password of the username to access the target source. | [optional] 
**TargetSourceId** | Pointer to **NullableInt64** | Specifies the target Protection Source id where the volumes will be mounted. NOTE: The source that was backed up and the mount target must be the same type, for example if the source is a VMware VM, then the mount target must also be a VMware VM. The mount target must be registered on the Cohesity Cluster. | [optional] 
**UseExistingAgent** | Pointer to **NullableBool** | Optional setting that determines if this will use an existing agent on the target vm to bring disks online. | [optional] 
**Username** | Pointer to **NullableString** | Specifies username to access the target source. | [optional] 
**VolumeNames** | Pointer to **[]string** | Array of Volume Names.  Optionally specify the names of volumes to mount. If none are specified, all volumes of the Server are mounted on the target. To get the names of the volumes, call the GET /public/restore/vms/volumesInformation operation. | [optional] 

## Methods

### NewMountVolumesParameters

`func NewMountVolumesParameters() *MountVolumesParameters`

NewMountVolumesParameters instantiates a new MountVolumesParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMountVolumesParametersWithDefaults

`func NewMountVolumesParametersWithDefaults() *MountVolumesParameters`

NewMountVolumesParametersWithDefaults instantiates a new MountVolumesParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBringDisksOnline

`func (o *MountVolumesParameters) GetBringDisksOnline() bool`

GetBringDisksOnline returns the BringDisksOnline field if non-nil, zero value otherwise.

### GetBringDisksOnlineOk

`func (o *MountVolumesParameters) GetBringDisksOnlineOk() (*bool, bool)`

GetBringDisksOnlineOk returns a tuple with the BringDisksOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBringDisksOnline

`func (o *MountVolumesParameters) SetBringDisksOnline(v bool)`

SetBringDisksOnline sets BringDisksOnline field to given value.

### HasBringDisksOnline

`func (o *MountVolumesParameters) HasBringDisksOnline() bool`

HasBringDisksOnline returns a boolean if a field has been set.

### SetBringDisksOnlineNil

`func (o *MountVolumesParameters) SetBringDisksOnlineNil(b bool)`

 SetBringDisksOnlineNil sets the value for BringDisksOnline to be an explicit nil

### UnsetBringDisksOnline
`func (o *MountVolumesParameters) UnsetBringDisksOnline()`

UnsetBringDisksOnline ensures that no value is present for BringDisksOnline, not even an explicit nil
### GetPassword

`func (o *MountVolumesParameters) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MountVolumesParameters) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MountVolumesParameters) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *MountVolumesParameters) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *MountVolumesParameters) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *MountVolumesParameters) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetTargetSourceId

`func (o *MountVolumesParameters) GetTargetSourceId() int64`

GetTargetSourceId returns the TargetSourceId field if non-nil, zero value otherwise.

### GetTargetSourceIdOk

`func (o *MountVolumesParameters) GetTargetSourceIdOk() (*int64, bool)`

GetTargetSourceIdOk returns a tuple with the TargetSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSourceId

`func (o *MountVolumesParameters) SetTargetSourceId(v int64)`

SetTargetSourceId sets TargetSourceId field to given value.

### HasTargetSourceId

`func (o *MountVolumesParameters) HasTargetSourceId() bool`

HasTargetSourceId returns a boolean if a field has been set.

### SetTargetSourceIdNil

`func (o *MountVolumesParameters) SetTargetSourceIdNil(b bool)`

 SetTargetSourceIdNil sets the value for TargetSourceId to be an explicit nil

### UnsetTargetSourceId
`func (o *MountVolumesParameters) UnsetTargetSourceId()`

UnsetTargetSourceId ensures that no value is present for TargetSourceId, not even an explicit nil
### GetUseExistingAgent

`func (o *MountVolumesParameters) GetUseExistingAgent() bool`

GetUseExistingAgent returns the UseExistingAgent field if non-nil, zero value otherwise.

### GetUseExistingAgentOk

`func (o *MountVolumesParameters) GetUseExistingAgentOk() (*bool, bool)`

GetUseExistingAgentOk returns a tuple with the UseExistingAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseExistingAgent

`func (o *MountVolumesParameters) SetUseExistingAgent(v bool)`

SetUseExistingAgent sets UseExistingAgent field to given value.

### HasUseExistingAgent

`func (o *MountVolumesParameters) HasUseExistingAgent() bool`

HasUseExistingAgent returns a boolean if a field has been set.

### SetUseExistingAgentNil

`func (o *MountVolumesParameters) SetUseExistingAgentNil(b bool)`

 SetUseExistingAgentNil sets the value for UseExistingAgent to be an explicit nil

### UnsetUseExistingAgent
`func (o *MountVolumesParameters) UnsetUseExistingAgent()`

UnsetUseExistingAgent ensures that no value is present for UseExistingAgent, not even an explicit nil
### GetUsername

`func (o *MountVolumesParameters) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *MountVolumesParameters) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *MountVolumesParameters) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *MountVolumesParameters) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *MountVolumesParameters) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *MountVolumesParameters) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetVolumeNames

`func (o *MountVolumesParameters) GetVolumeNames() []string`

GetVolumeNames returns the VolumeNames field if non-nil, zero value otherwise.

### GetVolumeNamesOk

`func (o *MountVolumesParameters) GetVolumeNamesOk() (*[]string, bool)`

GetVolumeNamesOk returns a tuple with the VolumeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeNames

`func (o *MountVolumesParameters) SetVolumeNames(v []string)`

SetVolumeNames sets VolumeNames field to given value.

### HasVolumeNames

`func (o *MountVolumesParameters) HasVolumeNames() bool`

HasVolumeNames returns a boolean if a field has been set.

### SetVolumeNamesNil

`func (o *MountVolumesParameters) SetVolumeNamesNil(b bool)`

 SetVolumeNamesNil sets the value for VolumeNames to be an explicit nil

### UnsetVolumeNames
`func (o *MountVolumesParameters) UnsetVolumeNames()`

UnsetVolumeNames ensures that no value is present for VolumeNames, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


