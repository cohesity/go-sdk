# ConstructRestoreMetaInfoOracleParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbName** | **NullableString** | Specifies the name of the Oracle database that we restore to. | 
**BaseDir** | **NullableString** | Specifies the base directory of Oracle at destination. | 
**HomeDir** | **NullableString** | Specifies the home directory of Oracle at destination. | 
**DbFileDestination** | Pointer to **NullableString** | Specifies the location to put the database files(datafiles, logfiles etc.) | [optional] 
**IsClone** | **NullableBool** | Specifies whether operation is clone or not | 
**IsGranularRestore** | Pointer to **NullableBool** | Specifies whether the operation is granular restore or not. | [optional] 

## Methods

### NewConstructRestoreMetaInfoOracleParams

`func NewConstructRestoreMetaInfoOracleParams(dbName NullableString, baseDir NullableString, homeDir NullableString, isClone NullableBool, ) *ConstructRestoreMetaInfoOracleParams`

NewConstructRestoreMetaInfoOracleParams instantiates a new ConstructRestoreMetaInfoOracleParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConstructRestoreMetaInfoOracleParamsWithDefaults

`func NewConstructRestoreMetaInfoOracleParamsWithDefaults() *ConstructRestoreMetaInfoOracleParams`

NewConstructRestoreMetaInfoOracleParamsWithDefaults instantiates a new ConstructRestoreMetaInfoOracleParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbName

`func (o *ConstructRestoreMetaInfoOracleParams) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *ConstructRestoreMetaInfoOracleParams) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *ConstructRestoreMetaInfoOracleParams) SetDbName(v string)`

SetDbName sets DbName field to given value.


### SetDbNameNil

`func (o *ConstructRestoreMetaInfoOracleParams) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *ConstructRestoreMetaInfoOracleParams) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil
### GetBaseDir

`func (o *ConstructRestoreMetaInfoOracleParams) GetBaseDir() string`

GetBaseDir returns the BaseDir field if non-nil, zero value otherwise.

### GetBaseDirOk

`func (o *ConstructRestoreMetaInfoOracleParams) GetBaseDirOk() (*string, bool)`

GetBaseDirOk returns a tuple with the BaseDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDir

`func (o *ConstructRestoreMetaInfoOracleParams) SetBaseDir(v string)`

SetBaseDir sets BaseDir field to given value.


### SetBaseDirNil

`func (o *ConstructRestoreMetaInfoOracleParams) SetBaseDirNil(b bool)`

 SetBaseDirNil sets the value for BaseDir to be an explicit nil

### UnsetBaseDir
`func (o *ConstructRestoreMetaInfoOracleParams) UnsetBaseDir()`

UnsetBaseDir ensures that no value is present for BaseDir, not even an explicit nil
### GetHomeDir

`func (o *ConstructRestoreMetaInfoOracleParams) GetHomeDir() string`

GetHomeDir returns the HomeDir field if non-nil, zero value otherwise.

### GetHomeDirOk

`func (o *ConstructRestoreMetaInfoOracleParams) GetHomeDirOk() (*string, bool)`

GetHomeDirOk returns a tuple with the HomeDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeDir

`func (o *ConstructRestoreMetaInfoOracleParams) SetHomeDir(v string)`

SetHomeDir sets HomeDir field to given value.


### SetHomeDirNil

`func (o *ConstructRestoreMetaInfoOracleParams) SetHomeDirNil(b bool)`

 SetHomeDirNil sets the value for HomeDir to be an explicit nil

### UnsetHomeDir
`func (o *ConstructRestoreMetaInfoOracleParams) UnsetHomeDir()`

UnsetHomeDir ensures that no value is present for HomeDir, not even an explicit nil
### GetDbFileDestination

`func (o *ConstructRestoreMetaInfoOracleParams) GetDbFileDestination() string`

GetDbFileDestination returns the DbFileDestination field if non-nil, zero value otherwise.

### GetDbFileDestinationOk

`func (o *ConstructRestoreMetaInfoOracleParams) GetDbFileDestinationOk() (*string, bool)`

GetDbFileDestinationOk returns a tuple with the DbFileDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbFileDestination

`func (o *ConstructRestoreMetaInfoOracleParams) SetDbFileDestination(v string)`

SetDbFileDestination sets DbFileDestination field to given value.

### HasDbFileDestination

`func (o *ConstructRestoreMetaInfoOracleParams) HasDbFileDestination() bool`

HasDbFileDestination returns a boolean if a field has been set.

### SetDbFileDestinationNil

`func (o *ConstructRestoreMetaInfoOracleParams) SetDbFileDestinationNil(b bool)`

 SetDbFileDestinationNil sets the value for DbFileDestination to be an explicit nil

### UnsetDbFileDestination
`func (o *ConstructRestoreMetaInfoOracleParams) UnsetDbFileDestination()`

UnsetDbFileDestination ensures that no value is present for DbFileDestination, not even an explicit nil
### GetIsClone

`func (o *ConstructRestoreMetaInfoOracleParams) GetIsClone() bool`

GetIsClone returns the IsClone field if non-nil, zero value otherwise.

### GetIsCloneOk

`func (o *ConstructRestoreMetaInfoOracleParams) GetIsCloneOk() (*bool, bool)`

GetIsCloneOk returns a tuple with the IsClone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClone

`func (o *ConstructRestoreMetaInfoOracleParams) SetIsClone(v bool)`

SetIsClone sets IsClone field to given value.


### SetIsCloneNil

`func (o *ConstructRestoreMetaInfoOracleParams) SetIsCloneNil(b bool)`

 SetIsCloneNil sets the value for IsClone to be an explicit nil

### UnsetIsClone
`func (o *ConstructRestoreMetaInfoOracleParams) UnsetIsClone()`

UnsetIsClone ensures that no value is present for IsClone, not even an explicit nil
### GetIsGranularRestore

`func (o *ConstructRestoreMetaInfoOracleParams) GetIsGranularRestore() bool`

GetIsGranularRestore returns the IsGranularRestore field if non-nil, zero value otherwise.

### GetIsGranularRestoreOk

`func (o *ConstructRestoreMetaInfoOracleParams) GetIsGranularRestoreOk() (*bool, bool)`

GetIsGranularRestoreOk returns a tuple with the IsGranularRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGranularRestore

`func (o *ConstructRestoreMetaInfoOracleParams) SetIsGranularRestore(v bool)`

SetIsGranularRestore sets IsGranularRestore field to given value.

### HasIsGranularRestore

`func (o *ConstructRestoreMetaInfoOracleParams) HasIsGranularRestore() bool`

HasIsGranularRestore returns a boolean if a field has been set.

### SetIsGranularRestoreNil

`func (o *ConstructRestoreMetaInfoOracleParams) SetIsGranularRestoreNil(b bool)`

 SetIsGranularRestoreNil sets the value for IsGranularRestore to be an explicit nil

### UnsetIsGranularRestore
`func (o *ConstructRestoreMetaInfoOracleParams) UnsetIsGranularRestore()`

UnsetIsGranularRestore ensures that no value is present for IsGranularRestore, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


