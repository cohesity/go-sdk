# RecoverIsilonNasVolumeParamsIsilonTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewSourceConfig** | Pointer to [**NullableRecoverIsilonToIsilonVolumeTargetParamsNewSourceConfig**](RecoverIsilonToIsilonVolumeTargetParamsNewSourceConfig.md) |  | [optional] 
**OriginalSourceConfig** | Pointer to [**NullableRecoverIsilonToIsilonVolumeTargetParamsOriginalSourceConfig**](RecoverIsilonToIsilonVolumeTargetParamsOriginalSourceConfig.md) |  | [optional] 
**RecoverToNewSource** | **bool** | Specifies the parameter whether the recovery should be performed to a new or the original Isilon target. | 

## Methods

### NewRecoverIsilonNasVolumeParamsIsilonTargetParams

`func NewRecoverIsilonNasVolumeParamsIsilonTargetParams(recoverToNewSource bool, ) *RecoverIsilonNasVolumeParamsIsilonTargetParams`

NewRecoverIsilonNasVolumeParamsIsilonTargetParams instantiates a new RecoverIsilonNasVolumeParamsIsilonTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverIsilonNasVolumeParamsIsilonTargetParamsWithDefaults

`func NewRecoverIsilonNasVolumeParamsIsilonTargetParamsWithDefaults() *RecoverIsilonNasVolumeParamsIsilonTargetParams`

NewRecoverIsilonNasVolumeParamsIsilonTargetParamsWithDefaults instantiates a new RecoverIsilonNasVolumeParamsIsilonTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewSourceConfig

`func (o *RecoverIsilonNasVolumeParamsIsilonTargetParams) GetNewSourceConfig() RecoverIsilonToIsilonVolumeTargetParamsNewSourceConfig`

GetNewSourceConfig returns the NewSourceConfig field if non-nil, zero value otherwise.

### GetNewSourceConfigOk

`func (o *RecoverIsilonNasVolumeParamsIsilonTargetParams) GetNewSourceConfigOk() (*RecoverIsilonToIsilonVolumeTargetParamsNewSourceConfig, bool)`

GetNewSourceConfigOk returns a tuple with the NewSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSourceConfig

`func (o *RecoverIsilonNasVolumeParamsIsilonTargetParams) SetNewSourceConfig(v RecoverIsilonToIsilonVolumeTargetParamsNewSourceConfig)`

SetNewSourceConfig sets NewSourceConfig field to given value.

### HasNewSourceConfig

`func (o *RecoverIsilonNasVolumeParamsIsilonTargetParams) HasNewSourceConfig() bool`

HasNewSourceConfig returns a boolean if a field has been set.

### SetNewSourceConfigNil

`func (o *RecoverIsilonNasVolumeParamsIsilonTargetParams) SetNewSourceConfigNil(b bool)`

 SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil

### UnsetNewSourceConfig
`func (o *RecoverIsilonNasVolumeParamsIsilonTargetParams) UnsetNewSourceConfig()`

UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
### GetOriginalSourceConfig

`func (o *RecoverIsilonNasVolumeParamsIsilonTargetParams) GetOriginalSourceConfig() RecoverIsilonToIsilonVolumeTargetParamsOriginalSourceConfig`

GetOriginalSourceConfig returns the OriginalSourceConfig field if non-nil, zero value otherwise.

### GetOriginalSourceConfigOk

`func (o *RecoverIsilonNasVolumeParamsIsilonTargetParams) GetOriginalSourceConfigOk() (*RecoverIsilonToIsilonVolumeTargetParamsOriginalSourceConfig, bool)`

GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSourceConfig

`func (o *RecoverIsilonNasVolumeParamsIsilonTargetParams) SetOriginalSourceConfig(v RecoverIsilonToIsilonVolumeTargetParamsOriginalSourceConfig)`

SetOriginalSourceConfig sets OriginalSourceConfig field to given value.

### HasOriginalSourceConfig

`func (o *RecoverIsilonNasVolumeParamsIsilonTargetParams) HasOriginalSourceConfig() bool`

HasOriginalSourceConfig returns a boolean if a field has been set.

### SetOriginalSourceConfigNil

`func (o *RecoverIsilonNasVolumeParamsIsilonTargetParams) SetOriginalSourceConfigNil(b bool)`

 SetOriginalSourceConfigNil sets the value for OriginalSourceConfig to be an explicit nil

### UnsetOriginalSourceConfig
`func (o *RecoverIsilonNasVolumeParamsIsilonTargetParams) UnsetOriginalSourceConfig()`

UnsetOriginalSourceConfig ensures that no value is present for OriginalSourceConfig, not even an explicit nil
### GetRecoverToNewSource

`func (o *RecoverIsilonNasVolumeParamsIsilonTargetParams) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *RecoverIsilonNasVolumeParamsIsilonTargetParams) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *RecoverIsilonNasVolumeParamsIsilonTargetParams) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


