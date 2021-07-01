# ExchangeDAGDatabase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseCopyInfoList** | Pointer to [**[]ExchangeDatabaseCopyInfo**](ExchangeDatabaseCopyInfo.md) | Specifies about all the copies of this DAG database. This include active and passive copy of the database. | [optional] 
**Guid** | Pointer to **NullableString** | Specifies the guid of the database. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the database. | [optional] 

## Methods

### NewExchangeDAGDatabase

`func NewExchangeDAGDatabase() *ExchangeDAGDatabase`

NewExchangeDAGDatabase instantiates a new ExchangeDAGDatabase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeDAGDatabaseWithDefaults

`func NewExchangeDAGDatabaseWithDefaults() *ExchangeDAGDatabase`

NewExchangeDAGDatabaseWithDefaults instantiates a new ExchangeDAGDatabase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseCopyInfoList

`func (o *ExchangeDAGDatabase) GetDatabaseCopyInfoList() []ExchangeDatabaseCopyInfo`

GetDatabaseCopyInfoList returns the DatabaseCopyInfoList field if non-nil, zero value otherwise.

### GetDatabaseCopyInfoListOk

`func (o *ExchangeDAGDatabase) GetDatabaseCopyInfoListOk() (*[]ExchangeDatabaseCopyInfo, bool)`

GetDatabaseCopyInfoListOk returns a tuple with the DatabaseCopyInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseCopyInfoList

`func (o *ExchangeDAGDatabase) SetDatabaseCopyInfoList(v []ExchangeDatabaseCopyInfo)`

SetDatabaseCopyInfoList sets DatabaseCopyInfoList field to given value.

### HasDatabaseCopyInfoList

`func (o *ExchangeDAGDatabase) HasDatabaseCopyInfoList() bool`

HasDatabaseCopyInfoList returns a boolean if a field has been set.

### SetDatabaseCopyInfoListNil

`func (o *ExchangeDAGDatabase) SetDatabaseCopyInfoListNil(b bool)`

 SetDatabaseCopyInfoListNil sets the value for DatabaseCopyInfoList to be an explicit nil

### UnsetDatabaseCopyInfoList
`func (o *ExchangeDAGDatabase) UnsetDatabaseCopyInfoList()`

UnsetDatabaseCopyInfoList ensures that no value is present for DatabaseCopyInfoList, not even an explicit nil
### GetGuid

`func (o *ExchangeDAGDatabase) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *ExchangeDAGDatabase) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *ExchangeDAGDatabase) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *ExchangeDAGDatabase) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### SetGuidNil

`func (o *ExchangeDAGDatabase) SetGuidNil(b bool)`

 SetGuidNil sets the value for Guid to be an explicit nil

### UnsetGuid
`func (o *ExchangeDAGDatabase) UnsetGuid()`

UnsetGuid ensures that no value is present for Guid, not even an explicit nil
### GetName

`func (o *ExchangeDAGDatabase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExchangeDAGDatabase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExchangeDAGDatabase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExchangeDAGDatabase) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ExchangeDAGDatabase) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ExchangeDAGDatabase) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


