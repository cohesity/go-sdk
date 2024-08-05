# RecoverOracleGranularRestoreInfoPdbRestoreParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DropDuplicatePDB** | Pointer to **NullableBool** | Specifies if the PDB should be ignored if a PDB already exists with same name. | [optional] 
**IncludeInRestore** | Pointer to **NullableBool** | Specifies whether to restore or skip the provided PDBs list. | [optional] 
**PdbObjects** | Pointer to [**[]OraclePdbObjectInfo**](OraclePdbObjectInfo.md) | Specifies list of PDB objects to restore. | [optional] 
**RenamePdbMap** | Pointer to [**[]KeyValuePair**](KeyValuePair.md) | Specifies the new PDB name mapping to existing PDBs. | [optional] 
**RestoreToExistingCdb** | Pointer to **NullableBool** | Specifies if pdbs should be restored to an existing CDB. | [optional] 
**SourceCdbKeystorePassword** | Pointer to **NullableString** | Specifies the keystore password of the source CDB. | [optional] 
**TargetCdbKeystorePassword** | Pointer to **NullableString** | Specifies the keystore password of the target CDB. | [optional] 

## Methods

### NewRecoverOracleGranularRestoreInfoPdbRestoreParams

`func NewRecoverOracleGranularRestoreInfoPdbRestoreParams() *RecoverOracleGranularRestoreInfoPdbRestoreParams`

NewRecoverOracleGranularRestoreInfoPdbRestoreParams instantiates a new RecoverOracleGranularRestoreInfoPdbRestoreParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverOracleGranularRestoreInfoPdbRestoreParamsWithDefaults

`func NewRecoverOracleGranularRestoreInfoPdbRestoreParamsWithDefaults() *RecoverOracleGranularRestoreInfoPdbRestoreParams`

NewRecoverOracleGranularRestoreInfoPdbRestoreParamsWithDefaults instantiates a new RecoverOracleGranularRestoreInfoPdbRestoreParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDropDuplicatePDB

`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) GetDropDuplicatePDB() bool`

GetDropDuplicatePDB returns the DropDuplicatePDB field if non-nil, zero value otherwise.

### GetDropDuplicatePDBOk

`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) GetDropDuplicatePDBOk() (*bool, bool)`

GetDropDuplicatePDBOk returns a tuple with the DropDuplicatePDB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropDuplicatePDB

`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) SetDropDuplicatePDB(v bool)`

SetDropDuplicatePDB sets DropDuplicatePDB field to given value.

### HasDropDuplicatePDB

`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) HasDropDuplicatePDB() bool`

HasDropDuplicatePDB returns a boolean if a field has been set.

### SetDropDuplicatePDBNil

`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) SetDropDuplicatePDBNil(b bool)`

 SetDropDuplicatePDBNil sets the value for DropDuplicatePDB to be an explicit nil

### UnsetDropDuplicatePDB
`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) UnsetDropDuplicatePDB()`

UnsetDropDuplicatePDB ensures that no value is present for DropDuplicatePDB, not even an explicit nil
### GetIncludeInRestore

`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) GetIncludeInRestore() bool`

GetIncludeInRestore returns the IncludeInRestore field if non-nil, zero value otherwise.

### GetIncludeInRestoreOk

`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) GetIncludeInRestoreOk() (*bool, bool)`

GetIncludeInRestoreOk returns a tuple with the IncludeInRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInRestore

`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) SetIncludeInRestore(v bool)`

SetIncludeInRestore sets IncludeInRestore field to given value.

### HasIncludeInRestore

`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) HasIncludeInRestore() bool`

HasIncludeInRestore returns a boolean if a field has been set.

### SetIncludeInRestoreNil

`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) SetIncludeInRestoreNil(b bool)`

 SetIncludeInRestoreNil sets the value for IncludeInRestore to be an explicit nil

### UnsetIncludeInRestore
`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) UnsetIncludeInRestore()`

UnsetIncludeInRestore ensures that no value is present for IncludeInRestore, not even an explicit nil
### GetPdbObjects

`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) GetPdbObjects() []OraclePdbObjectInfo`

GetPdbObjects returns the PdbObjects field if non-nil, zero value otherwise.

### GetPdbObjectsOk

`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) GetPdbObjectsOk() (*[]OraclePdbObjectInfo, bool)`

GetPdbObjectsOk returns a tuple with the PdbObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdbObjects

`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) SetPdbObjects(v []OraclePdbObjectInfo)`

SetPdbObjects sets PdbObjects field to given value.

### HasPdbObjects

`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) HasPdbObjects() bool`

HasPdbObjects returns a boolean if a field has been set.

### SetPdbObjectsNil

`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) SetPdbObjectsNil(b bool)`

 SetPdbObjectsNil sets the value for PdbObjects to be an explicit nil

### UnsetPdbObjects
`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) UnsetPdbObjects()`

