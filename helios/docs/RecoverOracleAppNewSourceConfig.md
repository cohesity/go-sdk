# RecoverOracleAppNewSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the source id of target host where databases will be recovered. This source id can be a physical host or virtual machine. | 
**RecoveryTarget** | Pointer to **NullableString** | Specifies if recovery target is a database or a view. | [optional] 
**RecoverDatabaseParams** | Pointer to [**NullableRecoverOracleNewTargetDatabaseConfig**](RecoverOracleNewTargetDatabaseConfig.md) | Specifies recovery parameters when recovering to a database | [optional] 
**RecoverViewParams** | Pointer to [**NullableRecoverOracleNewTargetViewConfig**](RecoverOracleNewTargetViewConfig.md) | Specifies recovery parameters when recovering to a view. | [optional] 

## Methods

### NewRecoverOracleAppNewSourceConfig

`func NewRecoverOracleAppNewSourceConfig(host NullableRecoveryObjectIdentifier, ) *RecoverOracleAppNewSourceConfig`

NewRecoverOracleAppNewSourceConfig instantiates a new RecoverOracleAppNewSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverOracleAppNewSourceConfigWithDefaults

`func NewRecoverOracleAppNewSourceConfigWithDefaults() *RecoverOracleAppNewSourceConfig`

NewRecoverOracleAppNewSourceConfigWithDefaults instantiates a new RecoverOracleAppNewSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *RecoverOracleAppNewSourceConfig) GetHost() RecoveryObjectIdentifier`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RecoverOracleAppNewSourceConfig) GetHostOk() (*RecoveryObjectIdentifier, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RecoverOracleAppNewSourceConfig) SetHost(v RecoveryObjectIdentifier)`

SetHost sets Host field to given value.


### SetHostNil

`func (o *RecoverOracleAppNewSourceConfig) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *RecoverOracleAppNewSourceConfig) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetRecoveryTarget

`func (o *RecoverOracleAppNewSourceConfig) GetRecoveryTarget() string`

GetRecoveryTarget returns the RecoveryTarget field if non-nil, zero value otherwise.

### GetRecoveryTargetOk

`func (o *RecoverOracleAppNewSourceConfig) GetRecoveryTargetOk() (*string, bool)`

GetRecoveryTargetOk returns a tuple with the RecoveryTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTarget

`func (o *RecoverOracleAppNewSourceConfig) SetRecoveryTarget(v string)`

SetRecoveryTarget sets RecoveryTarget field to given value.

### HasRecoveryTarget

`func (o *RecoverOracleAppNewSourceConfig) HasRecoveryTarget() bool`

HasRecoveryTarget returns a boolean if a field has been set.

### SetRecoveryTargetNil

`func (o *RecoverOracleAppNewSourceConfig) SetRecoveryTargetNil(b bool)`

 SetRecoveryTargetNil sets the value for RecoveryTarget to be an explicit nil

### UnsetRecoveryTarget
`func (o *RecoverOracleAppNewSourceConfig) UnsetRecoveryTarget()`

UnsetRecoveryTarget ensures that no value is present for RecoveryTarget, not even an explicit nil
### GetRecoverDatabaseParams

`func (o *RecoverOracleAppNewSourceConfig) GetRecoverDatabaseParams() RecoverOracleNewTargetDatabaseConfig`

GetRecoverDatabaseParams returns the RecoverDatabaseParams field if non-nil, zero value otherwise.

### GetRecoverDatabaseParamsOk

`func (o *RecoverOracleAppNewSourceConfig) GetRecoverDatabaseParamsOk() (*RecoverOracleNewTargetDatabaseConfig, bool)`

GetRecoverDatabaseParamsOk returns a tuple with the RecoverDatabaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverDatabaseParams

`func (o *RecoverOracleAppNewSourceConfig) SetRecoverDatabaseParams(v RecoverOracleNewTargetDatabaseConfig)`

SetRecoverDatabaseParams sets RecoverDatabaseParams field to given value.

### HasRecoverDatabaseParams

`func (o *RecoverOracleAppNewSourceConfig) HasRecoverDatabaseParams() bool`

HasRecoverDatabaseParams returns a boolean if a field has been set.

### SetRecoverDatabaseParamsNil

`func (o *RecoverOracleAppNewSourceConfig) SetRecoverDatabaseParamsNil(b bool)`

 SetRecoverDatabaseParamsNil sets the value for RecoverDatabaseParams to be an explicit nil

### UnsetRecoverDatabaseParams
`func (o *RecoverOracleAppNewSourceConfig) UnsetRecoverDatabaseParams()`

UnsetRecoverDatabaseParams ensures that no value is present for RecoverDatabaseParams, not even an explicit nil
### GetRecoverViewParams

`func (o *RecoverOracleAppNewSourceConfig) GetRecoverViewParams() RecoverOracleNewTargetViewConfig`

GetRecoverViewParams returns the RecoverViewParams field if non-nil, zero value otherwise.

### GetRecoverViewParamsOk

`func (o *RecoverOracleAppNewSourceConfig) GetRecoverViewParamsOk() (*RecoverOracleNewTargetViewConfig, bool)`

GetRecoverViewParamsOk returns a tuple with the RecoverViewParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverViewParams

`func (o *RecoverOracleAppNewSourceConfig) SetRecoverViewParams(v RecoverOracleNewTargetViewConfig)`

SetRecoverViewParams sets RecoverViewParams field to given value.

### HasRecoverViewParams

`func (o *RecoverOracleAppNewSourceConfig) HasRecoverViewParams() bool`

HasRecoverViewParams returns a boolean if a field has been set.

### SetRecoverViewParamsNil

`func (o *RecoverOracleAppNewSourceConfig) SetRecoverViewParamsNil(b bool)`

 SetRecoverViewParamsNil sets the value for RecoverViewParams to be an explicit nil

### UnsetRecoverViewParams
`func (o *RecoverOracleAppNewSourceConfig) UnsetRecoverViewParams()`

UnsetRecoverViewParams ensures that no value is present for RecoverViewParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


