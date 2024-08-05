# PhysicalTargetParamsForMountVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MountToOriginalTarget** | **NullableBool** | Specifies whether to mount to the original target. If true, originalTargetConfig must be specified. If false, newTargetConfig must be specified. | 
**MountedVolumeMapping** | Pointer to [**[]MountedVolumeMapping**](MountedVolumeMapping.md) | Specifies the mapping of original volumes and mounted volumes | [optional] [readonly] 
**NewTargetConfig** | Pointer to [**NullablePhysicalTargetParamsForMountVolumeNewTargetConfig**](PhysicalTargetParamsForMountVolumeNewTargetConfig.md) |  | [optional] 
**OriginalTargetConfig** | Pointer to [**NullablePhysicalTargetParamsForMountVolumeOriginalTargetConfig**](PhysicalTargetParamsForMountVolumeOriginalTargetConfig.md) |  | [optional] 
**ReadOnlyMount** | Pointer to **NullableBool** | Specifies whether to perform a read-only mount. Default is false. | [optional] 
**VlanConfig** | Pointer to [**NullableHyperVTargetParamsForMountVolumeVlanConfig**](HyperVTargetParamsForMountVolumeVlanConfig.md) |  | [optional] 
**VolumeNames** | Pointer to **[]string** | Specifies the names of volumes that need to be mounted. If this is not specified then all volumes that are part of the source VM will be mounted on the target VM. | [optional] 

## Methods

### NewPhysicalTargetParamsForMountVolume

`func NewPhysicalTargetParamsForMountVolume(mountToOriginalTarget NullableBool, ) *PhysicalTargetParamsForMountVolume`

NewPhysicalTargetParamsForMountVolume instantiates a new PhysicalTargetParamsForMountVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalTargetParamsForMountVolumeWithDefaults

`func NewPhysicalTargetParamsForMountVolumeWithDefaults() *PhysicalTargetParamsForMountVolume`

NewPhysicalTargetParamsForMountVolumeWithDefaults instantiates a new PhysicalTargetParamsForMountVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMountToOriginalTarget

`func (o *PhysicalTargetParamsForMountVolume) GetMountToOriginalTarget() bool`

GetMountToOriginalTarget returns the MountToOriginalTarget field if non-nil, zero value otherwise.

### GetMountToOriginalTargetOk

`func (o *PhysicalTargetParamsForMountVolume) GetMountToOriginalTargetOk() (*bool, bool)`

GetMountToOriginalTargetOk returns a tuple with the MountToOriginalTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountToOriginalTarget

`func (o *PhysicalTargetParamsForMountVolume) SetMountToOriginalTarget(v bool)`

SetMountToOriginalTarget sets MountToOriginalTarget field to given value.


### SetMountToOriginalTargetNil

`func (o *PhysicalTargetParamsForMountVolume) SetMountToOriginalTargetNil(b bool)`

 SetMountToOriginalTargetNil sets the value for MountToOriginalTarget to be an explicit nil

### UnsetMountToOriginalTarget
`func (o *PhysicalTargetParamsForMountVolume) UnsetMountToOriginalTarget()`

UnsetMountToOriginalTarget ensures that no value is present for MountToOriginalTarget, not even an explicit nil
### GetMountedVolumeMapping

`func (o *PhysicalTargetParamsForMountVolume) GetMountedVolumeMapping() []MountedVolumeMapping`

GetMountedVolumeMapping returns the MountedVolumeMapping field if non-nil, zero value otherwise.

### GetMountedVolumeMappingOk

`func (o *PhysicalTargetParamsForMountVolume) GetMountedVolumeMappingOk() (*[]MountedVolumeMapping, bool)`

GetMountedVolumeMappingOk returns a tuple with the MountedVolumeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountedVolumeMapping

`func (o *PhysicalTargetParamsForMountVolume) SetMountedVolumeMapping(v []MountedVolumeMapping)`

SetMountedVolumeMapping sets MountedVolumeMapping field to given value.

### HasMountedVolumeMapping

`func (o *PhysicalTargetParamsForMountVolume) HasMountedVolumeMapping() bool`

HasMountedVolumeMapping returns a boolean if a field has been set.

### SetMountedVolumeMappingNil

`func (o *PhysicalTargetParamsForMountVolume) SetMountedVolumeMappingNil(b bool)`

 SetMountedVolumeMappingNil sets the value for MountedVolumeMapping to be an explicit nil

### UnsetMountedVolumeMapping
`func (o *PhysicalTargetParamsForMountVolume) UnsetMountedVolumeMapping()`

UnsetMountedVolumeMapping ensures that no value is present for MountedVolumeMapping, not even an explicit nil
### GetNewTargetConfig

`func (o *PhysicalTargetParamsForMountVolume) GetNewTargetConfig() PhysicalTargetParamsForMountVolumeNewTargetConfig`

GetNewTargetConfig returns the NewTargetConfig field if non-nil, zero value otherwise.

### GetNewTargetConfigOk

`func (o *PhysicalTargetParamsForMountVolume) GetNewTargetConfigOk() (*PhysicalTargetParamsForMountVolumeNewTargetConfig, bool)`

GetNewTargetConfigOk returns a tuple with the NewTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTargetConfig

`func (o *PhysicalTargetParamsForMountVolume) SetNewTargetConfig(v PhysicalTargetParamsForMountVolumeNewTargetConfig)`

SetNewTargetConfig sets NewTargetConfig field to given value.

### HasNewTargetConfig

