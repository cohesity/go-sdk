# AzureSqlObjectProtectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CopyDatabase** | Pointer to **NullableBool** | If set to true, a copy of the database is created during backup, and the backup is performed from the copied database. This backup will be transactionally consistent. If set to false, the backup is performed from the production database while transactions are in progress. In this case, the backup will be transactionally inconsistent, and recovery can fail or the recovered database may be in an inconsistent state. | [optional] 
**CopyDatabaseSku** | Pointer to [**AzureSqlSkuOptions**](AzureSqlSkuOptions.md) |  | [optional] 
**DiskType** | Pointer to **NullableString** | Specifies azure managed storage disk to be used for object protection. By default Premium LRS is being used to support Azure SQL workloads. | [optional] 
**ExcludeObjectIds** | Pointer to **[]int64** | Specifies the ids of the objects to be excluded in the Object Protection. This can be used to ignore specific objects under a parent object which has been included for protection. | [optional] 
**Objects** | Pointer to [**[]AzureObjectLevelParams**](AzureObjectLevelParams.md) | Specifies the objects to be protected. | [optional] 
**SqlPackageOptions** | Pointer to [**AzureSqlPackageOptions**](AzureSqlPackageOptions.md) |  | [optional] 
**TempDiskSizeGb** | Pointer to **NullableInt32** | Specifies the size of the disk we will attach to rigel to use for exporting this DB(in GB). | [optional] 

## Methods

### NewAzureSqlObjectProtectionParams

`func NewAzureSqlObjectProtectionParams() *AzureSqlObjectProtectionParams`

NewAzureSqlObjectProtectionParams instantiates a new AzureSqlObjectProtectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureSqlObjectProtectionParamsWithDefaults

`func NewAzureSqlObjectProtectionParamsWithDefaults() *AzureSqlObjectProtectionParams`

NewAzureSqlObjectProtectionParamsWithDefaults instantiates a new AzureSqlObjectProtectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCopyDatabase

`func (o *AzureSqlObjectProtectionParams) GetCopyDatabase() bool`

GetCopyDatabase returns the CopyDatabase field if non-nil, zero value otherwise.

### GetCopyDatabaseOk

`func (o *AzureSqlObjectProtectionParams) GetCopyDatabaseOk() (*bool, bool)`

GetCopyDatabaseOk returns a tuple with the CopyDatabase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyDatabase

`func (o *AzureSqlObjectProtectionParams) SetCopyDatabase(v bool)`

SetCopyDatabase sets CopyDatabase field to given value.

### HasCopyDatabase

`func (o *AzureSqlObjectProtectionParams) HasCopyDatabase() bool`

HasCopyDatabase returns a boolean if a field has been set.

### SetCopyDatabaseNil

`func (o *AzureSqlObjectProtectionParams) SetCopyDatabaseNil(b bool)`

 SetCopyDatabaseNil sets the value for CopyDatabase to be an explicit nil

### UnsetCopyDatabase
`func (o *AzureSqlObjectProtectionParams) UnsetCopyDatabase()`

UnsetCopyDatabase ensures that no value is present for CopyDatabase, not even an explicit nil
### GetCopyDatabaseSku

`func (o *AzureSqlObjectProtectionParams) GetCopyDatabaseSku() AzureSqlSkuOptions`

GetCopyDatabaseSku returns the CopyDatabaseSku field if non-nil, zero value otherwise.

### GetCopyDatabaseSkuOk

`func (o *AzureSqlObjectProtectionParams) GetCopyDatabaseSkuOk() (*AzureSqlSkuOptions, bool)`

GetCopyDatabaseSkuOk returns a tuple with the CopyDatabaseSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyDatabaseSku

`func (o *AzureSqlObjectProtectionParams) SetCopyDatabaseSku(v AzureSqlSkuOptions)`

SetCopyDatabaseSku sets CopyDatabaseSku field to given value.

### HasCopyDatabaseSku

`func (o *AzureSqlObjectProtectionParams) HasCopyDatabaseSku() bool`

HasCopyDatabaseSku returns a boolean if a field has been set.

### GetDiskType

