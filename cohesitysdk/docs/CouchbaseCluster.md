# CouchbaseCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Seeds** | Pointer to **[]string** | Seeds of this Couchbase Cluster. | [optional] 

## Methods

### NewCouchbaseCluster

`func NewCouchbaseCluster() *CouchbaseCluster`

NewCouchbaseCluster instantiates a new CouchbaseCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouchbaseClusterWithDefaults

`func NewCouchbaseClusterWithDefaults() *CouchbaseCluster`

NewCouchbaseClusterWithDefaults instantiates a new CouchbaseCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeeds

`func (o *CouchbaseCluster) GetSeeds() []string`

GetSeeds returns the Seeds field if non-nil, zero value otherwise.

### GetSeedsOk

`func (o *CouchbaseCluster) GetSeedsOk() (*[]string, bool)`

GetSeedsOk returns a tuple with the Seeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeeds

`func (o *CouchbaseCluster) SetSeeds(v []string)`

SetSeeds sets Seeds field to given value.

### HasSeeds

`func (o *CouchbaseCluster) HasSeeds() bool`

HasSeeds returns a boolean if a field has been set.

### SetSeedsNil

`func (o *CouchbaseCluster) SetSeedsNil(b bool)`

 SetSeedsNil sets the value for Seeds to be an explicit nil

### UnsetSeeds
`func (o *CouchbaseCluster) UnsetSeeds()`

UnsetSeeds ensures that no value is present for Seeds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


