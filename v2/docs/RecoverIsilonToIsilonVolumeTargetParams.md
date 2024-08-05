# RecoverIsilonToIsilonVolumeTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewSourceConfig** | Pointer to [**NullableRecoverIsilonToIsilonVolumeTargetParamsNewSourceConfig**](RecoverIsilonToIsilonVolumeTargetParamsNewSourceConfig.md) |  | [optional] 
**OriginalSourceConfig** | Pointer to [**NullableRecoverIsilonToIsilonVolumeTargetParamsOriginalSourceConfig**](RecoverIsilonToIsilonVolumeTargetParamsOriginalSourceConfig.md) |  | [optional] 
**RecoverToNewSource** | **bool** | Specifies the parameter whether the recovery should be performed to a new or the original Isilon target. | 

## Methods

### NewRecoverIsilonToIsilonVolumeTargetParams

`func NewRecoverIsilonToIsilonVolumeTargetParams(recoverToNewSource bool, ) *RecoverIsilonToIsilonVolumeTargetParams`

NewRecoverIsilonToIsilonVolumeTargetParams instantiates a new RecoverIsilonToIsilonVolumeTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverIsilonToIsilonVolumeTargetParamsWithDefaults

`func NewRecoverIsilonToIsilonVolumeTargetParamsWithDefaults() *RecoverIsilonToIsilonVolumeTargetParams`

NewRecoverIsilonToIsilonVolumeTargetParamsWithDefaults instantiates a new RecoverIsilonToIsilonVolumeTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewSourceConfig

`func (o *RecoverIsilonToIsilonVolumeTargetParams) GetNewSourceConfig() RecoverIsilonToIsilonVolumeTargetParamsNewSourceConfig`

GetNewSourceConfig returns the NewSourceConfig field if non-nil, zero value otherwise.

### GetNewSourceConfigOk

`func (o *RecoverIsilonToIsilonVolumeTargetParams) GetNewSourceConfigOk() (*RecoverIsilonToIsilonVolumeTargetParamsNewSourceConfig, bool)`

GetNewSourceConfigOk returns a tuple with the NewSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSourceConfig

`func (o *RecoverIsilonToIsilonVolumeTargetParams) SetNewSourceConfig(v RecoverIsilonToIsilonVolumeTargetParamsNewSourceConfig)`

SetNewSourceConfig sets NewSourceConfig field to given value.

### HasNewSourceConfig

`func (o *RecoverIsilonToIsilonVolumeTargetParams) HasNewSourceConfig() bool`

HasNewSourceConfig returns a boolean if a field has been set.

### SetNewSourceConfigNil

`func (o *RecoverIsilonToIsilonVolumeTargetParams) SetNewSourceConfigNil(b bool)`

 SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil

### UnsetNewSourceConfig
`func (o *RecoverIsilonToIsilonVolumeTargetParams) UnsetNewSourceConfig()`

UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
### GetOriginalSourceConfig

`func (o *RecoverIsilonToIsilonVolumeTargetParams) GetOriginalSourceConfig() RecoverIsilonToIsilonVolumeTargetParamsOriginalSourceConfig`

GetOriginalSourceConfig returns the OriginalSourceConfig field if non-nil, zero value otherwise.

### GetOriginalSourceConfigOk

`func (o *RecoverIsilonToIsilonVolumeTargetParams) GetOriginalSourceConfigOk() (*RecoverIsilonToIsilonVolumeTargetParamsOriginalSourceConfig, bool)`

GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSourceConfig

`func (o *RecoverIsilonToIsilonVolumeTargetParams) SetOriginalSourceConfig(v RecoverIsilonToIsilonVolumeTargetParamsOriginalSourceConfig)`

SetOriginalSourceConfig sets OriginalSourceConfig field to given value.

### HasOriginalSourceConfig

`func (o *RecoverIsilonToIsilonVolumeTargetParams) HasOriginalSourceConfig() bool`

HasOriginalSourceConfig returns a boolean if a field has been set.

### SetOriginalSourceConfigNil

`func (o *RecoverIsilonToIsilonVolumeTargetParams) SetOriginalSourceConfigNil(b bool)`

 SetOriginalSourceConfigNil sets the value for OriginalSourceConfig to be an explicit nil

### UnsetOriginalSourceConfig
`func (o *RecoverIsilonToIsilonVolumeTargetParams) UnsetOriginalSourceConfig()`

UnsetOriginalSourceConfig ensures that no value is present for OriginalSourceConfig, not even an explicit nil
### GetRecoverToNewSource

`func (o *RecoverIsilonToIsilonVolumeTargetParams) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *RecoverIsilonToIsilonVolumeTargetParams) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *RecoverIsilonToIsilonVolumeTargetParams) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


