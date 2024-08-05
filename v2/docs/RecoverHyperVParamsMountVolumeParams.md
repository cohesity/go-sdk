# RecoverHyperVParamsMountVolumeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HypervTargetParams** | Pointer to [**NullableMountHyperVVolumeParamsHypervTargetParams**](MountHyperVVolumeParamsHypervTargetParams.md) |  | [optional] 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 

## Methods

### NewRecoverHyperVParamsMountVolumeParams

`func NewRecoverHyperVParamsMountVolumeParams(targetEnvironment string, ) *RecoverHyperVParamsMountVolumeParams`

NewRecoverHyperVParamsMountVolumeParams instantiates a new RecoverHyperVParamsMountVolumeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverHyperVParamsMountVolumeParamsWithDefaults

`func NewRecoverHyperVParamsMountVolumeParamsWithDefaults() *RecoverHyperVParamsMountVolumeParams`

NewRecoverHyperVParamsMountVolumeParamsWithDefaults instantiates a new RecoverHyperVParamsMountVolumeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHypervTargetParams

`func (o *RecoverHyperVParamsMountVolumeParams) GetHypervTargetParams() MountHyperVVolumeParamsHypervTargetParams`

GetHypervTargetParams returns the HypervTargetParams field if non-nil, zero value otherwise.

### GetHypervTargetParamsOk

`func (o *RecoverHyperVParamsMountVolumeParams) GetHypervTargetParamsOk() (*MountHyperVVolumeParamsHypervTargetParams, bool)`

GetHypervTargetParamsOk returns a tuple with the HypervTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervTargetParams

`func (o *RecoverHyperVParamsMountVolumeParams) SetHypervTargetParams(v MountHyperVVolumeParamsHypervTargetParams)`

SetHypervTargetParams sets HypervTargetParams field to given value.

### HasHypervTargetParams

`func (o *RecoverHyperVParamsMountVolumeParams) HasHypervTargetParams() bool`

HasHypervTargetParams returns a boolean if a field has been set.

### SetHypervTargetParamsNil

`func (o *RecoverHyperVParamsMountVolumeParams) SetHypervTargetParamsNil(b bool)`

 SetHypervTargetParamsNil sets the value for HypervTargetParams to be an explicit nil

### UnsetHypervTargetParams
`func (o *RecoverHyperVParamsMountVolumeParams) UnsetHypervTargetParams()`

UnsetHypervTargetParams ensures that no value is present for HypervTargetParams, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverHyperVParamsMountVolumeParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverHyperVParamsMountVolumeParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverHyperVParamsMountVolumeParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


