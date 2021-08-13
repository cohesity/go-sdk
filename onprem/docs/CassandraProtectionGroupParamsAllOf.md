# CassandraProtectionGroupParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataCenters** | Pointer to **[]string** | Only the specified data centers will be considered while taking backup. The keyspaces having replication strategy &#39;Simple&#39; can be backed up only if all the datacenters for the cassandra cluster are specified. For any keyspace having replication strategy as &#39;Network&#39;, all the associated data centers should be specified. | [optional] 

## Methods

### NewCassandraProtectionGroupParamsAllOf

`func NewCassandraProtectionGroupParamsAllOf() *CassandraProtectionGroupParamsAllOf`

NewCassandraProtectionGroupParamsAllOf instantiates a new CassandraProtectionGroupParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCassandraProtectionGroupParamsAllOfWithDefaults

`func NewCassandraProtectionGroupParamsAllOfWithDefaults() *CassandraProtectionGroupParamsAllOf`

NewCassandraProtectionGroupParamsAllOfWithDefaults instantiates a new CassandraProtectionGroupParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataCenters

`func (o *CassandraProtectionGroupParamsAllOf) GetDataCenters() []string`

GetDataCenters returns the DataCenters field if non-nil, zero value otherwise.

### GetDataCentersOk

`func (o *CassandraProtectionGroupParamsAllOf) GetDataCentersOk() (*[]string, bool)`

GetDataCentersOk returns a tuple with the DataCenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenters

`func (o *CassandraProtectionGroupParamsAllOf) SetDataCenters(v []string)`

SetDataCenters sets DataCenters field to given value.

### HasDataCenters

`func (o *CassandraProtectionGroupParamsAllOf) HasDataCenters() bool`

HasDataCenters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


