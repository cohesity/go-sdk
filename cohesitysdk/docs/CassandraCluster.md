# CassandraCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryHost** | Pointer to **NullableString** | Primary host from this Cassandra cluster. | [optional] 
**Seeds** | Pointer to **[]string** | Seeds of this Cassandra Cluster. | [optional] 

## Methods

### NewCassandraCluster

`func NewCassandraCluster() *CassandraCluster`

NewCassandraCluster instantiates a new CassandraCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCassandraClusterWithDefaults

`func NewCassandraClusterWithDefaults() *CassandraCluster`

NewCassandraClusterWithDefaults instantiates a new CassandraCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimaryHost

`func (o *CassandraCluster) GetPrimaryHost() string`

GetPrimaryHost returns the PrimaryHost field if non-nil, zero value otherwise.

### GetPrimaryHostOk

`func (o *CassandraCluster) GetPrimaryHostOk() (*string, bool)`

GetPrimaryHostOk returns a tuple with the PrimaryHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryHost

`func (o *CassandraCluster) SetPrimaryHost(v string)`

SetPrimaryHost sets PrimaryHost field to given value.

### HasPrimaryHost

`func (o *CassandraCluster) HasPrimaryHost() bool`

HasPrimaryHost returns a boolean if a field has been set.

### SetPrimaryHostNil

`func (o *CassandraCluster) SetPrimaryHostNil(b bool)`

 SetPrimaryHostNil sets the value for PrimaryHost to be an explicit nil

### UnsetPrimaryHost
`func (o *CassandraCluster) UnsetPrimaryHost()`

UnsetPrimaryHost ensures that no value is present for PrimaryHost, not even an explicit nil
### GetSeeds

`func (o *CassandraCluster) GetSeeds() []string`

GetSeeds returns the Seeds field if non-nil, zero value otherwise.

### GetSeedsOk

`func (o *CassandraCluster) GetSeedsOk() (*[]string, bool)`

GetSeedsOk returns a tuple with the Seeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeeds

`func (o *CassandraCluster) SetSeeds(v []string)`

SetSeeds sets Seeds field to given value.

### HasSeeds

`func (o *CassandraCluster) HasSeeds() bool`

HasSeeds returns a boolean if a field has been set.

### SetSeedsNil

`func (o *CassandraCluster) SetSeedsNil(b bool)`

 SetSeedsNil sets the value for Seeds to be an explicit nil

### UnsetSeeds
`func (o *CassandraCluster) UnsetSeeds()`

UnsetSeeds ensures that no value is present for Seeds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


