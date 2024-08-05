# OraclePdbRestoreParams

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

### NewOraclePdbRestoreParams

`func NewOraclePdbRestoreParams() *OraclePdbRestoreParams`

NewOraclePdbRestoreParams instantiates a new OraclePdbRestoreParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOraclePdbRestoreParamsWithDefaults

`func NewOraclePdbRestoreParamsWithDefaults() *OraclePdbRestoreParams`

NewOraclePdbRestoreParamsWithDefaults instantiates a new OraclePdbRestoreParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDropDuplicatePDB

`func (o *OraclePdbRestoreParams) GetDropDuplicatePDB() bool`

GetDropDuplicatePDB returns the DropDuplicatePDB field if non-nil, zero value otherwise.

### GetDropDuplicatePDBOk

`func (o *OraclePdbRestoreParams) GetDropDuplicatePDBOk() (*bool, bool)`

GetDropDuplicatePDBOk returns a tuple with the DropDuplicatePDB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropDuplicatePDB

`func (o *OraclePdbRestoreParams) SetDropDuplicatePDB(v bool)`

SetDropDuplicatePDB sets DropDuplicatePDB field to given value.

### HasDropDuplicatePDB

`func (o *OraclePdbRestoreParams) HasDropDuplicatePDB() bool`

HasDropDuplicatePDB returns a boolean if a field has been set.

### SetDropDuplicatePDBNil

`func (o *OraclePdbRestoreParams) SetDropDuplicatePDBNil(b bool)`

 SetDropDuplicatePDBNil sets the value for DropDuplicatePDB to be an explicit nil

### UnsetDropDuplicatePDB
`func (o *OraclePdbRestoreParams) UnsetDropDuplicatePDB()`

UnsetDropDuplicatePDB ensures that no value is present for DropDuplicatePDB, not even an explicit nil
### GetIncludeInRestore

`func (o *OraclePdbRestoreParams) GetIncludeInRestore() bool`

GetIncludeInRestore returns the IncludeInRestore field if non-nil, zero value otherwise.

### GetIncludeInRestoreOk

`func (o *OraclePdbRestoreParams) GetIncludeInRestoreOk() (*bool, bool)`

GetIncludeInRestoreOk returns a tuple with the IncludeInRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInRestore

`func (o *OraclePdbRestoreParams) SetIncludeInRestore(v bool)`

SetIncludeInRestore sets IncludeInRestore field to given value.

### HasIncludeInRestore

`func (o *OraclePdbRestoreParams) HasIncludeInRestore() bool`

HasIncludeInRestore returns a boolean if a field has been set.

### SetIncludeInRestoreNil

`func (o *OraclePdbRestoreParams) SetIncludeInRestoreNil(b bool)`

 SetIncludeInRestoreNil sets the value for IncludeInRestore to be an explicit nil

### UnsetIncludeInRestore
`func (o *OraclePdbRestoreParams) UnsetIncludeInRestore()`

UnsetIncludeInRestore ensures that no value is present for IncludeInRestore, not even an explicit nil
### GetPdbObjects

`func (o *OraclePdbRestoreParams) GetPdbObjects() []OraclePdbObjectInfo`

GetPdbObjects returns the PdbObjects field if non-nil, zero value otherwise.

### GetPdbObjectsOk

`func (o *OraclePdbRestoreParams) GetPdbObjectsOk() (*[]OraclePdbObjectInfo, bool)`

GetPdbObjectsOk returns a tuple with the PdbObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdbObjects

`func (o *OraclePdbRestoreParams) SetPdbObjects(v []OraclePdbObjectInfo)`

SetPdbObjects sets PdbObjects field to given value.

### HasPdbObjects

`func (o *OraclePdbRestoreParams) HasPdbObjects() bool`

HasPdbObjects returns a boolean if a field has been set.

### SetPdbObjectsNil

`func (o *OraclePdbRestoreParams) SetPdbObjectsNil(b bool)`

 SetPdbObjectsNil sets the value for PdbObjects to be an explicit nil

### UnsetPdbObjects
`func (o *OraclePdbRestoreParams) UnsetPdbObjects()`

UnsetPdbObjects ensures that no value is present for PdbObjects, not even an explicit nil
### GetRenamePdbMap

`func (o *OraclePdbRestoreParams) GetRenamePdbMap() []KeyValuePair`

GetRenamePdbMap returns the RenamePdbMap field if non-nil, zero value otherwise.

### GetRenamePdbMapOk

`func (o *OraclePdbRestoreParams) GetRenamePdbMapOk() (*[]KeyValuePair, bool)`

GetRenamePdbMapOk returns a tuple with the RenamePdbMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenamePdbMap

`func (o *OraclePdbRestoreParams) SetRenamePdbMap(v []KeyValuePair)`

SetRenamePdbMap sets RenamePdbMap field to given value.

### HasRenamePdbMap

