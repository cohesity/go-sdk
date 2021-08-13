# ExchangeRecoverDatabaseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseSource** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the source id of Exchange database which has to be recovered. | 
**RecoverToNewSource** | **bool** | Specifies the parameter whether the recovery should be performed to a new or an existing Source Target. | 
**RecoveryTargetConfig** | Pointer to [**NullableExchangeDatabaseRecoveryTargetConfig**](ExchangeDatabaseRecoveryTargetConfig.md) | Specifies the recovery target configuration if recovery has to be done to a different location which is different from original source. | [optional] 
**RestoreType** | **NullableString** | Specifies the type of exchange restore. | 
**ViewOptions** | Pointer to [**ViewOptions**](ViewOptions.md) |  | [optional] 

## Methods

### NewExchangeRecoverDatabaseParams

`func NewExchangeRecoverDatabaseParams(databaseSource NullableRecoveryObjectIdentifier, recoverToNewSource bool, restoreType NullableString, ) *ExchangeRecoverDatabaseParams`

NewExchangeRecoverDatabaseParams instantiates a new ExchangeRecoverDatabaseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeRecoverDatabaseParamsWithDefaults

`func NewExchangeRecoverDatabaseParamsWithDefaults() *ExchangeRecoverDatabaseParams`

NewExchangeRecoverDatabaseParamsWithDefaults instantiates a new ExchangeRecoverDatabaseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseSource

`func (o *ExchangeRecoverDatabaseParams) GetDatabaseSource() RecoveryObjectIdentifier`

GetDatabaseSource returns the DatabaseSource field if non-nil, zero value otherwise.

### GetDatabaseSourceOk

`func (o *ExchangeRecoverDatabaseParams) GetDatabaseSourceOk() (*RecoveryObjectIdentifier, bool)`

GetDatabaseSourceOk returns a tuple with the DatabaseSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseSource

`func (o *ExchangeRecoverDatabaseParams) SetDatabaseSource(v RecoveryObjectIdentifier)`

SetDatabaseSource sets DatabaseSource field to given value.


### SetDatabaseSourceNil

`func (o *ExchangeRecoverDatabaseParams) SetDatabaseSourceNil(b bool)`

 SetDatabaseSourceNil sets the value for DatabaseSource to be an explicit nil

### UnsetDatabaseSource
`func (o *ExchangeRecoverDatabaseParams) UnsetDatabaseSource()`

UnsetDatabaseSource ensures that no value is present for DatabaseSource, not even an explicit nil
### GetRecoverToNewSource

`func (o *ExchangeRecoverDatabaseParams) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *ExchangeRecoverDatabaseParams) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *ExchangeRecoverDatabaseParams) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.


### GetRecoveryTargetConfig

`func (o *ExchangeRecoverDatabaseParams) GetRecoveryTargetConfig() ExchangeDatabaseRecoveryTargetConfig`

GetRecoveryTargetConfig returns the RecoveryTargetConfig field if non-nil, zero value otherwise.

### GetRecoveryTargetConfigOk

`func (o *ExchangeRecoverDatabaseParams) GetRecoveryTargetConfigOk() (*ExchangeDatabaseRecoveryTargetConfig, bool)`

GetRecoveryTargetConfigOk returns a tuple with the RecoveryTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTargetConfig

`func (o *ExchangeRecoverDatabaseParams) SetRecoveryTargetConfig(v ExchangeDatabaseRecoveryTargetConfig)`

SetRecoveryTargetConfig sets RecoveryTargetConfig field to given value.

### HasRecoveryTargetConfig

`func (o *ExchangeRecoverDatabaseParams) HasRecoveryTargetConfig() bool`

HasRecoveryTargetConfig returns a boolean if a field has been set.

### SetRecoveryTargetConfigNil

`func (o *ExchangeRecoverDatabaseParams) SetRecoveryTargetConfigNil(b bool)`

 SetRecoveryTargetConfigNil sets the value for RecoveryTargetConfig to be an explicit nil

### UnsetRecoveryTargetConfig
`func (o *ExchangeRecoverDatabaseParams) UnsetRecoveryTargetConfig()`

UnsetRecoveryTargetConfig ensures that no value is present for RecoveryTargetConfig, not even an explicit nil
### GetRestoreType

`func (o *ExchangeRecoverDatabaseParams) GetRestoreType() string`

GetRestoreType returns the RestoreType field if non-nil, zero value otherwise.

### GetRestoreTypeOk

`func (o *ExchangeRecoverDatabaseParams) GetRestoreTypeOk() (*string, bool)`

GetRestoreTypeOk returns a tuple with the RestoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreType

`func (o *ExchangeRecoverDatabaseParams) SetRestoreType(v string)`

SetRestoreType sets RestoreType field to given value.


### SetRestoreTypeNil

`func (o *ExchangeRecoverDatabaseParams) SetRestoreTypeNil(b bool)`

 SetRestoreTypeNil sets the value for RestoreType to be an explicit nil

### UnsetRestoreType
`func (o *ExchangeRecoverDatabaseParams) UnsetRestoreType()`

UnsetRestoreType ensures that no value is present for RestoreType, not even an explicit nil
### GetViewOptions

`func (o *ExchangeRecoverDatabaseParams) GetViewOptions() ViewOptions`

GetViewOptions returns the ViewOptions field if non-nil, zero value otherwise.

### GetViewOptionsOk

`func (o *ExchangeRecoverDatabaseParams) GetViewOptionsOk() (*ViewOptions, bool)`

GetViewOptionsOk returns a tuple with the ViewOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewOptions

`func (o *ExchangeRecoverDatabaseParams) SetViewOptions(v ViewOptions)`

SetViewOptions sets ViewOptions field to given value.

### HasViewOptions

`func (o *ExchangeRecoverDatabaseParams) HasViewOptions() bool`

HasViewOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


