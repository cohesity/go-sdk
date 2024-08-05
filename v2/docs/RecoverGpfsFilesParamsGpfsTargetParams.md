# RecoverGpfsFilesParamsGpfsTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewSourceConfig** | Pointer to [**NullableRecoverGpfsToGpfsFilesTargetParamsNewSourceConfig**](RecoverGpfsToGpfsFilesTargetParamsNewSourceConfig.md) |  | [optional] 
**OriginalSourceConfig** | Pointer to [**NullableRecoverGpfsToGpfsFilesTargetParamsOriginalSourceConfig**](RecoverGpfsToGpfsFilesTargetParamsOriginalSourceConfig.md) |  | [optional] 
**RecoverToNewSource** | **bool** | Specifies the parameter whether the recovery should be performed to a new or the original GPFS target. | 

## Methods

### NewRecoverGpfsFilesParamsGpfsTargetParams

`func NewRecoverGpfsFilesParamsGpfsTargetParams(recoverToNewSource bool, ) *RecoverGpfsFilesParamsGpfsTargetParams`

NewRecoverGpfsFilesParamsGpfsTargetParams instantiates a new RecoverGpfsFilesParamsGpfsTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverGpfsFilesParamsGpfsTargetParamsWithDefaults

`func NewRecoverGpfsFilesParamsGpfsTargetParamsWithDefaults() *RecoverGpfsFilesParamsGpfsTargetParams`

NewRecoverGpfsFilesParamsGpfsTargetParamsWithDefaults instantiates a new RecoverGpfsFilesParamsGpfsTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewSourceConfig

`func (o *RecoverGpfsFilesParamsGpfsTargetParams) GetNewSourceConfig() RecoverGpfsToGpfsFilesTargetParamsNewSourceConfig`

GetNewSourceConfig returns the NewSourceConfig field if non-nil, zero value otherwise.

### GetNewSourceConfigOk

`func (o *RecoverGpfsFilesParamsGpfsTargetParams) GetNewSourceConfigOk() (*RecoverGpfsToGpfsFilesTargetParamsNewSourceConfig, bool)`

GetNewSourceConfigOk returns a tuple with the NewSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSourceConfig

`func (o *RecoverGpfsFilesParamsGpfsTargetParams) SetNewSourceConfig(v RecoverGpfsToGpfsFilesTargetParamsNewSourceConfig)`

SetNewSourceConfig sets NewSourceConfig field to given value.

### HasNewSourceConfig

`func (o *RecoverGpfsFilesParamsGpfsTargetParams) HasNewSourceConfig() bool`

HasNewSourceConfig returns a boolean if a field has been set.

### SetNewSourceConfigNil

`func (o *RecoverGpfsFilesParamsGpfsTargetParams) SetNewSourceConfigNil(b bool)`

 SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil

### UnsetNewSourceConfig
`func (o *RecoverGpfsFilesParamsGpfsTargetParams) UnsetNewSourceConfig()`

UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
### GetOriginalSourceConfig

`func (o *RecoverGpfsFilesParamsGpfsTargetParams) GetOriginalSourceConfig() RecoverGpfsToGpfsFilesTargetParamsOriginalSourceConfig`

GetOriginalSourceConfig returns the OriginalSourceConfig field if non-nil, zero value otherwise.

### GetOriginalSourceConfigOk

`func (o *RecoverGpfsFilesParamsGpfsTargetParams) GetOriginalSourceConfigOk() (*RecoverGpfsToGpfsFilesTargetParamsOriginalSourceConfig, bool)`

GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSourceConfig

`func (o *RecoverGpfsFilesParamsGpfsTargetParams) SetOriginalSourceConfig(v RecoverGpfsToGpfsFilesTargetParamsOriginalSourceConfig)`

SetOriginalSourceConfig sets OriginalSourceConfig field to given value.

### HasOriginalSourceConfig

`func (o *RecoverGpfsFilesParamsGpfsTargetParams) HasOriginalSourceConfig() bool`

HasOriginalSourceConfig returns a boolean if a field has been set.

### SetOriginalSourceConfigNil

`func (o *RecoverGpfsFilesParamsGpfsTargetParams) SetOriginalSourceConfigNil(b bool)`

 SetOriginalSourceConfigNil sets the value for OriginalSourceConfig to be an explicit nil

### UnsetOriginalSourceConfig
`func (o *RecoverGpfsFilesParamsGpfsTargetParams) UnsetOriginalSourceConfig()`

UnsetOriginalSourceConfig ensures that no value is present for OriginalSourceConfig, not even an explicit nil
### GetRecoverToNewSource

`func (o *RecoverGpfsFilesParamsGpfsTargetParams) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *RecoverGpfsFilesParamsGpfsTargetParams) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *RecoverGpfsFilesParamsGpfsTargetParams) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


