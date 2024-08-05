# MountPhysicalVolumeParamsPhysicalTargetParams

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

### NewMountPhysicalVolumeParamsPhysicalTargetParams

`func NewMountPhysicalVolumeParamsPhysicalTargetParams(mountToOriginalTarget NullableBool, ) *MountPhysicalVolumeParamsPhysicalTargetParams`

NewMountPhysicalVolumeParamsPhysicalTargetParams instantiates a new MountPhysicalVolumeParamsPhysicalTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMountPhysicalVolumeParamsPhysicalTargetParamsWithDefaults

`func NewMountPhysicalVolumeParamsPhysicalTargetParamsWithDefaults() *MountPhysicalVolumeParamsPhysicalTargetParams`

NewMountPhysicalVolumeParamsPhysicalTargetParamsWithDefaults instantiates a new MountPhysicalVolumeParamsPhysicalTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMountToOriginalTarget

`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) GetMountToOriginalTarget() bool`

GetMountToOriginalTarget returns the MountToOriginalTarget field if non-nil, zero value otherwise.

### GetMountToOriginalTargetOk

`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) GetMountToOriginalTargetOk() (*bool, bool)`

GetMountToOriginalTargetOk returns a tuple with the MountToOriginalTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountToOriginalTarget

`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) SetMountToOriginalTarget(v bool)`

SetMountToOriginalTarget sets MountToOriginalTarget field to given value.


### SetMountToOriginalTargetNil

`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) SetMountToOriginalTargetNil(b bool)`

 SetMountToOriginalTargetNil sets the value for MountToOriginalTarget to be an explicit nil

### UnsetMountToOriginalTarget
`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) UnsetMountToOriginalTarget()`

UnsetMountToOriginalTarget ensures that no value is present for MountToOriginalTarget, not even an explicit nil
### GetMountedVolumeMapping

`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) GetMountedVolumeMapping() []MountedVolumeMapping`

GetMountedVolumeMapping returns the MountedVolumeMapping field if non-nil, zero value otherwise.

### GetMountedVolumeMappingOk

`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) GetMountedVolumeMappingOk() (*[]MountedVolumeMapping, bool)`

GetMountedVolumeMappingOk returns a tuple with the MountedVolumeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountedVolumeMapping

`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) SetMountedVolumeMapping(v []MountedVolumeMapping)`

SetMountedVolumeMapping sets MountedVolumeMapping field to given value.

### HasMountedVolumeMapping

`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) HasMountedVolumeMapping() bool`

HasMountedVolumeMapping returns a boolean if a field has been set.

### SetMountedVolumeMappingNil

`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) SetMountedVolumeMappingNil(b bool)`

 SetMountedVolumeMappingNil sets the value for MountedVolumeMapping to be an explicit nil

### UnsetMountedVolumeMapping
`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) UnsetMountedVolumeMapping()`

UnsetMountedVolumeMapping ensures that no value is present for MountedVolumeMapping, not even an explicit nil
### GetNewTargetConfig

`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) GetNewTargetConfig() PhysicalTargetParamsForMountVolumeNewTargetConfig`

GetNewTargetConfig returns the NewTargetConfig field if non-nil, zero value otherwise.

### GetNewTargetConfigOk

`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) GetNewTargetConfigOk() (*PhysicalTargetParamsForMountVolumeNewTargetConfig, bool)`

GetNewTargetConfigOk returns a tuple with the NewTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTargetConfig

`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) SetNewTargetConfig(v PhysicalTargetParamsForMountVolumeNewTargetConfig)`

SetNewTargetConfig sets NewTargetConfig field to given value.

### HasNewTargetConfig

`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) HasNewTargetConfig() bool`

HasNewTargetConfig returns a boolean if a field has been set.

### SetNewTargetConfigNil

`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) SetNewTargetConfigNil(b bool)`

 SetNewTargetConfigNil sets the value for NewTargetConfig to be an explicit nil

