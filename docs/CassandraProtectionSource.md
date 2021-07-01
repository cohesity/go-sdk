# CassandraProtectionSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterInfo** | Pointer to [**CassandraCluster**](CassandraCluster.md) |  | [optional] 
**KeyspaceInfo** | Pointer to [**CassandraKeyspace**](CassandraKeyspace.md) |  | [optional] 
**Name** | Pointer to **NullableString** | Specifies the instance name of the Cassandra entity. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of the managed Object in Cassandra Protection Source. Replication strategy options for a keyspace. &#39;kCluster&#39; indicates a Cassandra cluster distributed over several physical nodes. &#39;kKeyspace&#39; indicates a Keyspace enclosing one or more tables. &#39;kTable&#39; indicates a Table in the Cassandra environment. | [optional] 
**Uuid** | Pointer to **NullableString** | Specifies the UUID for the Cassandra entity. Note : For each entity an ID unique within top level entity should be assigned by imanis backend. Example, UUID for a table can be the string &lt;keyspace_name&gt;.&lt;table_name&gt; | [optional] 

## Methods

### NewCassandraProtectionSource

`func NewCassandraProtectionSource() *CassandraProtectionSource`

NewCassandraProtectionSource instantiates a new CassandraProtectionSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCassandraProtectionSourceWithDefaults

`func NewCassandraProtectionSourceWithDefaults() *CassandraProtectionSource`

NewCassandraProtectionSourceWithDefaults instantiates a new CassandraProtectionSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterInfo

`func (o *CassandraProtectionSource) GetClusterInfo() CassandraCluster`

GetClusterInfo returns the ClusterInfo field if non-nil, zero value otherwise.

### GetClusterInfoOk

`func (o *CassandraProtectionSource) GetClusterInfoOk() (*CassandraCluster, bool)`

GetClusterInfoOk returns a tuple with the ClusterInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterInfo

`func (o *CassandraProtectionSource) SetClusterInfo(v CassandraCluster)`

SetClusterInfo sets ClusterInfo field to given value.

### HasClusterInfo

`func (o *CassandraProtectionSource) HasClusterInfo() bool`

HasClusterInfo returns a boolean if a field has been set.

### GetKeyspaceInfo

`func (o *CassandraProtectionSource) GetKeyspaceInfo() CassandraKeyspace`

GetKeyspaceInfo returns the KeyspaceInfo field if non-nil, zero value otherwise.

### GetKeyspaceInfoOk

`func (o *CassandraProtectionSource) GetKeyspaceInfoOk() (*CassandraKeyspace, bool)`

GetKeyspaceInfoOk returns a tuple with the KeyspaceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyspaceInfo

`func (o *CassandraProtectionSource) SetKeyspaceInfo(v CassandraKeyspace)`

SetKeyspaceInfo sets KeyspaceInfo field to given value.

### HasKeyspaceInfo

`func (o *CassandraProtectionSource) HasKeyspaceInfo() bool`

HasKeyspaceInfo returns a boolean if a field has been set.

### GetName

`func (o *CassandraProtectionSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CassandraProtectionSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CassandraProtectionSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CassandraProtectionSource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CassandraProtectionSource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CassandraProtectionSource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *CassandraProtectionSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CassandraProtectionSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CassandraProtectionSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CassandraProtectionSource) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *CassandraProtectionSource) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CassandraProtectionSource) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUuid

`func (o *CassandraProtectionSource) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CassandraProtectionSource) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CassandraProtectionSource) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CassandraProtectionSource) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *CassandraProtectionSource) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *CassandraProtectionSource) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


