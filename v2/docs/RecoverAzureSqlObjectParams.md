# RecoverAzureSqlObjectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewDatabaseName** | Pointer to **NullableString** | Specifies the new name to which the object should be renamed to after the recovery. | [optional] 
**OverwriteDatabase** | Pointer to **NullableBool** | Set to true to overwrite an existing object at the destination. If set to false, and the same object exists at the destination, then recovery will fail for that object. | [optional] 
**RestoredDatabaseSku** | Pointer to [**AzureSqlSkuOptions**](AzureSqlSkuOptions.md) |  | [optional] 
**SqlPackageOptions** | Pointer to [**AzureSqlPackageOptions**](AzureSqlPackageOptions.md) |  | [optional] 

## Methods

### NewRecoverAzureSqlObjectParams

`func NewRecoverAzureSqlObjectParams() *RecoverAzureSqlObjectParams`

NewRecoverAzureSqlObjectParams instantiates a new RecoverAzureSqlObjectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverAzureSqlObjectParamsWithDefaults

`func NewRecoverAzureSqlObjectParamsWithDefaults() *RecoverAzureSqlObjectParams`

NewRecoverAzureSqlObjectParamsWithDefaults instantiates a new RecoverAzureSqlObjectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewDatabaseName

`func (o *RecoverAzureSqlObjectParams) GetNewDatabaseName() string`

GetNewDatabaseName returns the NewDatabaseName field if non-nil, zero value otherwise.

### GetNewDatabaseNameOk

`func (o *RecoverAzureSqlObjectParams) GetNewDatabaseNameOk() (*string, bool)`

GetNewDatabaseNameOk returns a tuple with the NewDatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewDatabaseName

`func (o *RecoverAzureSqlObjectParams) SetNewDatabaseName(v string)`

SetNewDatabaseName sets NewDatabaseName field to given value.

### HasNewDatabaseName

`func (o *RecoverAzureSqlObjectParams) HasNewDatabaseName() bool`

HasNewDatabaseName returns a boolean if a field has been set.

### SetNewDatabaseNameNil

`func (o *RecoverAzureSqlObjectParams) SetNewDatabaseNameNil(b bool)`

 SetNewDatabaseNameNil sets the value for NewDatabaseName to be an explicit nil

### UnsetNewDatabaseName
`func (o *RecoverAzureSqlObjectParams) UnsetNewDatabaseName()`

UnsetNewDatabaseName ensures that no value is present for NewDatabaseName, not even an explicit nil
### GetOverwriteDatabase

`func (o *RecoverAzureSqlObjectParams) GetOverwriteDatabase() bool`

GetOverwriteDatabase returns the OverwriteDatabase field if non-nil, zero value otherwise.

### GetOverwriteDatabaseOk

`func (o *RecoverAzureSqlObjectParams) GetOverwriteDatabaseOk() (*bool, bool)`

GetOverwriteDatabaseOk returns a tuple with the OverwriteDatabase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteDatabase

`func (o *RecoverAzureSqlObjectParams) SetOverwriteDatabase(v bool)`

SetOverwriteDatabase sets OverwriteDatabase field to given value.

### HasOverwriteDatabase

`func (o *RecoverAzureSqlObjectParams) HasOverwriteDatabase() bool`

HasOverwriteDatabase returns a boolean if a field has been set.

### SetOverwriteDatabaseNil

`func (o *RecoverAzureSqlObjectParams) SetOverwriteDatabaseNil(b bool)`

 SetOverwriteDatabaseNil sets the value for OverwriteDatabase to be an explicit nil

### UnsetOverwriteDatabase
`func (o *RecoverAzureSqlObjectParams) UnsetOverwriteDatabase()`

UnsetOverwriteDatabase ensures that no value is present for OverwriteDatabase, not even an explicit nil
### GetRestoredDatabaseSku

`func (o *RecoverAzureSqlObjectParams) GetRestoredDatabaseSku() AzureSqlSkuOptions`

GetRestoredDatabaseSku returns the RestoredDatabaseSku field if non-nil, zero value otherwise.

### GetRestoredDatabaseSkuOk

`func (o *RecoverAzureSqlObjectParams) GetRestoredDatabaseSkuOk() (*AzureSqlSkuOptions, bool)`

GetRestoredDatabaseSkuOk returns a tuple with the RestoredDatabaseSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoredDatabaseSku

`func (o *RecoverAzureSqlObjectParams) SetRestoredDatabaseSku(v AzureSqlSkuOptions)`

SetRestoredDatabaseSku sets RestoredDatabaseSku field to given value.

### HasRestoredDatabaseSku

`func (o *RecoverAzureSqlObjectParams) HasRestoredDatabaseSku() bool`

HasRestoredDatabaseSku returns a boolean if a field has been set.

### GetSqlPackageOptions

`func (o *RecoverAzureSqlObjectParams) GetSqlPackageOptions() AzureSqlPackageOptions`

GetSqlPackageOptions returns the SqlPackageOptions field if non-nil, zero value otherwise.

### GetSqlPackageOptionsOk

`func (o *RecoverAzureSqlObjectParams) GetSqlPackageOptionsOk() (*AzureSqlPackageOptions, bool)`

GetSqlPackageOptionsOk returns a tuple with the SqlPackageOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlPackageOptions

`func (o *RecoverAzureSqlObjectParams) SetSqlPackageOptions(v AzureSqlPackageOptions)`

SetSqlPackageOptions sets SqlPackageOptions field to given value.

### HasSqlPackageOptions

`func (o *RecoverAzureSqlObjectParams) HasSqlPackageOptions() bool`

HasSqlPackageOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


