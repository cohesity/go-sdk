# SqlSourceId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDateMsecs** | Pointer to **NullableInt64** | Specifies a unique identifier generated from the date the database is created or renamed. Cohesity uses this identifier in combination with the databaseId to uniquely identify a database. | [optional] 
**DatabaseId** | Pointer to **NullableInt64** | Specifies a unique id of the database but only for the life of the database. SQL Server may reuse database ids. Cohesity uses the createDateMsecs in combination with this databaseId to uniquely identify a database. | [optional] 
**InstanceId** | Pointer to **[]int32** | Array of bytes that stores the SQL Server Instance id.  Specifies unique id for the SQL Server instance. This id does not change during the life of the instance. | [optional] 

## Methods

### NewSqlSourceId

`func NewSqlSourceId() *SqlSourceId`

NewSqlSourceId instantiates a new SqlSourceId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlSourceIdWithDefaults

`func NewSqlSourceIdWithDefaults() *SqlSourceId`

NewSqlSourceIdWithDefaults instantiates a new SqlSourceId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDateMsecs

`func (o *SqlSourceId) GetCreatedDateMsecs() int64`

GetCreatedDateMsecs returns the CreatedDateMsecs field if non-nil, zero value otherwise.

### GetCreatedDateMsecsOk

`func (o *SqlSourceId) GetCreatedDateMsecsOk() (*int64, bool)`

GetCreatedDateMsecsOk returns a tuple with the CreatedDateMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateMsecs

`func (o *SqlSourceId) SetCreatedDateMsecs(v int64)`

SetCreatedDateMsecs sets CreatedDateMsecs field to given value.

### HasCreatedDateMsecs

`func (o *SqlSourceId) HasCreatedDateMsecs() bool`

HasCreatedDateMsecs returns a boolean if a field has been set.

### SetCreatedDateMsecsNil

`func (o *SqlSourceId) SetCreatedDateMsecsNil(b bool)`

 SetCreatedDateMsecsNil sets the value for CreatedDateMsecs to be an explicit nil

### UnsetCreatedDateMsecs
`func (o *SqlSourceId) UnsetCreatedDateMsecs()`

UnsetCreatedDateMsecs ensures that no value is present for CreatedDateMsecs, not even an explicit nil
### GetDatabaseId

`func (o *SqlSourceId) GetDatabaseId() int64`

GetDatabaseId returns the DatabaseId field if non-nil, zero value otherwise.

### GetDatabaseIdOk

`func (o *SqlSourceId) GetDatabaseIdOk() (*int64, bool)`

GetDatabaseIdOk returns a tuple with the DatabaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseId

`func (o *SqlSourceId) SetDatabaseId(v int64)`

SetDatabaseId sets DatabaseId field to given value.

### HasDatabaseId

`func (o *SqlSourceId) HasDatabaseId() bool`

HasDatabaseId returns a boolean if a field has been set.

### SetDatabaseIdNil

`func (o *SqlSourceId) SetDatabaseIdNil(b bool)`

 SetDatabaseIdNil sets the value for DatabaseId to be an explicit nil

### UnsetDatabaseId
`func (o *SqlSourceId) UnsetDatabaseId()`

UnsetDatabaseId ensures that no value is present for DatabaseId, not even an explicit nil
### GetInstanceId

`func (o *SqlSourceId) GetInstanceId() []int32`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *SqlSourceId) GetInstanceIdOk() (*[]int32, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *SqlSourceId) SetInstanceId(v []int32)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *SqlSourceId) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### SetInstanceIdNil

`func (o *SqlSourceId) SetInstanceIdNil(b bool)`

 SetInstanceIdNil sets the value for InstanceId to be an explicit nil

### UnsetInstanceId
`func (o *SqlSourceId) UnsetInstanceId()`

UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


