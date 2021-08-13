# OracleCloneRefreshTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbName** | **NullableString** | Specifies the name of the cloned database. | 
**HomeFolder** | **NullableString** | Specifies the home folder for the cloned database. | 
**BaseFolder** | **NullableString** | Specifies the base folder of Oracle installation on the target host. | 
**Sga** | Pointer to **NullableString** | Specifies the System Global Area (SGA) for the clone database. | [optional] 
**DbVersion** | Pointer to **NullableString** | Specifies the version of the Oracle database. | [optional] 

## Methods

### NewOracleCloneRefreshTask

`func NewOracleCloneRefreshTask(dbName NullableString, homeFolder NullableString, baseFolder NullableString, ) *OracleCloneRefreshTask`

NewOracleCloneRefreshTask instantiates a new OracleCloneRefreshTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleCloneRefreshTaskWithDefaults

`func NewOracleCloneRefreshTaskWithDefaults() *OracleCloneRefreshTask`

NewOracleCloneRefreshTaskWithDefaults instantiates a new OracleCloneRefreshTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbName

`func (o *OracleCloneRefreshTask) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *OracleCloneRefreshTask) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *OracleCloneRefreshTask) SetDbName(v string)`

SetDbName sets DbName field to given value.


### SetDbNameNil

`func (o *OracleCloneRefreshTask) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *OracleCloneRefreshTask) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil
### GetHomeFolder

`func (o *OracleCloneRefreshTask) GetHomeFolder() string`

GetHomeFolder returns the HomeFolder field if non-nil, zero value otherwise.

### GetHomeFolderOk

`func (o *OracleCloneRefreshTask) GetHomeFolderOk() (*string, bool)`

GetHomeFolderOk returns a tuple with the HomeFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeFolder

`func (o *OracleCloneRefreshTask) SetHomeFolder(v string)`

SetHomeFolder sets HomeFolder field to given value.


### SetHomeFolderNil

`func (o *OracleCloneRefreshTask) SetHomeFolderNil(b bool)`

 SetHomeFolderNil sets the value for HomeFolder to be an explicit nil

### UnsetHomeFolder
`func (o *OracleCloneRefreshTask) UnsetHomeFolder()`

UnsetHomeFolder ensures that no value is present for HomeFolder, not even an explicit nil
### GetBaseFolder

`func (o *OracleCloneRefreshTask) GetBaseFolder() string`

GetBaseFolder returns the BaseFolder field if non-nil, zero value otherwise.

### GetBaseFolderOk

`func (o *OracleCloneRefreshTask) GetBaseFolderOk() (*string, bool)`

GetBaseFolderOk returns a tuple with the BaseFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseFolder

`func (o *OracleCloneRefreshTask) SetBaseFolder(v string)`

SetBaseFolder sets BaseFolder field to given value.


### SetBaseFolderNil

`func (o *OracleCloneRefreshTask) SetBaseFolderNil(b bool)`

 SetBaseFolderNil sets the value for BaseFolder to be an explicit nil

### UnsetBaseFolder
`func (o *OracleCloneRefreshTask) UnsetBaseFolder()`

UnsetBaseFolder ensures that no value is present for BaseFolder, not even an explicit nil
### GetSga

`func (o *OracleCloneRefreshTask) GetSga() string`

GetSga returns the Sga field if non-nil, zero value otherwise.

### GetSgaOk

`func (o *OracleCloneRefreshTask) GetSgaOk() (*string, bool)`

GetSgaOk returns a tuple with the Sga field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSga

`func (o *OracleCloneRefreshTask) SetSga(v string)`

SetSga sets Sga field to given value.

### HasSga

`func (o *OracleCloneRefreshTask) HasSga() bool`

HasSga returns a boolean if a field has been set.

### SetSgaNil

`func (o *OracleCloneRefreshTask) SetSgaNil(b bool)`

 SetSgaNil sets the value for Sga to be an explicit nil

### UnsetSga
`func (o *OracleCloneRefreshTask) UnsetSga()`

UnsetSga ensures that no value is present for Sga, not even an explicit nil
### GetDbVersion

`func (o *OracleCloneRefreshTask) GetDbVersion() string`

GetDbVersion returns the DbVersion field if non-nil, zero value otherwise.

### GetDbVersionOk

`func (o *OracleCloneRefreshTask) GetDbVersionOk() (*string, bool)`

GetDbVersionOk returns a tuple with the DbVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbVersion

`func (o *OracleCloneRefreshTask) SetDbVersion(v string)`

SetDbVersion sets DbVersion field to given value.

### HasDbVersion

`func (o *OracleCloneRefreshTask) HasDbVersion() bool`

HasDbVersion returns a boolean if a field has been set.

### SetDbVersionNil

`func (o *OracleCloneRefreshTask) SetDbVersionNil(b bool)`

 SetDbVersionNil sets the value for DbVersion to be an explicit nil

### UnsetDbVersion
`func (o *OracleCloneRefreshTask) UnsetDbVersion()`

UnsetDbVersion ensures that no value is present for DbVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


