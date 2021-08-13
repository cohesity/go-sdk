# PhysicalTargetParamsForRecoverVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MountTarget** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the target entity where the volumes are being mounted. | 
**VolumeMapping** | [**[]RecoverVolumeMapping**](RecoverVolumeMapping.md) | Specifies the mapping from source volumes to destination volumes. | 
**ForceUnmountVolume** | Pointer to **NullableBool** | Specifies whether volume would be dismounted first during LockVolume failure. If not specified, default is false. | [optional] 
**VlanConfig** | Pointer to [**NullableRecoveryVlanConfig**](RecoveryVlanConfig.md) | Specifies VLAN Params associated with the recovered. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client&#39;s (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery. | [optional] 

## Methods

### NewPhysicalTargetParamsForRecoverVolume

`func NewPhysicalTargetParamsForRecoverVolume(mountTarget NullableRecoveryObjectIdentifier, volumeMapping []RecoverVolumeMapping, ) *PhysicalTargetParamsForRecoverVolume`

NewPhysicalTargetParamsForRecoverVolume instantiates a new PhysicalTargetParamsForRecoverVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalTargetParamsForRecoverVolumeWithDefaults

`func NewPhysicalTargetParamsForRecoverVolumeWithDefaults() *PhysicalTargetParamsForRecoverVolume`

NewPhysicalTargetParamsForRecoverVolumeWithDefaults instantiates a new PhysicalTargetParamsForRecoverVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMountTarget

`func (o *PhysicalTargetParamsForRecoverVolume) GetMountTarget() RecoveryObjectIdentifier`

GetMountTarget returns the MountTarget field if non-nil, zero value otherwise.

### GetMountTargetOk

`func (o *PhysicalTargetParamsForRecoverVolume) GetMountTargetOk() (*RecoveryObjectIdentifier, bool)`

GetMountTargetOk returns a tuple with the MountTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountTarget

`func (o *PhysicalTargetParamsForRecoverVolume) SetMountTarget(v RecoveryObjectIdentifier)`

SetMountTarget sets MountTarget field to given value.


### SetMountTargetNil

`func (o *PhysicalTargetParamsForRecoverVolume) SetMountTargetNil(b bool)`

 SetMountTargetNil sets the value for MountTarget to be an explicit nil

### UnsetMountTarget
`func (o *PhysicalTargetParamsForRecoverVolume) UnsetMountTarget()`

UnsetMountTarget ensures that no value is present for MountTarget, not even an explicit nil
### GetVolumeMapping

`func (o *PhysicalTargetParamsForRecoverVolume) GetVolumeMapping() []RecoverVolumeMapping`

GetVolumeMapping returns the VolumeMapping field if non-nil, zero value otherwise.

### GetVolumeMappingOk

`func (o *PhysicalTargetParamsForRecoverVolume) GetVolumeMappingOk() (*[]RecoverVolumeMapping, bool)`

GetVolumeMappingOk returns a tuple with the VolumeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeMapping

`func (o *PhysicalTargetParamsForRecoverVolume) SetVolumeMapping(v []RecoverVolumeMapping)`

SetVolumeMapping sets VolumeMapping field to given value.


### SetVolumeMappingNil

`func (o *PhysicalTargetParamsForRecoverVolume) SetVolumeMappingNil(b bool)`

 SetVolumeMappingNil sets the value for VolumeMapping to be an explicit nil

### UnsetVolumeMapping
`func (o *PhysicalTargetParamsForRecoverVolume) UnsetVolumeMapping()`

UnsetVolumeMapping ensures that no value is present for VolumeMapping, not even an explicit nil
### GetForceUnmountVolume

`func (o *PhysicalTargetParamsForRecoverVolume) GetForceUnmountVolume() bool`

GetForceUnmountVolume returns the ForceUnmountVolume field if non-nil, zero value otherwise.

### GetForceUnmountVolumeOk

`func (o *PhysicalTargetParamsForRecoverVolume) GetForceUnmountVolumeOk() (*bool, bool)`

GetForceUnmountVolumeOk returns a tuple with the ForceUnmountVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceUnmountVolume

`func (o *PhysicalTargetParamsForRecoverVolume) SetForceUnmountVolume(v bool)`

SetForceUnmountVolume sets ForceUnmountVolume field to given value.

### HasForceUnmountVolume

`func (o *PhysicalTargetParamsForRecoverVolume) HasForceUnmountVolume() bool`

HasForceUnmountVolume returns a boolean if a field has been set.

### SetForceUnmountVolumeNil

`func (o *PhysicalTargetParamsForRecoverVolume) SetForceUnmountVolumeNil(b bool)`

 SetForceUnmountVolumeNil sets the value for ForceUnmountVolume to be an explicit nil

### UnsetForceUnmountVolume
`func (o *PhysicalTargetParamsForRecoverVolume) UnsetForceUnmountVolume()`

UnsetForceUnmountVolume ensures that no value is present for ForceUnmountVolume, not even an explicit nil
### GetVlanConfig

`func (o *PhysicalTargetParamsForRecoverVolume) GetVlanConfig() RecoveryVlanConfig`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *PhysicalTargetParamsForRecoverVolume) GetVlanConfigOk() (*RecoveryVlanConfig, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *PhysicalTargetParamsForRecoverVolume) SetVlanConfig(v RecoveryVlanConfig)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *PhysicalTargetParamsForRecoverVolume) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.

### SetVlanConfigNil

`func (o *PhysicalTargetParamsForRecoverVolume) SetVlanConfigNil(b bool)`

 SetVlanConfigNil sets the value for VlanConfig to be an explicit nil

### UnsetVlanConfig
`func (o *PhysicalTargetParamsForRecoverVolume) UnsetVlanConfig()`

UnsetVlanConfig ensures that no value is present for VlanConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


