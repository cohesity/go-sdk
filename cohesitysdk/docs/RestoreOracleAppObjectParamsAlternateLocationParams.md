# RestoreOracleAppObjectParamsAlternateLocationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseDir** | Pointer to **NullableString** | Base directory of Oracle at destination. Example : /u01/app/oracle | [optional] 
**DatabaseFileDestination** | Pointer to **NullableString** | Location to put the database files(datafiles, logfiles etc.). | [optional] 
**HomeDir** | Pointer to **NullableString** | Home directory of Oracle at destination. Example : /u01/app/oracle/product/11.2.0.3/db_1 | [optional] 
**NewDatabaseName** | Pointer to **NullableString** | The name of the Oracle database that we restore to. | [optional] 
**NewSidDeprecated** | Pointer to **NullableString** | Deprecated field SID of new Oracle database. | [optional] 
**OracleDbConfig** | Pointer to [**OracleDBConfig**](OracleDBConfig.md) |  | [optional] 

## Methods

### NewRestoreOracleAppObjectParamsAlternateLocationParams

`func NewRestoreOracleAppObjectParamsAlternateLocationParams() *RestoreOracleAppObjectParamsAlternateLocationParams`

NewRestoreOracleAppObjectParamsAlternateLocationParams instantiates a new RestoreOracleAppObjectParamsAlternateLocationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreOracleAppObjectParamsAlternateLocationParamsWithDefaults

`func NewRestoreOracleAppObjectParamsAlternateLocationParamsWithDefaults() *RestoreOracleAppObjectParamsAlternateLocationParams`

NewRestoreOracleAppObjectParamsAlternateLocationParamsWithDefaults instantiates a new RestoreOracleAppObjectParamsAlternateLocationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseDir

`func (o *RestoreOracleAppObjectParamsAlternateLocationParams) GetBaseDir() string`

GetBaseDir returns the BaseDir field if non-nil, zero value otherwise.

### GetBaseDirOk

`func (o *RestoreOracleAppObjectParamsAlternateLocationParams) GetBaseDirOk() (*string, bool)`

GetBaseDirOk returns a tuple with the BaseDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDir

`func (o *RestoreOracleAppObjectParamsAlternateLocationParams) SetBaseDir(v string)`

SetBaseDir sets BaseDir field to given value.

### HasBaseDir

`func (o *RestoreOracleAppObjectParamsAlternateLocationParams) HasBaseDir() bool`

HasBaseDir returns a boolean if a field has been set.

### SetBaseDirNil

`func (o *RestoreOracleAppObjectParamsAlternateLocationParams) SetBaseDirNil(b bool)`

 SetBaseDirNil sets the value for BaseDir to be an explicit nil

### UnsetBaseDir
`func (o *RestoreOracleAppObjectParamsAlternateLocationParams) UnsetBaseDir()`

UnsetBaseDir ensures that no value is present for BaseDir, not even an explicit nil
### GetDatabaseFileDestination

`func (o *RestoreOracleAppObjectParamsAlternateLocationParams) GetDatabaseFileDestination() string`

GetDatabaseFileDestination returns the DatabaseFileDestination field if non-nil, zero value otherwise.

### GetDatabaseFileDestinationOk

`func (o *RestoreOracleAppObjectParamsAlternateLocationParams) GetDatabaseFileDestinationOk() (*string, bool)`

GetDatabaseFileDestinationOk returns a tuple with the DatabaseFileDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseFileDestination

`func (o *RestoreOracleAppObjectParamsAlternateLocationParams) SetDatabaseFileDestination(v string)`

SetDatabaseFileDestination sets DatabaseFileDestination field to given value.

### HasDatabaseFileDestination

`func (o *RestoreOracleAppObjectParamsAlternateLocationParams) HasDatabaseFileDestination() bool`

HasDatabaseFileDestination returns a boolean if a field has been set.

### SetDatabaseFileDestinationNil

`func (o *RestoreOracleAppObjectParamsAlternateLocationParams) SetDatabaseFileDestinationNil(b bool)`

 SetDatabaseFileDestinationNil sets the value for DatabaseFileDestination to be an explicit nil

### UnsetDatabaseFileDestination
`func (o *RestoreOracleAppObjectParamsAlternateLocationParams) UnsetDatabaseFileDestination()`

UnsetDatabaseFileDestination ensures that no value is present for DatabaseFileDestination, not even an explicit nil
### GetHomeDir

`func (o *RestoreOracleAppObjectParamsAlternateLocationParams) GetHomeDir() string`

GetHomeDir returns the HomeDir field if non-nil, zero value otherwise.

### GetHomeDirOk

