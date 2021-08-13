# RecoverPureSanVolumeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding target params must be filled out. | 
**PureTargetParams** | Pointer to [**NullableRecoverPureVolumeTargetParams**](RecoverPureVolumeTargetParams.md) | Specifies the parameters of the Pure SAN volume to recover to. | [optional] 

## Methods

### NewRecoverPureSanVolumeParams

`func NewRecoverPureSanVolumeParams(targetEnvironment string, ) *RecoverPureSanVolumeParams`

NewRecoverPureSanVolumeParams instantiates a new RecoverPureSanVolumeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverPureSanVolumeParamsWithDefaults

`func NewRecoverPureSanVolumeParamsWithDefaults() *RecoverPureSanVolumeParams`

NewRecoverPureSanVolumeParamsWithDefaults instantiates a new RecoverPureSanVolumeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetEnvironment

`func (o *RecoverPureSanVolumeParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverPureSanVolumeParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverPureSanVolumeParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.


### GetPureTargetParams

`func (o *RecoverPureSanVolumeParams) GetPureTargetParams() RecoverPureVolumeTargetParams`

GetPureTargetParams returns the PureTargetParams field if non-nil, zero value otherwise.

### GetPureTargetParamsOk

`func (o *RecoverPureSanVolumeParams) GetPureTargetParamsOk() (*RecoverPureVolumeTargetParams, bool)`

GetPureTargetParamsOk returns a tuple with the PureTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPureTargetParams

`func (o *RecoverPureSanVolumeParams) SetPureTargetParams(v RecoverPureVolumeTargetParams)`

SetPureTargetParams sets PureTargetParams field to given value.

### HasPureTargetParams

`func (o *RecoverPureSanVolumeParams) HasPureTargetParams() bool`

HasPureTargetParams returns a boolean if a field has been set.

### SetPureTargetParamsNil

`func (o *RecoverPureSanVolumeParams) SetPureTargetParamsNil(b bool)`

 SetPureTargetParamsNil sets the value for PureTargetParams to be an explicit nil

### UnsetPureTargetParams
`func (o *RecoverPureSanVolumeParams) UnsetPureTargetParams()`

UnsetPureTargetParams ensures that no value is present for PureTargetParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


