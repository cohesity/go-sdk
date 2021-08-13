# MountHyperVVolumeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 
**HypervTargetParams** | Pointer to [**NullableHyperVTargetParamsForMountVolume**](HyperVTargetParamsForMountVolume.md) | Specifies the params for recovering to a HyperV target. | [optional] 

## Methods

### NewMountHyperVVolumeParams

`func NewMountHyperVVolumeParams(targetEnvironment string, ) *MountHyperVVolumeParams`

NewMountHyperVVolumeParams instantiates a new MountHyperVVolumeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMountHyperVVolumeParamsWithDefaults

`func NewMountHyperVVolumeParamsWithDefaults() *MountHyperVVolumeParams`

NewMountHyperVVolumeParamsWithDefaults instantiates a new MountHyperVVolumeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetEnvironment

`func (o *MountHyperVVolumeParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *MountHyperVVolumeParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *MountHyperVVolumeParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.


### GetHypervTargetParams

`func (o *MountHyperVVolumeParams) GetHypervTargetParams() HyperVTargetParamsForMountVolume`

GetHypervTargetParams returns the HypervTargetParams field if non-nil, zero value otherwise.

### GetHypervTargetParamsOk

`func (o *MountHyperVVolumeParams) GetHypervTargetParamsOk() (*HyperVTargetParamsForMountVolume, bool)`

GetHypervTargetParamsOk returns a tuple with the HypervTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervTargetParams

`func (o *MountHyperVVolumeParams) SetHypervTargetParams(v HyperVTargetParamsForMountVolume)`

SetHypervTargetParams sets HypervTargetParams field to given value.

### HasHypervTargetParams

`func (o *MountHyperVVolumeParams) HasHypervTargetParams() bool`

HasHypervTargetParams returns a boolean if a field has been set.

### SetHypervTargetParamsNil

`func (o *MountHyperVVolumeParams) SetHypervTargetParamsNil(b bool)`

 SetHypervTargetParamsNil sets the value for HypervTargetParams to be an explicit nil

### UnsetHypervTargetParams
`func (o *MountHyperVVolumeParams) UnsetHypervTargetParams()`

UnsetHypervTargetParams ensures that no value is present for HypervTargetParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


