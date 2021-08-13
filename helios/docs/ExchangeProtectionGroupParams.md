# ExchangeProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | [**[]ExchangeProtectionGroupObjectParams**](ExchangeProtectionGroupObjectParams.md) | Specifies the list of object ids to be protected. | 
**ExcludeDatabaseIds** | Pointer to **[]int64** | Specifies the list of IDs of the databases to not be protected by this Protection Group. This can be used to ignore specific databases under Exchange Server/DAG which has been included for protection. | [optional] 
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 
**BackupsCopyOnly** | Pointer to **NullableBool** | Specifies whether the backups should be copy-only. | [optional] 

## Methods

### NewExchangeProtectionGroupParams

`func NewExchangeProtectionGroupParams(objects []ExchangeProtectionGroupObjectParams, ) *ExchangeProtectionGroupParams`

NewExchangeProtectionGroupParams instantiates a new ExchangeProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeProtectionGroupParamsWithDefaults

`func NewExchangeProtectionGroupParamsWithDefaults() *ExchangeProtectionGroupParams`

NewExchangeProtectionGroupParamsWithDefaults instantiates a new ExchangeProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *ExchangeProtectionGroupParams) GetObjects() []ExchangeProtectionGroupObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *ExchangeProtectionGroupParams) GetObjectsOk() (*[]ExchangeProtectionGroupObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *ExchangeProtectionGroupParams) SetObjects(v []ExchangeProtectionGroupObjectParams)`

SetObjects sets Objects field to given value.


### GetExcludeDatabaseIds

`func (o *ExchangeProtectionGroupParams) GetExcludeDatabaseIds() []int64`

GetExcludeDatabaseIds returns the ExcludeDatabaseIds field if non-nil, zero value otherwise.

### GetExcludeDatabaseIdsOk

`func (o *ExchangeProtectionGroupParams) GetExcludeDatabaseIdsOk() (*[]int64, bool)`

GetExcludeDatabaseIdsOk returns a tuple with the ExcludeDatabaseIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeDatabaseIds

`func (o *ExchangeProtectionGroupParams) SetExcludeDatabaseIds(v []int64)`

SetExcludeDatabaseIds sets ExcludeDatabaseIds field to given value.

### HasExcludeDatabaseIds

`func (o *ExchangeProtectionGroupParams) HasExcludeDatabaseIds() bool`

HasExcludeDatabaseIds returns a boolean if a field has been set.

### GetIndexingPolicy

`func (o *ExchangeProtectionGroupParams) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *ExchangeProtectionGroupParams) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *ExchangeProtectionGroupParams) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *ExchangeProtectionGroupParams) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetBackupsCopyOnly

`func (o *ExchangeProtectionGroupParams) GetBackupsCopyOnly() bool`

GetBackupsCopyOnly returns the BackupsCopyOnly field if non-nil, zero value otherwise.

### GetBackupsCopyOnlyOk

`func (o *ExchangeProtectionGroupParams) GetBackupsCopyOnlyOk() (*bool, bool)`

GetBackupsCopyOnlyOk returns a tuple with the BackupsCopyOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupsCopyOnly

`func (o *ExchangeProtectionGroupParams) SetBackupsCopyOnly(v bool)`

SetBackupsCopyOnly sets BackupsCopyOnly field to given value.

### HasBackupsCopyOnly

`func (o *ExchangeProtectionGroupParams) HasBackupsCopyOnly() bool`

HasBackupsCopyOnly returns a boolean if a field has been set.

### SetBackupsCopyOnlyNil

`func (o *ExchangeProtectionGroupParams) SetBackupsCopyOnlyNil(b bool)`

 SetBackupsCopyOnlyNil sets the value for BackupsCopyOnly to be an explicit nil

### UnsetBackupsCopyOnly
`func (o *ExchangeProtectionGroupParams) UnsetBackupsCopyOnly()`

UnsetBackupsCopyOnly ensures that no value is present for BackupsCopyOnly, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


