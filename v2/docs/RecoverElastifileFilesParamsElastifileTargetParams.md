# RecoverElastifileFilesParamsElastifileTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewSourceConfig** | Pointer to [**NullableRecoverElastifileToElastifileFilesTargetParamsNewSourceConfig**](RecoverElastifileToElastifileFilesTargetParamsNewSourceConfig.md) |  | [optional] 
**OriginalSourceConfig** | Pointer to [**NullableRecoverElastifileToElastifileFilesTargetParamsOriginalSourceConfig**](RecoverElastifileToElastifileFilesTargetParamsOriginalSourceConfig.md) |  | [optional] 
**RecoverToNewSource** | **bool** | Specifies the parameter whether the recovery should be performed to a new or the original Elastifile target. | 

## Methods

### NewRecoverElastifileFilesParamsElastifileTargetParams

`func NewRecoverElastifileFilesParamsElastifileTargetParams(recoverToNewSource bool, ) *RecoverElastifileFilesParamsElastifileTargetParams`

NewRecoverElastifileFilesParamsElastifileTargetParams instantiates a new RecoverElastifileFilesParamsElastifileTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverElastifileFilesParamsElastifileTargetParamsWithDefaults

`func NewRecoverElastifileFilesParamsElastifileTargetParamsWithDefaults() *RecoverElastifileFilesParamsElastifileTargetParams`

NewRecoverElastifileFilesParamsElastifileTargetParamsWithDefaults instantiates a new RecoverElastifileFilesParamsElastifileTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewSourceConfig

`func (o *RecoverElastifileFilesParamsElastifileTargetParams) GetNewSourceConfig() RecoverElastifileToElastifileFilesTargetParamsNewSourceConfig`

GetNewSourceConfig returns the NewSourceConfig field if non-nil, zero value otherwise.

### GetNewSourceConfigOk

`func (o *RecoverElastifileFilesParamsElastifileTargetParams) GetNewSourceConfigOk() (*RecoverElastifileToElastifileFilesTargetParamsNewSourceConfig, bool)`

GetNewSourceConfigOk returns a tuple with the NewSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSourceConfig

`func (o *RecoverElastifileFilesParamsElastifileTargetParams) SetNewSourceConfig(v RecoverElastifileToElastifileFilesTargetParamsNewSourceConfig)`

SetNewSourceConfig sets NewSourceConfig field to given value.

### HasNewSourceConfig

`func (o *RecoverElastifileFilesParamsElastifileTargetParams) HasNewSourceConfig() bool`

HasNewSourceConfig returns a boolean if a field has been set.

### SetNewSourceConfigNil

`func (o *RecoverElastifileFilesParamsElastifileTargetParams) SetNewSourceConfigNil(b bool)`

 SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil

### UnsetNewSourceConfig
`func (o *RecoverElastifileFilesParamsElastifileTargetParams) UnsetNewSourceConfig()`

UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
### GetOriginalSourceConfig

`func (o *RecoverElastifileFilesParamsElastifileTargetParams) GetOriginalSourceConfig() RecoverElastifileToElastifileFilesTargetParamsOriginalSourceConfig`

GetOriginalSourceConfig returns the OriginalSourceConfig field if non-nil, zero value otherwise.

### GetOriginalSourceConfigOk

`func (o *RecoverElastifileFilesParamsElastifileTargetParams) GetOriginalSourceConfigOk() (*RecoverElastifileToElastifileFilesTargetParamsOriginalSourceConfig, bool)`

GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSourceConfig

`func (o *RecoverElastifileFilesParamsElastifileTargetParams) SetOriginalSourceConfig(v RecoverElastifileToElastifileFilesTargetParamsOriginalSourceConfig)`

SetOriginalSourceConfig sets OriginalSourceConfig field to given value.

### HasOriginalSourceConfig

`func (o *RecoverElastifileFilesParamsElastifileTargetParams) HasOriginalSourceConfig() bool`

HasOriginalSourceConfig returns a boolean if a field has been set.

### SetOriginalSourceConfigNil

`func (o *RecoverElastifileFilesParamsElastifileTargetParams) SetOriginalSourceConfigNil(b bool)`

 SetOriginalSourceConfigNil sets the value for OriginalSourceConfig to be an explicit nil

### UnsetOriginalSourceConfig
`func (o *RecoverElastifileFilesParamsElastifileTargetParams) UnsetOriginalSourceConfig()`

UnsetOriginalSourceConfig ensures that no value is present for OriginalSourceConfig, not even an explicit nil
### GetRecoverToNewSource

`func (o *RecoverElastifileFilesParamsElastifileTargetParams) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *RecoverElastifileFilesParamsElastifileTargetParams) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *RecoverElastifileFilesParamsElastifileTargetParams) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


