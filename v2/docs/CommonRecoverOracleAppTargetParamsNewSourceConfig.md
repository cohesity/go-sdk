# CommonRecoverOracleAppTargetParamsNewSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | [**NullableRecoverOracleAppNewSourceConfigHost**](RecoverOracleAppNewSourceConfigHost.md) |  | 
**RecoverDatabaseParams** | Pointer to [**NullableRecoverOracleAppNewSourceConfigRecoverDatabaseParams**](RecoverOracleAppNewSourceConfigRecoverDatabaseParams.md) |  | [optional] 
**RecoverViewParams** | Pointer to [**NullableRecoverOracleAppNewSourceConfigRecoverViewParams**](RecoverOracleAppNewSourceConfigRecoverViewParams.md) |  | [optional] 
**RecoveryTarget** | Pointer to **NullableString** | Specifies if recovery target is a database or a view. | [optional] 

## Methods

### NewCommonRecoverOracleAppTargetParamsNewSourceConfig

`func NewCommonRecoverOracleAppTargetParamsNewSourceConfig(host NullableRecoverOracleAppNewSourceConfigHost, ) *CommonRecoverOracleAppTargetParamsNewSourceConfig`

NewCommonRecoverOracleAppTargetParamsNewSourceConfig instantiates a new CommonRecoverOracleAppTargetParamsNewSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonRecoverOracleAppTargetParamsNewSourceConfigWithDefaults

`func NewCommonRecoverOracleAppTargetParamsNewSourceConfigWithDefaults() *CommonRecoverOracleAppTargetParamsNewSourceConfig`

NewCommonRecoverOracleAppTargetParamsNewSourceConfigWithDefaults instantiates a new CommonRecoverOracleAppTargetParamsNewSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *CommonRecoverOracleAppTargetParamsNewSourceConfig) GetHost() RecoverOracleAppNewSourceConfigHost`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CommonRecoverOracleAppTargetParamsNewSourceConfig) GetHostOk() (*RecoverOracleAppNewSourceConfigHost, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CommonRecoverOracleAppTargetParamsNewSourceConfig) SetHost(v RecoverOracleAppNewSourceConfigHost)`

SetHost sets Host field to given value.


### SetHostNil

`func (o *CommonRecoverOracleAppTargetParamsNewSourceConfig) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *CommonRecoverOracleAppTargetParamsNewSourceConfig) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetRecoverDatabaseParams

`func (o *CommonRecoverOracleAppTargetParamsNewSourceConfig) GetRecoverDatabaseParams() RecoverOracleAppNewSourceConfigRecoverDatabaseParams`

GetRecoverDatabaseParams returns the RecoverDatabaseParams field if non-nil, zero value otherwise.

### GetRecoverDatabaseParamsOk

`func (o *CommonRecoverOracleAppTargetParamsNewSourceConfig) GetRecoverDatabaseParamsOk() (*RecoverOracleAppNewSourceConfigRecoverDatabaseParams, bool)`

GetRecoverDatabaseParamsOk returns a tuple with the RecoverDatabaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverDatabaseParams

`func (o *CommonRecoverOracleAppTargetParamsNewSourceConfig) SetRecoverDatabaseParams(v RecoverOracleAppNewSourceConfigRecoverDatabaseParams)`

SetRecoverDatabaseParams sets RecoverDatabaseParams field to given value.

### HasRecoverDatabaseParams

`func (o *CommonRecoverOracleAppTargetParamsNewSourceConfig) HasRecoverDatabaseParams() bool`

HasRecoverDatabaseParams returns a boolean if a field has been set.

### SetRecoverDatabaseParamsNil

`func (o *CommonRecoverOracleAppTargetParamsNewSourceConfig) SetRecoverDatabaseParamsNil(b bool)`

 SetRecoverDatabaseParamsNil sets the value for RecoverDatabaseParams to be an explicit nil

### UnsetRecoverDatabaseParams
`func (o *CommonRecoverOracleAppTargetParamsNewSourceConfig) UnsetRecoverDatabaseParams()`

UnsetRecoverDatabaseParams ensures that no value is present for RecoverDatabaseParams, not even an explicit nil
### GetRecoverViewParams

`func (o *CommonRecoverOracleAppTargetParamsNewSourceConfig) GetRecoverViewParams() RecoverOracleAppNewSourceConfigRecoverViewParams`

GetRecoverViewParams returns the RecoverViewParams field if non-nil, zero value otherwise.

### GetRecoverViewParamsOk

`func (o *CommonRecoverOracleAppTargetParamsNewSourceConfig) GetRecoverViewParamsOk() (*RecoverOracleAppNewSourceConfigRecoverViewParams, bool)`

GetRecoverViewParamsOk returns a tuple with the RecoverViewParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverViewParams

`func (o *CommonRecoverOracleAppTargetParamsNewSourceConfig) SetRecoverViewParams(v RecoverOracleAppNewSourceConfigRecoverViewParams)`

SetRecoverViewParams sets RecoverViewParams field to given value.

### HasRecoverViewParams

`func (o *CommonRecoverOracleAppTargetParamsNewSourceConfig) HasRecoverViewParams() bool`

HasRecoverViewParams returns a boolean if a field has been set.

### SetRecoverViewParamsNil

`func (o *CommonRecoverOracleAppTargetParamsNewSourceConfig) SetRecoverViewParamsNil(b bool)`

 SetRecoverViewParamsNil sets the value for RecoverViewParams to be an explicit nil

### UnsetRecoverViewParams
`func (o *CommonRecoverOracleAppTargetParamsNewSourceConfig) UnsetRecoverViewParams()`

UnsetRecoverViewParams ensures that no value is present for RecoverViewParams, not even an explicit nil
### GetRecoveryTarget

`func (o *CommonRecoverOracleAppTargetParamsNewSourceConfig) GetRecoveryTarget() string`

GetRecoveryTarget returns the RecoveryTarget field if non-nil, zero value otherwise.

### GetRecoveryTargetOk

`func (o *CommonRecoverOracleAppTargetParamsNewSourceConfig) GetRecoveryTargetOk() (*string, bool)`

GetRecoveryTargetOk returns a tuple with the RecoveryTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTarget

`func (o *CommonRecoverOracleAppTargetParamsNewSourceConfig) SetRecoveryTarget(v string)`

SetRecoveryTarget sets RecoveryTarget field to given value.

### HasRecoveryTarget

`func (o *CommonRecoverOracleAppTargetParamsNewSourceConfig) HasRecoveryTarget() bool`

HasRecoveryTarget returns a boolean if a field has been set.

### SetRecoveryTargetNil

`func (o *CommonRecoverOracleAppTargetParamsNewSourceConfig) SetRecoveryTargetNil(b bool)`

 SetRecoveryTargetNil sets the value for RecoveryTarget to be an explicit nil

### UnsetRecoveryTarget
`func (o *CommonRecoverOracleAppTargetParamsNewSourceConfig) UnsetRecoveryTarget()`

UnsetRecoveryTarget ensures that no value is present for RecoveryTarget, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


