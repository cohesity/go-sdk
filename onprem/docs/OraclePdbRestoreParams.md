# OraclePdbRestoreParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DropDuplicatePDB** | Pointer to **NullableBool** | Specifies if the PDB should be ignored if a PDB already exists with same name. | [optional] 
**PdbObjects** | Pointer to [**[]OraclePdbObjectInfo**](OraclePdbObjectInfo.md) | Specifies list of PDB objects to restore. | [optional] 
**RestoreToExistingCdb** | Pointer to **NullableBool** | Specifies if pdbs should be restored to an existing CDB. | [optional] 

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


