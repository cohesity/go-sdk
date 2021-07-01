# RecoverVolumesParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForceUnmountVolume** | Pointer to **NullableBool** | Whether volume would be dismounted first during LockVolume failure | [optional] 
**MappingVec** | Pointer to [**[]RecoverVolumesParamsMapping**](RecoverVolumesParamsMapping.md) | Contains the volume mapping data that defines the restore task. | [optional] 
**TargetEntity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 

## Methods

### NewRecoverVolumesParams

`func NewRecoverVolumesParams() *RecoverVolumesParams`

NewRecoverVolumesParams instantiates a new RecoverVolumesParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverVolumesParamsWithDefaults

`func NewRecoverVolumesParamsWithDefaults() *RecoverVolumesParams`

NewRecoverVolumesParamsWithDefaults instantiates a new RecoverVolumesParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForceUnmountVolume

`func (o *RecoverVolumesParams) GetForceUnmountVolume() bool`

GetForceUnmountVolume returns the ForceUnmountVolume field if non-nil, zero value otherwise.

### GetForceUnmountVolumeOk

`func (o *RecoverVolumesParams) GetForceUnmountVolumeOk() (*bool, bool)`

GetForceUnmountVolumeOk returns a tuple with the ForceUnmountVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceUnmountVolume

`func (o *RecoverVolumesParams) SetForceUnmountVolume(v bool)`

SetForceUnmountVolume sets ForceUnmountVolume field to given value.

### HasForceUnmountVolume

`func (o *RecoverVolumesParams) HasForceUnmountVolume() bool`

HasForceUnmountVolume returns a boolean if a field has been set.

### SetForceUnmountVolumeNil

`func (o *RecoverVolumesParams) SetForceUnmountVolumeNil(b bool)`

 SetForceUnmountVolumeNil sets the value for ForceUnmountVolume to be an explicit nil

### UnsetForceUnmountVolume
`func (o *RecoverVolumesParams) UnsetForceUnmountVolume()`

UnsetForceUnmountVolume ensures that no value is present for ForceUnmountVolume, not even an explicit nil
### GetMappingVec

`func (o *RecoverVolumesParams) GetMappingVec() []RecoverVolumesParamsMapping`

GetMappingVec returns the MappingVec field if non-nil, zero value otherwise.

### GetMappingVecOk

`func (o *RecoverVolumesParams) GetMappingVecOk() (*[]RecoverVolumesParamsMapping, bool)`

GetMappingVecOk returns a tuple with the MappingVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingVec

`func (o *RecoverVolumesParams) SetMappingVec(v []RecoverVolumesParamsMapping)`

SetMappingVec sets MappingVec field to given value.

### HasMappingVec

`func (o *RecoverVolumesParams) HasMappingVec() bool`

HasMappingVec returns a boolean if a field has been set.

### SetMappingVecNil

`func (o *RecoverVolumesParams) SetMappingVecNil(b bool)`

 SetMappingVecNil sets the value for MappingVec to be an explicit nil

### UnsetMappingVec
`func (o *RecoverVolumesParams) UnsetMappingVec()`

UnsetMappingVec ensures that no value is present for MappingVec, not even an explicit nil
### GetTargetEntity

`func (o *RecoverVolumesParams) GetTargetEntity() EntityProto`

GetTargetEntity returns the TargetEntity field if non-nil, zero value otherwise.

### GetTargetEntityOk

`func (o *RecoverVolumesParams) GetTargetEntityOk() (*EntityProto, bool)`

GetTargetEntityOk returns a tuple with the TargetEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEntity

`func (o *RecoverVolumesParams) SetTargetEntity(v EntityProto)`

SetTargetEntity sets TargetEntity field to given value.

### HasTargetEntity

`func (o *RecoverVolumesParams) HasTargetEntity() bool`

HasTargetEntity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


