# TenantStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupList** | Pointer to [**[]StatsGroup**](StatsGroup.md) | Specifies a list of groups associated to this tenant (organization). | [optional] 
**Id** | Pointer to **NullableString** | Specifies the id of the tenant (organization). | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the tenant (organization). | [optional] 
**SchemaInfoList** | Pointer to [**[]UsageSchemaInfo**](UsageSchemaInfo.md) | Specifies a list of schemaInfos of the tenant (organization). | [optional] 
**Stats** | Pointer to [**DataUsageStats**](DataUsageStats.md) |  | [optional] 

## Methods

### NewTenantStats

`func NewTenantStats() *TenantStats`

NewTenantStats instantiates a new TenantStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantStatsWithDefaults

`func NewTenantStatsWithDefaults() *TenantStats`

NewTenantStatsWithDefaults instantiates a new TenantStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupList

`func (o *TenantStats) GetGroupList() []StatsGroup`

GetGroupList returns the GroupList field if non-nil, zero value otherwise.

### GetGroupListOk

`func (o *TenantStats) GetGroupListOk() (*[]StatsGroup, bool)`

GetGroupListOk returns a tuple with the GroupList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupList

`func (o *TenantStats) SetGroupList(v []StatsGroup)`

SetGroupList sets GroupList field to given value.

### HasGroupList

`func (o *TenantStats) HasGroupList() bool`

HasGroupList returns a boolean if a field has been set.

### SetGroupListNil

`func (o *TenantStats) SetGroupListNil(b bool)`

 SetGroupListNil sets the value for GroupList to be an explicit nil

### UnsetGroupList
`func (o *TenantStats) UnsetGroupList()`

UnsetGroupList ensures that no value is present for GroupList, not even an explicit nil
### GetId

`func (o *TenantStats) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantStats) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantStats) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TenantStats) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TenantStats) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TenantStats) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *TenantStats) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantStats) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantStats) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TenantStats) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TenantStats) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TenantStats) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSchemaInfoList

`func (o *TenantStats) GetSchemaInfoList() []UsageSchemaInfo`

GetSchemaInfoList returns the SchemaInfoList field if non-nil, zero value otherwise.

### GetSchemaInfoListOk

`func (o *TenantStats) GetSchemaInfoListOk() (*[]UsageSchemaInfo, bool)`

GetSchemaInfoListOk returns a tuple with the SchemaInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaInfoList

`func (o *TenantStats) SetSchemaInfoList(v []UsageSchemaInfo)`

SetSchemaInfoList sets SchemaInfoList field to given value.

### HasSchemaInfoList

`func (o *TenantStats) HasSchemaInfoList() bool`

HasSchemaInfoList returns a boolean if a field has been set.

### SetSchemaInfoListNil

`func (o *TenantStats) SetSchemaInfoListNil(b bool)`

 SetSchemaInfoListNil sets the value for SchemaInfoList to be an explicit nil

### UnsetSchemaInfoList
`func (o *TenantStats) UnsetSchemaInfoList()`

UnsetSchemaInfoList ensures that no value is present for SchemaInfoList, not even an explicit nil
### GetStats

`func (o *TenantStats) GetStats() DataUsageStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *TenantStats) GetStatsOk() (*DataUsageStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *TenantStats) SetStats(v DataUsageStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *TenantStats) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


