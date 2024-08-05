# RecoverOracleAppNewSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | [**NullableRecoverOracleAppNewSourceConfigHost**](RecoverOracleAppNewSourceConfigHost.md) |  | 
**RecoverDatabaseParams** | Pointer to [**NullableRecoverOracleAppNewSourceConfigRecoverDatabaseParams**](RecoverOracleAppNewSourceConfigRecoverDatabaseParams.md) |  | [optional] 
**RecoverViewParams** | Pointer to [**NullableRecoverOracleAppNewSourceConfigRecoverViewParams**](RecoverOracleAppNewSourceConfigRecoverViewParams.md) |  | [optional] 
**RecoveryTarget** | Pointer to **NullableString** | Specifies if recovery target is a database or a view. | [optional] 

## Methods

### NewRecoverOracleAppNewSourceConfig

`func NewRecoverOracleAppNewSourceConfig(host NullableRecoverOracleAppNewSourceConfigHost, ) *RecoverOracleAppNewSourceConfig`

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

`func (o *RecoverOracleAppNewSourceConfig) GetHost() RecoverOracleAppNewSourceConfigHost`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RecoverOracleAppNewSourceConfig) GetHostOk() (*RecoverOracleAppNewSourceConfigHost, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RecoverOracleAppNewSourceConfig) SetHost(v RecoverOracleAppNewSourceConfigHost)`

SetHost sets Host field to given value.


### SetHostNil

`func (o *RecoverOracleAppNewSourceConfig) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *RecoverOracleAppNewSourceConfig) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetRecoverDatabaseParams

`func (o *RecoverOracleAppNewSourceConfig) GetRecoverDatabaseParams() RecoverOracleAppNewSourceConfigRecoverDatabaseParams`

GetRecoverDatabaseParams returns the RecoverDatabaseParams field if non-nil, zero value otherwise.

### GetRecoverDatabaseParamsOk

`func (o *RecoverOracleAppNewSourceConfig) GetRecoverDatabaseParamsOk() (*RecoverOracleAppNewSourceConfigRecoverDatabaseParams, bool)`

GetRecoverDatabaseParamsOk returns a tuple with the RecoverDatabaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverDatabaseParams

`func (o *RecoverOracleAppNewSourceConfig) SetRecoverDatabaseParams(v RecoverOracleAppNewSourceConfigRecoverDatabaseParams)`

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

`func (o *RecoverOracleAppNewSourceConfig) GetRecoverViewParams() RecoverOracleAppNewSourceConfigRecoverViewParams`

GetRecoverViewParams returns the RecoverViewParams field if non-nil, zero value otherwise.

### GetRecoverViewParamsOk

`func (o *RecoverOracleAppNewSourceConfig) GetRecoverViewParamsOk() (*RecoverOracleAppNewSourceConfigRecoverViewParams, bool)`

GetRecoverViewParamsOk returns a tuple with the RecoverViewParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverViewParams

`func (o *RecoverOracleAppNewSourceConfig) SetRecoverViewParams(v RecoverOracleAppNewSourceConfigRecoverViewParams)`

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


