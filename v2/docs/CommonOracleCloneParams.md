# CommonOracleCloneParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseFolder** | **NullableString** | Specifies the base folder of Oracle installation on the target host. | 
**DbName** | **NullableString** | Specifies the name of the cloned database. | 
**HomeFolder** | **NullableString** | Specifies the home folder for the cloned database. | 
**Sga** | Pointer to **NullableString** | Specifies the System Global Area (SGA) for the clone database. | [optional] 

## Methods

### NewCommonOracleCloneParams

`func NewCommonOracleCloneParams(baseFolder NullableString, dbName NullableString, homeFolder NullableString, ) *CommonOracleCloneParams`

NewCommonOracleCloneParams instantiates a new CommonOracleCloneParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonOracleCloneParamsWithDefaults

`func NewCommonOracleCloneParamsWithDefaults() *CommonOracleCloneParams`

NewCommonOracleCloneParamsWithDefaults instantiates a new CommonOracleCloneParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseFolder

`func (o *CommonOracleCloneParams) GetBaseFolder() string`

GetBaseFolder returns the BaseFolder field if non-nil, zero value otherwise.

### GetBaseFolderOk

`func (o *CommonOracleCloneParams) GetBaseFolderOk() (*string, bool)`

GetBaseFolderOk returns a tuple with the BaseFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseFolder

`func (o *CommonOracleCloneParams) SetBaseFolder(v string)`

SetBaseFolder sets BaseFolder field to given value.


### SetBaseFolderNil

`func (o *CommonOracleCloneParams) SetBaseFolderNil(b bool)`

 SetBaseFolderNil sets the value for BaseFolder to be an explicit nil

### UnsetBaseFolder
`func (o *CommonOracleCloneParams) UnsetBaseFolder()`

UnsetBaseFolder ensures that no value is present for BaseFolder, not even an explicit nil
### GetDbName

`func (o *CommonOracleCloneParams) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *CommonOracleCloneParams) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *CommonOracleCloneParams) SetDbName(v string)`

SetDbName sets DbName field to given value.


### SetDbNameNil

`func (o *CommonOracleCloneParams) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *CommonOracleCloneParams) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil
### GetHomeFolder

`func (o *CommonOracleCloneParams) GetHomeFolder() string`

GetHomeFolder returns the HomeFolder field if non-nil, zero value otherwise.

### GetHomeFolderOk

`func (o *CommonOracleCloneParams) GetHomeFolderOk() (*string, bool)`

GetHomeFolderOk returns a tuple with the HomeFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeFolder

`func (o *CommonOracleCloneParams) SetHomeFolder(v string)`

SetHomeFolder sets HomeFolder field to given value.


### SetHomeFolderNil

`func (o *CommonOracleCloneParams) SetHomeFolderNil(b bool)`

 SetHomeFolderNil sets the value for HomeFolder to be an explicit nil

### UnsetHomeFolder
`func (o *CommonOracleCloneParams) UnsetHomeFolder()`

UnsetHomeFolder ensures that no value is present for HomeFolder, not even an explicit nil
### GetSga

`func (o *CommonOracleCloneParams) GetSga() string`

GetSga returns the Sga field if non-nil, zero value otherwise.

### GetSgaOk

`func (o *CommonOracleCloneParams) GetSgaOk() (*string, bool)`

GetSgaOk returns a tuple with the Sga field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSga

`func (o *CommonOracleCloneParams) SetSga(v string)`

SetSga sets Sga field to given value.

### HasSga

`func (o *CommonOracleCloneParams) HasSga() bool`

HasSga returns a boolean if a field has been set.

### SetSgaNil

`func (o *CommonOracleCloneParams) SetSgaNil(b bool)`

 SetSgaNil sets the value for Sga to be an explicit nil

### UnsetSga
`func (o *CommonOracleCloneParams) UnsetSga()`

UnsetSga ensures that no value is present for Sga, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


