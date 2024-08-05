# RecoverNetappFilesParamsNetappTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewSourceConfig** | Pointer to [**NullableRecoverNetappToNetappFilesTargetParamsNewSourceConfig**](RecoverNetappToNetappFilesTargetParamsNewSourceConfig.md) |  | [optional] 
**OriginalSourceConfig** | Pointer to [**NullableRecoverNetappToNetappFilesTargetParamsOriginalSourceConfig**](RecoverNetappToNetappFilesTargetParamsOriginalSourceConfig.md) |  | [optional] 
**RecoverToNewSource** | **bool** | Specifies the parameter whether the recovery should be performed to a new or the original Netapp target. | 

## Methods

### NewRecoverNetappFilesParamsNetappTargetParams

`func NewRecoverNetappFilesParamsNetappTargetParams(recoverToNewSource bool, ) *RecoverNetappFilesParamsNetappTargetParams`

NewRecoverNetappFilesParamsNetappTargetParams instantiates a new RecoverNetappFilesParamsNetappTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverNetappFilesParamsNetappTargetParamsWithDefaults

`func NewRecoverNetappFilesParamsNetappTargetParamsWithDefaults() *RecoverNetappFilesParamsNetappTargetParams`

NewRecoverNetappFilesParamsNetappTargetParamsWithDefaults instantiates a new RecoverNetappFilesParamsNetappTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewSourceConfig

`func (o *RecoverNetappFilesParamsNetappTargetParams) GetNewSourceConfig() RecoverNetappToNetappFilesTargetParamsNewSourceConfig`

GetNewSourceConfig returns the NewSourceConfig field if non-nil, zero value otherwise.

### GetNewSourceConfigOk

`func (o *RecoverNetappFilesParamsNetappTargetParams) GetNewSourceConfigOk() (*RecoverNetappToNetappFilesTargetParamsNewSourceConfig, bool)`

GetNewSourceConfigOk returns a tuple with the NewSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSourceConfig

`func (o *RecoverNetappFilesParamsNetappTargetParams) SetNewSourceConfig(v RecoverNetappToNetappFilesTargetParamsNewSourceConfig)`

SetNewSourceConfig sets NewSourceConfig field to given value.

### HasNewSourceConfig

`func (o *RecoverNetappFilesParamsNetappTargetParams) HasNewSourceConfig() bool`

HasNewSourceConfig returns a boolean if a field has been set.

### SetNewSourceConfigNil

`func (o *RecoverNetappFilesParamsNetappTargetParams) SetNewSourceConfigNil(b bool)`

 SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil

### UnsetNewSourceConfig
`func (o *RecoverNetappFilesParamsNetappTargetParams) UnsetNewSourceConfig()`

UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
### GetOriginalSourceConfig

`func (o *RecoverNetappFilesParamsNetappTargetParams) GetOriginalSourceConfig() RecoverNetappToNetappFilesTargetParamsOriginalSourceConfig`

GetOriginalSourceConfig returns the OriginalSourceConfig field if non-nil, zero value otherwise.

### GetOriginalSourceConfigOk

`func (o *RecoverNetappFilesParamsNetappTargetParams) GetOriginalSourceConfigOk() (*RecoverNetappToNetappFilesTargetParamsOriginalSourceConfig, bool)`

GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSourceConfig

`func (o *RecoverNetappFilesParamsNetappTargetParams) SetOriginalSourceConfig(v RecoverNetappToNetappFilesTargetParamsOriginalSourceConfig)`

SetOriginalSourceConfig sets OriginalSourceConfig field to given value.

### HasOriginalSourceConfig

`func (o *RecoverNetappFilesParamsNetappTargetParams) HasOriginalSourceConfig() bool`

HasOriginalSourceConfig returns a boolean if a field has been set.

### SetOriginalSourceConfigNil

`func (o *RecoverNetappFilesParamsNetappTargetParams) SetOriginalSourceConfigNil(b bool)`

 SetOriginalSourceConfigNil sets the value for OriginalSourceConfig to be an explicit nil

### UnsetOriginalSourceConfig
`func (o *RecoverNetappFilesParamsNetappTargetParams) UnsetOriginalSourceConfig()`

UnsetOriginalSourceConfig ensures that no value is present for OriginalSourceConfig, not even an explicit nil
### GetRecoverToNewSource

`func (o *RecoverNetappFilesParamsNetappTargetParams) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *RecoverNetappFilesParamsNetappTargetParams) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *RecoverNetappFilesParamsNetappTargetParams) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


