# MountVolumesParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HypervParams** | Pointer to [**MountVolumesHyperVParams**](MountVolumesHyperVParams.md) |  | [optional] 
**ReadonlyMount** | Pointer to **NullableBool** | Allows the caller to force the Agent to perform a read-only mount. This is not usually required and we want to give customers the ability to mutate this mount for test/dev purposes. | [optional] 
**TargetEntity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**UseExistingAgent** | Pointer to **NullableBool** | Whether this will use an existing agent on the target vm to do a restore operation. | [optional] 
**VmwareParams** | Pointer to [**MountVolumesVMwareParams**](MountVolumesVMwareParams.md) |  | [optional] 
**VolumeNameVec** | Pointer to **[]string** | Optional names of volumes that need to be mounted. The names here correspond to the volume names obtained by Iris from Yoda as part of VMVolumeInfo call. NOTE: If this is not specified then all volumes that are part of the server will be mounted on the target entity. | [optional] 

## Methods

### NewMountVolumesParams

`func NewMountVolumesParams() *MountVolumesParams`

NewMountVolumesParams instantiates a new MountVolumesParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMountVolumesParamsWithDefaults

`func NewMountVolumesParamsWithDefaults() *MountVolumesParams`

NewMountVolumesParamsWithDefaults instantiates a new MountVolumesParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHypervParams

`func (o *MountVolumesParams) GetHypervParams() MountVolumesHyperVParams`

GetHypervParams returns the HypervParams field if non-nil, zero value otherwise.

### GetHypervParamsOk

`func (o *MountVolumesParams) GetHypervParamsOk() (*MountVolumesHyperVParams, bool)`

GetHypervParamsOk returns a tuple with the HypervParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervParams

`func (o *MountVolumesParams) SetHypervParams(v MountVolumesHyperVParams)`

SetHypervParams sets HypervParams field to given value.

### HasHypervParams

`func (o *MountVolumesParams) HasHypervParams() bool`

HasHypervParams returns a boolean if a field has been set.

### GetReadonlyMount

`func (o *MountVolumesParams) GetReadonlyMount() bool`

GetReadonlyMount returns the ReadonlyMount field if non-nil, zero value otherwise.

### GetReadonlyMountOk

`func (o *MountVolumesParams) GetReadonlyMountOk() (*bool, bool)`

GetReadonlyMountOk returns a tuple with the ReadonlyMount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonlyMount

`func (o *MountVolumesParams) SetReadonlyMount(v bool)`

SetReadonlyMount sets ReadonlyMount field to given value.

### HasReadonlyMount

`func (o *MountVolumesParams) HasReadonlyMount() bool`

HasReadonlyMount returns a boolean if a field has been set.

### SetReadonlyMountNil

`func (o *MountVolumesParams) SetReadonlyMountNil(b bool)`

 SetReadonlyMountNil sets the value for ReadonlyMount to be an explicit nil

### UnsetReadonlyMount
`func (o *MountVolumesParams) UnsetReadonlyMount()`

UnsetReadonlyMount ensures that no value is present for ReadonlyMount, not even an explicit nil
### GetTargetEntity

`func (o *MountVolumesParams) GetTargetEntity() EntityProto`

GetTargetEntity returns the TargetEntity field if non-nil, zero value otherwise.

### GetTargetEntityOk

`func (o *MountVolumesParams) GetTargetEntityOk() (*EntityProto, bool)`

GetTargetEntityOk returns a tuple with the TargetEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEntity

`func (o *MountVolumesParams) SetTargetEntity(v EntityProto)`

SetTargetEntity sets TargetEntity field to given value.

### HasTargetEntity

`func (o *MountVolumesParams) HasTargetEntity() bool`

HasTargetEntity returns a boolean if a field has been set.

### GetUseExistingAgent

`func (o *MountVolumesParams) GetUseExistingAgent() bool`

GetUseExistingAgent returns the UseExistingAgent field if non-nil, zero value otherwise.

### GetUseExistingAgentOk

`func (o *MountVolumesParams) GetUseExistingAgentOk() (*bool, bool)`

GetUseExistingAgentOk returns a tuple with the UseExistingAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseExistingAgent

`func (o *MountVolumesParams) SetUseExistingAgent(v bool)`

SetUseExistingAgent sets UseExistingAgent field to given value.

### HasUseExistingAgent

`func (o *MountVolumesParams) HasUseExistingAgent() bool`

HasUseExistingAgent returns a boolean if a field has been set.

### SetUseExistingAgentNil

`func (o *MountVolumesParams) SetUseExistingAgentNil(b bool)`

 SetUseExistingAgentNil sets the value for UseExistingAgent to be an explicit nil

### UnsetUseExistingAgent
`func (o *MountVolumesParams) UnsetUseExistingAgent()`

UnsetUseExistingAgent ensures that no value is present for UseExistingAgent, not even an explicit nil
### GetVmwareParams

`func (o *MountVolumesParams) GetVmwareParams() MountVolumesVMwareParams`

GetVmwareParams returns the VmwareParams field if non-nil, zero value otherwise.

### GetVmwareParamsOk

`func (o *MountVolumesParams) GetVmwareParamsOk() (*MountVolumesVMwareParams, bool)`

GetVmwareParamsOk returns a tuple with the VmwareParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareParams

`func (o *MountVolumesParams) SetVmwareParams(v MountVolumesVMwareParams)`

SetVmwareParams sets VmwareParams field to given value.

### HasVmwareParams

`func (o *MountVolumesParams) HasVmwareParams() bool`

HasVmwareParams returns a boolean if a field has been set.

### GetVolumeNameVec

`func (o *MountVolumesParams) GetVolumeNameVec() []string`

GetVolumeNameVec returns the VolumeNameVec field if non-nil, zero value otherwise.

### GetVolumeNameVecOk

`func (o *MountVolumesParams) GetVolumeNameVecOk() (*[]string, bool)`

GetVolumeNameVecOk returns a tuple with the VolumeNameVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeNameVec

`func (o *MountVolumesParams) SetVolumeNameVec(v []string)`

SetVolumeNameVec sets VolumeNameVec field to given value.

### HasVolumeNameVec

`func (o *MountVolumesParams) HasVolumeNameVec() bool`

HasVolumeNameVec returns a boolean if a field has been set.

### SetVolumeNameVecNil

`func (o *MountVolumesParams) SetVolumeNameVecNil(b bool)`

 SetVolumeNameVecNil sets the value for VolumeNameVec to be an explicit nil

### UnsetVolumeNameVec
`func (o *MountVolumesParams) UnsetVolumeNameVec()`

UnsetVolumeNameVec ensures that no value is present for VolumeNameVec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