### UnsetNewTargetConfig
`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) UnsetNewTargetConfig()`

UnsetNewTargetConfig ensures that no value is present for NewTargetConfig, not even an explicit nil
### GetOriginalTargetConfig

`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) GetOriginalTargetConfig() PhysicalTargetParamsForMountVolumeOriginalTargetConfig`

GetOriginalTargetConfig returns the OriginalTargetConfig field if non-nil, zero value otherwise.

### GetOriginalTargetConfigOk

`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) GetOriginalTargetConfigOk() (*PhysicalTargetParamsForMountVolumeOriginalTargetConfig, bool)`

GetOriginalTargetConfigOk returns a tuple with the OriginalTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTargetConfig

`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) SetOriginalTargetConfig(v PhysicalTargetParamsForMountVolumeOriginalTargetConfig)`

SetOriginalTargetConfig sets OriginalTargetConfig field to given value.

### HasOriginalTargetConfig

`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) HasOriginalTargetConfig() bool`

HasOriginalTargetConfig returns a boolean if a field has been set.

### SetOriginalTargetConfigNil

`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) SetOriginalTargetConfigNil(b bool)`

 SetOriginalTargetConfigNil sets the value for OriginalTargetConfig to be an explicit nil

### UnsetOriginalTargetConfig
`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) UnsetOriginalTargetConfig()`

UnsetOriginalTargetConfig ensures that no value is present for OriginalTargetConfig, not even an explicit nil
### GetReadOnlyMount

`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) GetReadOnlyMount() bool`

GetReadOnlyMount returns the ReadOnlyMount field if non-nil, zero value otherwise.

### GetReadOnlyMountOk

`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) GetReadOnlyMountOk() (*bool, bool)`

GetReadOnlyMountOk returns a tuple with the ReadOnlyMount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlyMount

`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) SetReadOnlyMount(v bool)`

SetReadOnlyMount sets ReadOnlyMount field to given value.

### HasReadOnlyMount

`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) HasReadOnlyMount() bool`

HasReadOnlyMount returns a boolean if a field has been set.

### SetReadOnlyMountNil

`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) SetReadOnlyMountNil(b bool)`

 SetReadOnlyMountNil sets the value for ReadOnlyMount to be an explicit nil

### UnsetReadOnlyMount
`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) UnsetReadOnlyMount()`

UnsetReadOnlyMount ensures that no value is present for ReadOnlyMount, not even an explicit nil
### GetVlanConfig

`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) GetVlanConfig() HyperVTargetParamsForMountVolumeVlanConfig`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) GetVlanConfigOk() (*HyperVTargetParamsForMountVolumeVlanConfig, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) SetVlanConfig(v HyperVTargetParamsForMountVolumeVlanConfig)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.

### SetVlanConfigNil

`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) SetVlanConfigNil(b bool)`

 SetVlanConfigNil sets the value for VlanConfig to be an explicit nil

### UnsetVlanConfig
`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) UnsetVlanConfig()`

UnsetVlanConfig ensures that no value is present for VlanConfig, not even an explicit nil
### GetVolumeNames

`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) GetVolumeNames() []string`

GetVolumeNames returns the VolumeNames field if non-nil, zero value otherwise.

### GetVolumeNamesOk

`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) GetVolumeNamesOk() (*[]string, bool)`

GetVolumeNamesOk returns a tuple with the VolumeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeNames

`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) SetVolumeNames(v []string)`

SetVolumeNames sets VolumeNames field to given value.

### HasVolumeNames

`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) HasVolumeNames() bool`

HasVolumeNames returns a boolean if a field has been set.

### SetVolumeNamesNil

`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) SetVolumeNamesNil(b bool)`

 SetVolumeNamesNil sets the value for VolumeNames to be an explicit nil

### UnsetVolumeNames
`func (o *MountPhysicalVolumeParamsPhysicalTargetParams) UnsetVolumeNames()`

UnsetVolumeNames ensures that no value is present for VolumeNames, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


