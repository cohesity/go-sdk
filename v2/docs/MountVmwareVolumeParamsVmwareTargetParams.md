# MountVmwareVolumeParamsVmwareTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MountToOriginalTarget** | **NullableBool** | Specifies whether to mount to the original target. If true, originalTargetConfig must be specified. If false, newTargetConfig must be specified. | 
**MountedVolumeMapping** | Pointer to [**[]MountedVolumeMapping**](MountedVolumeMapping.md) | Specifies the mapping of original volumes and mounted volumes | [optional] [readonly] 
**NewTargetConfig** | Pointer to [**NullableVmwareTargetParamsForMountVolumeNewTargetConfig**](VmwareTargetParamsForMountVolumeNewTargetConfig.md) |  | [optional] 
**OriginalTargetConfig** | Pointer to [**NullableVmwareTargetParamsForMountVolumeOriginalTargetConfig**](VmwareTargetParamsForMountVolumeOriginalTargetConfig.md) |  | [optional] 
**ReadOnlyMount** | Pointer to **NullableBool** | Specifies whether to perform a read-only mount. Default is false. | [optional] 
**VlanConfig** | Pointer to [**NullableHyperVTargetParamsForMountVolumeVlanConfig**](HyperVTargetParamsForMountVolumeVlanConfig.md) |  | [optional] 
**VolumeNames** | Pointer to **[]string** | Specifies the names of volumes that need to be mounted. If this is not specified then all volumes that are part of the source VM will be mounted on the target VM. | [optional] 

## Methods

### NewMountVmwareVolumeParamsVmwareTargetParams

`func NewMountVmwareVolumeParamsVmwareTargetParams(mountToOriginalTarget NullableBool, ) *MountVmwareVolumeParamsVmwareTargetParams`

NewMountVmwareVolumeParamsVmwareTargetParams instantiates a new MountVmwareVolumeParamsVmwareTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMountVmwareVolumeParamsVmwareTargetParamsWithDefaults

`func NewMountVmwareVolumeParamsVmwareTargetParamsWithDefaults() *MountVmwareVolumeParamsVmwareTargetParams`

NewMountVmwareVolumeParamsVmwareTargetParamsWithDefaults instantiates a new MountVmwareVolumeParamsVmwareTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMountToOriginalTarget

`func (o *MountVmwareVolumeParamsVmwareTargetParams) GetMountToOriginalTarget() bool`

GetMountToOriginalTarget returns the MountToOriginalTarget field if non-nil, zero value otherwise.

### GetMountToOriginalTargetOk

`func (o *MountVmwareVolumeParamsVmwareTargetParams) GetMountToOriginalTargetOk() (*bool, bool)`

GetMountToOriginalTargetOk returns a tuple with the MountToOriginalTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountToOriginalTarget

`func (o *MountVmwareVolumeParamsVmwareTargetParams) SetMountToOriginalTarget(v bool)`

SetMountToOriginalTarget sets MountToOriginalTarget field to given value.


### SetMountToOriginalTargetNil

`func (o *MountVmwareVolumeParamsVmwareTargetParams) SetMountToOriginalTargetNil(b bool)`

 SetMountToOriginalTargetNil sets the value for MountToOriginalTarget to be an explicit nil

### UnsetMountToOriginalTarget
`func (o *MountVmwareVolumeParamsVmwareTargetParams) UnsetMountToOriginalTarget()`

UnsetMountToOriginalTarget ensures that no value is present for MountToOriginalTarget, not even an explicit nil
### GetMountedVolumeMapping

`func (o *MountVmwareVolumeParamsVmwareTargetParams) GetMountedVolumeMapping() []MountedVolumeMapping`

GetMountedVolumeMapping returns the MountedVolumeMapping field if non-nil, zero value otherwise.

### GetMountedVolumeMappingOk

`func (o *MountVmwareVolumeParamsVmwareTargetParams) GetMountedVolumeMappingOk() (*[]MountedVolumeMapping, bool)`

