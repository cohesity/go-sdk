# HiveTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApproxSizeBytes** | Pointer to **NullableInt64** | Specifies the approx size of the table in bytes. | [optional] 
**CreatedOn** | Pointer to **NullableInt64** | Specifies the created on, epoch millis. | [optional] 
**IsTransactionalTable** | Pointer to **NullableBool** | Specifies if this is a transactional table. | [optional] 
**Owner** | Pointer to **NullableString** | Specifies the owner of the table. | [optional] 
**TableType** | Pointer to **NullableString** | Specifies the type of table ex. MANAGED,VIRTUAL etc. Specifies the type of an Hive table. &#39;kManaged&#39; indicates a MANAGED Hive table. &#39;kExternal&#39; indicates a EXTERNAL Hive table. &#39;kVirtual&#39; indicates a VIRTUAL Hive tablet. &#39;kIndex&#39; indicates a INDEX Hive table. | [optional] 

## Methods

### NewHiveTable

`func NewHiveTable() *HiveTable`

NewHiveTable instantiates a new HiveTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHiveTableWithDefaults

`func NewHiveTableWithDefaults() *HiveTable`

NewHiveTableWithDefaults instantiates a new HiveTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApproxSizeBytes

`func (o *HiveTable) GetApproxSizeBytes() int64`

GetApproxSizeBytes returns the ApproxSizeBytes field if non-nil, zero value otherwise.

### GetApproxSizeBytesOk

`func (o *HiveTable) GetApproxSizeBytesOk() (*int64, bool)`

GetApproxSizeBytesOk returns a tuple with the ApproxSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproxSizeBytes

`func (o *HiveTable) SetApproxSizeBytes(v int64)`

SetApproxSizeBytes sets ApproxSizeBytes field to given value.

### HasApproxSizeBytes

`func (o *HiveTable) HasApproxSizeBytes() bool`

HasApproxSizeBytes returns a boolean if a field has been set.

### SetApproxSizeBytesNil

`func (o *HiveTable) SetApproxSizeBytesNil(b bool)`

 SetApproxSizeBytesNil sets the value for ApproxSizeBytes to be an explicit nil

### UnsetApproxSizeBytes
`func (o *HiveTable) UnsetApproxSizeBytes()`

UnsetApproxSizeBytes ensures that no value is present for ApproxSizeBytes, not even an explicit nil
### GetCreatedOn

`func (o *HiveTable) GetCreatedOn() int64`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *HiveTable) GetCreatedOnOk() (*int64, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *HiveTable) SetCreatedOn(v int64)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *HiveTable) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### SetCreatedOnNil

`func (o *HiveTable) SetCreatedOnNil(b bool)`

 SetCreatedOnNil sets the value for CreatedOn to be an explicit nil

### UnsetCreatedOn
`func (o *HiveTable) UnsetCreatedOn()`

UnsetCreatedOn ensures that no value is present for CreatedOn, not even an explicit nil
### GetIsTransactionalTable

`func (o *HiveTable) GetIsTransactionalTable() bool`

GetIsTransactionalTable returns the IsTransactionalTable field if non-nil, zero value otherwise.

### GetIsTransactionalTableOk

`func (o *HiveTable) GetIsTransactionalTableOk() (*bool, bool)`

GetIsTransactionalTableOk returns a tuple with the IsTransactionalTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTransactionalTable

`func (o *HiveTable) SetIsTransactionalTable(v bool)`

SetIsTransactionalTable sets IsTransactionalTable field to given value.

### HasIsTransactionalTable

`func (o *HiveTable) HasIsTransactionalTable() bool`

HasIsTransactionalTable returns a boolean if a field has been set.

### SetIsTransactionalTableNil

`func (o *HiveTable) SetIsTransactionalTableNil(b bool)`

 SetIsTransactionalTableNil sets the value for IsTransactionalTable to be an explicit nil

### UnsetIsTransactionalTable
`func (o *HiveTable) UnsetIsTransactionalTable()`

UnsetIsTransactionalTable ensures that no value is present for IsTransactionalTable, not even an explicit nil
### GetOwner

`func (o *HiveTable) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *HiveTable) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *HiveTable) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *HiveTable) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *HiveTable) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *HiveTable) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetTableType

`func (o *HiveTable) GetTableType() string`

GetTableType returns the TableType field if non-nil, zero value otherwise.

### GetTableTypeOk

`func (o *HiveTable) GetTableTypeOk() (*string, bool)`

GetTableTypeOk returns a tuple with the TableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableType

`func (o *HiveTable) SetTableType(v string)`

SetTableType sets TableType field to given value.

### HasTableType

`func (o *HiveTable) HasTableType() bool`

HasTableType returns a boolean if a field has been set.

### SetTableTypeNil

`func (o *HiveTable) SetTableTypeNil(b bool)`

 SetTableTypeNil sets the value for TableType to be an explicit nil

### UnsetTableType
`func (o *HiveTable) UnsetTableType()`

UnsetTableType ensures that no value is present for TableType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


