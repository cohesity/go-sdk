# OracleCloneTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreScript** | Pointer to [**CommonPreBackupScriptParams**](CommonPreBackupScriptParams.md) |  | [optional] 
**PostScript** | Pointer to [**CommonPostBackupScriptParams**](CommonPostBackupScriptParams.md) |  | [optional] 
**DbName** | **NullableString** | Specifies the name of the cloned database. | 
**HomeFolder** | **NullableString** | Specifies the home folder for the cloned database. | 
**BaseFolder** | **NullableString** | Specifies the base folder of Oracle installation on the target host. | 
**Sga** | Pointer to **NullableString** | Specifies the System Global Area (SGA) for the clone database. | [optional] 

## Methods

### NewOracleCloneTask

`func NewOracleCloneTask(dbName NullableString, homeFolder NullableString, baseFolder NullableString, ) *OracleCloneTask`

NewOracleCloneTask instantiates a new OracleCloneTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleCloneTaskWithDefaults

`func NewOracleCloneTaskWithDefaults() *OracleCloneTask`

NewOracleCloneTaskWithDefaults instantiates a new OracleCloneTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreScript

`func (o *OracleCloneTask) GetPreScript() CommonPreBackupScriptParams`

GetPreScript returns the PreScript field if non-nil, zero value otherwise.

### GetPreScriptOk

`func (o *OracleCloneTask) GetPreScriptOk() (*CommonPreBackupScriptParams, bool)`

GetPreScriptOk returns a tuple with the PreScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreScript

`func (o *OracleCloneTask) SetPreScript(v CommonPreBackupScriptParams)`

SetPreScript sets PreScript field to given value.

### HasPreScript

`func (o *OracleCloneTask) HasPreScript() bool`

HasPreScript returns a boolean if a field has been set.

### GetPostScript

`func (o *OracleCloneTask) GetPostScript() CommonPostBackupScriptParams`

GetPostScript returns the PostScript field if non-nil, zero value otherwise.

### GetPostScriptOk

`func (o *OracleCloneTask) GetPostScriptOk() (*CommonPostBackupScriptParams, bool)`

GetPostScriptOk returns a tuple with the PostScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostScript

`func (o *OracleCloneTask) SetPostScript(v CommonPostBackupScriptParams)`

SetPostScript sets PostScript field to given value.

### HasPostScript

`func (o *OracleCloneTask) HasPostScript() bool`

HasPostScript returns a boolean if a field has been set.

### GetDbName

`func (o *OracleCloneTask) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *OracleCloneTask) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *OracleCloneTask) SetDbName(v string)`

SetDbName sets DbName field to given value.


### SetDbNameNil

`func (o *OracleCloneTask) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *OracleCloneTask) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil
### GetHomeFolder

`func (o *OracleCloneTask) GetHomeFolder() string`

GetHomeFolder returns the HomeFolder field if non-nil, zero value otherwise.

### GetHomeFolderOk

`func (o *OracleCloneTask) GetHomeFolderOk() (*string, bool)`

GetHomeFolderOk returns a tuple with the HomeFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeFolder

`func (o *OracleCloneTask) SetHomeFolder(v string)`

SetHomeFolder sets HomeFolder field to given value.


### SetHomeFolderNil

`func (o *OracleCloneTask) SetHomeFolderNil(b bool)`

 SetHomeFolderNil sets the value for HomeFolder to be an explicit nil

### UnsetHomeFolder
`func (o *OracleCloneTask) UnsetHomeFolder()`

UnsetHomeFolder ensures that no value is present for HomeFolder, not even an explicit nil
### GetBaseFolder

`func (o *OracleCloneTask) GetBaseFolder() string`

GetBaseFolder returns the BaseFolder field if non-nil, zero value otherwise.

### GetBaseFolderOk

`func (o *OracleCloneTask) GetBaseFolderOk() (*string, bool)`

GetBaseFolderOk returns a tuple with the BaseFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseFolder

`func (o *OracleCloneTask) SetBaseFolder(v string)`

SetBaseFolder sets BaseFolder field to given value.


### SetBaseFolderNil

`func (o *OracleCloneTask) SetBaseFolderNil(b bool)`

 SetBaseFolderNil sets the value for BaseFolder to be an explicit nil

### UnsetBaseFolder
`func (o *OracleCloneTask) UnsetBaseFolder()`

UnsetBaseFolder ensures that no value is present for BaseFolder, not even an explicit nil
### GetSga

`func (o *OracleCloneTask) GetSga() string`

GetSga returns the Sga field if non-nil, zero value otherwise.

### GetSgaOk

`func (o *OracleCloneTask) GetSgaOk() (*string, bool)`

GetSgaOk returns a tuple with the Sga field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSga

`func (o *OracleCloneTask) SetSga(v string)`

SetSga sets Sga field to given value.

### HasSga

`func (o *OracleCloneTask) HasSga() bool`

HasSga returns a boolean if a field has been set.

### SetSgaNil

`func (o *OracleCloneTask) SetSgaNil(b bool)`

 SetSgaNil sets the value for Sga to be an explicit nil

### UnsetSga
`func (o *OracleCloneTask) UnsetSga()`

UnsetSga ensures that no value is present for Sga, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