UnsetPdbObjects ensures that no value is present for PdbObjects, not even an explicit nil
### GetRenamePdbMap

`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) GetRenamePdbMap() []KeyValuePair`

GetRenamePdbMap returns the RenamePdbMap field if non-nil, zero value otherwise.

### GetRenamePdbMapOk

`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) GetRenamePdbMapOk() (*[]KeyValuePair, bool)`

GetRenamePdbMapOk returns a tuple with the RenamePdbMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenamePdbMap

`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) SetRenamePdbMap(v []KeyValuePair)`

SetRenamePdbMap sets RenamePdbMap field to given value.

### HasRenamePdbMap

`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) HasRenamePdbMap() bool`

HasRenamePdbMap returns a boolean if a field has been set.

### SetRenamePdbMapNil

`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) SetRenamePdbMapNil(b bool)`

 SetRenamePdbMapNil sets the value for RenamePdbMap to be an explicit nil

### UnsetRenamePdbMap
`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) UnsetRenamePdbMap()`

UnsetRenamePdbMap ensures that no value is present for RenamePdbMap, not even an explicit nil
### GetRestoreToExistingCdb

`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) GetRestoreToExistingCdb() bool`

GetRestoreToExistingCdb returns the RestoreToExistingCdb field if non-nil, zero value otherwise.

### GetRestoreToExistingCdbOk

`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) GetRestoreToExistingCdbOk() (*bool, bool)`

GetRestoreToExistingCdbOk returns a tuple with the RestoreToExistingCdb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreToExistingCdb

`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) SetRestoreToExistingCdb(v bool)`

SetRestoreToExistingCdb sets RestoreToExistingCdb field to given value.

### HasRestoreToExistingCdb

`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) HasRestoreToExistingCdb() bool`

HasRestoreToExistingCdb returns a boolean if a field has been set.

### SetRestoreToExistingCdbNil

`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) SetRestoreToExistingCdbNil(b bool)`

 SetRestoreToExistingCdbNil sets the value for RestoreToExistingCdb to be an explicit nil

### UnsetRestoreToExistingCdb
`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) UnsetRestoreToExistingCdb()`

UnsetRestoreToExistingCdb ensures that no value is present for RestoreToExistingCdb, not even an explicit nil
### GetSourceCdbKeystorePassword

`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) GetSourceCdbKeystorePassword() string`

GetSourceCdbKeystorePassword returns the SourceCdbKeystorePassword field if non-nil, zero value otherwise.

### GetSourceCdbKeystorePasswordOk

`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) GetSourceCdbKeystorePasswordOk() (*string, bool)`

GetSourceCdbKeystorePasswordOk returns a tuple with the SourceCdbKeystorePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCdbKeystorePassword

`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) SetSourceCdbKeystorePassword(v string)`

SetSourceCdbKeystorePassword sets SourceCdbKeystorePassword field to given value.

### HasSourceCdbKeystorePassword

`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) HasSourceCdbKeystorePassword() bool`

HasSourceCdbKeystorePassword returns a boolean if a field has been set.

### SetSourceCdbKeystorePasswordNil

`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) SetSourceCdbKeystorePasswordNil(b bool)`

 SetSourceCdbKeystorePasswordNil sets the value for SourceCdbKeystorePassword to be an explicit nil

### UnsetSourceCdbKeystorePassword
`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) UnsetSourceCdbKeystorePassword()`

UnsetSourceCdbKeystorePassword ensures that no value is present for SourceCdbKeystorePassword, not even an explicit nil
### GetTargetCdbKeystorePassword

`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) GetTargetCdbKeystorePassword() string`

GetTargetCdbKeystorePassword returns the TargetCdbKeystorePassword field if non-nil, zero value otherwise.

### GetTargetCdbKeystorePasswordOk

`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) GetTargetCdbKeystorePasswordOk() (*string, bool)`

GetTargetCdbKeystorePasswordOk returns a tuple with the TargetCdbKeystorePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCdbKeystorePassword

`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) SetTargetCdbKeystorePassword(v string)`

SetTargetCdbKeystorePassword sets TargetCdbKeystorePassword field to given value.

### HasTargetCdbKeystorePassword

`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) HasTargetCdbKeystorePassword() bool`

HasTargetCdbKeystorePassword returns a boolean if a field has been set.

### SetTargetCdbKeystorePasswordNil

`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) SetTargetCdbKeystorePasswordNil(b bool)`

 SetTargetCdbKeystorePasswordNil sets the value for TargetCdbKeystorePassword to be an explicit nil

### UnsetTargetCdbKeystorePassword
`func (o *RecoverOracleGranularRestoreInfoPdbRestoreParams) UnsetTargetCdbKeystorePassword()`

UnsetTargetCdbKeystorePassword ensures that no value is present for TargetCdbKeystorePassword, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


