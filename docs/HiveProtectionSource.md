# HiveProtectionSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Specifies the instance name of the Hive entity. | [optional] 
**TableInfo** | Pointer to [**HiveTable**](HiveTable.md) |  | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of the managed Object in Hive Protection Source. Specifies the type of an Hive source entity. &#39;kCluster&#39; indicates a Hive cluster distributed over several physical nodes. &#39;kDatabase&#39; indicates a Database in the Hive environment. &#39;kTable&#39; indicates a Table in the Hive environment. | [optional] 
**Uuid** | Pointer to **NullableString** | Specifies the UUID for the Hive entity. | [optional] 

## Methods

### NewHiveProtectionSource

`func NewHiveProtectionSource() *HiveProtectionSource`

NewHiveProtectionSource instantiates a new HiveProtectionSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHiveProtectionSourceWithDefaults

`func NewHiveProtectionSourceWithDefaults() *HiveProtectionSource`

NewHiveProtectionSourceWithDefaults instantiates a new HiveProtectionSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HiveProtectionSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HiveProtectionSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HiveProtectionSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HiveProtectionSource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HiveProtectionSource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HiveProtectionSource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTableInfo

`func (o *HiveProtectionSource) GetTableInfo() HiveTable`

GetTableInfo returns the TableInfo field if non-nil, zero value otherwise.

### GetTableInfoOk

`func (o *HiveProtectionSource) GetTableInfoOk() (*HiveTable, bool)`

GetTableInfoOk returns a tuple with the TableInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableInfo

`func (o *HiveProtectionSource) SetTableInfo(v HiveTable)`

SetTableInfo sets TableInfo field to given value.

### HasTableInfo

`func (o *HiveProtectionSource) HasTableInfo() bool`

HasTableInfo returns a boolean if a field has been set.

### GetType

`func (o *HiveProtectionSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HiveProtectionSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HiveProtectionSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HiveProtectionSource) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *HiveProtectionSource) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *HiveProtectionSource) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUuid

`func (o *HiveProtectionSource) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *HiveProtectionSource) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *HiveProtectionSource) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *HiveProtectionSource) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *HiveProtectionSource) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *HiveProtectionSource) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


