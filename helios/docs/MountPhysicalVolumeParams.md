# MountPhysicalVolumeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 
**PhysicalTargetParams** | Pointer to [**NullablePhysicalTargetParamsForMountVolume**](PhysicalTargetParamsForMountVolume.md) | Specifies the params for recovering to a physical target. | [optional] 

## Methods

### NewMountPhysicalVolumeParams

`func NewMountPhysicalVolumeParams(targetEnvironment string, ) *MountPhysicalVolumeParams`

NewMountPhysicalVolumeParams instantiates a new MountPhysicalVolumeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMountPhysicalVolumeParamsWithDefaults

`func NewMountPhysicalVolumeParamsWithDefaults() *MountPhysicalVolumeParams`

NewMountPhysicalVolumeParamsWithDefaults instantiates a new MountPhysicalVolumeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetEnvironment

`func (o *MountPhysicalVolumeParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *MountPhysicalVolumeParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *MountPhysicalVolumeParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.


### GetPhysicalTargetParams

`func (o *MountPhysicalVolumeParams) GetPhysicalTargetParams() PhysicalTargetParamsForMountVolume`

GetPhysicalTargetParams returns the PhysicalTargetParams field if non-nil, zero value otherwise.

### GetPhysicalTargetParamsOk

`func (o *MountPhysicalVolumeParams) GetPhysicalTargetParamsOk() (*PhysicalTargetParamsForMountVolume, bool)`

GetPhysicalTargetParamsOk returns a tuple with the PhysicalTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalTargetParams

`func (o *MountPhysicalVolumeParams) SetPhysicalTargetParams(v PhysicalTargetParamsForMountVolume)`

SetPhysicalTargetParams sets PhysicalTargetParams field to given value.

### HasPhysicalTargetParams

`func (o *MountPhysicalVolumeParams) HasPhysicalTargetParams() bool`

HasPhysicalTargetParams returns a boolean if a field has been set.

### SetPhysicalTargetParamsNil

`func (o *MountPhysicalVolumeParams) SetPhysicalTargetParamsNil(b bool)`

 SetPhysicalTargetParamsNil sets the value for PhysicalTargetParams to be an explicit nil

### UnsetPhysicalTargetParams
`func (o *MountPhysicalVolumeParams) UnsetPhysicalTargetParams()`

UnsetPhysicalTargetParams ensures that no value is present for PhysicalTargetParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


