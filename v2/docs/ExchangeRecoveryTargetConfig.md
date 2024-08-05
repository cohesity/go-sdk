# ExchangeRecoveryTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseDirectoryLocation** | Pointer to **NullableString** | Specifies the directory where to put the database data files. Missing directory will be automatically created. | [optional] 
**DatabaseName** | Pointer to **NullableString** | Specifies a new name for the restored database. | [optional] 
**LogDirectoryLocation** | Pointer to **NullableString** | Specifies the directory where to put the database log files. Missing directory will be automatically created. | [optional] 
**MountDatabase** | Pointer to **NullableBool** | Specifies whether to mount the database after successful recovery. | [optional] 
**RestoreAsRecoveryDB** | Pointer to **NullableBool** | Specifies whether to restore the Database as Recovery database. | [optional] 
**RollForwardRecovery** | Pointer to **NullableBool** | Specifies whether to use the latest logs on Exchange Server to perform roll-forward recovery. | [optional] 
**Source** | Pointer to [**NullableExchangeRecoveryTargetConfigSource**](ExchangeRecoveryTargetConfigSource.md) |  | [optional] 

## Methods

### NewExchangeRecoveryTargetConfig

`func NewExchangeRecoveryTargetConfig() *ExchangeRecoveryTargetConfig`

NewExchangeRecoveryTargetConfig instantiates a new ExchangeRecoveryTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeRecoveryTargetConfigWithDefaults

`func NewExchangeRecoveryTargetConfigWithDefaults() *ExchangeRecoveryTargetConfig`

NewExchangeRecoveryTargetConfigWithDefaults instantiates a new ExchangeRecoveryTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseDirectoryLocation

`func (o *ExchangeRecoveryTargetConfig) GetDatabaseDirectoryLocation() string`

GetDatabaseDirectoryLocation returns the DatabaseDirectoryLocation field if non-nil, zero value otherwise.

### GetDatabaseDirectoryLocationOk

`func (o *ExchangeRecoveryTargetConfig) GetDatabaseDirectoryLocationOk() (*string, bool)`

GetDatabaseDirectoryLocationOk returns a tuple with the DatabaseDirectoryLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseDirectoryLocation

`func (o *ExchangeRecoveryTargetConfig) SetDatabaseDirectoryLocation(v string)`

SetDatabaseDirectoryLocation sets DatabaseDirectoryLocation field to given value.

### HasDatabaseDirectoryLocation

`func (o *ExchangeRecoveryTargetConfig) HasDatabaseDirectoryLocation() bool`

HasDatabaseDirectoryLocation returns a boolean if a field has been set.

### SetDatabaseDirectoryLocationNil

`func (o *ExchangeRecoveryTargetConfig) SetDatabaseDirectoryLocationNil(b bool)`

 SetDatabaseDirectoryLocationNil sets the value for DatabaseDirectoryLocation to be an explicit nil

### UnsetDatabaseDirectoryLocation
`func (o *ExchangeRecoveryTargetConfig) UnsetDatabaseDirectoryLocation()`

UnsetDatabaseDirectoryLocation ensures that no value is present for DatabaseDirectoryLocation, not even an explicit nil
### GetDatabaseName

