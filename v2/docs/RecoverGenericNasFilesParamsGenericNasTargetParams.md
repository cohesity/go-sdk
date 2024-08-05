# RecoverGenericNasFilesParamsGenericNasTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewSourceConfig** | Pointer to [**NullableRecoverGenericNasToGenericNasFilesTargetParamsNewSourceConfig**](RecoverGenericNasToGenericNasFilesTargetParamsNewSourceConfig.md) |  | [optional] 
**OriginalSourceConfig** | Pointer to [**NullableRecoverGenericNasToGenericNasFilesTargetParamsOriginalSourceConfig**](RecoverGenericNasToGenericNasFilesTargetParamsOriginalSourceConfig.md) |  | [optional] 
**RecoverToNewSource** | **bool** | Specifies the parameter whether the recovery should be performed to a new or the original Generic Nas target. | 

## Methods

### NewRecoverGenericNasFilesParamsGenericNasTargetParams

`func NewRecoverGenericNasFilesParamsGenericNasTargetParams(recoverToNewSource bool, ) *RecoverGenericNasFilesParamsGenericNasTargetParams`

NewRecoverGenericNasFilesParamsGenericNasTargetParams instantiates a new RecoverGenericNasFilesParamsGenericNasTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverGenericNasFilesParamsGenericNasTargetParamsWithDefaults

`func NewRecoverGenericNasFilesParamsGenericNasTargetParamsWithDefaults() *RecoverGenericNasFilesParamsGenericNasTargetParams`

NewRecoverGenericNasFilesParamsGenericNasTargetParamsWithDefaults instantiates a new RecoverGenericNasFilesParamsGenericNasTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewSourceConfig

`func (o *RecoverGenericNasFilesParamsGenericNasTargetParams) GetNewSourceConfig() RecoverGenericNasToGenericNasFilesTargetParamsNewSourceConfig`

GetNewSourceConfig returns the NewSourceConfig field if non-nil, zero value otherwise.

### GetNewSourceConfigOk

`func (o *RecoverGenericNasFilesParamsGenericNasTargetParams) GetNewSourceConfigOk() (*RecoverGenericNasToGenericNasFilesTargetParamsNewSourceConfig, bool)`

GetNewSourceConfigOk returns a tuple with the NewSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSourceConfig

`func (o *RecoverGenericNasFilesParamsGenericNasTargetParams) SetNewSourceConfig(v RecoverGenericNasToGenericNasFilesTargetParamsNewSourceConfig)`

SetNewSourceConfig sets NewSourceConfig field to given value.

### HasNewSourceConfig

`func (o *RecoverGenericNasFilesParamsGenericNasTargetParams) HasNewSourceConfig() bool`

HasNewSourceConfig returns a boolean if a field has been set.

### SetNewSourceConfigNil

`func (o *RecoverGenericNasFilesParamsGenericNasTargetParams) SetNewSourceConfigNil(b bool)`

 SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil

### UnsetNewSourceConfig
`func (o *RecoverGenericNasFilesParamsGenericNasTargetParams) UnsetNewSourceConfig()`

UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
### GetOriginalSourceConfig

`func (o *RecoverGenericNasFilesParamsGenericNasTargetParams) GetOriginalSourceConfig() RecoverGenericNasToGenericNasFilesTargetParamsOriginalSourceConfig`

GetOriginalSourceConfig returns the OriginalSourceConfig field if non-nil, zero value otherwise.

### GetOriginalSourceConfigOk

`func (o *RecoverGenericNasFilesParamsGenericNasTargetParams) GetOriginalSourceConfigOk() (*RecoverGenericNasToGenericNasFilesTargetParamsOriginalSourceConfig, bool)`

GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSourceConfig

`func (o *RecoverGenericNasFilesParamsGenericNasTargetParams) SetOriginalSourceConfig(v RecoverGenericNasToGenericNasFilesTargetParamsOriginalSourceConfig)`

SetOriginalSourceConfig sets OriginalSourceConfig field to given value.

### HasOriginalSourceConfig

`func (o *RecoverGenericNasFilesParamsGenericNasTargetParams) HasOriginalSourceConfig() bool`

HasOriginalSourceConfig returns a boolean if a field has been set.

### SetOriginalSourceConfigNil

`func (o *RecoverGenericNasFilesParamsGenericNasTargetParams) SetOriginalSourceConfigNil(b bool)`

 SetOriginalSourceConfigNil sets the value for OriginalSourceConfig to be an explicit nil

### UnsetOriginalSourceConfig
`func (o *RecoverGenericNasFilesParamsGenericNasTargetParams) UnsetOriginalSourceConfig()`

UnsetOriginalSourceConfig ensures that no value is present for OriginalSourceConfig, not even an explicit nil
### GetRecoverToNewSource

`func (o *RecoverGenericNasFilesParamsGenericNasTargetParams) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *RecoverGenericNasFilesParamsGenericNasTargetParams) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *RecoverGenericNasFilesParamsGenericNasTargetParams) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


