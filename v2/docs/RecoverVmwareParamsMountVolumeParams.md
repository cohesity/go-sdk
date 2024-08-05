# RecoverVmwareParamsMountVolumeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 
**VmwareTargetParams** | Pointer to [**NullableMountVmwareVolumeParamsVmwareTargetParams**](MountVmwareVolumeParamsVmwareTargetParams.md) |  | [optional] 

## Methods

### NewRecoverVmwareParamsMountVolumeParams

`func NewRecoverVmwareParamsMountVolumeParams(targetEnvironment string, ) *RecoverVmwareParamsMountVolumeParams`

NewRecoverVmwareParamsMountVolumeParams instantiates a new RecoverVmwareParamsMountVolumeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverVmwareParamsMountVolumeParamsWithDefaults

`func NewRecoverVmwareParamsMountVolumeParamsWithDefaults() *RecoverVmwareParamsMountVolumeParams`

NewRecoverVmwareParamsMountVolumeParamsWithDefaults instantiates a new RecoverVmwareParamsMountVolumeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetEnvironment

`func (o *RecoverVmwareParamsMountVolumeParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverVmwareParamsMountVolumeParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverVmwareParamsMountVolumeParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.


### GetVmwareTargetParams

`func (o *RecoverVmwareParamsMountVolumeParams) GetVmwareTargetParams() MountVmwareVolumeParamsVmwareTargetParams`

GetVmwareTargetParams returns the VmwareTargetParams field if non-nil, zero value otherwise.

### GetVmwareTargetParamsOk

`func (o *RecoverVmwareParamsMountVolumeParams) GetVmwareTargetParamsOk() (*MountVmwareVolumeParamsVmwareTargetParams, bool)`

GetVmwareTargetParamsOk returns a tuple with the VmwareTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareTargetParams

`func (o *RecoverVmwareParamsMountVolumeParams) SetVmwareTargetParams(v MountVmwareVolumeParamsVmwareTargetParams)`

SetVmwareTargetParams sets VmwareTargetParams field to given value.

### HasVmwareTargetParams

`func (o *RecoverVmwareParamsMountVolumeParams) HasVmwareTargetParams() bool`

HasVmwareTargetParams returns a boolean if a field has been set.

### SetVmwareTargetParamsNil

`func (o *RecoverVmwareParamsMountVolumeParams) SetVmwareTargetParamsNil(b bool)`

 SetVmwareTargetParamsNil sets the value for VmwareTargetParams to be an explicit nil

### UnsetVmwareTargetParams
`func (o *RecoverVmwareParamsMountVolumeParams) UnsetVmwareTargetParams()`

UnsetVmwareTargetParams ensures that no value is present for VmwareTargetParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


