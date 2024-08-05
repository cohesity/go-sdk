# RecoverPhysicalVolumeParamsPhysicalTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForceUnmountVolume** | Pointer to **NullableBool** | Specifies whether volume would be dismounted first during LockVolume failure. If not specified, default is false. | [optional] 
**MountTarget** | [**NullablePhysicalTargetParamsForRecoverVolumeMountTarget**](PhysicalTargetParamsForRecoverVolumeMountTarget.md) |  | 
**VlanConfig** | Pointer to [**NullableAcropolisTargetParamsForRecoverVmVlanConfig**](AcropolisTargetParamsForRecoverVmVlanConfig.md) |  | [optional] 
**VolumeMapping** | [**[]RecoverVolumeMapping**](RecoverVolumeMapping.md) | Specifies the mapping from source volumes to destination volumes. | 

## Methods

### NewRecoverPhysicalVolumeParamsPhysicalTargetParams

`func NewRecoverPhysicalVolumeParamsPhysicalTargetParams(mountTarget NullablePhysicalTargetParamsForRecoverVolumeMountTarget, volumeMapping []RecoverVolumeMapping, ) *RecoverPhysicalVolumeParamsPhysicalTargetParams`

NewRecoverPhysicalVolumeParamsPhysicalTargetParams instantiates a new RecoverPhysicalVolumeParamsPhysicalTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverPhysicalVolumeParamsPhysicalTargetParamsWithDefaults

`func NewRecoverPhysicalVolumeParamsPhysicalTargetParamsWithDefaults() *RecoverPhysicalVolumeParamsPhysicalTargetParams`

NewRecoverPhysicalVolumeParamsPhysicalTargetParamsWithDefaults instantiates a new RecoverPhysicalVolumeParamsPhysicalTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForceUnmountVolume

`func (o *RecoverPhysicalVolumeParamsPhysicalTargetParams) GetForceUnmountVolume() bool`

GetForceUnmountVolume returns the ForceUnmountVolume field if non-nil, zero value otherwise.

### GetForceUnmountVolumeOk

`func (o *RecoverPhysicalVolumeParamsPhysicalTargetParams) GetForceUnmountVolumeOk() (*bool, bool)`

GetForceUnmountVolumeOk returns a tuple with the ForceUnmountVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceUnmountVolume

`func (o *RecoverPhysicalVolumeParamsPhysicalTargetParams) SetForceUnmountVolume(v bool)`

SetForceUnmountVolume sets ForceUnmountVolume field to given value.

### HasForceUnmountVolume

`func (o *RecoverPhysicalVolumeParamsPhysicalTargetParams) HasForceUnmountVolume() bool`

HasForceUnmountVolume returns a boolean if a field has been set.

### SetForceUnmountVolumeNil

`func (o *RecoverPhysicalVolumeParamsPhysicalTargetParams) SetForceUnmountVolumeNil(b bool)`

 SetForceUnmountVolumeNil sets the value for ForceUnmountVolume to be an explicit nil

### UnsetForceUnmountVolume
`func (o *RecoverPhysicalVolumeParamsPhysicalTargetParams) UnsetForceUnmountVolume()`

UnsetForceUnmountVolume ensures that no value is present for ForceUnmountVolume, not even an explicit nil
### GetMountTarget

`func (o *RecoverPhysicalVolumeParamsPhysicalTargetParams) GetMountTarget() PhysicalTargetParamsForRecoverVolumeMountTarget`

GetMountTarget returns the MountTarget field if non-nil, zero value otherwise.

### GetMountTargetOk

`func (o *RecoverPhysicalVolumeParamsPhysicalTargetParams) GetMountTargetOk() (*PhysicalTargetParamsForRecoverVolumeMountTarget, bool)`

GetMountTargetOk returns a tuple with the MountTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountTarget

`func (o *RecoverPhysicalVolumeParamsPhysicalTargetParams) SetMountTarget(v PhysicalTargetParamsForRecoverVolumeMountTarget)`

SetMountTarget sets MountTarget field to given value.


### SetMountTargetNil

`func (o *RecoverPhysicalVolumeParamsPhysicalTargetParams) SetMountTargetNil(b bool)`

 SetMountTargetNil sets the value for MountTarget to be an explicit nil

### UnsetMountTarget
`func (o *RecoverPhysicalVolumeParamsPhysicalTargetParams) UnsetMountTarget()`

UnsetMountTarget ensures that no value is present for MountTarget, not even an explicit nil
### GetVlanConfig

`func (o *RecoverPhysicalVolumeParamsPhysicalTargetParams) GetVlanConfig() AcropolisTargetParamsForRecoverVmVlanConfig`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *RecoverPhysicalVolumeParamsPhysicalTargetParams) GetVlanConfigOk() (*AcropolisTargetParamsForRecoverVmVlanConfig, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *RecoverPhysicalVolumeParamsPhysicalTargetParams) SetVlanConfig(v AcropolisTargetParamsForRecoverVmVlanConfig)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *RecoverPhysicalVolumeParamsPhysicalTargetParams) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.

### SetVlanConfigNil

`func (o *RecoverPhysicalVolumeParamsPhysicalTargetParams) SetVlanConfigNil(b bool)`

 SetVlanConfigNil sets the value for VlanConfig to be an explicit nil

### UnsetVlanConfig
`func (o *RecoverPhysicalVolumeParamsPhysicalTargetParams) UnsetVlanConfig()`

UnsetVlanConfig ensures that no value is present for VlanConfig, not even an explicit nil
### GetVolumeMapping

`func (o *RecoverPhysicalVolumeParamsPhysicalTargetParams) GetVolumeMapping() []RecoverVolumeMapping`

GetVolumeMapping returns the VolumeMapping field if non-nil, zero value otherwise.

### GetVolumeMappingOk

`func (o *RecoverPhysicalVolumeParamsPhysicalTargetParams) GetVolumeMappingOk() (*[]RecoverVolumeMapping, bool)`

GetVolumeMappingOk returns a tuple with the VolumeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeMapping

`func (o *RecoverPhysicalVolumeParamsPhysicalTargetParams) SetVolumeMapping(v []RecoverVolumeMapping)`

SetVolumeMapping sets VolumeMapping field to given value.


### SetVolumeMappingNil

`func (o *RecoverPhysicalVolumeParamsPhysicalTargetParams) SetVolumeMappingNil(b bool)`

 SetVolumeMappingNil sets the value for VolumeMapping to be an explicit nil

### UnsetVolumeMapping
`func (o *RecoverPhysicalVolumeParamsPhysicalTargetParams) UnsetVolumeMapping()`

UnsetVolumeMapping ensures that no value is present for VolumeMapping, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


