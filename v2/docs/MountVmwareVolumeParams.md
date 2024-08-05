# MountVmwareVolumeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 
**VmwareTargetParams** | Pointer to [**NullableMountVmwareVolumeParamsVmwareTargetParams**](MountVmwareVolumeParamsVmwareTargetParams.md) |  | [optional] 

## Methods

### NewMountVmwareVolumeParams

`func NewMountVmwareVolumeParams(targetEnvironment string, ) *MountVmwareVolumeParams`

NewMountVmwareVolumeParams instantiates a new MountVmwareVolumeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMountVmwareVolumeParamsWithDefaults

`func NewMountVmwareVolumeParamsWithDefaults() *MountVmwareVolumeParams`

NewMountVmwareVolumeParamsWithDefaults instantiates a new MountVmwareVolumeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetEnvironment

`func (o *MountVmwareVolumeParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *MountVmwareVolumeParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *MountVmwareVolumeParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.


### GetVmwareTargetParams

`func (o *MountVmwareVolumeParams) GetVmwareTargetParams() MountVmwareVolumeParamsVmwareTargetParams`

GetVmwareTargetParams returns the VmwareTargetParams field if non-nil, zero value otherwise.

### GetVmwareTargetParamsOk

`func (o *MountVmwareVolumeParams) GetVmwareTargetParamsOk() (*MountVmwareVolumeParamsVmwareTargetParams, bool)`

GetVmwareTargetParamsOk returns a tuple with the VmwareTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareTargetParams

`func (o *MountVmwareVolumeParams) SetVmwareTargetParams(v MountVmwareVolumeParamsVmwareTargetParams)`

SetVmwareTargetParams sets VmwareTargetParams field to given value.

### HasVmwareTargetParams

`func (o *MountVmwareVolumeParams) HasVmwareTargetParams() bool`

HasVmwareTargetParams returns a boolean if a field has been set.

### SetVmwareTargetParamsNil

`func (o *MountVmwareVolumeParams) SetVmwareTargetParamsNil(b bool)`

 SetVmwareTargetParamsNil sets the value for VmwareTargetParams to be an explicit nil

### UnsetVmwareTargetParams
`func (o *MountVmwareVolumeParams) UnsetVmwareTargetParams()`

UnsetVmwareTargetParams ensures that no value is present for VmwareTargetParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