`func (o *PhysicalTargetParamsForMountVolume) HasNewTargetConfig() bool`

HasNewTargetConfig returns a boolean if a field has been set.

### SetNewTargetConfigNil

`func (o *PhysicalTargetParamsForMountVolume) SetNewTargetConfigNil(b bool)`

 SetNewTargetConfigNil sets the value for NewTargetConfig to be an explicit nil

### UnsetNewTargetConfig
`func (o *PhysicalTargetParamsForMountVolume) UnsetNewTargetConfig()`

UnsetNewTargetConfig ensures that no value is present for NewTargetConfig, not even an explicit nil
### GetOriginalTargetConfig

`func (o *PhysicalTargetParamsForMountVolume) GetOriginalTargetConfig() PhysicalTargetParamsForMountVolumeOriginalTargetConfig`

GetOriginalTargetConfig returns the OriginalTargetConfig field if non-nil, zero value otherwise.

### GetOriginalTargetConfigOk

`func (o *PhysicalTargetParamsForMountVolume) GetOriginalTargetConfigOk() (*PhysicalTargetParamsForMountVolumeOriginalTargetConfig, bool)`

GetOriginalTargetConfigOk returns a tuple with the OriginalTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTargetConfig

`func (o *PhysicalTargetParamsForMountVolume) SetOriginalTargetConfig(v PhysicalTargetParamsForMountVolumeOriginalTargetConfig)`

SetOriginalTargetConfig sets OriginalTargetConfig field to given value.

### HasOriginalTargetConfig

`func (o *PhysicalTargetParamsForMountVolume) HasOriginalTargetConfig() bool`

HasOriginalTargetConfig returns a boolean if a field has been set.

### SetOriginalTargetConfigNil

`func (o *PhysicalTargetParamsForMountVolume) SetOriginalTargetConfigNil(b bool)`

 SetOriginalTargetConfigNil sets the value for OriginalTargetConfig to be an explicit nil

### UnsetOriginalTargetConfig
`func (o *PhysicalTargetParamsForMountVolume) UnsetOriginalTargetConfig()`

UnsetOriginalTargetConfig ensures that no value is present for OriginalTargetConfig, not even an explicit nil
### GetReadOnlyMount

`func (o *PhysicalTargetParamsForMountVolume) GetReadOnlyMount() bool`

GetReadOnlyMount returns the ReadOnlyMount field if non-nil, zero value otherwise.

### GetReadOnlyMountOk

`func (o *PhysicalTargetParamsForMountVolume) GetReadOnlyMountOk() (*bool, bool)`

GetReadOnlyMountOk returns a tuple with the ReadOnlyMount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlyMount

`func (o *PhysicalTargetParamsForMountVolume) SetReadOnlyMount(v bool)`

SetReadOnlyMount sets ReadOnlyMount field to given value.

### HasReadOnlyMount

`func (o *PhysicalTargetParamsForMountVolume) HasReadOnlyMount() bool`

HasReadOnlyMount returns a boolean if a field has been set.

### SetReadOnlyMountNil

`func (o *PhysicalTargetParamsForMountVolume) SetReadOnlyMountNil(b bool)`

 SetReadOnlyMountNil sets the value for ReadOnlyMount to be an explicit nil

### UnsetReadOnlyMount
`func (o *PhysicalTargetParamsForMountVolume) UnsetReadOnlyMount()`

UnsetReadOnlyMount ensures that no value is present for ReadOnlyMount, not even an explicit nil
### GetVlanConfig

`func (o *PhysicalTargetParamsForMountVolume) GetVlanConfig() HyperVTargetParamsForMountVolumeVlanConfig`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *PhysicalTargetParamsForMountVolume) GetVlanConfigOk() (*HyperVTargetParamsForMountVolumeVlanConfig, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *PhysicalTargetParamsForMountVolume) SetVlanConfig(v HyperVTargetParamsForMountVolumeVlanConfig)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *PhysicalTargetParamsForMountVolume) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.

### SetVlanConfigNil

`func (o *PhysicalTargetParamsForMountVolume) SetVlanConfigNil(b bool)`

 SetVlanConfigNil sets the value for VlanConfig to be an explicit nil

### UnsetVlanConfig
`func (o *PhysicalTargetParamsForMountVolume) UnsetVlanConfig()`

UnsetVlanConfig ensures that no value is present for VlanConfig, not even an explicit nil
### GetVolumeNames

`func (o *PhysicalTargetParamsForMountVolume) GetVolumeNames() []string`

GetVolumeNames returns the VolumeNames field if non-nil, zero value otherwise.

### GetVolumeNamesOk

`func (o *PhysicalTargetParamsForMountVolume) GetVolumeNamesOk() (*[]string, bool)`

GetVolumeNamesOk returns a tuple with the VolumeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeNames

`func (o *PhysicalTargetParamsForMountVolume) SetVolumeNames(v []string)`

SetVolumeNames sets VolumeNames field to given value.

### HasVolumeNames

`func (o *PhysicalTargetParamsForMountVolume) HasVolumeNames() bool`

HasVolumeNames returns a boolean if a field has been set.

### SetVolumeNamesNil

`func (o *PhysicalTargetParamsForMountVolume) SetVolumeNamesNil(b bool)`

 SetVolumeNamesNil sets the value for VolumeNames to be an explicit nil

### UnsetVolumeNames
`func (o *PhysicalTargetParamsForMountVolume) UnsetVolumeNames()`

UnsetVolumeNames ensures that no value is present for VolumeNames, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


