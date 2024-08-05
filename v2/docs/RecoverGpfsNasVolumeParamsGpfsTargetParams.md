# RecoverGpfsNasVolumeParamsGpfsTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewSourceConfig** | Pointer to [**NullableRecoverGpfsToGpfsVolumeTargetParamsNewSourceConfig**](RecoverGpfsToGpfsVolumeTargetParamsNewSourceConfig.md) |  | [optional] 
**OriginalSourceConfig** | Pointer to [**NullableRecoverGpfsToGpfsVolumeTargetParamsOriginalSourceConfig**](RecoverGpfsToGpfsVolumeTargetParamsOriginalSourceConfig.md) |  | [optional] 
**RecoverToNewSource** | **bool** | Specifies the parameter whether the recovery should be performed to a new or the original GPFS target. | 

## Methods

### NewRecoverGpfsNasVolumeParamsGpfsTargetParams

`func NewRecoverGpfsNasVolumeParamsGpfsTargetParams(recoverToNewSource bool, ) *RecoverGpfsNasVolumeParamsGpfsTargetParams`

NewRecoverGpfsNasVolumeParamsGpfsTargetParams instantiates a new RecoverGpfsNasVolumeParamsGpfsTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverGpfsNasVolumeParamsGpfsTargetParamsWithDefaults

`func NewRecoverGpfsNasVolumeParamsGpfsTargetParamsWithDefaults() *RecoverGpfsNasVolumeParamsGpfsTargetParams`

NewRecoverGpfsNasVolumeParamsGpfsTargetParamsWithDefaults instantiates a new RecoverGpfsNasVolumeParamsGpfsTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewSourceConfig

`func (o *RecoverGpfsNasVolumeParamsGpfsTargetParams) GetNewSourceConfig() RecoverGpfsToGpfsVolumeTargetParamsNewSourceConfig`

GetNewSourceConfig returns the NewSourceConfig field if non-nil, zero value otherwise.

### GetNewSourceConfigOk

`func (o *RecoverGpfsNasVolumeParamsGpfsTargetParams) GetNewSourceConfigOk() (*RecoverGpfsToGpfsVolumeTargetParamsNewSourceConfig, bool)`

GetNewSourceConfigOk returns a tuple with the NewSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSourceConfig

`func (o *RecoverGpfsNasVolumeParamsGpfsTargetParams) SetNewSourceConfig(v RecoverGpfsToGpfsVolumeTargetParamsNewSourceConfig)`

SetNewSourceConfig sets NewSourceConfig field to given value.

### HasNewSourceConfig

`func (o *RecoverGpfsNasVolumeParamsGpfsTargetParams) HasNewSourceConfig() bool`

HasNewSourceConfig returns a boolean if a field has been set.

### SetNewSourceConfigNil

`func (o *RecoverGpfsNasVolumeParamsGpfsTargetParams) SetNewSourceConfigNil(b bool)`

 SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil

### UnsetNewSourceConfig
`func (o *RecoverGpfsNasVolumeParamsGpfsTargetParams) UnsetNewSourceConfig()`

UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
### GetOriginalSourceConfig

`func (o *RecoverGpfsNasVolumeParamsGpfsTargetParams) GetOriginalSourceConfig() RecoverGpfsToGpfsVolumeTargetParamsOriginalSourceConfig`

GetOriginalSourceConfig returns the OriginalSourceConfig field if non-nil, zero value otherwise.

### GetOriginalSourceConfigOk

`func (o *RecoverGpfsNasVolumeParamsGpfsTargetParams) GetOriginalSourceConfigOk() (*RecoverGpfsToGpfsVolumeTargetParamsOriginalSourceConfig, bool)`

GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSourceConfig

`func (o *RecoverGpfsNasVolumeParamsGpfsTargetParams) SetOriginalSourceConfig(v RecoverGpfsToGpfsVolumeTargetParamsOriginalSourceConfig)`

SetOriginalSourceConfig sets OriginalSourceConfig field to given value.

### HasOriginalSourceConfig

`func (o *RecoverGpfsNasVolumeParamsGpfsTargetParams) HasOriginalSourceConfig() bool`

HasOriginalSourceConfig returns a boolean if a field has been set.

### SetOriginalSourceConfigNil

`func (o *RecoverGpfsNasVolumeParamsGpfsTargetParams) SetOriginalSourceConfigNil(b bool)`

 SetOriginalSourceConfigNil sets the value for OriginalSourceConfig to be an explicit nil

### UnsetOriginalSourceConfig
`func (o *RecoverGpfsNasVolumeParamsGpfsTargetParams) UnsetOriginalSourceConfig()`

UnsetOriginalSourceConfig ensures that no value is present for OriginalSourceConfig, not even an explicit nil
### GetRecoverToNewSource

`func (o *RecoverGpfsNasVolumeParamsGpfsTargetParams) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *RecoverGpfsNasVolumeParamsGpfsTargetParams) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *RecoverGpfsNasVolumeParamsGpfsTargetParams) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


