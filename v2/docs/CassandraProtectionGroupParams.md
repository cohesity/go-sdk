# CassandraProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BandwidthMBPS** | Pointer to **NullableInt64** | Specifies the maximum network bandwidth that each concurrent IO Stream can use for exchanging data with the cluster. | [optional] 
**Concurrency** | Pointer to **NullableInt32** | Specifies the maximum number of concurrent IO Streams that will be created to exchange data with the cluster. | [optional] 
**CustomSourceName** | Pointer to **NullableString** | The user specified name for the Source on which this protection was run. | [optional] [readonly] 
**ExcludeObjectIds** | Pointer to **[]int64** | Specifies the objects to be excluded in the Protection Group. | [optional] 
**ExcludeObjectlist** | Pointer to **[]string** | Specifies the list of fully qualified name of the entities to exclude for protection. | [optional] 
**IncludeObjectlist** | Pointer to **[]string** | Specifies the list of fully qualified name of the entities to include for protection. | [optional] 
**Objects** | Pointer to [**[]NoSqlProtectionGroupObjectParams**](NoSqlProtectionGroupObjectParams.md) | Specifies the objects to be included in the Protection Group. | [optional] 
**OverwriteExcludeObjectlist** | Pointer to **NullableBool** | If disabled - The excludeObjectlist is merged with the existing exclude_sources_vec, preserving any existing elements while incorporating new ones. | [optional] [default to true]
**OverwriteIncludeObjectlist** | Pointer to **NullableBool** | If disabled - The includeObjectlist is merged with the existing sources_vec, preserving any existing elements while incorporating new ones. | [optional] [default to true]
**SourceId** | Pointer to **NullableInt64** | Object ID of the Source on which this protection was run . | [optional] [readonly] 
**SourceName** | Pointer to **NullableString** | Specifies the name of the Source on which this protection was run. | [optional] [readonly] 
**DataCenters** | Pointer to **[]string** | Only the specified data centers will be considered while taking backup. The keyspaces having replication strategy &#39;Simple&#39; can be backed up only if all the datacenters for the cassandra cluster are specified. For any keyspace having replication strategy as &#39;Network&#39;, all the associated data centers should be specified. | [optional] 
**IsLogBackup** | Pointer to **NullableBool** | Specifies the type of job for Cassandra. If true, only log backup job will be scheduled for the source. This requires a policy with log Backup option enabled. | [optional] 
**IsSystemKeyspaceBackup** | Pointer to **NullableBool** | Specifies whether this ia a system keyspace backup job. | [optional] 

## Methods

### NewCassandraProtectionGroupParams

`func NewCassandraProtectionGroupParams() *CassandraProtectionGroupParams`

NewCassandraProtectionGroupParams instantiates a new CassandraProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCassandraProtectionGroupParamsWithDefaults

`func NewCassandraProtectionGroupParamsWithDefaults() *CassandraProtectionGroupParams`

NewCassandraProtectionGroupParamsWithDefaults instantiates a new CassandraProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandwidthMBPS

`func (o *CassandraProtectionGroupParams) GetBandwidthMBPS() int64`

GetBandwidthMBPS returns the BandwidthMBPS field if non-nil, zero value otherwise.

### GetBandwidthMBPSOk

`func (o *CassandraProtectionGroupParams) GetBandwidthMBPSOk() (*int64, bool)`

GetBandwidthMBPSOk returns a tuple with the BandwidthMBPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthMBPS

`func (o *CassandraProtectionGroupParams) SetBandwidthMBPS(v int64)`

SetBandwidthMBPS sets BandwidthMBPS field to given value.

### HasBandwidthMBPS

`func (o *CassandraProtectionGroupParams) HasBandwidthMBPS() bool`

HasBandwidthMBPS returns a boolean if a field has been set.

### SetBandwidthMBPSNil

`func (o *CassandraProtectionGroupParams) SetBandwidthMBPSNil(b bool)`

 SetBandwidthMBPSNil sets the value for BandwidthMBPS to be an explicit nil

### UnsetBandwidthMBPS
`func (o *CassandraProtectionGroupParams) UnsetBandwidthMBPS()`

UnsetBandwidthMBPS ensures that no value is present for BandwidthMBPS, not even an explicit nil
### GetConcurrency