`func (o *RestoreOracleAppObjectParamsAlternateLocationParams) GetHomeDirOk() (*string, bool)`

GetHomeDirOk returns a tuple with the HomeDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeDir

`func (o *RestoreOracleAppObjectParamsAlternateLocationParams) SetHomeDir(v string)`

SetHomeDir sets HomeDir field to given value.

### HasHomeDir

`func (o *RestoreOracleAppObjectParamsAlternateLocationParams) HasHomeDir() bool`

HasHomeDir returns a boolean if a field has been set.

### SetHomeDirNil

`func (o *RestoreOracleAppObjectParamsAlternateLocationParams) SetHomeDirNil(b bool)`

 SetHomeDirNil sets the value for HomeDir to be an explicit nil

### UnsetHomeDir
`func (o *RestoreOracleAppObjectParamsAlternateLocationParams) UnsetHomeDir()`

UnsetHomeDir ensures that no value is present for HomeDir, not even an explicit nil
### GetNewDatabaseName

`func (o *RestoreOracleAppObjectParamsAlternateLocationParams) GetNewDatabaseName() string`

GetNewDatabaseName returns the NewDatabaseName field if non-nil, zero value otherwise.

### GetNewDatabaseNameOk

`func (o *RestoreOracleAppObjectParamsAlternateLocationParams) GetNewDatabaseNameOk() (*string, bool)`

GetNewDatabaseNameOk returns a tuple with the NewDatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewDatabaseName

`func (o *RestoreOracleAppObjectParamsAlternateLocationParams) SetNewDatabaseName(v string)`

SetNewDatabaseName sets NewDatabaseName field to given value.

### HasNewDatabaseName

`func (o *RestoreOracleAppObjectParamsAlternateLocationParams) HasNewDatabaseName() bool`

HasNewDatabaseName returns a boolean if a field has been set.

### SetNewDatabaseNameNil

`func (o *RestoreOracleAppObjectParamsAlternateLocationParams) SetNewDatabaseNameNil(b bool)`

 SetNewDatabaseNameNil sets the value for NewDatabaseName to be an explicit nil

### UnsetNewDatabaseName
`func (o *RestoreOracleAppObjectParamsAlternateLocationParams) UnsetNewDatabaseName()`

UnsetNewDatabaseName ensures that no value is present for NewDatabaseName, not even an explicit nil
### GetNewSidDeprecated

`func (o *RestoreOracleAppObjectParamsAlternateLocationParams) GetNewSidDeprecated() string`

GetNewSidDeprecated returns the NewSidDeprecated field if non-nil, zero value otherwise.

### GetNewSidDeprecatedOk

`func (o *RestoreOracleAppObjectParamsAlternateLocationParams) GetNewSidDeprecatedOk() (*string, bool)`

GetNewSidDeprecatedOk returns a tuple with the NewSidDeprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSidDeprecated

`func (o *RestoreOracleAppObjectParamsAlternateLocationParams) SetNewSidDeprecated(v string)`

SetNewSidDeprecated sets NewSidDeprecated field to given value.

### HasNewSidDeprecated

`func (o *RestoreOracleAppObjectParamsAlternateLocationParams) HasNewSidDeprecated() bool`

HasNewSidDeprecated returns a boolean if a field has been set.

### SetNewSidDeprecatedNil

`func (o *RestoreOracleAppObjectParamsAlternateLocationParams) SetNewSidDeprecatedNil(b bool)`

 SetNewSidDeprecatedNil sets the value for NewSidDeprecated to be an explicit nil

### UnsetNewSidDeprecated
`func (o *RestoreOracleAppObjectParamsAlternateLocationParams) UnsetNewSidDeprecated()`

UnsetNewSidDeprecated ensures that no value is present for NewSidDeprecated, not even an explicit nil
### GetOracleDbConfig

`func (o *RestoreOracleAppObjectParamsAlternateLocationParams) GetOracleDbConfig() OracleDBConfig`

GetOracleDbConfig returns the OracleDbConfig field if non-nil, zero value otherwise.

### GetOracleDbConfigOk

`func (o *RestoreOracleAppObjectParamsAlternateLocationParams) GetOracleDbConfigOk() (*OracleDBConfig, bool)`

GetOracleDbConfigOk returns a tuple with the OracleDbConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleDbConfig

`func (o *RestoreOracleAppObjectParamsAlternateLocationParams) SetOracleDbConfig(v OracleDBConfig)`

SetOracleDbConfig sets OracleDbConfig field to given value.

### HasOracleDbConfig

`func (o *RestoreOracleAppObjectParamsAlternateLocationParams) HasOracleDbConfig() bool`

HasOracleDbConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


