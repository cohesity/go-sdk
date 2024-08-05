# HyperVTargetParamsForMountVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MountToOriginalTarget** | **NullableBool** | Specifies whether to mount to the original target. If true, originalTargetConfig must be specified. If false, newTargetConfig must be specified. | 
**MountedVolumeMapping** | Pointer to [**[]MountedVolumeMapping**](MountedVolumeMapping.md) | Specifies the mapping of original volumes and mounted volumes | [optional] [readonly] 
**NewTargetConfig** | Pointer to [**NullableHyperVTargetParamsForMountVolumeNewTargetConfig**](HyperVTargetParamsForMountVolumeNewTargetConfig.md) |  | [optional] 
**OriginalTargetConfig** | Pointer to [**NullableHyperVTargetParamsForMountVolumeOriginalTargetConfig**](HyperVTargetParamsForMountVolumeOriginalTargetConfig.md) |  | [optional] 
**ReadOnlyMount** | Pointer to **NullableBool** | Specifies whether to perform a read-only mount. Default is false. | [optional] 
**VlanConfig** | Pointer to [**NullableHyperVTargetParamsForMountVolumeVlanConfig**](HyperVTargetParamsForMountVolumeVlanConfig.md) |  | [optional] 
**VolumeNames** | Pointer to **[]string** | Specifies the names of volumes that need to be mounted. If this is not specified then all volumes that are part of the source VM will be mounted on the target VM. | [optional] 

## Methods

### NewHyperVTargetParamsForMountVolume

`func NewHyperVTargetParamsForMountVolume(mountToOriginalTarget NullableBool, ) *HyperVTargetParamsForMountVolume`

NewHyperVTargetParamsForMountVolume instantiates a new HyperVTargetParamsForMountVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperVTargetParamsForMountVolumeWithDefaults

`func NewHyperVTargetParamsForMountVolumeWithDefaults() *HyperVTargetParamsForMountVolume`

NewHyperVTargetParamsForMountVolumeWithDefaults instantiates a new HyperVTargetParamsForMountVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMountToOriginalTarget

`func (o *HyperVTargetParamsForMountVolume) GetMountToOriginalTarget() bool`

GetMountToOriginalTarget returns the MountToOriginalTarget field if non-nil, zero value otherwise.

### GetMountToOriginalTargetOk

`func (o *HyperVTargetParamsForMountVolume) GetMountToOriginalTargetOk() (*bool, bool)`

GetMountToOriginalTargetOk returns a tuple with the MountToOriginalTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountToOriginalTarget

`func (o *HyperVTargetParamsForMountVolume) SetMountToOriginalTarget(v bool)`

SetMountToOriginalTarget sets MountToOriginalTarget field to given value.


### SetMountToOriginalTargetNil

`func (o *HyperVTargetParamsForMountVolume) SetMountToOriginalTargetNil(b bool)`

 SetMountToOriginalTargetNil sets the value for MountToOriginalTarget to be an explicit nil

### UnsetMountToOriginalTarget
`func (o *HyperVTargetParamsForMountVolume) UnsetMountToOriginalTarget()`

UnsetMountToOriginalTarget ensures that no value is present for MountToOriginalTarget, not even an explicit nil
### GetMountedVolumeMapping

`func (o *HyperVTargetParamsForMountVolume) GetMountedVolumeMapping() []MountedVolumeMapping`

GetMountedVolumeMapping returns the MountedVolumeMapping field if non-nil, zero value otherwise.

### GetMountedVolumeMappingOk

`func (o *HyperVTargetParamsForMountVolume) GetMountedVolumeMappingOk() (*[]MountedVolumeMapping, bool)`

GetMountedVolumeMappingOk returns a tuple with the MountedVolumeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountedVolumeMapping

`func (o *HyperVTargetParamsForMountVolume) SetMountedVolumeMapping(v []MountedVolumeMapping)`

SetMountedVolumeMapping sets MountedVolumeMapping field to given value.

### HasMountedVolumeMapping

`func (o *HyperVTargetParamsForMountVolume) HasMountedVolumeMapping() bool`

HasMountedVolumeMapping returns a boolean if a field has been set.

### SetMountedVolumeMappingNil

`func (o *HyperVTargetParamsForMountVolume) SetMountedVolumeMappingNil(b bool)`

 SetMountedVolumeMappingNil sets the value for MountedVolumeMapping to be an explicit nil

### UnsetMountedVolumeMapping
`func (o *HyperVTargetParamsForMountVolume) UnsetMountedVolumeMapping()`

UnsetMountedVolumeMapping ensures that no value is present for MountedVolumeMapping, not even an explicit nil
### GetNewTargetConfig

`func (o *HyperVTargetParamsForMountVolume) GetNewTargetConfig() HyperVTargetParamsForMountVolumeNewTargetConfig`

GetNewTargetConfig returns the NewTargetConfig field if non-nil, zero value otherwise.

### GetNewTargetConfigOk

`func (o *HyperVTargetParamsForMountVolume) GetNewTargetConfigOk() (*HyperVTargetParamsForMountVolumeNewTargetConfig, bool)`

GetNewTargetConfigOk returns a tuple with the NewTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTargetConfig

`func (o *HyperVTargetParamsForMountVolume) SetNewTargetConfig(v HyperVTargetParamsForMountVolumeNewTargetConfig)`

SetNewTargetConfig sets NewTargetConfig field to given value.

### HasNewTargetConfig

