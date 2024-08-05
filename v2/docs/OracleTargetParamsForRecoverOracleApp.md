# OracleTargetParamsForRecoverOracleApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewSourceConfig** | Pointer to [**NullableCommonRecoverOracleAppTargetParamsNewSourceConfig**](CommonRecoverOracleAppTargetParamsNewSourceConfig.md) |  | [optional] 
**OriginalSourceConfig** | Pointer to [**NullableCommonRecoverOracleAppTargetParamsOriginalSourceConfig**](CommonRecoverOracleAppTargetParamsOriginalSourceConfig.md) |  | [optional] 
**RecoverToNewSource** | **bool** | Specifies the parameter whether the recovery should be performed to a new source or an original Source Target. | 

## Methods

### NewOracleTargetParamsForRecoverOracleApp

`func NewOracleTargetParamsForRecoverOracleApp(recoverToNewSource bool, ) *OracleTargetParamsForRecoverOracleApp`

NewOracleTargetParamsForRecoverOracleApp instantiates a new OracleTargetParamsForRecoverOracleApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleTargetParamsForRecoverOracleAppWithDefaults

`func NewOracleTargetParamsForRecoverOracleAppWithDefaults() *OracleTargetParamsForRecoverOracleApp`

NewOracleTargetParamsForRecoverOracleAppWithDefaults instantiates a new OracleTargetParamsForRecoverOracleApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewSourceConfig

`func (o *OracleTargetParamsForRecoverOracleApp) GetNewSourceConfig() CommonRecoverOracleAppTargetParamsNewSourceConfig`

GetNewSourceConfig returns the NewSourceConfig field if non-nil, zero value otherwise.

### GetNewSourceConfigOk

`func (o *OracleTargetParamsForRecoverOracleApp) GetNewSourceConfigOk() (*CommonRecoverOracleAppTargetParamsNewSourceConfig, bool)`

GetNewSourceConfigOk returns a tuple with the NewSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSourceConfig

`func (o *OracleTargetParamsForRecoverOracleApp) SetNewSourceConfig(v CommonRecoverOracleAppTargetParamsNewSourceConfig)`

SetNewSourceConfig sets NewSourceConfig field to given value.

### HasNewSourceConfig

`func (o *OracleTargetParamsForRecoverOracleApp) HasNewSourceConfig() bool`

HasNewSourceConfig returns a boolean if a field has been set.

### SetNewSourceConfigNil

`func (o *OracleTargetParamsForRecoverOracleApp) SetNewSourceConfigNil(b bool)`

 SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil

### UnsetNewSourceConfig
`func (o *OracleTargetParamsForRecoverOracleApp) UnsetNewSourceConfig()`

UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
### GetOriginalSourceConfig

`func (o *OracleTargetParamsForRecoverOracleApp) GetOriginalSourceConfig() CommonRecoverOracleAppTargetParamsOriginalSourceConfig`

GetOriginalSourceConfig returns the OriginalSourceConfig field if non-nil, zero value otherwise.

### GetOriginalSourceConfigOk

`func (o *OracleTargetParamsForRecoverOracleApp) GetOriginalSourceConfigOk() (*CommonRecoverOracleAppTargetParamsOriginalSourceConfig, bool)`

GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSourceConfig

`func (o *OracleTargetParamsForRecoverOracleApp) SetOriginalSourceConfig(v CommonRecoverOracleAppTargetParamsOriginalSourceConfig)`

SetOriginalSourceConfig sets OriginalSourceConfig field to given value.

### HasOriginalSourceConfig

`func (o *OracleTargetParamsForRecoverOracleApp) HasOriginalSourceConfig() bool`

HasOriginalSourceConfig returns a boolean if a field has been set.

### SetOriginalSourceConfigNil

`func (o *OracleTargetParamsForRecoverOracleApp) SetOriginalSourceConfigNil(b bool)`

 SetOriginalSourceConfigNil sets the value for OriginalSourceConfig to be an explicit nil

### UnsetOriginalSourceConfig
`func (o *OracleTargetParamsForRecoverOracleApp) UnsetOriginalSourceConfig()`

UnsetOriginalSourceConfig ensures that no value is present for OriginalSourceConfig, not even an explicit nil
### GetRecoverToNewSource

`func (o *OracleTargetParamsForRecoverOracleApp) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *OracleTargetParamsForRecoverOracleApp) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *OracleTargetParamsForRecoverOracleApp) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