`func (o *OraclePdbRestoreParams) HasRenamePdbMap() bool`

HasRenamePdbMap returns a boolean if a field has been set.

### SetRenamePdbMapNil

`func (o *OraclePdbRestoreParams) SetRenamePdbMapNil(b bool)`

 SetRenamePdbMapNil sets the value for RenamePdbMap to be an explicit nil

### UnsetRenamePdbMap
`func (o *OraclePdbRestoreParams) UnsetRenamePdbMap()`

UnsetRenamePdbMap ensures that no value is present for RenamePdbMap, not even an explicit nil
### GetRestoreToExistingCdb

`func (o *OraclePdbRestoreParams) GetRestoreToExistingCdb() bool`

GetRestoreToExistingCdb returns the RestoreToExistingCdb field if non-nil, zero value otherwise.

### GetRestoreToExistingCdbOk

`func (o *OraclePdbRestoreParams) GetRestoreToExistingCdbOk() (*bool, bool)`

GetRestoreToExistingCdbOk returns a tuple with the RestoreToExistingCdb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreToExistingCdb

`func (o *OraclePdbRestoreParams) SetRestoreToExistingCdb(v bool)`

SetRestoreToExistingCdb sets RestoreToExistingCdb field to given value.

### HasRestoreToExistingCdb

`func (o *OraclePdbRestoreParams) HasRestoreToExistingCdb() bool`

HasRestoreToExistingCdb returns a boolean if a field has been set.

### SetRestoreToExistingCdbNil

`func (o *OraclePdbRestoreParams) SetRestoreToExistingCdbNil(b bool)`

 SetRestoreToExistingCdbNil sets the value for RestoreToExistingCdb to be an explicit nil

### UnsetRestoreToExistingCdb
`func (o *OraclePdbRestoreParams) UnsetRestoreToExistingCdb()`

UnsetRestoreToExistingCdb ensures that no value is present for RestoreToExistingCdb, not even an explicit nil
### GetSourceCdbKeystorePassword

`func (o *OraclePdbRestoreParams) GetSourceCdbKeystorePassword() string`

GetSourceCdbKeystorePassword returns the SourceCdbKeystorePassword field if non-nil, zero value otherwise.

### GetSourceCdbKeystorePasswordOk

`func (o *OraclePdbRestoreParams) GetSourceCdbKeystorePasswordOk() (*string, bool)`

GetSourceCdbKeystorePasswordOk returns a tuple with the SourceCdbKeystorePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCdbKeystorePassword

`func (o *OraclePdbRestoreParams) SetSourceCdbKeystorePassword(v string)`

SetSourceCdbKeystorePassword sets SourceCdbKeystorePassword field to given value.

### HasSourceCdbKeystorePassword

`func (o *OraclePdbRestoreParams) HasSourceCdbKeystorePassword() bool`

HasSourceCdbKeystorePassword returns a boolean if a field has been set.

### SetSourceCdbKeystorePasswordNil

`func (o *OraclePdbRestoreParams) SetSourceCdbKeystorePasswordNil(b bool)`

 SetSourceCdbKeystorePasswordNil sets the value for SourceCdbKeystorePassword to be an explicit nil

### UnsetSourceCdbKeystorePassword
`func (o *OraclePdbRestoreParams) UnsetSourceCdbKeystorePassword()`

UnsetSourceCdbKeystorePassword ensures that no value is present for SourceCdbKeystorePassword, not even an explicit nil
### GetTargetCdbKeystorePassword

`func (o *OraclePdbRestoreParams) GetTargetCdbKeystorePassword() string`

GetTargetCdbKeystorePassword returns the TargetCdbKeystorePassword field if non-nil, zero value otherwise.

### GetTargetCdbKeystorePasswordOk

`func (o *OraclePdbRestoreParams) GetTargetCdbKeystorePasswordOk() (*string, bool)`

GetTargetCdbKeystorePasswordOk returns a tuple with the TargetCdbKeystorePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCdbKeystorePassword

`func (o *OraclePdbRestoreParams) SetTargetCdbKeystorePassword(v string)`

SetTargetCdbKeystorePassword sets TargetCdbKeystorePassword field to given value.

### HasTargetCdbKeystorePassword

`func (o *OraclePdbRestoreParams) HasTargetCdbKeystorePassword() bool`

HasTargetCdbKeystorePassword returns a boolean if a field has been set.

### SetTargetCdbKeystorePasswordNil

`func (o *OraclePdbRestoreParams) SetTargetCdbKeystorePasswordNil(b bool)`

 SetTargetCdbKeystorePasswordNil sets the value for TargetCdbKeystorePassword to be an explicit nil

### UnsetTargetCdbKeystorePassword
`func (o *OraclePdbRestoreParams) UnsetTargetCdbKeystorePassword()`

UnsetTargetCdbKeystorePassword ensures that no value is present for TargetCdbKeystorePassword, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


