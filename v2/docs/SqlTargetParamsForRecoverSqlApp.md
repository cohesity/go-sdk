# SqlTargetParamsForRecoverSqlApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewSourceConfig** | Pointer to [**NullableCommonRecoverSqlAppTargetParamsNewSourceConfig**](CommonRecoverSqlAppTargetParamsNewSourceConfig.md) |  | [optional] 
**OriginalSourceConfig** | Pointer to [**NullableCommonRecoverSqlAppTargetParamsOriginalSourceConfig**](CommonRecoverSqlAppTargetParamsOriginalSourceConfig.md) |  | [optional] 
**RecoverToNewSource** | **bool** | Specifies the parameter whether the recovery should be performed to a new sources or an original Source Target. | 

## Methods

### NewSqlTargetParamsForRecoverSqlApp

`func NewSqlTargetParamsForRecoverSqlApp(recoverToNewSource bool, ) *SqlTargetParamsForRecoverSqlApp`

NewSqlTargetParamsForRecoverSqlApp instantiates a new SqlTargetParamsForRecoverSqlApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlTargetParamsForRecoverSqlAppWithDefaults

`func NewSqlTargetParamsForRecoverSqlAppWithDefaults() *SqlTargetParamsForRecoverSqlApp`

NewSqlTargetParamsForRecoverSqlAppWithDefaults instantiates a new SqlTargetParamsForRecoverSqlApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewSourceConfig

`func (o *SqlTargetParamsForRecoverSqlApp) GetNewSourceConfig() CommonRecoverSqlAppTargetParamsNewSourceConfig`

GetNewSourceConfig returns the NewSourceConfig field if non-nil, zero value otherwise.

### GetNewSourceConfigOk

`func (o *SqlTargetParamsForRecoverSqlApp) GetNewSourceConfigOk() (*CommonRecoverSqlAppTargetParamsNewSourceConfig, bool)`

GetNewSourceConfigOk returns a tuple with the NewSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSourceConfig

`func (o *SqlTargetParamsForRecoverSqlApp) SetNewSourceConfig(v CommonRecoverSqlAppTargetParamsNewSourceConfig)`

SetNewSourceConfig sets NewSourceConfig field to given value.

### HasNewSourceConfig

`func (o *SqlTargetParamsForRecoverSqlApp) HasNewSourceConfig() bool`

HasNewSourceConfig returns a boolean if a field has been set.

### SetNewSourceConfigNil

`func (o *SqlTargetParamsForRecoverSqlApp) SetNewSourceConfigNil(b bool)`

 SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil

### UnsetNewSourceConfig
`func (o *SqlTargetParamsForRecoverSqlApp) UnsetNewSourceConfig()`

UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
### GetOriginalSourceConfig

`func (o *SqlTargetParamsForRecoverSqlApp) GetOriginalSourceConfig() CommonRecoverSqlAppTargetParamsOriginalSourceConfig`

GetOriginalSourceConfig returns the OriginalSourceConfig field if non-nil, zero value otherwise.

### GetOriginalSourceConfigOk

`func (o *SqlTargetParamsForRecoverSqlApp) GetOriginalSourceConfigOk() (*CommonRecoverSqlAppTargetParamsOriginalSourceConfig, bool)`

GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSourceConfig

`func (o *SqlTargetParamsForRecoverSqlApp) SetOriginalSourceConfig(v CommonRecoverSqlAppTargetParamsOriginalSourceConfig)`

SetOriginalSourceConfig sets OriginalSourceConfig field to given value.

### HasOriginalSourceConfig

`func (o *SqlTargetParamsForRecoverSqlApp) HasOriginalSourceConfig() bool`

HasOriginalSourceConfig returns a boolean if a field has been set.

### SetOriginalSourceConfigNil

`func (o *SqlTargetParamsForRecoverSqlApp) SetOriginalSourceConfigNil(b bool)`

 SetOriginalSourceConfigNil sets the value for OriginalSourceConfig to be an explicit nil

### UnsetOriginalSourceConfig
`func (o *SqlTargetParamsForRecoverSqlApp) UnsetOriginalSourceConfig()`

UnsetOriginalSourceConfig ensures that no value is present for OriginalSourceConfig, not even an explicit nil
### GetRecoverToNewSource

`func (o *SqlTargetParamsForRecoverSqlApp) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *SqlTargetParamsForRecoverSqlApp) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *SqlTargetParamsForRecoverSqlApp) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


