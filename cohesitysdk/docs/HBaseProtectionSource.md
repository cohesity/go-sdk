# HBaseProtectionSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Specifies the instance name of the HBase entity. | [optional] 
**TableInfo** | Pointer to [**HBaseTable**](HBaseTable.md) |  | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of the managed Object in HBase Protection Source. Specifies the type of an HBase source entity. &#39;kCluster&#39; indicates a HBase cluster distributed over several physical nodes. &#39;kNamespace&#39; indicates a Namespace in the HBase environment. &#39;kTable&#39; indicates a Table in the HBase environment. | [optional] 
**Uuid** | Pointer to **NullableString** | Specifies the UUID for the HBase entity. | [optional] 

## Methods

### NewHBaseProtectionSource

`func NewHBaseProtectionSource() *HBaseProtectionSource`

NewHBaseProtectionSource instantiates a new HBaseProtectionSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHBaseProtectionSourceWithDefaults

`func NewHBaseProtectionSourceWithDefaults() *HBaseProtectionSource`

NewHBaseProtectionSourceWithDefaults instantiates a new HBaseProtectionSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HBaseProtectionSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HBaseProtectionSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HBaseProtectionSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HBaseProtectionSource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HBaseProtectionSource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HBaseProtectionSource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTableInfo

`func (o *HBaseProtectionSource) GetTableInfo() HBaseTable`

GetTableInfo returns the TableInfo field if non-nil, zero value otherwise.

### GetTableInfoOk

`func (o *HBaseProtectionSource) GetTableInfoOk() (*HBaseTable, bool)`

GetTableInfoOk returns a tuple with the TableInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableInfo

`func (o *HBaseProtectionSource) SetTableInfo(v HBaseTable)`

SetTableInfo sets TableInfo field to given value.

### HasTableInfo

`func (o *HBaseProtectionSource) HasTableInfo() bool`

HasTableInfo returns a boolean if a field has been set.

### GetType

`func (o *HBaseProtectionSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HBaseProtectionSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HBaseProtectionSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HBaseProtectionSource) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *HBaseProtectionSource) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *HBaseProtectionSource) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUuid

`func (o *HBaseProtectionSource) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *HBaseProtectionSource) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *HBaseProtectionSource) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *HBaseProtectionSource) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *HBaseProtectionSource) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *HBaseProtectionSource) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


