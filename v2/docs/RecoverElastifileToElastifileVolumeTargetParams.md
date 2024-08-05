# RecoverElastifileToElastifileVolumeTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewSourceConfig** | Pointer to [**NullableRecoverElastifileToElastifileVolumeTargetParamsNewSourceConfig**](RecoverElastifileToElastifileVolumeTargetParamsNewSourceConfig.md) |  | [optional] 
**OriginalSourceConfig** | Pointer to [**NullableRecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig**](RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig.md) |  | [optional] 
**RecoverToNewSource** | **bool** | Specifies the parameter whether the recovery should be performed to a new or the original Elastifile target. | 

## Methods

### NewRecoverElastifileToElastifileVolumeTargetParams

`func NewRecoverElastifileToElastifileVolumeTargetParams(recoverToNewSource bool, ) *RecoverElastifileToElastifileVolumeTargetParams`

NewRecoverElastifileToElastifileVolumeTargetParams instantiates a new RecoverElastifileToElastifileVolumeTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverElastifileToElastifileVolumeTargetParamsWithDefaults

`func NewRecoverElastifileToElastifileVolumeTargetParamsWithDefaults() *RecoverElastifileToElastifileVolumeTargetParams`

NewRecoverElastifileToElastifileVolumeTargetParamsWithDefaults instantiates a new RecoverElastifileToElastifileVolumeTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewSourceConfig

`func (o *RecoverElastifileToElastifileVolumeTargetParams) GetNewSourceConfig() RecoverElastifileToElastifileVolumeTargetParamsNewSourceConfig`

GetNewSourceConfig returns the NewSourceConfig field if non-nil, zero value otherwise.

### GetNewSourceConfigOk

`func (o *RecoverElastifileToElastifileVolumeTargetParams) GetNewSourceConfigOk() (*RecoverElastifileToElastifileVolumeTargetParamsNewSourceConfig, bool)`

GetNewSourceConfigOk returns a tuple with the NewSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSourceConfig

`func (o *RecoverElastifileToElastifileVolumeTargetParams) SetNewSourceConfig(v RecoverElastifileToElastifileVolumeTargetParamsNewSourceConfig)`

SetNewSourceConfig sets NewSourceConfig field to given value.

### HasNewSourceConfig

`func (o *RecoverElastifileToElastifileVolumeTargetParams) HasNewSourceConfig() bool`

HasNewSourceConfig returns a boolean if a field has been set.

### SetNewSourceConfigNil

`func (o *RecoverElastifileToElastifileVolumeTargetParams) SetNewSourceConfigNil(b bool)`

 SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil

### UnsetNewSourceConfig
`func (o *RecoverElastifileToElastifileVolumeTargetParams) UnsetNewSourceConfig()`

UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
### GetOriginalSourceConfig

`func (o *RecoverElastifileToElastifileVolumeTargetParams) GetOriginalSourceConfig() RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig`

GetOriginalSourceConfig returns the OriginalSourceConfig field if non-nil, zero value otherwise.

### GetOriginalSourceConfigOk

`func (o *RecoverElastifileToElastifileVolumeTargetParams) GetOriginalSourceConfigOk() (*RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig, bool)`

GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSourceConfig

`func (o *RecoverElastifileToElastifileVolumeTargetParams) SetOriginalSourceConfig(v RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig)`

SetOriginalSourceConfig sets OriginalSourceConfig field to given value.

### HasOriginalSourceConfig

`func (o *RecoverElastifileToElastifileVolumeTargetParams) HasOriginalSourceConfig() bool`

HasOriginalSourceConfig returns a boolean if a field has been set.

### SetOriginalSourceConfigNil

`func (o *RecoverElastifileToElastifileVolumeTargetParams) SetOriginalSourceConfigNil(b bool)`

 SetOriginalSourceConfigNil sets the value for OriginalSourceConfig to be an explicit nil

### UnsetOriginalSourceConfig
`func (o *RecoverElastifileToElastifileVolumeTargetParams) UnsetOriginalSourceConfig()`

UnsetOriginalSourceConfig ensures that no value is present for OriginalSourceConfig, not even an explicit nil
### GetRecoverToNewSource

`func (o *RecoverElastifileToElastifileVolumeTargetParams) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *RecoverElastifileToElastifileVolumeTargetParams) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *RecoverElastifileToElastifileVolumeTargetParams) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


