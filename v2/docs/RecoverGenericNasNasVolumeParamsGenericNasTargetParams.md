# RecoverGenericNasNasVolumeParamsGenericNasTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewSourceConfig** | Pointer to [**NullableRecoverGenericNasToGenericNasVolumeTargetParamsNewSourceConfig**](RecoverGenericNasToGenericNasVolumeTargetParamsNewSourceConfig.md) |  | [optional] 
**OriginalSourceConfig** | Pointer to [**NullableRecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig**](RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig.md) |  | [optional] 
**RecoverToNewSource** | **bool** | Specifies the parameter whether the recovery should be performed to a new or the original Generic NAS target. | 

## Methods

### NewRecoverGenericNasNasVolumeParamsGenericNasTargetParams

`func NewRecoverGenericNasNasVolumeParamsGenericNasTargetParams(recoverToNewSource bool, ) *RecoverGenericNasNasVolumeParamsGenericNasTargetParams`

NewRecoverGenericNasNasVolumeParamsGenericNasTargetParams instantiates a new RecoverGenericNasNasVolumeParamsGenericNasTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverGenericNasNasVolumeParamsGenericNasTargetParamsWithDefaults

`func NewRecoverGenericNasNasVolumeParamsGenericNasTargetParamsWithDefaults() *RecoverGenericNasNasVolumeParamsGenericNasTargetParams`

NewRecoverGenericNasNasVolumeParamsGenericNasTargetParamsWithDefaults instantiates a new RecoverGenericNasNasVolumeParamsGenericNasTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewSourceConfig

`func (o *RecoverGenericNasNasVolumeParamsGenericNasTargetParams) GetNewSourceConfig() RecoverGenericNasToGenericNasVolumeTargetParamsNewSourceConfig`

GetNewSourceConfig returns the NewSourceConfig field if non-nil, zero value otherwise.

### GetNewSourceConfigOk

`func (o *RecoverGenericNasNasVolumeParamsGenericNasTargetParams) GetNewSourceConfigOk() (*RecoverGenericNasToGenericNasVolumeTargetParamsNewSourceConfig, bool)`

GetNewSourceConfigOk returns a tuple with the NewSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSourceConfig

`func (o *RecoverGenericNasNasVolumeParamsGenericNasTargetParams) SetNewSourceConfig(v RecoverGenericNasToGenericNasVolumeTargetParamsNewSourceConfig)`

SetNewSourceConfig sets NewSourceConfig field to given value.

### HasNewSourceConfig

`func (o *RecoverGenericNasNasVolumeParamsGenericNasTargetParams) HasNewSourceConfig() bool`

HasNewSourceConfig returns a boolean if a field has been set.

### SetNewSourceConfigNil

`func (o *RecoverGenericNasNasVolumeParamsGenericNasTargetParams) SetNewSourceConfigNil(b bool)`

 SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil

### UnsetNewSourceConfig
`func (o *RecoverGenericNasNasVolumeParamsGenericNasTargetParams) UnsetNewSourceConfig()`

UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
### GetOriginalSourceConfig

`func (o *RecoverGenericNasNasVolumeParamsGenericNasTargetParams) GetOriginalSourceConfig() RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig`

GetOriginalSourceConfig returns the OriginalSourceConfig field if non-nil, zero value otherwise.

### GetOriginalSourceConfigOk

`func (o *RecoverGenericNasNasVolumeParamsGenericNasTargetParams) GetOriginalSourceConfigOk() (*RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig, bool)`

GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSourceConfig

`func (o *RecoverGenericNasNasVolumeParamsGenericNasTargetParams) SetOriginalSourceConfig(v RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig)`

SetOriginalSourceConfig sets OriginalSourceConfig field to given value.

### HasOriginalSourceConfig

`func (o *RecoverGenericNasNasVolumeParamsGenericNasTargetParams) HasOriginalSourceConfig() bool`

HasOriginalSourceConfig returns a boolean if a field has been set.

### SetOriginalSourceConfigNil

`func (o *RecoverGenericNasNasVolumeParamsGenericNasTargetParams) SetOriginalSourceConfigNil(b bool)`

 SetOriginalSourceConfigNil sets the value for OriginalSourceConfig to be an explicit nil

### UnsetOriginalSourceConfig
`func (o *RecoverGenericNasNasVolumeParamsGenericNasTargetParams) UnsetOriginalSourceConfig()`

UnsetOriginalSourceConfig ensures that no value is present for OriginalSourceConfig, not even an explicit nil
### GetRecoverToNewSource

`func (o *RecoverGenericNasNasVolumeParamsGenericNasTargetParams) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *RecoverGenericNasNasVolumeParamsGenericNasTargetParams) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *RecoverGenericNasNasVolumeParamsGenericNasTargetParams) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