`func (o *CassandraProtectionGroupParams) GetConcurrency() int32`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *CassandraProtectionGroupParams) GetConcurrencyOk() (*int32, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *CassandraProtectionGroupParams) SetConcurrency(v int32)`

SetConcurrency sets Concurrency field to given value.

### HasConcurrency

`func (o *CassandraProtectionGroupParams) HasConcurrency() bool`

HasConcurrency returns a boolean if a field has been set.

### SetConcurrencyNil

`func (o *CassandraProtectionGroupParams) SetConcurrencyNil(b bool)`

 SetConcurrencyNil sets the value for Concurrency to be an explicit nil

### UnsetConcurrency
`func (o *CassandraProtectionGroupParams) UnsetConcurrency()`

UnsetConcurrency ensures that no value is present for Concurrency, not even an explicit nil
### GetCustomSourceName

`func (o *CassandraProtectionGroupParams) GetCustomSourceName() string`

GetCustomSourceName returns the CustomSourceName field if non-nil, zero value otherwise.

### GetCustomSourceNameOk

`func (o *CassandraProtectionGroupParams) GetCustomSourceNameOk() (*string, bool)`

GetCustomSourceNameOk returns a tuple with the CustomSourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSourceName

`func (o *CassandraProtectionGroupParams) SetCustomSourceName(v string)`

SetCustomSourceName sets CustomSourceName field to given value.

### HasCustomSourceName

`func (o *CassandraProtectionGroupParams) HasCustomSourceName() bool`

HasCustomSourceName returns a boolean if a field has been set.

### SetCustomSourceNameNil

`func (o *CassandraProtectionGroupParams) SetCustomSourceNameNil(b bool)`

 SetCustomSourceNameNil sets the value for CustomSourceName to be an explicit nil

### UnsetCustomSourceName
`func (o *CassandraProtectionGroupParams) UnsetCustomSourceName()`

UnsetCustomSourceName ensures that no value is present for CustomSourceName, not even an explicit nil
### GetExcludeObjectIds

`func (o *CassandraProtectionGroupParams) GetExcludeObjectIds() []int64`

GetExcludeObjectIds returns the ExcludeObjectIds field if non-nil, zero value otherwise.

### GetExcludeObjectIdsOk

`func (o *CassandraProtectionGroupParams) GetExcludeObjectIdsOk() (*[]int64, bool)`

GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeObjectIds

`func (o *CassandraProtectionGroupParams) SetExcludeObjectIds(v []int64)`

SetExcludeObjectIds sets ExcludeObjectIds field to given value.

### HasExcludeObjectIds

`func (o *CassandraProtectionGroupParams) HasExcludeObjectIds() bool`

HasExcludeObjectIds returns a boolean if a field has been set.

### SetExcludeObjectIdsNil

`func (o *CassandraProtectionGroupParams) SetExcludeObjectIdsNil(b bool)`

 SetExcludeObjectIdsNil sets the value for ExcludeObjectIds to be an explicit nil

### UnsetExcludeObjectIds
`func (o *CassandraProtectionGroupParams) UnsetExcludeObjectIds()`

UnsetExcludeObjectIds ensures that no value is present for ExcludeObjectIds, not even an explicit nil
### GetExcludeObjectlist

`func (o *CassandraProtectionGroupParams) GetExcludeObjectlist() []string`

GetExcludeObjectlist returns the ExcludeObjectlist field if non-nil, zero value otherwise.

### GetExcludeObjectlistOk

`func (o *CassandraProtectionGroupParams) GetExcludeObjectlistOk() (*[]string, bool)`

GetExcludeObjectlistOk returns a tuple with the ExcludeObjectlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeObjectlist

`func (o *CassandraProtectionGroupParams) SetExcludeObjectlist(v []string)`

SetExcludeObjectlist sets ExcludeObjectlist field to given value.

### HasExcludeObjectlist

`func (o *CassandraProtectionGroupParams) HasExcludeObjectlist() bool`

HasExcludeObjectlist returns a boolean if a field has been set.

### SetExcludeObjectlistNil

`func (o *CassandraProtectionGroupParams) SetExcludeObjectlistNil(b bool)`

 SetExcludeObjectlistNil sets the value for ExcludeObjectlist to be an explicit nil

### UnsetExcludeObjectlist
`func (o *CassandraProtectionGroupParams) UnsetExcludeObjectlist()`

UnsetExcludeObjectlist ensures that no value is present for ExcludeObjectlist, not even an explicit nil
### GetIncludeObjectlist

`func (o *CassandraProtectionGroupParams) GetIncludeObjectlist() []string`

GetIncludeObjectlist returns the IncludeObjectlist field if non-nil, zero value otherwise.

### GetIncludeObjectlistOk

`func (o *CassandraProtectionGroupParams) GetIncludeObjectlistOk() (*[]string, bool)`

GetIncludeObjectlistOk returns a tuple with the IncludeObjectlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeObjectlist

`func (o *CassandraProtectionGroupParams) SetIncludeObjectlist(v []string)`

SetIncludeObjectlist sets IncludeObjectlist field to given value.

### HasIncludeObjectlist

`func (o *CassandraProtectionGroupParams) HasIncludeObjectlist() bool`

HasIncludeObjectlist returns a boolean if a field has been set.

### SetIncludeObjectlistNil

`func (o *CassandraProtectionGroupParams) SetIncludeObjectlistNil(b bool)`

 SetIncludeObjectlistNil sets the value for IncludeObjectlist to be an explicit nil

### UnsetIncludeObjectlist
`func (o *CassandraProtectionGroupParams) UnsetIncludeObjectlist()`

UnsetIncludeObjectlist ensures that no value is present for IncludeObjectlist, not even an explicit nil
### GetObjects

`func (o *CassandraProtectionGroupParams) GetObjects() []NoSqlProtectionGroupObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *CassandraProtectionGroupParams) GetObjectsOk() (*[]NoSqlProtectionGroupObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *CassandraProtectionGroupParams) SetObjects(v []NoSqlProtectionGroupObjectParams)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *CassandraProtectionGroupParams) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetOverwriteExcludeObjectlist

`func (o *CassandraProtectionGroupParams) GetOverwriteExcludeObjectlist() bool`

GetOverwriteExcludeObjectlist returns the OverwriteExcludeObjectlist field if non-nil, zero value otherwise.

### GetOverwriteExcludeObjectlistOk

`func (o *CassandraProtectionGroupParams) GetOverwriteExcludeObjectlistOk() (*bool, bool)`

GetOverwriteExcludeObjectlistOk returns a tuple with the OverwriteExcludeObjectlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteExcludeObjectlist

`func (o *CassandraProtectionGroupParams) SetOverwriteExcludeObjectlist(v bool)`

SetOverwriteExcludeObjectlist sets OverwriteExcludeObjectlist field to given value.

### HasOverwriteExcludeObjectlist

`func (o *CassandraProtectionGroupParams) HasOverwriteExcludeObjectlist() bool`

HasOverwriteExcludeObjectlist returns a boolean if a field has been set.

### SetOverwriteExcludeObjectlistNil

`func (o *CassandraProtectionGroupParams) SetOverwriteExcludeObjectlistNil(b bool)`

 SetOverwriteExcludeObjectlistNil sets the value for OverwriteExcludeObjectlist to be an explicit nil

### UnsetOverwriteExcludeObjectlist
`func (o *CassandraProtectionGroupParams) UnsetOverwriteExcludeObjectlist()`

UnsetOverwriteExcludeObjectlist ensures that no value is present for OverwriteExcludeObjectlist, not even an explicit nil
### GetOverwriteIncludeObjectlist

`func (o *CassandraProtectionGroupParams) GetOverwriteIncludeObjectlist() bool`

GetOverwriteIncludeObjectlist returns the OverwriteIncludeObjectlist field if non-nil, zero value otherwise.

### GetOverwriteIncludeObjectlistOk

`func (o *CassandraProtectionGroupParams) GetOverwriteIncludeObjectlistOk() (*bool, bool)`

GetOverwriteIncludeObjectlistOk returns a tuple with the OverwriteIncludeObjectlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteIncludeObjectlist

`func (o *CassandraProtectionGroupParams) SetOverwriteIncludeObjectlist(v bool)`

SetOverwriteIncludeObjectlist sets OverwriteIncludeObjectlist field to given value.

### HasOverwriteIncludeObjectlist

`func (o *CassandraProtectionGroupParams) HasOverwriteIncludeObjectlist() bool`

HasOverwriteIncludeObjectlist returns a boolean if a field has been set.

### SetOverwriteIncludeObjectlistNil

