# StorageDomainStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudSpillVaultId** | Pointer to **NullableInt64** | Specifies the cloud spill vault id of the view box (storage domain). | [optional] 
**GroupList** | Pointer to [**[]StatsGroup**](StatsGroup.md) | Specifies a list of groups associated to this view box (storage domain). | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies the id of the view box (storage domain). | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the view box (storage domain). | [optional] 
**QuotaHardLimitBytes** | Pointer to **NullableInt64** | Specifies the hard limit of physical quota of the view box (storage domain). | [optional] 
**SchemaInfoList** | Pointer to [**[]UsageSchemaInfo**](UsageSchemaInfo.md) | Specifies a list of schemaInfos of the view box (storage domain). | [optional] 
**Stats** | Pointer to [**DataUsageStats**](DataUsageStats.md) |  | [optional] 

## Methods

### NewStorageDomainStats

`func NewStorageDomainStats() *StorageDomainStats`

NewStorageDomainStats instantiates a new StorageDomainStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageDomainStatsWithDefaults

`func NewStorageDomainStatsWithDefaults() *StorageDomainStats`

NewStorageDomainStatsWithDefaults instantiates a new StorageDomainStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudSpillVaultId

`func (o *StorageDomainStats) GetCloudSpillVaultId() int64`

GetCloudSpillVaultId returns the CloudSpillVaultId field if non-nil, zero value otherwise.

### GetCloudSpillVaultIdOk

`func (o *StorageDomainStats) GetCloudSpillVaultIdOk() (*int64, bool)`

GetCloudSpillVaultIdOk returns a tuple with the CloudSpillVaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudSpillVaultId

`func (o *StorageDomainStats) SetCloudSpillVaultId(v int64)`

SetCloudSpillVaultId sets CloudSpillVaultId field to given value.

### HasCloudSpillVaultId

`func (o *StorageDomainStats) HasCloudSpillVaultId() bool`

HasCloudSpillVaultId returns a boolean if a field has been set.

### SetCloudSpillVaultIdNil

`func (o *StorageDomainStats) SetCloudSpillVaultIdNil(b bool)`

 SetCloudSpillVaultIdNil sets the value for CloudSpillVaultId to be an explicit nil

### UnsetCloudSpillVaultId
`func (o *StorageDomainStats) UnsetCloudSpillVaultId()`

UnsetCloudSpillVaultId ensures that no value is present for CloudSpillVaultId, not even an explicit nil
### GetGroupList

`func (o *StorageDomainStats) GetGroupList() []StatsGroup`

GetGroupList returns the GroupList field if non-nil, zero value otherwise.

### GetGroupListOk

`func (o *StorageDomainStats) GetGroupListOk() (*[]StatsGroup, bool)`

GetGroupListOk returns a tuple with the GroupList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupList

`func (o *StorageDomainStats) SetGroupList(v []StatsGroup)`

SetGroupList sets GroupList field to given value.

### HasGroupList

`func (o *StorageDomainStats) HasGroupList() bool`

HasGroupList returns a boolean if a field has been set.

### SetGroupListNil

`func (o *StorageDomainStats) SetGroupListNil(b bool)`

 SetGroupListNil sets the value for GroupList to be an explicit nil

### UnsetGroupList
`func (o *StorageDomainStats) UnsetGroupList()`

UnsetGroupList ensures that no value is present for GroupList, not even an explicit nil
### GetId

`func (o *StorageDomainStats) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorageDomainStats) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorageDomainStats) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *StorageDomainStats) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *StorageDomainStats) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *StorageDomainStats) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *StorageDomainStats) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageDomainStats) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageDomainStats) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageDomainStats) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *StorageDomainStats) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *StorageDomainStats) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetQuotaHardLimitBytes

`func (o *StorageDomainStats) GetQuotaHardLimitBytes() int64`

GetQuotaHardLimitBytes returns the QuotaHardLimitBytes field if non-nil, zero value otherwise.

### GetQuotaHardLimitBytesOk

`func (o *StorageDomainStats) GetQuotaHardLimitBytesOk() (*int64, bool)`

GetQuotaHardLimitBytesOk returns a tuple with the QuotaHardLimitBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaHardLimitBytes

`func (o *StorageDomainStats) SetQuotaHardLimitBytes(v int64)`

SetQuotaHardLimitBytes sets QuotaHardLimitBytes field to given value.

### HasQuotaHardLimitBytes

`func (o *StorageDomainStats) HasQuotaHardLimitBytes() bool`

HasQuotaHardLimitBytes returns a boolean if a field has been set.

### SetQuotaHardLimitBytesNil

`func (o *StorageDomainStats) SetQuotaHardLimitBytesNil(b bool)`

 SetQuotaHardLimitBytesNil sets the value for QuotaHardLimitBytes to be an explicit nil

### UnsetQuotaHardLimitBytes
`func (o *StorageDomainStats) UnsetQuotaHardLimitBytes()`

UnsetQuotaHardLimitBytes ensures that no value is present for QuotaHardLimitBytes, not even an explicit nil
### GetSchemaInfoList

`func (o *StorageDomainStats) GetSchemaInfoList() []UsageSchemaInfo`

GetSchemaInfoList returns the SchemaInfoList field if non-nil, zero value otherwise.

### GetSchemaInfoListOk

`func (o *StorageDomainStats) GetSchemaInfoListOk() (*[]UsageSchemaInfo, bool)`

GetSchemaInfoListOk returns a tuple with the SchemaInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaInfoList

`func (o *StorageDomainStats) SetSchemaInfoList(v []UsageSchemaInfo)`

SetSchemaInfoList sets SchemaInfoList field to given value.

### HasSchemaInfoList

`func (o *StorageDomainStats) HasSchemaInfoList() bool`

HasSchemaInfoList returns a boolean if a field has been set.

### SetSchemaInfoListNil

`func (o *StorageDomainStats) SetSchemaInfoListNil(b bool)`

 SetSchemaInfoListNil sets the value for SchemaInfoList to be an explicit nil

### UnsetSchemaInfoList
`func (o *StorageDomainStats) UnsetSchemaInfoList()`

UnsetSchemaInfoList ensures that no value is present for SchemaInfoList, not even an explicit nil
### GetStats

`func (o *StorageDomainStats) GetStats() DataUsageStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *StorageDomainStats) GetStatsOk() (*DataUsageStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *StorageDomainStats) SetStats(v DataUsageStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *StorageDomainStats) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