`func (o *HyperVTargetParamsForMountVolume) HasNewTargetConfig() bool`

HasNewTargetConfig returns a boolean if a field has been set.

### SetNewTargetConfigNil

`func (o *HyperVTargetParamsForMountVolume) SetNewTargetConfigNil(b bool)`

 SetNewTargetConfigNil sets the value for NewTargetConfig to be an explicit nil

### UnsetNewTargetConfig
`func (o *HyperVTargetParamsForMountVolume) UnsetNewTargetConfig()`

UnsetNewTargetConfig ensures that no value is present for NewTargetConfig, not even an explicit nil
### GetOriginalTargetConfig

`func (o *HyperVTargetParamsForMountVolume) GetOriginalTargetConfig() HyperVTargetParamsForMountVolumeOriginalTargetConfig`

GetOriginalTargetConfig returns the OriginalTargetConfig field if non-nil, zero value otherwise.

### GetOriginalTargetConfigOk

`func (o *HyperVTargetParamsForMountVolume) GetOriginalTargetConfigOk() (*HyperVTargetParamsForMountVolumeOriginalTargetConfig, bool)`

GetOriginalTargetConfigOk returns a tuple with the OriginalTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTargetConfig

`func (o *HyperVTargetParamsForMountVolume) SetOriginalTargetConfig(v HyperVTargetParamsForMountVolumeOriginalTargetConfig)`

SetOriginalTargetConfig sets OriginalTargetConfig field to given value.

### HasOriginalTargetConfig

`func (o *HyperVTargetParamsForMountVolume) HasOriginalTargetConfig() bool`

HasOriginalTargetConfig returns a boolean if a field has been set.

### SetOriginalTargetConfigNil

`func (o *HyperVTargetParamsForMountVolume) SetOriginalTargetConfigNil(b bool)`

 SetOriginalTargetConfigNil sets the value for OriginalTargetConfig to be an explicit nil

### UnsetOriginalTargetConfig
`func (o *HyperVTargetParamsForMountVolume) UnsetOriginalTargetConfig()`

UnsetOriginalTargetConfig ensures that no value is present for OriginalTargetConfig, not even an explicit nil
### GetReadOnlyMount

`func (o *HyperVTargetParamsForMountVolume) GetReadOnlyMount() bool`

GetReadOnlyMount returns the ReadOnlyMount field if non-nil, zero value otherwise.

### GetReadOnlyMountOk

`func (o *HyperVTargetParamsForMountVolume) GetReadOnlyMountOk() (*bool, bool)`

GetReadOnlyMountOk returns a tuple with the ReadOnlyMount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlyMount

`func (o *HyperVTargetParamsForMountVolume) SetReadOnlyMount(v bool)`

SetReadOnlyMount sets ReadOnlyMount field to given value.

### HasReadOnlyMount

`func (o *HyperVTargetParamsForMountVolume) HasReadOnlyMount() bool`

HasReadOnlyMount returns a boolean if a field has been set.

### SetReadOnlyMountNil

`func (o *HyperVTargetParamsForMountVolume) SetReadOnlyMountNil(b bool)`

 SetReadOnlyMountNil sets the value for ReadOnlyMount to be an explicit nil

### UnsetReadOnlyMount
`func (o *HyperVTargetParamsForMountVolume) UnsetReadOnlyMount()`

UnsetReadOnlyMount ensures that no value is present for ReadOnlyMount, not even an explicit nil
### GetVlanConfig

`func (o *HyperVTargetParamsForMountVolume) GetVlanConfig() HyperVTargetParamsForMountVolumeVlanConfig`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *HyperVTargetParamsForMountVolume) GetVlanConfigOk() (*HyperVTargetParamsForMountVolumeVlanConfig, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *HyperVTargetParamsForMountVolume) SetVlanConfig(v HyperVTargetParamsForMountVolumeVlanConfig)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *HyperVTargetParamsForMountVolume) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.

### SetVlanConfigNil

`func (o *HyperVTargetParamsForMountVolume) SetVlanConfigNil(b bool)`

 SetVlanConfigNil sets the value for VlanConfig to be an explicit nil

### UnsetVlanConfig
`func (o *HyperVTargetParamsForMountVolume) UnsetVlanConfig()`

UnsetVlanConfig ensures that no value is present for VlanConfig, not even an explicit nil
### GetVolumeNames

`func (o *HyperVTargetParamsForMountVolume) GetVolumeNames() []string`

GetVolumeNames returns the VolumeNames field if non-nil, zero value otherwise.

### GetVolumeNamesOk

`func (o *HyperVTargetParamsForMountVolume) GetVolumeNamesOk() (*[]string, bool)`

GetVolumeNamesOk returns a tuple with the VolumeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeNames

`func (o *HyperVTargetParamsForMountVolume) SetVolumeNames(v []string)`

SetVolumeNames sets VolumeNames field to given value.

### HasVolumeNames

`func (o *HyperVTargetParamsForMountVolume) HasVolumeNames() bool`

HasVolumeNames returns a boolean if a field has been set.

### SetVolumeNamesNil

`func (o *HyperVTargetParamsForMountVolume) SetVolumeNamesNil(b bool)`

 SetVolumeNamesNil sets the value for VolumeNames to be an explicit nil

### UnsetVolumeNames
`func (o *HyperVTargetParamsForMountVolume) UnsetVolumeNames()`

UnsetVolumeNames ensures that no value is present for VolumeNames, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