GetMountedVolumeMappingOk returns a tuple with the MountedVolumeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountedVolumeMapping

`func (o *MountVmwareVolumeParamsVmwareTargetParams) SetMountedVolumeMapping(v []MountedVolumeMapping)`

SetMountedVolumeMapping sets MountedVolumeMapping field to given value.

### HasMountedVolumeMapping

`func (o *MountVmwareVolumeParamsVmwareTargetParams) HasMountedVolumeMapping() bool`

HasMountedVolumeMapping returns a boolean if a field has been set.

### SetMountedVolumeMappingNil

`func (o *MountVmwareVolumeParamsVmwareTargetParams) SetMountedVolumeMappingNil(b bool)`

 SetMountedVolumeMappingNil sets the value for MountedVolumeMapping to be an explicit nil

### UnsetMountedVolumeMapping
`func (o *MountVmwareVolumeParamsVmwareTargetParams) UnsetMountedVolumeMapping()`

UnsetMountedVolumeMapping ensures that no value is present for MountedVolumeMapping, not even an explicit nil
### GetNewTargetConfig

`func (o *MountVmwareVolumeParamsVmwareTargetParams) GetNewTargetConfig() VmwareTargetParamsForMountVolumeNewTargetConfig`

GetNewTargetConfig returns the NewTargetConfig field if non-nil, zero value otherwise.

### GetNewTargetConfigOk

`func (o *MountVmwareVolumeParamsVmwareTargetParams) GetNewTargetConfigOk() (*VmwareTargetParamsForMountVolumeNewTargetConfig, bool)`

GetNewTargetConfigOk returns a tuple with the NewTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTargetConfig

`func (o *MountVmwareVolumeParamsVmwareTargetParams) SetNewTargetConfig(v VmwareTargetParamsForMountVolumeNewTargetConfig)`

SetNewTargetConfig sets NewTargetConfig field to given value.

### HasNewTargetConfig

`func (o *MountVmwareVolumeParamsVmwareTargetParams) HasNewTargetConfig() bool`

HasNewTargetConfig returns a boolean if a field has been set.

### SetNewTargetConfigNil

`func (o *MountVmwareVolumeParamsVmwareTargetParams) SetNewTargetConfigNil(b bool)`

 SetNewTargetConfigNil sets the value for NewTargetConfig to be an explicit nil

### UnsetNewTargetConfig
`func (o *MountVmwareVolumeParamsVmwareTargetParams) UnsetNewTargetConfig()`

UnsetNewTargetConfig ensures that no value is present for NewTargetConfig, not even an explicit nil
### GetOriginalTargetConfig

`func (o *MountVmwareVolumeParamsVmwareTargetParams) GetOriginalTargetConfig() VmwareTargetParamsForMountVolumeOriginalTargetConfig`

GetOriginalTargetConfig returns the OriginalTargetConfig field if non-nil, zero value otherwise.

### GetOriginalTargetConfigOk

`func (o *MountVmwareVolumeParamsVmwareTargetParams) GetOriginalTargetConfigOk() (*VmwareTargetParamsForMountVolumeOriginalTargetConfig, bool)`

GetOriginalTargetConfigOk returns a tuple with the OriginalTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTargetConfig

`func (o *MountVmwareVolumeParamsVmwareTargetParams) SetOriginalTargetConfig(v VmwareTargetParamsForMountVolumeOriginalTargetConfig)`

SetOriginalTargetConfig sets OriginalTargetConfig field to given value.

### HasOriginalTargetConfig

`func (o *MountVmwareVolumeParamsVmwareTargetParams) HasOriginalTargetConfig() bool`

HasOriginalTargetConfig returns a boolean if a field has been set.

### SetOriginalTargetConfigNil

`func (o *MountVmwareVolumeParamsVmwareTargetParams) SetOriginalTargetConfigNil(b bool)`

 SetOriginalTargetConfigNil sets the value for OriginalTargetConfig to be an explicit nil

