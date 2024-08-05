# OracleCloneObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseFolder** | **NullableString** | Specifies the base folder of Oracle installation on the target host. | 
**DbName** | **NullableString** | Specifies the name of the cloned database. | 
**HomeFolder** | **NullableString** | Specifies the home folder for the cloned database. | 
**Sga** | Pointer to **NullableString** | Specifies the System Global Area (SGA) for the clone database. | [optional] 
**DbVersion** | Pointer to **NullableString** | Specifies the version of the Oracle database. | [optional] 

## Methods

### NewOracleCloneObject

`func NewOracleCloneObject(baseFolder NullableString, dbName NullableString, homeFolder NullableString, ) *OracleCloneObject`

NewOracleCloneObject instantiates a new OracleCloneObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleCloneObjectWithDefaults

`func NewOracleCloneObjectWithDefaults() *OracleCloneObject`

NewOracleCloneObjectWithDefaults instantiates a new OracleCloneObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseFolder

`func (o *OracleCloneObject) GetBaseFolder() string`

GetBaseFolder returns the BaseFolder field if non-nil, zero value otherwise.

### GetBaseFolderOk

`func (o *OracleCloneObject) GetBaseFolderOk() (*string, bool)`

GetBaseFolderOk returns a tuple with the BaseFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseFolder

`func (o *OracleCloneObject) SetBaseFolder(v string)`

SetBaseFolder sets BaseFolder field to given value.


### SetBaseFolderNil

`func (o *OracleCloneObject) SetBaseFolderNil(b bool)`

 SetBaseFolderNil sets the value for BaseFolder to be an explicit nil

### UnsetBaseFolder
`func (o *OracleCloneObject) UnsetBaseFolder()`

UnsetBaseFolder ensures that no value is present for BaseFolder, not even an explicit nil
### GetDbName

`func (o *OracleCloneObject) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *OracleCloneObject) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *OracleCloneObject) SetDbName(v string)`

SetDbName sets DbName field to given value.


### SetDbNameNil

`func (o *OracleCloneObject) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *OracleCloneObject) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil
### GetHomeFolder

`func (o *OracleCloneObject) GetHomeFolder() string`

GetHomeFolder returns the HomeFolder field if non-nil, zero value otherwise.

### GetHomeFolderOk

`func (o *OracleCloneObject) GetHomeFolderOk() (*string, bool)`

GetHomeFolderOk returns a tuple with the HomeFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeFolder

`func (o *OracleCloneObject) SetHomeFolder(v string)`

SetHomeFolder sets HomeFolder field to given value.


### SetHomeFolderNil

`func (o *OracleCloneObject) SetHomeFolderNil(b bool)`

 SetHomeFolderNil sets the value for HomeFolder to be an explicit nil

### UnsetHomeFolder
`func (o *OracleCloneObject) UnsetHomeFolder()`

UnsetHomeFolder ensures that no value is present for HomeFolder, not even an explicit nil
### GetSga

`func (o *OracleCloneObject) GetSga() string`

GetSga returns the Sga field if non-nil, zero value otherwise.

### GetSgaOk

`func (o *OracleCloneObject) GetSgaOk() (*string, bool)`

GetSgaOk returns a tuple with the Sga field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSga

`func (o *OracleCloneObject) SetSga(v string)`

SetSga sets Sga field to given value.

### HasSga

`func (o *OracleCloneObject) HasSga() bool`

HasSga returns a boolean if a field has been set.

### SetSgaNil

`func (o *OracleCloneObject) SetSgaNil(b bool)`

 SetSgaNil sets the value for Sga to be an explicit nil

### UnsetSga
`func (o *OracleCloneObject) UnsetSga()`

UnsetSga ensures that no value is present for Sga, not even an explicit nil
### GetDbVersion

`func (o *OracleCloneObject) GetDbVersion() string`

GetDbVersion returns the DbVersion field if non-nil, zero value otherwise.

### GetDbVersionOk

`func (o *OracleCloneObject) GetDbVersionOk() (*string, bool)`

GetDbVersionOk returns a tuple with the DbVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbVersion

`func (o *OracleCloneObject) SetDbVersion(v string)`

SetDbVersion sets DbVersion field to given value.

### HasDbVersion

`func (o *OracleCloneObject) HasDbVersion() bool`

HasDbVersion returns a boolean if a field has been set.

### SetDbVersionNil

`func (o *OracleCloneObject) SetDbVersionNil(b bool)`

 SetDbVersionNil sets the value for DbVersion to be an explicit nil

### UnsetDbVersion
`func (o *OracleCloneObject) UnsetDbVersion()`

UnsetDbVersion ensures that no value is present for DbVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