`func (o *AzureSqlObjectProtectionParams) GetDiskType() string`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *AzureSqlObjectProtectionParams) GetDiskTypeOk() (*string, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *AzureSqlObjectProtectionParams) SetDiskType(v string)`

SetDiskType sets DiskType field to given value.

### HasDiskType

`func (o *AzureSqlObjectProtectionParams) HasDiskType() bool`

HasDiskType returns a boolean if a field has been set.

### SetDiskTypeNil

`func (o *AzureSqlObjectProtectionParams) SetDiskTypeNil(b bool)`

 SetDiskTypeNil sets the value for DiskType to be an explicit nil

### UnsetDiskType
`func (o *AzureSqlObjectProtectionParams) UnsetDiskType()`

UnsetDiskType ensures that no value is present for DiskType, not even an explicit nil
### GetExcludeObjectIds

`func (o *AzureSqlObjectProtectionParams) GetExcludeObjectIds() []int64`

GetExcludeObjectIds returns the ExcludeObjectIds field if non-nil, zero value otherwise.

### GetExcludeObjectIdsOk

`func (o *AzureSqlObjectProtectionParams) GetExcludeObjectIdsOk() (*[]int64, bool)`

GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeObjectIds

`func (o *AzureSqlObjectProtectionParams) SetExcludeObjectIds(v []int64)`

SetExcludeObjectIds sets ExcludeObjectIds field to given value.

### HasExcludeObjectIds

`func (o *AzureSqlObjectProtectionParams) HasExcludeObjectIds() bool`

HasExcludeObjectIds returns a boolean if a field has been set.

### SetExcludeObjectIdsNil

`func (o *AzureSqlObjectProtectionParams) SetExcludeObjectIdsNil(b bool)`

 SetExcludeObjectIdsNil sets the value for ExcludeObjectIds to be an explicit nil

### UnsetExcludeObjectIds
`func (o *AzureSqlObjectProtectionParams) UnsetExcludeObjectIds()`

UnsetExcludeObjectIds ensures that no value is present for ExcludeObjectIds, not even an explicit nil
### GetObjects

`func (o *AzureSqlObjectProtectionParams) GetObjects() []AzureObjectLevelParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *AzureSqlObjectProtectionParams) GetObjectsOk() (*[]AzureObjectLevelParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *AzureSqlObjectProtectionParams) SetObjects(v []AzureObjectLevelParams)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *AzureSqlObjectProtectionParams) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetSqlPackageOptions

`func (o *AzureSqlObjectProtectionParams) GetSqlPackageOptions() AzureSqlPackageOptions`

GetSqlPackageOptions returns the SqlPackageOptions field if non-nil, zero value otherwise.

### GetSqlPackageOptionsOk

`func (o *AzureSqlObjectProtectionParams) GetSqlPackageOptionsOk() (*AzureSqlPackageOptions, bool)`

GetSqlPackageOptionsOk returns a tuple with the SqlPackageOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlPackageOptions

`func (o *AzureSqlObjectProtectionParams) SetSqlPackageOptions(v AzureSqlPackageOptions)`

SetSqlPackageOptions sets SqlPackageOptions field to given value.

### HasSqlPackageOptions

`func (o *AzureSqlObjectProtectionParams) HasSqlPackageOptions() bool`

HasSqlPackageOptions returns a boolean if a field has been set.

### GetTempDiskSizeGb

`func (o *AzureSqlObjectProtectionParams) GetTempDiskSizeGb() int32`

GetTempDiskSizeGb returns the TempDiskSizeGb field if non-nil, zero value otherwise.

### GetTempDiskSizeGbOk

`func (o *AzureSqlObjectProtectionParams) GetTempDiskSizeGbOk() (*int32, bool)`

GetTempDiskSizeGbOk returns a tuple with the TempDiskSizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempDiskSizeGb

`func (o *AzureSqlObjectProtectionParams) SetTempDiskSizeGb(v int32)`

SetTempDiskSizeGb sets TempDiskSizeGb field to given value.

### HasTempDiskSizeGb

`func (o *AzureSqlObjectProtectionParams) HasTempDiskSizeGb() bool`

HasTempDiskSizeGb returns a boolean if a field has been set.

### SetTempDiskSizeGbNil

`func (o *AzureSqlObjectProtectionParams) SetTempDiskSizeGbNil(b bool)`

 SetTempDiskSizeGbNil sets the value for TempDiskSizeGb to be an explicit nil

### UnsetTempDiskSizeGb
`func (o *AzureSqlObjectProtectionParams) UnsetTempDiskSizeGb()`

UnsetTempDiskSizeGb ensures that no value is present for TempDiskSizeGb, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


