# RecoverFlashbladeNasVolumeParamsFlashbladeTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewSourceConfig** | Pointer to [**NullableRecoverFlashbladeToFlashbladeVolumeTargetParamsNewSourceConfig**](RecoverFlashbladeToFlashbladeVolumeTargetParamsNewSourceConfig.md) |  | [optional] 
**OriginalSourceConfig** | Pointer to [**NullableRecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig**](RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig.md) |  | [optional] 
**RecoverToNewSource** | **bool** | Specifies the parameter whether the recovery should be performed to a new or the original Flashblade target. | 

## Methods

### NewRecoverFlashbladeNasVolumeParamsFlashbladeTargetParams

`func NewRecoverFlashbladeNasVolumeParamsFlashbladeTargetParams(recoverToNewSource bool, ) *RecoverFlashbladeNasVolumeParamsFlashbladeTargetParams`

NewRecoverFlashbladeNasVolumeParamsFlashbladeTargetParams instantiates a new RecoverFlashbladeNasVolumeParamsFlashbladeTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverFlashbladeNasVolumeParamsFlashbladeTargetParamsWithDefaults

`func NewRecoverFlashbladeNasVolumeParamsFlashbladeTargetParamsWithDefaults() *RecoverFlashbladeNasVolumeParamsFlashbladeTargetParams`

NewRecoverFlashbladeNasVolumeParamsFlashbladeTargetParamsWithDefaults instantiates a new RecoverFlashbladeNasVolumeParamsFlashbladeTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewSourceConfig

`func (o *RecoverFlashbladeNasVolumeParamsFlashbladeTargetParams) GetNewSourceConfig() RecoverFlashbladeToFlashbladeVolumeTargetParamsNewSourceConfig`

GetNewSourceConfig returns the NewSourceConfig field if non-nil, zero value otherwise.

### GetNewSourceConfigOk

`func (o *RecoverFlashbladeNasVolumeParamsFlashbladeTargetParams) GetNewSourceConfigOk() (*RecoverFlashbladeToFlashbladeVolumeTargetParamsNewSourceConfig, bool)`

GetNewSourceConfigOk returns a tuple with the NewSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSourceConfig

`func (o *RecoverFlashbladeNasVolumeParamsFlashbladeTargetParams) SetNewSourceConfig(v RecoverFlashbladeToFlashbladeVolumeTargetParamsNewSourceConfig)`

SetNewSourceConfig sets NewSourceConfig field to given value.

### HasNewSourceConfig

`func (o *RecoverFlashbladeNasVolumeParamsFlashbladeTargetParams) HasNewSourceConfig() bool`

HasNewSourceConfig returns a boolean if a field has been set.

### SetNewSourceConfigNil

`func (o *RecoverFlashbladeNasVolumeParamsFlashbladeTargetParams) SetNewSourceConfigNil(b bool)`

 SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil

### UnsetNewSourceConfig
`func (o *RecoverFlashbladeNasVolumeParamsFlashbladeTargetParams) UnsetNewSourceConfig()`

UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
### GetOriginalSourceConfig

`func (o *RecoverFlashbladeNasVolumeParamsFlashbladeTargetParams) GetOriginalSourceConfig() RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig`

GetOriginalSourceConfig returns the OriginalSourceConfig field if non-nil, zero value otherwise.

### GetOriginalSourceConfigOk

`func (o *RecoverFlashbladeNasVolumeParamsFlashbladeTargetParams) GetOriginalSourceConfigOk() (*RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig, bool)`

GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSourceConfig

`func (o *RecoverFlashbladeNasVolumeParamsFlashbladeTargetParams) SetOriginalSourceConfig(v RecoverFlashbladeToFlashbladeVolumeTargetParamsOriginalSourceConfig)`

SetOriginalSourceConfig sets OriginalSourceConfig field to given value.

### HasOriginalSourceConfig

`func (o *RecoverFlashbladeNasVolumeParamsFlashbladeTargetParams) HasOriginalSourceConfig() bool`

HasOriginalSourceConfig returns a boolean if a field has been set.

### SetOriginalSourceConfigNil

`func (o *RecoverFlashbladeNasVolumeParamsFlashbladeTargetParams) SetOriginalSourceConfigNil(b bool)`

 SetOriginalSourceConfigNil sets the value for OriginalSourceConfig to be an explicit nil

### UnsetOriginalSourceConfig
`func (o *RecoverFlashbladeNasVolumeParamsFlashbladeTargetParams) UnsetOriginalSourceConfig()`

UnsetOriginalSourceConfig ensures that no value is present for OriginalSourceConfig, not even an explicit nil
### GetRecoverToNewSource

`func (o *RecoverFlashbladeNasVolumeParamsFlashbladeTargetParams) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *RecoverFlashbladeNasVolumeParamsFlashbladeTargetParams) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *RecoverFlashbladeNasVolumeParamsFlashbladeTargetParams) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