`func (o *CassandraProtectionGroupParams) SetOverwriteIncludeObjectlistNil(b bool)`

 SetOverwriteIncludeObjectlistNil sets the value for OverwriteIncludeObjectlist to be an explicit nil

### UnsetOverwriteIncludeObjectlist
`func (o *CassandraProtectionGroupParams) UnsetOverwriteIncludeObjectlist()`

UnsetOverwriteIncludeObjectlist ensures that no value is present for OverwriteIncludeObjectlist, not even an explicit nil
### GetSourceId

`func (o *CassandraProtectionGroupParams) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *CassandraProtectionGroupParams) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *CassandraProtectionGroupParams) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *CassandraProtectionGroupParams) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *CassandraProtectionGroupParams) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *CassandraProtectionGroupParams) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *CassandraProtectionGroupParams) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *CassandraProtectionGroupParams) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *CassandraProtectionGroupParams) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *CassandraProtectionGroupParams) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *CassandraProtectionGroupParams) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *CassandraProtectionGroupParams) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetDataCenters

`func (o *CassandraProtectionGroupParams) GetDataCenters() []string`

GetDataCenters returns the DataCenters field if non-nil, zero value otherwise.

### GetDataCentersOk

`func (o *CassandraProtectionGroupParams) GetDataCentersOk() (*[]string, bool)`

GetDataCentersOk returns a tuple with the DataCenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenters

`func (o *CassandraProtectionGroupParams) SetDataCenters(v []string)`

SetDataCenters sets DataCenters field to given value.

### HasDataCenters

`func (o *CassandraProtectionGroupParams) HasDataCenters() bool`

HasDataCenters returns a boolean if a field has been set.

### GetIsLogBackup

`func (o *CassandraProtectionGroupParams) GetIsLogBackup() bool`

GetIsLogBackup returns the IsLogBackup field if non-nil, zero value otherwise.

### GetIsLogBackupOk

`func (o *CassandraProtectionGroupParams) GetIsLogBackupOk() (*bool, bool)`

GetIsLogBackupOk returns a tuple with the IsLogBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLogBackup

`func (o *CassandraProtectionGroupParams) SetIsLogBackup(v bool)`

SetIsLogBackup sets IsLogBackup field to given value.

### HasIsLogBackup

`func (o *CassandraProtectionGroupParams) HasIsLogBackup() bool`

HasIsLogBackup returns a boolean if a field has been set.

### SetIsLogBackupNil

`func (o *CassandraProtectionGroupParams) SetIsLogBackupNil(b bool)`

 SetIsLogBackupNil sets the value for IsLogBackup to be an explicit nil

### UnsetIsLogBackup
`func (o *CassandraProtectionGroupParams) UnsetIsLogBackup()`

UnsetIsLogBackup ensures that no value is present for IsLogBackup, not even an explicit nil
### GetIsSystemKeyspaceBackup

`func (o *CassandraProtectionGroupParams) GetIsSystemKeyspaceBackup() bool`

GetIsSystemKeyspaceBackup returns the IsSystemKeyspaceBackup field if non-nil, zero value otherwise.

### GetIsSystemKeyspaceBackupOk

`func (o *CassandraProtectionGroupParams) GetIsSystemKeyspaceBackupOk() (*bool, bool)`

GetIsSystemKeyspaceBackupOk returns a tuple with the IsSystemKeyspaceBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemKeyspaceBackup

`func (o *CassandraProtectionGroupParams) SetIsSystemKeyspaceBackup(v bool)`

SetIsSystemKeyspaceBackup sets IsSystemKeyspaceBackup field to given value.

### HasIsSystemKeyspaceBackup

`func (o *CassandraProtectionGroupParams) HasIsSystemKeyspaceBackup() bool`

HasIsSystemKeyspaceBackup returns a boolean if a field has been set.

### SetIsSystemKeyspaceBackupNil

`func (o *CassandraProtectionGroupParams) SetIsSystemKeyspaceBackupNil(b bool)`

 SetIsSystemKeyspaceBackupNil sets the value for IsSystemKeyspaceBackup to be an explicit nil

### UnsetIsSystemKeyspaceBackup
`func (o *CassandraProtectionGroupParams) UnsetIsSystemKeyspaceBackup()`

UnsetIsSystemKeyspaceBackup ensures that no value is present for IsSystemKeyspaceBackup, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