### UnsetOriginalTargetConfig
`func (o *MountVmwareVolumeParamsVmwareTargetParams) UnsetOriginalTargetConfig()`

UnsetOriginalTargetConfig ensures that no value is present for OriginalTargetConfig, not even an explicit nil
### GetReadOnlyMount

`func (o *MountVmwareVolumeParamsVmwareTargetParams) GetReadOnlyMount() bool`

GetReadOnlyMount returns the ReadOnlyMount field if non-nil, zero value otherwise.

### GetReadOnlyMountOk

`func (o *MountVmwareVolumeParamsVmwareTargetParams) GetReadOnlyMountOk() (*bool, bool)`

GetReadOnlyMountOk returns a tuple with the ReadOnlyMount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlyMount

`func (o *MountVmwareVolumeParamsVmwareTargetParams) SetReadOnlyMount(v bool)`

SetReadOnlyMount sets ReadOnlyMount field to given value.

### HasReadOnlyMount

`func (o *MountVmwareVolumeParamsVmwareTargetParams) HasReadOnlyMount() bool`

HasReadOnlyMount returns a boolean if a field has been set.

### SetReadOnlyMountNil

`func (o *MountVmwareVolumeParamsVmwareTargetParams) SetReadOnlyMountNil(b bool)`

 SetReadOnlyMountNil sets the value for ReadOnlyMount to be an explicit nil

### UnsetReadOnlyMount
`func (o *MountVmwareVolumeParamsVmwareTargetParams) UnsetReadOnlyMount()`

UnsetReadOnlyMount ensures that no value is present for ReadOnlyMount, not even an explicit nil
### GetVlanConfig

`func (o *MountVmwareVolumeParamsVmwareTargetParams) GetVlanConfig() HyperVTargetParamsForMountVolumeVlanConfig`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *MountVmwareVolumeParamsVmwareTargetParams) GetVlanConfigOk() (*HyperVTargetParamsForMountVolumeVlanConfig, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *MountVmwareVolumeParamsVmwareTargetParams) SetVlanConfig(v HyperVTargetParamsForMountVolumeVlanConfig)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *MountVmwareVolumeParamsVmwareTargetParams) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.

### SetVlanConfigNil

`func (o *MountVmwareVolumeParamsVmwareTargetParams) SetVlanConfigNil(b bool)`

 SetVlanConfigNil sets the value for VlanConfig to be an explicit nil

### UnsetVlanConfig
`func (o *MountVmwareVolumeParamsVmwareTargetParams) UnsetVlanConfig()`

UnsetVlanConfig ensures that no value is present for VlanConfig, not even an explicit nil
### GetVolumeNames

`func (o *MountVmwareVolumeParamsVmwareTargetParams) GetVolumeNames() []string`

GetVolumeNames returns the VolumeNames field if non-nil, zero value otherwise.

### GetVolumeNamesOk

`func (o *MountVmwareVolumeParamsVmwareTargetParams) GetVolumeNamesOk() (*[]string, bool)`

GetVolumeNamesOk returns a tuple with the VolumeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeNames

`func (o *MountVmwareVolumeParamsVmwareTargetParams) SetVolumeNames(v []string)`

SetVolumeNames sets VolumeNames field to given value.

### HasVolumeNames

`func (o *MountVmwareVolumeParamsVmwareTargetParams) HasVolumeNames() bool`

HasVolumeNames returns a boolean if a field has been set.

### SetVolumeNamesNil

`func (o *MountVmwareVolumeParamsVmwareTargetParams) SetVolumeNamesNil(b bool)`

 SetVolumeNamesNil sets the value for VolumeNames to be an explicit nil

### UnsetVolumeNames
`func (o *MountVmwareVolumeParamsVmwareTargetParams) UnsetVolumeNames()`

UnsetVolumeNames ensures that no value is present for VolumeNames, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