`func (o *ExchangeRecoveryTargetConfig) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *ExchangeRecoveryTargetConfig) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *ExchangeRecoveryTargetConfig) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *ExchangeRecoveryTargetConfig) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.

### SetDatabaseNameNil

`func (o *ExchangeRecoveryTargetConfig) SetDatabaseNameNil(b bool)`

 SetDatabaseNameNil sets the value for DatabaseName to be an explicit nil

### UnsetDatabaseName
`func (o *ExchangeRecoveryTargetConfig) UnsetDatabaseName()`

UnsetDatabaseName ensures that no value is present for DatabaseName, not even an explicit nil
### GetLogDirectoryLocation

`func (o *ExchangeRecoveryTargetConfig) GetLogDirectoryLocation() string`

GetLogDirectoryLocation returns the LogDirectoryLocation field if non-nil, zero value otherwise.

### GetLogDirectoryLocationOk

`func (o *ExchangeRecoveryTargetConfig) GetLogDirectoryLocationOk() (*string, bool)`

GetLogDirectoryLocationOk returns a tuple with the LogDirectoryLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDirectoryLocation

`func (o *ExchangeRecoveryTargetConfig) SetLogDirectoryLocation(v string)`

SetLogDirectoryLocation sets LogDirectoryLocation field to given value.

### HasLogDirectoryLocation

`func (o *ExchangeRecoveryTargetConfig) HasLogDirectoryLocation() bool`

HasLogDirectoryLocation returns a boolean if a field has been set.

### SetLogDirectoryLocationNil

`func (o *ExchangeRecoveryTargetConfig) SetLogDirectoryLocationNil(b bool)`

 SetLogDirectoryLocationNil sets the value for LogDirectoryLocation to be an explicit nil

### UnsetLogDirectoryLocation
`func (o *ExchangeRecoveryTargetConfig) UnsetLogDirectoryLocation()`

UnsetLogDirectoryLocation ensures that no value is present for LogDirectoryLocation, not even an explicit nil
### GetMountDatabase

`func (o *ExchangeRecoveryTargetConfig) GetMountDatabase() bool`

GetMountDatabase returns the MountDatabase field if non-nil, zero value otherwise.

### GetMountDatabaseOk

`func (o *ExchangeRecoveryTargetConfig) GetMountDatabaseOk() (*bool, bool)`

GetMountDatabaseOk returns a tuple with the MountDatabase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountDatabase

`func (o *ExchangeRecoveryTargetConfig) SetMountDatabase(v bool)`

SetMountDatabase sets MountDatabase field to given value.

### HasMountDatabase

`func (o *ExchangeRecoveryTargetConfig) HasMountDatabase() bool`

HasMountDatabase returns a boolean if a field has been set.

### SetMountDatabaseNil

`func (o *ExchangeRecoveryTargetConfig) SetMountDatabaseNil(b bool)`

 SetMountDatabaseNil sets the value for MountDatabase to be an explicit nil

### UnsetMountDatabase
`func (o *ExchangeRecoveryTargetConfig) UnsetMountDatabase()`

UnsetMountDatabase ensures that no value is present for MountDatabase, not even an explicit nil
### GetRestoreAsRecoveryDB

`func (o *ExchangeRecoveryTargetConfig) GetRestoreAsRecoveryDB() bool`

GetRestoreAsRecoveryDB returns the RestoreAsRecoveryDB field if non-nil, zero value otherwise.

### GetRestoreAsRecoveryDBOk

`func (o *ExchangeRecoveryTargetConfig) GetRestoreAsRecoveryDBOk() (*bool, bool)`

GetRestoreAsRecoveryDBOk returns a tuple with the RestoreAsRecoveryDB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreAsRecoveryDB

`func (o *ExchangeRecoveryTargetConfig) SetRestoreAsRecoveryDB(v bool)`

SetRestoreAsRecoveryDB sets RestoreAsRecoveryDB field to given value.

### HasRestoreAsRecoveryDB

`func (o *ExchangeRecoveryTargetConfig) HasRestoreAsRecoveryDB() bool`

HasRestoreAsRecoveryDB returns a boolean if a field has been set.

### SetRestoreAsRecoveryDBNil

`func (o *ExchangeRecoveryTargetConfig) SetRestoreAsRecoveryDBNil(b bool)`

 SetRestoreAsRecoveryDBNil sets the value for RestoreAsRecoveryDB to be an explicit nil

### UnsetRestoreAsRecoveryDB
`func (o *ExchangeRecoveryTargetConfig) UnsetRestoreAsRecoveryDB()`

UnsetRestoreAsRecoveryDB ensures that no value is present for RestoreAsRecoveryDB, not even an explicit nil
### GetRollForwardRecovery

`func (o *ExchangeRecoveryTargetConfig) GetRollForwardRecovery() bool`

GetRollForwardRecovery returns the RollForwardRecovery field if non-nil, zero value otherwise.

### GetRollForwardRecoveryOk

`func (o *ExchangeRecoveryTargetConfig) GetRollForwardRecoveryOk() (*bool, bool)`

GetRollForwardRecoveryOk returns a tuple with the RollForwardRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollForwardRecovery

`func (o *ExchangeRecoveryTargetConfig) SetRollForwardRecovery(v bool)`

SetRollForwardRecovery sets RollForwardRecovery field to given value.

### HasRollForwardRecovery

`func (o *ExchangeRecoveryTargetConfig) HasRollForwardRecovery() bool`

HasRollForwardRecovery returns a boolean if a field has been set.

### SetRollForwardRecoveryNil

`func (o *ExchangeRecoveryTargetConfig) SetRollForwardRecoveryNil(b bool)`

 SetRollForwardRecoveryNil sets the value for RollForwardRecovery to be an explicit nil

### UnsetRollForwardRecovery
`func (o *ExchangeRecoveryTargetConfig) UnsetRollForwardRecovery()`

UnsetRollForwardRecovery ensures that no value is present for RollForwardRecovery, not even an explicit nil
### GetSource

`func (o *ExchangeRecoveryTargetConfig) GetSource() ExchangeRecoveryTargetConfigSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ExchangeRecoveryTargetConfig) GetSourceOk() (*ExchangeRecoveryTargetConfigSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ExchangeRecoveryTargetConfig) SetSource(v ExchangeRecoveryTargetConfigSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *ExchangeRecoveryTargetConfig) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *ExchangeRecoveryTargetConfig) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *ExchangeRecoveryTargetConfig) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


