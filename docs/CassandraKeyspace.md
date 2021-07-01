# CassandraKeyspace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChildrenCount** | Pointer to **NullableInt32** | Number of documents in this bucket. | [optional] 
**DcList** | Pointer to **[]string** | If the replication strategy is set as kNetwork, then dc_list will have a list of data centers to which the keyspace is being replicated to. | [optional] 
**ReplicationStrategy** | Pointer to **NullableString** | Replication stragegy for the keyspace. Specifies the type of an Cassandra source entity. | [optional] 

## Methods

### NewCassandraKeyspace

`func NewCassandraKeyspace() *CassandraKeyspace`

NewCassandraKeyspace instantiates a new CassandraKeyspace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCassandraKeyspaceWithDefaults

`func NewCassandraKeyspaceWithDefaults() *CassandraKeyspace`

NewCassandraKeyspaceWithDefaults instantiates a new CassandraKeyspace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildrenCount

`func (o *CassandraKeyspace) GetChildrenCount() int32`

GetChildrenCount returns the ChildrenCount field if non-nil, zero value otherwise.

### GetChildrenCountOk

`func (o *CassandraKeyspace) GetChildrenCountOk() (*int32, bool)`

GetChildrenCountOk returns a tuple with the ChildrenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildrenCount

`func (o *CassandraKeyspace) SetChildrenCount(v int32)`

SetChildrenCount sets ChildrenCount field to given value.

### HasChildrenCount

`func (o *CassandraKeyspace) HasChildrenCount() bool`

HasChildrenCount returns a boolean if a field has been set.

### SetChildrenCountNil

`func (o *CassandraKeyspace) SetChildrenCountNil(b bool)`

 SetChildrenCountNil sets the value for ChildrenCount to be an explicit nil

### UnsetChildrenCount
`func (o *CassandraKeyspace) UnsetChildrenCount()`

UnsetChildrenCount ensures that no value is present for ChildrenCount, not even an explicit nil
### GetDcList

`func (o *CassandraKeyspace) GetDcList() []string`

GetDcList returns the DcList field if non-nil, zero value otherwise.

### GetDcListOk

`func (o *CassandraKeyspace) GetDcListOk() (*[]string, bool)`

GetDcListOk returns a tuple with the DcList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcList

`func (o *CassandraKeyspace) SetDcList(v []string)`

SetDcList sets DcList field to given value.

### HasDcList

`func (o *CassandraKeyspace) HasDcList() bool`

HasDcList returns a boolean if a field has been set.

### SetDcListNil

`func (o *CassandraKeyspace) SetDcListNil(b bool)`

 SetDcListNil sets the value for DcList to be an explicit nil

### UnsetDcList
`func (o *CassandraKeyspace) UnsetDcList()`

UnsetDcList ensures that no value is present for DcList, not even an explicit nil
### GetReplicationStrategy

`func (o *CassandraKeyspace) GetReplicationStrategy() string`

GetReplicationStrategy returns the ReplicationStrategy field if non-nil, zero value otherwise.

### GetReplicationStrategyOk

`func (o *CassandraKeyspace) GetReplicationStrategyOk() (*string, bool)`

GetReplicationStrategyOk returns a tuple with the ReplicationStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationStrategy

`func (o *CassandraKeyspace) SetReplicationStrategy(v string)`

SetReplicationStrategy sets ReplicationStrategy field to given value.

### HasReplicationStrategy

`func (o *CassandraKeyspace) HasReplicationStrategy() bool`

HasReplicationStrategy returns a boolean if a field has been set.

### SetReplicationStrategyNil

`func (o *CassandraKeyspace) SetReplicationStrategyNil(b bool)`

 SetReplicationStrategyNil sets the value for ReplicationStrategy to be an explicit nil

### UnsetReplicationStrategy
`func (o *CassandraKeyspace) UnsetReplicationStrategy()`

UnsetReplicationStrategy ensures that no value is present for ReplicationStrategy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


